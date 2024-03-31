package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sort"
)

var wsChan = make(chan WsPayload)

var clients = make(map[WebSocketConnection]string)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		// 전부다 허용
		if origin == "http://localhost:8080" {
			return true
		}
		return true
	},
}

type WebSocketConnection struct {
	*websocket.Conn
}

// Action : What Server should do (e.g. Join, Message, Leave, ...)
// WsJsonResponse defines the response sent back from websocket
type WsJsonResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

type WsPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}

// 사용자가 들어와서 웹소켓을 연결하게됨
func WsEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	// TODO 에러 발생시 로그 찍고, 연결 실패 응답 보내기
	if err != nil {
		log.Println("wesocket upgrade failed", err)
		return
	}

	log.Printf("Client Connected to Endpoint")

	conn := WebSocketConnection{Conn: ws}

	go ListenForWs(&conn) // 고루틴으로 듣기 시작
}

// 요청을 보낼때 웹소켓을 통해 메시지를 보내게됨
func ListenForWs(conn *WebSocketConnection) {
	log.Println("Listening to websocket~~")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error = ", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

// wsChan 채널로부터 메세지를 지속적으로 수신하고 받은 메세지를 처리하여 모든 클라이언트에게 브로드 캐스트하는 역할
func ListenToWsChannel() {
	var response WsJsonResponse

	for {
		e := <-wsChan

		//response.Action = "Got here"
		//response.Message = fmt.Sprintf("Some message, and action was %s", e.Action)
		//broadcastToAll(response)

		switch e.Action {
		case "username":
			// get a list of all clients and send it back
			clients[e.Conn] = e.Username
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			broadcastToAll(response)

		case "left":
			response.Action = "list_users"
			delete(clients, e.Conn)
			log.Println("Client left" + e.Username)

			users := getUserList()
			response.ConnectedUsers = users
			broadcastToAll(response)

		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s", e.Username, e.Message)
			broadcastToAll(response)
		}
	}
}

func getUserList() []string {
	var userList []string
	for _, x := range clients {
		if x != "" {
			userList = append(userList, x)
		}
	}
	sort.Strings(userList)
	return userList
}

func broadcastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("websocket err")
			_ = client.Close()
			delete(clients, client)
		}
	}
}
