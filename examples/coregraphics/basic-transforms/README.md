# CoreGraphics Affine Transform Examples

This example demonstrates how to use CoreGraphics affine transforms in Go.

## What it demonstrates

- Creating identity, translation, scale, and rotation transforms
- Combining multiple transforms using CGAffineTransformConcat
- Inverting transforms
- Chaining transforms together
- Applying transforms to points

## Running the example

```bash
go run main.go
```

## Key Concepts

### Affine Transforms

Affine transforms are represented by a 3x3 matrix that can perform:
- Translation (moving)
- Scaling (resizing)
- Rotation
- Skewing/shearing

The transform matrix has 6 components: a, b, c, d, tx, ty

```
| a  b  0 |
| c  d  0 |
| tx ty 1 |
```

### Common Operations

- **Translation**: Moves objects by a specified distance
- **Scale**: Resizes objects by a specified factor
- **Rotation**: Rotates objects by a specified angle (in radians)
- **Concat**: Combines two transforms (order matters!)
- **Invert**: Creates the inverse of a transform

### Transform Order

When concatenating transforms, order matters:
- `scale → translate` produces different results than `translate → scale`
- Transforms are typically applied right-to-left when concatenated

## References

- [Apple CoreGraphics Documentation](https://developer.apple.com/documentation/coregraphics)
- [Affine Transformations](https://developer.apple.com/documentation/coregraphics/cgaffinetransform)
