package article

import "time"

type Frontmatter struct {
	Slug        string         `yaml:"slug" json:"slug"`
	Kind        string         `yaml:"kind" json:"kind"`
	Title       string         `yaml:"title" json:"title"`
	Description string         `yaml:"description" json:"description,omitempty"`
	Tags        []string       `yaml:"tags" json:"tags,omitempty"`
	CoverImage  string         `yaml:"coverImage" json:"-"`
	PublishedAt time.Time      `yaml:"publishedAt" json:"publishedAt"`
	Extra       map[string]any `yaml:",inline" json:"extra,omitempty"`
}

type ParsedFile struct {
	Meta     Frontmatter
	Sections []RawSection
}

type Article struct {
	Meta          Frontmatter  `json:"meta"`
	CoverImageURL string       `json:"coverImageUrl,omitempty"`
	Sections      []Section    `json:"sections"`
	FileHash      string       `json:"-"`
	CommitSHA     string       `json:"-"`
	Raw           []RawSection `json:"-"`
}

// shortened for list display
type Summary struct {
	Slug        string    `json:"slug"`
	Kind        string    `json:"kind"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	PublishedAt time.Time `json:"publishedAt"`
}

func (a *Article) InternalSection(key string) (RawSection, bool) {
	if a == nil || key == "" {
		return RawSection{}, false
	}
	for _, s := range a.Raw {
		if s.Key == key && s.Visibility == VisibilityInternal {
			return s, true
		}
	}
	return RawSection{}, false
}
