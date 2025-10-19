// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

package imageio

import "unsafe"

// CFArrayRef is a CoreGraphics opaque type.
type CFArrayRef unsafe.Pointer

// CFDataRef is a CoreGraphics opaque type.
type CFDataRef unsafe.Pointer

// CFDictionaryRef is a CoreGraphics opaque type.
type CFDictionaryRef unsafe.Pointer

// CFURLRef is a CoreGraphics opaque type.
type CFURLRef unsafe.Pointer

// CGImageDestinationRef is a CoreGraphics opaque type.
type CGImageDestinationRef unsafe.Pointer

// CGImageMetadataRef is a CoreGraphics opaque type.
type CGImageMetadataRef unsafe.Pointer

// CGImageSourceRef is a CoreGraphics opaque type.
type CGImageSourceRef unsafe.Pointer

// CGMutableImageMetadataRef is a CoreGraphics opaque type.
type CGMutableImageMetadataRef unsafe.Pointer


// Common CoreGraphics struct types
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

// Common type aliases
type Range = CGPoint  // NSRange
type Size = CGSize    // NSSize
type Point = CGPoint  // NSPoint
type Rect = CGRect    // NSRect
type TimeInterval = float64  // NSTimeInterval


