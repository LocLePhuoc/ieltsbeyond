package model

import "time"

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
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Category    Category  `json:"category"`
	Tags        []string  `json:"tags"`
	Date        time.Time `json:"date"`
	CoverImage  string    `json:"coverImage"`
	Content     string    `json:"-"`                      // raw markdown
	HTMLContent string    `json:"htmlContent,omitempty"` // rendered HTML
}

type CategoryInfo struct {
	Slug        Category `json:"slug"`
	Description string   `json:"description"`
}
