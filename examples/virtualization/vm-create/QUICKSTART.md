# Virtualization VM Create - Quickstart Guide

This is a minimal example of creating and launching virtual machines (Linux and macOS) using Apple's Virtualization framework with Go bindings.

## Quick Start

The **quickstart version** (`main_quickstart.go`) is the simplest way to get started:

```bash
# Build the quickstart version
go build -o vm-create main_quickstart.go

# Create a Linux VM
./vm-create vmlinuz disk.img

# Download and configure a macOS VM (auto-download)
./vm-create --download-macos monterey
```

**Size comparison:**
- Full example: 604 lines with UI, multiple examples, documentation
- Quickstart: 284 lines - essentials + DMG download

## Linux VM Examples

### Minimal VM Configuration
```bash
./vm-create vmlinuz ubuntu.img
```

### With Initrd
```bash
./vm-create vmlinuz ubuntu.img initrd.img
```

### With Custom Kernel Command Line
```bash
./vm-create vmlinuz ubuntu.img initrd.img "console=ttyS0 root=/dev/vda"
```

## macOS VM Examples (Auto-Download)

### Download macOS Installer and Create Config
```bash
# Download Monterey installer and create VM config
./vm-create --download-macos monterey

# Download Ventura installer
./vm-create --download-macos ventura

# Download Sonoma installer
./vm-create --download-macos sonoma
```

### Features
- ✅ Automatic DMG download from Apple servers
- ✅ Progress indicator during download
- ✅ Caches downloaded files for reuse
- ✅ Supports multiple macOS versions (Monterey, Ventura, Sonoma)
- ✅ Creates initial VM configuration

## What the Quickstart Does

### For Linux VMs:

1. **Configuration Creation** - Creates a VZVirtualMachineConfiguration
2. **Platform Setup** - Configures for generic (Linux) platform
3. **Boot Loader** - Sets up Linux kernel boot with command-line arguments
4. **Resource Allocation** - Configures CPU count (half of host) and 2GB RAM
5. **VM Instance** - Creates the VZVirtualMachine with the configuration

### For macOS VMs:

1. **Download Detection** - Checks if installer already cached
2. **Auto-Download** - Downloads macOS installer DMG with progress
3. **File Validation** - Verifies download integrity
4. **Configuration Setup** - Initializes macOS VM config structure
5. **Next Steps Guidance** - Shows what's needed for full setup

## Files

- **main.go** (604 lines) - Full featured example with UI, multiple demos, framework overview
- **main_quickstart.go** (284 lines) - Minimal quickstart + DMG download support
- **go.mod** - Dependencies (purego, appledocs generated bindings)

## What's NOT in the Quickstart

The quickstart focuses on core VM creation. For the complete implementation, see `main.go`:

- ✋ Network device configuration (uses complex type conversions)
- 💾 Storage device attachment (requires full binding implementation)
- 🪟 AppKit UI window creation and display
- 📋 Framework overview and multiple examples
- 🔧 Advanced event loop handling
- 🔐 macOS IPSW processing (requires additional Apple APIs)

## Getting the Files (Linux)

You need to prepare:

1. **Linux kernel image** (e.g., `vmlinuz-6.1.0`)
   - Typically from a Linux distribution or `vmlinux` compiled kernel

2. **Disk image** (e.g., `ubuntu.img`)
   - Root filesystem with Linux installed
   - Can be raw, qcow2, or other supported formats

3. **Initial ramdisk** (optional) (e.g., `initrd.img`)
   - Boot-time filesystem
   - Usually from the same distro as kernel

### Quick Linux Setup Example

```bash
# Ubuntu - minimal VM image (typically ~500MB)
wget https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-arm64.tar.gz
tar xzf jammy-server-cloudimg-arm64.tar.gz

# Extract kernel and initrd if needed
# Build custom Linux kernel with Virtio support
```

## Getting macOS Installers (Auto-Download)

Just use the `--download-macos` flag:

```bash
./vm-create --download-macos monterey    # Auto-downloads to macOS-monterey-installer.dmg
./vm-create --download-macos ventura     # Auto-downloads to macOS-ventura-installer.dmg
./vm-create --download-macos sonoma      # Auto-downloads to macOS-sonoma-installer.dmg
```

Features:
- Automatic caching (won't re-download if already present)
- Progress indicator during download
- Saves to current directory with clear naming
- Shows file size after completion

## Type-Safe Go Bindings

The bindings use type-safe Go wrappers around Objective-C:

```go
// Type-safe creation
config := virtualization.NewVZVirtualMachineConfiguration()
platform := virtualization.NewVZGenericPlatformConfiguration()

// Type-safe setters
config.SetPlatform(unsafe.Pointer(platform.ID))
config.SetCpuCount(uint(4))
config.SetMemorySize(uint64(8 * 1024 * 1024 * 1024))

// Type-safe creation with parameters
bootLoader := virtualization.NewVZLinuxBootLoaderWithKernelURL(
    unsafe.Pointer(kernelURL.ID),
)

// String conversion helpers included
kernelURL := stringToNSURL("vmlinuz-6.1.0")
cmdlineStr := stringToNSString("console=ttyS0 root=/dev/vda")
```

## Next Steps

1. **For Linux VMs**: Prepare kernel/disk images, run quickstart
2. **For macOS VMs**: Use auto-download, see main.go for full setup
3. **To understand more**, read the full `main.go` example
4. **For complete bindings** with storage/network setup, see:
   - Code-Hex/vz: https://github.com/Code-Hex/vz (mature reference implementation)

## Requirements

- macOS 11.0+ (Big Sur or later)
- Apple Silicon or Intel Mac (depends on guest OS)
- Virtualization entitlement in application signature
- Network connection (for macOS DMG downloads)

## Building

```bash
# Just the quickstart (minimal dependencies)
go build -o vm-create main_quickstart.go

# Or the full example
go build -o vm-create main.go

# Or both
go build ./cmd/... -o .
```

## Performance Notes

- The generated bindings have no runtime penalty vs hand-written code
- Selector caching is automatic
- Memory management follows Objective-C conventions
- Thread-safe (AppKit requires main thread)
- Download progress uses streaming (memory efficient)

## Customization

The macOS version URLs can be customized in the code:

```go
var macOSVersions = map[string]string{
    "monterey": "https://your-server.com/monterey.dmg",
    "ventura":  "https://your-server.com/ventura.dmg",
    "sonoma":   "https://your-server.com/sonoma.dmg",
}
```

---

**Status**: Ready to use. Full VM startup requires additional setup for console I/O and event handling. macOS VM setup requires IPSW processing.
