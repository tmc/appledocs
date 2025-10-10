# Swift/Go Interop Example

EXPERIMENTAL: Calling Swift code from Go using purego instead of cgo.

## Overview

This example demonstrates calling Swift functions from Go without using cgo, leveraging purego's ability to call native functions via dlsym.

## Swift ABI Considerations

Swift has a complex ABI that differs from C:
- **Name Mangling**: Swift uses name mangling (like C++)
- **Calling Convention**: Swift uses its own calling convention
- **Memory Management**: Swift uses ARC (Automatic Reference Counting)
- **Value Types**: Structs are passed by value with specific rules

## Strategy

For simple interop:
1. Use `@_cdecl` attribute in Swift to export C-compatible symbols
2. Keep signatures simple (primitives, pointers)
3. Use purego to dlopen the dylib and call functions

## Example Structure

- `hello.swift` - Simple Swift library with C exports
- `main.go` - Go program calling Swift via purego
- `Makefile` - Build instructions

## Building

```bash
# Compile Swift to dynamic library
swiftc -emit-library hello.swift -o libhello.dylib

# Run Go program
go run main.go
```

## Limitations

- No automatic bridging of Swift types
- Must use C-compatible types at the boundary
- Swift classes/protocols require Objective-C runtime
- For complex interop, consider Swift/ObjC bridge + purego ObjC calls
