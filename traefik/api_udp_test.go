package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetUdpRouters tests GetUdpRouters method
func TestGetUdpRouters(t *testing.T) {
	tests := []struct {
		name       string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid UDP routers response",
			serverResp: `{"udp-router1":{"entryPoints":["udp-5000"]},"udp-router2":{"entryPoints":["udp-5001"]}}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "empty UDP routers",
			serverResp: `{}`,
			statusCode: http.StatusOK,
			wantErr:    false,
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
			resp, err := client.GetUdpRouters()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUdpRouters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetUdpRouters() returned nil response")
			}
		})
	}
}

// TestGetUdpRouter tests GetUdpRouter method
func TestGetUdpRouter(t *testing.T) {
	tests := []struct {
		name       string
		routerName string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid single UDP router",
			routerName: "udp-router-1",
			serverResp: `{"entryPoints":["udp-5000"],"service":"udp-service"}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "UDP router not found",
			routerName: "nonexistent",
			serverResp: `{}`,
			statusCode: http.StatusNotFound,
			wantErr:    false,
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
			resp, err := client.GetUdpRouter(tt.routerName)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUdpRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetUdpRouter() returned nil response")
			}
		})
	}
}

// TestGetUdpServices tests GetUdpServices method
func TestGetUdpServices(t *testing.T) {
	tests := []struct {
		name       string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid UDP services response",
			serverResp: `{"udp-service1":{"loadBalancer":{}},"udp-service2":{"loadBalancer":{}}}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "empty UDP services",
			serverResp: `{}`,
			statusCode: http.StatusOK,
			wantErr:    false,
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
			resp, err := client.GetUdpServices()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUdpServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetUdpServices() returned nil response")
			}
		})
	}
}

// TestGetUdpService tests GetUdpService method (note: currently all UDP methods return same response due to bug)
// See ID-003 in backlog.md
func TestGetUdpService(t *testing.T) {
	tests := []struct {
		name        string
		serviceName string
		serverResp  string
		statusCode  int
		wantErr     bool
	}{
		{
			name:        "valid UDP services list (currently returns all)",
			serviceName: "udp-service-1",
			serverResp:  `{"udp-service1":{"loadBalancer":{"servers":[{"address":"192.168.1.1:5000"}]}}}`,
			statusCode:  http.StatusOK,
			wantErr:     false,
		},
		{
			name:        "empty UDP services",
			serviceName: "nonexistent",
			serverResp:  `{}`,
			statusCode:  http.StatusOK,
			wantErr:     false,
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
			resp, err := client.GetUdpService(tt.serviceName)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUdpService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetUdpService() returned nil response")
			}
		})
	}
}
