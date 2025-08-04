package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/jhphon0730/dairify/internal/auth"
	"github.com/jhphon0730/dairify/internal/dto"
	"github.com/jhphon0730/dairify/internal/repository"
	"github.com/jhphon0730/dairify/pkg/apperror"
	"github.com/jhphon0730/dairify/pkg/utils"
)

// UserService 인터페이스는 사용자 관련 서비스의 메서드를 정의합니다.
type UserService interface {
	SignupUser(ctx context.Context, userSignupDTO dto.UserSignupDTO) (int64, int, error)
	SigninUser(ctx context.Context, userSigninDTO dto.UserSigninDTO) (*dto.UserSigninResponseDTO, int, error)
	Profile(ctx context.Context, userID int64) (*dto.UserProfileResponseDTO, int, error)
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

// Signup 함수는 새로운 사용자를 등록합니다.
func (s *userService) SignupUser(ctx context.Context, userSignupDTO dto.UserSignupDTO) (int64, int, error) {
	if err := userSignupDTO.CheckIsValidInput(); err != nil {
		return 0, http.StatusBadRequest, err
	}

	// 비밀번호 암호화를 위한 로직
	hashedPassword, err := utils.GenerateHashPassword(userSignupDTO.Password)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	userSignupDTO.Password = hashedPassword

	signupID, err := s.userRepository.CreateUser(ctx, userSignupDTO)
	if err != nil {
		if errors.Is(err, apperror.ErrUserSignupDuplicateEmail) {
			return 0, http.StatusConflict, err
		}
		if errors.Is(err, apperror.ErrUserSignupDuplicateUserName) {
			return 0, http.StatusConflict, err
		}
		return 0, http.StatusInternalServerError, err
	}

	return signupID, http.StatusCreated, nil
}

// SigninUser 함수는 사용자를 로그인합니다.
func (s *userService) SigninUser(ctx context.Context, userSigninDTO dto.UserSigninDTO) (*dto.UserSigninResponseDTO, int, error) {
	if err := userSigninDTO.CheckIsValidInput(); err != nil {
		return nil, http.StatusBadRequest, err
	}

	user, err := s.userRepository.FindUserByUsername(ctx, userSigninDTO.Username)
	if errors.Is(err, apperror.ErrUserNotFound) {
		return nil, http.StatusUnauthorized, apperror.ErrUserSigninInvalidUserName
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// 비밀번호 검증
	if err := utils.CompareHashAndPassword(user.Password, userSigninDTO.Password); err != nil {
		return nil, http.StatusUnauthorized, apperror.ErrUserSigninInvalidPassword
	}

	// JWT 토큰 생성 ( access, refresh )
	acessToken, err := auth.GenerateJWTToken(user.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &dto.UserSigninResponseDTO{
		AccessToken:  acessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, http.StatusOK, nil
}

// Profile 함수는 사용자의 프로필 정보를 반환합니다.
func (s *userService) Profile(ctx context.Context, userID int64) (*dto.UserProfileResponseDTO, int, error) {
	user, err := s.userRepository.FindUserByUserID(ctx, userID)
	if errors.Is(err, apperror.ErrUserNotFound) {
		return nil, http.StatusNotFound, err
	}

	return &dto.UserProfileResponseDTO{
		User: user,
	}, http.StatusOK, nil
}
