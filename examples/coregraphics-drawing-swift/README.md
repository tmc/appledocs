# CoreGraphics Drawing via Swift

EXPERIMENTAL: Call CoreGraphics from Go through Swift (not Objective-C) using purego.

## Overview

This example demonstrates:
1. Swift wrapping CoreGraphics C APIs
2. Exporting Swift functions with `@_cdecl` for Go
3. Calling Swift→CG from Go via purego (no cgo!)
4. Comparing with existing ObjC purego bindings

## Why Swift instead of ObjC?

- **Direct CG API access**: CoreGraphics is C, Swift can call it directly
- **Modern syntax**: Swift has better ergonomics than ObjC
- **Value types**: Swift structs can wrap CG types efficiently
- **No runtime overhead**: With `@_cdecl`, it's just function calls

## Architecture

```
Go (purego) → Swift dylib (@_cdecl) → CoreGraphics (C API)
```

Compare with ObjC approach:
```
Go (purego) → ObjC runtime → ObjC wrapper → CoreGraphics (C API)
```

## Shared Library

This example uses the shared Swift CoreGraphics wrapper located at:
```
../../generated/swift/frameworks/coregraphics/
```

This allows multiple examples to reuse the same Swift→CG bindings without duplication.

## Building

```bash
make           # Build Swift library
make test      # Run Go tests
make example   # Run example program
```

## Files

- `coregraphics_swift.swift` - Swift wrapper for CG drawing
- `main.go` - Go program calling Swift CG functions
- `test.go` - Tests comparing output
