package server

import (
	"context"
	"net/http"
	"time"
)

// Server 인터페이스는 서버의 기본 동작을 정의합니다.
type Server interface {
	RunServer() error
	Shutdown(ctx context.Context) error
}

// 이 구조체는 Server 인터페이스를 구현합니다.
type server struct {
	httpServer *http.Server
}

// NewServer는 새로운 Server 인스턴스를 생성합니다.
func NewServer(PORT string) Server {
	mux := http.NewServeMux()
	SetupRoutes(mux)

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

// RunServer는 서버를 시작합니다.
func (s *server) RunServer() error {
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Shutdown은 서버를 안전하게 종료합니다.
func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
