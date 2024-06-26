package ch_dto

import (
	"ws/internal/chatroom/ch_domain"
	"ws/internal/common/util"
)

type ChatroomDTO struct {
	ID           int                      `json:"ID"`
	Type         ch_domain.RoomType       `json:"type"`
	Participants []ChatroomParticipantDTO `json:"participants"`
	Messages     []ChatroomMessageDTO     `json:"messages"`
}

type ChatroomParticipantDTO struct {
	ID           int    `json:"ID"`
	Name         string `json:"name"`
	ProfileImage string `json:"profileImage"`
}

type ChatroomMessageDTO struct {
	ID           int                    `json:"ID"`
	RoomID       int                    `json:"roomID"`
	Content      string                 `json:"content"`
	Type         ch_domain.MessageType  `json:"type"`
	FileLocation string                 `json:"fileLocation"`
	Participant  ChatroomParticipantDTO `json:"participant"`
	Time         string                 `json:"time"`
}

func NewChatroomMessageDto(message *ch_domain.ChatroomMessage) *ChatroomMessageDTO {
	return &ChatroomMessageDTO{
		ID:           message.ID,
		RoomID:       message.RoomID,
		Content:      message.Content,
		Type:         message.Type,
		FileLocation: message.FileLocation,
		Participant: ChatroomParticipantDTO{
			ID:           message.Participant.ID,
			Name:         message.Participant.Name,
			ProfileImage: message.Participant.ProfileImage,
		},
		Time: util.ConvertDateToString(message.Time),
	}
}

func NewChatroomDTO(chatroom *ch_domain.Chatroom, chatroomMessages []ch_domain.ChatroomMessage) *ChatroomDTO {
	dto := &ChatroomDTO{
		ID:   chatroom.ID,
		Type: chatroom.Type,
	}

	for _, participant := range chatroom.Participants {
		dto.Participants = append(dto.Participants, ChatroomParticipantDTO{
			ID:           participant.ID,
			Name:         participant.Name,
			ProfileImage: participant.ProfileImage,
		})
	}

	for _, message := range chatroomMessages {
		dto.Messages = append(dto.Messages, *NewChatroomMessageDto(&message))
	}

	return dto
}
