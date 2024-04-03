package chatroom

import (
	"net/http"
	"ws/internal/util/converter"
	"ws/internal/util/template"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetSingleChatroom(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid")
	// TODO uid 체크로직, uid가 int가 아닐 경우 에러처리, id int 변환 실패시 에러처리
	accessUserId, _ := uid.(int)
	id := r.URL.Query().Get(":id")

	otherUserId, _ := converter.ConvertToInt(id)

	chatroomDto := h.service.GetChatroomByUserId(accessUserId, otherUserId)

	template.RenderWithHeader(w, "chatroom", &chatroomDto)
}

func (h *Handler) GetChatList(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid")
	userId, _ := converter.ConvertToInt(uid)

	chatListDto := h.service.GetChatListByUserId(userId)

	template.RenderWithHeader(w, "chatlist", &chatListDto)
}
