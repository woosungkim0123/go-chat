package handlers

import (
	"github.com/bmizerany/pat"
	"net/http"
	"ws/internal/middleware"
)

func Routes() http.Handler {
	mux := pat.New()
	mux.Get("/", middleware.AuthMiddleware(http.HandlerFunc(Home)))
	mux.Get("/login", http.HandlerFunc(Login))
	mux.Post("/login", http.HandlerFunc(DoLogin))

	mux.Get("/ws", http.HandlerFunc(WsEndPoint))
	return mux
}
