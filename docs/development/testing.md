# Testing Guide

## Test Structure

Table-driven tests for all public methods:

```go
func TestGetHttpRouters(t *testing.T) {
    tests := []struct {
        name       string
        serverResp string
        wantErr    bool
    }{
        {"valid", `{}`, false},
        {"error", `error`, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test code here
        })
    }
}
```

## Coverage

- **Threshold**: 70% (CI enforces)
- **Target**: 80%
- **Check**: `make coverage`

## Mocking

Use `httptest.NewServer` for HTTP mocking:

```go
server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{}`))
}))
defer server.Close()

client, _ := BuildTraefik(server.URL, false)
```

See [test files](../../traefik/*_test.go) for examples.
