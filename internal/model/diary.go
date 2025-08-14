package model

// Diary는 일기(다이어리) 모델을 나타냅니다.
type Diary struct {
	ID         int64   `json:"id"`
	CreatorID  int64   `json:"creator_id"`
	CategoryID *int64  `json:"category_id,omitempty"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	IsDeleted  bool    `json:"is_deleted"`
	DeletedAt  *string `json:"deleted_at,omitempty"`
}
