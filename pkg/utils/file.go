package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jhphon0730/dairify/pkg/apperror"
)

// ensureDir는 지정한 경로의 디렉터리를 생성합니다(이미 존재하면 통과).
func ensureDir(dir string) error {
	return os.MkdirAll(dir, DIR_MODE)
}

// generateUniqueName는 접두사와 확장자를 사용해 충돌 가능성이 낮은 파일명을 생성합니다.
func generateUniqueName(prefix, ext string) string {
	return fmt.Sprintf("%s%d%s", prefix, time.Now().UnixNano(), ext)
}

// saveDiaryImageToDisk 함수는 이미지를 디스크에 저장하고 경로를 반환합니다.
func saveDiaryImageToDisk(dir, prefix, fileName, contentType string, content []byte) (string, error) {
	if err := ensureDir(dir); err != nil {
		return "", err
	}
	ext := pickImageExt(fileName, contentType)
	name := generateUniqueName(prefix, ext)
	fullPath := filepath.Join(dir, name)

	if err := os.WriteFile(fullPath, content, FILE_MODE); err != nil {
		return "", apperror.ErrUploadFailedInternalServerError
	}
	return fullPath, nil
}

func removeFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
