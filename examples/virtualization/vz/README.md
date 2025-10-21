# vz - macOS Virtual Machine Runner

A complete macOS virtual machine implementation using generated Virtualization framework bindings.

## Overview

This is a pure-Go (no cgo) implementation of a macOS VM runner, equivalent to [Code-Hex/vz/example/macOS](https://github.com/Code-Hex/vz/tree/main/example/macOS) but using auto-generated bindings from Apple documentation.

## Features

- ✅ macOS guest support (Apple Silicon)
- ✅ Full VM configuration (CPU, memory, storage, network, graphics)
- ✅ Shared directories with host
- ✅ Graphics device with display configuration
- ✅ Network device (NAT)
- ✅ Audio devices (input/output)
- ✅ Input devices (keyboard, trackpad, mouse)
- ✅ Platform persistence (hardware model, machine ID)
- ⏳ macOS installation support (planned)
- ⏳ Recovery mode boot (planned)

## Quick Start

### 1. Initialize VM Platform Configuration

```bash
# Create the platform configuration (hardware model, machine ID, etc.)
go run . -reinit
```

This creates `~/VM.bundle/` with:
- `HardwareModel` - Persisted hardware model
- `MachineIdentifier` - Unique machine identifier
- `AuxiliaryStorage` - NVRAM and other auxiliary data

### 2. Create Disk Image

```bash
# Create a 64GB disk (default size)
go run . -new-disk

# Or specify custom size
go run . -new-disk -disk-size 128
```

### 3. Run the VM

```bash
# Basic run
go run .

# With shared directory
go run . -shared ~/Documents -mount-tag docs

# Custom disk
go run . -disk-path /path/to/disk.img
```

## Usage

```
Usage of vz:
  -install
        run in install mode
  -install-version string
        specific macOS version to install
  -shared string
        path to directory to share with VM
  -mount-tag string
        tag name for shared directory (default "shared")
  -auto-mount
        auto-mount shared directory in macOS 13+
  -reinit
        reinitialize VM platform configuration
  -new-disk
        create fresh disk image
  -disk-size uint
        disk size in GiB (default 64)
  -disk-path string
        path to custom disk image
  -recovery
        boot in recovery mode
```

## VM Configuration

The implementation automatically configures:

### Platform
- **Hardware Model**: Persisted from supported models
- **Machine Identifier**: Unique identifier (like a serial number)
- **Auxiliary Storage**: NVRAM, boot settings, etc.

### Resources
- **CPU**: `NumCPU - 1` (respects framework min/max limits)
- **Memory**: 4GB (respects framework min/max limits)

### Devices

#### Storage
- Primary disk at `~/VM.bundle/Disk.img` (64GB default)
- VirtIO block device
- Read-write access

#### Graphics
- Mac graphics device
- 1920x1200 @ 80 PPI display
- Full acceleration support

#### Network
- VirtIO network device
- NAT attachment for internet access
- Automatic DHCP configuration

#### Input
- Mac keyboard configuration
- USB screen coordinate pointing device
- Mac trackpad (if available)

#### Audio
- VirtIO sound device
- Input stream (microphone)
- Output stream (speakers)

### Shared Directories
```bash
# Share a directory with the VM
go run . -shared ~/Documents -mount-tag docs

# Inside macOS guest, mount with:
mkdir ~/shared
mount -t virtiofs docs ~/shared
```

## Architecture

### File Structure

```
vz/
├── main.go        # Entry point, CLI, and lifecycle management
├── config.go      # VM configuration setup
├── bundle.go      # VM bundle path management
├── go.mod         # Module dependencies
└── README.md      # This file
```

### VM Bundle Structure

```
~/VM.bundle/
├── Disk.img              # Virtual disk (64GB default)
├── AuxiliaryStorage      # NVRAM and auxiliary data
├── HardwareModel         # Persisted hardware model
├── MachineIdentifier     # Machine identifier (UUID)
└── RestoreImage.ipsw     # macOS restore image (optional)
```

## Implementation Details

### Generated Bindings

Uses bindings from `github.com/tmc/appledocs/generated/virtualization`:

```go
import (
    "github.com/tmc/appledocs/generated/foundation"
    "github.com/tmc/appledocs/generated/virtualization"
)

// Create VM configuration
config := virtualization.NewVZVirtualMachineConfiguration()
config.SetBootLoader(unsafe.Pointer(bootLoader.ID))
config.SetPlatform(unsafe.Pointer(platform.ID))
config.SetCpuCount(cpuCount)
config.SetMemorySize(memorySize)

// Create VM
vm := virtualization.NewVZVirtualMachineWithConfiguration(
    unsafe.Pointer(config.ID),
)
```

### Objective-C Interop

Uses [ebitengine/purego](https://github.com/ebitengine/purego) for Objective-C runtime access:

```go
import "github.com/ebitengine/purego/objc"

// String conversion
nsStr := objc.Send[objc.ID](
    objc.ID(objc.GetClass("NSString")),
    objc.RegisterName("stringWithUTF8String:"),
    unsafe.Pointer(unsafe.StringData(goString)),
)

// Array creation
array := objc.Send[objc.ID](
    objc.ID(objc.GetClass("NSArray")),
    objc.RegisterName("arrayWithObjects:count:"),
    unsafe.Pointer(&objects[0]),
    uintptr(len(objects)),
)
```

### Memory Management

The generated bindings automatically handle Autorelease:

```go
// Constructors call Autorelease() automatically
bootLoader := virtualization.NewVZMacOSBootLoader()
// bootLoader is autoreleased

// Manual retain/release if needed
// objc.Send[objc.ID](bootLoader.ID, objc.RegisterName("retain"))
// objc.Send[objc.ID](bootLoader.ID, objc.RegisterName("release"))
```

## Comparison with Code-Hex/vz

| Feature | Code-Hex/vz | This Implementation |
|---------|-------------|-------------------|
| Language | Go + cgo | Pure Go |
| Objective-C | Manual wrappers | Auto-generated bindings |
| Build | Requires cgo | Pure Go build |
| API Style | Functional options | Property setters |
| Updates | Manual | Regenerate bindings |
| Platform | All platforms (with cgo) | macOS only |

### Code Comparison

#### Code-Hex/vz
```go
config, err := vz.NewVirtualMachineConfiguration(
    bootLoader,
    computeCPUCount(),
    computeMemorySize(),
)
config.SetPlatformVirtualMachineConfiguration(platformConfig)
```

#### This Implementation
```go
config := virtualization.NewVZVirtualMachineConfiguration()
config.SetBootLoader(unsafe.Pointer(bootLoader.ID))
config.SetPlatform(unsafe.Pointer(platform.ID))
config.SetCpuCount(computeCPUCount())
config.SetMemorySize(computeMemorySize())
```

## Building

```bash
# Build binary
go build -o vz

# Run directly
go run .

# Install to $GOPATH/bin
go install
```

## Requirements

- macOS 11.0+ (Big Sur or later)
- Apple Silicon or Intel with virtualization support
- Virtualization framework entitlement (for signed apps)

## Limitations

### Current Implementation

1. **No Installation Support**: Cannot download/install macOS yet
   - Workaround: Use `-reinit` to set up platform, then manually install macOS

2. **No Completion Handlers**: Async operations not fully implemented
   - VM start/stop work but without proper completion callbacks

3. **No State Monitoring**: VM state changes not monitored
   - Cannot detect running/stopped/error states dynamically

4. **No Recovery Mode**: Recovery boot not implemented
   - Flag exists but functionality pending

5. **No Graphics Output**: No window/framebuffer integration
   - VM runs headless; need VZVirtualMachineView integration

### Planned Enhancements

- [ ] macOS installation from restore images
- [ ] Progress tracking for downloads and installation
- [ ] State change notification handling
- [ ] Completion handler support
- [ ] Recovery mode boot
- [ ] Graphics window integration
- [ ] Save/restore VM state
- [ ] Snapshot support

## Troubleshooting

### "Platform configuration not found"

```bash
# Run initialization first
go run . -reinit
```

### "Failed to create disk"

```bash
# Ensure ~/VM.bundle exists and is writable
mkdir -p ~/VM.bundle
chmod 755 ~/VM.bundle
```

### "Failed to load hardware model"

```bash
# Reinitialize platform configuration
go run . -reinit
```

### VM doesn't start

Check that:
1. Platform configuration exists (`-reinit` if not)
2. Disk image exists (`-new-disk` if not)
3. You have virtualization entitlements (for signed apps)
4. macOS is installed on the disk (or use `-install`)

## Examples

### Basic Setup and Run

```bash
# 1. Initialize platform
go run . -reinit

# 2. Create disk
go run . -new-disk -disk-size 64

# 3. Install macOS (when supported)
# go run . -install

# 4. Run VM
go run .
```

### With Shared Directory

```bash
# Share Documents folder
go run . -shared ~/Documents -mount-tag docs

# Inside macOS guest:
mkdir ~/shared
mount -t virtiofs docs ~/shared
ls ~/shared  # See host Documents
```

### Custom Disk

```bash
# Use existing disk
go run . -disk-path ~/path/to/macos.img

# Or create fresh disk
go run . -new-disk -disk-size 128
```

### Multiple VMs

Change `VMPaths` in `bundle.go` to use different bundle directories:

```go
// Edit bundle.go
bundlePath := filepath.Join(home, "VM-dev.bundle")  // Different VM
```

## Development

### Regenerate Bindings

```bash
cd /path/to/appledocs
./cmd/generate-framework-bindings/generate-framework-bindings \
  -framework Virtualization \
  -output generated
```

### Debug Mode

Add logging to see Objective-C calls:

```go
import "log"

// Before objc.Send calls
log.Printf("Calling %s on %v", selector, object)
```

### Testing

```bash
# Build
go build

# Run tests (when added)
go test ./...
```

## References

- [Apple Virtualization Framework](https://developer.apple.com/documentation/virtualization)
- [Code-Hex/vz](https://github.com/Code-Hex/vz) - Reference implementation
- [ebitengine/purego](https://github.com/ebitengine/purego) - Pure Go Objective-C runtime
- [appledocs](https://github.com/tmc/appledocs) - Binding generator

## License

See main repository LICENSE.
