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

	if err := http.ListenAndServe(serverInfo.Host+":"+serverInfo.Port, handlers.Routes()); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
