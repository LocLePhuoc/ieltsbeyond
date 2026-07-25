package post

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/google/uuid"
	"ieltsbeyond/internal/model"
)

var (
	ErrNotFound      = errors.New("post not found")
	ErrDuplicateSlug = errors.New("slug already exists")
	ErrInvalidInput  = errors.New("invalid post input")
)

type Store interface {
	ListPublished(ctx context.Context, category *model.Category, limit int) ([]Post, error)
	GetPublishedBySlug(ctx context.Context, slug string) (*Post, error)
	ListAdmin(ctx context.Context) ([]Post, error)
	GetAdmin(ctx context.Context, id uuid.UUID) (*Post, error)
	Create(ctx context.Context, input UpsertPostInput) (*Post, error)
	Update(ctx context.Context, id uuid.UUID, input UpsertPostInput) (*Post, error)
	Publish(ctx context.Context, id uuid.UUID) (*Post, error)
	Unpublish(ctx context.Context, id uuid.UUID) (*Post, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func ValidateInput(input UpsertPostInput) error {
	if strings.TrimSpace(input.Slug) == "" || strings.TrimSpace(input.Title) == "" || strings.TrimSpace(input.Summary) == "" || strings.TrimSpace(input.Category) == "" || len(input.ContentJSON) == 0 || strings.TrimSpace(input.ContentHTML) == "" {
		return ErrInvalidInput
	}
	if !model.IsValidCategory(input.Category) {
		return ErrInvalidInput
	}
	if !jsonLooksValid(input.ContentJSON) {
		return ErrInvalidInput
	}
	return nil
}

func jsonLooksValid(raw []byte) bool {
	return json.Valid(raw)
}
