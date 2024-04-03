package chatroom

import "time"

type Chatroom struct {
	RoomID   int    `json:"roomId"`
	RoomType string `json:"roomType"`
	//Participants []User            `json:"participants"`
	Messages []ChatroomMessage `json:"messages"`
}

type ChatroomMessage struct {
	Message string `json:"message"`
	//User    User      `json:"user"`
	Time time.Time `json:"time"`
}
