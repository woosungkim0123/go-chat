package handlers

import (
	"ws/config/database"
	ah "ws/internal/auth/handler"
	"ws/internal/auth/repository"
	"ws/internal/auth/service"
	ch "ws/internal/chatroom/handler"
	repository2 "ws/internal/chatroom/repository"
	cs "ws/internal/chatroom/service"
	wh "ws/internal/ui/handler"
)

type Container struct {
	DBManager       *database.DBManager
	AuthHandler     *ah.AuthHandler
	AuthService     *service.AuthService
	ChatroomHandler *ch.ChatroomHandler
	ChatroomService *cs.ChatroomService
	WebHandler      *wh.WebHandler
}

func NewContainer() *Container {
	dbManager := database.NewDBManager("my.db")

	authRepository := repository.NewAuthRepository(dbManager.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := ah.NewAuthHandler(authService)

	chatroomRepository := repository2.NewRepository(dbManager.DB)
	chatroomService := cs.NewChatroomService(chatroomRepository, authService)
	chatroomHandler := ch.NewChatroomHandler(chatroomService, authService)

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
