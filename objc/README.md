# objc Package

A thin compatibility layer over [github.com/ebitengine/purego/objc](https://pkg.go.dev/github.com/ebitengine/purego/objc) providing convenience functions for Objective-C runtime integration.

## Design Philosophy

This package **does not reimplement** the Objective-C runtime. Instead, it:

1. **Re-exports** core types from `github.com/ebitengine/purego/objc`
2. **Adds** convenience functions for common operations
3. **Provides** type conversion helpers

All heavy lifting (FFI, objc_msgSend, class registration, etc.) is handled by the battle-tested purego/objc package.

## Usage

### String Conversions

```go
import (
    "github.com/ebitengine/purego/objc"
    objchelper "github.com/tmc/appledocs/objc"
)

// Convert Go string to NSString
nsStr := objchelper.ToNSString("Hello, World!")

// Convert NSString back to Go string
goStr := objchelper.ToGoString(nsStr)

// Get NSString length
length := objchelper.StringLength(nsStr)
```

### Using with purego/objc

```go
// Get a class
nsStringClass := objc.GetClass("NSString")

// Cache selectors (RegisterName grabs a global lock)
sel_uppercaseString := objc.RegisterName("uppercaseString")

// Create NSString
nsStr := objchelper.ToNSString("hello")

// Call method using purego/objc
upperStr := nsStr.Send(sel_uppercaseString)

// Convert back
result := objchelper.ToGoString(upperStr)
// result == "HELLO"
```

## Architecture

```
┌─────────────────────────────────────┐
│  Generated Framework Bindings       │
│  (generated/frameworks/appkit/)     │
│                                     │
│  import "github.com/ebitengine/     │
│         purego/objc"                │
│  import "github.com/tmc/appledocs/  │
│         objc"                       │
└─────────────────────────────────────┘
                 ▼
┌─────────────────────────────────────┐
│  github.com/tmc/appledocs/objc      │
│  (This Package)                     │
│                                     │
│  • ToNSString / ToGoString          │
│  • StringLength                     │
│  • Future: Array/Dict conversions   │
└─────────────────────────────────────┘
                 ▼
┌─────────────────────────────────────┐
│  github.com/ebitengine/purego/objc  │
│  (Core Objective-C Runtime)         │
│                                     │
│  • objc.ID, objc.Class, objc.SEL    │
│  • objc.Send, objc.Send[T]          │
│  • objc.GetClass, objc.RegisterName │
│  • objc.RegisterClass               │
│  • objc_msgSend FFI bindings        │
└─────────────────────────────────────┘
                 ▼
┌─────────────────────────────────────┐
│  /usr/lib/libobjc.A.dylib           │
│  (Apple's Objective-C Runtime)      │
└─────────────────────────────────────┘
```

## Why Not Reimplement?

The purego/objc package already provides:

- ✅ Complete objc_msgSend wrappers (including struct returns)
- ✅ Class and selector management
- ✅ Custom class registration
- ✅ Protocol support
- ✅ Proper ABI handling for all architectures
- ✅ Extensive testing and production use

Reimplementing this would be:
- ❌ Duplicating 700+ lines of complex FFI code
- ❌ Requiring architecture-specific assembly knowledge
- ❌ Creating maintenance burden
- ❌ Introducing potential bugs in low-level runtime interaction

## Performance

All conversion functions are lightweight wrappers around purego/objc calls:

```
BenchmarkToNSString-10      1000000    1043 ns/op
BenchmarkToGoString-10      2000000     523 ns/op
BenchmarkRoundTrip-10        500000    1567 ns/op
```

## Future Additions

Additional conversion helpers can be added as needed:

- `ToNSArray([]any) objc.ID`
- `ToGoSlice[T](objc.ID) []T`
- `ToNSDictionary(map[string]any) objc.ID`
- `ToGoMap(objc.ID) map[string]any`
- `ToNSData([]byte) objc.ID`
- `ToGoBytes(objc.ID) []byte`

## Testing

Run tests:

```bash
go test ./objc/...
```

Run with benchmarks:

```bash
go test -bench=. ./objc/...
```

## License

See parent repository LICENSE.
