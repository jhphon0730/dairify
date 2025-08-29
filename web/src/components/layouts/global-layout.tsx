import { useEffect } from "react"
import { Outlet, useNavigate } from "react-router-dom"
import Swal from "sweetalert2"

import { IsLoggedIn } from "@/api/api"
import { useAuthStore } from "@/store/auth_store"

const GlobalLayout = () => {
  const navigate = useNavigate()
  const { clearUser, user } = useAuthStore((state) => state)

  // useEffect로 URL이 바뀔 때 마다 확인
  useEffect(() => {
    console.info("Checking login status...")
    const checkLoginStatus = () => {
      if (!IsLoggedIn() || !user) {
        Swal.fire({
          title: "로그인 필요",
          text: "로그인이 필요합니다.",
          icon: "warning",
          confirmButtonText: "확인",
        })
        clearUser()
        navigate("/auth/signin")
      }
    }

    checkLoginStatus()
  }, [])

  return (
    <div className="min-h-screen bg-background">
      {/* Header */}
      <header className="text-center space-y-2 mx-0 mt-6 mb-2">
        <h1 className="text-3xl font-bold text-gray-600">Dairify</h1>
        <p className="text-muted-foreground">나만의 소중한 일기장</p>
      </header>
      <main className="container mx-auto p-2">
        <Outlet />
      </main>
    </div>
  )
}

export default GlobalLayout;