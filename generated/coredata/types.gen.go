// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata


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


