package dto

import "github.com/jhphon0730/dairify/internal/model"

// GetDiariesByCreatorIDResponseDTO 구조체는 일기 목록 조회 응답을 위한 DTO입니다.
type GetDiariesByCreatorIDResponseDTO struct {
	Diaries []model.Diary `json:"diaries"`
}
