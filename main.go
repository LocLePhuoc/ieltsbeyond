package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"ieltsbeyond/internal/handler"
	"ieltsbeyond/internal/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const webDist = "web/dist"

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required. Copy .env.example, start Postgres, and run migrations/import.")
	}

	ctx := context.Background()
	db, err := postgres.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	defer db.Close()

	store := postgres.NewPostRepository(db)
	publicHandler := handler.NewPublicHandler(store)
	adminHandler := handler.NewAdminHandler(store, os.Getenv("ADMIN_TOKEN"))

	writingTaskRepo := postgres.NewWritingTaskRepository(db)
	writingTaskhandler := handler.NewWritingTaskHandler(*writingTaskRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	// JSON API
	r.Route("/api", func(r chi.Router) {
		r.Get("/categories", publicHandler.HandleCategories)
		r.Get("/posts", publicHandler.HandlePosts)
		r.Get("/posts/{slug}", publicHandler.HandlePostBySlug)
		r.Route("/admin", adminHandler.Routes)
		r.Get("/writing/task1", writingTaskhandler.HandlerGetAllTask1)
		r.Get("/writing/task2", writingTaskhandler.HandlerGetAllTask2)
	})

	// Static content assets (cover images, etc.)
	staticServer := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", staticServer))

	// React SPA
	r.Get("/*", spaHandler(webDist))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
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
