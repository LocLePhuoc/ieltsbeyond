package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/loclp/personal-website/internal/markdown"
	"github.com/loclp/personal-website/internal/model"
)

// FilePostRepository reads blog posts from markdown files in a content directory.
type FilePostRepository struct {
	contentDir string
	mu         sync.RWMutex
	posts      []model.Post
}

// NewFilePostRepository creates a new file-based repository and loads all posts.
func NewFilePostRepository(contentDir string) (*FilePostRepository, error) {
	r := &FilePostRepository{contentDir: contentDir}
	if err := r.load(); err != nil {
		return nil, fmt.Errorf("failed to load posts: %w", err)
	}
	return r, nil
}

func (r *FilePostRepository) load() error {
	var posts []model.Post

	entries, err := os.ReadDir(r.contentDir)
	if err != nil {
		return fmt.Errorf("failed to read content directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		categoryName := entry.Name()
		if !model.IsValidCategory(categoryName) {
			continue
		}

		categoryDir := filepath.Join(r.contentDir, categoryName)
		files, err := os.ReadDir(categoryDir)
		if err != nil {
			return fmt.Errorf("failed to read category directory %s: %w", categoryName, err)
		}

		for _, file := range files {
			if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
				continue
			}

			filePath := filepath.Join(categoryDir, file.Name())
			raw, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", filePath, err)
			}

			fm, body, err := markdown.Parse(raw)
			if err != nil {
				return fmt.Errorf("failed to parse %s: %w", filePath, err)
			}

			htmlContent, err := markdown.Render(body)
			if err != nil {
				return fmt.Errorf("failed to render %s: %w", filePath, err)
			}

			slug := strings.TrimSuffix(file.Name(), ".md")
			date, _ := markdown.ParseDate(fm.Date)

			coverImage := fm.CoverImage
			if coverImage == "" {
				coverImage = "/static/images/placeholder.svg"
			}

			posts = append(posts, model.Post{
				Slug:        slug,
				Title:       fm.Title,
				Summary:     fm.Summary,
				Category:    model.Category(categoryName),
				Tags:        fm.Tags,
				Date:        date,
				CoverImage:  coverImage,
				Content:     body,
				HTMLContent: htmlContent,
			})
		}
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	r.mu.Lock()
	r.posts = posts
	r.mu.Unlock()

	return nil
}

// Reload re-reads all posts from disk.
func (r *FilePostRepository) Reload() error {
	return r.load()
}

func (r *FilePostRepository) GetAllPosts() ([]model.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]model.Post, len(r.posts))
	copy(result, r.posts)
	return result, nil
}

func (r *FilePostRepository) GetPostsByCategory(category model.Category) ([]model.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []model.Post
	for _, p := range r.posts {
		if p.Category == category {
			result = append(result, p)
		}
	}
	return result, nil
}

func (r *FilePostRepository) GetPostBySlug(slug string) (*model.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, p := range r.posts {
		if p.Slug == slug {
			post := p
			return &post, nil
		}
	}
	return nil, fmt.Errorf("post not found: %s", slug)
}

func (r *FilePostRepository) GetRecentPosts(limit int) ([]model.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if limit > len(r.posts) {
		limit = len(r.posts)
	}

	result := make([]model.Post, limit)
	copy(result, r.posts[:limit])
	return result, nil
}
