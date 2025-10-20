# CoreGraphics Geometry Examples

This example demonstrates working with CoreGraphics geometry types in Go.

## What it demonstrates

- Creating and using CGPoint, CGSize, and CGRect
- Calculating rectangle properties (area, corners)
- Working with aspect ratios
- Using type aliases (Point, Size, Rect)
- Common screen sizes and their dimensions

## Running the example

```bash
go run main.go
```

## Key Concepts

### CGPoint

Represents a point in 2D coordinate space with X and Y components:

```go
point := coregraphics.CGPoint{X: 10, Y: 20}
```

### CGSize

Represents dimensions with Width and Height:

```go
size := coregraphics.CGSize{Width: 100, Height: 200}
```

### CGRect

Represents a rectangle with an origin point and size:

```go
rect := coregraphics.CGRect{
    Origin: coregraphics.CGPoint{X: 0, Y: 0},
    Size:   coregraphics.CGSize{Width: 100, Height: 100},
}
```

### Type Aliases

CoreGraphics provides convenient type aliases for compatibility:

- `Point` = `CGPoint`
- `Size` = `CGSize`
- `Rect` = `CGRect`

These can be used interchangeably.

## Coordinate System

CoreGraphics uses a coordinate system where:
- Origin (0,0) is typically at the top-left on macOS/iOS screens
- X increases to the right
- Y increases downward (screen coordinates) or upward (PDF/print coordinates)

## Common Use Cases

- Window and view positioning
- Hit testing and bounds checking
- Layout calculations
- Graphics rendering
- Screen/display calculations

## References

- [CGPoint Documentation](https://developer.apple.com/documentation/corefoundation/cgpoint)
- [CGSize Documentation](https://developer.apple.com/documentation/corefoundation/cgsize)
- [CGRect Documentation](https://developer.apple.com/documentation/corefoundation/cgrect)
