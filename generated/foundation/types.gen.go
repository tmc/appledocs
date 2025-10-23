// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Foundation-specific types

// Foundation geometry types - compatible with NSPoint, NSSize, NSRect, NSRange
type Point struct {
	X float64
	Y float64
}

type Size struct {
	Width  float64
	Height float64
}

type Rect struct {
	Origin Point
	Size   Size
}

type Range struct {
	Location int
	Length   int
}

// RectEdge defines which edge of a rectangle.
type RectEdge int

const (
	RectEdgeMinX RectEdge = 0
	RectEdgeMinY RectEdge = 1
	RectEdgeMaxX RectEdge = 2
	RectEdgeMaxY RectEdge = 3
)


