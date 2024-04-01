package web_socket

import (
	"fmt"
	"log"
	"strconv"
	"ws/internal/service/userService"
)

var wsChan = make(chan WsJsonRequest)

var ActiveChatrooms = make(map[int]ChatroomSession)

// 요청을 보낼때 웹소켓을 통해 메시지를 보내게됨
func ListenForWs(conn *WebSocketConnection) {
	log.Println("Listening to websocket~~")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error = ", fmt.Sprintf("%v", r))
		}
	}()

	var request WsJsonRequest
	for {
		err := conn.ReadJSON(&request)
		if err != nil {
			continue
		}
		request.Conn = conn
		wsChan <- request
	}
}

// wsChan 채널로부터 메세지를 지속적으로 수신하고 받은 메세지를 처리하여 모든 클라이언트에게 브로드 캐스트하는 역할
func ListenToWsChannel() {
	log.Println("Listening to ws channel")
	var response WsJsonResponse

	for {
		request := <-wsChan // --blocking

		switch request.Action {
		case "join":
			roomId, userId, err := parseIds(request.RoomId, request.UserId)
			if err != nil {
				log.Println(err)
				continue
			}
			accessUser := userService.FindUser(userId)
			if accessUser == nil {
				log.Printf("User not found: %d\n", userId)
				continue
			}

			// 채팅방 세션이 있으면 해당 세션을 반환하고, 없으면 새로운 채팅방 세션을 생성
			chatroomSession, exists := ActiveChatrooms[roomId]
			if !exists {
				chatroomSession = ChatroomSession{
					Participants: make(map[int]UserSession),
				}
				ActiveChatrooms[roomId] = chatroomSession
			}

			// 사용자 세션 추가
			chatroomSession.Participants[userId] = UserSession{Conn: request.Conn, Username: accessUser.Name}
			log.Printf("User %d joined room %d", userId, roomId)

			response.Action = "join"
			response.ConnectedUserId = request.UserId
			response.Message = fmt.Sprintf("User %d joined room %d", userId, roomId)
			broadcastToRoom(request, response)

		case "username":
			// get a list of all clients and send it back
			//clients[e.Conn] = e.Username
			//users := getUserList()
			response.Action = "list_users"
			response.ConnectedUserId = request.UserId
			broadcastToRoom(request, response)

		case "left":
			response.Action = "list_users"
			//delete(clients, request.Conn)
			//log.Println("Client left" + e.Username)

			//users := getUserList()
			//response.ConnectedUsers = users
			broadcastToRoom(request, response)

		case "broadcast":
			response.Action = "broadcast"
			response.ConnectedUserId = request.UserId
			response.Message = request.Message
			broadcastToRoom(request, response)
		}
	}
}

func broadcastToRoom(request WsJsonRequest, response WsJsonResponse) {
	roomId, err := strconv.Atoi(request.RoomId) // RoomId를 int로 변환
	if err != nil {
		log.Printf("Invalid room ID: %s\n", request.RoomId)
		return
	}

	chatroomSession, exists := ActiveChatrooms[roomId]
	if !exists {
		log.Printf("Room ID %d not found\n", roomId)
		return
	}

	var toDelete []int
	for userID, userSession := range chatroomSession.Participants {
		err := userSession.Conn.WriteJSON(response)
		if err != nil {
			log.Println("Error sending message to websocket:", err)
			toDelete = append(toDelete, userID)
			continue
		}
	}

	// 클라이언트와 연결이 유효하지 않다고 판단되면 해당 클라이언트를 삭제합니다.
	for _, userID := range toDelete {
		delete(chatroomSession.Participants, userID) // 안전하게 삭제
		log.Printf("Removed user %d due to error\n", userID)
	}
}

func parseIds(roomIdStr, userIdStr string) (int, int, error) {
	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid room ID: %s", roomIdStr)
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return roomId, 0, fmt.Errorf("Invalid user ID: %s", userIdStr)
	}

	return roomId, userId, nil
}

//func getUserList() []string {
//	var userList []string
//	for _, x := range clients {
//		if x != "" {
//			userList = append(userList, x)
//		}
//	}
//	sort.Strings(userList)
//	return userList
//}
