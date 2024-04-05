package web_socket

//var wsChan = make(chan WsJsonRequest)
//
//var ActiveChatrooms = make(map[int]ChatroomSession)
//
//// 요청을 보낼때 웹소켓을 통해 메시지를 보내게됨
//func ListenForWs(conn *WebSocketConnection) {
//	log.Println("Listening to websocket~~")
//	defer func() {
//		if r := recover(); r != nil {
//			log.Println("Error = ", fmt.Sprintf("%v", r))
//		}
//	}()
//
//	var request WsJsonRequest
//	for {
//		apperror := conn.ReadJSON(&request)
//		if apperror != nil {
//			continue
//		}
//		request.Conn = conn
//		wsChan <- request
//	}
//}
//
//// wsChan 채널로부터 메세지를 지속적으로 수신하고 받은 메세지를 처리하여 모든 클라이언트에게 브로드 캐스트하는 역할
//func ListenToWsChannel() {
//	log.Println("Listening to ws channel")
//	var response WsJsonResponse
//
//	for {
//		request := <-wsChan // --blocking
//
//		switch request.Action {
//		case "join":
//			roomId, userId, apperror := parseIds(request.RoomId, request.UserId)
//			if apperror != nil {
//				log.Println(apperror)
//				continue
//			}
//			accessUser := auth.FindUser(userId)
//			if accessUser == nil {
//				log.Printf("User not found: %d\n", userId)
//				continue
//			}
//
//			// 채팅방 세션이 있으면 해당 세션을 반환하고, 없으면 새로운 채팅방 세션을 생성
//			chatroomSession, exists := ActiveChatrooms[roomId]
//			if !exists {
//				chatroomSession = ChatroomSession{
//					Participants: make(map[int]UserSession),
//				}
//				ActiveChatrooms[roomId] = chatroomSession
//			}
//
//			// 사용자 세션 추가
//			chatroomSession.Participants[userId] = UserSession{Conn: request.Conn, Username: accessUser.Name}
//			log.Printf("User %d joined room %d", userId, roomId)
//
//			location, apperror := time.LoadLocation("Asia/Seoul")
//			currentTime := time.Now().UTC().In(location)
//			currentTimeFormat := currentTime.Format("1/02 15:04:05")
//			response.Action = "join"
//			response.User = UserSocketDto{ID: accessUser.ID, Name: accessUser.Name}
//			response.Time = currentTimeFormat
//			response.Message = fmt.Sprintf("User %d joined room %d", userId, roomId)
//			broadcastToRoom(request, response)
//
//		case "left":
//			roomId, userId, apperror := parseIds(request.RoomId, request.UserId)
//			if apperror != nil {
//				log.Println(apperror)
//				continue
//			}
//
//			leftUser := auth.FindUser(userId)
//			if leftUser == nil {
//				log.Printf("User not found: %d\n", userId)
//				continue
//			}
//
//			chatroomSession, exists := ActiveChatrooms[roomId]
//			if !exists {
//				log.Printf("Room ID %d not found\n", roomId)
//				return
//			}
//
//			delete(chatroomSession.Participants, userId)
//
//			response.Action = "left"
//			response.User = UserSocketDto{ID: leftUser.ID, Name: leftUser.Name}
//			response.Message = fmt.Sprintf("User %d left room %d", userId, roomId)
//			broadcastToRoom(request, response)
//
//		case "broadcast":
//			roomId, userId, apperror := parseIds(request.RoomId, request.UserId)
//			if apperror != nil {
//				log.Println(apperror)
//				continue
//			}
//			log.Println("Broadcasting message to room", roomId)
//			accessUser := auth.FindUser(userId)
//			if accessUser == nil {
//				log.Printf("User not found: %d\n", userId)
//				continue
//			}
//			location, apperror := time.LoadLocation("Asia/Seoul")
//			currentTime := time.Now().UTC().In(location)
//
//			// 데이터베이스에 저장(임시적으로 Json)
//			chatroomMessage := domain.ChatroomMessage{
//				Message: request.Message,
//				User:    *accessUser,
//				Time:    currentTime,
//			}
//			chatService.SaveMessage(roomId, chatroomMessage)
//
//			response.Action = "broadcast"
//
//			currentTimeFormat := currentTime.Format("1/02 15:04:05")
//			response.User = UserSocketDto{ID: accessUser.ID, Name: accessUser.Name}
//			response.Time = currentTimeFormat
//			response.Message = request.Message
//			broadcastToRoom(request, response)
//		}
//	}
//}

//func broadcastToRoom(request WsJsonRequest, response WsJsonResponse) {
//	roomId, apperror := strconv.Atoi(request.RoomId) // RoomId를 int로 변환
//	if apperror != nil {
//		log.Printf("Invalid room ID: %s\n", request.RoomId)
//		return
//	}
//
//	chatroomSession, exists := ActiveChatrooms[roomId]
//	if !exists {
//		log.Printf("Room ID %d not found\n", roomId)
//		return
//	}
//
//	var toDelete []int
//	for uId, userSession := range chatroomSession.Participants {
//		apperror := userSession.Conn.WriteJSON(response)
//		if apperror != nil {
//			log.Println("Error sending message to websocket:", apperror)
//			toDelete = append(toDelete, uId)
//			continue
//		}
//	}
//
//	// 클라이언트와 연결이 유효하지 않다고 판단되면 해당 클라이언트를 삭제합니다.
//	for _, uId := range toDelete {
//		delete(chatroomSession.Participants, uId) // 안전하게 삭제
//		log.Printf("Removed user %d due to error\n", uId)
//	}
//}
//
//func parseIds(roomIdStr, userIdStr string) (int, int, error) {
//	roomId, apperror := strconv.Atoi(roomIdStr)
//	if apperror != nil {
//		return 0, 0, fmt.Errorf("Invalid room ID: %s", roomIdStr)
//	}
//
//	userId, apperror := strconv.Atoi(userIdStr)
//	if apperror != nil {
//		return roomId, 0, fmt.Errorf("Invalid user ID: %s", userIdStr)
//	}
//
//	return roomId, userId, nil
//}
