#!/bin/bash
# Create gen.go files for all framework packages

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
FRAMEWORKS_DIR="$PROJECT_ROOT/generated/frameworks"

# Get list of all framework directories
for framework_dir in "$FRAMEWORKS_DIR"/*; do
    if [[ ! -d "$framework_dir" ]]; then
        continue
    fi

    framework_name=$(basename "$framework_dir")

    # Convert to proper case for framework name (e.g., fskit -> FSKit)
    case "$framework_name" in
        appkit) FRAMEWORK="AppKit" ;;
        avfoundation) FRAMEWORK="AVFoundation" ;;
        coreaudio) FRAMEWORK="CoreAudio" ;;
        corefoundation) FRAMEWORK="CoreFoundation" ;;
        coregraphics) FRAMEWORK="CoreGraphics" ;;
        corelocation) FRAMEWORK="CoreLocation" ;;
        coremedia) FRAMEWORK="CoreMedia" ;;
        coreml) FRAMEWORK="CoreML" ;;
        corevideo) FRAMEWORK="CoreVideo" ;;
        createml) FRAMEWORK="CreateML" ;;
        diskarbitration) FRAMEWORK="DiskArbitration" ;;
        foundation) FRAMEWORK="Foundation" ;;
        fskit) FRAMEWORK="FSKit" ;;
        imageio) FRAMEWORK="ImageIO" ;;
        iokit) FRAMEWORK="IOKit" ;;
        metal) FRAMEWORK="Metal" ;;
        metalkit) FRAMEWORK="MetalKit" ;;
        naturallanguage) FRAMEWORK="NaturalLanguage" ;;
        network) FRAMEWORK="Network" ;;
        networkextension) FRAMEWORK="NetworkExtension" ;;
        oslog) FRAMEWORK="OSLog" ;;
        quartzcore) FRAMEWORK="QuartzCore" ;;
        scenekit) FRAMEWORK="SceneKit" ;;
        security) FRAMEWORK="Security" ;;
        spritekit) FRAMEWORK="SpriteKit" ;;
        systemconfiguration) FRAMEWORK="SystemConfiguration" ;;
        vision) FRAMEWORK="Vision" ;;
        xpc) FRAMEWORK="XPC" ;;
        combine) FRAMEWORK="Combine" ;;
        *)
            echo "Warning: Unknown framework '$framework_name', skipping"
            continue
            ;;
    esac

    echo "Creating gen.go for $framework_name ($FRAMEWORK)"

    # Create gen.go
    cat > "$framework_dir/gen.go" <<'EOF'
//go:build ignore

package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Remove generated files
	generatedFiles := []string{
		"doc.go",
		"types.gen.go",
		"functions.gen.go",
		"loader.gen.go",
	}

	for _, file := range generatedFiles {
		if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
			log.Printf("Warning: failed to remove %s: %v", file, err)
		}
	}

	// Re-generate bindings using generate-framework-bindings directly
	// This skips the mirroring step and uses cached JSON files
	cmd := exec.Command("generate-framework-bindings", "-framework", "FRAMEWORK_NAME", "-output", "generated/frameworks")

	// Set working directory to project root
	// We're in generated/frameworks/PACKAGE_NAME, so go up 3 levels
	if wd, err := os.Getwd(); err == nil {
		projectRoot := filepath.Join(wd, "..", "..", "..")
		cmd.Dir = projectRoot
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to regenerate FRAMEWORK_NAME bindings: %v", err)
	}

	log.Println("Successfully regenerated FRAMEWORK_NAME bindings")
}
EOF

    # Replace placeholders
    sed -i '' "s/FRAMEWORK_NAME/$FRAMEWORK/g" "$framework_dir/gen.go"
    sed -i '' "s/PACKAGE_NAME/$framework_name/g" "$framework_dir/gen.go"

    # Add go:generate directive to doc.go if it exists and doesn't have it
    if [[ -f "$framework_dir/doc.go" ]]; then
        if ! grep -q "^//go:generate" "$framework_dir/doc.go"; then
            # Add go:generate directive at the top
            echo "//go:generate go run gen.go" > "$framework_dir/doc.go.tmp"
            echo "" >> "$framework_dir/doc.go.tmp"
            cat "$framework_dir/doc.go" >> "$framework_dir/doc.go.tmp"
            mv "$framework_dir/doc.go.tmp" "$framework_dir/doc.go"
            echo "  ✓ Added go:generate directive to doc.go"
        fi
    fi
done

echo ""
echo "Created gen.go files for all frameworks"
echo "You can regenerate any framework with: cd generated/frameworks/FRAMEWORK && go generate"
