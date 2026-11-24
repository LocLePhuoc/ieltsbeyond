package postgres

import (
	"context"
	"errors"
	"fmt"
	"ieltsbeyond/internal/article"
	"ieltsbeyond/internal/utils"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ article.Repository = (*ArticleRepository)(nil)

type ArticleRepository struct {
	db *pgxpool.Pool
}

func NewArticleRepository(db *pgxpool.Pool) *ArticleRepository {
	return &ArticleRepository{db: db}
}

const qGetArticleBySlug = `
SELECT
    slug, kind, title, description, tags, cover_image_key, extra,
    sections, raw_sections, file_hash, commit_sha, published_at
FROM articles
WHERE slug = $1
  AND unpublished_at IS NULL
  AND published_at <= now()`

func (s *ArticleRepository) GetBySlug(ctx context.Context, slug string) (*article.Article, error) {
	rows, err := s.db.Query(ctx, qGetArticleBySlug, slug)
	if err != nil {
		return nil, err
	}
	row, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[articleRow])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return row.toArticle()

}

func (r *ArticleRepository) List(ctx context.Context, f article.ListFilter) ([]article.Article, error) {
	panic("TO IMPLEMENT")
}

func (r *ArticleRepository) Upsert(ctx context.Context, a *article.Article) error {
	panic("TO IMPLEMENT")
}

func (r *ArticleRepository) Unpublish(ctx context.Context, slug string) error {
	panic("TO IMPLEMENT")
}

type articleRow struct {
	Slug          string    `db:"slug"`
	Kind          string    `db:"kind"`
	Title         string    `db:"title"`
	Description   string    `db:"description"`
	Tags          []string  `db:"tags"`
	CoverImageKey *string   `db:"cover_image_key"`
	Extra         []byte    `db:"extra"`
	Sections      []byte    `db:"sections"`
	RawSections   []byte    `db:"raw_sections"`
	FileHash      string    `db:"file_hash"`
	CommitSHA     string    `db:"commit_sha"`
	PublishedAt   time.Time `db:"published_at"`
}

func (r *articleRow) toArticle() (*article.Article, error) {
	a := &article.Article{
		Meta: article.Frontmatter{
			Slug:        r.Slug,
			Kind:        r.Kind,
			Title:       r.Title,
			Description: r.Description,
			Tags:        r.Tags,
			PublishedAt: r.PublishedAt,
		},
		FileHash:  r.FileHash,
		CommitSHA: r.CommitSHA,
	}

	if r.CoverImageKey != nil {
		a.Meta.CoverImage = *r.CoverImageKey
	}

	if err := utils.DecodeJSON(r.Extra, &a.Meta.Extra); err != nil {
		return nil, fmt.Errorf("article %q: extra: %w", r.Slug, err)
	}
	if err := utils.DecodeJSON(r.Sections, &a.Sections); err != nil {
		return nil, fmt.Errorf("article %q: sections: %w", r.Slug, err)
	}
	if err := utils.DecodeJSON(r.RawSections, &a.Raw); err != nil {
		return nil, fmt.Errorf("article %q: raw_sections: %w", r.Slug, err)
	}

	// Always return array, so that JSON sent to frontend is [] instead of null
	if a.Sections == nil {
		a.Sections = []article.Section{}
	}
	return a, nil
}
