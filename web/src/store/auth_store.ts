// 인증 상태 관리를 위한 zustand 스토어

import { create } from "zustand"
import { persist, createJSONStorage } from "zustand/middleware"
import type { User } from "@/type/auth"

// 스토리지 키 상수
const STORAGE_KEY_AUTH = "auth/user"

// 인증 상태 인터페이스
interface AuthState {
  // 현재 로그인한 사용자 정보
  user: User | null
  // 사용자 정보 설정
  setUser: (user: User) => void
  // 사용자 정보 초기화
  clearUser: () => void
}

// 사용자 정보를 영속화(persist)하여 새로고침 시에도 유지
export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      // 초기 사용자 정보
      user: null,
      // 사용자 정보 설정 함수
      setUser: (user) => set({ user }),
      // 사용자 정보 초기화 함수
      clearUser: () => set({ user: null }),
    }),
    {
      // 로컬 스토리지 키
      name: STORAGE_KEY_AUTH,
      // 로컬 스토리지 사용
      storage: createJSONStorage(() => localStorage),
      // user 필드만 저장하도록 제한
      partialize: (state) => ({ user: state.user }),
    }
  )
)