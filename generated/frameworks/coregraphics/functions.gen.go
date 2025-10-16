// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// CoreGraphics Functions (727 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CGColorSpaceRelease                                     func(CGColorSpaceRef)
	_CGContextSetInterpolationQuality                        func(CGContextRef, unsafe.Pointer)
	_CGAcquireDisplayFadeReservation                         func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGAffineTransformConcat                                 func(CGAffineTransform, CGAffineTransform) CGAffineTransform
	_CGAffineTransformDecompose                              func(CGAffineTransform) unsafe.Pointer
	_CGAffineTransformEqualToTransform                       func(CGAffineTransform, CGAffineTransform) bool
	_CGAffineTransformInvert                                 func(CGAffineTransform) CGAffineTransform
	_CGAffineTransformIsIdentity                             func(CGAffineTransform) bool
	_CGAffineTransformMake                                   func(CGFloat, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat) CGAffineTransform
	_CGAffineTransformMakeRotation                           func(CGFloat) CGAffineTransform
	_CGAffineTransformMakeScale                              func(CGFloat, CGFloat) CGAffineTransform
	_CGAffineTransformMakeTranslation                        func(CGFloat, CGFloat) CGAffineTransform
	_CGAffineTransformMakeWithComponents                     func(unsafe.Pointer) CGAffineTransform
	_CGAffineTransformRotate                                 func(CGAffineTransform, CGFloat) CGAffineTransform
	_CGAffineTransformScale                                  func(CGAffineTransform, CGFloat, CGFloat) CGAffineTransform
	_CGAffineTransformTranslate                              func(CGAffineTransform, CGFloat, CGFloat) CGAffineTransform
	_CGAssociateMouseAndMouseCursorPosition                  func(unsafe.Pointer) unsafe.Pointer
	_CGBeginDisplayConfiguration                             func(unsafe.Pointer) unsafe.Pointer
	_CGBitmapContextCreateAdaptive                           func(uintptr, uintptr, unsafe.Pointer, bool) CGContextRef
	_CGCancelDisplayConfiguration                            func(CGDisplayConfigRef) unsafe.Pointer
	_CGCaptureAllDisplays                                    func() unsafe.Pointer
	_CGCaptureAllDisplaysWithOptions                         func(unsafe.Pointer) unsafe.Pointer
	_CGColorGetAlpha                                         func(CGColorRef) CGFloat
	_CGColorGetColorSpace                                    func(CGColorRef) CGColorSpaceRef
	_CGColorGetContentHeadroom                               func(CGColorRef) float32
	_CGColorCreateCopyByMatchingToColorSpace                 func(CGColorSpaceRef, unsafe.Pointer, CGColorRef, unsafe.Pointer) CGColorRef
	_CGColorCreateCopy                                       func(CGColorRef) CGColorRef
	_CGColorCreateCopyWithAlpha                              func(CGColorRef, CGFloat) CGColorRef
	_CGColorCreate                                           func(CGColorSpaceRef, unsafe.Pointer) CGColorRef
	_CGColorCreateGenericCMYK                                func(CGFloat, CGFloat, CGFloat, CGFloat, CGFloat) CGColorRef
	_CGColorCreateGenericGrayGamma2_2                        func(CGFloat, CGFloat) CGColorRef
	_CGColorCreateGenericGray                                func(CGFloat, CGFloat) CGColorRef
	_CGColorCreateWithContentHeadroom                        func(float32, CGColorSpaceRef, CGFloat, CGFloat, CGFloat, CGFloat) CGColorRef
	_CGColorCreateWithPattern                                func(CGColorSpaceRef, CGPatternRef, unsafe.Pointer) CGColorRef
	_CGColorCreateGenericRGB                                 func(CGFloat, CGFloat, CGFloat, CGFloat) CGColorRef
	_CGColorCreateSRGB                                       func(CGFloat, CGFloat, CGFloat, CGFloat) CGColorRef
	_CGColorGetNumberOfComponents                            func(CGColorRef) uintptr
	_CGColorGetPattern                                       func(CGColorRef) CGPatternRef
	_CGColorGetTypeID                                        func() unsafe.Pointer
	_CGColorConversionInfoConvertData                        func(CGColorConversionInfoRef, uintptr, uintptr, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_CGColorConversionInfoCreateWithOptions                  func(CGColorSpaceRef, CGColorSpaceRef, unsafe.Pointer) CGColorConversionInfoRef
	_CGColorConversionInfoCreate                             func(CGColorSpaceRef, CGColorSpaceRef) CGColorConversionInfoRef
	_CGColorConversionInfoCreateForToneMapping               func(CGColorSpaceRef, float32, CGColorSpaceRef, float32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGColorConversionInfoRef
	_CGColorConversionInfoGetTypeID                          func() unsafe.Pointer
	_CGColorConversionInfoCreateFromList                     func(unsafe.Pointer, CGColorSpaceRef, unsafe.Pointer, unsafe.Pointer) CGColorConversionInfoRef
	_CGColorConversionInfoCreateFromListWithArguments        func(unsafe.Pointer, CGColorSpaceRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGColorConversionInfoRef
	_CGColorEqualToColor                                     func(CGColorRef, CGColorRef) bool
	_CGColorGetComponents                                    func(CGColorRef) unsafe.Pointer
	_CGColorGetConstantColor                                 func(unsafe.Pointer) CGColorRef
	_CGColorRelease                                          func(CGColorRef)
	_CGColorRetain                                           func(CGColorRef) CGColorRef
	_CGColorSpaceGetBaseColorSpace                           func(CGColorSpaceRef) CGColorSpaceRef
	_CGColorSpaceCopyICCData                                 func(CGColorSpaceRef) unsafe.Pointer
	_CGColorSpaceCopyPropertyList                            func(CGColorSpaceRef) unsafe.Pointer
	_CGColorSpaceCopyICCProfile                              func(CGColorSpaceRef) unsafe.Pointer
	_CGColorSpaceCreateCalibratedGray                        func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, CGFloat) CGColorSpaceRef
	_CGColorSpaceCreateCalibratedRGB                         func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGColorSpaceRef
	_CGColorSpaceCreateICCBased                              func(uintptr, unsafe.Pointer, CGDataProviderRef, CGColorSpaceRef) CGColorSpaceRef
	_CGColorSpaceCreateWithICCData                           func(unsafe.Pointer) CGColorSpaceRef
	_CGColorSpaceCreateWithICCProfile                        func(unsafe.Pointer) CGColorSpaceRef
	_CGColorSpaceCreateIndexed                               func(CGColorSpaceRef, uintptr, unsafe.Pointer) CGColorSpaceRef
	_CGColorSpaceCreateLab                                   func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGColorSpaceRef
	_CGColorSpaceCreateWithName                              func(unsafe.Pointer) CGColorSpaceRef
	_CGColorSpaceCreatePattern                               func(CGColorSpaceRef) CGColorSpaceRef
	_CGColorSpaceCreateWithPlatformColorSpace                func(unsafe.Pointer) CGColorSpaceRef
	_CGColorSpaceCreateWithPropertyList                      func(unsafe.Pointer) CGColorSpaceRef
	_CGColorSpaceIsHDR                                       func(CGColorSpaceRef) bool
	_CGColorSpaceIsWideGamutRGB                              func(CGColorSpaceRef) bool
	_CGColorSpaceGetModel                                    func(CGColorSpaceRef) unsafe.Pointer
	_CGColorSpaceCopyName                                    func(CGColorSpaceRef) unsafe.Pointer
	_CGColorSpaceGetNumberOfComponents                       func(CGColorSpaceRef) uintptr
	_CGColorSpaceSupportsOutput                              func(CGColorSpaceRef) bool
	_CGColorSpaceGetTypeID                                   func() unsafe.Pointer
	_CGColorSpaceCopyBaseColorSpace                          func(CGColorSpaceRef) CGColorSpaceRef
	_CGColorSpaceCreateCopyWithStandardRange                 func(CGColorSpaceRef) CGColorSpaceRef
	_CGColorSpaceCreateDeviceCMYK                            func() CGColorSpaceRef
	_CGColorSpaceCreateDeviceGray                            func() CGColorSpaceRef
	_CGColorSpaceCreateDeviceRGB                             func() CGColorSpaceRef
	_CGColorSpaceCreateExtended                              func(CGColorSpaceRef) CGColorSpaceRef
	_CGColorSpaceCreateExtendedLinearized                    func(CGColorSpaceRef) CGColorSpaceRef
	_CGColorSpaceCreateLinearized                            func(CGColorSpaceRef) CGColorSpaceRef
	_CGColorSpaceCreateWithColorSyncProfile                  func(unsafe.Pointer, unsafe.Pointer) CGColorSpaceRef
	_CGColorSpaceGetColorTable                               func(CGColorSpaceRef, unsafe.Pointer)
	_CGColorSpaceGetColorTableCount                          func(CGColorSpaceRef) uintptr
	_CGColorSpaceGetName                                     func(CGColorSpaceRef) unsafe.Pointer
	_CGColorSpaceIsHLGBased                                  func(CGColorSpaceRef) bool
	_CGColorSpaceIsPQBased                                   func(CGColorSpaceRef) bool
	_CGColorSpaceRetain                                      func(CGColorSpaceRef) CGColorSpaceRef
	_CGColorSpaceUsesExtendedRange                           func(CGColorSpaceRef) bool
	_CGColorSpaceUsesITUR_2100TF                             func(CGColorSpaceRef) bool
	_CGCompleteDisplayConfiguration                          func(CGDisplayConfigRef, unsafe.Pointer) unsafe.Pointer
	_CGConfigureDisplayFadeEffect                            func(CGDisplayConfigRef, unsafe.Pointer, unsafe.Pointer, float32, float32, float32) unsafe.Pointer
	_CGConfigureDisplayMirrorOfDisplay                       func(CGDisplayConfigRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGConfigureDisplayMode                                  func(CGDisplayConfigRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGConfigureDisplayOrigin                                func(CGDisplayConfigRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGConfigureDisplayStereoOperation                       func(CGDisplayConfigRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGConfigureDisplayWithDisplayMode                       func(CGDisplayConfigRef, unsafe.Pointer, CGDisplayModeRef, unsafe.Pointer) unsafe.Pointer
	_CGPDFContextAddDestinationAtPoint                       func(CGContextRef, unsafe.Pointer, CGPoint)
	_CGPDFContextAddDocumentMetadata                         func(CGContextRef, unsafe.Pointer)
	_CGContextAddEllipseInRect                               func(CGContextRef, CGRect)
	_CGContextAddPath                                        func(CGContextRef, CGPathRef)
	_CGContextAddRect                                        func(CGContextRef, CGRect)
	_CGBitmapContextGetAlphaInfo                             func(CGContextRef) unsafe.Pointer
	_CGPDFContextBeginPage                                   func(CGContextRef, unsafe.Pointer)
	_CGContextBeginPage                                      func(CGContextRef, unsafe.Pointer)
	_CGContextBeginPath                                      func(CGContextRef)
	_CGContextBeginTransparencyLayer                         func(CGContextRef, unsafe.Pointer)
	_CGContextBeginTransparencyLayerWithRect                 func(CGContextRef, CGRect, unsafe.Pointer)
	_CGBitmapContextGetBitmapInfo                            func(CGContextRef) unsafe.Pointer
	_CGBitmapContextGetBitsPerComponent                      func(CGContextRef) uintptr
	_CGBitmapContextGetBitsPerPixel                          func(CGContextRef) uintptr
	_CGContextGetClipBoundingBox                             func(CGContextRef) CGRect
	_CGContextGetPathBoundingBox                             func(CGContextRef) CGRect
	_CGBitmapContextGetBytesPerRow                           func(CGContextRef) uintptr
	_CGContextClearRect                                      func(CGContextRef, CGRect)
	_CGContextClipToRect                                     func(CGContextRef, CGRect)
	_CGContextClipToMask                                     func(CGContextRef, CGRect, CGImageRef)
	_CGPDFContextClose                                       func(CGContextRef)
	_CGContextClosePath                                      func(CGContextRef)
	_CGBitmapContextGetColorSpace                            func(CGContextRef) CGColorSpaceRef
	_CGContextConcatCTM                                      func(CGContextRef, CGAffineTransform)
	_CGContextConvertSizeToDeviceSpace                       func(CGContextRef, CGSize) CGSize
	_CGContextConvertPointToDeviceSpace                      func(CGContextRef, CGPoint) CGPoint
	_CGContextConvertRectToDeviceSpace                       func(CGContextRef, CGRect) CGRect
	_CGContextConvertRectToUserSpace                         func(CGContextRef, CGRect) CGRect
	_CGContextConvertPointToUserSpace                        func(CGContextRef, CGPoint) CGPoint
	_CGContextConvertSizeToUserSpace                         func(CGContextRef, CGSize) CGSize
	_CGContextGetCTM                                         func(CGContextRef) CGAffineTransform
	_CGContextGetPathCurrentPoint                            func(CGContextRef) CGPoint
	_CGBitmapContextGetData                                  func(CGContextRef) unsafe.Pointer
	_CGContextDrawLinearGradient                             func(CGContextRef, CGGradientRef, CGPoint, CGPoint, unsafe.Pointer)
	_CGContextDrawPDFPage                                    func(CGContextRef, CGPDFPageRef)
	_CGContextDrawPath                                       func(CGContextRef, unsafe.Pointer)
	_CGContextDrawRadialGradient                             func(CGContextRef, CGGradientRef, CGPoint, CGFloat, CGPoint, CGFloat, unsafe.Pointer)
	_CGContextDrawShading                                    func(CGContextRef, CGShadingRef)
	_CGPDFContextEndPage                                     func(CGContextRef)
	_CGContextEndPage                                        func(CGContextRef)
	_CGContextEndTransparencyLayer                           func(CGContextRef)
	_CGContextFillRect                                       func(CGContextRef, CGRect)
	_CGContextFillEllipseInRect                              func(CGContextRef, CGRect)
	_CGContextFlush                                          func(CGContextRef)
	_CGBitmapContextGetHeight                                func(CGContextRef) uintptr
	_CGPDFContextCreateWithURL                               func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGContextRef
	_CGPDFContextCreate                                      func(CGDataConsumerRef, unsafe.Pointer, unsafe.Pointer) CGContextRef
	_CGBitmapContextCreate                                   func(unsafe.Pointer, uintptr, uintptr, uintptr, uintptr, CGColorSpaceRef, unsafe.Pointer) CGContextRef
	_CGBitmapContextCreateWithData                           func(unsafe.Pointer, uintptr, uintptr, uintptr, uintptr, CGColorSpaceRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGContextRef
	_CGContextGetInterpolationQuality                        func(CGContextRef) unsafe.Pointer
	_CGContextIsPathEmpty                                    func(CGContextRef) bool
	_CGBitmapContextCreateImage                              func(CGContextRef) CGImageRef
	_CGContextCopyPath                                       func(CGContextRef) CGPathRef
	_CGContextPathContainsPoint                              func(CGContextRef, CGPoint, unsafe.Pointer) bool
	_CGContextReplacePathWithStrokedPath                     func(CGContextRef)
	_CGContextResetClip                                      func(CGContextRef)
	_CGContextRestoreGState                                  func(CGContextRef)
	_CGContextRotateCTM                                      func(CGContextRef, CGFloat)
	_CGContextSaveGState                                     func(CGContextRef)
	_CGContextScaleCTM                                       func(CGContextRef, CGFloat, CGFloat)
	_CGContextSelectFont                                     func(CGContextRef, unsafe.Pointer, CGFloat, unsafe.Pointer)
	_CGContextSetAllowsAntialiasing                          func(CGContextRef, bool)
	_CGContextSetAllowsFontSmoothing                         func(CGContextRef, bool)
	_CGContextSetAllowsFontSubpixelPositioning               func(CGContextRef, bool)
	_CGContextSetAllowsFontSubpixelQuantization              func(CGContextRef, bool)
	_CGContextSetAlpha                                       func(CGContextRef, CGFloat)
	_CGContextSetBlendMode                                   func(CGContextRef, unsafe.Pointer)
	_CGContextSetCharacterSpacing                            func(CGContextRef, CGFloat)
	_CGPDFContextSetDestinationForRect                       func(CGContextRef, unsafe.Pointer, CGRect)
	_CGContextSetEDRTargetHeadroom                           func(CGContextRef, float32) bool
	_CGContextSetFillColor                                   func(CGContextRef, unsafe.Pointer)
	_CGContextSetFillColorWithColor                          func(CGContextRef, CGColorRef)
	_CGContextSetCMYKFillColor                               func(CGContextRef, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat)
	_CGContextSetGrayFillColor                               func(CGContextRef, CGFloat, CGFloat)
	_CGContextSetRGBFillColor                                func(CGContextRef, CGFloat, CGFloat, CGFloat, CGFloat)
	_CGContextSetFillColorSpace                              func(CGContextRef, CGColorSpaceRef)
	_CGContextSetFillPattern                                 func(CGContextRef, CGPatternRef, unsafe.Pointer)
	_CGContextSetFlatness                                    func(CGContextRef, CGFloat)
	_CGContextSetFont                                        func(CGContextRef, CGFontRef)
	_CGContextSetFontSize                                    func(CGContextRef, CGFloat)
	_CGContextSetLineCap                                     func(CGContextRef, unsafe.Pointer)
	_CGContextSetLineJoin                                    func(CGContextRef, unsafe.Pointer)
	_CGContextSetLineWidth                                   func(CGContextRef, CGFloat)
	_CGContextSetMiterLimit                                  func(CGContextRef, CGFloat)
	_CGContextSetPatternPhase                                func(CGContextRef, CGSize)
	_CGContextSetRenderingIntent                             func(CGContextRef, unsafe.Pointer)
	_CGContextSetShadow                                      func(CGContextRef, CGSize, CGFloat)
	_CGContextSetShadowWithColor                             func(CGContextRef, CGSize, CGFloat, CGColorRef)
	_CGContextSetShouldAntialias                             func(CGContextRef, bool)
	_CGContextSetShouldSmoothFonts                           func(CGContextRef, bool)
	_CGContextSetShouldSubpixelPositionFonts                 func(CGContextRef, bool)
	_CGContextSetShouldSubpixelQuantizeFonts                 func(CGContextRef, bool)
	_CGContextSetStrokeColorWithColor                        func(CGContextRef, CGColorRef)
	_CGContextSetStrokeColor                                 func(CGContextRef, unsafe.Pointer)
	_CGContextSetCMYKStrokeColor                             func(CGContextRef, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat)
	_CGContextSetGrayStrokeColor                             func(CGContextRef, CGFloat, CGFloat)
	_CGContextSetRGBStrokeColor                              func(CGContextRef, CGFloat, CGFloat, CGFloat, CGFloat)
	_CGContextSetStrokeColorSpace                            func(CGContextRef, CGColorSpaceRef)
	_CGContextSetStrokePattern                               func(CGContextRef, CGPatternRef, unsafe.Pointer)
	_CGContextSetTextDrawingMode                             func(CGContextRef, unsafe.Pointer)
	_CGPDFContextSetURLForRect                               func(CGContextRef, unsafe.Pointer, CGRect)
	_CGContextShowGlyphs                                     func(CGContextRef, unsafe.Pointer, uintptr)
	_CGContextShowGlyphsAtPoint                              func(CGContextRef, CGFloat, CGFloat, unsafe.Pointer, uintptr)
	_CGContextShowGlyphsWithAdvances                         func(CGContextRef, unsafe.Pointer, unsafe.Pointer, uintptr)
	_CGContextShowText                                       func(CGContextRef, unsafe.Pointer, uintptr)
	_CGContextShowTextAtPoint                                func(CGContextRef, CGFloat, CGFloat, unsafe.Pointer, uintptr)
	_CGContextStrokeRect                                     func(CGContextRef, CGRect)
	_CGContextStrokeRectWithWidth                            func(CGContextRef, CGRect, CGFloat)
	_CGContextStrokeEllipseInRect                            func(CGContextRef, CGRect)
	_CGContextStrokePath                                     func(CGContextRef)
	_CGContextSynchronize                                    func(CGContextRef)
	_CGContextSynchronizeAttributes                          func(CGContextRef)
	_CGContextGetTextMatrix                                  func(CGContextRef) CGAffineTransform
	_CGContextTranslateCTM                                   func(CGContextRef, CGFloat, CGFloat)
	_CGContextGetTypeID                                      func() unsafe.Pointer
	_CGContextGetUserSpaceToDeviceSpaceTransform             func(CGContextRef) CGAffineTransform
	_CGBitmapContextGetWidth                                 func(CGContextRef) uintptr
	_CGContextAddArc                                         func(CGContextRef, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat, int)
	_CGContextAddArcToPoint                                  func(CGContextRef, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat)
	_CGContextAddCurveToPoint                                func(CGContextRef, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat)
	_CGContextAddLineToPoint                                 func(CGContextRef, CGFloat, CGFloat)
	_CGContextAddLines                                       func(CGContextRef, unsafe.Pointer, uintptr)
	_CGContextAddQuadCurveToPoint                            func(CGContextRef, CGFloat, CGFloat, CGFloat, CGFloat)
	_CGContextAddRects                                       func(CGContextRef, unsafe.Pointer, uintptr)
	_CGContextClip                                           func(CGContextRef)
	_CGContextClipToRects                                    func(CGContextRef, unsafe.Pointer, uintptr)
	_CGContextDrawConicGradient                              func(CGContextRef, CGGradientRef, CGPoint, CGFloat)
	_CGContextDrawImage                                      func(CGContextRef, CGRect, CGImageRef)
	_CGContextDrawImageApplyingToneMapping                   func(CGContextRef, CGRect, CGImageRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGContextDrawLayerAtPoint                               func(CGContextRef, CGPoint, CGLayerRef)
	_CGContextDrawLayerInRect                                func(CGContextRef, CGRect, CGLayerRef)
	_CGContextDrawPDFDocument                                func(CGContextRef, CGRect, CGPDFDocumentRef, int)
	_CGContextDrawTiledImage                                 func(CGContextRef, CGRect, CGImageRef)
	_CGContextEOClip                                         func(CGContextRef)
	_CGContextEOFillPath                                     func(CGContextRef)
	_CGContextFillPath                                       func(CGContextRef)
	_CGContextFillRects                                      func(CGContextRef, unsafe.Pointer, uintptr)
	_CGContextGetContentToneMappingInfo                      func(CGContextRef) unsafe.Pointer
	_CGContextGetEDRTargetHeadroom                           func(CGContextRef) float32
	_CGContextGetTextPosition                                func(CGContextRef) CGPoint
	_CGContextMoveToPoint                                    func(CGContextRef, CGFloat, CGFloat)
	_CGContextRelease                                        func(CGContextRef)
	_CGContextRetain                                         func(CGContextRef) CGContextRef
	_CGContextSetContentToneMappingInfo                      func(CGContextRef, unsafe.Pointer)
	_CGContextSetLineDash                                    func(CGContextRef, CGFloat, unsafe.Pointer, uintptr)
	_CGContextSetTextMatrix                                  func(CGContextRef, CGAffineTransform)
	_CGContextSetTextPosition                                func(CGContextRef, CGFloat, CGFloat)
	_CGContextShowGlyphsAtPositions                          func(CGContextRef, unsafe.Pointer, unsafe.Pointer, uintptr)
	_CGContextStrokeLineSegments                             func(CGContextRef, unsafe.Pointer, uintptr)
	_CGConvertColorDataWithFormat                            func(uintptr, uintptr, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_CGCursorIsDrawnInFramebuffer                            func() unsafe.Pointer
	_CGCursorIsVisible                                       func() unsafe.Pointer
	_CGDataConsumerCreateWithCFData                          func(unsafe.Pointer) CGDataConsumerRef
	_CGDataConsumerCreate                                    func(unsafe.Pointer, unsafe.Pointer) CGDataConsumerRef
	_CGDataConsumerCreateWithURL                             func(unsafe.Pointer) CGDataConsumerRef
	_CGDataConsumerGetTypeID                                 func() unsafe.Pointer
	_CGDataConsumerRelease                                   func(CGDataConsumerRef)
	_CGDataConsumerRetain                                    func(CGDataConsumerRef) CGDataConsumerRef
	_CGDataProviderCopyData                                  func(CGDataProviderRef) unsafe.Pointer
	_CGDataProviderGetInfo                                   func(CGDataProviderRef) unsafe.Pointer
	_CGDataProviderCreateWithCFData                          func(unsafe.Pointer) CGDataProviderRef
	_CGDataProviderCreateWithData                            func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer) CGDataProviderRef
	_CGDataProviderCreateDirect                              func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGDataProviderRef
	_CGDataProviderCreateWithFilename                        func(unsafe.Pointer) CGDataProviderRef
	_CGDataProviderCreateSequential                          func(unsafe.Pointer, unsafe.Pointer) CGDataProviderRef
	_CGDataProviderCreateWithURL                             func(unsafe.Pointer) CGDataProviderRef
	_CGDataProviderGetTypeID                                 func() unsafe.Pointer
	_CGDataProviderRelease                                   func(CGDataProviderRef)
	_CGDataProviderRetain                                    func(CGDataProviderRef) CGDataProviderRef
	_CGDirectDisplayCopyCurrentMetalDevice                   func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayAvailableModes                                 func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayBestModeForParameters                          func(unsafe.Pointer, uintptr, uintptr, uintptr, unsafe.Pointer) unsafe.Pointer
	_CGDisplayBestModeForParametersAndRefreshRate            func(unsafe.Pointer, uintptr, uintptr, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGDisplayBounds                                         func(unsafe.Pointer) CGRect
	_CGDisplayCapture                                        func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayCaptureWithOptions                             func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGDisplayCopyAllDisplayModes                            func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGDisplayCopyColorSpace                                 func(unsafe.Pointer) CGColorSpaceRef
	_CGDisplayCopyDisplayMode                                func(unsafe.Pointer) CGDisplayModeRef
	_CGDisplayCreateImage                                    func(unsafe.Pointer) CGImageRef
	_CGDisplayCreateImageForRect                             func(unsafe.Pointer, CGRect) CGImageRef
	_CGDisplayCurrentMode                                    func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayFade                                           func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, float32, float32, float32, unsafe.Pointer) unsafe.Pointer
	_CGDisplayFadeOperationInProgress                        func() unsafe.Pointer
	_CGDisplayGammaTableCapacity                             func(unsafe.Pointer) uint32
	_CGDisplayGetDrawingContext                              func(unsafe.Pointer) CGContextRef
	_CGDisplayHideCursor                                     func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIDToOpenGLDisplayMask                          func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIOServicePort                                  func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIsActive                                       func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIsAlwaysInMirrorSet                            func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIsAsleep                                       func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIsBuiltin                                      func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIsCaptured                                     func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIsInHWMirrorSet                                func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIsInMirrorSet                                  func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIsMain                                         func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIsOnline                                       func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayIsStereo                                       func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayMirrorsDisplay                                 func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayModeGetHeight                                  func(CGDisplayModeRef) uintptr
	_CGDisplayModeGetIODisplayModeID                         func(CGDisplayModeRef) unsafe.Pointer
	_CGDisplayModeGetIOFlags                                 func(CGDisplayModeRef) uint32
	_CGDisplayModeIsUsableForDesktopGUI                      func(CGDisplayModeRef) bool
	_CGDisplayModeCopyPixelEncoding                          func(CGDisplayModeRef) unsafe.Pointer
	_CGDisplayModeGetPixelHeight                             func(CGDisplayModeRef) uintptr
	_CGDisplayModeGetPixelWidth                              func(CGDisplayModeRef) uintptr
	_CGDisplayModeGetRefreshRate                             func(CGDisplayModeRef) float64
	_CGDisplayModeGetTypeID                                  func() unsafe.Pointer
	_CGDisplayModeGetWidth                                   func(CGDisplayModeRef) uintptr
	_CGDisplayModeRelease                                    func(CGDisplayModeRef)
	_CGDisplayModeRetain                                     func(CGDisplayModeRef) CGDisplayModeRef
	_CGDisplayModelNumber                                    func(unsafe.Pointer) uint32
	_CGDisplayMoveCursorToPoint                              func(unsafe.Pointer, CGPoint) unsafe.Pointer
	_CGDisplayPixelsHigh                                     func(unsafe.Pointer) uintptr
	_CGDisplayPixelsWide                                     func(unsafe.Pointer) uintptr
	_CGDisplayPrimaryDisplay                                 func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayRegisterReconfigurationCallback                func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGDisplayRelease                                        func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayRemoveReconfigurationCallback                  func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGDisplayRestoreColorSyncSettings                       func()
	_CGDisplayRotation                                       func(unsafe.Pointer) float64
	_CGDisplayScreenSize                                     func(unsafe.Pointer) CGSize
	_CGDisplaySerialNumber                                   func(unsafe.Pointer) uint32
	_CGDisplaySetDisplayMode                                 func(unsafe.Pointer, CGDisplayModeRef, unsafe.Pointer) unsafe.Pointer
	_CGDisplaySetStereoOperation                             func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGDisplayShowCursor                                     func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayStreamCreateWithDispatchQueue                  func(unsafe.Pointer, uintptr, uintptr, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGDisplayStreamRef
	_CGDisplayStreamCreate                                   func(unsafe.Pointer, uintptr, uintptr, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGDisplayStreamRef
	_CGDisplayStreamGetRunLoopSource                         func(CGDisplayStreamRef) unsafe.Pointer
	_CGDisplayStreamStart                                    func(CGDisplayStreamRef) unsafe.Pointer
	_CGDisplayStreamStop                                     func(CGDisplayStreamRef) unsafe.Pointer
	_CGDisplayStreamGetTypeID                                func() unsafe.Pointer
	_CGDisplayStreamUpdateGetDropCount                       func(CGDisplayStreamUpdateRef) uintptr
	_CGDisplayStreamUpdateGetMovedRectsDelta                 func(CGDisplayStreamUpdateRef, unsafe.Pointer, unsafe.Pointer)
	_CGDisplayStreamUpdateGetRects                           func(CGDisplayStreamUpdateRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGDisplayStreamUpdateCreateMergedUpdate                 func(CGDisplayStreamUpdateRef, CGDisplayStreamUpdateRef) CGDisplayStreamUpdateRef
	_CGDisplayStreamUpdateGetTypeID                          func() unsafe.Pointer
	_CGDisplaySwitchToMode                                   func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGDisplayUnitNumber                                     func(unsafe.Pointer) uint32
	_CGDisplayUsesOpenGLAcceleration                         func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayVendorNumber                                   func(unsafe.Pointer) uint32
	_CGEXRToneMappingGammaGetDefaultOptions                  func() unsafe.Pointer
	_CGEnableEventStateCombining                             func(unsafe.Pointer) unsafe.Pointer
	_CGErrorSetCallback                                      func(unsafe.Pointer)
	_CGEventCreateCopy                                       func(CGEventRef) CGEventRef
	_CGEventGetFlags                                         func(CGEventRef) unsafe.Pointer
	_CGEventGetDoubleValueField                              func(CGEventRef, unsafe.Pointer) float64
	_CGEventGetIntegerValueField                             func(CGEventRef, unsafe.Pointer) unsafe.Pointer
	_CGEventCreateKeyboardEvent                              func(CGEventSourceRef, unsafe.Pointer, bool) CGEventRef
	_CGEventCreateMouseEvent                                 func(CGEventSourceRef, unsafe.Pointer, CGPoint, unsafe.Pointer) CGEventRef
	_CGEventCreateScrollWheelEvent2                          func(CGEventSourceRef, unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGEventRef
	_CGEventCreate                                           func(CGEventSourceRef) CGEventRef
	_CGEventCreateFromData                                   func(unsafe.Pointer, unsafe.Pointer) CGEventRef
	_CGEventKeyboardGetUnicodeString                         func(CGEventRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CGEventKeyboardSetUnicodeString                         func(CGEventRef, unsafe.Pointer, unsafe.Pointer)
	_CGEventGetLocation                                      func(CGEventRef) CGPoint
	_CGEventPost                                             func(unsafe.Pointer, CGEventRef)
	_CGEventPostToPSN                                        func(unsafe.Pointer, CGEventRef)
	_CGEventPostToPid                                        func(unsafe.Pointer, CGEventRef)
	_CGEventSetDoubleValueField                              func(CGEventRef, unsafe.Pointer, float64)
	_CGEventSetIntegerValueField                             func(CGEventRef, unsafe.Pointer, unsafe.Pointer)
	_CGEventSetSource                                        func(CGEventRef, CGEventSourceRef)
	_CGEventTapCreate                                        func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGEventTapCreateForPSN                                  func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGEventTapCreateForPid                                  func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGEventTapEnable                                        func(unsafe.Pointer, bool)
	_CGEventTapIsEnabled                                     func(unsafe.Pointer) bool
	_CGEventTapPostEvent                                     func(unsafe.Pointer, CGEventRef)
	_CGEventGetTimestamp                                     func(CGEventRef) unsafe.Pointer
	_CGEventGetType                                          func(CGEventRef) unsafe.Pointer
	_CGEventGetTypeID                                        func() unsafe.Pointer
	_CGEventGetUnflippedLocation                             func(CGEventRef) CGPoint
	_CGEventCreateData                                       func(unsafe.Pointer, CGEventRef) unsafe.Pointer
	_CGEventCreateScrollWheelEvent                           func(CGEventSourceRef, unsafe.Pointer, uint32, unsafe.Pointer) CGEventRef
	_CGEventSetFlags                                         func(CGEventRef, unsafe.Pointer)
	_CGEventSetLocation                                      func(CGEventRef, CGPoint)
	_CGEventSetTimestamp                                     func(CGEventRef, unsafe.Pointer)
	_CGEventSetType                                          func(CGEventRef, unsafe.Pointer)
	_CGEventSourceButtonState                                func(unsafe.Pointer, unsafe.Pointer) bool
	_CGEventSourceCounterForEventType                        func(unsafe.Pointer, unsafe.Pointer) uint32
	_CGEventSourceFlagsState                                 func(unsafe.Pointer) unsafe.Pointer
	_CGEventSourceGetLocalEventsFilterDuringSuppressionState func(CGEventSourceRef, unsafe.Pointer) unsafe.Pointer
	_CGEventCreateSourceFromEvent                            func(CGEventRef) CGEventSourceRef
	_CGEventSourceCreate                                     func(unsafe.Pointer) CGEventSourceRef
	_CGEventSourceKeyState                                   func(unsafe.Pointer, unsafe.Pointer) bool
	_CGEventSourceGetKeyboardType                            func(CGEventSourceRef) unsafe.Pointer
	_CGEventSourceGetLocalEventsSuppressionInterval          func(CGEventSourceRef) unsafe.Pointer
	_CGEventSourceGetPixelsPerLine                           func(CGEventSourceRef) float64
	_CGEventSourceSecondsSinceLastEventType                  func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGEventSourceSetLocalEventsFilterDuringSuppressionState func(CGEventSourceRef, unsafe.Pointer, unsafe.Pointer)
	_CGEventSourceGetSourceStateID                           func(CGEventSourceRef) unsafe.Pointer
	_CGEventSourceGetTypeID                                  func() unsafe.Pointer
	_CGEventSourceGetUserData                                func(CGEventSourceRef) unsafe.Pointer
	_CGEventSourceSetKeyboardType                            func(CGEventSourceRef, unsafe.Pointer)
	_CGEventSourceSetLocalEventsSuppressionInterval          func(CGEventSourceRef, unsafe.Pointer)
	_CGEventSourceSetPixelsPerLine                           func(CGEventSourceRef, float64)
	_CGEventSourceSetUserData                                func(CGEventSourceRef, unsafe.Pointer)
	_CGFontGetAscent                                         func(CGFontRef) int
	_CGFontCanCreatePostScriptSubset                         func(CGFontRef, unsafe.Pointer) bool
	_CGFontGetCapHeight                                      func(CGFontRef) int
	_CGFontCreateCopyWithVariations                          func(CGFontRef, unsafe.Pointer) CGFontRef
	_CGFontCreatePostScriptEncoding                          func(CGFontRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGFontCreatePostScriptSubset                            func(CGFontRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGFontGetDescent                                        func(CGFontRef) int
	_CGFontGetFontBBox                                       func(CGFontRef) CGRect
	_CGFontCopyFullName                                      func(CGFontRef) unsafe.Pointer
	_CGFontGetGlyphAdvances                                  func(CGFontRef, unsafe.Pointer, uintptr, unsafe.Pointer) bool
	_CGFontGetGlyphBBoxes                                    func(CGFontRef, unsafe.Pointer, uintptr, unsafe.Pointer) bool
	_CGFontGetGlyphWithGlyphName                             func(CGFontRef, unsafe.Pointer) unsafe.Pointer
	_CGFontCreateWithFontName                                func(unsafe.Pointer) CGFontRef
	_CGFontCreateWithDataProvider                            func(CGDataProviderRef) CGFontRef
	_CGFontGetItalicAngle                                    func(CGFontRef) CGFloat
	_CGFontGetLeading                                        func(CGFontRef) int
	_CGFontCopyGlyphNameForGlyph                             func(CGFontRef, unsafe.Pointer) unsafe.Pointer
	_CGFontGetNumberOfGlyphs                                 func(CGFontRef) uintptr
	_CGFontCopyPostScriptName                                func(CGFontRef) unsafe.Pointer
	_CGFontGetStemV                                          func(CGFontRef) CGFloat
	_CGFontCopyTableForTag                                   func(CGFontRef, uint32) unsafe.Pointer
	_CGFontCopyTableTags                                     func(CGFontRef) unsafe.Pointer
	_CGFontGetTypeID                                         func() unsafe.Pointer
	_CGFontGetUnitsPerEm                                     func(CGFontRef) int
	_CGFontCopyVariationAxes                                 func(CGFontRef) unsafe.Pointer
	_CGFontCopyVariations                                    func(CGFontRef) unsafe.Pointer
	_CGFontGetXHeight                                        func(CGFontRef) int
	_CGFontCreateWithPlatformFont                            func(unsafe.Pointer) CGFontRef
	_CGFontRelease                                           func(CGFontRef)
	_CGFontRetain                                            func(CGFontRef) CGFontRef
	_CGFunctionCreate                                        func(unsafe.Pointer, uintptr, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) CGFunctionRef
	_CGFunctionGetTypeID                                     func() unsafe.Pointer
	_CGFunctionRelease                                       func(CGFunctionRef)
	_CGFunctionRetain                                        func(CGFunctionRef) CGFunctionRef
	_CGGetActiveDisplayList                                  func(uint32, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGGetDisplayTransferByFormula                           func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGGetDisplayTransferByTable                             func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGGetDisplaysWithOpenGLDisplayMask                      func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGGetDisplaysWithPoint                                  func(CGPoint, uint32, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGGetDisplaysWithRect                                   func(CGRect, uint32, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGGetEventTapList                                       func(uint32, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGGetLastMouseDelta                                     func(unsafe.Pointer, unsafe.Pointer)
	_CGGetOnlineDisplayList                                  func(uint32, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGGradientGetContentHeadroom                            func(CGGradientRef) float32
	_CGGradientCreateWithColorComponents                     func(CGColorSpaceRef, unsafe.Pointer, unsafe.Pointer, uintptr) CGGradientRef
	_CGGradientCreateWithColors                              func(CGColorSpaceRef, unsafe.Pointer, unsafe.Pointer) CGGradientRef
	_CGGradientCreateWithContentHeadroom                     func(float32, CGColorSpaceRef, unsafe.Pointer, unsafe.Pointer, uintptr) CGGradientRef
	_CGGradientGetTypeID                                     func() unsafe.Pointer
	_CGGradientRelease                                       func(CGGradientRef)
	_CGGradientRetain                                        func(CGGradientRef) CGGradientRef
	_CGImageGetAlphaInfo                                     func(CGImageRef) unsafe.Pointer
	_CGImageGetBitmapInfo                                    func(CGImageRef) unsafe.Pointer
	_CGImageGetBitsPerComponent                              func(CGImageRef) uintptr
	_CGImageGetBitsPerPixel                                  func(CGImageRef) uintptr
	_CGImageGetByteOrderInfo                                 func(CGImageRef) unsafe.Pointer
	_CGImageGetBytesPerRow                                   func(CGImageRef) uintptr
	_CGImageCalculateContentAverageLightLevel                func(CGImageRef) float32
	_CGImageCalculateContentHeadroom                         func(CGImageRef) float32
	_CGImageGetColorSpace                                    func(CGImageRef) CGColorSpaceRef
	_CGImageContainsImageSpecificToneMappingMetadata         func(CGImageRef) bool
	_CGImageGetContentAverageLightLevel                      func(CGImageRef) float32
	_CGImageGetContentHeadroom                               func(CGImageRef) float32
	_CGImageCreateCopy                                       func(CGImageRef) CGImageRef
	_CGImageCreateCopyWithColorSpace                         func(CGImageRef, CGColorSpaceRef) CGImageRef
	_CGImageCreateCopyWithContentAverageLightLevel           func(CGImageRef, float32) CGImageRef
	_CGImageCreateCopyWithCalculatedHDRStats                 func(CGImageRef) CGImageRef
	_CGImageCreateWithImageInRect                            func(CGImageRef, CGRect) CGImageRef
	_CGImageGetDataProvider                                  func(CGImageRef) CGDataProviderRef
	_CGImageGetDecode                                        func(CGImageRef) unsafe.Pointer
	_CGImageGetHeight                                        func(CGImageRef) uintptr
	_CGImageCreateWithContentHeadroom                        func(float32, uintptr, uintptr, uintptr, uintptr, uintptr, CGColorSpaceRef, unsafe.Pointer, CGDataProviderRef, unsafe.Pointer, bool, unsafe.Pointer) CGImageRef
	_CGImageCreateWithJPEGDataProvider                       func(CGDataProviderRef, unsafe.Pointer, bool, unsafe.Pointer) CGImageRef
	_CGImageMaskCreate                                       func(uintptr, uintptr, uintptr, uintptr, uintptr, CGDataProviderRef, unsafe.Pointer, bool) CGImageRef
	_CGImageCreateWithPNGDataProvider                        func(CGDataProviderRef, unsafe.Pointer, bool, unsafe.Pointer) CGImageRef
	_CGImageCreate                                           func(uintptr, uintptr, uintptr, uintptr, uintptr, CGColorSpaceRef, unsafe.Pointer, CGDataProviderRef, unsafe.Pointer, bool, unsafe.Pointer) CGImageRef
	_CGWindowListCreateImageFromArray                        func(CGRect, unsafe.Pointer, unsafe.Pointer) CGImageRef
	_CGImageIsMask                                           func(CGImageRef) bool
	_CGImageCreateWithMask                                   func(CGImageRef, CGImageRef) CGImageRef
	_CGImageGetPixelFormatInfo                               func(CGImageRef) unsafe.Pointer
	_CGImageGetRenderingIntent                               func(CGImageRef) unsafe.Pointer
	_CGImageGetShouldInterpolate                             func(CGImageRef) bool
	_CGImageShouldToneMap                                    func(CGImageRef) bool
	_CGImageGetTypeID                                        func() unsafe.Pointer
	_CGImageGetUTType                                        func(CGImageRef) unsafe.Pointer
	_CGImageGetWidth                                         func(CGImageRef) uintptr
	_CGImageCreateCopyWithContentHeadroom                    func(float32, CGImageRef) CGImageRef
	_CGImageCreateWithMaskingColors                          func(CGImageRef, unsafe.Pointer) CGImageRef
	_CGImageRelease                                          func(CGImageRef)
	_CGImageRetain                                           func(CGImageRef) CGImageRef
	_CGInhibitLocalEvents                                    func(unsafe.Pointer) unsafe.Pointer
	_CGLayerGetContext                                       func(CGLayerRef) CGContextRef
	_CGLayerCreateWithContext                                func(CGContextRef, CGSize, unsafe.Pointer) CGLayerRef
	_CGLayerGetSize                                          func(CGLayerRef) CGSize
	_CGLayerGetTypeID                                        func() unsafe.Pointer
	_CGLayerRelease                                          func(CGLayerRef)
	_CGLayerRetain                                           func(CGLayerRef) CGLayerRef
	_CGMainDisplayID                                         func() unsafe.Pointer
	_CGPathCloseSubpath                                      func(CGMutablePathRef)
	_CGPathCreateMutable                                     func() CGMutablePathRef
	_CGOpenGLDisplayMaskToDisplayID                          func(unsafe.Pointer) unsafe.Pointer
	_CGPDFArrayApplyBlock                                    func(CGPDFArrayRef, unsafe.Pointer, unsafe.Pointer)
	_CGPDFArrayGetArray                                      func(CGPDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetBoolean                                    func(CGPDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetCount                                      func(CGPDFArrayRef) uintptr
	_CGPDFArrayGetDictionary                                 func(CGPDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetInteger                                    func(CGPDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetName                                       func(CGPDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetNull                                       func(CGPDFArrayRef, uintptr) bool
	_CGPDFArrayGetNumber                                     func(CGPDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetObject                                     func(CGPDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetStream                                     func(CGPDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetString                                     func(CGPDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFContentStreamCreateWithPage                        func(CGPDFPageRef) CGPDFContentStreamRef
	_CGPDFContentStreamCreateWithStream                      func(CGPDFStreamRef, CGPDFDictionaryRef, CGPDFContentStreamRef) CGPDFContentStreamRef
	_CGPDFContentStreamGetResource                           func(CGPDFContentStreamRef, unsafe.Pointer, unsafe.Pointer) CGPDFObjectRef
	_CGPDFContentStreamGetStreams                            func(CGPDFContentStreamRef) unsafe.Pointer
	_CGPDFContentStreamRelease                               func(CGPDFContentStreamRef)
	_CGPDFContentStreamRetain                                func(CGPDFContentStreamRef) CGPDFContentStreamRef
	_CGPDFContextBeginTag                                    func(CGContextRef, unsafe.Pointer, unsafe.Pointer)
	_CGPDFContextEndTag                                      func(CGContextRef)
	_CGPDFContextSetIDTree                                   func(CGContextRef, CGPDFDictionaryRef)
	_CGPDFContextSetOutline                                  func(CGContextRef, unsafe.Pointer)
	_CGPDFContextSetPageTagStructureTree                     func(CGContextRef, unsafe.Pointer)
	_CGPDFContextSetParentTree                               func(CGContextRef, CGPDFDictionaryRef)
	_CGPDFDictionaryApplyBlock                               func(CGPDFDictionaryRef, unsafe.Pointer, unsafe.Pointer)
	_CGPDFDictionaryApplyFunction                            func(CGPDFDictionaryRef, unsafe.Pointer, unsafe.Pointer)
	_CGPDFDictionaryGetArray                                 func(CGPDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetBoolean                               func(CGPDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetCount                                 func(CGPDFDictionaryRef) uintptr
	_CGPDFDictionaryGetDictionary                            func(CGPDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetInteger                               func(CGPDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetName                                  func(CGPDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetNumber                                func(CGPDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetObject                                func(CGPDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetStream                                func(CGPDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetString                                func(CGPDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDocumentGetAccessPermissions                       func(CGPDFDocumentRef) unsafe.Pointer
	_CGPDFDocumentAllowsCopying                              func(CGPDFDocumentRef) bool
	_CGPDFDocumentAllowsPrinting                             func(CGPDFDocumentRef) bool
	_CGPDFDocumentGetCatalog                                 func(CGPDFDocumentRef) CGPDFDictionaryRef
	_CGPDFDocumentGetID                                      func(CGPDFDocumentRef) CGPDFArrayRef
	_CGPDFDocumentGetVersion                                 func(CGPDFDocumentRef, unsafe.Pointer, unsafe.Pointer)
	_CGPDFDocumentGetInfo                                    func(CGPDFDocumentRef) CGPDFDictionaryRef
	_CGPDFDocumentCreateWithURL                              func(unsafe.Pointer) CGPDFDocumentRef
	_CGPDFDocumentCreateWithProvider                         func(CGDataProviderRef) CGPDFDocumentRef
	_CGPDFDocumentIsEncrypted                                func(CGPDFDocumentRef) bool
	_CGPDFDocumentIsUnlocked                                 func(CGPDFDocumentRef) bool
	_CGPDFDocumentGetNumberOfPages                           func(CGPDFDocumentRef) uintptr
	_CGPDFDocumentGetOutline                                 func(CGPDFDocumentRef) unsafe.Pointer
	_CGPDFDocumentGetPage                                    func(CGPDFDocumentRef, uintptr) CGPDFPageRef
	_CGPDFDocumentGetTypeID                                  func() unsafe.Pointer
	_CGPDFDocumentUnlockWithPassword                         func(CGPDFDocumentRef, unsafe.Pointer) bool
	_CGPDFDocumentGetArtBox                                  func(CGPDFDocumentRef, int) CGRect
	_CGPDFDocumentGetBleedBox                                func(CGPDFDocumentRef, int) CGRect
	_CGPDFDocumentGetCropBox                                 func(CGPDFDocumentRef, int) CGRect
	_CGPDFDocumentGetMediaBox                                func(CGPDFDocumentRef, int) CGRect
	_CGPDFDocumentGetRotationAngle                           func(CGPDFDocumentRef, int) int
	_CGPDFDocumentGetTrimBox                                 func(CGPDFDocumentRef, int) CGRect
	_CGPDFDocumentRelease                                    func(CGPDFDocumentRef)
	_CGPDFDocumentRetain                                     func(CGPDFDocumentRef) CGPDFDocumentRef
	_CGPDFObjectGetType                                      func(CGPDFObjectRef) unsafe.Pointer
	_CGPDFObjectGetValue                                     func(CGPDFObjectRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFOperatorTableCreate                                func() CGPDFOperatorTableRef
	_CGPDFOperatorTableRelease                               func(CGPDFOperatorTableRef)
	_CGPDFOperatorTableRetain                                func(CGPDFOperatorTableRef) CGPDFOperatorTableRef
	_CGPDFOperatorTableSetCallback                           func(CGPDFOperatorTableRef, unsafe.Pointer, unsafe.Pointer)
	_CGPDFPageGetDictionary                                  func(CGPDFPageRef) CGPDFDictionaryRef
	_CGPDFPageGetDocument                                    func(CGPDFPageRef) CGPDFDocumentRef
	_CGPDFPageGetBoxRect                                     func(CGPDFPageRef, unsafe.Pointer) CGRect
	_CGPDFPageGetDrawingTransform                            func(CGPDFPageRef, unsafe.Pointer, CGRect, int, bool) CGAffineTransform
	_CGPDFPageGetPageNumber                                  func(CGPDFPageRef) uintptr
	_CGPDFPageGetRotationAngle                               func(CGPDFPageRef) int
	_CGPDFPageGetTypeID                                      func() unsafe.Pointer
	_CGPDFPageRelease                                        func(CGPDFPageRef)
	_CGPDFPageRetain                                         func(CGPDFPageRef) CGPDFPageRef
	_CGPDFScannerCreate                                      func(CGPDFContentStreamRef, CGPDFOperatorTableRef, unsafe.Pointer) CGPDFScannerRef
	_CGPDFScannerGetContentStream                            func(CGPDFScannerRef) CGPDFContentStreamRef
	_CGPDFScannerPopArray                                    func(CGPDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopBoolean                                  func(CGPDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopDictionary                               func(CGPDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopInteger                                  func(CGPDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopName                                     func(CGPDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopNumber                                   func(CGPDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopObject                                   func(CGPDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopStream                                   func(CGPDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopString                                   func(CGPDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerRelease                                     func(CGPDFScannerRef)
	_CGPDFScannerRetain                                      func(CGPDFScannerRef) CGPDFScannerRef
	_CGPDFScannerScan                                        func(CGPDFScannerRef) bool
	_CGPDFScannerStop                                        func(CGPDFScannerRef)
	_CGPDFStreamCopyData                                     func(CGPDFStreamRef, unsafe.Pointer) unsafe.Pointer
	_CGPDFStreamGetDictionary                                func(CGPDFStreamRef) CGPDFDictionaryRef
	_CGPDFStringCopyDate                                     func(CGPDFStringRef) unsafe.Pointer
	_CGPDFStringCopyTextString                               func(CGPDFStringRef) unsafe.Pointer
	_CGPDFStringGetBytePtr                                   func(CGPDFStringRef) unsafe.Pointer
	_CGPDFStringGetLength                                    func(CGPDFStringRef) uintptr
	_CGPDFTagTypeGetName                                     func(unsafe.Pointer) unsafe.Pointer
	_CGPSConverterAbort                                      func(CGPSConverterRef) bool
	_CGPSConverterConvert                                    func(CGPSConverterRef, CGDataProviderRef, CGDataConsumerRef, unsafe.Pointer) bool
	_CGPSConverterCreate                                     func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGPSConverterRef
	_CGPSConverterIsConverting                               func(CGPSConverterRef) bool
	_CGPSConverterGetTypeID                                  func() unsafe.Pointer
	_CGPathApply                                             func(CGPathRef, unsafe.Pointer, unsafe.Pointer)
	_CGPathApplyWithBlock                                    func(CGPathRef, unsafe.Pointer)
	_CGPathGetBoundingBox                                    func(CGPathRef) CGRect
	_CGPathGetPathBoundingBox                                func(CGPathRef) CGRect
	_CGPathCreateCopy                                        func(CGPathRef) CGPathRef
	_CGPathCreateCopyByTransformingPath                      func(CGPathRef, unsafe.Pointer) CGPathRef
	_CGPathGetCurrentPoint                                   func(CGPathRef) CGPoint
	_CGPathCreateWithEllipseInRect                           func(CGRect, unsafe.Pointer) CGPathRef
	_CGPathCreateWithRect                                    func(CGRect, unsafe.Pointer) CGPathRef
	_CGPathCreateWithRoundedRect                             func(CGRect, CGFloat, CGFloat, unsafe.Pointer) CGPathRef
	_CGPathIsEmpty                                           func(CGPathRef) bool
	_CGPathIsRect                                            func(CGPathRef, unsafe.Pointer) bool
	_CGPathCreateMutableCopy                                 func(CGPathRef) CGMutablePathRef
	_CGPathCreateMutableCopyByTransformingPath               func(CGPathRef, unsafe.Pointer) CGMutablePathRef
	_CGPathGetTypeID                                         func() unsafe.Pointer
	_CGPathAddArc                                            func(CGMutablePathRef, unsafe.Pointer, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat, bool)
	_CGPathAddArcToPoint                                     func(CGMutablePathRef, unsafe.Pointer, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat)
	_CGPathAddCurveToPoint                                   func(CGMutablePathRef, unsafe.Pointer, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat)
	_CGPathAddEllipseInRect                                  func(CGMutablePathRef, unsafe.Pointer, CGRect)
	_CGPathAddLineToPoint                                    func(CGMutablePathRef, unsafe.Pointer, CGFloat, CGFloat)
	_CGPathAddLines                                          func(CGMutablePathRef, unsafe.Pointer, unsafe.Pointer, uintptr)
	_CGPathAddPath                                           func(CGMutablePathRef, unsafe.Pointer, CGPathRef)
	_CGPathAddQuadCurveToPoint                               func(CGMutablePathRef, unsafe.Pointer, CGFloat, CGFloat, CGFloat, CGFloat)
	_CGPathAddRect                                           func(CGMutablePathRef, unsafe.Pointer, CGRect)
	_CGPathAddRects                                          func(CGMutablePathRef, unsafe.Pointer, unsafe.Pointer, uintptr)
	_CGPathAddRelativeArc                                    func(CGMutablePathRef, unsafe.Pointer, CGFloat, CGFloat, CGFloat, CGFloat, CGFloat)
	_CGPathAddRoundedRect                                    func(CGMutablePathRef, unsafe.Pointer, CGRect, CGFloat, CGFloat)
	_CGPathContainsPoint                                     func(CGPathRef, unsafe.Pointer, CGPoint, bool) bool
	_CGPathCreateCopyByDashingPath                           func(CGPathRef, unsafe.Pointer, CGFloat, unsafe.Pointer, uintptr) CGPathRef
	_CGPathCreateCopyByFlattening                            func(CGPathRef, CGFloat) CGPathRef
	_CGPathCreateCopyByIntersectingPath                      func(CGPathRef, CGPathRef, bool) CGPathRef
	_CGPathCreateCopyByNormalizing                           func(CGPathRef, bool) CGPathRef
	_CGPathCreateCopyByStrokingPath                          func(CGPathRef, unsafe.Pointer, CGFloat, unsafe.Pointer, unsafe.Pointer, CGFloat) CGPathRef
	_CGPathCreateCopyBySubtractingPath                       func(CGPathRef, CGPathRef, bool) CGPathRef
	_CGPathCreateCopyBySymmetricDifferenceOfPath             func(CGPathRef, CGPathRef, bool) CGPathRef
	_CGPathCreateCopyByUnioningPath                          func(CGPathRef, CGPathRef, bool) CGPathRef
	_CGPathCreateCopyOfLineByIntersectingPath                func(CGPathRef, CGPathRef, bool) CGPathRef
	_CGPathCreateCopyOfLineBySubtractingPath                 func(CGPathRef, CGPathRef, bool) CGPathRef
	_CGPathCreateSeparateComponents                          func(CGPathRef, bool) unsafe.Pointer
	_CGPathEqualToPath                                       func(CGPathRef, CGPathRef) bool
	_CGPathIntersectsPath                                    func(CGPathRef, CGPathRef, bool) bool
	_CGPathMoveToPoint                                       func(CGMutablePathRef, unsafe.Pointer, CGFloat, CGFloat)
	_CGPathRelease                                           func(CGPathRef)
	_CGPathRetain                                            func(CGPathRef) CGPathRef
	_CGPatternCreate                                         func(unsafe.Pointer, CGRect, CGAffineTransform, CGFloat, CGFloat, unsafe.Pointer, bool, unsafe.Pointer) CGPatternRef
	_CGPatternGetTypeID                                      func() unsafe.Pointer
	_CGPatternRelease                                        func(CGPatternRef)
	_CGPatternRetain                                         func(CGPatternRef) CGPatternRef
	_CGPointApplyAffineTransform                             func(CGPoint, CGAffineTransform) CGPoint
	_CGPointCreateDictionaryRepresentation                   func(CGPoint) unsafe.Pointer
	_CGPointEqualToPoint                                     func(CGPoint, CGPoint) bool
	_CGPointMakeWithDictionaryRepresentation                 func(unsafe.Pointer, unsafe.Pointer) bool
	_CGPostKeyboardEvent                                     func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGPostMouseEvent                                        func(CGPoint, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGPostScrollWheelEvent                                  func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGPreflightListenEventAccess                            func() bool
	_CGPreflightPostEventAccess                              func() bool
	_CGPreflightScreenCaptureAccess                          func() bool
	_CGRectApplyAffineTransform                              func(CGRect, CGAffineTransform) CGRect
	_CGRectContainsPoint                                     func(CGRect, CGPoint) bool
	_CGRectContainsRect                                      func(CGRect, CGRect) bool
	_CGRectCreateDictionaryRepresentation                    func(CGRect) unsafe.Pointer
	_CGRectDivide                                            func(CGRect, unsafe.Pointer, unsafe.Pointer, CGFloat, unsafe.Pointer)
	_CGRectEqualToRect                                       func(CGRect, CGRect) bool
	_CGRectGetHeight                                         func(CGRect) CGFloat
	_CGRectGetMaxX                                           func(CGRect) CGFloat
	_CGRectGetMaxY                                           func(CGRect) CGFloat
	_CGRectGetMidX                                           func(CGRect) CGFloat
	_CGRectGetMidY                                           func(CGRect) CGFloat
	_CGRectGetMinX                                           func(CGRect) CGFloat
	_CGRectGetMinY                                           func(CGRect) CGFloat
	_CGRectGetWidth                                          func(CGRect) CGFloat
	_CGRectInset                                             func(CGRect, CGFloat, CGFloat) CGRect
	_CGRectIntegral                                          func(CGRect) CGRect
	_CGRectIntersection                                      func(CGRect, CGRect) CGRect
	_CGRectIntersectsRect                                    func(CGRect, CGRect) bool
	_CGRectIsEmpty                                           func(CGRect) bool
	_CGRectIsInfinite                                        func(CGRect) bool
	_CGRectIsNull                                            func(CGRect) bool
	_CGRectMakeWithDictionaryRepresentation                  func(unsafe.Pointer, unsafe.Pointer) bool
	_CGRectOffset                                            func(CGRect, CGFloat, CGFloat) CGRect
	_CGRectStandardize                                       func(CGRect) CGRect
	_CGRectUnion                                             func(CGRect, CGRect) CGRect
	_CGRegisterScreenRefreshCallback                         func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGReleaseAllDisplays                                    func() unsafe.Pointer
	_CGReleaseDisplayFadeReservation                         func(unsafe.Pointer) unsafe.Pointer
	_CGReleaseScreenRefreshRects                             func(unsafe.Pointer)
	_CGRenderingBufferLockBytePtr                            func(CGRenderingBufferProviderRef) unsafe.Pointer
	_CGRenderingBufferProviderCreate                         func(unsafe.Pointer, uintptr) CGRenderingBufferProviderRef
	_CGRenderingBufferProviderCreateWithCFData               func(unsafe.Pointer) CGRenderingBufferProviderRef
	_CGRenderingBufferProviderGetSize                        func(CGRenderingBufferProviderRef) uintptr
	_CGRenderingBufferProviderGetTypeID                      func() unsafe.Pointer
	_CGRenderingBufferUnlockBytePtr                          func(CGRenderingBufferProviderRef)
	_CGRequestListenEventAccess                              func() bool
	_CGRequestPostEventAccess                                func() bool
	_CGRequestScreenCaptureAccess                            func() bool
	_CGRestorePermanentDisplayConfiguration                  func()
	_CGScreenRegisterMoveCallback                            func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGScreenUnregisterMoveCallback                          func(unsafe.Pointer, unsafe.Pointer)
	_CGSessionCopyCurrentDictionary                          func() unsafe.Pointer
	_CGSetDisplayTransferByByteTable                         func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGSetDisplayTransferByFormula                           func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGSetDisplayTransferByTable                             func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGSetLocalEventsFilterDuringSuppressionState            func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGSetLocalEventsSuppressionInterval                     func(unsafe.Pointer) unsafe.Pointer
	_CGShadingGetContentHeadroom                             func(CGShadingRef) float32
	_CGShadingCreateAxialWithContentHeadroom                 func(float32, CGColorSpaceRef, CGPoint, CGPoint, CGFunctionRef, bool, bool) CGShadingRef
	_CGShadingCreateAxial                                    func(CGColorSpaceRef, CGPoint, CGPoint, CGFunctionRef, bool, bool) CGShadingRef
	_CGShadingCreateRadialWithContentHeadroom                func(float32, CGColorSpaceRef, CGPoint, CGFloat, CGPoint, CGFloat, CGFunctionRef, bool, bool) CGShadingRef
	_CGShadingCreateRadial                                   func(CGColorSpaceRef, CGPoint, CGFloat, CGPoint, CGFloat, CGFunctionRef, bool, bool) CGShadingRef
	_CGShadingGetTypeID                                      func() unsafe.Pointer
	_CGShadingRelease                                        func(CGShadingRef)
	_CGShadingRetain                                         func(CGShadingRef) CGShadingRef
	_CGShieldingWindowID                                     func(unsafe.Pointer) unsafe.Pointer
	_CGShieldingWindowLevel                                  func() unsafe.Pointer
	_CGSizeApplyAffineTransform                              func(CGSize, CGAffineTransform) CGSize
	_CGSizeCreateDictionaryRepresentation                    func(CGSize) unsafe.Pointer
	_CGSizeEqualToSize                                       func(CGSize, CGSize) bool
	_CGSizeMakeWithDictionaryRepresentation                  func(unsafe.Pointer, unsafe.Pointer) bool
	_CGUnregisterScreenRefreshCallback                       func(unsafe.Pointer, unsafe.Pointer)
	_CGWaitForScreenRefreshRects                             func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGWaitForScreenUpdateRects                              func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGWarpMouseCursorPosition                               func(CGPoint) unsafe.Pointer
	_CGWindowLevelForKey                                     func(unsafe.Pointer) unsafe.Pointer
	_CGWindowListCopyWindowInfo                              func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGWindowListCreate                                      func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CGWindowListCreateDescriptionFromArray                  func(unsafe.Pointer) unsafe.Pointer
	_CGWindowListCreateImage                                 func(CGRect, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CGImageRef
	_CGWindowServerCFMachPort                                func() unsafe.Pointer
	_CGWindowServerCreateServerPort                          func() unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	tryRegister(&_CGColorSpaceRelease, lib, "CGColorSpaceRelease")
	tryRegister(&_CGContextSetInterpolationQuality, lib, "CGContextSetInterpolationQuality")
	tryRegister(&_CGAcquireDisplayFadeReservation, lib, "CGAcquireDisplayFadeReservation")
	tryRegister(&_CGAffineTransformConcat, lib, "CGAffineTransformConcat")
	tryRegister(&_CGAffineTransformDecompose, lib, "CGAffineTransformDecompose")
	tryRegister(&_CGAffineTransformEqualToTransform, lib, "CGAffineTransformEqualToTransform")
	tryRegister(&_CGAffineTransformInvert, lib, "CGAffineTransformInvert")
	tryRegister(&_CGAffineTransformIsIdentity, lib, "CGAffineTransformIsIdentity")
	tryRegister(&_CGAffineTransformMake, lib, "CGAffineTransformMake")
	tryRegister(&_CGAffineTransformMakeRotation, lib, "CGAffineTransformMakeRotation")
	tryRegister(&_CGAffineTransformMakeScale, lib, "CGAffineTransformMakeScale")
	tryRegister(&_CGAffineTransformMakeTranslation, lib, "CGAffineTransformMakeTranslation")
	tryRegister(&_CGAffineTransformMakeWithComponents, lib, "CGAffineTransformMakeWithComponents")
	tryRegister(&_CGAffineTransformRotate, lib, "CGAffineTransformRotate")
	tryRegister(&_CGAffineTransformScale, lib, "CGAffineTransformScale")
	tryRegister(&_CGAffineTransformTranslate, lib, "CGAffineTransformTranslate")
	tryRegister(&_CGAssociateMouseAndMouseCursorPosition, lib, "CGAssociateMouseAndMouseCursorPosition")
	tryRegister(&_CGBeginDisplayConfiguration, lib, "CGBeginDisplayConfiguration")
	tryRegister(&_CGBitmapContextCreateAdaptive, lib, "CGBitmapContextCreateAdaptive")
	tryRegister(&_CGCancelDisplayConfiguration, lib, "CGCancelDisplayConfiguration")
	tryRegister(&_CGCaptureAllDisplays, lib, "CGCaptureAllDisplays")
	tryRegister(&_CGCaptureAllDisplaysWithOptions, lib, "CGCaptureAllDisplaysWithOptions")
	tryRegister(&_CGColorGetAlpha, lib, "CGColorGetAlpha")
	tryRegister(&_CGColorGetColorSpace, lib, "CGColorGetColorSpace")
	tryRegister(&_CGColorGetContentHeadroom, lib, "CGColorGetContentHeadroom")
	tryRegister(&_CGColorCreateCopyByMatchingToColorSpace, lib, "CGColorCreateCopyByMatchingToColorSpace")
	tryRegister(&_CGColorCreateCopy, lib, "CGColorCreateCopy")
	tryRegister(&_CGColorCreateCopyWithAlpha, lib, "CGColorCreateCopyWithAlpha")
	tryRegister(&_CGColorCreate, lib, "CGColorCreate")
	tryRegister(&_CGColorCreateGenericCMYK, lib, "CGColorCreateGenericCMYK")
	tryRegister(&_CGColorCreateGenericGrayGamma2_2, lib, "CGColorCreateGenericGrayGamma2_2")
	tryRegister(&_CGColorCreateGenericGray, lib, "CGColorCreateGenericGray")
	tryRegister(&_CGColorCreateWithContentHeadroom, lib, "CGColorCreateWithContentHeadroom")
	tryRegister(&_CGColorCreateWithPattern, lib, "CGColorCreateWithPattern")
	tryRegister(&_CGColorCreateGenericRGB, lib, "CGColorCreateGenericRGB")
	tryRegister(&_CGColorCreateSRGB, lib, "CGColorCreateSRGB")
	tryRegister(&_CGColorGetNumberOfComponents, lib, "CGColorGetNumberOfComponents")
	tryRegister(&_CGColorGetPattern, lib, "CGColorGetPattern")
	tryRegister(&_CGColorGetTypeID, lib, "CGColorGetTypeID")
	tryRegister(&_CGColorConversionInfoConvertData, lib, "CGColorConversionInfoConvertData")
	tryRegister(&_CGColorConversionInfoCreateWithOptions, lib, "CGColorConversionInfoCreateWithOptions")
	tryRegister(&_CGColorConversionInfoCreate, lib, "CGColorConversionInfoCreate")
	tryRegister(&_CGColorConversionInfoCreateForToneMapping, lib, "CGColorConversionInfoCreateForToneMapping")
	tryRegister(&_CGColorConversionInfoGetTypeID, lib, "CGColorConversionInfoGetTypeID")
	tryRegister(&_CGColorConversionInfoCreateFromList, lib, "CGColorConversionInfoCreateFromList")
	tryRegister(&_CGColorConversionInfoCreateFromListWithArguments, lib, "CGColorConversionInfoCreateFromListWithArguments")
	tryRegister(&_CGColorEqualToColor, lib, "CGColorEqualToColor")
	tryRegister(&_CGColorGetComponents, lib, "CGColorGetComponents")
	tryRegister(&_CGColorGetConstantColor, lib, "CGColorGetConstantColor")
	tryRegister(&_CGColorRelease, lib, "CGColorRelease")
	tryRegister(&_CGColorRetain, lib, "CGColorRetain")
	tryRegister(&_CGColorSpaceGetBaseColorSpace, lib, "CGColorSpaceGetBaseColorSpace")
	tryRegister(&_CGColorSpaceCopyICCData, lib, "CGColorSpaceCopyICCData")
	tryRegister(&_CGColorSpaceCopyPropertyList, lib, "CGColorSpaceCopyPropertyList")
	tryRegister(&_CGColorSpaceCopyICCProfile, lib, "CGColorSpaceCopyICCProfile")
	tryRegister(&_CGColorSpaceCreateCalibratedGray, lib, "CGColorSpaceCreateCalibratedGray")
	tryRegister(&_CGColorSpaceCreateCalibratedRGB, lib, "CGColorSpaceCreateCalibratedRGB")
	tryRegister(&_CGColorSpaceCreateICCBased, lib, "CGColorSpaceCreateICCBased")
	tryRegister(&_CGColorSpaceCreateWithICCData, lib, "CGColorSpaceCreateWithICCData")
	tryRegister(&_CGColorSpaceCreateWithICCProfile, lib, "CGColorSpaceCreateWithICCProfile")
	tryRegister(&_CGColorSpaceCreateIndexed, lib, "CGColorSpaceCreateIndexed")
	tryRegister(&_CGColorSpaceCreateLab, lib, "CGColorSpaceCreateLab")
	tryRegister(&_CGColorSpaceCreateWithName, lib, "CGColorSpaceCreateWithName")
	tryRegister(&_CGColorSpaceCreatePattern, lib, "CGColorSpaceCreatePattern")
	tryRegister(&_CGColorSpaceCreateWithPlatformColorSpace, lib, "CGColorSpaceCreateWithPlatformColorSpace")
	tryRegister(&_CGColorSpaceCreateWithPropertyList, lib, "CGColorSpaceCreateWithPropertyList")
	tryRegister(&_CGColorSpaceIsHDR, lib, "CGColorSpaceIsHDR")
	tryRegister(&_CGColorSpaceIsWideGamutRGB, lib, "CGColorSpaceIsWideGamutRGB")
	tryRegister(&_CGColorSpaceGetModel, lib, "CGColorSpaceGetModel")
	tryRegister(&_CGColorSpaceCopyName, lib, "CGColorSpaceCopyName")
	tryRegister(&_CGColorSpaceGetNumberOfComponents, lib, "CGColorSpaceGetNumberOfComponents")
	tryRegister(&_CGColorSpaceSupportsOutput, lib, "CGColorSpaceSupportsOutput")
	tryRegister(&_CGColorSpaceGetTypeID, lib, "CGColorSpaceGetTypeID")
	tryRegister(&_CGColorSpaceCopyBaseColorSpace, lib, "CGColorSpaceCopyBaseColorSpace")
	tryRegister(&_CGColorSpaceCreateCopyWithStandardRange, lib, "CGColorSpaceCreateCopyWithStandardRange")
	tryRegister(&_CGColorSpaceCreateDeviceCMYK, lib, "CGColorSpaceCreateDeviceCMYK")
	tryRegister(&_CGColorSpaceCreateDeviceGray, lib, "CGColorSpaceCreateDeviceGray")
	tryRegister(&_CGColorSpaceCreateDeviceRGB, lib, "CGColorSpaceCreateDeviceRGB")
	tryRegister(&_CGColorSpaceCreateExtended, lib, "CGColorSpaceCreateExtended")
	tryRegister(&_CGColorSpaceCreateExtendedLinearized, lib, "CGColorSpaceCreateExtendedLinearized")
	tryRegister(&_CGColorSpaceCreateLinearized, lib, "CGColorSpaceCreateLinearized")
	tryRegister(&_CGColorSpaceCreateWithColorSyncProfile, lib, "CGColorSpaceCreateWithColorSyncProfile")
	tryRegister(&_CGColorSpaceGetColorTable, lib, "CGColorSpaceGetColorTable")
	tryRegister(&_CGColorSpaceGetColorTableCount, lib, "CGColorSpaceGetColorTableCount")
	tryRegister(&_CGColorSpaceGetName, lib, "CGColorSpaceGetName")
	tryRegister(&_CGColorSpaceIsHLGBased, lib, "CGColorSpaceIsHLGBased")
	tryRegister(&_CGColorSpaceIsPQBased, lib, "CGColorSpaceIsPQBased")
	tryRegister(&_CGColorSpaceRetain, lib, "CGColorSpaceRetain")
	tryRegister(&_CGColorSpaceUsesExtendedRange, lib, "CGColorSpaceUsesExtendedRange")
	tryRegister(&_CGColorSpaceUsesITUR_2100TF, lib, "CGColorSpaceUsesITUR_2100TF")
	tryRegister(&_CGCompleteDisplayConfiguration, lib, "CGCompleteDisplayConfiguration")
	tryRegister(&_CGConfigureDisplayFadeEffect, lib, "CGConfigureDisplayFadeEffect")
	tryRegister(&_CGConfigureDisplayMirrorOfDisplay, lib, "CGConfigureDisplayMirrorOfDisplay")
	tryRegister(&_CGConfigureDisplayMode, lib, "CGConfigureDisplayMode")
	tryRegister(&_CGConfigureDisplayOrigin, lib, "CGConfigureDisplayOrigin")
	tryRegister(&_CGConfigureDisplayStereoOperation, lib, "CGConfigureDisplayStereoOperation")
	tryRegister(&_CGConfigureDisplayWithDisplayMode, lib, "CGConfigureDisplayWithDisplayMode")
	tryRegister(&_CGPDFContextAddDestinationAtPoint, lib, "CGPDFContextAddDestinationAtPoint")
	tryRegister(&_CGPDFContextAddDocumentMetadata, lib, "CGPDFContextAddDocumentMetadata")
	tryRegister(&_CGContextAddEllipseInRect, lib, "CGContextAddEllipseInRect")
	tryRegister(&_CGContextAddPath, lib, "CGContextAddPath")
	tryRegister(&_CGContextAddRect, lib, "CGContextAddRect")
	tryRegister(&_CGBitmapContextGetAlphaInfo, lib, "CGBitmapContextGetAlphaInfo")
	tryRegister(&_CGPDFContextBeginPage, lib, "CGPDFContextBeginPage")
	tryRegister(&_CGContextBeginPage, lib, "CGContextBeginPage")
	tryRegister(&_CGContextBeginPath, lib, "CGContextBeginPath")
	tryRegister(&_CGContextBeginTransparencyLayer, lib, "CGContextBeginTransparencyLayer")
	tryRegister(&_CGContextBeginTransparencyLayerWithRect, lib, "CGContextBeginTransparencyLayerWithRect")
	tryRegister(&_CGBitmapContextGetBitmapInfo, lib, "CGBitmapContextGetBitmapInfo")
	tryRegister(&_CGBitmapContextGetBitsPerComponent, lib, "CGBitmapContextGetBitsPerComponent")
	tryRegister(&_CGBitmapContextGetBitsPerPixel, lib, "CGBitmapContextGetBitsPerPixel")
	tryRegister(&_CGContextGetClipBoundingBox, lib, "CGContextGetClipBoundingBox")
	tryRegister(&_CGContextGetPathBoundingBox, lib, "CGContextGetPathBoundingBox")
	tryRegister(&_CGBitmapContextGetBytesPerRow, lib, "CGBitmapContextGetBytesPerRow")
	tryRegister(&_CGContextClearRect, lib, "CGContextClearRect")
	tryRegister(&_CGContextClipToRect, lib, "CGContextClipToRect")
	tryRegister(&_CGContextClipToMask, lib, "CGContextClipToMask")
	tryRegister(&_CGPDFContextClose, lib, "CGPDFContextClose")
	tryRegister(&_CGContextClosePath, lib, "CGContextClosePath")
	tryRegister(&_CGBitmapContextGetColorSpace, lib, "CGBitmapContextGetColorSpace")
	tryRegister(&_CGContextConcatCTM, lib, "CGContextConcatCTM")
	tryRegister(&_CGContextConvertSizeToDeviceSpace, lib, "CGContextConvertSizeToDeviceSpace")
	tryRegister(&_CGContextConvertPointToDeviceSpace, lib, "CGContextConvertPointToDeviceSpace")
	tryRegister(&_CGContextConvertRectToDeviceSpace, lib, "CGContextConvertRectToDeviceSpace")
	tryRegister(&_CGContextConvertRectToUserSpace, lib, "CGContextConvertRectToUserSpace")
	tryRegister(&_CGContextConvertPointToUserSpace, lib, "CGContextConvertPointToUserSpace")
	tryRegister(&_CGContextConvertSizeToUserSpace, lib, "CGContextConvertSizeToUserSpace")
	tryRegister(&_CGContextGetCTM, lib, "CGContextGetCTM")
	tryRegister(&_CGContextGetPathCurrentPoint, lib, "CGContextGetPathCurrentPoint")
	tryRegister(&_CGBitmapContextGetData, lib, "CGBitmapContextGetData")
	tryRegister(&_CGContextDrawLinearGradient, lib, "CGContextDrawLinearGradient")
	tryRegister(&_CGContextDrawPDFPage, lib, "CGContextDrawPDFPage")
	tryRegister(&_CGContextDrawPath, lib, "CGContextDrawPath")
	tryRegister(&_CGContextDrawRadialGradient, lib, "CGContextDrawRadialGradient")
	tryRegister(&_CGContextDrawShading, lib, "CGContextDrawShading")
	tryRegister(&_CGPDFContextEndPage, lib, "CGPDFContextEndPage")
	tryRegister(&_CGContextEndPage, lib, "CGContextEndPage")
	tryRegister(&_CGContextEndTransparencyLayer, lib, "CGContextEndTransparencyLayer")
	tryRegister(&_CGContextFillRect, lib, "CGContextFillRect")
	tryRegister(&_CGContextFillEllipseInRect, lib, "CGContextFillEllipseInRect")
	tryRegister(&_CGContextFlush, lib, "CGContextFlush")
	tryRegister(&_CGBitmapContextGetHeight, lib, "CGBitmapContextGetHeight")
	tryRegister(&_CGPDFContextCreateWithURL, lib, "CGPDFContextCreateWithURL")
	tryRegister(&_CGPDFContextCreate, lib, "CGPDFContextCreate")
	tryRegister(&_CGBitmapContextCreate, lib, "CGBitmapContextCreate")
	tryRegister(&_CGBitmapContextCreateWithData, lib, "CGBitmapContextCreateWithData")
	tryRegister(&_CGContextGetInterpolationQuality, lib, "CGContextGetInterpolationQuality")
	tryRegister(&_CGContextIsPathEmpty, lib, "CGContextIsPathEmpty")
	tryRegister(&_CGBitmapContextCreateImage, lib, "CGBitmapContextCreateImage")
	tryRegister(&_CGContextCopyPath, lib, "CGContextCopyPath")
	tryRegister(&_CGContextPathContainsPoint, lib, "CGContextPathContainsPoint")
	tryRegister(&_CGContextReplacePathWithStrokedPath, lib, "CGContextReplacePathWithStrokedPath")
	tryRegister(&_CGContextResetClip, lib, "CGContextResetClip")
	tryRegister(&_CGContextRestoreGState, lib, "CGContextRestoreGState")
	tryRegister(&_CGContextRotateCTM, lib, "CGContextRotateCTM")
	tryRegister(&_CGContextSaveGState, lib, "CGContextSaveGState")
	tryRegister(&_CGContextScaleCTM, lib, "CGContextScaleCTM")
	tryRegister(&_CGContextSelectFont, lib, "CGContextSelectFont")
	tryRegister(&_CGContextSetAllowsAntialiasing, lib, "CGContextSetAllowsAntialiasing")
	tryRegister(&_CGContextSetAllowsFontSmoothing, lib, "CGContextSetAllowsFontSmoothing")
	tryRegister(&_CGContextSetAllowsFontSubpixelPositioning, lib, "CGContextSetAllowsFontSubpixelPositioning")
	tryRegister(&_CGContextSetAllowsFontSubpixelQuantization, lib, "CGContextSetAllowsFontSubpixelQuantization")
	tryRegister(&_CGContextSetAlpha, lib, "CGContextSetAlpha")
	tryRegister(&_CGContextSetBlendMode, lib, "CGContextSetBlendMode")
	tryRegister(&_CGContextSetCharacterSpacing, lib, "CGContextSetCharacterSpacing")
	tryRegister(&_CGPDFContextSetDestinationForRect, lib, "CGPDFContextSetDestinationForRect")
	tryRegister(&_CGContextSetEDRTargetHeadroom, lib, "CGContextSetEDRTargetHeadroom")
	tryRegister(&_CGContextSetFillColor, lib, "CGContextSetFillColor")
	tryRegister(&_CGContextSetFillColorWithColor, lib, "CGContextSetFillColorWithColor")
	tryRegister(&_CGContextSetCMYKFillColor, lib, "CGContextSetCMYKFillColor")
	tryRegister(&_CGContextSetGrayFillColor, lib, "CGContextSetGrayFillColor")
	tryRegister(&_CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")
	tryRegister(&_CGContextSetFillColorSpace, lib, "CGContextSetFillColorSpace")
	tryRegister(&_CGContextSetFillPattern, lib, "CGContextSetFillPattern")
	tryRegister(&_CGContextSetFlatness, lib, "CGContextSetFlatness")
	tryRegister(&_CGContextSetFont, lib, "CGContextSetFont")
	tryRegister(&_CGContextSetFontSize, lib, "CGContextSetFontSize")
	tryRegister(&_CGContextSetLineCap, lib, "CGContextSetLineCap")
	tryRegister(&_CGContextSetLineJoin, lib, "CGContextSetLineJoin")
	tryRegister(&_CGContextSetLineWidth, lib, "CGContextSetLineWidth")
	tryRegister(&_CGContextSetMiterLimit, lib, "CGContextSetMiterLimit")
	tryRegister(&_CGContextSetPatternPhase, lib, "CGContextSetPatternPhase")
	tryRegister(&_CGContextSetRenderingIntent, lib, "CGContextSetRenderingIntent")
	tryRegister(&_CGContextSetShadow, lib, "CGContextSetShadow")
	tryRegister(&_CGContextSetShadowWithColor, lib, "CGContextSetShadowWithColor")
	tryRegister(&_CGContextSetShouldAntialias, lib, "CGContextSetShouldAntialias")
	tryRegister(&_CGContextSetShouldSmoothFonts, lib, "CGContextSetShouldSmoothFonts")
	tryRegister(&_CGContextSetShouldSubpixelPositionFonts, lib, "CGContextSetShouldSubpixelPositionFonts")
	tryRegister(&_CGContextSetShouldSubpixelQuantizeFonts, lib, "CGContextSetShouldSubpixelQuantizeFonts")
	tryRegister(&_CGContextSetStrokeColorWithColor, lib, "CGContextSetStrokeColorWithColor")
	tryRegister(&_CGContextSetStrokeColor, lib, "CGContextSetStrokeColor")
	tryRegister(&_CGContextSetCMYKStrokeColor, lib, "CGContextSetCMYKStrokeColor")
	tryRegister(&_CGContextSetGrayStrokeColor, lib, "CGContextSetGrayStrokeColor")
	tryRegister(&_CGContextSetRGBStrokeColor, lib, "CGContextSetRGBStrokeColor")
	tryRegister(&_CGContextSetStrokeColorSpace, lib, "CGContextSetStrokeColorSpace")
	tryRegister(&_CGContextSetStrokePattern, lib, "CGContextSetStrokePattern")
	tryRegister(&_CGContextSetTextDrawingMode, lib, "CGContextSetTextDrawingMode")
	tryRegister(&_CGPDFContextSetURLForRect, lib, "CGPDFContextSetURLForRect")
	tryRegister(&_CGContextShowGlyphs, lib, "CGContextShowGlyphs")
	tryRegister(&_CGContextShowGlyphsAtPoint, lib, "CGContextShowGlyphsAtPoint")
	tryRegister(&_CGContextShowGlyphsWithAdvances, lib, "CGContextShowGlyphsWithAdvances")
	tryRegister(&_CGContextShowText, lib, "CGContextShowText")
	tryRegister(&_CGContextShowTextAtPoint, lib, "CGContextShowTextAtPoint")
	tryRegister(&_CGContextStrokeRect, lib, "CGContextStrokeRect")
	tryRegister(&_CGContextStrokeRectWithWidth, lib, "CGContextStrokeRectWithWidth")
	tryRegister(&_CGContextStrokeEllipseInRect, lib, "CGContextStrokeEllipseInRect")
	tryRegister(&_CGContextStrokePath, lib, "CGContextStrokePath")
	tryRegister(&_CGContextSynchronize, lib, "CGContextSynchronize")
	tryRegister(&_CGContextSynchronizeAttributes, lib, "CGContextSynchronizeAttributes")
	tryRegister(&_CGContextGetTextMatrix, lib, "CGContextGetTextMatrix")
	tryRegister(&_CGContextTranslateCTM, lib, "CGContextTranslateCTM")
	tryRegister(&_CGContextGetTypeID, lib, "CGContextGetTypeID")
	tryRegister(&_CGContextGetUserSpaceToDeviceSpaceTransform, lib, "CGContextGetUserSpaceToDeviceSpaceTransform")
	tryRegister(&_CGBitmapContextGetWidth, lib, "CGBitmapContextGetWidth")
	tryRegister(&_CGContextAddArc, lib, "CGContextAddArc")
	tryRegister(&_CGContextAddArcToPoint, lib, "CGContextAddArcToPoint")
	tryRegister(&_CGContextAddCurveToPoint, lib, "CGContextAddCurveToPoint")
	tryRegister(&_CGContextAddLineToPoint, lib, "CGContextAddLineToPoint")
	tryRegister(&_CGContextAddLines, lib, "CGContextAddLines")
	tryRegister(&_CGContextAddQuadCurveToPoint, lib, "CGContextAddQuadCurveToPoint")
	tryRegister(&_CGContextAddRects, lib, "CGContextAddRects")
	tryRegister(&_CGContextClip, lib, "CGContextClip")
	tryRegister(&_CGContextClipToRects, lib, "CGContextClipToRects")
	tryRegister(&_CGContextDrawConicGradient, lib, "CGContextDrawConicGradient")
	tryRegister(&_CGContextDrawImage, lib, "CGContextDrawImage")
	tryRegister(&_CGContextDrawImageApplyingToneMapping, lib, "CGContextDrawImageApplyingToneMapping")
	tryRegister(&_CGContextDrawLayerAtPoint, lib, "CGContextDrawLayerAtPoint")
	tryRegister(&_CGContextDrawLayerInRect, lib, "CGContextDrawLayerInRect")
	tryRegister(&_CGContextDrawPDFDocument, lib, "CGContextDrawPDFDocument")
	tryRegister(&_CGContextDrawTiledImage, lib, "CGContextDrawTiledImage")
	tryRegister(&_CGContextEOClip, lib, "CGContextEOClip")
	tryRegister(&_CGContextEOFillPath, lib, "CGContextEOFillPath")
	tryRegister(&_CGContextFillPath, lib, "CGContextFillPath")
	tryRegister(&_CGContextFillRects, lib, "CGContextFillRects")
	tryRegister(&_CGContextGetContentToneMappingInfo, lib, "CGContextGetContentToneMappingInfo")
	tryRegister(&_CGContextGetEDRTargetHeadroom, lib, "CGContextGetEDRTargetHeadroom")
	tryRegister(&_CGContextGetTextPosition, lib, "CGContextGetTextPosition")
	tryRegister(&_CGContextMoveToPoint, lib, "CGContextMoveToPoint")
	tryRegister(&_CGContextRelease, lib, "CGContextRelease")
	tryRegister(&_CGContextRetain, lib, "CGContextRetain")
	tryRegister(&_CGContextSetContentToneMappingInfo, lib, "CGContextSetContentToneMappingInfo")
	tryRegister(&_CGContextSetLineDash, lib, "CGContextSetLineDash")
	tryRegister(&_CGContextSetTextMatrix, lib, "CGContextSetTextMatrix")
	tryRegister(&_CGContextSetTextPosition, lib, "CGContextSetTextPosition")
	tryRegister(&_CGContextShowGlyphsAtPositions, lib, "CGContextShowGlyphsAtPositions")
	tryRegister(&_CGContextStrokeLineSegments, lib, "CGContextStrokeLineSegments")
	tryRegister(&_CGConvertColorDataWithFormat, lib, "CGConvertColorDataWithFormat")
	tryRegister(&_CGCursorIsDrawnInFramebuffer, lib, "CGCursorIsDrawnInFramebuffer")
	tryRegister(&_CGCursorIsVisible, lib, "CGCursorIsVisible")
	tryRegister(&_CGDataConsumerCreateWithCFData, lib, "CGDataConsumerCreateWithCFData")
	tryRegister(&_CGDataConsumerCreate, lib, "CGDataConsumerCreate")
	tryRegister(&_CGDataConsumerCreateWithURL, lib, "CGDataConsumerCreateWithURL")
	tryRegister(&_CGDataConsumerGetTypeID, lib, "CGDataConsumerGetTypeID")
	tryRegister(&_CGDataConsumerRelease, lib, "CGDataConsumerRelease")
	tryRegister(&_CGDataConsumerRetain, lib, "CGDataConsumerRetain")
	tryRegister(&_CGDataProviderCopyData, lib, "CGDataProviderCopyData")
	tryRegister(&_CGDataProviderGetInfo, lib, "CGDataProviderGetInfo")
	tryRegister(&_CGDataProviderCreateWithCFData, lib, "CGDataProviderCreateWithCFData")
	tryRegister(&_CGDataProviderCreateWithData, lib, "CGDataProviderCreateWithData")
	tryRegister(&_CGDataProviderCreateDirect, lib, "CGDataProviderCreateDirect")
	tryRegister(&_CGDataProviderCreateWithFilename, lib, "CGDataProviderCreateWithFilename")
	tryRegister(&_CGDataProviderCreateSequential, lib, "CGDataProviderCreateSequential")
	tryRegister(&_CGDataProviderCreateWithURL, lib, "CGDataProviderCreateWithURL")
	tryRegister(&_CGDataProviderGetTypeID, lib, "CGDataProviderGetTypeID")
	tryRegister(&_CGDataProviderRelease, lib, "CGDataProviderRelease")
	tryRegister(&_CGDataProviderRetain, lib, "CGDataProviderRetain")
	tryRegister(&_CGDirectDisplayCopyCurrentMetalDevice, lib, "CGDirectDisplayCopyCurrentMetalDevice")
	tryRegister(&_CGDisplayAvailableModes, lib, "CGDisplayAvailableModes")
	tryRegister(&_CGDisplayBestModeForParameters, lib, "CGDisplayBestModeForParameters")
	tryRegister(&_CGDisplayBestModeForParametersAndRefreshRate, lib, "CGDisplayBestModeForParametersAndRefreshRate")
	tryRegister(&_CGDisplayBounds, lib, "CGDisplayBounds")
	tryRegister(&_CGDisplayCapture, lib, "CGDisplayCapture")
	tryRegister(&_CGDisplayCaptureWithOptions, lib, "CGDisplayCaptureWithOptions")
	tryRegister(&_CGDisplayCopyAllDisplayModes, lib, "CGDisplayCopyAllDisplayModes")
	tryRegister(&_CGDisplayCopyColorSpace, lib, "CGDisplayCopyColorSpace")
	tryRegister(&_CGDisplayCopyDisplayMode, lib, "CGDisplayCopyDisplayMode")
	tryRegister(&_CGDisplayCreateImage, lib, "CGDisplayCreateImage")
	tryRegister(&_CGDisplayCreateImageForRect, lib, "CGDisplayCreateImageForRect")
	tryRegister(&_CGDisplayCurrentMode, lib, "CGDisplayCurrentMode")
	tryRegister(&_CGDisplayFade, lib, "CGDisplayFade")
	tryRegister(&_CGDisplayFadeOperationInProgress, lib, "CGDisplayFadeOperationInProgress")
	tryRegister(&_CGDisplayGammaTableCapacity, lib, "CGDisplayGammaTableCapacity")
	tryRegister(&_CGDisplayGetDrawingContext, lib, "CGDisplayGetDrawingContext")
	tryRegister(&_CGDisplayHideCursor, lib, "CGDisplayHideCursor")
	tryRegister(&_CGDisplayIDToOpenGLDisplayMask, lib, "CGDisplayIDToOpenGLDisplayMask")
	tryRegister(&_CGDisplayIOServicePort, lib, "CGDisplayIOServicePort")
	tryRegister(&_CGDisplayIsActive, lib, "CGDisplayIsActive")
	tryRegister(&_CGDisplayIsAlwaysInMirrorSet, lib, "CGDisplayIsAlwaysInMirrorSet")
	tryRegister(&_CGDisplayIsAsleep, lib, "CGDisplayIsAsleep")
	tryRegister(&_CGDisplayIsBuiltin, lib, "CGDisplayIsBuiltin")
	tryRegister(&_CGDisplayIsCaptured, lib, "CGDisplayIsCaptured")
	tryRegister(&_CGDisplayIsInHWMirrorSet, lib, "CGDisplayIsInHWMirrorSet")
	tryRegister(&_CGDisplayIsInMirrorSet, lib, "CGDisplayIsInMirrorSet")
	tryRegister(&_CGDisplayIsMain, lib, "CGDisplayIsMain")
	tryRegister(&_CGDisplayIsOnline, lib, "CGDisplayIsOnline")
	tryRegister(&_CGDisplayIsStereo, lib, "CGDisplayIsStereo")
	tryRegister(&_CGDisplayMirrorsDisplay, lib, "CGDisplayMirrorsDisplay")
	tryRegister(&_CGDisplayModeGetHeight, lib, "CGDisplayModeGetHeight")
	tryRegister(&_CGDisplayModeGetIODisplayModeID, lib, "CGDisplayModeGetIODisplayModeID")
	tryRegister(&_CGDisplayModeGetIOFlags, lib, "CGDisplayModeGetIOFlags")
	tryRegister(&_CGDisplayModeIsUsableForDesktopGUI, lib, "CGDisplayModeIsUsableForDesktopGUI")
	tryRegister(&_CGDisplayModeCopyPixelEncoding, lib, "CGDisplayModeCopyPixelEncoding")
	tryRegister(&_CGDisplayModeGetPixelHeight, lib, "CGDisplayModeGetPixelHeight")
	tryRegister(&_CGDisplayModeGetPixelWidth, lib, "CGDisplayModeGetPixelWidth")
	tryRegister(&_CGDisplayModeGetRefreshRate, lib, "CGDisplayModeGetRefreshRate")
	tryRegister(&_CGDisplayModeGetTypeID, lib, "CGDisplayModeGetTypeID")
	tryRegister(&_CGDisplayModeGetWidth, lib, "CGDisplayModeGetWidth")
	tryRegister(&_CGDisplayModeRelease, lib, "CGDisplayModeRelease")
	tryRegister(&_CGDisplayModeRetain, lib, "CGDisplayModeRetain")
	tryRegister(&_CGDisplayModelNumber, lib, "CGDisplayModelNumber")
	tryRegister(&_CGDisplayMoveCursorToPoint, lib, "CGDisplayMoveCursorToPoint")
	tryRegister(&_CGDisplayPixelsHigh, lib, "CGDisplayPixelsHigh")
	tryRegister(&_CGDisplayPixelsWide, lib, "CGDisplayPixelsWide")
	tryRegister(&_CGDisplayPrimaryDisplay, lib, "CGDisplayPrimaryDisplay")
	tryRegister(&_CGDisplayRegisterReconfigurationCallback, lib, "CGDisplayRegisterReconfigurationCallback")
	tryRegister(&_CGDisplayRelease, lib, "CGDisplayRelease")
	tryRegister(&_CGDisplayRemoveReconfigurationCallback, lib, "CGDisplayRemoveReconfigurationCallback")
	tryRegister(&_CGDisplayRestoreColorSyncSettings, lib, "CGDisplayRestoreColorSyncSettings")
	tryRegister(&_CGDisplayRotation, lib, "CGDisplayRotation")
	tryRegister(&_CGDisplayScreenSize, lib, "CGDisplayScreenSize")
	tryRegister(&_CGDisplaySerialNumber, lib, "CGDisplaySerialNumber")
	tryRegister(&_CGDisplaySetDisplayMode, lib, "CGDisplaySetDisplayMode")
	tryRegister(&_CGDisplaySetStereoOperation, lib, "CGDisplaySetStereoOperation")
	tryRegister(&_CGDisplayShowCursor, lib, "CGDisplayShowCursor")
	tryRegister(&_CGDisplayStreamCreateWithDispatchQueue, lib, "CGDisplayStreamCreateWithDispatchQueue")
	tryRegister(&_CGDisplayStreamCreate, lib, "CGDisplayStreamCreate")
	tryRegister(&_CGDisplayStreamGetRunLoopSource, lib, "CGDisplayStreamGetRunLoopSource")
	tryRegister(&_CGDisplayStreamStart, lib, "CGDisplayStreamStart")
	tryRegister(&_CGDisplayStreamStop, lib, "CGDisplayStreamStop")
	tryRegister(&_CGDisplayStreamGetTypeID, lib, "CGDisplayStreamGetTypeID")
	tryRegister(&_CGDisplayStreamUpdateGetDropCount, lib, "CGDisplayStreamUpdateGetDropCount")
	tryRegister(&_CGDisplayStreamUpdateGetMovedRectsDelta, lib, "CGDisplayStreamUpdateGetMovedRectsDelta")
	tryRegister(&_CGDisplayStreamUpdateGetRects, lib, "CGDisplayStreamUpdateGetRects")
	tryRegister(&_CGDisplayStreamUpdateCreateMergedUpdate, lib, "CGDisplayStreamUpdateCreateMergedUpdate")
	tryRegister(&_CGDisplayStreamUpdateGetTypeID, lib, "CGDisplayStreamUpdateGetTypeID")
	tryRegister(&_CGDisplaySwitchToMode, lib, "CGDisplaySwitchToMode")
	tryRegister(&_CGDisplayUnitNumber, lib, "CGDisplayUnitNumber")
	tryRegister(&_CGDisplayUsesOpenGLAcceleration, lib, "CGDisplayUsesOpenGLAcceleration")
	tryRegister(&_CGDisplayVendorNumber, lib, "CGDisplayVendorNumber")
	tryRegister(&_CGEXRToneMappingGammaGetDefaultOptions, lib, "CGEXRToneMappingGammaGetDefaultOptions")
	tryRegister(&_CGEnableEventStateCombining, lib, "CGEnableEventStateCombining")
	tryRegister(&_CGErrorSetCallback, lib, "CGErrorSetCallback")
	tryRegister(&_CGEventCreateCopy, lib, "CGEventCreateCopy")
	tryRegister(&_CGEventGetFlags, lib, "CGEventGetFlags")
	tryRegister(&_CGEventGetDoubleValueField, lib, "CGEventGetDoubleValueField")
	tryRegister(&_CGEventGetIntegerValueField, lib, "CGEventGetIntegerValueField")
	tryRegister(&_CGEventCreateKeyboardEvent, lib, "CGEventCreateKeyboardEvent")
	tryRegister(&_CGEventCreateMouseEvent, lib, "CGEventCreateMouseEvent")
	tryRegister(&_CGEventCreateScrollWheelEvent2, lib, "CGEventCreateScrollWheelEvent2")
	tryRegister(&_CGEventCreate, lib, "CGEventCreate")
	tryRegister(&_CGEventCreateFromData, lib, "CGEventCreateFromData")
	tryRegister(&_CGEventKeyboardGetUnicodeString, lib, "CGEventKeyboardGetUnicodeString")
	tryRegister(&_CGEventKeyboardSetUnicodeString, lib, "CGEventKeyboardSetUnicodeString")
	tryRegister(&_CGEventGetLocation, lib, "CGEventGetLocation")
	tryRegister(&_CGEventPost, lib, "CGEventPost")
	tryRegister(&_CGEventPostToPSN, lib, "CGEventPostToPSN")
	tryRegister(&_CGEventPostToPid, lib, "CGEventPostToPid")
	tryRegister(&_CGEventSetDoubleValueField, lib, "CGEventSetDoubleValueField")
	tryRegister(&_CGEventSetIntegerValueField, lib, "CGEventSetIntegerValueField")
	tryRegister(&_CGEventSetSource, lib, "CGEventSetSource")
	tryRegister(&_CGEventTapCreate, lib, "CGEventTapCreate")
	tryRegister(&_CGEventTapCreateForPSN, lib, "CGEventTapCreateForPSN")
	tryRegister(&_CGEventTapCreateForPid, lib, "CGEventTapCreateForPid")
	tryRegister(&_CGEventTapEnable, lib, "CGEventTapEnable")
	tryRegister(&_CGEventTapIsEnabled, lib, "CGEventTapIsEnabled")
	tryRegister(&_CGEventTapPostEvent, lib, "CGEventTapPostEvent")
	tryRegister(&_CGEventGetTimestamp, lib, "CGEventGetTimestamp")
	tryRegister(&_CGEventGetType, lib, "CGEventGetType")
	tryRegister(&_CGEventGetTypeID, lib, "CGEventGetTypeID")
	tryRegister(&_CGEventGetUnflippedLocation, lib, "CGEventGetUnflippedLocation")
	tryRegister(&_CGEventCreateData, lib, "CGEventCreateData")
	tryRegister(&_CGEventCreateScrollWheelEvent, lib, "CGEventCreateScrollWheelEvent")
	tryRegister(&_CGEventSetFlags, lib, "CGEventSetFlags")
	tryRegister(&_CGEventSetLocation, lib, "CGEventSetLocation")
	tryRegister(&_CGEventSetTimestamp, lib, "CGEventSetTimestamp")
	tryRegister(&_CGEventSetType, lib, "CGEventSetType")
	tryRegister(&_CGEventSourceButtonState, lib, "CGEventSourceButtonState")
	tryRegister(&_CGEventSourceCounterForEventType, lib, "CGEventSourceCounterForEventType")
	tryRegister(&_CGEventSourceFlagsState, lib, "CGEventSourceFlagsState")
	tryRegister(&_CGEventSourceGetLocalEventsFilterDuringSuppressionState, lib, "CGEventSourceGetLocalEventsFilterDuringSuppressionState")
	tryRegister(&_CGEventCreateSourceFromEvent, lib, "CGEventCreateSourceFromEvent")
	tryRegister(&_CGEventSourceCreate, lib, "CGEventSourceCreate")
	tryRegister(&_CGEventSourceKeyState, lib, "CGEventSourceKeyState")
	tryRegister(&_CGEventSourceGetKeyboardType, lib, "CGEventSourceGetKeyboardType")
	tryRegister(&_CGEventSourceGetLocalEventsSuppressionInterval, lib, "CGEventSourceGetLocalEventsSuppressionInterval")
	tryRegister(&_CGEventSourceGetPixelsPerLine, lib, "CGEventSourceGetPixelsPerLine")
	tryRegister(&_CGEventSourceSecondsSinceLastEventType, lib, "CGEventSourceSecondsSinceLastEventType")
	tryRegister(&_CGEventSourceSetLocalEventsFilterDuringSuppressionState, lib, "CGEventSourceSetLocalEventsFilterDuringSuppressionState")
	tryRegister(&_CGEventSourceGetSourceStateID, lib, "CGEventSourceGetSourceStateID")
	tryRegister(&_CGEventSourceGetTypeID, lib, "CGEventSourceGetTypeID")
	tryRegister(&_CGEventSourceGetUserData, lib, "CGEventSourceGetUserData")
	tryRegister(&_CGEventSourceSetKeyboardType, lib, "CGEventSourceSetKeyboardType")
	tryRegister(&_CGEventSourceSetLocalEventsSuppressionInterval, lib, "CGEventSourceSetLocalEventsSuppressionInterval")
	tryRegister(&_CGEventSourceSetPixelsPerLine, lib, "CGEventSourceSetPixelsPerLine")
	tryRegister(&_CGEventSourceSetUserData, lib, "CGEventSourceSetUserData")
	tryRegister(&_CGFontGetAscent, lib, "CGFontGetAscent")
	tryRegister(&_CGFontCanCreatePostScriptSubset, lib, "CGFontCanCreatePostScriptSubset")
	tryRegister(&_CGFontGetCapHeight, lib, "CGFontGetCapHeight")
	tryRegister(&_CGFontCreateCopyWithVariations, lib, "CGFontCreateCopyWithVariations")
	tryRegister(&_CGFontCreatePostScriptEncoding, lib, "CGFontCreatePostScriptEncoding")
	tryRegister(&_CGFontCreatePostScriptSubset, lib, "CGFontCreatePostScriptSubset")
	tryRegister(&_CGFontGetDescent, lib, "CGFontGetDescent")
	tryRegister(&_CGFontGetFontBBox, lib, "CGFontGetFontBBox")
	tryRegister(&_CGFontCopyFullName, lib, "CGFontCopyFullName")
	tryRegister(&_CGFontGetGlyphAdvances, lib, "CGFontGetGlyphAdvances")
	tryRegister(&_CGFontGetGlyphBBoxes, lib, "CGFontGetGlyphBBoxes")
	tryRegister(&_CGFontGetGlyphWithGlyphName, lib, "CGFontGetGlyphWithGlyphName")
	tryRegister(&_CGFontCreateWithFontName, lib, "CGFontCreateWithFontName")
	tryRegister(&_CGFontCreateWithDataProvider, lib, "CGFontCreateWithDataProvider")
	tryRegister(&_CGFontGetItalicAngle, lib, "CGFontGetItalicAngle")
	tryRegister(&_CGFontGetLeading, lib, "CGFontGetLeading")
	tryRegister(&_CGFontCopyGlyphNameForGlyph, lib, "CGFontCopyGlyphNameForGlyph")
	tryRegister(&_CGFontGetNumberOfGlyphs, lib, "CGFontGetNumberOfGlyphs")
	tryRegister(&_CGFontCopyPostScriptName, lib, "CGFontCopyPostScriptName")
	tryRegister(&_CGFontGetStemV, lib, "CGFontGetStemV")
	tryRegister(&_CGFontCopyTableForTag, lib, "CGFontCopyTableForTag")
	tryRegister(&_CGFontCopyTableTags, lib, "CGFontCopyTableTags")
	tryRegister(&_CGFontGetTypeID, lib, "CGFontGetTypeID")
	tryRegister(&_CGFontGetUnitsPerEm, lib, "CGFontGetUnitsPerEm")
	tryRegister(&_CGFontCopyVariationAxes, lib, "CGFontCopyVariationAxes")
	tryRegister(&_CGFontCopyVariations, lib, "CGFontCopyVariations")
	tryRegister(&_CGFontGetXHeight, lib, "CGFontGetXHeight")
	tryRegister(&_CGFontCreateWithPlatformFont, lib, "CGFontCreateWithPlatformFont")
	tryRegister(&_CGFontRelease, lib, "CGFontRelease")
	tryRegister(&_CGFontRetain, lib, "CGFontRetain")
	tryRegister(&_CGFunctionCreate, lib, "CGFunctionCreate")
	tryRegister(&_CGFunctionGetTypeID, lib, "CGFunctionGetTypeID")
	tryRegister(&_CGFunctionRelease, lib, "CGFunctionRelease")
	tryRegister(&_CGFunctionRetain, lib, "CGFunctionRetain")
	tryRegister(&_CGGetActiveDisplayList, lib, "CGGetActiveDisplayList")
	tryRegister(&_CGGetDisplayTransferByFormula, lib, "CGGetDisplayTransferByFormula")
	tryRegister(&_CGGetDisplayTransferByTable, lib, "CGGetDisplayTransferByTable")
	tryRegister(&_CGGetDisplaysWithOpenGLDisplayMask, lib, "CGGetDisplaysWithOpenGLDisplayMask")
	tryRegister(&_CGGetDisplaysWithPoint, lib, "CGGetDisplaysWithPoint")
	tryRegister(&_CGGetDisplaysWithRect, lib, "CGGetDisplaysWithRect")
	tryRegister(&_CGGetEventTapList, lib, "CGGetEventTapList")
	tryRegister(&_CGGetLastMouseDelta, lib, "CGGetLastMouseDelta")
	tryRegister(&_CGGetOnlineDisplayList, lib, "CGGetOnlineDisplayList")
	tryRegister(&_CGGradientGetContentHeadroom, lib, "CGGradientGetContentHeadroom")
	tryRegister(&_CGGradientCreateWithColorComponents, lib, "CGGradientCreateWithColorComponents")
	tryRegister(&_CGGradientCreateWithColors, lib, "CGGradientCreateWithColors")
	tryRegister(&_CGGradientCreateWithContentHeadroom, lib, "CGGradientCreateWithContentHeadroom")
	tryRegister(&_CGGradientGetTypeID, lib, "CGGradientGetTypeID")
	tryRegister(&_CGGradientRelease, lib, "CGGradientRelease")
	tryRegister(&_CGGradientRetain, lib, "CGGradientRetain")
	tryRegister(&_CGImageGetAlphaInfo, lib, "CGImageGetAlphaInfo")
	tryRegister(&_CGImageGetBitmapInfo, lib, "CGImageGetBitmapInfo")
	tryRegister(&_CGImageGetBitsPerComponent, lib, "CGImageGetBitsPerComponent")
	tryRegister(&_CGImageGetBitsPerPixel, lib, "CGImageGetBitsPerPixel")
	tryRegister(&_CGImageGetByteOrderInfo, lib, "CGImageGetByteOrderInfo")
	tryRegister(&_CGImageGetBytesPerRow, lib, "CGImageGetBytesPerRow")
	tryRegister(&_CGImageCalculateContentAverageLightLevel, lib, "CGImageCalculateContentAverageLightLevel")
	tryRegister(&_CGImageCalculateContentHeadroom, lib, "CGImageCalculateContentHeadroom")
	tryRegister(&_CGImageGetColorSpace, lib, "CGImageGetColorSpace")
	tryRegister(&_CGImageContainsImageSpecificToneMappingMetadata, lib, "CGImageContainsImageSpecificToneMappingMetadata")
	tryRegister(&_CGImageGetContentAverageLightLevel, lib, "CGImageGetContentAverageLightLevel")
	tryRegister(&_CGImageGetContentHeadroom, lib, "CGImageGetContentHeadroom")
	tryRegister(&_CGImageCreateCopy, lib, "CGImageCreateCopy")
	tryRegister(&_CGImageCreateCopyWithColorSpace, lib, "CGImageCreateCopyWithColorSpace")
	tryRegister(&_CGImageCreateCopyWithContentAverageLightLevel, lib, "CGImageCreateCopyWithContentAverageLightLevel")
	tryRegister(&_CGImageCreateCopyWithCalculatedHDRStats, lib, "CGImageCreateCopyWithCalculatedHDRStats")
	tryRegister(&_CGImageCreateWithImageInRect, lib, "CGImageCreateWithImageInRect")
	tryRegister(&_CGImageGetDataProvider, lib, "CGImageGetDataProvider")
	tryRegister(&_CGImageGetDecode, lib, "CGImageGetDecode")
	tryRegister(&_CGImageGetHeight, lib, "CGImageGetHeight")
	tryRegister(&_CGImageCreateWithContentHeadroom, lib, "CGImageCreateWithContentHeadroom")
	tryRegister(&_CGImageCreateWithJPEGDataProvider, lib, "CGImageCreateWithJPEGDataProvider")
	tryRegister(&_CGImageMaskCreate, lib, "CGImageMaskCreate")
	tryRegister(&_CGImageCreateWithPNGDataProvider, lib, "CGImageCreateWithPNGDataProvider")
	tryRegister(&_CGImageCreate, lib, "CGImageCreate")
	tryRegister(&_CGWindowListCreateImageFromArray, lib, "CGWindowListCreateImageFromArray")
	tryRegister(&_CGImageIsMask, lib, "CGImageIsMask")
	tryRegister(&_CGImageCreateWithMask, lib, "CGImageCreateWithMask")
	tryRegister(&_CGImageGetPixelFormatInfo, lib, "CGImageGetPixelFormatInfo")
	tryRegister(&_CGImageGetRenderingIntent, lib, "CGImageGetRenderingIntent")
	tryRegister(&_CGImageGetShouldInterpolate, lib, "CGImageGetShouldInterpolate")
	tryRegister(&_CGImageShouldToneMap, lib, "CGImageShouldToneMap")
	tryRegister(&_CGImageGetTypeID, lib, "CGImageGetTypeID")
	tryRegister(&_CGImageGetUTType, lib, "CGImageGetUTType")
	tryRegister(&_CGImageGetWidth, lib, "CGImageGetWidth")
	tryRegister(&_CGImageCreateCopyWithContentHeadroom, lib, "CGImageCreateCopyWithContentHeadroom")
	tryRegister(&_CGImageCreateWithMaskingColors, lib, "CGImageCreateWithMaskingColors")
	tryRegister(&_CGImageRelease, lib, "CGImageRelease")
	tryRegister(&_CGImageRetain, lib, "CGImageRetain")
	tryRegister(&_CGInhibitLocalEvents, lib, "CGInhibitLocalEvents")
	tryRegister(&_CGLayerGetContext, lib, "CGLayerGetContext")
	tryRegister(&_CGLayerCreateWithContext, lib, "CGLayerCreateWithContext")
	tryRegister(&_CGLayerGetSize, lib, "CGLayerGetSize")
	tryRegister(&_CGLayerGetTypeID, lib, "CGLayerGetTypeID")
	tryRegister(&_CGLayerRelease, lib, "CGLayerRelease")
	tryRegister(&_CGLayerRetain, lib, "CGLayerRetain")
	tryRegister(&_CGMainDisplayID, lib, "CGMainDisplayID")
	tryRegister(&_CGPathCloseSubpath, lib, "CGPathCloseSubpath")
	tryRegister(&_CGPathCreateMutable, lib, "CGPathCreateMutable")
	tryRegister(&_CGOpenGLDisplayMaskToDisplayID, lib, "CGOpenGLDisplayMaskToDisplayID")
	tryRegister(&_CGPDFArrayApplyBlock, lib, "CGPDFArrayApplyBlock")
	tryRegister(&_CGPDFArrayGetArray, lib, "CGPDFArrayGetArray")
	tryRegister(&_CGPDFArrayGetBoolean, lib, "CGPDFArrayGetBoolean")
	tryRegister(&_CGPDFArrayGetCount, lib, "CGPDFArrayGetCount")
	tryRegister(&_CGPDFArrayGetDictionary, lib, "CGPDFArrayGetDictionary")
	tryRegister(&_CGPDFArrayGetInteger, lib, "CGPDFArrayGetInteger")
	tryRegister(&_CGPDFArrayGetName, lib, "CGPDFArrayGetName")
	tryRegister(&_CGPDFArrayGetNull, lib, "CGPDFArrayGetNull")
	tryRegister(&_CGPDFArrayGetNumber, lib, "CGPDFArrayGetNumber")
	tryRegister(&_CGPDFArrayGetObject, lib, "CGPDFArrayGetObject")
	tryRegister(&_CGPDFArrayGetStream, lib, "CGPDFArrayGetStream")
	tryRegister(&_CGPDFArrayGetString, lib, "CGPDFArrayGetString")
	tryRegister(&_CGPDFContentStreamCreateWithPage, lib, "CGPDFContentStreamCreateWithPage")
	tryRegister(&_CGPDFContentStreamCreateWithStream, lib, "CGPDFContentStreamCreateWithStream")
	tryRegister(&_CGPDFContentStreamGetResource, lib, "CGPDFContentStreamGetResource")
	tryRegister(&_CGPDFContentStreamGetStreams, lib, "CGPDFContentStreamGetStreams")
	tryRegister(&_CGPDFContentStreamRelease, lib, "CGPDFContentStreamRelease")
	tryRegister(&_CGPDFContentStreamRetain, lib, "CGPDFContentStreamRetain")
	tryRegister(&_CGPDFContextBeginTag, lib, "CGPDFContextBeginTag")
	tryRegister(&_CGPDFContextEndTag, lib, "CGPDFContextEndTag")
	tryRegister(&_CGPDFContextSetIDTree, lib, "CGPDFContextSetIDTree")
	tryRegister(&_CGPDFContextSetOutline, lib, "CGPDFContextSetOutline")
	tryRegister(&_CGPDFContextSetPageTagStructureTree, lib, "CGPDFContextSetPageTagStructureTree")
	tryRegister(&_CGPDFContextSetParentTree, lib, "CGPDFContextSetParentTree")
	tryRegister(&_CGPDFDictionaryApplyBlock, lib, "CGPDFDictionaryApplyBlock")
	tryRegister(&_CGPDFDictionaryApplyFunction, lib, "CGPDFDictionaryApplyFunction")
	tryRegister(&_CGPDFDictionaryGetArray, lib, "CGPDFDictionaryGetArray")
	tryRegister(&_CGPDFDictionaryGetBoolean, lib, "CGPDFDictionaryGetBoolean")
	tryRegister(&_CGPDFDictionaryGetCount, lib, "CGPDFDictionaryGetCount")
	tryRegister(&_CGPDFDictionaryGetDictionary, lib, "CGPDFDictionaryGetDictionary")
	tryRegister(&_CGPDFDictionaryGetInteger, lib, "CGPDFDictionaryGetInteger")
	tryRegister(&_CGPDFDictionaryGetName, lib, "CGPDFDictionaryGetName")
	tryRegister(&_CGPDFDictionaryGetNumber, lib, "CGPDFDictionaryGetNumber")
	tryRegister(&_CGPDFDictionaryGetObject, lib, "CGPDFDictionaryGetObject")
	tryRegister(&_CGPDFDictionaryGetStream, lib, "CGPDFDictionaryGetStream")
	tryRegister(&_CGPDFDictionaryGetString, lib, "CGPDFDictionaryGetString")
	tryRegister(&_CGPDFDocumentGetAccessPermissions, lib, "CGPDFDocumentGetAccessPermissions")
	tryRegister(&_CGPDFDocumentAllowsCopying, lib, "CGPDFDocumentAllowsCopying")
	tryRegister(&_CGPDFDocumentAllowsPrinting, lib, "CGPDFDocumentAllowsPrinting")
	tryRegister(&_CGPDFDocumentGetCatalog, lib, "CGPDFDocumentGetCatalog")
	tryRegister(&_CGPDFDocumentGetID, lib, "CGPDFDocumentGetID")
	tryRegister(&_CGPDFDocumentGetVersion, lib, "CGPDFDocumentGetVersion")
	tryRegister(&_CGPDFDocumentGetInfo, lib, "CGPDFDocumentGetInfo")
	tryRegister(&_CGPDFDocumentCreateWithURL, lib, "CGPDFDocumentCreateWithURL")
	tryRegister(&_CGPDFDocumentCreateWithProvider, lib, "CGPDFDocumentCreateWithProvider")
	tryRegister(&_CGPDFDocumentIsEncrypted, lib, "CGPDFDocumentIsEncrypted")
	tryRegister(&_CGPDFDocumentIsUnlocked, lib, "CGPDFDocumentIsUnlocked")
	tryRegister(&_CGPDFDocumentGetNumberOfPages, lib, "CGPDFDocumentGetNumberOfPages")
	tryRegister(&_CGPDFDocumentGetOutline, lib, "CGPDFDocumentGetOutline")
	tryRegister(&_CGPDFDocumentGetPage, lib, "CGPDFDocumentGetPage")
	tryRegister(&_CGPDFDocumentGetTypeID, lib, "CGPDFDocumentGetTypeID")
	tryRegister(&_CGPDFDocumentUnlockWithPassword, lib, "CGPDFDocumentUnlockWithPassword")
	tryRegister(&_CGPDFDocumentGetArtBox, lib, "CGPDFDocumentGetArtBox")
	tryRegister(&_CGPDFDocumentGetBleedBox, lib, "CGPDFDocumentGetBleedBox")
	tryRegister(&_CGPDFDocumentGetCropBox, lib, "CGPDFDocumentGetCropBox")
	tryRegister(&_CGPDFDocumentGetMediaBox, lib, "CGPDFDocumentGetMediaBox")
	tryRegister(&_CGPDFDocumentGetRotationAngle, lib, "CGPDFDocumentGetRotationAngle")
	tryRegister(&_CGPDFDocumentGetTrimBox, lib, "CGPDFDocumentGetTrimBox")
	tryRegister(&_CGPDFDocumentRelease, lib, "CGPDFDocumentRelease")
	tryRegister(&_CGPDFDocumentRetain, lib, "CGPDFDocumentRetain")
	tryRegister(&_CGPDFObjectGetType, lib, "CGPDFObjectGetType")
	tryRegister(&_CGPDFObjectGetValue, lib, "CGPDFObjectGetValue")
	tryRegister(&_CGPDFOperatorTableCreate, lib, "CGPDFOperatorTableCreate")
	tryRegister(&_CGPDFOperatorTableRelease, lib, "CGPDFOperatorTableRelease")
	tryRegister(&_CGPDFOperatorTableRetain, lib, "CGPDFOperatorTableRetain")
	tryRegister(&_CGPDFOperatorTableSetCallback, lib, "CGPDFOperatorTableSetCallback")
	tryRegister(&_CGPDFPageGetDictionary, lib, "CGPDFPageGetDictionary")
	tryRegister(&_CGPDFPageGetDocument, lib, "CGPDFPageGetDocument")
	tryRegister(&_CGPDFPageGetBoxRect, lib, "CGPDFPageGetBoxRect")
	tryRegister(&_CGPDFPageGetDrawingTransform, lib, "CGPDFPageGetDrawingTransform")
	tryRegister(&_CGPDFPageGetPageNumber, lib, "CGPDFPageGetPageNumber")
	tryRegister(&_CGPDFPageGetRotationAngle, lib, "CGPDFPageGetRotationAngle")
	tryRegister(&_CGPDFPageGetTypeID, lib, "CGPDFPageGetTypeID")
	tryRegister(&_CGPDFPageRelease, lib, "CGPDFPageRelease")
	tryRegister(&_CGPDFPageRetain, lib, "CGPDFPageRetain")
	tryRegister(&_CGPDFScannerCreate, lib, "CGPDFScannerCreate")
	tryRegister(&_CGPDFScannerGetContentStream, lib, "CGPDFScannerGetContentStream")
	tryRegister(&_CGPDFScannerPopArray, lib, "CGPDFScannerPopArray")
	tryRegister(&_CGPDFScannerPopBoolean, lib, "CGPDFScannerPopBoolean")
	tryRegister(&_CGPDFScannerPopDictionary, lib, "CGPDFScannerPopDictionary")
	tryRegister(&_CGPDFScannerPopInteger, lib, "CGPDFScannerPopInteger")
	tryRegister(&_CGPDFScannerPopName, lib, "CGPDFScannerPopName")
	tryRegister(&_CGPDFScannerPopNumber, lib, "CGPDFScannerPopNumber")
	tryRegister(&_CGPDFScannerPopObject, lib, "CGPDFScannerPopObject")
	tryRegister(&_CGPDFScannerPopStream, lib, "CGPDFScannerPopStream")
	tryRegister(&_CGPDFScannerPopString, lib, "CGPDFScannerPopString")
	tryRegister(&_CGPDFScannerRelease, lib, "CGPDFScannerRelease")
	tryRegister(&_CGPDFScannerRetain, lib, "CGPDFScannerRetain")
	tryRegister(&_CGPDFScannerScan, lib, "CGPDFScannerScan")
	tryRegister(&_CGPDFScannerStop, lib, "CGPDFScannerStop")
	tryRegister(&_CGPDFStreamCopyData, lib, "CGPDFStreamCopyData")
	tryRegister(&_CGPDFStreamGetDictionary, lib, "CGPDFStreamGetDictionary")
	tryRegister(&_CGPDFStringCopyDate, lib, "CGPDFStringCopyDate")
	tryRegister(&_CGPDFStringCopyTextString, lib, "CGPDFStringCopyTextString")
	tryRegister(&_CGPDFStringGetBytePtr, lib, "CGPDFStringGetBytePtr")
	tryRegister(&_CGPDFStringGetLength, lib, "CGPDFStringGetLength")
	tryRegister(&_CGPDFTagTypeGetName, lib, "CGPDFTagTypeGetName")
	tryRegister(&_CGPSConverterAbort, lib, "CGPSConverterAbort")
	tryRegister(&_CGPSConverterConvert, lib, "CGPSConverterConvert")
	tryRegister(&_CGPSConverterCreate, lib, "CGPSConverterCreate")
	tryRegister(&_CGPSConverterIsConverting, lib, "CGPSConverterIsConverting")
	tryRegister(&_CGPSConverterGetTypeID, lib, "CGPSConverterGetTypeID")
	tryRegister(&_CGPathApply, lib, "CGPathApply")
	tryRegister(&_CGPathApplyWithBlock, lib, "CGPathApplyWithBlock")
	tryRegister(&_CGPathGetBoundingBox, lib, "CGPathGetBoundingBox")
	tryRegister(&_CGPathGetPathBoundingBox, lib, "CGPathGetPathBoundingBox")
	tryRegister(&_CGPathCreateCopy, lib, "CGPathCreateCopy")
	tryRegister(&_CGPathCreateCopyByTransformingPath, lib, "CGPathCreateCopyByTransformingPath")
	tryRegister(&_CGPathGetCurrentPoint, lib, "CGPathGetCurrentPoint")
	tryRegister(&_CGPathCreateWithEllipseInRect, lib, "CGPathCreateWithEllipseInRect")
	tryRegister(&_CGPathCreateWithRect, lib, "CGPathCreateWithRect")
	tryRegister(&_CGPathCreateWithRoundedRect, lib, "CGPathCreateWithRoundedRect")
	tryRegister(&_CGPathIsEmpty, lib, "CGPathIsEmpty")
	tryRegister(&_CGPathIsRect, lib, "CGPathIsRect")
	tryRegister(&_CGPathCreateMutableCopy, lib, "CGPathCreateMutableCopy")
	tryRegister(&_CGPathCreateMutableCopyByTransformingPath, lib, "CGPathCreateMutableCopyByTransformingPath")
	tryRegister(&_CGPathGetTypeID, lib, "CGPathGetTypeID")
	tryRegister(&_CGPathAddArc, lib, "CGPathAddArc")
	tryRegister(&_CGPathAddArcToPoint, lib, "CGPathAddArcToPoint")
	tryRegister(&_CGPathAddCurveToPoint, lib, "CGPathAddCurveToPoint")
	tryRegister(&_CGPathAddEllipseInRect, lib, "CGPathAddEllipseInRect")
	tryRegister(&_CGPathAddLineToPoint, lib, "CGPathAddLineToPoint")
	tryRegister(&_CGPathAddLines, lib, "CGPathAddLines")
	tryRegister(&_CGPathAddPath, lib, "CGPathAddPath")
	tryRegister(&_CGPathAddQuadCurveToPoint, lib, "CGPathAddQuadCurveToPoint")
	tryRegister(&_CGPathAddRect, lib, "CGPathAddRect")
	tryRegister(&_CGPathAddRects, lib, "CGPathAddRects")
	tryRegister(&_CGPathAddRelativeArc, lib, "CGPathAddRelativeArc")
	tryRegister(&_CGPathAddRoundedRect, lib, "CGPathAddRoundedRect")
	tryRegister(&_CGPathContainsPoint, lib, "CGPathContainsPoint")
	tryRegister(&_CGPathCreateCopyByDashingPath, lib, "CGPathCreateCopyByDashingPath")
	tryRegister(&_CGPathCreateCopyByFlattening, lib, "CGPathCreateCopyByFlattening")
	tryRegister(&_CGPathCreateCopyByIntersectingPath, lib, "CGPathCreateCopyByIntersectingPath")
	tryRegister(&_CGPathCreateCopyByNormalizing, lib, "CGPathCreateCopyByNormalizing")
	tryRegister(&_CGPathCreateCopyByStrokingPath, lib, "CGPathCreateCopyByStrokingPath")
	tryRegister(&_CGPathCreateCopyBySubtractingPath, lib, "CGPathCreateCopyBySubtractingPath")
	tryRegister(&_CGPathCreateCopyBySymmetricDifferenceOfPath, lib, "CGPathCreateCopyBySymmetricDifferenceOfPath")
	tryRegister(&_CGPathCreateCopyByUnioningPath, lib, "CGPathCreateCopyByUnioningPath")
	tryRegister(&_CGPathCreateCopyOfLineByIntersectingPath, lib, "CGPathCreateCopyOfLineByIntersectingPath")
	tryRegister(&_CGPathCreateCopyOfLineBySubtractingPath, lib, "CGPathCreateCopyOfLineBySubtractingPath")
	tryRegister(&_CGPathCreateSeparateComponents, lib, "CGPathCreateSeparateComponents")
	tryRegister(&_CGPathEqualToPath, lib, "CGPathEqualToPath")
	tryRegister(&_CGPathIntersectsPath, lib, "CGPathIntersectsPath")
	tryRegister(&_CGPathMoveToPoint, lib, "CGPathMoveToPoint")
	tryRegister(&_CGPathRelease, lib, "CGPathRelease")
	tryRegister(&_CGPathRetain, lib, "CGPathRetain")
	tryRegister(&_CGPatternCreate, lib, "CGPatternCreate")
	tryRegister(&_CGPatternGetTypeID, lib, "CGPatternGetTypeID")
	tryRegister(&_CGPatternRelease, lib, "CGPatternRelease")
	tryRegister(&_CGPatternRetain, lib, "CGPatternRetain")
	tryRegister(&_CGPointApplyAffineTransform, lib, "CGPointApplyAffineTransform")
	tryRegister(&_CGPointCreateDictionaryRepresentation, lib, "CGPointCreateDictionaryRepresentation")
	tryRegister(&_CGPointEqualToPoint, lib, "CGPointEqualToPoint")
	tryRegister(&_CGPointMakeWithDictionaryRepresentation, lib, "CGPointMakeWithDictionaryRepresentation")
	tryRegister(&_CGPostKeyboardEvent, lib, "CGPostKeyboardEvent")
	tryRegister(&_CGPostMouseEvent, lib, "CGPostMouseEvent")
	tryRegister(&_CGPostScrollWheelEvent, lib, "CGPostScrollWheelEvent")
	tryRegister(&_CGPreflightListenEventAccess, lib, "CGPreflightListenEventAccess")
	tryRegister(&_CGPreflightPostEventAccess, lib, "CGPreflightPostEventAccess")
	tryRegister(&_CGPreflightScreenCaptureAccess, lib, "CGPreflightScreenCaptureAccess")
	tryRegister(&_CGRectApplyAffineTransform, lib, "CGRectApplyAffineTransform")
	tryRegister(&_CGRectContainsPoint, lib, "CGRectContainsPoint")
	tryRegister(&_CGRectContainsRect, lib, "CGRectContainsRect")
	tryRegister(&_CGRectCreateDictionaryRepresentation, lib, "CGRectCreateDictionaryRepresentation")
	tryRegister(&_CGRectDivide, lib, "CGRectDivide")
	tryRegister(&_CGRectEqualToRect, lib, "CGRectEqualToRect")
	tryRegister(&_CGRectGetHeight, lib, "CGRectGetHeight")
	tryRegister(&_CGRectGetMaxX, lib, "CGRectGetMaxX")
	tryRegister(&_CGRectGetMaxY, lib, "CGRectGetMaxY")
	tryRegister(&_CGRectGetMidX, lib, "CGRectGetMidX")
	tryRegister(&_CGRectGetMidY, lib, "CGRectGetMidY")
	tryRegister(&_CGRectGetMinX, lib, "CGRectGetMinX")
	tryRegister(&_CGRectGetMinY, lib, "CGRectGetMinY")
	tryRegister(&_CGRectGetWidth, lib, "CGRectGetWidth")
	tryRegister(&_CGRectInset, lib, "CGRectInset")
	tryRegister(&_CGRectIntegral, lib, "CGRectIntegral")
	tryRegister(&_CGRectIntersection, lib, "CGRectIntersection")
	tryRegister(&_CGRectIntersectsRect, lib, "CGRectIntersectsRect")
	tryRegister(&_CGRectIsEmpty, lib, "CGRectIsEmpty")
	tryRegister(&_CGRectIsInfinite, lib, "CGRectIsInfinite")
	tryRegister(&_CGRectIsNull, lib, "CGRectIsNull")
	tryRegister(&_CGRectMakeWithDictionaryRepresentation, lib, "CGRectMakeWithDictionaryRepresentation")
	tryRegister(&_CGRectOffset, lib, "CGRectOffset")
	tryRegister(&_CGRectStandardize, lib, "CGRectStandardize")
	tryRegister(&_CGRectUnion, lib, "CGRectUnion")
	tryRegister(&_CGRegisterScreenRefreshCallback, lib, "CGRegisterScreenRefreshCallback")
	tryRegister(&_CGReleaseAllDisplays, lib, "CGReleaseAllDisplays")
	tryRegister(&_CGReleaseDisplayFadeReservation, lib, "CGReleaseDisplayFadeReservation")
	tryRegister(&_CGReleaseScreenRefreshRects, lib, "CGReleaseScreenRefreshRects")
	tryRegister(&_CGRenderingBufferLockBytePtr, lib, "CGRenderingBufferLockBytePtr")
	tryRegister(&_CGRenderingBufferProviderCreate, lib, "CGRenderingBufferProviderCreate")
	tryRegister(&_CGRenderingBufferProviderCreateWithCFData, lib, "CGRenderingBufferProviderCreateWithCFData")
	tryRegister(&_CGRenderingBufferProviderGetSize, lib, "CGRenderingBufferProviderGetSize")
	tryRegister(&_CGRenderingBufferProviderGetTypeID, lib, "CGRenderingBufferProviderGetTypeID")
	tryRegister(&_CGRenderingBufferUnlockBytePtr, lib, "CGRenderingBufferUnlockBytePtr")
	tryRegister(&_CGRequestListenEventAccess, lib, "CGRequestListenEventAccess")
	tryRegister(&_CGRequestPostEventAccess, lib, "CGRequestPostEventAccess")
	tryRegister(&_CGRequestScreenCaptureAccess, lib, "CGRequestScreenCaptureAccess")
	tryRegister(&_CGRestorePermanentDisplayConfiguration, lib, "CGRestorePermanentDisplayConfiguration")
	tryRegister(&_CGScreenRegisterMoveCallback, lib, "CGScreenRegisterMoveCallback")
	tryRegister(&_CGScreenUnregisterMoveCallback, lib, "CGScreenUnregisterMoveCallback")
	tryRegister(&_CGSessionCopyCurrentDictionary, lib, "CGSessionCopyCurrentDictionary")
	tryRegister(&_CGSetDisplayTransferByByteTable, lib, "CGSetDisplayTransferByByteTable")
	tryRegister(&_CGSetDisplayTransferByFormula, lib, "CGSetDisplayTransferByFormula")
	tryRegister(&_CGSetDisplayTransferByTable, lib, "CGSetDisplayTransferByTable")
	tryRegister(&_CGSetLocalEventsFilterDuringSuppressionState, lib, "CGSetLocalEventsFilterDuringSuppressionState")
	tryRegister(&_CGSetLocalEventsSuppressionInterval, lib, "CGSetLocalEventsSuppressionInterval")
	tryRegister(&_CGShadingGetContentHeadroom, lib, "CGShadingGetContentHeadroom")
	tryRegister(&_CGShadingCreateAxialWithContentHeadroom, lib, "CGShadingCreateAxialWithContentHeadroom")
	tryRegister(&_CGShadingCreateAxial, lib, "CGShadingCreateAxial")
	tryRegister(&_CGShadingCreateRadialWithContentHeadroom, lib, "CGShadingCreateRadialWithContentHeadroom")
	tryRegister(&_CGShadingCreateRadial, lib, "CGShadingCreateRadial")
	tryRegister(&_CGShadingGetTypeID, lib, "CGShadingGetTypeID")
	tryRegister(&_CGShadingRelease, lib, "CGShadingRelease")
	tryRegister(&_CGShadingRetain, lib, "CGShadingRetain")
	tryRegister(&_CGShieldingWindowID, lib, "CGShieldingWindowID")
	tryRegister(&_CGShieldingWindowLevel, lib, "CGShieldingWindowLevel")
	tryRegister(&_CGSizeApplyAffineTransform, lib, "CGSizeApplyAffineTransform")
	tryRegister(&_CGSizeCreateDictionaryRepresentation, lib, "CGSizeCreateDictionaryRepresentation")
	tryRegister(&_CGSizeEqualToSize, lib, "CGSizeEqualToSize")
	tryRegister(&_CGSizeMakeWithDictionaryRepresentation, lib, "CGSizeMakeWithDictionaryRepresentation")
	tryRegister(&_CGUnregisterScreenRefreshCallback, lib, "CGUnregisterScreenRefreshCallback")
	tryRegister(&_CGWaitForScreenRefreshRects, lib, "CGWaitForScreenRefreshRects")
	tryRegister(&_CGWaitForScreenUpdateRects, lib, "CGWaitForScreenUpdateRects")
	tryRegister(&_CGWarpMouseCursorPosition, lib, "CGWarpMouseCursorPosition")
	tryRegister(&_CGWindowLevelForKey, lib, "CGWindowLevelForKey")
	tryRegister(&_CGWindowListCopyWindowInfo, lib, "CGWindowListCopyWindowInfo")
	tryRegister(&_CGWindowListCreate, lib, "CGWindowListCreate")
	tryRegister(&_CGWindowListCreateDescriptionFromArray, lib, "CGWindowListCreateDescriptionFromArray")
	tryRegister(&_CGWindowListCreateImage, lib, "CGWindowListCreateImage")
	tryRegister(&_CGWindowServerCFMachPort, lib, "CGWindowServerCFMachPort")
	tryRegister(&_CGWindowServerCreateServerPort, lib, "CGWindowServerCreateServerPort")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}

// Decrements the retain count of a color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408855-cgcolorspacerelease
func CGColorSpaceRelease(p0 CGColorSpaceRef) {
	_CGColorSpaceRelease(p0)
}

// Sets the level of interpolation quality for a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455656-cgcontextsetinterpolationquality
func CGContextSetInterpolationQuality(p0 CGContextRef, p1 unsafe.Pointer) {
	_CGContextSetInterpolationQuality(p0, p1)
}

// Reserves the fade hardware for a specified time interval. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgacquiredisplayfadereservation(_:_:)
func CGAcquireDisplayFadeReservation(seconds unsafe.Pointer, token unsafe.Pointer) unsafe.Pointer {
	return _CGAcquireDisplayFadeReservation(seconds, token)
}

// Returns an affine transformation matrix constructed by combining two existing affine transforms. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformconcat(_:_:)
func CGAffineTransformConcat(t1 CGAffineTransform, t2 CGAffineTransform) CGAffineTransform {
	return _CGAffineTransformConcat(t1, t2)
}

// CGAffineTransformDecompose is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformdecompose
func CGAffineTransformDecompose(transform CGAffineTransform) unsafe.Pointer {
	return _CGAffineTransformDecompose(transform)
}

// Checks whether two affine transforms are equal. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformequaltotransform(_:_:)
func CGAffineTransformEqualToTransform(t1 CGAffineTransform, t2 CGAffineTransform) bool {
	return _CGAffineTransformEqualToTransform(t1, t2)
}

// Returns an affine transformation matrix constructed by inverting an existing affine transform. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransforminvert(_:)
func CGAffineTransformInvert(t CGAffineTransform) CGAffineTransform {
	return _CGAffineTransformInvert(t)
}

// Checks whether an affine transform is the identity transform. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformisidentity(_:)
func CGAffineTransformIsIdentity(t CGAffineTransform) bool {
	return _CGAffineTransformIsIdentity(t)
}

// Returns an affine transformation matrix constructed from values you provide. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformmake(_:_:_:_:_:_:)
func CGAffineTransformMake(a CGFloat, b CGFloat, c CGFloat, d CGFloat, tx CGFloat, ty CGFloat) CGAffineTransform {
	return _CGAffineTransformMake(a, b, c, d, tx, ty)
}

// Returns an affine transformation matrix constructed from a rotation value you provide. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformmakerotation(_:)
func CGAffineTransformMakeRotation(angle CGFloat) CGAffineTransform {
	return _CGAffineTransformMakeRotation(angle)
}

// Returns an affine transformation matrix constructed from scaling values you provide. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformmakescale(_:_:)
func CGAffineTransformMakeScale(sx CGFloat, sy CGFloat) CGAffineTransform {
	return _CGAffineTransformMakeScale(sx, sy)
}

// Returns an affine transformation matrix constructed from translation values you provide. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformmaketranslation(_:_:)
func CGAffineTransformMakeTranslation(tx CGFloat, ty CGFloat) CGAffineTransform {
	return _CGAffineTransformMakeTranslation(tx, ty)
}

// CGAffineTransformMakeWithComponents is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformmakewithcomponents
func CGAffineTransformMakeWithComponents(components unsafe.Pointer) CGAffineTransform {
	return _CGAffineTransformMakeWithComponents(components)
}

// Returns an affine transformation matrix constructed by rotating an existing affine transform. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformrotate(_:_:)
func CGAffineTransformRotate(t CGAffineTransform, angle CGFloat) CGAffineTransform {
	return _CGAffineTransformRotate(t, angle)
}

// Returns an affine transformation matrix constructed by scaling an existing affine transform. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformscale(_:_:_:)
func CGAffineTransformScale(t CGAffineTransform, sx CGFloat, sy CGFloat) CGAffineTransform {
	return _CGAffineTransformScale(t, sx, sy)
}

// Returns an affine transformation matrix constructed by translating an existing affine transform. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgaffinetransformtranslate(_:_:_:)
func CGAffineTransformTranslate(t CGAffineTransform, tx CGFloat, ty CGFloat) CGAffineTransform {
	return _CGAffineTransformTranslate(t, tx, ty)
}

// Connects or disconnects the mouse and cursor while an application is in the foreground. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgassociatemouseandmousecursorposition(_:)
func CGAssociateMouseAndMouseCursorPosition(connected unsafe.Pointer) unsafe.Pointer {
	return _CGAssociateMouseAndMouseCursorPosition(connected)
}

// Begins a new set of display configuration changes. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgbegindisplayconfiguration(_:)
func CGBeginDisplayConfiguration(config unsafe.Pointer) unsafe.Pointer {
	return _CGBeginDisplayConfiguration(config)
}

// CGBitmapContextCreateAdaptive is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgbitmapcontextcreateadaptive
func CGBitmapContextCreateAdaptive(width uintptr, height uintptr, auxiliaryInfo unsafe.Pointer, onResolve bool) CGContextRef {
	return _CGBitmapContextCreateAdaptive(width, height, auxiliaryInfo, onResolve)
}

// Cancels a set of display configuration changes. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcanceldisplayconfiguration(_:)
func CGCancelDisplayConfiguration(config CGDisplayConfigRef) unsafe.Pointer {
	return _CGCancelDisplayConfiguration(config)
}

// Obtains exclusive use of all active displays, preventing other applications and system services from using the display or changing its configuration. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcapturealldisplays()
func CGCaptureAllDisplays() unsafe.Pointer {
	return _CGCaptureAllDisplays()
}

// Captures all attached displays, using the specified options. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcapturealldisplayswithoptions(_:)
func CGCaptureAllDisplaysWithOptions(options unsafe.Pointer) unsafe.Pointer {
	return _CGCaptureAllDisplaysWithOptions(options)
}

// Returns the value of the alpha component associated with a color. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/alpha
func CGColorGetAlpha(color CGColorRef) CGFloat {
	return _CGColorGetAlpha(color)
}

// Returns the color space associated with a color. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/colorspace
func CGColorGetColorSpace(color CGColorRef) CGColorSpaceRef {
	return _CGColorGetColorSpace(color)
}

// CGColorGetContentHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/contentheadroom
func CGColorGetContentHeadroom(color CGColorRef) float32 {
	return _CGColorGetContentHeadroom(color)
}

// Creates a new color in a different color space that matches the provided color. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/converted(to:intent:options:)
func CGColorCreateCopyByMatchingToColorSpace(p0 CGColorSpaceRef, intent unsafe.Pointer, color CGColorRef, options unsafe.Pointer) CGColorRef {
	return _CGColorCreateCopyByMatchingToColorSpace(p0, intent, color, options)
}

// Creates a copy of an existing color. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/copy()
func CGColorCreateCopy(color CGColorRef) CGColorRef {
	return _CGColorCreateCopy(color)
}

// Creates a copy of an existing color, substituting a new alpha value. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/copy(alpha:)
func CGColorCreateCopyWithAlpha(color CGColorRef, alpha CGFloat) CGColorRef {
	return _CGColorCreateCopyWithAlpha(color, alpha)
}

// Creates a color using a list of intensity values (including alpha) and an associated color space. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/init(colorspace:components:)
func CGColorCreate(space CGColorSpaceRef, components unsafe.Pointer) CGColorRef {
	return _CGColorCreate(space, components)
}

// Creates a color in the Generic CMYK color space. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/init(genericcmykcyan:magenta:yellow:black:alpha:)
func CGColorCreateGenericCMYK(cyan CGFloat, magenta CGFloat, yellow CGFloat, black CGFloat, alpha CGFloat) CGColorRef {
	return _CGColorCreateGenericCMYK(cyan, magenta, yellow, black, alpha)
}

// Creates a color in the Generic gray color space with a gamma ramp of 2.2. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/init(genericgraygamma2_2gray:alpha:)
func CGColorCreateGenericGrayGamma2_2(gray CGFloat, alpha CGFloat) CGColorRef {
	return _CGColorCreateGenericGrayGamma2_2(gray, alpha)
}

// Creates a color in the Generic gray color space. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/init(gray:alpha:)
func CGColorCreateGenericGray(gray CGFloat, alpha CGFloat) CGColorRef {
	return _CGColorCreateGenericGray(gray, alpha)
}

// CGColorCreateWithContentHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/init(headroom:colorspace:red:green:blue:alpha:)
func CGColorCreateWithContentHeadroom(headroom float32, space CGColorSpaceRef, red CGFloat, green CGFloat, blue CGFloat, alpha CGFloat) CGColorRef {
	return _CGColorCreateWithContentHeadroom(headroom, space, red, green, blue, alpha)
}

// Creates a color using a list of intensity values (including alpha), a pattern color space, and a pattern. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/init(patternspace:pattern:components:)
func CGColorCreateWithPattern(space CGColorSpaceRef, pattern CGPatternRef, components unsafe.Pointer) CGColorRef {
	return _CGColorCreateWithPattern(space, pattern, components)
}

// Creates a color in the Generic RGB color space. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/init(red:green:blue:alpha:)
func CGColorCreateGenericRGB(red CGFloat, green CGFloat, blue CGFloat, alpha CGFloat) CGColorRef {
	return _CGColorCreateGenericRGB(red, green, blue, alpha)
}

// Creates a color in the sRGB color space. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/init(srgbred:green:blue:alpha:)
func CGColorCreateSRGB(red CGFloat, green CGFloat, blue CGFloat, alpha CGFloat) CGColorRef {
	return _CGColorCreateSRGB(red, green, blue, alpha)
}

// Returns the number of color components (including alpha) associated with a color. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/numberofcomponents
func CGColorGetNumberOfComponents(color CGColorRef) uintptr {
	return _CGColorGetNumberOfComponents(color)
}

// Returns the pattern associated with a color in a pattern color space. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/pattern
func CGColorGetPattern(color CGColorRef) CGPatternRef {
	return _CGColorGetPattern(color)
}

// Returns the Core Foundation type identifier for a color data type. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolor/typeid
func CGColorGetTypeID() unsafe.Pointer {
	return _CGColorGetTypeID()
}

// CGColorConversionInfoConvertData is a CoreGraphics function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorconversioninfo/convert(width:height:to:format:from:format:options:)
func CGColorConversionInfoConvertData(info CGColorConversionInfoRef, width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format unsafe.Pointer, src_data unsafe.Pointer, src_format unsafe.Pointer, options unsafe.Pointer) bool {
	return _CGColorConversionInfoConvertData(info, width, height, dst_data, dst_format, src_data, src_format, options)
}

// CGColorConversionInfoCreateWithOptions is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.14.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorconversioninfo/init(optionssrc:dst:options:)
func CGColorConversionInfoCreateWithOptions(src CGColorSpaceRef, dst CGColorSpaceRef, options unsafe.Pointer) CGColorConversionInfoRef {
	return _CGColorConversionInfoCreateWithOptions(src, dst, options)
}

// Creates a conversion between two specified color spaces. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorconversioninfo/init(src:dst:)
func CGColorConversionInfoCreate(src CGColorSpaceRef, dst CGColorSpaceRef) CGColorConversionInfoRef {
	return _CGColorConversionInfoCreate(src, dst)
}

// CGColorConversionInfoCreateForToneMapping is a CoreGraphics function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorconversioninfo/init(src:srcheadroom:dst:dstheadroom:tonemapping:options:_:)
func CGColorConversionInfoCreateForToneMapping(from CGColorSpaceRef, source_headroom float32, to CGColorSpaceRef, target_headroom float32, method unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) CGColorConversionInfoRef {
	return _CGColorConversionInfoCreateForToneMapping(from, source_headroom, to, target_headroom, method, options, error)
}

// Returns the Core Foundation type identifier for a color conversion info data type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorconversioninfo/typeid
func CGColorConversionInfoGetTypeID() unsafe.Pointer {
	return _CGColorConversionInfoGetTypeID()
}

// Creates a conversion between an arbitrary number of specified color spaces. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorconversioninfocreatefromlist
func CGColorConversionInfoCreateFromList(options unsafe.Pointer, p1 CGColorSpaceRef, p2 unsafe.Pointer, p3 unsafe.Pointer) CGColorConversionInfoRef {
	return _CGColorConversionInfoCreateFromList(options, p1, p2, p3)
}

// CGColorConversionInfoCreateFromListWithArguments is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorconversioninfocreatefromlistwitharguments
func CGColorConversionInfoCreateFromListWithArguments(options unsafe.Pointer, p1 CGColorSpaceRef, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer) CGColorConversionInfoRef {
	return _CGColorConversionInfoCreateFromListWithArguments(options, p1, p2, p3, p4)
}

// Indicates whether two colors are equal. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorequaltocolor
func CGColorEqualToColor(color1 CGColorRef, color2 CGColorRef) bool {
	return _CGColorEqualToColor(color1, color2)
}

// Returns the values of the color components (including alpha) associated with a color. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorgetcomponents
func CGColorGetComponents(color CGColorRef) unsafe.Pointer {
	return _CGColorGetComponents(color)
}

// Returns a color object that represents a constant color. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorgetconstantcolor
func CGColorGetConstantColor(colorName unsafe.Pointer) CGColorRef {
	return _CGColorGetConstantColor(colorName)
}

// Decrements the retain count of a color. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorrelease
func CGColorRelease(color CGColorRef) {
	_CGColorRelease(color)
}

// Increments the retain count of a color. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorretain
func CGColorRetain(color CGColorRef) CGColorRef {
	return _CGColorRetain(color)
}

// Returns the base color space of a pattern or indexed color space. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/basecolorspace
func CGColorSpaceGetBaseColorSpace(space CGColorSpaceRef) CGColorSpaceRef {
	return _CGColorSpaceGetBaseColorSpace(space)
}

// Returns a copy of the ICC profile data of the provided color space. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/copyiccdata()
func CGColorSpaceCopyICCData(space CGColorSpaceRef) unsafe.Pointer {
	return _CGColorSpaceCopyICCData(space)
}

// Returns a copy of the color space’s properties. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/copypropertylist()
func CGColorSpaceCopyPropertyList(space CGColorSpaceRef) unsafe.Pointer {
	return _CGColorSpaceCopyPropertyList(space)
}

// Returns a copy of the ICC profile of the provided color space. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/iccdata
func CGColorSpaceCopyICCProfile(space CGColorSpaceRef) unsafe.Pointer {
	return _CGColorSpaceCopyICCProfile(space)
}

// Creates a calibrated grayscale color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/init(calibratedgraywhitepoint:blackpoint:gamma:)
func CGColorSpaceCreateCalibratedGray(whitePoint unsafe.Pointer, p1 unsafe.Pointer, blackPoint unsafe.Pointer, p3 unsafe.Pointer, gamma CGFloat) CGColorSpaceRef {
	return _CGColorSpaceCreateCalibratedGray(whitePoint, p1, blackPoint, p3, gamma)
}

// Creates a calibrated RGB color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/init(calibratedrgbwhitepoint:blackpoint:gamma:matrix:)
func CGColorSpaceCreateCalibratedRGB(whitePoint unsafe.Pointer, p1 unsafe.Pointer, blackPoint unsafe.Pointer, p3 unsafe.Pointer, gamma unsafe.Pointer, p5 unsafe.Pointer, matrix unsafe.Pointer, p7 unsafe.Pointer) CGColorSpaceRef {
	return _CGColorSpaceCreateCalibratedRGB(whitePoint, p1, blackPoint, p3, gamma, p5, matrix, p7)
}

// Creates a device-independent color space that is defined according to the ICC color profile specification. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/init(iccbasedncomponents:range:profile:alternate:)
func CGColorSpaceCreateICCBased(nComponents uintptr, range_ unsafe.Pointer, profile CGDataProviderRef, alternate CGColorSpaceRef) CGColorSpaceRef {
	return _CGColorSpaceCreateICCBased(nComponents, range_, profile, alternate)
}

// Creates an ICC-based color space using the ICC profile contained in the specified data. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/init(iccdata:)
func CGColorSpaceCreateWithICCData(data unsafe.Pointer) CGColorSpaceRef {
	return _CGColorSpaceCreateWithICCData(data)
}

// Creates an ICC-based color space using the ICC profile contained in the specified data. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/init(iccprofiledata:)
func CGColorSpaceCreateWithICCProfile(data unsafe.Pointer) CGColorSpaceRef {
	return _CGColorSpaceCreateWithICCProfile(data)
}

// Creates an indexed color space, consisting of colors specified by a color lookup table. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/init(indexedbasespace:last:colortable:)
func CGColorSpaceCreateIndexed(baseSpace CGColorSpaceRef, lastIndex uintptr, colorTable unsafe.Pointer) CGColorSpaceRef {
	return _CGColorSpaceCreateIndexed(baseSpace, lastIndex, colorTable)
}

// Creates a device-independent color space that is relative to human color perception, according to the CIE L*a*b* standard. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/init(labwhitepoint:blackpoint:range:)
func CGColorSpaceCreateLab(whitePoint unsafe.Pointer, p1 unsafe.Pointer, blackPoint unsafe.Pointer, p3 unsafe.Pointer, range_ unsafe.Pointer, p5 unsafe.Pointer) CGColorSpaceRef {
	return _CGColorSpaceCreateLab(whitePoint, p1, blackPoint, p3, range_, p5)
}

// Creates a specified type of Quartz color space. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/init(name:)
func CGColorSpaceCreateWithName(name unsafe.Pointer) CGColorSpaceRef {
	return _CGColorSpaceCreateWithName(name)
}

// Creates a pattern color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/init(patternbasespace:)
func CGColorSpaceCreatePattern(baseSpace CGColorSpaceRef) CGColorSpaceRef {
	return _CGColorSpaceCreatePattern(baseSpace)
}

// Creates a platform-specific color space. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/init(platformcolorspaceref:)
func CGColorSpaceCreateWithPlatformColorSpace(ref unsafe.Pointer) CGColorSpaceRef {
	return _CGColorSpaceCreateWithPlatformColorSpace(ref)
}

// Creates a color space from a property list. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/init(propertylistplist:)
func CGColorSpaceCreateWithPropertyList(plist unsafe.Pointer) CGColorSpaceRef {
	return _CGColorSpaceCreateWithPropertyList(plist)
}

// CGColorSpaceIsHDR is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/ishdr()
func CGColorSpaceIsHDR(p0 CGColorSpaceRef) bool {
	return _CGColorSpaceIsHDR(p0)
}

// Returns whether the RGB color space covers a significant portion of the NTSC color gamut. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/iswidegamutrgb
func CGColorSpaceIsWideGamutRGB(p0 CGColorSpaceRef) bool {
	return _CGColorSpaceIsWideGamutRGB(p0)
}

// Returns the color space model of the provided color space. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/model
func CGColorSpaceGetModel(space CGColorSpaceRef) unsafe.Pointer {
	return _CGColorSpaceGetModel(space)
}

// Returns the name used to create the specified color space. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/name
func CGColorSpaceCopyName(space CGColorSpaceRef) unsafe.Pointer {
	return _CGColorSpaceCopyName(space)
}

// Returns the number of color components in a color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/numberofcomponents
func CGColorSpaceGetNumberOfComponents(space CGColorSpaceRef) uintptr {
	return _CGColorSpaceGetNumberOfComponents(space)
}

// Returns a Boolean indicating whether the color space can be used as a destination color space. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/supportsoutput
func CGColorSpaceSupportsOutput(space CGColorSpaceRef) bool {
	return _CGColorSpaceSupportsOutput(space)
}

// Returns the Core Foundation type identifier for Quartz color spaces. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspace/typeid
func CGColorSpaceGetTypeID() unsafe.Pointer {
	return _CGColorSpaceGetTypeID()
}

// CGColorSpaceCopyBaseColorSpace is a CoreGraphics function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacecopybasecolorspace(_:)
func CGColorSpaceCopyBaseColorSpace(space CGColorSpaceRef) CGColorSpaceRef {
	return _CGColorSpaceCopyBaseColorSpace(space)
}

// CGColorSpaceCreateCopyWithStandardRange is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacecreatecopywithstandardrange(_:)
func CGColorSpaceCreateCopyWithStandardRange(space CGColorSpaceRef) CGColorSpaceRef {
	return _CGColorSpaceCreateCopyWithStandardRange(space)
}

// Creates a device-dependent CMYK color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacecreatedevicecmyk()
func CGColorSpaceCreateDeviceCMYK() CGColorSpaceRef {
	return _CGColorSpaceCreateDeviceCMYK()
}

// Creates a device-dependent grayscale color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacecreatedevicegray()
func CGColorSpaceCreateDeviceGray() CGColorSpaceRef {
	return _CGColorSpaceCreateDeviceGray()
}

// Creates a device-dependent RGB color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacecreatedevicergb()
func CGColorSpaceCreateDeviceRGB() CGColorSpaceRef {
	return _CGColorSpaceCreateDeviceRGB()
}

// CGColorSpaceCreateExtended is a CoreGraphics function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacecreateextended(_:)
func CGColorSpaceCreateExtended(space CGColorSpaceRef) CGColorSpaceRef {
	return _CGColorSpaceCreateExtended(space)
}

// CGColorSpaceCreateExtendedLinearized is a CoreGraphics function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacecreateextendedlinearized(_:)
func CGColorSpaceCreateExtendedLinearized(space CGColorSpaceRef) CGColorSpaceRef {
	return _CGColorSpaceCreateExtendedLinearized(space)
}

// CGColorSpaceCreateLinearized is a CoreGraphics function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacecreatelinearized(_:)
func CGColorSpaceCreateLinearized(space CGColorSpaceRef) CGColorSpaceRef {
	return _CGColorSpaceCreateLinearized(space)
}

// CGColorSpaceCreateWithColorSyncProfile is a CoreGraphics function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacecreatewithcolorsyncprofile(_:_:)
func CGColorSpaceCreateWithColorSyncProfile(p0 unsafe.Pointer, options unsafe.Pointer) CGColorSpaceRef {
	return _CGColorSpaceCreateWithColorSyncProfile(p0, options)
}

// Copies the entries in the color table of an indexed color space. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacegetcolortable
func CGColorSpaceGetColorTable(space CGColorSpaceRef, table unsafe.Pointer) {
	_CGColorSpaceGetColorTable(space, table)
}

// Returns the number of entries in the color table of an indexed color space. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacegetcolortablecount
func CGColorSpaceGetColorTableCount(space CGColorSpaceRef) uintptr {
	return _CGColorSpaceGetColorTableCount(space)
}

// CGColorSpaceGetName is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacegetname
func CGColorSpaceGetName(space CGColorSpaceRef) unsafe.Pointer {
	return _CGColorSpaceGetName(space)
}

// CGColorSpaceIsHLGBased is a CoreGraphics function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspaceishlgbased(_:)
func CGColorSpaceIsHLGBased(s CGColorSpaceRef) bool {
	return _CGColorSpaceIsHLGBased(s)
}

// CGColorSpaceIsPQBased is a CoreGraphics function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspaceispqbased(_:)
func CGColorSpaceIsPQBased(s CGColorSpaceRef) bool {
	return _CGColorSpaceIsPQBased(s)
}

// Increments the retain count of a color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspaceretain
func CGColorSpaceRetain(space CGColorSpaceRef) CGColorSpaceRef {
	return _CGColorSpaceRetain(space)
}

// CGColorSpaceUsesExtendedRange is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspaceusesextendedrange(_:)
func CGColorSpaceUsesExtendedRange(space CGColorSpaceRef) bool {
	return _CGColorSpaceUsesExtendedRange(space)
}

// CGColorSpaceUsesITUR_2100TF is a CoreGraphics function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspaceusesitur_2100tf(_:)
func CGColorSpaceUsesITUR_2100TF(p0 CGColorSpaceRef) bool {
	return _CGColorSpaceUsesITUR_2100TF(p0)
}

// Completes a set of display configuration changes. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcompletedisplayconfiguration(_:_:)
func CGCompleteDisplayConfiguration(config CGDisplayConfigRef, option unsafe.Pointer) unsafe.Pointer {
	return _CGCompleteDisplayConfiguration(config, option)
}

// Modifies the settings of the built-in fade effect that occurs during a display configuration. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgconfiguredisplayfadeeffect(_:_:_:_:_:_:)
func CGConfigureDisplayFadeEffect(config CGDisplayConfigRef, fadeOutSeconds unsafe.Pointer, fadeInSeconds unsafe.Pointer, fadeRed float32, fadeGreen float32, fadeBlue float32) unsafe.Pointer {
	return _CGConfigureDisplayFadeEffect(config, fadeOutSeconds, fadeInSeconds, fadeRed, fadeGreen, fadeBlue)
}

// Changes the configuration of a mirroring set. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgconfiguredisplaymirrorofdisplay(_:_:_:)
func CGConfigureDisplayMirrorOfDisplay(config CGDisplayConfigRef, display unsafe.Pointer, master unsafe.Pointer) unsafe.Pointer {
	return _CGConfigureDisplayMirrorOfDisplay(config, display, master)
}

// Configures the display mode of a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgconfiguredisplaymode(_:_:_:)
func CGConfigureDisplayMode(config CGDisplayConfigRef, display unsafe.Pointer, mode unsafe.Pointer) unsafe.Pointer {
	return _CGConfigureDisplayMode(config, display, mode)
}

// Configures the origin of a display relative to the global display coordinate space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgconfiguredisplayorigin(_:_:_:_:)
func CGConfigureDisplayOrigin(config CGDisplayConfigRef, display unsafe.Pointer, x unsafe.Pointer, y unsafe.Pointer) unsafe.Pointer {
	return _CGConfigureDisplayOrigin(config, display, x, y)
}

// Enables or disables stereo operation for a display, as part of a display configuration. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgconfiguredisplaystereooperation(_:_:_:_:)
func CGConfigureDisplayStereoOperation(config CGDisplayConfigRef, display unsafe.Pointer, stereo unsafe.Pointer, forceBlueLine unsafe.Pointer) unsafe.Pointer {
	return _CGConfigureDisplayStereoOperation(config, display, stereo, forceBlueLine)
}

// Configures the display mode of a display. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgconfiguredisplaywithdisplaymode(_:_:_:_:)
func CGConfigureDisplayWithDisplayMode(config CGDisplayConfigRef, display unsafe.Pointer, mode CGDisplayModeRef, options unsafe.Pointer) unsafe.Pointer {
	return _CGConfigureDisplayWithDisplayMode(config, display, mode, options)
}

// Sets a destination to jump to when a point in the current page of a PDF graphics context is clicked. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/adddestination(_:at:)
func CGPDFContextAddDestinationAtPoint(context CGContextRef, name unsafe.Pointer, point CGPoint) {
	_CGPDFContextAddDestinationAtPoint(context, name, point)
}

// Associates custom metadata with the PDF document. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/adddocumentmetadata(_:)
func CGPDFContextAddDocumentMetadata(context CGContextRef, metadata unsafe.Pointer) {
	_CGPDFContextAddDocumentMetadata(context, metadata)
}

// Adds an ellipse that fits inside the specified rectangle. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/addellipse(in:)
func CGContextAddEllipseInRect(c CGContextRef, rect CGRect) {
	_CGContextAddEllipseInRect(c, rect)
}

// Adds a previously created path object to the current path in a graphics context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/addpath(_:)
func CGContextAddPath(c CGContextRef, path CGPathRef) {
	_CGContextAddPath(c, path)
}

// Adds a rectangular path to the current path. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/addrect(_:)
func CGContextAddRect(c CGContextRef, rect CGRect) {
	_CGContextAddRect(c, rect)
}

// Returns the alpha information associated with the context, which indicates how a bitmap context handles the alpha component. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/alphainfo
func CGBitmapContextGetAlphaInfo(context CGContextRef) unsafe.Pointer {
	return _CGBitmapContextGetAlphaInfo(context)
}

// Begins a new page in a PDF graphics context. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/beginpdfpage(_:)
func CGPDFContextBeginPage(context CGContextRef, pageInfo unsafe.Pointer) {
	_CGPDFContextBeginPage(context, pageInfo)
}

// Starts a new page in a page-based graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/beginpage(mediabox:)
func CGContextBeginPage(c CGContextRef, mediaBox unsafe.Pointer) {
	_CGContextBeginPage(c, mediaBox)
}

// Creates a new empty path in a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/beginpath()
func CGContextBeginPath(c CGContextRef) {
	_CGContextBeginPath(c)
}

// Begins a transparency layer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/begintransparencylayer(auxiliaryinfo:)
func CGContextBeginTransparencyLayer(c CGContextRef, auxiliaryInfo unsafe.Pointer) {
	_CGContextBeginTransparencyLayer(c, auxiliaryInfo)
}

// Begins a transparency layer whose contents are bounded by the specified rectangle. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/begintransparencylayer(in:auxiliaryinfo:)
func CGContextBeginTransparencyLayerWithRect(c CGContextRef, rect CGRect, auxInfo unsafe.Pointer) {
	_CGContextBeginTransparencyLayerWithRect(c, rect, auxInfo)
}

// Obtains the bitmap information associated with a bitmap graphics context. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/bitmapinfo
func CGBitmapContextGetBitmapInfo(context CGContextRef) unsafe.Pointer {
	return _CGBitmapContextGetBitmapInfo(context)
}

// Returns the bits per component of a bitmap context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/bitspercomponent
func CGBitmapContextGetBitsPerComponent(context CGContextRef) uintptr {
	return _CGBitmapContextGetBitsPerComponent(context)
}

// Returns the bits per pixel of a bitmap context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/bitsperpixel
func CGBitmapContextGetBitsPerPixel(context CGContextRef) uintptr {
	return _CGBitmapContextGetBitsPerPixel(context)
}

// Returns the bounding box of a clipping path. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/boundingboxofclippath
func CGContextGetClipBoundingBox(c CGContextRef) CGRect {
	return _CGContextGetClipBoundingBox(c)
}

// Returns the smallest rectangle that contains the current path. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/boundingboxofpath
func CGContextGetPathBoundingBox(c CGContextRef) CGRect {
	return _CGContextGetPathBoundingBox(c)
}

// Returns the bytes per row of a bitmap context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/bytesperrow
func CGBitmapContextGetBytesPerRow(context CGContextRef) uintptr {
	return _CGBitmapContextGetBytesPerRow(context)
}

// Paints a transparent rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/clear(_:)
func CGContextClearRect(c CGContextRef, rect CGRect) {
	_CGContextClearRect(c, rect)
}

// Sets the clipping path to the intersection of the current clipping path with the area defined by the specified rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/clip(to:)-7cbwq
func CGContextClipToRect(c CGContextRef, rect CGRect) {
	_CGContextClipToRect(c, rect)
}

// Maps a mask into the specified rectangle and intersects it with the current clipping area of the graphics context. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/clip(to:mask:)
func CGContextClipToMask(c CGContextRef, rect CGRect, mask CGImageRef) {
	_CGContextClipToMask(c, rect, mask)
}

// Closes a PDF document. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/closepdf()
func CGPDFContextClose(context CGContextRef) {
	_CGPDFContextClose(context)
}

// Closes and terminates the current path’s subpath. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/closepath()
func CGContextClosePath(c CGContextRef) {
	_CGContextClosePath(c)
}

// Returns the color space of a bitmap context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/colorspace
func CGBitmapContextGetColorSpace(context CGContextRef) CGColorSpaceRef {
	return _CGBitmapContextGetColorSpace(context)
}

// Transforms the user coordinate system in a context using a specified matrix. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/concatenate(_:)
func CGContextConcatCTM(c CGContextRef, transform CGAffineTransform) {
	_CGContextConcatCTM(c, transform)
}

// Returns a size that is transformed from user space coordinates to device space coordinates. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/converttodevicespace(_:)-224h2
func CGContextConvertSizeToDeviceSpace(c CGContextRef, size CGSize) CGSize {
	return _CGContextConvertSizeToDeviceSpace(c, size)
}

// Returns a point that is transformed from user space coordinates to device space coordinates. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/converttodevicespace(_:)-53m7u
func CGContextConvertPointToDeviceSpace(c CGContextRef, point CGPoint) CGPoint {
	return _CGContextConvertPointToDeviceSpace(c, point)
}

// Returns a rectangle that is transformed from user space coordinate to device space coordinates. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/converttodevicespace(_:)-91x5g
func CGContextConvertRectToDeviceSpace(c CGContextRef, rect CGRect) CGRect {
	return _CGContextConvertRectToDeviceSpace(c, rect)
}

// Returns a rectangle that is transformed from device space coordinate to user space coordinates. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/converttouserspace(_:)-1hk5r
func CGContextConvertRectToUserSpace(c CGContextRef, rect CGRect) CGRect {
	return _CGContextConvertRectToUserSpace(c, rect)
}

// Returns a point that is transformed from device space coordinates to user space coordinates. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/converttouserspace(_:)-3mtg3
func CGContextConvertPointToUserSpace(c CGContextRef, point CGPoint) CGPoint {
	return _CGContextConvertPointToUserSpace(c, point)
}

// Returns a size that is transformed from device space coordinates to user space coordinates. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/converttouserspace(_:)-693ur
func CGContextConvertSizeToUserSpace(c CGContextRef, size CGSize) CGSize {
	return _CGContextConvertSizeToUserSpace(c, size)
}

// Returns the current transformation matrix. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/ctm
func CGContextGetCTM(c CGContextRef) CGAffineTransform {
	return _CGContextGetCTM(c)
}

// Returns the current point in a non-empty path. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/currentpointofpath
func CGContextGetPathCurrentPoint(c CGContextRef) CGPoint {
	return _CGContextGetPathCurrentPoint(c)
}

// Returns a pointer to the image data associated with a bitmap context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/data
func CGBitmapContextGetData(context CGContextRef) unsafe.Pointer {
	return _CGBitmapContextGetData(context)
}

// Paints a gradient fill that varies along the line defined by the provided starting and ending points. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/drawlineargradient(_:start:end:options:)
func CGContextDrawLinearGradient(c CGContextRef, gradient CGGradientRef, startPoint CGPoint, endPoint CGPoint, options unsafe.Pointer) {
	_CGContextDrawLinearGradient(c, gradient, startPoint, endPoint, options)
}

// Draws the content of a PDF page into the current graphics context. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/drawpdfpage(_:)
func CGContextDrawPDFPage(c CGContextRef, page CGPDFPageRef) {
	_CGContextDrawPDFPage(c, page)
}

// Draws the current path using the provided drawing mode. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/drawpath(using:)
func CGContextDrawPath(c CGContextRef, mode unsafe.Pointer) {
	_CGContextDrawPath(c, mode)
}

// Paints a gradient fill that varies along the area defined by the provided starting and ending circles. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/drawradialgradient(_:startcenter:startradius:endcenter:endradius:options:)
func CGContextDrawRadialGradient(c CGContextRef, gradient CGGradientRef, startCenter CGPoint, startRadius CGFloat, endCenter CGPoint, endRadius CGFloat, options unsafe.Pointer) {
	_CGContextDrawRadialGradient(c, gradient, startCenter, startRadius, endCenter, endRadius, options)
}

// Fills the clipping path of a context with the specified shading. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/drawshading(_:)
func CGContextDrawShading(c CGContextRef, shading CGShadingRef) {
	_CGContextDrawShading(c, shading)
}

// Ends the current page in the PDF graphics context. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/endpdfpage()
func CGPDFContextEndPage(context CGContextRef) {
	_CGPDFContextEndPage(context)
}

// Ends the current page in a page-based graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/endpage()
func CGContextEndPage(c CGContextRef) {
	_CGContextEndPage(c)
}

// Ends a transparency layer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/endtransparencylayer()
func CGContextEndTransparencyLayer(c CGContextRef) {
	_CGContextEndTransparencyLayer(c)
}

// Paints the area contained within the provided rectangle, using the fill color in the current graphics state. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/fill(_:)-7a0rk
func CGContextFillRect(c CGContextRef, rect CGRect) {
	_CGContextFillRect(c, rect)
}

// Paints the area of the ellipse that fits inside the provided rectangle, using the fill color in the current graphics state. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/fillellipse(in:)
func CGContextFillEllipseInRect(c CGContextRef, rect CGRect) {
	_CGContextFillEllipseInRect(c, rect)
}

// Forces all pending drawing operations in a window context to be rendered immediately to the destination device. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/flush()
func CGContextFlush(c CGContextRef) {
	_CGContextFlush(c)
}

// Returns the height in pixels of a bitmap context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/height
func CGBitmapContextGetHeight(context CGContextRef) uintptr {
	return _CGBitmapContextGetHeight(context)
}

// Creates a URL-based PDF graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/init(_:mediabox:_:)
func CGPDFContextCreateWithURL(url unsafe.Pointer, mediaBox unsafe.Pointer, auxiliaryInfo unsafe.Pointer) CGContextRef {
	return _CGPDFContextCreateWithURL(url, mediaBox, auxiliaryInfo)
}

// Creates a PDF graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/init(consumer:mediabox:_:)
func CGPDFContextCreate(consumer CGDataConsumerRef, mediaBox unsafe.Pointer, auxiliaryInfo unsafe.Pointer) CGContextRef {
	return _CGPDFContextCreate(consumer, mediaBox, auxiliaryInfo)
}

// CGBitmapContextCreate is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/init(data:width:height:bitspercomponent:bytesperrow:space:bitmapinfo:)-10b3i
func CGBitmapContextCreate(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo unsafe.Pointer) CGContextRef {
	return _CGBitmapContextCreate(data, width, height, bitsPerComponent, bytesPerRow, space, bitmapInfo)
}

// CGBitmapContextCreateWithData is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/init(data:width:height:bitspercomponent:bytesperrow:space:bitmapinfo:releasecallback:releaseinfo:)-4yzt5
func CGBitmapContextCreateWithData(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo unsafe.Pointer, releaseCallback unsafe.Pointer, releaseInfo unsafe.Pointer) CGContextRef {
	return _CGBitmapContextCreateWithData(data, width, height, bitsPerComponent, bytesPerRow, space, bitmapInfo, releaseCallback, releaseInfo)
}

// Returns the current level of interpolation quality for a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/interpolationquality
func CGContextGetInterpolationQuality(c CGContextRef) unsafe.Pointer {
	return _CGContextGetInterpolationQuality(c)
}

// Indicates whether the current path contains any subpaths. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/ispathempty
func CGContextIsPathEmpty(c CGContextRef) bool {
	return _CGContextIsPathEmpty(c)
}

// Creates and returns a CGImage from the pixel data in a bitmap graphics context. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/makeimage()
func CGBitmapContextCreateImage(context CGContextRef) CGImageRef {
	return _CGBitmapContextCreateImage(context)
}

// Returns a path object built from the current path information in a graphics context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/path
func CGContextCopyPath(c CGContextRef) CGPathRef {
	return _CGContextCopyPath(c)
}

// Checks to see whether the specified point is contained in the current path. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/pathcontains(_:mode:)
func CGContextPathContainsPoint(c CGContextRef, point CGPoint, mode unsafe.Pointer) bool {
	return _CGContextPathContainsPoint(c, point, mode)
}

// Replaces the path in the graphics context with the stroked version of the path. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/replacepathwithstrokedpath()
func CGContextReplacePathWithStrokedPath(c CGContextRef) {
	_CGContextReplacePathWithStrokedPath(c)
}

// CGContextResetClip is a CoreGraphics function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/resetclip()
func CGContextResetClip(c CGContextRef) {
	_CGContextResetClip(c)
}

// Sets the current graphics state to the state most recently saved. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/restoregstate()
func CGContextRestoreGState(c CGContextRef) {
	_CGContextRestoreGState(c)
}

// Rotates the user coordinate system in a context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/rotate(by:)
func CGContextRotateCTM(c CGContextRef, angle CGFloat) {
	_CGContextRotateCTM(c, angle)
}

// Pushes a copy of the current graphics state onto the graphics state stack for the context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/savegstate()
func CGContextSaveGState(c CGContextRef) {
	_CGContextSaveGState(c)
}

// Changes the scale of the user coordinate system in a context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/scaleby(x:y:)
func CGContextScaleCTM(c CGContextRef, sx CGFloat, sy CGFloat) {
	_CGContextScaleCTM(c, sx, sy)
}

// Sets the font and font size in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/selectfont(name:size:textencoding:)
func CGContextSelectFont(c CGContextRef, name unsafe.Pointer, size CGFloat, textEncoding unsafe.Pointer) {
	_CGContextSelectFont(c, name, size, textEncoding)
}

// Sets whether or not to allow antialiasing for a graphics context. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setallowsantialiasing(_:)
func CGContextSetAllowsAntialiasing(c CGContextRef, allowsAntialiasing bool) {
	_CGContextSetAllowsAntialiasing(c, allowsAntialiasing)
}

// Sets whether or not to allow font smoothing for a graphics context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setallowsfontsmoothing(_:)
func CGContextSetAllowsFontSmoothing(c CGContextRef, allowsFontSmoothing bool) {
	_CGContextSetAllowsFontSmoothing(c, allowsFontSmoothing)
}

// Sets whether or not to allow subpixel positioning for a graphics context. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setallowsfontsubpixelpositioning(_:)
func CGContextSetAllowsFontSubpixelPositioning(c CGContextRef, allowsFontSubpixelPositioning bool) {
	_CGContextSetAllowsFontSubpixelPositioning(c, allowsFontSubpixelPositioning)
}

// Sets whether or not to allow subpixel quantization for a graphics context. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setallowsfontsubpixelquantization(_:)
func CGContextSetAllowsFontSubpixelQuantization(c CGContextRef, allowsFontSubpixelQuantization bool) {
	_CGContextSetAllowsFontSubpixelQuantization(c, allowsFontSubpixelQuantization)
}

// Sets the opacity level for objects drawn in a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setalpha(_:)
func CGContextSetAlpha(c CGContextRef, alpha CGFloat) {
	_CGContextSetAlpha(c, alpha)
}

// Sets how sample values are composited by a graphics context. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setblendmode(_:)
func CGContextSetBlendMode(c CGContextRef, mode unsafe.Pointer) {
	_CGContextSetBlendMode(c, mode)
}

// Sets the current character spacing. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setcharacterspacing(_:)
func CGContextSetCharacterSpacing(c CGContextRef, spacing CGFloat) {
	_CGContextSetCharacterSpacing(c, spacing)
}

// Sets a destination to jump to when a rectangle in the current PDF page is clicked. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setdestination(_:for:)
func CGPDFContextSetDestinationForRect(context CGContextRef, name unsafe.Pointer, rect CGRect) {
	_CGPDFContextSetDestinationForRect(context, name, rect)
}

// CGContextSetEDRTargetHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setedrtargetheadroom(_:)
func CGContextSetEDRTargetHeadroom(c CGContextRef, headroom float32) bool {
	return _CGContextSetEDRTargetHeadroom(c, headroom)
}

// Sets the current fill color. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setfillcolor(_:)-756dy
func CGContextSetFillColor(c CGContextRef, components unsafe.Pointer) {
	_CGContextSetFillColor(c, components)
}

// Sets the current fill color in a graphics context, using a CGColor. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setfillcolor(_:)-8lhn8
func CGContextSetFillColorWithColor(c CGContextRef, color CGColorRef) {
	_CGContextSetFillColorWithColor(c, color)
}

// Sets the current fill color to a value in the DeviceCMYK color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setfillcolor(cyan:magenta:yellow:black:alpha:)
func CGContextSetCMYKFillColor(c CGContextRef, cyan CGFloat, magenta CGFloat, yellow CGFloat, black CGFloat, alpha CGFloat) {
	_CGContextSetCMYKFillColor(c, cyan, magenta, yellow, black, alpha)
}

// Sets the current fill color to a value in the DeviceGray color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setfillcolor(gray:alpha:)
func CGContextSetGrayFillColor(c CGContextRef, gray CGFloat, alpha CGFloat) {
	_CGContextSetGrayFillColor(c, gray, alpha)
}

// Sets the current fill color to a value in the DeviceRGB color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setfillcolor(red:green:blue:alpha:)
func CGContextSetRGBFillColor(c CGContextRef, red CGFloat, green CGFloat, blue CGFloat, alpha CGFloat) {
	_CGContextSetRGBFillColor(c, red, green, blue, alpha)
}

// Sets the fill color space in a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setfillcolorspace(_:)
func CGContextSetFillColorSpace(c CGContextRef, space CGColorSpaceRef) {
	_CGContextSetFillColorSpace(c, space)
}

// Sets the fill pattern in the specified graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setfillpattern(_:colorcomponents:)
func CGContextSetFillPattern(c CGContextRef, pattern CGPatternRef, components unsafe.Pointer) {
	_CGContextSetFillPattern(c, pattern, components)
}

// Sets the accuracy of curved paths in a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setflatness(_:)
func CGContextSetFlatness(c CGContextRef, flatness CGFloat) {
	_CGContextSetFlatness(c, flatness)
}

// Sets the platform font in a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setfont(_:)
func CGContextSetFont(c CGContextRef, font CGFontRef) {
	_CGContextSetFont(c, font)
}

// Sets the current font size. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setfontsize(_:)
func CGContextSetFontSize(c CGContextRef, size CGFloat) {
	_CGContextSetFontSize(c, size)
}

// Sets the style for the endpoints of lines drawn in a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setlinecap(_:)
func CGContextSetLineCap(c CGContextRef, cap unsafe.Pointer) {
	_CGContextSetLineCap(c, cap)
}

// Sets the style for the joins of connected lines in a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setlinejoin(_:)
func CGContextSetLineJoin(c CGContextRef, join unsafe.Pointer) {
	_CGContextSetLineJoin(c, join)
}

// Sets the line width for a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setlinewidth(_:)
func CGContextSetLineWidth(c CGContextRef, width CGFloat) {
	_CGContextSetLineWidth(c, width)
}

// Sets the miter limit for the joins of connected lines in a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setmiterlimit(_:)
func CGContextSetMiterLimit(c CGContextRef, limit CGFloat) {
	_CGContextSetMiterLimit(c, limit)
}

// Sets the pattern phase of a context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setpatternphase(_:)
func CGContextSetPatternPhase(c CGContextRef, phase CGSize) {
	_CGContextSetPatternPhase(c, phase)
}

// Sets the rendering intent in the current graphics state. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setrenderingintent(_:)
func CGContextSetRenderingIntent(c CGContextRef, intent unsafe.Pointer) {
	_CGContextSetRenderingIntent(c, intent)
}

// Enables shadowing in a graphics context. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setshadow(offset:blur:)
func CGContextSetShadow(c CGContextRef, offset CGSize, blur CGFloat) {
	_CGContextSetShadow(c, offset, blur)
}

// Enables shadowing with color a graphics context. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setshadow(offset:blur:color:)
func CGContextSetShadowWithColor(c CGContextRef, offset CGSize, blur CGFloat, color CGColorRef) {
	_CGContextSetShadowWithColor(c, offset, blur, color)
}

// Sets antialiasing on or off for a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setshouldantialias(_:)
func CGContextSetShouldAntialias(c CGContextRef, shouldAntialias bool) {
	_CGContextSetShouldAntialias(c, shouldAntialias)
}

// Enables or disables font smoothing in a graphics context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setshouldsmoothfonts(_:)
func CGContextSetShouldSmoothFonts(c CGContextRef, shouldSmoothFonts bool) {
	_CGContextSetShouldSmoothFonts(c, shouldSmoothFonts)
}

// Enables or disables subpixel positioning in a graphics context. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setshouldsubpixelpositionfonts(_:)
func CGContextSetShouldSubpixelPositionFonts(c CGContextRef, shouldSubpixelPositionFonts bool) {
	_CGContextSetShouldSubpixelPositionFonts(c, shouldSubpixelPositionFonts)
}

// Enables or disables subpixel quantization in a graphics context. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setshouldsubpixelquantizefonts(_:)
func CGContextSetShouldSubpixelQuantizeFonts(c CGContextRef, shouldSubpixelQuantizeFonts bool) {
	_CGContextSetShouldSubpixelQuantizeFonts(c, shouldSubpixelQuantizeFonts)
}

// Sets the current stroke color in a context, using a CGColor. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setstrokecolor(_:)-1sskg
func CGContextSetStrokeColorWithColor(c CGContextRef, color CGColorRef) {
	_CGContextSetStrokeColorWithColor(c, color)
}

// Sets the current stroke color. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setstrokecolor(_:)-4pd8p
func CGContextSetStrokeColor(c CGContextRef, components unsafe.Pointer) {
	_CGContextSetStrokeColor(c, components)
}

// Sets the current stroke color to a value in the DeviceCMYK color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setstrokecolor(cyan:magenta:yellow:black:alpha:)
func CGContextSetCMYKStrokeColor(c CGContextRef, cyan CGFloat, magenta CGFloat, yellow CGFloat, black CGFloat, alpha CGFloat) {
	_CGContextSetCMYKStrokeColor(c, cyan, magenta, yellow, black, alpha)
}

// Sets the current stroke color to a value in the DeviceGray color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setstrokecolor(gray:alpha:)
func CGContextSetGrayStrokeColor(c CGContextRef, gray CGFloat, alpha CGFloat) {
	_CGContextSetGrayStrokeColor(c, gray, alpha)
}

// Sets the current stroke color to a value in the DeviceRGB color space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setstrokecolor(red:green:blue:alpha:)
func CGContextSetRGBStrokeColor(c CGContextRef, red CGFloat, green CGFloat, blue CGFloat, alpha CGFloat) {
	_CGContextSetRGBStrokeColor(c, red, green, blue, alpha)
}

// Sets the stroke color space in a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setstrokecolorspace(_:)
func CGContextSetStrokeColorSpace(c CGContextRef, space CGColorSpaceRef) {
	_CGContextSetStrokeColorSpace(c, space)
}

// Sets the stroke pattern in the specified graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/setstrokepattern(_:colorcomponents:)
func CGContextSetStrokePattern(c CGContextRef, pattern CGPatternRef, components unsafe.Pointer) {
	_CGContextSetStrokePattern(c, pattern, components)
}

// Sets the current text drawing mode. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/settextdrawingmode(_:)
func CGContextSetTextDrawingMode(c CGContextRef, mode unsafe.Pointer) {
	_CGContextSetTextDrawingMode(c, mode)
}

// Sets the URL associated with a rectangle in a PDF graphics context. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/seturl(_:for:)
func CGPDFContextSetURLForRect(context CGContextRef, url unsafe.Pointer, rect CGRect) {
	_CGPDFContextSetURLForRect(context, url, rect)
}

// Displays an array of glyphs at the current text position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/showglyphs(g:count:)
func CGContextShowGlyphs(c CGContextRef, g unsafe.Pointer, count uintptr) {
	_CGContextShowGlyphs(c, g, count)
}

// Displays an array of glyphs at a position you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/showglyphsatpoint(x:y:glyphs:count:)
func CGContextShowGlyphsAtPoint(c CGContextRef, x CGFloat, y CGFloat, glyphs unsafe.Pointer, count uintptr) {
	_CGContextShowGlyphsAtPoint(c, x, y, glyphs, count)
}

// Draws an array of glyphs with varying offsets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/showglyphswithadvances(glyphs:advances:count:)
func CGContextShowGlyphsWithAdvances(c CGContextRef, glyphs unsafe.Pointer, advances unsafe.Pointer, count uintptr) {
	_CGContextShowGlyphsWithAdvances(c, glyphs, advances, count)
}

// Displays a character array at the current text position, a point specified by the current text matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/showtext(string:length:)
func CGContextShowText(c CGContextRef, string unsafe.Pointer, length uintptr) {
	_CGContextShowText(c, string, length)
}

// Displays a character string at a position you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/showtextatpoint(x:y:string:length:)
func CGContextShowTextAtPoint(c CGContextRef, x CGFloat, y CGFloat, string unsafe.Pointer, length uintptr) {
	_CGContextShowTextAtPoint(c, x, y, string, length)
}

// Paints a rectangular path. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/stroke(_:)
func CGContextStrokeRect(c CGContextRef, rect CGRect) {
	_CGContextStrokeRect(c, rect)
}

// Paints a rectangular path, using the specified line width. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/stroke(_:width:)
func CGContextStrokeRectWithWidth(c CGContextRef, rect CGRect, width CGFloat) {
	_CGContextStrokeRectWithWidth(c, rect, width)
}

// Strokes an ellipse that fits inside the specified rectangle. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/strokeellipse(in:)
func CGContextStrokeEllipseInRect(c CGContextRef, rect CGRect) {
	_CGContextStrokeEllipseInRect(c, rect)
}

// Paints a line along the current path. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/strokepath()
func CGContextStrokePath(c CGContextRef) {
	_CGContextStrokePath(c)
}

// Marks a window context for update. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/synchronize()
func CGContextSynchronize(c CGContextRef) {
	_CGContextSynchronize(c)
}

// CGContextSynchronizeAttributes is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/synchronizeattributes()
func CGContextSynchronizeAttributes(c CGContextRef) {
	_CGContextSynchronizeAttributes(c)
}

// Returns the current text matrix. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/textmatrix
func CGContextGetTextMatrix(c CGContextRef) CGAffineTransform {
	return _CGContextGetTextMatrix(c)
}

// Changes the origin of the user coordinate system in a context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/translateby(x:y:)
func CGContextTranslateCTM(c CGContextRef, tx CGFloat, ty CGFloat) {
	_CGContextTranslateCTM(c, tx, ty)
}

// Returns the type identifier for a graphics context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/typeid
func CGContextGetTypeID() unsafe.Pointer {
	return _CGContextGetTypeID()
}

// Returns an affine transform that maps user space coordinates to device space coordinates. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/userspacetodevicespacetransform
func CGContextGetUserSpaceToDeviceSpaceTransform(c CGContextRef) CGAffineTransform {
	return _CGContextGetUserSpaceToDeviceSpaceTransform(c)
}

// Returns the width in pixels of a bitmap context. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontext/width
func CGBitmapContextGetWidth(context CGContextRef) uintptr {
	return _CGBitmapContextGetWidth(context)
}

// Adds an arc of a circle to the current path, possibly preceded by a straight line segment [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextaddarc
func CGContextAddArc(c CGContextRef, x CGFloat, y CGFloat, radius CGFloat, startAngle CGFloat, endAngle CGFloat, clockwise int) {
	_CGContextAddArc(c, x, y, radius, startAngle, endAngle, clockwise)
}

// Adds an arc of a circle to the current path, using a radius and tangent points. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextaddarctopoint
func CGContextAddArcToPoint(c CGContextRef, x1 CGFloat, y1 CGFloat, x2 CGFloat, y2 CGFloat, radius CGFloat) {
	_CGContextAddArcToPoint(c, x1, y1, x2, y2, radius)
}

// Appends a cubic Bézier curve from the current point, using the provided control points and end point . [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextaddcurvetopoint
func CGContextAddCurveToPoint(c CGContextRef, cp1x CGFloat, cp1y CGFloat, cp2x CGFloat, cp2y CGFloat, x CGFloat, y CGFloat) {
	_CGContextAddCurveToPoint(c, cp1x, cp1y, cp2x, cp2y, x, y)
}

// Appends a straight line segment from the current point to the provided point . [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextaddlinetopoint
func CGContextAddLineToPoint(c CGContextRef, x CGFloat, y CGFloat) {
	_CGContextAddLineToPoint(c, x, y)
}

// Adds a sequence of connected straight-line segments to the current path. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextaddlines
func CGContextAddLines(c CGContextRef, points unsafe.Pointer, count uintptr) {
	_CGContextAddLines(c, points, count)
}

// Appends a quadratic Bézier curve from the current point, using a control point and an end point you specify. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextaddquadcurvetopoint
func CGContextAddQuadCurveToPoint(c CGContextRef, cpx CGFloat, cpy CGFloat, x CGFloat, y CGFloat) {
	_CGContextAddQuadCurveToPoint(c, cpx, cpy, x, y)
}

// Adds a set of rectangular paths to the current path. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextaddrects
func CGContextAddRects(c CGContextRef, rects unsafe.Pointer, count uintptr) {
	_CGContextAddRects(c, rects, count)
}

// Modifies the current clipping path, using the nonzero winding number rule. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextclip
func CGContextClip(c CGContextRef) {
	_CGContextClip(c)
}

// Sets the clipping path to the intersection of the current clipping path with the region defined by an array of rectangles. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextcliptorects
func CGContextClipToRects(c CGContextRef, rects unsafe.Pointer, count uintptr) {
	_CGContextClipToRects(c, rects, count)
}

// CGContextDrawConicGradient is a CoreGraphics function. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextdrawconicgradient(_:_:_:_:)
func CGContextDrawConicGradient(c CGContextRef, gradient CGGradientRef, center CGPoint, angle CGFloat) {
	_CGContextDrawConicGradient(c, gradient, center, angle)
}

// Draws an image into a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextdrawimage
func CGContextDrawImage(c CGContextRef, rect CGRect, image CGImageRef) {
	_CGContextDrawImage(c, rect, image)
}

// CGContextDrawImageApplyingToneMapping is a CoreGraphics function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextdrawimageapplyingtonemapping
func CGContextDrawImageApplyingToneMapping(c CGContextRef, r CGRect, image CGImageRef, method unsafe.Pointer, options unsafe.Pointer) bool {
	return _CGContextDrawImageApplyingToneMapping(c, r, image, method, options)
}

// Draws the contents of a CGLayer object at the specified point. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextdrawlayeratpoint
func CGContextDrawLayerAtPoint(context CGContextRef, point CGPoint, layer CGLayerRef) {
	_CGContextDrawLayerAtPoint(context, point, layer)
}

// Draws the contents of a layer object into the specified rectangle. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextdrawlayerinrect
func CGContextDrawLayerInRect(context CGContextRef, rect CGRect, layer CGLayerRef) {
	_CGContextDrawLayerInRect(context, rect, layer)
}

// CGContextDrawPDFDocument is a CoreGraphics function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextdrawpdfdocument
func CGContextDrawPDFDocument(c CGContextRef, rect CGRect, document CGPDFDocumentRef, page int) {
	_CGContextDrawPDFDocument(c, rect, document, page)
}

// Repeatedly draws an image, scaled to the provided rectangle, to fill the current clip region. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextdrawtiledimage
func CGContextDrawTiledImage(c CGContextRef, rect CGRect, image CGImageRef) {
	_CGContextDrawTiledImage(c, rect, image)
}

// Modifies the current clipping path, using the even-odd rule. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontexteoclip
func CGContextEOClip(c CGContextRef) {
	_CGContextEOClip(c)
}

// Paints the area within the current path, using the even-odd fill rule. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontexteofillpath
func CGContextEOFillPath(c CGContextRef) {
	_CGContextEOFillPath(c)
}

// Paints the area within the current path, using the nonzero winding number rule. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextfillpath
func CGContextFillPath(c CGContextRef) {
	_CGContextFillPath(c)
}

// Paints the areas contained within the provided rectangles, using the fill color in the current graphics state. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextfillrects
func CGContextFillRects(c CGContextRef, rects unsafe.Pointer, count uintptr) {
	_CGContextFillRects(c, rects, count)
}

// CGContextGetContentToneMappingInfo is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextgetcontenttonemappinginfo
func CGContextGetContentToneMappingInfo(c CGContextRef) unsafe.Pointer {
	return _CGContextGetContentToneMappingInfo(c)
}

// CGContextGetEDRTargetHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextgetedrtargetheadroom(_:)
func CGContextGetEDRTargetHeadroom(c CGContextRef) float32 {
	return _CGContextGetEDRTargetHeadroom(c)
}

// CGContextGetTextPosition is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextgettextposition
func CGContextGetTextPosition(c CGContextRef) CGPoint {
	return _CGContextGetTextPosition(c)
}

// Begins a new subpath at the point you specify. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextmovetopoint
func CGContextMoveToPoint(c CGContextRef, x CGFloat, y CGFloat) {
	_CGContextMoveToPoint(c, x, y)
}

// Decrements the retain count of a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextrelease
func CGContextRelease(c CGContextRef) {
	_CGContextRelease(c)
}

// Increments the retain count of a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextretain
func CGContextRetain(c CGContextRef) CGContextRef {
	return _CGContextRetain(c)
}

// CGContextSetContentToneMappingInfo is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextsetcontenttonemappinginfo
func CGContextSetContentToneMappingInfo(c CGContextRef, info unsafe.Pointer) {
	_CGContextSetContentToneMappingInfo(c, info)
}

// Sets the pattern for dashed lines in a graphics context. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextsetlinedash
func CGContextSetLineDash(c CGContextRef, phase CGFloat, lengths unsafe.Pointer, count uintptr) {
	_CGContextSetLineDash(c, phase, lengths, count)
}

// Sets the current text matrix. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextsettextmatrix
func CGContextSetTextMatrix(c CGContextRef, t CGAffineTransform) {
	_CGContextSetTextMatrix(c, t)
}

// Sets the location at which text is drawn. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextsettextposition
func CGContextSetTextPosition(c CGContextRef, x CGFloat, y CGFloat) {
	_CGContextSetTextPosition(c, x, y)
}

// Draws glyphs at the provided position. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextshowglyphsatpositions
func CGContextShowGlyphsAtPositions(c CGContextRef, glyphs unsafe.Pointer, Lpositions unsafe.Pointer, count uintptr) {
	_CGContextShowGlyphsAtPositions(c, glyphs, Lpositions, count)
}

// Strokes a sequence of line segments. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcontextstrokelinesegments
func CGContextStrokeLineSegments(c CGContextRef, points unsafe.Pointer, count uintptr) {
	_CGContextStrokeLineSegments(c, points, count)
}

// CGConvertColorDataWithFormat is a CoreGraphics function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgconvertcolordatawithformat(_:_:_:_:_:_:_:)
func CGConvertColorDataWithFormat(width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format unsafe.Pointer, src_data unsafe.Pointer, src_format unsafe.Pointer, options unsafe.Pointer) bool {
	return _CGConvertColorDataWithFormat(width, height, dst_data, dst_format, src_data, src_format, options)
}

// Returns a Boolean value indicating whether the mouse cursor is drawn in framebuffer memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcursorisdrawninframebuffer()
func CGCursorIsDrawnInFramebuffer() unsafe.Pointer {
	return _CGCursorIsDrawnInFramebuffer()
}

// Returns a Boolean value indicating whether the mouse cursor is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcursorisvisible()
func CGCursorIsVisible() unsafe.Pointer {
	return _CGCursorIsVisible()
}

// Creates a data consumer that writes to a CFData object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataconsumer/init(data:)
func CGDataConsumerCreateWithCFData(data unsafe.Pointer) CGDataConsumerRef {
	return _CGDataConsumerCreateWithCFData(data)
}

// Creates a data consumer that uses callback functions to write data. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataconsumer/init(info:cbks:)
func CGDataConsumerCreate(info unsafe.Pointer, cbks unsafe.Pointer) CGDataConsumerRef {
	return _CGDataConsumerCreate(info, cbks)
}

// Creates a data consumer that writes data to a location specified by a URL. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataconsumer/init(url:)
func CGDataConsumerCreateWithURL(url unsafe.Pointer) CGDataConsumerRef {
	return _CGDataConsumerCreateWithURL(url)
}

// Returns the Core Foundation type identifier for Core Graphics data consumers. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataconsumer/typeid
func CGDataConsumerGetTypeID() unsafe.Pointer {
	return _CGDataConsumerGetTypeID()
}

// Decrements the retain count of a data consumer. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataconsumerrelease
func CGDataConsumerRelease(consumer CGDataConsumerRef) {
	_CGDataConsumerRelease(consumer)
}

// Increments the retain count of a data consumer. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataconsumerretain
func CGDataConsumerRetain(consumer CGDataConsumerRef) CGDataConsumerRef {
	return _CGDataConsumerRetain(consumer)
}

// Returns a copy of the provider’s data. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovider/data
func CGDataProviderCopyData(provider CGDataProviderRef) unsafe.Pointer {
	return _CGDataProviderCopyData(provider)
}

// CGDataProviderGetInfo is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovider/info
func CGDataProviderGetInfo(provider CGDataProviderRef) unsafe.Pointer {
	return _CGDataProviderGetInfo(provider)
}

// Creates a data provider that reads from a CFData object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovider/init(data:)
func CGDataProviderCreateWithCFData(data unsafe.Pointer) CGDataProviderRef {
	return _CGDataProviderCreateWithCFData(data)
}

// Creates a direct-access data provider that uses data your program supplies. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovider/init(datainfo:data:size:releasedata:)
func CGDataProviderCreateWithData(info unsafe.Pointer, data unsafe.Pointer, size uintptr, releaseData unsafe.Pointer) CGDataProviderRef {
	return _CGDataProviderCreateWithData(info, data, size, releaseData)
}

// Creates a direct-access data provider. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovider/init(directinfo:size:callbacks:)
func CGDataProviderCreateDirect(info unsafe.Pointer, size unsafe.Pointer, callbacks unsafe.Pointer) CGDataProviderRef {
	return _CGDataProviderCreateDirect(info, size, callbacks)
}

// Creates a direct-access data provider that uses a file to supply data. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovider/init(filename:)
func CGDataProviderCreateWithFilename(filename unsafe.Pointer) CGDataProviderRef {
	return _CGDataProviderCreateWithFilename(filename)
}

// Creates a sequential-access data provider. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovider/init(sequentialinfo:callbacks:)
func CGDataProviderCreateSequential(info unsafe.Pointer, callbacks unsafe.Pointer) CGDataProviderRef {
	return _CGDataProviderCreateSequential(info, callbacks)
}

// Creates a direct-access data provider that uses a URL to supply data. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovider/init(url:)
func CGDataProviderCreateWithURL(url unsafe.Pointer) CGDataProviderRef {
	return _CGDataProviderCreateWithURL(url)
}

// Returns the Core Foundation type identifier for data providers. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataprovider/typeid
func CGDataProviderGetTypeID() unsafe.Pointer {
	return _CGDataProviderGetTypeID()
}

// Decrements the retain count of a data provider. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataproviderrelease
func CGDataProviderRelease(provider CGDataProviderRef) {
	_CGDataProviderRelease(provider)
}

// Increments the retain count of a data provider. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdataproviderretain
func CGDataProviderRetain(provider CGDataProviderRef) CGDataProviderRef {
	return _CGDataProviderRetain(provider)
}

// Returns the GPU device instance that’s currently driving a display. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdirectdisplaycopycurrentmetaldevice(_:)
func CGDirectDisplayCopyCurrentMetalDevice(display unsafe.Pointer) unsafe.Pointer {
	return _CGDirectDisplayCopyCurrentMetalDevice(display)
}

// Returns information about the currently available display modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayavailablemodes(_:)
func CGDisplayAvailableModes(dsp unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayAvailableModes(dsp)
}

// Returns information about the display mode closest to a specified depth and screen size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaybestmodeforparameters(_:_:_:_:_:)
func CGDisplayBestModeForParameters(display unsafe.Pointer, bitsPerPixel uintptr, width uintptr, height uintptr, exactMatch unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayBestModeForParameters(display, bitsPerPixel, width, height, exactMatch)
}

// Returns information about the display mode closest to a specified depth, screen size, and refresh rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaybestmodeforparametersandrefreshrate(_:_:_:_:_:_:)
func CGDisplayBestModeForParametersAndRefreshRate(display unsafe.Pointer, bitsPerPixel uintptr, width uintptr, height uintptr, refreshRate unsafe.Pointer, exactMatch unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayBestModeForParametersAndRefreshRate(display, bitsPerPixel, width, height, refreshRate, exactMatch)
}

// Returns the bounds of a display in the global display coordinate space. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaybounds(_:)
func CGDisplayBounds(display unsafe.Pointer) CGRect {
	return _CGDisplayBounds(display)
}

// Obtains exclusive use of a display, preventing other applications and system services from using the display or changing its configuration. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaycapture(_:)
func CGDisplayCapture(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayCapture(display)
}

// Obtains exclusive use of a display for an application using the options you specify. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaycapturewithoptions(_:_:)
func CGDisplayCaptureWithOptions(display unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayCaptureWithOptions(display, options)
}

// Returns information about the currently available display modes. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaycopyalldisplaymodes(_:_:)
func CGDisplayCopyAllDisplayModes(display unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayCopyAllDisplayModes(display, options)
}

// Returns the color space for a display. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaycopycolorspace(_:)
func CGDisplayCopyColorSpace(display unsafe.Pointer) CGColorSpaceRef {
	return _CGDisplayCopyColorSpace(display)
}

// Returns information about a display’s current configuration. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaycopydisplaymode(_:)
func CGDisplayCopyDisplayMode(display unsafe.Pointer) CGDisplayModeRef {
	return _CGDisplayCopyDisplayMode(display)
}

// Returns an image containing the contents of the specified display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaycreateimage(_:)
func CGDisplayCreateImage(displayID unsafe.Pointer) CGImageRef {
	return _CGDisplayCreateImage(displayID)
}

// Returns an image containing the contents of a portion of the specified display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaycreateimage(_:rect:)
func CGDisplayCreateImageForRect(display unsafe.Pointer, rect CGRect) CGImageRef {
	return _CGDisplayCreateImageForRect(display, rect)
}

// Returns information about the current display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaycurrentmode(_:)
func CGDisplayCurrentMode(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayCurrentMode(display)
}

// Performs a single fade operation. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayfade(_:_:_:_:_:_:_:_:)
func CGDisplayFade(token unsafe.Pointer, duration unsafe.Pointer, startBlend unsafe.Pointer, endBlend unsafe.Pointer, redBlend float32, greenBlend float32, blueBlend float32, synchronous unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayFade(token, duration, startBlend, endBlend, redBlend, greenBlend, blueBlend, synchronous)
}

// Returns a Boolean value indicating whether a fade operation is currently in progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayfadeoperationinprogress()
func CGDisplayFadeOperationInProgress() unsafe.Pointer {
	return _CGDisplayFadeOperationInProgress()
}

// Returns the capacity, or number of entries, in the gamma table for a display. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaygammatablecapacity(_:)
func CGDisplayGammaTableCapacity(display unsafe.Pointer) uint32 {
	return _CGDisplayGammaTableCapacity(display)
}

// Returns a graphics context suitable for drawing to a captured display. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaygetdrawingcontext(_:)
func CGDisplayGetDrawingContext(display unsafe.Pointer) CGContextRef {
	return _CGDisplayGetDrawingContext(display)
}

// Hides the mouse cursor, and increments the hide cursor count. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayhidecursor(_:)
func CGDisplayHideCursor(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayHideCursor(display)
}

// Maps a display ID to an OpenGL display mask. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayidtoopengldisplaymask(_:)
func CGDisplayIDToOpenGLDisplayMask(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIDToOpenGLDisplayMask(display)
}

// Returns the I/O Kit service port of the specified display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayioserviceport(_:)
func CGDisplayIOServicePort(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIOServicePort(display)
}

// Returns a Boolean value indicating whether a display is active. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayisactive(_:)
func CGDisplayIsActive(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIsActive(display)
}

// Returns a Boolean value indicating whether a display is always in a mirroring set. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayisalwaysinmirrorset(_:)
func CGDisplayIsAlwaysInMirrorSet(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIsAlwaysInMirrorSet(display)
}

// Returns a Boolean value indicating whether a display is sleeping (and is therefore not drawable). [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayisasleep(_:)
func CGDisplayIsAsleep(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIsAsleep(display)
}

// Returns a Boolean value indicating whether a display is built-in, such as the internal display in portable systems. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayisbuiltin(_:)
func CGDisplayIsBuiltin(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIsBuiltin(display)
}

// Returns a Boolean value indicating whether a display is captured. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayiscaptured(_:)
func CGDisplayIsCaptured(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIsCaptured(display)
}

// Returns a Boolean value indicating whether a display is in a hardware mirroring set. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayisinhwmirrorset(_:)
func CGDisplayIsInHWMirrorSet(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIsInHWMirrorSet(display)
}

// Returns a Boolean value indicating whether a display is in a mirroring set. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayisinmirrorset(_:)
func CGDisplayIsInMirrorSet(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIsInMirrorSet(display)
}

// Returns a Boolean value indicating whether a display is the main display. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayismain(_:)
func CGDisplayIsMain(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIsMain(display)
}

// Returns a Boolean value indicating whether a display is connected or online. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayisonline(_:)
func CGDisplayIsOnline(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIsOnline(display)
}

// Returns a Boolean value indicating whether a display is running in a stereo graphics mode. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayisstereo(_:)
func CGDisplayIsStereo(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayIsStereo(display)
}

// For a secondary display in a mirroring set, returns the primary display. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymirrorsdisplay(_:)
func CGDisplayMirrorsDisplay(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayMirrorsDisplay(display)
}

// Returns the height of the specified display mode. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymode/height
func CGDisplayModeGetHeight(mode CGDisplayModeRef) uintptr {
	return _CGDisplayModeGetHeight(mode)
}

// Returns the I/O Kit display mode ID of the specified display mode. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymode/iodisplaymodeid
func CGDisplayModeGetIODisplayModeID(mode CGDisplayModeRef) unsafe.Pointer {
	return _CGDisplayModeGetIODisplayModeID(mode)
}

// Returns the I/O Kit flags of the specified display mode. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymode/ioflags
func CGDisplayModeGetIOFlags(mode CGDisplayModeRef) uint32 {
	return _CGDisplayModeGetIOFlags(mode)
}

// Returns a Boolean value indicating whether the specified display mode is usable for a desktop graphical user interface. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymode/isusablefordesktopgui()
func CGDisplayModeIsUsableForDesktopGUI(mode CGDisplayModeRef) bool {
	return _CGDisplayModeIsUsableForDesktopGUI(mode)
}

// Returns the pixel encoding of the specified display mode. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymode/pixelencoding
func CGDisplayModeCopyPixelEncoding(mode CGDisplayModeRef) unsafe.Pointer {
	return _CGDisplayModeCopyPixelEncoding(mode)
}

// CGDisplayModeGetPixelHeight is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymode/pixelheight
func CGDisplayModeGetPixelHeight(mode CGDisplayModeRef) uintptr {
	return _CGDisplayModeGetPixelHeight(mode)
}

// CGDisplayModeGetPixelWidth is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymode/pixelwidth
func CGDisplayModeGetPixelWidth(mode CGDisplayModeRef) uintptr {
	return _CGDisplayModeGetPixelWidth(mode)
}

// Returns the refresh rate of the specified display mode. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymode/refreshrate
func CGDisplayModeGetRefreshRate(mode CGDisplayModeRef) float64 {
	return _CGDisplayModeGetRefreshRate(mode)
}

// Returns the type identifier of Quartz display modes. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymode/typeid
func CGDisplayModeGetTypeID() unsafe.Pointer {
	return _CGDisplayModeGetTypeID()
}

// Returns the width of the specified display mode. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymode/width
func CGDisplayModeGetWidth(mode CGDisplayModeRef) uintptr {
	return _CGDisplayModeGetWidth(mode)
}

// Releases a Core Graphics display mode. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymoderelease
func CGDisplayModeRelease(mode CGDisplayModeRef) {
	_CGDisplayModeRelease(mode)
}

// Retains a Core Graphics display mode. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymoderetain
func CGDisplayModeRetain(mode CGDisplayModeRef) CGDisplayModeRef {
	return _CGDisplayModeRetain(mode)
}

// Returns the model number of a display monitor. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymodelnumber(_:)
func CGDisplayModelNumber(display unsafe.Pointer) uint32 {
	return _CGDisplayModelNumber(display)
}

// Moves the mouse cursor to a specified point relative to the upper-left corner of the display. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaymovecursortopoint(_:_:)
func CGDisplayMoveCursorToPoint(display unsafe.Pointer, point CGPoint) unsafe.Pointer {
	return _CGDisplayMoveCursorToPoint(display, point)
}

// Returns the display height in pixel units. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaypixelshigh(_:)
func CGDisplayPixelsHigh(display unsafe.Pointer) uintptr {
	return _CGDisplayPixelsHigh(display)
}

// Returns the display width in pixel units. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaypixelswide(_:)
func CGDisplayPixelsWide(display unsafe.Pointer) uintptr {
	return _CGDisplayPixelsWide(display)
}

// Returns the primary display in a hardware mirroring set. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayprimarydisplay(_:)
func CGDisplayPrimaryDisplay(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayPrimaryDisplay(display)
}

// Registers a callback function to be invoked whenever a local display is reconfigured. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayregisterreconfigurationcallback(_:_:)
func CGDisplayRegisterReconfigurationCallback(callback unsafe.Pointer, userInfo unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayRegisterReconfigurationCallback(callback, userInfo)
}

// Releases a captured display. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayrelease(_:)
func CGDisplayRelease(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayRelease(display)
}

// Removes the registration of a callback function that’s invoked whenever a local display is reconfigured. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayremovereconfigurationcallback(_:_:)
func CGDisplayRemoveReconfigurationCallback(callback unsafe.Pointer, userInfo unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayRemoveReconfigurationCallback(callback, userInfo)
}

// Restores the gamma tables to the values in the user’s ColorSync display profile. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayrestorecolorsyncsettings()
func CGDisplayRestoreColorSyncSettings() {
	_CGDisplayRestoreColorSyncSettings()
}

// Returns the rotation angle of a display in degrees. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayrotation(_:)
func CGDisplayRotation(display unsafe.Pointer) float64 {
	return _CGDisplayRotation(display)
}

// Returns the width and height of a display in millimeters. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayscreensize(_:)
func CGDisplayScreenSize(display unsafe.Pointer) CGSize {
	return _CGDisplayScreenSize(display)
}

// Returns the serial number of a display monitor. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayserialnumber(_:)
func CGDisplaySerialNumber(display unsafe.Pointer) uint32 {
	return _CGDisplaySerialNumber(display)
}

// Switches a display to a different mode. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaysetdisplaymode(_:_:_:)
func CGDisplaySetDisplayMode(display unsafe.Pointer, mode CGDisplayModeRef, options unsafe.Pointer) unsafe.Pointer {
	return _CGDisplaySetDisplayMode(display, mode, options)
}

// Immediately enables or disables stereo operation for a display. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaysetstereooperation(_:_:_:_:)
func CGDisplaySetStereoOperation(display unsafe.Pointer, stereo unsafe.Pointer, forceBlueLine unsafe.Pointer, option unsafe.Pointer) unsafe.Pointer {
	return _CGDisplaySetStereoOperation(display, stereo, forceBlueLine, option)
}

// Decrements the hide cursor count, and shows the mouse cursor if the count is  [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayshowcursor(_:)
func CGDisplayShowCursor(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayShowCursor(display)
}

// Creates a new display stream whose updates are delivered to a dispatch queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystream/init(dispatchqueuedisplay:outputwidth:outputheight:pixelformat:properties:queue:handler:)
func CGDisplayStreamCreateWithDispatchQueue(display unsafe.Pointer, outputWidth uintptr, outputHeight uintptr, pixelFormat unsafe.Pointer, properties unsafe.Pointer, queue unsafe.Pointer, handler unsafe.Pointer) CGDisplayStreamRef {
	return _CGDisplayStreamCreateWithDispatchQueue(display, outputWidth, outputHeight, pixelFormat, properties, queue, handler)
}

// Creates a new display stream to be used with a  [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystream/init(display:outputwidth:outputheight:pixelformat:properties:handler:)
func CGDisplayStreamCreate(display unsafe.Pointer, outputWidth uintptr, outputHeight uintptr, pixelFormat unsafe.Pointer, properties unsafe.Pointer, handler unsafe.Pointer) CGDisplayStreamRef {
	return _CGDisplayStreamCreate(display, outputWidth, outputHeight, pixelFormat, properties, handler)
}

// Gets the run loop source for a display stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystream/runloopsource
func CGDisplayStreamGetRunLoopSource(displayStream CGDisplayStreamRef) unsafe.Pointer {
	return _CGDisplayStreamGetRunLoopSource(displayStream)
}

// Tells a stream to start sending updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystream/start()
func CGDisplayStreamStart(displayStream CGDisplayStreamRef) unsafe.Pointer {
	return _CGDisplayStreamStart(displayStream)
}

// Tells a stream to stop sending updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystream/stop()
func CGDisplayStreamStop(displayStream CGDisplayStreamRef) unsafe.Pointer {
	return _CGDisplayStreamStop(displayStream)
}

// Returns the type identifier of a Quartz display stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystream/typeid
func CGDisplayStreamGetTypeID() unsafe.Pointer {
	return _CGDisplayStreamGetTypeID()
}

// Returns the number of frames that have been dropped since the last call to your update handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystreamupdate/dropcount
func CGDisplayStreamUpdateGetDropCount(updateRef CGDisplayStreamUpdateRef) uintptr {
	return _CGDisplayStreamUpdateGetDropCount(updateRef)
}

// Return the movement delta values for a single update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystreamupdate/getmovedrectsdelta(dx:dy:)
func CGDisplayStreamUpdateGetMovedRectsDelta(updateRef CGDisplayStreamUpdateRef, dx unsafe.Pointer, dy unsafe.Pointer) {
	_CGDisplayStreamUpdateGetMovedRectsDelta(updateRef, dx, dy)
}

// Returns an array of rectangles that describe where the frame has changed since the previous frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystreamupdate/getrects(_:rectcount:)
func CGDisplayStreamUpdateGetRects(updateRef CGDisplayStreamUpdateRef, rectType unsafe.Pointer, rectCount unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayStreamUpdateGetRects(updateRef, rectType, rectCount)
}

// Combines two updates into a new update that includes the metadata for both source updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystreamupdate/init(mergedupdatefirstupdate:secondupdate:)
func CGDisplayStreamUpdateCreateMergedUpdate(firstUpdate CGDisplayStreamUpdateRef, secondUpdate CGDisplayStreamUpdateRef) CGDisplayStreamUpdateRef {
	return _CGDisplayStreamUpdateCreateMergedUpdate(firstUpdate, secondUpdate)
}

// Returns the type identifier of a Quartz display stream update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystreamupdate/typeid
func CGDisplayStreamUpdateGetTypeID() unsafe.Pointer {
	return _CGDisplayStreamUpdateGetTypeID()
}

// Switches a display to a different mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayswitchtomode(_:_:)
func CGDisplaySwitchToMode(display unsafe.Pointer, mode unsafe.Pointer) unsafe.Pointer {
	return _CGDisplaySwitchToMode(display, mode)
}

// Returns the logical unit number of a display. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayunitnumber(_:)
func CGDisplayUnitNumber(display unsafe.Pointer) uint32 {
	return _CGDisplayUnitNumber(display)
}

// Returns a Boolean value indicating whether Quartz is using OpenGL-based window acceleration (Quartz Extreme) to render in a display. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayusesopenglacceleration(_:)
func CGDisplayUsesOpenGLAcceleration(display unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayUsesOpenGLAcceleration(display)
}

// Returns the vendor number of the specified display’s monitor. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayvendornumber(_:)
func CGDisplayVendorNumber(display unsafe.Pointer) uint32 {
	return _CGDisplayVendorNumber(display)
}

// CGEXRToneMappingGammaGetDefaultOptions is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgexrtonemappinggammagetdefaultoptions
func CGEXRToneMappingGammaGetDefaultOptions() unsafe.Pointer {
	return _CGEXRToneMappingGammaGetDefaultOptions()
}

// Enables or disables the merging of actual key and mouse state with the application-specified state in a synthetic event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgenableeventstatecombining(_:)
func CGEnableEventStateCombining(combineState unsafe.Pointer) unsafe.Pointer {
	return _CGEnableEventStateCombining(combineState)
}

// CGErrorSetCallback is a CoreGraphics function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgerrorsetcallback(_:)
func CGErrorSetCallback(callback unsafe.Pointer) {
	_CGErrorSetCallback(callback)
}

// Returns a copy of an existing Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/copy()
func CGEventCreateCopy(event CGEventRef) CGEventRef {
	return _CGEventCreateCopy(event)
}

// Returns the event flags of a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/flags
func CGEventGetFlags(event CGEventRef) unsafe.Pointer {
	return _CGEventGetFlags(event)
}

// Returns the floating-point value of a field in a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/getdoublevaluefield(_:)
func CGEventGetDoubleValueField(event CGEventRef, field unsafe.Pointer) float64 {
	return _CGEventGetDoubleValueField(event, field)
}

// Returns the integer value of a field in a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/getintegervaluefield(_:)
func CGEventGetIntegerValueField(event CGEventRef, field unsafe.Pointer) unsafe.Pointer {
	return _CGEventGetIntegerValueField(event, field)
}

// Returns a new Quartz keyboard event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/init(keyboardeventsource:virtualkey:keydown:)
func CGEventCreateKeyboardEvent(source CGEventSourceRef, virtualKey unsafe.Pointer, keyDown bool) CGEventRef {
	return _CGEventCreateKeyboardEvent(source, virtualKey, keyDown)
}

// Returns a new Quartz mouse event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/init(mouseeventsource:mousetype:mousecursorposition:mousebutton:)
func CGEventCreateMouseEvent(source CGEventSourceRef, mouseType unsafe.Pointer, mouseCursorPosition CGPoint, mouseButton unsafe.Pointer) CGEventRef {
	return _CGEventCreateMouseEvent(source, mouseType, mouseCursorPosition, mouseButton)
}

// CGEventCreateScrollWheelEvent2 is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/init(scrollwheelevent2source:units:wheelcount:wheel1:wheel2:wheel3:)
func CGEventCreateScrollWheelEvent2(source CGEventSourceRef, units unsafe.Pointer, wheelCount uint32, wheel1 unsafe.Pointer, wheel2 unsafe.Pointer, wheel3 unsafe.Pointer) CGEventRef {
	return _CGEventCreateScrollWheelEvent2(source, units, wheelCount, wheel1, wheel2, wheel3)
}

// Returns a new Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/init(source:)
func CGEventCreate(source CGEventSourceRef) CGEventRef {
	return _CGEventCreate(source)
}

// Returns a Quartz event created from a flattened data representation of the event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/init(withdataallocator:data:)
func CGEventCreateFromData(allocator unsafe.Pointer, data unsafe.Pointer) CGEventRef {
	return _CGEventCreateFromData(allocator, data)
}

// Returns the Unicode string associated with a Quartz keyboard event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/keyboardgetunicodestring(maxstringlength:actualstringlength:unicodestring:)
func CGEventKeyboardGetUnicodeString(event CGEventRef, maxStringLength unsafe.Pointer, actualStringLength unsafe.Pointer, unicodeString unsafe.Pointer) {
	_CGEventKeyboardGetUnicodeString(event, maxStringLength, actualStringLength, unicodeString)
}

// Sets the Unicode string associated with a Quartz keyboard event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/keyboardsetunicodestring(stringlength:unicodestring:)
func CGEventKeyboardSetUnicodeString(event CGEventRef, stringLength unsafe.Pointer, unicodeString unsafe.Pointer) {
	_CGEventKeyboardSetUnicodeString(event, stringLength, unicodeString)
}

// Returns the location of a Quartz mouse event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/location
func CGEventGetLocation(event CGEventRef) CGPoint {
	return _CGEventGetLocation(event)
}

// Posts a Quartz event into the event stream at a specified location. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/post(tap:)
func CGEventPost(tap unsafe.Pointer, event CGEventRef) {
	_CGEventPost(tap, event)
}

// Posts a Quartz event into the event stream for a specific application. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/posttopsn(processserialnumber:)
func CGEventPostToPSN(processSerialNumber unsafe.Pointer, event CGEventRef) {
	_CGEventPostToPSN(processSerialNumber, event)
}

// CGEventPostToPid is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/posttopid(_:)
func CGEventPostToPid(pid unsafe.Pointer, event CGEventRef) {
	_CGEventPostToPid(pid, event)
}

// Sets the floating-point value of a field in a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/setdoublevaluefield(_:value:)
func CGEventSetDoubleValueField(event CGEventRef, field unsafe.Pointer, value float64) {
	_CGEventSetDoubleValueField(event, field, value)
}

// Sets the integer value of a field in a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/setintegervaluefield(_:value:)
func CGEventSetIntegerValueField(event CGEventRef, field unsafe.Pointer, value unsafe.Pointer) {
	_CGEventSetIntegerValueField(event, field, value)
}

// Sets the event source of a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/setsource(_:)
func CGEventSetSource(event CGEventRef, source CGEventSourceRef) {
	_CGEventSetSource(event, source)
}

// Creates an event tap. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/tapcreate(tap:place:options:eventsofinterest:callback:userinfo:)
func CGEventTapCreate(tap unsafe.Pointer, place unsafe.Pointer, options unsafe.Pointer, eventsOfInterest unsafe.Pointer, callback unsafe.Pointer, userInfo unsafe.Pointer) unsafe.Pointer {
	return _CGEventTapCreate(tap, place, options, eventsOfInterest, callback, userInfo)
}

// Creates an event tap for a specified process. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/tapcreateforpsn(processserialnumber:place:options:eventsofinterest:callback:userinfo:)
func CGEventTapCreateForPSN(processSerialNumber unsafe.Pointer, place unsafe.Pointer, options unsafe.Pointer, eventsOfInterest unsafe.Pointer, callback unsafe.Pointer, userInfo unsafe.Pointer) unsafe.Pointer {
	return _CGEventTapCreateForPSN(processSerialNumber, place, options, eventsOfInterest, callback, userInfo)
}

// CGEventTapCreateForPid is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/tapcreateforpid(pid:place:options:eventsofinterest:callback:userinfo:)
func CGEventTapCreateForPid(pid unsafe.Pointer, place unsafe.Pointer, options unsafe.Pointer, eventsOfInterest unsafe.Pointer, callback unsafe.Pointer, userInfo unsafe.Pointer) unsafe.Pointer {
	return _CGEventTapCreateForPid(pid, place, options, eventsOfInterest, callback, userInfo)
}

// Enables or disables an event tap. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/tapenable(tap:enable:)
func CGEventTapEnable(tap unsafe.Pointer, enable bool) {
	_CGEventTapEnable(tap, enable)
}

// Returns a Boolean value indicating whether an event tap is enabled. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/tapisenabled(tap:)
func CGEventTapIsEnabled(tap unsafe.Pointer) bool {
	return _CGEventTapIsEnabled(tap)
}

// Posts a Quartz event from an event tap into the event stream. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/tappostevent(_:)
func CGEventTapPostEvent(proxy unsafe.Pointer, event CGEventRef) {
	_CGEventTapPostEvent(proxy, event)
}

// Returns the timestamp of a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/timestamp
func CGEventGetTimestamp(event CGEventRef) unsafe.Pointer {
	return _CGEventGetTimestamp(event)
}

// Returns the event type of a Quartz event (left mouse down, for example). [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/type
func CGEventGetType(event CGEventRef) unsafe.Pointer {
	return _CGEventGetType(event)
}

// Returns the type identifier for the opaque type  [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/typeid
func CGEventGetTypeID() unsafe.Pointer {
	return _CGEventGetTypeID()
}

// Returns the location of a Quartz mouse event. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgevent/unflippedlocation
func CGEventGetUnflippedLocation(event CGEventRef) CGPoint {
	return _CGEventGetUnflippedLocation(event)
}

// Returns a flattened data representation of a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventcreatedata
func CGEventCreateData(allocator unsafe.Pointer, event CGEventRef) unsafe.Pointer {
	return _CGEventCreateData(allocator, event)
}

// Returns a new Quartz scrolling event. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventcreatescrollwheelevent
func CGEventCreateScrollWheelEvent(source CGEventSourceRef, units unsafe.Pointer, wheelCount uint32, wheel1 unsafe.Pointer) CGEventRef {
	return _CGEventCreateScrollWheelEvent(source, units, wheelCount, wheel1)
}

// Sets the event flags of a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsetflags
func CGEventSetFlags(event CGEventRef, flags unsafe.Pointer) {
	_CGEventSetFlags(event, flags)
}

// Sets the location of a Quartz mouse event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsetlocation
func CGEventSetLocation(event CGEventRef, location CGPoint) {
	_CGEventSetLocation(event, location)
}

// Sets the timestamp of a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsettimestamp
func CGEventSetTimestamp(event CGEventRef, timestamp unsafe.Pointer) {
	_CGEventSetTimestamp(event, timestamp)
}

// Sets the event type of a Quartz event (left mouse down, for example). [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsettype
func CGEventSetType(event CGEventRef, type_ unsafe.Pointer) {
	_CGEventSetType(event, type_)
}

// Returns a Boolean value indicating the current button state of a Quartz event source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/buttonstate(_:button:)
func CGEventSourceButtonState(stateID unsafe.Pointer, button unsafe.Pointer) bool {
	return _CGEventSourceButtonState(stateID, button)
}

// Returns a count of events of a given type seen since the window server started. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/counterforeventtype(_:eventtype:)
func CGEventSourceCounterForEventType(stateID unsafe.Pointer, eventType unsafe.Pointer) uint32 {
	return _CGEventSourceCounterForEventType(stateID, eventType)
}

// Returns the current flags of a Quartz event source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/flagsstate(_:)
func CGEventSourceFlagsState(stateID unsafe.Pointer) unsafe.Pointer {
	return _CGEventSourceFlagsState(stateID)
}

// Returns the mask that indicates which classes of local hardware events are enabled during event suppression. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/getlocaleventsfilterduringsuppressionstate(_:)
func CGEventSourceGetLocalEventsFilterDuringSuppressionState(source CGEventSourceRef, state unsafe.Pointer) unsafe.Pointer {
	return _CGEventSourceGetLocalEventsFilterDuringSuppressionState(source, state)
}

// Returns a Quartz event source created from an existing Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/init(event:)
func CGEventCreateSourceFromEvent(event CGEventRef) CGEventSourceRef {
	return _CGEventCreateSourceFromEvent(event)
}

// Returns a Quartz event source created with a specified source state. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/init(stateid:)
func CGEventSourceCreate(stateID unsafe.Pointer) CGEventSourceRef {
	return _CGEventSourceCreate(stateID)
}

// Returns a Boolean value indicating the current keyboard state of a Quartz event source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/keystate(_:key:)
func CGEventSourceKeyState(stateID unsafe.Pointer, key unsafe.Pointer) bool {
	return _CGEventSourceKeyState(stateID, key)
}

// Returns the keyboard type to be used with a Quartz event source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/keyboardtype
func CGEventSourceGetKeyboardType(source CGEventSourceRef) unsafe.Pointer {
	return _CGEventSourceGetKeyboardType(source)
}

// Returns the interval that local hardware events may be suppressed following the posting of a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/localeventssuppressioninterval
func CGEventSourceGetLocalEventsSuppressionInterval(source CGEventSourceRef) unsafe.Pointer {
	return _CGEventSourceGetLocalEventsSuppressionInterval(source)
}

// Gets the scale of pixels per line in a scrolling event source. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/pixelsperline
func CGEventSourceGetPixelsPerLine(source CGEventSourceRef) float64 {
	return _CGEventSourceGetPixelsPerLine(source)
}

// Returns the elapsed time since the last event for a Quartz event source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/secondssincelasteventtype(_:eventtype:)
func CGEventSourceSecondsSinceLastEventType(stateID unsafe.Pointer, eventType unsafe.Pointer) unsafe.Pointer {
	return _CGEventSourceSecondsSinceLastEventType(stateID, eventType)
}

// Sets the mask that indicates which classes of local hardware events are enabled during event suppression. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/setlocaleventsfilterduringsuppressionstate(_:state:)
func CGEventSourceSetLocalEventsFilterDuringSuppressionState(source CGEventSourceRef, filter unsafe.Pointer, state unsafe.Pointer) {
	_CGEventSourceSetLocalEventsFilterDuringSuppressionState(source, filter, state)
}

// Returns the source state associated with a Quartz event source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/sourcestateid
func CGEventSourceGetSourceStateID(source CGEventSourceRef) unsafe.Pointer {
	return _CGEventSourceGetSourceStateID(source)
}

// Returns the type identifier for the opaque type  [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/typeid
func CGEventSourceGetTypeID() unsafe.Pointer {
	return _CGEventSourceGetTypeID()
}

// Returns the 64-bit user-specified data for a Quartz event source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsource/userdata
func CGEventSourceGetUserData(source CGEventSourceRef) unsafe.Pointer {
	return _CGEventSourceGetUserData(source)
}

// Sets the keyboard type to be used with a Quartz event source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsourcesetkeyboardtype
func CGEventSourceSetKeyboardType(source CGEventSourceRef, keyboardType unsafe.Pointer) {
	_CGEventSourceSetKeyboardType(source, keyboardType)
}

// Sets the interval that local hardware events may be suppressed following the posting of a Quartz event. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsourcesetlocaleventssuppressioninterval
func CGEventSourceSetLocalEventsSuppressionInterval(source CGEventSourceRef, seconds unsafe.Pointer) {
	_CGEventSourceSetLocalEventsSuppressionInterval(source, seconds)
}

// Sets the scale of pixels per line in a scrolling event source. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsourcesetpixelsperline
func CGEventSourceSetPixelsPerLine(source CGEventSourceRef, pixelsPerLine float64) {
	_CGEventSourceSetPixelsPerLine(source, pixelsPerLine)
}

// Sets the 64-bit user-specified data for a Quartz event source. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsourcesetuserdata
func CGEventSourceSetUserData(source CGEventSourceRef, userData unsafe.Pointer) {
	_CGEventSourceSetUserData(source, userData)
}

// Returns the ascent of a font. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/ascent
func CGFontGetAscent(font CGFontRef) int {
	return _CGFontGetAscent(font)
}

// Determines whether Core Graphics can create a subset of the font in PostScript format. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/cancreatepostscriptsubset(_:)
func CGFontCanCreatePostScriptSubset(font CGFontRef, format unsafe.Pointer) bool {
	return _CGFontCanCreatePostScriptSubset(font, format)
}

// Returns the cap height of a font. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/capheight
func CGFontGetCapHeight(font CGFontRef) int {
	return _CGFontGetCapHeight(font)
}

// Creates a copy of a font using a variation specification dictionary. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/copy(withvariations:)
func CGFontCreateCopyWithVariations(font CGFontRef, variations unsafe.Pointer) CGFontRef {
	return _CGFontCreateCopyWithVariations(font, variations)
}

// Creates a PostScript encoding of a font. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/createpostscriptencoding(encoding:)
func CGFontCreatePostScriptEncoding(font CGFontRef, encoding unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _CGFontCreatePostScriptEncoding(font, encoding, p2)
}

// Creates a subset of the font in the specified PostScript format. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/createpostscriptsubset(subsetname:format:glyphs:count:encoding:)
func CGFontCreatePostScriptSubset(font CGFontRef, subsetName unsafe.Pointer, format unsafe.Pointer, glyphs unsafe.Pointer, count uintptr, encoding unsafe.Pointer, p6 unsafe.Pointer) unsafe.Pointer {
	return _CGFontCreatePostScriptSubset(font, subsetName, format, glyphs, count, encoding, p6)
}

// Returns the descent of a font. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/descent
func CGFontGetDescent(font CGFontRef) int {
	return _CGFontGetDescent(font)
}

// Returns the bounding box of a font. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/fontbbox
func CGFontGetFontBBox(font CGFontRef) CGRect {
	return _CGFontGetFontBBox(font)
}

// Returns the full name associated with a font object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/fullname
func CGFontCopyFullName(font CGFontRef) unsafe.Pointer {
	return _CGFontCopyFullName(font)
}

// Gets the advance width of each glyph in the provided array. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/getglyphadvances(glyphs:count:advances:)
func CGFontGetGlyphAdvances(font CGFontRef, glyphs unsafe.Pointer, count uintptr, advances unsafe.Pointer) bool {
	return _CGFontGetGlyphAdvances(font, glyphs, count, advances)
}

// Get the bounding box of each glyph in an array. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/getglyphbboxes(glyphs:count:bboxes:)
func CGFontGetGlyphBBoxes(font CGFontRef, glyphs unsafe.Pointer, count uintptr, bboxes unsafe.Pointer) bool {
	return _CGFontGetGlyphBBoxes(font, glyphs, count, bboxes)
}

// Returns the glyph for the glyph name associated with the specified font object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/getglyphwithglyphname(name:)
func CGFontGetGlyphWithGlyphName(font CGFontRef, name unsafe.Pointer) unsafe.Pointer {
	return _CGFontGetGlyphWithGlyphName(font, name)
}

// Creates a font object corresponding to the font specified by a PostScript or full name. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/init(_:)-1p4b
func CGFontCreateWithFontName(name unsafe.Pointer) CGFontRef {
	return _CGFontCreateWithFontName(name)
}

// Creates a font object from data supplied from a data provider. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/init(_:)-9aour
func CGFontCreateWithDataProvider(provider CGDataProviderRef) CGFontRef {
	return _CGFontCreateWithDataProvider(provider)
}

// Returns the italic angle of a font. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/italicangle
func CGFontGetItalicAngle(font CGFontRef) CGFloat {
	return _CGFontGetItalicAngle(font)
}

// Returns the leading of a font. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/leading
func CGFontGetLeading(font CGFontRef) int {
	return _CGFontGetLeading(font)
}

// Returns the glyph name of the specified glyph in the specified font. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/name(for:)
func CGFontCopyGlyphNameForGlyph(font CGFontRef, glyph unsafe.Pointer) unsafe.Pointer {
	return _CGFontCopyGlyphNameForGlyph(font, glyph)
}

// Returns the number of glyphs in a font. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/numberofglyphs
func CGFontGetNumberOfGlyphs(font CGFontRef) uintptr {
	return _CGFontGetNumberOfGlyphs(font)
}

// Obtains the PostScript name of a font. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/postscriptname
func CGFontCopyPostScriptName(font CGFontRef) unsafe.Pointer {
	return _CGFontCopyPostScriptName(font)
}

// Returns the thickness of the dominant vertical stems of glyphs in a font. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/stemv
func CGFontGetStemV(font CGFontRef) CGFloat {
	return _CGFontGetStemV(font)
}

// Returns the font table that corresponds to the provided tag. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/table(for:)
func CGFontCopyTableForTag(font CGFontRef, tag uint32) unsafe.Pointer {
	return _CGFontCopyTableForTag(font, tag)
}

// Returns an array of tags that correspond to the font tables for a font. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/tabletags
func CGFontCopyTableTags(font CGFontRef) unsafe.Pointer {
	return _CGFontCopyTableTags(font)
}

// Returns the Core Foundation type identifier for Core Graphics fonts. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/typeid
func CGFontGetTypeID() unsafe.Pointer {
	return _CGFontGetTypeID()
}

// Returns the number of glyph space units per em for the provided font. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/unitsperem
func CGFontGetUnitsPerEm(font CGFontRef) int {
	return _CGFontGetUnitsPerEm(font)
}

// Returns an array of the variation axis dictionaries for a font. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/variationaxes
func CGFontCopyVariationAxes(font CGFontRef) unsafe.Pointer {
	return _CGFontCopyVariationAxes(font)
}

// Returns the variation specification dictionary for a font. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/variations
func CGFontCopyVariations(font CGFontRef) unsafe.Pointer {
	return _CGFontCopyVariations(font)
}

// Returns the x-height of a font. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfont/xheight
func CGFontGetXHeight(font CGFontRef) int {
	return _CGFontGetXHeight(font)
}

// Creates a font object from an Apple Type Services (ATS) font. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfontcreatewithplatformfont
func CGFontCreateWithPlatformFont(platformFontReference unsafe.Pointer) CGFontRef {
	return _CGFontCreateWithPlatformFont(platformFontReference)
}

// Decrements the retain count of a font. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfontrelease
func CGFontRelease(font CGFontRef) {
	_CGFontRelease(font)
}

// Increments the retain count of a font. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfontretain
func CGFontRetain(font CGFontRef) CGFontRef {
	return _CGFontRetain(font)
}

// Creates a Core Graphics function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfunction/init(info:domaindimension:domain:rangedimension:range:callbacks:)
func CGFunctionCreate(info unsafe.Pointer, domainDimension uintptr, domain unsafe.Pointer, rangeDimension uintptr, range_ unsafe.Pointer, callbacks unsafe.Pointer) CGFunctionRef {
	return _CGFunctionCreate(info, domainDimension, domain, rangeDimension, range_, callbacks)
}

// Returns the type identifier for Core Graphics function objects. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfunction/typeid
func CGFunctionGetTypeID() unsafe.Pointer {
	return _CGFunctionGetTypeID()
}

// Decrements the retain count of a function object. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfunctionrelease
func CGFunctionRelease(function CGFunctionRef) {
	_CGFunctionRelease(function)
}

// Increments the retain count of a function object. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfunctionretain
func CGFunctionRetain(function CGFunctionRef) CGFunctionRef {
	return _CGFunctionRetain(function)
}

// Provides a list of displays that are active for drawing. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggetactivedisplaylist(_:_:_:)
func CGGetActiveDisplayList(maxDisplays uint32, activeDisplays unsafe.Pointer, displayCount unsafe.Pointer) unsafe.Pointer {
	return _CGGetActiveDisplayList(maxDisplays, activeDisplays, displayCount)
}

// Gets the coefficients of the gamma transfer formula for a display. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggetdisplaytransferbyformula(_:_:_:_:_:_:_:_:_:_:)
func CGGetDisplayTransferByFormula(display unsafe.Pointer, redMin unsafe.Pointer, redMax unsafe.Pointer, redGamma unsafe.Pointer, greenMin unsafe.Pointer, greenMax unsafe.Pointer, greenGamma unsafe.Pointer, blueMin unsafe.Pointer, blueMax unsafe.Pointer, blueGamma unsafe.Pointer) unsafe.Pointer {
	return _CGGetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
}

// Gets the values in the RGB gamma tables for a display. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggetdisplaytransferbytable(_:_:_:_:_:_:)
func CGGetDisplayTransferByTable(display unsafe.Pointer, capacity uint32, redTable unsafe.Pointer, greenTable unsafe.Pointer, blueTable unsafe.Pointer, sampleCount unsafe.Pointer) unsafe.Pointer {
	return _CGGetDisplayTransferByTable(display, capacity, redTable, greenTable, blueTable, sampleCount)
}

// Provides a list of displays that corresponds to the bits set in an OpenGL display mask. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggetdisplayswithopengldisplaymask(_:_:_:_:)
func CGGetDisplaysWithOpenGLDisplayMask(mask unsafe.Pointer, maxDisplays uint32, displays unsafe.Pointer, matchingDisplayCount unsafe.Pointer) unsafe.Pointer {
	return _CGGetDisplaysWithOpenGLDisplayMask(mask, maxDisplays, displays, matchingDisplayCount)
}

// Provides a list of online displays with bounds that include the specified point. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggetdisplayswithpoint(_:_:_:_:)
func CGGetDisplaysWithPoint(point CGPoint, maxDisplays uint32, displays unsafe.Pointer, matchingDisplayCount unsafe.Pointer) unsafe.Pointer {
	return _CGGetDisplaysWithPoint(point, maxDisplays, displays, matchingDisplayCount)
}

// Gets a list of online displays with bounds that intersect the specified rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggetdisplayswithrect(_:_:_:_:)
func CGGetDisplaysWithRect(rect CGRect, maxDisplays uint32, displays unsafe.Pointer, matchingDisplayCount unsafe.Pointer) unsafe.Pointer {
	return _CGGetDisplaysWithRect(rect, maxDisplays, displays, matchingDisplayCount)
}

// Gets a list of currently installed event taps. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggeteventtaplist(_:_:_:)
func CGGetEventTapList(maxNumberOfTaps uint32, tapList unsafe.Pointer, eventTapCount unsafe.Pointer) unsafe.Pointer {
	return _CGGetEventTapList(maxNumberOfTaps, tapList, eventTapCount)
}

// Reports the change in mouse position since the last mouse movement event received by the application. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggetlastmousedelta
func CGGetLastMouseDelta(deltaX unsafe.Pointer, deltaY unsafe.Pointer) {
	_CGGetLastMouseDelta(deltaX, deltaY)
}

// Provides a list of displays that are online (active, mirrored, or sleeping). [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggetonlinedisplaylist(_:_:_:)
func CGGetOnlineDisplayList(maxDisplays uint32, onlineDisplays unsafe.Pointer, displayCount unsafe.Pointer) unsafe.Pointer {
	return _CGGetOnlineDisplayList(maxDisplays, onlineDisplays, displayCount)
}

// CGGradientGetContentHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggradient/contentheadroom
func CGGradientGetContentHeadroom(gradient CGGradientRef) float32 {
	return _CGGradientGetContentHeadroom(gradient)
}

// Creates a CGGradient object from a color space and the provided color components and locations. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggradient/init(colorspace:colorcomponents:locations:count:)
func CGGradientCreateWithColorComponents(space CGColorSpaceRef, components unsafe.Pointer, locations unsafe.Pointer, count uintptr) CGGradientRef {
	return _CGGradientCreateWithColorComponents(space, components, locations, count)
}

// Creates a gradient object from a color space and the provided color objects and locations. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggradient/init(colorsspace:colors:locations:)
func CGGradientCreateWithColors(space CGColorSpaceRef, colors unsafe.Pointer, locations unsafe.Pointer) CGGradientRef {
	return _CGGradientCreateWithColors(space, colors, locations)
}

// CGGradientCreateWithContentHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggradient/init(headroom:colorspace:colorcomponents:locations:count:)
func CGGradientCreateWithContentHeadroom(headroom float32, space CGColorSpaceRef, components unsafe.Pointer, locations unsafe.Pointer, count uintptr) CGGradientRef {
	return _CGGradientCreateWithContentHeadroom(headroom, space, components, locations, count)
}

// Returns the Core Foundation type identifier for CGGradient objects. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggradient/typeid
func CGGradientGetTypeID() unsafe.Pointer {
	return _CGGradientGetTypeID()
}

// Decrements the retain count of a CGGradient object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggradientrelease
func CGGradientRelease(gradient CGGradientRef) {
	_CGGradientRelease(gradient)
}

// Increments the retain count of a CGGradient object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggradientretain
func CGGradientRetain(gradient CGGradientRef) CGGradientRef {
	return _CGGradientRetain(gradient)
}

// Returns the alpha channel information for a bitmap image. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/alphainfo
func CGImageGetAlphaInfo(image CGImageRef) unsafe.Pointer {
	return _CGImageGetAlphaInfo(image)
}

// Returns the bitmap information for a bitmap image. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/bitmapinfo
func CGImageGetBitmapInfo(image CGImageRef) unsafe.Pointer {
	return _CGImageGetBitmapInfo(image)
}

// Returns the number of bits allocated for a single color component of a bitmap image. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/bitspercomponent
func CGImageGetBitsPerComponent(image CGImageRef) uintptr {
	return _CGImageGetBitsPerComponent(image)
}

// Returns the number of bits allocated for a single pixel in a bitmap image. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/bitsperpixel
func CGImageGetBitsPerPixel(image CGImageRef) uintptr {
	return _CGImageGetBitsPerPixel(image)
}

// CGImageGetByteOrderInfo is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/byteorderinfo
func CGImageGetByteOrderInfo(image CGImageRef) unsafe.Pointer {
	return _CGImageGetByteOrderInfo(image)
}

// Returns the number of bytes allocated for a single row of a bitmap image. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/bytesperrow
func CGImageGetBytesPerRow(image CGImageRef) uintptr {
	return _CGImageGetBytesPerRow(image)
}

// CGImageCalculateContentAverageLightLevel is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/calculatedcontentaveragelightlevel
func CGImageCalculateContentAverageLightLevel(image CGImageRef) float32 {
	return _CGImageCalculateContentAverageLightLevel(image)
}

// CGImageCalculateContentHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/calculatedcontentheadroom
func CGImageCalculateContentHeadroom(image CGImageRef) float32 {
	return _CGImageCalculateContentHeadroom(image)
}

// Return the color space for a bitmap image. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/colorspace
func CGImageGetColorSpace(image CGImageRef) CGColorSpaceRef {
	return _CGImageGetColorSpace(image)
}

// CGImageContainsImageSpecificToneMappingMetadata is a CoreGraphics function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/containsimagespecifictonemappingmetadata
func CGImageContainsImageSpecificToneMappingMetadata(image CGImageRef) bool {
	return _CGImageContainsImageSpecificToneMappingMetadata(image)
}

// CGImageGetContentAverageLightLevel is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/contentaveragelightlevel
func CGImageGetContentAverageLightLevel(image CGImageRef) float32 {
	return _CGImageGetContentAverageLightLevel(image)
}

// CGImageGetContentHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/contentheadroom
func CGImageGetContentHeadroom(image CGImageRef) float32 {
	return _CGImageGetContentHeadroom(image)
}

// Creates a copy of a bitmap image. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/copy()
func CGImageCreateCopy(image CGImageRef) CGImageRef {
	return _CGImageCreateCopy(image)
}

// Creates a copy of a bitmap image, replacing its colorspace. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/copy(colorspace:)
func CGImageCreateCopyWithColorSpace(image CGImageRef, space CGColorSpaceRef) CGImageRef {
	return _CGImageCreateCopyWithColorSpace(image, space)
}

// CGImageCreateCopyWithContentAverageLightLevel is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/copy(contentaveragelightlevel:)
func CGImageCreateCopyWithContentAverageLightLevel(image CGImageRef, avll float32) CGImageRef {
	return _CGImageCreateCopyWithContentAverageLightLevel(image, avll)
}

// CGImageCreateCopyWithCalculatedHDRStats is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/copywithcalculatedhdrstats()
func CGImageCreateCopyWithCalculatedHDRStats(image CGImageRef) CGImageRef {
	return _CGImageCreateCopyWithCalculatedHDRStats(image)
}

// Creates a bitmap image using the data contained within a subregion of an existing bitmap image. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/cropping(to:)
func CGImageCreateWithImageInRect(image CGImageRef, rect CGRect) CGImageRef {
	return _CGImageCreateWithImageInRect(image, rect)
}

// Returns the data provider for a bitmap image or image mask. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/dataprovider
func CGImageGetDataProvider(image CGImageRef) CGDataProviderRef {
	return _CGImageGetDataProvider(image)
}

// Returns the decode array for a bitmap image. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/decode
func CGImageGetDecode(image CGImageRef) unsafe.Pointer {
	return _CGImageGetDecode(image)
}

// Returns the height of a bitmap image. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/height
func CGImageGetHeight(image CGImageRef) uintptr {
	return _CGImageGetHeight(image)
}

// CGImageCreateWithContentHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/init(headroom:width:height:bitspercomponent:bitsperpixel:bytesperrow:space:bitmapinfo:provider:decode:shouldinterpolate:intent:)
func CGImageCreateWithContentHeadroom(headroom float32, width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo unsafe.Pointer, provider CGDataProviderRef, decode unsafe.Pointer, shouldInterpolate bool, intent unsafe.Pointer) CGImageRef {
	return _CGImageCreateWithContentHeadroom(headroom, width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, bitmapInfo, provider, decode, shouldInterpolate, intent)
}

// Creates a bitmap image using JPEG-encoded data supplied by a data provider. [Full Topic]
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/init(jpegdataprovidersource:decode:shouldinterpolate:intent:)
func CGImageCreateWithJPEGDataProvider(source CGDataProviderRef, decode unsafe.Pointer, shouldInterpolate bool, intent unsafe.Pointer) CGImageRef {
	return _CGImageCreateWithJPEGDataProvider(source, decode, shouldInterpolate, intent)
}

// Creates a bitmap image mask from data supplied by a data provider. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/init(maskwidth:height:bitspercomponent:bitsperpixel:bytesperrow:provider:decode:shouldinterpolate:)
func CGImageMaskCreate(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, provider CGDataProviderRef, decode unsafe.Pointer, shouldInterpolate bool) CGImageRef {
	return _CGImageMaskCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, provider, decode, shouldInterpolate)
}

// Creates a bitmap image using PNG-encoded data supplied by a data provider. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/init(pngdataprovidersource:decode:shouldinterpolate:intent:)
func CGImageCreateWithPNGDataProvider(source CGDataProviderRef, decode unsafe.Pointer, shouldInterpolate bool, intent unsafe.Pointer) CGImageRef {
	return _CGImageCreateWithPNGDataProvider(source, decode, shouldInterpolate, intent)
}

// Creates a bitmap image from data supplied by a data provider. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/init(width:height:bitspercomponent:bitsperpixel:bytesperrow:space:bitmapinfo:provider:decode:shouldinterpolate:intent:)
func CGImageCreate(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo unsafe.Pointer, provider CGDataProviderRef, decode unsafe.Pointer, shouldInterpolate bool, intent unsafe.Pointer) CGImageRef {
	return _CGImageCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, bitmapInfo, provider, decode, shouldInterpolate, intent)
}

// Returns a composite image of the specified windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/init(windowlistfromarrayscreenbounds:windowarray:imageoption:)
func CGWindowListCreateImageFromArray(screenBounds CGRect, windowArray unsafe.Pointer, imageOption unsafe.Pointer) CGImageRef {
	return _CGWindowListCreateImageFromArray(screenBounds, windowArray, imageOption)
}

// Returns whether a bitmap image is an image mask. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/ismask
func CGImageIsMask(image CGImageRef) bool {
	return _CGImageIsMask(image)
}

// Creates a bitmap image from an existing image and an image mask. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/masking(_:)
func CGImageCreateWithMask(image CGImageRef, mask CGImageRef) CGImageRef {
	return _CGImageCreateWithMask(image, mask)
}

// CGImageGetPixelFormatInfo is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/pixelformatinfo
func CGImageGetPixelFormatInfo(image CGImageRef) unsafe.Pointer {
	return _CGImageGetPixelFormatInfo(image)
}

// Returns the rendering intent setting for a bitmap image. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/renderingintent
func CGImageGetRenderingIntent(image CGImageRef) unsafe.Pointer {
	return _CGImageGetRenderingIntent(image)
}

// Returns the interpolation setting for a bitmap image. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/shouldinterpolate
func CGImageGetShouldInterpolate(image CGImageRef) bool {
	return _CGImageGetShouldInterpolate(image)
}

// CGImageShouldToneMap is a CoreGraphics function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/shouldtonemap
func CGImageShouldToneMap(image CGImageRef) bool {
	return _CGImageShouldToneMap(image)
}

// Returns the type identifier for CGImage objects. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/typeid
func CGImageGetTypeID() unsafe.Pointer {
	return _CGImageGetTypeID()
}

// The Universal Type Identifier for the image. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/uttype
func CGImageGetUTType(image CGImageRef) unsafe.Pointer {
	return _CGImageGetUTType(image)
}

// Returns the width of a bitmap image, in pixels. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimage/width
func CGImageGetWidth(image CGImageRef) uintptr {
	return _CGImageGetWidth(image)
}

// CGImageCreateCopyWithContentHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimagecreatecopywithcontentheadroom(_:_:)
func CGImageCreateCopyWithContentHeadroom(headroom float32, image CGImageRef) CGImageRef {
	return _CGImageCreateCopyWithContentHeadroom(headroom, image)
}

// Creates a bitmap image by masking an existing bitmap image with the provided color values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimagecreatewithmaskingcolors
func CGImageCreateWithMaskingColors(image CGImageRef, components unsafe.Pointer) CGImageRef {
	return _CGImageCreateWithMaskingColors(image, components)
}

// Decrements the retain count of a bitmap image. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimagerelease
func CGImageRelease(image CGImageRef) {
	_CGImageRelease(image)
}

// Increments the retain count of a bitmap image. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimageretain
func CGImageRetain(image CGImageRef) CGImageRef {
	return _CGImageRetain(image)
}

// Turns off local hardware events in the current session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cginhibitlocalevents(_:)
func CGInhibitLocalEvents(inhibit unsafe.Pointer) unsafe.Pointer {
	return _CGInhibitLocalEvents(inhibit)
}

// Returns the graphics context associated with a layer object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cglayer/context
func CGLayerGetContext(layer CGLayerRef) CGContextRef {
	return _CGLayerGetContext(layer)
}

// Creates a layer object that is associated with a graphics context. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cglayer/init(_:size:auxiliaryinfo:)
func CGLayerCreateWithContext(context CGContextRef, size CGSize, auxiliaryInfo unsafe.Pointer) CGLayerRef {
	return _CGLayerCreateWithContext(context, size, auxiliaryInfo)
}

// Returns the width and height of a layer object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cglayer/size
func CGLayerGetSize(layer CGLayerRef) CGSize {
	return _CGLayerGetSize(layer)
}

// Returns the unique type identifier used for  [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cglayer/typeid
func CGLayerGetTypeID() unsafe.Pointer {
	return _CGLayerGetTypeID()
}

// Decrements the retain count of a layer object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cglayerrelease
func CGLayerRelease(layer CGLayerRef) {
	_CGLayerRelease(layer)
}

// Increments the retain count of a layer object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cglayerretain
func CGLayerRetain(layer CGLayerRef) CGLayerRef {
	return _CGLayerRetain(layer)
}

// Returns the display ID of the main display. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgmaindisplayid()
func CGMainDisplayID() unsafe.Pointer {
	return _CGMainDisplayID()
}

// Closes and completes a subpath in a mutable graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgmutablepath/closesubpath()
func CGPathCloseSubpath(path CGMutablePathRef) {
	_CGPathCloseSubpath(path)
}

// Creates a mutable graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgmutablepath/init()
func CGPathCreateMutable() CGMutablePathRef {
	return _CGPathCreateMutable()
}

// Maps an OpenGL display mask to a display ID. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgopengldisplaymasktodisplayid(_:)
func CGOpenGLDisplayMaskToDisplayID(mask unsafe.Pointer) unsafe.Pointer {
	return _CGOpenGLDisplayMaskToDisplayID(mask)
}

// CGPDFArrayApplyBlock is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarrayapplyblock(_:_:_:)
func CGPDFArrayApplyBlock(array CGPDFArrayRef, block unsafe.Pointer, info unsafe.Pointer) {
	_CGPDFArrayApplyBlock(array, block, info)
}

// Returns whether an object at a given index in a PDF array is another PDF array and, if so, retrieves that array. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarraygetarray(_:_:_:)
func CGPDFArrayGetArray(array CGPDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetArray(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF Boolean and, if so, retrieves that Boolean. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarraygetboolean(_:_:_:)
func CGPDFArrayGetBoolean(array CGPDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetBoolean(array, index, value)
}

// Returns the number of items in a PDF array. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarraygetcount(_:)
func CGPDFArrayGetCount(array CGPDFArrayRef) uintptr {
	return _CGPDFArrayGetCount(array)
}

// Returns whether an object at a given index in a PDF array is a PDF dictionary and, if so, retrieves that dictionary. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarraygetdictionary(_:_:_:)
func CGPDFArrayGetDictionary(array CGPDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetDictionary(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF integer and, if so, retrieves that object. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarraygetinteger(_:_:_:)
func CGPDFArrayGetInteger(array CGPDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetInteger(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF name reference (represented as a constant C string) and, if so, retrieves that name. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarraygetname(_:_:_:)
func CGPDFArrayGetName(array CGPDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetName(array, index, value)
}

// Returns whether an object at a given index in a Quartz PDF array is a PDF null. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarraygetnull(_:_:)
func CGPDFArrayGetNull(array CGPDFArrayRef, index uintptr) bool {
	return _CGPDFArrayGetNull(array, index)
}

// Returns whether an object at a given index in a PDF array is a PDF number and, if so, retrieves that object. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarraygetnumber(_:_:_:)
func CGPDFArrayGetNumber(array CGPDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetNumber(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF object and, if so, retrieves that object. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarraygetobject(_:_:_:)
func CGPDFArrayGetObject(array CGPDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetObject(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF stream and, if so, retrieves that stream. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarraygetstream(_:_:_:)
func CGPDFArrayGetStream(array CGPDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetStream(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF string and, if so, retrieves that string. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfarraygetstring(_:_:_:)
func CGPDFArrayGetString(array CGPDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetString(array, index, value)
}

// Creates a content stream object from a PDF page object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontentstreamcreatewithpage(_:)
func CGPDFContentStreamCreateWithPage(page CGPDFPageRef) CGPDFContentStreamRef {
	return _CGPDFContentStreamCreateWithPage(page)
}

// Creates a PDF content stream object from an existing PDF content stream object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontentstreamcreatewithstream(_:_:_:)
func CGPDFContentStreamCreateWithStream(stream CGPDFStreamRef, streamResources CGPDFDictionaryRef, parent CGPDFContentStreamRef) CGPDFContentStreamRef {
	return _CGPDFContentStreamCreateWithStream(stream, streamResources, parent)
}

// Gets the specified resource from a PDF content stream object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontentstreamgetresource(_:_:_:)
func CGPDFContentStreamGetResource(cs CGPDFContentStreamRef, category unsafe.Pointer, name unsafe.Pointer) CGPDFObjectRef {
	return _CGPDFContentStreamGetResource(cs, category, name)
}

// Gets the array of PDF content streams contained in a PDF content stream object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontentstreamgetstreams(_:)
func CGPDFContentStreamGetStreams(cs CGPDFContentStreamRef) unsafe.Pointer {
	return _CGPDFContentStreamGetStreams(cs)
}

// Decrements the retain count of a PDF content stream object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontentstreamrelease(_:)
func CGPDFContentStreamRelease(cs CGPDFContentStreamRef) {
	_CGPDFContentStreamRelease(cs)
}

// Increments the retain count of a PDF content stream object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontentstreamretain(_:)
func CGPDFContentStreamRetain(cs CGPDFContentStreamRef) CGPDFContentStreamRef {
	return _CGPDFContentStreamRetain(cs)
}

// CGPDFContextBeginTag is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontextbegintag(_:_:_:)
func CGPDFContextBeginTag(context CGContextRef, tagType unsafe.Pointer, tagProperties unsafe.Pointer) {
	_CGPDFContextBeginTag(context, tagType, tagProperties)
}

// CGPDFContextEndTag is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontextendtag(_:)
func CGPDFContextEndTag(context CGContextRef) {
	_CGPDFContextEndTag(context)
}

// CGPDFContextSetIDTree is a CoreGraphics function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontextsetidtree(_:_:)
func CGPDFContextSetIDTree(context CGContextRef, IDTreeDictionary CGPDFDictionaryRef) {
	_CGPDFContextSetIDTree(context, IDTreeDictionary)
}

// CGPDFContextSetOutline is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontextsetoutline(_:_:)
func CGPDFContextSetOutline(context CGContextRef, outline unsafe.Pointer) {
	_CGPDFContextSetOutline(context, outline)
}

// CGPDFContextSetPageTagStructureTree is a CoreGraphics function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontextsetpagetagstructuretree(_:_:)
func CGPDFContextSetPageTagStructureTree(context CGContextRef, pageTagStructureTreeDictionary unsafe.Pointer) {
	_CGPDFContextSetPageTagStructureTree(context, pageTagStructureTreeDictionary)
}

// CGPDFContextSetParentTree is a CoreGraphics function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfcontextsetparenttree(_:_:)
func CGPDFContextSetParentTree(context CGContextRef, parentTreeDictionary CGPDFDictionaryRef) {
	_CGPDFContextSetParentTree(context, parentTreeDictionary)
}

// CGPDFDictionaryApplyBlock is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionaryapplyblock(_:_:_:)
func CGPDFDictionaryApplyBlock(dict CGPDFDictionaryRef, block unsafe.Pointer, info unsafe.Pointer) {
	_CGPDFDictionaryApplyBlock(dict, block, info)
}

// Applies a function to each entry in a dictionary. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionaryapplyfunction(_:_:_:)
func CGPDFDictionaryApplyFunction(dict CGPDFDictionaryRef, function unsafe.Pointer, info unsafe.Pointer) {
	_CGPDFDictionaryApplyFunction(dict, function, info)
}

// Returns whether there is a PDF array associated with a specified key in a PDF dictionary and, if so, retrieves that array. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionarygetarray(_:_:_:)
func CGPDFDictionaryGetArray(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetArray(dict, key, value)
}

// Returns whether there is a PDF Boolean value associated with a specified key in a PDF dictionary and, if so, retrieves the Boolean value. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionarygetboolean(_:_:_:)
func CGPDFDictionaryGetBoolean(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetBoolean(dict, key, value)
}

// Returns the number of entries in a PDF dictionary. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionarygetcount(_:)
func CGPDFDictionaryGetCount(dict CGPDFDictionaryRef) uintptr {
	return _CGPDFDictionaryGetCount(dict)
}

// Returns whether there is another PDF dictionary associated with a specified key in a PDF dictionary and, if so, retrieves that dictionary. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionarygetdictionary(_:_:_:)
func CGPDFDictionaryGetDictionary(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetDictionary(dict, key, value)
}

// Returns whether there is a PDF integer associated with a specified key in a PDF dictionary and, if so, retrieves that integer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionarygetinteger(_:_:_:)
func CGPDFDictionaryGetInteger(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetInteger(dict, key, value)
}

// Returns whether an object with a specified key in a PDF dictionary is a PDF name reference (represented as a constant C string) and, if so, retrieves that name. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionarygetname(_:_:_:)
func CGPDFDictionaryGetName(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetName(dict, key, value)
}

// Returns whether there is a PDF number associated with a specified key in a PDF dictionary and, if so, retrieves that number. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionarygetnumber(_:_:_:)
func CGPDFDictionaryGetNumber(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetNumber(dict, key, value)
}

// Returns whether there is a PDF object associated with a specified key in a PDF dictionary and, if so, retrieves that object. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionarygetobject(_:_:_:)
func CGPDFDictionaryGetObject(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetObject(dict, key, value)
}

// Returns whether there is a PDF stream associated with a specified key in a PDF dictionary and, if so, retrieves that stream. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionarygetstream(_:_:_:)
func CGPDFDictionaryGetStream(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetStream(dict, key, value)
}

// Returns whether there is a PDF string associated with a specified key in a PDF dictionary and, if so, retrieves that string. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdictionarygetstring(_:_:_:)
func CGPDFDictionaryGetString(dict CGPDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetString(dict, key, value)
}

// CGPDFDocumentGetAccessPermissions is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/accesspermissions
func CGPDFDocumentGetAccessPermissions(document CGPDFDocumentRef) unsafe.Pointer {
	return _CGPDFDocumentGetAccessPermissions(document)
}

// Returns whether the specified PDF document allows copying. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/allowscopying
func CGPDFDocumentAllowsCopying(document CGPDFDocumentRef) bool {
	return _CGPDFDocumentAllowsCopying(document)
}

// Returns whether a PDF document allows printing. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/allowsprinting
func CGPDFDocumentAllowsPrinting(document CGPDFDocumentRef) bool {
	return _CGPDFDocumentAllowsPrinting(document)
}

// Returns the document catalog of a Core Graphics PDF document. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/catalog
func CGPDFDocumentGetCatalog(document CGPDFDocumentRef) CGPDFDictionaryRef {
	return _CGPDFDocumentGetCatalog(document)
}

// Gets the file identifier for a PDF document. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/fileidentifier
func CGPDFDocumentGetID(document CGPDFDocumentRef) CGPDFArrayRef {
	return _CGPDFDocumentGetID(document)
}

// Returns the major and minor version numbers of a Core Graphics PDF document. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/getversion(majorversion:minorversion:)
func CGPDFDocumentGetVersion(document CGPDFDocumentRef, majorVersion unsafe.Pointer, minorVersion unsafe.Pointer) {
	_CGPDFDocumentGetVersion(document, majorVersion, minorVersion)
}

// Gets the information dictionary for a PDF document. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/info
func CGPDFDocumentGetInfo(document CGPDFDocumentRef) CGPDFDictionaryRef {
	return _CGPDFDocumentGetInfo(document)
}

// Creates a Core Graphics PDF document using data specified by a URL. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/init(_:)-2gtsd
func CGPDFDocumentCreateWithURL(url unsafe.Pointer) CGPDFDocumentRef {
	return _CGPDFDocumentCreateWithURL(url)
}

// Creates a Core Graphics PDF document using a data provider. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/init(_:)-gbq6
func CGPDFDocumentCreateWithProvider(provider CGDataProviderRef) CGPDFDocumentRef {
	return _CGPDFDocumentCreateWithProvider(provider)
}

// Returns whether the specified PDF file is encrypted. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/isencrypted
func CGPDFDocumentIsEncrypted(document CGPDFDocumentRef) bool {
	return _CGPDFDocumentIsEncrypted(document)
}

// Returns whether the specified PDF document is currently unlocked. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/isunlocked
func CGPDFDocumentIsUnlocked(document CGPDFDocumentRef) bool {
	return _CGPDFDocumentIsUnlocked(document)
}

// Returns the number of pages in a PDF document. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/numberofpages
func CGPDFDocumentGetNumberOfPages(document CGPDFDocumentRef) uintptr {
	return _CGPDFDocumentGetNumberOfPages(document)
}

// CGPDFDocumentGetOutline is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/outline
func CGPDFDocumentGetOutline(document CGPDFDocumentRef) unsafe.Pointer {
	return _CGPDFDocumentGetOutline(document)
}

// Returns a page from a Core Graphics PDF document. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/page(at:)
func CGPDFDocumentGetPage(document CGPDFDocumentRef, pageNumber uintptr) CGPDFPageRef {
	return _CGPDFDocumentGetPage(document, pageNumber)
}

// Returns the type identifier for Core Graphics PDF documents. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/typeid
func CGPDFDocumentGetTypeID() unsafe.Pointer {
	return _CGPDFDocumentGetTypeID()
}

// Unlocks an encrypted PDF document when a valid password is supplied. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocument/unlockwithpassword(_:)
func CGPDFDocumentUnlockWithPassword(document CGPDFDocumentRef, password unsafe.Pointer) bool {
	return _CGPDFDocumentUnlockWithPassword(document, password)
}

// Returns the art box of a page in a PDF document. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocumentgetartbox
func CGPDFDocumentGetArtBox(document CGPDFDocumentRef, page int) CGRect {
	return _CGPDFDocumentGetArtBox(document, page)
}

// Returns the bleed box of a page in a PDF document. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocumentgetbleedbox
func CGPDFDocumentGetBleedBox(document CGPDFDocumentRef, page int) CGRect {
	return _CGPDFDocumentGetBleedBox(document, page)
}

// Returns the crop box of a page in a PDF document. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocumentgetcropbox
func CGPDFDocumentGetCropBox(document CGPDFDocumentRef, page int) CGRect {
	return _CGPDFDocumentGetCropBox(document, page)
}

// Returns the media box of a page in a PDF document. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocumentgetmediabox
func CGPDFDocumentGetMediaBox(document CGPDFDocumentRef, page int) CGRect {
	return _CGPDFDocumentGetMediaBox(document, page)
}

// Returns the rotation angle of a page in a PDF document. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocumentgetrotationangle
func CGPDFDocumentGetRotationAngle(document CGPDFDocumentRef, page int) int {
	return _CGPDFDocumentGetRotationAngle(document, page)
}

// Returns the trim box of a page in a PDF document. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocumentgettrimbox
func CGPDFDocumentGetTrimBox(document CGPDFDocumentRef, page int) CGRect {
	return _CGPDFDocumentGetTrimBox(document, page)
}

// Decrements the retain count of a PDF document. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocumentrelease
func CGPDFDocumentRelease(document CGPDFDocumentRef) {
	_CGPDFDocumentRelease(document)
}

// Increments the retain count of a Core Graphics PDF document. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdocumentretain
func CGPDFDocumentRetain(document CGPDFDocumentRef) CGPDFDocumentRef {
	return _CGPDFDocumentRetain(document)
}

// Returns the PDF type identifier of an object. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfobjectgettype(_:)
func CGPDFObjectGetType(object CGPDFObjectRef) unsafe.Pointer {
	return _CGPDFObjectGetType(object)
}

// Returns whether an object is of a given type and if it is, retrieves its value. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfobjectgetvalue(_:_:_:)
func CGPDFObjectGetValue(object CGPDFObjectRef, type_ unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFObjectGetValue(object, type_, value)
}

// Creates an empty PDF operator table. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfoperatortablecreate()
func CGPDFOperatorTableCreate() CGPDFOperatorTableRef {
	return _CGPDFOperatorTableCreate()
}

// Decrements the retain count of a CGPDFOperatorTable object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfoperatortablerelease(_:)
func CGPDFOperatorTableRelease(table CGPDFOperatorTableRef) {
	_CGPDFOperatorTableRelease(table)
}

// Increments the retain count of a CGPDFOperatorTable object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfoperatortableretain(_:)
func CGPDFOperatorTableRetain(table CGPDFOperatorTableRef) CGPDFOperatorTableRef {
	return _CGPDFOperatorTableRetain(table)
}

// Sets a callback function for a PDF operator. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfoperatortablesetcallback(_:_:_:)
func CGPDFOperatorTableSetCallback(table CGPDFOperatorTableRef, name unsafe.Pointer, callback unsafe.Pointer) {
	_CGPDFOperatorTableSetCallback(table, name, callback)
}

// Returns the dictionary of a PDF page. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfpage/dictionary
func CGPDFPageGetDictionary(page CGPDFPageRef) CGPDFDictionaryRef {
	return _CGPDFPageGetDictionary(page)
}

// Returns the document for a page. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfpage/document
func CGPDFPageGetDocument(page CGPDFPageRef) CGPDFDocumentRef {
	return _CGPDFPageGetDocument(page)
}

// Returns the rectangle that represents a type of box for a content region or page dimensions of a PDF page. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfpage/getboxrect(_:)
func CGPDFPageGetBoxRect(page CGPDFPageRef, box unsafe.Pointer) CGRect {
	return _CGPDFPageGetBoxRect(page, box)
}

// Returns the affine transform that maps a box to a given rectangle on a PDF page. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfpage/getdrawingtransform(_:rect:rotate:preserveaspectratio:)
func CGPDFPageGetDrawingTransform(page CGPDFPageRef, box unsafe.Pointer, rect CGRect, rotate int, preserveAspectRatio bool) CGAffineTransform {
	return _CGPDFPageGetDrawingTransform(page, box, rect, rotate, preserveAspectRatio)
}

// Returns the page number of the specified PDF page. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfpage/pagenumber
func CGPDFPageGetPageNumber(page CGPDFPageRef) uintptr {
	return _CGPDFPageGetPageNumber(page)
}

// Returns the rotation angle of a PDF page, in degrees. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfpage/rotationangle
func CGPDFPageGetRotationAngle(page CGPDFPageRef) int {
	return _CGPDFPageGetRotationAngle(page)
}

// Returns the CFType ID for PDF page objects. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfpage/typeid
func CGPDFPageGetTypeID() unsafe.Pointer {
	return _CGPDFPageGetTypeID()
}

// Decrements the retain count of a PDF page. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfpagerelease
func CGPDFPageRelease(page CGPDFPageRef) {
	_CGPDFPageRelease(page)
}

// Increments the retain count of a PDF page. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfpageretain
func CGPDFPageRetain(page CGPDFPageRef) CGPDFPageRef {
	return _CGPDFPageRetain(page)
}

// Creates a PDF scanner. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannercreate(_:_:_:)
func CGPDFScannerCreate(cs CGPDFContentStreamRef, table CGPDFOperatorTableRef, info unsafe.Pointer) CGPDFScannerRef {
	return _CGPDFScannerCreate(cs, table, info)
}

// Returns the content stream associated with a PDF scanner object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannergetcontentstream(_:)
func CGPDFScannerGetContentStream(scanner CGPDFScannerRef) CGPDFContentStreamRef {
	return _CGPDFScannerGetContentStream(scanner)
}

// Retrieves an array object from the scanner stack. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerpoparray(_:_:)
func CGPDFScannerPopArray(scanner CGPDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopArray(scanner, value)
}

// Retrieves a Boolean object from the scanner stack. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerpopboolean(_:_:)
func CGPDFScannerPopBoolean(scanner CGPDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopBoolean(scanner, value)
}

// Retrieves a PDF dictionary object from the scanner stack. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerpopdictionary(_:_:)
func CGPDFScannerPopDictionary(scanner CGPDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopDictionary(scanner, value)
}

// Retrieves an integer object from the scanner stack. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerpopinteger(_:_:)
func CGPDFScannerPopInteger(scanner CGPDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopInteger(scanner, value)
}

// Retrieves a character string from the scanner stack. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerpopname(_:_:)
func CGPDFScannerPopName(scanner CGPDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopName(scanner, value)
}

// Retrieves a real value object from the scanner stack. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerpopnumber(_:_:)
func CGPDFScannerPopNumber(scanner CGPDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopNumber(scanner, value)
}

// Retrieves an object from the scanner stack. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerpopobject(_:_:)
func CGPDFScannerPopObject(scanner CGPDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopObject(scanner, value)
}

// Retrieves a PDF stream object from the scanner stack. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerpopstream(_:_:)
func CGPDFScannerPopStream(scanner CGPDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopStream(scanner, value)
}

// Retrieves a string object from the scanner stack. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerpopstring(_:_:)
func CGPDFScannerPopString(scanner CGPDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopString(scanner, value)
}

// Decrements the retain count of a scanner object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerrelease(_:)
func CGPDFScannerRelease(scanner CGPDFScannerRef) {
	_CGPDFScannerRelease(scanner)
}

// Increments the retain count of a scanner object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerretain(_:)
func CGPDFScannerRetain(scanner CGPDFScannerRef) CGPDFScannerRef {
	return _CGPDFScannerRetain(scanner)
}

// Parses the content stream of a PDF scanner object. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerscan(_:)
func CGPDFScannerScan(scanner CGPDFScannerRef) bool {
	return _CGPDFScannerScan(scanner)
}

// CGPDFScannerStop is a CoreGraphics function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfscannerstop(_:)
func CGPDFScannerStop(s CGPDFScannerRef) {
	_CGPDFScannerStop(s)
}

// Returns the data associated with a PDF stream. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfstreamcopydata(_:_:)
func CGPDFStreamCopyData(stream CGPDFStreamRef, format unsafe.Pointer) unsafe.Pointer {
	return _CGPDFStreamCopyData(stream, format)
}

// Returns the dictionary associated with a PDF stream. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfstreamgetdictionary(_:)
func CGPDFStreamGetDictionary(stream CGPDFStreamRef) CGPDFDictionaryRef {
	return _CGPDFStreamGetDictionary(stream)
}

// Converts a string to a date. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfstringcopydate(_:)
func CGPDFStringCopyDate(string CGPDFStringRef) unsafe.Pointer {
	return _CGPDFStringCopyDate(string)
}

// Returns a CFString object that represents a PDF string as a text string. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfstringcopytextstring(_:)
func CGPDFStringCopyTextString(string CGPDFStringRef) unsafe.Pointer {
	return _CGPDFStringCopyTextString(string)
}

// Returns a pointer to the bytes of a PDF string. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfstringgetbyteptr(_:)
func CGPDFStringGetBytePtr(string CGPDFStringRef) unsafe.Pointer {
	return _CGPDFStringGetBytePtr(string)
}

// Returns the number of bytes in a PDF string. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfstringgetlength(_:)
func CGPDFStringGetLength(string CGPDFStringRef) uintptr {
	return _CGPDFStringGetLength(string)
}

// CGPDFTagTypeGetName is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdftagtype/name
func CGPDFTagTypeGetName(tagType unsafe.Pointer) unsafe.Pointer {
	return _CGPDFTagTypeGetName(tagType)
}

// Tells a PostScript converter to abort a conversion at the next available opportunity. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverter/abort()
func CGPSConverterAbort(converter CGPSConverterRef) bool {
	return _CGPSConverterAbort(converter)
}

// Uses a PostScript converter to convert PostScript data to PDF data. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverter/convert(_:consumer:options:)
func CGPSConverterConvert(converter CGPSConverterRef, provider CGDataProviderRef, consumer CGDataConsumerRef, options unsafe.Pointer) bool {
	return _CGPSConverterConvert(converter, provider, consumer, options)
}

// Creates a new PostScript converter. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverter/init(info:callbacks:options:)
func CGPSConverterCreate(info unsafe.Pointer, callbacks unsafe.Pointer, options unsafe.Pointer) CGPSConverterRef {
	return _CGPSConverterCreate(info, callbacks, options)
}

// Checks whether the converter is currently converting data. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverter/isconverting
func CGPSConverterIsConverting(converter CGPSConverterRef) bool {
	return _CGPSConverterIsConverting(converter)
}

// Returns the Core Foundation type identifier for PostScript converters. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpsconverter/typeid
func CGPSConverterGetTypeID() unsafe.Pointer {
	return _CGPSConverterGetTypeID()
}

// For each element in a graphics path, calls a custom applier function. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/apply(info:function:)
func CGPathApply(path CGPathRef, info unsafe.Pointer, function unsafe.Pointer) {
	_CGPathApply(path, info, function)
}

// CGPathApplyWithBlock is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/applywithblock(_:)
func CGPathApplyWithBlock(path CGPathRef, block unsafe.Pointer) {
	_CGPathApplyWithBlock(path, block)
}

// Returns the bounding box containing all points in a graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/boundingbox
func CGPathGetBoundingBox(path CGPathRef) CGRect {
	return _CGPathGetBoundingBox(path)
}

// Returns the bounding box of a graphics path. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/boundingboxofpath
func CGPathGetPathBoundingBox(path CGPathRef) CGRect {
	return _CGPathGetPathBoundingBox(path)
}

// Creates an immutable copy of a graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/copy()
func CGPathCreateCopy(path CGPathRef) CGPathRef {
	return _CGPathCreateCopy(path)
}

// Creates an immutable copy of a graphics path transformed by a transformation matrix. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/copy(using:)
func CGPathCreateCopyByTransformingPath(path CGPathRef, transform unsafe.Pointer) CGPathRef {
	return _CGPathCreateCopyByTransformingPath(path, transform)
}

// Returns the current point in a graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/currentpoint
func CGPathGetCurrentPoint(path CGPathRef) CGPoint {
	return _CGPathGetCurrentPoint(path)
}

// Create an immutable path of an ellipse. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/init(ellipsein:transform:)
func CGPathCreateWithEllipseInRect(rect CGRect, transform unsafe.Pointer) CGPathRef {
	return _CGPathCreateWithEllipseInRect(rect, transform)
}

// Create an immutable path of a rectangle. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/init(rect:transform:)
func CGPathCreateWithRect(rect CGRect, transform unsafe.Pointer) CGPathRef {
	return _CGPathCreateWithRect(rect, transform)
}

// Create an immutable path of a rounded rectangle. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/init(roundedrect:cornerwidth:cornerheight:transform:)
func CGPathCreateWithRoundedRect(rect CGRect, cornerWidth CGFloat, cornerHeight CGFloat, transform unsafe.Pointer) CGPathRef {
	return _CGPathCreateWithRoundedRect(rect, cornerWidth, cornerHeight, transform)
}

// Indicates whether or not a graphics path is empty. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/isempty
func CGPathIsEmpty(path CGPathRef) bool {
	return _CGPathIsEmpty(path)
}

// Indicates whether or not a graphics path represents a rectangle. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/isrect(_:)
func CGPathIsRect(path CGPathRef, rect unsafe.Pointer) bool {
	return _CGPathIsRect(path, rect)
}

// Creates a mutable copy of an existing graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/mutablecopy()
func CGPathCreateMutableCopy(path CGPathRef) CGMutablePathRef {
	return _CGPathCreateMutableCopy(path)
}

// Creates a mutable copy of a graphics path transformed by a transformation matrix. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/mutablecopy(using:)
func CGPathCreateMutableCopyByTransformingPath(path CGPathRef, transform unsafe.Pointer) CGMutablePathRef {
	return _CGPathCreateMutableCopyByTransformingPath(path, transform)
}

// Returns the Core Foundation type identifier for Core Graphics paths. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpath/typeid
func CGPathGetTypeID() unsafe.Pointer {
	return _CGPathGetTypeID()
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddarc
func CGPathAddArc(path CGMutablePathRef, m unsafe.Pointer, x CGFloat, y CGFloat, radius CGFloat, startAngle CGFloat, endAngle CGFloat, clockwise bool) {
	_CGPathAddArc(path, m, x, y, radius, startAngle, endAngle, clockwise)
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddarctopoint
func CGPathAddArcToPoint(path CGMutablePathRef, m unsafe.Pointer, x1 CGFloat, y1 CGFloat, x2 CGFloat, y2 CGFloat, radius CGFloat) {
	_CGPathAddArcToPoint(path, m, x1, y1, x2, y2, radius)
}

// Appends a cubic Bézier curve to a mutable graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddcurvetopoint
func CGPathAddCurveToPoint(path CGMutablePathRef, m unsafe.Pointer, cp1x CGFloat, cp1y CGFloat, cp2x CGFloat, cp2y CGFloat, x CGFloat, y CGFloat) {
	_CGPathAddCurveToPoint(path, m, cp1x, cp1y, cp2x, cp2y, x, y)
}

// Adds to a path an ellipse that fits inside a rectangle. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddellipseinrect
func CGPathAddEllipseInRect(path CGMutablePathRef, m unsafe.Pointer, rect CGRect) {
	_CGPathAddEllipseInRect(path, m, rect)
}

// Appends a line segment to a mutable graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddlinetopoint
func CGPathAddLineToPoint(path CGMutablePathRef, m unsafe.Pointer, x CGFloat, y CGFloat) {
	_CGPathAddLineToPoint(path, m, x, y)
}

// Appends an array of new line segments to a mutable graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddlines
func CGPathAddLines(path CGMutablePathRef, m unsafe.Pointer, points unsafe.Pointer, count uintptr) {
	_CGPathAddLines(path, m, points, count)
}

// Appends a path to onto a mutable graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddpath
func CGPathAddPath(path1 CGMutablePathRef, m unsafe.Pointer, path2 CGPathRef) {
	_CGPathAddPath(path1, m, path2)
}

// Appends a quadratic Bézier curve to a mutable graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddquadcurvetopoint
func CGPathAddQuadCurveToPoint(path CGMutablePathRef, m unsafe.Pointer, cpx CGFloat, cpy CGFloat, x CGFloat, y CGFloat) {
	_CGPathAddQuadCurveToPoint(path, m, cpx, cpy, x, y)
}

// Appends a rectangle to a mutable graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddrect
func CGPathAddRect(path CGMutablePathRef, m unsafe.Pointer, rect CGRect) {
	_CGPathAddRect(path, m, rect)
}

// Appends an array of rectangles to a mutable graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddrects
func CGPathAddRects(path CGMutablePathRef, m unsafe.Pointer, rects unsafe.Pointer, count uintptr) {
	_CGPathAddRects(path, m, rects, count)
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddrelativearc
func CGPathAddRelativeArc(path CGMutablePathRef, matrix unsafe.Pointer, x CGFloat, y CGFloat, radius CGFloat, startAngle CGFloat, delta CGFloat) {
	_CGPathAddRelativeArc(path, matrix, x, y, radius, startAngle, delta)
}

// Appends a rounded rectangle to a mutable graphics path. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathaddroundedrect
func CGPathAddRoundedRect(path CGMutablePathRef, transform unsafe.Pointer, rect CGRect, cornerWidth CGFloat, cornerHeight CGFloat) {
	_CGPathAddRoundedRect(path, transform, rect, cornerWidth, cornerHeight)
}

// Checks whether a point is contained in a graphics path. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcontainspoint
func CGPathContainsPoint(path CGPathRef, m unsafe.Pointer, point CGPoint, eoFill bool) bool {
	return _CGPathContainsPoint(path, m, point, eoFill)
}

// Creates a dashed copy of another path. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcreatecopybydashingpath
func CGPathCreateCopyByDashingPath(path CGPathRef, transform unsafe.Pointer, phase CGFloat, lengths unsafe.Pointer, count uintptr) CGPathRef {
	return _CGPathCreateCopyByDashingPath(path, transform, phase, lengths, count)
}

// CGPathCreateCopyByFlattening is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcreatecopybyflattening
func CGPathCreateCopyByFlattening(path CGPathRef, flatteningThreshold CGFloat) CGPathRef {
	return _CGPathCreateCopyByFlattening(path, flatteningThreshold)
}

// CGPathCreateCopyByIntersectingPath is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcreatecopybyintersectingpath
func CGPathCreateCopyByIntersectingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	return _CGPathCreateCopyByIntersectingPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateCopyByNormalizing is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcreatecopybynormalizing
func CGPathCreateCopyByNormalizing(path CGPathRef, evenOddFillRule bool) CGPathRef {
	return _CGPathCreateCopyByNormalizing(path, evenOddFillRule)
}

// Creates a stroked copy of another path. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcreatecopybystrokingpath
func CGPathCreateCopyByStrokingPath(path CGPathRef, transform unsafe.Pointer, lineWidth CGFloat, lineCap unsafe.Pointer, lineJoin unsafe.Pointer, miterLimit CGFloat) CGPathRef {
	return _CGPathCreateCopyByStrokingPath(path, transform, lineWidth, lineCap, lineJoin, miterLimit)
}

// CGPathCreateCopyBySubtractingPath is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcreatecopybysubtractingpath
func CGPathCreateCopyBySubtractingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	return _CGPathCreateCopyBySubtractingPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateCopyBySymmetricDifferenceOfPath is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcreatecopybysymmetricdifferenceofpath
func CGPathCreateCopyBySymmetricDifferenceOfPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	return _CGPathCreateCopyBySymmetricDifferenceOfPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateCopyByUnioningPath is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcreatecopybyunioningpath
func CGPathCreateCopyByUnioningPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	return _CGPathCreateCopyByUnioningPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateCopyOfLineByIntersectingPath is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcreatecopyoflinebyintersectingpath
func CGPathCreateCopyOfLineByIntersectingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	return _CGPathCreateCopyOfLineByIntersectingPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateCopyOfLineBySubtractingPath is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcreatecopyoflinebysubtractingpath
func CGPathCreateCopyOfLineBySubtractingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	return _CGPathCreateCopyOfLineBySubtractingPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateSeparateComponents is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathcreateseparatecomponents
func CGPathCreateSeparateComponents(path CGPathRef, evenOddFillRule bool) unsafe.Pointer {
	return _CGPathCreateSeparateComponents(path, evenOddFillRule)
}

// Indicates whether two graphics paths are equivalent. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathequaltopath
func CGPathEqualToPath(path1 CGPathRef, path2 CGPathRef) bool {
	return _CGPathEqualToPath(path1, path2)
}

// CGPathIntersectsPath is a CoreGraphics function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathintersectspath
func CGPathIntersectsPath(path1 CGPathRef, path2 CGPathRef, evenOddFillRule bool) bool {
	return _CGPathIntersectsPath(path1, path2, evenOddFillRule)
}

// Starts a new subpath at a specified location in a mutable graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathmovetopoint
func CGPathMoveToPoint(path CGMutablePathRef, m unsafe.Pointer, x CGFloat, y CGFloat) {
	_CGPathMoveToPoint(path, m, x, y)
}

// Decrements the retain count of a graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathrelease
func CGPathRelease(path CGPathRef) {
	_CGPathRelease(path)
}

// Increments the retain count of a graphics path. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathretain
func CGPathRetain(path CGPathRef) CGPathRef {
	return _CGPathRetain(path)
}

// Creates a pattern object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpattern/init(info:bounds:matrix:xstep:ystep:tiling:iscolored:callbacks:)
func CGPatternCreate(info unsafe.Pointer, bounds CGRect, matrix CGAffineTransform, xStep CGFloat, yStep CGFloat, tiling unsafe.Pointer, isColored bool, callbacks unsafe.Pointer) CGPatternRef {
	return _CGPatternCreate(info, bounds, matrix, xStep, yStep, tiling, isColored, callbacks)
}

// Returns the type identifier for Core Graphics patterns. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpattern/typeid
func CGPatternGetTypeID() unsafe.Pointer {
	return _CGPatternGetTypeID()
}

// Decrements the retain count of a Core Graphics pattern. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpatternrelease
func CGPatternRelease(pattern CGPatternRef) {
	_CGPatternRelease(pattern)
}

// Increments the retain count of a Core Graphics pattern. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpatternretain
func CGPatternRetain(pattern CGPatternRef) CGPatternRef {
	return _CGPatternRetain(pattern)
}

// Returns the point resulting from an affine transformation of an existing point. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpointapplyaffinetransform(_:_:)
func CGPointApplyAffineTransform(point CGPoint, t CGAffineTransform) CGPoint {
	return _CGPointApplyAffineTransform(point, t)
}

// Returns a dictionary representation of the specified point. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpointcreatedictionaryrepresentation(_:)
func CGPointCreateDictionaryRepresentation(point CGPoint) unsafe.Pointer {
	return _CGPointCreateDictionaryRepresentation(point)
}

// Returns whether two points are equal. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpointequaltopoint(_:_:)
func CGPointEqualToPoint(point1 CGPoint, point2 CGPoint) bool {
	return _CGPointEqualToPoint(point1, point2)
}

// Fills in a point using the contents of the specified dictionary. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpointmakewithdictionaryrepresentation(_:_:)
func CGPointMakeWithDictionaryRepresentation(dict unsafe.Pointer, point unsafe.Pointer) bool {
	return _CGPointMakeWithDictionaryRepresentation(dict, point)
}

// Synthesizes a low-level keyboard event on the local machine. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpostkeyboardevent(_:_:_:)
func CGPostKeyboardEvent(keyChar unsafe.Pointer, virtualKey unsafe.Pointer, keyDown unsafe.Pointer) unsafe.Pointer {
	return _CGPostKeyboardEvent(keyChar, virtualKey, keyDown)
}

// Synthesizes a low-level mouse-button event on the local machine. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpostmouseevent
func CGPostMouseEvent(mouseCursorPosition CGPoint, updateMouseCursorPosition unsafe.Pointer, buttonCount unsafe.Pointer, mouseButtonDown unsafe.Pointer) unsafe.Pointer {
	return _CGPostMouseEvent(mouseCursorPosition, updateMouseCursorPosition, buttonCount, mouseButtonDown)
}

// Synthesizes a low-level scrolling event on the local machine. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpostscrollwheelevent
func CGPostScrollWheelEvent(wheelCount unsafe.Pointer, wheel1 unsafe.Pointer) unsafe.Pointer {
	return _CGPostScrollWheelEvent(wheelCount, wheel1)
}

// CGPreflightListenEventAccess is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpreflightlisteneventaccess()
func CGPreflightListenEventAccess() bool {
	return _CGPreflightListenEventAccess()
}

// CGPreflightPostEventAccess is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpreflightposteventaccess()
func CGPreflightPostEventAccess() bool {
	return _CGPreflightPostEventAccess()
}

// CGPreflightScreenCaptureAccess is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpreflightscreencaptureaccess()
func CGPreflightScreenCaptureAccess() bool {
	return _CGPreflightScreenCaptureAccess()
}

// Applies an affine transform to a rectangle. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectapplyaffinetransform(_:_:)
func CGRectApplyAffineTransform(rect CGRect, t CGAffineTransform) CGRect {
	return _CGRectApplyAffineTransform(rect, t)
}

// Returns whether a rectangle contains a specified point. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectcontainspoint(_:_:)
func CGRectContainsPoint(rect CGRect, point CGPoint) bool {
	return _CGRectContainsPoint(rect, point)
}

// Returns whether the first rectangle contains the second rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectcontainsrect(_:_:)
func CGRectContainsRect(rect1 CGRect, rect2 CGRect) bool {
	return _CGRectContainsRect(rect1, rect2)
}

// Returns a dictionary representation of the provided rectangle. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectcreatedictionaryrepresentation(_:)
func CGRectCreateDictionaryRepresentation(p0 CGRect) unsafe.Pointer {
	return _CGRectCreateDictionaryRepresentation(p0)
}

// Divides a source rectangle into two component rectangles. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectdivide
func CGRectDivide(rect CGRect, slice unsafe.Pointer, remainder unsafe.Pointer, amount CGFloat, edge unsafe.Pointer) {
	_CGRectDivide(rect, slice, remainder, amount, edge)
}

// Returns whether two rectangles are equal in size and position. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectequaltorect(_:_:)
func CGRectEqualToRect(rect1 CGRect, rect2 CGRect) bool {
	return _CGRectEqualToRect(rect1, rect2)
}

// Returns the height of a rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectgetheight(_:)
func CGRectGetHeight(rect CGRect) CGFloat {
	return _CGRectGetHeight(rect)
}

// Returns the largest value of the x-coordinate for the rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectgetmaxx(_:)
func CGRectGetMaxX(rect CGRect) CGFloat {
	return _CGRectGetMaxX(rect)
}

// Returns the largest value for the y-coordinate of the rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectgetmaxy(_:)
func CGRectGetMaxY(rect CGRect) CGFloat {
	return _CGRectGetMaxY(rect)
}

// Returns the x- coordinate that establishes the center of a rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectgetmidx(_:)
func CGRectGetMidX(rect CGRect) CGFloat {
	return _CGRectGetMidX(rect)
}

// Returns the y-coordinate that establishes the center of the rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectgetmidy(_:)
func CGRectGetMidY(rect CGRect) CGFloat {
	return _CGRectGetMidY(rect)
}

// Returns the smallest value for the x-coordinate of the rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectgetminx(_:)
func CGRectGetMinX(rect CGRect) CGFloat {
	return _CGRectGetMinX(rect)
}

// Returns the smallest value for the y-coordinate of the rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectgetminy(_:)
func CGRectGetMinY(rect CGRect) CGFloat {
	return _CGRectGetMinY(rect)
}

// Returns the width of a rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectgetwidth(_:)
func CGRectGetWidth(rect CGRect) CGFloat {
	return _CGRectGetWidth(rect)
}

// Returns a rectangle that is smaller or larger than the source rectangle, with the same center point. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectinset(_:_:_:)
func CGRectInset(rect CGRect, dx CGFloat, dy CGFloat) CGRect {
	return _CGRectInset(rect, dx, dy)
}

// Returns the smallest rectangle that results from converting the source rectangle values to integers. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectintegral(_:)
func CGRectIntegral(rect CGRect) CGRect {
	return _CGRectIntegral(rect)
}

// Returns the intersection of two rectangles. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectintersection(_:_:)
func CGRectIntersection(r1 CGRect, r2 CGRect) CGRect {
	return _CGRectIntersection(r1, r2)
}

// Returns whether two rectangles intersect. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectintersectsrect(_:_:)
func CGRectIntersectsRect(rect1 CGRect, rect2 CGRect) bool {
	return _CGRectIntersectsRect(rect1, rect2)
}

// Returns whether a rectangle has zero width or height, or is a null rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectisempty(_:)
func CGRectIsEmpty(rect CGRect) bool {
	return _CGRectIsEmpty(rect)
}

// Returns whether a rectangle is infinite. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectisinfinite(_:)
func CGRectIsInfinite(rect CGRect) bool {
	return _CGRectIsInfinite(rect)
}

// Returns whether the rectangle is equal to the null rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectisnull(_:)
func CGRectIsNull(rect CGRect) bool {
	return _CGRectIsNull(rect)
}

// Fills in a rectangle using the contents of the specified dictionary. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectmakewithdictionaryrepresentation(_:_:)
func CGRectMakeWithDictionaryRepresentation(dict unsafe.Pointer, rect unsafe.Pointer) bool {
	return _CGRectMakeWithDictionaryRepresentation(dict, rect)
}

// Returns a rectangle with an origin that is offset from that of the source rectangle. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectoffset(_:_:_:)
func CGRectOffset(rect CGRect, dx CGFloat, dy CGFloat) CGRect {
	return _CGRectOffset(rect, dx, dy)
}

// Returns a rectangle with a positive width and height. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectstandardize(_:)
func CGRectStandardize(rect CGRect) CGRect {
	return _CGRectStandardize(rect)
}

// Returns the smallest rectangle that contains the two source rectangles. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectunion(_:_:)
func CGRectUnion(r1 CGRect, r2 CGRect) CGRect {
	return _CGRectUnion(r1, r2)
}

// Registers a callback function to be invoked when local displays are refreshed or modified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgregisterscreenrefreshcallback(_:_:)
func CGRegisterScreenRefreshCallback(callback unsafe.Pointer, userInfo unsafe.Pointer) unsafe.Pointer {
	return _CGRegisterScreenRefreshCallback(callback, userInfo)
}

// Releases all captured displays. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgreleasealldisplays()
func CGReleaseAllDisplays() unsafe.Pointer {
	return _CGReleaseAllDisplays()
}

// Releases a display fade reservation, and unfades the display if needed. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgreleasedisplayfadereservation(_:)
func CGReleaseDisplayFadeReservation(token unsafe.Pointer) unsafe.Pointer {
	return _CGReleaseDisplayFadeReservation(token)
}

// Deallocates a list of rectangles that represent changed areas on local displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgreleasescreenrefreshrects(_:)
func CGReleaseScreenRefreshRects(rects unsafe.Pointer) {
	_CGReleaseScreenRefreshRects(rects)
}

// CGRenderingBufferLockBytePtr is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrenderingbufferlockbyteptr
func CGRenderingBufferLockBytePtr(provider CGRenderingBufferProviderRef) unsafe.Pointer {
	return _CGRenderingBufferLockBytePtr(provider)
}

// CGRenderingBufferProviderCreate is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrenderingbufferprovidercreate
func CGRenderingBufferProviderCreate(info unsafe.Pointer, size uintptr) CGRenderingBufferProviderRef {
	return _CGRenderingBufferProviderCreate(info, size)
}

// CGRenderingBufferProviderCreateWithCFData is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrenderingbufferprovidercreatewithcfdata
func CGRenderingBufferProviderCreateWithCFData(data unsafe.Pointer) CGRenderingBufferProviderRef {
	return _CGRenderingBufferProviderCreateWithCFData(data)
}

// CGRenderingBufferProviderGetSize is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrenderingbufferprovidergetsize
func CGRenderingBufferProviderGetSize(provider CGRenderingBufferProviderRef) uintptr {
	return _CGRenderingBufferProviderGetSize(provider)
}

// CGRenderingBufferProviderGetTypeID is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrenderingbufferprovidergettypeid
func CGRenderingBufferProviderGetTypeID() unsafe.Pointer {
	return _CGRenderingBufferProviderGetTypeID()
}

// CGRenderingBufferUnlockBytePtr is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrenderingbufferunlockbyteptr
func CGRenderingBufferUnlockBytePtr(provider CGRenderingBufferProviderRef) {
	_CGRenderingBufferUnlockBytePtr(provider)
}

// CGRequestListenEventAccess is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrequestlisteneventaccess()
func CGRequestListenEventAccess() bool {
	return _CGRequestListenEventAccess()
}

// CGRequestPostEventAccess is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrequestposteventaccess()
func CGRequestPostEventAccess() bool {
	return _CGRequestPostEventAccess()
}

// CGRequestScreenCaptureAccess is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrequestscreencaptureaccess()
func CGRequestScreenCaptureAccess() bool {
	return _CGRequestScreenCaptureAccess()
}

// Restores the permanent display configuration settings for the current user. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrestorepermanentdisplayconfiguration()
func CGRestorePermanentDisplayConfiguration() {
	_CGRestorePermanentDisplayConfiguration()
}

// Registers a callback function to be invoked when an area of the display is moved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgscreenregistermovecallback(_:_:)
func CGScreenRegisterMoveCallback(callback unsafe.Pointer, userInfo unsafe.Pointer) unsafe.Pointer {
	return _CGScreenRegisterMoveCallback(callback, userInfo)
}

// Removes a previously registered callback function invoked when an area of the display is moved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgscreenunregistermovecallback(_:_:)
func CGScreenUnregisterMoveCallback(callback unsafe.Pointer, userInfo unsafe.Pointer) {
	_CGScreenUnregisterMoveCallback(callback, userInfo)
}

// Returns information about the caller’s window server session. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgsessioncopycurrentdictionary()
func CGSessionCopyCurrentDictionary() unsafe.Pointer {
	return _CGSessionCopyCurrentDictionary()
}

// Sets the byte values in the 8-bit RGB gamma tables for a display. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgsetdisplaytransferbybytetable(_:_:_:_:_:)
func CGSetDisplayTransferByByteTable(display unsafe.Pointer, tableSize uint32, redTable unsafe.Pointer, greenTable unsafe.Pointer, blueTable unsafe.Pointer) unsafe.Pointer {
	return _CGSetDisplayTransferByByteTable(display, tableSize, redTable, greenTable, blueTable)
}

// Sets the gamma function for a display by specifying the coefficients of the gamma transfer formula. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgsetdisplaytransferbyformula(_:_:_:_:_:_:_:_:_:_:)
func CGSetDisplayTransferByFormula(display unsafe.Pointer, redMin unsafe.Pointer, redMax unsafe.Pointer, redGamma unsafe.Pointer, greenMin unsafe.Pointer, greenMax unsafe.Pointer, greenGamma unsafe.Pointer, blueMin unsafe.Pointer, blueMax unsafe.Pointer, blueGamma unsafe.Pointer) unsafe.Pointer {
	return _CGSetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
}

// Sets the color gamma function for a display by specifying the values in the RGB gamma tables. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgsetdisplaytransferbytable(_:_:_:_:_:)
func CGSetDisplayTransferByTable(display unsafe.Pointer, tableSize uint32, redTable unsafe.Pointer, greenTable unsafe.Pointer, blueTable unsafe.Pointer) unsafe.Pointer {
	return _CGSetDisplayTransferByTable(display, tableSize, redTable, greenTable, blueTable)
}

// Filters local hardware events from the keyboard and mouse during the short interval after a synthetic event is posted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgsetlocaleventsfilterduringsuppressionstate(_:_:)
func CGSetLocalEventsFilterDuringSuppressionState(filter unsafe.Pointer, state unsafe.Pointer) unsafe.Pointer {
	return _CGSetLocalEventsFilterDuringSuppressionState(filter, state)
}

// Sets the time interval in seconds that local hardware events are suppressed after posting a synthetic event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgsetlocaleventssuppressioninterval(_:)
func CGSetLocalEventsSuppressionInterval(seconds unsafe.Pointer) unsafe.Pointer {
	return _CGSetLocalEventsSuppressionInterval(seconds)
}

// CGShadingGetContentHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgshading/contentheadroom
func CGShadingGetContentHeadroom(shading CGShadingRef) float32 {
	return _CGShadingGetContentHeadroom(shading)
}

// CGShadingCreateAxialWithContentHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgshading/init(axialheadroom:space:start:end:function:extendstart:extendend:)
func CGShadingCreateAxialWithContentHeadroom(headroom float32, space CGColorSpaceRef, start CGPoint, end CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	return _CGShadingCreateAxialWithContentHeadroom(headroom, space, start, end, function, extendStart, extendEnd)
}

// Creates a shading object to use for axial shading. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgshading/init(axialspace:start:end:function:extendstart:extendend:)
func CGShadingCreateAxial(space CGColorSpaceRef, start CGPoint, end CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	return _CGShadingCreateAxial(space, start, end, function, extendStart, extendEnd)
}

// CGShadingCreateRadialWithContentHeadroom is a CoreGraphics function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgshading/init(radialheadroom:space:start:startradius:end:endradius:function:extendstart:extendend:)
func CGShadingCreateRadialWithContentHeadroom(headroom float32, space CGColorSpaceRef, start CGPoint, startRadius CGFloat, end CGPoint, endRadius CGFloat, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	return _CGShadingCreateRadialWithContentHeadroom(headroom, space, start, startRadius, end, endRadius, function, extendStart, extendEnd)
}

// Creates a shading object to use for radial shading. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgshading/init(radialspace:start:startradius:end:endradius:function:extendstart:extendend:)
func CGShadingCreateRadial(space CGColorSpaceRef, start CGPoint, startRadius CGFloat, end CGPoint, endRadius CGFloat, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	return _CGShadingCreateRadial(space, start, startRadius, end, endRadius, function, extendStart, extendEnd)
}

// Returns the Core Foundation type identifier for Core Graphics shading objects. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgshading/typeid
func CGShadingGetTypeID() unsafe.Pointer {
	return _CGShadingGetTypeID()
}

// Decrements the retain count of a shading object. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgshadingrelease
func CGShadingRelease(shading CGShadingRef) {
	_CGShadingRelease(shading)
}

// Increments the retain count of a shading object. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgshadingretain
func CGShadingRetain(shading CGShadingRef) CGShadingRef {
	return _CGShadingRetain(shading)
}

// Returns the window ID of the shield window for a captured display. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgshieldingwindowid(_:)
func CGShieldingWindowID(display unsafe.Pointer) unsafe.Pointer {
	return _CGShieldingWindowID(display)
}

// Returns the window level of the shield window for a captured display. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgshieldingwindowlevel()
func CGShieldingWindowLevel() unsafe.Pointer {
	return _CGShieldingWindowLevel()
}

// Returns the height and width resulting from a transformation of an existing height and width. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgsizeapplyaffinetransform(_:_:)
func CGSizeApplyAffineTransform(size CGSize, t CGAffineTransform) CGSize {
	return _CGSizeApplyAffineTransform(size, t)
}

// Returns a dictionary representation of the specified size. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgsizecreatedictionaryrepresentation(_:)
func CGSizeCreateDictionaryRepresentation(size CGSize) unsafe.Pointer {
	return _CGSizeCreateDictionaryRepresentation(size)
}

// Returns whether two sizes are equal. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgsizeequaltosize(_:_:)
func CGSizeEqualToSize(size1 CGSize, size2 CGSize) bool {
	return _CGSizeEqualToSize(size1, size2)
}

// Fills in a size using the contents of the specified dictionary. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgsizemakewithdictionaryrepresentation(_:_:)
func CGSizeMakeWithDictionaryRepresentation(dict unsafe.Pointer, size unsafe.Pointer) bool {
	return _CGSizeMakeWithDictionaryRepresentation(dict, size)
}

// Removes a previously registered callback function invoked when local displays are refreshed or modified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgunregisterscreenrefreshcallback(_:_:)
func CGUnregisterScreenRefreshCallback(callback unsafe.Pointer, userInfo unsafe.Pointer) {
	_CGUnregisterScreenRefreshCallback(callback, userInfo)
}

// Waits for screen refresh operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwaitforscreenrefreshrects(_:_:)
func CGWaitForScreenRefreshRects(rects unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	return _CGWaitForScreenRefreshRects(rects, count)
}

// Waits for screen update operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwaitforscreenupdaterects(_:_:_:_:_:)
func CGWaitForScreenUpdateRects(requestedOperations unsafe.Pointer, currentOperation unsafe.Pointer, rects unsafe.Pointer, rectCount unsafe.Pointer, delta unsafe.Pointer) unsafe.Pointer {
	return _CGWaitForScreenUpdateRects(requestedOperations, currentOperation, rects, rectCount, delta)
}

// Moves the mouse cursor without generating events. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwarpmousecursorposition(_:)
func CGWarpMouseCursorPosition(newCursorPosition CGPoint) unsafe.Pointer {
	return _CGWarpMouseCursorPosition(newCursorPosition)
}

// Returns the window level that corresponds to one of the standard window types. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowlevelforkey(_:)
func CGWindowLevelForKey(key unsafe.Pointer) unsafe.Pointer {
	return _CGWindowLevelForKey(key)
}

// Generates and returns information about the selected windows in the current user session. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowlistcopywindowinfo(_:_:)
func CGWindowListCopyWindowInfo(option unsafe.Pointer, relativeToWindow unsafe.Pointer) unsafe.Pointer {
	return _CGWindowListCopyWindowInfo(option, relativeToWindow)
}

// Returns the list of window IDs associated with the specified windows in the current user session. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowlistcreate
func CGWindowListCreate(option unsafe.Pointer, relativeToWindow unsafe.Pointer) unsafe.Pointer {
	return _CGWindowListCreate(option, relativeToWindow)
}

// Generates and returns information about windows with the specified window IDs. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowlistcreatedescriptionfromarray(_:)
func CGWindowListCreateDescriptionFromArray(windowArray unsafe.Pointer) unsafe.Pointer {
	return _CGWindowListCreateDescriptionFromArray(windowArray)
}

// Returns a composite image based on a dynamically generated list of windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowlistcreateimage(_:_:_:_:)
func CGWindowListCreateImage(screenBounds CGRect, listOption unsafe.Pointer, windowID unsafe.Pointer, imageOption unsafe.Pointer) CGImageRef {
	return _CGWindowListCreateImage(screenBounds, listOption, windowID, imageOption)
}

// Returns a Core Foundation Mach port (CFMachPort) that corresponds to the macOS window server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowservercfmachport()
func CGWindowServerCFMachPort() unsafe.Pointer {
	return _CGWindowServerCFMachPort()
}

// CGWindowServerCreateServerPort is a CoreGraphics function. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowservercreateserverport()
func CGWindowServerCreateServerPort() unsafe.Pointer {
	return _CGWindowServerCreateServerPort()
}
