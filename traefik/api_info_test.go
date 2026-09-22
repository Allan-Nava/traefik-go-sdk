package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetApiOverview tests GetApiOverview method
func TestGetApiOverview(t *testing.T) {
	tests := []struct {
		name       string
		serverResp string
		wantErr    bool
	}{
		{
			name:       "valid response",
			serverResp: `{"providers":["docker"],"middlewares":["retry"]}`,
			wantErr:    false,
		},
		{
			name:       "empty response",
			serverResp: `{}`,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.serverResp))
			}))
			defer server.Close()

			client, _ := BuildTraefik(server.URL, false)
			resp, err := client.GetApiOverview()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetApiOverview() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetApiOverview() returned nil response")
			}
		})
	}
}

// TestGetApiVersion tests GetApiVersion method
func TestGetApiVersion(t *testing.T) {
	tests := []struct {
		name       string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid version response",
			serverResp: `{"Version":"v2.0","Codename":"Traefik"}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "not found (Traefik < 2.x)",
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
			resp, err := client.GetApiVersion()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetApiVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetApiVersion() returned nil response")
			}
		})
	}
}

// TestGetApiRawData tests GetApiRawData method
func TestGetApiRawData(t *testing.T) {
	tests := []struct {
		name       string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid raw data response",
			serverResp: `{"http":{"routers":{}},"tcp":{"routers":{}}}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "not found (Traefik < 2.x)",
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
			resp, err := client.GetApiRawData()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetApiRawData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetApiRawData() returned nil response")
			}
		})
	}
}
