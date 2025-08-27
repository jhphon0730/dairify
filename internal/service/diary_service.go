package service

import (
	"context"
	"errors"
	"mime/multipart"
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
	UpdateDiary(ctx context.Context, updateDTO dto.UpdateDiaryDTO, diaryID int64, creatorID int64) (int, error)
	UploadDiaryImage(ctx context.Context, files []*multipart.FileHeader, diaryID int64, creatorID int64) ([]*model.DiaryImage, int, error)
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
	if err := diary.Validate(); err != nil {
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

	// 일기에 해당하는 이미지들도 함께 조회
	images, err := s.diaryRepository.GetImagesByDiaryID(ctx, diary.ID)
	if err != nil {
		diary.Images = nil
		if errors.Is(err, apperror.ErrDiaryImageNotFound) {
			return diary, http.StatusOK, nil
		}
		return nil, http.StatusInternalServerError, apperror.ErrDiaryGetInternal
	}
	diary.Images = images

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

func (s *diaryService) UpdateDiary(ctx context.Context, updateDTO dto.UpdateDiaryDTO, diaryID int64, creatorID int64) (int, error) {
	if err := updateDTO.Validate(); err != nil {
		return http.StatusBadRequest, err
	}

	// 일기 조회를 먼저 수행하여 업데이터 하려는 다이어리가 존재하는 지 확인
	diary := &model.Diary{ID: diaryID}
	err := s.diaryRepository.GetDiaryByID(ctx, diary)
	if err != nil {
		if errors.Is(err, apperror.ErrDiaryNotFound) {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, apperror.ErrDiaryGetInternal
	}

	// 다이어리와, 사용자의 권한을 확인합니다.
	if diary.CreatorID != creatorID {
		return http.StatusForbidden, apperror.ErrDiaryUpdateForbidden
	}

	diary = updateDTO.ToModel()
	diary.ID = diaryID

	if err := s.diaryRepository.UpdateDiary(ctx, diary); err != nil {
		return http.StatusInternalServerError, apperror.ErrDiaryUpdateInternal
	}

	return http.StatusOK, nil
}

// UploadDiaryImage 함수는 다이어리 이미지를 업로드하고 저장된 경로를 반환합니다.
func (s *diaryService) UploadDiaryImage(ctx context.Context, files []*multipart.FileHeader, diaryID int64, creatorID int64) ([]*model.DiaryImage, int, error) {
	diary := &model.Diary{ID: diaryID}
	err := s.diaryRepository.GetDiaryByID(ctx, diary)
	if err != nil {
		if errors.Is(err, apperror.ErrDiaryNotFound) {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, apperror.ErrDiaryGetInternal
	}

	// 다이어리와, 사용자의 권한을 확인합니다.
	if diary.CreatorID != creatorID {
		return nil, http.StatusForbidden, apperror.ErrDiaryUpdateForbidden
	}

	diaryImages, err := s.diaryRepository.UploadDiaryImage(ctx, files, diaryID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return diaryImages, http.StatusOK, nil
}
