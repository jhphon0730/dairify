package repository

import (
	"github.com/jhphon0730/dairify/internal/database"
	"github.com/jhphon0730/dairify/internal/dto"
)

// UserRepository 인터페이스는 사용자 관련 데이터베이스 작업을 정의합니다.
type UserRepository interface {
	CreateUser(userSignupDTO dto.UserSignupDTO) (int64, error)
}

// userRepository 구조체는 UserRepository 인터페이스를 구현합니다.
type userRepository struct {
	db *database.DB
}

// NewUserRepository 함수는 UserRepository 인터페이스의 구현체를 반환합니다.
func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser 함수는 새로운 사용자를 데이터베이스에 추가합니다.
func (r *userRepository) CreateUser(userSignupDTO dto.UserSignupDTO) (int64, error) {
	var id int64
	query := `
		INSERT INTO users (username, nickname, password_hash, email)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	if err := r.db.DB.QueryRow(query, userSignupDTO.Username, userSignupDTO.Nickname, userSignupDTO.Password, userSignupDTO.Email).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
