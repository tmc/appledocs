// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/tmc/appledocs/generated/objectivec"

// TimeInterval for non-CoreGraphics frameworks
type TimeInterval = float64  // NSTimeInterval
// Foundation-specific types

// Foundation geometry types - aliases to objectivec base types
// This avoids circular imports while maintaining API compatibility
type Point = objectivec.Point
type Size = objectivec.Size
type Rect = objectivec.Rect
type Range = objectivec.Range

// RectEdge defines which edge of a rectangle.
type RectEdge int

const (
	RectEdgeMinX RectEdge = 0
	RectEdgeMinY RectEdge = 1
	RectEdgeMaxX RectEdge = 2
	RectEdgeMaxY RectEdge = 3
)


