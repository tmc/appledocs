#!/bin/bash

# Check Go documentation completeness
# This script verifies that all exported types, functions, and methods have documentation

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "Checking Go documentation..."

# Find all Go files (excluding tests and vendor)
GO_FILES=$(find . -name "*.go" -not -path "./vendor/*" -not -path "./.cache/*" -not -name "*_test.go")

MISSING_DOCS=0
TOTAL_EXPORTED=0

for file in $GO_FILES; do
    # Check for exported types, functions, and methods without documentation
    while IFS= read -r line; do
        if [[ -n "$line" ]]; then
            ((TOTAL_EXPORTED++))
            
            # Extract the identifier name
            identifier=$(echo "$line" | awk '{print $2}' | sed 's/[({].*//')
            
            # Check if the previous line is a comment
            line_num=$(grep -n "^$line" "$file" | cut -d: -f1)
            prev_line_num=$((line_num - 1))
            
            if [[ $prev_line_num -gt 0 ]]; then
                prev_line=$(sed -n "${prev_line_num}p" "$file")
                if [[ ! "$prev_line" =~ ^[[:space:]]*// ]]; then
                    echo -e "${RED}Missing documentation:${NC} $file:$line_num - $identifier"
                    ((MISSING_DOCS++))
                fi
            else
                echo -e "${RED}Missing documentation:${NC} $file:$line_num - $identifier"
                ((MISSING_DOCS++))
            fi
        fi
    done < <(grep -E "^(type|func|var|const) [A-Z]" "$file" || true)
done

echo ""
echo "Documentation Summary:"
echo "Total exported items: $TOTAL_EXPORTED"
echo "Missing documentation: $MISSING_DOCS"

if [[ $MISSING_DOCS -gt 0 ]]; then
    coverage=$((100 * (TOTAL_EXPORTED - MISSING_DOCS) / TOTAL_EXPORTED))
    echo -e "${YELLOW}Documentation coverage: ${coverage}%${NC}"
    
    # Fail if coverage is below threshold
    if [[ $coverage -lt 80 ]]; then
        echo -e "${RED}Documentation coverage is below 80% threshold${NC}"
        exit 1
    else
        echo -e "${YELLOW}Warning: Some documentation is missing but coverage is acceptable${NC}"
    fi
else
    echo -e "${GREEN}All exported items are documented!${NC}"
fi

# Check for TODO/FIXME comments
echo ""
echo "Checking for TODO/FIXME comments..."
TODO_COUNT=$(grep -r "TODO\|FIXME" --include="*.go" --exclude-dir=vendor --exclude-dir=.cache . | wc -l || echo "0")

if [[ $TODO_COUNT -gt 0 ]]; then
    echo -e "${YELLOW}Found $TODO_COUNT TODO/FIXME comments${NC}"
    grep -r "TODO\|FIXME" --include="*.go" --exclude-dir=vendor --exclude-dir=.cache . || true
fi

echo ""
echo "Documentation check complete!"