package handlers

import (
	"github.com/bmizerany/pat"
	"net/http"
	"ws/internal/middleware"
)

type Router struct {
	Container *Container
	mux       *pat.PatternServeMux
}

func NewRouter(container *Container) *Router {
	mux := pat.New()
	return &Router{Container: container, mux: mux}
}

func (r *Router) Routes() http.Handler {
	r.setWebRouter("/")
	r.setAuthRouter("/auth")
	r.setChatRouter("/chatroom")
	r.mux.Get("/ws", http.HandlerFunc(WsEndPoint))
	return r.mux
}

func (r *Router) setWebRouter(url string) {
	wh := r.Container.WebHandler
	r.mux.Get(url, middleware.AuthMiddleware(http.HandlerFunc(wh.Home)))
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
	r.mux.Get(url+"/single", middleware.AuthMiddleware(http.HandlerFunc(ch.GetSingleChatroom)))
}
