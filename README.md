# IELTS Knowledge Hub

A Go + React blog/CMS for IELTS content.

This project now has:

- a Go backend serving the JSON API and React SPA
- a React public blog frontend
- a Postgres-backed post store
- a simple admin CMS with Tiptap rich-text editing
- a markdown import command for existing `content/<category>/*.md` posts

## What was added

### Database/CMS backend

Added a new post module under:

```txt
internal/post/
```

It owns:

- public post queries
- admin post commands
- draft/published status rules
- slug uniqueness handling
- database row mapping
- bearer-token admin auth

Public routes are preserved:

```txt
GET /api/categories
GET /api/posts
GET /api/posts?category=<category>
GET /api/posts?limit=<n>
GET /api/posts/{slug}
```

Admin routes were added under `/api/admin`:

```txt
GET    /api/admin/posts
GET    /api/admin/posts/{id}
POST   /api/admin/posts
PUT    /api/admin/posts/{id}
POST   /api/admin/posts/{id}/publish
POST   /api/admin/posts/{id}/unpublish
DELETE /api/admin/posts/{id}
```

Admin requests require:

```txt
Authorization: Bearer <ADMIN_TOKEN>
```

### Postgres

Added:

```txt
docker-compose.yml
migrations/000001_create_posts.up.sql
migrations/000001_create_posts.down.sql
cmd/migrate/main.go
```

The `posts` table stores:

- Tiptap JSON as `content_json`
- rendered HTML as `content_html`
- draft/published status
- slug, title, summary, category, tags, cover image, timestamps

### Markdown import

Added:

```txt
cmd/import-markdown/main.go
```

This imports existing markdown posts from:

```txt
content/<category>/*.md
```

It parses frontmatter, renders markdown to HTML, and inserts/updates published posts in Postgres.

### Admin frontend

Added:

```txt
web/src/admin/
web/src/lib/adminApi.ts
web/src/components/TiptapRenderer.tsx
```

Admin pages:

```txt
/admin/posts
/admin/posts/new
/admin/posts/:id/edit
```

The editor uses Tiptap and saves both:

- `contentJSON` from `editor.getJSON()`
- `contentHTML` from `editor.getHTML()`

Public post display now renders saved Tiptap JSON using a read-only Tiptap renderer instead of only dumping raw HTML.

## Requirements

- Go
- Node.js + npm
- Docker + Docker Compose for local Postgres

> Note: Docker was not available in the coding environment, so database startup/import could not be fully run there. Go and frontend builds were validated.

## Environment

Create a local `.env` file:

```bash
cp .env.example .env
```

Default local values:

```env
DATABASE_URL=postgres://ieltsbeyond:ieltsbeyond@localhost:5432/ieltsbeyond?sslmode=disable
ADMIN_TOKEN=local-dev-token
PORT=3000
```

Use `local-dev-token` in the admin UI token box unless you change `ADMIN_TOKEN`.

## First-time setup

Install frontend dependencies:

```bash
make setup
```

Start Postgres:

```bash
make db-up
```

Run migrations:

```bash
make migrate-up
```

Import existing markdown content:

```bash
make import-markdown
```

## Run locally

Use two terminals.

Terminal 1 — backend API:

```bash
make dev-api
```

Terminal 2 — frontend dev server:

```bash
make dev-web
```

Open the frontend URL printed by Vite, commonly:

```txt
http://localhost:5173
```

or, if 5173 is busy:

```txt
http://localhost:5174
```

Admin CMS:

```txt
http://localhost:5173/admin/posts
```

or:

```txt
http://localhost:5174/admin/posts
```

Enter the admin token:

```txt
local-dev-token
```

## Common workflow

Create a post:

1. Open `/admin/posts`.
2. Enter the admin bearer token.
3. Click **New post**.
4. Fill in slug, category, title, summary, tags, cover image.
5. Write the post body in the Tiptap editor.
6. Click **Save draft**.
7. Click **Publish** when ready.

Only published posts appear on the public blog.

## Makefile commands

```txt
make setup             install frontend dependencies
make db-up             start local Postgres
make db-down           stop local Postgres
make migrate-up        run database migrations
make migrate-down      roll back one migration
make import-markdown   import content/*.md posts into Postgres
make dev-api           run Go backend on PORT, default 3000
make dev-web           run Vite frontend dev server
make dev               db-up + migrate-up + import-markdown + dev-api
make web               build React frontend
make build             build frontend and Go binary
make run               run production build
make clean             remove build artifacts
```

## Validation

These checks passed after the CMS work:

```bash
go test ./...
cd web && npm run build
cd web && npm run lint
```

The frontend build may warn that the JS chunk is larger than 500 KB because Tiptap adds editor code. This is currently only a warning.

## Notes

- The current categories are still `tech`, `gaming`, and `travelling` from the original project state.
- Imported markdown posts are public immediately.
- Imported markdown posts get placeholder Tiptap JSON; newly created CMS posts have real Tiptap JSON.
- Image upload is not implemented yet; `coverImage` is currently a URL/path string.
