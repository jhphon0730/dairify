package server

import (
	"context"
	"net/http"
	"time"

	"github.com/jhphon0730/dairify/internal/database"
)

// 상수는 파일 상단에 모아 선언합니다.
const (
	// 서버 타임아웃 설정(초)
	READ_TIMEOUT_SEC  = 10
	WRITE_TIMEOUT_SEC = 10
	IDLE_TIMEOUT_SEC  = 30

	// CORS 설정
	CORS_ALLOW_ORIGIN      = "*" // 예: "https://example.com"
	CORS_ALLOW_METHODS     = "GET,POST,PUT,PATCH,DELETE,OPTIONS"
	CORS_ALLOW_HEADERS     = "Content-Type,Authorization"
	CORS_EXPOSE_HEADERS    = "" // 노출할 헤더가 없으면 빈 문자열 유지
	CORS_ALLOW_CREDENTIALS = false
	CORS_MAX_AGE           = "86400" // 24시간

	// 미디어 경로 설정
	MEDIA_DIR          = "./media"
	MEDIA_ROUTE_PREFIX = "/media/"
)

// Server 인터페이스는 서버의 기본 동작을 정의합니다.
type Server interface {
	settingMediaDir()

	RunServer() error
	Shutdown(ctx context.Context) error
}

// 이 구조체는 Server 인터페이스를 구현합니다.
type server struct {
	httpServer *http.Server
	mux        *http.ServeMux
}

// NewServer는 새로운 Server 인스턴스를 생성합니다.
func NewServer(port string, db *database.DB) Server {
	// 라우터 생성
	mux := http.NewServeMux()

	// 라우팅 설정
	SetupRoutes(mux, db)

	// CORS 미들웨어로 핸들러 래핑
	handler := withCors(mux)

	// http.Server 생성
	s := &server{
		httpServer: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  time.Duration(READ_TIMEOUT_SEC) * time.Second,
			WriteTimeout: time.Duration(WRITE_TIMEOUT_SEC) * time.Second,
			IdleTimeout:  time.Duration(IDLE_TIMEOUT_SEC) * time.Second,
		},
		mux: mux,
	}

	return s
}

// withCors는 CORS 헤더를 추가하는 미들웨어를 반환합니다.
func withCors(next http.Handler) http.Handler {
	// CORS 처리는 미들웨어로 분리하여 단일 책임을 유지합니다.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 공통 CORS 헤더 설정
		w.Header().Set("Access-Control-Allow-Origin", CORS_ALLOW_ORIGIN)
		w.Header().Set("Access-Control-Allow-Methods", CORS_ALLOW_METHODS)
		w.Header().Set("Access-Control-Allow-Headers", CORS_ALLOW_HEADERS)
		w.Header().Set("Access-Control-Max-Age", CORS_MAX_AGE)

		// 자격 증명 허용이 필요한 경우에만 헤더 추가
		if CORS_ALLOW_CREDENTIALS {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// 노출할 헤더가 지정된 경우에만 헤더 추가
		if CORS_EXPOSE_HEADERS != "" {
			w.Header().Set("Access-Control-Expose-Headers", CORS_EXPOSE_HEADERS)
		}

		// Preflight 요청(OPTIONS)은 바로 응답 후 종료
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 다음 핸들러로 위임
		next.ServeHTTP(w, r)
	})
}

// settingMediaDir는 미디어 디렉토리를 설정합니다.
func (s *server) settingMediaDir() {
	// 정적 파일 서버 설정
	fs := http.FileServer(http.Dir(MEDIA_DIR))
	s.mux.Handle(MEDIA_ROUTE_PREFIX, http.StripPrefix(MEDIA_ROUTE_PREFIX, fs))
}

// RunServer는 서버를 시작합니다.
func (s *server) RunServer() error {
	// 서버 시작 전 미디어 경로 등록
	s.settingMediaDir()

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Shutdown은 서버를 안전하게 종료합니다.
func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
