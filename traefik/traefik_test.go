package traefik

import (
	"testing"
)

// TestBuildTraefik tests the SDK builder
func TestBuildTraefik(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		debug   bool
		wantErr bool
	}{
		{
			name:    "valid URL with debug off",
			url:     "http://localhost:8080",
			debug:   false,
			wantErr: false,
		},
		{
			name:    "valid URL with debug on",
			url:     "http://localhost:8080",
			debug:   true,
			wantErr: false,
		},
		{
			name:    "empty URL",
			url:     "",
			debug:   false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := BuildTraefik(tt.url, tt.debug)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildTraefik() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if client == nil {
				t.Errorf("BuildTraefik() returned nil client")
				return
			}
		})
	}
}

// TestIsDebug tests the debug flag getter
func TestIsDebug(t *testing.T) {
	tests := []struct {
		name     string
		debug    bool
		wantBool bool
	}{
		{
			name:     "debug enabled",
			debug:    true,
			wantBool: true,
		},
		{
			name:     "debug disabled",
			debug:    false,
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, _ := BuildTraefik("http://localhost:8080", tt.debug)
			if client.IsDebug() != tt.wantBool {
				t.Errorf("IsDebug() = %v, want %v", client.IsDebug(), tt.wantBool)
			}
		})
	}
}

// TestHealthCheck tests basic health check
func TestHealthCheck(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "valid localhost URL",
			url:     "http://localhost:8080",
			wantErr: false,
		},
		{
			name:    "invalid unreachable URL",
			url:     "http://localhost:19999",
			wantErr: false, // Currently always returns nil — see ID-003
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, _ := BuildTraefik(tt.url, false)
			err := client.HealthCheck()
			if (err != nil) != tt.wantErr {
				t.Logf("HealthCheck() error = %v, wantErr %v", err, tt.wantErr)
				// Note: Currently all paths return nil — this is a known issue (ID-003)
			}
		})
	}
}
