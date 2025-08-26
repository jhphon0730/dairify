package apperror

import "errors"

var (
	ErrImageInvalidContentType         = errors.New("지원하지 않는 이미지 형식입니다")
	ErrImageIsRequired                 = errors.New("업로드할 이미지가 필요합니다")
	ErrUploadFailedInternalServerError = errors.New("내부 서버 오류로 업로드에 실패했습니다")
	ErrImageInvalidSize                = errors.New("유효하지 않은 이미지 크기입니다")

	ErrEmptySlice = errors.New("빈 슬라이스입니다")
)
