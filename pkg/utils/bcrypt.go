package utils

import (
	"github.com/jhphon0730/dairify/internal/config"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	cfg := config.GetConfig()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), InterfaceToInt(cfg.BCRYPT_COST))
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
