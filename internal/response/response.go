package response

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse 구조체는 성공 응답의 JSON 구조를 정의합니다.
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse 구조체는 에러 응답의 JSON 구조를 정의합니다.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Success 함수는 성공 응답을 표준화된 JSON 형식으로 반환합니다.
func Success(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: message,
		Data:    data,
	})
}

// Error 함수는 에러 응답을 표준화된 JSON 형식으로 반환합니다.
func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
	})
}
