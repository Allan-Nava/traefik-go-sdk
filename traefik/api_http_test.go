package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetHttpRouters tests GetHttpRouters method
func TestGetHttpRouters(t *testing.T) {
	tests := []struct {
		name       string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid routers response",
			serverResp: `{"router1":{"entryPoints":["web"]},"router2":{"entryPoints":["websecure"]}}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "empty routers",
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
			resp, err := client.GetHttpRouters()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetHttpRouters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetHttpRouters() returned nil response")
			}
		})
	}
}

// TestGetHttpRouter tests GetHttpRouter method
func TestGetHttpRouter(t *testing.T) {
	tests := []struct {
		name       string
		routerName string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid single router",
			routerName: "my-router",
			serverResp: `{"entryPoints":["web"],"service":"my-service"}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "router not found",
			routerName: "nonexistent",
			serverResp: `{"error":"not found"}`,
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
			resp, err := client.GetHttpRouter(tt.routerName)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetHttpRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetHttpRouter() returned nil response")
			}
		})
	}
}

// TestGetHttpServices tests GetHttpServices method
func TestGetHttpServices(t *testing.T) {
	tests := []struct {
		name       string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid services response",
			serverResp: `{"service1":{"loadBalancer":{}},"service2":{"loadBalancer":{}}}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "empty services",
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
			resp, err := client.GetHttpServices()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetHttpServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetHttpServices() returned nil response")
			}
		})
	}
}

// TestGetHttpMiddlewares tests GetHttpMiddlewares method
func TestGetHttpMiddlewares(t *testing.T) {
	tests := []struct {
		name       string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid middlewares response",
			serverResp: `{"auth-basic":{"basicAuth":{}},"retry":{"retry":{}}}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "empty middlewares",
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
			resp, err := client.GetHttpMiddlewares()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetHttpMiddlewares() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetHttpMiddlewares() returned nil response")
			}
		})
	}
}
