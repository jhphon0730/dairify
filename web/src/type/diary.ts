export interface Diary {
  id: number;
  creator_id: number; // 작성자 ID
  category_id: number;
  title: string;
  content: string;
  created_at: string;
  updated_at: string;
}

export interface GetDiariesDTO {
  title?: string; // 제목으로 필터링
  category_id?: number; // 카테고리 ID로 필터링
}

export interface GetDiariesResponse {
  diaries: Diary[];
}