# Swift ABI and Calling Conventions

Notes on calling Swift from Go using purego.

## Name Mangling

Swift uses name mangling similar to C++. To export C-compatible symbols, use:

```swift
@_cdecl("my_function_name")
public func myFunction() {
    // ...
}
```

This exports the function with a plain C symbol name that can be found with `dlsym`.

## Calling Convention

### Simple Cases (Works with purego)

When using `@_cdecl`, Swift functions use the C calling convention:
- **Primitives**: Int32, UInt64, Float, Double, Bool - pass directly
- **Pointers**: UnsafePointer<T>, UnsafeMutablePointer<T> - pass as Go pointers
- **C Strings**: UnsafePointer<CChar> - pass as null-terminated byte arrays

### Complex Cases (Harder with purego)

Without `@_cdecl`, Swift uses its own calling convention:
- **Value Types**: Structs passed via registers or stack based on size
- **Reference Types**: Classes use ARC (retain/release protocol)
- **Generics**: Witness tables and metadata pointers
- **Protocol Types**: Existential containers with vtables

## Memory Management

### Simple Approach
Use `@_cdecl` and manage memory manually:
- Swift allocates with `strdup()` → Go must free (but purego doesn't expose C free)
- Go allocates → Swift must not retain

### Full ARC Interop
Would require:
1. Understanding Swift's ARC conventions
2. Calling `swift_retain` and `swift_release` runtime functions
3. Managing retain counts manually from Go

## Type Mapping

### Primitives
| Swift | C | Go |
|-------|---|-----|
| Int32 | int32_t | int32 |
| UInt64 | uint64_t | uint64 |
| Float | float | float32 |
| Double | double | float64 |
| Bool | bool | bool |

### Pointers
| Swift | C | Go |
|-------|---|-----|
| UnsafePointer<T> | const T* | *T |
| UnsafeMutablePointer<T> | T* | *T |
| UnsafeRawPointer | const void* | unsafe.Pointer |

### Strings
Swift String is complex (not a C string). Options:
1. Use `UnsafePointer<CChar>` and convert in Swift
2. Return C strings with `strdup()` (caller must free)
3. Pass fixed-size buffers

## Objective-C Bridge

For complex Swift types, consider:
1. Swift class → Objective-C class (`@objc`)
2. Call via ObjC runtime using purego
3. Already have patterns for this in the binding generator

## Recommended Approach

### For Simple Libraries
Use `@_cdecl` with C-compatible types:
```swift
@_cdecl("my_simple_func")
public func mySimpleFunc(_ x: Int32) -> Int32 {
    return x * 2
}
```

### For Complex Libraries
1. Create Swift → ObjC bridge
2. Use existing ObjC binding patterns
3. Leverage `@objc` attribute on Swift classes

## Example: Bridging Swift Class

```swift
// Swift side
@objc public class Calculator: NSObject {
    @objc public func add(_ a: Int, _ b: Int) -> Int {
        return a + b
    }
}

@_cdecl("create_calculator")
public func createCalculator() -> UnsafeMutableRawPointer {
    let calc = Calculator()
    return Unmanaged.passRetained(calc).toOpaque()
}

@_cdecl("calculator_add")
public func calculatorAdd(_ ptr: UnsafeMutableRawPointer, _ a: Int32, _ b: Int32) -> Int32 {
    let calc = Unmanaged<Calculator>.fromOpaque(ptr).takeUnretainedValue()
    return Int32(calc.add(Int(a), Int(b)))
}

@_cdecl("release_calculator")
public func releaseCalculator(_ ptr: UnsafeMutableRawPointer) {
    Unmanaged<Calculator>.fromOpaque(ptr).release()
}
```

This pattern:
- Uses C calling convention
- Manages ARC explicitly with Unmanaged
- Passes opaque pointers to Go
- Works perfectly with purego

## Tools for Investigation

```bash
# View exported symbols
nm -gU libhello.dylib

# View mangled Swift symbols
nm libhello.dylib | grep swift

# Demangle Swift symbols
swift demangle <mangled_name>

# Check calling convention
otool -tV libhello.dylib
```

## References

- [Swift ABI Stability Manifesto](https://github.com/apple/swift/blob/main/docs/ABIStabilityManifesto.md)
- [Swift Calling Convention](https://github.com/apple/swift/blob/main/docs/CallingConvention.rst)
- [Name Mangling](https://github.com/apple/swift/blob/main/docs/ABI/Mangling.rst)
