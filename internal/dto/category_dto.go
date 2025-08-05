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

// CheckIsValidInput 함수는 카테고리 생성 입력 값을 확인해주는 함수입니다.
func (d *CreateCategoryDTO) CheckIsValidInput() error {
	if d.Name == "" {
		return apperror.ErrCategoryNameRequired
	}
	if d.CreatorID <= 0 {
		return apperror.ErrCategoryCreatorIDRequired
	}

	return nil
}

// CreateCategoryResponse 구조체는 카테고리 생성 응답을 위한 데이터 전송 객체입니다.
type CreateCategoryResponseDTO struct {
	Category *model.Category `json:"category"`
}
