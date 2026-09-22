# Architecture Overview

## Layers

**Public API**: ITraefikClient interface with methods for all endpoints

**Transport**: resty HTTP client for actual requests

**Configuration**: BuildTraefik builder pattern

## Design Decisions

See [decisions.md](decisions.md) for detailed rationale.
