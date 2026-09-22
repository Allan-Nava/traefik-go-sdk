# AGENTS.md — traefik-go-sdk

Questo file definisce le regole operative per gli agent (Claude, Copilot, altri tool AI) quando lavorano in questo repository.

**Traefik Go SDK**: libreria Go per interagire con l'API di Traefik. Struttura semplice, API pubblica stabile (v1.x), CI/CD minimal (Go build/test).

## Regole di lavoro (SEMPRE)

### TDD (Test-Driven Development) — OBBLIGATORIO
- **Test PRIMA del codice**: spec → test (`*_test.go`) → implementazione → commit logico. Mai "implemento e scrivo test dopo".
- **Coverage minimo 80%**: `go test -cover ./...` deve ritornare ≥80%. PR che abbassano coverage sono bloccate in CI.
- **Test structure**: Table-driven tests per metodi pubblici. Mock `*http.Client` con `httptest`. Resty client mockabile via dependency injection.
- **CI automation**: workflow `go-test.yml` esegue `go test -v ./... -coverprofile=coverage.out`. Fallisce se coverage < 80%. Badge coverage nel README (codecov.io / coveralls.io).

### Versionamento e release
- **Ogni commit** "features", "fix rilevanti", "aggiunte API" merita **tag semver** (`vX.Y.Z`) + **entry in `CHANGELOG.md`**. Bump `minor` per nuove API / comportamenti, `patch` per fix / miglioramenti. Tag committato subito dopo il commit logico. **Esenti**: auto-commit su `.claude/settings.json` (hook permessi), commit `report:` della CI.

### Commit e push
- **MAI `git push`** — lo fa l'utente. **MAI `Co-Authored-By`** nei commit.

### Documentazione API
- **Documentare sempre**: ogni nuova API pubblica va documentata in `README.md` (Usage/Example), aggiornare il CHANGELOG contestualmente (nuova sezione `[X.Y.Z] - YYYY-MM-DD`). Se cambiano error types o comportamenti di metodi esistenti = **breaking change** = minor version.

### Backlog automation
- **Sorgente unica**: `docs/backlog.md` (item con `id` stabile). Sync idempotente verso GitHub Issues (`[backlog]` label).
- **Workflow**:
  1. Feature/fix richiesta → issue GitHub (`#NNN`)
  2. Aggiungere a `docs/backlog.md`: `- [ID-001] Description [backlog]`
  3. PR menciona issue: `Closes #NNN` → auto-close su merge
  4. CHANGELOG riporta issue chiusi
- **Helper script**: `scripts/sync-backlog.sh` (cron nightly, GitHub Issues ↔ backlog.md)

### Dipendenze
- **`go.mod`/`go.sum`**: sempre tidy. PR che aggiungono dipendenze **devono giustificarle**. Valutare: esiste in stdlib? Esiste già un sostituto nel progetto?

### Code style
- **`gofmt`, `go vet`, `go fmt ./...`** prima di committare. Receiver brevi, interfacce uppercase, no underscore prefixes. Niente custom linting per ora.

## Regole tecniche

- **resty Client**: unica dipendenza HTTP usata. Ritorno standard `(*resty.Response, error)` — il caller decide interpreting status code. Non aggiungere wrapper custom (JSONMarshal, retry logic) senza richiesta esplicita.

- **Builder pattern**: `BuildTraefik(url string, debug bool)` è l'entry point. Accetta URL generico (il caller passa `http://traefik-api:8080`, SDK aggiunge `/api/...` negli endpoint). Non assumere port o schema — lato caller.

- **Error handling**: Ritorno sempre tupla `(response, error)`. Se resty fallisce (network error), ritorna `(nil, err)`. Se HTTP è 4xx/5xx, ritorna `(resp, nil)` — responsabilità del caller parare lo status. Documentare eccezioni (es., `HealthCheck()` ritorna errore solo se network down).

- **Debug mode**: flag `debug` accende `SetDebug(true)` su resty (log stdout). Non aggiungere:
  - opzioni di file output
  - rotating log
  - syslog
  - livelli customizzati (info, debug, warn)
  Senza richiesta esplicita.

- **API versioning**: metodi `GetApiVersion()`, `GetApiRawData()` are **experimental** (v1.4+). Documentare che richiedono Traefik 2.x+. In future, versioning tramite struct field (`Version string` in response) o nuovo metodo `GetApiMetadata()` — niente "v2" client parallelato.

- **Typo storico**: "haivision" in commenti (`// init haivision`) — NON è intenzionale (copia-incolla da altro SDK). Se tocchi quel blocco, correggilo; se lo vedi in PR altrui, segnalalo gentilmente.

## Pattern di intervento (TDD-first)

### 1. Spec → Test → Implementation
```
Issue #NNN: "Add GetApiVersion() method"
├─ Step 1: Scrivi test (ROSSO)
│  └─ tests/api_test.go: TestGetApiVersion() table-driven
├─ Step 2: Implementa (GREEN)
│  └─ api_info.go: func (o *traefikSdk) GetApiVersion() ...
├─ Step 3: Coverage ≥80%
│  └─ go test -cover ./... → 82%
└─ Step 4: Commit logico
   └─ git commit -m "feat: add GetApiVersion() to ITraefikClient

   - Queries /api/version endpoint
   - Returns version struct with semantic versioning
   - Edge case: Traefik < 2.x returns 404 (tested)
   Closes #NNN"
```

### 2. Scope validation
Nuova API? Bug fix? Dependency bump? Chiarirlo nel commit message e issue.

### 3. Build & test locale (TDD-aware)
```bash
# 1. Test prima di implementazione
go test -v ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

# 2. Tidy + style
go mod tidy
go fmt ./...
go vet ./...

# 3. Build
go build -v ./...

# 4. Retest + race detector
go test -v ./... -race
```

### 4. Coverage check
- **Nuova API** = test obbligatorio (table-driven: happy path + edge case)
- **Bug fix** = aggiungi test che dimostra il bug (rosso) → fix → verde
- **Coverage report**: `go tool cover -html=coverage.out` — visualizzare gap
- **Reject se < 80%**: CI fallisce, PR bloccata

### 5. Documentazione (DDD — Documentation-Driven Development)
- **README.md**: aggiungere example se nuova API pubblica
- **CHANGELOG.md**: nuova sezione `[X.Y.Z] - YYYY-MM-DD` con entry (`- Added GetApiVersion() method`)
- **docs/backlog.md**: aggiornare item risolti (es. `- [ID-001] Add GetApiVersion() [DONE]`)
- **Commit message**: titolo + corpo (se multi-line) = WHY (risolvere quale problema?), non WHAT (il diff lo mostra)

### 6. Diff review (pre-merge)
- Niente `.mod`/`.sum` accidentali
- Niente debug print / commented code
- Metodi **completi** (no half-implementation, no TODO comments)
- `ITraefikClient` aggiornata se nuovi metodi pubblici
- **Test PRIMA del codice** nel history (git log --oneline)

### 7. Merge
Linear history (`--no-ff` se squash, o single commit). Log pulito per git blame future.

## Decisions log (storico)

- **TDD obbligatorio** (2026-09-22): Coverage ≥80% in CI, test-first workflow. Rationale: SDK pubblica — rischio API instabile è alto; TDD riduce regression e documenta behavior; coverage badge = trust indicator per consumers.
- **Backlog automation** (2026-09-22): `docs/backlog.md` + GitHub Issues sync. Rationale: sorgente unica, evita TODO sparsi nel codice; script nightly mantiene sync idempotente.
- **go.mod**: Go 1.18+ (no 1.17). Resty per HTTP (alternativa: net/http std, deciso resty per built-in retry/timeout/header handling).
- **Interfaccia pubblica**: `ITraefikClient` vs `TraefikClient` concrete — interfaccia per testability (mock in consumer code).
- **Tag release**: semver stretto (1.0.0 = first stable, breaking = +minor, features = +patch). No alpha/beta per ora.

## Puntatori

- **Traefik API**: https://doc.traefik.io/traefik/operations/api/
- **Go idioms**: https://golang.org/doc/effective_go
- **Resty docs**: https://pkg.go.dev/github.com/go-resty/resty/v2
- **GitHub Actions**: `.github/workflows/` (go-build.yml, go-test.yml) — update matrix Go versions se needful
- **Issues/backlog**: GitHub Issues (no separate backlog.md per ora)
