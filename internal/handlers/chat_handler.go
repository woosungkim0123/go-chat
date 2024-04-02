package handlers

import (
	"net/http"
	"ws/internal/service/chatService"
	"ws/internal/util/converter"
	"ws/internal/util/template"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid")
	// TODO uid 체크로직, uid가 int가 아닐 경우 에러처리, id int 변환 실패시 에러처리
	accessUserId, _ := uid.(int)
	id := r.URL.Query().Get(":id")

	otherUserId, _ := converter.ConvertToInt(id)

	chatroomDto := chatService.GetChatroomByUserId(accessUserId, otherUserId)

	template.RenderWithHeader(w, "chatroom", &chatroomDto)
}

func ChatList(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid")
	userId, _ := converter.ConvertToInt(uid)

	chatlist := chatService.GetChatList(userId)

	template.RenderWithHeader(w, "chatlist", &chatlist)
}
