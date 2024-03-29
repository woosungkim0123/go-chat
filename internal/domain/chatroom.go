package domain

type Chatroom struct {
	RoomID       int    `json:"roomId"`
	RoomType     string `json:"roomType"`
	Participants []User `json:"participants"`
}
