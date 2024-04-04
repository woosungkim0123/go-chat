package handler

import (
	"net/http"
	"ws/internal/auth/service"
	"ws/internal/ui/dto"
	"ws/internal/util/template"
)

type WebHandler struct {
	authService *service.AuthService
	loginPage   string
}

func NewWebHandler(service *service.AuthService) *WebHandler {
	return &WebHandler{authService: service, loginPage: "/auth/login"}
}

func (h *WebHandler) Home(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID")

	profile, err := h.authService.GetMyProfile(userID.(int))
	if err != nil {
		http.Redirect(w, r, h.loginPage+"?error=not_found_user", http.StatusSeeOther)
		return
	}

	userList := h.authService.GetUserListWithoutSelf(userID.(int))

	template.RenderWithHeader(w, "home", dto.HomePageDto{Profile: profile, Users: userList})
}
