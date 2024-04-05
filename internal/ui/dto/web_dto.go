package dto

import "ws/internal/auth/dto"

type HomePageDto struct {
	Profile *dto.ProfileDto  `json:"profile"`
	Users   []dto.ProfileDto `json:"users"`
}
