package dto

import (
	"strings"

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
		return apperror.Err_USER_SIGNUP_USERNAME_REQUIRED
	}

	if strings.TrimSpace(d.Nickname) == "" {
		return apperror.Err_USER_SIGNUP_NICKNAME_REQUIRED
	}

	if strings.TrimSpace(d.Password) == "" {
		return apperror.Err_USER_SIGNUP_PASSWORD_REQUIRED
	}

	if strings.TrimSpace(d.Email) == "" {
		return apperror.Err_USER_SIGNUP_EMAIL_REQUIRED
	}

	return nil
}

type UserSignupResponseDTO struct {
	SignupID int64 `json:"signup_id"`
}
