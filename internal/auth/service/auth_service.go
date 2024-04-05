package service

import (
	"log"
	"ws/internal/auth/domain"
	"ws/internal/auth/dto"
	"ws/internal/auth/repository"
	"ws/internal/common/apperror"
)

type AuthService struct {
	repository *repository.AuthRepository
}

func NewAuthService(repository *repository.AuthRepository) *AuthService {
	return &AuthService{repository: repository}
}

// Login 로그인 - 비밀번호 로직 생략
func (s *AuthService) Login(loginID string) (*dto.LoginDto, *apperror.CustomError) {
	user, err := s.findUserByLoginID(loginID)
	if err != nil {
		return nil, err
	}

	log.Printf("user logged in: %v", user)

	return &dto.LoginDto{ID: user.ID, Name: user.Name}, nil
}

// GetUserProfile 내 프로필 가져오기
func (s *AuthService) GetUserProfile(userID int) (*dto.ProfileDto, *apperror.CustomError) {
	user, err := s.findUserByID(userID)
	if err != nil {
		return nil, err
	}

	return &dto.ProfileDto{ID: user.ID, Name: user.Name, ProfileImage: user.ProfileImage}, nil
}

// GetUserListWithoutSelf 자신을 제외한 사용자 리스트 가져오기
func (s *AuthService) GetUserListWithoutSelf(userID int) []dto.ProfileDto {
	users, err := s.repository.GetUserListWithoutSelf(userID)
	if err != nil {
		log.Printf("failed to get user list: %v", err)
		return []dto.ProfileDto{}
	}

	return s.convertUsersToProfileDtoList(users)
}

func (s *AuthService) findUserByID(userID int) (*domain.User, *apperror.CustomError) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) findUserByLoginID(loginID string) (*domain.User, *apperror.CustomError) {
	return s.repository.FindUserByLoginID(loginID)
}

func (s *AuthService) FindUserByID(userID int) (*domain.User, *apperror.CustomError) {
	return s.repository.FindByID(userID)
}

func (s *AuthService) convertUsersToProfileDtoList(users []domain.User) []dto.ProfileDto {
	userDtoList := make([]dto.ProfileDto, len(users))
	for i, user := range users {
		userDtoList[i] = dto.ProfileDto{
			ID:           user.ID,
			Name:         user.Name,
			ProfileImage: user.ProfileImage,
		}
	}
	return userDtoList
}
