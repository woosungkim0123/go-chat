package handlers

import (
	"net/http"
	"ws/internal/dto"
	"ws/internal/service/userService"
	"ws/internal/util/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// 세션이 있으면?
	if r.Context().Value("session") != nil {

	}

	template.Render(w, "login", nil)
}

func DoLogin(w http.ResponseWriter, r *http.Request) {
	// 폼 파싱
	err := r.ParseForm()
	if err != nil {
		// 에러 시 /login 으로 리디렉션
		http.Redirect(w, r, "?error=wrong_argument", http.StatusSeeOther)
		return
	}

	// "userid" 필드 값 추출
	userid := r.FormValue("userid") // 또는 r.Form.Get("userid")

	// userid 값 검증
	if userid == "" {
		// userid가 누락된 경우, 에러와 함께 리디렉션
		http.Redirect(w, r, "?error=missing_userid", http.StatusSeeOther)
		return
	}

	// 여기서 로그인 로직 구현...

	// 모든 처리가 성공적으로 끝나면, 메인 페이지로 리디렉션
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

type HomeData struct {
	Title string
	Users []dto.UserDto
}

func Home(w http.ResponseWriter, r *http.Request) {
	users := userService.GetUserList()

	template.RenderWithHeader(w, "home", HomeData{
		Title: "Home",
		Users: users,
	})
}
