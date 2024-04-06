package ch_dto

import (
	"ws/internal/auth/domain"
	"ws/internal/chatroom/ch_domain"
	"ws/internal/common/util"
)

type MineChatroomRequestDTO struct {
	RoomID       int                   `json:"roomID"`
	Content      string                `json:"content"`
	Type         ch_domain.MessageType `json:"type"`
	FileLocation string                `json:"fileLocation"`
}

func (dto *MineChatroomRequestDTO) ToDomain(user *domain.User) *ch_domain.ChatroomMessage {
	return &ch_domain.ChatroomMessage{
		RoomID:       dto.RoomID,
		Content:      dto.Content,
		Type:         dto.Type,
		FileLocation: dto.FileLocation,
		Participant: ch_domain.ChatroomParticipant{
			ID:           user.ID,
			Name:         user.Name,
			ProfileImage: user.ProfileImage,
		},
		Time: util.GetCurrentDate(),
	}
}
