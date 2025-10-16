# DarwinKit-Style Bindings End-to-End Test

This example demonstrates the complete workflow for generating and using DarwinKit-style Objective-C bindings.

## What This Demonstrates

1. **Basic objc runtime access** - Direct use of `github.com/progrium/darwinkit/objc` to call Objective-C methods
2. **Generated bindings workflow** - How to generate DarwinKit-compatible bindings from Apple documentation
3. **Type-safe Go API** - Using generated bindings that feel like native Go

## Running the Basic Test

The basic test demonstrates raw objc runtime usage:

```bash
go run main.go
```

This will:
- Load the AppKit framework using purego
- Get the NSColor class from the Objective-C runtime
- Call the `+[NSColor redColor]` class method
- Get the object's description
- Verify selector response

## Generating DarwinKit-Style Bindings

To generate bindings for AppKit classes (future work - requires method parsing):

```bash
# Generate bindings for a specific class
cd ../../
go run ./cmd/generate-framework-bindings \
  -framework AppKit \
  -variant darwinkit \
  -output ./examples/darwinkit-bindings-test/appkit

# This would generate:
# appkit/
#   doc.gen.go          - Package documentation
#   gen.go              - go:generate directive
#   color.gen.go        - NSColor bindings
#   application.gen.go  - NSApplication bindings
#   ...
```

## Expected Generated Code Structure

For NSColor, the generated code would look like:

```go
package appkit

import (
    "unsafe"
    "github.com/progrium/darwinkit/objc"
)

// The class instance for the [Color] class.
var ColorClass _ColorClass

func init() {
    ColorClass = _ColorClass{objc.GetClass("NSColor")}
}

type _ColorClass struct {
    objc.Class
}

// An interface definition for the [Color] class.
type IColor interface {
    objc.IObject
    // Methods...
}

type Color struct {
    objc.Object
}

func ColorFrom(ptr unsafe.Pointer) Color {
    return Color{Object: objc.ObjectFrom(ptr)}
}

// Class method: +[NSColor redColor]
func (cc _ColorClass) RedColor() Color {
    rv := objc.Call[Color](cc, objc.Sel("redColor"))
    return rv
}

// Convenience constructor
func Color_RedColor() Color {
    return ColorClass.RedColor()
}
```

## Architecture

The binding generation follows this flow:

1. **Parse Apple Documentation** → Extract class, method, property info from JSON
2. **Apply Templates** → Use templates_darwinkit.txtar to generate Go code
3. **Generate Files** → One `.gen.go` file per class
4. **Compile & Use** → Import and use type-safe Go API

## Next Steps

To fully demonstrate end-to-end:

1. Implement method discovery from Apple documentation
2. Generate actual class files with methods
3. Create a demo using the generated bindings
4. Compare with hand-written darwinkit bindings for quality

## Current Status

✅ Basic objc runtime integration working
✅ Template system for DarwinKit-style code
✅ Parser supports methods and properties
⏳ Need method discovery from Apple docs
⏳ Need integration test with generated code
