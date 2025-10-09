#!/bin/bash
# Tidy all example go.mod files

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
cd "$PROJECT_ROOT"

for dir in examples/*/; do
  example=$(basename "$dir")
  case "$example" in
    codegen-demo|dispatch-gcd|drawing-generated-bindings|extract-methods|\
    helloworld-*|list-*|platform-analysis|search-symbols|security-keychain|\
    fskit-simple|coregraphics-drawing)
      continue
      ;;
  esac
  echo "Tidying $example..."
  cd "$dir"
  go mod tidy
  cd "$PROJECT_ROOT"
done

echo "Done!"
