package handlers

import (
	"github.com/bmizerany/pat"
	"net/http"
	"ws/internal/middleware"
)

func Routes() http.Handler {
	mux := pat.New()
	mux.Get("/", middleware.AuthMiddleware(http.HandlerFunc(Home)))

	authRouter("/login", mux)

	chatRouter("/chat", mux)
	mux.Get("/ws", http.HandlerFunc(WsEndPoint))
	return mux
}

func authRouter(routerUrl string, mux *pat.PatternServeMux) {
	mux.Get(routerUrl, http.HandlerFunc(Login))
	mux.Post(routerUrl, http.HandlerFunc(DoLogin))
	mux.Get("/logout", http.HandlerFunc(DoLogout))
}

func chatRouter(routerUrl string, mux *pat.PatternServeMux) {
	mux.Get(routerUrl, middleware.AuthMiddleware(http.HandlerFunc(ChatList)))
	mux.Get(routerUrl+"/:id", middleware.AuthMiddleware(http.HandlerFunc(Chat)))

}
