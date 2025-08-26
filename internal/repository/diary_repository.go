package repository

import (
	"context"
	"database/sql"
	"errors"
	"mime/multipart"
	"net/url"

	"github.com/jhphon0730/dairify/internal/database"
	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/pkg/apperror"
	"github.com/jhphon0730/dairify/pkg/utils"
)

// DiaryRepository는 일기 관련 데이터베이스 작업을 처리하는 인터페이스입니다.
type DiaryRepository interface {
	GetDiaryByID(ctx context.Context, diary *model.Diary) error
	GetDiariesByCreatorID(ctx context.Context, creatorID int64, params url.Values) ([]model.Diary, error)
	CreateDiary(ctx context.Context, diary *model.Diary) error
	DeleteDiary(ctx context.Context, diaryID int64, creatorID int64) error
	UpdateDiary(ctx context.Context, diary *model.Diary) error
	UploadDiaryImage(ctx context.Context, file []*multipart.FileHeader, diaryID int64) ([]*model.DiaryImage, error)
}

// diaryRepository 구조체는 DiaryRepository 인터페이스를 구현합니다.
type diaryRepository struct {
	db *database.DB
}

// NewDiaryRepository 함수는 DiaryRepository 인터페이스의 구현체를 반환합니다.
func NewDiaryRepository(db *database.DB) DiaryRepository {
	return &diaryRepository{
		db: db,
	}
}

// GetDiariesByCreatorID 함수는 주어진 생성자 ID로 일기 목록을 조회합니다.
func (r *diaryRepository) GetDiariesByCreatorID(ctx context.Context, creatorID int64, params url.Values) ([]model.Diary, error) {
	var diaries []model.Diary

	// 소프트 삭제된 레코드는 제외
	query := "SELECT id, title, content, creator_id, category_id, created_at, updated_at, is_deleted, deleted_at FROM diaries WHERE creator_id = $1 AND is_deleted = FALSE"
	args := []interface{}{creatorID}
	argIdx := 2 // $2부터 시작

	// 카테고리 필터링 추가
	if v := params.Get("category_id"); v != "" {
		query += " AND category_id = $" + utils.InterfaceToString(argIdx)
		args = append(args, v)
		argIdx++
	}

	// 제목 필터링 추가 (LIKE 검색)
	if v := params.Get("title"); v != "" {
		// 부분 일치 검색을 위해 %%를 양쪽에 붙임
		query += " AND title LIKE $" + utils.InterfaceToString(argIdx)
		args = append(args, "%"+v+"%")
		argIdx++
	}

	// 정렬 조건 추가
	query += " ORDER BY created_at DESC"

	rows, err := r.db.DB.QueryContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var diary model.Diary
		if err := rows.Scan(&diary.ID, &diary.Title, &diary.Content, &diary.CreatorID, &diary.CategoryID, &diary.CreatedAt, &diary.UpdatedAt, &diary.IsDeleted, &diary.DeletedAt); err != nil {
			return nil, err
		}
		diaries = append(diaries, diary)
	}

	// 결과 반환
	return diaries, nil
}

// CreateDiary 함수는 새로운 일기를 생성합니다.
func (r *diaryRepository) CreateDiary(ctx context.Context, diary *model.Diary) error {
	query := "INSERT INTO diaries (title, content, creator_id, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.DB.QueryRowContext(ctx, query, diary.Title, diary.Content, diary.CreatorID, diary.CategoryID).Scan(&diary.ID)
	if err != nil {
		return apperror.ErrDiaryCreateInternal
	}
	return nil
}

// GetDiaryByID 함수는 ID로 일기를 조회합니다.
func (r *diaryRepository) GetDiaryByID(ctx context.Context, diary *model.Diary) error {
	query := "SELECT id, title, content, creator_id, category_id, created_at, updated_at, is_deleted, deleted_at FROM diaries WHERE id = $1 AND is_deleted = FALSE"
	if err := r.db.DB.QueryRowContext(ctx, query, diary.ID).Scan(&diary.ID, &diary.Title, &diary.Content, &diary.CreatorID, &diary.CategoryID, &diary.CreatedAt, &diary.UpdatedAt, &diary.IsDeleted, &diary.DeletedAt); err != nil {
		// 조회 실패 시에는 id가 이상한 값이거나, 해당 일기가 존재하지 않는 경우
		if errors.Is(err, sql.ErrNoRows) {
			return apperror.ErrDiaryNotFound
		}
		return apperror.ErrDiaryGetInternal
	}
	return nil
}

// DeleteDiary 함수는 일기를 소프트 삭제 처리합니다.
func (r *diaryRepository) DeleteDiary(ctx context.Context, diaryID int64, creatorID int64) error {
	// 작성자 조건을 추가하여 다른 사용자의 일기 삭제 방지
	query := "UPDATE diaries SET is_deleted = TRUE, deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND creator_id = $2 AND is_deleted = FALSE"
	res, err := r.db.DB.ExecContext(ctx, query, diaryID, creatorID)
	if err != nil {
		return apperror.ErrDiaryDeleteInternal
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return apperror.ErrDiaryDeleteInternal
	}
	if rows == 0 {
		return apperror.ErrDiaryNotFound
	}
	return nil
}

// UpdateDiary 함수는 일기를 업데이트합니다.
func (r *diaryRepository) UpdateDiary(ctx context.Context, diary *model.Diary) error {
	query := "UPDATE diaries SET title = $1, content = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3 AND is_deleted = FALSE"
	res, err := r.db.DB.ExecContext(ctx, query, diary.Title, diary.Content, diary.ID)
	if err != nil {
		return apperror.ErrDiaryUpdateInternal
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return apperror.ErrDiaryUpdateInternal
	}
	if rows == 0 {
		return apperror.ErrDiaryNotFound
	}
	return nil
}

// UploadDiaryImage 함수는 다이어리 이미지를 업로드하고 저장된 경로를 반환합니다.
func (r *diaryRepository) UploadDiaryImage(ctx context.Context, files []*multipart.FileHeader, diaryID int64) ([]*model.DiaryImage, error) {
	var diaryImages []*model.DiaryImage

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, apperror.ErrDiaryImageUploadInternal
	}
	defer func() {
		_ = tx.Rollback()
	}()

	query := "INSERT INTO images (diary_id, file_path, file_name, content_type, file_size) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	for _, file := range files {
		// 이미지 업로드를 수행하고 결과 객체를 반환받음
		diaryImage, err := utils.UploadDiaryImage(file, diaryID)
		if err != nil {
			utils.RemoveDiaryImages(diaryImages)
			return nil, err
		}
		// 업로드된 이미지 정보를 슬라이스에 추가
		diaryImages = append(diaryImages, diaryImage)

		// RETURNING id를 사용하므로 QueryRowContext로 id를 스캔
		if err := tx.QueryRowContext(ctx, query, diaryID, diaryImage.FilePath, diaryImage.FileName, diaryImage.ContentType, diaryImage.FileSize).Scan(&diaryImage.ID); err != nil {
			utils.RemoveDiaryImages(diaryImages)
			return nil, apperror.ErrDiaryImageUploadInternal
		}
	}

	return diaryImages, tx.Commit()
}
