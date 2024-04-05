package ch_domain

import (
	"time"
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
