package ch_dto

import (
	udomain "ws/internal/auth/domain"
	"ws/internal/auth/dto"
	"ws/internal/chatroom/ch_domain"
)

type ChatroomPageDTO struct {
	ID           int                      `json:"ID"`
	Type         ch_domain.RoomType       `json:"type"`
	Participants []ChatroomParticipantDTO `json:"participants"`
	Messages     []ChatroomMessageDTO     `json:"messages"`
	AccessUser   *dto.ProfileDto          `json:"accessUser"`
}

func NewChatroomPageDTO(chatroomDTO *ChatroomDTO, accessUser *udomain.User) *ChatroomPageDTO {
	return &ChatroomPageDTO{
		ID:           chatroomDTO.ID,
		Type:         chatroomDTO.Type,
		Participants: chatroomDTO.Participants,
		Messages:     chatroomDTO.Messages,
		AccessUser:   &dto.ProfileDto{ID: accessUser.ID, Name: accessUser.Name, ProfileImage: accessUser.ProfileImage},
	}
}
