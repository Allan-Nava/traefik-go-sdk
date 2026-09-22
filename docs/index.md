# Traefik Go SDK

**A modern Go SDK for Traefik reverse proxy and load balancer**

[![Go Reference](https://pkg.go.dev/badge/github.com/Allan-Nava/traefik-go-sdk.svg)](https://pkg.go.dev/github.com/Allan-Nava/traefik-go-sdk)
[![Go build](https://github.com/Allan-Nava/traefik-go-sdk/actions/workflows/go-build.yml/badge.svg)](https://github.com/Allan-Nava/traefik-go-sdk/actions/workflows/go-build.yml)
[![Go test](https://github.com/Allan-Nava/traefik-go-sdk/actions/workflows/go-test.yml/badge.svg)](https://github.com/Allan-Nava/traefik-go-sdk/actions/workflows/go-test.yml)
[![Coverage Status](https://img.shields.io/badge/coverage-71%25-yellowgreen)](docs/coverage.md)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## Features

✨ **Comprehensive API Coverage**
- HTTP routers, services, middlewares
- TCP and UDP routers and services
- Entrypoint listing and diagnostics
- API overview and health checks

🧪 **Production-Ready**
- 50+ test cases (71.1% coverage)
- TDD-first development
- Pre-commit hooks for code quality
- GitHub Actions CI/CD

🛡️ **Secure by Default**
- Branch protection on main
- Require PR reviews and status checks
- Comprehensive error handling
- Race detector enabled

📚 **Well-Documented**
- Extensive API documentation
- Developer guides and examples
- Architecture decision records
- Clear troubleshooting guide

---

## Quick Start

### Installation

```bash
go get github.com/Allan-Nava/traefik-go-sdk
```

### Usage Example

```go
package main

import (
	"fmt"
	"log"
	
	"github.com/Allan-Nava/traefik-go-sdk/traefik"
)

func main() {
	// Create a Traefik client
	client, err := traefik.BuildTraefik("http://localhost:8080", false)
	if err != nil {
		log.Fatal(err)
	}

	// Get all HTTP routers
	routers, err := client.GetHttpRouters()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Response: %s\n", routers.String())
}
```

See [Quick Start Guide](guides/02-quick-start.md) for more examples.

---

## Project Status

### Current Release
- **Version**: v0.2.0 (TDD Infrastructure + GitHub Automation)
- **Coverage**: 71.1% (target: 80%)
- **Status**: Active Development

### v1.0.0 Roadmap
Target: October 2026

- ✅ Comprehensive test suite (50+ tests)
- ✅ TDD enforcement (pre-commit hooks)
- ✅ GitHub branch protection
- ⏳ Reach 80% coverage
- ⏳ Fix HealthCheck() error handling
- ⏳ Create production documentation

See [Milestone v1.0.0](architecture/roadmap.md) for detailed roadmap.

---

## Why Traefik Go SDK?

### Problem
Traefik's API is powerful but required manual HTTP calls. No native Go integration.

### Solution
A type-safe, tested Go SDK that:
- Reduces boilerplate HTTP client code
- Provides compile-time type checking
- Includes comprehensive test coverage
- Follows Go best practices

### Use Cases
- Automated configuration management
- Health monitoring and diagnostics
- Integration with other Go tools
- CI/CD pipeline automation

---

## Architecture

The SDK is organized in layers:

```
├── Public API (ITraefikClient interface)
│   ├── HTTP endpoints (routers, services, middlewares)
│   ├── TCP endpoints
│   ├── UDP endpoints
│   └── Info endpoints (health, overview, version)
│
├── Transport Layer (resty HTTP client)
│   ├── restyGet() - HTTP GET requests
│   └── restyPost() - HTTP POST requests
│
└── Configuration
    ├── Builder pattern (BuildTraefik)
    └── Debug mode support
```

See [Architecture Overview](architecture/overview.md) for detailed design.

---

## Development

### Quick Setup

```bash
# Clone and setup
git clone https://github.com/Allan-Nava/traefik-go-sdk.git
cd traefik-go-sdk

# Enable git hooks (TDD enforcement)
git config core.hooksPath .githooks

# Run tests
make test

# Check coverage
make coverage

# Full lint (fmt + vet + test + coverage)
make lint
```

### Workflow

This project uses **TDD (Test-Driven Development)**:

1. Write test first (RED)
2. Implement code (GREEN)
3. Ensure coverage ≥70% (CI threshold)
4. Commit with pre-commit hooks

See [TDD Workflow](development/tdd.md) for details.

---

## Contributing

Contributions are welcome! Please read our [Contributing Guide](development/contributing.md) first.

**Key principles:**
- TDD: Test first, always
- Coverage: ≥70% (targeting 80%)
- Documentation: Update alongside code
- Semantic Versioning: Major.Minor.Patch

---

## License

This project is licensed under the MIT License — see [LICENSE](../LICENSE) file for details.

---

## Resources

- **API Documentation**: [Traefik Official Docs](https://doc.traefik.io/traefik/operations/api/)
- **Go Packages**: [pkg.go.dev](https://pkg.go.dev/github.com/Allan-Nava/traefik-go-sdk)
- **GitHub Issues**: [Report Bugs](https://github.com/Allan-Nava/traefik-go-sdk/issues)
- **Discussions**: [Ask Questions](https://github.com/Allan-Nava/traefik-go-sdk/discussions)

---

## Support

Having issues? Check [Troubleshooting](troubleshooting.md) or open a GitHub issue.

**Need help?**
- 📖 Read the [docs](guides/01-installation.md)
- 🐛 Check [existing issues](https://github.com/Allan-Nava/traefik-go-sdk/issues)
- 💬 Start a [discussion](https://github.com/Allan-Nava/traefik-go-sdk/discussions)
