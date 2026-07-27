package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"ieltsbeyond/internal/post"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ post.Repository = (*PostRepository)(nil)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{db: db}
}

func Connect(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}

func (s *PostRepository) ListPublished(ctx context.Context, category *post.Category, limit int) ([]post.Post, error) {
	query := baseSelect() + " WHERE status = 'published'"
	args := []any{}
	if category != nil {
		args = append(args, string(*category))
		query += fmt.Sprintf(" AND category = $%d", len(args))
	}
	query += " ORDER BY published_at DESC NULLS LAST, created_at DESC"
	if limit >= 0 {
		args = append(args, limit)
		query += fmt.Sprintf(" LIMIT $%d", len(args))
	}
	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanPosts(rows)
}

func (s *PostRepository) GetPublishedBySlug(ctx context.Context, slug string) (*post.Post, error) {
	row := s.db.QueryRow(ctx, baseSelect()+" WHERE slug = $1 AND status = 'published'", slug)
	return scanPost(row)
}

func (s *PostRepository) ListAdmin(ctx context.Context) ([]post.Post, error) {
	rows, err := s.db.Query(ctx, baseSelect()+" ORDER BY updated_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanPosts(rows)
}

func (s *PostRepository) GetAdmin(ctx context.Context, id uuid.UUID) (*post.Post, error) {
	row := s.db.QueryRow(ctx, baseSelect()+" WHERE id = $1", id)
	return scanPost(row)
}

func (s *PostRepository) Create(ctx context.Context, input post.UpsertPostInput) (*post.Post, error) {
	if err := ValidateInput(input); err != nil {
		return nil, err
	}
	id := uuid.New()
	cover := input.CoverImage
	if strings.TrimSpace(cover) == "" {
		cover = "/static/images/placeholder.svg"
	}
	tagsJSON, _ := json.Marshal(input.Tags)
	row := s.db.QueryRow(ctx, `WITH inserted AS (
		INSERT INTO posts (id, slug, title, summary, category, tags, cover_image, status, content_json, content_html)
		VALUES ($1, $2, $3, $4, $5, $6::jsonb, $7, 'draft', $8, $9)
		RETURNING id
	) `+baseSelect()+` WHERE id = (SELECT id FROM inserted)`, id, input.Slug, input.Title, input.Summary, input.Category, tagsJSON, cover, input.ContentJSON, input.ContentHTML)
	return scanPost(row)
}

func (s *PostRepository) Update(ctx context.Context, id uuid.UUID, input post.UpsertPostInput) (*post.Post, error) {
	if err := ValidateInput(input); err != nil {
		return nil, err
	}
	cover := input.CoverImage
	if strings.TrimSpace(cover) == "" {
		cover = "/static/images/placeholder.svg"
	}
	tagsJSON, _ := json.Marshal(input.Tags)
	row := s.db.QueryRow(ctx, `WITH updated AS (
		UPDATE posts
		SET slug = $2, title = $3, summary = $4, category = $5, tags = $6::jsonb, cover_image = $7, content_json = $8, content_html = $9, updated_at = now()
		WHERE id = $1
		RETURNING id
	) `+baseSelect()+` WHERE id = (SELECT id FROM updated)`, id, input.Slug, input.Title, input.Summary, input.Category, tagsJSON, cover, input.ContentJSON, input.ContentHTML)
	return scanPost(row)
}

func (s *PostRepository) Publish(ctx context.Context, id uuid.UUID) (*post.Post, error) {
	row := s.db.QueryRow(ctx, `WITH updated AS (
		UPDATE posts SET status = 'published', published_at = COALESCE(published_at, now()), updated_at = now()
		WHERE id = $1 RETURNING id
	) `+baseSelect()+` WHERE id = (SELECT id FROM updated)`, id)
	return scanPost(row)
}

func (s *PostRepository) Unpublish(ctx context.Context, id uuid.UUID) (*post.Post, error) {
	row := s.db.QueryRow(ctx, `WITH updated AS (
		UPDATE posts SET status = 'draft', updated_at = now()
		WHERE id = $1 RETURNING id
	) `+baseSelect()+` WHERE id = (SELECT id FROM updated)`, id)
	return scanPost(row)
}

func (s *PostRepository) Delete(ctx context.Context, id uuid.UUID) error {
	cmd, err := s.db.Exec(ctx, "DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return post.ErrNotFound
	}
	return nil
}

func baseSelect() string {
	return `SELECT id, slug, title, summary, category, tags, cover_image, status, content_json, content_html, created_at, updated_at, published_at FROM posts`
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanPost(row rowScanner) (*post.Post, error) {
	var p post.Post
	var tagsBytes []byte
	if err := row.Scan(&p.ID, &p.Slug, &p.Title, &p.Summary, &p.Category, &tagsBytes, &p.CoverImage, &p.Status, &p.ContentJSON, &p.ContentHTML, &p.CreatedAt, &p.UpdatedAt, &p.PublishedAt); err != nil {
		return nil, mapPgError(err)
	}
	if len(tagsBytes) > 0 {
		_ = json.Unmarshal(tagsBytes, &p.Tags)
	}
	if p.Tags == nil {
		p.Tags = []string{}
	}
	return &p, nil
}

func scanPosts(rows pgx.Rows) ([]post.Post, error) {
	var posts []post.Post
	for rows.Next() {
		p, err := scanPost(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, *p)
	}
	if err := rows.Err(); err != nil {
		return nil, mapPgError(err)
	}
	return posts, nil
}

func mapPgError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return post.ErrDuplicateSlug
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return post.ErrDuplicateSlug
		}
	}
	return err
}

func ValidateInput(input post.UpsertPostInput) error {
	if strings.TrimSpace(input.Slug) == "" || strings.TrimSpace(input.Title) == "" || strings.TrimSpace(input.Summary) == "" || strings.TrimSpace(input.Category) == "" || len(input.ContentJSON) == 0 || strings.TrimSpace(input.ContentHTML) == "" {
		return post.ErrInvalidInput
	}
	if !post.IsValidCategory(input.Category) {
		return post.ErrInvalidInput
	}
	if !jsonLooksValid(input.ContentJSON) {
		return post.ErrInvalidInput
	}
	return nil
}

func jsonLooksValid(raw []byte) bool {
	return json.Valid(raw)
}
