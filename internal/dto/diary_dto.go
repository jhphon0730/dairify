package dto

import (
	"strings"

	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

// GetDiariesByCreatorIDResponseDTO 구조체는 일기 목록 조회 응답을 위한 DTO입니다.
type GetDiariesByCreatorIDResponseDTO struct {
	Diaries []model.Diary `json:"diaries"`
}

// CreateDiaryDTO 구조체는 새로운 일기를 생성하기 위한 DTO입니다.
type CreateDiaryDTO struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID *int64 `json:"category_id"`
}

// CreateDiaryDTO 구조체는 새로운 일기를 생성하기 위한 DTO입니다.
func (dto *CreateDiaryDTO) CheckIsValidInput() error {
	if strings.TrimSpace(dto.Title) == "" {
		return apperror.ErrDiaryCreateTitleRequired
	}
	if strings.TrimSpace(dto.Content) == "" {
		return apperror.ErrDiaryCreateContentRequired
	}
	return nil
}

// ToModel 함수는 CreateDiaryDTO를 model.Diary로 변환합니다.
func (dto *CreateDiaryDTO) ToModel(creatorID int64) *model.Diary {
	d := &model.Diary{
		Title:      dto.Title,
		Content:    dto.Content,
		CreatorID:  creatorID,
		CategoryID: dto.CategoryID,
	}

	return d
}

type CreateDiaryResponseDTO struct {
	Diary *model.Diary `json:"diary"`
}

// GetDiaryByIDResponseDTO 구조체는 단일 일기 조회 응답을 위한 DTO입니다.
type GetDiaryByIDResponseDTO struct {
	Diary *model.Diary `json:"diary"`
}
