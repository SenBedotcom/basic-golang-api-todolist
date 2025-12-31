.PHONY: help run build test clean docker-up docker-down install

run:
	go run main.go

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	go mod download
	go mod tidy

build: ## Build the application
	go build -o bin/todo-api cmdapi/main.go

test: ## Run tests
	go test -v ./...

clean: ## Clean build files
	rm -rf bin/

docker-up: ## Start PostgreSQL with Docker Compose
	docker-compose up -d

docker-down: ## Stop PostgreSQL with Docker Compose
	docker-compose down

docker-logs: ## View PostgreSQL logs
	docker-compose logs -f postgres

setup: docker-up install ## Setup the project (start DB and install dependencies)
	@echo "Waiting for PostgreSQL to be ready..."
	@sleep 5
	@echo "Setup complete! Run 'make run' to start the application"
