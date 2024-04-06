package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewAPIResponse(success bool, message string, data interface{}) *APIResponse {
	return &APIResponse{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func (r *APIResponse) SendJSON(w http.ResponseWriter, statusCode int) {
	// 응답 헤더 설정
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Println("Error encoding JSON response:", err)
	}
}
