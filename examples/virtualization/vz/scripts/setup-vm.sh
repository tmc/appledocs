#!/bin/bash
# Setup script for vz macOS VM
# This script initializes a new VM from scratch

set -e

echo "==================================="
echo "vz - macOS VM Setup Script"
echo "==================================="
echo ""

# Check if we're in the right directory
if [ ! -f "main.go" ]; then
    echo "Error: Must run from vz directory"
    exit 1
fi

# Parse arguments
DISK_SIZE=64
SHARED_DIR=""
MOUNT_TAG="shared"

while [[ $# -gt 0 ]]; do
    case $1 in
        --disk-size)
            DISK_SIZE="$2"
            shift 2
            ;;
        --shared)
            SHARED_DIR="$2"
            shift 2
            ;;
        --mount-tag)
            MOUNT_TAG="$2"
            shift 2
            ;;
        --help)
            echo "Usage: $0 [options]"
            echo ""
            echo "Options:"
            echo "  --disk-size N      Disk size in GiB (default: 64)"
            echo "  --shared PATH      Directory to share with VM"
            echo "  --mount-tag TAG    Tag for shared directory (default: shared)"
            echo "  --help             Show this help"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            echo "Run with --help for usage"
            exit 1
            ;;
    esac
done

# Step 1: Check if VM bundle exists
echo "Step 1: Checking VM bundle..."
if [ -d ~/VM.bundle ]; then
    echo "⚠️  VM bundle already exists at ~/VM.bundle"
    read -p "Remove existing bundle? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "Creating backup..."
        BACKUP=~/VM.bundle.backup-$(date +%Y%m%d-%H%M%S)
        mv ~/VM.bundle "$BACKUP"
        echo "✓ Backed up to $BACKUP"
    else
        echo "Keeping existing bundle. Skipping initialization."
        exit 0
    fi
fi

# Step 2: Initialize platform configuration
echo ""
echo "Step 2: Initializing platform configuration..."
echo "This creates hardware model and machine identifier"
echo ""

go run . -reinit

if [ $? -ne 0 ]; then
    echo "❌ Failed to initialize platform"
    exit 1
fi

echo "✓ Platform initialized"

# Step 3: Create disk image
echo ""
echo "Step 3: Creating disk image (${DISK_SIZE}GB)..."
echo ""

go run . -new-disk -disk-size="$DISK_SIZE"

if [ $? -ne 0 ]; then
    echo "❌ Failed to create disk"
    exit 1
fi

echo "✓ Disk created"

# Step 4: Summary
echo ""
echo "==================================="
echo "Setup Complete!"
echo "==================================="
echo ""
echo "VM Bundle: ~/VM.bundle/"
echo "  - HardwareModel: ✓"
echo "  - MachineIdentifier: ✓"
echo "  - Disk.img: ✓ (${DISK_SIZE}GB)"
echo ""

if [ -n "$SHARED_DIR" ]; then
    echo "Shared directory configured: $SHARED_DIR"
    echo "Mount tag: $MOUNT_TAG"
    echo ""
    echo "To mount in macOS guest:"
    echo "  mkdir ~/shared"
    echo "  mount -t virtiofs $MOUNT_TAG ~/shared"
    echo ""
fi

echo "Next steps:"
echo "  1. Install macOS (when supported):"
echo "     go run . -install"
echo ""
echo "  2. Or run the VM (if macOS already installed):"
if [ -n "$SHARED_DIR" ]; then
    echo "     go run . -shared \"$SHARED_DIR\" -mount-tag \"$MOUNT_TAG\""
else
    echo "     go run ."
fi
echo ""
echo "Note: Build currently blocked on Foundation bindings"
echo "See STATUS.md for details"
