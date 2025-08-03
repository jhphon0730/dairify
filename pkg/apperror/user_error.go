package apperror

import "errors"

var (
	Err_USER_SIGNUP_USERNAME_REQUIRED = errors.New("ID는 필수 입력값입니다")
	Err_USER_SIGNUP_NICKNAME_REQUIRED = errors.New("닉네임은 필수 입력값입니다")
	Err_USER_SIGNUP_PASSWORD_REQUIRED = errors.New("비밀번호는 필수 입력값입니다")
	Err_USER_SIGNUP_EMAIL_REQUIRED    = errors.New("이메일은 필수 입력값입니다")
)
