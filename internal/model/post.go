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
	Slug        string
	Title       string
	Summary     string
	Category    Category
	Tags        []string
	Date        time.Time
	CoverImage  string
	Content     string // raw markdown
	HTMLContent string // rendered HTML
}
