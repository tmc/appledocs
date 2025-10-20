// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec


// TimeInterval for non-CoreGraphics frameworks
type TimeInterval = float64  // NSTimeInterval

// Base geometry types - compatible with NSPoint, NSSize, NSRect, NSRange
// These are defined in objectivec to avoid circular imports with foundation

// Point represents a point in a two-dimensional coordinate system.
type Point struct {
	X float64
	Y float64
}

// Size represents width and height values.
type Size struct {
	Width  float64
	Height float64
}

// Rect represents a rectangle.
type Rect struct {
	Origin Point
	Size   Size
}

// Range represents a range of sequential items, identified by location and length.
type Range struct {
	Location int
	Length   int
}

