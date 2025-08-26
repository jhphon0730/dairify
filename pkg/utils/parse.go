package utils

import (
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/jhphon0730/dairify/pkg/apperror"
)

// ParseMultipartForm은 multipart/form-data 형식의 요청을 파싱합니다.
func ParseMultipartForm(r *http.Request) error {
	if err := r.ParseMultipartForm(MAX_UPLOAD_MEMORY); err != nil {
		return err
	}
	return nil
}

// ParseImagesByFileHeader는 multipart.FileHeader를 기반으로 이미지를 파싱합니다.
func ParseImagesByFileHeader(fileHeader *multipart.FileHeader) (*ParseFileHeaderImages, error) {
	// 개별 파트의 Content-Type 확인 (image/* 인지 검사)
	contentType := fileHeader.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, CONTENT_TYPE_IMAGE_PREFIX) {
		return nil, apperror.ErrImageInvalidContentType
	}

	// 파일 크기 유효성 검사
	if fileHeader.Size <= 0 || fileHeader.Size > int64(MAX_IMAGE_SIZE) {
		return nil, apperror.ErrImageInvalidSize
	}

	// 파일 열기
	file, err := fileHeader.Open()
	if err != nil {
		return nil, apperror.ErrUploadFailedInternalServerError
	}

	// 파일 내용 읽기 (과도한 읽기를 방지하기 위해 LimitReader 사용)
	data, err := io.ReadAll(io.LimitReader(file, int64(MAX_IMAGE_SIZE)+1))
	_ = file.Close()
	if err != nil {
		return nil, apperror.ErrUploadFailedInternalServerError
	}

	// 읽은 데이터 크기 최종 확인
	if len(data) == 0 || len(data) > MAX_IMAGE_SIZE {
		return nil, apperror.ErrImageInvalidSize
	}

	return &ParseFileHeaderImages{
		FileName:    fileHeader.Filename,
		ContentType: contentType,
		Size:        fileHeader.Size,
		Content:     data,
	}, nil
}
