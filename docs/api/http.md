# HTTP API Reference

HTTP routers, services, and middlewares.

## Routers

### GetHttpRouters()

Lists all HTTP routers.

```go
routers, err := client.GetHttpRouters()
if err != nil {
    log.Fatal(err)
}
fmt.Println(routers.String())
```

**Response**:
```json
{
  "router1": {"entryPoints": ["web"], "service": "service1"},
  "router2": {"entryPoints": ["websecure"], "service": "service2"}
}
```

---

### GetHttpRouter(name)

Get a specific HTTP router.

```go
router, err := client.GetHttpRouter("my-router")
```

**Parameters**:
- `name` (string): Router name

---

## Services

### GetHttpServices()

Lists all HTTP services.

```go
services, err := client.GetHttpServices()
```

---

## Middlewares

### GetHttpMiddlewares()

Lists all HTTP middlewares.

```go
middlewares, err := client.GetHttpMiddlewares()
```

---

## See Also

- [TCP API](tcp.md)
- [UDP API](udp.md)
- [Traefik Docs](https://doc.traefik.io/traefik/operations/api/)
