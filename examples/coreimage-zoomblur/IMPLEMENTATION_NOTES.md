# CoreImage ZoomBlur Example - Implementation Notes

## Summary

This example demonstrates using the generated CoreImage bindings to apply a CIZoomBlur filter effect to an image. It showcases both the current API and what an ideal API would look like with protocol wrapper types.

## Files Created

- `main.go` - Complete working example with extensive comments
- `go.mod` - Module definition with local replacement
- `README.md` - User-facing documentation
- `IMPLEMENTATION_NOTES.md` - This file

## Current API Challenges

### 1. Type Safety Issues

**Problem**: Many APIs return `unsafe.Pointer` instead of typed objects
```go
// Returns unsafe.Pointer instead of CIImage
inputImage := coreimage.ImageClass.ImageWithContentsOfURL(unsafe.Pointer(inputURL.ID))

// Returns unsafe.Pointer instead of Filter
filterPtr := coreimage.FilterClass.ZoomBlurFilter()
```

**Impact**: Requires manual type conversion and provides no compile-time safety

### 2. Key-Value Coding Verbosity

**Problem**: Setting filter parameters requires KVC with string keys
```go
filter.SetValueForKey(objc.ID(inputImage), "inputImage")
```

**Impact**:
- No type safety
- Runtime errors for typos
- Difficult to discover available parameters
- Verbose compared to typed setters

### 3. Value Wrapping

**Problem**: Go primitives must be wrapped in Objective-C types
```go
amount := 50.0
amountValue := foundation.NewNumberWithDouble(unsafe.Pointer(&amount))
filter.SetValueForKey(amountValue.ID, "inputAmount")
```

**Impact**: Tedious and error-prone

### 4. Manual Object Creation

**Problem**: No high-level wrappers for common types like CIVector
```go
vectorClass := objc.GetClass("CIVector")
sel := objc.RegisterName("vectorWithX:Y:")
centerVector := objc.Send[objc.ID](objc.ID(vectorClass), sel, centerX, centerY)
filter.SetValueForKey(centerVector, "inputCenter")
```

**Impact**: Requires deep Objective-C knowledge

## Ideal API Design

### Type-Safe Filter Creation
```go
// Instead of:
filterPtr := coreimage.FilterClass.ZoomBlurFilter()
filter := coreimage.FilterFrom(filterPtr)

// Should be:
filter := coreimage.NewZoomBlurFilter()
```

### Type-Safe Parameter Setting
```go
// Instead of:
filter.SetValueForKey(objc.ID(inputImage), "inputImage")
amount := 50.0
amountValue := foundation.NewNumberWithDouble(unsafe.Pointer(&amount))
filter.SetValueForKey(amountValue.ID, "inputAmount")

// Should be:
filter.SetInputImage(inputImage)
filter.SetAmount(50.0)
```

### Helper Types
```go
// Instead of:
vectorClass := objc.GetClass("CIVector")
sel := objc.RegisterName("vectorWithX:Y:")
centerVector := objc.Send[objc.ID](objc.ID(vectorClass), sel, centerX, centerY)

// Should be:
centerVector := coreimage.NewVector(centerX, centerY)
filter.SetCenter(centerVector)
```

## Implementation Roadmap

To achieve the ideal API, we need to:

### 1. Protocol Wrapper Generation

Parse CoreImage protocol definitions and generate wrapper types:

```go
// From CIZoomBlur protocol, generate:
type ZoomBlurFilter struct {
    Filter // Embed base Filter
}

func NewZoomBlurFilter() ZoomBlurFilter {
    ptr := FilterClass.ZoomBlurFilter()
    return ZoomBlurFilter{FilterFrom(ptr)}
}
```

### 2. Property Setter Generation

For each protocol property, generate type-safe setters:

```go
func (f ZoomBlurFilter) SetInputImage(image Image) {
    f.SetValueForKey(image.ID, "inputImage")
}

func (f ZoomBlurFilter) SetAmount(amount float64) {
    num := foundation.NewNumberWithDouble(unsafe.Pointer(&amount))
    f.SetValueForKey(num.ID, "inputAmount")
}

func (f ZoomBlurFilter) SetCenter(center Vector) {
    f.SetValueForKey(center.ID, "inputCenter")
}
```

### 3. Property Getter Generation

```go
func (f ZoomBlurFilter) Amount() float64 {
    val := f.ValueForKey("inputAmount")
    num := foundation.NumberFrom(unsafe.Pointer(val))
    return num.DoubleValue()
}
```

### 4. Helper Type Wrappers

```go
// Vector wrapper
type Vector struct {
    objectivec.Object
}

func NewVector(x, y float64) Vector {
    class := objc.GetClass("CIVector")
    sel := objc.RegisterName("vectorWithX:Y:")
    id := objc.Send[objc.ID](objc.ID(class), sel, x, y)
    return Vector{objectivec.Object{id}}
}
```

### 5. Template Updates

Update `templates.txtar` to include protocol wrapper generation alongside existing class generation.

## Build and Test

```bash
# Build
cd ~/go/src/github.com/tmc/appledocs/examples/coreimage-zoomblur
go build

# Test (requires an input image)
./coreimage-zoomblur input.jpg output.png
```

## Key Learnings

1. **KVC is fundamental**: All filter parameter setting uses Key-Value Coding
2. **Type wrapping is pervasive**: Go primitives need Objective-C wrappers
3. **Protocols define filter APIs**: Each filter type has a protocol with properties
4. **CIContext handles rendering**: Central object for all rendering operations
5. **File I/O through CIContext**: Simpler than manual CGImage handling

## Comparison with DarwinKit

DarwinKit achieves the ideal API through:
- Protocol-based code generation
- Type-safe wrapper structs
- Property mapping to Go types
- Helper type libraries

Our approach should be similar but integrated with our existing pipeline.

## Next Steps

1. Add protocol wrapper generation to `cmd/generate-framework-bindings`
2. Create helper library for common types (CIVector, CIColor, etc.)
3. Generate type-safe wrappers for all CoreImage filters
4. Update templates to include protocol wrappers
5. Add property override system for better type mapping

## References

- [CIFilter Documentation](https://developer.apple.com/documentation/coreimage/cifilter)
- [CIZoomBlur Protocol](https://developer.apple.com/documentation/coreimage/cizoomblur)
- [DarwinKit CoreImage](https://github.com/progrium/darwinkit/tree/main/macos/coreimage)
- [Key-Value Coding Guide](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/)
