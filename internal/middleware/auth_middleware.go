package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/jhphon0730/dairify/internal/auth"
	"github.com/jhphon0730/dairify/internal/response"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

// 컨텍스트 키 상수 정의
const (
	USER_ID_CTX_KEY userIDContextKey = iota // 사용자 ID 컨텍스트 키
)

// 컨텍스트 키 타입 정의
type userIDContextKey int

// AuthMiddleware는 JWT 토큰을 검증하고 사용자 ID를 컨텍스트에 추가합니다.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Authorization 헤더에서 토큰 추출
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.Error(w, http.StatusUnauthorized, apperror.ErrAuthRequiredToken.Error())
			return
		}

		// 토큰에서 사용자 ID 추출
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			response.Error(w, http.StatusUnauthorized, apperror.ErrAuthRequiredToken.Error())
			return
		}
		claims, err := auth.ValidateAndParseJWT(token)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, apperror.ErrAuthInvalidToken.Error())
			return
		}

		// 컨텍스트에 사용자 ID 추가
		userID := claims.UserID
		ctx := context.WithValue(r.Context(), USER_ID_CTX_KEY, int64(userID)) // 사용자 정의 타입 키 사용
		next(w, r.WithContext(ctx))
	}
}

// GetUserIDFromContext는 컨텍스트에서 사용자 ID를 반환합니다.
func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(USER_ID_CTX_KEY).(int64) // 사용자 정의 타입 키 사용
	return userID, ok
}
