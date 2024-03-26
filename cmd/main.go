package main

import (
	"log"
	"net/http"
	"ws/config"
	"ws/internal/handlers"
)

func main() {
	go handlers.ListenToWsChannel()

	startServer()
}

func startServer() {
	serverInfo := config.Configuration.ServerInfo
	log.Printf("Starting server on %s:%s", serverInfo.Host, serverInfo.Port)

	err := http.ListenAndServe(serverInfo.Host+":"+serverInfo.Port, handlers.Routes())
	if err != nil {
		log.Printf("Error starting server: %s", err)
	}
}
