# Framework Bindings Generator

Generate Go bindings for macOS/iOS frameworks from Apple's documentation JSON files.

## Overview

This tool generates Go bindings by parsing Apple's official documentation JSON files and creating type-safe Go interfaces for C frameworks.

## Features

- **Framework-by-framework generation** - Generate bindings for any Apple framework
- **Multiple output files** - Organized `.gen.go` files for types, loaders, and functions
- **Type mapping** - Automatic C to Go type conversion
- **Documentation extraction** - Preserves function signatures and comments
- **Multiple styles** - Support for purego, darwinkit, and simple binding styles

## Usage

```bash
go run cmd/generate-framework-bindings/main.go \
  -framework CoreGraphics \
  -input output/tutorials/data/documentation \
  -output generated/frameworks \
  -style purego
```

### Flags

- `-framework` - Framework name (e.g., CoreGraphics, Foundation, AppKit)
- `-input` - Input directory containing Apple documentation JSON files
- `-output` - Output directory for generated bindings
- `-style` - Binding style: `purego`, `darwinkit`, or `simple`

## Generated Files

The generator creates three `.gen.go` files per framework:

### 1. `types.gen.go`
Type definitions for the framework:
```go
// Fundamental types
type CGFloat float64

// Opaque reference types
type CGContextRef unsafe.Pointer
type CGColorRef unsafe.Pointer

// Geometric types
type CGPoint struct {
    X, Y CGFloat
}
```

### 2. `loader.gen.go`
Framework loading and initialization:
```go
var lib uintptr

func init() {
    var err error
    lib, err = purego.Dlopen("/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
    if err != nil {
        panic(err)
    }
}
```

### 3. `functions.gen.go`
Function declarations discovered from documentation:
```go
// CGContextSetRGBFillColor(c CGContextRef, red, green, blue, alpha CGFloat)
// CGContextStrokePath(c CGContextRef)
// CGContextFillRect(c CGContextRef, rect CGRect)
```

## Examples

### Generate CoreGraphics Bindings

```bash
go run cmd/generate-framework-bindings/main.go -framework CoreGraphics
```

Output:
```
generated/frameworks/coregraphics/
├── types.gen.go       (794 bytes)
├── loader.gen.go      (388 bytes)
└── functions.gen.go   (2.4 KB)
```

### Generate Foundation Bindings

```bash
go run cmd/generate-framework-bindings/main.go -framework Foundation
```

### Generate Multiple Frameworks

```bash
for fw in CoreGraphics CoreFoundation Foundation AppKit; do
  go run cmd/generate-framework-bindings/main.go -framework $fw
done
```

## Supported Frameworks

The generator can process any framework in Apple's documentation:

- **Graphics**: CoreGraphics, CoreImage, CoreAnimation, Metal
- **UI**: AppKit, UIKit, SwiftUI
- **Foundation**: Foundation, CoreFoundation
- **Media**: AVFoundation, CoreAudio, CoreVideo
- **And many more...**

## Output Styles

### Purego Style (Default)
Generates bindings for use with `github.com/ebitengine/purego`:
- Type aliases for opaque pointers
- Direct function bindings
- Minimal overhead

### Darwinkit Style
Object-oriented wrappers compatible with darwinkit:
- Struct wrappers with methods
- Type-safe interfaces
- Helper functions

### Simple Style
Raw bindings with maximum control:
- Direct C function mappings
- Explicit type conversions
- Low-level access

## Integration

To use generated bindings in your project:

1. **Generate bindings**:
   ```bash
   go run cmd/generate-framework-bindings/main.go -framework CoreGraphics
   ```

2. **Copy to your project**:
   ```bash
   cp -r generated/frameworks/coregraphics your-project/pkg/
   ```

3. **Import and use**:
   ```go
   import "your-project/pkg/coregraphics"

   // Use the lib variable and types from generated files
   ```

4. **Add function bindings** (in your own code):
   ```go
   var CGContextSetRGBFillColor func(c coregraphics.CGContextRef, r, g, b, a coregraphics.CGFloat)

   func init() {
       purego.RegisterLibFunc(&CGContextSetRGBFillColor, coregraphics.lib, "CGContextSetRGBFillColor")
   }
   ```

## Architecture

```
Input (Apple Docs JSON)
         ↓
    Parser/Analyzer
         ↓
    Type System
         ↓
  Template Generator
         ↓
 Output (.gen.go files)
```

### Parsing Process

1. **Scan** framework directory for JSON files
2. **Parse** JSON to extract function declarations
3. **Filter** Objective-C variants for C functions
4. **Map** C types to Go types
5. **Generate** organized Go source files

### Type Mapping

Common C to Go type mappings:

| C Type | Go Type |
|--------|---------|
| `CGFloat` | `float64` |
| `CGContextRef` | `unsafe.Pointer` |
| `size_t` | `uint` |
| `bool` | `bool` |
| `void *` | `unsafe.Pointer` |

## Limitations

Current limitations and future improvements:

- **Function signatures** - Some complex signatures need manual cleanup
- **Block types** - Objective-C blocks not fully supported
- **Variadic functions** - Not currently handled
- **Constants/Enums** - Not yet extracted from documentation
- **Structs** - Only common types like CGRect generated

## Contributing

To extend the generator:

1. Add new type mappings in `mapCTypeToGo()`
2. Add framework-specific types in `generateTypesFile()`
3. Enhance parser for complex signatures in `parseDeclaration()`
4. Add new output styles in `generate*Bindings()`

## See Also

- [Main Documentation](../../README.md)
- [Code Generation v2](../generate-objc-bindings-v2/README.md)
- [Example: CoreGraphics Drawing](../../examples/coregraphics-drawing/)
- [Apple Documentation](https://developer.apple.com/documentation/)

## License

Same as parent project.
