CREATE TABLE IF NOT EXISTS articles (
  slug          TEXT PRIMARY KEY,
  kind          TEXT NOT NULL,
  title         TEXT NOT NULL,
  description   TEXT NOT NULL DEFAULT '',
  tags          TEXT[] NOT NULL DEFAULT '{}',
  cover_image   TEXT NOT NULL DEFAULT '',
  extra         JSONB NOT NULL DEFAULT '{}'::jsonb,
  sections      JSONB NOT NULL DEFAULT '[]'::jsonb,
  raw_sections  JSONB NOT NULL DEFAULT '[]'::jsonb,
  file_hash     TEXT NOT NULL,
  commit_sha    TEXT NOT NULL DEFAULT '',
  published_at  TIMESTAMPTZ NOT NULL,
  unpublished_at TIMESTAMPTZ,
  created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS articles_kind_published_at_idx
  ON articles (kind, published_at DESC) WHERE unpublished_at IS NULL;

CREATE INDEX IF NOT EXISTS articles_tags_idx
  ON articles USING GIN (tags);