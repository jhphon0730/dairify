package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	"github.com/jhphon0730/dairify/internal/database"
	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/pkg/utils"
)

// DiaryRepository는 일기 관련 데이터베이스 작업을 처리하는 인터페이스입니다.
type DiaryRepository interface {
	GetDiariesByCreatorID(ctx context.Context, creatorID int64, params url.Values) ([]model.Diary, error)
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

	query := "SELECT id, title, content, creator_id, created_at, updated_at FROM diaries WHERE creator_id = $1"
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

	fmt.Println("Executing query:", query, "with args:", args)

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
		if err := rows.Scan(&diary.ID, &diary.Title, &diary.Content, &diary.CreatorID, &diary.CreatedAt, &diary.UpdatedAt); err != nil {
			return nil, err
		}
		diaries = append(diaries, diary)
	}

	// 결과 반환
	return diaries, nil
}
