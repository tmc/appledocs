// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia

import "unsafe"

// CFAllocatorRef is a CoreGraphics opaque type.
type CFAllocatorRef unsafe.Pointer

// CFArrayRef is a CoreGraphics opaque type.
type CFArrayRef unsafe.Pointer

// CFDictionaryRef is a CoreGraphics opaque type.
type CFDictionaryRef unsafe.Pointer

// CFStringRef is a CoreGraphics opaque type.
type CFStringRef unsafe.Pointer

// CMBufferQueueRef is a CoreGraphics opaque type.
type CMBufferQueueRef unsafe.Pointer

// CMBufferRef is a CoreGraphics opaque type.
type CMBufferRef unsafe.Pointer

// CMClockOrTimebaseRef is a CoreGraphics opaque type.
type CMClockOrTimebaseRef unsafe.Pointer

// CMMutableTagCollectionRef is a CoreGraphics opaque type.
type CMMutableTagCollectionRef unsafe.Pointer

// CMSampleBufferRef is a CoreGraphics opaque type.
type CMSampleBufferRef unsafe.Pointer

// CMTaggedBufferGroupRef is a CoreGraphics opaque type.
type CMTaggedBufferGroupRef unsafe.Pointer

// CMTimebaseRef is a CoreGraphics opaque type.
type CMTimebaseRef unsafe.Pointer


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


