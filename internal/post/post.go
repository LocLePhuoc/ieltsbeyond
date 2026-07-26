package post

import (
	"context"
	"errors"
	"ieltsbeyond/internal/model"

	"github.com/google/uuid"
)


var (
	ErrNotFound = errors.New("post not found")
	ErrDuplicateSlug = errors.New("slug already exists")
	ErrInvalidInput  = errors.New("invalid post input")
)


type Repository interface {
	ListPublished(ctx context.Context, category *model.Category, limit int) ([]model.Post, error)
	GetPublishedBySlug(ctx context.Context, slug string) (*model.Post, error)
	ListAdmin(ctx context.Context) ([]model.Post, error)
	GetAdmin(ctx context.Context, id uuid.UUID) (*model.Post, error)
	Create(ctx context.Context, input model.UpsertPostInput) (*model.Post, error)
	Update(ctx context.Context, id uuid.UUID, input model.UpsertPostInput) (*model.Post, error)
	Publish(ctx context.Context, id uuid.UUID) (*model.Post, error)
	Unpublish(ctx context.Context, id uuid.UUID) (*model.Post, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
