package ch_dto

import (
	"time"
	"ws/internal/chatroom/ch_domain"
)

type ChatroomWithLastMessageDTO struct {
	RoomID      int                `json:"roomID"`
	RoomType    ch_domain.RoomType `json:"roomType"`
	LastMessage *LastMessageDTO    `json:"lastMessage,omitempty"`
}

type LastMessageDTO struct {
	ID          int                       `json:"ID"`
	Content     string                    `json:"content"`
	Type        ch_domain.MessageType     `json:"type"`
	Participant LastMessageParticipantDTO `json:"participant"`
	Time        time.Time                 `json:"time"`
}

type LastMessageParticipantDTO struct {
	ID int `json:"ID"`
}

func NewChatroomWithLastMessageDTO(chatroom *ch_domain.Chatroom, message *ch_domain.ChatroomMessage) *ChatroomWithLastMessageDTO {
	dto := &ChatroomWithLastMessageDTO{
		RoomID:   chatroom.ID,
		RoomType: chatroom.Type,
	}

	if message != nil {
		dto.LastMessage = &LastMessageDTO{
			ID:      message.ID,
			Content: message.Content,
			Type:    message.Type,
			Participant: LastMessageParticipantDTO{
				ID: message.Participant.ID,
			},
			Time: message.Time,
		}
	}

	return dto
}
