# Contributing Guide

We welcome contributions!

## Workflow

1. Fork repository
2. Create feature branch: `git checkout -b feat/your-feature`
3. Follow TDD: test first, then code
4. Ensure `make lint` passes
5. Create Pull Request
6. Wait for review + CI
7. Merge (requires 1 approval)

## Commit Rules

- TDD: Test must precede code
- Semver: Bump minor for features, patch for fixes
- Message: Brief title, then why (not what)

Example:
```
feat: add GetApiVersion() method

- Queries /api/version endpoint
- Tested on Traefik 2.0+
- Closes #123
```

## Code Style

- `gofmt` formatting
- `go vet` compliance
- Table-driven tests
- Error propagation

See [CLAUDE.md](../../CLAUDE.md) for details.
