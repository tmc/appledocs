# photosframework

Swift-like Go bindings for Photos framework extensions.

## Overview

This package provides idiomatic Go APIs for Swift extensions to the Photos framework that are not available in the Objective-C API.

## Features

- ✅ Type-safe Go wrappers around Swift Photos extensions
- ✅ Automatic memory management with finalizers
- ✅ Iterator support for persistent changes
- ✅ No cgo required (uses purego)
- ✅ Swift-like API design

## Installation

```bash
go get github.com/tmc/appledocs/photosframework
```

## Prerequisites

The `libPhotosSwift.dylib` library must be in your library path or current directory.

### Building the Swift Library

```bash
cd ../generated/swift/frameworks/photos
swift build -c release
```

### Running Tests

The library must be in your current directory or library path:

```bash
# Copy library to test directory
cp ../generated/swift/frameworks/photos/.build/release/libPhotosSwift.dylib .

# Run tests
GOWORK=off go test -v
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/tmc/appledocs/photosframework"
)

func main() {
    // Test the library is loaded
    photosframework.TestHello()

    // Get shared photo library
    library := photosframework.SharedPhotoLibrary()
    defer library.Release()

    fmt.Println("Photo library:", library)
}
```

### Working with Persistent Changes

```go
// Get persistent change fetch result (from elsewhere)
var fetchResult *photosframework.PersistentChangeFetchResult

// Create iterator
iter := fetchResult.Iterator()
defer iter.Release()

// Iterate through changes
for {
    change := iter.Next()
    if change == nil {
        break
    }
    defer change.Release()

    // Process change
    fmt.Println("Got change:", change)
}
```

### Removing Assets from Project

```go
// Get project change request and assets (from elsewhere)
var request *photosframework.ProjectChangeRequest
var assets *photosframework.FetchResult

// Remove assets using Swift extension method
request.RemoveAssets(assets)
```

## API Design

### Memory Management

All types implement automatic memory management:

- Pointers are automatically released when garbage collected (via `runtime.SetFinalizer`)
- Manual `Release()` methods available for explicit cleanup
- Safe to call `Release()` multiple times

### Type Safety

Each Photos type is wrapped in a Go struct:

- `PhotoLibrary` - PHPhotoLibrary
- `FetchResult` - PHFetchResult<PHAsset>
- `ProjectChangeRequest` - PHProjectChangeRequest
- `PersistentChangeFetchResult` - PHPersistentChangeFetchResult
- `PersistentChangeIterator` - PHPersistentChangeFetchResult.Iterator
- `PersistentChange` - PHPersistentChange

### Swift Extensions Exposed

These methods are Swift extensions not available in Objective-C:

1. **PHProjectChangeRequest.removeAssets(_:)**
   ```go
   request.RemoveAssets(assets)
   ```

2. **PHPersistentChangeFetchResult.makeIterator()**
   ```go
   iter := fetchResult.Iterator()
   ```

3. **PHPersistentChangeFetchResult.Iterator.next()**
   ```go
   change := iter.Next()
   ```

## Architecture

```
Go Application
    ↓
photosframework (this package)
    ↓ purego FFI
libPhotosSwift.dylib
    ↓ @_cdecl exports
Swift Extensions (PhotosSwift.swift)
    ↓
Photos Framework
```

## Compared to darwinkit

This package complements `github.com/progrium/darwinkit/macos/photos`:

| Feature | darwinkit | photosframework |
|---------|-----------|-----------------|
| Objective-C APIs | ✅ Complete | ❌ Not included |
| Swift extensions | ❌ Not available | ✅ This package |
| Approach | objc runtime | Swift wrappers |
| Memory | Manual | Automatic |

**Use together:**
```go
import (
    "github.com/progrium/darwinkit/macos/photos"  // ObjC APIs
    pf "github.com/tmc/appledocs/photosframework"  // Swift extensions
)
```

## Generated From

This package was automatically generated from:

- **Source:** Photos.swiftinterface
- **Parser:** tools/swift-interface-parser
- **Wrapper:** generated/swift/frameworks/photos/PhotosSwift.swift
- **Generator:** Code generation tools

## Performance

- ✅ No cgo overhead
- ✅ Direct C function calls via purego
- ✅ Native Swift performance
- ✅ Automatic memory management with minimal overhead

## Contributing

To add more Photos APIs:

1. Parse Photos.swiftinterface with `swift-interface-parser`
2. Add @_cdecl wrappers to PhotosSwift.swift
3. Regenerate this package
4. Add Go wrapper methods
5. Update tests

## License

Same as parent project.
