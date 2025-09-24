#!/usr/bin/env bash
set -euo pipefail

# Build binary into ./bin
mkdir -p "$(dirname "$0")/../bin"
GOFLAGS=${GOFLAGS:-}
GOOS=${GOOS:-}
GOARCH=${GOARCH:-}

echo "Building shop-api..."
GOFLAGS="$GOFLAGS" GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$(dirname "$0")/../bin/shop-api" ./cmd/shop-api

echo "OK -> bin/shop-api"
