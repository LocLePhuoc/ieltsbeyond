-include .env
export

.PHONY: setup web build run dev-api dev-web db-up db-down migrate-up migrate-down import-markdown dev clean

# Install dependencies
setup:
	cd web && npm install

# Build the React frontend
web:
	cd web && npm run build

# Build everything (frontend + Go binary)
build: web
	go build -o bin/server ./main.go

# Run the production build
run: build
	./bin/server

# Run the Go API/SPA server (serves web/dist as built)
dev-api:
	go run main.go

# Run the Vite dev server with hot reload; proxies /api and /static to
# the Go server on :3000, so run `make dev-api` in another terminal too.
dev-web:
	cd web && npm run dev

# Start local Postgres
db-up:
	docker compose up -d postgres

# Stop local Postgres
db-down:
	docker compose down

# Run database migrations
migrate-up:
	go run ./cmd/migrate -direction up

# Roll back one database migration
migrate-down:
	go run ./cmd/migrate -direction down

# Import existing markdown content into Postgres
import-markdown:
	go run ./cmd/import-markdown

# Start database, migrate, import content, and run API
# Run `make dev-web` separately for the Vite frontend.
dev: db-up migrate-up import-markdown dev-api

# Clean build artifacts
clean:
	rm -rf bin web/dist
