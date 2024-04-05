package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/bbolt"
	"log"
	"strconv"
	"ws/internal/auth/domain"
	apperror2 "ws/internal/common/apperror"
)

type AuthRepository struct {
	db   *bbolt.DB
	user string
}

func NewAuthRepository(db *bbolt.DB) *AuthRepository {
	return &AuthRepository{db: db, user: "user"}
}

// FindUserByLoginID 로그인 ID로 사용자 찾기
// (bolt DB 값을 찾을 때 순회 중단하려면 에러를 반환해야 함)
func (r *AuthRepository) FindUserByLoginID(LoginID string) (*domain.User, *apperror2.CustomError) {
	var user *domain.User

	dbError := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(r.user))
		if b == nil {
			log.Printf("bucket not found: %s", r.user)
			return fmt.Errorf(string(apperror2.NotFoundBucket))
		}

		var FoundUserAndStopIterator = errors.New("found user")
		foundUserError := b.ForEach(func(k, v []byte) error {
			var tempUser domain.User
			if jsonError := json.Unmarshal(v, &tempUser); jsonError != nil {
				log.Printf("failed to unmarshal user data: %v", jsonError)
				return fmt.Errorf(string(apperror2.FailJsonUnmarshal))
			}

			if tempUser.LoginID == LoginID {
				user = &tempUser
				return FoundUserAndStopIterator
			}

			return nil
		})

		if foundUserError == nil {
			log.Printf("no user found with login ID: %s", LoginID)
			return fmt.Errorf(string(apperror2.NotFoundUserByLoginID))
		} else if errors.Is(foundUserError, FoundUserAndStopIterator) {
			return nil // 사용자를 찾았으므로, 에러 없음
		} else {
			return foundUserError // 다른 에러 처리
		}
	})

	if dbError != nil {
		if dbError.Error() == string(apperror2.NotFoundUserByLoginID) {
			return nil, &apperror2.CustomError{Code: apperror2.NotFoundUserByLoginID, Message: "LoginID에 해당하는 유저를 찾을 수 없습니다."}
		}
		return nil, &apperror2.CustomError{Code: apperror2.DataBaseProblem, Message: "데이터베이스에 문제가 발생했습니다."}
	}

	return user, nil
}

// FindByID userID로 사용자 찾기
func (r *AuthRepository) FindByID(userID int) (*domain.User, *apperror2.CustomError) {
	var user *domain.User

	dbError := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(r.user))
		if b == nil {
			log.Printf("bucket not found: %s", r.user)
			return fmt.Errorf(string(apperror2.NotFoundBucket))
		}

		userData := b.Get([]byte(strconv.Itoa(userID)))
		if userData == nil {
			log.Printf("no user found with ID: %d", userID)
			return fmt.Errorf(string(apperror2.NotFoundUserByID))
		}

		if err := json.Unmarshal(userData, &user); err != nil {
			log.Printf("failed to unmarshal user data: %v", err)
			return fmt.Errorf(string(apperror2.FailJsonUnmarshal))
		}
		return nil
	})

	if dbError != nil {
		if dbError.Error() == string(apperror2.NotFoundUserByID) {
			return nil, &apperror2.CustomError{Code: apperror2.NotFoundUserByID, Message: "ID에 해당하는 유저를 찾을 수 없습니다."}
		}
		return nil, &apperror2.CustomError{Code: apperror2.DataBaseProblem, Message: "데이터베이스에서 유저를 찾는 중 문제가 발생했습니다."}
	}

	return user, nil
}

// GetUserListWithoutSelf 자기 자신을 제외한 모든 사용자 가져오기
func (r *AuthRepository) GetUserListWithoutSelf(userID int) ([]domain.User, *apperror2.CustomError) {
	var users []domain.User

	dbError := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(r.user))
		if b == nil {
			log.Printf("bucket not found: %s", r.user)
			return fmt.Errorf(string(apperror2.NotFoundBucket))
		}

		return b.ForEach(func(k, v []byte) error {
			var user domain.User
			if err := json.Unmarshal(v, &user); err != nil {
				log.Printf("failed to unmarshal user data: %v", err)
				return fmt.Errorf(string(apperror2.FailJsonUnmarshal))
			}

			if user.ID != userID {
				users = append(users, user)
			}

			return nil
		})
	})

	if dbError != nil {
		return nil, &apperror2.CustomError{Code: apperror2.DataBaseProblem, Message: "데이터베이스에서 유저를 찾는 중 문제가 발생했습니다."}
	}

	return users, nil
}
