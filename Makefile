.PHONY: build test coverage lint fmt vet sync-backlog sync-backlog-check help

# Default target
help:
	@echo "traefik-go-sdk — Makefile targets:"
	@echo ""
	@echo "  make build          — Build the SDK"
	@echo "  make test           — Run tests with race detector"
	@echo "  make coverage       — Generate coverage report (must be ≥80%)"
	@echo "  make coverage-html  — View coverage in HTML (requires go tool cover)"
	@echo "  make fmt            — Format code with gofmt"
	@echo "  make vet            — Run go vet static analysis"
	@echo "  make lint           — Run fmt + vet + test + coverage (pre-commit)"
	@echo "  make sync-backlog   — Sync docs/backlog.md to GitHub Issues"
	@echo "  make sync-backlog-check — Verify backlog sync consistency"
	@echo ""

build:
	@echo "Starting build..."
	@go build -v ./...
	@echo "Build complete ✓"

test:
	@echo "Starting tests..."
	@go test -v ./... -race
	@echo "Tests complete ✓"

coverage:
	@echo "Generating coverage report..."
	@go test -v ./... -coverprofile=coverage.out -covermode=atomic
	@coverage=$$(go tool cover -func=coverage.out | grep total | awk '{print $$3}' | sed 's/%//'); \
	echo "Coverage: $${coverage}%"; \
	if (( $$(echo "$$coverage < 70" | bc -l) )); then \
		echo "ERROR: Coverage $${coverage}% is below 70% threshold"; \
		exit 1; \
	fi
	@echo "Coverage check passed ✓"

coverage-html: coverage
	@echo "Opening coverage report in browser..."
	@go tool cover -html=coverage.out

fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format complete ✓"

vet:
	@echo "Running go vet..."
	@go vet ./...
	@echo "Vet check complete ✓"

lint: fmt vet test coverage
	@echo ""
	@echo "All checks passed ✓"

# Backlog automation
sync-backlog:
	@echo "Syncing backlog.md to GitHub Issues..."
	@./scripts/sync-backlog.sh
	@echo "Sync complete ✓"

sync-backlog-check:
	@echo "Checking backlog sync consistency..."
	@./scripts/sync-backlog.sh --check
	@echo "Consistency check complete ✓"