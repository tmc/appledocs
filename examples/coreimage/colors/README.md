# CoreImage Color Examples

This example demonstrates how to create and work with CIColor objects in Go.

## What it demonstrates

- Creating colors using RGB values
- Creating colors with alpha (RGBA)
- Common predefined colors
- Converting web colors (hex) to CIColor
- Creating colors from string representations
- Proper color value ranges (0.0 to 1.0)

## Running the example

```bash
go run main.go
```

## Key Concepts

### CIColor

CIColor represents colors in Core Image's color space. Unlike system colors (NSColor/UIColor), CIColor objects are:
- Device-independent
- Immutable
- Optimized for image processing
- Used as inputs to Core Image filters

### Color Value Ranges

CoreImage uses floating-point values from 0.0 to 1.0 for color components:
- `0.0` = minimum intensity (black for RGB)
- `1.0` = maximum intensity (white for RGB, fully opaque for alpha)

### Converting Hex Colors

Web colors use 0-255 ranges, so divide by 255.0:

```go
// #FF6347 (Tomato)
color := coreimage.NewColorWithRedGreenBlue(
    0xFF / 255.0,  // Red: 255/255 = 1.0
    0x63 / 255.0,  // Green: 99/255 = 0.388
    0x47 / 255.0,  // Blue: 71/255 = 0.278
)
```

### Alpha Channel

The alpha channel controls transparency:
- `1.0` = fully opaque
- `0.5` = 50% transparent
- `0.0` = fully transparent

## Common Use Cases

CIColor objects are primarily used with Core Image filters:

- **Color Filters**: Adjust hue, saturation, brightness
- **Blend Modes**: Combine images with specific colors
- **Tinting**: Apply color overlays to images
- **Color Generation**: Create solid color images
- **Color Correction**: White balance, temperature adjustments

## Color Spaces

CIColor supports different color spaces:
- RGB (default)
- sRGB
- Linear RGB
- Device-specific color spaces

Use `NewColorWithRedGreenBlueColorSpace` to specify a custom color space.

## Memory Management

CIColor objects are Objective-C objects that conform to ARC (Automatic Reference Counting):
- No manual retain/release needed in most cases
- Go's garbage collector handles cleanup
- Use autorelease pools for intensive color creation loops

## References

- [CIColor Documentation](https://developer.apple.com/documentation/coreimage/cicolor)
- [Core Image Programming Guide](https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_intro/ci_intro.html)
- [Color Management in Core Image](https://developer.apple.com/documentation/coreimage/cicolor/color_space_objects)
