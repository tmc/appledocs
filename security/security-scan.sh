#!/bin/bash
# Security scanning script for appledocs container images

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
IMAGE_NAME="${1:-appledocs:latest}"
REPORT_DIR="security-reports"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

# Create report directory
mkdir -p "${REPORT_DIR}"

echo -e "${GREEN}Starting security scan for image: ${IMAGE_NAME}${NC}"
echo "Reports will be saved to: ${REPORT_DIR}"
echo "Timestamp: ${TIMESTAMP}"
echo

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to run scan and check results
run_scan() {
    local scanner="$1"
    local command="$2"
    local report_file="$3"
    
    echo -e "${YELLOW}Running ${scanner} scan...${NC}"
    
    if eval "${command}"; then
        echo -e "${GREEN}✓ ${scanner} scan completed successfully${NC}"
        return 0
    else
        echo -e "${RED}✗ ${scanner} scan found issues${NC}"
        return 1
    fi
}

# Track overall scan status
SCAN_FAILED=0

# 1. Trivy Scan
if command_exists trivy; then
    # Vulnerability scan
    run_scan "Trivy vulnerability" \
        "trivy image --severity HIGH,CRITICAL --format json --output ${REPORT_DIR}/trivy-vuln-${TIMESTAMP}.json ${IMAGE_NAME}" \
        "${REPORT_DIR}/trivy-vuln-${TIMESTAMP}.json" || SCAN_FAILED=1
    
    # Configuration scan
    run_scan "Trivy config" \
        "trivy image --security-checks config --format json --output ${REPORT_DIR}/trivy-config-${TIMESTAMP}.json ${IMAGE_NAME}" \
        "${REPORT_DIR}/trivy-config-${TIMESTAMP}.json" || SCAN_FAILED=1
    
    # Generate human-readable report
    trivy image --severity HIGH,CRITICAL --format table ${IMAGE_NAME} > "${REPORT_DIR}/trivy-summary-${TIMESTAMP}.txt"
else
    echo -e "${YELLOW}Trivy not found. Install with: brew install aquasecurity/trivy/trivy${NC}"
fi

echo

# 2. Snyk Scan (if available)
if command_exists snyk; then
    if [ -n "${SNYK_TOKEN:-}" ]; then
        run_scan "Snyk" \
            "snyk container test ${IMAGE_NAME} --json > ${REPORT_DIR}/snyk-${TIMESTAMP}.json" \
            "${REPORT_DIR}/snyk-${TIMESTAMP}.json" || SCAN_FAILED=1
    else
        echo -e "${YELLOW}Snyk found but SNYK_TOKEN not set. Skipping Snyk scan.${NC}"
    fi
else
    echo -e "${YELLOW}Snyk not found. Install with: npm install -g snyk${NC}"
fi

echo

# 3. Grype Scan (Anchore)
if command_exists grype; then
    run_scan "Grype" \
        "grype ${IMAGE_NAME} -o json > ${REPORT_DIR}/grype-${TIMESTAMP}.json" \
        "${REPORT_DIR}/grype-${TIMESTAMP}.json" || SCAN_FAILED=1
else
    echo -e "${YELLOW}Grype not found. Install with: brew install grype${NC}"
fi

echo

# 4. Docker Scout (if available)
if command_exists docker-scout; then
    run_scan "Docker Scout" \
        "docker scout cves ${IMAGE_NAME} --format json > ${REPORT_DIR}/scout-${TIMESTAMP}.json" \
        "${REPORT_DIR}/scout-${TIMESTAMP}.json" || SCAN_FAILED=1
else
    echo -e "${YELLOW}Docker Scout not found. Enable with: docker scout quickview${NC}"
fi

echo

# 5. Hadolint - Dockerfile Linter
if command_exists hadolint; then
    DOCKERFILES=$(find . -name "Dockerfile*" -type f)
    for dockerfile in ${DOCKERFILES}; do
        echo -e "${YELLOW}Running Hadolint on ${dockerfile}...${NC}"
        if hadolint "${dockerfile}" > "${REPORT_DIR}/hadolint-$(basename ${dockerfile})-${TIMESTAMP}.txt" 2>&1; then
            echo -e "${GREEN}✓ Hadolint check passed for ${dockerfile}${NC}"
        else
            echo -e "${RED}✗ Hadolint found issues in ${dockerfile}${NC}"
            SCAN_FAILED=1
        fi
    done
else
    echo -e "${YELLOW}Hadolint not found. Install with: brew install hadolint${NC}"
fi

echo

# 6. Container Structure Test
if command_exists container-structure-test; then
    if [ -f "container-structure-test.yaml" ]; then
        run_scan "Container Structure Test" \
            "container-structure-test test --image ${IMAGE_NAME} --config container-structure-test.yaml --output json > ${REPORT_DIR}/structure-test-${TIMESTAMP}.json" \
            "${REPORT_DIR}/structure-test-${TIMESTAMP}.json" || SCAN_FAILED=1
    else
        echo -e "${YELLOW}Container structure test config not found. Skipping.${NC}"
    fi
else
    echo -e "${YELLOW}Container Structure Test not found. Install from: https://github.com/GoogleContainerTools/container-structure-test${NC}"
fi

echo

# 7. Generate Summary Report
echo -e "${YELLOW}Generating summary report...${NC}"

cat > "${REPORT_DIR}/summary-${TIMESTAMP}.md" << EOF
# Security Scan Summary

**Image**: ${IMAGE_NAME}  
**Date**: $(date)  
**Scan ID**: ${TIMESTAMP}

## Scan Results

| Scanner | Status | Report |
|---------|--------|--------|
EOF

# Add results to summary
for report in "${REPORT_DIR}"/*-"${TIMESTAMP}".*; do
    if [ -f "$report" ]; then
        scanner=$(basename "$report" | cut -d'-' -f1)
        echo "| ${scanner} | ✓ Completed | [View]($(basename "$report")) |" >> "${REPORT_DIR}/summary-${TIMESTAMP}.md"
    fi
done

echo "" >> "${REPORT_DIR}/summary-${TIMESTAMP}.md"

# Add vulnerability counts if Trivy was run
if [ -f "${REPORT_DIR}/trivy-vuln-${TIMESTAMP}.json" ]; then
    echo "## Vulnerability Summary" >> "${REPORT_DIR}/summary-${TIMESTAMP}.md"
    echo "" >> "${REPORT_DIR}/summary-${TIMESTAMP}.md"
    
    # Use jq to parse JSON if available
    if command_exists jq; then
        CRITICAL=$(jq '[.Results[]?.Vulnerabilities[]? | select(.Severity == "CRITICAL")] | length' "${REPORT_DIR}/trivy-vuln-${TIMESTAMP}.json" 2>/dev/null || echo "0")
        HIGH=$(jq '[.Results[]?.Vulnerabilities[]? | select(.Severity == "HIGH")] | length' "${REPORT_DIR}/trivy-vuln-${TIMESTAMP}.json" 2>/dev/null || echo "0")
        
        echo "- **Critical**: ${CRITICAL}" >> "${REPORT_DIR}/summary-${TIMESTAMP}.md"
        echo "- **High**: ${HIGH}" >> "${REPORT_DIR}/summary-${TIMESTAMP}.md"
    fi
fi

echo -e "${GREEN}✓ Summary report saved to: ${REPORT_DIR}/summary-${TIMESTAMP}.md${NC}"

# 8. Create latest symlinks for CI/CD integration
ln -sf "summary-${TIMESTAMP}.md" "${REPORT_DIR}/summary-latest.md"
ln -sf "trivy-vuln-${TIMESTAMP}.json" "${REPORT_DIR}/trivy-vuln-latest.json" 2>/dev/null || true

echo
if [ ${SCAN_FAILED} -eq 0 ]; then
    echo -e "${GREEN}✓ All security scans completed successfully!${NC}"
    exit 0
else
    echo -e "${RED}✗ Some security scans found issues. Please review the reports in ${REPORT_DIR}${NC}"
    exit 1
fi