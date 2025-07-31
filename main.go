package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jhphon0730/dairify/internal/config"
	"github.com/jhphon0730/dairify/internal/database"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// config 설정 초기화
	config, err := config.LoadConfig()
	if err != nil || config == nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 데이터베이스 연결 및 스키마 적용
	db, err := database.NewDB()
	if err != nil || db == nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close() // 서버 종료 시 DB 연결 닫기

	// 서버 옵션 설정
	PORT := config.Port
	MOD := config.AppEnv

	// HTTP 서버 설정
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is healthy"))
	})

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: mux,
	}

	// OS 종료 신호 처리
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// 서버 실행
	go func() {
		log.Printf("Server running on port %s in %s mode", PORT, MOD)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// 종료 신호 대기
	<-c
	log.Println("Shutting down server...")

	// 서버 종료
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	log.Println("Server stopped")
}
