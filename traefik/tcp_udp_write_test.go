package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TCP Write Operations Tests

func TestCreateTcpRouter(t *testing.T) {
	tests := []struct {
		name       string
		routerName string
		config     interface{}
		statusCode int
		wantErr    bool
	}{
		{
			name:       "successful TCP router creation",
			routerName: "tcp-router",
			config: map[string]interface{}{
				"entryPoints": []string{"tcpep"},
				"service":     "tcp-service",
				"rule":        "HostSNI(`example.com`)",
			},
			statusCode: http.StatusCreated,
			wantErr:    false,
		},
		{
			name:       "TCP router creation with conflict",
			routerName: "tcp-router",
			config: map[string]interface{}{
				"entryPoints": []string{"tcpep"},
				"service":     "tcp-service",
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
			resp, err := client.CreateTcpRouter(tt.routerName, tt.config)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTcpRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp != nil && resp.StatusCode() != tt.statusCode {
				t.Errorf("Expected status %d, got %d", tt.statusCode, resp.StatusCode())
			}
		})
	}
}

func TestUpdateTcpRouter(t *testing.T) {
	tests := []struct {
		name       string
		routerName string
		config     interface{}
		statusCode int
		wantErr    bool
	}{
		{
			name:       "successful TCP router update",
			routerName: "tcp-router",
			config: map[string]interface{}{
				"entryPoints": []string{"tcpep", "tcpep2"},
				"service":     "tcp-service-2",
				"rule":        "HostSNI(`example.com`)",
			},
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "update non-existent TCP router",
			routerName: "non-existent",
			config: map[string]interface{}{
				"service": "tcp-service",
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
			resp, err := client.UpdateTcpRouter(tt.routerName, tt.config)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTcpRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp != nil && resp.StatusCode() != tt.statusCode {
				t.Errorf("Expected status %d, got %d", tt.statusCode, resp.StatusCode())
			}
		})
	}
}

func TestDeleteTcpRouter(t *testing.T) {
	tests := []struct {
		name       string
		routerName string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "successful TCP router deletion",
			routerName: "tcp-router",
			statusCode: http.StatusNoContent,
			wantErr:    false,
		},
		{
			name:       "delete non-existent TCP router",
			routerName: "non-existent",
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
			resp, err := client.DeleteTcpRouter(tt.routerName)

			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTcpRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp != nil && resp.StatusCode() != tt.statusCode {
				t.Errorf("Expected status %d, got %d", tt.statusCode, resp.StatusCode())
			}
		})
	}
}

// TCP Service Write Operations

func TestCreateTcpService(t *testing.T) {
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
		"loadBalancer": map[string]interface{}{
			"servers": []map[string]interface{}{
				{"address": "localhost:5000"},
			},
		},
	}
	resp, err := client.CreateTcpService("tcp-service", config)

	if err != nil {
		t.Errorf("CreateTcpService() error = %v", err)
	}

	if resp == nil {
		t.Errorf("CreateTcpService() returned nil")
	}
}

func TestUpdateTcpService(t *testing.T) {
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
		"loadBalancer": map[string]interface{}{
			"servers": []map[string]interface{}{
				{"address": "localhost:5001"},
				{"address": "localhost:5002"},
			},
		},
	}
	resp, err := client.UpdateTcpService("tcp-service", config)

	if err != nil {
		t.Errorf("UpdateTcpService() error = %v", err)
	}

	if resp == nil {
		t.Errorf("UpdateTcpService() returned nil")
	}
}

func TestDeleteTcpService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.DeleteTcpService("tcp-service")

	if err != nil {
		t.Errorf("DeleteTcpService() error = %v", err)
	}

	if resp == nil {
		t.Errorf("DeleteTcpService() returned nil")
	}
}

// UDP Write Operations Tests

func TestCreateUdpRouter(t *testing.T) {
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
		"entryPoints": []string{"udpep"},
		"service":     "udp-service",
	}
	resp, err := client.CreateUdpRouter("udp-router", config)

	if err != nil {
		t.Errorf("CreateUdpRouter() error = %v", err)
	}

	if resp == nil {
		t.Errorf("CreateUdpRouter() returned nil")
	}
}

func TestUpdateUdpRouter(t *testing.T) {
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
		"entryPoints": []string{"udpep"},
		"service":     "udp-service-2",
	}
	resp, err := client.UpdateUdpRouter("udp-router", config)

	if err != nil {
		t.Errorf("UpdateUdpRouter() error = %v", err)
	}

	if resp == nil {
		t.Errorf("UpdateUdpRouter() returned nil")
	}
}

func TestDeleteUdpRouter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.DeleteUdpRouter("udp-router")

	if err != nil {
		t.Errorf("DeleteUdpRouter() error = %v", err)
	}

	if resp == nil {
		t.Errorf("DeleteUdpRouter() returned nil")
	}
}

func TestCreateUdpService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	config := map[string]interface{}{
		"loadBalancer": map[string]interface{}{
			"servers": []map[string]interface{}{
				{"address": "localhost:5000"},
			},
		},
	}
	resp, err := client.CreateUdpService("udp-service", config)

	if err != nil {
		t.Errorf("CreateUdpService() error = %v", err)
	}

	if resp == nil {
		t.Errorf("CreateUdpService() returned nil")
	}
}

func TestUpdateUdpService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	config := map[string]interface{}{
		"loadBalancer": map[string]interface{}{
			"servers": []map[string]interface{}{
				{"address": "localhost:5001"},
			},
		},
	}
	resp, err := client.UpdateUdpService("udp-service", config)

	if err != nil {
		t.Errorf("UpdateUdpService() error = %v", err)
	}

	if resp == nil {
		t.Errorf("UpdateUdpService() returned nil")
	}
}

func TestDeleteUdpService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.DeleteUdpService("udp-service")

	if err != nil {
		t.Errorf("DeleteUdpService() error = %v", err)
	}

	if resp == nil {
		t.Errorf("DeleteUdpService() returned nil")
	}
}
