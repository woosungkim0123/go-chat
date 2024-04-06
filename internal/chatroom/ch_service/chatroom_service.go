package ch_service

import (
	"log"
	udomain "ws/internal/auth/domain"
	"ws/internal/auth/service"
	"ws/internal/chatroom/ch_domain"
	"ws/internal/chatroom/ch_dto"
	"ws/internal/chatroom/ch_repository"
	"ws/internal/common/apperror"
)

type ChatroomService struct {
	chatroomRepository *ch_repository.ChatroomRepository
	authService        *service.AuthService
}

func NewChatroomService(chatroomRepository *ch_repository.ChatroomRepository, authService *service.AuthService) *ChatroomService {
	return &ChatroomService{chatroomRepository: chatroomRepository, authService: authService}
}

func (s *ChatroomService) GetMineChatroom(accessUser *udomain.User) (*ch_dto.ChatroomDTO, *apperror.CustomError) {
	chatroom, err := s.getMyChatroom(accessUser)
	if err != nil {
		return nil, err
	}
	var chatroomMessages []ch_domain.ChatroomMessage
	chatroomMessages, err = s.getChatroomMessages(chatroom.ID)
	if err != nil {
		return nil, err
	}

	return ch_dto.NewChatroomDTO(chatroom, chatroomMessages), nil
}

func (s *ChatroomService) GetChatroomByUserID(accessUser *udomain.User, opponentUserID int) (*ch_dto.ChatroomDTO, *apperror.CustomError) {
	if s.isAccessMineChatroom(accessUser, opponentUserID) {
		return nil, &apperror.CustomError{Code: apperror.WrongAccessMineChatroom, Message: "잘못된 접근입니다."}
	}

	// 값이 있으면 string을 int로 변환하는걸로, 그리고 값이 없으면 그냥 나둠
	/*
		if _, err := s.authService.GetUserProfile(otherUserID); err != nil {
			fmt.Errorf("상대방 유저 정보를 가져오는데 실패했습니다: %v", err)
		}
		fmt.Print(otherUserID)
		otherUser, err2 := s.authService.GetUserProfile(otherUserID)
		if err2 != nil {
			fmt.Errorf("상대방 유저 정보를 가져오는데 실패했습니다.")
		}
		fmt.Print(otherUser)
	*/
	//
	//// 값이 없거나 유저가 없거나 잘못되면 그냥 없는걸로 간주하고 자신만있는 채팅방
	//s.repository.findChatroomByUserID(accessUserID, opponentUserID)
	//
	//room := findChatRoom(accessUserId, otherUserId)
	//if room != nil {
	//	return convertChatroomDto(room, accessUserId)
	//}
	//
	//room = createChatRoom(accessUserId, otherUserId)
	//
	//chatroomDto := convertChatroomDto(room, accessUserId)
	//
	//return chatroomDto
	return nil, nil
}

func (s *ChatroomService) SaveMessage(chatroomMessage *ch_domain.ChatroomMessage) (*ch_dto.ChatroomMessageDTO, *apperror.CustomError) {
	message, err := s.chatroomRepository.SaveMessage(chatroomMessage)
	if err != nil {
		return nil, err
	}

	return ch_dto.NewChatroomMessageDto(message), nil
}

func (s *ChatroomService) isAccessMineChatroom(accessUser *udomain.User, opponentUserID int) bool {
	return opponentUserID == 0 || accessUser.ID == opponentUserID
}

func (s *ChatroomService) getMyChatroom(user *udomain.User) (*ch_domain.Chatroom, *apperror.CustomError) {
	chatroom, err := s.chatroomRepository.GetMineChatroom(user.ID)
	if err != nil {
		if err.Code == apperror.NotFoundMineChatroom {
			chatroom, err = s.createMineChatroom(user)
			if err != nil {
				log.Printf("내 채팅방 생성에 실패했습니다: %v", err)
				return nil, err
			}
		} else {
			log.Printf("내 채팅방을 가져오는데 실패했습니다: %v", err)
			return nil, err
		}
	}
	return chatroom, nil
}

func (s *ChatroomService) createMineChatroom(user *udomain.User) (*ch_domain.Chatroom, *apperror.CustomError) {
	chatroom := ch_domain.Chatroom{
		Type: ch_domain.Mine,
		Participants: []ch_domain.ChatroomParticipant{
			{ID: user.ID, Name: user.Name, ProfileImage: user.ProfileImage},
		},
	}

	err := s.chatroomRepository.AddChatroom(&chatroom)
	if err != nil {
		log.Printf("채팅방을 생성하는데 실패했습니다: %v", err)
		return nil, err
	}

	return &chatroom, nil
}

func (s *ChatroomService) getChatroomMessages(chatroomID int) ([]ch_domain.ChatroomMessage, *apperror.CustomError) {
	chatroomMessages, err := s.chatroomRepository.GetChatroomMessages(chatroomID)
	if err != nil {
		log.Printf("채팅방 메시지를 가져오는데 실패했습니다: %v", err)
		return nil, err
	}
	return chatroomMessages, nil
}

//func (s *Service) SaveMessage(roomId int, chatroomMessage domain.ChatroomMessage) {
//	saveChatroom(roomId, chatroomMessage)
//}
//
//func (s *Service) GetChatListByUserId(userID int) []dto2.ChatroomWithLastMessageDTO {
//	chatroomListDto := s.repository.GetChatroomListByUserId(userID)
//
//	return chatroomListDto
//}

//func getUserChatList(userId int) []domain.Chatroom {
//	allChatroom := chatRepository.GetAllChatroom()
//	var userChatList []domain.Chatroom
//	for _, room := range allChatroom {
//		if contains(room.Participants, userId) {
//			userChatList = append(userChatList, room)
//		}
//	}
//	return userChatList
//}
//
//func saveChatroom(roomId int, chatroomMessage domain.ChatroomMessage) {
//	allChatroom := chatRepository.GetAllChatroom()
//	for i, room := range allChatroom {
//		if room.RoomID == roomId {
//			allChatroom[i].Messages = append(room.Messages, chatroomMessage)
//		}
//	}
//	jsonReader.Write("internal/store/json/chatroom.json", allChatroom)
//}
//
//func findChatRoom(accessUserId, otherUserId int) *domain.Chatroom {
//	allChatroom := chatRepository.GetAllChatroom()
//	for _, room := range allChatroom {
//		if contains(room.Participants, accessUserId) && contains(room.Participants, otherUserId) {
//			return &room
//		}
//	}
//	return nil
//}
//
//func contains(users []auth.User, userID int) bool {
//	for _, user := range users {
//		if user.ID == userID {
//			return true
//		}
//	}
//	return false
//}
//
//func createChatRoom(accessUserId, otherUserId int) *domain.Chatroom {
//	accessUser := auth.FindUser(accessUserId)
//	otherUser := auth.FindUser(otherUserId)
//
//	room := domain.Chatroom{
//		RoomID:           len(chatRepository.GetAllChatroom()) + 1,
//		domain2.RoomType: "Single",
//		Participants: []auth.User{
//			*accessUser,
//			*otherUser,
//		},
//		Messages: []domain.ChatroomMessage{},
//	}
//	chatRepository.AddChatroom(room)
//
//	return &room
//}
//
//func convertChatroomDto(room *domain.Chatroom, accessUserId int) *dto.ChatroomDto {
//	var userDtos []dto.UserDto
//	var accessUserDto dto.UserDto
//
//	for _, user := range room.Participants {
//		if user.ID == accessUserId {
//			accessUserDto = dto.UserDto{ID: user.ID, Name: user.Name}
//		}
//		userDtos = append(userDtos, dto.UserDto{ID: user.ID, Name: user.Name})
//	}
//
//	var chatMessageListDto []dto.ChatMessageDto
//
//	for _, m := range room.Messages {
//		chatMessageListDto = append(chatMessageListDto, dto.ChatMessageDto{
//			Message: m.Message,
//			User:    dto.UserDto{ID: m.User.ID, Name: m.User.Name},
//			Time:    m.Time.Format("1/02 15:04:05"),
//		})
//	}
//
//	return &dto.ChatroomDto{
//		RoomId:           room.RoomID,
//		domain2.RoomType: room.RoomType,
//		Participants:     userDtos,
//		AccessUser:       accessUserDto,
//		Messages:         chatMessageListDto,
//	}
//}
