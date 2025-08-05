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

	// 중복 확인
	existingCategory, err := s.categoryRepository.GetCategoryByNameAndCreatorID(ctx, createCategoryDTO.Name, createCategoryDTO.CreatorID)
	if err != nil && !errors.Is(err, apperror.ErrCategoryNotFound) {
		return nil, http.StatusInternalServerError, apperror.ErrCreateFailedInternalServerError
	}

	if existingCategory != nil {
		return nil, http.StatusConflict, apperror.ErrCategoryCreateDuplicateName
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
