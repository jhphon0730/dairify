"use client"

import type React from "react"

import { useEffect, useState } from "react"
import Swal from "sweetalert2"

import type { Diary } from "@/type/diary"
import { GetDiaries } from "@/api/diary"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Search, Plus, Calendar, BookOpen } from "lucide-react"

const MainPage = () => {
  const [diaries, setDiaries] = useState<Diary[]>([])
  const [searchTitle, setSearchTitle] = useState<string>("")
  const [searchCategory, setSearchCategory] = useState<string | undefined>(undefined)
  const [isLoading, setIsLoading] = useState<boolean>(false)

  useEffect(() => {
    handleGetDiaries("", undefined)
  }, [])

  const handleChangeSearchTitle = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTitle(() => event.target.value)
  }

  const handleGetDiaries = async (searchT: string, searchC: number | undefined) => {
    setIsLoading(() => true)

    const res = await GetDiaries({
      title: searchT,
      category_id: searchC,
    })

    if (res.error) {
      Swal.fire({
        title: "일기 불러오기 실패",
        text: res.error,
      })
      return
    }

    setDiaries(() => res.data.diaries)
    setIsLoading(() => false)
  }

  const handleSearchSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    handleGetDiaries(searchTitle, searchCategory ? Number.parseInt(searchCategory) : undefined)
  }

  const handleResetSearch = () => {
    setSearchTitle("")
    setSearchCategory(undefined)
    handleGetDiaries("", undefined)
  }

  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return date.toLocaleDateString("ko-KR", {
      year: "numeric",
      month: "long",
      day: "numeric",
    })
  }

  const truncateContent = (content: string, maxLength = 100) => {
    return content.length > maxLength ? content.substring(0, maxLength) + "..." : content
  }

  return (
    <div className="max-w-md mx-auto p-4 space-y-6 min-h-screen bg-background">
      {/* Header */}
      <div className="text-center space-y-2">
        <h1 className="text-3xl font-bold text-gray-600">Dairify</h1>
        <p className="text-muted-foreground">나만의 소중한 일기장</p>
      </div>

      {/* Search Form */}
      <form onSubmit={handleSearchSubmit} className="space-y-4">
        <div className="relative">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground h-4 w-4" />
          <Input
            type="text"
            placeholder="일기 제목으로 검색..."
            value={searchTitle}
            onChange={handleChangeSearchTitle}
            className="pl-10 bg-white border-gray-200 focus:border-gray-400"
          />
        </div>
        <div className="flex gap-2">
          <Button type="submit" className="flex-1 bg-gray-500 hover:bg-gray-600 text-white" disabled={isLoading}>
            {isLoading ? "검색 중..." : "검색"}
          </Button>
          <Button
            type="button"
            variant="outline"
            onClick={handleResetSearch}
            className="border-gray-200 text-gray-600 hover:bg-gray-50 bg-transparent"
          >
            초기화
          </Button>
        </div>
      </form>

      {/* Add New Diary Button */}
      <Button className="w-full bg-gray-600 hover:bg-gray-700 text-white py-3" size="lg">
        <Plus className="mr-2 h-5 w-5" />새 일기 작성
      </Button>

      {/* Diary List */}
      <div className="space-y-4">
        {diaries && diaries.length === 0 && !isLoading ? (
          <Card className="bg-white border-gray-100">
            <CardContent className="flex flex-col items-center justify-center py-12 text-center">
              <BookOpen className="h-12 w-12 text-gray-300 mb-4" />
              <p className="text-gray-600 mb-2">아직 작성된 일기가 없습니다</p>
              <p className="text-sm text-gray-500">첫 번째 일기를 작성해보세요!</p>
            </CardContent>
          </Card>
        ) : (
          diaries && diaries.map((diary) => (
            <Card
              key={diary.id}
              className="bg-white border-gray-100 hover:shadow-md transition-shadow cursor-pointer hover:border-gray-200"
            >
              <CardHeader className="pb-3">
                <div className="flex items-start justify-between">
                  <CardTitle className="text-lg text-gray-800 text-balance">{diary.title}</CardTitle>
                  <Badge variant="secondary" className="ml-2 shrink-0 bg-gray-100 text-gray-700">
                    카테고리 {diary.category_id}
                  </Badge>
                </div>
                <div className="flex items-center text-sm text-gray-500">
                  <Calendar className="mr-1 h-4 w-4" />
                  {formatDate(diary.created_at)}
                </div>
              </CardHeader>
              <CardContent className="pt-0">
                <CardDescription className="text-gray-600 text-pretty">
                  {truncateContent(diary.content)}
                </CardDescription>
                <Button variant="ghost" size="sm" className="mt-3 p-0 h-auto text-gray-600 hover:text-gray-700">
                  자세히 보기 →
                </Button>
              </CardContent>
            </Card>
          ))
        )}
      </div>
    </div>
  )
}

export default MainPage
