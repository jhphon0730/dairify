package auth

import (
	"github.com/jhphon0730/dairify/internal/config"
	"github.com/jhphon0730/dairify/pkg/apperror"

	"github.com/golang-jwt/jwt/v5"

	"time"
)

var (
	jwtSecret []byte = []byte(config.GetConfig().JWT_SECRET)
)

type TokenClaims struct {
	UserID int64 `json:"userID"`
	jwt.RegisteredClaims
}

// GenerateJWTToken 함수는 사용자 ID를 기반으로 JWT 토큰을 생성합니다.
func GenerateJWTToken(userID int64) (string, error) {
	claims := TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // 1시간 후 만료
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// GenerateRefreshToken 함수는 사용자 ID를 기반으로 리프레시 토큰을 생성합니다.
func GenerateRefreshToken(userID int64) (string, error) {
	claims := TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // 7일 후 만료
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateAndParseJWT 함수는 JWT 토큰을 검증하고 파싱하여 클레임을 반환합니다.
func ValidateAndParseJWT(tokenString string) (*TokenClaims, error) {
	// 토큰을 파싱하고 클레임을 추출
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperror.ErrJWTInvalidTokenSigningMethod
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 토큰이 유효한지 확인하고, 유효한 경우 클레임을 반환
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		// 만료 여부 확인
		if claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, apperror.ErrJWTExpiredToken
		}
		return claims, nil
	}

	return nil, apperror.ErrJWTInvalidRequest
}
