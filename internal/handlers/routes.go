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
	r.mux.Get("/", middleware.AuthMiddleware(http.HandlerFunc(Home)))
	r.setAuthRouter("/user")
	r.setChatRouter("/chatroom")
	r.mux.Get("/ws", http.HandlerFunc(WsEndPoint))
	return r.mux
}

func (r *Router) setAuthRouter(url string) {
	r.mux.Get(url, http.HandlerFunc(Login))
	r.mux.Post(url, http.HandlerFunc(DoLogin))
	r.mux.Get("/logout", http.HandlerFunc(DoLogout))
}

func (r *Router) setChatRouter(url string) {
	ch := r.Container.ChatroomHandler
	r.mux.Get(url, middleware.AuthMiddleware(http.HandlerFunc(ch.GetChatList)))
	r.mux.Get(url+"/:id", middleware.AuthMiddleware(http.HandlerFunc(ch.GetSingleChatroom)))
}
