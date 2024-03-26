package handlers

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func Routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(Home))
	mux.Get("/ws", http.HandlerFunc(WsEndPoint))
	return mux
}
