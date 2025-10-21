#!/bin/bash
# Status script for vz macOS VM
# Shows current VM bundle status and configuration

echo "==================================="
echo "vz - VM Status"
echo "==================================="
echo ""

# Check VM bundle
if [ ! -d ~/VM.bundle ]; then
    echo "❌ No VM bundle found at ~/VM.bundle"
    echo ""
    echo "Run './scripts/setup-vm.sh' to create one"
    exit 1
fi

echo "📁 VM Bundle: ~/VM.bundle/"
echo ""

# Platform configuration
echo "Platform Configuration:"
echo "----------------------"

if [ -f ~/VM.bundle/HardwareModel ]; then
    SIZE=$(du -h ~/VM.bundle/HardwareModel | cut -f1)
    echo "  Hardware Model:      ✓ ($SIZE)"
else
    echo "  Hardware Model:      ❌ Missing"
fi

if [ -f ~/VM.bundle/MachineIdentifier ]; then
    SIZE=$(du -h ~/VM.bundle/MachineIdentifier | cut -f1)
    echo "  Machine Identifier:  ✓ ($SIZE)"
else
    echo "  Machine Identifier:  ❌ Missing"
fi

if [ -f ~/VM.bundle/AuxiliaryStorage ]; then
    SIZE=$(du -h ~/VM.bundle/AuxiliaryStorage | cut -f1)
    echo "  Auxiliary Storage:   ✓ ($SIZE)"
else
    echo "  Auxiliary Storage:   ⏸  (Created on first run)"
fi

echo ""

# Disk image
echo "Storage:"
echo "--------"

if [ -f ~/VM.bundle/Disk.img ]; then
    SIZE=$(du -h ~/VM.bundle/Disk.img | cut -f1)
    SIZE_BYTES=$(du -b ~/VM.bundle/Disk.img | cut -f1)
    SIZE_GB=$((SIZE_BYTES / 1024 / 1024 / 1024))
    echo "  Disk Image:          ✓ ($SIZE / ${SIZE_GB}GB)"
else
    echo "  Disk Image:          ❌ Missing"
fi

echo ""

# Restore image
echo "Installation:"
echo "------------"

if [ -f ~/VM.bundle/RestoreImage.ipsw ]; then
    SIZE=$(du -h ~/VM.bundle/RestoreImage.ipsw | cut -f1)
    echo "  Restore Image:       ✓ ($SIZE)"
else
    echo "  Restore Image:       ❌ Missing"
fi

echo ""

# Total size
TOTAL=$(du -sh ~/VM.bundle/ | cut -f1)
echo "Total Bundle Size: $TOTAL"

echo ""

# Readiness check
echo "Readiness Check:"
echo "---------------"

READY=true

if [ ! -f ~/VM.bundle/HardwareModel ] || [ ! -f ~/VM.bundle/MachineIdentifier ]; then
    echo "  ❌ Platform not initialized"
    echo "     Run: go run . -reinit"
    READY=false
else
    echo "  ✓ Platform initialized"
fi

if [ ! -f ~/VM.bundle/Disk.img ]; then
    echo "  ❌ Disk not created"
    echo "     Run: go run . -new-disk"
    READY=false
else
    echo "  ✓ Disk created"
fi

if [ ! -f ~/VM.bundle/RestoreImage.ipsw ]; then
    echo "  ⚠️  No restore image (macOS not installed)"
    echo "     This is OK if macOS is already installed on disk"
fi

echo ""

if [ "$READY" = true ]; then
    echo "Status: ✓ Ready to run"
    echo ""
    echo "To start VM:"
    echo "  go run ."
else
    echo "Status: ⚠️  Setup required"
    echo ""
    echo "To complete setup:"
    echo "  ./scripts/setup-vm.sh"
fi

echo ""
echo "Note: Build currently blocked on Foundation bindings"
echo "See STATUS.md for details"
