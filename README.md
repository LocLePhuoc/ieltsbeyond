# IELTS Beyond

A Go + React app for IELTS Writing practice with AI-assisted assessment.

- Go backend serving the JSON API and the React SPA
- Writing Task 1 / Task 2 prompts stored in Postgres
- Submissions and assessments stored in MongoDB
- Task 1 chart images stored in S3-compatible object storage
- LLM-based assessment through OpenRouter

## API

```txt
GET  /api/writing/task1
GET  /api/writing/task1/{id}
GET  /api/writing/task1/{id}/assess?submission_id=<id>
POST /api/writing/task1/{id}/submit
GET  /api/writing/task2
```

## Requirements

- Go
- Node.js + npm
- Postgres reachable via `DATABASE_URL`
- MongoDB reachable via `SUBMISSION_REPO_URL`

## Environment

```bash
cp .env.example .env
```

Fill in `DATABASE_URL` (add `?sslmode=disable` for a local Postgres without SSL), storage, MongoDB and LLM settings.

## Setup

```bash
make setup        # install frontend dependencies
make migrate-up   # run database migrations
```

## Run locally

Terminal 1 — backend: `make dev-api`
Terminal 2 — frontend: `make dev-web` (Vite proxies `/api` to the Go server on :3000)

## Makefile commands

```txt
make setup         install frontend dependencies
make migrate-up    run database migrations
make migrate-down  roll back one migration
make dev-api       run Go backend on PORT, default 3000
make dev-web       run Vite frontend dev server
make dev           migrate-up + dev-api
make web           build React frontend
make build         build frontend and Go binary
make run           run production build
make clean         remove build artifacts
```

## Validation

```bash
go build ./... && go vet ./...
cd web && npm run build && npm run lint
```
