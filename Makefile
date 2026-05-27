.DEFAULT_GOAL := help
GO          ?= go
GOLANGCI    ?= golangci-lint
BINARY      := go-vrf

.PHONY: help run build test test-integration lint fmt vet check tidy clean

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'

run: ## Run the server
	$(GO) run .

build: ## Build the binary
	$(GO) build -o $(BINARY) .

test: ## Run unit tests (excludes live NSX-T integration tests)
	$(GO) test ./...

test-integration: ## Run integration tests against a live NSX-T (needs creds)
	$(GO) test -tags=integration ./...

lint: ## Run golangci-lint
	$(GOLANGCI) run ./...

fmt: ## Format code
	$(GOLANGCI) fmt ./...

vet: ## Run go vet
	$(GO) vet ./...

check: fmt vet lint test ## Run the full pre-PR check suite

tidy: ## Tidy go.mod / go.sum
	$(GO) mod tidy

clean: ## Remove build artifacts
	rm -f $(BINARY)
