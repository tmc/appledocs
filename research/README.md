# AppKit Swift/Go FFI Research

This directory contains research findings for calling Swift code from Go for the appledocs binding generation project.

## Documents

### [swift-go-ffi-investigation.md](swift-go-ffi-investigation.md)
**Main research findings** - Comprehensive investigation of Swift/Go FFI mechanisms.

**Key Results:**
- ✅ Swift @_cdecl functions work perfectly with purego (~174ns overhead)
- ✅ Mangled Swift symbols are callable (experimental)
- ❌ Pure Swift frameworks (Charts) don't ship as loadable dylibs
- **Recommended**: purego + @_cdecl for production use

### [swift-wrapper-example.md](swift-wrapper-example.md)
**Practical pattern** for wrapping pure Swift frameworks with @_cdecl exports.

Shows how to:
- Create Swift wrapper libraries
- Compile and distribute wrappers
- Generate Go bindings for wrappers
- Potential for automation

## Test Code

Working examples are in `/tmp/swift-ffi-test/`:
- `SwiftLib.swift` - Test Swift library
- `test_purego.go` - purego FFI (✅ all tests pass)
- `test_cgo.go` - CGo FFI (✅ all tests pass)
- `test_mangled.go` - Mangled symbols (⚠️ experimental)
- `benchmark_purego_test.go` - Performance tests

## Quick Start

```bash
# Reproduce findings
cd /tmp/swift-ffi-test
swiftc -emit-library -o libswifttest.dylib -emit-module -module-name SwiftTest SwiftLib.swift
go run test_purego.go
go test -bench=. -benchmem
```

## Recommendation for appledocs

1. **Immediate**: Focus on Objective-C frameworks (current approach works)
2. **Phase 1**: Add support for Swift frameworks with @_cdecl exports
3. **Phase 2**: Create wrapper generator for pure Swift frameworks
4. **Phase 3**: Experimental mangled symbol support for advanced users

## Related Beads

- appledocs-394: Investigate Swift/Go FFI mechanisms (this research)
- appledocs-396: Prototype Swift binding generation

## Date

Research conducted: 2025-10-21
