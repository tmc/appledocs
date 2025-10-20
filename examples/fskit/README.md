# FSKit Framework Bindings Demo

This example demonstrates that the FSKit framework can be loaded using purego-based bindings without cgo, with basic end-to-end testing.

## Overview

FSKit is Apple's File System Kit framework introduced in macOS 15.4 (Sequoia) for building file system extensions. It enables:
- User-space file system implementations
- Custom file system protocols
- Integration with macOS file system stack
- System Extensions for file systems

## Building

```bash
go build .
```

## Running

```bash
# Show framework overview
./fskit

# Run end-to-end tests
./fskit -e2e
```

## Features Demonstrated

1. **Framework Loading** - Demonstrates FSKit framework can be loaded via purego
2. **Basic Testing** - Simple e2e tests to verify framework accessibility
3. **No CGO Required** - Pure Go bindings without C dependencies

## E2E Tests

The `-e2e` flag runs basic tests that verify:

1. **Framework loads successfully** - FSKit framework can be loaded
2. **Package is accessible** - Generated bindings are importable

These tests ensure that the Go bindings are correctly generated and the framework is available on the system.

## FSKit Framework

FSKit provides APIs for creating file system extensions that:

### Core Capabilities
- **File System Extensions** - Create custom file systems in user space
- **Volume Management** - Handle volume lifecycle (mount, unmount, format)
- **File Operations** - Implement read, write, create, delete operations
- **Directory Operations** - Handle directory listing and navigation
- **Metadata Management** - Manage file attributes and extended attributes
- **Error Handling** - Convert between error types (POSIX, Mach, Cocoa)

### Integration Points
- **Finder Integration** - Custom file systems appear in Finder
- **Time Machine Support** - Enable backups for custom file systems
- **Spotlight Integration** - Index custom file system content
- **Quick Look Support** - Preview files in custom formats
- **Tags and Labels** - Support macOS file tagging

### System Extension Model
FSKit uses the System Extensions architecture:
- Runs in user space (not kernel)
- Sandboxed for security
- Managed via System Preferences
- Requires user approval for installation

## Use Cases

### Network File Systems
- SMB/CIFS clients
- NFS clients
- WebDAV
- Cloud storage (Dropbox, Google Drive, etc.)

### Archive File Systems
- ZIP file systems
- ISO mounting
- TAR archives
- Custom archive formats

### Encryption Layers
- Encrypted volumes
- Per-file encryption
- Cloud storage encryption

### Virtual File Systems
- Union file systems
- Overlay file systems
- RAM disks
- Caching layers

### Special Purpose
- Git file system (filesystem view of git repos)
- Database as file system
- API as file system
- Process file system (/proc-like)

## Requirements

- macOS 15.4+ (Sequoia or later)
- Go 1.24.1 or later
- FSKit framework (included in macOS 15.4+)
- System Extensions entitlement (for actual file system extensions)

## Limitations

This example demonstrates framework loading only. It does not:
- Create actual file system extensions
- Implement file operations
- Handle volume management
- Register with System Extensions

### Creating Real File System Extensions

To create a functional file system extension, you would need to:

1. **Implement Extension Protocol**
   - Subclass FSFileSystem
   - Implement FSFileSystemOperations
   - Handle volume lifecycle events

2. **Handle File Operations**
   - Read/write operations
   - Create/delete operations
   - Directory operations
   - Metadata management

3. **System Extension Setup**
   - Create System Extension target
   - Configure Info.plist
   - Add FSKit entitlements
   - Handle activation requests

4. **Testing and Deployment**
   - Test with System Extensions
   - Handle user approval
   - Debug extension crashes
   - Monitor system logs

## FSKit Architecture

```
┌─────────────────────────────────────┐
│        User Applications            │
│     (Finder, Terminal, Apps)        │
└────────────────┬────────────────────┘
                 │
┌────────────────┴────────────────────┐
│        macOS File System Stack      │
│           (VFS Layer)               │
└────────────────┬────────────────────┘
                 │
┌────────────────┴────────────────────┐
│         FSKit Framework             │
│  (User Space File System Manager)   │
└────────────────┬────────────────────┘
                 │
┌────────────────┴────────────────────┐
│    Your File System Extension       │
│    (System Extension Bundle)        │
└─────────────────────────────────────┘
```

## Error Handling

FSKit provides error conversion functions:

- `fs_errorForPOSIXError` - Convert POSIX errno to FSError
- `fs_errorForMachError` - Convert Mach error to FSError
- `fs_errorForCocoaError` - Convert NSError to FSError

These ensure consistent error handling across different error domains.

## Security Considerations

File system extensions:
- Run in user space with restricted permissions
- Require System Extensions entitlement
- Need user approval for installation
- Are sandboxed for security
- Can be uninstalled via System Preferences
- Are monitored by the system for crashes

## References

- [FSKit Framework](https://developer.apple.com/documentation/fskit)
- [File System Extensions](https://developer.apple.com/documentation/fskit/file_system_extensions)
- [Building a File System Extension](https://developer.apple.com/documentation/fskit/building_a_file_system_extension)
- [System Extensions](https://developer.apple.com/documentation/systemextensions)
- [System Extension Lifecycle](https://developer.apple.com/documentation/systemextensions/system_extension_lifecycle)

## See Also

- `../fskit-simple/` - Minimal FSKit loading example
- `../../generated/fskit/` - Generated FSKit bindings
