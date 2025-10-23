// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation
import (
	"unsafe"
)


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
type Bool = int

type Status = int

type StreamDelegate = int


// Manual fallback types for API collection classes
// These classes come from API collections and aren't in the parsed class list
type AppleEventSendOptions = int
type AppleEventManagerSuspensionID = int
type ByteCountFormatterCountStyle = int
type ByteCountFormatterUnits = int
type DecodingFailurePolicy = int
type ProgressFileOperationKind = int
type ProgressKind = int
type ProgressUserInfoKey = int
type StringEncoding = int
type StringTransform = int


