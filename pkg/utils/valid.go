package utils

import (
	"net/http"
	"strings"

	"github.com/jhphon0730/dairify/pkg/apperror"
)

// ValidateImageUpload는 이미지 업로드 요청의 유효성을 검사합니다.
func ValidateImageUpload(r *http.Request) error {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		return apperror.ErrImageInvalidContentType
	}
	return nil
}
