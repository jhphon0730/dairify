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
          <SelectTrigger className="h-12 bg-input border-border focus:border-primary/50 focus:ring-ring shadow-sm rounded-xl">
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

      <Button className="w-full h-14 bg-secondary hover:bg-secondary/90 text-secondary-foreground shadow-md hover:shadow-lg transition-all duration-200 rounded-xl font-semibold text-lg" size="lg" >새 일기 작성</Button>

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
              className="bg-gradient-to-br from-card via-card to-card/95 border border-border/50 hover:border-primary/30 shadow-md hover:shadow-xl hover:shadow-primary/5 transition-all duration-300 cursor-pointer rounded-2xl group overflow-hidden relative backdrop-blur-sm"
            >
              <div className="absolute inset-0 bg-gradient-to-br from-primary/[0.02] via-transparent to-secondary/[0.02] opacity-0 group-hover:opacity-100 transition-opacity duration-300" />

              <CardHeader className="pb-4 relative z-10">
                <div className="flex items-start justify-between gap-3">
                  <CardTitle className="text-xl font-bold text-card-foreground text-balance group-hover:text-primary transition-colors duration-300 leading-tight">
                    {diary.title}
                  </CardTitle>
                  <Badge className="shrink-0 bg-gradient-to-r from-secondary/20 to-secondary/10 text-secondary border border-secondary/30 rounded-full px-4 py-1.5 text-xs font-semibold shadow-sm group-hover:shadow-md group-hover:scale-105 transition-all duration-200 text-black">
                    {getCategoryName(diary.category_id)}
                  </Badge>
                </div>
                <div className="flex items-center text-sm text-muted-foreground/80 mt-2">
                  <div className="flex items-center bg-muted/30 rounded-full px-3 py-1.5">
                    <Calendar className="mr-2 h-4 w-4 text-primary/70" />
                    <span className="font-medium">{formatDate(diary.created_at)}</span>
                  </div>
                </div>
              </CardHeader>

              <CardContent className="pt-0 relative z-10">
                <CardDescription className="text-card-foreground/85 text-pretty leading-relaxed mb-6 text-base">
                  {truncateContent(diary.content)}
                </CardDescription>

                <div className="flex justify-between items-center">
                  <Button
                    variant="ghost"
                    size="sm"
                    className="group/btn bg-gradient-to-r from-primary/10 to-primary/5 hover:from-primary/20 hover:to-primary/10 text-primary hover:text-primary border border-primary/20 hover:border-primary/30 font-semibold px-4 py-2 rounded-full transition-all duration-200 hover:shadow-md hover:shadow-primary/20"
                  >
                    <span className="group-hover/btn:translate-x-0.5 transition-transform duration-200">
                      자세히 보기
                    </span>
                  </Button>

                  <div className="flex items-center text-xs text-muted-foreground/60">
                    <div className="w-2 h-2 rounded-full bg-primary/30 mr-2" />
                    <span>{diary.content.length}자</span>
                  </div>
                </div>
              </CardContent>
            </Card>
          ))
        )}
      </div>
    </div>
  )
}

export default MainPage
