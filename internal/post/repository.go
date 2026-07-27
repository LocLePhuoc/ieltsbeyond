package post

import (
	"context"
	"errors"

	"github.com/google/uuid"
)


var (
	ErrNotFound = errors.New("post not found")
	ErrDuplicateSlug = errors.New("slug already exists")
	ErrInvalidInput  = errors.New("invalid post input")
)


type Repository interface {
	ListPublished(ctx context.Context, category *Category, limit int) ([]Post, error)
	GetPublishedBySlug(ctx context.Context, slug string) (*Post, error)
	ListAdmin(ctx context.Context) ([]Post, error)
	GetAdmin(ctx context.Context, id uuid.UUID) (*Post, error)
	Create(ctx context.Context, input UpsertPostInput) (*Post, error)
	Update(ctx context.Context, id uuid.UUID, input UpsertPostInput) (*Post, error)
	Publish(ctx context.Context, id uuid.UUID) (*Post, error)
	Unpublish(ctx context.Context, id uuid.UUID) (*Post, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
