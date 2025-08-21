package dto

import (
	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

// CreateCategoryDTO 구조체는 카테고리 생성을 위한 데이터 전송 객체입니다.
type CreateCategoryDTO struct {
	Name      string `json:"name" binding:"required"`
	CreatorID int64  `json:"creator_id"`
}

// Validate 함수는 카테고리 생성 입력 값을 확인해주는 함수입니다.
func (d *CreateCategoryDTO) Validate() error {
	if d.Name == "" {
		return apperror.ErrCategoryNameRequired
	}
	if d.CreatorID <= 0 {
		return apperror.ErrCategoryCreatorIDRequired
	}

	return nil
}

// ToModel 함수는 CreateCategoryDTO를 model.Category로 변환합니다.
func (d *CreateCategoryDTO) ToModel() *model.Category {
	return &model.Category{
		Name:      d.Name,
		CreatorID: d.CreatorID,
	}
}

// CreateCategoryResponse 구조체는 카테고리 생성 응답을 위한 데이터 전송 객체입니다.
type CreateCategoryResponseDTO struct {
	Category *model.Category `json:"category"`
}

// GetCategoriesResponseDTO 구조체는 카테고리 목록 조회 응답을 위한 데이터 전송 객체입니다.
type GetCategoriesResponseDTO struct {
	Categories []model.Category `json:"categories"`
}

// UpdateCategoryDTO 구조체는 카테고리 이름 업데이트를 위한 데이터 전송 객체입니다.
type UpdateCategoryDTO struct {
	ID        int64  `json:"id"`
	CreatorID int64  `json:"creator_id"`
	Name      string `json:"name" binding:"required"`
}

// Validate 함수는 카테고리 이름 업데이트 입력 값을 확인해주는 함수입니다.
func (d *UpdateCategoryDTO) Validate() error {
	if d.Name == "" {
		return apperror.ErrCategoryNameRequired
	}
	return nil
}

// UpdateCategoryResponseDTO 구조체는 카테고리 이름 업데이트 응답을 위한 데이터 전송 객체입니다.
type UpdateCategoryResponseDTO struct {
	Category *model.Category `json:"category"`
}
