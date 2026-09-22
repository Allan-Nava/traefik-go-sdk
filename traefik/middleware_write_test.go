package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateHttpMiddleware(t *testing.T) {
	tests := []struct {
		name       string
		middlewareName string
		config     interface{}
		statusCode int
		wantErr    bool
	}{
		{
			name:       "successful HTTP middleware creation",
			middlewareName: "auth-middleware",
			config: map[string]interface{}{
				"basicAuth": map[string]interface{}{
					"users": []string{"user:password"},
				},
			},
			statusCode: http.StatusCreated,
			wantErr:    false,
		},
		{
			name:       "middleware creation with conflict",
			middlewareName: "auth-middleware",
			config: map[string]interface{}{
				"basicAuth": map[string]interface{}{},
			},
			statusCode: http.StatusConflict,
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
			resp, err := client.CreateHttpMiddleware(tt.middlewareName, tt.config)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateHttpMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp != nil && resp.StatusCode() != tt.statusCode {
				t.Errorf("Expected status %d, got %d", tt.statusCode, resp.StatusCode())
			}
		})
	}
}

func TestUpdateHttpMiddleware(t *testing.T) {
	tests := []struct {
		name       string
		middlewareName string
		config     interface{}
		statusCode int
		wantErr    bool
	}{
		{
			name:       "successful HTTP middleware update",
			middlewareName: "auth-middleware",
			config: map[string]interface{}{
				"basicAuth": map[string]interface{}{
					"users": []string{"user1:password1", "user2:password2"},
				},
			},
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "update non-existent middleware",
			middlewareName: "non-existent",
			config: map[string]interface{}{
				"basicAuth": map[string]interface{}{},
			},
			statusCode: http.StatusNotFound,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "PUT" {
					t.Errorf("Expected PUT, got %s", r.Method)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(`{}`))
			}))
			defer server.Close()

			client, _ := BuildTraefik(server.URL, false)
			resp, err := client.UpdateHttpMiddleware(tt.middlewareName, tt.config)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateHttpMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp != nil && resp.StatusCode() != tt.statusCode {
				t.Errorf("Expected status %d, got %d", tt.statusCode, resp.StatusCode())
			}
		})
	}
}

func TestDeleteHttpMiddleware(t *testing.T) {
	tests := []struct {
		name       string
		middlewareName string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "successful HTTP middleware deletion",
			middlewareName: "auth-middleware",
			statusCode: http.StatusNoContent,
			wantErr:    false,
		},
		{
			name:       "delete non-existent middleware",
			middlewareName: "non-existent",
			statusCode: http.StatusNotFound,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "DELETE" {
					t.Errorf("Expected DELETE, got %s", r.Method)
				}
				w.WriteHeader(tt.statusCode)
			}))
			defer server.Close()

			client, _ := BuildTraefik(server.URL, false)
			resp, err := client.DeleteHttpMiddleware(tt.middlewareName)

			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteHttpMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp != nil && resp.StatusCode() != tt.statusCode {
				t.Errorf("Expected status %d, got %d", tt.statusCode, resp.StatusCode())
			}
		})
	}
}

// TCP Middleware Write Operations

func TestCreateTcpMiddleware(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	config := map[string]interface{}{
		"ipWhiteList": map[string]interface{}{
			"sourceRange": []string{"10.0.0.0/8"},
		},
	}
	resp, err := client.CreateTcpMiddleware("tcp-middleware", config)

	if err != nil {
		t.Errorf("CreateTcpMiddleware() error = %v", err)
	}

	if resp == nil {
		t.Errorf("CreateTcpMiddleware() returned nil")
	}
}

func TestUpdateTcpMiddleware(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	config := map[string]interface{}{
		"ipWhiteList": map[string]interface{}{
			"sourceRange": []string{"10.0.0.0/8", "172.16.0.0/12"},
		},
	}
	resp, err := client.UpdateTcpMiddleware("tcp-middleware", config)

	if err != nil {
		t.Errorf("UpdateTcpMiddleware() error = %v", err)
	}

	if resp == nil {
		t.Errorf("UpdateTcpMiddleware() returned nil")
	}
}

func TestDeleteTcpMiddleware(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.DeleteTcpMiddleware("tcp-middleware")

	if err != nil {
		t.Errorf("DeleteTcpMiddleware() error = %v", err)
	}

	if resp == nil {
		t.Errorf("DeleteTcpMiddleware() returned nil")
	}
}
