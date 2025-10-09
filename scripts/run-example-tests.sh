#!/bin/bash
# Run all framework examples in test mode
# This tests that the generated bindings can be imported and frameworks can be loaded

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
cd "$PROJECT_ROOT"

echo "Running framework binding tests..."
echo "=================================="
echo ""

PASSED=0
FAILED=0
TOTAL=0

for example_dir in examples/*/; do
    example=$(basename "$example_dir")

    # Skip non-framework examples
    case "$example" in
        codegen-demo|dispatch-gcd|drawing-generated-bindings|extract-methods|\
        helloworld-*|list-*|platform-analysis|search-symbols|security-keychain|\
        fskit-simple|coregraphics-drawing)
            continue
            ;;
    esac

    ((TOTAL++))
    echo -n "Testing $example... "

    cd "$example_dir"

    # Test that the example compiles and runs
    if timeout 5 go run main.go > /tmp/test-$example.out 2>&1; then
        echo "✓ PASS"
        ((PASSED++))
    else
        echo "✗ FAIL"
        echo "  Output: $(head -3 /tmp/test-$example.out 2>/dev/null | tr '\n' ' ')"
        ((FAILED++))
    fi

    cd "$PROJECT_ROOT"
done

echo ""
echo "=================================="
echo "Test Summary:"
echo "  Passed: $PASSED/$TOTAL"
echo "  Failed: $FAILED/$TOTAL"
echo "=================================="

if [ $FAILED -gt 0 ]; then
    echo ""
    echo "Failed tests:"
    for example_dir in examples/*/; do
        example=$(basename "$example_dir")
        if [ -f "/tmp/test-$example.out" ] && ! grep -q "✓" "/tmp/test-$example.out" 2>/dev/null; then
            echo "  $example: $(head -1 /tmp/test-$example.out 2>/dev/null)"
        fi
    done
    exit 1
fi

echo ""
echo "All tests passed!"
