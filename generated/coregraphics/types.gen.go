// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import "unsafe"

// CGColorConversionInfoRef is a CoreGraphics opaque type.
type CGColorConversionInfoRef unsafe.Pointer

// CGColorRef is a CoreGraphics opaque type.
type CGColorRef unsafe.Pointer

// CGColorSpaceRef is a CoreGraphics opaque type.
type CGColorSpaceRef unsafe.Pointer

// CGContextRef is a CoreGraphics opaque type.
type CGContextRef unsafe.Pointer

// CGDataConsumerRef is a CoreGraphics opaque type.
type CGDataConsumerRef unsafe.Pointer

// CGDataProviderRef is a CoreGraphics opaque type.
type CGDataProviderRef unsafe.Pointer

// CGDisplayConfigRef is a CoreGraphics opaque type.
type CGDisplayConfigRef unsafe.Pointer

// CGDisplayModeRef is a CoreGraphics opaque type.
type CGDisplayModeRef unsafe.Pointer

// CGDisplayStreamRef is a CoreGraphics opaque type.
type CGDisplayStreamRef unsafe.Pointer

// CGDisplayStreamUpdateRef is a CoreGraphics opaque type.
type CGDisplayStreamUpdateRef unsafe.Pointer

// CGEventRef is a CoreGraphics opaque type.
type CGEventRef unsafe.Pointer

// CGEventSourceRef is a CoreGraphics opaque type.
type CGEventSourceRef unsafe.Pointer

// CGFontRef is a CoreGraphics opaque type.
type CGFontRef unsafe.Pointer

// CGFunctionRef is a CoreGraphics opaque type.
type CGFunctionRef unsafe.Pointer

// CGGradientRef is a CoreGraphics opaque type.
type CGGradientRef unsafe.Pointer

// CGImageRef is a CoreGraphics opaque type.
type CGImageRef unsafe.Pointer

// CGLayerRef is a CoreGraphics opaque type.
type CGLayerRef unsafe.Pointer

// CGMutablePathRef is a CoreGraphics opaque type.
type CGMutablePathRef unsafe.Pointer

// CGPDFArrayRef is a CoreGraphics opaque type.
type CGPDFArrayRef unsafe.Pointer

// CGPDFContentStreamRef is a CoreGraphics opaque type.
type CGPDFContentStreamRef unsafe.Pointer

// CGPDFDictionaryRef is a CoreGraphics opaque type.
type CGPDFDictionaryRef unsafe.Pointer

// CGPDFDocumentRef is a CoreGraphics opaque type.
type CGPDFDocumentRef unsafe.Pointer

// CGPDFObjectRef is a CoreGraphics opaque type.
type CGPDFObjectRef unsafe.Pointer

// CGPDFOperatorTableRef is a CoreGraphics opaque type.
type CGPDFOperatorTableRef unsafe.Pointer

// CGPDFPageRef is a CoreGraphics opaque type.
type CGPDFPageRef unsafe.Pointer

// CGPDFScannerRef is a CoreGraphics opaque type.
type CGPDFScannerRef unsafe.Pointer

// CGPDFStreamRef is a CoreGraphics opaque type.
type CGPDFStreamRef unsafe.Pointer

// CGPDFStringRef is a CoreGraphics opaque type.
type CGPDFStringRef unsafe.Pointer

// CGPSConverterRef is a CoreGraphics opaque type.
type CGPSConverterRef unsafe.Pointer

// CGPathRef is a CoreGraphics opaque type.
type CGPathRef unsafe.Pointer

// CGPatternRef is a CoreGraphics opaque type.
type CGPatternRef unsafe.Pointer

// CGRenderingBufferProviderRef is a CoreGraphics opaque type.
type CGRenderingBufferProviderRef unsafe.Pointer

// CGShadingRef is a CoreGraphics opaque type.
type CGShadingRef unsafe.Pointer


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


