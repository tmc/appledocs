// Code generated DO NOT EDIT - Manual geometry type definitions

package coregraphics

// CGRect represents a rectangle in a two-dimensional coordinate system.
//
// In the default Core Graphics coordinate space, the origin is located
// in the lower-left corner of the rectangle and the rectangle extends
// towards the upper-right corner. If the context has a flipped-coordinate
// space—often the case on iOS—the origin is in the upper-left corner
// and the rectangle extends towards the lower-right corner.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRect
type Rect struct {
	Origin Point
	Size   Size
}

// Point represents a point in a two-dimensional coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
type Point struct {
	X Float
	Y Float
}

// Size represents width and height values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGSize
type Size struct {
	Width  Float
	Height Float
}

// Float is the basic type for floating-point scalar values in Core Graphics.
// On 64-bit systems, CGFloat is defined as double. On 32-bit systems, it's float.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGFloat
type Float = float64

// AffineTransform represents an affine transformation matrix for use in drawing 2D graphics.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGAffineTransform
type AffineTransform struct {
	A  Float
	B  Float
	C  Float
	D  Float
	TX Float
	TY Float
}
