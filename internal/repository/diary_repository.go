package repository

import "github.com/jhphon0730/dairify/internal/database"

// DiaryRepository는 일기 관련 데이터베이스 작업을 처리하는 인터페이스입니다.
type DiaryRepository interface {
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
