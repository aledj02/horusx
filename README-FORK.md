# horusx — fork interno de Horusx

> Fork privado de [ZupIT/horusx](https://github.com/ZupIT/horusec) para uso interno.
> Namespace de imágenes: **`aledj02/horusx-{python,go,shell,js}`**. Baseline tag upstream: `v2.10.0-fork.0` (commit `873d4104`). Tags propios desde Fase 5: **SemVer estándar** (ej. `v2.10.1`) — sin el suffix `-fork`.
> Objeto del fork: modernizar el toolchain **Go 1.24 + docker SDK v24**, habilitar **custom rules JSON** para **Python / Go / Shell / Swift**, y refresh de herramientas Docker (bandit, safety, gosec, nancy, shellcheck, node) cuando la Fase 5 se aplique.
>
> Changelog completo en [`CHANGELOG-FORK.md`](./CHANGELOG-FORK.md). Plan de trabajo de referenciaci[ó]n en [`migration-plan.md`](./migration-plan.md).

---

## 1. Build from scratch

### 1.1 Requisitos
- **Go 1.21+** (recomendado: **1.24**) — `go version`
- **Docker** con `buildx` (solo para Fase 5)
- `golangci-lint` (opcional, para lint)
- Submodules inicializados (solo si quieres los tests que usen `examples/`):
  ```bash
  git submodule update --init --recursive
  ```

> ⚠️ `examples/` es un **git submodule** (`horusx-examples-vulnerabilities`). Si no lo inicializas, los tests de `formatters/{csharp,dotnet_cli,elixir,generic,go/nancy,javascript,python/safety,ruby}` fallan (no son regresiones).

### 1.2 Compilar y correr tests

```bash
go mod download          # ya viene con go.sum resuelto para go 1.24
go build -o bin/horusx .
go test ./internal/services/custom_rules/...
go test ./internal/services/engines/{python,go,shell}/...
go test ./internal/services/formatters/
go test ./internal/controllers/analyzer/...
```

Binario: `./bin/horusx --help`.

### 1.3 Quick E2E (sin Docker)

```bash
# fixtures del repo (1 file vulnerable por lenguaje)
ls examples/{python,go,shell,swift}/example1/

# scan LIGERO — solo custom rules (no invoca Docker)
./bin/horusx start \
  -p examples/python/example1 \
  --custom-rules-path custom_rules/python.json \
  -o txt
```

Salida esperada en `examples/python/example1/.horusx.json`:

```json
"rule_id": "HS-PYTHON-100"
```

Para los 4 lenguages de golpe usa el combinado:
```
--custom-rules-path custom_rules/horusx-custom-rules.json
```

---

## 2. Reglas custom JSON

**Carpetas**: `custom_rules/` (NO `rules/` como se lee en algunos templates — se mantuvo `custom_rules/` por no romper la E2E del repo).

**Archivos**:
- `python.json` — 16 reglas (`HS-PYTHON-100` … `115`)
- `go.json` — 11 reglas (`HS-GO-100` … `110`)
- `shell.json` — 12 reglas (`HS-SHELL-100` … `111`)
- `swift.json` — 11 reglas (`HS-SWIFT-100` … `110`) — complementa los 24 built-in
- `horusx-custom-rules.json` — combinado de los 50

### 2.1 Schema (array de objetos, **no** dict)

```json
{
  "id": "HS-PYTHON-100",
  "name": "Deserialización insegura de pickle",
  "description": "pickle.loads / pickle.load puede ejecutar código remoto. CWE-502.",
  "language": "Python",
  "severity": "CRITICAL",
  "confidence": "HIGH",
  "type": "Regular",
  "expressions": ["\\bpickle\\.loads\\s*\\(", "\\bpickle\\.load\\s*\\("]
}
```

| Campo         | Tipo       | Notas                                                                                              |
|---|---|---|
| `id`          | `string`  | Uppercase. `HS-{LANG}-N`, `N ≥ 100` para evitar colisión con built-ins.                            |
| `language`    | `string`  | Uno de: `Python` \| `Go` \| `Shell` \| `Swift`. **Exacto** (case-sensitive `ToString()` del enum). |
| `severity`    | `string`  | `CRITICAL` \| `HIGH` \| `MEDIUM` \| `LOW` \| `INFO` \| `UNKNOWN`                                    |
| `confidence`  | `string`  | `HIGH` \| `MEDIUM` \| `LOW`                                                                         |
| `type`        | `string`  | `Regular` (default, OR) \| `OrMatch` (igual a `Regular`) \| `AndMatch` (todas deben matchear) \| `NotMatch` (reporta files que NO matchean — raro) |
| `expressions` | `string[]`| **RE2** (Go regexp). Sin lookahead (`(?=)`), sin backrefs (`\1`).                                  |

### 2.2 Agregar una regla

1. Copia el JSON del lenguaje y agrega el objeto al **array** (no dentro de un dict).
2. Regenerar el combinado:
   ```bash
   python3 - <<'PY'
   import json, os
   d = "custom_rules"
   out = []
   for f in ["python.json", "go.json", "shell.json", "swift.json"]:
       with open(os.path.join(d, f)) as fh:
           out.extend(json.load(fh))
   out.sort(key=lambda r: {"Python":0,"Go":1,"Shell":2,"Swift":3}[r["language"]])
   with open(os.path.join(d, "horusx-custom-rules.json"), "w") as fh:
       json.dump(out, fh, indent=2); fh.write("\n")
   PY
   ```
3. Smoke test:
   ```bash
   /usr/local/bin/horusx start -p /tmp/vulnerable \
     --custom-rules-path custom_rules/horusx-custom-rules.json
   ```

### 2.3 Validación de reglas RE2

Verificar que todas las expresiones compilen bajo Go `regexp` (RE2) **antes** de commit — si no, el service falla en runtime.

```bash
go test ./internal/services/custom_rules/ -run TestParseLanguageFile -v
```

(El caso `TestE2E_CustormRulesDischargeFixture` en `internal/services/custom_rules/e2e_test.go` también dispara una regla por lenguaje — es la prueba de que el JSON → engine → finding fluye.)

---

## 3. Point at private registry

El fork apunta a `horusx/*` en Docker Hub (Docker Hub público, user `horusx`), por D8/D9. Para apuntar a un **private registry** (si en el futuro se cambia):

### 3.1 `internal/enums/images/images.go`

```go
const (
    Python     = "aledj02/horusx-python:<tag>"
    Go         = "aledj02/horusx-go:<tag>"
    Shell      = "aledj02/horusx-shell:<tag>"
    Javascript = "aledj02/horusx-js:<tag>"
)
```

### 3.2 Credenciales

```bash
docker login <registry>
# docker/docker SDK lee ~/.docker/config.json — no hay flag para eso
```

### 3.3 Build & push (cuando se aplique Fase 5)

```bash
docker buildx create --name horusx --use
docker buildx build --platform linux/amd64,linux/arm64 \
  -t aledj02/horusx-python:v2.10.1 \
  -t aledj02/horusx-python:latest \
  -f internal/services/formatters/python/deployments/Dockerfile . \
  --push
```

Ver `migration-plan.md` Fase 5.1-5.7 para el template de las 4 imágenes.

---

## 4. Differences vs upstream (ZupIT/horusx)

| Ámbito | Upstream | Fork `horusx` |
|---|---|---|
| Go | 1.17 | **1.24** |
| docker SDK | 20.10.9 | **24.0.9+incompatible** |
| Custom rules | C#/Java/JS/Kotlin/Leak/Nginx/YAML | + **Python, Go, Shell, Swift** |
| Reglas Python | 0 built-in (solo bandit/safety viejas) | 16 JSON + (Docker tool cuando F5) |
| Reglas Go      | 0 built-in (solo gosec/nancy viejas)   | 11 JSON |
| Reglas Shell   | 0 built-in (solo shellcheck vieja)      | 12 JSON |
| Reglas Swift   | 24 built-in | 24 built-in + **11 JSON** |
| Docker images  | `horuszup/*` | `aledj02/horusx-*` (Fase 5 ✅ publicada) |
| Prefix de reg  | `HS-*` | `HS-*` (iguales) — los custom son `HS-{PY|GO|SH|SW}-100+` |
| Cosign / firma | Sí (upstream) | No (D5) |
| Fork           | `ZupIT/horusx` | `horusx/` (privado) |

---

## 5. Troubleshooting

| Symptom | Causa probable | Resolución |
|---|---|---|
| `lstat examples/<lang>/example1: no such file or directory` en tests | `examples/` submodule no inicializado | `git submodule update --init --recursive`, o ignorar esos tests (no son regresiones) |
| `no rules available for language` | `--custom-rules-path` no apunta al JSON correcto, o `language` del JSON no case-matches el enum | Verificar que `language` sea `Python`/`Go`/`Shell`/`Swift` (exacto, no `python`, no `golang`) |
| `invalid regular expression` | Expressions no RE2-validas (lookahead, backrefs) | Validar con `go test -run TestParseLanguageFile` |
| Docker tool (bandit/gosec/etc.) no dispara | Sin network, o image `horusx/*` no publicada | Fase 5 pendiente — hasta entonces solo disparan las custom rules. Los tools de Docker corren en paralelo (no excluyentes) |
| `pull access denied for horuszup/horusx-*` | Upstream no publicó el tag | `docker pull horuszup/<image>:<tag>` manualmente antes de `start` |
| `examples` submodule not found | Clon sin `--recursive` | `git submodule update --init --recursive` o clonar con `--recurse-submodules` |

---

## 6. Distribución (Fase 7)

### 6.1 Instalar

```bash
curl -fsSL https://raw.githubusercontent.com/aledj02/horusx/main/install.sh | sh
curl -fsSL https://raw.githubusercontent.com/aledj02/horusx/main/install.sh | sh -s -- --version v2.10.1
horusx version
```

### 6.2 Build desde cero (sin goreleaser)

```bash
V=$(git describe --tags --abbrev=0); C=$(git rev-parse --short HEAD); D=$(date -u +%Y-%m-%dT%H:%M:%SZ)
LF="-s -w -X github.com/ZupIT/horusec/cmd/app/version.Version=${V} -X github.com/ZupIT/horusec/cmd/app/version.Commit=${C} -X github.com/ZupIT/horusec/cmd/app/version.Date=${D}"
CGO_ENABLED=0 go build -trimpath -ldflags "$LF" -o horusx_darwin_arm64 ./cmd/app   # host
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "$LF" -o horusx_linux_amd64 ./cmd/app
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags "$LF" -o horusx_linux_arm64 ./cmd/app
```

### 6.3 Publicar release (goreleaser)

```bash
goreleaser release --clean            # lee la tag actual, genera bins+tar.gz, sube assets a GH
goreleaser build --snapshot --clean   # solo build local, no publica
```

Configuración en [`.goreleaser.yml`](./.goreleaser.yml): 3 plataformas (`darwin/arm64`, `linux/amd64`, `linux/arm64`), `CGO_ENABLED=0`, version inyectada por ldflags, `custom_rules/` embebido en cada archive.

### 6.4 Custom rules JSON

Los JSON viven en [`custom_rules/`](./custom_rules/) y se embeben en cada release. Para añadir reglas: editar `custom_rules/{python,go,shell,swift}.json` (array de objetos, campo `id`, regex RE2, IDs `HS-{PY|GO|SH|SW}-100+`) y usar con `horusx start -c custom_rules/python.json`.

### 6.5 Update tool versions

Reescribir el Dockerfile por lenguaje en `internal/services/formatters/<lang>/deployments/Dockerfile` (versines pinadas), re-build + push con `docker buildx build --platform linux/amd64,linux/arm64 --push`, y bump de `images.go` si cambia tag. Ver §3.

---

## 7. Licencia

Apache-2.0 (igual que upstream).
