# Swift vs Objective-C purego Bindings Comparison

## Approach Comparison

### Swift Approach (This Example)

```
Go (purego) → Swift dylib (@_cdecl) → CoreGraphics C API
```

**Pros:**
- ✅ **Direct CG access**: Swift calls CG C APIs directly
- ✅ **No runtime overhead**: No ObjC runtime, just function calls
- ✅ **Modern syntax**: Swift is more ergonomic than ObjC
- ✅ **Value types**: Swift structs efficiently wrap CG types
- ✅ **Type safety**: Swift's type system catches errors at compile time
- ✅ **Memory managed**: Swift ARC with explicit `Unmanaged` for control
- ✅ **Simple exports**: `@_cdecl` creates clean C symbols

**Cons:**
- ❌ **Must build dylib**: Requires Swift compiler
- ❌ **Explicit exports**: Each function needs `@_cdecl`
- ❌ **Manual ARC**: Must use `Unmanaged` for object lifetime
- ❌ **Limited to C types**: Boundary must be C-compatible

### Objective-C Approach (Existing Bindings)

```
Go (purego) → ObjC Runtime → ObjC wrapper → CoreGraphics C API
```

**Pros:**
- ✅ **No compilation**: Uses system ObjC runtime
- ✅ **Dynamic**: Can discover classes/methods at runtime
- ✅ **Rich ecosystem**: Access to all Cocoa APIs
- ✅ **Well documented**: ObjC runtime behavior is stable

**Cons:**
- ❌ **Runtime overhead**: Message dispatch, retain/release
- ❌ **Complex setup**: Need to understand ObjC runtime
- ❌ **Verbose**: ObjC syntax and patterns
- ❌ **More layers**: Go → ObjC runtime → ObjC → C API

## Performance

### Swift (This Approach)
- **Function call**: Direct C call via dlsym
- **Memory**: Explicit control with Unmanaged
- **Overhead**: Minimal (just the dylib boundary)

### Objective-C (Existing)
- **Function call**: `objc_msgSend` dispatch
- **Memory**: Automatic retain/release
- **Overhead**: Message dispatch + retain/release cycles

**Winner**: Swift for pure CG operations (no messaging overhead)

## Use Cases

### When to Use Swift Approach
1. **Pure drawing operations**: CoreGraphics, Quartz, Metal
2. **Performance critical**: No ObjC runtime overhead wanted
3. **C-like APIs**: Wrapping C frameworks
4. **Greenfield projects**: Starting fresh

### When to Use ObjC Approach
1. **Cocoa integration**: AppKit, UIKit, Foundation classes
2. **Dynamic features**: Runtime introspection needed
3. **Existing code**: Already using ObjC runtime
4. **Complex types**: Need automatic bridging

## Code Comparison

### Creating a Context

**Swift:**
```swift
@_cdecl("cg_create_bitmap_context")
public func cgCreateBitmapContext(_ width: Int32, _ height: Int32) -> UnsafeMutableRawPointer? {
    let context = CGContext(...)
    return Unmanaged.passRetained(context as AnyObject).toOpaque()
}
```

```go
var createContext func(int32, int32) uintptr
purego.RegisterLibFunc(&createContext, lib, "cg_create_bitmap_context")
ctx := createContext(400, 400)
```

**ObjC (Hypothetical):**
```go
// Get NSGraphicsContext class
class := objc.GetClass("NSGraphicsContext")

// Create context via message send
sel := objc.GetSelector("graphicsContextWithBitmapImageRep:")
ctx := objc.MsgSend(class, sel, bitmapRep)
```

### Drawing a Rectangle

**Swift:**
```go
ctx.SetFillColor(cg, 0.2, 0.4, 0.8, 1.0)
ctx.FillRect(cg, 50, 50, 100, 100)
```

**ObjC (Hypothetical):**
```go
// Create NSColor
colorClass := objc.GetClass("NSColor")
colorSel := objc.GetSelector("colorWithRed:green:blue:alpha:")
color := objc.MsgSend(colorClass, colorSel, 0.2, 0.4, 0.8, 1.0)

// Set fill color
objc.MsgSend(color, objc.GetSelector("set"))

// Create NSRect and fill
// ... more message passing
```

## Recommendation

### For CoreGraphics Drawing
**Use Swift approach** - It's:
- Simpler
- Faster
- More direct
- Easier to maintain

### For General macOS/iOS APIs
**Use ObjC approach** - You need:
- Dynamic runtime features
- Cocoa class integration
- Existing patterns

## Hybrid Approach

The best solution might be **both**:
1. Use Swift for performance-critical CG operations
2. Use ObjC runtime for Cocoa integration
3. Mix as needed in the same application

Example:
```go
// Create window with ObjC
window := cocoa.NewWindow(...)

// Draw with Swift CG
cg := NewSwiftCG()
ctx := cg.CreateContext(400, 400)
// ... fast drawing ...
```

## Conclusion

This Swift approach demonstrates that:
1. **You can bypass ObjC entirely** for C-based frameworks
2. **purego works with any dylib**, not just ObjC runtime
3. **Swift's `@_cdecl` is perfect** for Go interop
4. **Performance is excellent** - no runtime overhead

For **CoreGraphics specifically**, the Swift approach is **superior** to going through the ObjC runtime, as it eliminates unnecessary layers while providing a clean, type-safe interface.
