# Makefile for go-reloaded project
APP=go-reloaded

.PHONY: help run build test testcases fmt vet clean clear-cache

# =============================================================================
# 🧭 HELP SYSTEM
# =============================================================================
# You can add "##" comments after each target; make help will print them.
# Usage:
#   make help
# =============================================================================
help: ## Show this help
	@echo "Usage: make <target>"
	@echo
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# =============================================================================
# 🎯 COMMANDS
# =============================================================================

run: ## Run the project with sample data and print result
	@echo "Running sample pipeline..."
	go run ./cmd/$(APP) testdata/samples/sample.txt result.txt
	@echo "-------------------------------"
	@cat result.txt

build: ## Build the binary into ./bin/
	@echo "Building binary..."
	go build -o bin/$(APP) ./cmd/$(APP)

test: ## Run all Go tests
	@echo "Running tests..."
	go test ./...

testcases: ## Run comprehensive test cases (24 cases)
	@echo "Running comprehensive test cases..."
	go test -v -run TestCases

bench: ## Run benchmark tests with memory stats
	@echo "Running benchmarks..."
	go test -bench=. -benchmem

fmt: ## Format all Go code
	@echo "Formatting code..."
	go fmt ./...

vet: ## Run static analysis (vet)
	@echo "Running go vet..."
	go vet ./...

clean: ## Remove build artifacts
	@echo "Cleaning up..."
	rm -rf bin result.txt

clear-cache: ## Clear Go build and test cache
	@echo "Clearing Go cache..."
	go clean -cache -testcache -modcache
