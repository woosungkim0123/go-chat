package ws_dto

import (
	"ws/internal/chatroom/ch_domain"
	"ws/internal/chatroom/ch_dto"
	"ws/internal/web_socket/ws_domain"
)

type WsJsonResponse struct {
	Action string                     `json:"action"`
	Data   *ch_dto.ChatroomMessageDTO `json:"data"`
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
