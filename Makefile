.DEFAULT_GOAL := build

build: ## Build ezpass
	CGO_ENABLED=0 go build .

install: ## Installs local builds
	CGO_ENABLED=0 go install -v .

clean: ## Clean compiled binary / misc artifacts
	rm -f ezpass

generate: ## Generate wordlist
	go generate ./...

test: ## Run tests
	go test ./...

lint: ## Run linters
	golangci-lint run ./...

lint-fix: ## Run linters and formatters, atttempt autofixes
	golangci-lint run ./... --fix

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*## "} /^[a-zA-Z0-9_-]+:.*## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build clean generate test lint lint-fix help
