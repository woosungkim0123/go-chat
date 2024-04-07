package ch_dto

import (
	"time"
	"ws/internal/chatroom/ch_domain"
)

type ChatroomWithLastMessageDTO struct {
	ID           int                      `json:"ID"`
	Type         ch_domain.RoomType       `json:"type"`
	Participants []ChatroomParticipantDTO `json:"participants"`
	LastMessage  *LastMessageDTO          `json:"lastMessage,omitempty"`
}

type LastMessageDTO struct {
	ID          int                    `json:"ID"`
	Content     string                 `json:"content"`
	Type        ch_domain.MessageType  `json:"type"`
	Participant ChatroomParticipantDTO `json:"participant"`
	Time        time.Time              `json:"time"`
}

func NewChatroomWithLastMessageDTO(chatroom *ch_domain.Chatroom, message *ch_domain.ChatroomMessage) *ChatroomWithLastMessageDTO {
	dto := &ChatroomWithLastMessageDTO{
		ID:           chatroom.ID,
		Type:         chatroom.Type,
		Participants: make([]ChatroomParticipantDTO, 0),
	}

	for _, participant := range chatroom.Participants {
		dto.Participants = append(dto.Participants, ChatroomParticipantDTO{
			ID:           participant.ID,
			Name:         participant.Name,
			ProfileImage: participant.ProfileImage,
		})
	}

	if message != nil {
		dto.LastMessage = &LastMessageDTO{
			ID:      message.ID,
			Content: message.Content,
			Type:    message.Type,
			Participant: ChatroomParticipantDTO{
				ID:           message.Participant.ID,
				Name:         message.Participant.Name,
				ProfileImage: message.Participant.ProfileImage,
			},
			Time: message.Time,
		}
	}

	return dto
}
