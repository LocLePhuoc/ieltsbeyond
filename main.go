package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"ieltsbeyond/internal/handler"
	"ieltsbeyond/internal/repository"
)

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

	// Static files
	fileServer := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Pages
	r.Get("/", h.HandleHome)
	r.Get("/posts/{slug}", h.HandlePost)
	r.Get("/{category}", h.HandleCategory)

	// Catch-all 404
	r.NotFound(h.HandleNotFound)

	log.Println("Server starting on http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
