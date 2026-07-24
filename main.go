package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"ieltsbeyond/internal/handler"
	"ieltsbeyond/internal/repository"
)

const webDist = "web/dist"

func main() {
	repo, err := repository.NewFilePostRepository("content")
	if err != nil {
		log.Fatalf("Failed to load posts: %v", err)
	}

	h := handler.New(repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	// JSON API
	r.Route("/api", func(r chi.Router) {
		r.Get("/categories", h.HandleCategories)
		r.Get("/posts", h.HandlePosts)
		r.Get("/posts/{slug}", h.HandlePostBySlug)
	})

	// Static content assets (cover images, etc.)
	staticServer := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", staticServer))

	// React SPA
	r.Get("/*", spaHandler(webDist))

	log.Println("Server starting on http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// spaHandler serves files from the React build directory, falling back to
// index.html for any path that doesn't match a real file so client-side
// routing (react-router) can take over.
func spaHandler(distDir string) http.HandlerFunc {
	fileServer := http.FileServer(http.Dir(distDir))

	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(distDir, filepath.Clean(r.URL.Path))

		info, err := os.Stat(path)
		if err != nil || info.IsDir() {
			http.ServeFile(w, r, filepath.Join(distDir, "index.html"))
			return
		}

		fileServer.ServeHTTP(w, r)
	}
}
