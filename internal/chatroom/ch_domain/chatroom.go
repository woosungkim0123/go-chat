package ch_domain

import (
	"time"
	"ws/internal/auth/domain"
)

type Chatroom struct {
	ID           int                   `json:"ID"`
	Type         RoomType              `json:"type"`
	Participants []ChatroomParticipant `json:"participants"`
}

type ChatroomMessage struct {
	ID           int                 `json:"ID"`
	RoomID       int                 `json:"roomID"`
	Content      string              `json:"content"`
	Type         MessageType         `json:"type"`
	FileLocation string              `json:"fileLocation"`
	Participant  ChatroomParticipant `json:"participant"`
	Time         time.Time           `json:"time"`
}

type ChatroomParticipant struct {
	ID           int    `json:"ID"`
	Name         string `json:"name"`
	ProfileImage string `json:"profileImage"`
}

func NewChatroom(roomType RoomType, users []domain.User) *Chatroom {
	participants := make([]ChatroomParticipant, 0)
	for _, user := range users {
		participants = append(participants, ChatroomParticipant{ID: user.ID, Name: user.Name, ProfileImage: user.ProfileImage})
	}

	return &Chatroom{Type: roomType, Participants: participants}
}
