package handler

import (
	"log"
	"net/http"
	service2 "ws/internal/auth/service"
	"ws/internal/common/apperror"
	"ws/internal/common/template"
	"ws/internal/config"
)

type AuthHandler struct {
	service   *service2.AuthService
	loginPage string
}

func NewAuthHandler(service *service2.AuthService) *AuthHandler {
	return &AuthHandler{service: service, loginPage: "/auth/login"}
}

func (*AuthHandler) GetLoginPage(w http.ResponseWriter, _ *http.Request) {
	template.Render(w, "login", nil)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, "ws-session")
	if err := r.ParseForm(); err != nil {
		http.Redirect(w, r, h.loginPage+"?error=wrong_argument", http.StatusSeeOther)
		return
	}

	loginID := r.FormValue("loginID")
	if loginID == "" {
		http.Redirect(w, r, h.loginPage+"?error=missing_userid", http.StatusSeeOther)
		return
	}

	loginDto, loginError := h.service.Login(loginID)
	if loginError != nil {
		if loginError.Code == apperror.NotFoundUserByLoginID {
			http.Redirect(w, r, h.loginPage+"?error=not_found_user", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, h.loginPage+"?error=server_error", http.StatusSeeOther)
		return
	}

	session.Values["userID"] = loginDto.ID
	if sessionError := session.Save(r, w); sessionError != nil {
		http.Redirect(w, r, h.loginPage+"?error=server_error", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, "ws-session")
	session.Options.MaxAge = -1

	if err := session.Save(r, w); err != nil {
		log.Printf("Failed to save session: %v", err)
	}

	http.Redirect(w, r, h.loginPage+"?error=logout", http.StatusSeeOther)
}
