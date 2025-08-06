package apperror

import "errors"

var (
	ErrCategoryCreateDuplicateName     = errors.New("중복된 카테고리 이름입니다")
	ErrCategoryNotFound                = errors.New("카테고리를 찾을 수 없습니다")
	ErrCategoryCreateFailed            = errors.New("카테고리 생성에 실패했습니다")
	ErrCreateFailedInternalServerError = errors.New("서버 내부 오류로 카테고리 생성에 실패했습니다")
	ErrGetFailedInternalServerError    = errors.New("서버 내부 오류로 카테고리 조회에 실패했습니다")
	ErrUpdateFailedInternalServerError = errors.New("서버 내부 오류로 카테고리 업데이트에 실패했습니다")

	ErrCategoryNameRequired      = errors.New("카테고리 이름은 필수입니다")
	ErrCategoryCreatorIDRequired = errors.New("카테고리 생성자 ID는 필수입니다")
	ErrCategoryUpdateForbidden   = errors.New("카테고리 업데이트 권한이 없습니다")
	ErrCategoryIDIsRequired      = errors.New("카테고리 ID는 필수입니다")
)
