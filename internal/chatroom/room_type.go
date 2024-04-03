package chatroom

type RoomType string

const (
	Single RoomType = "single"
	Group  RoomType = "group"
	Open   RoomType = "open"
)
