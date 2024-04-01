package web_socket

type WsJsonResponse struct {
	Action          string `json:"action"`
	ConnectedUserId string `json:"connectedUserId"`
	Message         string `json:"message"`
	MessageType     string `json:"message_type"`
}

type WsJsonRequest struct {
	Action  string               `json:"action"`
	UserId  string               `json:"userId"`
	RoomId  string               `json:"roomId"`
	Message string               `json:"message"`
	Conn    *WebSocketConnection `json:"-"`
}
