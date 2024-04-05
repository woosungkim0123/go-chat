package ch_domain

type MessageType string

const (
	Text  MessageType = "text"
	Emoji MessageType = "emoji"
	Image MessageType = "image"
	File  MessageType = "file"
)
