package apperror

import "errors"

var (
	ErrAuthRequiredToken = errors.New("인증이 필요한 요청입니다")
	ErrAuthInvalidToken  = errors.New("유효하지 않은 토큰입니다")
)
