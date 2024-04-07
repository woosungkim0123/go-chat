package ws_service

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"ws/internal/auth/service"
	"ws/internal/chatroom/ch_domain"
	"ws/internal/chatroom/ch_service"
	"ws/internal/common/util"
	"ws/internal/web_socket/ws_domain"
	"ws/internal/web_socket/ws_dto"
)

type WebSocketService struct {
	wsChan             chan ws_dto.WsJsonRequest
	ActiveChatroomList map[int]ws_domain.ChatroomSession
	authService        *service.AuthService
	chatService        *ch_service.ChatroomService
}

func NewWebSocketService(authService *service.AuthService, chatService *ch_service.ChatroomService) *WebSocketService {
	return &WebSocketService{
		wsChan:             make(chan ws_dto.WsJsonRequest),
		ActiveChatroomList: make(map[int]ws_domain.ChatroomSession),
		authService:        authService,
		chatService:        chatService,
	}
}

func (s *WebSocketService) ListenForWebSocket(conn *ws_domain.WebSocketConnection) {
	log.Println("Listening to websocket~~")

	var request ws_dto.WsJsonRequest
	for {
		err := conn.ReadJSON(&request)
		if err != nil {
			log.Println("Error reading JSON:", err)
			if err := conn.Close(); err != nil {
				log.Println("Error closing connection:", err)
			}
			return
		}
		request.Conn = conn
		s.wsChan <- request
	}
}

// ListenToWsChannel wsChan 채널로부터 메세지를 지속적으로 수신하고 받은 메세지를 처리하여 모든 클라이언트에게 브로드 캐스트하는 역할
func (s *WebSocketService) ListenToWsChannel() {
	var response ws_dto.WsJsonResponse
	for {
		request := <-s.wsChan
		log.Printf("Listening to ws channel\n")
		switch request.Action {
		case "join":
			roomID, userID, err := s.convertRoomIDAndUserID(request.RoomID, request.UserID)
			if err != nil {
				continue
			}

			accessUser, appError := s.authService.FindUserByID(userID)
			if appError != nil {
				log.Printf("User not found: %d\n", userID)
				continue
			}

			// 채팅방 세션이 있으면 해당 세션을 반환하고, 없으면 새로운 채팅방 세션을 생성
			chatroomSession, exists := s.ActiveChatroomList[roomID]
			if !exists {
				chatroomSession = ws_domain.ChatroomSession{
					Participants: make(map[int]ws_domain.UserSession),
				}
				s.ActiveChatroomList[roomID] = chatroomSession
			}

			// 사용자 세션 추가
			chatroomSession.Participants[userID] = ws_domain.UserSession{Conn: request.Conn, UserName: accessUser.Name}
			log.Printf("User %d joined room %d", userID, roomID)

			response.Action = "join"
			response.Message = fmt.Sprintf("User %d joined room %d", userID, roomID)
			s.broadcastToRoom(&request, &response)

		case "left":
			roomID, userID, err := s.convertRoomIDAndUserID(request.RoomID, request.UserID)
			if err != nil {
				continue
			}

			_, appError := s.authService.FindUserByID(userID)
			if appError != nil {
				log.Printf("User not found: %d\n", userID)
				continue
			}

			chatroomSession, exists := s.ActiveChatroomList[roomID]
			if !exists {
				log.Printf("room ID %d not found\n", roomID)
				return
			}

			delete(chatroomSession.Participants, userID)

			response.Action = "left"
			response.Message = fmt.Sprintf("User %d left room %d", userID, roomID)

			s.broadcastToRoom(&request, &response)

		case "broadcast":
			roomID, userID, err := s.convertRoomIDAndUserID(request.RoomID, request.UserID)
			if err != nil {
				continue
			}

			log.Println("Broadcasting message to room", roomID)
			accessUser, appError := s.authService.FindUserByID(userID)
			if appError != nil {
				log.Printf("User not found: %d\n", userID)
				continue
			}

			chatroomMessage := ch_domain.ChatroomMessage{
				RoomID:       roomID,
				Content:      request.Content,
				Type:         request.Type,
				FileLocation: request.FileLocation,
				Participant:  ch_domain.ChatroomParticipant{ID: accessUser.ID, Name: accessUser.Name, ProfileImage: accessUser.ProfileImage},
				Time:         util.GetCurrentDate(),
			}

			messageDto, savedError := s.chatService.SaveMessage(&chatroomMessage)
			if savedError != nil {
				log.Println("Error saving message to database:", savedError)
				continue
			}

			response.Action = "broadcast"
			response.Message = fmt.Sprintf("User %d broadcasted message to room %d", userID, roomID)
			response.Data = messageDto

			s.broadcastToRoom(&request, &response)
		}
	}
}

func (s *WebSocketService) broadcastToRoom(request *ws_dto.WsJsonRequest, response *ws_dto.WsJsonResponse) {
	roomID, err := s.convertStringToInt(request.RoomID)
	if err != nil {
		return
	}

	chatroomSession, exists := s.ActiveChatroomList[roomID]
	if !exists {
		log.Printf("room ID %d not found\n", roomID)
		return
	}

	var willDeleteSession []int
	for uID, userSession := range chatroomSession.Participants {
		writeJsonError := userSession.Conn.WriteJSON(response)
		if writeJsonError != nil {
			log.Println("Error sending message to websocket:", err)
			willDeleteSession = append(willDeleteSession, uID)
			continue
		}
	}

	// 클라이언트와 연결이 유효하지 않다고 판단되면 해당 클라이언트를 삭제합니다.
	for _, uID := range willDeleteSession {
		delete(chatroomSession.Participants, uID)
		log.Printf("removed user %d due to error\n", uID)
	}
}

func (s *WebSocketService) convertRoomIDAndUserID(roomIdStr, userIdStr string) (int, int, error) {
	roomID, roomIDError := s.convertStringToInt(roomIdStr)
	if roomIDError != nil {
		return 0, 0, errors.New("Invalid room ID: " + roomIdStr)
	}
	userID, userIDError := s.convertStringToInt(userIdStr)
	if userIDError != nil {
		return 0, 0, errors.New("Invalid user ID: " + userIdStr)
	}
	return roomID, userID, nil
}

func (s *WebSocketService) convertStringToInt(IDStr string) (int, error) {
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		log.Println("Error converting string to int: ", err)
		return 0, err
	}
	return ID, nil
}
