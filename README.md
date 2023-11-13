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
Here is a simple example demonstrating how to use the Traefik Go SDK to update Traefik's dynamic configuration:
```go

import (
	"fmt"
	"github.com/Allan-Nava/traefik-go-sdk"
)

func main() {
	// Create a Traefik client
	client, err := traefik.BuildTraefik("http://traefik-api-url")
	if err != nil {
		panic(err)
	}
}

```

For more detailed examples and [API documentation](https://doc.traefik.io/traefik/operations/api/), refer to the GoDoc.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License
This Traefik Go SDK is open-source software licensed under the MIT License.