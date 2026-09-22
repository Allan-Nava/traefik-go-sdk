package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test exporting configuration to JSON
func TestExportConfiguration(t *testing.T) {
	expectedConfig := `{"http":{"routers":{"router1":{"rule":"Host(example.com)","service":"service1"}},"services":{"service1":{"loadBalancer":{"servers":[{"url":"http://localhost:8080"}]}}}},"tcp":{"routers":{},"services":{}}}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedConfig))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.ExportConfiguration()

	if err != nil {
		t.Errorf("ExportConfiguration() error = %v", err)
	}

	if resp == nil {
		t.Errorf("ExportConfiguration() returned nil")
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode())
	}
}

// Test importing configuration from JSON
func TestImportConfiguration(t *testing.T) {
	configToImport := map[string]interface{}{
		"http": map[string]interface{}{
			"routers": map[string]interface{}{
				"router1": map[string]interface{}{
					"rule":    "Host(`example.com`)",
					"service": "service1",
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
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status":"imported"}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.ImportConfiguration(configToImport)

	if err != nil {
		t.Errorf("ImportConfiguration() error = %v", err)
	}

	if resp == nil {
		t.Errorf("ImportConfiguration() returned nil")
	}

	if resp.StatusCode() != http.StatusCreated {
		t.Errorf("Expected 201, got %d", resp.StatusCode())
	}
}

// Test creating configuration backup
func TestBackupConfiguration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		timestamp := r.URL.Query().Get("timestamp")
		if timestamp == "" {
			t.Errorf("Expected timestamp query param")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Disposition", "attachment; filename=traefik-backup-"+timestamp+".json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.BackupConfiguration("2026-09-22T15:35:00Z")

	if err != nil {
		t.Errorf("BackupConfiguration() error = %v", err)
	}

	if resp == nil {
		t.Errorf("BackupConfiguration() returned nil")
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode())
	}
}

// Test restoring configuration from backup
func TestRestoreConfiguration(t *testing.T) {
	backup := map[string]interface{}{
		"version": "2.0.0",
		"http": map[string]interface{}{
			"routers": map[string]interface{}{},
			"services": map[string]interface{}{},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"restored"}`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.RestoreConfiguration(backup)

	if err != nil {
		t.Errorf("RestoreConfiguration() error = %v", err)
	}

	if resp == nil {
		t.Errorf("RestoreConfiguration() returned nil")
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode())
	}
}
