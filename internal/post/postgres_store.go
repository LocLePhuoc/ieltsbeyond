package post

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"ieltsbeyond/internal/model"
)

type PostgresStore struct {
	db *pgxpool.Pool
}

func NewPostgresStore(db *pgxpool.Pool) *PostgresStore {
	return &PostgresStore{db: db}
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

func (s *PostgresStore) ListPublished(ctx context.Context, category *model.Category, limit int) ([]Post, error) {
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

func (s *PostgresStore) GetPublishedBySlug(ctx context.Context, slug string) (*Post, error) {
	row := s.db.QueryRow(ctx, baseSelect()+" WHERE slug = $1 AND status = 'published'", slug)
	return scanPost(row)
}

func (s *PostgresStore) ListAdmin(ctx context.Context) ([]Post, error) {
	rows, err := s.db.Query(ctx, baseSelect()+" ORDER BY updated_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanPosts(rows)
}

func (s *PostgresStore) GetAdmin(ctx context.Context, id uuid.UUID) (*Post, error) {
	row := s.db.QueryRow(ctx, baseSelect()+" WHERE id = $1", id)
	return scanPost(row)
}

func (s *PostgresStore) Create(ctx context.Context, input UpsertPostInput) (*Post, error) {
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

func (s *PostgresStore) Update(ctx context.Context, id uuid.UUID, input UpsertPostInput) (*Post, error) {
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

func (s *PostgresStore) Publish(ctx context.Context, id uuid.UUID) (*Post, error) {
	row := s.db.QueryRow(ctx, `WITH updated AS (
		UPDATE posts SET status = 'published', published_at = COALESCE(published_at, now()), updated_at = now()
		WHERE id = $1 RETURNING id
	) `+baseSelect()+` WHERE id = (SELECT id FROM updated)`, id)
	return scanPost(row)
}

func (s *PostgresStore) Unpublish(ctx context.Context, id uuid.UUID) (*Post, error) {
	row := s.db.QueryRow(ctx, `WITH updated AS (
		UPDATE posts SET status = 'draft', updated_at = now()
		WHERE id = $1 RETURNING id
	) `+baseSelect()+` WHERE id = (SELECT id FROM updated)`, id)
	return scanPost(row)
}

func (s *PostgresStore) Delete(ctx context.Context, id uuid.UUID) error {
	cmd, err := s.db.Exec(ctx, "DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func baseSelect() string {
	return `SELECT id, slug, title, summary, category, tags, cover_image, status, content_json, content_html, created_at, updated_at, published_at FROM posts`
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanPost(row rowScanner) (*Post, error) {
	var p Post
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

func scanPosts(rows pgx.Rows) ([]Post, error) {
	var posts []Post
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
		return ErrNotFound
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return ErrDuplicateSlug
		}
	}
	return err
}
