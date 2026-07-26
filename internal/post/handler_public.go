package post

import (
	"errors"
	"ieltsbeyond/internal/utils"
	"log"
	"net/http"
	"strconv"

	"ieltsbeyond/internal/model"

	"github.com/go-chi/chi/v5"
)

type PublicHandler struct {
	store Store
}

func NewPublicHandler(store Store) *PublicHandler {
	return &PublicHandler{store: store}
}

func (h *PublicHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	infos := make([]model.CategoryInfo, len(model.AllCategories))
	for i, c := range model.AllCategories {
		infos[i] = model.CategoryInfo{Slug: c, Description: model.CategoryDescriptions[c]}
	}
	utils.WriteJSON(w, http.StatusOK, infos)
}

func (h *PublicHandler) HandlePosts(w http.ResponseWriter, r *http.Request) {
	categorySlug := r.URL.Query().Get("category")
	limit := -1
	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		parsed, err := strconv.Atoi(limitParam)
		if err != nil || parsed < 0 {
			utils.WriteError(w, http.StatusBadRequest, "invalid limit")
			return
		}
		limit = parsed
	}

	var category *model.Category
	if categorySlug != "" {
		if !model.IsValidCategory(categorySlug) {
			utils.WriteError(w, http.StatusBadRequest, "invalid category")
			return
		}
		c := model.Category(categorySlug)
		category = &c
	}

	posts, err := h.store.ListPublished(r.Context(), category, limit)
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "failed to load posts")
		return
	}

	out := make([]model.PublicPost, len(posts))
	for i, p := range posts {
		out[i] = p.Public(false)
	}
	utils.WriteJSON(w, http.StatusOK, out)
}

func (h *PublicHandler) HandlePostBySlug(w http.ResponseWriter, r *http.Request) {
	p, err := h.store.GetPublishedBySlug(r.Context(), chi.URLParam(r, "slug"))
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			utils.WriteError(w, http.StatusNotFound, "post not found")
			return
		}
		log.Printf("Error getting post: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "failed to load post")
		return
	}
	utils.WriteJSON(w, http.StatusOK, p.Public(true))
}
