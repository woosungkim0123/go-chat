package dto

type ChatroomDto struct {
	RoomID       int       `json:"roomId"`
	RoomType     string    `json:"roomType"`
	Participants []UserDto `json:"participants"`
	AccessUser   UserDto   `json:"accessUser"`
	OtherUser    UserDto   `json:"otherUser"`
}
