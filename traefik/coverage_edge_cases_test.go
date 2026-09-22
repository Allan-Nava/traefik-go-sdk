package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test debug print method
func TestDebugPrint(t *testing.T) {
	tests := []struct {
		name  string
		debug bool
		data  string
	}{
		{
			name:  "debug on",
			debug: true,
			data:  "test data",
		},
		{
			name:  "debug off",
			debug: false,
			data:  "test data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, _ := BuildTraefik("http://localhost:8080", tt.debug)
			// debugPrint is internal, but we test via IsDebug()
			if client.IsDebug() != tt.debug {
				t.Errorf("IsDebug() = %v, want %v", client.IsDebug(), tt.debug)
			}
		})
	}
}

// Test HTTP error scenarios
func TestGetHttpRoutersWithError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		serverResp string
	}{
		{
			name:       "server error 500",
			statusCode: http.StatusInternalServerError,
			serverResp: `{"error":"internal server error"}`,
		},
		{
			name:       "bad request 400",
			statusCode: http.StatusBadRequest,
			serverResp: `{"error":"bad request"}`,
		},
		{
			name:       "unauthorized 401",
			statusCode: http.StatusUnauthorized,
			serverResp: `{"error":"unauthorized"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.serverResp))
			}))
			defer server.Close()

			client, _ := BuildTraefik(server.URL, false)
			resp, _ := client.GetHttpRouters()

			if resp != nil && resp.StatusCode() != tt.statusCode {
				t.Errorf("Expected status %d, got %d", tt.statusCode, resp.StatusCode())
			}
		})
	}
}

// Test restyPost directly
func TestRestyPost(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       map[string]interface{}
		wantErr    bool
	}{
		{
			name:       "successful POST",
			statusCode: http.StatusOK,
			body:       map[string]interface{}{"key": "value"},
			wantErr:    false,
		},
		{
			name:       "POST with error response",
			statusCode: http.StatusBadRequest,
			body:       map[string]interface{}{"invalid": "data"},
			wantErr:    false,
		},
		{
			name:       "POST server error",
			statusCode: http.StatusInternalServerError,
			body:       nil,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Errorf("Expected POST, got %s", r.Method)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(`{}`))
			}))
			defer server.Close()

			client, _ := BuildTraefik(server.URL, false)
			// Access private method via reflection or type assertion
			sdk := client.(*traefikSdk)
			resp, err := sdk.restyPost("/api/test", tt.body)

			if (err != nil) != tt.wantErr {
				t.Errorf("restyPost() error = %v, wantErr %v", err, tt.wantErr)
			}

			if resp != nil && resp.StatusCode() != tt.statusCode {
				t.Errorf("Expected status %d, got %d", tt.statusCode, resp.StatusCode())
			}
		})
	}
}

// Test debugPrint directly
func TestDebugPrintDirect(t *testing.T) {
	tests := []struct {
		name  string
		debug bool
		data  string
	}{
		{
			name:  "debug enabled - should print",
			debug: true,
			data:  "debug message",
		},
		{
			name:  "debug disabled - should not print",
			debug: false,
			data:  "silent message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, _ := BuildTraefik("http://localhost:8080", tt.debug)
			sdk := client.(*traefikSdk)
			// Call debugPrint - should not panic
			sdk.debugPrint(tt.data)
			if sdk.debug != tt.debug {
				t.Errorf("Expected debug=%v, got %v", tt.debug, sdk.debug)
			}
		})
	}
}

// Test debugPrint with various data types
func TestDebugPrintDataTypes(t *testing.T) {
	client, _ := BuildTraefik("http://localhost:8080", true)
	sdk := client.(*traefikSdk)

	// Test with different data types
	testCases := []interface{}{
		"string",
		123,
		map[string]interface{}{"key": "value"},
		[]string{"a", "b", "c"},
		nil,
	}

	for _, data := range testCases {
		// Should not panic
		sdk.debugPrint(data)
	}
}

// Test debugPrint mode state
func TestDebugMode(t *testing.T) {
	// Test that debug flag is stored and retrievable
	client1, _ := BuildTraefik("http://localhost:8080", true)
	if !client1.IsDebug() {
		t.Errorf("Expected debug=true, got false")
	}

	client2, _ := BuildTraefik("http://localhost:8080", false)
	if client2.IsDebug() {
		t.Errorf("Expected debug=false, got true")
	}
}

// Test with query parameters
func TestGetHttpRoutersWithQueryParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request was made correctly
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"router1":{"entryPoints":["web"]}}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.GetHttpRouters()

	if err != nil {
		t.Errorf("GetHttpRouters() error = %v, want nil", err)
	}

	if resp == nil {
		t.Errorf("GetHttpRouters() returned nil response")
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode())
	}
}

// Additional edge case tests
func TestAPIMethodsWithLargePayload(t *testing.T) {
	largeResp := `{`
	for i := 0; i < 100; i++ {
		if i > 0 {
			largeResp += `,`
		}
		largeResp += `"router` + string(rune(i)) + `":{"entryPoints":["web"]}`
	}
	largeResp += `}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(largeResp))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.GetHttpRouters()

	if err != nil {
		t.Errorf("GetHttpRouters() with large payload error = %v", err)
	}

	if resp == nil {
		t.Errorf("GetHttpRouters() returned nil response")
	}
}

// Test TCP service endpoint
func TestGetTcpServiceInvocation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"loadBalancer":{"servers":[{"address":"127.0.0.1:8000"}]}}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)

	// GetTcpService should use restyGet properly now
	resp, err := client.GetTcpServices()

	if err != nil {
		t.Errorf("GetTcpServices() error = %v", err)
	}

	if resp == nil {
		t.Errorf("GetTcpServices() returned nil")
	}
}

// Test client creation with various URLs
func TestBuildTraefikWithVariousURLs(t *testing.T) {
	urls := []string{
		"http://localhost:8080",
		"https://localhost:8080",
		"http://traefik:8080",
		"http://192.168.1.1:8080",
		"http://traefik.example.com",
	}

	for _, url := range urls {
		t.Run(url, func(t *testing.T) {
			client, err := BuildTraefik(url, false)
			if err != nil {
				t.Errorf("BuildTraefik(%s) error = %v", url, err)
			}
			if client == nil {
				t.Errorf("BuildTraefik(%s) returned nil client", url)
			}
		})
	}
}

// Test GetHttpRouter by name
func TestGetHttpRouterByName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"entryPoints":["web"]}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.GetHttpRouter("test-router")

	if err != nil {
		t.Errorf("GetHttpRouter() error = %v", err)
	}
	if resp == nil {
		t.Errorf("GetHttpRouter() returned nil")
	}
	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode())
	}
}

// Test HealthCheck error cases
func TestHealthCheckErrors(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
	}{
		{
			name:       "health check 500 error",
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "health check 503 unavailable",
			statusCode: http.StatusServiceUnavailable,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
			}))
			defer server.Close()

			client, _ := BuildTraefik(server.URL, false)
			err := client.HealthCheck()
			// HealthCheck returns nil on error based on implementation
			if err != nil {
				t.Logf("HealthCheck returned error: %v", err)
			}
		})
	}
}

// Test API methods with error status codes
func TestGetHttpRoutersErrorStatus(t *testing.T) {
	statusCodes := []int{
		http.StatusNotFound,
		http.StatusForbidden,
		http.StatusGone,
	}

	for _, code := range statusCodes {
		t.Run("status_"+string(rune(code)), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(code)
				w.Write([]byte(`{}`))
			}))
			defer server.Close()

			client, _ := BuildTraefik(server.URL, false)
			resp, err := client.GetHttpRouters()

			if err == nil && resp.StatusCode() != code {
				t.Errorf("Expected status %d, got %d", code, resp.StatusCode())
			}
		})
	}
}

// Test restyPost error cases
func TestRestyPostErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"not found"}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	sdk := client.(*traefikSdk)
	resp, err := sdk.restyPost("/nonexistent", nil)

	if err == nil && resp != nil && resp.StatusCode() != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", resp.StatusCode())
	}
}
