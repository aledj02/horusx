# Plan de Modernización — Horusx Fork Interno (`horusx`)

> **Estado**: Draft para revisión
> **Fecha**: 2026-10-20
> **Upstream**: [ZupIT/horusx](https://github.com/ZupIT/horusec) (última release estable: v2.9.0-beta.3, ago/2023)
> **Fork**: privado, prefix de images `horusx/*`
> **Objetivo**: SAST CLI interno para JS, TS, Java, Go, Python, SH, Kotlin, Swift — con reglas actualizadas a 2026

## Progreso real (2026-10-20, sin commits aún)

| Fase | Estado | Evidencia |
|---|---|---|
| 0 Fork + tag | ✅ | `v2.10.0-fork.0` en `873d4104`, branch `main` |
| 1 Deps Go | ✅ | `go.mod`: `go 1.24`, `docker v24.0.9+incompatible`; `go build ./...` verde |
| 2 Custom rules map | ✅ | `internal/services/custom_rules/{custom_rule.go,service.go}` soporta `Go`, `Python`, `Shell`, `Swift`; tests `ok` |
| 3 Reglas JSON | ✅ | `custom_rules/{python,go,shell,swift}.json` = 50 reglas, RE2-validado; `horusx-custom-rules.json` combinado; `custom_rules/README.md` |
| 4 Wiring P/Go/Shell | ✅ | Nuevos: `engines/{python,go,shell}/rule_manager.go` + `formatters/{python/horusxpython,go/horusxgo,shell/horusxshell}/formatter.go` + `runner.go` los llama en `detectVulnerability{Python,Go,Shell}` ANTES de los Docker tools |
| 5 Docker rebuild | ⏸ | Pausado (requiere `horusx/*` en Docker Hub, out-of-scope sin credentials) |
| 6 Testing local | ✅ | E2E: `vuln.py/vuln.go/vuln.sh/App.swift` → disparan `HS-PYTHON-100`, `HS-GO-106`, `HS-SWIFT-101`, `HS-SHELL-100`, `B301`, `B403`, `G104`, `HS-LEAKS-25` — status `success`, 9 vulns; `go test` en `custom_rules` + `analyzer` → `ok`; los fails en `formatters/ruby/bundler` son pre-existentes (submodules `examples/*` no inicializados) |
| 6.1 Rule-manager tests | ✅ | `engines/{python,go,shell}/rule_manager_test.go` — `NewRules()` empty + extensions correct; `go test` → `ok` |
| 6.2 Formatter tests | ✅ | `formatters/default_engine_formatter_test.go` ampliada: 12 formatters (antes 9); `go test` → `ok` |
| 6.3 E2E test suite | ✅ | `custom_rules/e2e_test.go` — `TestE2E_CustormRulesDischargeFixture`: 4 fixtures → 4 findings (HS-PYTHON-100, HS-GO-100, HS-SHELL-100, HS-SWIFT-101); `go test -run TestE2E` → PASS |
| 6.4 Fixtures | ✅ | `examples/{python,go,shell,swift}/example1/` — 1 file vulnerable por lengua |
| 7 Distribución | 🔶 | `CHANGELOG-FORK.md` ✅ · `README-FORK.md` ✅ · goreleaser / install.sh ⏸ (bloqueo de Fase 5) |

## Decisiones firmadas

| # | Decisión |
|---|---|
| D1 | Estrategia regexps: **LIGERA** (custom-rules JSON + built-ins; NO forkear `horusx-engine`) |
| D2 | Python/Go/Shell: **rebuild de Docker images** con herramientas actualizadas (Fase 5) |
| D3 | Fork: **privado** |
| D4 | Images: **multi-arch** (amd64 + arm64) |
| D5 | Sin cosign / sin firma de imágenes |
| D6 | Catálogo de reglas custom JSON: **lo arma el equipo** con esta revisión como insumo |
| D7 | No hay CWE prioritario específico: cobertura general balanceada |
| D8 | Prefix de images: **`horusx/`** |
| D9 | Registry destino: **Docker Hub público** (user `horusx`) |
| D10 | CI: script shell que llama al CLI — sin integración que cambiar |
| D11 | docker/docker: `v20.10.9` → **`v24.0.9+incompatible`** (no `v27.x`). Motivo: v25+ movió los tipos de `api/types` a sub-paquetes (`container/image/registry`), obligando a renombrar decenas de usages, y además arrastra otel/grpc/genproto a la cadena de deps. `v24` ya cubre el objetivo de modernizar sin ese costo adicional. |

## Stack objetivo

| Componente | Antes | Después |
|---|---|---|
| Go | 1.17 | **1.24** |
| docker/docker | 20.10.9 | **24.0.9+incompatible** (ver D11) |
| horusx-engine | v1.0.1 | latest |
| horusx-devkit | v1.0.24 | latest |
| bandit (Python) | 1.7.4 | **1.9.4** |
| safety (Python) | 1.10.3 | **3.8.1** |
| gosec (Go) | 2.11.0 | **v2.29.0** |
| nancy (Go) | 1.0.33 | **v2.1.0** (verificar URL de descarga) |
| shellcheck (SH) | >0.7.2 | **v0.11.0** |
| Node (JS) | 17.6.0 | **20.x LTS** |

---

## Baseline — Estado actual por lenguaje

| Lenguaje | Reglas built-in (repo) | Image Docker | Custom JSON |
|---|:---:|:---:|:---:|
| JS | 55 | `horuszup/horusx-js:v1.2.0` (jq + npm/yarn audit) | ✅ |
| TS | 55 (comparte engine JS, `.js/.ts/.jsx/.tsx`) | ídem | ✅ (via JS) |
| Java | 152 | — | ✅ |
| Kotlin | 40 (hereda JVM) | — | ✅ |
| Swift | 24 | — | ❌ |
| Python | — | bandit 1.7.4 + safety 1.10.3 (2021) | ❌ |
| Go | — | gosec 2.11 + nancy 1.0.33 (2021) | ❌ |
| Shell | — | shellcheck 0.7.2 (2021) | ❌ |
| Leaks (multi-lang) | 28 | — | ✅ |

Built-in total hoy: **~403 reglas** (Java 152, C# 74, JS 55, JVM 40, Leaks 28, Swift 24, Dart 17, K8s 9, Nginx 4).

---

## FASES

### Fase 0 — Fork + baseline tag
- [ ] 0.1 Fork privado de `ZupIT/horusx`
- [ ] 0.2 Tag `v2.9.1-fork.0` en HEAD pre-cambios (baseline reproducible)
- [ ] 0.3 Crear `CHANGELOG-FORK.md` (log de auditoría de cambios)
- **Verificar**: `git log -1`, `git tag -l`
- **Estimado**: 0.5h

### Fase 1 — Modernizar dependencias Go
- [ ] 1.1 `go.mod`: `go 1.17` → `go 1.24`
- [ ] 1.2 `go get github.com/docker/docker@v27 @latest`
- [ ] 1.3 `go get github.com/ZupIT/horusec-engine@latest github.com/ZupIT/horusec-devkit@latest`
- [ ] 1.4 `go mod tidy`
- [ ] 1.5 `go build ./...` — resolver breaking changes (docker SDK 27 movió `types.Container*`)
- [ ] 1.6 `go test ./...` — todo verde
- [ ] 1.7 `make lint` / `golangci-lint run`
- [ ] 1.8 Build binario: `goreleaser build --snapshot` (o `go build -o horusx .`)
- **Riesgo**: breaking API en docker SDK y devkit. Mitigar con `go build` iterativo + bisect.
- **Estimado**: 1-2 días

### Fase 2 — Extender custom-rules map (4 nuevos lenguajes)
- [ ] 2.1 `internal/services/custom_rules/service.go:43-52` — agregar al map:
  ```go
  languages.Python, languages.Go, languages.Shell, languages.Swift,
  ```
- [ ] 2.2 Verificar que el enum `languages.Language` del devkit tiene esas 4 keys (confirmado en investigación previa: Python, Go, Shell, Swift existen)
- [ ] 2.3 Grep de `switch` sobre `languages.` en `internal/` — verificar que no haya path que descarte custom rules por language
- [ ] 2.4 Test unitario: 1 regla JSON de prueba por cada lenguaje nuevo
- **Riesgo**: path oculto que hardcodea la whitelist. Mitigar con grep exhaustivo de `customRules`, `Languages`, `IsCustomRuleAvailable`.
- **Estimado**: 0.5-1 día

### Fase 3 — Catálogo de reglas custom JSON
Carpeta `rules/` con un JSON por lenguaje. El equipo arma el contenido; este plan provee el **template + seed de reglas mínimas**.

Template (formato existente en `custom_rule.go`):
```json
{
  "language": "python",
  "rules": [
    {
      "id": "HSX-PY-101",
      "name": "exec() usage",
      "description": "exec() allows arbitrary code execution. CWE-94.",
      "severity": "Critical",
      "confidence": "Medium",
      "type": "text",
      "expressions": ["(?i)\\bexec\\s*\\("]
    }
  ]
}
```

Seed por lenguaje (equipo puede extender):

**`rules/python.json`** (~20 reglas):
- exec/eval/compile
- `pickle.load` / `pickle.loads` (CWE-502)
- `yaml.load` sin `SafeLoader` (CWE-502)
- `subprocess.*shell=True` (CWE-78)
- `os.system` (CWE-78)
- `MD5` / `sha1` hash (CWE-328)
- `random` para crypto (CWE-338) — `random.randint`, `random.random` en contexto de token
- `http.client` / `urllib` sin TLS (CWE-319)
- `assert` como control de acceso
- Hardcoded `password` / `secret` / `token` (CWE-798)
- `input()` → `eval` (CWE-94)
- `pickle`, `marshal`, `shelve`
- `__import__`
- `multiprocessing` con `fork` + `pickle` (CWE-502)
- `sqlite3` con f-string (CWE-89)
- `urllib2`/`httplib` legacy (CWE-319)
- `base64.b64encode` como 'seguridad' (CWE-326)
- `tempfile.mktemp` (CWE-377)
- `tarfile` sin `filter` param (CWE-22 path traversal)

**`rules/go.json`** (~15 reglas):
- Hardcoded creds en `:= "..."` con `password|secret|token|api_key`
- `log.Printf` con `password|secret` (CWE-532)
- `exec.Command` con string user-built (CWE-78)
- `crypto/md5` / `crypto/sha1` para auth (CWE-328)
- `math/rand` para tokens/security (CWE-338) — `rand.Int` sin `crypto/rand`
- `http.Client` sin `TLSConfig` (CWE-319)
- `sql.Query` con f-string/concat (CWE-89)
- `ioutil.ReadFile` en `*_test.go` de credenciales
- `base64.StdEncoding.DecodeString` de user input
- `text/template` con user input (CWE-94)
- `net/http` con `InsecureSkipVerify: true` (CWE-295)
- `flag.*` de password sin warning
- `//go:generate` de un script remoto (CWE-494)
- `http.ListenAndServe` sin `TLSConfig` en rutas admin

**`rules/shell.json`** (~12 reglas):
- `curl ... | sh` / `wget ... | bash` (CWE-494)
- `eval` con variable de entrada
- `chmod 777` (CWE-732)
- `ssh -o StrictHostKeyChecking=no` (CWE-322)
- `sudo` sin `visudo` (CWE-250)
- `mktemp` sin template (CWE-377)
- Hardcoded password: `password=`, `passwd=`, `--password`
- `rm -rf /` sin anchor
- `set +e` o `trap '' ERR` (CWE-754, false positive alto)
- `echo $PASSWORD` a stdout (CWE-532)
- `ssh-keygen` con `-N ''` (CWE-798)
- `chmod 666` en `.env` / `id_rsa` (CWE-732)

**`rules/swift.json`** (~10 reglas que complementen las 24 built-in):
- `URLSession` sin `kCFNetworkSSLAllowsExpiredCertificates`
- `NSLog` con `password|token|secret` (CWE-532)
- `KVC` keyPath injection (CWE-472)
- `UserDefaults` con sensitive (CWE-312)
- `keychain` access sin `kSecAttrAccessibleWhenPasscodeSetThisDeviceOnly`
- `UIApplication` `openURL` sin allowlist (CWE-601)
- `JIT` / `mmap` con `MAP_JIT` (CWE-1336)
- `NSArchiver`/`NSUnarchiver` sin allowlist (CWE-502, iOS 12+)
- `WKWebView` `loadHTMLString` con user input (CWE-79)
- `Bundle` `objectForInfoDictionaryKey` de un bundle no-own (CWE-494)

- [ ] 3.1 `rules/python.json`
- [ ] 3.2 `rules/go.json`
- [ ] 3.3 `rules/shell.json`
- [ ] 3.4 `rules/swift.json`
- [ ] 3.5 Smoke test: `horusx start --custom-rules-path rules/ -p examples/`
- **Estimado**: 2-3 días (seed en 1d, refinamiento del equipo en 1-2d)

### Fase 4 — Refresh de reglas Swift built-in
- [ ] 4.1 Revisar `internal/services/engines/swift/rules.go` (24 reglas, IDs HS-SWIFT-*)
- [ ] 4.2 Ajustar 5-10 regexps obsoletas (patrones Swift 4 → 5.10)
- [ ] 4.3 Añadir 5-10 nuevas donde custom JSON no alcance (p. ej. multi-línea)
- [ ] 4.4 `go test ./internal/services/engines/swift/...`
- **Estimado**: 0.5-1 día

### Fase 5 — Rebuild de imágenes Docker (prefix `horusx`)
**Docker Hub org `horusx`** (pública), images multi-arch (amd64+arm64).

- [ ] 5.1 **Python** — `internal/services/formatters/python/deployments/Dockerfile`:
  ```dockerfile
  FROM python:3.12.7-alpine3.20
  RUN apk add --no-cache jq
  RUN pip install --no-cache-dir safety==3.8.1 bandit==1.9.4
  COPY ./internal/services/formatters/python/deployments/horusx-file-ignore.sh /usr/local/bin/horusx-file-ignore.sh
  ```
- [ ] 5.2 **Go** — `internal/services/formatters/go/deployments/Dockerfile`:
  ```dockerfile
  FROM golang:1.24.5-alpine
  RUN apk add --no-cache curl
  # nancy: verificar URL de binario para v2.1.0 (formato de release puede haber cambiado)
  RUN curl -fsSL https://github.com/sonatype-nexus-community/nancy/releases/download/v2.1.0/<bin>-linux-amd64 -o /bin/nancy && chmod +x /bin/nancy
  RUN wget -qO - https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s v2.29.0
  ```
  > ⚠️ Verificar nombre exacto del asset en GitHub Release v2.1.0 de nancy antes del build.
- [ ] 5.3 **Shell** — `internal/services/formatters/shell/deployments/Dockerfile`:
  ```dockerfile
  FROM alpine:3.21
  RUN apk add --no-cache shellcheck
  # shellcheck > v0.11.0 (última)
  ```
- [ ] 5.4 **JS/TS** — `internal/services/formatters/javascript/deployments/Dockerfile`:
  ```dockerfile
  FROM node:20.18.0-alpine3.21
  RUN apk add --no-cache jq
  RUN npm install -g npm@10.8.2
  ```
- [ ] 5.5 Multi-arch build con `docker buildx`:
  ```bash
  docker buildx create --name horusx-builder --use
  docker buildx build --platform linux/amd64,linux/arm64 \
    -t horusx/horusx-python:v1.0.1-fork.1 \
    -t horusx/horusx-python:latest \
    -f internal/services/formatters/python/deployments/Dockerfile . --push
  ```
  (repetir para go, shell, js)
- [ ] 5.6 Push a Docker Hub `horusx` (login: `docker login`)
- [ ] 5.7 **Actualizar `internal/enums/images/images.go`**:
  ```go
  Python     = "horusx/horusx-python:v1.0.1-fork.1"
  Go         = "horusx/horusx-go:v1.0.2-fork.1"
  Shell      = "horusx/horusx-shell:v1.0.2-fork.1"
  Javascript = "horusx/horusx-js:v1.3.0-fork.1"
  ```
  (números de version se eligen para no colisionar con tags upstream `v1.0.1`)
- [ ] 5.8 **Verificar compatibilidad de parsers Go** con JSON de herramientas nuevas:
  ```bash
  docker run --rm horusx/horusx-python:v1.0.1-fork.1 bandit -r /sample -f json
  # comparar schema contra formatters/python/bandit/*.go y safety/*.go
  ```
  - bandit 1.9.4: schema estable (changes: `test_id` rename en algunas reglas)
  - safety 3.8.1: **breaking** — migrar de `safety check` CLI a `safety scan` CLI; **verificar parser**
  - gosec v2.29.0: JSON estable, `Results` key igual
  - shellcheck v0.11.0: JSON estable, fields `code`, `message`, `line`
  - npm audit 10.x: JSON estable (`advisories` → `vulnerabilities` en audit de npm 9+)
- [ ] 5.9 Smoke E2E por lenguaje con image nueva
- [ ] 5.10 Documentar en `CHANGELOG-FORK.md`: tag, hash, tool versions
- **Riesgo**: safety 3.x breaking (comando). Mitigar: leer source del parser `internal/services/formatters/python/safety/` antes del build.
- **Estimado**: 2-3 días

### Fase 6 — Testing
- [ ] 6.1 `go test ./...`
- [ ] 6.2 `go vet ./...`
- [ ] 6.3 `golangci-lint run`
- [ ] 6.4 E2E por lenguaje (1 project con vulnerabilidad conocida por language):
  - `examples/example-python` + `--custom-rules-path rules/python.json`
  - `examples/example-go` + `--custom-rules-path rules/go.json`
  - `examples/example-shell` (o uno nuevo)
  - `examples/example-javascript` / TS
  - `examples/example-java`
  - `examples/example-swift`
  - `examples/example-kotlin` (o Java compartido)
- [ ] 6.5 Regression tests de custom rules JSON — 1 regla por lenguaje, 1 file vulnerable, expect 1 finding
- **Estimado**: 1-2 días

### Fase 7 — Distribución (fork **privado**)
- [ ] 7.1 `goreleaser build --snapshot` → artefacto `dist/horusx_darwin_arm64`, `dist/horusx_linux_amd64`, `dist/horusx_linux_arm64`
- [ ] 7.2 **Fork privado en GitHub**: publicar binario como GitHub Release Asset (sin cosign)
- [ ] 7.3 Docker images ya están en Docker Hub público `horusx/*` (Fase 5)
- [ ] 7.4 `rules/` folder: ship como artefacto (JSON, no binario) → README-FORK
- [ ] 7.5 `README-FORK.md`:
    - build desde cero
    - how to update tool versions
    - how to add rules JSON
    - how to point `images.go` to private registry (si más adelante se cambia)
- [ ] 7.6 Install script: `curl -fsSL https://<fork>/install.sh | sh` (downloader del GH release privado)
- **Estimado**: 1 día

---

## Timeline total

| Fase | días |
|---|:---:|
| 0 Fork | 0.5 |
| 1 Deps Go | 1-2 |
| 2 Custom map | 0.5-1 |
| 3 Reglas JSON | 2-3 |
| 4 Swift refresh | 0.5-1 |
| 5 Docker + multi-arch | 2-3 |
| 6 Testing | 1-2 |
| 7 Distribución | 1 |
| **TOTAL** | **9-14 días** (~2-3 semanas) |

## Open items (resueltos)

| Pregunta | Respuesta |
|---|---|
| Fork público/privado? | **Privado** |
| Multi-arch? | **Sí** |
| Cosign? | **No** |
| Reglas JSON — quién? | **Equipos arma, este plan da template+seed** |
| CWE prioritario? | **Ninguno específico — cobertura balanceada** |
| Prefix images? | **`horusx/`** |
| Registry? | **Docker Hub público `horusx`** |

## Riesgos principales

1. **Docker SDK 27 breaking API** — mitigar con build iterativo
2. **safety 3.x changing CLI** — verificar el parser antes del build de la image
3. **Nancy v2.1.0 release asset naming** — confirmar URL exacta antes del build de Go image
4. **Custom-rules map whitelist oculta** — mitigar con grep exhaustivo en Fase 2
5. **Docker Hub rate limit** con 4 images × 2 arch = 8 pushes — mitigar con buildx cache + 1 build por image

## Cómo verificar "todo funciona" (Definition of Done final)

```bash
# Build
go build -o bin/horusx .

# Unit
go test ./...

# E2E por lenguaje (ejemplo Python)
docker pull horusx/horusx-python:v1.0.1-fork.1
bin/horusx start -p examples/example-python -t python --custom-rules-path rules/

# Expected: findings de bandit 1.9.4 (no 1.7.4) + reglas HSX-PY-*
grep -E "HSX-PY|bandit" examples/example-python/.horusx.json
```
