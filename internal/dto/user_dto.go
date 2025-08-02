package dto

import (
	"errors"
	"strings"
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
		return errors.New("username is required")
	}

	if strings.TrimSpace(d.Nickname) == "" {
		return errors.New("nickname is required")
	}

	if strings.TrimSpace(d.Password) == "" {
		return errors.New("password is required")
	}

	if strings.TrimSpace(d.Email) == "" {
		return errors.New("email is required")
	}

	return nil
}

type UserSignupResponseDTO struct {
	SignupID int64 `json:"signup_id"`
}
