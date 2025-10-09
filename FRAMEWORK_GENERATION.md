# Framework Binding Generation

This document describes the automated framework binding generation system for macOS frameworks.

## Overview

The `appledocs` tool generates Go bindings for macOS frameworks from Apple's official documentation. These bindings use [purego](https://github.com/ebitengine/purego) to call C functions without requiring cgo.

## Key Features

### Version-Aware Bindings

Generated bindings include comprehensive platform availability information:

```go
// CGBitmapInfoMake creates bitmap information from component values.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+
```

### Package Documentation

Each framework package includes:
- Minimum macOS version requirement
- Framework path
- Package-level documentation

```go
// Package coregraphics provides Go bindings for the CoreGraphics framework.
//
// Minimum macOS version: 10.0
// Framework path: /System/Library/Frameworks/CoreGraphics.framework/CoreGraphics
package coregraphics

const MinMacOSVersion = "10.0"
const FrameworkPath = "/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics"
```

## Generation Scripts

### Priority Frameworks

Generate bindings for commonly-used frameworks:

```bash
./scripts/generate-priority-frameworks.sh
```

This includes 29 high-priority frameworks:
- Core: Foundation, CoreFoundation, CoreGraphics, AppKit
- Graphics: QuartzCore, Metal, MetalKit, SpriteKit, SceneKit
- Media: AVFoundation, CoreAudio, CoreMedia, CoreVideo, ImageIO
- System: Security, SystemConfiguration, IOKit, DiskArbitration
- Modern: Combine, CoreML, CreateML, NaturalLanguage, Vision, CoreLocation
- File System: FSKit
- Networking: Network, NetworkExtension
- Developer: XPC, OSLog

### All Frameworks

Generate bindings for all 270+ public frameworks:

```bash
./scripts/generate-all-frameworks.sh
```

Both scripts:
- Skip already-generated frameworks
- Use 5-minute timeout per framework
- Provide progress tracking
- Generate summary statistics

## Generated Package Structure

```
generated/frameworks/
└── coregraphics/
    ├── doc.go            # Package documentation with version requirements
    ├── types.gen.go      # Type definitions
    ├── functions.gen.go  # Function declarations with availability
    └── loader.gen.go     # Framework loading utilities
```

## Implementation Details

### Version Comparison

The generator uses semantic version comparison:
- `10.10` > `10.9` (not lexicographic)
- Handles major.minor versioning
- Determines minimum framework version across all APIs

### Platform Support

Automatically handles all Apple platforms:
- macOS
- iOS
- iPadOS
- tvOS
- watchOS
- visionOS
- Mac Catalyst

New platforms are automatically supported without code changes due to map-based platform tracking.

### Caching

The generator uses a local cache (`.cache/`) for Apple documentation:
- 100% cache hit rate for repeated generation
- Significant performance improvement
- Shared across all framework generations

## Usage Example

```go
package main

import (
    "github.com/tmc/appledocs/generated/frameworks/coregraphics"
)

func main() {
    // Generated bindings include all version information
    // OS handles symbol resolution at runtime

    // Use framework APIs through generated bindings
    // ...
}
```

## Design Principles

Following Russ Cox-style Go design:

1. **Simplicity**: Single package per framework, no build tags
2. **Data-driven**: Maps instead of fixed struct fields
3. **Testability**: Comprehensive test coverage for version handling
4. **Flexibility**: Automatic support for new platforms
5. **Documentation**: Clear godoc comments with availability info

## Testing

Run generator tests:

```bash
cd cmd/generate-framework-bindings
go test -v
```

Tests cover:
- Version parsing (semantic comparison)
- Platform availability extraction
- Minimum version calculation
- Edge cases (deprecated, beta, unavailable APIs)

## Performance

Generation performance (with cache):
- ~15-20 files/second processing rate
- 100% cache hit rate (after initial download)
- 5-minute timeout per framework
- Parallel processing for independent frameworks possible

## Future Enhancements

Potential improvements:
- Parallel framework generation
- Incremental updates (only regenerate changed APIs)
- Version-specific type definitions
- API evolution tracking
- Compatibility matrix generation
- Migration guide generation

## References

- [purego](https://github.com/ebitengine/purego) - Pure Go FFI library
- [Apple Documentation](https://developer.apple.com/documentation/)
- [VERSION_STRATEGY.md](VERSION_STRATEGY.md) - Multi-version API design
- [ROADMAP.md](ROADMAP.md) - Project roadmap
