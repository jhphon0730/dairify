package apperror

import "errors"

var (
	ErrDiaryCreateTitleRequired   = errors.New("일기 제목은 필수입니다")
	ErrDiaryCreateContentRequired = errors.New("일기 내용은 필수입니다")
)
