package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/jhphon0730/dairify/internal/config"
	"github.com/jhphon0730/dairify/pkg/apperror"
	"github.com/jhphon0730/dairify/pkg/utils"
)

const (
	USER_TOKEN_KEY = "user:%d:token"
)

// UserRedis 인터페이스는 사용자 관련 Redis 작업을 정의합니다.
type UserRedis interface {
	SetUserToken(ctx context.Context, userID int64, token string) error
	GetUserToken(ctx context.Context, userID int64) (string, error)
	DeleteUserToken(ctx context.Context, userID int64) error
	Close() error
}

// userRedis 구조체는 Redis 클라이언트를 포함하고 UserRedis 인터페이스를 구현합니다.
type userRedis struct {
	client *redis.Client
}

var (
	user_once     sync.Once
	user_instance UserRedis
)

// NewUserRedis 함수는 Redis 클라이언트를 초기화하고 UserRedis 인스턴스를 생성합니다.
func NewUserRedis(ctx context.Context) error {
	cfg := config.GetConfig()
	db := utils.InterfaceToInt(cfg.Redis.USER_DB)

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       db,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	user_instance = &userRedis{
		client: client,
	}

	return nil
}

// GetUserRedis 함수는 UserRedis 인스턴스를 반환합니다. 싱글턴 패턴을 사용하여 인스턴스를 생성합니다.
func GetUserRedis(ctx context.Context) (UserRedis, error) {
	var err error
	user_once.Do(func() {
		err = NewUserRedis(ctx)
	})

	if err != nil {
		return nil, err
	}

	if user_instance == nil {
		return nil, apperror.ErrUserRedisIsNil
	}

	return user_instance, nil
}

// SetUserToken 함수는 사용자 토큰을 Redis에 저장합니다.
func (r *userRedis) SetUserToken(ctx context.Context, userID int64, token string) error {
	key := fmt.Sprintf(USER_TOKEN_KEY, userID)
	return r.client.Set(ctx, key, token, config.GetConfig().Redis.AccessTokenExpiry).Err()
}

// GetUserToken 함수는 사용자 ID에 해당하는 토큰을 Redis에서 조회합니다.
func (r *userRedis) GetUserToken(ctx context.Context, userID int64) (string, error) {
	key := fmt.Sprintf(USER_TOKEN_KEY, userID)
	token, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", apperror.ErrAuthRequiredToken
	}
	if err != nil {
		return "", err
	}
	return token, nil
}

// DeleteUserToken 함수는 사용자 ID에 해당하는 토큰을 Redis에서 삭제합니다.
func (r *userRedis) DeleteUserToken(ctx context.Context, userID int64) error {
	key := fmt.Sprintf(USER_TOKEN_KEY, userID)
	return r.client.Del(ctx, key).Err()
}

// Close 함수는 Redis 클라이언트를 종료합니다.
func (r *userRedis) Close() error {
	return r.client.Close()
}
