# Backlog — traefik-go-sdk

Sorgente unica per feature/fix pianificati. Sincronizzato nightly verso GitHub Issues (`[backlog]` label).

## Formato
- `[ID-NNN]` — identificativo stabile (idempotente)
- `[backlog]` — da fare
- `[in-progress]` — in sviluppo (PR aperta)
- `[done]` — completato (PR merged)

---

## Active backlog

### High priority

- `[ID-001]` Add comprehensive unit tests with ≥80% coverage `[backlog]` — No tests found; TDD-first. Target: all public API methods + edge cases (404, timeout, malformed JSON).
- `[ID-002]` Create CHANGELOG.md from git history `[backlog]` — Keep a Changelog format. Retroactive entries from recent commits (bb71140 → cff631c).
- `[ID-003]` Fix HealthCheck() error handling `[backlog]` — Currently returns `(nil, nil)` always. Should propagate resty errors.
- `[ID-004]` Correct haivision typo in codebase `[backlog]` — Stray comment "init haivision" (copia-incolla). Replace with "traefik".

### Medium priority

- `[ID-005]` Upgrade Go version floor to 1.20+ `[backlog]` — 1.18 EOL Aug 2024. Update CI matrix; drop 1.18/1.19.
- `[ID-006]` Document GetApiVersion() and GetApiRawData() edge cases `[backlog]` — Experimental methods (v1.4+). Require Traefik 2.x+.
- `[ID-007]` Add code coverage badge to README `[backlog]` — Link codecov.io or coveralls.io. CI workflow publishes coverage.
- `[ID-008]` Create docs/ARCHITECTURE.md `[backlog]` — Component overview (traefik.go, api_http.go, resty integration), design decisions.

### Low priority / Future

- `[ID-009]` Consider net/http alternative to resty `[future]` — No urgency; resty works fine. Revisit if dependency concerns arise.
- `[ID-010]` Add context.Context support to API methods `[future]` — Async-friendly, timeout control. Post-1.0 feature.
- `[ID-011]` Support multiple Traefik instances (client pool) `[future]` — BuildTraefik() builds 1 client. Pool pattern post-1.0.

---

## Completed

(none yet — this is the initial backlog)

---

## Notes

- Issues synced nightly via `scripts/sync-backlog.sh`
- PR should mention `Closes #NNN` (GitHub auto-close on merge)
- CHANGELOG entries reference resolved backlog items
