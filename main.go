package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jhphon0730/dairify/internal/config"
	"github.com/jhphon0730/dairify/internal/database"
)

func main() {
	// config 설정 초기화
	config, err := config.LoadConfig()
	if err != nil || config == nil {
		log.Fatalln("Failed to load config:", err)
	}

	// 데이터베이스 연결 및 스키마 적용
	if d, err := database.NewDB(); err != nil || d == nil {
		log.Fatalln("Failed to initialize database:", err)
	}

	// 서버 옵션 설정
	PORT := config.Port
	MOD := config.AppEnv

	// 여기에 서버 추가 함수

	// OS 종료 신호 처리
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// 서버 실행
	go func() {
		// 서버 실행 로직 추가
		log.Printf("Server running on port %s in %s mode", PORT, MOD)
	}()

	// OS 종료 신호 처리
	<-c
	log.Println("Shutting down server...")
}
