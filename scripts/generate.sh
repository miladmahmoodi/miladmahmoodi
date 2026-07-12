#!/usr/bin/env bash
# Generate README.md using Forge.
# Used by the GitHub Action and locally via `make generate`.

set -euo pipefail

FORGE="${FORGE_BIN:-forge}"

if ! command -v "$FORGE" &>/dev/null; then
  echo "forge not found — installing..."
  go install github.com/miladmahmoodi/forge@latest
fi

echo "Running forge build..."
"$FORGE" build --config config.yml

echo "Done — README.md regenerated."
