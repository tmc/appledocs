// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import "github.com/ebitengine/purego"

// CoreGraphics Functions (210 total)
//
// This file contains executable function bindings automatically registered via purego.
// All functions are ready to use after package initialization.

// CGContextSetInterpolationQuality is available on macOS 10.0+
var CGContextSetInterpolationQuality func(unsafe.Pointer)

// CGAcquireDisplayFadeReservation is available on macOS 10.2+
var CGAcquireDisplayFadeReservation func(seconds unsafe.Pointer, token unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGAffineTransformConcat is available on macOS 10.0+
var CGAffineTransformConcat func(t1 CGAffineTransform, t2 unsafe.Pointer, unsafe.Pointer) CGAffineTransform

// CGAffineTransformMake is available on macOS 10.0+
var CGAffineTransformMake func(a CGFloat, b unsafe.Pointer, c unsafe.Pointer, d unsafe.Pointer, tx unsafe.Pointer, ty unsafe.Pointer, unsafe.Pointer) CGAffineTransform

// CGAffineTransformMakeTranslation is available on macOS 10.0+
var CGAffineTransformMakeTranslation func(tx CGFloat, ty unsafe.Pointer, unsafe.Pointer) CGAffineTransform

// CGAffineTransformScale is available on macOS 10.0+
var CGAffineTransformScale func(t CGAffineTransform, sx unsafe.Pointer, sy unsafe.Pointer, unsafe.Pointer) CGAffineTransform

// CGAffineTransformTranslate is available on macOS 10.0+
var CGAffineTransformTranslate func(t CGAffineTransform, tx unsafe.Pointer, ty unsafe.Pointer, unsafe.Pointer) CGAffineTransform

// CGBitmapContextCreateAdaptive is available on macOS 26.0+
var CGBitmapContextCreateAdaptive func(width uintptr, height unsafe.Pointer, auxiliaryInfo unsafe.Pointer, onResolve unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, onAllocate unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, onFree unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, onError unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGContextRef

// CGColorGetColorSpace is available on macOS 10.3+
var CGColorGetColorSpace func(color CGColorRef, unsafe.Pointer) CGColorSpaceRef

// CGColorGetContentHeadroom is available on macOS 26.0+
var CGColorGetContentHeadroom func(color CGColorRef, unsafe.Pointer) float32

// CGColorCreate is available on macOS 10.3+
var CGColorCreate func(space CGColorSpaceRef, components unsafe.Pointer, unsafe.Pointer) CGColorRef

// CGColorCreateGenericCMYK is available on macOS 10.5+
var CGColorCreateGenericCMYK func(cyan CGFloat, magenta unsafe.Pointer, yellow unsafe.Pointer, black unsafe.Pointer, alpha unsafe.Pointer, unsafe.Pointer) CGColorRef

// CGColorCreateWithContentHeadroom is available on macOS 26.0+
var CGColorCreateWithContentHeadroom func(headroom float32, space unsafe.Pointer, red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, alpha unsafe.Pointer, unsafe.Pointer) CGColorRef

// CGColorGetPattern is available on macOS 10.3+
var CGColorGetPattern func(color CGColorRef, unsafe.Pointer) CGPatternRef

// CGColorConversionInfoConvertData is available on macOS 15.0+
var CGColorConversionInfoConvertData func(info CGColorConversionInfoRef, width unsafe.Pointer, height unsafe.Pointer, dst_data unsafe.Pointer, dst_format unsafe.Pointer, src_data unsafe.Pointer, src_format unsafe.Pointer, options unsafe.Pointer, unsafe.Pointer) bool

// CGColorConversionInfoCreateForToneMapping is available on macOS 15.0+
var CGColorConversionInfoCreateForToneMapping func(from CGColorSpaceRef, source_headroom unsafe.Pointer, to unsafe.Pointer, target_headroom unsafe.Pointer, method unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer, unsafe.Pointer) CGColorConversionInfoRef

// CGColorConversionInfoCreateFromList is available on macOS 10.12+
var CGColorConversionInfoCreateFromList func(options unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGColorConversionInfoRef

// CGColorRelease is available on macOS 10.3+
var CGColorRelease func(color CGColorRef, unsafe.Pointer)

// CGColorSpaceCreateCalibratedRGB is available on macOS 10.0+
var CGColorSpaceCreateCalibratedRGB func(whitePoint unsafe.Pointer, blackPoint unsafe.Pointer, gamma unsafe.Pointer, matrix unsafe.Pointer, unsafe.Pointer) CGColorSpaceRef

// CGColorSpaceCreatePattern is available on macOS 10.0+
var CGColorSpaceCreatePattern func(baseSpace CGColorSpaceRef, unsafe.Pointer) CGColorSpaceRef

// CGColorSpaceIsWideGamutRGB is available on macOS 10.12+
var CGColorSpaceIsWideGamutRGB func(CGColorSpaceRef, unsafe.Pointer) bool

// CGColorSpaceCreateLinearized is available on macOS 11.0+
var CGColorSpaceCreateLinearized func(space CGColorSpaceRef, unsafe.Pointer) CGColorSpaceRef

// CGColorSpaceIsHLGBased is available on macOS 12.0+
var CGColorSpaceIsHLGBased func(s CGColorSpaceRef, unsafe.Pointer) bool

// CGColorSpaceUsesExtendedRange is available on macOS 10.12+
var CGColorSpaceUsesExtendedRange func(space CGColorSpaceRef, unsafe.Pointer) bool

// CGContextAddPath is available on macOS 10.2+
var CGContextAddPath func(c CGContextRef, path unsafe.Pointer, unsafe.Pointer)

// CGBitmapContextGetBytesPerRow is available on macOS 10.2+
var CGBitmapContextGetBytesPerRow func(context CGContextRef, unsafe.Pointer) uintptr

// CGContextDrawShading is available on macOS 10.2+
var CGContextDrawShading func(c CGContextRef, shading unsafe.Pointer, unsafe.Pointer)

// CGPDFContextEndPage is available on macOS 10.4+
var CGPDFContextEndPage func(context CGContextRef, unsafe.Pointer)

// CGBitmapContextCreate is available on macOS 10.0+
var CGBitmapContextCreate func(data unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, bitsPerComponent unsafe.Pointer, bytesPerRow unsafe.Pointer, space unsafe.Pointer, bitmapInfo unsafe.Pointer, unsafe.Pointer) CGContextRef

// CGBitmapContextCreateWithData is available on macOS 10.6+
var CGBitmapContextCreateWithData func(data unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, bitsPerComponent unsafe.Pointer, bytesPerRow unsafe.Pointer, space unsafe.Pointer, bitmapInfo unsafe.Pointer, releaseCallback unsafe.Pointer, releaseInfo unsafe.Pointer, unsafe.Pointer) CGContextRef

// CGContextGetInterpolationQuality is available on macOS 10.0+
var CGContextGetInterpolationQuality func(c CGContextRef, unsafe.Pointer) unsafe.Pointer

// CGContextRotateCTM is available on macOS 10.0+
var CGContextRotateCTM func(c CGContextRef, angle unsafe.Pointer, unsafe.Pointer)

// CGContextScaleCTM is available on macOS 10.0+
var CGContextScaleCTM func(c CGContextRef, sx unsafe.Pointer, sy unsafe.Pointer, unsafe.Pointer)

// CGPDFContextSetDestinationForRect is available on macOS 10.4+
var CGPDFContextSetDestinationForRect func(context CGContextRef, name unsafe.Pointer, rect unsafe.Pointer, unsafe.Pointer)

// CGContextSetShouldSubpixelPositionFonts is available on macOS 10.5+
var CGContextSetShouldSubpixelPositionFonts func(c CGContextRef, shouldSubpixelPositionFonts unsafe.Pointer, unsafe.Pointer)

// CGContextSetStrokeColor is available on macOS 10.0+
var CGContextSetStrokeColor func(c CGContextRef, components unsafe.Pointer, unsafe.Pointer)

// CGContextSynchronizeAttributes is available on macOS 26.0+
var CGContextSynchronizeAttributes func(c CGContextRef, unsafe.Pointer)

// CGContextGetTextMatrix is available on macOS 10.0+
var CGContextGetTextMatrix func(c CGContextRef, unsafe.Pointer) CGAffineTransform

// CGContextGetContentToneMappingInfo is available on macOS 26.0+
var CGContextGetContentToneMappingInfo func(c CGContextRef, unsafe.Pointer) unsafe.Pointer

// CGContextMoveToPoint is available on macOS 10.0+
var CGContextMoveToPoint func(c CGContextRef, x unsafe.Pointer, y unsafe.Pointer, unsafe.Pointer)

// CGContextSetContentToneMappingInfo is available on macOS 26.0+
var CGContextSetContentToneMappingInfo func(c CGContextRef, info unsafe.Pointer, unsafe.Pointer)

// CGContextSetLineDash is available on macOS 10.0+
var CGContextSetLineDash func(c CGContextRef, phase unsafe.Pointer, lengths unsafe.Pointer, count unsafe.Pointer, unsafe.Pointer)

// CGContextStrokeLineSegments is available on macOS 10.4+
var CGContextStrokeLineSegments func(c CGContextRef, points unsafe.Pointer, count unsafe.Pointer, unsafe.Pointer)

// CGDataConsumerCreateWithCFData is available on macOS 10.4+
var CGDataConsumerCreateWithCFData func(data unsafe.Pointer, unsafe.Pointer) CGDataConsumerRef

// CGDataConsumerCreateWithURL is available on macOS 10.0+
var CGDataConsumerCreateWithURL func(url unsafe.Pointer, unsafe.Pointer) CGDataConsumerRef

// CGDataConsumerRelease is available on macOS 10.0+
var CGDataConsumerRelease func(consumer CGDataConsumerRef, unsafe.Pointer)

// CGDataConsumerRetain is available on macOS 10.0+
var CGDataConsumerRetain func(consumer CGDataConsumerRef, unsafe.Pointer) CGDataConsumerRef

// CGDataProviderCopyData is available on macOS 10.3+
var CGDataProviderCopyData func(provider CGDataProviderRef, unsafe.Pointer) unsafe.Pointer

// CGDataProviderGetInfo is available on macOS 10.13+
var CGDataProviderGetInfo func(provider CGDataProviderRef, unsafe.Pointer) unsafe.Pointer

// CGDataProviderCreateWithData is available on macOS 10.0+
var CGDataProviderCreateWithData func(info unsafe.Pointer, data unsafe.Pointer, size unsafe.Pointer, releaseData unsafe.Pointer, unsafe.Pointer) CGDataProviderRef

// CGDataProviderCreateDirect is available on macOS 10.5+
var CGDataProviderCreateDirect func(info unsafe.Pointer, size unsafe.Pointer, callbacks unsafe.Pointer, unsafe.Pointer) CGDataProviderRef

// CGDataProviderCreateWithFilename is available on macOS 10.0+
var CGDataProviderCreateWithFilename func(filename unsafe.Pointer, unsafe.Pointer) CGDataProviderRef

// CGDataProviderCreateSequential is available on macOS 10.5+
var CGDataProviderCreateSequential func(info unsafe.Pointer, callbacks unsafe.Pointer, unsafe.Pointer) CGDataProviderRef

// CGDataProviderRetain is available on macOS 10.0+
var CGDataProviderRetain func(provider CGDataProviderRef, unsafe.Pointer) CGDataProviderRef

// CGDisplayCaptureWithOptions is available on macOS 10.3+
var CGDisplayCaptureWithOptions func(display unsafe.Pointer, options unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGDisplayIsActive is available on macOS 10.2+
var CGDisplayIsActive func(display unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGDisplayModeGetPixelHeight is available on macOS 10.8+
var CGDisplayModeGetPixelHeight func(mode CGDisplayModeRef, unsafe.Pointer) uintptr

// CGDisplayRemoveReconfigurationCallback is available on macOS 10.3+
var CGDisplayRemoveReconfigurationCallback func(callback unsafe.Pointer, userInfo unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

var CGDisplayStreamUpdateGetRects func(updateRef CGDisplayStreamUpdateRef, rectType unsafe.Pointer, rectCount unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGEventCreateMouseEvent is available on macOS 10.4+
var CGEventCreateMouseEvent func(source CGEventSourceRef, mouseType unsafe.Pointer, mouseCursorPosition unsafe.Pointer, mouseButton unsafe.Pointer, unsafe.Pointer) CGEventRef

// CGEventKeyboardGetUnicodeString is available on macOS 10.4+
var CGEventKeyboardGetUnicodeString func(event CGEventRef, maxStringLength unsafe.Pointer, actualStringLength unsafe.Pointer, unicodeString unsafe.Pointer, unsafe.Pointer)

// CGEventKeyboardSetUnicodeString is available on macOS 10.4+
var CGEventKeyboardSetUnicodeString func(event CGEventRef, stringLength unsafe.Pointer, unicodeString unsafe.Pointer, unsafe.Pointer)

// CGEventTapCreate is available on macOS 10.4+
var CGEventTapCreate func(tap unsafe.Pointer, place unsafe.Pointer, options unsafe.Pointer, eventsOfInterest unsafe.Pointer, callback unsafe.Pointer, userInfo unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGEventTapCreateForPSN is available on macOS 10.4+
var CGEventTapCreateForPSN func(processSerialNumber unsafe.Pointer, place unsafe.Pointer, options unsafe.Pointer, eventsOfInterest unsafe.Pointer, callback unsafe.Pointer, userInfo unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGEventTapPostEvent is available on macOS 10.4+
var CGEventTapPostEvent func(proxy unsafe.Pointer, event unsafe.Pointer, unsafe.Pointer)

// CGEventCreateScrollWheelEvent is available on macOS 10.5+
var CGEventCreateScrollWheelEvent func(source CGEventSourceRef, units unsafe.Pointer, wheelCount unsafe.Pointer, wheel1 unsafe.Pointer, unsafe.Pointer) CGEventRef

// CGEventSetFlags is available on macOS 10.4+
var CGEventSetFlags func(event CGEventRef, flags unsafe.Pointer, unsafe.Pointer)

// CGEventSourceCounterForEventType is available on macOS 10.4+
var CGEventSourceCounterForEventType func(stateID unsafe.Pointer, eventType unsafe.Pointer, unsafe.Pointer) uint32

// CGEventSourceFlagsState is available on macOS 10.4+
var CGEventSourceFlagsState func(stateID unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGEventSourceGetLocalEventsSuppressionInterval is available on macOS 10.4+
var CGEventSourceGetLocalEventsSuppressionInterval func(source CGEventSourceRef, unsafe.Pointer) unsafe.Pointer

// CGFontGetCapHeight is available on macOS 10.5+
var CGFontGetCapHeight func(font CGFontRef, unsafe.Pointer) int

// CGFontCreateCopyWithVariations is available on macOS 10.4+
var CGFontCreateCopyWithVariations func(font CGFontRef, variations unsafe.Pointer, unsafe.Pointer) CGFontRef

// CGFontCreatePostScriptEncoding is available on macOS 10.4+
var CGFontCreatePostScriptEncoding func(font CGFontRef, encoding unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGFontGetDescent is available on macOS 10.5+
var CGFontGetDescent func(font CGFontRef, unsafe.Pointer) int

// CGFontGetGlyphWithGlyphName is available on macOS 10.5+
var CGFontGetGlyphWithGlyphName func(font CGFontRef, name unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGFontGetLeading is available on macOS 10.5+
var CGFontGetLeading func(font CGFontRef, unsafe.Pointer) int

// CGFontGetStemV is available on macOS 10.5+
var CGFontGetStemV func(font CGFontRef, unsafe.Pointer) CGFloat

// CGFontCopyTableForTag is available on macOS 10.5+
var CGFontCopyTableForTag func(font CGFontRef, tag unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGFontGetXHeight is available on macOS 10.5+
var CGFontGetXHeight func(font CGFontRef, unsafe.Pointer) int

// CGFontCreateWithPlatformFont is available on macOS 10.0+ (Deprecated in 10.6)
var CGFontCreateWithPlatformFont func(platformFontReference unsafe.Pointer, unsafe.Pointer) CGFontRef

// CGFontRelease is available on macOS 10.0+
var CGFontRelease func(font CGFontRef, unsafe.Pointer)

// CGFunctionCreate is available on macOS 10.2+
var CGFunctionCreate func(info unsafe.Pointer, domainDimension unsafe.Pointer, domain unsafe.Pointer, rangeDimension unsafe.Pointer, range unsafe.Pointer, callbacks unsafe.Pointer, unsafe.Pointer) CGFunctionRef

// CGFunctionRelease is available on macOS 10.2+
var CGFunctionRelease func(function CGFunctionRef, unsafe.Pointer)

// CGFunctionRetain is available on macOS 10.2+
var CGFunctionRetain func(function CGFunctionRef, unsafe.Pointer) CGFunctionRef

// CGGradientGetContentHeadroom is available on macOS 26.0+
var CGGradientGetContentHeadroom func(gradient CGGradientRef, unsafe.Pointer) float32

// CGGradientCreateWithColorComponents is available on macOS 10.5+
var CGGradientCreateWithColorComponents func(space CGColorSpaceRef, components unsafe.Pointer, locations unsafe.Pointer, count unsafe.Pointer, unsafe.Pointer) CGGradientRef

// CGGradientCreateWithColors is available on macOS 10.5+
var CGGradientCreateWithColors func(space CGColorSpaceRef, colors unsafe.Pointer, locations unsafe.Pointer, unsafe.Pointer) CGGradientRef

// CGGradientCreateWithContentHeadroom is available on macOS 26.0+
var CGGradientCreateWithContentHeadroom func(headroom float32, space unsafe.Pointer, components unsafe.Pointer, locations unsafe.Pointer, count unsafe.Pointer, unsafe.Pointer) CGGradientRef

// CGGradientRelease is available on macOS 10.5+
var CGGradientRelease func(gradient CGGradientRef, unsafe.Pointer)

// CGGradientRetain is available on macOS 10.5+
var CGGradientRetain func(gradient CGGradientRef, unsafe.Pointer) CGGradientRef

// CGImageGetAlphaInfo is available on macOS 10.0+
var CGImageGetAlphaInfo func(image CGImageRef, unsafe.Pointer) unsafe.Pointer

// CGImageGetBitsPerComponent is available on macOS 10.0+
var CGImageGetBitsPerComponent func(image CGImageRef, unsafe.Pointer) uintptr

// CGImageGetByteOrderInfo is available on macOS 10.14+
var CGImageGetByteOrderInfo func(image CGImageRef, unsafe.Pointer) unsafe.Pointer

// CGImageGetBytesPerRow is available on macOS 10.0+
var CGImageGetBytesPerRow func(image CGImageRef, unsafe.Pointer) uintptr

// CGImageCalculateContentAverageLightLevel is available on macOS 26.0+
var CGImageCalculateContentAverageLightLevel func(image CGImageRef, unsafe.Pointer) float32

// CGImageCreateWithJPEGDataProvider is available on macOS 10.1+
var CGImageCreateWithJPEGDataProvider func(source CGDataProviderRef, decode unsafe.Pointer, shouldInterpolate unsafe.Pointer, intent unsafe.Pointer, unsafe.Pointer) CGImageRef

// CGImageCreate is available on macOS 10.0+
var CGImageCreate func(width uintptr, height unsafe.Pointer, bitsPerComponent unsafe.Pointer, bitsPerPixel unsafe.Pointer, bytesPerRow unsafe.Pointer, space unsafe.Pointer, bitmapInfo unsafe.Pointer, provider unsafe.Pointer, decode unsafe.Pointer, shouldInterpolate unsafe.Pointer, intent unsafe.Pointer, unsafe.Pointer) CGImageRef

var CGWindowListCreateImageFromArray func(screenBounds CGRect, windowArray unsafe.Pointer, imageOption unsafe.Pointer, unsafe.Pointer) CGImageRef

// CGImageIsMask is available on macOS 10.0+
var CGImageIsMask func(image CGImageRef, unsafe.Pointer) bool

// CGImageCreateWithMask is available on macOS 10.4+
var CGImageCreateWithMask func(image CGImageRef, mask unsafe.Pointer, unsafe.Pointer) CGImageRef

// CGImageRetain is available on macOS 10.0+
var CGImageRetain func(image CGImageRef, unsafe.Pointer) CGImageRef

// CGLayerGetContext is available on macOS 10.4+
var CGLayerGetContext func(layer CGLayerRef, unsafe.Pointer) CGContextRef

// CGLayerCreateWithContext is available on macOS 10.4+
var CGLayerCreateWithContext func(context CGContextRef, size unsafe.Pointer, auxiliaryInfo unsafe.Pointer, unsafe.Pointer) CGLayerRef

// CGLayerGetSize is available on macOS 10.4+
var CGLayerGetSize func(layer CGLayerRef, unsafe.Pointer) CGSize

// CGLayerRelease is available on macOS 10.4+
var CGLayerRelease func(layer CGLayerRef, unsafe.Pointer)

// CGLayerRetain is available on macOS 10.4+
var CGLayerRetain func(layer CGLayerRef, unsafe.Pointer) CGLayerRef

// CGPathCloseSubpath is available on macOS 10.2+
var CGPathCloseSubpath func(path CGMutablePathRef, unsafe.Pointer)

// CGPDFArrayGetBoolean is available on macOS 10.3+
var CGPDFArrayGetBoolean func(array CGPDFArrayRef, index unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFArrayGetDictionary is available on macOS 10.3+
var CGPDFArrayGetDictionary func(array CGPDFArrayRef, index unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFArrayGetInteger is available on macOS 10.3+
var CGPDFArrayGetInteger func(array CGPDFArrayRef, index unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFArrayGetNull is available on macOS 10.3+
var CGPDFArrayGetNull func(array CGPDFArrayRef, index unsafe.Pointer, unsafe.Pointer) bool

// CGPDFArrayGetObject is available on macOS 10.3+
var CGPDFArrayGetObject func(array CGPDFArrayRef, index unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFArrayGetString is available on macOS 10.3+
var CGPDFArrayGetString func(array CGPDFArrayRef, index unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFContentStreamGetResource is available on macOS 10.4+
var CGPDFContentStreamGetResource func(cs CGPDFContentStreamRef, category unsafe.Pointer, name unsafe.Pointer, unsafe.Pointer) CGPDFObjectRef

// CGPDFContentStreamRelease is available on macOS 10.4+
var CGPDFContentStreamRelease func(cs CGPDFContentStreamRef, unsafe.Pointer)

// CGPDFContentStreamRetain is available on macOS 10.4+
var CGPDFContentStreamRetain func(cs CGPDFContentStreamRef, unsafe.Pointer) CGPDFContentStreamRef

// CGPDFContextBeginTag is available on macOS 10.15+
var CGPDFContextBeginTag func(context CGContextRef, tagType unsafe.Pointer, tagProperties unsafe.Pointer, unsafe.Pointer)

// CGPDFDictionaryApplyFunction is available on macOS 10.3+
var CGPDFDictionaryApplyFunction func(dict CGPDFDictionaryRef, function unsafe.Pointer, info unsafe.Pointer, unsafe.Pointer)

// CGPDFDictionaryGetArray is available on macOS 10.3+
var CGPDFDictionaryGetArray func(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFDictionaryGetBoolean is available on macOS 10.3+
var CGPDFDictionaryGetBoolean func(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFDictionaryGetCount is available on macOS 10.3+
var CGPDFDictionaryGetCount func(dict CGPDFDictionaryRef, unsafe.Pointer) uintptr

// CGPDFDictionaryGetDictionary is available on macOS 10.3+
var CGPDFDictionaryGetDictionary func(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFDictionaryGetObject is available on macOS 10.3+
var CGPDFDictionaryGetObject func(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFDictionaryGetStream is available on macOS 10.3+
var CGPDFDictionaryGetStream func(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFDictionaryGetString is available on macOS 10.3+
var CGPDFDictionaryGetString func(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFDocumentGetAccessPermissions is available on macOS 10.13+
var CGPDFDocumentGetAccessPermissions func(document CGPDFDocumentRef, unsafe.Pointer) unsafe.Pointer

// CGPDFDocumentAllowsCopying is available on macOS 10.2+
var CGPDFDocumentAllowsCopying func(document CGPDFDocumentRef, unsafe.Pointer) bool

// CGPDFDocumentGetID is available on macOS 10.4+
var CGPDFDocumentGetID func(document CGPDFDocumentRef, unsafe.Pointer) CGPDFArrayRef

// CGPDFDocumentCreateWithProvider is available on macOS 10.0+
var CGPDFDocumentCreateWithProvider func(provider CGDataProviderRef, unsafe.Pointer) CGPDFDocumentRef

// CGPDFDocumentGetNumberOfPages is available on macOS 10.0+
var CGPDFDocumentGetNumberOfPages func(document CGPDFDocumentRef, unsafe.Pointer) uintptr

// CGPDFDocumentGetOutline is available on macOS 10.13+
var CGPDFDocumentGetOutline func(document CGPDFDocumentRef, unsafe.Pointer) unsafe.Pointer

// CGPDFDocumentUnlockWithPassword is available on macOS 10.2+
var CGPDFDocumentUnlockWithPassword func(document CGPDFDocumentRef, password unsafe.Pointer, unsafe.Pointer) bool

// CGPDFDocumentGetCropBox is available on macOS 10.0+ (Deprecated in 10.5)
var CGPDFDocumentGetCropBox func(document CGPDFDocumentRef, page unsafe.Pointer, unsafe.Pointer) CGRect

// CGPDFDocumentGetTrimBox is available on macOS 10.0+ (Deprecated in 10.5)
var CGPDFDocumentGetTrimBox func(document CGPDFDocumentRef, page unsafe.Pointer, unsafe.Pointer) CGRect

// CGPDFDocumentRetain is available on macOS 10.0+
var CGPDFDocumentRetain func(document CGPDFDocumentRef, unsafe.Pointer) CGPDFDocumentRef

// CGPDFObjectGetType is available on macOS 10.3+
var CGPDFObjectGetType func(object CGPDFObjectRef, unsafe.Pointer) unsafe.Pointer

// CGPDFObjectGetValue is available on macOS 10.3+
var CGPDFObjectGetValue func(object CGPDFObjectRef, type unsafe.Pointer, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFOperatorTableRetain is available on macOS 10.4+
var CGPDFOperatorTableRetain func(table CGPDFOperatorTableRef, unsafe.Pointer) CGPDFOperatorTableRef

// CGPDFOperatorTableSetCallback is available on macOS 10.4+
var CGPDFOperatorTableSetCallback func(table CGPDFOperatorTableRef, name unsafe.Pointer, callback unsafe.Pointer, unsafe.Pointer)

// CGPDFPageGetDictionary is available on macOS 10.3+
var CGPDFPageGetDictionary func(page CGPDFPageRef, unsafe.Pointer) CGPDFDictionaryRef

// CGPDFPageGetDocument is available on macOS 10.3+
var CGPDFPageGetDocument func(page CGPDFPageRef, unsafe.Pointer) CGPDFDocumentRef

// CGPDFPageGetBoxRect is available on macOS 10.3+
var CGPDFPageGetBoxRect func(page CGPDFPageRef, box unsafe.Pointer, unsafe.Pointer) CGRect

// CGPDFPageGetDrawingTransform is available on macOS 10.3+
var CGPDFPageGetDrawingTransform func(page CGPDFPageRef, box unsafe.Pointer, rect unsafe.Pointer, rotate unsafe.Pointer, preserveAspectRatio unsafe.Pointer, unsafe.Pointer) CGAffineTransform

// CGPDFPageGetPageNumber is available on macOS 10.3+
var CGPDFPageGetPageNumber func(page CGPDFPageRef, unsafe.Pointer) uintptr

// CGPDFPageGetRotationAngle is available on macOS 10.3+
var CGPDFPageGetRotationAngle func(page CGPDFPageRef, unsafe.Pointer) int

// CGPDFPageRelease is available on macOS 10.3+
var CGPDFPageRelease func(page CGPDFPageRef, unsafe.Pointer)

// CGPDFPageRetain is available on macOS 10.3+
var CGPDFPageRetain func(page CGPDFPageRef, unsafe.Pointer) CGPDFPageRef

// CGPDFScannerCreate is available on macOS 10.4+
var CGPDFScannerCreate func(cs CGPDFContentStreamRef, table unsafe.Pointer, info unsafe.Pointer, unsafe.Pointer) CGPDFScannerRef

// CGPDFScannerPopArray is available on macOS 10.4+
var CGPDFScannerPopArray func(scanner CGPDFScannerRef, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFScannerPopDictionary is available on macOS 10.4+
var CGPDFScannerPopDictionary func(scanner CGPDFScannerRef, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFScannerPopName is available on macOS 10.4+
var CGPDFScannerPopName func(scanner CGPDFScannerRef, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFScannerPopNumber is available on macOS 10.4+
var CGPDFScannerPopNumber func(scanner CGPDFScannerRef, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFScannerPopString is available on macOS 10.4+
var CGPDFScannerPopString func(scanner CGPDFScannerRef, value unsafe.Pointer, unsafe.Pointer) bool

// CGPDFScannerRelease is available on macOS 10.4+
var CGPDFScannerRelease func(scanner CGPDFScannerRef, unsafe.Pointer)

// CGPDFScannerRetain is available on macOS 10.4+
var CGPDFScannerRetain func(scanner CGPDFScannerRef, unsafe.Pointer) CGPDFScannerRef

// CGPDFStreamCopyData is available on macOS 10.3+
var CGPDFStreamCopyData func(stream CGPDFStreamRef, format unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGPDFStreamGetDictionary is available on macOS 10.3+
var CGPDFStreamGetDictionary func(stream CGPDFStreamRef, unsafe.Pointer) CGPDFDictionaryRef

// CGPDFStringCopyDate is available on macOS 10.4+
var CGPDFStringCopyDate func(string CGPDFStringRef, unsafe.Pointer) unsafe.Pointer

// CGPDFStringCopyTextString is available on macOS 10.3+
var CGPDFStringCopyTextString func(string CGPDFStringRef, unsafe.Pointer) unsafe.Pointer

// CGPDFStringGetBytePtr is available on macOS 10.3+
var CGPDFStringGetBytePtr func(string CGPDFStringRef, unsafe.Pointer) unsafe.Pointer

// CGPDFStringGetLength is available on macOS 10.3+
var CGPDFStringGetLength func(string CGPDFStringRef, unsafe.Pointer) uintptr

// CGPSConverterAbort is available on macOS 10.3+
var CGPSConverterAbort func(converter CGPSConverterRef, unsafe.Pointer) bool

// CGPSConverterConvert is available on macOS 10.3+
var CGPSConverterConvert func(converter CGPSConverterRef, provider unsafe.Pointer, consumer unsafe.Pointer, options unsafe.Pointer, unsafe.Pointer) bool

// CGPSConverterCreate is available on macOS 10.3+
var CGPSConverterCreate func(info unsafe.Pointer, callbacks unsafe.Pointer, options unsafe.Pointer, unsafe.Pointer) CGPSConverterRef

// CGPathGetBoundingBox is available on macOS 10.2+
var CGPathGetBoundingBox func(path CGPathRef, unsafe.Pointer) CGRect

// CGPathGetCurrentPoint is available on macOS 10.2+
var CGPathGetCurrentPoint func(path CGPathRef, unsafe.Pointer) CGPoint

// CGPathCreateWithRect is available on macOS 10.5+
var CGPathCreateWithRect func(rect CGRect, transform unsafe.Pointer, unsafe.Pointer) CGPathRef

// CGPathCreateWithRoundedRect is available on macOS 10.9+
var CGPathCreateWithRoundedRect func(rect CGRect, cornerWidth unsafe.Pointer, cornerHeight unsafe.Pointer, transform unsafe.Pointer, unsafe.Pointer) CGPathRef

// CGPathIsRect is available on macOS 10.2+
var CGPathIsRect func(path CGPathRef, rect unsafe.Pointer, unsafe.Pointer) bool

// CGPathCreateMutableCopy is available on macOS 10.2+
var CGPathCreateMutableCopy func(path CGPathRef, unsafe.Pointer) CGMutablePathRef

// CGPathAddArcToPoint is available on macOS 10.2+
var CGPathAddArcToPoint func(path CGMutablePathRef, m unsafe.Pointer, x1 unsafe.Pointer, y1 unsafe.Pointer, x2 unsafe.Pointer, y2 unsafe.Pointer, radius unsafe.Pointer, unsafe.Pointer)

// CGPathAddCurveToPoint is available on macOS 10.2+
var CGPathAddCurveToPoint func(path CGMutablePathRef, m unsafe.Pointer, cp1x unsafe.Pointer, cp1y unsafe.Pointer, cp2x unsafe.Pointer, cp2y unsafe.Pointer, x unsafe.Pointer, y unsafe.Pointer, unsafe.Pointer)

// CGPathAddLineToPoint is available on macOS 10.2+
var CGPathAddLineToPoint func(path CGMutablePathRef, m unsafe.Pointer, x unsafe.Pointer, y unsafe.Pointer, unsafe.Pointer)

// CGPathAddLines is available on macOS 10.2+
var CGPathAddLines func(path CGMutablePathRef, m unsafe.Pointer, points unsafe.Pointer, count unsafe.Pointer, unsafe.Pointer)

// CGPathAddPath is available on macOS 10.2+
var CGPathAddPath func(path1 CGMutablePathRef, m unsafe.Pointer, path2 unsafe.Pointer, unsafe.Pointer)

// CGPathAddQuadCurveToPoint is available on macOS 10.2+
var CGPathAddQuadCurveToPoint func(path CGMutablePathRef, m unsafe.Pointer, cpx unsafe.Pointer, cpy unsafe.Pointer, x unsafe.Pointer, y unsafe.Pointer, unsafe.Pointer)

// CGPathCreateCopyByDashingPath is available on macOS 10.7+
var CGPathCreateCopyByDashingPath func(path CGPathRef, transform unsafe.Pointer, phase unsafe.Pointer, lengths unsafe.Pointer, count unsafe.Pointer, unsafe.Pointer) CGPathRef

// CGPathCreateCopyByFlattening is available on macOS 13.0+
var CGPathCreateCopyByFlattening func(path CGPathRef, flatteningThreshold unsafe.Pointer, unsafe.Pointer) CGPathRef

// CGPathCreateCopyByStrokingPath is available on macOS 10.7+
var CGPathCreateCopyByStrokingPath func(path CGPathRef, transform unsafe.Pointer, lineWidth unsafe.Pointer, lineCap unsafe.Pointer, lineJoin unsafe.Pointer, miterLimit unsafe.Pointer, unsafe.Pointer) CGPathRef

// CGPathCreateCopyBySymmetricDifferenceOfPath is available on macOS 13.0+
var CGPathCreateCopyBySymmetricDifferenceOfPath func(path CGPathRef, maskPath unsafe.Pointer, evenOddFillRule unsafe.Pointer, unsafe.Pointer) CGPathRef

// CGPathMoveToPoint is available on macOS 10.2+
var CGPathMoveToPoint func(path CGMutablePathRef, m unsafe.Pointer, x unsafe.Pointer, y unsafe.Pointer, unsafe.Pointer)

// CGPatternCreate is available on macOS 10.0+
var CGPatternCreate func(info unsafe.Pointer, bounds unsafe.Pointer, matrix unsafe.Pointer, xStep unsafe.Pointer, yStep unsafe.Pointer, tiling unsafe.Pointer, isColored unsafe.Pointer, callbacks unsafe.Pointer, unsafe.Pointer) CGPatternRef

// CGPatternRelease is available on macOS 10.0+
var CGPatternRelease func(pattern CGPatternRef, unsafe.Pointer)

// CGPatternRetain is available on macOS 10.0+
var CGPatternRetain func(pattern CGPatternRef, unsafe.Pointer) CGPatternRef

// CGPostMouseEvent is available on macOS 10.0+ (Deprecated in 10.6)
var CGPostMouseEvent func(mouseCursorPosition CGPoint, updateMouseCursorPosition unsafe.Pointer, buttonCount unsafe.Pointer, mouseButtonDown unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGRectApplyAffineTransform is available on macOS 10.4+
var CGRectApplyAffineTransform func(rect CGRect, t unsafe.Pointer, unsafe.Pointer) CGRect

// CGRectEqualToRect is available on macOS 10.0+
var CGRectEqualToRect func(rect1 CGRect, rect2 unsafe.Pointer, unsafe.Pointer) bool

// CGRectGetMaxY is available on macOS 10.0+
var CGRectGetMaxY func(rect CGRect, unsafe.Pointer) CGFloat

// CGRectGetMinX is available on macOS 10.0+
var CGRectGetMinX func(rect CGRect, unsafe.Pointer) CGFloat

// CGRectGetMinY is available on macOS 10.0+
var CGRectGetMinY func(rect CGRect, unsafe.Pointer) CGFloat

// CGRectIsNull is available on macOS 10.0+
var CGRectIsNull func(rect CGRect, unsafe.Pointer) bool

// CGRectOffset is available on macOS 10.0+
var CGRectOffset func(rect CGRect, dx unsafe.Pointer, dy unsafe.Pointer, unsafe.Pointer) CGRect

// CGRenderingBufferLockBytePtr is available on macOS 26.0+
var CGRenderingBufferLockBytePtr func(provider CGRenderingBufferProviderRef, unsafe.Pointer) unsafe.Pointer

// CGRenderingBufferProviderCreate is available on macOS 26.0+
var CGRenderingBufferProviderCreate func(info unsafe.Pointer, size unsafe.Pointer, lockPointer unsafe.Pointer, info unsafe.Pointer, unlockPointer unsafe.Pointer, info unsafe.Pointer, pointer unsafe.Pointer, releaseInfo unsafe.Pointer, info unsafe.Pointer, unsafe.Pointer) CGRenderingBufferProviderRef

// CGRenderingBufferProviderCreateWithCFData is available on macOS 26.0+
var CGRenderingBufferProviderCreateWithCFData func(data unsafe.Pointer, unsafe.Pointer) CGRenderingBufferProviderRef

// CGRenderingBufferProviderGetSize is available on macOS 26.0+
var CGRenderingBufferProviderGetSize func(provider CGRenderingBufferProviderRef, unsafe.Pointer) uintptr

// CGRenderingBufferUnlockBytePtr is available on macOS 26.0+
var CGRenderingBufferUnlockBytePtr func(provider CGRenderingBufferProviderRef, unsafe.Pointer)

// CGShadingGetContentHeadroom is available on macOS 26.0+
var CGShadingGetContentHeadroom func(shading CGShadingRef, unsafe.Pointer) float32

// CGShadingCreateAxialWithContentHeadroom is available on macOS 26.0+
var CGShadingCreateAxialWithContentHeadroom func(headroom float32, space unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer, function unsafe.Pointer, extendStart unsafe.Pointer, extendEnd unsafe.Pointer, unsafe.Pointer) CGShadingRef

// CGShadingCreateAxial is available on macOS 10.2+
var CGShadingCreateAxial func(space CGColorSpaceRef, start unsafe.Pointer, end unsafe.Pointer, function unsafe.Pointer, extendStart unsafe.Pointer, extendEnd unsafe.Pointer, unsafe.Pointer) CGShadingRef

// CGShadingCreateRadialWithContentHeadroom is available on macOS 26.0+
var CGShadingCreateRadialWithContentHeadroom func(headroom float32, space unsafe.Pointer, start unsafe.Pointer, startRadius unsafe.Pointer, end unsafe.Pointer, endRadius unsafe.Pointer, function unsafe.Pointer, extendStart unsafe.Pointer, extendEnd unsafe.Pointer, unsafe.Pointer) CGShadingRef

// CGShadingCreateRadial is available on macOS 10.2+
var CGShadingCreateRadial func(space CGColorSpaceRef, start unsafe.Pointer, startRadius unsafe.Pointer, end unsafe.Pointer, endRadius unsafe.Pointer, function unsafe.Pointer, extendStart unsafe.Pointer, extendEnd unsafe.Pointer, unsafe.Pointer) CGShadingRef

// CGShadingRelease is available on macOS 10.2+
var CGShadingRelease func(shading CGShadingRef, unsafe.Pointer)

// CGShadingRetain is available on macOS 10.2+
var CGShadingRetain func(shading CGShadingRef, unsafe.Pointer) CGShadingRef

// CGSizeApplyAffineTransform is available on macOS 10.0+
var CGSizeApplyAffineTransform func(size CGSize, t unsafe.Pointer, unsafe.Pointer) CGSize

// CGSizeCreateDictionaryRepresentation is available on macOS 10.5+
var CGSizeCreateDictionaryRepresentation func(size CGSize, unsafe.Pointer) unsafe.Pointer

// CGSizeMakeWithDictionaryRepresentation is available on macOS 10.5+
var CGSizeMakeWithDictionaryRepresentation func(dict unsafe.Pointer, size unsafe.Pointer, unsafe.Pointer) bool

var CGWaitForScreenRefreshRects func(rects unsafe.Pointer, count unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGWindowListCopyWindowInfo is available on macOS 10.5+
var CGWindowListCopyWindowInfo func(option unsafe.Pointer, relativeToWindow unsafe.Pointer, unsafe.Pointer) unsafe.Pointer

// CGWindowListCreateDescriptionFromArray is available on macOS 10.5+
var CGWindowListCreateDescriptionFromArray func(windowArray unsafe.Pointer, unsafe.Pointer) unsafe.Pointer


// registerFunctions registers all framework functions with purego
func registerFunctions() {
	purego.RegisterLibFunc(&CGContextSetInterpolationQuality, lib, "CGContextSetInterpolationQuality")
	purego.RegisterLibFunc(&CGAcquireDisplayFadeReservation, lib, "CGAcquireDisplayFadeReservation")
	purego.RegisterLibFunc(&CGAffineTransformConcat, lib, "CGAffineTransformConcat")
	purego.RegisterLibFunc(&CGAffineTransformMake, lib, "CGAffineTransformMake")
	purego.RegisterLibFunc(&CGAffineTransformMakeTranslation, lib, "CGAffineTransformMakeTranslation")
	purego.RegisterLibFunc(&CGAffineTransformScale, lib, "CGAffineTransformScale")
	purego.RegisterLibFunc(&CGAffineTransformTranslate, lib, "CGAffineTransformTranslate")
	purego.RegisterLibFunc(&CGBitmapContextCreateAdaptive, lib, "CGBitmapContextCreateAdaptive")
	purego.RegisterLibFunc(&CGColorGetColorSpace, lib, "CGColorGetColorSpace")
	purego.RegisterLibFunc(&CGColorGetContentHeadroom, lib, "CGColorGetContentHeadroom")
	purego.RegisterLibFunc(&CGColorCreate, lib, "CGColorCreate")
	purego.RegisterLibFunc(&CGColorCreateGenericCMYK, lib, "CGColorCreateGenericCMYK")
	purego.RegisterLibFunc(&CGColorCreateWithContentHeadroom, lib, "CGColorCreateWithContentHeadroom")
	purego.RegisterLibFunc(&CGColorGetPattern, lib, "CGColorGetPattern")
	purego.RegisterLibFunc(&CGColorConversionInfoConvertData, lib, "CGColorConversionInfoConvertData")
	purego.RegisterLibFunc(&CGColorConversionInfoCreateForToneMapping, lib, "CGColorConversionInfoCreateForToneMapping")
	purego.RegisterLibFunc(&CGColorConversionInfoCreateFromList, lib, "CGColorConversionInfoCreateFromList")
	purego.RegisterLibFunc(&CGColorRelease, lib, "CGColorRelease")
	purego.RegisterLibFunc(&CGColorSpaceCreateCalibratedRGB, lib, "CGColorSpaceCreateCalibratedRGB")
	purego.RegisterLibFunc(&CGColorSpaceCreatePattern, lib, "CGColorSpaceCreatePattern")
	purego.RegisterLibFunc(&CGColorSpaceIsWideGamutRGB, lib, "CGColorSpaceIsWideGamutRGB")
	purego.RegisterLibFunc(&CGColorSpaceCreateLinearized, lib, "CGColorSpaceCreateLinearized")
	purego.RegisterLibFunc(&CGColorSpaceIsHLGBased, lib, "CGColorSpaceIsHLGBased")
	purego.RegisterLibFunc(&CGColorSpaceUsesExtendedRange, lib, "CGColorSpaceUsesExtendedRange")
	purego.RegisterLibFunc(&CGContextAddPath, lib, "CGContextAddPath")
	purego.RegisterLibFunc(&CGBitmapContextGetBytesPerRow, lib, "CGBitmapContextGetBytesPerRow")
	purego.RegisterLibFunc(&CGContextDrawShading, lib, "CGContextDrawShading")
	purego.RegisterLibFunc(&CGPDFContextEndPage, lib, "CGPDFContextEndPage")
	purego.RegisterLibFunc(&CGBitmapContextCreate, lib, "CGBitmapContextCreate")
	purego.RegisterLibFunc(&CGBitmapContextCreateWithData, lib, "CGBitmapContextCreateWithData")
	purego.RegisterLibFunc(&CGContextGetInterpolationQuality, lib, "CGContextGetInterpolationQuality")
	purego.RegisterLibFunc(&CGContextRotateCTM, lib, "CGContextRotateCTM")
	purego.RegisterLibFunc(&CGContextScaleCTM, lib, "CGContextScaleCTM")
	purego.RegisterLibFunc(&CGPDFContextSetDestinationForRect, lib, "CGPDFContextSetDestinationForRect")
	purego.RegisterLibFunc(&CGContextSetShouldSubpixelPositionFonts, lib, "CGContextSetShouldSubpixelPositionFonts")
	purego.RegisterLibFunc(&CGContextSetStrokeColor, lib, "CGContextSetStrokeColor")
	purego.RegisterLibFunc(&CGContextSynchronizeAttributes, lib, "CGContextSynchronizeAttributes")
	purego.RegisterLibFunc(&CGContextGetTextMatrix, lib, "CGContextGetTextMatrix")
	purego.RegisterLibFunc(&CGContextGetContentToneMappingInfo, lib, "CGContextGetContentToneMappingInfo")
	purego.RegisterLibFunc(&CGContextMoveToPoint, lib, "CGContextMoveToPoint")
	purego.RegisterLibFunc(&CGContextSetContentToneMappingInfo, lib, "CGContextSetContentToneMappingInfo")
	purego.RegisterLibFunc(&CGContextSetLineDash, lib, "CGContextSetLineDash")
	purego.RegisterLibFunc(&CGContextStrokeLineSegments, lib, "CGContextStrokeLineSegments")
	purego.RegisterLibFunc(&CGDataConsumerCreateWithCFData, lib, "CGDataConsumerCreateWithCFData")
	purego.RegisterLibFunc(&CGDataConsumerCreateWithURL, lib, "CGDataConsumerCreateWithURL")
	purego.RegisterLibFunc(&CGDataConsumerRelease, lib, "CGDataConsumerRelease")
	purego.RegisterLibFunc(&CGDataConsumerRetain, lib, "CGDataConsumerRetain")
	purego.RegisterLibFunc(&CGDataProviderCopyData, lib, "CGDataProviderCopyData")
	purego.RegisterLibFunc(&CGDataProviderGetInfo, lib, "CGDataProviderGetInfo")
	purego.RegisterLibFunc(&CGDataProviderCreateWithData, lib, "CGDataProviderCreateWithData")
	purego.RegisterLibFunc(&CGDataProviderCreateDirect, lib, "CGDataProviderCreateDirect")
	purego.RegisterLibFunc(&CGDataProviderCreateWithFilename, lib, "CGDataProviderCreateWithFilename")
	purego.RegisterLibFunc(&CGDataProviderCreateSequential, lib, "CGDataProviderCreateSequential")
	purego.RegisterLibFunc(&CGDataProviderRetain, lib, "CGDataProviderRetain")
	purego.RegisterLibFunc(&CGDisplayCaptureWithOptions, lib, "CGDisplayCaptureWithOptions")
	purego.RegisterLibFunc(&CGDisplayIsActive, lib, "CGDisplayIsActive")
	purego.RegisterLibFunc(&CGDisplayModeGetPixelHeight, lib, "CGDisplayModeGetPixelHeight")
	purego.RegisterLibFunc(&CGDisplayRemoveReconfigurationCallback, lib, "CGDisplayRemoveReconfigurationCallback")
	purego.RegisterLibFunc(&CGDisplayStreamUpdateGetRects, lib, "CGDisplayStreamUpdateGetRects")
	purego.RegisterLibFunc(&CGEventCreateMouseEvent, lib, "CGEventCreateMouseEvent")
	purego.RegisterLibFunc(&CGEventKeyboardGetUnicodeString, lib, "CGEventKeyboardGetUnicodeString")
	purego.RegisterLibFunc(&CGEventKeyboardSetUnicodeString, lib, "CGEventKeyboardSetUnicodeString")
	purego.RegisterLibFunc(&CGEventTapCreate, lib, "CGEventTapCreate")
	purego.RegisterLibFunc(&CGEventTapCreateForPSN, lib, "CGEventTapCreateForPSN")
	purego.RegisterLibFunc(&CGEventTapPostEvent, lib, "CGEventTapPostEvent")
	purego.RegisterLibFunc(&CGEventCreateScrollWheelEvent, lib, "CGEventCreateScrollWheelEvent")
	purego.RegisterLibFunc(&CGEventSetFlags, lib, "CGEventSetFlags")
	purego.RegisterLibFunc(&CGEventSourceCounterForEventType, lib, "CGEventSourceCounterForEventType")
	purego.RegisterLibFunc(&CGEventSourceFlagsState, lib, "CGEventSourceFlagsState")
	purego.RegisterLibFunc(&CGEventSourceGetLocalEventsSuppressionInterval, lib, "CGEventSourceGetLocalEventsSuppressionInterval")
	purego.RegisterLibFunc(&CGFontGetCapHeight, lib, "CGFontGetCapHeight")
	purego.RegisterLibFunc(&CGFontCreateCopyWithVariations, lib, "CGFontCreateCopyWithVariations")
	purego.RegisterLibFunc(&CGFontCreatePostScriptEncoding, lib, "CGFontCreatePostScriptEncoding")
	purego.RegisterLibFunc(&CGFontGetDescent, lib, "CGFontGetDescent")
	purego.RegisterLibFunc(&CGFontGetGlyphWithGlyphName, lib, "CGFontGetGlyphWithGlyphName")
	purego.RegisterLibFunc(&CGFontGetLeading, lib, "CGFontGetLeading")
	purego.RegisterLibFunc(&CGFontGetStemV, lib, "CGFontGetStemV")
	purego.RegisterLibFunc(&CGFontCopyTableForTag, lib, "CGFontCopyTableForTag")
	purego.RegisterLibFunc(&CGFontGetXHeight, lib, "CGFontGetXHeight")
	purego.RegisterLibFunc(&CGFontCreateWithPlatformFont, lib, "CGFontCreateWithPlatformFont")
	purego.RegisterLibFunc(&CGFontRelease, lib, "CGFontRelease")
	purego.RegisterLibFunc(&CGFunctionCreate, lib, "CGFunctionCreate")
	purego.RegisterLibFunc(&CGFunctionRelease, lib, "CGFunctionRelease")
	purego.RegisterLibFunc(&CGFunctionRetain, lib, "CGFunctionRetain")
	purego.RegisterLibFunc(&CGGradientGetContentHeadroom, lib, "CGGradientGetContentHeadroom")
	purego.RegisterLibFunc(&CGGradientCreateWithColorComponents, lib, "CGGradientCreateWithColorComponents")
	purego.RegisterLibFunc(&CGGradientCreateWithColors, lib, "CGGradientCreateWithColors")
	purego.RegisterLibFunc(&CGGradientCreateWithContentHeadroom, lib, "CGGradientCreateWithContentHeadroom")
	purego.RegisterLibFunc(&CGGradientRelease, lib, "CGGradientRelease")
	purego.RegisterLibFunc(&CGGradientRetain, lib, "CGGradientRetain")
	purego.RegisterLibFunc(&CGImageGetAlphaInfo, lib, "CGImageGetAlphaInfo")
	purego.RegisterLibFunc(&CGImageGetBitsPerComponent, lib, "CGImageGetBitsPerComponent")
	purego.RegisterLibFunc(&CGImageGetByteOrderInfo, lib, "CGImageGetByteOrderInfo")
	purego.RegisterLibFunc(&CGImageGetBytesPerRow, lib, "CGImageGetBytesPerRow")
	purego.RegisterLibFunc(&CGImageCalculateContentAverageLightLevel, lib, "CGImageCalculateContentAverageLightLevel")
	purego.RegisterLibFunc(&CGImageCreateWithJPEGDataProvider, lib, "CGImageCreateWithJPEGDataProvider")
	purego.RegisterLibFunc(&CGImageCreate, lib, "CGImageCreate")
	purego.RegisterLibFunc(&CGWindowListCreateImageFromArray, lib, "CGWindowListCreateImageFromArray")
	purego.RegisterLibFunc(&CGImageIsMask, lib, "CGImageIsMask")
	purego.RegisterLibFunc(&CGImageCreateWithMask, lib, "CGImageCreateWithMask")
	purego.RegisterLibFunc(&CGImageRetain, lib, "CGImageRetain")
	purego.RegisterLibFunc(&CGLayerGetContext, lib, "CGLayerGetContext")
	purego.RegisterLibFunc(&CGLayerCreateWithContext, lib, "CGLayerCreateWithContext")
	purego.RegisterLibFunc(&CGLayerGetSize, lib, "CGLayerGetSize")
	purego.RegisterLibFunc(&CGLayerRelease, lib, "CGLayerRelease")
	purego.RegisterLibFunc(&CGLayerRetain, lib, "CGLayerRetain")
	purego.RegisterLibFunc(&CGPathCloseSubpath, lib, "CGPathCloseSubpath")
	purego.RegisterLibFunc(&CGPDFArrayGetBoolean, lib, "CGPDFArrayGetBoolean")
	purego.RegisterLibFunc(&CGPDFArrayGetDictionary, lib, "CGPDFArrayGetDictionary")
	purego.RegisterLibFunc(&CGPDFArrayGetInteger, lib, "CGPDFArrayGetInteger")
	purego.RegisterLibFunc(&CGPDFArrayGetNull, lib, "CGPDFArrayGetNull")
	purego.RegisterLibFunc(&CGPDFArrayGetObject, lib, "CGPDFArrayGetObject")
	purego.RegisterLibFunc(&CGPDFArrayGetString, lib, "CGPDFArrayGetString")
	purego.RegisterLibFunc(&CGPDFContentStreamGetResource, lib, "CGPDFContentStreamGetResource")
	purego.RegisterLibFunc(&CGPDFContentStreamRelease, lib, "CGPDFContentStreamRelease")
	purego.RegisterLibFunc(&CGPDFContentStreamRetain, lib, "CGPDFContentStreamRetain")
	purego.RegisterLibFunc(&CGPDFContextBeginTag, lib, "CGPDFContextBeginTag")
	purego.RegisterLibFunc(&CGPDFDictionaryApplyFunction, lib, "CGPDFDictionaryApplyFunction")
	purego.RegisterLibFunc(&CGPDFDictionaryGetArray, lib, "CGPDFDictionaryGetArray")
	purego.RegisterLibFunc(&CGPDFDictionaryGetBoolean, lib, "CGPDFDictionaryGetBoolean")
	purego.RegisterLibFunc(&CGPDFDictionaryGetCount, lib, "CGPDFDictionaryGetCount")
	purego.RegisterLibFunc(&CGPDFDictionaryGetDictionary, lib, "CGPDFDictionaryGetDictionary")
	purego.RegisterLibFunc(&CGPDFDictionaryGetObject, lib, "CGPDFDictionaryGetObject")
	purego.RegisterLibFunc(&CGPDFDictionaryGetStream, lib, "CGPDFDictionaryGetStream")
	purego.RegisterLibFunc(&CGPDFDictionaryGetString, lib, "CGPDFDictionaryGetString")
	purego.RegisterLibFunc(&CGPDFDocumentGetAccessPermissions, lib, "CGPDFDocumentGetAccessPermissions")
	purego.RegisterLibFunc(&CGPDFDocumentAllowsCopying, lib, "CGPDFDocumentAllowsCopying")
	purego.RegisterLibFunc(&CGPDFDocumentGetID, lib, "CGPDFDocumentGetID")
	purego.RegisterLibFunc(&CGPDFDocumentCreateWithProvider, lib, "CGPDFDocumentCreateWithProvider")
	purego.RegisterLibFunc(&CGPDFDocumentGetNumberOfPages, lib, "CGPDFDocumentGetNumberOfPages")
	purego.RegisterLibFunc(&CGPDFDocumentGetOutline, lib, "CGPDFDocumentGetOutline")
	purego.RegisterLibFunc(&CGPDFDocumentUnlockWithPassword, lib, "CGPDFDocumentUnlockWithPassword")
	purego.RegisterLibFunc(&CGPDFDocumentGetCropBox, lib, "CGPDFDocumentGetCropBox")
	purego.RegisterLibFunc(&CGPDFDocumentGetTrimBox, lib, "CGPDFDocumentGetTrimBox")
	purego.RegisterLibFunc(&CGPDFDocumentRetain, lib, "CGPDFDocumentRetain")
	purego.RegisterLibFunc(&CGPDFObjectGetType, lib, "CGPDFObjectGetType")
	purego.RegisterLibFunc(&CGPDFObjectGetValue, lib, "CGPDFObjectGetValue")
	purego.RegisterLibFunc(&CGPDFOperatorTableRetain, lib, "CGPDFOperatorTableRetain")
	purego.RegisterLibFunc(&CGPDFOperatorTableSetCallback, lib, "CGPDFOperatorTableSetCallback")
	purego.RegisterLibFunc(&CGPDFPageGetDictionary, lib, "CGPDFPageGetDictionary")
	purego.RegisterLibFunc(&CGPDFPageGetDocument, lib, "CGPDFPageGetDocument")
	purego.RegisterLibFunc(&CGPDFPageGetBoxRect, lib, "CGPDFPageGetBoxRect")
	purego.RegisterLibFunc(&CGPDFPageGetDrawingTransform, lib, "CGPDFPageGetDrawingTransform")
	purego.RegisterLibFunc(&CGPDFPageGetPageNumber, lib, "CGPDFPageGetPageNumber")
	purego.RegisterLibFunc(&CGPDFPageGetRotationAngle, lib, "CGPDFPageGetRotationAngle")
	purego.RegisterLibFunc(&CGPDFPageRelease, lib, "CGPDFPageRelease")
	purego.RegisterLibFunc(&CGPDFPageRetain, lib, "CGPDFPageRetain")
	purego.RegisterLibFunc(&CGPDFScannerCreate, lib, "CGPDFScannerCreate")
	purego.RegisterLibFunc(&CGPDFScannerPopArray, lib, "CGPDFScannerPopArray")
	purego.RegisterLibFunc(&CGPDFScannerPopDictionary, lib, "CGPDFScannerPopDictionary")
	purego.RegisterLibFunc(&CGPDFScannerPopName, lib, "CGPDFScannerPopName")
	purego.RegisterLibFunc(&CGPDFScannerPopNumber, lib, "CGPDFScannerPopNumber")
	purego.RegisterLibFunc(&CGPDFScannerPopString, lib, "CGPDFScannerPopString")
	purego.RegisterLibFunc(&CGPDFScannerRelease, lib, "CGPDFScannerRelease")
	purego.RegisterLibFunc(&CGPDFScannerRetain, lib, "CGPDFScannerRetain")
	purego.RegisterLibFunc(&CGPDFStreamCopyData, lib, "CGPDFStreamCopyData")
	purego.RegisterLibFunc(&CGPDFStreamGetDictionary, lib, "CGPDFStreamGetDictionary")
	purego.RegisterLibFunc(&CGPDFStringCopyDate, lib, "CGPDFStringCopyDate")
	purego.RegisterLibFunc(&CGPDFStringCopyTextString, lib, "CGPDFStringCopyTextString")
	purego.RegisterLibFunc(&CGPDFStringGetBytePtr, lib, "CGPDFStringGetBytePtr")
	purego.RegisterLibFunc(&CGPDFStringGetLength, lib, "CGPDFStringGetLength")
	purego.RegisterLibFunc(&CGPSConverterAbort, lib, "CGPSConverterAbort")
	purego.RegisterLibFunc(&CGPSConverterConvert, lib, "CGPSConverterConvert")
	purego.RegisterLibFunc(&CGPSConverterCreate, lib, "CGPSConverterCreate")
	purego.RegisterLibFunc(&CGPathGetBoundingBox, lib, "CGPathGetBoundingBox")
	purego.RegisterLibFunc(&CGPathGetCurrentPoint, lib, "CGPathGetCurrentPoint")
	purego.RegisterLibFunc(&CGPathCreateWithRect, lib, "CGPathCreateWithRect")
	purego.RegisterLibFunc(&CGPathCreateWithRoundedRect, lib, "CGPathCreateWithRoundedRect")
	purego.RegisterLibFunc(&CGPathIsRect, lib, "CGPathIsRect")
	purego.RegisterLibFunc(&CGPathCreateMutableCopy, lib, "CGPathCreateMutableCopy")
	purego.RegisterLibFunc(&CGPathAddArcToPoint, lib, "CGPathAddArcToPoint")
	purego.RegisterLibFunc(&CGPathAddCurveToPoint, lib, "CGPathAddCurveToPoint")
	purego.RegisterLibFunc(&CGPathAddLineToPoint, lib, "CGPathAddLineToPoint")
	purego.RegisterLibFunc(&CGPathAddLines, lib, "CGPathAddLines")
	purego.RegisterLibFunc(&CGPathAddPath, lib, "CGPathAddPath")
	purego.RegisterLibFunc(&CGPathAddQuadCurveToPoint, lib, "CGPathAddQuadCurveToPoint")
	purego.RegisterLibFunc(&CGPathCreateCopyByDashingPath, lib, "CGPathCreateCopyByDashingPath")
	purego.RegisterLibFunc(&CGPathCreateCopyByFlattening, lib, "CGPathCreateCopyByFlattening")
	purego.RegisterLibFunc(&CGPathCreateCopyByStrokingPath, lib, "CGPathCreateCopyByStrokingPath")
	purego.RegisterLibFunc(&CGPathCreateCopyBySymmetricDifferenceOfPath, lib, "CGPathCreateCopyBySymmetricDifferenceOfPath")
	purego.RegisterLibFunc(&CGPathMoveToPoint, lib, "CGPathMoveToPoint")
	purego.RegisterLibFunc(&CGPatternCreate, lib, "CGPatternCreate")
	purego.RegisterLibFunc(&CGPatternRelease, lib, "CGPatternRelease")
	purego.RegisterLibFunc(&CGPatternRetain, lib, "CGPatternRetain")
	purego.RegisterLibFunc(&CGPostMouseEvent, lib, "CGPostMouseEvent")
	purego.RegisterLibFunc(&CGRectApplyAffineTransform, lib, "CGRectApplyAffineTransform")
	purego.RegisterLibFunc(&CGRectEqualToRect, lib, "CGRectEqualToRect")
	purego.RegisterLibFunc(&CGRectGetMaxY, lib, "CGRectGetMaxY")
	purego.RegisterLibFunc(&CGRectGetMinX, lib, "CGRectGetMinX")
	purego.RegisterLibFunc(&CGRectGetMinY, lib, "CGRectGetMinY")
	purego.RegisterLibFunc(&CGRectIsNull, lib, "CGRectIsNull")
	purego.RegisterLibFunc(&CGRectOffset, lib, "CGRectOffset")
	purego.RegisterLibFunc(&CGRenderingBufferLockBytePtr, lib, "CGRenderingBufferLockBytePtr")
	purego.RegisterLibFunc(&CGRenderingBufferProviderCreate, lib, "CGRenderingBufferProviderCreate")
	purego.RegisterLibFunc(&CGRenderingBufferProviderCreateWithCFData, lib, "CGRenderingBufferProviderCreateWithCFData")
	purego.RegisterLibFunc(&CGRenderingBufferProviderGetSize, lib, "CGRenderingBufferProviderGetSize")
	purego.RegisterLibFunc(&CGRenderingBufferUnlockBytePtr, lib, "CGRenderingBufferUnlockBytePtr")
	purego.RegisterLibFunc(&CGShadingGetContentHeadroom, lib, "CGShadingGetContentHeadroom")
	purego.RegisterLibFunc(&CGShadingCreateAxialWithContentHeadroom, lib, "CGShadingCreateAxialWithContentHeadroom")
	purego.RegisterLibFunc(&CGShadingCreateAxial, lib, "CGShadingCreateAxial")
	purego.RegisterLibFunc(&CGShadingCreateRadialWithContentHeadroom, lib, "CGShadingCreateRadialWithContentHeadroom")
	purego.RegisterLibFunc(&CGShadingCreateRadial, lib, "CGShadingCreateRadial")
	purego.RegisterLibFunc(&CGShadingRelease, lib, "CGShadingRelease")
	purego.RegisterLibFunc(&CGShadingRetain, lib, "CGShadingRetain")
	purego.RegisterLibFunc(&CGSizeApplyAffineTransform, lib, "CGSizeApplyAffineTransform")
	purego.RegisterLibFunc(&CGSizeCreateDictionaryRepresentation, lib, "CGSizeCreateDictionaryRepresentation")
	purego.RegisterLibFunc(&CGSizeMakeWithDictionaryRepresentation, lib, "CGSizeMakeWithDictionaryRepresentation")
	purego.RegisterLibFunc(&CGWaitForScreenRefreshRects, lib, "CGWaitForScreenRefreshRects")
	purego.RegisterLibFunc(&CGWindowListCopyWindowInfo, lib, "CGWindowListCopyWindowInfo")
	purego.RegisterLibFunc(&CGWindowListCreateDescriptionFromArray, lib, "CGWindowListCreateDescriptionFromArray")
}
