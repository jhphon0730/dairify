package repository

import "github.com/jhphon0730/dairify/internal/database"

// UserRepository 인터페이스는 사용자 관련 데이터베이스 작업을 정의합니다.
type UserRepository interface {
}

// userRepository 구조체는 UserRepository 인터페이스를 구현합니다.
type userRepository struct {
	db *database.DB
}

// NewUserRepository 함수는 UserRepository 인터페이스의 구현체를 반환합니다.
func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{db: db}
}
