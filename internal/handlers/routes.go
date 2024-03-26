package handlers

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func Routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(Login))
	mux.Post("/login", http.HandlerFunc(DoLogin))

	mux.Get("/ws", http.HandlerFunc(WsEndPoint))
	return mux
}
