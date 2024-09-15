# Simple Makefile for a Go project

# Build the application
all: build

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/air-verse/air@latest
	@pnpm install
	@asdf reshim golang

build-deps:
	@echo "Building deps..."
	@templ generate
	@pnpm build:css

build: build-deps
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run: build-deps
	@go run cmd/api/main.go

# Test the application
test: build-deps
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

templ-watch:
	@templ generate --watch --proxy="http://localhost:8080" --cmd="go run cmd/api/main.go"

.PHONY: all build run test clean watch
