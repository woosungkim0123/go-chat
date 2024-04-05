package main

import (
	"log"
	"net/http"
	"ws/internal/config"
	"ws/internal/config/database"
	"ws/internal/config/di"
	"ws/internal/config/router"
)

func main() {
	startServer()
}

func startServer() {
	serverConfig := config.Configuration.ServerConfig
	log.Printf("Starting server on %s:%s", serverConfig.Host, serverConfig.Port)

	container := di.NewContainer()
	defer container.DBManager.Close()

	database.NewInitializer(container.DBManager).Init() // 초기 데이터 생성 (스키마, 임시 데이터)

	go container.WebSocketService.ListenToWsChannel()

	if err := http.ListenAndServe(":"+serverConfig.Port, router.NewRouter(container).Routes()); err != nil {
		log.Printf("Error starting server: %s", err)
		panic(err)
	}
}
