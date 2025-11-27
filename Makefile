.DEFAULT_GOAL := build

BUILD_ARGS?=-ldflags="-s -w" -trimpath

build: ## Build ezpass
	CGO_ENABLED=0 go build $(BUILD_ARGS) .

install: ## Installs local builds
	CGO_ENABLED=0 go install $(BUILD_ARGS) -v .

clean: ## Clean compiled binary / misc artifacts
	rm -f ezpass

generate: ## Generate wordlist
	go generate ./...

test: ## Run tests
	go test ./...

bench: ## Run benchmark tests
	go test ./words -bench . -benchmem

lint: ## Run linters
	golangci-lint run ./...

lint-fix: ## Run linters and formatters, atttempt autofixes
	golangci-lint run ./... --fix

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*## "} /^[a-zA-Z0-9_-]+:.*## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build clean generate test bench lint lint-fix help
