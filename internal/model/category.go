package model

type Category struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatorID int64  `json:"creator_id"` // CreatorID는 카테고리를 생성한 사용자의 ID입니다.
	CreatedAt string `json:"created_at"`
}
