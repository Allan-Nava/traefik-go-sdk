package traefik

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Advanced filtering test for HTTP routers by rule
func TestGetHttpRoutersByRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		rule := r.URL.Query().Get("rule")
		if rule != "Host(`example.com`)" {
			t.Errorf("Expected rule query param, got %s", rule)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"name":"router1","rule":"Host(example.com)"}]`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.GetHttpRoutersByRule("Host(`example.com`)")

	if err != nil {
		t.Errorf("GetHttpRoutersByRule() error = %v", err)
	}

	if resp == nil {
		t.Errorf("GetHttpRoutersByRule() returned nil")
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode())
	}
}

// Advanced filtering test for TCP routers by entrypoint
func TestGetTcpRoutersByEntryPoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		ep := r.URL.Query().Get("entryPoint")
		if ep != "tcpep" {
			t.Errorf("Expected entryPoint query param, got %s", ep)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"name":"tcp-router1","entryPoints":["tcpep"]}]`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.GetTcpRoutersByEntryPoint("tcpep")

	if err != nil {
		t.Errorf("GetTcpRoutersByEntryPoint() error = %v", err)
	}

	if resp == nil {
		t.Errorf("GetTcpRoutersByEntryPoint() returned nil")
	}
}

// Advanced filtering test for UDP routers by entrypoint
func TestGetUdpRoutersByEntryPoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		ep := r.URL.Query().Get("entryPoint")
		if ep != "udpep" {
			t.Errorf("Expected entryPoint query param, got %s", ep)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"name":"udp-router1","entryPoints":["udpep"]}]`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.GetUdpRoutersByEntryPoint("udpep")

	if err != nil {
		t.Errorf("GetUdpRoutersByEntryPoint() error = %v", err)
	}

	if resp == nil {
		t.Errorf("GetUdpRoutersByEntryPoint() returned nil")
	}
}

// Advanced filtering test for services by router
func TestGetServicesByRouter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		router := r.URL.Query().Get("router")
		if router != "my-router" {
			t.Errorf("Expected router query param, got %s", router)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"name":"my-service"}]`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.GetServicesByRouter("my-router")

	if err != nil {
		t.Errorf("GetServicesByRouter() error = %v", err)
	}

	if resp == nil {
		t.Errorf("GetServicesByRouter() returned nil")
	}
}

// Test middleware lookup by name/type
func TestGetHttpMiddlewareByType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		mwType := r.URL.Query().Get("type")
		if mwType != "basicAuth" {
			t.Errorf("Expected type query param, got %s", mwType)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"name":"auth1","type":"basicAuth"}]`))
	}))
	defer server.Close()

	client, _ := BuildTraefik(server.URL, false)
	resp, err := client.GetHttpMiddlewareByType("basicAuth")

	if err != nil {
		t.Errorf("GetHttpMiddlewareByType() error = %v", err)
	}

	if resp == nil {
		t.Errorf("GetHttpMiddlewareByType() returned nil")
	}
}
