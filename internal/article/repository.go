package article

import "context"

type Repository interface {
	GetBySlug(ctx context.Context, slug string) (*Article, error)
	List(ctx context.Context, f ListFilter) ([]Article, error)
	Upsert(ctx context.Context, a *Article) error
	Unpublish(ctx context.Context, slug string) error
}

type ListFilter struct {
	Kind   string
	Tag    string
	Limit  int
	Offset int
}
