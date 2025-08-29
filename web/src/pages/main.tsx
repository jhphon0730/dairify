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
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Search, Calendar, BookOpen, SearchX, PenTool, X } from "lucide-react"

// 상수: 매직 넘버를 의미 있는 이름으로 관리
const MAX_SNIPPET_LENGTH = 100
const DATE_LOCALE = "ko-KR"
const SKELETON_COUNT = 4

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

  // 카테고리 목록 불러오기
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

  // 검색 제목 입력값 변경 처리
  const handleChangeSearchTitle = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTitle(() => event.target.value)
  }

  // 카테고리 칩 선택 처리
  const handleChangeSearchCategory = (value: string) => {
    setSearchCategory(() => (value === "all" ? undefined : value))
  }

  // 일기 목록 불러오기
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

  // 검색 폼 제출 처리
  const handleSearchSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setHasSearched(true)
    handleGetDiaries(searchTitle, searchCategory ? Number.parseInt(searchCategory) : undefined)
  }

  // 검색 초기화 처리
  const handleResetSearch = () => {
    setSearchTitle("")
    setSearchCategory(undefined)
    setHasSearched(false)
    handleGetDiaries("", undefined)
  }

  // 날짜 포맷 변환
  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return date.toLocaleDateString(DATE_LOCALE, {
      year: "numeric",
      month: "long",
      day: "numeric",
    })
  }

  // 본문 미리보기 길이 제한
  const truncateContent = (content: string, maxLength = MAX_SNIPPET_LENGTH) => {
    return content.length > maxLength ? content.substring(0, maxLength) + "..." : content
  }

  // 카테고리 이름 조회
  const getCategoryName = (categoryId: number) => {
    const category = categories.find((cat) => cat.id === categoryId)
    return category ? category.name : `카테고리 ${categoryId}`
  }

  // 로딩 스켈레톤 카드 UI
  const renderSkeletonCard = () => (
    <Card className="bg-white/80 border border-border/60 shadow-sm rounded-2xl overflow-hidden">
      <CardHeader className="pb-4">
        <div className="h-5 w-40 bg-muted/50 rounded animate-pulse" />
        <div className="mt-3 h-4 w-24 bg-muted/40 rounded animate-pulse" />
      </CardHeader>
      <CardContent>
        <div className="space-y-2">
          <div className="h-4 w-full bg-muted/50 rounded animate-pulse" />
          <div className="h-4 w-5/6 bg-muted/40 rounded animate-pulse" />
          <div className="h-4 w-2/3 bg-muted/30 rounded animate-pulse" />
        </div>
        <div className="mt-6 h-8 w-24 bg-muted/50 rounded-full animate-pulse" />
      </CardContent>
    </Card>
  )

  return (
    <div className="relative mx-auto max-w-md min-h-screen bg-gradient-to-b from-white to-[#F5F8FB]">
      {/* 상단 고정 영역: 검색 및 카테고리 칩 */}
      <div className="sticky top-0 z-20 backdrop-blur supports-[backdrop-filter]:bg-white/70 bg-white/90 border-b border-border/60">
        <form onSubmit={handleSearchSubmit} className="px-4 py-3 space-y-3">
          {/* 검색 인풋 */}
          <div className="relative">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground h-5 w-5" />
            <Input
              type="text"
              placeholder="일기 제목으로 검색"
              value={searchTitle}
              onChange={handleChangeSearchTitle}
              className="pl-11 pr-10 h-11 bg-white border-border focus:border-primary/50 focus:ring-ring shadow-sm rounded-xl"
            />
            {searchTitle && (
              <button
                type="button"
                aria-label="검색어 초기화"
                className="absolute right-2.5 top-1/2 -translate-y-1/2 p-1 rounded-full hover:bg-muted/60"
                onClick={handleResetSearch}
              >
                <X className="h-4 w-4 text-muted-foreground" />
              </button>
            )}
          </div>

          {/* 카테고리 칩 가로 스크롤 */}
          <div className="overflow-x-auto -mx-1">
            <div className="flex gap-2 px-1 pb-2 whitespace-nowrap">
              {[
                { id: "all", name: "전체" },
                ...categories.map((c) => ({ id: String(c.id), name: c.name })),
              ].map((c) => {
                const isActive = (searchCategory ?? "all") === c.id
                return (
                  <button
                    key={c.id}
                    type="button"
                    onClick={() => handleChangeSearchCategory(c.id)}
                    className={
                      `px-3 py-2 rounded-full text-sm transition-colors border ` +
                      (isActive
                        ? "bg-primary text-primary-foreground border-primary"
                        : "bg-white text-foreground border-border hover:bg-muted")
                    }
                  >
                    {c.name}
                  </button>
                )
              })}
            </div>
          </div>

          {/* 액션 버튼들 */}
          <div className="flex gap-2">
            <Button
              type="submit"
              className="flex-1 h-11 bg-primary hover:bg-primary/90 text-primary-foreground shadow-sm rounded-xl"
              disabled={isLoading}
            >
              {isLoading ? "검색 중..." : "검색"}
            </Button>
            <Button
              type="button"
              variant="outline"
              onClick={handleResetSearch}
              className="h-11 px-5 border-border hover:bg-muted rounded-xl"
            >
              초기화
            </Button>
          </div>
        </form>
      </div>

      {/* 컨텐츠 영역 */}
      <div className="px-4 py-4 space-y-4 pb-28">
        {/* 로딩 스켈레톤 */}
        {isLoading && (
          Array.from({ length: SKELETON_COUNT }).map((_, idx) => (
            <div key={`skeleton-${idx}`}>{renderSkeletonCard()}</div>
          ))
        )}

        {diaries && diaries.length === 0 && !isLoading ? (
          <Card className="bg-white border-border/60 shadow-sm rounded-2xl">
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
              className="bg-white/95 border border-border/60 hover:border-primary/30 shadow-sm hover:shadow-md transition-all duration-200 cursor-pointer rounded-2xl group overflow-hidden relative"
            >
              <div className="absolute inset-0 bg-gradient-to-br from-primary/[0.03] via-transparent to-secondary/[0.03] opacity-0 group-hover:opacity-100 transition-opacity duration-300" />

              <CardHeader className="pb-4 relative z-10">
                <div className="flex items-start justify-between gap-3">
                  <CardTitle className="text-xl font-bold text-card-foreground text-balance group-hover:text-primary transition-colors duration-300 leading-tight">
                    {diary.title}
                  </CardTitle>
                  <Badge className="shrink-0 bg-primary/10 text-primary border border-primary/20 rounded-full px-3 py-1.5 text-xs font-semibold">
                    {getCategoryName(diary.category_id)}
                  </Badge>
                </div>
                <div className="flex items-center text-sm text-muted-foreground/80 mt-2">
                  <Calendar className="mr-2 h-4 w-4 text-primary/70" />
                  <span className="font-medium">{formatDate(diary.created_at)}</span>
                </div>
              </CardHeader>

              <CardContent className="pt-0 relative z-10">
                <CardDescription className="text-card-foreground/85 leading-relaxed mb-6 text-[15px]">
                  {truncateContent(diary.content)}
                </CardDescription>

                <div className="flex justify-between items-center">
                  <Button
                    variant="ghost"
                    size="sm"
                    className="group/btn text-primary hover:bg-primary/10 font-semibold px-3 py-2 rounded-full"
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

      {/* 하단 플로팅 작성 버튼 (모바일 전용) */}
      <div className="fixed bottom-4 left-1/2 -translate-x-1/2 z-30 w-[min(92%,420px)] px-4">
        <Button
          className="w-full h-14 rounded-full bg-gradient-to-r from-primary to-primary/90 text-primary-foreground shadow-lg active:scale-[0.99]"
          size="lg"
        >
          <div className="flex items-center justify-center gap-2">
            <PenTool className="h-5 w-5" />
            <span className="font-semibold">새 일기 작성</span>
          </div>
        </Button>
      </div>
    </div>
  )
}

export default MainPage
