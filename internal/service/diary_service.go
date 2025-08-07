package service

import (
	"context"
	"net/url"

	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/internal/repository"
)

// DiaryService는 일기 관련 비즈니스 로직을 처리하는 인터페이스입니다.
type DiaryService interface {
	GetDiariesByCreatorID(ctx context.Context, creatorID int64, params url.Values) ([]model.Diary, error)
}

// diaryService 구조체는 DiaryService 인터페이스를 구현합니다.
type diaryService struct {
	diaryRepository repository.DiaryRepository
}

// NewDiaryService 함수는 DiaryService 인터페이스의 구현체를 반환합니다.
func NewDiaryService(diaryRepository repository.DiaryRepository) DiaryService {
	return &diaryService{
		diaryRepository: diaryRepository,
	}
}

// GetDiariesByCreatorID 함수는 주어진 생성자 ID로 일기 목록을 조회합니다.
func (s *diaryService) GetDiariesByCreatorID(ctx context.Context, creatorID int64, params url.Values) ([]model.Diary, error) {
	return s.diaryRepository.GetDiariesByCreatorID(ctx, creatorID, params)
}
