package userService

import (
	"ws/internal/dto"
	"ws/internal/util/jsonReader"
)

func GetUserList() []dto.UserDto {
	var users []dto.UserDto
	jsonReader.ReadAndConvert("internal/store/json/users.json", &users)
	return users
}
