package dto

import (
	"errors"
	"strings"
)

type UserSignupDTO struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

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
