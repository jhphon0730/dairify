package middleware

import (
	"net/http"
)

// ChainLoggingWithAuthMiddleware 함수는 로깅과 사용자 인증 미들웨어를 한 번에 적용합니다.
func ChainLoggingWithAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return LoggingMiddleware(AuthMiddleware(next))
}
