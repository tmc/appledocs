#!/bin/bash
# Generate Go bindings for all macOS frameworks

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
cd "$PROJECT_ROOT"

# List of frameworks to skip (private/internal frameworks starting with _)
# Also skip frameworks known to cause issues
SKIP_FRAMEWORKS=(
    "_*"  # Skip all private frameworks
    "Python.framework"
    "Kernel.framework"
    "System.framework"
)

# Get list of all public frameworks
FRAMEWORKS_DIR="/System/Library/Frameworks"
FRAMEWORKS=()

while IFS= read -r framework; do
    # Remove .framework suffix
    name="${framework%.framework}"

    # Skip if starts with underscore
    if [[ "$name" == _* ]]; then
        continue
    fi

    FRAMEWORKS+=("$name")
done < <(ls "$FRAMEWORKS_DIR" | grep -v '^_')

echo "Found ${#FRAMEWORKS[@]} public frameworks"
echo ""

# Generate bindings for each framework
SUCCESS=0
FAILED=0
SKIPPED=0

for framework in "${FRAMEWORKS[@]}"; do
    echo "========================================="
    echo "Processing: $framework"
    echo "========================================="

    # Check if already generated
    OUTPUT_DIR="$PROJECT_ROOT/generated/frameworks/$(echo "$framework" | tr '[:upper:]' '[:lower:]')"
    if [[ -f "$OUTPUT_DIR/doc.go" ]]; then
        echo "✓ Already generated, skipping"
        ((SKIPPED++))
        continue
    fi

    # Try to generate bindings
    if timeout 300 ./appledocs generate-framework "$framework" 2>&1 | tee "/tmp/generate-$framework.log"; then
        echo "✓ Successfully generated $framework"
        ((SUCCESS++))
    else
        EXIT_CODE=$?
        if [[ $EXIT_CODE -eq 124 ]]; then
            echo "✗ Timeout generating $framework (> 5 minutes)"
        else
            echo "✗ Failed to generate $framework (exit code: $EXIT_CODE)"
        fi
        ((FAILED++))
    fi

    echo ""
done

echo "========================================="
echo "Summary:"
echo "  Success: $SUCCESS"
echo "  Failed:  $FAILED"
echo "  Skipped: $SKIPPED"
echo "  Total:   ${#FRAMEWORKS[@]}"
echo "========================================="
