package router

import (
	"github.com/bmizerany/pat"
	"net/http"
	"ws/internal/common/middleware"
	"ws/internal/config/di"
)

type Router struct {
	Container *di.Container
	mux       *pat.PatternServeMux
}

func NewRouter(container *di.Container) *Router {
	mux := pat.New()
	return &Router{Container: container, mux: mux}
}

func (r *Router) Routes() http.Handler {

	r.setWebRouter("/")
	r.setFileRouter("/static/")
	r.setAuthRouter("/auth")
	r.setChatRouter("/chatroom")
	r.setWebSocketRouter("/ws")
	return r.mux
}

func (r *Router) setWebRouter(url string) {
	wh := r.Container.WebHandler
	r.mux.Get(url, middleware.AuthMiddleware(http.HandlerFunc(wh.Home)))
}

func (r *Router) setFileRouter(url string) {
	r.mux.Get(url, http.StripPrefix(url, http.FileServer(http.Dir("public"))))
}

func (r *Router) setAuthRouter(url string) {
	ah := r.Container.AuthHandler
	r.mux.Get(url+"/login", http.HandlerFunc(ah.GetLoginPage))
	r.mux.Post(url+"/login", http.HandlerFunc(ah.Login))
	r.mux.Get("/logout", http.HandlerFunc(ah.Logout))
}

func (r *Router) setChatRouter(url string) {
	ch := r.Container.ChatroomHandler
	r.mux.Get(url, middleware.AuthMiddleware(http.HandlerFunc(ch.GetChatList)))
	r.mux.Get(url+"/mine", middleware.AuthMiddleware(http.HandlerFunc(ch.GetMineChatroom)))
	r.mux.Get(url+"/single/:userID", middleware.AuthMiddleware(http.HandlerFunc(ch.GetSingleChatroom)))
	r.mux.Post(url+"/mine", middleware.AuthMiddleware(http.HandlerFunc(ch.AddMineChatroomMessage)))
}

func (r *Router) setWebSocketRouter(url string) {
	wh := r.Container.WebSocketHandler
	r.mux.Get(url, http.HandlerFunc(wh.GetWebSocketEndPoint))
}
