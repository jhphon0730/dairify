package database

import (
	"log"
	"strings"
	"io/ioutil"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/jhphon0730/dairify/internal/config"
)

var (
	db *sql.DB
)

// NewDB는 db 인스턴스 변수를 생성해주는 함수
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

// applySchema 함수는 
func applySchema(db *sql.DB) error {
	schema, err := ioutil.ReadFile("migrations/schema.sql")
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

func (db *DB) Close() error {
	return db.DB.Close()
}
