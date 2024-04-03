package chatroom

import (
	"ws/internal/domain"
	"ws/internal/dto"
	"ws/internal/service/userService"
	"ws/internal/store/chatRepository"
	"ws/internal/util/jsonReader"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetChatroomByUserId(accessUserId, otherUserId int) *dto.ChatroomDto {
	room := findChatRoom(accessUserId, otherUserId)
	if room != nil {
		return convertChatroomDto(room, accessUserId)
	}

	room = createChatRoom(accessUserId, otherUserId)

	chatroomDto := convertChatroomDto(room, accessUserId)

	return chatroomDto
}

func (s *Service) SaveMessage(roomId int, chatroomMessage domain.ChatroomMessage) {
	saveChatroom(roomId, chatroomMessage)
}

func (s *Service) GetChatListByUserId(userID int) []ChatroomWithLastMessageDTO {
	chatroomListDto := s.repository.GetChatroomListByUserId(userID)

	return chatroomListDto
}

func getUserChatList(userId int) []domain.Chatroom {
	allChatroom := chatRepository.GetAllChatroom()
	var userChatList []domain.Chatroom
	for _, room := range allChatroom {
		if contains(room.Participants, userId) {
			userChatList = append(userChatList, room)
		}
	}
	return userChatList
}

func saveChatroom(roomId int, chatroomMessage domain.ChatroomMessage) {
	allChatroom := chatRepository.GetAllChatroom()
	for i, room := range allChatroom {
		if room.RoomID == roomId {
			allChatroom[i].Messages = append(room.Messages, chatroomMessage)
		}
	}
	jsonReader.Write("internal/store/json/chatroom.json", allChatroom)
}

func findChatRoom(accessUserId, otherUserId int) *domain.Chatroom {
	allChatroom := chatRepository.GetAllChatroom()
	for _, room := range allChatroom {
		if contains(room.Participants, accessUserId) && contains(room.Participants, otherUserId) {
			return &room
		}
	}
	return nil
}

func contains(users []domain.User, userID int) bool {
	for _, user := range users {
		if user.Id == userID {
			return true
		}
	}
	return false
}

func createChatRoom(accessUserId, otherUserId int) *domain.Chatroom {
	accessUser := userService.FindUser(accessUserId)
	otherUser := userService.FindUser(otherUserId)

	room := domain.Chatroom{
		RoomID:   len(chatRepository.GetAllChatroom()) + 1,
		RoomType: "Single",
		Participants: []domain.User{
			*accessUser,
			*otherUser,
		},
		Messages: []domain.ChatroomMessage{},
	}
	chatRepository.AddChatroom(room)

	return &room
}

func convertChatroomDto(room *domain.Chatroom, accessUserId int) *dto.ChatroomDto {
	var userDtos []dto.UserDto
	var accessUserDto dto.UserDto

	for _, user := range room.Participants {
		if user.Id == accessUserId {
			accessUserDto = dto.UserDto{Id: user.Id, Name: user.Name}
		}
		userDtos = append(userDtos, dto.UserDto{Id: user.Id, Name: user.Name})
	}

	var chatMessageListDto []dto.ChatMessageDto

	for _, m := range room.Messages {
		chatMessageListDto = append(chatMessageListDto, dto.ChatMessageDto{
			Message: m.Message,
			User:    dto.UserDto{Id: m.User.Id, Name: m.User.Name},
			Time:    m.Time.Format("1/02 15:04:05"),
		})
	}

	return &dto.ChatroomDto{
		RoomId:       room.RoomID,
		RoomType:     room.RoomType,
		Participants: userDtos,
		AccessUser:   accessUserDto,
		Messages:     chatMessageListDto,
	}
}
