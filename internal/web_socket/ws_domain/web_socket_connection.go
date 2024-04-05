package ws_domain

import (
	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	*websocket.Conn
}

type UserSession struct {
	Conn     *WebSocketConnection
	UserName string
}

type ChatroomSession struct {
	Participants map[int]UserSession // key: 사용자 ID, value: 사용자 세션
}
