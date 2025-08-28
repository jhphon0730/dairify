package server

import (
	"context"
	"net/http"
	"time"

	"github.com/jhphon0730/dairify/internal/database"
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
}

// NewServer는 새로운 Server 인스턴스를 생성합니다.
func NewServer(PORT string, db *database.DB) Server {
	mux := http.NewServeMux()
	SetupRoutes(mux, db)

	return &server{
		httpServer: &http.Server{
			Addr:         ":" + PORT,
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  30 * time.Second,
		},
	}
}

// settingMediaDir는 미디어 디렉토리를 설정합니다.
func (s *server) settingMediaDir() {
	fs := http.FileServer(http.Dir("./media"))
	s.httpServer.Handler.(*http.ServeMux).Handle("/media/", http.StripPrefix("/media/", fs))
}

// RunServer는 서버를 시작합니다.
func (s *server) RunServer() error {
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
