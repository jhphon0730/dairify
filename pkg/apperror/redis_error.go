package apperror

import "errors"

var (
	ErrUserRedisIsNil        = errors.New("세션 서버 연결에 실패했습니다")
	ErrUserRedisInvalidToken = errors.New("유효하지 않은 세션 토큰입니다")
)
