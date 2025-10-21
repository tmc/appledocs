# macOS Virtualization Examples

This directory contains examples of using the Apple Virtualization framework through
generated Go bindings from `github.com/tmc/appledocs/generated/virtualization`.

## Overview

These examples demonstrate how to create and configure virtual machines on macOS using
pure Go bindings (via purego) instead of cgo. The bindings are automatically generated
from Apple's official documentation.

## Examples

### 1. vz/ - Complete macOS VM Runner ⭐

**Status**: ✅ Complete implementation (⚠️ blocked on Foundation bindings)

A full-featured macOS VM implementation equivalent to [Code-Hex/vz](https://github.com/Code-Hex/vz).

```bash
cd vz
make setup          # Initialize VM
make run            # Run VM (once Foundation is fixed)
```

**Features**:
- Complete CLI with 10+ flags
- Platform configuration (hardware model, machine ID)
- Full device support (storage, graphics, network, audio, input)
- Shared directories via VirtioFS
- Comprehensive documentation and scripts

**Size**: ~1,600 lines of code + documentation

See [vz/README.md](vz/README.md) for details.

### 2. vm-create/ - Quick Start Examples

**Status**: ✅ Working examples

Simpler examples for getting started:

- **macos_vm.go** - macOS VM configuration using generated bindings
- **main_quickstart.go** - Simplified Linux VM example
- **main.go** - Full-featured VM implementation

**Size**: ~400 lines total

## Quick Start

### Option 1: Full VM (vz/)

```bash
cd vz
./scripts/setup-vm.sh --disk-size 64
# Currently blocked - see vz/STATUS.md
```

### Option 2: Quick Example (vm-create/)

```bash
cd vm-create
go build -o vm-create main_quickstart.go
# Currently blocked - see ../vz/STATUS.md
```

## Comparison: Generated Bindings vs Code-Hex/vz

These implementations are based on the excellent [Code-Hex/vz](https://github.com/Code-Hex/vz)
project but use our generated bindings instead of cgo.

### Code-Hex/vz Approach (cgo)
```go
import "github.com/Code-Hex/vz/v3"

// Uses cgo and Objective-C wrapper code
config, err := vz.NewVirtualMachineConfiguration(
    bootLoader,
    computeCPUCount(),
    computeMemorySize(),
)
```

### Generated Bindings Approach (purego)
```go
import "github.com/tmc/appledocs/generated/virtualization"

// Uses pure Go with purego for Objective-C runtime
config := virtualization.NewVZVirtualMachineConfiguration()
config.SetBootLoader(unsafe.Pointer(bootLoader.ID))
config.SetCpuCount(cpuCount)
config.SetMemorySize(memorySize)
```

## Key Differences

1. **No cgo** - Generated bindings use [ebitengine/purego](https://github.com/ebitengine/purego) for Objective-C interop
2. **Type-safe** - All Objective-C classes are wrapped in Go types
3. **Auto-generated** - Bindings are generated from Apple documentation, not hand-written
4. **Property-style setters** - Uses `SetX()` methods instead of constructor parameters

## Building

```bash
cd vm-create
go build -o vm-create
```

## Usage

### Create macOS VM Configuration

```go
package main

import (
    "github.com/tmc/appledocs/examples/virtualization/vm-create"
)

func main() {
    // Creates a macOS VM configuration with:
    // - macOS bootloader
    // - Mac platform (hardware model, machine ID, auxiliary storage)
    // - Graphics device (1920x1200 @ 80 PPI)
    // - Network (NAT)
    // - Storage (64GB disk)
    // - Keyboard and trackpad
    config := main.CreateVMConfig()

    // Use config to create and run VM
    // vm := virtualization.NewVZVirtualMachineWithConfiguration(unsafe.Pointer(config.ID))
    // vm.Start()
}
```

### VM Bundle Structure

The example creates a VM bundle at `~/VM.bundle/`:

```
~/VM.bundle/
├── Disk.img              # Virtual disk image (64GB)
├── AuxiliaryStorage      # macOS auxiliary storage
├── HardwareModel         # Saved hardware model
├── MachineIdentifier     # Machine identifier
└── RestoreImage.ipsw     # macOS restore image (optional)
```

## Features Demonstrated

### Platform Configuration
- ✅ Mac hardware model persistence
- ✅ Machine identifier creation/loading
- ✅ Auxiliary storage setup

### Devices
- ✅ Graphics device with display configuration
- ✅ Block storage devices
- ✅ Network devices (NAT)
- ✅ Pointing devices (mouse + trackpad)
- ✅ Keyboard configuration

### Resource Management
- ✅ CPU count (respects min/max limits)
- ✅ Memory size (respects min/max limits)
- ✅ Disk image creation

## Implementation Notes

### Objective-C Interop

The generated bindings use purego for Objective-C runtime access:

```go
// Class access
class := objc.GetClass("VZVirtualMachine")

// Method calls
objc.Send[ReturnType](objectID, selector, args...)

// String conversion
nsStr := objc.Send[objc.ID](
    objc.ID(objc.GetClass("NSString")),
    objc.RegisterName("stringWithUTF8String:"),
    unsafe.Pointer(unsafe.StringData(goString)),
)
```

### Memory Management

Objects returned from `New*` constructors are autoreleased:

```go
bootLoader := virtualization.NewVZMacOSBootLoader()
// bootLoader.Autorelease() is called automatically in constructor
```

### Type Wrapping

Objective-C types are wrapped in Go structs:

```go
type VZVirtualMachine struct {
    VZObject  // Base class inheritance
    ID objc.ID // Objective-C object ID
}
```

## Next Steps

To create a fully functional macOS VM, you would need to add:

1. **VM Lifecycle Management**
   - State change notification handling
   - Start/stop/pause/resume
   - Error handling

2. **Display Output**
   - VZVirtualMachineView integration
   - Graphics framebuffer access
   - Window management

3. **macOS Installation**
   - Restore image download (VZMacOSRestoreImage)
   - Installation process handling
   - Progress tracking

4. **Advanced Features**
   - Shared directories
   - Audio devices
   - Serial console
   - Network block devices

See the [Code-Hex/vz example/macOS](https://github.com/Code-Hex/vz/tree/main/example/macOS)
for a complete implementation of these features.

## Directory Structure

```
virtualization/
├── vz/                          # Complete VM implementation
│   ├── main.go                  # Entry point (393 lines)
│   ├── config.go                # Configuration (380 lines)
│   ├── bundle.go                # Path management
│   ├── delegate_example.go      # Delegate patterns
│   ├── Makefile                 # Build automation
│   ├── scripts/                 # Helper scripts
│   │   ├── setup-vm.sh         # VM initialization
│   │   ├── clean-vm.sh         # Cleanup
│   │   └── status-vm.sh        # Status check
│   ├── README.md                # Usage guide
│   ├── STATUS.md                # Implementation status
│   ├── CONTRIBUTING.md          # Contribution guide
│   └── IMPLEMENTATION_SUMMARY.md # Architecture
├── vm-create/                   # Quick examples
│   ├── macos_vm.go             # macOS example
│   └── main_quickstart.go      # Linux example
├── README.md                    # This file
└── SUMMARY.md                   # First implementation summary
```

## Current Status

### vz/

**Implementation**: ✅ Complete
**Build**: ⚠️ Blocked on Foundation framework bindings
**Documentation**: ✅ Comprehensive

**Blocking Issue**: The Foundation framework bindings have compilation errors that prevent building. See [vz/STATUS.md](vz/STATUS.md) for detailed analysis and solution.

Once Foundation is fixed:
1. `cd vz && make build` will work
2. All functionality will be testable
3. Full VM lifecycle will be runnable

### vm-create/

**Status**: Same Foundation blocking issue affects these examples.

## Contributing

See [vz/CONTRIBUTING.md](vz/CONTRIBUTING.md) for how to contribute.

**Priority**: Fix Foundation framework bindings to unblock all examples.

## References

- [Apple Virtualization Framework](https://developer.apple.com/documentation/virtualization)
- [Code-Hex/vz](https://github.com/Code-Hex/vz) - Reference implementation
- [ebitengine/purego](https://github.com/ebitengine/purego) - Pure Go Objective-C runtime
- [appledocs](https://github.com/tmc/appledocs) - Binding generator

## License

This example code is provided as-is for educational purposes. See the main repository
LICENSE for details.
