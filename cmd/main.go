package main

import (
	"log"
	"net/http"
	"ws/config"
	"ws/internal/handlers"
	"ws/internal/handlers/web_socket"
)

func main() {
	go web_socket.ListenToWsChannel()
	startServer()
}

func startServer() {
	serverConfig := config.Configuration.ServerConfig
	log.Printf("Starting server on %s:%s", serverConfig.Host, serverConfig.Port)

	container := handlers.NewContainer()
	defer container.DBManager.Close()

	if err := http.ListenAndServe(":"+serverConfig.Port, handlers.NewRouter(container).Routes()); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
