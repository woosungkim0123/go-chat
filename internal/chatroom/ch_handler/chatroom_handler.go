package ch_handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"ws/internal/auth/domain"
	aservice "ws/internal/auth/service"
	"ws/internal/chatroom/ch_dto"
	"ws/internal/chatroom/ch_service"
	"ws/internal/common/apperror"
	"ws/internal/common/converter"
	"ws/internal/common/response"
	"ws/internal/common/template"
	"ws/internal/common/util"
)

type ChatroomHandler struct {
	chatroomService *ch_service.ChatroomService
	authService     *aservice.AuthService
}

func NewChatroomHandler(chatroomService *ch_service.ChatroomService, authService *aservice.AuthService) *ChatroomHandler {
	return &ChatroomHandler{chatroomService: chatroomService, authService: authService}
}

// GetMineChatroom 자신만의 채팅방을 가져온다.
func (h *ChatroomHandler) GetMineChatroom(w http.ResponseWriter, r *http.Request) {
	accessUser := h.getAccessUser(w, r)

	// TODO error handler -> 404나 500에러 페이지로
	chatroomDTO, err := h.chatroomService.GetMineChatroom(accessUser)
	if err != nil {
		log.Printf("failed to get single chatroom: %v", err)
	}

	template.RenderWithHeader(w, "chatroom_mine", ch_dto.NewChatroomPageDTO(chatroomDTO, accessUser))
}

// AddMineChatroomMessage 자신만의 채팅방에 메세지를 추가한다.
func (h *ChatroomHandler) AddMineChatroomMessage(w http.ResponseWriter, r *http.Request) {
	var msg ch_dto.MineChatroomRequestDTO
	if err := util.GetBodyData(r, &msg); err != nil {
		response.NewAPIResponse(false, "failed to add chatroom message", nil).SendJSON(w, http.StatusBadRequest)
		return
	}
	accessUser := h.getAccessUser(w, r)

	chatroomMessageDto, err := h.chatroomService.SaveMessage(msg.ToDomain(accessUser))

	if err != nil {
		response.NewAPIResponse(false, "failed to add chatroom message", nil).SendJSON(w, http.StatusInternalServerError)
		return
	}

	response.NewAPIResponse(true, "success", chatroomMessageDto).SendJSON(w, http.StatusOK)
}

// GetSingleChatroom 상대방과의 채팅방을 가져온다.
func (h *ChatroomHandler) GetSingleChatroom(w http.ResponseWriter, r *http.Request) {
	accessUser := h.getAccessUser(w, r)
	opponentUserID := r.URL.Query().Get(":userID")

	// TODO error handler -> 404나 500에러 페이지로
	chatroomDTO, err := h.chatroomService.GetSingleChatroom(accessUser, h.convertStringToInt(opponentUserID))
	if err != nil {
		if err.Code == apperror.WrongAccessMineChatroom {
			http.Redirect(w, r, "/chatroom/mine", http.StatusSeeOther)
			return
		}
		log.Printf("failed to get single chatroom: %v", err)
	}

	template.RenderWithHeader(w, "chatroom", ch_dto.NewChatroomPageDTO(chatroomDTO, accessUser))
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

func (h *ChatroomHandler) getAccessUser(_ http.ResponseWriter, r *http.Request) *domain.User {
	accessUserID := r.Context().Value("userID").(int)
	user, err := h.authService.FindUserByID(accessUserID)
	if err != nil {
		log.Printf("failed to get access user: %v", err)
		// TODO redirect
		// http.Redirect(w, r, h.loginPage+"?error=not_found_user", http.StatusSeeOther)
	}
	return user
}
