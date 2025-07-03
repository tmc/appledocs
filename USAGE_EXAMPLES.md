# Apple Docs Mirror Tool - Usage Examples

## Overview

This document provides comprehensive, real-world usage examples for the `appledocs` tool. Each example includes the command, expected output, and practical use cases.

## Table of Contents

- [Basic Operations](#basic-operations)
- [Advanced Crawling Scenarios](#advanced-crawling-scenarios)
- [Validation and Quality Assurance](#validation-and-quality-assurance)
- [Performance Optimization](#performance-optimization)
- [CI/CD Integration](#cicd-integration)
- [Monitoring and Metrics](#monitoring-and-metrics)
- [Content Generation](#content-generation)
- [Troubleshooting Workflows](#troubleshooting-workflows)

## Basic Operations

### 1. First-Time Setup and Basic Crawl

```bash
# Create a basic mirror of Apple documentation
./appledocs -mode crawl -verbose

# Expected output:
# 2024/01/15 10:00:00 Starting crawl from https://developer.apple.com/tutorials/data/documentation/technologies.json
# 2024/01/15 10:00:01 Found 156 technologies to process
# 2024/01/15 10:00:02 Progress: 1 files processed | Cache: 0 hits (0.0%), 1 misses
# ...
# 2024/01/15 10:05:30 Final statistics:
#   - Runtime: 5m30s
#   - Processed: 2,847 JSON files (8.6 files/sec)
#   - Cache: 0 hits (0.0%), 2,847 misses
#   - Data: 45.2 MB downloaded
```

**Use Case**: Initial setup for development teams wanting local Apple documentation access.

### 2. Framework-Specific Documentation

```bash
# Crawl only SwiftUI documentation
./appledocs -mode crawl -entry-point SwiftUI -verbose

# Expected output:
# 2024/01/15 10:00:00 Fetching URLs for framework: SwiftUI
# 2024/01/15 10:00:01 Found 1,243 unique data URLs
# 2024/01/15 10:00:02 Progress: 1 files processed | Cache: 0 hits (0.0%), 1 misses
# ...
# 2024/01/15 10:02:45 Final statistics:
#   - Runtime: 2m45s
#   - Processed: 1,243 JSON files (7.5 files/sec)
#   - Content: 1 frameworks, 89 classes, 1,153 methods
```

**Use Case**: Teams focusing on specific frameworks who want targeted documentation.

### 3. Quick URL Discovery

```bash
# Discover all available URLs without downloading
./appledocs -print-urls | head -20

# Expected output:
# https://developer.apple.com/tutorials/data/documentation/Accelerate.json
# https://developer.apple.com/tutorials/data/documentation/AVFoundation.json
# https://developer.apple.com/tutorials/data/documentation/CoreData.json
# https://developer.apple.com/tutorials/data/documentation/SwiftUI.json
# ...
```

**Use Case**: Explore available documentation before committing to a full crawl.

## Advanced Crawling Scenarios

### 4. High-Performance Bulk Crawling

```bash
# High-concurrency crawl with rate limiting
./appledocs -mode crawl \
  -concurrency 30 \
  -rate-limit 75 \
  -timeout 45s \
  -max-time 3h \
  -export-metrics bulk-crawl-metrics.json \
  -verbose

# Expected output:
# 2024/01/15 10:00:00 Rate limiting enabled: 75 requests/second
# 2024/01/15 10:00:00 Starting crawl with 30 workers
# 2024/01/15 10:00:05 Progress: 187 files processed | Cache: 0 hits (0.0%), 187 misses | 12.3 MB/s
# 2024/01/15 10:00:10 Progress: 423 files processed | Cache: 0 hits (0.0%), 423 misses | 15.7 MB/s
# ...
# 2024/01/15 10:45:20 Final statistics:
#   - Runtime: 45m20s
#   - Processed: 15,842 JSON files (5.8 files/sec)
#   - Data: 892.3 MB downloaded, 0.0 MB from cache (19.7 MB/s)
#   - Content: 87 frameworks, 2,341 classes, 13,414 methods
```

**Use Case**: Enterprise environments with high-bandwidth networks needing complete documentation mirrors.

### 5. Incremental Updates with Cache Validation

```bash
# Update existing cache with validation
./appledocs -mode crawl \
  -validate-cache \
  -checksum-validation \
  -concurrency 10 \
  -rate-limit 20 \
  -export-metrics incremental-update.json

# Expected output:
# 2024/01/15 10:00:00 Validating cache integrity...
# 2024/01/15 10:00:05 Cache validation completed successfully
# 2024/01/15 10:00:05 checked 8,234 files: 0 corrupted, 1,203 new
# 2024/01/15 10:00:06 Starting crawl from https://developer.apple.com/tutorials/data/documentation/technologies.json
# 2024/01/15 10:00:10 Progress: 45 files processed | Cache: 8,189 hits (99.5%), 45 misses
# ...
# 2024/01/15 10:05:30 Final statistics:
#   - Runtime: 5m30s
#   - Processed: 8,234 JSON files (25.0 files/sec)
#   - Cache: 7,031 hits (85.4%), 1,203 misses
#   - Data: 34.5 MB downloaded, 256.8 MB from cache (95.2 MB/s)
```

**Use Case**: Daily or weekly updates to keep documentation current while minimizing bandwidth usage.

### 6. Filtered Crawling for Specific Use Cases

```bash
# Crawl without symbol-level docs to reduce size
./appledocs -mode crawl \
  -skip-symbols \
  -exclude-paths "legacy,deprecated,en-US/docs/Mozilla" \
  -concurrency 15 \
  -verbose

# Expected output:
# 2024/01/15 10:00:00 Starting crawl with symbol filtering enabled
# 2024/01/15 10:00:05 Progress: 89 files processed | Skipped symbol URLs: 1,456
# 2024/01/15 10:00:10 Progress: 234 files processed | Skipped symbol URLs: 3,892
# ...
# 2024/01/15 10:15:30 Final statistics:
#   - Runtime: 15m30s
#   - Processed: 3,421 JSON files (3.7 files/sec)
#   - Skipped symbol URLs: 12,398
#   - Data: 156.7 MB downloaded (50% size reduction)
```

**Use Case**: Documentation for overview and high-level API understanding without detailed method documentation.

## Validation and Quality Assurance

### 7. Comprehensive Cache Validation

```bash
# Thorough validation of existing cache
./appledocs -validate-cache -checksum-validation -verbose

# Expected output:
# 2024/01/15 10:00:00 Validating cache integrity...
# 2024/01/15 10:00:00 Using enhanced checksum-based validation...
# 2024/01/15 10:00:15 checked 8,234 files: 2 corrupted, 0 new
# 2024/01/15 10:00:15 Cache validation errors found:
#   - validation error in .cache/developer.apple.com/tutorials/data/documentation/SwiftUI.json at integrity: file checksum mismatch - file may be corrupted or tampered
#   - validation error in .cache/developer.apple.com/tutorials/data/documentation/UIKit.json at integrity: file checksum mismatch - file may be corrupted or tampered
```

**Use Case**: Quality assurance before important releases or after network issues.

### 8. Data Integrity Verification

```bash
# Force refresh with enhanced validation
./appledocs -mode crawl \
  -force \
  -checksum-validation \
  -export-metrics validation-metrics.json \
  -log-level debug

# Expected output:
# 2024/01/15 10:00:00 DEBUG: Data validation enabled for all downloads
# 2024/01/15 10:00:05 Data validation failed for "https://developer.apple.com/tutorials/data/documentation/broken.json":
#   - File contains binary data
# 2024/01/15 10:00:05 Added bad URL due to validation failure: https://developer.apple.com/tutorials/data/documentation/broken.json
# ...
# 2024/01/15 10:30:00 Final statistics:
#   - Errors: 23 (validation failures caught and handled)
```

**Use Case**: Ensuring data quality in mission-critical documentation systems.

## Performance Optimization

### 9. Bandwidth-Limited Environment

```bash
# Optimize for slow or metered connections
./appledocs -mode crawl \
  -concurrency 3 \
  -rate-limit 5 \
  -delay 200ms \
  -timeout 90s \
  -skip-symbols \
  -exclude-paths "media,assets,images"

# Expected output:
# 2024/01/15 10:00:00 Rate limiting enabled: 5 requests/second
# 2024/01/15 10:00:00 Using 200ms delay between URL processing
# 2024/01/15 10:00:10 Progress: 12 files processed | Cache: 0 hits (0.0%), 12 misses | 0.8 MB/s
# ...
# 2024/01/15 11:30:00 Final statistics:
#   - Runtime: 1h30m
#   - Processed: 2,341 JSON files (0.4 files/sec)
#   - Data: 67.8 MB downloaded (0.8 MB/s average)
```

**Use Case**: Mobile hotspots, satellite connections, or metered bandwidth environments.

### 10. Memory-Optimized Processing

```bash
# Process large documentation sets with memory constraints
./appledocs -mode crawl \
  -concurrency 5 \
  -rate-limit 10 \
  -timeout 30s \
  -pretty=false \
  -log-level warn

# Expected output:
# 2024/01/15 10:00:00 Starting memory-optimized crawl
# 2024/01/15 10:00:10 Progress: 67 files processed | Cache: 0 hits (0.0%), 67 misses
# ...
# Memory usage stays consistent throughout run
```

**Use Case**: Constrained environments like small VPS instances or embedded systems.

## CI/CD Integration

### 11. Automated Daily Documentation Updates

```bash
#!/bin/bash
# daily-docs-update.sh

# Update documentation with validation and metrics
./appledocs -mode crawl \
  -validate-cache \
  -checksum-validation \
  -concurrency 20 \
  -rate-limit 40 \
  -max-time 2h \
  -export-metrics "metrics-$(date +%Y-%m-%d).json" \
  -verbose

# Generate user-friendly outputs
./appledocs -mode html
./appledocs -mode markdown

# Expected output:
# 2024/01/15 03:00:00 Daily documentation update started
# 2024/01/15 03:00:05 Cache validation completed successfully
# 2024/01/15 03:15:30 Crawling completed: 1,203 new files, 7,031 cached
# 2024/01/15 03:16:00 HTML index generated: output/index.html
# 2024/01/15 03:17:30 Markdown documentation generated: markdown/
# 2024/01/15 03:17:30 Metrics exported: metrics-2024-01-15.json
```

**Use Case**: Automated systems maintaining up-to-date documentation repositories.

### 12. Build Pipeline Integration

```bash
# build-docs.sh for CI pipeline
set -e

echo "Starting documentation build..."

# Quick validation
./appledocs -validate-cache || {
    echo "Cache validation failed, performing full refresh"
    ./appledocs -mode crawl -force -concurrency 10
}

# Generate documentation artifacts
./appledocs -mode all -md-output ./dist/docs

# Validate output
[ -f "./dist/docs/index.md" ] || {
    echo "Documentation generation failed"
    exit 1
}

echo "Documentation build completed successfully"

# Expected output:
# Starting documentation build...
# Cache validation completed successfully
# HTML index generation completed
# Markdown generation completed
# Documentation build completed successfully
```

**Use Case**: Integration into software release pipelines.

## Monitoring and Metrics

### 13. Comprehensive Metrics Collection

```bash
# Detailed metrics collection for monitoring
./appledocs -mode crawl \
  -concurrency 25 \
  -rate-limit 50 \
  -export-metrics detailed-metrics.json \
  -verbose

# Analyze metrics after completion
jq '.cache_hit_rate' detailed-metrics.json
# Output: 87.3

jq '.download_rate_mbps' detailed-metrics.json  
# Output: 15.7

jq '.http_errors' detailed-metrics.json
# Output: {"404": 12, "503": 3}

jq '.framework_count, .class_count, .method_count' detailed-metrics.json
# Output: 87
# Output: 2341  
# Output: 13414
```

**Use Case**: Performance monitoring and capacity planning for documentation systems.

### 14. Health Check and Monitoring

```bash
# Simple health check script
#!/bin/bash
# health-check.sh

METRICS_FILE="/tmp/appledocs-health.json"

# Run quick validation
./appledocs -validate-cache -export-metrics "$METRICS_FILE" 2>/dev/null

if [ $? -eq 0 ]; then
    ERRORS=$(jq '.errors' "$METRICS_FILE")
    if [ "$ERRORS" -eq 0 ]; then
        echo "HEALTHY: Documentation cache is valid"
        exit 0
    else
        echo "WARNING: $ERRORS validation errors found"
        exit 1
    fi
else
    echo "CRITICAL: Documentation validation failed"
    exit 2
fi

# Expected outputs:
# HEALTHY: Documentation cache is valid (exit 0)
# WARNING: 3 validation errors found (exit 1) 
# CRITICAL: Documentation validation failed (exit 2)
```

**Use Case**: Monitoring systems and health checks.

## Content Generation

### 15. Multi-Format Documentation Generation

```bash
# Generate all documentation formats
./appledocs -mode all \
  -output ./apple-docs \
  -md-output ./documentation \
  -verbose

# Generate specialized EndpointSecurity docs
make endpointsecurity

# Expected output:
# 2024/01/15 10:00:00 Starting HTML index generation
# 2024/01/15 10:00:15 HTML index generation completed: ./apple-docs/index.html
# 2024/01/15 10:00:15 Starting Markdown generation
# 2024/01/15 10:01:45 Markdown generation completed: ./documentation/
# 2024/01/15 10:01:45 Generated specialized EndpointSecurity reference

# Verify outputs
ls -la ./apple-docs/index.html
ls -la ./documentation/index.md
ls -la ./EndpointSecurity.md
```

**Use Case**: Creating comprehensive documentation websites and offline references.

### 16. Custom Documentation Subsets

```bash
# Generate documentation for specific frameworks
./appledocs -print-urls -entry-point SwiftUI | \
while read url; do
    echo "Processing: $url"
    # Custom processing logic here
done

# Expected output:
# Processing: https://developer.apple.com/tutorials/data/documentation/SwiftUI.json
# Processing: https://developer.apple.com/tutorials/data/documentation/SwiftUI/View.json
# Processing: https://developer.apple.com/tutorials/data/documentation/SwiftUI/Button.json
# ...
```

**Use Case**: Creating focused documentation for specific teams or projects.

## Troubleshooting Workflows

### 17. Diagnosing Network Issues

```bash
# Debug network connectivity issues
./appledocs -mode crawl \
  -concurrency 1 \
  -rate-limit 1 \
  -timeout 60s \
  -log-level debug \
  -verbose \
  -entry-point SwiftUI

# Expected debug output:
# 2024/01/15 10:00:00 DEBUG: Creating request for https://developer.apple.com/tutorials/data/documentation/SwiftUI.json
# 2024/01/15 10:00:00 DEBUG: Adding browser-like headers
# 2024/01/15 10:00:01 DEBUG: Request completed in 1.2s
# 2024/01/15 10:00:01 DEBUG: Response size: 45,678 bytes
# 2024/01/15 10:00:01 Cache miss for "https://developer.apple.com/tutorials/data/documentation/SwiftUI.json"
```

**Use Case**: Troubleshooting connectivity issues or server problems.

### 18. Recovery from Partial Failures

```bash
# Recover from interrupted crawl
./appledocs -mode crawl \
  -validate-cache \
  -force=false \
  -concurrency 10 \
  -export-metrics recovery-metrics.json

# Check what was recovered
jq '.cache_hits, .cache_misses' recovery-metrics.json
# Output: 5432  (already cached)
# Output: 1234  (newly downloaded)

# Clean up any corruption
./appledocs -validate-cache -checksum-validation
```

**Use Case**: Resuming after network interruptions or system crashes.

### 19. Performance Debugging

```bash
# Analyze performance bottlenecks
./appledocs -mode crawl \
  -concurrency 1 \
  -rate-limit 0 \
  -export-metrics perf-debug.json \
  -verbose \
  -entry-point SwiftUI

# Analyze results
echo "Average response time:" $(jq '.avg_response_time' perf-debug.json)
echo "Download rate:" $(jq '.download_rate_mbps' perf-debug.json)
echo "Cache hit rate:" $(jq '.cache_hit_rate' perf-debug.json)

# Expected output:
# Average response time: "892ms"
# Download rate: 3.4
# Cache hit rate: 23.5
```

**Use Case**: Optimizing performance for specific network conditions.

### 20. Data Quality Analysis

```bash
# Comprehensive data quality check
./appledocs -validate-cache \
  -checksum-validation \
  -log-level debug \
  -verbose > validation-report.txt 2>&1

# Analyze validation report
grep "validation error" validation-report.txt | wc -l
grep "corrupted" validation-report.txt
grep "checksum mismatch" validation-report.txt

# Generate summary
echo "Validation Summary:"
echo "Total errors: $(grep -c "validation error" validation-report.txt)"
echo "Corrupted files: $(grep -c "corrupted" validation-report.txt)"
echo "Checksum mismatches: $(grep -c "checksum mismatch" validation-report.txt)"

# Expected output:
# Validation Summary:
# Total errors: 7
# Corrupted files: 2
# Checksum mismatches: 2
```

**Use Case**: Ensuring data integrity before deploying documentation updates.

## Advanced Integration Examples

### 21. GraphQL API Server Usage

```bash
# Start the GraphQL API server
cd cmd/appledocs-gql
go build
./appledocs-gql -port 8080 -cache ../../.cache

# Query the API
curl "http://localhost:8080/api/document?path=tutorials/data/documentation/SwiftUI.json" | jq '.title'
# Output: "SwiftUI"

curl "http://localhost:8080/api/search?q=Button&limit=5" | jq '.results[].title'
# Output: "Button", "ButtonStyle", "ButtonConfiguration", etc.
```

**Use Case**: Providing programmatic access to documentation for custom applications.

### 22. Docker Container Integration

```bash
# Build documentation in container
docker run --rm -v $(pwd):/workspace -w /workspace golang:1.21 bash -c "
  go build -o appledocs
  ./appledocs -mode all -output /workspace/docs
"

# Use in multi-stage build
# Dockerfile
FROM golang:1.21 AS builder
COPY . /src
WORKDIR /src
RUN go build -o appledocs

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /src/appledocs /usr/local/bin/
VOLUME ["/docs", "/cache"]
CMD ["appledocs", "-mode", "all", "-output", "/docs", "-cache", "/cache"]
```

**Use Case**: Containerized documentation generation for cloud deployments.

## Best Practices Summary

### Recommended Commands by Use Case

**Development Team (Daily Use)**:
```bash
./appledocs -mode crawl -concurrency 10 -rate-limit 20 -skip-symbols
```

**Enterprise Mirror (Complete)**:
```bash
./appledocs -mode all -concurrency 30 -rate-limit 75 -validate-cache -export-metrics enterprise-metrics.json
```

**CI/CD Pipeline**:
```bash
./appledocs -mode crawl -validate-cache -checksum-validation -max-time 1h -export-metrics ci-metrics.json
```

**Quality Assurance**:
```bash
./appledocs -validate-cache -checksum-validation -verbose -log-level debug
```

**Bandwidth Constrained**:
```bash
./appledocs -mode crawl -concurrency 3 -rate-limit 5 -skip-symbols -exclude-paths "media,assets"
```

These examples demonstrate the flexibility and power of the appledocs tool across various real-world scenarios, from basic documentation mirroring to enterprise-scale deployments with comprehensive monitoring and validation.