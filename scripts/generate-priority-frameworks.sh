#!/bin/bash
# Generate Go bindings for high-priority macOS frameworks
# These are the most commonly used frameworks in macOS development

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
cd "$PROJECT_ROOT"

# High-priority frameworks for macOS development
PRIORITY_FRAMEWORKS=(
    # Core frameworks
    "Foundation"
    "CoreFoundation"
    "CoreGraphics"
    "AppKit"

    # UI and graphics
    "QuartzCore"
    "Metal"
    "MetalKit"
    "SpriteKit"
    "SceneKit"

    # Media
    "AVFoundation"
    "CoreAudio"
    "CoreMedia"
    "CoreVideo"
    "ImageIO"

    # System
    "Security"
    "SystemConfiguration"
    "IOKit"
    "DiskArbitration"

    # Modern APIs
    "Combine"
    "CoreML"
    "CreateML"
    "NaturalLanguage"
    "Vision"
    "CoreLocation"

    # File system
    "FSKit"

    # Networking
    "Network"
    "NetworkExtension"

    # Developer tools
    "XPC"
    "OSLog"
)

echo "Generating bindings for ${#PRIORITY_FRAMEWORKS[@]} priority frameworks"
echo ""

SUCCESS=0
FAILED=0
SKIPPED=0

for framework in "${PRIORITY_FRAMEWORKS[@]}"; do
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

    # Try to generate bindings with timeout
    # Use generate-framework-bindings directly to skip re-mirroring
    if timeout 300 generate-framework-bindings -framework "$framework" -output "generated/frameworks" 2>&1 | tee "/tmp/generate-$framework.log"; then
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
echo "  Total:   ${#PRIORITY_FRAMEWORKS[@]}"
echo "========================================="
