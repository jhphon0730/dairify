"use client"

import { useEffect, useMemo, useState } from "react"
import { useNavigate, useParams } from "react-router-dom"
import Swal from "sweetalert2"

import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Carousel, CarouselContent, CarouselItem, CarouselNext, CarouselPrevious } from "@/components/ui/carousel"
import { Calendar, ChevronLeft, Trash2, Pencil } from "lucide-react"

import type { DiaryDetail } from "@/type/diary"
import type { Category } from "@/type/category"
import { GetDiaryById } from "@/api/diary"
import { GetCategories } from "@/api/category"

// 상수: 파일 상단에 모아 선언
const DATE_LOCALE = "ko-KR"
const IMAGE_BASE_URL = import.meta.env.VITE_API_URL?.replace(/\/$/, "") || "http://localhost:8080"

const DiaryDetailPage = () => {
  // 상태: 상세 데이터, 카테고리, 로딩
  const [diary, setDiary] = useState<DiaryDetail | null>(null)
  const [categories, setCategories] = useState<Category[]>([])
  const [isLoading, setIsLoading] = useState<boolean>(false)

  const navigate = useNavigate()
  const params = useParams()

  // 카테고리 ID -> 이름 매핑 메모이제이션
  const categoryMap = useMemo(() => {
    const map = new Map<number, string>()
    categories.forEach((c) => map.set(c.id, c.name))
    return map
  }, [categories])

  // 상세 데이터 및 카테고리 로드
  useEffect(() => {
    const idRaw = params["*"] || params.id // 라우팅에 따라 안전하게 ID 추출
    const id = Number(idRaw)
    if (!id || Number.isNaN(id)) {
      Swal.fire({ title: "잘못된 접근", text: "유효하지 않은 일기 ID입니다." })
      navigate(-1)
      return
    }

    const run = async () => {
      try {
        setIsLoading(true)
        const [detailRes, categoryRes] = await Promise.all([
          GetDiaryById(id),
          GetCategories(),
        ])

        if (detailRes.error) {
          Swal.fire({ title: "불러오기 실패", text: detailRes.error })
          navigate(-1)
          return
        }

        if (categoryRes.error) {
          Swal.fire({ title: "카테고리 불러오기 실패", text: categoryRes.error })
        } else {
          setCategories(categoryRes.data.categories)
        }

        setDiary(detailRes.data.diary)
      } finally {
        setIsLoading(false)
      }
    }

    run()
  }, [navigate, params])

  // 날짜 포맷 함수 (단일 책임)
  const formatDate = (iso: string) => {
    const d = new Date(iso)
    return d.toLocaleDateString(DATE_LOCALE, { year: "numeric", month: "long", day: "numeric" })
  }

  // 카테고리 이름 조회 (없을 경우 미지정)
  const getCategoryName = (cid?: number | null) => {
    if (!cid) return "미지정"
    return categoryMap.get(cid) ?? `카테고리 ${cid}`
  }

  // 뒤로 가기 (네비게이션)
  const handleGoBack = () => {
    navigate(-1)
  }

  // 편집 이동 (추후 구현 연결)
  const handleEdit = () => {
    Swal.fire({ title: "준비 중", text: "편집 기능은 곧 제공될 예정입니다." })
  }

  // 삭제 처리 (확인 모달만 구현)
  const handleDelete = async () => {
    const confirm = await Swal.fire({
      title: "삭제하시겠습니까?",
      text: "삭제 후 되돌릴 수 없습니다.",
      showCancelButton: true,
      confirmButtonText: "삭제",
      cancelButtonText: "취소",
      confirmButtonColor: "#ef4444",
    })
    if (confirm.isConfirmed) {
      Swal.fire({ title: "준비 중", text: "삭제 기능은 곧 제공될 예정입니다." })
    }
  }

  // 이미지 URL 조합 (서버 정적 경로 기준)
  const resolveImageUrl = (filePath: string) => {
    // 서버는 /media/ 하위에 파일을 서빙함
    if (filePath.startsWith("http")) return filePath
    return `${IMAGE_BASE_URL}/${filePath.replace(/^\/+/, "")}`
  }

  return (
    <div className="relative mx-auto max-w-md min-h-screen bg-gradient-to-b from-white to-[#F5F8FB] pb-20">
      {/* 상단 바 */}
      <div className="sticky top-0 z-20 flex items-center gap-3 px-4 py-3 border-b border-border/60 bg-white/90 backdrop-blur supports-[backdrop-filter]:bg-white/70">
        <Button variant="ghost" size="sm" className="px-2" onClick={handleGoBack}>
          <ChevronLeft className="h-5 w-5" />
        </Button>
        <div className="text-base font-semibold">일기 상세</div>
      </div>

      {/* 본문 */}
      <div className="px-4 py-4 space-y-4">
        {/* 로딩 상태 */}
        {isLoading && (
          <Card className="bg-white/90 border border-border/60 rounded-2xl">
            <CardContent className="p-4 space-y-3">
              <div className="h-6 w-1/2 bg-muted/50 rounded animate-pulse" />
              <div className="h-4 w-28 bg-muted/40 rounded animate-pulse" />
              <div className="h-24 w-full bg-muted/40 rounded animate-pulse" />
            </CardContent>
          </Card>
        )}

        {!isLoading && diary && (
          <Card className="bg-white/95 border border-border/60 rounded-2xl overflow-hidden">
            <CardContent className="p-5 space-y-5">
              {/* 상단 중앙 정렬 */}
              <h1 className="text-2xl font-bold text-center leading-snug">
                {diary.title}
              </h1>

              {/* 제목 우측 밑(작게), 모바일에서는 중앙 제목 아래 우측 정렬 */}
              <div className="flex justify-center">
                <div className="text-center">
                  <div className="flex items-center flex-col gap-2">
                    <Badge className="bg-primary/10 text-primary border border-primary/20 rounded-full px-3 py-1.5 text-xs font-semibold">
                      {getCategoryName(diary.category_id ?? null)}
                    </Badge>
                    <div className="flex items-center text-sm text-muted-foreground/80">
                      <Calendar className="mr-1.5 h-4 w-4 text-primary/70" />
                      <span className="font-medium">{formatDate(diary.created_at)}</span>
                    </div>
                  </div>
                </div>
              </div>

              {/* 이미지 슬라이더 */}
              {diary.images && diary.images.length > 0 && (
                <div className="relative flex items-center justify-center">
                  <Carousel className="rounded-xl w-full max-w-4/6">
                    <CarouselContent>
                      {diary.images.map((img) => (
                        <CarouselItem key={img.id}>
                          <img
                            src={resolveImageUrl(img.file_path)}
                            alt={img.file_name}
                            className="w-full aspect-[4/3] object-cover"
                          />
                        </CarouselItem>
                      ))}
                    </CarouselContent>
                    <CarouselPrevious />
                    <CarouselNext />
                  </Carousel>
                </div>
              )}

              {/* 본문 */}
              <div className="whitespace-pre-wrap leading-relaxed text-[15px] text-card-foreground/90">
                {diary.content}
              </div>

              {/* 액션 버튼 */}
              <div className="flex gap-2">
                <Button variant="outline" className="flex-1" onClick={handleEdit}>
                  <Pencil className="h-4 w-4 mr-1.5" />
                  편집
                </Button>
                <Button variant="destructive" className="flex-1" onClick={handleDelete}>
                  <Trash2 className="h-4 w-4 mr-1.5" />
                  삭제
                </Button>
              </div>
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  )
}

export default DiaryDetailPage
