#!/bin/sh
# Installs the latest gcli release for Linux or macOS.
#
#   curl -fsSL https://raw.githubusercontent.com/daflecardoso/gcli/main/install.sh | sh
#
# Install a specific version by setting GCLI_VERSION for the piped-to
# shell (not curl, which wouldn't pass it along):
#
#   curl -fsSL https://raw.githubusercontent.com/daflecardoso/gcli/main/install.sh | GCLI_VERSION=v0.1.0 sh
set -eu

REPO="daflecardoso/gcli"
INSTALL_DIR="${GCLI_INSTALL_DIR:-$HOME/.local/bin}"
VERSION="${GCLI_VERSION:-latest}"

info() { printf '\033[1;36m==>\033[0m %s\n' "$1"; }
die() { printf '\033[1;31merror:\033[0m %s\n' "$1" >&2; exit 1; }

os() {
  case "$(uname -s)" in
    Linux) echo "linux" ;;
    Darwin) echo "darwin" ;;
    *) die "unsupported OS: $(uname -s). Use install.ps1 on Windows." ;;
  esac
}

arch() {
  case "$(uname -m)" in
    x86_64 | amd64) echo "amd64" ;;
    arm64 | aarch64) echo "arm64" ;;
    *) die "unsupported architecture: $(uname -m)" ;;
  esac
}

main() {
  need_cmd curl
  need_cmd tar

  GOOS="$(os)"
  GOARCH="$(arch)"
  ASSET="gcli_${GOOS}_${GOARCH}.tar.gz"
  if [ "$VERSION" = "latest" ]; then
    URL="https://github.com/${REPO}/releases/latest/download/${ASSET}"
  else
    URL="https://github.com/${REPO}/releases/download/${VERSION}/${ASSET}"
  fi

  TMP_DIR="$(mktemp -d)"
  trap 'rm -rf "$TMP_DIR"' EXIT

  info "Downloading ${ASSET} (${VERSION})..."
  curl -fsSL "$URL" -o "$TMP_DIR/gcli.tar.gz" \
    || die "download failed: $URL (no release built for ${GOOS}/${GOARCH}, or ${VERSION} doesn't exist?)"

  tar -xzf "$TMP_DIR/gcli.tar.gz" -C "$TMP_DIR"

  mkdir -p "$INSTALL_DIR"
  mv "$TMP_DIR/gcli" "$INSTALL_DIR/gcli"
  chmod +x "$INSTALL_DIR/gcli"

  info "Installed to ${INSTALL_DIR}/gcli"

  case ":$PATH:" in
    *":$INSTALL_DIR:"*) ;;
    *)
      printf '\n\033[1;33mNote:\033[0m %s is not on your PATH.\n' "$INSTALL_DIR"
      printf 'Add this to your shell profile:\n\n  export PATH="%s:$PATH"\n\n' "$INSTALL_DIR"
      ;;
  esac

  info "Run 'gcli' inside a git repo with a gcli.json config to get started."
}

need_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    die "missing required command: $1"
  fi
}

main "$@"
