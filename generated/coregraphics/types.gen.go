// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics
import (
	"unsafe"
)


// CoreGraphics struct types
type CGFloat = float64

type CGPoint struct {
	X CGFloat
	Y CGFloat
}

type CGSize struct {
	Width  CGFloat
	Height CGFloat
}

type CGRect struct {
	Origin CGPoint
	Size   CGSize
}

type CGAffineTransform struct {
	A  CGFloat
	B  CGFloat
	C  CGFloat
	D  CGFloat
	Tx CGFloat
	Ty CGFloat
}

type CGVector struct {
	DX CGFloat
	DY CGFloat
}

// Common type aliases
type Range = CGPoint  // NSRange
type Size = CGSize    // NSSize
type Point = CGPoint  // NSPoint
type Rect = CGRect    // NSRect

// Fallback type aliases for undefined types
// These types are referenced in method signatures but not fully documented.
// Using unsafe.Pointer as fallback to allow code generation.
type CGAffineTransformComponents unsafe.Pointer

type CGColorBufferFormat unsafe.Pointer

type CGColorDataFormat unsafe.Pointer

type CGContentToneMappingInfo unsafe.Pointer

type CGDisplayStreamFrameAvailableHandler unsafe.Pointer

type CGPDFArrayApplierBlock unsafe.Pointer

type CGPDFDictionaryApplierBlock unsafe.Pointer

type CGPathApplyBlock unsafe.Pointer

type CGRectEdge unsafe.Pointer

type CGScreenUpdateMoveDelta unsafe.Pointer



