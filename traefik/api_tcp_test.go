package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetTcpRouters tests GetTcpRouters method
func TestGetTcpRouters(t *testing.T) {
	tests := []struct {
		name       string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid TCP routers response",
			serverResp: `{"tcp-router1":{"entryPoints":["tcp-8000"]},"tcp-router2":{"entryPoints":["tcp-8001"]}}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "empty TCP routers",
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
			resp, err := client.GetTcpRouters()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetTcpRouters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetTcpRouters() returned nil response")
			}
		})
	}
}

// TestGetTcpRouter tests GetTcpRouter method
func TestGetTcpRouter(t *testing.T) {
	tests := []struct {
		name       string
		routerName string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid single TCP router",
			routerName: "tcp-router-1",
			serverResp: `{"entryPoints":["tcp-8000"],"service":"tcp-service"}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "TCP router not found",
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
			resp, err := client.GetTcpRouter(tt.routerName)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetTcpRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetTcpRouter() returned nil response")
			}
		})
	}
}

// TestGetTcpServices tests GetTcpServices method
func TestGetTcpServices(t *testing.T) {
	tests := []struct {
		name       string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid TCP services response",
			serverResp: `{"tcp-service1":{"loadBalancer":{}},"tcp-service2":{"loadBalancer":{}}}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "empty TCP services",
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
			resp, err := client.GetTcpServices()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetTcpServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetTcpServices() returned nil response")
			}
		})
	}
}
