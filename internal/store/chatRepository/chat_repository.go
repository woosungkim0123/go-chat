package chatRepository

import (
	"ws/internal/domain"
	"ws/internal/util/jsonReader"
)

func GetAllChatroom() []domain.Chatroom {
	var chatroomList []domain.Chatroom
	jsonReader.ReadAndConvert("internal/store/json/chatroom.json", &chatroomList)
	return chatroomList
}

func AddChatroom(chatroom domain.Chatroom) {
	chatroomList := GetAllChatroom()
	chatroomList = append(chatroomList, chatroom)
	jsonReader.Write("internal/store/json/chatroom.json", chatroomList)
}
