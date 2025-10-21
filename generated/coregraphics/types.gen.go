// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics
import (
	"unsafe"
)

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
type CFAllocatorRef unsafe.Pointer

type CFArrayRef unsafe.Pointer

type CFDataRef unsafe.Pointer

type CFDateRef unsafe.Pointer

type CFDictionaryRef unsafe.Pointer

type CFErrorRef unsafe.Pointer

type CFMachPortRef unsafe.Pointer

type CFMutableDataRef unsafe.Pointer

type CFPropertyListRef unsafe.Pointer

type CFRunLoopSourceRef unsafe.Pointer

type CFStringRef unsafe.Pointer

type CFTimeInterval unsafe.Pointer

type CFTypeID unsafe.Pointer

type CFTypeRef unsafe.Pointer

type CFURLRef unsafe.Pointer

type CGAffineTransformComponents unsafe.Pointer

type CGBitmapContextReleaseDataCallback unsafe.Pointer

type CGButtonCount unsafe.Pointer

type CGCharCode unsafe.Pointer

type CGColorBufferFormat unsafe.Pointer

type CGColorDataFormat unsafe.Pointer

type CGContentToneMappingInfo unsafe.Pointer

type CGDataConsumerCallbacks unsafe.Pointer

type CGDataProviderDirectCallbacks unsafe.Pointer

type CGDataProviderReleaseDataCallback unsafe.Pointer

type CGDataProviderSequentialCallbacks unsafe.Pointer

type CGDirectDisplayID unsafe.Pointer

type CGDisplayBlendFraction unsafe.Pointer

type CGDisplayFadeInterval unsafe.Pointer

type CGDisplayFadeReservationToken unsafe.Pointer

type CGDisplayReconfigurationCallBack unsafe.Pointer

type CGDisplayReservationInterval unsafe.Pointer

type CGDisplayStreamFrameAvailableHandler unsafe.Pointer

type CGErrorCallback unsafe.Pointer

type CGEventMask unsafe.Pointer

type CGEventSourceKeyboardType unsafe.Pointer

type CGEventTapCallBack unsafe.Pointer

type CGEventTapInformation unsafe.Pointer

type CGEventTapProxy unsafe.Pointer

type CGEventTimestamp unsafe.Pointer

type CGFunctionCallbacks unsafe.Pointer

type CGGammaValue unsafe.Pointer

type CGGlyph unsafe.Pointer

type CGKeyCode unsafe.Pointer

type CGOpenGLDisplayMask unsafe.Pointer

type CGPDFArrayApplierBlock unsafe.Pointer

type CGPDFBoolean unsafe.Pointer

type CGPDFDictionaryApplierBlock unsafe.Pointer

type CGPDFDictionaryApplierFunction unsafe.Pointer

type CGPDFInteger unsafe.Pointer

type CGPDFOperatorCallback unsafe.Pointer

type CGPDFReal unsafe.Pointer

type CGPSConverterCallbacks unsafe.Pointer

type CGPathApplierFunction unsafe.Pointer

type CGPathApplyBlock unsafe.Pointer

type CGPatternCallbacks unsafe.Pointer

type CGRectEdge unsafe.Pointer

type CGRefreshRate unsafe.Pointer

type CGScreenRefreshCallback unsafe.Pointer

type CGScreenUpdateMoveCallback unsafe.Pointer

type CGScreenUpdateMoveDelta unsafe.Pointer

type CGWheelCount unsafe.Pointer

type CGWindowID unsafe.Pointer

type CGWindowLevel unsafe.Pointer

type ColorSyncProfileRef unsafe.Pointer

type MTLDevice unsafe.Pointer

type UniChar unsafe.Pointer

type UniCharCount unsafe.Pointer



