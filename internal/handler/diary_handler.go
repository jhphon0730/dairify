package handler

import (
	"net/http"

	"github.com/jhphon0730/dairify/internal/dto"
	"github.com/jhphon0730/dairify/internal/middleware"
	"github.com/jhphon0730/dairify/internal/response"
	"github.com/jhphon0730/dairify/internal/service"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

// DiaryHandler는 일기 관련 HTTP 요청을 처리하는 인터페이스입니다.
type DiaryHandler interface {
	GetDiariesByCreatorID(w http.ResponseWriter, r *http.Request)
}

// diaryHandler 구조체는 DiaryHandler 인터페이스를 구현합니다.
type diaryHandler struct {
	diaryService service.DiaryService
}

// NewDiaryHandler 함수는 DiaryHandler 인터페이스의 구현체를 반환합니다.
func NewDiaryHandler(diaryService service.DiaryService) DiaryHandler {
	return &diaryHandler{
		diaryService: diaryService,
	}
}

// GetDiariesByCreatorID 함수는 주어진 생성자 ID로 일기 목록을 조회하는 HTTP 핸들러입니다.
func (h *diaryHandler) GetDiariesByCreatorID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, apperror.ErrAuthUnauthorized.Error())
		return
	}

	params := r.URL.Query()
	diaries, err := h.diaryService.GetDiariesByCreatorID(r.Context(), userID, params)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := dto.GetDiariesByCreatorIDResponseDTO{
		Diaries: diaries,
	}

	response.Success(w, http.StatusOK, "Diary list retrieved successfully", res)
}
