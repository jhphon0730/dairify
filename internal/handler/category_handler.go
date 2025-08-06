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

// CategoryHandler는 카테고리 관련 HTTP 요청을 처리하는 인터페이스입니다.
type CategoryHandler interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	GetCategoriesByCreatorID(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
}

// categoryHandler 구조체는 CategoryHandler 인터페이스를 구현합니다.
type categoryHandler struct {
	categoryService service.CategoryService
}

// NewCategoryHandler 함수는 CategoryHandler 인터페이스의 구현체를 반환합니다.
func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

/* CreateCategory 함수는 새로운 카테고리를 생성하는 HTTP 핸들러입니다. */
func (h *categoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	var createCategoryDTO dto.CreateCategoryDTO
	if err := json.NewDecoder(r.Body).Decode(&createCategoryDTO); err != nil && err.Error() != "EOF" {
		response.Error(w, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	createCategoryDTO.CreatorID = userID

	category, statusCode, err := h.categoryService.CreateCategory(r.Context(), createCategoryDTO)
	if err != nil {
		response.Error(w, statusCode, err.Error())
		return
	}

	res := dto.CreateCategoryResponseDTO{
		Category: category,
	}

	response.Success(w, http.StatusCreated, "Category created successfully", res)
}

// GetCategoriesByCreatorID 함수는 주어진 생성자 ID로 카테고리 목록을 조회하는 HTTP 핸들러입니다.
func (h *categoryHandler) GetCategoriesByCreatorID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, apperror.ErrAuthUnauthorized.Error())
		return
	}

	categories, statusCode, err := h.categoryService.GetCategoriesByCreatorID(r.Context(), userID)
	if err != nil {
		response.Error(w, statusCode, "Error retrieving categories: "+err.Error())
		return
	}

	res := dto.GetCategoriesResponseDTO{
		Categories: categories,
	}

	response.Success(w, http.StatusOK, "Categories retrieved successfully", res)
}

// UpdateCategory 함수는 카테고리 이름을 업데이트하는 HTTP 핸들러입니다.
func (h *categoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response.Error(w, http.StatusMethodNotAllowed, apperror.ErrHttpMethodNotAllowed.Error())
		return
	}

	id := r.PathValue("id")
	if id == "" {
		response.Error(w, http.StatusBadRequest, apperror.ErrCategoryIDIsRequired.Error())
		return
	}

	var updateCategoryDTO dto.UpdateCategoryDTO
	if err := json.NewDecoder(r.Body).Decode(&updateCategoryDTO); err != nil && err.Error() != "EOF" {
		response.Error(w, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, apperror.ErrAuthUnauthorized.Error())
		return
	}

	// updateCategoryDTO에 ID와 CreatorID를 설정
	updateCategoryDTO.ID = utils.InterfaceToInt64(id)
	updateCategoryDTO.CreatorID = userID

	category, statusCode, err := h.categoryService.UpdateCategoryName(r.Context(), updateCategoryDTO)
	if err != nil {
		response.Error(w, statusCode, "Error updating category: "+err.Error())
		return
	}

	res := dto.UpdateCategoryResponseDTO{
		Category: category,
	}

	response.Success(w, http.StatusOK, "Category updated successfully", res)
}
