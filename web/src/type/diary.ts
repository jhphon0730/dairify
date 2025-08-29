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

// 상세 조회 시 내려오는 일기 이미지 타입
export interface DiaryImage {
  id: number;
  diary_id: number;
  file_path: string;
  file_name: string;
  content_type: string;
  file_size: number;
  created_at: string;
  updated_at: string;
}

// 상세 조회 전용 일기 타입 (카테고리, 이미지가 옵션일 수 있음)
export interface DiaryDetail {
  id: number;
  creator_id: number;
  category_id?: number | null;
  title: string;
  content: string;
  created_at: string;
  updated_at: string;
  images?: DiaryImage[];
}

export interface GetDiaryByIdResponse {
  diary: DiaryDetail;
}