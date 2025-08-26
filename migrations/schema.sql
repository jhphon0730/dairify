CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    nickname VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    creator_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_category_per_user UNIQUE (name, creator_id)
);

CREATE TABLE IF NOT EXISTS diaries (
    id SERIAL PRIMARY KEY,
    creator_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT chk_title CHECK (LENGTH(title) > 0),
    CONSTRAINT chk_content CHECK (LENGTH(content) > 0)
);

CREATE INDEX IF NOT EXISTS idx_diaries_creator_id ON diaries(creator_id);
CREATE INDEX IF NOT EXISTS idx_diaries_category_id ON diaries(category_id);

-- 소프트 삭제 컬럼이 과거 스키마에 없었던 경우를 위한 보강 (idempotent)
ALTER TABLE diaries ADD COLUMN IF NOT EXISTS is_deleted BOOLEAN NOT NULL DEFAULT FALSE;
ALTER TABLE diaries ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP NULL;
-- 컬럼 존재가 보장된 이후 인덱스 생성
CREATE INDEX IF NOT EXISTS idx_diaries_is_deleted ON diaries(is_deleted);
CREATE INDEX IF NOT EXISTS idx_diaries_is_deleted2 ON diaries(is_deleted, creator_id);

-- 이미지 메타데이터 테이블 (1:N: diary -> images)
CREATE TABLE IF NOT EXISTS "images" (
  "id" BIGSERIAL PRIMARY KEY,
  "diary_id" BIGINT NOT NULL,
  "file_path" TEXT NOT NULL,
  "file_name" TEXT NOT NULL,
  "content_type" TEXT NOT NULL,
  "file_size" BIGINT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT "fk_images_diary"
    FOREIGN KEY ("diary_id") REFERENCES "diaries"("id")
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_images_diary_id" ON "images"("diary_id");