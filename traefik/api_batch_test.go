package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestApplyConfiguration tests applying a complete configuration (batch operation)
func TestApplyConfiguration(t *testing.T) {
	tests := []struct {
		name       string
		config     interface{}
		statusCode int
		wantErr    bool
	}{
		{
			name: "successful configuration apply",
			config: map[string]interface{}{
				"http": map[string]interface{}{
					"routers": map[string]interface{}{
						"router1": map[string]interface{}{
							"entryPoints": []string{"web"},
							"service":     "service1",
							"rule":        "Host(`example.com`)",
						},
					},
					"services": map[string]interface{}{
						"service1": map[string]interface{}{
							"loadBalancer": map[string]interface{}{
								"servers": []map[string]interface{}{
									{"url": "http://localhost:8080"},
								},
							},
						},
					},
				},
			},
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "apply with tcp routers",
			config: map[string]interface{}{
				"tcp": map[string]interface{}{
					"routers": map[string]interface{}{
						"tcp-router": map[string]interface{}{
							"entryPoints": []string{"tcpep"},
							"service":     "tcp-service",
							"rule":        "HostSNI(`example.com`)",
						},
					},
					"services": map[string]interface{}{
						"tcp-service": map[string]interface{}{
							"loadBalancer": map[string]interface{}{
								"servers": []map[string]interface{}{
									{"address": "localhost:5000"},
								},
							},
						},
					},
				},
			},
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "apply with invalid config",
			config: map[string]interface{}{
				"invalid": "data",
			},
			statusCode: http.StatusBadRequest,
			wantErr:    false,
		},
		{
			name:       "apply with nil config",
			config:     nil,
			statusCode: http.StatusBadRequest,
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
			resp, err := client.ApplyConfiguration(tt.config)

			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("ApplyConfiguration() returned nil response")
			}
		})
	}
}

// TestGetConfiguration tests retrieving the complete configuration
func TestGetConfiguration(t *testing.T) {
	completeConfig := `{
		"http": {
			"routers": {
				"router1": {"entryPoints":["web"],"service":"service1"}
			},
			"services": {
				"service1": {"loadBalancer":{"servers":[{"url":"http://localhost:8080"}]}}
			}
		},
		"tcp": {
			"routers": {
				"tcp-router": {"entryPoints":["tcpep"],"service":"tcp-service"}
			},
			"services": {
				"tcp-service": {"loadBalancer":{"servers":[{"address":"localhost:5000"}]}}
			}
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(completeConfig))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.GetConfiguration()

	if err != nil {
		t.Errorf("GetConfiguration() error = %v", err)
	}

	if resp == nil {
		t.Errorf("GetConfiguration() returned nil")
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode())
	}
}

// TestValidateConfiguration tests validating a configuration before applying
func TestValidateConfiguration(t *testing.T) {
	tests := []struct {
		name       string
		config     interface{}
		statusCode int
		valid      bool
	}{
		{
			name: "valid configuration",
			config: map[string]interface{}{
				"http": map[string]interface{}{
					"routers": map[string]interface{}{
						"router1": map[string]interface{}{
							"entryPoints": []string{"web"},
							"service":     "service1",
							"rule":        "Host(`example.com`)",
						},
					},
					"services": map[string]interface{}{
						"service1": map[string]interface{}{
							"loadBalancer": map[string]interface{}{},
						},
					},
				},
			},
			statusCode: http.StatusOK,
			valid:      true,
		},
		{
			name: "invalid configuration",
			config: map[string]interface{}{
				"http": map[string]interface{}{
					"routers": map[string]interface{}{
						"router1": map[string]interface{}{
							// Missing required fields
						},
					},
				},
			},
			statusCode: http.StatusBadRequest,
			valid:      false,
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
				w.Write([]byte(`{"valid":true}`))
			}))
			defer server.Close()

			client, _ := BuildTraefik(server.URL, false)
			resp, err := client.ValidateConfiguration(tt.config)

			if err != nil {
				t.Errorf("ValidateConfiguration() error = %v", err)
			}

			if resp == nil {
				t.Errorf("ValidateConfiguration() returned nil")
			}

			if resp.StatusCode() != tt.statusCode {
				t.Errorf("Expected %d, got %d", tt.statusCode, resp.StatusCode())
			}
		})
	}
}

// TestResetConfiguration tests resetting configuration to default
func TestResetConfiguration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.ResetConfiguration()

	if err != nil {
		t.Errorf("ResetConfiguration() error = %v", err)
	}

	if resp == nil {
		t.Errorf("ResetConfiguration() returned nil")
	}

	if resp.StatusCode() != http.StatusNoContent {
		t.Errorf("Expected 204, got %d", resp.StatusCode())
	}
}
