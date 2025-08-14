package apperror

import "errors"

var (
	ErrDiaryGetInternal    = errors.New("서버 내부 오류로 일기 조회에 실패했습니다")
	ErrDiaryCreateInternal = errors.New("서버 내부 오류로 일기 생성에 실패했습니다")

	ErrDiaryCreateTitleRequired   = errors.New("일기 제목은 필수 입력값입니다")
	ErrDiaryCreateContentRequired = errors.New("일기 내용은 필수 입력값입니다")

	ErrDiaryNotFound = errors.New("해당 일기를 찾을 수 없습니다")
)
