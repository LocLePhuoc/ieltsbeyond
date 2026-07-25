package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"ieltsbeyond/internal/markdown"
	"ieltsbeyond/internal/model"
	"ieltsbeyond/internal/post"
)

func main() {
	contentDir := flag.String("content", "content", "content directory")
	flag.Parse()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	ctx := context.Background()
	db, err := post.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	defer db.Close()

	placeholderJSON, _ := json.Marshal(map[string]any{
		"type": "doc",
		"content": []map[string]any{{
			"type": "paragraph",
			"content": []map[string]string{{
				"type": "text",
				"text": "Imported from markdown. Edit this post in the CMS to normalize Tiptap content.",
			}},
		}},
	})

	count := 0
	for _, category := range model.AllCategories {
		categoryDir := filepath.Join(*contentDir, string(category))
		files, err := os.ReadDir(categoryDir)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			log.Fatalf("Failed to read %s: %v", categoryDir, err)
		}
		for _, file := range files {
			if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
				continue
			}
			path := filepath.Join(categoryDir, file.Name())
			raw, err := os.ReadFile(path)
			if err != nil {
				log.Fatalf("Failed to read %s: %v", path, err)
			}
			fm, body, err := markdown.Parse(raw)
			if err != nil {
				log.Fatalf("Failed to parse %s: %v", path, err)
			}
			if fm == nil {
				log.Printf("Skipping %s: missing frontmatter", path)
				continue
			}
			html, err := markdown.Render(body)
			if err != nil {
				log.Fatalf("Failed to render %s: %v", path, err)
			}
			publishedAt, err := markdown.ParseDate(fm.Date)
			if err != nil {
				publishedAt = time.Now()
			}
			cover := fm.CoverImage
			if cover == "" {
				cover = "/static/images/placeholder.svg"
			}
			slug := strings.TrimSuffix(file.Name(), ".md")

			tagsJSON, _ := json.Marshal(fm.Tags)
			_, err = db.Exec(ctx, `
				INSERT INTO posts (id, slug, title, summary, category, tags, cover_image, status, content_json, content_html, published_at)
				VALUES ($1, $2, $3, $4, $5, $6::jsonb, $7, 'published', $8, $9, $10)
				ON CONFLICT (slug) DO UPDATE SET
				  title = EXCLUDED.title,
				  summary = EXCLUDED.summary,
				  category = EXCLUDED.category,
				  tags = EXCLUDED.tags,
				  cover_image = EXCLUDED.cover_image,
				  status = 'published',
				  content_html = EXCLUDED.content_html,
				  published_at = EXCLUDED.published_at,
				  updated_at = now()
			`, uuid.New(), slug, fm.Title, fm.Summary, string(category), tagsJSON, cover, placeholderJSON, html, publishedAt)
			if err != nil {
				log.Fatalf("Failed to import %s: %v", path, err)
			}
			count++
		}
	}
	log.Printf("Imported %d markdown posts", count)
}
