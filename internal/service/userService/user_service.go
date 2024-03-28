package userService

import (
	"errors"
	"log"
	"ws/internal/apperrors"
	"ws/internal/domain"
	"ws/internal/dto"
	"ws/internal/util/jsonReader"
)

// 비밀번호 로직 생략
func DoLogin(userId string) (*dto.UserDto, *apperrors.CustomError) {
	user, err := findByUserId(userId)
	if err != nil {
		log.Printf("login error: %v", err)
		return nil, &apperrors.CustomError{Code: apperrors.NotFoundUserNameError, Message: err.Error()}
	}

	log.Println("user logged in:", user.UserId)

	var userDto = dto.UserDto{Id: user.Id, Name: user.Name}

	return &userDto, nil
}

func findByUserId(userId string) (*domain.User, error) {
	users := getUserList()
	for _, u := range users {
		if u.UserId == userId {
			return &u, nil
		}
	}

	return nil, errors.New("user not found")
}

func getUserList() []domain.User {
	var users []domain.User
	jsonReader.ReadAndConvert("internal/store/json/users.json", &users)

	return users
}
