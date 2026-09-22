package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestUpdateHttpRouter tests updating an HTTP router
func TestUpdateHttpRouter(t *testing.T) {
	tests := []struct {
		name       string
		routerName string
		body       map[string]interface{}
		statusCode int
		wantErr    bool
	}{
		{
			name:       "successful update",
			routerName: "my-router",
			body: map[string]interface{}{
				"entryPoints": []string{"web", "websecure"},
				"service":     "my-service",
				"rule":        "Host(`example.com`)",
			},
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "update with middleware",
			routerName: "secure-router",
			body: map[string]interface{}{
				"entryPoints": []string{"websecure"},
				"service":     "secure-service",
				"rule":        "Host(`api.example.com`)",
				"middlewares": []string{"auth-basic", "rate-limit"},
			},
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "update non-existent router",
			routerName: "nonexistent",
			body: map[string]interface{}{
				"service": "service",
			},
			statusCode: http.StatusNotFound,
			wantErr:    false,
		},
		{
			name:       "update with invalid body",
			routerName: "bad-router",
			body:       nil,
			statusCode: http.StatusBadRequest,
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
				w.Write([]byte(`{"entryPoints":["web"]}`))
			}))
			defer server.Close()

			client, _ := BuildTraefik(server.URL, false)
			resp, err := client.UpdateHttpRouter(tt.routerName, tt.body)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateHttpRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("UpdateHttpRouter() returned nil response")
			}

			if resp != nil && resp.StatusCode() != tt.statusCode {
				t.Errorf("Expected status %d, got %d", tt.statusCode, resp.StatusCode())
			}
		})
	}
}

// TestCreateHttpRouter tests creating a new HTTP router
func TestCreateHttpRouter(t *testing.T) {
	tests := []struct {
		name       string
		routerName string
		body       map[string]interface{}
		statusCode int
		wantErr    bool
	}{
		{
			name:       "successful creation",
			routerName: "new-router",
			body: map[string]interface{}{
				"entryPoints": []string{"web"},
				"service":     "new-service",
				"rule":        "Host(`new.example.com`)",
			},
			statusCode: http.StatusCreated,
			wantErr:    false,
		},
		{
			name:       "creation with conflict",
			routerName: "existing-router",
			body: map[string]interface{}{
				"service": "service",
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
				w.Write([]byte(`{"entryPoints":["web"]}`))
			}))
			defer server.Close()

			client, _ := BuildTraefik(server.URL, false)
			resp, err := client.CreateHttpRouter(tt.routerName, tt.body)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateHttpRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("CreateHttpRouter() returned nil response")
			}
		})
	}
}

// TestDeleteHttpRouter tests deleting an HTTP router
func TestDeleteHttpRouter(t *testing.T) {
	tests := []struct {
		name       string
		routerName string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "successful deletion",
			routerName: "router-to-delete",
			statusCode: http.StatusNoContent,
			wantErr:    false,
		},
		{
			name:       "delete non-existent",
			routerName: "nonexistent",
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
			resp, err := client.DeleteHttpRouter(tt.routerName)

			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteHttpRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("DeleteHttpRouter() returned nil response")
			}
		})
	}
}

// TestUpdateHttpService tests updating an HTTP service
func TestUpdateHttpService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"loadBalancer":{"servers":[{"url":"http://localhost:8080"}]}}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	serviceBody := map[string]interface{}{
		"loadBalancer": map[string]interface{}{
			"servers": []map[string]interface{}{
				{"url": "http://localhost:8080"},
				{"url": "http://localhost:8081"},
			},
		},
	}

	resp, err := client.UpdateHttpService("my-service", serviceBody)

	if err != nil {
		t.Errorf("UpdateHttpService() error = %v", err)
	}

	if resp == nil {
		t.Errorf("UpdateHttpService() returned nil")
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode())
	}
}

// TestCreateHttpService tests creating an HTTP service
func TestCreateHttpService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"loadBalancer":{}}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	serviceBody := map[string]interface{}{
		"loadBalancer": map[string]interface{}{
			"servers": []map[string]interface{}{
				{"url": "http://backend:8080"},
			},
		},
	}

	resp, err := client.CreateHttpService("backend-service", serviceBody)

	if err != nil {
		t.Errorf("CreateHttpService() error = %v", err)
	}

	if resp == nil {
		t.Errorf("CreateHttpService() returned nil")
	}
}

// TestDeleteHttpService tests deleting an HTTP service
func TestDeleteHttpService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.DeleteHttpService("service-to-delete")

	if err != nil {
		t.Errorf("DeleteHttpService() error = %v", err)
	}

	if resp == nil {
		t.Errorf("DeleteHttpService() returned nil")
	}
}
