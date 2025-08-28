const API_BASE_URL =
  import.meta.env.VITE_API_BASE_URL || "http://localhost:8080/";

// tokenManager 는 토큰을 저장하고 관리하는 객체입니다.
export const tokenManager = {
  getToken: (): string | null => {
    return localStorage.getItem("token");
  },

  setToken: (token: string): void => {
    return localStorage.setItem("token", token);
  },

  clearToken: (): void => {
    return localStorage.removeItem("token");
  },
};
