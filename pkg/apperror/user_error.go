package apperror

import "errors"

var (
	ErrUserSignupUserNameRequired = errors.New("ID는 필수 입력값입니다")
	ErrUserSignupNicknameRequired = errors.New("닉네임은 필수 입력값입니다")
	ErrUserSignupPasswordRequired = errors.New("비밀번호는 필수 입력값입니다")
	ErrUserSignupEmailRequired    = errors.New("이메일은 필수 입력값입니다")

	ErrUserSignupDuplicateUserName = errors.New("이미 사용 중인 사용자 ID입니다")
	ErrUserSignupDuplicateEmail    = errors.New("이미 사용 중인 이메일입니다")
)
