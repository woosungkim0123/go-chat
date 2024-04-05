package di

import (
	ah "ws/internal/auth/handler"
	"ws/internal/auth/repository"
	"ws/internal/auth/service"
	"ws/internal/chatroom/ch_handler"
	"ws/internal/chatroom/ch_repository"
	"ws/internal/chatroom/ch_service"
	"ws/internal/config"
	"ws/internal/config/database"
	wh "ws/internal/ui/handler"
	"ws/internal/web_socket/ws_handler"
	"ws/internal/web_socket/ws_service"
)

type Container struct {
	DBManager        *database.DBManager
	AuthHandler      *ah.AuthHandler
	AuthService      *service.AuthService
	ChatroomHandler  *ch_handler.ChatroomHandler
	ChatroomService  *ch_service.ChatroomService
	WebHandler       *wh.WebHandler
	WebSocketService *ws_service.WebSocketService
	WebSocketHandler *ws_handler.WebSocketHandler
}

func NewContainer() *Container {
	dbManager := database.NewDBManager("my.db")

	authRepository := repository.NewAuthRepository(dbManager.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := ah.NewAuthHandler(authService)

	chatroomRepository := ch_repository.NewChatroomRepository(dbManager.DB)
	chatroomService := ch_service.NewChatroomService(chatroomRepository, authService)
	chatroomHandler := ch_handler.NewChatroomHandler(chatroomService, authService)

	webHandler := wh.NewWebHandler(authService)

	webSocketService := ws_service.NewWebSocketService(authService, chatroomService)
	webSocketHandler := ws_handler.NewWebSocketHandler(webSocketService, config.UpgradeConnection)

	return &Container{
		DBManager:        dbManager,
		AuthHandler:      authHandler,
		AuthService:      authService,
		ChatroomHandler:  chatroomHandler,
		ChatroomService:  chatroomService,
		WebHandler:       webHandler,
		WebSocketService: webSocketService,
		WebSocketHandler: webSocketHandler,
	}
}
