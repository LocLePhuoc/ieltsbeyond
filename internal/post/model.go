package post

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"ieltsbeyond/internal/model"
)

const (
	StatusDraft     = "draft"
	StatusPublished = "published"
)

type Post struct {
	ID          uuid.UUID       `json:"id"`
	Slug        string          `json:"slug"`
	Title       string          `json:"title"`
	Summary     string          `json:"summary"`
	Category    model.Category  `json:"category"`
	Tags        []string        `json:"tags"`
	CoverImage  string          `json:"coverImage"`
	Status      string          `json:"status"`
	ContentJSON json.RawMessage `json:"contentJSON"`
	ContentHTML string          `json:"contentHTML"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	PublishedAt *time.Time      `json:"publishedAt,omitempty"`
}

type PublicPost struct {
	Slug        string          `json:"slug"`
	Title       string          `json:"title"`
	Summary     string          `json:"summary"`
	Category    model.Category  `json:"category"`
	Tags        []string        `json:"tags"`
	Date        time.Time       `json:"date"`
	CoverImage  string          `json:"coverImage"`
	HTMLContent string          `json:"htmlContent,omitempty"`
	ContentJSON json.RawMessage `json:"contentJSON,omitempty"`
}

type UpsertPostInput struct {
	Slug        string          `json:"slug"`
	Title       string          `json:"title"`
	Summary     string          `json:"summary"`
	Category    string          `json:"category"`
	Tags        []string        `json:"tags"`
	CoverImage  string          `json:"coverImage"`
	ContentJSON json.RawMessage `json:"contentJSON"`
	ContentHTML string          `json:"contentHTML"`
}

func (p Post) Public(includeHTML bool) PublicPost {
	date := p.CreatedAt
	if p.PublishedAt != nil {
		date = *p.PublishedAt
	}
	out := PublicPost{
		Slug:       p.Slug,
		Title:      p.Title,
		Summary:    p.Summary,
		Category:   p.Category,
		Tags:       p.Tags,
		Date:       date,
		CoverImage: p.CoverImage,
	}
	if includeHTML {
		out.HTMLContent = p.ContentHTML
		out.ContentJSON = p.ContentJSON
	}
	return out
}
