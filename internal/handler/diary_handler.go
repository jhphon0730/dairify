package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jhphon0730/dairify/internal/dto"
	"github.com/jhphon0730/dairify/internal/middleware"
	"github.com/jhphon0730/dairify/internal/response"
	"github.com/jhphon0730/dairify/internal/service"
	"github.com/jhphon0730/dairify/pkg/apperror"
	"github.com/jhphon0730/dairify/pkg/utils"
)

const ()

// DiaryHandler는 일기 관련 HTTP 요청을 처리하는 인터페이스입니다.
type DiaryHandler interface {
	GetDiaryByID(w http.ResponseWriter, r *http.Request)
	GetDiariesByCreatorID(w http.ResponseWriter, r *http.Request)
	CreateDiary(w http.ResponseWriter, r *http.Request)
	DeleteDiary(w http.ResponseWriter, r *http.Request)
	UpdateDiary(w http.ResponseWriter, r *http.Request)
	UploadDiaryImage(w http.ResponseWriter, r *http.Request)
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

// GetDiaryByID 함수는 ID로 단일 일기를 조회하는 HTTP 핸들러입니다.
func (h *diaryHandler) GetDiaryByID(w http.ResponseWriter, r *http.Request) {
	// 메서드 검증
	if r.Method != http.MethodGet {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	// 경로 변수에서 id 추출 (예: /detail/{id}/)
	idParam := r.PathValue("id")
	if idParam == "" {
		response.Error(w, http.StatusBadRequest, apperror.ErrDiaryNotFound.Error())
		return
	}

	// 사용자 인증 정보 확인
	if _, ok := middleware.GetUserIDFromContext(r.Context()); !ok {
		response.Error(w, http.StatusUnauthorized, apperror.ErrAuthUnauthorized.Error())
		return
	}

	// 문자열 ID -> int64 변환
	diaryID := utils.InterfaceToInt64(idParam)
	if diaryID == 0 { // 0은 유효하지 않은 ID로 간주
		response.Error(w, http.StatusBadRequest, apperror.ErrDiaryNotFound.Error())
		return
	}

	diary, status, err := h.diaryService.GetDiaryByID(r.Context(), diaryID)
	if err != nil {
		response.Error(w, status, err.Error())
		return
	}

	res := dto.GetDiaryByIDResponseDTO{Diary: diary}
	response.Success(w, status, "Diary retrieved successfully", res)
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
	diaries, status, err := h.diaryService.GetDiariesByCreatorID(r.Context(), userID, params)
	if err != nil {
		response.Error(w, status, err.Error())
		return
	}

	res := dto.GetDiariesByCreatorIDResponseDTO{
		Diaries: diaries,
	}

	response.Success(w, status, "Diary list retrieved successfully", res)
}

// CreateDiary 함수는 새로운 일기를 생성하는 HTTP 핸들러입니다.
func (h *diaryHandler) CreateDiary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, apperror.ErrAuthUnauthorized.Error())
		return
	}

	var createDiaryDTO dto.CreateDiaryDTO
	if err := json.NewDecoder(r.Body).Decode(&createDiaryDTO); err != nil && err.Error() != "EOF" {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	diary, status, err := h.diaryService.CreateDiary(r.Context(), createDiaryDTO, userID)
	if err != nil {
		response.Error(w, status, err.Error())
		return
	}

	res := dto.CreateDiaryResponseDTO{
		Diary: diary,
	}

	response.Success(w, status, "Diary created successfully", res)
}

// DeleteDiary 함수는 일기를 소프트 삭제 처리하는 HTTP 핸들러입니다.
func (h *diaryHandler) DeleteDiary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, apperror.ErrAuthUnauthorized.Error())
		return
	}

	idParam := r.PathValue("id")
	if idParam == "" {
		response.Error(w, http.StatusBadRequest, apperror.ErrDiaryNotFound.Error())
		return
	}
	diaryID := utils.InterfaceToInt64(idParam)
	if diaryID <= 0 {
		response.Error(w, http.StatusBadRequest, apperror.ErrDiaryNotFound.Error())
		return
	}

	status, err := h.diaryService.DeleteDiary(r.Context(), diaryID, userID)
	if err != nil {
		response.Error(w, status, err.Error())
		return
	}

	response.Success(w, status, "Diary deleted successfully", nil)
}

// UpdateDiary 함수는 일기를 수정하는 HTTP 핸들러입니다.
func (h *diaryHandler) UpdateDiary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	var updateDiaryDTO dto.UpdateDiaryDTO
	if err := json.NewDecoder(r.Body).Decode(&updateDiaryDTO); err != nil && err.Error() != "EOF" {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, apperror.ErrAuthUnauthorized.Error())
		return
	}

	idParam := r.PathValue("id")
	if idParam == "" {
		response.Error(w, http.StatusBadRequest, apperror.ErrDiaryNotFound.Error())
		return
	}
	diaryID := utils.InterfaceToInt64(idParam)
	if diaryID <= 0 {
		response.Error(w, http.StatusBadRequest, apperror.ErrDiaryNotFound.Error())
		return
	}

	status, err := h.diaryService.UpdateDiary(r.Context(), updateDiaryDTO, diaryID, userID)
	if err != nil {
		response.Error(w, status, err.Error())
		return
	}

	response.Success(w, status, "Diary updated successfully", nil)
}

// UploadDiaryImage 함수는 일기 이미지 업로드를 처리하는 HTTP 핸들러입니다.
func (h *diaryHandler) UploadDiaryImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, apperror.ErrAuthUnauthorized.Error())
		return
	}

	idParam := r.PathValue("id")
	if idParam == "" {
		response.Error(w, http.StatusBadRequest, apperror.ErrDiaryNotFound.Error())
		return
	}
	diaryID := utils.InterfaceToInt64(idParam)
	if diaryID <= 0 {
		response.Error(w, http.StatusBadRequest, apperror.ErrDiaryNotFound.Error())
		return
	}

	// Content-Type 검증 (multipart/form-data 여야 함)
	if err := utils.ValidateImageUpload(r); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// multipart/form-data 파싱
	if err := utils.ParseMultipartForm(r); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// 파일 목록 획득
	files, err := utils.GetDiaryUploadedFiles(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	diaryImages, status, err := h.diaryService.UploadDiaryImage(r.Context(), files, diaryID, userID)
	if err != nil {
		response.Error(w, status, err.Error())
		return
	}

	res := dto.UploadDiaryImageResponseDTO{
		Images: diaryImages,
	}

	response.Success(w, status, "Diary images uploaded successfully", res)
}
