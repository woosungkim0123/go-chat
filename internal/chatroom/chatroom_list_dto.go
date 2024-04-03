package chatroom

import (
	"time"
)

type ChatroomWithLastMessageDTO struct {
	RoomID      int             `json:"roomId"`
	RoomType    RoomType        `json:"roomType"`
	LastMessage *LastMessageDTO `json:"lastMessage,omitempty"`
}

type LastMessageDTO struct {
	ID          int                       `json:"Id"`
	Content     string                    `json:"content"`
	Type        MessageType               `json:"type"`
	Participant LastMessageParticipantDTO `json:"participant"`
	Time        time.Time                 `json:"time"`
}

type LastMessageParticipantDTO struct {
	ID int `json:"id"`
}

func NewChatroomWithLastMessageDTO(chatroom *Chatroom, message *ChatroomMessage) *ChatroomWithLastMessageDTO {
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
				ID: message.Participant.Id,
			},
			Time: message.Time,
		}
	}

	return dto
}
