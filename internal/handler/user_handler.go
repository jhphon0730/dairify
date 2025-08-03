package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jhphon0730/dairify/internal/dto"
	"github.com/jhphon0730/dairify/internal/response"
	"github.com/jhphon0730/dairify/internal/service"
)

// UserHandler 인터페이스는 사용자 관련 핸들러의 메서드를 정의합니다.
type UserHandler interface {
	SignupUser(w http.ResponseWriter, r *http.Request)
}

// userHandler 구조체는 UserHandler 인터페이스를 구현합니다.
type userHandler struct {
	userService service.UserService
}

// NewUserHandler 함수는 UserHandler 인터페이스의 구현체를 반환합니다.
func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

/*
SignupUser 함수는 새로운 사용자를 등록하는 핸들러입니다.
curl -X POST http://localhost:8080/api/v1/users/signup/ \
-H "Content-Type: application/json" \
-d '{"username":"testuser","nickname":"Test","password":"password123","email":"test@example.com"}'
*/
func (h *userHandler) SignupUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// body로 Input 받기
	var inp dto.UserSignupDTO
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		response.Error(w, http.StatusBadRequest, "Bad Request: "+err.Error())
		return
	}

	// Service 함수 호출
	signupID, status, err := h.userService.SignupUser(r.Context(), inp)
	if err != nil {
		response.Error(w, status, "Error: "+err.Error())
		return
	}

	res := dto.UserSignupResponseDTO{
		SignupID: signupID,
	}

	response.Success(w, status, "User signed up successfully", res)
}
