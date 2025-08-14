package service

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/jhphon0730/dairify/internal/dto"
	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/internal/repository"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

// DiaryService는 일기 관련 비즈니스 로직을 처리하는 인터페이스입니다.
type DiaryService interface {
	GetDiaryByID(ctx context.Context, diaryID int64) (*model.Diary, int, error)
	GetDiariesByCreatorID(ctx context.Context, creatorID int64, params url.Values) ([]model.Diary, int, error)
	CreateDiary(ctx context.Context, diary dto.CreateDiaryDTO, creatorID int64) (*model.Diary, int, error)
	DeleteDiary(ctx context.Context, diaryID int64, creatorID int64) (int, error)
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
func (s *diaryService) GetDiariesByCreatorID(ctx context.Context, creatorID int64, params url.Values) ([]model.Diary, int, error) {
	diaries, err := s.diaryRepository.GetDiariesByCreatorID(ctx, creatorID, params)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return diaries, http.StatusOK, nil
}

// CreateDiary 함수는 새로운 일기를 생성합니다.
func (s *diaryService) CreateDiary(ctx context.Context, diary dto.CreateDiaryDTO, creatorID int64) (*model.Diary, int, error) {
	if err := diary.CheckIsValidInput(); err != nil {
		return nil, http.StatusBadRequest, err
	}

	diaryModel := diary.ToModel(creatorID)
	if err := s.diaryRepository.CreateDiary(ctx, diaryModel); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return diaryModel, http.StatusCreated, nil
}

// GetDiaryByID 함수는 ID로 일기를 조회합니다.
func (s *diaryService) GetDiaryByID(ctx context.Context, diaryID int64) (*model.Diary, int, error) {
	// 조회 대상 일기 모델 생성
	diary := &model.Diary{ID: diaryID}

	// 레포지토리 호출 후 오류 분기
	if err := s.diaryRepository.GetDiaryByID(ctx, diary); err != nil {
		if errors.Is(err, apperror.ErrDiaryNotFound) {
			// 존재하지 않는 일기
			return nil, http.StatusNotFound, err
		}
		// 내부 서버 오류 - 명시적 에러 반환
		return nil, http.StatusInternalServerError, apperror.ErrDiaryGetInternal
	}

	return diary, http.StatusOK, nil
}

// DeleteDiary 함수는 일기를 소프트 삭제합니다.
func (s *diaryService) DeleteDiary(ctx context.Context, diaryID int64, creatorID int64) (int, error) {
	if diaryID <= 0 {
		return http.StatusBadRequest, apperror.ErrDiaryNotFound
	}
	if err := s.diaryRepository.DeleteDiary(ctx, diaryID, creatorID); err != nil {
		if errors.Is(err, apperror.ErrDiaryNotFound) {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, apperror.ErrDiaryDeleteInternal
	}
	return http.StatusOK, nil
}
