package main

import (
	"log"

	"github.com/jhphon0730/dairify/internal/config"
	"github.com/jhphon0730/dairify/internal/database"
)

func main() {
	// config 설정 초기화
	if c, err := config.LoadConfig(); err != nil || c == nil {
		log.Fatalln("Failed to load config:", err)
	}

	// 데이터베이스 연결 및 스키마 적용
	if d, err := database.NewDB(); err != nil || d == nil {
		log.Fatalln("Failed to initialize database:", err)
	}
}
