package middleware

import (
	"log"
	"net/http"
)

// loggingResponseWriter는 http.ResponseWriter를 래핑하여 상태 코드를 기록합니다.
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Write 메서드는 응답 본문을 작성하고 상태 코드를 기록합니다.
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware는 HTTP 요청과 응답을 로깅하는 미들웨어입니다.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 요청 정보 로깅
		log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		// ResponseWriter 래핑
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// 다음 핸들러 호출
		next(lrw, r)

		// 응답 정보 로깅
		log.Printf("Responded to request: %s %s with status code %d", r.Method, r.URL.Path, lrw.statusCode)
	}
}
