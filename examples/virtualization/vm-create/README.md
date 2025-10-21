# Virtualization Framework Example - VM Create

This example demonstrates Apple's Virtualization framework bindings with two implementations using build tags.

## Implementations

### V1: Generated Bindings (Default)
Uses the generated bindings from `github.com/tmc/appledocs/generated/virtualization` with limited functionality.

**Build:** `go build` or `go build -tags '!v2'`
**Run:** `go run .` or `go run -tags '!v2' .`

**Features:**
- Basic framework overview (20 examples)
- Minimal VM configuration demonstration
- Limited device support due to missing setter methods
- Simple AppKit window integration

### V2: Enhanced with objc.Send
Reimplements full VM functionality using generated bindings + direct `objc.Send` calls to work around missing methods.

**Build:** `go build -tags v2`
**Run:** `go run -tags v2 .`

**Features:**
- ✓ Full Linux VM support with all device types
- ✓ Automatic CPU and memory configuration via `objc.Send`
- ✓ VirtIO block device for disk storage
- ✓ NAT network device for internet connectivity
- ✓ VirtIO graphics device with scanout
- ✓ AppKit window with VZVirtualMachineView
- ✓ Complete VM configuration and startup

## Usage

### Show Framework Overview
```bash
# V1 (default) - Shows extensive framework documentation
go run .

# V2 - Shows enhanced capabilities
go run -tags v2 .
```

### Start a VM with UI

**V1 (Limited):**
```bash
go run . -start -kernel <kernel_path> -disk <disk_path>
```

**V2 (Full Featured):**
```bash
go run -tags v2 . -start -kernel <kernel_path> -disk <disk_path> [-initrd <initrd_path>] [-cmdline <cmdline>]
```

**Example:**
```bash
go run -tags v2 . -start \
  -kernel vmlinuz-6.1.0 \
  -disk ubuntu.img \
  -cmdline "console=ttyS0 root=/dev/vda"
```

## Command-Line Flags

- `-start` - Start a VM with UI display (requires `-kernel` and `-disk`)
- `-kernel <path>` - Path to Linux kernel image
- `-disk <path>` - Path to disk image for VM
- `-initrd <path>` - Path to initrd image (optional)
- `-cmdline <string>` - Linux kernel command line (default: "console=ttyS0")
- `-e2e` - Run end-to-end tests

## Implementation Details

### V1: main.go (`!v2` build tag)
- Uses only generated bindings
- Limited by missing methods in generator output
- Good for demonstration and documentation
- Cannot fully configure or start VMs

### V2: main_v2.go (`v2` build tag)
- Uses generated bindings + direct `objc.Send` calls
- Works around missing setters:
  - `setCPUCount:` via objc.Send
  - `setMemorySize:` via objc.Send
  - `setAttachment:` via objc.Send for devices
  - `setVirtualMachine:` via objc.Send for VM view
  - Array manipulation via NSMutableArray
- Demonstrates how to achieve full functionality without Code-Hex/vz dependency

### Key Techniques in V2

**Setting CPU/Memory:**
```go
objc.Send[bool](
    config.ID,
    objc.RegisterName("setCPUCount:"),
    cpuCount,
)
```

**Creating Arrays:**
```go
arrayClass := objc.GetClass("NSMutableArray")
arrayAlloc := objc.Send[objc.ID](objc.ID(arrayClass), objc.RegisterName("alloc"))
arrayID := objc.Send[objc.ID](arrayAlloc, objc.RegisterName("init"))
```

**Setting Device Attachments:**
```go
objc.Send[bool](
    blockDevice.ID,
    objc.RegisterName("setAttachment:"),
    unsafe.Pointer(attachID),
)
```

## Comparison

| Feature | V1 (!v2) | V2 (v2) |
|---------|----------|---------|
| Framework documentation | ✓ (20 examples) | ✓ (concise) |
| VM configuration | Partial | ✓ Complete |
| CPU/Memory config | ✗ | ✓ via objc.Send |
| Storage devices | ✗ | ✓ via objc.Send |
| Network devices | ✗ | ✓ via objc.Send |
| Graphics devices | ✗ | ✓ via objc.Send |
| VZVirtualMachineView | ✗ | ✓ via objc.Send |
| VM startup | ✗ | ✓ |
| Window display | Basic | ✓ with VM view |

## Requirements

- macOS 11.0+ (Big Sur or later)
- Go 1.21+
- For actual VM startup:
  - Linux kernel image (vmlinuz)
  - Disk image (raw, qcow2, etc.)
  - Optional: initrd image
  - Virtualization entitlement (for app bundles)

## Notes

- V1 demonstrates the current state of generated bindings
- V2 shows how to work around limitations using objc.Send
- V2 is the **ideal implementation** that achieves full functionality
- Neither requires Code-Hex/vz dependency
- V2 serves as a reference for improving the binding generator

## Future Enhancements

The V2 implementation demonstrates what the binding generator should produce:
1. Property setters for all properties
2. Complete constructor variants
3. Proper interface conformance
4. Helper functions for common patterns (array creation, etc.)

## References

- [Apple Virtualization Framework Documentation](https://developer.apple.com/documentation/virtualization)
- [purego](https://github.com/ebitengine/purego) - Pure Go Objective-C runtime
- [appledocs binding generator](../../cmd/generate-framework-bindings/)
