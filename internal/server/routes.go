package server

import (
	"net/http"

	"github.com/jhphon0730/dairify/internal/database"
)

// SetupRoutes는 HTTP 라우트를 설정합니다.
func SetupRoutes(mux *http.ServeMux, db *database.DB) {
	// HTTP 연결 상태 확인 라우트 설정
	RegisterHealthRoutes(mux)
}

// RegisterHealthRoutes는 헬스 체크 라우트를 등록합니다.
func RegisterHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is healthy"))
	})
}
