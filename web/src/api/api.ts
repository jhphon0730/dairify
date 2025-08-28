import type { FetchOptions } from "@/type/api"

const VITE_API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/"

const defaultHeaders = {
	"Content-Type": "application/json",
}

export const tokenManager = {
  getToken(): string | null {
    return localStorage.getItem("auth_token")
  },

  setToken(token: string): void {
    localStorage.setItem("auth_token", token)
  },

  removeToken(): void {
    localStorage.removeItem("auth_token")
  },

  isTokenExpired(token: string): boolean {
    try {
      const payload = JSON.parse(atob(token.split(".")[1]))
      return payload.exp * 1000 < Date.now() // 토큰이 만료되었는지 확인
    } catch {
      return true
    }
  },
}

// 로그인 상태인지 Manager로 확인하는 함수
export const IsLoggedIn = (): boolean => {
  const token = tokenManager.getToken()
  return token !== null && !tokenManager.isTokenExpired(token)
}

// JWT 없이 요청
export const FetchWithOutAuth = async (url: string, options: FetchOptions = {}) => {
	const mergeOptions = {
		...options,
		headers: {
			...defaultHeaders,
			...options.headers,
		},
	}
	const res = await fetch(`${VITE_API_URL}${url}`, mergeOptions)

	return await res.json()
}

// JWT 포함 요청
export const FetchWithAuth = async (url: string, options: FetchOptions = {}) => {
	try {
		const token = tokenManager.getToken()
		const mergeOptions = {
			...options,
			headers: {
				...defaultHeaders,
				Authorization: `Bearer ${token}`,
				...options.headers,
			},
		}

		const res = await fetch(`${VITE_API_URL}${url}`, mergeOptions)

		// 토큰 만료 (401 Unauthorized) 처리
		if (res.status === 401) {
			handleTokenExpiration()
			throw new Error("Your session has expired. Please log in again.")
		}

		return await res.json()
	} catch (error) {
		console.error("FetchWithAuth Error:", error)
		throw error // 에러를 다시 던져서 호출한 쪽에서 핸들링할 수 있도록 함
	}
}

// JWT + FormData 요청 (파일 업로드 등)
export const FetchWithAuthFormData = async (url: string, options: FetchOptions = {}) => {
	try {
		const token = tokenManager.getToken()
		const mergeOptions = {
			...options,
			headers: {
				Authorization: `Bearer ${token}`,
				...options.headers, // Content-Type 제거 (자동 설정됨)
			},
		}

		const res = await fetch(`${VITE_API_URL}${url}`, mergeOptions)

		// 토큰 만료 (401 Unauthorized) 처리
		if (res.status === 401) {
			handleTokenExpiration()
			throw new Error("Your session has expired. Please log in again.")
		}

		return await res.json()
	} catch (error) {
		console.error("FetchWithAuthFormData Error:", error)
		throw error
	}
}

/* 토큰 만료 시 로그아웃 + 리디렉션
 */
const handleTokenExpiration = () => {
  // logout()
  tokenManager.removeToken()
	window.location.replace("/signin")
}