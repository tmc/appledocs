// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import "unsafe"

// CoreGraphics Types

// Fundamental types
type CGFloat float64

// Opaque reference types
type CGColorConversionInfoRef unsafe.Pointer
type CGColorRef unsafe.Pointer
type CGColorSpaceRef unsafe.Pointer
type CGContextRef unsafe.Pointer
type CGDataConsumerRef unsafe.Pointer
type CGDataProviderRef unsafe.Pointer
type CGDisplayConfigRef unsafe.Pointer
type CGDisplayModeRef unsafe.Pointer
type CGDisplayStreamRef unsafe.Pointer
type CGDisplayStreamUpdateRef unsafe.Pointer
type CGEventRef unsafe.Pointer
type CGEventSourceRef unsafe.Pointer
type CGFontRef unsafe.Pointer
type CGFunctionRef unsafe.Pointer
type CGGradientRef unsafe.Pointer
type CGImageRef unsafe.Pointer
type CGLayerRef unsafe.Pointer
type CGMutablePathRef unsafe.Pointer
type CGPDFArrayRef unsafe.Pointer
type CGPDFContentStreamRef unsafe.Pointer
type CGPDFDictionaryRef unsafe.Pointer
type CGPDFDocumentRef unsafe.Pointer
type CGPDFObjectRef unsafe.Pointer
type CGPDFOperatorTableRef unsafe.Pointer
type CGPDFPageRef unsafe.Pointer
type CGPDFScannerRef unsafe.Pointer
type CGPDFStreamRef unsafe.Pointer
type CGPDFStringRef unsafe.Pointer
type CGPSConverterRef unsafe.Pointer
type CGPathRef unsafe.Pointer
type CGPatternRef unsafe.Pointer
type CGRenderingBufferProviderRef unsafe.Pointer
type CGShadingRef unsafe.Pointer

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
