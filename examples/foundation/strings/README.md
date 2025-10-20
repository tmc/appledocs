# Foundation NSString Examples

This example demonstrates working with NSString objects in Go using Foundation bindings.

## What it demonstrates

- Creating NSString objects
- Converting between NSString and Go strings
- String length queries
- String comparison
- Case conversion (uppercase, lowercase, capitalized)
- Substring searching
- Prefix and suffix checking
- Common string patterns

## Running the example

```bash
go run main.go
```

## Key Concepts

### NSString

NSString is Foundation's immutable string class. Key characteristics:
- Immutable (cannot be modified after creation)
- Unicode-aware (handles international characters)
- Reference-counted (automatic memory management)
- Thread-safe (can be shared across threads)

### Creating NSStrings

```go
// Empty string
str := foundation.NewString()

// From Go string (using helper)
str := goStringToNS("Hello, World!")
```

### String Conversions

Converting between NSString and Go string requires the Objective-C bridge:

```go
// NSString → Go string
func nsStringToGo(nsStr foundation.String) string {
    cStrSel := objc.RegisterName("UTF8String")
    cStr := nsStr.ID.Send(cStrSel)
    // Convert C string bytes to Go string
    ...
}

// Go string → NSString
func goStringToNS(s string) foundation.String {
    cStr := append([]byte(s), 0) // null-terminated
    sel := objc.RegisterName("stringWithUTF8String:")
    class := objc.GetClass("NSString")
    result := objc.ID(class).Send(sel, unsafe.Pointer(&cStr[0]))
    return foundation.StringFrom(unsafe.Pointer(result))
}
```

## Common Operations

### Length

```go
lengthSel := objc.RegisterName("length")
length := str.ID.Send(lengthSel)
```

### Comparison

```go
compareSel := objc.RegisterName("compare:")
result := str1.ID.Send(compareSel, str2.ID)
// result: negative = less than, 0 = equal, positive = greater than
```

### Case Conversion

```go
uppercaseSel := objc.RegisterName("uppercaseString")
uppercase := str.ID.Send(uppercaseSel)

lowercaseSel := objc.RegisterName("lowercaseString")
lowercase := str.ID.Send(lowercaseSel)

capitalizedSel := objc.RegisterName("capitalizedString")
capitalized := str.ID.Send(capitalizedSel)
```

### Substring Search

```go
containsSel := objc.RegisterName("containsString:")
contains := str.ID.Send(containsSel, searchStr.ID)
// contains != 0 means found
```

### Prefix/Suffix

```go
hasPrefix := objc.RegisterName("hasPrefix:")
hasSuffix := objc.RegisterName("hasSuffix:")

startsWithDoc := str.ID.Send(hasPrefix, prefix.ID) != 0
endsWithPdf := str.ID.Send(hasSuffix, suffix.ID) != 0
```

## Use Cases

### File Path Operations

```go
path := goStringToNS("/Users/username/Documents/file.txt")
// Check extension, directory, etc.
```

### Bundle Identifiers

```go
bundleID := goStringToNS("com.example.myapp")
// Validate format, extract components
```

### User Input Validation

```go
email := goStringToNS(userInput)
// Check format, length, etc.
```

### Localization

NSString supports localization through NSLocalizedString and related functions.

## NSString vs Go strings

| Feature | NSString | Go string |
|---------|----------|-----------|
| Mutability | Immutable | Immutable |
| Encoding | UTF-16 internally | UTF-8 |
| Thread-safe | Yes | Yes |
| Memory | Reference counted | Value type |
| Methods | Rich ObjC API | Simple, functional |
| Performance | Slower (ObjC calls) | Faster (native) |

**When to use NSString:**
- Interfacing with Objective-C/Foundation APIs
- Need NSString-specific methods
- Working with localization
- Interop with AppKit/UIKit

**When to use Go strings:**
- Pure Go code
- Performance-critical paths
- Simple string operations
- No Foundation dependency needed

## Best Practices

1. **Minimize conversions**: Convert once at boundaries
2. **Cache selectors**: Register selectors once, reuse
3. **Use Go strings**: When possible, for better performance
4. **Helper functions**: Create reusable conversion helpers
5. **Error handling**: Check for nil/zero IDs

## Advanced Features

NSString also supports:
- Regular expressions (via NSRegularExpression)
- Locale-aware comparison
- Path manipulation
- URL encoding/decoding
- Character set operations
- Line and paragraph iteration

## References

- [NSString Documentation](https://developer.apple.com/documentation/foundation/nsstring)
- [String Programming Guide](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/introStrings.html)
- [Unicode in NSString](https://developer.apple.com/documentation/foundation/nsstring/1413865-length)
