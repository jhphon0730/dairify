package dto

import (
	"strings"

	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

// GetDiariesByCreatorIDResponseDTO 구조체는 일기 목록 조회 응답 DTO입니다.
type GetDiariesByCreatorIDResponseDTO struct {
	Diaries []model.Diary `json:"diaries"`
}

// CreateDiaryDTO 구조체는 신규 일기 생성 요청 DTO입니다.
type CreateDiaryDTO struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID *int64 `json:"category_id"`
}

// Validate 함수는 CreateDiaryDTO의 입력 유효성을 검사합니다.
func (dto *CreateDiaryDTO) Validate() error {
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
	return &model.Diary{
		Title:      dto.Title,
		Content:    dto.Content,
		CreatorID:  creatorID,
		CategoryID: dto.CategoryID,
	}
}

// CreateDiaryResponseDTO 구조체는 일기 생성 응답 DTO입니다.
type CreateDiaryResponseDTO struct {
	Diary *model.Diary `json:"diary"`
}

// GetDiaryByIDResponseDTO 구조체는 단일 일기 조회 응답 DTO입니다.
type GetDiaryByIDResponseDTO struct {
	Diary *model.Diary `json:"diary"`
}

// UpdateDiaryDTO 구조체는 일기 수정 요청 DTO입니다.
type UpdateDiaryDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Validate 함수는 UpdateDiaryDTO의 입력 유효성을 검사합니다.
func (dto *UpdateDiaryDTO) Validate() error {
	if strings.TrimSpace(dto.Title) == "" {
		return apperror.ErrDiaryUpdateTitleRequired
	}
	if strings.TrimSpace(dto.Content) == "" {
		return apperror.ErrDiaryUpdateContentRequired
	}
	return nil
}

// ToModel 함수는 UpdateDiaryDTO를 model.Diary로 변환합니다.
func (dto *UpdateDiaryDTO) ToModel() *model.Diary {
	return &model.Diary{
		Title:   dto.Title,
		Content: dto.Content,
	}
}

// UpdateDiaryResponseDTO 구조체는 일기 수정 응답 DTO입니다.
type UpdateDiaryResponseDTO struct {
	Diary *model.Diary `json:"diary"`
}
