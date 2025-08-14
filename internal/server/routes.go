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
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	diaryRepository := repository.NewDiaryRepository(db)
	diaryService := service.NewDiaryService(diaryRepository)

	userHandler := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	diaryHandler := handler.NewDiaryHandler(diaryService)

	// HTTP 연결 상태 확인 라우트 설정
	RegisterHealthRoutes(mux)

	RegisterUserRoutes(mux, userHandler)
	RegisterCategoryRoutes(mux, categoryHandler)
	RegisterDiaryRoutes(mux, diaryHandler)
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
	api_v1_users := http.NewServeMux()

	api_v1_users.HandleFunc("/signup/", middleware.LoggingMiddleware(userHandler.SignupUser))                // 회원가입
	api_v1_users.HandleFunc("/signin/", middleware.LoggingMiddleware(userHandler.SigninUser))                // 로그인
	api_v1_users.HandleFunc("/signout/", middleware.ChainLoggingWithAuthMiddleware(userHandler.SignoutUser)) // 로그아웃
	api_v1_users.HandleFunc("/profile/", middleware.ChainLoggingWithAuthMiddleware(userHandler.ProfileUser)) // 프로필 조회

	mux.Handle("/api/v1/users/", http.StripPrefix("/api/v1/users", api_v1_users))
}

// RegisterCateogoryRoutes는 카테고리 관련 라우트를 등록합니다.
func RegisterCategoryRoutes(mux *http.ServeMux, categoryHandler handler.CategoryHandler) {
	api_v1_categories := http.NewServeMux()

	api_v1_categories.HandleFunc("/create/", middleware.ChainLoggingWithAuthMiddleware(categoryHandler.CreateCategory))         // 카테고리 생성
	api_v1_categories.HandleFunc("/list/", middleware.ChainLoggingWithAuthMiddleware(categoryHandler.GetCategoriesByCreatorID)) // 카테고리 목록 조회
	api_v1_categories.HandleFunc("/update/{id}/", middleware.ChainLoggingWithAuthMiddleware(categoryHandler.UpdateCategory))    // 카테고리 이름 업데이트
	api_v1_categories.HandleFunc("/delete/{id}/", middleware.ChainLoggingWithAuthMiddleware(categoryHandler.DeleteCategory))    // 카테고리 삭제

	mux.Handle("/api/v1/categories/", http.StripPrefix("/api/v1/categories", api_v1_categories))
}

// RegisterDiaryRoutes는 일기 관련 라우트를 등록합니다.
func RegisterDiaryRoutes(mux *http.ServeMux, diaryHandler handler.DiaryHandler) {
	api_v1_diaries := http.NewServeMux()

	api_v1_diaries.HandleFunc("/list/", middleware.ChainLoggingWithAuthMiddleware(diaryHandler.GetDiariesByCreatorID)) // 일기 목록 조회
	api_v1_diaries.HandleFunc("/create/", middleware.ChainLoggingWithAuthMiddleware(diaryHandler.CreateDiary))         // 일기 생성
	api_v1_diaries.HandleFunc("/detail/{id}/", middleware.ChainLoggingWithAuthMiddleware(diaryHandler.GetDiaryByID))   // 일기 단건 조회

	mux.Handle("/api/v1/diaries/", http.StripPrefix("/api/v1/diaries", api_v1_diaries))
}
