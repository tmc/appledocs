// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// Common selectors cached for performance
var (
	selAlloc       = objc.RegisterName("alloc")
	selNew         = objc.RegisterName("new")
	selInit        = objc.RegisterName("init")
	selAutorelease = objc.RegisterName("autorelease")
)

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

// Aliases for common use - NSRange, NSPoint, NSSize, NSRect
type Range = CGPoint  // For NSRange - structured as {location, length}
type Point = CGPoint
type Size = CGSize
type Rect = CGRect


