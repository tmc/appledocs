# FSKit Framework Bindings Demo (Simple)

This example demonstrates that the FSKit framework can be loaded using purego-based bindings without cgo.

## Overview

FSKit is Apple's File System Kit framework for building file system extensions on macOS. It provides:
- File system extension development APIs
- Volume management
- File system metadata handling
- Integration with macOS file system stack

## Building

```bash
go build .
```

## Running

```bash
# Show framework overview
./fskit-simple
```

## Features Demonstrated

1. **Framework Loading** - Loading FSKit framework via purego
2. **Available Functions** - Error conversion utilities
3. **Minimum Requirements** - macOS 15.4+

## Available FSKit Functions

This demo shows that the following FSKit functions are available through the generated bindings:

- `fs_errorForCocoaError` - Convert Cocoa error to FSKit error
- `fs_errorForMachError` - Convert Mach error to FSKit error
- `fs_errorForPOSIXError` - Convert POSIX error to FSKit error

## Requirements

- macOS 15.4+ (Sequoia)
- Go 1.24.1 or later
- FSKit framework (included in macOS 15.4+)

## Limitations

This is a minimal example that demonstrates framework loading only. It does not:
- Create file system extensions
- Implement volume management
- Handle file operations

For full FSKit functionality, you would need to:
- Implement file system extension protocols
- Handle volume lifecycle
- Manage file system metadata
- Integrate with macOS System Extensions

## File System Extensions

FSKit enables creating file system extensions that:
- Run in user space (not kernel space)
- Provide custom file system implementations
- Integrate with macOS Finder
- Support Time Machine backups
- Enable Spotlight indexing

## Use Cases

- Custom file systems (archive formats, network protocols)
- Virtual file systems (FUSE alternatives)
- File system overlays
- Encryption layers
- Compression file systems
- Network file systems

## References

- [FSKit Framework](https://developer.apple.com/documentation/fskit)
- [File System Extensions](https://developer.apple.com/documentation/fskit/file_system_extensions)
- [Building a File System Extension](https://developer.apple.com/documentation/fskit/building_a_file_system_extension)
- [System Extensions](https://developer.apple.com/documentation/systemextensions)

## See Also

- `../fskit/` - More comprehensive FSKit example with e2e tests
