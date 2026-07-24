package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"ieltsbeyond/internal/model"
	"ieltsbeyond/internal/repository"
	"ieltsbeyond/templates"
)

type Handler struct {
	repo repository.PostRepository
}

func New(repo repository.PostRepository) *Handler {
	return &Handler{repo: repo}
}

func isHTMX(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

func (h *Handler) render(w http.ResponseWriter, r *http.Request, content templ.Component, title string, activeTab string) {
	if isHTMX(r) {
		// HTMX request: return content fragment + sidebar OOB update
		if err := content.Render(r.Context(), w); err != nil {
			log.Printf("Error rendering content: %v", err)
		}
		if err := templates.Footer().Render(r.Context(), w); err != nil {
			log.Printf("Error rendering footer: %v", err)
		}
		if err := templates.SidebarOOB(activeTab).Render(r.Context(), w); err != nil {
			log.Printf("Error rendering sidebar OOB: %v", err)
		}
	} else {
		// Direct navigation: render full page
		page := templates.Layout(title, activeTab, content)
		if err := page.Render(r.Context(), w); err != nil {
			log.Printf("Error rendering page: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	posts, err := h.repo.GetRecentPosts(6)
	if err != nil {
		log.Printf("Error getting recent posts: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	content := templates.Home(posts)
	h.render(w, r, content, "My Blog", "home")
}

func (h *Handler) HandleCategory(w http.ResponseWriter, r *http.Request) {
	categorySlug := chi.URLParam(r, "category")

	if !model.IsValidCategory(categorySlug) {
		h.HandleNotFound(w, r)
		return
	}

	category := model.Category(categorySlug)
	posts, err := h.repo.GetPostsByCategory(category)
	if err != nil {
		log.Printf("Error getting posts for category %s: %v", categorySlug, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	description := model.CategoryDescriptions[category]
	content := templates.Category(category, description, posts)
	title := strings.ToUpper(categorySlug[:1]) + categorySlug[1:] + " - My Blog"
	h.render(w, r, content, title, categorySlug)
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	post, err := h.repo.GetPostBySlug(slug)
	if err != nil {
		h.HandleNotFound(w, r)
		return
	}

	content := templates.PostDetail(post)
	title := post.Title + " - My Blog"
	h.render(w, r, content, title, string(post.Category))
}

func (h *Handler) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	content := templates.NotFound()
	h.render(w, r, content, "Not Found - My Blog", "")
}
