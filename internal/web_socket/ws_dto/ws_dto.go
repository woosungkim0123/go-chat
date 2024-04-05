package ws_dto

import (
	"ws/internal/chatroom/ch_domain"
	"ws/internal/web_socket/ws_domain"
)

type WsJsonResponse struct {
	Action      string        `json:"action"`
	User        UserSocketDto `json:"user"`
	Message     string        `json:"message"`
	MessageType string        `json:"message_type"`
	Time        string        `json:"time"`
}

type WsJsonRequest struct {
	Action       string                         `json:"action"`
	UserID       string                         `json:"userID"`
	RoomID       string                         `json:"roomID"`
	Content      string                         `json:"content"`
	Type         ch_domain.MessageType          `json:"type"`
	FileLocation string                         `json:"fileLocation"`
	Conn         *ws_domain.WebSocketConnection `json:"-"`
}

type UserSocketDto struct {
	ID           int    `json:"ID"`
	Name         string `json:"name"`
	ProfileImage string `json:"profileImage"`
}
