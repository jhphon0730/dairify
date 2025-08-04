package apperror

import "errors"

var (
	ErrJWTExpiredToken              = errors.New("로그인 세션이 만료되었습니다")
	ErrJWTInvalidTokenSigningMethod = errors.New("올바르지 않은 토큰 서명 방법입니다")
	ErrJWTInvalidRequest            = errors.New("올바르지 않은 요청입니다")

	ErrUserNotFound     = errors.New("사용자를 찾을 수 없습니다")
	ErrAuthUnauthorized = errors.New("인증되지 않은 요청입니다")

	ErrUserSignupUserNameRequired = errors.New("ID는 필수 입력값입니다")
	ErrUserSignupNicknameRequired = errors.New("닉네임은 필수 입력값입니다")
	ErrUserSignupPasswordRequired = errors.New("비밀번호는 필수 입력값입니다")
	ErrUserSignupEmailRequired    = errors.New("이메일은 필수 입력값입니다")

	ErrUserSignupDuplicateUserName = errors.New("이미 사용 중인 사용자 ID입니다")
	ErrUserSignupDuplicateEmail    = errors.New("이미 사용 중인 이메일입니다")

	ErrUserSigninInvalidUserName = errors.New("사용자 ID가 올바르지 않습니다")
	ErrUserSigninInvalidPassword = errors.New("비밀번호가 올바르지 않습니다")
)
