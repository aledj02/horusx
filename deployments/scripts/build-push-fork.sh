#!/usr/bin/env bash
# fork aledj02/horusx — build + push de las 8 imágenes de tools que antes venían de horuszup/horusec-*
#
# Uso:
#   bash deployments/scripts/build-push-fork.sh                 # build+push las 8 con tag v2.12.0 + latest
#   TAG=v9.9.9 bash deployments/scripts/build-push-fork.sh      # tag custom
#   PUSH=0 bash deployments/scripts/build-push-fork.sh          # sólo build local (no push)
#
# Requiere: docker daemon corriendo, `docker login` a `aledj02`.
set -euo pipefail

TAG="${TAG:-v2.12.0}"
PUSH="${PUSH:-1}"
REGISTRY="aledj02"

# lang -> ruta del Dockerfile (build context = raíz del repo, porque leaks/ COPY un .toml relativo)
dockerfile_for() {
  case "$1" in
    c)       echo "internal/services/formatters/c/deployments/Dockerfile" ;;
    csharp)  echo "internal/services/formatters/csharp/deployments/Dockerfile" ;;
    elixir)  echo "internal/services/formatters/elixir/deployments/Dockerfile" ;;
    generic) echo "internal/services/formatters/generic/deployments/Dockerfile" ;;
    hcl)     echo "internal/services/formatters/hcl/deployments/Dockerfile" ;;
    leaks)   echo "internal/services/formatters/leaks/deployments/Dockerfile" ;;
    php)     echo "internal/services/formatters/php/deployments/Dockerfile" ;;
    ruby)    echo "internal/services/formatters/ruby/deployments/Dockerfile" ;;
    *)       echo "" ;;
  esac
}

for lang in c csharp elixir generic hcl leaks php ruby; do
  dockerfile="$(dockerfile_for "$lang")"
  if [ -z "$dockerfile" ] || [ ! -f "$dockerfile" ]; then
    echo "ERROR: Dockerfile no encontrado para $lang ($dockerfile)" >&2
    exit 1
  fi
  img="$REGISTRY/horusx-$lang:$TAG"
  echo "=================================================================="
  echo ">>> [$lang] building $img  (Dockerfile: $dockerfile)"
  echo "=================================================================="
  docker build -t "$img" -f "$dockerfile" .

  if [ "$PUSH" = "1" ]; then
    docker tag "$img" "$REGISTRY/horusx-$lang:latest"
    docker push "$REGISTRY/horusx-$lang:$TAG"
    docker push "$REGISTRY/horusx-$lang:latest"
  else
    docker tag "$img" "$REGISTRY/horusx-$lang:latest"
  fi
done

echo "=================================================================="
echo "DONE. Imágenes (registry $REGISTRY):"
for lang in c csharp elixir generic hcl leaks php ruby; do
  echo "  $REGISTRY/horusx-$lang:$TAG (+ :latest)"
done
echo "=================================================================="
echo "Siguiente paso: actualizar internal/enums/images/images.go a apuntar a $REGISTRY/horusx-*, "
echo "rebuild del CLI, re-release, y E2E con un proyecto de cada lenguaje."
