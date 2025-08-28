import { type ReactNode } from "react"

import { Outlet } from "react-router-dom"

// 로그인, 회원가입 페이지에서만 사용되는 레이아웃 컴포넌트
const AuthLayout = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4">
      <div className="w-full max-w-md">
        <div className="bg-white rounded-lg shadow-lg p-8">
          <div className="text-center mb-2 pb-2 border-b-2">
            <h1 className="text-3xl font-bold text-gray-900 mb-2">Dairify</h1>
            <p className="text-gray-600">직관적인 일기 작성 앱</p>
          </div>
          <Outlet />
        </div>
      </div>
    </div>
  )
}

export default AuthLayout