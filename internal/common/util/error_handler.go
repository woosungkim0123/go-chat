package util

import (
	"fmt"
	"log"
	"ws/internal/common/apperror"
)

func HandleError(msg string, code apperror.ErrorCode) error {
	log.Println(msg)
	return fmt.Errorf(string(code))
}
