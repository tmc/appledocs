# CoreImage Zoom Blur - Full AppKit Demo

This example demonstrates CoreImage's ZoomBlur filter in a complete AppKit application with a visual side-by-side comparison.

## Overview

The application creates a native macOS window displaying:
- **Left side**: Original image (solid color gradient)
- **Right side**: Same image with CIZoomBlur filter applied
- **Labels**: Descriptive text explaining the effect
- **Live rendering**: Real-time CoreImage processing

The CIZoomBlur filter simulates a zooming motion blur effect, creating a radial blur that emanates from a center point.

## Features

### Visual Demonstration
- Side-by-side comparison of original and blurred images
- Bordered image views with proper scaling
- Labeled sections with descriptions
- Clean, native macOS interface

### CoreImage Integration
- Creates sample image using CIColor
- Applies CIZoomBlur filter with custom parameters
- Renders using CIContext
- Converts CIImage → CGImage → NSImage for display

### AppKit Usage
- NSApplication with proper activation policy
- NSWindow with standard controls (close, minimize, resize)
- NSImageView for image display
- NSTextField for labels
- Main thread execution (required for AppKit)

## Building

```bash
cd ~/go/src/github.com/tmc/appledocs/examples/coreimage-zoomblur
go build
```

**Note**: There's currently a syntax error in the generated `coreimage` package that needs to be fixed:
- `generated/coreimage/constants.gen.go:262` - syntax error in grouped declaration
- `generated/coreimage/typedefs.gen.go:40` - syntax error in type declaration

Once the generator is fixed to produce valid Go code, the app will compile.

## Running

```bash
./coreimage-zoomblur
```

The application will:
1. Create a 1200x700 window
2. Generate a solid color image (blue/teal gradient)
3. Apply zoom blur with amount=30.0, center at (250, 200)
4. Display both images side-by-side
5. Run the AppKit event loop

## Code Structure

### Main Components

- **main()** - Sets up NSApplication and event loop
- **createWindow()** - Creates and configures the main window
- **createContentView()** - Builds the UI layout with image views and labels
- **createImageView()** - Creates NSImageView with border styling
- **createLabel()** - Creates centered NSTextField labels
- **createSampleImage()** - Generates a test image using CoreImage
- **applyZoomBlur()** - Applies the CIZoomBlur filter
- **ciImageToNSImage()** - Converts CIImage to NSImage for display

### Current API Challenges

The code demonstrates several challenges with the current API:

1. **Manual objc.Send calls** for many AppKit operations
   ```go
   objc.Send[objc.ID](app.ID, objc.RegisterName("setActivationPolicy:"), uint(0))
   ```

2. **unsafe.Pointer conversions** throughout
   ```go
   filter.SetValueForKey(objc.ID(inputImage), "inputImage")
   ```

3. **String-based KVC** for filter parameters
   ```go
   filter.SetValueForKey(amountValue.ID, "inputAmount")
   ```

4. **Manual NSNumber wrapping**
   ```go
   amountValue := foundation.NewNumberWithDouble(unsafe.Pointer(&amount))
   ```

5. **Manual CIVector creation**
   ```go
   vectorClass := objc.GetClass("CIVector")
   sel := objc.RegisterName("vectorWithX:Y:")
   centerVector := objc.Send[objc.ID](objc.ID(vectorClass), sel, centerX, centerY)
   ```

### Ideal API (Future Goal)

With protocol wrappers, the code would be much cleaner:

```go
// Instead of current verbose approach:
filterPtr := coreimage.FilterClass.ZoomBlurFilter()
filter := coreimage.FilterFrom(filterPtr)
filter.SetValueForKey(objc.ID(inputImage), "inputImage")
amount := 30.0
amountValue := foundation.NewNumberWithDouble(unsafe.Pointer(&amount))
filter.SetValueForKey(amountValue.ID, "inputAmount")

// Would become:
filter := coreimage.NewZoomBlurFilter()
filter.SetInputImage(inputImage)
filter.SetAmount(30.0)
filter.SetCenter(coreimage.NewVector(250, 200))
```

## Filter Parameters

The CIZoomBlur filter accepts:

- **inputImage** (CIImage): The image to blur
- **inputAmount** (float): The zoom blur radius
  - Default: 20.0
  - Range: 0.0 to infinite
  - Example uses: 30.0
- **inputCenter** (CIVector): The center point of the zoom effect
  - Default: Image center
  - Format: (x, y) coordinates
  - Example uses: (250, 200)

## Implementation Notes

### Thread Safety
AppKit requires all UI operations to run on the main thread:
```go
runtime.LockOSThread()
```

### Image Rendering Pipeline
1. Create CIImage (via CIColor or file)
2. Apply CIFilter (ZoomBlur)
3. Get output CIImage
4. Render to CGImage via CIContext
5. Convert to NSImage for AppKit display

### Memory Management
- Uses automatic reference counting (ARC)
- CIImage and CIFilter are immutable
- CIContext can be reused for multiple renders

## What This Demonstrates

### Current API State
- ✅ Full AppKit integration is possible
- ✅ CoreImage filters work correctly
- ✅ Type-safe where bindings exist
- ⚠️ Requires many manual objc.Send calls
- ⚠️ Lots of unsafe.Pointer usage
- ⚠️ No compile-time checking for filter parameters

### What We Need
1. **Protocol wrapper generation** for filters
2. **Helper types** (CIVector, CIColor wrappers)
3. **Better AppKit method coverage** in bindings
4. **Fix generator syntax errors** in constants and typedefs

### Comparison with DarwinKit

DarwinKit achieves cleaner API through:
- Protocol-based filter wrappers (e.g., `ZoomBlurFilter` type)
- Type-safe setter methods
- Automatic parameter conversion
- Complete AppKit method bindings

Our goal is to implement similar improvements while maintaining our generation pipeline.

## Next Steps

1. Fix syntax errors in coreimage generation
2. Add protocol wrapper generation to templates
3. Create helper type library (CIVector, CIColor, etc.)
4. Expand AppKit method coverage
5. Add property override system for better type mapping

## Related Files

- `main.go` - Complete AppKit application
- `IMPLEMENTATION_NOTES.md` - Technical details and API comparison
- `go.mod` - Module definition
- `../../generated/coreimage/` - Generated bindings (needs fix)
- `../../cmd/generate-framework-bindings/` - Code generator

## References

- [NSApplication Documentation](https://developer.apple.com/documentation/appkit/nsapplication)
- [CIFilter Documentation](https://developer.apple.com/documentation/coreimage/cifilter)
- [CIZoomBlur Documentation](https://developer.apple.com/documentation/coreimage/cizoomblur)
- [CIImage Documentation](https://developer.apple.com/documentation/coreimage/ciimage)
- [Key-Value Coding](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/)
