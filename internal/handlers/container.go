package handlers

import (
	"ws/config/database"
	"ws/internal/chatroom"
)

type Container struct {
	DBManager       *database.DBManager
	ChatroomHandler *chatroom.Handler
	ChatroomService *chatroom.Service
}

func NewContainer() *Container {
	dbManager := database.NewDBManager("my.db")

	chatroomRepository := chatroom.NewRepository(dbManager.DB)
	chatroomService := chatroom.NewService(chatroomRepository)
	chatroomHandler := chatroom.NewHandler(chatroomService)

	return &Container{
		DBManager:       dbManager,
		ChatroomHandler: chatroomHandler,
		ChatroomService: chatroomService,
	}
}
