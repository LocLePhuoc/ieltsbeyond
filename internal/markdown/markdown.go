package markdown

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

type Frontmatter struct {
	Title      string   `yaml:"title"`
	Summary    string   `yaml:"summary"`
	Date       string   `yaml:"date"`
	Tags       []string `yaml:"tags"`
	CoverImage string   `yaml:"cover_image"`
}

// Parse splits a markdown file into frontmatter and body content.
func Parse(raw []byte) (*Frontmatter, string, error) {
	content := string(raw)

	if !strings.HasPrefix(content, "---") {
		return nil, content, nil
	}

	// Find closing ---
	end := strings.Index(content[3:], "\n---")
	if end == -1 {
		return nil, content, nil
	}

	fmRaw := content[3 : end+3]
	body := content[end+3+4:] // skip past "\n---"

	var fm Frontmatter
	if err := yaml.Unmarshal([]byte(fmRaw), &fm); err != nil {
		return nil, "", fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	return &fm, strings.TrimSpace(body), nil
}

// ParseDate parses a date string in "2006-01-02" format.
func ParseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

// Render converts markdown text to HTML.
func Render(markdownBody string) (string, error) {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
			),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdownBody), &buf); err != nil {
		return "", fmt.Errorf("failed to render markdown: %w", err)
	}

	return buf.String(), nil
}
