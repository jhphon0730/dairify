package database

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"sync"

	_ "github.com/lib/pq"

	"github.com/jhphon0730/dairify/internal/config"
)

var (
	db   *sql.DB
	once sync.Once
)

// NewDB 함수는 데이터베이스 연결을 생성하고 스키마를 적용하는 함수
func NewDB() (*sql.DB, error) {
	cfg := config.GetConfig()

	connStr := "host=" + cfg.Postgres.DB_HOST +
		" port=" + cfg.Postgres.DB_PORT +
		" user=" + cfg.Postgres.DB_USER +
		" password=" + cfg.Postgres.DB_PASSWORD +
		" dbname=" + cfg.Postgres.DB_NAME +
		" sslmode=" + cfg.Postgres.SSL_MODE
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// 스키마 파일 읽기 및 실행
	if err := applySchema(db); err != nil {
		return nil, err
	}

	return db, nil
}

// GetDB 함수는 db 인스턴스를 반환하는 함수
func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = NewDB()
		if err != nil {
			log.Fatalln("Failed to connect to database:", err)
		}
	})

	return db
}

// applySchema 함수는 데이터베이스에 스키마를 적용하는 함수
func applySchema(db *sql.DB) error {
	schema, err := os.ReadFile("migrations/schema.sql")
	if err != nil {
		return err
	}

	// 쿼리를 ;로 분리
	queries := strings.Split(string(schema), ";")
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			return err
		}
	}

	log.Println("Schema applied successfully")
	return nil
}

// Close 함수는 db 인스턴스를 닫아주는 함수
func Close() error {
	return db.Close()
}
