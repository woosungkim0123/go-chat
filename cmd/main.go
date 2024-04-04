package main

import (
	"log"
	"net/http"
	"ws/config"
	"ws/config/database"
	"ws/internal/handlers"
)

func main() {
	//go web_socket.ListenToWsChannel()
	startServer()
}

func startServer() {
	serverConfig := config.Configuration.ServerConfig
	log.Printf("Starting server on %s:%s", serverConfig.Host, serverConfig.Port)

	container := handlers.NewContainer()
	defer container.DBManager.Close()

	database.NewInitializer(container.DBManager).Init() // 초기 데이터 생성 (스키마, 임시 데이터)

	if err := http.ListenAndServe(":"+serverConfig.Port, handlers.NewRouter(container).Routes()); err != nil {
		log.Printf("Error starting server: %s", err)
		panic(err)
	}
}
