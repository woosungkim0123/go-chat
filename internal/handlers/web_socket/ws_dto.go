package web_socket

import "ws/internal/auth/domain"

type WsJsonResponse struct {
	Action      string        `json:"action"`
	User        UserSocketDto `json:"user"`
	Message     string        `json:"message"`
	MessageType string        `json:"message_type"`
	Time        string        `json:"time"`
}

type WsJsonRequest struct {
	Action  string                      `json:"action"`
	UserId  string                      `json:"userId"`
	RoomId  string                      `json:"roomId"`
	Message string                      `json:"message"`
	Conn    *domain.WebSocketConnection `json:"-"`
}

type UserSocketDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
