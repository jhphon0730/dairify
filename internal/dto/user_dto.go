package dto

import (
	"strings"

	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

// UserSignupDTO 구조체는 사용자 등록을 위한 데이터 전송 객체입니다.
type UserSignupDTO struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// CheckIsValidInput 함수는 입력 값을 확인해주는 함수입니다.
func (d *UserSignupDTO) CheckIsValidInput() error {
	if strings.TrimSpace(d.Username) == "" {
		return apperror.ErrUserSignupUserNameRequired
	}

	if strings.TrimSpace(d.Nickname) == "" {
		return apperror.ErrUserSignupNicknameRequired
	}

	if strings.TrimSpace(d.Password) == "" {
		return apperror.ErrUserSignupPasswordRequired
	}

	if strings.TrimSpace(d.Email) == "" {
		return apperror.ErrUserSignupEmailRequired
	}

	return nil
}

// UserSignupResponseDTO 구조체는 사용자 등록 응답을 위한 데이터 전송 객체입니다.
type UserSignupResponseDTO struct {
	SignupID int64 `json:"signup_id"`
}

// UserSigninDTO 구조체는 사용자 로그인을 위한 데이터 전송 객체입니다.
type UserSigninDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckIsValidInput 함수는 사용자 로그인 입력 값을 확인해주는 함수입니다.
func (d *UserSigninDTO) CheckIsValidInput() error {
	if strings.TrimSpace(d.Username) == "" {
		return apperror.ErrUserSigninInvalidUserName
	}

	if strings.TrimSpace(d.Password) == "" {
		return apperror.ErrUserSigninInvalidPassword
	}

	return nil
}

// UserSigninResponseDTO 구조체는 사용자 로그인 응답을 위한 데이터 전송 객체입니다.
type UserSigninResponseDTO struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         *model.User `json:"user"`
}
