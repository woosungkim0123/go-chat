package handlers

import (
	"net/http"
	"ws/internal/dto"
	"ws/internal/service/userService"
	"ws/internal/util/template"
)

type HomeData struct {
	MyProfile *dto.UserDto
	Users     []dto.UserDto
}

func Home(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid").(int)

	myProfile := userService.GetMyProfile(uid)
	chatUserList := userService.GetChatList(uid)

	template.RenderWithHeader(w, "home", HomeData{
		MyProfile: myProfile,
		Users:     chatUserList,
	})
}
