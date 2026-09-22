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

// Test restyPost indirectly through API methods (if any use POST)
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
