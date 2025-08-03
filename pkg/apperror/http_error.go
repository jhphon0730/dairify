package apperror

import "errors"

var (
	Err_HTTP_METHOD_NOT_ALLOWED = errors.New("요청 메서드가 허용되지 않습니다")
)
