# Installation

Get up and running with the Traefik Go SDK in minutes.

## Prerequisites

- **Go 1.18+** (1.20+ recommended)
- **Traefik instance** running with API enabled
- **Network access** to your Traefik API endpoint

## Install via `go get`

```bash
go get github.com/Allan-Nava/traefik-go-sdk
```

This downloads the latest version and adds it to your `go.mod`.

## Verify Installation

Create a simple test file:

```go
package main

import (
    "fmt"
    "github.com/Allan-Nava/traefik-go-sdk/traefik"
)

func main() {
    client, _ := traefik.BuildTraefik("http://localhost:8080", false)
    fmt.Println(client.IsDebug()) // false
    fmt.Println("SDK installed successfully!")
}
```

Run it:

```bash
go run main.go
# Output: false
# Output: SDK installed successfully!
```

---

## Setting Up Traefik API

The SDK communicates with Traefik's API endpoint (default: `http://localhost:8080/api`).

### Enable API in Traefik

**Docker Compose Example:**

```yaml
version: '3.8'
services:
  traefik:
    image: traefik:latest
    command:
      - "--api.insecure=true"          # Enable API (dev only!)
      - "--api.dashboard=true"          # Enable dashboard
      - "--entrypoints.web.address=:80"
    ports:
      - "80:80"
      - "8080:8080"                     # API port
```

**Kubernetes Example:**

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: traefik-config
data:
  traefik.yaml: |
    api:
      insecure: true        # dev only!
      dashboard: true
    entryPoints:
      web:
        address: ":80"
      websecure:
        address: ":443"
```

### Verify API is Running

```bash
curl http://localhost:8080/api/overview
# Should return JSON with Traefik info
```

---

## Next Steps

- [Quick Start Guide](02-quick-start.md) — First real example
- [Configuration](03-configuration.md) — Customize SDK behavior
- [API Reference](../api/http.md) — Explore all endpoints

---

## Troubleshooting

### "Connection refused" error

**Problem**: `dial tcp 127.0.0.1:8080: connect: connection refused`

**Solution**: 
1. Verify Traefik is running
2. Check the API port (default: 8080)
3. Verify firewall allows connection

### "API not enabled" error

**Problem**: `404 Not Found` when calling API

**Solution**:
1. Enable API in Traefik config (`--api.insecure=true`)
2. Restart Traefik
3. Verify endpoint: `curl http://localhost:8080/api/version`

### DNS resolution issues

**Problem**: `getaddrinfo: nodename nor servname provided`

**Solution**:
1. Use IP instead of hostname: `http://192.168.1.100:8080`
2. Check DNS configuration
3. Try `localhost` instead of `127.0.0.1`

See [Troubleshooting Guide](../troubleshooting.md) for more help.
