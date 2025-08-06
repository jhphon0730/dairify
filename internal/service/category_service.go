package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/jhphon0730/dairify/internal/dto"
	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/internal/repository"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

// CategoryService 인터페이스는 카테고리 관련 서비스의 메서드를 정의합니다.
type CategoryService interface {
	CreateCategory(ctx context.Context, createCategoryDTO dto.CreateCategoryDTO) (*model.Category, int, error)
	GetCategoriesByCreatorID(ctx context.Context, creatorID int64) ([]model.Category, int, error)
	UpdateCategoryName(ctx context.Context, updateCategoryDTO dto.UpdateCategoryDTO) (*model.Category, int, error)
	DeleteCategory(ctx context.Context, categoryID int64, creatorID int64) (int, error)
}

// categoryService 구조체는 CategoryService 인터페이스를 구현합니다.
type categoryService struct {
	categoryRepository repository.CategoryRepository
}

// NewCategoryService 함수는 CategoryService 인터페이스의 구현체를 반환합니다.
func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

// CreateCategory 함수는 새로운 카테고리를 생성합니다.
func (s *categoryService) CreateCategory(ctx context.Context, createCategoryDTO dto.CreateCategoryDTO) (*model.Category, int, error) {
	if err := createCategoryDTO.CheckIsValidInput(); err != nil {
		return nil, http.StatusBadRequest, err
	}

	category := &model.Category{
		Name:      createCategoryDTO.Name,
		CreatorID: createCategoryDTO.CreatorID,
	}

	if err := s.categoryRepository.CreateCategory(ctx, category); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return category, http.StatusCreated, nil
}

// GetCategoriesByCreatorID 함수는 주어진 생성자 ID로 카테고리 목록을 조회합니다.
func (s *categoryService) GetCategoriesByCreatorID(ctx context.Context, creatorID int64) ([]model.Category, int, error) {
	categories, err := s.categoryRepository.GetCategoriesByCreatorID(ctx, creatorID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return categories, http.StatusOK, nil
}

// UpdateCategoryName 함수는 카테고리 이름을 업데이트합니다.
func (s *categoryService) UpdateCategoryName(ctx context.Context, updateCategoryDTO dto.UpdateCategoryDTO) (*model.Category, int, error) {
	if err := updateCategoryDTO.CheckIsValidInput(); err != nil {
		return nil, http.StatusBadRequest, err
	}

	category, err := s.categoryRepository.GetCategoryByID(ctx, updateCategoryDTO.ID, updateCategoryDTO.CreatorID)
	if err != nil {
		if errors.Is(err, apperror.ErrCategoryNotFound) {
			return nil, http.StatusNotFound, apperror.ErrCategoryNotFound
		}
		return nil, http.StatusInternalServerError, err
	}

	// 카테고리의 생성자 ID가 요청한 사용자 ID와 일치하는지 확인
	if category.CreatorID != updateCategoryDTO.CreatorID {
		return nil, http.StatusForbidden, apperror.ErrCategoryUpdateForbidden
	}

	// 카테고리 이름 업데이트
	category.Name = updateCategoryDTO.Name

	if err := s.categoryRepository.UpdateCategoryName(ctx, category); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return category, http.StatusOK, nil
}

// DeleteCategory 함수는 카테고리를 삭제합니다.
func (s *categoryService) DeleteCategory(ctx context.Context, categoryID int64, creatorID int64) (int, error) {
	if categoryID <= 0 {
		return http.StatusBadRequest, apperror.ErrCategoryIDIsRequired
	}

	category, err := s.categoryRepository.GetCategoryByID(ctx, categoryID, creatorID)
	if err != nil {
		if errors.Is(err, apperror.ErrCategoryNotFound) {
			return http.StatusNotFound, apperror.ErrCategoryNotFound
		}
		return http.StatusInternalServerError, err
	}

	// 카테고리의 생성자 ID가 요청한 사용자 ID와 일치하는지 확인
	if category.CreatorID != creatorID {
		return http.StatusForbidden, apperror.ErrCategoryDeleteForbidden
	}

	if err := s.categoryRepository.DeleteCategory(ctx, categoryID, creatorID); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
