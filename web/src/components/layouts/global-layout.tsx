import { type ReactNode } from "react"

interface GlobalLayoutProps {
  children: ReactNode
}

const GlobalLayout = ({children}: GlobalLayoutProps) => {
  return (
    <div className="min-h-screen bg-background">
      <main className="container mx-auto px-4 py-8">{children}</main>
    </div>
  )
}

export default GlobalLayout;