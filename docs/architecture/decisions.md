# Design Decisions

## Why resty over net/http?

**Decision**: Use resty HTTP client.

**Rationale**:
- Built-in retry logic
- Simplified header/query param handling
- Better error handling

**Alternative**: Direct net/http (more control, less convenience)

## Why interface over concrete struct?

**Decision**: Public API as ITraefikClient interface.

**Rationale**:
- Testability (easy to mock)
- Dependency inversion
- Future extensibility

## Semantic Versioning

**Decision**: Strict semver (1.0.0).

**Rationale**:
- Consumer trust
- Clear upgrade expectations
- Go module compatibility
