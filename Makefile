.PHONY: list-frameworks list-generated setup-frameworks setup-priority generate-priority generate-all clean-frameworks test help

# Default target
help:
	@echo "Available targets:"
	@echo "  list-frameworks     - List all available macOS frameworks"
	@echo "  list-generated      - List currently generated framework packages"
	@echo "  setup-frameworks    - Create directories and gen.go files for all frameworks"
	@echo "  setup-priority      - Create directories and gen.go for priority frameworks only"
	@echo "  generate-priority   - Generate bindings for priority frameworks"
	@echo "  generate-all        - Generate bindings for all frameworks"
	@echo "  clean-frameworks    - Remove all generated framework directories"
	@echo "  test                - Run tests"

# List all available macOS frameworks
list-frameworks:
	@echo "Public macOS frameworks (excluding private frameworks starting with _):"
	@ls /System/Library/Frameworks/ | grep -v '^_' | sed 's/\.framework//' | sort

# List currently generated framework packages
list-generated:
	@echo "Generated framework packages:"
	@ls -1 generated/frameworks/ 2>/dev/null || echo "No frameworks generated yet"
	@echo ""
	@echo "Frameworks with complete bindings (have doc.go):"
	@find generated/frameworks -name "doc.go" -exec dirname {} \; | xargs -n1 basename | sort

# Priority frameworks list
PRIORITY_FRAMEWORKS := Foundation CoreFoundation CoreGraphics AppKit QuartzCore Metal MetalKit SpriteKit SceneKit AVFoundation CoreAudio CoreMedia CoreVideo ImageIO Security SystemConfiguration IOKit DiskArbitration Combine CoreML CreateML NaturalLanguage Vision CoreLocation FSKit Network NetworkExtension XPC OSLog

# Create directories and gen.go files for priority frameworks
setup-priority:
	@echo "Setting up priority frameworks..."
	@for framework in $(PRIORITY_FRAMEWORKS); do \
		dir="generated/frameworks/$$(echo $$framework | tr '[:upper:]' '[:lower:]')"; \
		if [ ! -d "$$dir" ]; then \
			echo "Creating directory: $$dir"; \
			mkdir -p "$$dir"; \
		fi; \
	done
	@./scripts/create-gen-files.sh
	@echo "Priority frameworks setup complete"

# Create directories for all frameworks
setup-frameworks:
	@echo "Creating directories for all frameworks found in /System/Library/Frameworks..."
	@for framework_file in /System/Library/Frameworks/*.framework; do \
		framework=$$(basename "$$framework_file" .framework); \
		if [[ "$$framework" == _* ]]; then continue; fi; \
		dir="generated/frameworks/$$(echo $$framework | tr '[:upper:]' '[:lower:]')"; \
		if [ ! -d "$$dir" ]; then \
			echo "  Creating: $$dir"; \
			mkdir -p "$$dir"; \
		fi; \
	done
	@./scripts/create-gen-files.sh
	@echo "Setup complete. Run 'make generate-priority' or 'make generate-all' to generate bindings."

# Generate bindings for priority frameworks
generate-priority:
	@echo "Generating bindings for priority frameworks..."
	@./scripts/generate-priority-frameworks.sh

# Generate bindings for all frameworks
generate-all:
	@echo "Generating bindings for all frameworks..."
	@./scripts/generate-all-frameworks.sh

# Clean all generated framework directories
clean-frameworks:
	@echo "Removing all generated framework directories..."
	@rm -rf generated/frameworks/*
	@echo "Clean complete"

# Run tests
test:
	@echo "Running binding generator tests..."
	@cd cmd/generate-framework-bindings && go test -v
	@echo ""
	@echo "Testing generated bindings compilation..."
	@cd generated/frameworks/coregraphics && go build
	@echo "All tests passed"
