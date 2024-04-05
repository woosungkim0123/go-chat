package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var allowOrigin = "http://localhost:8080"

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if origin == allowOrigin {
			return true
		}
		return false
	},
}

// 사용자가 들어와서 웹소켓을 연결하게됨
func WsEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println("wesocket upgrade failed", err)
		return
	}

	log.Printf("Client Connected to Endpoint")
	fmt.Print(ws)

	//conn := web_socket.WebSocketConnection{Conn: ws}

	//go web_socket.ListenForWs(&conn) // 고루틴으로 듣기 시작
}
