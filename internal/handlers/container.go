package handlers

import (
	"ws/config/database"
	ah "ws/internal/auth/handler"
	"ws/internal/auth/repository"
	"ws/internal/auth/service"
	"ws/internal/chatroom"
	ch "ws/internal/chatroom/handler"
	wh "ws/internal/ui/handler"
)

type Container struct {
	DBManager       *database.DBManager
	AuthHandler     *ah.AuthHandler
	AuthService     *service.AuthService
	ChatroomHandler *ch.ChatroomHandler
	ChatroomService *chatroom.Service
	WebHandler      *wh.WebHandler
}

func NewContainer() *Container {
	dbManager := database.NewDBManager("my.db")

	authRepository := repository.NewAuthRepository(dbManager.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := ah.NewAuthHandler(authService)

	chatroomRepository := chatroom.NewRepository(dbManager.DB)
	chatroomService := chatroom.NewService(chatroomRepository)
	chatroomHandler := ch.NewChatroomHandler(chatroomService)

	webHandler := wh.NewWebHandler(authService)

	return &Container{
		DBManager:       dbManager,
		AuthHandler:     authHandler,
		AuthService:     authService,
		ChatroomHandler: chatroomHandler,
		ChatroomService: chatroomService,
		WebHandler:      webHandler,
	}
}
