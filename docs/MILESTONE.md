# Milestone — traefik-go-sdk v1.0.0

**Release**: Q4 2026 (target: end of October)  
**Status**: In progress  
**Goal**: Stable, fully-tested SDK for Traefik API interaction

---

## Milestone Goals

1. ✅ **Complete test suite** (≥80% coverage) — ID-001
2. ✅ **Production-ready documentation** — README, DEVELOPMENT.md, API docs
3. ✅ **Release infrastructure** — CHANGELOG, versioning, tagging
4. ✅ **CI/CD automation** — GitHub Actions, code coverage, branch protection
5. ✅ **Backlog automation** — GitHub Issues sync, tracking

---

## Issues (Backlog Items)

### High Priority (Blocking v1.0.0)

- [ ] **ID-001** — Add comprehensive unit tests with ≥80% coverage
  - Status: 71.1% (target 80%)
  - PR/commit: TBD
  
- [ ] **ID-002** — Create CHANGELOG.md from git history
  - Status: Not started
  - Retroactive entries: bb71140 → cff631c
  
- [ ] **ID-003** — Fix HealthCheck() error handling
  - Status: Identified (returns nil,nil always)
  - Impact: Critical for production use
  
- [ ] **ID-004** — Correct haivision typo in codebase
  - Status: Not started
  - Files: traefik.go comment

### Medium Priority (v1.0.0 nice-to-have)

- [ ] **ID-005** — Upgrade Go version floor to 1.20+
  - Status: Not started
  - Impact: Drop 1.18/1.19 from CI matrix
  
- [ ] **ID-006** — Document GetApiVersion() and GetApiRawData() edge cases
  - Status: Identified (require Traefik 2.x+)
  - Docs: Update README, DEVELOPMENT.md
  
- [ ] **ID-007** — Add code coverage badge to README
  - Status: Not started
  - Integration: codecov.io or coveralls.io
  
- [ ] **ID-008** — Create docs/ARCHITECTURE.md
  - Status: Not started
  - Content: Component overview, design decisions

### Future (post-v1.0.0)

- [ ] **ID-009** — Consider net/http alternative to resty
- [ ] **ID-010** — Add context.Context support to API methods
- [ ] **ID-011** — Support multiple Traefik instances (client pool)

---

## Success Criteria

✅ **All high-priority items completed**  
✅ **Coverage ≥80%**  
✅ **All tests passing on Go 1.18-1.21**  
✅ **Branch main protected** (1 PR review, CI required)  
✅ **CHANGELOG.md created** (Keep a Changelog format)  
✅ **Release tagged** as `v1.0.0`  
✅ **GitHub badge** (build, coverage, version)  

---

## Timeline

| Phase | Target Date | Items |
|-------|-------------|-------|
| Phase 1: Testing | 2026-09-30 | ID-001 (80% coverage) |
| Phase 2: Docs & Release | 2026-10-10 | ID-002, ID-004, ID-006, ID-007, ID-008 |
| Phase 3: Hardening | 2026-10-20 | ID-003, ID-005 |
| Phase 4: Release | 2026-10-31 | Tag v1.0.0, publish to pkg.go.dev |

---

## Tracking

- **Backlog source**: `docs/backlog.md` (single source of truth)
- **Sync**: Nightly via `scripts/sync-backlog.sh` (GitHub Issues ↔ backlog.md)
- **PR labels**: Use `[backlog]` label on issues tied to this milestone
- **Reference**: Mention issue in commit: `Closes #NNN`

---

## Notes

- Release will be tagged as `vX.Y.Z` with CHANGELOG entry
- Semver: 1.0.0 (first stable), future bumps: minor for features, patch for fixes
- No pre-releases (alpha/beta) for v1.0.0
- Deployment: `go get github.com/Allan-Nava/traefik-go-sdk@v1.0.0`

---

## Related

- **docs/backlog.md** — Item details
- **CHANGELOG.md** — Release notes (to create)
- **.github/BRANCH_PROTECTION.md** — Merge rules
- **CLAUDE.md** — Versioning rules
