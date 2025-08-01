package service

import "github.com/jhphon0730/dairify/internal/repository"

// UserService 인터페이스는 사용자 관련 서비스의 메서드를 정의합니다.
type UserService interface {
}

// userService 구조체는 UserService 인터페이스를 구현합니다.
type userService struct {
	userRepository repository.UserRepository
}

// NewUserService 함수는 UserService 인터페이스의 구현체를 반환합니다.
func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
