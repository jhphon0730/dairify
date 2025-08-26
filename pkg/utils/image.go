package utils

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/jhphon0730/dairify/internal/model"
	"github.com/jhphon0730/dairify/pkg/apperror"
)

const (
	// multipart form 필드명
	DIARY_IMAGE_FIELD_NAME = "images" // multipart 필드명

	// 저장 경로
	DIARY_IMAGE_UPLOAD_DIR = "media/uploads/diary" // 다이어리 이미지 저장 경로

	// 권한
	FILE_MODE = 0o644 // 파일 권한
	DIR_MODE  = 0o755 // 디렉터리 권한

	//  이미지 크기 및 타입/확장자
	MAX_UPLOAD_MEMORY         = 10 << 20 // 10MB
	MAX_IMAGE_SIZE            = MAX_UPLOAD_MEMORY
	CONTENT_TYPE_IMAGE_PREFIX = "image/"
	DEFAULT_IMAGE_EXT         = ".jpg" // 기본 이미지 확장자

	// 접두사
	DIARY_FILENAME_PREFIX = "diary_"
)

// ParseFileHeaderImages 구조체는 multipart.FileHeader를 파싱하여 이미지 정보를 추출합니다.
type ParseFileHeaderImages struct {
	FileName    string
	ContentType string
	Size        int64
	Content     []byte
}

// pickImageExt는 파일명과 Content-Type을 바탕으로 적절한 확장자를 결정합니다.
func pickImageExt(fileName, contentType string) string {
	ext := strings.ToLower(filepath.Ext(fileName))
	if ext != "" {
		return ext
	}
	switch strings.ToLower(contentType) {
	case "image/jpeg", "image/jpg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "image/webp":
		return ".webp"
	case "image/svg+xml":
		return ".svg"
	default:
		return DEFAULT_IMAGE_EXT
	}
}

// GetDiaryUploadedFiles는 업로드된 파일 목록을 반환합니다.
func GetDiaryUploadedFiles(r *http.Request) ([]*multipart.FileHeader, error) {
	files := r.MultipartForm.File[DIARY_IMAGE_FIELD_NAME]
	if len(files) == 0 {
		return nil, apperror.ErrImageIsRequired
	}
	return files, nil
}

// 다이어리 이미지를 업로드하여 디스크에 저장하고 저장된 경로 목록을 반환합니다.
func UploadDiaryImage(file *multipart.FileHeader, diaryID int64) (*model.DiaryImage, error) {
	// 파일 헤더 파싱
	img, err := ParseImagesByFileHeader(file)
	if err != nil {
		return nil, err
	}

	// 저장 디렉터리 준비
	if err := ensureDir(DIARY_IMAGE_UPLOAD_DIR); err != nil {
		return nil, err
	}

	// 이미지 저장
	fullPath, err := saveDiaryImageToDisk(DIARY_IMAGE_UPLOAD_DIR, DIARY_FILENAME_PREFIX, img.FileName, img.ContentType, img.Content)
	if err != nil {
		return nil, err
	}

	return &model.DiaryImage{
		DiaryID:     diaryID,
		FilePath:    fullPath,
		FileName:    img.FileName,
		ContentType: img.ContentType,
		FileSize:    img.Size,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// RemoveDiaryImages 함수는 다이어리 이미지 목록을 삭제합니다
func RemoveDiaryImages(files []*model.DiaryImage) error {
	if len(files) == 0 {
		return apperror.ErrEmptySlice
	}

	for _, file := range files {
		if err := removeFile(file.FilePath); err != nil {
			return err
		}
	}
	return nil
}
