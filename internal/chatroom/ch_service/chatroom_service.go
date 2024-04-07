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

func (s *ChatroomService) GetChatroomListByUserID(userID int) ([]ch_dto.ChatroomWithLastMessageDTO, *apperror.CustomError) {
	chatroomWithLastMessageDTO, err := s.chatroomRepository.GetChatroomListByUserID(userID)
	if err != nil {
		log.Printf("채팅방 리스트를 가져오는데 실패했습니다: %v", err)
		return nil, err
	}
	return chatroomWithLastMessageDTO, nil
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

func (s *ChatroomService) GetSingleChatroom(accessUser *udomain.User, opponentUserID int) (*ch_dto.ChatroomDTO, *apperror.CustomError) {
	if s.isAccessMineChatroom(accessUser, opponentUserID) {
		return nil, &apperror.CustomError{Code: apperror.WrongAccessMineChatroom, Message: "잘못된 접근입니다."}
	}
	opponentUser, userError := s.findUserByUserID(opponentUserID)
	if userError != nil {
		return nil, userError
	}
	chatroom, err := s.getSingleChatroom(accessUser, opponentUser)
	if err != nil {
		return nil, err
	}

	var chatroomMessages []ch_domain.ChatroomMessage
	chatroomMessages, err = s.getChatroomMessages(chatroom.ID)

	return ch_dto.NewChatroomDTO(chatroom, chatroomMessages), nil
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
	users := make([]udomain.User, 0)
	users = append(users, *user)

	if err != nil {
		if err.Code == apperror.NotFoundChatroom {
			chatroom, err = s.createChatroom(ch_domain.Mine, users)
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

func (s *ChatroomService) getSingleChatroom(accessUser *udomain.User, opponentUser *udomain.User) (*ch_domain.Chatroom, *apperror.CustomError) {
	chatroom, err := s.chatroomRepository.GetSingleChatroom(accessUser.ID, opponentUser.ID)
	users := make([]udomain.User, 0)
	users = append(users, *accessUser, *opponentUser)

	if err != nil {
		if err.Code == apperror.NotFoundChatroom {
			chatroom, err = s.createChatroom(ch_domain.Single, users)
			if err != nil {
				log.Printf("상대방 채팅방 생성에 실패했습니다: %v", err)
				return nil, err
			}
		} else {
			log.Printf("상대방 채팅방을 가져오는데 실패했습니다: %v", err)
			return nil, err
		}
	}
	return chatroom, nil
}

func (s *ChatroomService) createChatroom(roomType ch_domain.RoomType, users []udomain.User) (*ch_domain.Chatroom, *apperror.CustomError) {
	chatroom := ch_domain.NewChatroom(roomType, users)
	err := s.chatroomRepository.AddChatroom(chatroom)
	if err != nil {
		log.Printf("채팅방을 생성하는데 실패했습니다: %v", err)
		return nil, err
	}

	return chatroom, nil
}

func (s *ChatroomService) getChatroomMessages(chatroomID int) ([]ch_domain.ChatroomMessage, *apperror.CustomError) {
	chatroomMessages, err := s.chatroomRepository.GetChatroomMessages(chatroomID)
	if err != nil {
		log.Printf("채팅방 메시지를 가져오는데 실패했습니다: %v", err)
		return nil, err
	}
	return chatroomMessages, nil
}

func (s *ChatroomService) findUserByUserID(userID int) (*udomain.User, *apperror.CustomError) {
	user, err := s.authService.FindUserByID(userID)
	if err != nil {
		log.Printf("유저를 찾는데 실패했습니다. userID: %d, %v", userID, err)
		return nil, err
	}
	return user, nil
}
