package ws_handler

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"ws/internal/web_socket/ws_domain"
	"ws/internal/web_socket/ws_service"
)

type WebSocketHandler struct {
	webSocketService  *ws_service.WebSocketService
	upgradeConnection websocket.Upgrader
}

func NewWebSocketHandler(webSocketService *ws_service.WebSocketService, upgradeConnection websocket.Upgrader) *WebSocketHandler {
	return &WebSocketHandler{
		webSocketService:  webSocketService,
		upgradeConnection: upgradeConnection,
	}
}

func (h *WebSocketHandler) GetWebSocketEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := h.upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket upgrade failed", err)
		return
	}
	log.Printf("Client Connected to Endpoint")

	conn := ws_domain.WebSocketConnection{Conn: ws}

	go h.webSocketService.ListenForWebSocket(&conn) // 고루틴으로 듣기 시작
}
