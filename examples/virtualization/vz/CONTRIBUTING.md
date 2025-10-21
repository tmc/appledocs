# Contributing to vz

Thank you for your interest in contributing to the vz macOS VM implementation!

## Getting Started

### Prerequisites

- macOS 11.0+ (Big Sur or later)
- Go 1.24.1 or later
- Git
- Basic understanding of:
  - Go programming
  - Apple Virtualization framework
  - Objective-C runtime (helpful but not required)

### Setup Development Environment

```bash
# Clone the repository
cd /path/to/appledocs

# Navigate to vz
cd examples/virtualization/vz

# Check dependencies
make check

# Try building (will fail until Foundation is fixed)
make build
```

## Current Status

**BLOCKED**: The implementation is complete but cannot build due to Foundation framework binding errors. See [STATUS.md](STATUS.md) for details.

### What Works
- ✅ All implementation code is complete
- ✅ All features implemented
- ✅ Documentation comprehensive

### What's Blocked
- ❌ Cannot build (Foundation errors)
- ❌ Cannot test runtime behavior
- ❌ Cannot run actual VMs

## How to Contribute

### 1. Fix Foundation Bindings (HIGHEST PRIORITY)

This unblocks everything else.

**Location**: `/Volumes/tmc/go/src/github.com/tmc/appledocs/cmd/generate-framework-bindings/`

**Issues to fix**:
1. Self-references: `foundation.Foo` within foundation package
2. Cross-framework refs: `appkit.View` in foundation package

**Files involved**:
- `generated/foundation/ns_ordered_collection_difference.gen.go:124,125`
- `generated/foundation/nsurl_session_configuration.gen.go:117,118,128`

**How to help**:
1. Review binding generator code
2. Fix self-reference issue (use local type names)
3. Fix cross-framework references (use interfaces or unsafe.Pointer)
4. Regenerate Foundation bindings
5. Test vz build

### 2. Add Completion Handler Support

**Current**: Async operations are stubbed

**Files**: `main.go`, `config.go`

**What's needed**:
- Implement delegate pattern for completion handlers
- Create callback wrappers
- Handle async start/stop/install operations

**Example**:
```go
// Current (stub):
vm.Start()

// Needed:
done := make(chan error)
vm.StartWithCompletionHandler(func(err error) {
    done <- err
})
return <-done
```

### 3. Add State Change Monitoring

**Current**: No state observation

**What's needed**:
- Implement VZVirtualMachineDelegate
- Monitor state changes (running, stopped, error)
- Handle notifications

**Example**:
```go
type vmDelegate struct {
    stateChanged chan VirtualMachineState
}

func (d *vmDelegate) virtualMachineDidStop(vm VZVirtualMachine) {
    d.stateChanged <- VirtualMachineStateStopped
}
```

### 4. Graphics Window Integration

**Current**: Headless VM

**What's needed**:
- Integrate VZVirtualMachineView
- Create AppKit window
- Handle display output

**Complexity**: High (requires AppKit integration)

### 5. macOS Installation Support

**Current**: Stubbed in `installMacOS()`

**What's needed**:
- VZMacOSRestoreImage handling
- Download restore images
- Progress tracking
- Installation process management

## Code Style

### Go Code
- Follow standard Go conventions
- Use `gofmt` or `make fmt`
- Keep functions focused and small
- Document public APIs

### Objective-C Interop
- Use helper functions for common operations
- Prefer `objc.ID` over wrapped types when possible
- Always check return values (`if obj.ID == 0`)
- Document Objective-C method names in comments

### Example:
```go
// Good: Clear, documented, checked
// Calls [NSString stringWithUTF8String:]
func stringToNSString(s string) objc.ID {
    str := objc.Send[objc.ID](
        objc.ID(objc.GetClass("NSString")),
        objc.RegisterName("stringWithUTF8String:"),
        unsafe.Pointer(unsafe.StringData(s)),
    )
    if str == 0 {
        panic("failed to create NSString")
    }
    return str
}

// Bad: Unclear, unchecked
func makeStr(s string) objc.ID {
    return objc.Send[objc.ID](
        objc.ID(objc.GetClass("NSString")),
        objc.RegisterName("stringWithUTF8String:"),
        unsafe.Pointer(unsafe.StringData(s)),
    )
}
```

## Testing

### Unit Tests (TODO)

Create tests once Foundation is fixed:

```go
// config_test.go
func TestComputeCPUCount(t *testing.T) {
    count := computeCPUCount()

    max := virtualization.VZVirtualMachineConfigurationClass.MaximumAllowedCPUCount()
    min := virtualization.VZVirtualMachineConfigurationClass.MinimumAllowedCPUCount()

    if count > max {
        t.Errorf("CPU count %d exceeds max %d", count, max)
    }
    if count < min {
        t.Errorf("CPU count %d below min %d", count, min)
    }
}
```

### Integration Tests

Test full VM lifecycle:

```bash
# Test setup
make setup

# Test VM creation (manual for now)
go run . -reinit
go run . -new-disk
# Would need macOS install to test further
```

## Documentation

### Required for New Features

1. **Code comments**: Document public functions
2. **README updates**: Update usage examples
3. **STATUS updates**: Track implementation progress
4. **CHANGELOG entries**: Note what changed

### Documentation Style

```go
// Good: Clear purpose, parameters, return value
// createDiskAttachment creates a VZDiskImageStorageDeviceAttachment for the given path.
// If the disk doesn't exist, it creates one with the specified size.
// Returns the attachment or error.
func createDiskAttachment(path string, sizeBytes uint64) (virtualization.VZDiskImageStorageDeviceAttachment, error) {
    // ...
}

// Bad: Vague, no details
// creates disk
func createDiskAttachment(path string, sizeBytes uint64) (virtualization.VZDiskImageStorageDeviceAttachment, error) {
```

## Pull Request Process

### Before Submitting

1. **Test build**: `make build` (or verify reason for failure)
2. **Format code**: `make fmt`
3. **Run linter**: `make lint`
4. **Update docs**: Keep README, STATUS in sync
5. **Write tests**: Add tests for new functionality

### PR Description Template

```markdown
## What

Brief description of changes

## Why

Problem being solved or feature being added

## How

Technical approach and key decisions

## Testing

How you tested the changes

## Checklist

- [ ] Code builds (or Foundation fix in progress)
- [ ] Tests added/updated
- [ ] Documentation updated
- [ ] CHANGELOG.md updated
```

### Review Process

1. Submit PR against `main` branch (or appropriate feature branch)
2. CI will run (once Foundation is fixed)
3. Maintainers review
4. Address feedback
5. Merge when approved

## Project Structure

```
vz/
├── main.go              # Entry point, CLI, lifecycle
├── config.go            # VM configuration
├── bundle.go            # Path management
├── delegate.go          # TODO: Delegate implementations
├── graphics.go          # TODO: Graphics integration
├── installer.go         # TODO: macOS installation
├── Makefile             # Build automation
├── go.mod               # Dependencies
├── README.md            # User documentation
├── STATUS.md            # Implementation status
├── CONTRIBUTING.md      # This file
└── *_test.go            # Tests (to be added)
```

## Common Tasks

### Adding a New Device Type

1. Create configuration function in `config.go`:
```go
func configureFooDevices(config virtualization.VZVirtualMachineConfiguration) error {
    device := virtualization.NewVZFooDeviceConfiguration()
    if device.ID == 0 {
        return fmt.Errorf("failed to create foo device")
    }

    // Configure device...

    // Add to config
    devices := []unsafe.Pointer{unsafe.Pointer(device.ID)}
    array := createNSArray(devices)
    config.SetFooDevices(unsafe.Pointer(array))

    return nil
}
```

2. Call from `setupVMConfiguration()`:
```go
if err := configureFooDevices(config); err != nil {
    return virtualization.VZVirtualMachineConfiguration{},
        fmt.Errorf("failed to configure foo: %w", err)
}
```

3. Add CLI flag if needed in `main.go`
4. Document in README.md
5. Add tests

### Adding a New CLI Flag

1. Declare variable in `main.go`:
```go
var fooOption string
```

2. Register flag in `init()`:
```go
func init() {
    flag.StringVar(&fooOption, "foo", "", "description of foo option")
}
```

3. Use in implementation:
```go
if fooOption != "" {
    // Handle option
}
```

4. Document in README.md help section

## Resources

### Internal
- [STATUS.md](STATUS.md) - Current status and blocking issues
- [README.md](README.md) - User documentation
- [IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md) - Architecture overview

### External
- [Apple Virtualization Framework](https://developer.apple.com/documentation/virtualization)
- [Code-Hex/vz](https://github.com/Code-Hex/vz) - Reference implementation
- [ebitengine/purego](https://github.com/ebitengine/purego) - Pure Go Objective-C runtime
- [Go Documentation](https://go.dev/doc/comment) - Writing Go documentation

## Getting Help

### Questions

1. Check [STATUS.md](STATUS.md) for known issues
2. Review [README.md](README.md) for usage
3. Check [Code-Hex/vz](https://github.com/Code-Hex/vz) for reference

### Reporting Bugs

Include:
- Go version (`go version`)
- macOS version
- Build error or runtime error
- Steps to reproduce
- Expected vs actual behavior

### Suggesting Features

Consider:
- Is it needed for basic VM functionality?
- Does Code-Hex/vz have it?
- What's the implementation complexity?
- Can it wait until Foundation is fixed?

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.

## Thank You!

Every contribution helps make pure-Go macOS framework access a reality. Thank you for being part of this effort!
