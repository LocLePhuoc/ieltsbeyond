package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"ieltsbeyond/internal/model"
	"ieltsbeyond/internal/repository"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	repo repository.PostRepository
}

func New(repo repository.PostRepository) *Handler {
	return &Handler{repo: repo}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

// stripBody clears fields that list views don't need, keeping payloads small.
func stripBody(posts []model.Post) []model.Post {
	result := make([]model.Post, len(posts))
	for i, p := range posts {
		p.HTMLContent = ""
		result[i] = p
	}
	return result
}

func (h *Handler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	infos := make([]model.CategoryInfo, len(model.AllCategories))
	for i, c := range model.AllCategories {
		infos[i] = model.CategoryInfo{Slug: c, Description: model.CategoryDescriptions[c]}
	}
	writeJSON(w, http.StatusOK, infos)
}

func (h *Handler) HandlePosts(w http.ResponseWriter, r *http.Request) {
	categorySlug := r.URL.Query().Get("category")
	limitParam := r.URL.Query().Get("limit")

	var posts []model.Post
	var err error

	switch {
	case limitParam != "":
		limit, convErr := strconv.Atoi(limitParam)
		if convErr != nil || limit < 0 {
			writeError(w, http.StatusBadRequest, "invalid limit")
			return
		}
		posts, err = h.repo.GetRecentPosts(limit)
	case categorySlug != "":
		if !model.IsValidCategory(categorySlug) {
			writeError(w, http.StatusBadRequest, "invalid category")
			return
		}
		posts, err = h.repo.GetPostsByCategory(model.Category(categorySlug))
	default:
		posts, err = h.repo.GetAllPosts()
	}

	if err != nil {
		log.Printf("Error getting posts: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to load posts")
		return
	}

	writeJSON(w, http.StatusOK, stripBody(posts))
}

func (h *Handler) HandlePostBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	post, err := h.repo.GetPostBySlug(slug)
	if err != nil {
		writeError(w, http.StatusNotFound, "post not found")
		return
	}

	writeJSON(w, http.StatusOK, post)
}
