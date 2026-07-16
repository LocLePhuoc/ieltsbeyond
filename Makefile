.PHONY: setup generate css build run dev clean

GOBIN := $(shell go env GOPATH)/bin
export PATH := $(GOBIN):$(PATH)

# Install tools
setup:
	go install github.com/a-h/templ/cmd/templ@latest
	npm install -g tailwindcss @tailwindcss/typography 2>/dev/null || npx tailwindcss --help > /dev/null

# Generate templ Go files
generate:
	$(GOBIN)/templ generate

# Build Tailwind CSS
css:
	npx tailwindcss -i tailwind-input.css -o static/css/output.css --minify

# Build everything
build: generate css
	go build -o bin/server ./main.go

# Run the server
run: build
	./bin/server

# Development mode
dev: generate css
	go run main.go

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f static/css/output.css
	find . -name "*_templ.go" -delete
