package domain

type RoomType string

const (
	Mine   RoomType = "mine"
	Single RoomType = "single"
	Group  RoomType = "group"
	Open   RoomType = "open"
)
