package server

import (
	"context"
	"net/http"
	"time"
)

// Server 구조체는 서버의 설정과 상태를 관리합니다.
type Server struct {
	httpServer *http.Server
}

// NewServer는 새로운 Server 인스턴스를 생성합니다.
func NewServer(PORT string) *Server {
	mux := http.NewServeMux()
	SetupRoutes(mux)

	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + PORT,
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  30 * time.Second,
		},
	}
}

// RunServer는 서버를 시작합니다.
func (s *Server) RunServer() error {
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Shutdown은 서버를 안전하게 종료합니다.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
