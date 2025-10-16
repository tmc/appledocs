.PHONY: list-frameworks list-generated list-priority setup-frameworks setup-priority setup-pattern generate-priority generate-all generate generate-pattern clean-frameworks test help

# Color output
RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[0;33m
BLUE := \033[0;34m
NC := \033[0m # No Color

# Directories
FRAMEWORKS_DIR := /System/Library/Frameworks
GENERATED_DIR := generated/frameworks

# Priority frameworks list (most commonly used in macOS development)
PRIORITY_FRAMEWORKS := \
	Foundation CoreFoundation CoreGraphics AppKit \
	QuartzCore Metal MetalKit SpriteKit SceneKit \
	AVFoundation CoreAudio CoreMedia CoreVideo ImageIO \
	Security SystemConfiguration IOKit DiskArbitration \
	Combine CoreML CreateML NaturalLanguage Vision CoreLocation \
	FSKit Network NetworkExtension XPC OSLog

# Dynamically discover all public frameworks (excluding private frameworks starting with _)
ALL_FRAMEWORKS := $(shell ls $(FRAMEWORKS_DIR) | grep -v '^_' | sed 's/\.framework//' | sort)

# Get list of already generated frameworks
GENERATED_FRAMEWORKS := $(shell find $(GENERATED_DIR) -maxdepth 1 -type d -exec basename {} \; 2>/dev/null | grep -v '^frameworks$$' | sort)

# Convert framework name to lowercase directory name
to_lower = $(shell echo $(1) | tr '[:upper:]' '[:lower:]')

# Default target
help:
	@echo "$(BLUE)Available targets:$(NC)"
	@echo ""
	@echo "$(YELLOW)Listing:$(NC)"
	@echo "  $(GREEN)list-frameworks$(NC)            - List all available macOS frameworks"
	@echo "  $(GREEN)list-frameworks$(NC) PATTERN=re - List frameworks matching regex pattern"
	@echo "  $(GREEN)list-priority$(NC)              - List priority frameworks"
	@echo "  $(GREEN)list-generated$(NC)             - List currently generated framework packages"
	@echo ""
	@echo "$(YELLOW)Setup:$(NC)"
	@echo "  $(GREEN)setup-priority$(NC)             - Create directories for priority frameworks"
	@echo "  $(GREEN)setup-pattern$(NC) PATTERN=re   - Create directories for frameworks matching pattern"
	@echo "  $(GREEN)setup-frameworks$(NC)           - Create directories for all frameworks"
	@echo ""
	@echo "$(YELLOW)Generate:$(NC)"
	@echo "  $(GREEN)generate$(NC) FW=<name>         - Generate bindings for a specific framework"
	@echo "  $(GREEN)generate-pattern$(NC) PATTERN=re- Generate bindings for frameworks matching pattern"
	@echo "  $(GREEN)generate-priority$(NC)          - Generate bindings for priority frameworks"
	@echo "  $(GREEN)generate-all$(NC)               - Generate bindings for all frameworks"
	@echo ""
	@echo "$(YELLOW)Other:$(NC)"
	@echo "  $(GREEN)clean-frameworks$(NC)           - Remove all generated framework directories"
	@echo "  $(GREEN)test$(NC)                       - Run tests"
	@echo ""
	@echo "$(BLUE)Common patterns:$(NC)"
	@echo "  ^Core         - Frameworks starting with 'Core'"
	@echo "  Kit$$          - Frameworks ending with 'Kit'"
	@echo "  ^(Metal|Scene|Sprite)Kit$$ - Specific graphics frameworks"
	@echo ""
	@echo "$(BLUE)Examples:$(NC)"
	@echo "  make list-frameworks PATTERN='^Core'"
	@echo "  make setup-pattern PATTERN='Kit$$'"
	@echo "  make generate FW=Metal"
	@echo "  make generate-pattern PATTERN='^(CoreML|Vision)$$'"

# List all available macOS frameworks
# Usage: make list-frameworks [PATTERN=regex]
list-frameworks:
	@if [ -n "$(PATTERN)" ]; then \
		echo "$(BLUE)Frameworks matching '$(PATTERN)':$(NC)"; \
		count=0; \
		for fw in $(ALL_FRAMEWORKS); do \
			if echo "$$fw" | grep -qE '$(PATTERN)'; then \
				echo "  $$fw"; \
				count=$$((count + 1)); \
			fi; \
		done; \
		echo "$(BLUE)Total: $$count frameworks$(NC)"; \
	else \
		echo "$(BLUE)Public macOS frameworks ($(words $(ALL_FRAMEWORKS)) total):$(NC)"; \
		for fw in $(ALL_FRAMEWORKS); do echo "  $$fw"; done | sort; \
	fi

# List priority frameworks
list-priority:
	@echo "$(BLUE)Priority frameworks ($(words $(PRIORITY_FRAMEWORKS)) total):$(NC)"
	@for fw in $(PRIORITY_FRAMEWORKS); do echo "  $$fw"; done


# List currently generated framework packages
list-generated:
	@echo "$(BLUE)Generated framework packages:$(NC)"
	@if [ -d "$(GENERATED_DIR)" ]; then \
		ls -1 $(GENERATED_DIR) 2>/dev/null || echo "  No frameworks generated yet"; \
	else \
		echo "  No frameworks generated yet"; \
	fi
	@echo ""
	@echo "$(BLUE)Frameworks with complete bindings (have doc.go):$(NC)"
	@find $(GENERATED_DIR) -name "doc.go" -exec dirname {} \; 2>/dev/null | xargs -n1 basename | sort | sed 's/^/  /' || echo "  None"

# Setup directories for priority frameworks
setup-priority:
	@echo "$(BLUE)Setting up priority frameworks...$(NC)"
	@for framework in $(PRIORITY_FRAMEWORKS); do \
		dir="$(GENERATED_DIR)/$$(echo $$framework | tr '[:upper:]' '[:lower:]')"; \
		if [ ! -d "$$dir" ]; then \
			echo "  $(GREEN)Creating:$(NC) $$dir"; \
			mkdir -p "$$dir"; \
		fi; \
	done
	@./scripts/create-gen-files.sh
	@echo "$(GREEN)Priority frameworks setup complete$(NC)"

# Setup directories for frameworks matching a pattern
# Usage: make setup-pattern PATTERN='^Core'
setup-pattern:
	@if [ -z "$(PATTERN)" ]; then \
		echo "$(RED)Error: PATTERN parameter required$(NC)"; \
		echo "Usage: make setup-pattern PATTERN=<regex>"; \
		echo "Example: make setup-pattern PATTERN='^Core'"; \
		exit 1; \
	fi
	@echo "$(BLUE)Setting up frameworks matching '$(PATTERN)'...$(NC)"
	@count=0; \
	for framework in $(ALL_FRAMEWORKS); do \
		if echo "$$framework" | grep -qE "$(PATTERN)"; then \
			dir="$(GENERATED_DIR)/$$(echo $$framework | tr '[:upper:]' '[:lower:]')"; \
			if [ ! -d "$$dir" ]; then \
				echo "  $(GREEN)Creating:$(NC) $$dir"; \
				mkdir -p "$$dir"; \
				count=$$((count + 1)); \
			fi; \
		fi; \
	done; \
	echo "$(BLUE)Created $$count directories$(NC)"
	@./scripts/create-gen-files.sh
	@echo "$(GREEN)Setup complete$(NC)"

# Setup directories for all frameworks
setup-frameworks:
	@echo "$(BLUE)Creating directories for all $(words $(ALL_FRAMEWORKS)) frameworks...$(NC)"
	@for framework in $(ALL_FRAMEWORKS); do \
		dir="$(GENERATED_DIR)/$$(echo $$framework | tr '[:upper:]' '[:lower:]')"; \
		if [ ! -d "$$dir" ]; then \
			echo "  $(GREEN)Creating:$(NC) $$dir"; \
			mkdir -p "$$dir"; \
		fi; \
	done
	@./scripts/create-gen-files.sh
	@echo "$(GREEN)Setup complete$(NC). Run 'make generate-priority' or 'make generate-all' to generate bindings."

# Generate bindings for a specific framework
# Usage: make generate FW=Metal
generate:
	@if [ -z "$(FW)" ]; then \
		echo "$(RED)Error: FW parameter required$(NC)"; \
		echo "Usage: make generate FW=<FrameworkName>"; \
		echo "Example: make generate FW=Metal"; \
		exit 1; \
	fi
	@echo "$(BLUE)Generating bindings for $(FW)...$(NC)"
	@generate-framework-bindings -framework "$(FW)" -output "$(GENERATED_DIR)"
	@echo "$(GREEN)Successfully generated $(FW)$(NC)"

# Generate bindings for frameworks matching a pattern
# Usage: make generate-pattern PATTERN='^Core'
generate-pattern:
	@if [ -z "$(PATTERN)" ]; then \
		echo "$(RED)Error: PATTERN parameter required$(NC)"; \
		echo "Usage: make generate-pattern PATTERN=<regex>"; \
		echo "Example: make generate-pattern PATTERN='^Core'"; \
		exit 1; \
	fi
	@echo "$(BLUE)Generating bindings for frameworks matching '$(PATTERN)'...$(NC)"; \
	SUCCESS=0; \
	FAILED=0; \
	SKIPPED=0; \
	for framework in $(ALL_FRAMEWORKS); do \
		if echo "$$framework" | grep -qE "$(PATTERN)"; then \
			OUTPUT_DIR="$(GENERATED_DIR)/$$(echo $$framework | tr '[:upper:]' '[:lower:]')"; \
			if [ -f "$$OUTPUT_DIR/doc.go" ]; then \
				echo "  $(YELLOW)✓$(NC) $$framework (already generated)"; \
				SKIPPED=$$((SKIPPED + 1)); \
			else \
				echo "  $(BLUE)Generating:$(NC) $$framework"; \
				if timeout 300 generate-framework-bindings -framework "$$framework" -output "$(GENERATED_DIR)" 2>&1; then \
					echo "  $(GREEN)✓$(NC) $$framework"; \
					SUCCESS=$$((SUCCESS + 1)); \
				else \
					echo "  $(RED)✗$(NC) $$framework (failed)"; \
					FAILED=$$((FAILED + 1)); \
				fi; \
			fi; \
		fi; \
	done; \
	echo ""; \
	echo "$(BLUE)Summary: Success=$$SUCCESS Failed=$$FAILED Skipped=$$SKIPPED$(NC)"

# Generate bindings for priority frameworks
generate-priority:
	@echo "$(BLUE)Generating bindings for $(words $(PRIORITY_FRAMEWORKS)) priority frameworks...$(NC)"
	@./scripts/generate-priority-frameworks.sh

# Generate bindings for all frameworks
generate-all:
	@echo "$(BLUE)Generating bindings for all $(words $(ALL_FRAMEWORKS)) frameworks...$(NC)"
	@./scripts/generate-all-frameworks.sh

# Clean all generated framework directories
clean-frameworks:
	@echo "$(YELLOW)Removing all generated framework directories...$(NC)"
	@rm -rf $(GENERATED_DIR)/*
	@echo "$(GREEN)Clean complete$(NC)"

# Run tests
test:
	@echo "$(BLUE)Running binding generator tests...$(NC)"
	@cd cmd/generate-framework-bindings && go test -v
	@echo ""
	@echo "$(BLUE)Testing generated bindings compilation...$(NC)"
	@if [ -d "$(GENERATED_DIR)/coregraphics" ]; then \
		cd $(GENERATED_DIR)/coregraphics && go build; \
		echo "$(GREEN)All tests passed$(NC)"; \
	else \
		echo "$(YELLOW)Warning: CoreGraphics not generated yet$(NC)"; \
	fi
