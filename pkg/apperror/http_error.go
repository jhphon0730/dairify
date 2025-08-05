package apperror

import "errors"

var (
	ErrHttpMethodNotAllowed = errors.New("요청 메서드가 허용되지 않습니다")
	ErrInternalServerError  = errors.New("내부 서버 오류가 발생했습니다")
)
