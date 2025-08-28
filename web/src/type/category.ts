export interface Category {
  id: number;
  name: string;
  creator_id: number;
  created_at: string;
}

export interface GetCategoriesResponse {
  categories: Category[];
}