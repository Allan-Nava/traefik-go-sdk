# Configuration

Configure the Traefik Go SDK for your environment.

## BuildTraefik() Options

The SDK is configured via the `BuildTraefik()` builder:

```go
client, err := traefik.BuildTraefik(
    "http://localhost:8080",  // Traefik API URL
    false,                      // Debug mode (optional)
)
```

### Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `url` | string | - | Traefik API base URL (e.g., `http://localhost:8080`) |
| `debug` | bool | false | Enable HTTP request/response logging |

### Example: Production Config

```go
const traefikURL = "https://traefik.prod.example.com:8080"
client, err := traefik.BuildTraefik(traefikURL, false)
if err != nil {
    log.Fatal(err)
}
```

### Example: Development Config

```go
const traefikURL = "http://localhost:8080"
client, err := traefik.BuildTraefik(traefikURL, true) // Debug enabled
if err != nil {
    log.Fatal(err)
}
```

---

## Debug Mode

Enable debug mode to log all HTTP requests/responses:

```go
client, _ := traefik.BuildTraefik("http://localhost:8080", true)

// Now all requests print to stdout:
// => GET /api/http/routers HTTP/1.1
// => Host: localhost:8080
// <= HTTP/1.1 200 OK
// <= Content-Type: application/json
```

Check debug status:

```go
if client.IsDebug() {
    fmt.Println("Debug mode is ON")
}
```

---

## URL Configuration

### Local Development

```go
client, _ := traefik.BuildTraefik("http://localhost:8080", false)
```

### Remote Server

```go
client, _ := traefik.BuildTraefik("http://traefik.example.com:8080", false)
```

### HTTPS with Self-Signed Certificate

```go
client, _ := traefik.BuildTraefik("https://localhost:8080", false)
```

**Note**: The SDK uses the Go `net/http` client, which respects system CA certificates.

---

## Environment Variables

The SDK doesn't use environment variables directly, but you can:

```go
import "os"

url := os.Getenv("TRAEFIK_URL")
if url == "" {
    url = "http://localhost:8080"
}

client, _ := traefik.BuildTraefik(url, false)
```

---

## Timeout & Retry

The SDK uses the `resty` client with default timeouts. For custom configuration:

```go
// The SDK uses resty internally
// To customize timeouts, modify the SDK or use a proxy
client, _ := traefik.BuildTraefik("http://localhost:8080", false)
```

**Note**: Custom timeout support is on the [roadmap](../architecture/roadmap.md).

---

## Best Practices

✅ **Use environment variables for URLs:**
```go
url := os.Getenv("TRAEFIK_API_URL")
client, _ := traefik.BuildTraefik(url, false)
```

✅ **Enable debug only in development:**
```go
debug := os.Getenv("DEBUG") == "true"
client, _ := traefik.BuildTraefik(url, debug)
```

✅ **Validate URL before using:**
```go
if url == "" {
    log.Fatal("TRAEFIK_URL not set")
}
```

❌ **Don't hardcode URLs in production:**
```go
// BAD
client, _ := traefik.BuildTraefik("http://prod-traefik:8080", false)

// GOOD
client, _ := traefik.BuildTraefik(os.Getenv("TRAEFIK_API_URL"), false)
```

---

## Troubleshooting

**Q: Connection refused**
- Check Traefik is running on the correct host/port
- Verify firewall rules

**Q: 404 Not Found**
- Verify API is enabled in Traefik (`--api.insecure=true`)
- Check you're using the correct base URL

**Q: Slow responses**
- Enable debug mode to see timing
- Check network connectivity
- See [Roadmap](../architecture/roadmap.md) for context.Context support

See [Troubleshooting](../troubleshooting.md) for more help.
