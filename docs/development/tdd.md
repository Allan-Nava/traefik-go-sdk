# TDD Workflow

Test-Driven Development is mandatory.

## Workflow

1. **RED**: Write test (expect failure)
2. **GREEN**: Implement code (make test pass)
3. **REFACTOR**: Improve without changing behavior
4. **COMMIT**: TDD enforces via pre-commit hooks

## Example

```go
// 1. Write test (RED)
func TestGetApiVersion(t *testing.T) {
    client, _ := BuildTraefik("http://localhost:8080", false)
    resp, _ := client.GetApiVersion()
    if resp == nil {
        t.Fatal("expected response")
    }
}

// 2. Implement (GREEN)
func (o *traefikSdk) GetApiVersion() (*resty.Response, error) {
    return o.restyGet(API_VERSION, nil)
}

// 3. Run tests
go test -v ./...

// 4. Commit
git commit -m "feat: add GetApiVersion()"
```

See [docs/DEVELOPMENT.md](../DEVELOPMENT.md) for full guide.
