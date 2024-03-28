package handlers

import (
	"net/http"
	"ws/config"
	"ws/internal/apperrors"
	"ws/internal/service/userService"
	"ws/internal/util/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "login", nil)
}

func DoLogin(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, "ws-session")
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/?error=wrong_argument", http.StatusSeeOther)
		return
	}
	userId := r.FormValue("userid")
	if userId == "" {
		http.Redirect(w, r, "/login?error=missing_userid", http.StatusSeeOther)
		return
	}

	userDto, loginError := userService.DoLogin(userId)
	if loginError != nil {
		if loginError.Code == apperrors.NotFoundUserNameError {
			http.Redirect(w, r, "/login?error=not_found_user", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/login?error=login_failed", http.StatusSeeOther)
		return
	}

	session.Values["id"] = userDto.Id
	if err := session.Save(r, w); err != nil {
		http.Redirect(w, r, "/login?error=server_error", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
