import type { Response } from "@/type/api"
import type { GetCategoriesResponse } from '@/type/category'
import { FetchWithAuth } from "@/api/api";

const CATEGORY_GET_ENDPOINT = "api/v1/categories/list/";

export const GetCategories = async (): Promise<Response<GetCategoriesResponse>> => {
  const res = await FetchWithAuth(CATEGORY_GET_ENDPOINT, {
    method: 'GET',
  });

  return {
    data: res.data,
    error: res.error,
  };
};