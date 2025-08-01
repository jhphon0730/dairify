package handler

// UserHandler 인터페이스는 사용자 관련 핸들러의 메서드를 정의합니다.
type UserHandler interface {
}

// userHandler 구조체는 UserHandler 인터페이스를 구현합니다.
type userHandler struct {
}

// NewUserHandler 함수는 UserHandler 인터페이스의 구현체를 반환합니다.
func NewUserHandler() UserHandler {
	return &userHandler{}
}
