# Development Guide — traefik-go-sdk

This guide explains the TDD (Test-Driven Development) workflow and automation setup for contributing to traefik-go-sdk.

## Quick Start

### 1. Setup git hooks (TDD enforcement)

```bash
# Enable pre-commit hooks (auto-runs `make lint` before every commit)
git config core.hooksPath .githooks
```

Now, every `git commit` will run:
- `go fmt` (code formatting)
- `go vet` (static analysis)
- `go test -v ./... -race` (unit tests + race detector)
- Coverage check (must be ≥80%)

If any check fails, the commit is blocked. Fix and retry.

### 2. Development workflow (TDD-first)

#### For a new feature (e.g., "Add GetApiVersion() method"):

```bash
# 1. Write a test FIRST (red)
echo "
func TestGetApiVersion(t *testing.T) {
    client, _ := BuildTraefik(\"http://localhost:8080\", false)
    resp, err := client.GetApiVersion()
    if err != nil {
        t.Fatalf(\"unexpected error: %v\", err)
    }
    if resp == nil {
        t.Fatal(\"expected response, got nil\")
    }
}
" >> traefik/api_info_test.go

# 2. Run tests (expect RED)
make test

# 3. Implement the method (green)
# Edit traefik/api_info.go → add GetApiVersion() to struct + interface

# 4. Rerun tests (expect GREEN)
make test

# 5. Check coverage (must be ≥80%)
make coverage

# 6. Commit with a clear message
git add .
git commit -m "feat: add GetApiVersion() to ITraefikClient

- Queries /api/version endpoint
- Returns version struct with semantic versioning
- Tested: happy path + error cases

Closes #123"
```

#### For a bug fix (e.g., "Fix HealthCheck() error handling"):

```bash
# 1. Write a test that reproduces the bug (RED)
echo "
func TestHealthCheckNetworkError(t *testing.T) {
    client, _ := BuildTraefik(\"http://localhost:9999\", false)
    _, err := client.HealthCheck()
    if err == nil {
        t.Fatal(\"expected error (connection refused), got nil\")
    }
}
" >> traefik/traefik_test.go

# 2. Run tests (expect RED — bug confirmed)
make test

# 3. Fix the bug (green)
# Edit traefik/traefik.go → propagate resty errors in HealthCheck()

# 4. Rerun tests (expect GREEN)
make test

# 5. Check coverage
make coverage

# 6. Commit
git add .
git commit -m "fix: HealthCheck() now propagates network errors

Previously returned (nil, nil) for all cases.
Now returns (nil, error) when connection fails.

Closes #124"
```

## Makefile Targets

```bash
make help              # Show all available targets

make build            # Build the SDK
make test             # Run all tests with race detector (-race)
make coverage         # Generate coverage report + check ≥80% threshold
make coverage-html    # Open coverage report in browser

make fmt              # Format code (gofmt)
make vet              # Run go vet static analysis
make lint             # Full check: fmt + vet + test + coverage (pre-commit)

make sync-backlog     # Sync docs/backlog.md to GitHub Issues
make sync-backlog-check  # Verify backlog consistency
```

## Git Hook Behavior

The `.githooks/pre-commit` hook runs `make lint` automatically before every commit:

```
git commit -m "feat: add method"
  ↓
.githooks/pre-commit runs
  ↓
make lint (fmt → vet → test → coverage)
  ↓
✓ All pass → commit created
✗ Any fail → commit blocked, fix errors, retry
```

To skip hooks (not recommended, but possible):

```bash
git commit --no-verify -m "..."
```

## Test Structure (TDD style)

### Table-driven tests

```go
func TestGetHttpRouters(t *testing.T) {
    tests := []struct {
        name    string
        url     string
        wantErr bool
    }{
        {"valid URL", "http://localhost:8080", false},
        {"invalid URL", "", true},
        {"network error", "http://localhost:9999", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            client, _ := BuildTraefik(tt.url, false)
            _, err := client.GetHttpRouters()
            if (err != nil) != tt.wantErr {
                t.Errorf("got err=%v, want err=%v", err, tt.wantErr)
            }
        })
    }
}
```

### Mocking HTTP

Use `httptest` to mock Traefik API responses:

```go
import "net/http/httptest"

func TestGetApiVersion(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/api/version" {
            w.Header().Set("Content-Type", "application/json")
            fmt.Fprintf(w, `{"Version":"v2.0","Codename":"example"}`)
        }
    }))
    defer server.Close()
    
    client, _ := BuildTraefik(server.URL, false)
    resp, err := client.GetApiVersion()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    // ... assert resp
}
```

## Coverage Report

```bash
# Generate coverage
make coverage

# View in HTML
make coverage-html

# Manual check
go test -v ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Threshold**: Coverage must be ≥80%. CI will fail if lower.

## Backlog (docs/backlog.md)

The backlog is synced nightly to GitHub Issues:

```bash
# View current backlog
cat docs/backlog.md

# Sync manually (creates GitHub Issues)
make sync-backlog

# Check sync consistency
make sync-backlog-check
```

**Format**: Each item has a stable `[ID-NNN]` identifier, tied to GitHub issue `#NNN`.

```markdown
- `[ID-001]` Add comprehensive unit tests `[backlog]`
- `[ID-002]` Create CHANGELOG.md `[in-progress]`
- `[ID-003]` Fix HealthCheck() `[done]`
```

PR should mention issue: `Closes #001` → auto-closes on merge.

## Continuous Integration (CI)

GitHub Actions runs on every push/PR:

**Workflow**: `.github/workflows/go-test.yml`

```yaml
- Run tests on Go 1.18-1.21
- Check coverage ≥80%
- Upload coverage to Codecov
```

If CI fails:
1. Check the logs: GitHub Actions tab → job output
2. Run `make lint` locally to replicate
3. Fix errors and push again

## Tips & Tricks

### Watch mode (auto-rerun tests on file change)

```bash
# Install entr (file monitor)
brew install entr  # macOS
apt install entr   # Linux

# Watch and retest
find traefik -name "*.go" | entr make test
```

### Quick build without tests

```bash
go build -v ./...  # Skip tests
```

### Debug flag for HTTP logging

```go
client, _ := BuildTraefik("http://localhost:8080", true)  // true = debug mode
// Resty will log all HTTP requests/responses to stdout
```

## FAQ

**Q: Can I commit without running tests?**
A: No, `git commit` will run pre-commit hook and block if tests fail.

**Q: How do I skip the hook?**
A: `git commit --no-verify` (but don't make a habit of it).

**Q: What if coverage drops below 80%?**
A: Commit is blocked. Add tests to increase coverage.

**Q: Can I add a feature without tests?**
A: No, TDD is mandatory. Write test → green → commit.

**Q: How do I know if a backlog item is in GitHub Issues?**
A: Run `make sync-backlog-check`. Nightly cron syncs automatically.

---

For more, see `CLAUDE.md` and `AGENTS.md` in the repo root.
