package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Category string

const (
	CategoryTech       Category = "tech"
	CategoryGaming     Category = "gaming"
	CategoryTravelling Category = "travelling"
)

var AllCategories = []Category{CategoryTech, CategoryGaming, CategoryTravelling}

var CategoryDescriptions = map[Category]string{
	CategoryTech:       "Software engineering, tools, and the web.",
	CategoryGaming:     "Reviews, setups, and gaming culture.",
	CategoryTravelling: "Places explored and stories from the road.",
}

func IsValidCategory(s string) bool {
	for _, c := range AllCategories {
		if string(c) == s {
			return true
		}
	}
	return false
}

type Post struct {
	ID          uuid.UUID       `json:"id"`
	Slug        string          `json:"slug"`
	Title       string          `json:"title"`
	Summary     string          `json:"summary"`
	Category    Category  `json:"category"`
	Tags        []string        `json:"tags"`
	CoverImage  string          `json:"coverImage"`
	Status      string          `json:"status"`
	ContentJSON json.RawMessage `json:"contentJSON"`
	ContentHTML string          `json:"contentHTML"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	PublishedAt *time.Time      `json:"publishedAt,omitempty"`
}

type CategoryInfo struct {
	Slug        Category `json:"slug"`
	Description string   `json:"description"`
}

type PublicPost struct {
	Slug        string          `json:"slug"`
	Title       string          `json:"title"`
	Summary     string          `json:"summary"`
	Category    Category  `json:"category"`
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