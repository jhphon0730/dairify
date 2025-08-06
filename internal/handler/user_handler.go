package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jhphon0730/dairify/internal/dto"
	"github.com/jhphon0730/dairify/internal/middleware"
	"github.com/jhphon0730/dairify/internal/response"
	"github.com/jhphon0730/dairify/internal/service"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

// UserHandler 인터페이스는 사용자 관련 핸들러의 메서드를 정의합니다.
type UserHandler interface {
	SignupUser(w http.ResponseWriter, r *http.Request)
	SigninUser(w http.ResponseWriter, r *http.Request)
	SignoutUser(w http.ResponseWriter, r *http.Request)
	ProfileUser(w http.ResponseWriter, r *http.Request)
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

/* SignupUser 함수는 새로운 사용자를 등록하는 핸들러입니다. */
func (h *userHandler) SignupUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	// body로 Input 받기
	var inp dto.UserSignupDTO
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil && err.Error() != "EOF" {
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

/* SigninUser 함수는 사용자를 로그인하는 핸들러입니다. */
func (h *userHandler) SigninUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	// body로 Input 받기
	var inp dto.UserSigninDTO
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil && err.Error() != "EOF" {
		response.Error(w, http.StatusBadRequest, "Bad Request: "+err.Error())
		return
	}

	// Service 함수 호출
	accessToken, refreshToken, status, err := h.userService.SigninUser(r.Context(), inp)
	if err != nil {
		response.Error(w, status, "Error: "+err.Error())
		return
	}

	signinResponse := dto.UserSigninResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	response.Success(w, status, "User signed in successfully", signinResponse)
}

/* SignoutUser 함수는 사용자를 로그아웃하는 핸들러입니다. */
func (h *userHandler) SignoutUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, apperror.ErrAuthUnauthorized.Error())
		return
	}

	status, err := h.userService.SignoutUser(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Error: "+err.Error())
		return
	}

	response.Success(w, status, "User signed out successfully", nil)
}

/* ProfileUser 함수는 사용자의 프로필 정보를 조회하는 핸들러입니다. */
func (h *userHandler) ProfileUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, apperror.ErrAuthUnauthorized.Error())
		return
	}

	user, status, err := h.userService.Profile(r.Context(), userID)
	if err != nil {
		response.Error(w, status, "Error: "+err.Error())
		return
	}

	res := dto.UserProfileResponseDTO{
		User: user,
	}
	response.Success(w, status, "User profile retrieved successfully", res)
}
