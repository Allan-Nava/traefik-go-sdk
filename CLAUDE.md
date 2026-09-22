# CLAUDE.md — traefik-go-sdk

**Traefik Go SDK**: libreria Go per interagire con l'API di Traefik (reverse proxy/load balancer). Espone metodi Go per leggere/configurare router HTTP/TCP/UDP, servizi, middleware, entrypoint e diagnostic.

Repo: `github.com/Allan-Nava/traefik-go-sdk` · Go 1.18+ · Dipendenze: resty (HTTP), godotenv, validator

## Regole di lavoro (SEMPRE)

### TDD (Test-Driven Development) — OBBLIGATORIO
- **Test PRIMA del codice**: spec → test (`*_test.go`) → implementazione → commit logico. Mai "implemento e scrivo test dopo".
- **Coverage minimo 80%**: `go test -cover ./...` deve ritornare ≥80%. PR che abbassano coverage sono bloccate in CI.
- **Test structure**: Table-driven tests per metodi pubblici (input → expected output + error case). Esempio:
  ```go
  func TestGetHttpRouters(t *testing.T) {
      tests := []struct {
          name    string
          url     string
          wantErr bool
      }{
          {"valid URL", "http://localhost:8080", false},
          {"empty URL", "", true},
      }
      for _, tt := range tests { ... }
  }
  ```
- **Mocking**: Mock `*http.Client` con `httptest` per isolamento (no network in tests). Resty client mockabile via dependency injection se serve.
- **CI automation**: `go test -v ./... -coverprofile=coverage.out` nel workflow; fallisce se coverage < 80%. Badge coverage nel README (usa `codecov.io` o `coveralls.io` se vuoi).

### Versionamento e release
- **Ogni commit rilevante** ("features", "fix rilevanti", "aggiunte API") merita un **tag semver** (`vX.Y.Z`) + **entry in `CHANGELOG.md`** (formato "Keep a Changelog"). Bump `minor` per nuove API / comportamenti, `patch` per fix / miglioramenti di dettaglio. Helper: `git tag -a vX.Y.Z -m "Release X.Y.Z"` da committare. **Esenti** da tag/CHANGELOG: commit auto su `.claude/settings.json` (hook permessi) e report: CI.
  
### Commit e push
- **MAI `git push`** — lo fa sempre l'utente. **MAI `Co-Authored-By`** nei commit.

### Documentazione API
- **Documentare sempre** aggiunte API e breaking changes: ogni nuovo metodo, tipo di errore, o cambio di firma va documentato in `README.md` (Usage section, Example section) + voce in CHANGELOG. Test di integrazione nelle PR che toccano l'API pubblica.

### Backlog automation
- **Sorgente unica**: `docs/backlog.md` (Keep a Changelog-style, item con `id` stabile). Sync idempotente verso GitHub Issues (`[backlog]` label).
- **Workflow**:
  1. Feature/fix → issue GitHub con label `[backlog]`
  2. Sincronizzare in `docs/backlog.md` (riga: `- [ID-001] Item description [backlog]`)
  3. PR menciona issue (`Closes #123`) → auto-close su merge
  4. CHANGELOG riporta item risolti (es. "Closes #123 - Added GetApiVersion()")
- **Helper script** (`scripts/sync-backlog.sh`): cron nightly che sincronizza GitHub Issues ↔ backlog.md

### Dipendenze
- **`go.mod`/`go.sum`** sempre sincronizzati (`go mod tidy` prima di committare). PR che aggiungono dipendenze devono giustificarle (perché resty vs Go nativo? Perché validator?).

### Naming e style
- **Interfacce uppercase** (`ITraefikClient`), **receiver brevi** (`o *traefikSdk`), **metodi getter** (`IsDebug()`, `GetHttpRouters()`). Seguire `gofmt` e `go vet` — niente linter custom per ora.

## Note implementative

- **Base URL construction**: Builder `BuildTraefik(url string, debug bool)` accetta URL generico; costruire endpoint con `http.Client` / `*resty.Client` standard (no magic path prefix). Verificare che `/api` sia parte dell'URL passato dal caller.

- **Error handling**: Metodi ritornano `(*resty.Response, error)` — il caller decide se interpretare status HTTP. Non fare assumption su 200/4xx/5xx senza documentarlo.

- **Debug mode**: flag `debug` in struct, attivabile al build via `BuildTraefik(url, true)` — accende log resty verso stdout (`SetDebug`). Non aggiungere opzioni di file output senza richiesta.

- **Typo storico**: "haivision" in commenti (`// init haivision`) → correggere quando tocchi quel codice.

- **API versione / raw data**: metodi `GetApiVersion()` / `GetApiRawData()` aggiunti di recente (commit `30e3d4b`) — sono sperimentali per ora, documentare cautele su Traefik < 2.x .

## Workflow di revisione PR (TDD first)

1. **Test coverage check**: 
   ```bash
   go test -v ./... -coverprofile=coverage.out
   go tool cover -html=coverage.out  # visualizzare gap
   ```
   Deve ritornare ≥80%. PR che abbassano coverage = reject.

2. **Build locale** (su tutte Go 1.18-1.21):
   ```bash
   go mod tidy
   go fmt ./...
   go vet ./...
   go build -v ./...
   go test -v ./... -race  # detect data races
   ```

3. **Diff review**:
   - Niente `.mod`/`.sum` accidentali
   - Test DEVE precedere implementazione (verifica commit history)
   - Metodi pubblici hanno test corrispondenti
   - Niente commenti obsoleti o debug print

4. **Docs aggiornati**: se l'API cambia (firma, ritorno, endpoint), aggiornare README + CHANGELOG + `docs/backlog.md` contestualmente.

5. **Backlog sync**: PR chiude issue? Aggiornare `docs/backlog.md` (riga diventa completata / archiviata).

6. **Merge**: linear history. Log pulito per git blame future.

## Puntatori

- **Traefik API docs**: https://doc.traefik.io/traefik/operations/api/ (endpoint ufficiali, formato risposta)
- **Go best practices**: https://golang.org/doc/effective_go (naming, style, error handling)
- **Resty client**: https://pkg.go.dev/github.com/go-resty/resty/v2 (HTTP lib usata)
- **Backlog / roadmap**: issues GitHub nel repo
- **CI status**: GitHub Actions badge nel README (go-build, go-test)
