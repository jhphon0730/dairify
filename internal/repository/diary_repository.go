package repository

import (
	"context"
	"database/sql"
	"errors"
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

	query := "SELECT id, title, content, creator_id, category_id, created_at, updated_at FROM diaries WHERE creator_id = $1"
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
		if err := rows.Scan(&diary.ID, &diary.Title, &diary.Content, &diary.CreatorID, &diary.CategoryID, &diary.CreatedAt, &diary.UpdatedAt); err != nil {
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
	query := "SELECT id, title, content, creator_id, category_id, created_at, updated_at FROM diaries WHERE id = $1"
	if err := r.db.DB.QueryRowContext(ctx, query, diary.ID).Scan(&diary.ID, &diary.Title, &diary.Content, &diary.CreatorID, &diary.CategoryID, &diary.CreatedAt, &diary.UpdatedAt); err != nil {
		// 조회 실패 시에는 id가 이상한 값이거나, 해당 일기가 존재하지 않는 경우
		if errors.Is(err, sql.ErrNoRows) {
			return apperror.ErrDiaryNotFound
		}
		return apperror.ErrDiaryGetInternal
	}
	return nil
}
