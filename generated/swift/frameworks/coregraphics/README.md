# CoreGraphics Swift Bindings

Reusable Swift wrapper for CoreGraphics C API with `@_cdecl` exports for Go/purego consumption.

## Purpose

This library provides a clean, Go-callable interface to CoreGraphics by:
1. Wrapping CoreGraphics C APIs in idiomatic Swift
2. Exporting functions with `@_cdecl` for C compatibility
3. Managing memory with explicit `Unmanaged` for Go control

## Building

```bash
# Build dynamic library
swift build -c release

# Output location
.build/release/libCoreGraphicsSwift.dylib
```

## Using from Go

```go
import "github.com/ebitengine/purego"

lib, _ := purego.Dlopen("./libCoreGraphicsSwift.dylib", purego.RTLD_LAZY)
var cgCreateContext func(int32, int32) uintptr
purego.RegisterLibFunc(&cgCreateContext, lib, "cg_create_bitmap_context")
```

## API Coverage

- Bitmap context creation
- Fill/stroke colors
- Rectangle drawing
- Ellipse drawing
- Path operations
- PNG export

## Examples

See `examples/coregraphics-drawing-swift/` for complete usage example.
