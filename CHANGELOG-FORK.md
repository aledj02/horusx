# Horizon fork changelog (`horusx`)

> **Fork de**: [ZupIT/horusx](https://github.com/ZupIT/horusec) — commit `873d4104` (v2.9.0-beta.3 upstream era la referencia de baseline)
> **Tag baseline**: `v2.10.0-fork.0` (equivalente al último upstream)
> **Convención de versión desde F5**: SemVer estándar (`v2.10.1`, `v2.11.0`, …), **sin el sufijo `-fork.N`** (decision 2026-10-21).
> **Objetivo**: fork interno, SAST modernizado para **JS, TS, Java, Go, Python, Shell, Kotlin, Swift** — con custom rules RE2-validadas y Docker images frescas (cuando se aplique Fase 5).

> **Formato**: [keep-a-changelog](https://keepachangelog.com/es/1.1.0/) + [conventional commits](https://www.conventionalcommits.org/es/) de referencia.
> Cada bloque de cambio lista: **motivo**, **archivos**, **verificación**. Commits aún **sin push** (pendientes de aprobación del equipo).

---

## Fases aplicadas (2026-10-20, sin commits)

| # | Fase | Estado |
|---|---|---|
| F0 | Fork + tag baseline | ✅ |
| F1 | Deps Go modernizadas | ✅ |
| F2 | Custom-rules map extendido a Python/Go/Shell/Swift | ✅ |
| F3 | Catálogo de reglas JSON (50 reglas) | ✅ |
| F4 | Wiring de formatters (Python/Go/Shell — Swift ya existía) | ✅ |
| F5 | Rebuild de Docker images `horusx/*` | ⏸ pendiente (requiere account Docker Hub) |
| F6 | Testing local (unit + E2E) | ✅ |
| F7 | Distribución (binario + README-FORK) | 🔶 README pendiente |

---

## [Unreleased] — Fork v2.10.0-fork.1

### Added
- **Custom-rules support para 4 lenguajes que no lo tenían**: Python, Go, Shell, Swift.
  - `internal/services/custom_rules/custom_rule.go` — agregadas las `languages.Language` keys al map de allowed languages (antes solo C#/Java/JS/Kotlin/Leak/Nginx/YAML).
  - `internal/services/custom_rules/service.go` — map de rule ID prefixes (`HS-PYTHON-`, `HS-GO-`, `HS-SHELL-`, `HS-SWIFT-`) → validación de que el `id` del JSON coincida con el `language` (evita que un JSON Python sea cargado por `HS-GO-100` por error).
  - 3 nuevos rule managers LIGEROS (built-in `0` reglas):
    - `internal/services/engines/python/rule_manager.go` — extensions `.py`
    - `internal/services/engines/go/rule_manager.go` — extensions `.go`
    - `internal/services/engines/shell/rule_manager.go` — extensions `.sh`, `.bash`, `.zsh`, `.ksh`, `.ksh93`
  - 3 nuevos formatters (wrapper `NewDefaultFormatter` por lenguaje):
    - `internal/services/formatters/python/horusxpython/formatter.go`
    - `internal/services/formatters/go/horusxgo/formatter.go`
    - `internal/services/formatters/shell/horusxshell/formatter.go`
  - `custom_rules/` — carpeta con el catálogo:
    - `python.json` — 16 reglas (`HS-PYTHON-100` … `115`)
    - `go.json` — 11 reglas (`HS-GO-100` … `110`)
    - `shell.json` — 12 reglas (`HS-SHELL-100` … `111`)
    - `swift.json` — 11 reglas (`HS-SWIFT-100` … `110`) — complemente de los 24 built-in de Swift
    - `horusx-custom-rules.json` — combinado de los 4 (50 reglas total)
    - `README.md` — schema + uso + cómo generar el combinado
- `internal/controllers/analyzer/runner.go` — `detectVulnerability{Python,Go,Shell}` **ahora ejecutan el formatter custom-rules ANTES de los Docker tools** (bandit/safety/gosec/nancy/shellcheck). Así el scanner LIGERO sigue funcionando aunque Docker falle (Fase 5 aún pendiente).
- **Tests** — 6 nuevos archivos:
  - `internal/services/engines/python/rule_manager_test.go` — `NewRules()` → `GetAllRules()` empty (LIGERO), `GetAllExtensions() == [".py"]`
  - `internal/services/engines/go/rule_manager_test.go` — idem, `[".go"]`
  - `internal/services/engines/shell/rule_manager_test.go` — idem, `[".sh",".bash",".zsh",".ksh",".ksh93"]`
  - `internal/services/formatters/default_engine_formatter_test.go` — ampliada: 9 → 12 formatters (agregados Python, Go, Shell)
  - `internal/services/custom_rules/e2e_test.go` — **`TestE2E_CustormRulesDischargeFixture`**: 1 fixture por lenguaje (pickles, TLS skip, `curl|bash`, `NSKeyedUnarchiver`) → 1 finding cada uno usando el JSON real (no mock).
  - `internal/services/custom_rules/custom_rule_test.go` — nuevos casos de test para validación de `language` + `id` en Python/Go/Shell.
- `examples/{python,go,shell,swift}/example1/` — 4 fixtures E2E:
  - `pickle_loads.py` (`HS-PYTHON-100`)
  - `insecure_tls.go` (`HS-GO-100`)
  - `install_from_url.sh` (`HS-SHELL-100`)
  - `insecure_unarchive.swift` (`HS-SWIFT-101`)
- `CHANGELOG-FORK.md` (este archivo) — log de auditoría del fork.
- `migration-plan.md` — plan de trabajo del fork, con progreso por fase.

### Changed
- `go.mod` — **`go 1.17` → `go 1.24`**.
- `go.mod` — `github.com/docker/docker@v20.10.9` → **`v24.0.9+incompatible`**.
  - **Motivo** (de `D11` en `migration-plan.md`): v25+ movió los tipos de `api/types` a sub-packages (`container/image/registry`), obligando a renombrar decenas de usages y arrastrando otel/grpc/genproto a la cadena. v24 moderniza sin ese costo.
- `go.sum` — 706 lines re-resueltos por `go mod tidy` con las novas versiones.
- `internal/services/docker/docker_api.go` — adaptación a la API del SDK v24 (signatures levemente cambiantes en `client.Container*`).
- `internal/services/docker/client/docker_config_test.go` + `internal/utils/testutil/docker_mock.go` — mocks alineados a la firma v24.
- `internal/controllers/analyzer/analyzer_test.go` — mocks actualizados por el cambio de firma.
- `internal/services/custom_rules/custom_rule.go:146-148` — validación de que el `language` del JSON coincide con el `id` (regla nueva para los 4 lenguajes añadidos).

### Fixed
- **Bug** (descubierto en T2): el test E2E estaba infiriendo el lenguaje del nombre del test-case (`name[:3] == "Pyt"` / `"She"` / `"Swe"`). Fallaba para **Shell** (`"Shel" != "She"` por slice de 4) y **Swift** (`"Swi" != "Swe"`).
  - **Fix**: el struct `e2eCase` ahora lleva `lang languages.Language` explícito por caso. `detectLanguage(name)` eliminado.
  - **Archivo**: `internal/services/custom_rules/e2e_test.go:30-90`.
  - **Verificado**: `go test -run TestE2E_CustormRulesDischargeFixture` → PASS los 4 casos.

### Verified
- `go build ./...` → ok (todo el repo).
- `go vet ./internal/services/custom_rules/...` → clean.
- `go test ./internal/services/engines/{python,go,shell}/...` → ok.
- `go test ./internal/services/formatters/...` → ok (12 formatters, no panics).
- `go test ./internal/services/custom_rules/...` → ok (incluye E2E de 4 reglas).
- `go test ./internal/controllers/analyzer/...` → ok.
- `go test ./config/...` → ok.
- **Fails pre-existentes (no son regresiones)** — formatters que cargan fixtures `examples/<lang>/example*` desde el git submodule `horusx-examples-vulnerabilities` (no inicializado en el fork): `csharp/dotnet_cli`, `elixir/{mixaudit,sobelow}`, `generic/trivy`, `go/nancy`, `javascript/{npmaudit,yarnaudit}`, `python/safety`, `ruby/{brakeman,bundler}`. Ninguno de esos files aparece en `git status --short` de este fork.

### E2E (manual, sin binario)
```bash
/tmp/horusx start \
  -p /Users/naranjax/nx/hx-test \
  --custom-rules-path /Users/naranjax/nx/horusx/custom_rules/horusx-custom-rules.json \
  --output-format json --json-output-file /tmp/hx-out.json
```
Resultados (`analysisVulnerabilities[].vulnerabilities.rule_id`):
| rule_id    | file      | origen        |
|------------|-----------|---------------|
| `HS-PYTHON-100` | `vuln.py`   | custom rule (fork) |
| `HS-GO-106`     | `vuln.go`   | custom rule (fork) |
| `HS-SWIFT-101`  | `App.swift` | custom rule (fork) |
| `HS-SWIFT-20`   | `App.swift` | built-in (upstream) |
| `HS-SHELL-100`  | `vuln.sh`   | custom rule (fork) |
| `HS-LEAKS-25`   | `App.swift` | built-in (upstream) |
| `B301`, `B403`  | `vuln.py`   | bandit (Docker) |
| `G104`          | `vuln.go`   | gosec (Docker)    |

9 findings — los 4 custom rules nuevas disparan **en paralelo** a los tools de Docker.

### Not included (pendientes, Fase 5 + 7)
- **Docker images `horusx/*`** (Python bandit 1.9.4 + safety 3.8.1, Go gosec v2.29 + nancy v2.1.0, Shell shellcheck v0.11.0, JS/TS node 20.x) — requieren account Docker Hub `horusx` y push multi-arch.
- **`internal/enums/images/images.go`** — cuando se publiquen las images, se actualizan los 4 IDs.
- **`README-FORK.md`** — build from scratch, how to add rules, how to point at private registry.
- **Install script** `install.sh` (downloader de GH release privado) — Fase 7.
- **Cosign / firmar imágenes** — no (D5).
- **Commits + push** — pendientes de aprobación del equipo.

---

## [v2.12.0] — Rename `horusec` → `horusx`

**Fecha**: 2026-10-21 · **Tag**: `v2.12.0` · **Commit**: `9e6a772d`
**Fork**: `aledj02/horusx` · **Release**: `https://github.com/aledj02/horusx/releases/tag/v2.12.0`
**Assets**: `horusx_2.12.0_{darwin_arm64,linux_amd64,linux_arm64}.tar.gz` + checksums.

### BREAKING
- Binario: `horusec` → `horusx`.
- Env vars: `HORUSEC_CLI_*` → `HORUSX_CLI_*`.
- Config file: `horusec-config.json` → `horusx-config.json`.
- Packages Go: `horusecpython` → `horusxpython`, `internal/services/horusec_api` → `horusx_api`, `internal/services/formatters/*horusec*` → `*horusx*` (10 dirs).
- Custom-rules catalog: `custom_rules/horusec-custom-rules.json` → `horusx-custom-rules.json`.
- Workflows: `.github/workflows/update-horusec-*.yml` → `update-horusx-*.yml` (10).
- Install script env vars: `HORUX_*` → `HORUSX_*` (corruption del sed del rename).

### NOT BREAKING
- **`github.com/ZupIT/horusec`** (module path en `go.mod`) — NO se renombra (upstream).
- **`tools.HorusecEngine`** (`ZupIT/horusec-devkit` @ v1.0.24) — valor string literal `"HorusecEngine"` EXTERNO al fork; se preserva en `internal/entities/toolsconfig/tools_config.go:43,72`, `horusx-config.json:75`, y en los tests (`config_test.go:515`).
- **Docker images** — siguen siendo `aledj02/horusx-{python,go,shell,js}` (ya se llamaban `horusx`, sin cambio).

### Changed (148 files, 1288+/1285-)
- Branding: `Horusx` / `HORUSX` / `horusx` a lo largo de workflows, config, README, docs.
- `.goreleaser.yml` — binary name `horusx`.
- `install.sh` — BIN_NAME=`horusx`, env vars `HORUSX_{VERSION,BIN_DIR,REPO}`.
- `horusx-config.json` — nuevo filename (content re-branded).

### Verified
- `go build ./...` → EXIT=0.
- `gofmt -l .` → 0 files.
- `go test ./config/ ./internal/entities/toolsconfig/ ./internal/services/formatters/ ./internal/controllers/printresults/` → ok/ok/ok/ok.
- E2E: `HORUSX_BIN_DIR=/tmp/hx-e2e/bin HORUSX_VERSION=v2.12.0 bash install.sh` → `/tmp/hx-e2e/bin/horusx version` → `Version: 2.12.0, Git commit: 9e6a772d`.

### Not included
- `internal/services/formatters/*` y `e2e/*` pre-existing failures (por fixtures ausentes o docker pulls de `horuszup/*` que no existen) — NO causados por el rename.

---

## [v2.11.0] — Distribución (release pipeline)

**Fecha**: 2026-10-21 · **Tag**: `v2.11.0`
**Release**: `https://github.com/aledj02/horusx/releases/tag/v2.11.0`
**Assets**: `horusec_2.11.0_{darwin_arm64,linux_amd64,linux_arm64}.tar.gz` + checksums. (El binario aún se decía `horusec`; el rename a `horusx` llega en v2.12.0.)

### Added
- `.goreleaser.yml` — build matrix `darwin/arm64`, `linux/amd64`, `linux/arm64` + checksums.
- `install.sh` — downloader de 3 rutas: `gh release download` → GitHub API → `github.com/.../releases/download`.
- `.github/workflows/release.yml` — dispara goreleaser al push de tag `v*`.

### Verified
- Local: `goreleaser release --clean` → 3 binaries + checksums.
- E2E: binario descargado + `horusec version` reporta 2.11.0, commit correcto.

---

## [v2.10.1] — Docker images publicadas

**Fecha**: 2026-10-21 · **Tag**: `v2.10.1`
**Images**: `aledj02/horusx-{python,go,shell,js}:{v2.10.1,latest}` en Docker Hub.

### Added
- 4 Dockerfiles: `docker/python/Dockerfile`, `docker/go/Dockerfile`, `docker/shell/Dockerfile`, `docker/js/Dockerfile`.
- Publish multi-arch (amd64 + arm64) en Docker Hub.

### Verified
- Smoke: `docker run aledj02/horusx-{python,go,shell,js}:v2.10.1` responde.
- `internal/enums/images/images.go` actualizado con los 4 nuevos IDs.

### Not included
- Cosign / firma de imágenes (D5).

---

## [v2.10.0-fork.0] — Baseline (`873d4104`)

Este tag marca el **upstream** en el punto de fork. No hay cambios propios del fork todavía — es la referencia de "antes".

```
$ git rev-parse v2.10.0-fork.0
66181e7519040091f591bfb6926909db7fb347bc   # (peel del objeto de tag) → commit 873d4104
```

| Campo  | Valor                                            |
|---|---|
| commit | `873d4104a6aa89be8f86d93db8e416909d9add87` |
| msg   | `start:bugfix wrong interpretation of subversion of docker (#1156)` |
| Go    | 1.17 |
| docker SDK | 20.10.9 |

---

## Proceso de cambios (audit trail)

Todo cambio relevante quedó registrado como:
- **Issue/Plan**: `migration-plan.md` (cada fase con checkboxes `[ ]`/`[x]` y motivo de decisión).
- **Verificación**: al menos 1 test Go por funcionalidad + 1 E2E manual (Docker tools corriendo, o LIGERO si Docker no disponible).

### Decisiones (resumen — ver `migration-plan.md` §Decisiones firmadas)
| # | Decisión |
|---|---|
| D1 | **LIGERO**: custom-rules JSON + built-ins vacíos; no forkear `horusx-engine` |
| D2 | Python/Go/Shell: rebuild de Docker images (Fase 5) |
| D3 | Fork: privado |
| D4 | Images: multi-arch (amd64 + arm64) |
| D5 | Sin cosign / sin firma de imágenes |
| D6 | Catálogo de reglas custom JSON: lo arma el equipo |
| D7 | Sin CWE prioritario — cobertura general balanceada |
| D8 | Prefix images: `horusx/` |
| D9 | Registry destino: Docker Hub público (user `horusx`) |
| D10 | CI: script shell que llama al CLI — sin integración a cambiar |
| D11 | docker/docker: `v20.10.9` → **`v24.0.9+incompatible`** (no `v27.x`, ver motivo en §Changed) |

### Formato de IDs de regla
- Pattern: `HS-{LANG}-\d+` (uppercase, `LANG` ∈ `PYTHON | GO | SHELL | SWIFT`, `\d+` 100+ para evitar colisión con built-in).
- **Uppercase**: el devkit normaliza a `String(language)` + `ToUpper()` (ver `custom_rule.go:146-148`).
- Rango `100-115` por lenguaje (Python 16, Go 11, Shell 12, Swift 11).

### Formato de regla JSON
- **Array** de objetos (no dict con key `"rules"`).
- Campos: `id`, `name`, `description`, `language`, `severity`, `confidence`, `type` (`Regular`|`OrMatch`|`AndMatch`|`NotMatch`), `expressions` (RE2).
- `Regular` y `OrMatch` son equivalentes (OR) — `AndMatch` = todas deben matchear (1 finding por file). `NotMatch` = reporta files que no matchean (raro uso).
- **RE2 (Go regexp)**: soporta `(?i)`, `\b`, `\d`, `\w`, `(?m)`, `[^\s]`, `(?:...)`, `{n,m}`. **No soporta** lookahead/lookbehind/backrefs/atomic/possessive. 50 reglas validadas con `regexp.Compile`.

---

## Cómo generar el próximo tag

1. Asegurar `go build ./...` + `go test ./...` verdes.
2. Bump de version (SemVer, SIN suffix `-fork`):
   - **`v2.10.1`** (Fase 5 — Docker images `aledj02/horusx-{python,go,shell,js}:v2.10.1`)
   - `v2.11.0` (si el próximo release incluye nuevas reglas JSON u otras features)
   - `v2.10.2` (fixes menores)
3. ```bash
   git tag v2.10.1
   git push origin v2.10.1
   ```
4. Append un bloque nuevo arriba de este changelog (formato `[v2.10.1]` con fecha).
5. **Docker images** (cuando se pushen): `aledj02/horusx-{python,go,shell,js}:v2.10.1` + `:latest`.
