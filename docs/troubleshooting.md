# Troubleshooting Guide

## Common Issues

### Connection Refused

```
dial tcp 127.0.0.1:8080: connect: connection refused
```

**Cause**: Traefik is not running or API port is wrong.

**Solution**:
```bash
# Check if Traefik is running
docker ps | grep traefik

# Verify API port
curl http://localhost:8080/api/overview

# Check if firewall is blocking
sudo ufw allow 8080/tcp  # Linux
```

---

### 404 Not Found

```
HTTP/1.1 404 Not Found
```

**Cause**: API is not enabled in Traefik.

**Solution**:
```yaml
# In traefik config
api:
  insecure: true  # Enable API
  dashboard: true
```

Then restart Traefik:
```bash
docker restart traefik
```

---

### DNS Resolution Failed

```
getaddrinfo: nodename nor servname provided, or not known
```

**Cause**: Hostname cannot be resolved.

**Solution**:
```go
// Try with IP instead of hostname
client, _ := traefik.BuildTraefik("http://192.168.1.100:8080", false)

// Or use localhost
client, _ := traefik.BuildTraefik("http://localhost:8080", false)
```

---

### Response Parsing Error

```
json: cannot unmarshal object into Go value of type []string
```

**Cause**: Response format doesn't match expected structure.

**Solution**:
```go
// Check actual response format
routers, _ := client.GetHttpRouters()
fmt.Println(routers.String())  // Print raw JSON

// Then parse accordingly
var data map[string]interface{}
json.Unmarshal(routers.Body(), &data)
```

---

### Timeout

```
context deadline exceeded
```

**Cause**: Traefik API is slow or unreachable.

**Solution**:
1. Check network connectivity
2. Check Traefik is responding: `curl http://localhost:8080/api/overview`
3. Check for high CPU/memory on Traefik
4. Enable debug mode to see request timing

---

### HTTPS Certificate Error

```
x509: certificate signed by unknown authority
```

**Cause**: Self-signed certificate not in system CA store.

**Solution**:
```bash
# Add certificate to system trust store
# macOS:
security add-trusted-cert -d -r trustRoot -k ~/Library/Keychains/login.keychain cert.pem

# Linux:
sudo cp cert.pem /usr/local/share/ca-certificates/
sudo update-ca-certificates

# Then use HTTPS:
client, _ := traefik.BuildTraefik("https://localhost:8080", false)
```

---

## Debug Tips

### Enable Debug Logging

```go
client, _ := traefik.BuildTraefik("http://localhost:8080", true)
// Now you'll see all HTTP requests/responses
```

### Print Full Response

```go
resp, _ := client.GetHttpRouters()
fmt.Printf("Status: %d\n", resp.StatusCode())
fmt.Printf("Headers: %v\n", resp.Header())
fmt.Printf("Body: %s\n", resp.String())
```

### Check Traefik API Directly

```bash
# Health check
curl http://localhost:8080/api/overview

# Get version
curl http://localhost:8080/api/version

# Get routers
curl http://localhost:8080/api/http/routers
```

---

## Performance Issues

**SDK is slow?**

1. Check network latency:
```bash
ping traefik.example.com
```

2. Check Traefik performance:
```bash
# View Traefik metrics
curl http://localhost:8080/metrics
```

3. Enable caching in your application:
```go
// Cache responses for 5 seconds
type CachedClient struct {
	client traefik.ITraefikClient
	cache  map[string]CacheEntry
}
```

---

## Still Having Issues?

1. **Check [Docs](guides/01-installation.md)** — Most common issues are covered
2. **Enable [Debug Mode](#debug-tips)** — See exactly what's happening
3. **Open [GitHub Issue](https://github.com/Allan-Nava/traefik-go-sdk/issues)** — Include:
   - Your Go version (`go version`)
   - Traefik version (`curl http://localhost:8080/api/version`)
   - Full error message
   - Minimal reproduction code

---

## Contact & Support

- 📖 **Documentation**: Start with [Installation](guides/01-installation.md)
- 🐛 **Report Bugs**: [GitHub Issues](https://github.com/Allan-Nava/traefik-go-sdk/issues)
- 💬 **Ask Questions**: [GitHub Discussions](https://github.com/Allan-Nava/traefik-go-sdk/discussions)
- 📚 **API Docs**: [Traefik Official](https://doc.traefik.io/traefik/operations/api/)
