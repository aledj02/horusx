# horusex custom rules

Catálogo de reglas SAST adicionales para **Python, Go, Shell y Swift** — lenguajes que **no tienen built-in rules** en el repo y dependen de Docker images viejas (2021) del upstream.

Estas reglas se registran vía `--custom-rules-path` (ver [customRule JSON schema](#custom-rule-json-schema)).

## Usecase rápido

```bash
# Punto único (todos los lenguajes)
horusec analyze \
  --custom-rules-path ./custom_rules/horusec-custom-rules.json \
  --path ./myservice

# Punto por lenguaje (elige la lista según código del repo)
horusec analyze --custom-rules-path ./custom_rules/python.json  # repo Python
horusec analyze --custom-rules-path ./custom_rules/go.json      # repo Go
horusec analyze --custom-rules-path ./custom_rules/shell.json   # repo de scripts
horusec analyze --custom-rules-path ./custom_rules/swift.json   # repo iOS/macOS
```

O vía **config**:

```json
// horusec-config.json
{
  "custom_rules_path": "./custom_rules/horusec-custom-rules.json"
}
```

O vía **env**:

```
export HORUSEC_CLI_CUSTOM_RULES_PATH=./custom_rules/horusec-custom-rules.json
```

## Reglas por lenguaje (50 total)

| Lenguaje  | Reglas | ID range |
|-----------|:---:|----------|
| Python    | 16   | HS-PYTHON-100 … HS-PYTHON-115 |
| Go        | 11   | HS-GO-100 … HS-GO-110 |
| Shell     | 12   | HS-SHELL-100 … HS-SHELL-111 |
| Swift     | 11   | HS-SWIFT-100 … HS-SWIFT-110 |

**IDs `HS-SWIFT-100+`** por higiene: los built-in `HS-SWIFT-1..24` existen en `internal/services/engines/swift/rules.go`; usar `100+` evita colisiones de "duplicates" (ver `custom_rule.go:146-148` — `HS-SWIFT` no se valida contra builtins, pero conviene no chocar).

## Custom-rule JSON schema

Cada regla es un objeto en un **array**:

```json
{
  "id": "HS-PYTHON-100",
  "name": "Deserialización insegura de pickle",
  "description": "pickle.loads / pickle.load puede ejecutar código remoto arbitrario. CWE-502.",
  "language": "Python",
  "severity": "CRITICAL",
  "confidence": "HIGH",
  "type": "Regular",
  "expressions": [
    "\\bpickle\\.loads\\s*\\(",
    "\\bpickle\\.load\\s*\\("
  ]
}
```

| Campo         | Tipo      | Valores / notas                                                                                              |
|---------------|-----------|--------------------------------------------------------------------------------------------------------------|
| `id`          | string    | `HS-{LANG}-\d+`. **Uppercase**: `HS-PYTHON-N`, `HS-GO-N`, `HS-SHELL-N`, `HS-SWIFT-N`.                        |
| `name`        | string    | Libre. Recomendado: título corto, sin CWE.                                                                  |
| `description` | string    | Libre. Recomendable: 1-2 frases + CWE ref (`CWE-94`, `CWE-798`).                                            |
| `language`    | string    | `Python` \| `Go` \| `Shell` \| `Swift` — exacto, case-sensitive (valor `ToString()` del enum devkit).        |
| `severity`    | string    | `CRITICAL` \| `HIGH` \| `MEDIUM` \| `LOW` \| `INFO` \| `UNKNOWN` (uppercase).                                |
| `confidence`  | string    | `HIGH` \| `MEDIUM` \| `LOW` (uppercase).                                                                     |
| `type`        | string    | `Regular` \| `OrMatch` \| `AndMatch` \| `NotMatch`. **`Regular` y `OrMatch` son equivalentes** (OR). **`AndMatch`** = todas las `expressions` deben matchear (reporta 1 finding por file). **`NotMatch`** = reporta files que NO matchean (no recomendado). |
| `expressions` | `string[]`| 1 o más regex **RE2 (Go)**. Sin lookahead, sin backrefs. Sin `\d` en id.                                     |

## Reglas RE2 (Go) — soporte

**Soporta**: `(?i)` case-insensitive, `\b` word boundary, `\d`, `\w`, `(?m)` multiline, `[^\s]`, `(?:...)` non-capture, `(?:...)|` OR, `{n,m}` quantifier, lookbehind negativo `(?<!\w)` (útil p/ `eval\s*\(` sin confundir con `revert`).

**NO soportado**: lookahead `(?=)`, lookbehind general `(?<=)`, backrefs `\1`, atomic groups `(?>)`, possessive quantifier `++`.

## Falsos positivos

Son **regex**, no AST. Algunas reglas pueden matchear en tests, docs, o código de fixture. Recomendaciones:

- Filtrar por severity antes de fail el CI:
  ```bash
  # solo CRITICAL + HIGH fallan el build
  horusec analyze --custom-rules-path ... --severity HIGH --severity CRITICAL --failure
  ```
- Ajustar `confidence` a `MEDIUM` si el patrón es ambiguo (el report lo muestra con flag "medium confidence").
- Para desactivar una regla puntual: moverla a un archivo que no esté en `--path`.

## Agregar una regla nueva

1. Copia el template:
   ```bash
   cp custom_rules/python.json custom_rules/python_new.json
   ```
2. Abre el JSON y agrega tu rule (IDs `HS-PYTHON-900+` si quieres mantener 100-115).
3. **Genera el combinado**:
   ```bash
   python3 - <<'PY'
   import json, os
   d = "custom_rules"
   out = []
   for f in ["python.json", "go.json", "shell.json", "swift.json"]:
       with open(os.path.join(d, f)) as fh:
           out.extend(json.load(fh))
   out.sort(key=lambda r: {"Python":0,"Go":1,"Shell":2,"Swift":3}[r["language"]])
   with open(os.path.join(d, "horusec-custom-rules.json"), "w") as fh:
       json.dump(out, fh, indent=2)
       fh.write("\n")
   PY
   ```
4. Smoke-test:
   ```bash
   horusec analyze --path /tmp/vulnerable --custom-rules-path custom_rules/horusec-custom-rules.json
   ```

## Proveniencia

Estas reglas son **ad-hoc del fork privado `horusx/`**, no del upstream. El upstream de `ZupIT/horusec` no las distribuye. Si se adoptan al pipeline central, mantener un índice de CWEs cubiertos (ver `description` de cada regla).
