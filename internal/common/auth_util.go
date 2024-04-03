package common

import (
	"errors"
	"net/http"
)

func GetAuthentication(r *http.Request) (int, error) {
	userIdStr := r.Context().Value("uid")
	if userIdStr == nil {
		return 0, errors.New("uid is nil")
	}

	userId, ok := userIdStr.(int)
	if !ok {
		return 0, errors.New("uid is not int")
	}

	return userId, nil
}
