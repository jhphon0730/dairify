import { useEffect } from "react"
import { Outlet, useNavigate } from "react-router-dom"

import { IsLoggedIn } from "@/api/api"

const GlobalLayout = () => {
  const navigate = useNavigate()

  // useEffect로 URL이 바뀔 때 마다 확인
  useEffect(() => {
    console.info("Checking login status...")
    const checkLoginStatus = () => {
      if (!IsLoggedIn()) {
        navigate("/auth/signin")
      }
    }

    checkLoginStatus()
  }, [])

  return (
    <div className="min-h-screen bg-background">
      {/* Header */}
      <header className="text-center space-y-2 mx-0 my-4">
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