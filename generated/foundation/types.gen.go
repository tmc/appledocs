// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation
import (
	"unsafe"
)

// CFTypeRef is a CoreGraphics opaque type.
type CFTypeRef unsafe.Pointer


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


// Fallback type aliases for undefined types
// These types are referenced in method signatures but not fully documented.
// Using inferred base types as fallback to allow code generation.
type CFIndex = int

type CFStringEncoding = int

type CFTypeID = int

type CFUUIDBytes = int

type NSAffineTransformStruct = int

type NSAttributedStringCompletionHandler = int

type NSComparator = int

type NSComparisonResult = int

type NSDecimal = int

type NSEdgeInsets = int

type NSFontTraitMask = int

type NSFormattingUnitStyle = int

type NSHashEnumerator = int

type NSHashTableCallBacks = int

type NSKeyValueObservingOptions = int

type NSMapEnumerator = int

type NSMapTableKeyCallBacks = int

type NSMapTableValueCallBacks = int

type NSRunLoopMode = int

type NSTextAlignment = int

type NSWritingDirection = int

type UIEdgeInsets = int



