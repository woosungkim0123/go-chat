package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"ws/internal/auth/domain"
	aservice "ws/internal/auth/service"
	"ws/internal/chatroom/dto"
	"ws/internal/chatroom/service"
	"ws/internal/common/converter"
	"ws/internal/common/template"
)

type ChatroomHandler struct {
	chatroomService *service.ChatroomService
	authService     *aservice.AuthService
}

func NewChatroomHandler(chatroomService *service.ChatroomService, authService *aservice.AuthService) *ChatroomHandler {
	return &ChatroomHandler{chatroomService: chatroomService, authService: authService}
}

func (h *ChatroomHandler) GetSingleChatroom(w http.ResponseWriter, r *http.Request) {
	accessUser := h.getAccessUser(w, r)
	opponentUserID := r.URL.Query().Get("userID")

	// TODO error handler -> 404나 500에러 페이지로
	chatroomDTO, err := h.chatroomService.GetChatroomByUserID(accessUser, h.convertStringToInt(opponentUserID))
	if err != nil {
		log.Printf("failed to get single chatroom: %v", err)
	}

	template.RenderWithHeader(w, "chatroom", dto.NewChatroomPageDTO(chatroomDTO, accessUser))
}

func (h *ChatroomHandler) GetChatList(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid")
	userId, _ := converter.ConvertToInt(uid)

	fmt.Print(userId)

	//chatListDto := h.service.GetChatListByUserId(userId)

	template.RenderWithHeader(w, "chatlist", nil)
}

func (h *ChatroomHandler) convertStringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}

func (h *ChatroomHandler) getAccessUser(w http.ResponseWriter, r *http.Request) *domain.User {
	accessUserID := r.Context().Value("userID").(int)
	user, err := h.authService.GetUserByID(accessUserID)
	if err != nil {
		log.Printf("failed to get access user: %v", err)
		// TODO redirect
		// http.Redirect(w, r, h.loginPage+"?error=not_found_user", http.StatusSeeOther)
	}
	return user
}
