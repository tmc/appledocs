// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import "unsafe"

// CoreGraphics Types

// Fundamental types
type CGFloat float64

// Opaque reference types
type CGContextRef unsafe.Pointer
type CGColorRef unsafe.Pointer
type CGColorSpaceRef unsafe.Pointer
type CGPathRef unsafe.Pointer
type CGImageRef unsafe.Pointer
type CGDataProviderRef unsafe.Pointer
type CGFontRef unsafe.Pointer
type CGGradientRef unsafe.Pointer
type CGLayerRef unsafe.Pointer
type CGPDFDocumentRef unsafe.Pointer
type CGPDFPageRef unsafe.Pointer

// Geometric types
type CGPoint struct {
	X, Y CGFloat
}

type CGSize struct {
	Width, Height CGFloat
}

type CGRect struct {
	Origin CGPoint
	Size   CGSize
}

type CGAffineTransform struct {
	A, B, C, D, Tx, Ty CGFloat
}

