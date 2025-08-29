"use client"

import type React from "react"

import { useEffect, useState } from "react"
import Swal from "sweetalert2"

import type { Diary } from "@/type/diary"
import type { Category } from "@/type/category"
import { GetDiaries } from "@/api/diary"
import { GetCategories } from "@/api/category"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Search, Calendar, BookOpen, SearchX, PenTool } from "lucide-react"

const MainPage = () => {
  const [diaries, setDiaries] = useState<Diary[]>([])
  const [categories, setCategories] = useState<Category[]>([])
  const [searchTitle, setSearchTitle] = useState<string>("")
  const [searchCategory, setSearchCategory] = useState<string | undefined>(undefined)
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [hasSearched, setHasSearched] = useState<boolean>(false)

  useEffect(() => {
    handleGetDiaries("", undefined)
    handleGetCategories()
  }, [])

  const handleGetCategories = async () => {
    const res = await GetCategories()

    if (res.error) {
      Swal.fire({
        title: "카테고리 불러오기 실패",
        text: res.error,
      })
      return
    }

    setCategories(() => res.data.categories)
  }

  const handleChangeSearchTitle = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTitle(() => event.target.value)
  }

  const handleChangeSearchCategory = (value: string) => {
    setSearchCategory(() => (value === "all" ? undefined : value))
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
    setHasSearched(true)
    handleGetDiaries(searchTitle, searchCategory ? Number.parseInt(searchCategory) : undefined)
  }

  const handleResetSearch = () => {
    setSearchTitle("")
    setSearchCategory(undefined)
    setHasSearched(false)
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

  const getCategoryName = (categoryId: number) => {
    const category = categories.find((cat) => cat.id === categoryId)
    return category ? category.name : `카테고리 ${categoryId}`
  }

  return (
    <div className="max-w-md mx-auto p-4 space-y-6 min-h-screen bg-background">
      <form onSubmit={handleSearchSubmit} className="space-y-4">
        <div className="relative">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground h-5 w-5" />
          <Input
            type="text"
            placeholder="일기 제목으로 검색..."
            value={searchTitle}
            onChange={handleChangeSearchTitle}
            className="pl-11 h-12 bg-input border-border focus:border-primary/50 focus:ring-ring shadow-sm rounded-xl"
          />
        </div>

        <Select value={searchCategory || "all"} onValueChange={handleChangeSearchCategory}>
          <SelectTrigger className="h-12 bg-input border-border focus:border-primary/50 focus:ring-ring shadow-sm rounded-xl w-full">
            <SelectValue placeholder="카테고리 선택" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">전체 카테고리</SelectItem>
            {categories.map((category) => (
              <SelectItem key={category.id} value={category.id.toString()}>
                {category.name}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>

        <div className="flex gap-3">
          <Button
            type="submit"
            className="flex-1 h-12 bg-primary hover:bg-primary/90 text-primary-foreground shadow-md hover:shadow-lg transition-all duration-200 rounded-xl font-medium"
            disabled={isLoading}
          >
            {isLoading ? "검색 중..." : "검색"}
          </Button>
          <Button
            type="button"
            variant="outline"
            onClick={handleResetSearch}
            className="h-12 px-6 border-border hover:bg-muted hover:border-primary/30 transition-all duration-200 rounded-xl bg-transparent"
          >
            초기화
          </Button>
        </div>
      </form>

      <Button
        className="w-full h-14 bg-secondary hover:bg-secondary/90 text-secondary-foreground shadow-md hover:shadow-lg transition-all duration-200 rounded-xl font-semibold text-lg"
        size="lg"
      >
        새 일기 작성
      </Button>

      <div className="space-y-4">
        {diaries && diaries.length === 0 && !isLoading ? (
          <Card className="bg-card border-border shadow-sm rounded-xl">
            <CardContent className="flex flex-col items-center justify-center py-16 text-center">
              {hasSearched ? (
                <>
                  <SearchX className="h-16 w-16 text-muted-foreground/60 mb-6" />
                  <p className="text-foreground text-lg font-medium mb-2">검색 결과가 없습니다</p>
                  <p className="text-muted-foreground">다른 키워드로 검색해보세요</p>
                </>
              ) : (
                <>
                  <BookOpen className="h-16 w-16 text-muted-foreground/60 mb-6" />
                  <p className="text-foreground text-lg font-medium mb-2">아직 작성된 일기가 없습니다</p>
                  <p className="text-muted-foreground">첫 번째 일기를 작성해보세요!</p>
                </>
              )}
            </CardContent>
          </Card>
        ) : (
          diaries &&
          diaries.map((diary) => (
            <Card
              key={diary.id}
              className="bg-card border-border hover:shadow-lg hover:border-primary/20 transition-all duration-200 cursor-pointer rounded-xl group"
            >
              <CardHeader className="pb-3">
                <div className="flex items-start justify-between">
                  <CardTitle className="text-xl text-card-foreground text-balance group-hover:text-primary transition-colors duration-200">
                    {diary.title}
                  </CardTitle>
                  <Badge className="ml-2 shrink-0 bg-secondary/10 text-secondary border-secondary/20 rounded-lg px-3 py-1">
                    {getCategoryName(diary.category_id)}
                  </Badge>
                </div>
                <div className="flex items-center text-sm text-muted-foreground">
                  <Calendar className="mr-2 h-4 w-4" />
                  {formatDate(diary.created_at)}
                </div>
              </CardHeader>
              <CardContent className="pt-0">
                <CardDescription className="text-card-foreground/80 text-pretty leading-relaxed mb-4">
                  {truncateContent(diary.content)}
                </CardDescription>
                <Button
                  variant="ghost"
                  size="sm"
                  className="p-0 h-auto text-primary hover:text-primary/80 font-medium group-hover:translate-x-1 transition-all duration-200"
                >
                  자세히 보기
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
