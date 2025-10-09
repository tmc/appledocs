# FSKit Go Bindings

Go bindings for Apple's FSKit framework (macOS 15.4+) using purego.

## About FSKit

FSKit enables you to implement file systems that run in user space on macOS. With FSKit, you can extend macOS by enabling access to new types of file systems through app extensions compatible with Mac App Store distribution.

## Requirements

- **macOS 15.4 or later** (Sequoia)
- Go 1.20+
- [purego](https://github.com/ebitengine/purego) for cgo-free framework loading

## Installation

```bash
go get github.com/tmc/appledocs/generated/frameworks/fskit
```

## Available Functions

This package currently exposes 3 error conversion functions from FSKit:

- `fs_errorForCocoaError` - Creates an error object for Cocoa error codes
- `fs_errorForMachError` - Creates an error object for Mach error codes
- `fs_errorForPOSIXError` - Creates an error object for POSIX error codes

## Usage

```go
package main

import (
    _ "github.com/tmc/appledocs/generated/frameworks/fskit"
)

func main() {
    // Framework is automatically loaded via init()
    // You can now use FSKit functions
}
```

## Generated Code

These bindings are automatically generated from Apple's official documentation using the `generate-framework-bindings` tool. The generation process:

1. Parses Apple's JSON documentation for FSKit
2. Extracts function signatures and availability information
3. Generates type-safe Go bindings with purego
4. Includes version constraints (macOS 15.4+)

## Platform Availability

All FSKit APIs require:
- **macOS 15.4+** (Beta as of generation)

## Framework Structure

The generated package includes:

- `doc.go` - Package documentation with version info
- `types.gen.go` - Type definitions (placeholder for future types)
- `loader.gen.go` - Framework loading via purego
- `functions.gen.go` - Function declarations with availability info

## See Also

- [FSKit Documentation](https://developer.apple.com/documentation/fskit)
- [DarwinKit](https://github.com/progrium/darwinkit) - Comprehensive macOS framework bindings
- [purego](https://github.com/ebitengine/purego) - CGo-free syscall interface

## License

Generated code is provided as-is for interfacing with Apple's frameworks. Please refer to Apple's SDK license for framework usage terms.
