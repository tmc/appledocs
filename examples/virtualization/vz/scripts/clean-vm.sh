#!/bin/bash
# Clean script for vz macOS VM
# Removes VM bundle with optional backup

set -e

echo "==================================="
echo "vz - VM Cleanup Script"
echo "==================================="
echo ""

# Parse arguments
BACKUP=true
FORCE=false

while [[ $# -gt 0 ]]; do
    case $1 in
        --no-backup)
            BACKUP=false
            shift
            ;;
        --force)
            FORCE=true
            shift
            ;;
        --help)
            echo "Usage: $0 [options]"
            echo ""
            echo "Options:"
            echo "  --no-backup    Don't create backup before removing"
            echo "  --force        Don't ask for confirmation"
            echo "  --help         Show this help"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            echo "Run with --help for usage"
            exit 1
            ;;
    esac
done

# Check if VM bundle exists
if [ ! -d ~/VM.bundle ]; then
    echo "No VM bundle found at ~/VM.bundle"
    exit 0
fi

# Show what will be removed
echo "VM Bundle: ~/VM.bundle/"
ls -lh ~/VM.bundle/
echo ""

TOTAL_SIZE=$(du -sh ~/VM.bundle/ | cut -f1)
echo "Total size: $TOTAL_SIZE"
echo ""

# Confirm deletion
if [ "$FORCE" = false ]; then
    if [ "$BACKUP" = true ]; then
        read -p "Create backup and remove VM bundle? (y/N): " -n 1 -r
    else
        read -p "⚠️  Permanently remove VM bundle? (y/N): " -n 1 -r
    fi
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "Cancelled."
        exit 0
    fi
fi

# Create backup if requested
if [ "$BACKUP" = true ]; then
    BACKUP_PATH=~/VM.bundle.backup-$(date +%Y%m%d-%H%M%S)
    echo "Creating backup at $BACKUP_PATH..."
    cp -R ~/VM.bundle "$BACKUP_PATH"
    echo "✓ Backup created"
fi

# Remove VM bundle
echo "Removing VM bundle..."
rm -rf ~/VM.bundle

echo ""
echo "✓ VM bundle removed"

if [ "$BACKUP" = true ]; then
    echo ""
    echo "Backup available at: $BACKUP_PATH"
    echo "To restore: mv \"$BACKUP_PATH\" ~/VM.bundle"
fi
