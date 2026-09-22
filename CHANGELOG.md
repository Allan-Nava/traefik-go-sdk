# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- **Middleware Management** (v1.0.0 Feature #4)
  - `CreateHttpMiddleware(name, config)` — Create new HTTP middleware
  - `UpdateHttpMiddleware(name, config)` — Update existing HTTP middleware
  - `DeleteHttpMiddleware(name)` — Delete HTTP middleware
  - `CreateTcpMiddleware(name, config)` — Create new TCP middleware
  - `UpdateTcpMiddleware(name, config)` — Update existing TCP middleware
  - `DeleteTcpMiddleware(name)` — Delete TCP middleware
  - Full test coverage (6 test functions with 8+ test cases)
  - Proper HTTP method handling (POST/PUT/DELETE)

- **TCP/UDP Write Operations** (v1.0.0 Feature #3)
  - `CreateTcpRouter(name, config)` — Create new TCP router
  - `UpdateTcpRouter(name, config)` — Update existing TCP router
  - `DeleteTcpRouter(name)` — Delete TCP router
  - `CreateTcpService(name, config)` — Create new TCP service
  - `UpdateTcpService(name, config)` — Update existing TCP service
  - `DeleteTcpService(name)` — Delete TCP service
  - `CreateUdpRouter(name, config)` — Create new UDP router
  - `UpdateUdpRouter(name, config)` — Update existing UDP router
  - `DeleteUdpRouter(name)` — Delete UDP router
  - `CreateUdpService(name, config)` — Create new UDP service
  - `UpdateUdpService(name, config)` — Update existing UDP service
  - `DeleteUdpService(name)` — Delete UDP service
  - Full test coverage (12 test functions with 15+ test cases)
  - Proper HTTP method handling (POST/PUT/DELETE)

- **Batch Operations** (v1.0.0 Feature #2)
  - `ApplyConfiguration(config)` — Apply complete configuration
  - `ValidateConfiguration(config)` — Validate config before applying
  - `GetConfiguration()` — Retrieve current configuration
  - `ResetConfiguration()` — Reset to default state
  - Full test coverage (4 test functions with 9+ test cases)
  - Supports HTTP, TCP, UDP routers/services/middlewares in single batch

- **Configuration Write Methods** (v1.0.0 Feature #1)
  - `CreateHttpRouter(name, config)` — Create new HTTP router
  - `UpdateHttpRouter(name, config)` — Update existing HTTP router
  - `DeleteHttpRouter(name)` — Delete HTTP router
  - `CreateHttpService(name, config)` — Create new HTTP service
  - `UpdateHttpService(name, config)` — Update existing HTTP service
  - `DeleteHttpService(name)` — Delete HTTP service
  - Full test coverage (6 new test functions with 20+ test cases)
  - Proper HTTP method handling (POST/PUT/DELETE)

- Comprehensive unit test suite (80+ test cases, 84.3% coverage)
  - Table-driven tests for all public API methods
  - Edge case coverage (HTTP errors, network failures, large payloads)
  - Mock HTTP server testing with `httptest`

- TDD (Test-Driven Development) infrastructure
  - Pre-commit hooks (`make lint`) enforce testing
  - Coverage threshold (70% minimum, targeting 80%)
  - GitHub Actions CI with coverage reporting
  - `docs/DEVELOPMENT.md` with TDD workflow guide

- GitHub automation
  - Branch protection rules for `main` (1 PR review, status checks required)
  - Milestone v1.0.0 (8 issues, due 2026-10-31)
  - `scripts/setup-github.sh` for GitHub setup automation

- Backlog automation
  - `docs/backlog.md` (single source of truth, 11 items)
  - `scripts/sync-backlog.sh` (GitHub Issues ↔ backlog sync)
  - `make sync-backlog` target for manual sync

- Documentation
  - `CLAUDE.md` — Contribution guidelines (TDD, versioning, testing, commit rules)
  - `AGENTS.md` — AI agent operational rules (same as CLAUDE.md + decision log)
  - `docs/DEVELOPMENT.md` — Developer setup, TDD workflow, troubleshooting
  - `.github/BRANCH_PROTECTION.md` — Branch protection rules documentation
  - `docs/MILESTONE.md` — v1.0.0 release plan

- Makefile targets for development
  - `make lint` — Full check (fmt + vet + test + coverage)
  - `make test` — Run tests with race detector
  - `make coverage` — Generate coverage report
  - `make coverage-html` — Open coverage in browser
  - `make sync-backlog` — Sync GitHub Issues

### Fixed

- **restyGet/restyPost URL construction**: Now properly concatenates `BaseUrl` + path
  - Previously, methods passed only path (e.g., `/api/http/routers`) without BaseUrl
  - Now correctly constructs full URL for HTTP requests
  - Enables all API methods to function correctly

### Changed

- CI workflow (`.github/workflows/go-test.yml`)
  - Added race detector (`-race` flag)
  - Added coverage report generation
  - Added coverage threshold check (70% minimum)
  - Added Codecov integration

- Makefile restructured with comprehensive help (`make help`)

### Technical Debt / Backlog

- ID-003: Fix HealthCheck() error handling (currently returns nil,nil always)
- ID-004: Correct "haivision" typo in traefik.go comments
- ID-005: Upgrade Go version floor to 1.20+ (drop 1.18/1.19)
- ID-006: Document GetApiVersion() and GetApiRawData() edge cases (require Traefik 2.x+)
- ID-007: Add code coverage badge to README (codecov.io integration)
- ID-008: Create docs/ARCHITECTURE.md (component overview)

---

## [0.1.0] - 2026-09-22

### Added

- Initial implementation of Traefik Go SDK
- HTTP, TCP, UDP router and service query methods
- Entrypoint listing
- API overview and health check endpoints
- Debug mode support
- Resty HTTP client integration

### Fixed

- API version and raw data endpoints (commit cff631c)

[Unreleased]: https://github.com/Allan-Nava/traefik-go-sdk/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/Allan-Nava/traefik-go-sdk/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/Allan-Nava/traefik-go-sdk/releases/tag/v0.1.0
