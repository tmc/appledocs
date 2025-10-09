#!/bin/bash
# Test all example programs with and without -e2e flag

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
cd "$PROJECT_ROOT"

echo "Testing all framework examples..."
echo ""

PASSED=0
FAILED=0
FRAMEWORKS=$(ls -1 examples/)

for fw in $FRAMEWORKS; do
    echo "=========================================
"
    echo "Testing: $fw"
    echo "========================================="

    cd "examples/$fw"

    # Test without -e2e
    echo -n "  Basic run... "
    if timeout 10 go run main.go > /tmp/test-$fw.log 2>&1; then
        echo "✓ PASS"
        ((PASSED++))
    else
        echo "✗ FAIL"
        echo "    See /tmp/test-$fw.log for details"
        ((FAILED++))
    fi

    # Test with -e2e
    echo -n "  E2E tests... "
    if timeout 10 go run main.go -e2e > /tmp/test-$fw-e2e.log 2>&1; then
        echo "✓ PASS"
        ((PASSED++))
    else
        echo "✗ FAIL"
        echo "    See /tmp/test-$fw-e2e.log for details"
        ((FAILED++))
    fi

    cd "$PROJECT_ROOT"
    echo ""
done

echo "========================================="
echo "Summary:"
echo "  Passed: $PASSED"
echo "  Failed: $FAILED"
echo "  Total:  $((PASSED + FAILED))"
echo "========================================="

if [ $FAILED -gt 0 ]; then
    exit 1
fi
