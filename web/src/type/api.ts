// Response는 API 응답의 형식을 정의합니다.
export interface Response<T> {
	data: T
  message?: string
	error?: string
}

// FetchOptions는 API 요청의 형식을 정의합니다.
export interface FetchOptions {
	headers?: Record<string, string>
	method?: "GET" | "POST" | "PUT" | "PATCH" | "DELETE"
	body?: string | FormData
	cache?: "no-cache" | "default" | "reload" | "force-cache" | "only-if-cached"
}