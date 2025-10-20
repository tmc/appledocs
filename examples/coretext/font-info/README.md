# CoreText Font Information Example

This example demonstrates how to work with fonts using CoreText in Go.

## What it demonstrates

- Creating CTFont objects from font names
- Querying font properties (size, ascent, descent)
- Working with system fonts
- Converting between CoreFoundation and Go strings
- Proper memory management with CFRelease

## Running the example

```bash
go run main.go
```

## Key Concepts

### CoreText

CoreText is a low-level text layout and font handling framework on macOS/iOS. It provides:
- Font creation and management
- Text layout and rendering
- Glyph handling
- Typography features

### CTFont

CTFont represents a font with a specific size and attributes. It's immutable and thread-safe.

```go
// Create a font by name
font := coretext.CTFontCreateWithName(fontName, size, nil)
defer corefoundation.CFRelease(font)
```

### Font Metrics

Key font metrics:
- **Size**: The point size of the font
- **Ascent**: Distance from baseline to top of tallest glyph
- **Descent**: Distance from baseline to bottom (negative)
- **Line Height**: Total height (ascent + descent)
- **Leading**: Additional space between lines

### CoreFoundation Integration

CoreText uses CoreFoundation types:
- `CFStringRef` for strings
- `CFRelease` for memory management
- `CFStringGetCStringPtr` for string conversion

Always release CF objects when done:

```go
cfStr := corefoundation.CFStringCreateWithCString(nil, "Hello", encoding)
defer corefoundation.CFRelease(cfStr)
```

## Common System Fonts on macOS

- **.AppleSystemUIFont**: Default system UI font (San Francisco)
- **Helvetica**: Classic sans-serif
- **Helvetica Neue**: Modern sans-serif
- **Times**: Classic serif font
- **Courier**: Classic monospace
- **Monaco**: Monospace for code (classic)
- **Menlo**: Modern monospace for code
- **SF Mono**: Apple's monospace font

## Font Naming

Fonts can be referenced by:
1. **Family name**: "Helvetica"
2. **Full name**: "Helvetica Bold"
3. **PostScript name**: "Helvetica-Bold"
4. **System name**: ".AppleSystemUIFont"

## Memory Management

CoreText/CoreFoundation objects use manual reference counting:

```go
// Create increases reference count
font := coretext.CTFontCreateWithName(...)

// Must release when done
defer corefoundation.CFRelease(font)

// Copy functions also increase reference count
displayName := coretext.CTFontCopyDisplayName(font)
defer corefoundation.CFRelease(displayName)
```

## Use Cases

### Font Selection

Check if a font exists before using:

```go
font := coretext.CTFontCreateWithName(fontName, size, nil)
if font != nil {
    defer corefoundation.CFRelease(font)
    // Font is available
}
```

### Line Height Calculation

Calculate proper line spacing:

```go
ascent := coretext.CTFontGetAscent(font)
descent := coretext.CTFontGetDescent(font)
leading := coretext.CTFontGetLeading(font)
lineHeight := ascent + descent + leading
```

### Font Metrics for Layout

Use metrics for precise text positioning:

```go
// Position text baseline
y := topMargin + coretext.CTFontGetAscent(font)

// Calculate next line
y += lineHeight
```

### Custom Font Loading

Load fonts from files:

```go
// Create font descriptor from file
// Use CTFontCreateWithGraphicsFont
```

## Thread Safety

CTFont objects are thread-safe and can be:
- Shared across threads
- Cached and reused
- Created on background threads

## Best Practices

1. **Cache fonts**: Creating fonts can be expensive
2. **Release properly**: Always use defer for CFRelease
3. **Check nil**: Font creation can fail if font doesn't exist
4. **Use size carefully**: Font sizes are in points, not pixels
5. **Convert strings**: Use proper encoding (UTF-8) for CFString

## Advanced Features

CoreText also supports:
- Font collections
- Font descriptors (for font matching)
- Font variations (weight, width, slant)
- Glyph paths and rendering
- OpenType features
- Text layout and line breaking

## References

- [CTFont Documentation](https://developer.apple.com/documentation/coretext/ctfont)
- [CoreText Programming Guide](https://developer.apple.com/library/archive/documentation/StringsTextFonts/Conceptual/CoreText_Programming/Introduction/Introduction.html)
- [Font Handling](https://developer.apple.com/documentation/coretext/ctfont-q6r)
