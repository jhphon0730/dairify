package config

import (
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

// Postgres 구조체는 PostgreSQL 데이터베이스 연결 정보를 포함합니다.
type Postgres struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	SSL_MODE    string
	TIMEZONE    string
}

// Redis 구조체는 Redis 데이터베이스 연결 정보를 포함합니다.
type Redis struct {
	Host     string
	Port     string
	Password string

	USER_DB           int
	AccessTokenExpiry time.Duration
}

// Config 구조체는 애플리케이션의 설정 정보를 포함합니다.
type Config struct {
	AppEnv string
	Port   string

	BCRYPT_COST string
	JWT_SECRET  string
	CHAR_SET    string

	Postgres Postgres
	Redis    Redis
}

var (
	configInstance *Config
	once           sync.Once
)

// LoadConfig 함수는 환경 변수에서 설정 정보를 로드합니다.
func LoadConfig() (*Config, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		Port:   getEnv("PORT", "8080"),

		BCRYPT_COST: getEnv("BCRYPT_COST", "5"),

		Postgres: Postgres{
			DB_HOST:     getEnv("DB_HOST", "localhost"),
			DB_USER:     getEnv("DB_USER", "postgres"),
			DB_PASSWORD: getEnv("DB_PASSWORD", "postgres"),
			DB_NAME:     getEnv("DB_NAME", "dairify"),
			DB_PORT:     getEnv("DB_PORT", "5432"),
			SSL_MODE:    getEnv("SSL_MODE", "disable"),
			TIMEZONE:    getEnv("TIMEZONE", "Asia/Shanghai"),
		},
		Redis: Redis{
			Host:              getEnv("REDIS_HOST", "localhost"),
			Port:              getEnv("REDIS_PORT", "6379"),
			Password:          getEnv("REDIS_PASSWORD", ""),
			USER_DB:           0,
			AccessTokenExpiry: time.Hour,
		},
		JWT_SECRET: getEnv("JWT_SECRET", ""),
		CHAR_SET:   getEnv("CHAR_SET", "asdqwe123"),
	}, nil
}

// GetConfig 함수는 싱글턴 패턴을 사용하여 Config 인스턴스를 반환합니다.
func GetConfig() *Config {
	once.Do(func() {
		configInstance, _ = LoadConfig()
	})
	return configInstance
}

// getEnv 함수는 환경 변수에서 값을 가져오고, 없으면 기본값을 반환합니다.
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
