# Makefile for vz - macOS VM runner

.PHONY: help build run clean install test reinit new-disk check lint fmt

# Default target
help:
	@echo "vz - macOS Virtual Machine Runner"
	@echo ""
	@echo "Available targets:"
	@echo "  make build        Build the vz binary"
	@echo "  make run          Run the VM"
	@echo "  make install      Install to \$$GOPATH/bin"
	@echo "  make clean        Clean build artifacts"
	@echo "  make test         Run tests"
	@echo "  make check        Check if dependencies are ready"
	@echo "  make lint         Run linters"
	@echo "  make fmt          Format code"
	@echo ""
	@echo "VM Setup:"
	@echo "  make reinit       Initialize VM platform"
	@echo "  make new-disk     Create fresh disk image"
	@echo ""
	@echo "Examples:"
	@echo "  make run ARGS=-shared=~/Documents"
	@echo "  make run ARGS=\"-new-disk -disk-size=128\""

# Build the binary
build:
	@echo "Building vz..."
	go build -o vz

# Install to GOPATH/bin
install:
	@echo "Installing vz to \$$GOPATH/bin..."
	go install

# Run the VM
run: build
	@echo "Running VM..."
	./vz $(ARGS)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f vz
	go clean

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Check dependencies
check:
	@echo "Checking dependencies..."
	@go version
	@echo "Go version: OK"
	@if go list -m github.com/ebitengine/purego > /dev/null 2>&1; then \
		echo "purego: OK"; \
	else \
		echo "purego: MISSING - run 'go mod tidy'"; \
	fi
	@if go list -m github.com/tmc/appledocs > /dev/null 2>&1; then \
		echo "appledocs: OK"; \
	else \
		echo "appledocs: MISSING - check go.mod replace directive"; \
	fi
	@echo ""
	@echo "Build status:"
	@if go build -o /dev/null 2>&1; then \
		echo "  ✓ Builds successfully"; \
	else \
		echo "  ✗ Build fails (likely Foundation binding issues)"; \
		echo "  See STATUS.md for details"; \
	fi

# Lint code
lint:
	@echo "Running linters..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found, running basic checks..."; \
		go vet ./...; \
	fi

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Initialize VM platform
reinit:
	@echo "Reinitializing VM platform..."
	@if [ -f vz ]; then \
		./vz -reinit; \
	else \
		go run . -reinit; \
	fi

# Create new disk
new-disk:
	@echo "Creating new disk..."
	@if [ -f vz ]; then \
		./vz -new-disk $(if $(SIZE),-disk-size=$(SIZE),); \
	else \
		go run . -new-disk $(if $(SIZE),-disk-size=$(SIZE),); \
	fi

# Show VM bundle status
status:
	@echo "VM Bundle Status:"
	@echo "================="
	@if [ -d ~/VM.bundle ]; then \
		echo "Bundle directory: ✓ exists"; \
		if [ -f ~/VM.bundle/HardwareModel ]; then \
			echo "Hardware model:   ✓ exists"; \
		else \
			echo "Hardware model:   ✗ missing (run 'make reinit')"; \
		fi; \
		if [ -f ~/VM.bundle/MachineIdentifier ]; then \
			echo "Machine ID:       ✓ exists"; \
		else \
			echo "Machine ID:       ✗ missing (run 'make reinit')"; \
		fi; \
		if [ -f ~/VM.bundle/Disk.img ]; then \
			SIZE=$$(du -h ~/VM.bundle/Disk.img | cut -f1); \
			echo "Disk image:       ✓ exists ($$SIZE)"; \
		else \
			echo "Disk image:       ✗ missing (run 'make new-disk')"; \
		fi; \
		if [ -f ~/VM.bundle/AuxiliaryStorage ]; then \
			echo "Auxiliary:        ✓ exists"; \
		else \
			echo "Auxiliary:        ✗ missing (created on first run)"; \
		fi; \
	else \
		echo "Bundle directory: ✗ missing"; \
		echo ""; \
		echo "Run 'make reinit' to initialize"; \
	fi

# Quick setup (reinit + disk)
setup: reinit new-disk
	@echo ""
	@echo "✓ VM setup complete!"
	@echo "Run 'make run' to start the VM"

# Development: watch and rebuild on changes
watch:
	@echo "Watching for changes..."
	@while true; do \
		inotifywait -qq -r -e modify *.go 2>/dev/null || \
		fswatch -1 -r . 2>/dev/null || \
		(echo "Install fswatch for auto-rebuild" && exit 1); \
		make build; \
	done

# Show VM bundle contents
ls:
	@if [ -d ~/VM.bundle ]; then \
		ls -lh ~/VM.bundle/; \
	else \
		echo "VM bundle not found at ~/VM.bundle"; \
	fi

# Backup VM bundle
backup:
	@echo "Backing up VM bundle..."
	@BACKUP=~/VM.bundle.backup-$$(date +%Y%m%d-%H%M%S); \
	if [ -d ~/VM.bundle ]; then \
		cp -R ~/VM.bundle $$BACKUP && \
		echo "✓ Backup created at $$BACKUP"; \
	else \
		echo "✗ No VM bundle to backup"; \
	fi

# Restore VM bundle from backup
restore:
	@echo "Available backups:"
	@ls -dt ~/VM.bundle.backup-* 2>/dev/null || echo "No backups found"
	@echo ""
	@echo "To restore, run: cp -R ~/VM.bundle.backup-YYYYMMDD-HHMMSS ~/VM.bundle"
