import { type ReactNode, useEffect } from "react"
import { useNavigate } from "react-router-dom"

import { IsLoggedIn } from "@/api/api"

interface GlobalLayoutProps {
  children: ReactNode
}

const GlobalLayout = ({children}: GlobalLayoutProps) => {
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
      <main className="container mx-auto px-4 py-8">{children}</main>
    </div>
  )
}

export default GlobalLayout;