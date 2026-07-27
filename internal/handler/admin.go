package handler

import (
	"encoding/json"
	"errors"
	"ieltsbeyond/internal/post"
	"ieltsbeyond/internal/utils"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type AdminHandler struct {
	store      post.Repository
	adminToken string
}

func NewAdminHandler(repository post.Repository, adminToken string) *AdminHandler {
	return &AdminHandler{store: repository, adminToken: adminToken}
}

func (h *AdminHandler) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h.adminToken == "" {
			utils.WriteError(w, http.StatusUnauthorized, "admin token is not configured")
			return
		}
		got := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if got == "" || got != h.adminToken {
			utils.WriteError(w, http.StatusUnauthorized, "invalid admin token")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *AdminHandler) Routes(r chi.Router) {
	r.Use(h.Middleware)
	r.Get("/posts", h.ListPosts)
	r.Get("/posts/{id}", h.GetPost)
	r.Post("/posts", h.CreatePost)
	r.Put("/posts/{id}", h.UpdatePost)
	r.Post("/posts/{id}/publish", h.PublishPost)
	r.Post("/posts/{id}/unpublish", h.UnpublishPost)
	r.Delete("/posts/{id}", h.DeletePost)
}

func (h *AdminHandler) ListPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.store.ListAdmin(r.Context())
	if err != nil {
		log.Printf("Error listing admin posts: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "failed to load posts")
		return
	}
	utils.WriteJSON(w, http.StatusOK, posts)
}

func (h *AdminHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	p, err := h.store.GetAdmin(r.Context(), id)
	writePostOrError(w, p, err)
}

func (h *AdminHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	input, ok := decodeInput(w, r)
	if !ok {
		return
	}
	p, err := h.store.Create(r.Context(), input)
	if err != nil {
		writeStoreError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, p)
}

func (h *AdminHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	input, ok := decodeInput(w, r)
	if !ok {
		return
	}
	p, err := h.store.Update(r.Context(), id, input)
	writePostOrError(w, p, err)
}

func (h *AdminHandler) PublishPost(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	p, err := h.store.Publish(r.Context(), id)
	writePostOrError(w, p, err)
}

func (h *AdminHandler) UnpublishPost(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	p, err := h.store.Unpublish(r.Context(), id)
	writePostOrError(w, p, err)
}

func (h *AdminHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	if err := h.store.Delete(r.Context(), id); err != nil {
		writeStoreError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func parseID(w http.ResponseWriter, r *http.Request) (uuid.UUID, bool) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid post id")
		return uuid.Nil, false
	}
	return id, true
}

func decodeInput(w http.ResponseWriter, r *http.Request) (post.UpsertPostInput, bool) {
	defer r.Body.Close()
	var input post.UpsertPostInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return input, false
	}
	return input, true
}

func writePostOrError(w http.ResponseWriter, p *post.Post, err error) {
	if err != nil {
		writeStoreError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, p)
}

func writeStoreError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, post.ErrNotFound):
		utils.WriteError(w, http.StatusNotFound, "post not found")
	case errors.Is(err, post.ErrDuplicateSlug):
		utils.WriteError(w, http.StatusConflict, "slug already exists")
	case errors.Is(err, post.ErrInvalidInput):
		utils.WriteError(w, http.StatusBadRequest, "invalid post input")
	default:
		log.Printf("Post store error: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "post operation failed")
	}
}
