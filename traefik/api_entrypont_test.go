package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetEntrypoints tests GetEntrypoints method
func TestGetEntrypoints(t *testing.T) {
	tests := []struct {
		name       string
		serverResp string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "valid entrypoints response",
			serverResp: `{"web":{"address":":80"},"websecure":{"address":":443"}}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "empty entrypoints",
			serverResp: `{}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "single entrypoint",
			serverResp: `{"web":{"address":":80","forwardedHeaders":{"insecure":false}}}`,
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
			resp, err := client.GetEntrypoints()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetEntrypoints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("GetEntrypoints() returned nil response")
			}
		})
	}
}
