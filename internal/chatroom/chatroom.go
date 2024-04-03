package chatroom

import "time"

type Chatroom struct {
	ID           int                   `json:"roomId"`
	Type         RoomType              `json:"roomType"`
	Participants []ChatroomParticipant `json:"participants"`
}

type ChatroomMessage struct {
	ID           int                 `json:"id"`
	RoomID       int                 `json:"roomId"`
	Content      string              `json:"content"`
	Type         MessageType         `json:"type"`
	FileLocation string              `json:"fileLocation"`
	Participant  ChatroomParticipant `json:"participant"`
	Time         time.Time           `json:"time"`
}

type ChatroomParticipant struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	ProfileImage string `json:"profile_image"`
}
