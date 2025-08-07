package handler

import "github.com/jhphon0730/dairify/internal/service"

// DiaryHandler는 일기 관련 HTTP 요청을 처리하는 인터페이스입니다.
type DiaryHandler interface{}

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
