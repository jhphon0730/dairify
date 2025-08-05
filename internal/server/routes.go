package server

import (
	"net/http"

	"github.com/jhphon0730/dairify/internal/database"
	"github.com/jhphon0730/dairify/internal/handler"
	"github.com/jhphon0730/dairify/internal/middleware"
	"github.com/jhphon0730/dairify/internal/repository"
	"github.com/jhphon0730/dairify/internal/service"
)

// SetupRoutes는 HTTP 라우트를 설정합니다.
func SetupRoutes(mux *http.ServeMux, db *database.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	// HTTP 연결 상태 확인 라우트 설정
	RegisterHealthRoutes(mux)

	// 사용자 관련 라우트 설정
	RegisterUserRoutes(mux, userHandler)
}

// RegisterHealthRoutes는 헬스 체크 라우트를 등록합니다.
func RegisterHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is healthy"))
	})
}

// RegisterUserRoutes는 사용자 관련 라우트를 등록합니다.
func RegisterUserRoutes(mux *http.ServeMux, userHandler handler.UserHandler) {
	mux.HandleFunc("/api/v1/users/signup/", userHandler.SignupUser)                              // 회원가입
	mux.HandleFunc("/api/v1/users/signin/", userHandler.SigninUser)                              // 로그인
	mux.HandleFunc("/api/v1/users/signout/", middleware.AuthMiddleware(userHandler.SignoutUser)) // 로그아웃
	mux.HandleFunc("/api/v1/users/profile/", middleware.AuthMiddleware(userHandler.ProfileUser)) // 프로필 조회
}
