# Quick Start

Get your first API call working in 5 minutes.

## Basic Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/Allan-Nava/traefik-go-sdk/traefik"
)

func main() {
	// Create a client
	client, err := traefik.BuildTraefik("http://localhost:8080", false)
	if err != nil {
		log.Fatal(err)
	}

	// Check health
	err = client.HealthCheck()
	if err != nil {
		log.Fatal("Traefik is not responding")
	}

	// Get all HTTP routers
	routers, err := client.GetHttpRouters()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got response with status: %d\n", routers.StatusCode())
	fmt.Printf("Response body: %s\n", routers.String())
}
```

## Run It

```bash
go run main.go
```

Expected output:
```
Got response with status: 200
Response body: {"router1":{"entryPoints":["web"]},...}
```

---

## Common Tasks

### Check Health

```go
client, _ := traefik.BuildTraefik("http://localhost:8080", false)
err := client.HealthCheck()
if err != nil {
	log.Fatal("Traefik down")
}
fmt.Println("Traefik is healthy")
```

### Get All HTTP Routers

```go
routers, err := client.GetHttpRouters()
if err != nil {
	log.Fatal(err)
}
// routers is a *resty.Response
// Parse JSON in your code
```

### Get Specific Router

```go
router, err := client.GetHttpRouter("my-router")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Router details: %s\n", router.String())
```

### Get HTTP Services

```go
services, err := client.GetHttpServices()
if err != nil {
	log.Fatal(err)
}
```

### Get HTTP Middlewares

```go
middlewares, err := client.GetHttpMiddlewares()
if err != nil {
	log.Fatal(err)
}
```

### TCP Endpoints

```go
// Get all TCP routers
tcpRouters, err := client.GetTcpRouters()

// Get specific TCP router
router, err := client.GetTcpRouter("tcp-router-1")

// Get TCP services
tcpServices, err := client.GetTcpServices()
```

### UDP Endpoints

```go
// Get all UDP routers
udpRouters, err := client.GetUdpRouters()

// Get specific UDP router
router, err := client.GetUdpRouter("udp-router-1")

// Get UDP services
udpServices, err := client.GetUdpServices()

// Get specific UDP service
service, err := client.GetUdpService("udp-service-1")
```

### Get API Info

```go
// API overview
overview, err := client.GetApiOverview()

// API version (Traefik 2.x+ required)
version, err := client.GetApiVersion()

// Raw configuration data (Traefik 2.x+ required)
rawData, err := client.GetApiRawData()
```

### Enable Debug Mode

```go
// Creates client with debug logging enabled
client, err := traefik.BuildTraefik("http://localhost:8080", true)
// Now all HTTP requests are logged to stdout
```

---

## Parsing Responses

All API methods return `(*resty.Response, error)`. Parse them as needed:

### As JSON

```go
import "encoding/json"

type Router struct {
	EntryPoints []string `json:"entryPoints"`
	Service     string   `json:"service"`
}

routers, _ := client.GetHttpRouters()
var data map[string]Router
json.Unmarshal(routers.Body(), &data)
```

### As String

```go
routers, _ := client.GetHttpRouters()
fmt.Println(routers.String()) // Pretty-printed
```

### As Bytes

```go
routers, _ := client.GetHttpRouters()
raw := routers.Body() // []byte
```

---

## Error Handling

```go
resp, err := client.GetHttpRouters()

if err != nil {
	// Network error
	log.Printf("Request failed: %v", err)
	return
}

if resp.StatusCode() >= 400 {
	// HTTP error
	log.Printf("Got status %d: %s", resp.StatusCode(), resp.String())
	return
}

// Success
fmt.Println("Got response:", resp.String())
```

---

## Next Steps

- [Configuration Guide](03-configuration.md) — Customize SDK
- [API Reference](../api/http.md) — All endpoints
- [Development Guide](../development/tdd.md) — Contributing

---

## Troubleshooting

**Q: How do I connect to a remote Traefik?**
```go
client, _ := traefik.BuildTraefik("http://traefik.example.com:8080", false)
```

**Q: How do I handle authentication?**
Currently, the SDK doesn't support authentication. Use a proxy or network policies.

**Q: Can I use HTTPS?**
```go
client, _ := traefik.BuildTraefik("https://localhost:8080", false)
```

See [Troubleshooting Guide](../troubleshooting.md) for more.
