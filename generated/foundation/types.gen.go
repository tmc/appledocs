// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation
import (
	"unsafe"
)

// CFStringRef is a CoreGraphics opaque type.
type CFStringRef unsafe.Pointer


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


// Fallback type aliases for undefined types
// These types are referenced in method signatures but not fully documented.
// Using inferred base types as fallback to allow code generation.
type AffineTransformStruct = int

type AttributedStringCompletionHandler = int

type BinarySearchingOptions = int

type Comparator = int

type ComparisonResult = int

type Decimal = int

type EdgeInsets = int

type FontTraitMask = int

type HashEnumerator = int

type HashTableCallBacks = int

type KeyValueObservingOptions = int

type MapEnumerator = int

type MapTableKeyCallBacks = int

type MapTableValueCallBacks = int

type StringDrawingOptions = int

type StringRef = int

type TextAlignment = int

type UInteger = int

type URLHandle = int

type WritingDirection = int

type Zone = int



