import type { Response } from "@/type/api"
import type { GetDiariesDTO, GetDiariesResponse, GetDiaryByIdResponse } from '@/type/diary'
import { FetchWithAuth } from "@/api/api";

const DIARY_GET_ENDPOINT = "api/v1/diaries/list/";
const DIARY_DETAIL_ENDPOINT = (id: number | string) => `api/v1/diaries/detail/${id}/`;

export const GetDiaries = async (getProps: GetDiariesDTO): Promise<Response<GetDiariesResponse>> => {
  let search_query = new URLSearchParams();
  // 동적으로 Props에 따라 Query가 들어가는 식
  if (getProps.title) {
    search_query.append("title", getProps.title);
  } 
  if (getProps.category_id) {
    search_query.append("category_id", getProps.category_id.toString());
  }

  const res = await FetchWithAuth(DIARY_GET_ENDPOINT + "?" + search_query.toString(), {
    method: 'GET',
  })
  return {
    data: res.data,
    error: res.error,
  }
}

export const GetDiaryById = async (id: number): Promise<Response<GetDiaryByIdResponse>> => {
  const res = await FetchWithAuth(DIARY_DETAIL_ENDPOINT(id), { method: 'GET' })
  return {
    data: res.data,
    error: res.error,
  }
}