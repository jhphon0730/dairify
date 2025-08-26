package apperror

import "errors"

var (
	ErrDiaryGetInternal         = errors.New("서버 내부 오류로 일기 조회에 실패했습니다")
	ErrDiaryCreateInternal      = errors.New("서버 내부 오류로 일기 생성에 실패했습니다")
	ErrDiaryDeleteInternal      = errors.New("서버 내부 오류로 일기 삭제에 실패했습니다")
	ErrDiaryUpdateInternal      = errors.New("서버 내부 오류로 일기 수정에 실패했습니다")
	ErrDiaryImageUploadInternal = errors.New("서버 내부 오류로 일기 이미지 업로드에 실패했습니다")

	ErrDiaryCreateTitleRequired   = errors.New("일기 제목은 필수 입력값입니다")
	ErrDiaryCreateContentRequired = errors.New("일기 내용은 필수 입력값입니다")
	ErrDiaryUpdateTitleRequired   = errors.New("일기 제목은 필수 입력값입니다")
	ErrDiaryUpdateContentRequired = errors.New("일기 내용은 필수 입력값입니다")

	ErrDiaryNotFound          = errors.New("해당 일기를 찾을 수 없습니다")
	ErrDiaryNotFoundOrDeleted = errors.New("해당 일기가 존재하지 않거나 삭제되었습니다")
	ErrDiaryUpdateForbidden   = errors.New("해당 일기를 수정할 권한이 없습니다")
)
