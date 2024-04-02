package dto

type ChatroomDto struct {
	RoomId       int              `json:"roomId"`
	RoomType     string           `json:"roomType"`
	Participants []UserDto        `json:"participants"`
	AccessUser   UserDto          `json:"accessUser"`
	Messages     []ChatMessageDto `json:"messages"`
}

type ChatMessageDto struct {
	Message string  `json:"message"`
	User    UserDto `json:"user"`
	Time    string  `json:"time"`
}

type ChatroomListDto struct {
	RoomId   int     `json:"roomId"`
	RoomType string  `json:"roomType"`
	Audience UserDto `json:"audience"`
	Message  string  `json:"message"`
}
