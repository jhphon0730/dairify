package repository

import (
	"context"
	"database/sql"

	"github.com/jhphon0730/dairify/internal/database"
	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

// CategoryRepository 인터페이스는 카테고리 관련 데이터베이스 작업을 정의합니다.
type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *model.Category) error
	GetCategoryByNameAndCreatorID(ctx context.Context, name string, creatorID int64) (*model.Category, error)
}

// categoryRepository 구조체는 CategoryRepository 인터페이스를 구현합니다.
type categoryRepository struct {
	db *database.DB
}

// NewCategoryRepository 함수는 CategoryRepository 인터페이스의 구현체를 반환합니다.
func NewCategoryRepository(db *database.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

// CreateCategory 함수는 새로운 카테고리를 데이터베이스에 추가합니다.
func (r *categoryRepository) CreateCategory(ctx context.Context, category *model.Category) error {
	query := `
		INSERT INTO categories (name, creator_id)
		VALUES ($1, $2)
		RETURNING id
	`

	if err := r.db.DB.QueryRowContext(ctx, query, category.Name, category.CreatorID).Scan(&category.ID); err != nil {
		return apperror.ErrCreateFailedInternalServerError
	}

	return nil
}

// GetCategoryByNameAndCreatorID 함수는 주어진 이름과 생성자 ID로 카테고리를 조회합니다.
func (r *categoryRepository) GetCategoryByNameAndCreatorID(ctx context.Context, name string, creatorID int64) (*model.Category, error) {
	query := `
		SELECT id, name, creator_id, created_at
		FROM categories
		WHERE name = $1 AND creator_id = $2
	`

	row := r.db.DB.QueryRowContext(ctx, query, name, creatorID)

	var category model.Category
	if err := row.Scan(&category.ID, &category.Name, &category.CreatorID, &category.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.ErrCategoryNotFound
		}
		return nil, apperror.ErrGetFailedInternalServerError
	}

	return &category, nil
}
