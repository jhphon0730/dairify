package utils

import (
	"github.com/jhphon0730/dairify/internal/config"
	"golang.org/x/crypto/bcrypt"
)

// GenerateHash 함수는 주어진 비밀번호를 해시화하여 반환합니다.
func GenerateHashPassword(password string) (string, error) {
	cfg := config.GetConfig()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), InterfaceToInt(cfg.BCRYPT_COST))
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// CompareHashAndPassword 함수는 주어진 비밀번호와 해시된 비밀번호를 비교합니다.
func CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
