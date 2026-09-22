# Traefik go sdk
[![Go Reference](https://pkg.go.dev/badge/github.com/Allan-Nava/traefik-go-sdk.svg)](https://pkg.go.dev/github.com/Allan-Nava/traefik-go-sdk)
[![Go build](https://github.com/Allan-Nava/traefik-go-sdk/actions/workflows/go-build.yml/badge.svg)](https://github.com/Allan-Nava/traefik-go-sdk/actions/workflows/go-build.yml)

The Traefik Go SDK provides a convenient way to interact with Traefik, a modern reverse proxy and load balancer. This SDK is designed to simplify the integration of Traefik into your Go applications.

## Features

- **Dynamic Configuration:** Easily configure and manage Traefik's dynamic configuration from your Go application.
- **Programmatic Interaction:** Interact with Traefik programmatically, making it easier to automate configuration changes.

## Installation

To use the Traefik Go SDK in your project, you need to install it using the `go get` command:

```bash
go get github.com/Allan-Nava/traefik-go-sdk
```

## Usage

Here is a simple example demonstrating how to use the Traefik Go SDK:

```go
import (
	"github.com/Allan-Nava/traefik-go-sdk/traefik"
)

func main() {
	// Create a Traefik client
	client, err := traefik.BuildTraefik("http://localhost:8080", false)
	if err != nil {
		panic(err)
	}

	// Read router configuration
	routers, _ := client.GetHttpRouters()
	println("Routers:", routers.Status())

	// Create a new router
	newRouter := map[string]interface{}{
		"entryPoints": []string{"web"},
		"service":     "my-service",
		"rule":        "Host(`example.com`)",
	}
	resp, _ := client.CreateHttpRouter("my-router", newRouter)
	println("Created:", resp.Status())

	// Update router
	updatedRouter := map[string]interface{}{
		"entryPoints": []string{"web", "websecure"},
		"service":     "my-service",
		"rule":        "Host(`example.com`)",
	}
	client.UpdateHttpRouter("my-router", updatedRouter)

	// Delete router
	client.DeleteHttpRouter("my-router")
}
```

## API Methods

### Read Operations

- **HTTP Routers:** `GetHttpRouters()`, `GetHttpRouter(name)`
- **HTTP Services:** `GetHttpServices()`
- **HTTP Middlewares:** `GetHttpMiddlewares()`
- **TCP Routers:** `GetTcpRouters()`, `GetTcpRouter(name)`
- **TCP Services:** `GetTcpServices()`
- **UDP Routers:** `GetUdpRouters()`, `GetUdpRouter(name)`
- **UDP Services:** `GetUdpServices()`, `GetUdpService(name)`
- **Entry Points:** `GetEntrypoints()`
- **API Info:** `GetApiOverview()`, `GetApiVersion()`, `GetApiRawData()`

### Write Operations

- **HTTP Routers:** `CreateHttpRouter(name, config)`, `UpdateHttpRouter(name, config)`, `DeleteHttpRouter(name)`
- **HTTP Services:** `CreateHttpService(name, config)`, `UpdateHttpService(name, config)`, `DeleteHttpService(name)`
- **TCP Routers:** `CreateTcpRouter(name, config)`, `UpdateTcpRouter(name, config)`, `DeleteTcpRouter(name)`
- **TCP Services:** `CreateTcpService(name, config)`, `UpdateTcpService(name, config)`, `DeleteTcpService(name)`
- **UDP Routers:** `CreateUdpRouter(name, config)`, `UpdateUdpRouter(name, config)`, `DeleteUdpRouter(name)`
- **UDP Services:** `CreateUdpService(name, config)`, `UpdateUdpService(name, config)`, `DeleteUdpService(name)`
- **HTTP Middlewares:** `CreateHttpMiddleware(name, config)`, `UpdateHttpMiddleware(name, config)`, `DeleteHttpMiddleware(name)`
- **TCP Middlewares:** `CreateTcpMiddleware(name, config)`, `UpdateTcpMiddleware(name, config)`, `DeleteTcpMiddleware(name)`

### Batch Operations

- **Configuration Management:** `ApplyConfiguration(config)` — Apply complete config (HTTP + TCP + UDP)
- **Configuration Validation:** `ValidateConfiguration(config)` — Validate before applying
- **Configuration Retrieval:** `GetConfiguration()` — Get current raw data
- **Configuration Reset:** `ResetConfiguration()` — Reset to default state

### Advanced Filtering

- **Router Queries:** `GetHttpRoutersByRule(rule)`, `GetTcpRoutersByEntryPoint(ep)`, `GetUdpRoutersByEntryPoint(ep)`
- **Service Queries:** `GetServicesByRouter(name)` — Get services for a specific router
- **Middleware Queries:** `GetHttpMiddlewareByType(type)` — Get middlewares by type (basicAuth, ipWhiteList, etc.)

### Utility Methods

- **Health Check:** `HealthCheck()` — Verify connectivity to Traefik API
- **Debug Mode:** `IsDebug()` — Check if debug logging is enabled

For more detailed examples and [API documentation](https://doc.traefik.io/traefik/operations/api/), refer to the GoDoc.

## Documentation

The documentation site at https://allan-nava.github.io/traefik-go-sdk/ is automatically generated from this `README.md` file using a custom Node.js build script. The site is a single-page application with responsive design and dark mode support.

### Updating the Site

The site is built from `README.md` content organized by H2 headers (`## Section`). To update the documentation:

1. **Edit README.md** — Modify any section or add new H2 sections
2. **Build locally** — Run `npm run build` to generate `site/dist/index.html`
3. **Review** — Open `site/dist/index.html` in your browser to preview
4. **Commit & Push** — Standard Git workflow; CI will rebuild the site automatically

### Site Architecture

- **`site/build.mjs`** — Node.js build script that:
  - Parses `README.md` by H2 headers (each becomes a section)
  - Converts markdown to HTML using the `marked` library
  - Generates semantic HTML with JSON-LD structured data
  - Links file paths in code blocks to GitHub repository
  - Supports light/dark mode CSS variables
  
- **`site/package.json`** — Dependencies and build scripts
- **`site/dist/index.html`** — Generated single-page site (~10KB)
- **`.github/workflows/pages.yml`** — GitHub Actions builds and deploys to `gh-pages` branch

### Deployment

The site is deployed to GitHub Pages via:

```bash
npm run build
git subtree push --prefix site/dist origin gh-pages
```

This pushes only the `site/dist/` folder to the `gh-pages` orphan branch, which is served at the repository's GitHub Pages URL. The main repository history remains clean.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License
This Traefik Go SDK is open-source software licensed under the MIT License.