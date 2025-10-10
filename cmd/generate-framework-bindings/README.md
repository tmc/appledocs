# Framework Bindings Generator - Improvements

## Overview

Enhanced code generation pipeline for creating Go bindings from Apple's documentation JSON files.

## Key Improvements 

### 1. Error Handling
- `ErrorCollector` accumulates all errors with context
- Detailed error reporting by stage/file/symbol
- Strict mode option (-strict)

### 2. Enhanced Parsing
- **Blocks** - Objective-C closures
- **Methods** - Instance and class methods
- **Properties** - With attributes
- **Categories** - Extensions
- **Enums** - Constants (partial)

### 3. Better Documentation
- Full API doc extraction from JSON
- Proper GoDoc comments
- Parameter/return descriptions
- Code examples preserved

### 4. New Output Files
- `methods.gen.go` - Objective-C methods
- `properties.gen.go` - Property declarations
- `blocks.gen.go` - Block type definitions

## Files

- `main.go` - Original generator
- `main_improved.go` - Enhanced version
- `errors.go` - Error handling
- `parsers.go` - Enhanced parsers
- `docgen.go` - Documentation extraction

## Usage

```bash
# Basic
go run ./cmd/generate-framework-bindings/*.go -framework Foundation

# Verbose with error details
go run ./cmd/generate-framework-bindings/*.go -framework CoreGraphics -verbose

# Strict mode (fail on errors)
go run ./cmd/generate-framework-bindings/*.go -framework AppKit -strict

# EXPERIMENTAL: Swift interop mode (generates bindings for calling Swift via purego)
go run ./cmd/generate-framework-bindings/*.go -framework MySwiftFramework -swift-interop
```

## Swift Interop (Experimental)

The `-swift-interop` flag enables experimental Swift interop mode. This generates bindings for calling Swift code using purego instead of cgo.

**Key Requirements:**
- Swift functions must use `@_cdecl` attribute for C compatibility
- Only simple types supported (primitives, pointers, C strings)
- For complex types, use Swift → ObjC bridge

**Example:**
See `examples/swift-interop/` for a complete working example.

## Migration

To use improved generator:
```bash
cp cmd/generate-framework-bindings/main.go cmd/generate-framework-bindings/main_original.go
mv cmd/generate-framework-bindings/main_improved.go cmd/generate-framework-bindings/main.go
```

See `/tmp/generator-analysis.md` for detailed analysis.
