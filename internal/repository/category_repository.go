package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jhphon0730/dairify/internal/database"
	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

// CategoryRepository 인터페이스는 카테고리 관련 데이터베이스 작업을 정의합니다.
type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *model.Category) error
	GetCategoriesByCreatorID(ctx context.Context, creatorID int64) ([]model.Category, error)
	GetCategoryByID(ctx context.Context, id int64, creatorID int64) (*model.Category, error)
	UpdateCategoryName(ctx context.Context, category *model.Category) error
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
		ON CONFLICT (name, creator_id) DO NOTHING
		RETURNING id
	`

	err := r.db.DB.QueryRowContext(ctx, query, category.Name, category.CreatorID).Scan(&category.ID)
	if err != nil {
		// 만약 카테고리가 이미 존재한다면, 에러를 무시하고 nil을 반환합니다.
		if errors.Is(err, sql.ErrNoRows) {
			return apperror.ErrCategoryCreateDuplicateName
		}
		return err
	}
	return nil
}

// GetCategoriesByCreatorID 함수는 주어진 생성자 ID로 카테고리 목록을 조회합니다.
func (r *categoryRepository) GetCategoriesByCreatorID(ctx context.Context, creatorID int64) ([]model.Category, error) {
	var categories []model.Category
	query := `
		SELECT id, name, creator_id, created_at
		FROM categories
		WHERE creator_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.DB.QueryContext(ctx, query, creatorID)
	// 만약 카테고리가 존재하지 않는다면, 빈 슬라이스를 반환합니다.
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, apperror.ErrGetFailedInternalServerError
	}
	defer rows.Close()

	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatorID, &category.CreatedAt); err != nil {
			return nil, apperror.ErrGetFailedInternalServerError
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// FindCategoryByID 함수는 주어진 ID와 생성자 ID로 카테고리를 조회합니다.
func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int64, creatorID int64) (*model.Category, error) {
	query := `
		SELECT id, name, creator_id, created_at
		FROM categories
		WHERE id = $1 AND creator_id = $2
	`

	var category model.Category
	err := r.db.DB.QueryRowContext(ctx, query, id, creatorID).Scan(&category.ID, &category.Name, &category.CreatorID, &category.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrCategoryNotFound
		}
		return nil, apperror.ErrGetFailedInternalServerError
	}

	return &category, nil
}

// UpdateCategoryName 함수는 카테고리의 이름을 업데이트합니다.
func (r *categoryRepository) UpdateCategoryName(ctx context.Context, category *model.Category) error {
	query := `
		UPDATE categories
		SET name = $1
		WHERE id = $2 AND creator_id = $3
	`

	result, err := r.db.DB.ExecContext(ctx, query, category.Name, category.ID, category.CreatorID)
	if err != nil {
		return apperror.ErrUpdateFailedInternalServerError
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return apperror.ErrUpdateFailedInternalServerError
	}

	if rowsAffected == 0 {
		return apperror.ErrCategoryNotFound
	}

	return nil
}
