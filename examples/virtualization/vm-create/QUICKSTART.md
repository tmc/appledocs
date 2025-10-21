# VM Create - Quick Start Guide

## TL;DR

```bash
# See framework overview
go run .

# See enhanced features (V2)
go run -tags v2 .

# Launch actual VM (with kernel and disk)
go run -tags v2 . -start -kernel vmlinuz -disk disk.img
```

## Build Targets

| Command | What It Does |
|---------|-------------|
| `go build` | Build V1 (default) |
| `go build -tags v2` | Build V2 (enhanced) |
| `go run .` | Run V1 directly |
| `go run -tags v2 .` | Run V2 directly |

## Features

### V1 (Default - `go run .`)
- Shows 20+ framework capability examples
- Demonstrates basic VM configuration
- Demonstrates current generated bindings limitations

### V2 (Enhanced - `go run -tags v2 .`)
- Shows enhanced capabilities with workarounds
- Complete device configuration
- VM lifecycle management

## Flags

```bash
-start              # Start a VM with UI display
-kernel <path>      # Path to Linux kernel image
-disk <path>        # Path to disk image
-initrd <path>      # Optional initrd image
-cmdline <string>   # Kernel command line (default: "console=ttyS0")
-e2e                # Run end-to-end tests
-h, --help          # Show help
```

## Examples

### Show Framework Documentation
```bash
# V1: 20 examples of framework capabilities
go run .

# V2: Enhanced features list
go run -tags v2 .
```

### Start a VM
```bash
# Minimal (V2 only)
go run -tags v2 . -start -kernel vmlinuz -disk disk.img

# With initrd
go run -tags v2 . -start \
  -kernel vmlinuz \
  -disk disk.img \
  -initrd initramfs.img

# With custom kernel command line
go run -tags v2 . -start \
  -kernel vmlinuz \
  -disk disk.img \
  -cmdline "console=ttyS0 root=/dev/vda rw"
```

## What Happens When You Start a VM

1. Validates Virtualization framework is supported
2. Creates NSApplication
3. Configures VM:
   - Platform: Generic (for Linux)
   - Boot: Linux boot loader with kernel
   - Storage: VirtIO block device with your disk
   - Network: NAT (internet connectivity)
   - Graphics: VirtIO graphics device
   - Memory: 4GB
   - CPUs: NumCPU - 1 (reserved for host)
4. Creates window (800x600) with VM display
5. Shows window
6. Starts VM
7. Monitors state changes
8. Runs until you close the window

## Files Explained

| File | Purpose |
|------|---------|
| `main.go` | V1 implementation (generated bindings only) |
| `main_v2.go` | V2 implementation (enhanced with workarounds) |
| `README.md` | Complete documentation |
| `QUICKSTART.md` | This file |

## Build Tags Explained

| Tag | Meaning |
|-----|---------|
| `!v2` (default) | Use V1 implementation |
| `v2` | Use V2 implementation |

Think of them as feature flags:
- V1 = What's available today
- V2 = What's possible with workarounds

## Common Issues

### "kernel file not found"
```bash
# Make sure kernel file exists and path is correct
ls -la /path/to/vmlinuz
go run -tags v2 . -start -kernel /path/to/vmlinuz -disk /path/to/disk.img
```

### "disk image not found"
```bash
# Make sure disk image exists
ls -la /path/to/disk.img
go run -tags v2 . -start -kernel /path/to/vmlinuz -disk /path/to/disk.img
```

### Build fails
```bash
# Make sure you're using Go 1.21+
go version

# Clean and rebuild
go clean
go build -tags v2 -v
```

## Testing

Both V1 and V2 build successfully:
```bash
go build              # V1
go build -tags v2    # V2

# Both work
go run .             # V1
go run -tags v2 .    # V2
```

## Performance Notes

- V2 uses 500ms polling for VM state monitoring
- Window updates in real-time as VM runs
- Memory usage: ~4GB for VM + ~100MB for host app

## Next Steps

1. Get a Linux kernel image (vmlinuz)
2. Get or create a disk image with Linux installed
3. Run: `go run -tags v2 . -start -kernel vmlinuz -disk disk.img`
4. See the VM boot and run in the window

---

**Need more details?** See `README.md` for comprehensive documentation.
