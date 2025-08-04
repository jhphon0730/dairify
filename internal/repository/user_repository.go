package repository

import (
	"context"
	"database/sql"

	"github.com/jhphon0730/dairify/internal/database"
	"github.com/jhphon0730/dairify/internal/dto"
	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/pkg/apperror"
	"github.com/lib/pq"
)

// UserRepository 인터페이스는 사용자 관련 데이터베이스 작업을 정의합니다.
type UserRepository interface {
	CreateUser(cxt context.Context, userSignupDTO dto.UserSignupDTO) (int64, error)
	FindUserByUsername(ctx context.Context, username string) (*model.User, error)
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
func (r *userRepository) CreateUser(ctx context.Context, userSignupDTO dto.UserSignupDTO) (int64, error) {
	var id int64
	query := `
		INSERT INTO users (username, nickname, password, email)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	if err := r.db.DB.QueryRowContext(ctx, query, userSignupDTO.Username, userSignupDTO.Nickname, userSignupDTO.Password, userSignupDTO.Email).Scan(&id); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			if pqErr.Constraint == "users_username_key" {
				return 0, apperror.ErrUserSignupDuplicateUserName
			}

			if pqErr.Constraint == "users_email_key" {
				return 0, apperror.ErrUserSignupDuplicateEmail
			}
		}

		return 0, err
	}

	return id, nil
}

// FindUserByUsername 함수는 사용자 이름으로 사용자를 검색합니다.
func (r *userRepository) FindUserByUsername(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	query := `
		SELECT id, username, nickname, password, email, created_at
		FROM users
		WHERE username = $1
	`

	// 사용자 정보를 조회합니다.
	if err := r.db.DB.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.Nickname, &user.Password, &user.Email, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}
