#!/usr/bin/env sh
# aledj02/horusx — install the latest (or a pinned) horusx binary from GitHub Releases.
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/aledj02/horusx/main/install.sh | sh
#   curl -fsSL .../install.sh | sh -s -- --version v2.10.1
#   curl -fsSL .../install.sh | sh -s -- --bin-dir /usr/local/bin --version v2.10.1
#
# Env overrides:
#   HORUX_VERSION   same as --version
#   HORUX_BIN_DIR   same as --bin-dir (default: first writable of ~/.local/bin, /usr/local/bin, /usr/bin)
#   HORUX_REPO      default aledj02/horusx
set -eu

REPO="${HORUX_REPO:-aledj02/horusx}"
BIN_NAME="horusx"
VERSION=""
BIN_DIR="${HORUX_BIN_DIR:-}"
OS="$(uname -s)"; ARCH="$(uname -m)"

case "$OS" in
  Darwin) OS=darwin ;;
  Linux)  OS=linux  ;;
  *) echo "unsupported OS: $OS" >&2; exit 1 ;;
esac
case "$ARCH" in
  x86_64|amd64) ARCH=amd64 ;;
  arm64|aarch64) ARCH=arm64 ;;
  *) echo "unsupported ARCH: $ARCH" >&2; exit 1 ;;
esac

while [ $# -gt 0 ]; do
  case "$1" in
    --version)  VERSION="${2:?--version needs a value}"; shift 2 ;;
    --bin-dir)  BIN_DIR="${2:?--bin-dir needs a value}"; shift 2 ;;
    *) echo "unknown arg: $1" >&2; exit 1 ;;
  esac
done
[ -n "${HORUX_VERSION:-}" ] && VERSION="$HORUX_VERSION"

if [ -z "$VERSION" ]; then
  VERSION="$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep -o '"tag_name"[[:space:]]*:[[:space:]]*"[^"]*"' | head -n1 | sed -E 's/.*"([^"]+)"$/\1/')"
  [ -n "$VERSION" ] || { echo "could not resolve latest release tag" >&2; exit 1; }
fi
case "$VERSION" in v*) ;; *) VERSION="v${VERSION}" ;; esac

ASSET="${BIN_NAME}_${VERSION#v}_${OS}_${ARCH}.tar.gz"
AUTH_TOKEN="${GITHUB_TOKEN:-${GH_TOKEN:-}}"

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

download_asset() {
  if command -v gh >/dev/null 2>&1 && gh auth status >/dev/null 2>&1; then
    echo ">> downloading via gh from ${REPO} tag ${VERSION}"
    gh release download "$VERSION" --repo "$REPO" --pattern "$ASSET" --dir "$TMP"
  elif [ -n "$AUTH_TOKEN" ]; then
    echo ">> downloading via GitHub API (auth) from ${REPO}"
    ASSET_ID="$(curl -fsSL -H "Authorization: token $AUTH_TOKEN" \
        "https://api.github.com/repos/${REPO}/releases/tags/${VERSION}" \
        | python3 -c "import sys,json; print(next(a['id'] for a in json.load(sys.stdin)['assets'] if a['name']=='${ASSET}'))" 2>/dev/null)"
    [ -n "$ASSET_ID" ] || { echo "no matching asset ${ASSET} found" >&2; return 1; }
    curl -fL --progress-bar -H "Authorization: token $AUTH_TOKEN" \
        -H "Accept: application/octet-stream" \
        -o "$TMP/$ASSET" "https://api.github.com/repos/${REPO}/releases/assets/${ASSET_ID}"
  else
    echo ">> downloading from https://github.com/${REPO}/releases/download/${VERSION}/${ASSET}"
    curl -fL --progress-bar -o "$TMP/$ASSET" \
        "https://github.com/${REPO}/releases/download/${VERSION}/${ASSET}"
  fi
}
download_asset
tar -xzf "$TMP/$ASSET" -C "$TMP"

if [ -z "$BIN_DIR" ]; then
  for d in "$HOME/.local/bin" /usr/local/bin /usr/bin; do
    if [ -w "$d" ] || [ "$(id -u)" = "0" ]; then BIN_DIR="$d"; break; fi
  done
  [ -n "$BIN_DIR" ] || { echo "no writable bin dir found; try --bin-dir" >&2; exit 1; }
fi
mkdir -p "$BIN_DIR"
install -m 0755 "$TMP/$BIN_NAME" "$BIN_DIR/$BIN_NAME"

echo ">> installed $BIN_NAME $VERSION -> $BIN_DIR/$BIN_NAME"
"$BIN_DIR/$BIN_NAME" version || true
