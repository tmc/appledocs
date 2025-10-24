// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import (
	"unsafe"

	"github.com/ebitengine/purego"
	corefoundation "github.com/tmc/appledocs/generated/corefoundation"
)


// CoreGraphics Functions (727 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CGAcquireDisplayFadeReservation func(DisplayReservationInterval, unsafe.Pointer) Error
	_CGAffineTransformConcat func(AffineTransform, AffineTransform) AffineTransform
	_CGAffineTransformDecompose func(AffineTransform) corefoundation.AffineTransformComponents
	_CGAffineTransformEqualToTransform func(AffineTransform, AffineTransform) bool
	_CGAffineTransformInvert func(AffineTransform) AffineTransform
	_CGAffineTransformIsIdentity func(AffineTransform) bool
	_CGAffineTransformMake func(Float, Float, Float, Float, Float, Float) AffineTransform
	_CGAffineTransformMakeRotation func(Float) AffineTransform
	_CGAffineTransformMakeScale func(Float, Float) AffineTransform
	_CGAffineTransformMakeTranslation func(Float, Float) AffineTransform
	_CGAffineTransformMakeWithComponents func(corefoundation.AffineTransformComponents) AffineTransform
	_CGAffineTransformRotate func(AffineTransform, Float) AffineTransform
	_CGAffineTransformScale func(AffineTransform, Float, Float) AffineTransform
	_CGAffineTransformTranslate func(AffineTransform, Float, Float) AffineTransform
	_CGAssociateMouseAndMouseCursorPosition func(unsafe.Pointer) Error
	_CGBeginDisplayConfiguration func(unsafe.Pointer) Error
	_CGBitmapContextCreateAdaptive func(uintptr, uintptr, DictionaryRef, bool) ContextRef
	_CGCancelDisplayConfiguration func(DisplayConfigRef) Error
	_CGCaptureAllDisplays func() Error
	_CGCaptureAllDisplaysWithOptions func(CaptureOptions) Error
	_CGColorGetAlpha func(ColorRef) Float
	_CGColorGetColorSpace func(ColorRef) ColorSpaceRef
	_CGColorGetContentHeadroom func(ColorRef) float32
	_CGColorCreateCopyByMatchingToColorSpace func(ColorSpaceRef, ColorRenderingIntent, ColorRef, DictionaryRef) ColorRef
	_CGColorCreateCopy func(ColorRef) ColorRef
	_CGColorCreateCopyWithAlpha func(ColorRef, Float) ColorRef
	_CGColorCreate func(ColorSpaceRef, []float64) ColorRef
	_CGColorCreateGenericCMYK func(Float, Float, Float, Float, Float) ColorRef
	_CGColorCreateGenericGrayGamma2_2 func(Float, Float) ColorRef
	_CGColorCreateGenericGray func(Float, Float) ColorRef
	_CGColorCreateWithContentHeadroom func(float32, ColorSpaceRef, Float, Float, Float, Float) ColorRef
	_CGColorCreateWithPattern func(ColorSpaceRef, PatternRef, []float64) ColorRef
	_CGColorCreateGenericRGB func(Float, Float, Float, Float) ColorRef
	_CGColorCreateSRGB func(Float, Float, Float, Float) ColorRef
	_CGColorGetNumberOfComponents func(ColorRef) uintptr
	_CGColorGetPattern func(ColorRef) PatternRef
	_CGColorGetTypeID func() TypeID
	_CGColorConversionInfoConvertData func(ColorConversionInfoRef, uintptr, uintptr, unsafe.Pointer, CGColorBufferFormat, unsafe.Pointer, CGColorBufferFormat, DictionaryRef) bool
	_CGColorConversionInfoCreateWithOptions func(ColorSpaceRef, ColorSpaceRef, DictionaryRef) ColorConversionInfoRef
	_CGColorConversionInfoCreate func(ColorSpaceRef, ColorSpaceRef) ColorConversionInfoRef
	_CGColorConversionInfoCreateForToneMapping func(ColorSpaceRef, float32, ColorSpaceRef, float32, ToneMapping, DictionaryRef, unsafe.Pointer) ColorConversionInfoRef
	_CGColorConversionInfoGetTypeID func() TypeID
	_CGColorConversionInfoCreateFromList func(DictionaryRef, ColorSpaceRef, ColorConversionInfoTransformType, ColorRenderingIntent) ColorConversionInfoRef
	_CGColorConversionInfoCreateFromListWithArguments func(DictionaryRef, ColorSpaceRef, ColorConversionInfoTransformType, ColorRenderingIntent, unsafe.Pointer) ColorConversionInfoRef
	_CGColorEqualToColor func(ColorRef, ColorRef) bool
	_CGColorGetComponents func(ColorRef) []float64
	_CGColorGetConstantColor func(StringRef) ColorRef
	_CGColorRelease func(ColorRef)
	_CGColorRetain func(ColorRef) ColorRef
	_CGColorSpaceGetBaseColorSpace func(ColorSpaceRef) ColorSpaceRef
	_CGColorSpaceCopyICCData func(ColorSpaceRef) DataRef
	_CGColorSpaceCopyPropertyList func(ColorSpaceRef) PropertyListRef
	_CGColorSpaceCopyICCProfile func(ColorSpaceRef) DataRef
	_CGColorSpaceCreateCalibratedGray func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) ColorSpaceRef
	_CGColorSpaceCreateCalibratedRGB func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) ColorSpaceRef
	_CGColorSpaceCreateICCBased func(uintptr, []float64, DataProviderRef, ColorSpaceRef) ColorSpaceRef
	_CGColorSpaceCreateWithICCData func(TypeRef) ColorSpaceRef
	_CGColorSpaceCreateWithICCProfile func(DataRef) ColorSpaceRef
	_CGColorSpaceCreateIndexed func(ColorSpaceRef, uintptr, unsafe.Pointer) ColorSpaceRef
	_CGColorSpaceCreateLab func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) ColorSpaceRef
	_CGColorSpaceCreateWithName func(StringRef) ColorSpaceRef
	_CGColorSpaceCreatePattern func(ColorSpaceRef) ColorSpaceRef
	_CGColorSpaceCreateWithPlatformColorSpace func(unsafe.Pointer) ColorSpaceRef
	_CGColorSpaceCreateWithPropertyList func(PropertyListRef) ColorSpaceRef
	_CGColorSpaceIsHDR func(ColorSpaceRef) bool
	_CGColorSpaceIsWideGamutRGB func(ColorSpaceRef) bool
	_CGColorSpaceGetModel func(ColorSpaceRef) ColorSpaceModel
	_CGColorSpaceCopyName func(ColorSpaceRef) StringRef
	_CGColorSpaceGetNumberOfComponents func(ColorSpaceRef) uintptr
	_CGColorSpaceSupportsOutput func(ColorSpaceRef) bool
	_CGColorSpaceGetTypeID func() TypeID
	_CGColorSpaceCopyBaseColorSpace func(ColorSpaceRef) ColorSpaceRef
	_CGColorSpaceCreateCopyWithStandardRange func(ColorSpaceRef) ColorSpaceRef
	_CGColorSpaceCreateDeviceCMYK func() ColorSpaceRef
	_CGColorSpaceCreateDeviceGray func() ColorSpaceRef
	_CGColorSpaceCreateDeviceRGB func() ColorSpaceRef
	_CGColorSpaceCreateExtended func(ColorSpaceRef) ColorSpaceRef
	_CGColorSpaceCreateExtendedLinearized func(ColorSpaceRef) ColorSpaceRef
	_CGColorSpaceCreateLinearized func(ColorSpaceRef) ColorSpaceRef
	_CGColorSpaceCreateWithColorSyncProfile func(ColorSyncProfileRef, DictionaryRef) ColorSpaceRef
	_CGColorSpaceGetColorTable func(ColorSpaceRef, unsafe.Pointer)
	_CGColorSpaceGetColorTableCount func(ColorSpaceRef) uintptr
	_CGColorSpaceGetName func(ColorSpaceRef) StringRef
	_CGColorSpaceIsHLGBased func(ColorSpaceRef) bool
	_CGColorSpaceIsPQBased func(ColorSpaceRef) bool
	_CGColorSpaceRelease func(ColorSpaceRef)
	_CGColorSpaceRetain func(ColorSpaceRef) ColorSpaceRef
	_CGColorSpaceUsesExtendedRange func(ColorSpaceRef) bool
	_CGColorSpaceUsesITUR_2100TF func(ColorSpaceRef) bool
	_CGCompleteDisplayConfiguration func(DisplayConfigRef, ConfigureOption) Error
	_CGConfigureDisplayFadeEffect func(DisplayConfigRef, DisplayFadeInterval, DisplayFadeInterval, float32, float32, float32) Error
	_CGConfigureDisplayMirrorOfDisplay func(DisplayConfigRef, DirectDisplayID, DirectDisplayID) Error
	_CGConfigureDisplayMode func(DisplayConfigRef, DirectDisplayID, DictionaryRef) Error
	_CGConfigureDisplayOrigin func(DisplayConfigRef, DirectDisplayID, int32, int32) Error
	_CGConfigureDisplayStereoOperation func(DisplayConfigRef, DirectDisplayID, unsafe.Pointer, unsafe.Pointer) Error
	_CGConfigureDisplayWithDisplayMode func(DisplayConfigRef, DirectDisplayID, DisplayModeRef, DictionaryRef) Error
	_CGPDFContextAddDestinationAtPoint func(ContextRef, StringRef, Point)
	_CGPDFContextAddDocumentMetadata func(ContextRef, DataRef)
	_CGContextAddEllipseInRect func(ContextRef, Rect)
	_CGContextAddPath func(ContextRef, PathRef)
	_CGContextAddRect func(ContextRef, Rect)
	_CGBitmapContextGetAlphaInfo func(ContextRef) ImageAlphaInfo
	_CGPDFContextBeginPage func(ContextRef, DictionaryRef)
	_CGContextBeginPage func(ContextRef, unsafe.Pointer)
	_CGContextBeginPath func(ContextRef)
	_CGContextBeginTransparencyLayer func(ContextRef, DictionaryRef)
	_CGContextBeginTransparencyLayerWithRect func(ContextRef, Rect, DictionaryRef)
	_CGBitmapContextGetBitmapInfo func(ContextRef) BitmapInfo
	_CGBitmapContextGetBitsPerComponent func(ContextRef) uintptr
	_CGBitmapContextGetBitsPerPixel func(ContextRef) uintptr
	_CGContextGetClipBoundingBox func(ContextRef) Rect
	_CGContextGetPathBoundingBox func(ContextRef) Rect
	_CGBitmapContextGetBytesPerRow func(ContextRef) uintptr
	_CGContextClearRect func(ContextRef, Rect)
	_CGContextClipToRect func(ContextRef, Rect)
	_CGContextClipToMask func(ContextRef, Rect, ImageRef)
	_CGPDFContextClose func(ContextRef)
	_CGContextClosePath func(ContextRef)
	_CGBitmapContextGetColorSpace func(ContextRef) ColorSpaceRef
	_CGContextConcatCTM func(ContextRef, AffineTransform)
	_CGContextConvertSizeToDeviceSpace func(ContextRef, Size) Size
	_CGContextConvertPointToDeviceSpace func(ContextRef, Point) Point
	_CGContextConvertRectToDeviceSpace func(ContextRef, Rect) Rect
	_CGContextConvertRectToUserSpace func(ContextRef, Rect) Rect
	_CGContextConvertPointToUserSpace func(ContextRef, Point) Point
	_CGContextConvertSizeToUserSpace func(ContextRef, Size) Size
	_CGContextGetCTM func(ContextRef) AffineTransform
	_CGContextGetPathCurrentPoint func(ContextRef) Point
	_CGBitmapContextGetData func(ContextRef) unsafe.Pointer
	_CGContextDrawLinearGradient func(ContextRef, GradientRef, Point, Point, GradientDrawingOptions)
	_CGContextDrawPDFPage func(ContextRef, PDFPageRef)
	_CGContextDrawPath func(ContextRef, PathDrawingMode)
	_CGContextDrawRadialGradient func(ContextRef, GradientRef, Point, Float, Point, Float, GradientDrawingOptions)
	_CGContextDrawShading func(ContextRef, ShadingRef)
	_CGPDFContextEndPage func(ContextRef)
	_CGContextEndPage func(ContextRef)
	_CGContextEndTransparencyLayer func(ContextRef)
	_CGContextFillRect func(ContextRef, Rect)
	_CGContextFillEllipseInRect func(ContextRef, Rect)
	_CGContextFlush func(ContextRef)
	_CGBitmapContextGetHeight func(ContextRef) uintptr
	_CGPDFContextCreateWithURL func(URLRef, unsafe.Pointer, DictionaryRef) ContextRef
	_CGPDFContextCreate func(DataConsumerRef, unsafe.Pointer, DictionaryRef) ContextRef
	_CGBitmapContextCreate func(unsafe.Pointer, uintptr, uintptr, uintptr, uintptr, ColorSpaceRef, BitmapInfo) ContextRef
	_CGBitmapContextCreateWithData func(unsafe.Pointer, uintptr, uintptr, uintptr, uintptr, ColorSpaceRef, BitmapInfo, BitmapContextReleaseDataCallback, unsafe.Pointer) ContextRef
	_CGContextGetInterpolationQuality func(ContextRef) InterpolationQuality
	_CGContextIsPathEmpty func(ContextRef) bool
	_CGBitmapContextCreateImage func(ContextRef) ImageRef
	_CGContextCopyPath func(ContextRef) PathRef
	_CGContextPathContainsPoint func(ContextRef, Point, PathDrawingMode) bool
	_CGContextReplacePathWithStrokedPath func(ContextRef)
	_CGContextResetClip func(ContextRef)
	_CGContextRestoreGState func(ContextRef)
	_CGContextRotateCTM func(ContextRef, Float)
	_CGContextSaveGState func(ContextRef)
	_CGContextScaleCTM func(ContextRef, Float, Float)
	_CGContextSelectFont func(ContextRef, unsafe.Pointer, Float, TextEncoding)
	_CGContextSetAllowsAntialiasing func(ContextRef, bool)
	_CGContextSetAllowsFontSmoothing func(ContextRef, bool)
	_CGContextSetAllowsFontSubpixelPositioning func(ContextRef, bool)
	_CGContextSetAllowsFontSubpixelQuantization func(ContextRef, bool)
	_CGContextSetAlpha func(ContextRef, Float)
	_CGContextSetBlendMode func(ContextRef, BlendMode)
	_CGContextSetCharacterSpacing func(ContextRef, Float)
	_CGPDFContextSetDestinationForRect func(ContextRef, StringRef, Rect)
	_CGContextSetEDRTargetHeadroom func(ContextRef, float32) bool
	_CGContextSetFillColor func(ContextRef, []float64)
	_CGContextSetFillColorWithColor func(ContextRef, ColorRef)
	_CGContextSetCMYKFillColor func(ContextRef, Float, Float, Float, Float, Float)
	_CGContextSetGrayFillColor func(ContextRef, Float, Float)
	_CGContextSetRGBFillColor func(ContextRef, Float, Float, Float, Float)
	_CGContextSetFillColorSpace func(ContextRef, ColorSpaceRef)
	_CGContextSetFillPattern func(ContextRef, PatternRef, []float64)
	_CGContextSetFlatness func(ContextRef, Float)
	_CGContextSetFont func(ContextRef, FontRef)
	_CGContextSetFontSize func(ContextRef, Float)
	_CGContextSetLineCap func(ContextRef, LineCap)
	_CGContextSetLineJoin func(ContextRef, LineJoin)
	_CGContextSetLineWidth func(ContextRef, Float)
	_CGContextSetMiterLimit func(ContextRef, Float)
	_CGContextSetPatternPhase func(ContextRef, Size)
	_CGContextSetRenderingIntent func(ContextRef, ColorRenderingIntent)
	_CGContextSetShadow func(ContextRef, Size, Float)
	_CGContextSetShadowWithColor func(ContextRef, Size, Float, ColorRef)
	_CGContextSetShouldAntialias func(ContextRef, bool)
	_CGContextSetShouldSmoothFonts func(ContextRef, bool)
	_CGContextSetShouldSubpixelPositionFonts func(ContextRef, bool)
	_CGContextSetShouldSubpixelQuantizeFonts func(ContextRef, bool)
	_CGContextSetStrokeColorWithColor func(ContextRef, ColorRef)
	_CGContextSetStrokeColor func(ContextRef, []float64)
	_CGContextSetCMYKStrokeColor func(ContextRef, Float, Float, Float, Float, Float)
	_CGContextSetGrayStrokeColor func(ContextRef, Float, Float)
	_CGContextSetRGBStrokeColor func(ContextRef, Float, Float, Float, Float)
	_CGContextSetStrokeColorSpace func(ContextRef, ColorSpaceRef)
	_CGContextSetStrokePattern func(ContextRef, PatternRef, []float64)
	_CGContextSetTextDrawingMode func(ContextRef, TextDrawingMode)
	_CGPDFContextSetURLForRect func(ContextRef, URLRef, Rect)
	_CGContextShowGlyphs func(ContextRef, unsafe.Pointer, uintptr)
	_CGContextShowGlyphsAtPoint func(ContextRef, Float, Float, unsafe.Pointer, uintptr)
	_CGContextShowGlyphsWithAdvances func(ContextRef, unsafe.Pointer, unsafe.Pointer, uintptr)
	_CGContextShowText func(ContextRef, unsafe.Pointer, uintptr)
	_CGContextShowTextAtPoint func(ContextRef, Float, Float, unsafe.Pointer, uintptr)
	_CGContextStrokeRect func(ContextRef, Rect)
	_CGContextStrokeRectWithWidth func(ContextRef, Rect, Float)
	_CGContextStrokeEllipseInRect func(ContextRef, Rect)
	_CGContextStrokePath func(ContextRef)
	_CGContextSynchronize func(ContextRef)
	_CGContextSynchronizeAttributes func(ContextRef)
	_CGContextGetTextMatrix func(ContextRef) AffineTransform
	_CGContextTranslateCTM func(ContextRef, Float, Float)
	_CGContextGetTypeID func() TypeID
	_CGContextGetUserSpaceToDeviceSpaceTransform func(ContextRef) AffineTransform
	_CGBitmapContextGetWidth func(ContextRef) uintptr
	_CGContextAddArc func(ContextRef, Float, Float, Float, Float, Float, int)
	_CGContextAddArcToPoint func(ContextRef, Float, Float, Float, Float, Float)
	_CGContextAddCurveToPoint func(ContextRef, Float, Float, Float, Float, Float, Float)
	_CGContextAddLineToPoint func(ContextRef, Float, Float)
	_CGContextAddLines func(ContextRef, unsafe.Pointer, uintptr)
	_CGContextAddQuadCurveToPoint func(ContextRef, Float, Float, Float, Float)
	_CGContextAddRects func(ContextRef, unsafe.Pointer, uintptr)
	_CGContextClip func(ContextRef)
	_CGContextClipToRects func(ContextRef, unsafe.Pointer, uintptr)
	_CGContextDrawConicGradient func(ContextRef, GradientRef, Point, Float)
	_CGContextDrawImage func(ContextRef, Rect, ImageRef)
	_CGContextDrawImageApplyingToneMapping func(ContextRef, Rect, ImageRef, ToneMapping, DictionaryRef) bool
	_CGContextDrawLayerAtPoint func(ContextRef, Point, LayerRef)
	_CGContextDrawLayerInRect func(ContextRef, Rect, LayerRef)
	_CGContextDrawPDFDocument func(ContextRef, Rect, PDFDocumentRef, int)
	_CGContextDrawTiledImage func(ContextRef, Rect, ImageRef)
	_CGContextEOClip func(ContextRef)
	_CGContextEOFillPath func(ContextRef)
	_CGContextFillPath func(ContextRef)
	_CGContextFillRects func(ContextRef, unsafe.Pointer, uintptr)
	_CGContextGetContentToneMappingInfo func(ContextRef) CGContentToneMappingInfo
	_CGContextGetEDRTargetHeadroom func(ContextRef) float32
	_CGContextGetTextPosition func(ContextRef) Point
	_CGContextMoveToPoint func(ContextRef, Float, Float)
	_CGContextRelease func(ContextRef)
	_CGContextRetain func(ContextRef) ContextRef
	_CGContextSetContentToneMappingInfo func(ContextRef, CGContentToneMappingInfo)
	_CGContextSetInterpolationQuality func(ContextRef, InterpolationQuality)
	_CGContextSetLineDash func(ContextRef, Float, []float64, uintptr)
	_CGContextSetTextMatrix func(ContextRef, AffineTransform)
	_CGContextSetTextPosition func(ContextRef, Float, Float)
	_CGContextShowGlyphsAtPositions func(ContextRef, unsafe.Pointer, unsafe.Pointer, uintptr)
	_CGContextStrokeLineSegments func(ContextRef, unsafe.Pointer, uintptr)
	_CGConvertColorDataWithFormat func(uintptr, uintptr, unsafe.Pointer, CGColorDataFormat, unsafe.Pointer, CGColorDataFormat, DictionaryRef) bool
	_CGCursorIsDrawnInFramebuffer func() unsafe.Pointer
	_CGCursorIsVisible func() unsafe.Pointer
	_CGDataConsumerCreateWithCFData func(MutableDataRef) DataConsumerRef
	_CGDataConsumerCreate func(unsafe.Pointer, unsafe.Pointer) DataConsumerRef
	_CGDataConsumerCreateWithURL func(URLRef) DataConsumerRef
	_CGDataConsumerGetTypeID func() TypeID
	_CGDataConsumerRelease func(DataConsumerRef)
	_CGDataConsumerRetain func(DataConsumerRef) DataConsumerRef
	_CGDataProviderCopyData func(DataProviderRef) DataRef
	_CGDataProviderGetInfo func(DataProviderRef) unsafe.Pointer
	_CGDataProviderCreateWithCFData func(DataRef) DataProviderRef
	_CGDataProviderCreateWithData func(unsafe.Pointer, unsafe.Pointer, uintptr, DataProviderReleaseDataCallback) DataProviderRef
	_CGDataProviderCreateDirect func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) DataProviderRef
	_CGDataProviderCreateWithFilename func(unsafe.Pointer) DataProviderRef
	_CGDataProviderCreateSequential func(unsafe.Pointer, unsafe.Pointer) DataProviderRef
	_CGDataProviderCreateWithURL func(URLRef) DataProviderRef
	_CGDataProviderGetTypeID func() TypeID
	_CGDataProviderRelease func(DataProviderRef)
	_CGDataProviderRetain func(DataProviderRef) DataProviderRef
	_CGDirectDisplayCopyCurrentMetalDevice func(DirectDisplayID) unsafe.Pointer
	_CGDisplayAvailableModes func(DirectDisplayID) ArrayRef
	_CGDisplayBestModeForParameters func(DirectDisplayID, uintptr, uintptr, uintptr, unsafe.Pointer) DictionaryRef
	_CGDisplayBestModeForParametersAndRefreshRate func(DirectDisplayID, uintptr, uintptr, uintptr, RefreshRate, unsafe.Pointer) DictionaryRef
	_CGDisplayBounds func(DirectDisplayID) Rect
	_CGDisplayCapture func(DirectDisplayID) Error
	_CGDisplayCaptureWithOptions func(DirectDisplayID, CaptureOptions) Error
	_CGDisplayCopyAllDisplayModes func(DirectDisplayID, DictionaryRef) ArrayRef
	_CGDisplayCopyColorSpace func(DirectDisplayID) ColorSpaceRef
	_CGDisplayCopyDisplayMode func(DirectDisplayID) DisplayModeRef
	_CGDisplayCreateImage func(DirectDisplayID) ImageRef
	_CGDisplayCreateImageForRect func(DirectDisplayID, Rect) ImageRef
	_CGDisplayCurrentMode func(DirectDisplayID) DictionaryRef
	_CGDisplayFade func(DisplayFadeReservationToken, DisplayFadeInterval, DisplayBlendFraction, DisplayBlendFraction, float32, float32, float32, unsafe.Pointer) Error
	_CGDisplayFadeOperationInProgress func() unsafe.Pointer
	_CGDisplayGammaTableCapacity func(DirectDisplayID) uint32
	_CGDisplayGetDrawingContext func(DirectDisplayID) ContextRef
	_CGDisplayHideCursor func(DirectDisplayID) Error
	_CGDisplayIDToOpenGLDisplayMask func(DirectDisplayID) OpenGLDisplayMask
	_CGDisplayIOServicePort func(DirectDisplayID) unsafe.Pointer
	_CGDisplayIsActive func(DirectDisplayID) unsafe.Pointer
	_CGDisplayIsAlwaysInMirrorSet func(DirectDisplayID) unsafe.Pointer
	_CGDisplayIsAsleep func(DirectDisplayID) unsafe.Pointer
	_CGDisplayIsBuiltin func(DirectDisplayID) unsafe.Pointer
	_CGDisplayIsCaptured func(DirectDisplayID) unsafe.Pointer
	_CGDisplayIsInHWMirrorSet func(DirectDisplayID) unsafe.Pointer
	_CGDisplayIsInMirrorSet func(DirectDisplayID) unsafe.Pointer
	_CGDisplayIsMain func(DirectDisplayID) unsafe.Pointer
	_CGDisplayIsOnline func(DirectDisplayID) unsafe.Pointer
	_CGDisplayIsStereo func(DirectDisplayID) unsafe.Pointer
	_CGDisplayMirrorsDisplay func(DirectDisplayID) DirectDisplayID
	_CGDisplayModeGetHeight func(DisplayModeRef) uintptr
	_CGDisplayModeGetIODisplayModeID func(DisplayModeRef) int32
	_CGDisplayModeGetIOFlags func(DisplayModeRef) uint32
	_CGDisplayModeIsUsableForDesktopGUI func(DisplayModeRef) bool
	_CGDisplayModeCopyPixelEncoding func(DisplayModeRef) StringRef
	_CGDisplayModeGetPixelHeight func(DisplayModeRef) uintptr
	_CGDisplayModeGetPixelWidth func(DisplayModeRef) uintptr
	_CGDisplayModeGetRefreshRate func(DisplayModeRef) float64
	_CGDisplayModeGetTypeID func() TypeID
	_CGDisplayModeGetWidth func(DisplayModeRef) uintptr
	_CGDisplayModeRelease func(DisplayModeRef)
	_CGDisplayModeRetain func(DisplayModeRef) DisplayModeRef
	_CGDisplayModelNumber func(DirectDisplayID) uint32
	_CGDisplayMoveCursorToPoint func(DirectDisplayID, Point) Error
	_CGDisplayPixelsHigh func(DirectDisplayID) uintptr
	_CGDisplayPixelsWide func(DirectDisplayID) uintptr
	_CGDisplayPrimaryDisplay func(DirectDisplayID) DirectDisplayID
	_CGDisplayRegisterReconfigurationCallback func(DisplayReconfigurationCallBack, unsafe.Pointer) Error
	_CGDisplayRelease func(DirectDisplayID) Error
	_CGDisplayRemoveReconfigurationCallback func(DisplayReconfigurationCallBack, unsafe.Pointer) Error
	_CGDisplayRestoreColorSyncSettings func()
	_CGDisplayRotation func(DirectDisplayID) float64
	_CGDisplayScreenSize func(DirectDisplayID) Size
	_CGDisplaySerialNumber func(DirectDisplayID) uint32
	_CGDisplaySetDisplayMode func(DirectDisplayID, DisplayModeRef, DictionaryRef) Error
	_CGDisplaySetStereoOperation func(DirectDisplayID, unsafe.Pointer, unsafe.Pointer, ConfigureOption) Error
	_CGDisplayShowCursor func(DirectDisplayID) Error
	_CGDisplayStreamCreateWithDispatchQueue func(DirectDisplayID, uintptr, uintptr, int32, DictionaryRef, unsafe.Pointer, DisplayStreamFrameAvailableHandler) DisplayStreamRef
	_CGDisplayStreamCreate func(DirectDisplayID, uintptr, uintptr, int32, DictionaryRef, DisplayStreamFrameAvailableHandler) DisplayStreamRef
	_CGDisplayStreamGetRunLoopSource func(DisplayStreamRef) RunLoopSourceRef
	_CGDisplayStreamStart func(DisplayStreamRef) Error
	_CGDisplayStreamStop func(DisplayStreamRef) Error
	_CGDisplayStreamGetTypeID func() TypeID
	_CGDisplayStreamUpdateGetDropCount func(DisplayStreamUpdateRef) uintptr
	_CGDisplayStreamUpdateGetMovedRectsDelta func(DisplayStreamUpdateRef, []float64, []float64)
	_CGDisplayStreamUpdateGetRects func(DisplayStreamUpdateRef, DisplayStreamUpdateRectType, unsafe.Pointer) unsafe.Pointer
	_CGDisplayStreamUpdateCreateMergedUpdate func(DisplayStreamUpdateRef, DisplayStreamUpdateRef) DisplayStreamUpdateRef
	_CGDisplayStreamUpdateGetTypeID func() TypeID
	_CGDisplaySwitchToMode func(DirectDisplayID, DictionaryRef) Error
	_CGDisplayUnitNumber func(DirectDisplayID) uint32
	_CGDisplayUsesOpenGLAcceleration func(DirectDisplayID) unsafe.Pointer
	_CGDisplayVendorNumber func(DirectDisplayID) uint32
	_CGEXRToneMappingGammaGetDefaultOptions func() DictionaryRef
	_CGEnableEventStateCombining func(unsafe.Pointer) Error
	_CGErrorSetCallback func(ErrorCallback)
	_CGEventCreateCopy func(EventRef) EventRef
	_CGEventGetFlags func(EventRef) EventFlags
	_CGEventGetDoubleValueField func(EventRef, EventField) float64
	_CGEventGetIntegerValueField func(EventRef, EventField) int64
	_CGEventCreateKeyboardEvent func(EventSourceRef, KeyCode, bool) EventRef
	_CGEventCreateMouseEvent func(EventSourceRef, EventType, Point, MouseButton) EventRef
	_CGEventCreateScrollWheelEvent2 func(EventSourceRef, ScrollEventUnit, uint32, int32, int32, int32) EventRef
	_CGEventCreate func(EventSourceRef) EventRef
	_CGEventCreateFromData func(AllocatorRef, DataRef) EventRef
	_CGEventKeyboardGetUnicodeString func(EventRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CGEventKeyboardSetUnicodeString func(EventRef, unsafe.Pointer, unsafe.Pointer)
	_CGEventGetLocation func(EventRef) Point
	_CGEventPost func(EventTapLocation, EventRef)
	_CGEventPostToPSN func(unsafe.Pointer, EventRef)
	_CGEventPostToPid func(unsafe.Pointer, EventRef)
	_CGEventSetDoubleValueField func(EventRef, EventField, float64)
	_CGEventSetIntegerValueField func(EventRef, EventField, int64)
	_CGEventSetSource func(EventRef, EventSourceRef)
	_CGEventTapCreate func(EventTapLocation, EventTapPlacement, EventTapOptions, EventMask, EventTapCallBack, unsafe.Pointer) MachPortRef
	_CGEventTapCreateForPSN func(unsafe.Pointer, EventTapPlacement, EventTapOptions, EventMask, EventTapCallBack, unsafe.Pointer) MachPortRef
	_CGEventTapCreateForPid func(unsafe.Pointer, EventTapPlacement, EventTapOptions, EventMask, EventTapCallBack, unsafe.Pointer) MachPortRef
	_CGEventTapEnable func(MachPortRef, bool)
	_CGEventTapIsEnabled func(MachPortRef) bool
	_CGEventTapPostEvent func(EventTapProxy, EventRef)
	_CGEventGetTimestamp func(EventRef) EventTimestamp
	_CGEventGetType func(EventRef) EventType
	_CGEventGetTypeID func() TypeID
	_CGEventGetUnflippedLocation func(EventRef) Point
	_CGEventCreateData func(AllocatorRef, EventRef) DataRef
	_CGEventCreateScrollWheelEvent func(EventSourceRef, ScrollEventUnit, uint32, int32) EventRef
	_CGEventSetFlags func(EventRef, EventFlags)
	_CGEventSetLocation func(EventRef, Point)
	_CGEventSetTimestamp func(EventRef, EventTimestamp)
	_CGEventSetType func(EventRef, EventType)
	_CGEventSourceButtonState func(EventSourceStateID, MouseButton) bool
	_CGEventSourceCounterForEventType func(EventSourceStateID, EventType) uint32
	_CGEventSourceFlagsState func(EventSourceStateID) EventFlags
	_CGEventSourceGetLocalEventsFilterDuringSuppressionState func(EventSourceRef, EventSuppressionState) EventFilterMask
	_CGEventCreateSourceFromEvent func(EventRef) EventSourceRef
	_CGEventSourceCreate func(EventSourceStateID) EventSourceRef
	_CGEventSourceKeyState func(EventSourceStateID, KeyCode) bool
	_CGEventSourceGetKeyboardType func(EventSourceRef) EventSourceKeyboardType
	_CGEventSourceGetLocalEventsSuppressionInterval func(EventSourceRef) TimeInterval
	_CGEventSourceGetPixelsPerLine func(EventSourceRef) float64
	_CGEventSourceSecondsSinceLastEventType func(EventSourceStateID, EventType) TimeInterval
	_CGEventSourceSetLocalEventsFilterDuringSuppressionState func(EventSourceRef, EventFilterMask, EventSuppressionState)
	_CGEventSourceGetSourceStateID func(EventSourceRef) EventSourceStateID
	_CGEventSourceGetTypeID func() TypeID
	_CGEventSourceGetUserData func(EventSourceRef) int64
	_CGEventSourceSetKeyboardType func(EventSourceRef, EventSourceKeyboardType)
	_CGEventSourceSetLocalEventsSuppressionInterval func(EventSourceRef, TimeInterval)
	_CGEventSourceSetPixelsPerLine func(EventSourceRef, float64)
	_CGEventSourceSetUserData func(EventSourceRef, int64)
	_CGFontGetAscent func(FontRef) int
	_CGFontCanCreatePostScriptSubset func(FontRef, FontPostScriptFormat) bool
	_CGFontGetCapHeight func(FontRef) int
	_CGFontCreateCopyWithVariations func(FontRef, DictionaryRef) FontRef
	_CGFontCreatePostScriptEncoding func(FontRef, unsafe.Pointer, unsafe.Pointer) DataRef
	_CGFontCreatePostScriptSubset func(FontRef, StringRef, FontPostScriptFormat, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) DataRef
	_CGFontGetDescent func(FontRef) int
	_CGFontGetFontBBox func(FontRef) Rect
	_CGFontCopyFullName func(FontRef) StringRef
	_CGFontGetGlyphAdvances func(FontRef, unsafe.Pointer, uintptr, []int) bool
	_CGFontGetGlyphBBoxes func(FontRef, unsafe.Pointer, uintptr, unsafe.Pointer) bool
	_CGFontGetGlyphWithGlyphName func(FontRef, StringRef) Glyph
	_CGFontCreateWithFontName func(StringRef) FontRef
	_CGFontCreateWithDataProvider func(DataProviderRef) FontRef
	_CGFontGetItalicAngle func(FontRef) Float
	_CGFontGetLeading func(FontRef) int
	_CGFontCopyGlyphNameForGlyph func(FontRef, Glyph) StringRef
	_CGFontGetNumberOfGlyphs func(FontRef) uintptr
	_CGFontCopyPostScriptName func(FontRef) StringRef
	_CGFontGetStemV func(FontRef) Float
	_CGFontCopyTableForTag func(FontRef, uint32) DataRef
	_CGFontCopyTableTags func(FontRef) ArrayRef
	_CGFontGetTypeID func() TypeID
	_CGFontGetUnitsPerEm func(FontRef) int
	_CGFontCopyVariationAxes func(FontRef) ArrayRef
	_CGFontCopyVariations func(FontRef) DictionaryRef
	_CGFontGetXHeight func(FontRef) int
	_CGFontCreateWithPlatformFont func(unsafe.Pointer) FontRef
	_CGFontRelease func(FontRef)
	_CGFontRetain func(FontRef) FontRef
	_CGFunctionCreate func(unsafe.Pointer, uintptr, []float64, uintptr, []float64, unsafe.Pointer) FunctionRef
	_CGFunctionGetTypeID func() TypeID
	_CGFunctionRelease func(FunctionRef)
	_CGFunctionRetain func(FunctionRef) FunctionRef
	_CGGetActiveDisplayList func(uint32, unsafe.Pointer, []uint32) Error
	_CGGetDisplayTransferByFormula func(DirectDisplayID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) Error
	_CGGetDisplayTransferByTable func(DirectDisplayID, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, []uint32) Error
	_CGGetDisplaysWithOpenGLDisplayMask func(OpenGLDisplayMask, uint32, unsafe.Pointer, []uint32) Error
	_CGGetDisplaysWithPoint func(Point, uint32, unsafe.Pointer, []uint32) Error
	_CGGetDisplaysWithRect func(Rect, uint32, unsafe.Pointer, []uint32) Error
	_CGGetEventTapList func(uint32, unsafe.Pointer, []uint32) Error
	_CGGetLastMouseDelta func(unsafe.Pointer, unsafe.Pointer)
	_CGGetOnlineDisplayList func(uint32, unsafe.Pointer, []uint32) Error
	_CGGradientGetContentHeadroom func(GradientRef) float32
	_CGGradientCreateWithColorComponents func(ColorSpaceRef, []float64, []float64, uintptr) GradientRef
	_CGGradientCreateWithColors func(ColorSpaceRef, ArrayRef, []float64) GradientRef
	_CGGradientCreateWithContentHeadroom func(float32, ColorSpaceRef, []float64, []float64, uintptr) GradientRef
	_CGGradientGetTypeID func() TypeID
	_CGGradientRelease func(GradientRef)
	_CGGradientRetain func(GradientRef) GradientRef
	_CGImageGetAlphaInfo func(ImageRef) ImageAlphaInfo
	_CGImageGetBitmapInfo func(ImageRef) BitmapInfo
	_CGImageGetBitsPerComponent func(ImageRef) uintptr
	_CGImageGetBitsPerPixel func(ImageRef) uintptr
	_CGImageGetByteOrderInfo func(ImageRef) ImageByteOrderInfo
	_CGImageGetBytesPerRow func(ImageRef) uintptr
	_CGImageCalculateContentAverageLightLevel func(ImageRef) float32
	_CGImageCalculateContentHeadroom func(ImageRef) float32
	_CGImageGetColorSpace func(ImageRef) ColorSpaceRef
	_CGImageContainsImageSpecificToneMappingMetadata func(ImageRef) bool
	_CGImageGetContentAverageLightLevel func(ImageRef) float32
	_CGImageGetContentHeadroom func(ImageRef) float32
	_CGImageCreateCopy func(ImageRef) ImageRef
	_CGImageCreateCopyWithColorSpace func(ImageRef, ColorSpaceRef) ImageRef
	_CGImageCreateCopyWithContentAverageLightLevel func(ImageRef, float32) ImageRef
	_CGImageCreateCopyWithCalculatedHDRStats func(ImageRef) ImageRef
	_CGImageCreateWithImageInRect func(ImageRef, Rect) ImageRef
	_CGImageGetDataProvider func(ImageRef) DataProviderRef
	_CGImageGetDecode func(ImageRef) []float64
	_CGImageGetHeight func(ImageRef) uintptr
	_CGImageCreateWithContentHeadroom func(float32, uintptr, uintptr, uintptr, uintptr, uintptr, ColorSpaceRef, BitmapInfo, DataProviderRef, []float64, bool, ColorRenderingIntent) ImageRef
	_CGImageCreateWithJPEGDataProvider func(DataProviderRef, []float64, bool, ColorRenderingIntent) ImageRef
	_CGImageMaskCreate func(uintptr, uintptr, uintptr, uintptr, uintptr, DataProviderRef, []float64, bool) ImageRef
	_CGImageCreateWithPNGDataProvider func(DataProviderRef, []float64, bool, ColorRenderingIntent) ImageRef
	_CGImageCreate func(uintptr, uintptr, uintptr, uintptr, uintptr, ColorSpaceRef, BitmapInfo, DataProviderRef, []float64, bool, ColorRenderingIntent) ImageRef
	_CGWindowListCreateImageFromArray func(Rect, ArrayRef, WindowImageOption) ImageRef
	_CGImageIsMask func(ImageRef) bool
	_CGImageCreateWithMask func(ImageRef, ImageRef) ImageRef
	_CGImageGetPixelFormatInfo func(ImageRef) ImagePixelFormatInfo
	_CGImageGetRenderingIntent func(ImageRef) ColorRenderingIntent
	_CGImageGetShouldInterpolate func(ImageRef) bool
	_CGImageShouldToneMap func(ImageRef) bool
	_CGImageGetTypeID func() TypeID
	_CGImageGetUTType func(ImageRef) StringRef
	_CGImageGetWidth func(ImageRef) uintptr
	_CGImageCreateCopyWithContentHeadroom func(float32, ImageRef) ImageRef
	_CGImageCreateWithMaskingColors func(ImageRef, []float64) ImageRef
	_CGImageRelease func(ImageRef)
	_CGImageRetain func(ImageRef) ImageRef
	_CGInhibitLocalEvents func(unsafe.Pointer) Error
	_CGLayerGetContext func(LayerRef) ContextRef
	_CGLayerCreateWithContext func(ContextRef, Size, DictionaryRef) LayerRef
	_CGLayerGetSize func(LayerRef) Size
	_CGLayerGetTypeID func() TypeID
	_CGLayerRelease func(LayerRef)
	_CGLayerRetain func(LayerRef) LayerRef
	_CGMainDisplayID func() DirectDisplayID
	_CGPathCloseSubpath func(MutablePathRef)
	_CGPathCreateMutable func() MutablePathRef
	_CGOpenGLDisplayMaskToDisplayID func(OpenGLDisplayMask) DirectDisplayID
	_CGPDFArrayApplyBlock func(PDFArrayRef, PDFArrayApplierBlock, unsafe.Pointer)
	_CGPDFArrayGetArray func(PDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetBoolean func(PDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetCount func(PDFArrayRef) uintptr
	_CGPDFArrayGetDictionary func(PDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetInteger func(PDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetName func(PDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetNull func(PDFArrayRef, uintptr) bool
	_CGPDFArrayGetNumber func(PDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetObject func(PDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetStream func(PDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFArrayGetString func(PDFArrayRef, uintptr, unsafe.Pointer) bool
	_CGPDFContentStreamCreateWithPage func(PDFPageRef) PDFContentStreamRef
	_CGPDFContentStreamCreateWithStream func(PDFStreamRef, PDFDictionaryRef, PDFContentStreamRef) PDFContentStreamRef
	_CGPDFContentStreamGetResource func(PDFContentStreamRef, unsafe.Pointer, unsafe.Pointer) PDFObjectRef
	_CGPDFContentStreamGetStreams func(PDFContentStreamRef) ArrayRef
	_CGPDFContentStreamRelease func(PDFContentStreamRef)
	_CGPDFContentStreamRetain func(PDFContentStreamRef) PDFContentStreamRef
	_CGPDFContextBeginTag func(ContextRef, PDFTagType, DictionaryRef)
	_CGPDFContextEndTag func(ContextRef)
	_CGPDFContextSetIDTree func(ContextRef, PDFDictionaryRef)
	_CGPDFContextSetOutline func(ContextRef, DictionaryRef)
	_CGPDFContextSetPageTagStructureTree func(ContextRef, DictionaryRef)
	_CGPDFContextSetParentTree func(ContextRef, PDFDictionaryRef)
	_CGPDFDictionaryApplyBlock func(PDFDictionaryRef, PDFDictionaryApplierBlock, unsafe.Pointer)
	_CGPDFDictionaryApplyFunction func(PDFDictionaryRef, PDFDictionaryApplierFunction, unsafe.Pointer)
	_CGPDFDictionaryGetArray func(PDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetBoolean func(PDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetCount func(PDFDictionaryRef) uintptr
	_CGPDFDictionaryGetDictionary func(PDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetInteger func(PDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetName func(PDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetNumber func(PDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetObject func(PDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetStream func(PDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDictionaryGetString func(PDFDictionaryRef, unsafe.Pointer, unsafe.Pointer) bool
	_CGPDFDocumentGetAccessPermissions func(PDFDocumentRef) PDFAccessPermissions
	_CGPDFDocumentAllowsCopying func(PDFDocumentRef) bool
	_CGPDFDocumentAllowsPrinting func(PDFDocumentRef) bool
	_CGPDFDocumentGetCatalog func(PDFDocumentRef) PDFDictionaryRef
	_CGPDFDocumentGetID func(PDFDocumentRef) PDFArrayRef
	_CGPDFDocumentGetVersion func(PDFDocumentRef, []int, []int)
	_CGPDFDocumentGetInfo func(PDFDocumentRef) PDFDictionaryRef
	_CGPDFDocumentCreateWithURL func(URLRef) PDFDocumentRef
	_CGPDFDocumentCreateWithProvider func(DataProviderRef) PDFDocumentRef
	_CGPDFDocumentIsEncrypted func(PDFDocumentRef) bool
	_CGPDFDocumentIsUnlocked func(PDFDocumentRef) bool
	_CGPDFDocumentGetNumberOfPages func(PDFDocumentRef) uintptr
	_CGPDFDocumentGetOutline func(PDFDocumentRef) DictionaryRef
	_CGPDFDocumentGetPage func(PDFDocumentRef, uintptr) PDFPageRef
	_CGPDFDocumentGetTypeID func() TypeID
	_CGPDFDocumentUnlockWithPassword func(PDFDocumentRef, unsafe.Pointer) bool
	_CGPDFDocumentGetArtBox func(PDFDocumentRef, int) Rect
	_CGPDFDocumentGetBleedBox func(PDFDocumentRef, int) Rect
	_CGPDFDocumentGetCropBox func(PDFDocumentRef, int) Rect
	_CGPDFDocumentGetMediaBox func(PDFDocumentRef, int) Rect
	_CGPDFDocumentGetRotationAngle func(PDFDocumentRef, int) int
	_CGPDFDocumentGetTrimBox func(PDFDocumentRef, int) Rect
	_CGPDFDocumentRelease func(PDFDocumentRef)
	_CGPDFDocumentRetain func(PDFDocumentRef) PDFDocumentRef
	_CGPDFObjectGetType func(PDFObjectRef) PDFObjectType
	_CGPDFObjectGetValue func(PDFObjectRef, PDFObjectType, unsafe.Pointer) bool
	_CGPDFOperatorTableCreate func() PDFOperatorTableRef
	_CGPDFOperatorTableRelease func(PDFOperatorTableRef)
	_CGPDFOperatorTableRetain func(PDFOperatorTableRef) PDFOperatorTableRef
	_CGPDFOperatorTableSetCallback func(PDFOperatorTableRef, unsafe.Pointer, PDFOperatorCallback)
	_CGPDFPageGetDictionary func(PDFPageRef) PDFDictionaryRef
	_CGPDFPageGetDocument func(PDFPageRef) PDFDocumentRef
	_CGPDFPageGetBoxRect func(PDFPageRef, PDFBox) Rect
	_CGPDFPageGetDrawingTransform func(PDFPageRef, PDFBox, Rect, int, bool) AffineTransform
	_CGPDFPageGetPageNumber func(PDFPageRef) uintptr
	_CGPDFPageGetRotationAngle func(PDFPageRef) int
	_CGPDFPageGetTypeID func() TypeID
	_CGPDFPageRelease func(PDFPageRef)
	_CGPDFPageRetain func(PDFPageRef) PDFPageRef
	_CGPDFScannerCreate func(PDFContentStreamRef, PDFOperatorTableRef, unsafe.Pointer) PDFScannerRef
	_CGPDFScannerGetContentStream func(PDFScannerRef) PDFContentStreamRef
	_CGPDFScannerPopArray func(PDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopBoolean func(PDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopDictionary func(PDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopInteger func(PDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopName func(PDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopNumber func(PDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopObject func(PDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopStream func(PDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerPopString func(PDFScannerRef, unsafe.Pointer) bool
	_CGPDFScannerRelease func(PDFScannerRef)
	_CGPDFScannerRetain func(PDFScannerRef) PDFScannerRef
	_CGPDFScannerScan func(PDFScannerRef) bool
	_CGPDFScannerStop func(PDFScannerRef)
	_CGPDFStreamCopyData func(PDFStreamRef, unsafe.Pointer) DataRef
	_CGPDFStreamGetDictionary func(PDFStreamRef) PDFDictionaryRef
	_CGPDFStringCopyDate func(PDFStringRef) DateRef
	_CGPDFStringCopyTextString func(PDFStringRef) StringRef
	_CGPDFStringGetBytePtr func(PDFStringRef) unsafe.Pointer
	_CGPDFStringGetLength func(PDFStringRef) uintptr
	_CGPDFTagTypeGetName func(PDFTagType) unsafe.Pointer
	_CGPSConverterAbort func(PSConverterRef) bool
	_CGPSConverterConvert func(PSConverterRef, DataProviderRef, DataConsumerRef, DictionaryRef) bool
	_CGPSConverterCreate func(unsafe.Pointer, unsafe.Pointer, DictionaryRef) PSConverterRef
	_CGPSConverterIsConverting func(PSConverterRef) bool
	_CGPSConverterGetTypeID func() TypeID
	_CGPathApply func(PathRef, unsafe.Pointer, PathApplierFunction)
	_CGPathApplyWithBlock func(PathRef, PathApplyBlock)
	_CGPathGetBoundingBox func(PathRef) Rect
	_CGPathGetPathBoundingBox func(PathRef) Rect
	_CGPathCreateCopy func(PathRef) PathRef
	_CGPathCreateCopyByTransformingPath func(PathRef, unsafe.Pointer) PathRef
	_CGPathGetCurrentPoint func(PathRef) Point
	_CGPathCreateWithEllipseInRect func(Rect, unsafe.Pointer) PathRef
	_CGPathCreateWithRect func(Rect, unsafe.Pointer) PathRef
	_CGPathCreateWithRoundedRect func(Rect, Float, Float, unsafe.Pointer) PathRef
	_CGPathIsEmpty func(PathRef) bool
	_CGPathIsRect func(PathRef, unsafe.Pointer) bool
	_CGPathCreateMutableCopy func(PathRef) MutablePathRef
	_CGPathCreateMutableCopyByTransformingPath func(PathRef, unsafe.Pointer) MutablePathRef
	_CGPathGetTypeID func() TypeID
	_CGPathAddArc func(MutablePathRef, unsafe.Pointer, Float, Float, Float, Float, Float, bool)
	_CGPathAddArcToPoint func(MutablePathRef, unsafe.Pointer, Float, Float, Float, Float, Float)
	_CGPathAddCurveToPoint func(MutablePathRef, unsafe.Pointer, Float, Float, Float, Float, Float, Float)
	_CGPathAddEllipseInRect func(MutablePathRef, unsafe.Pointer, Rect)
	_CGPathAddLineToPoint func(MutablePathRef, unsafe.Pointer, Float, Float)
	_CGPathAddLines func(MutablePathRef, unsafe.Pointer, unsafe.Pointer, uintptr)
	_CGPathAddPath func(MutablePathRef, unsafe.Pointer, PathRef)
	_CGPathAddQuadCurveToPoint func(MutablePathRef, unsafe.Pointer, Float, Float, Float, Float)
	_CGPathAddRect func(MutablePathRef, unsafe.Pointer, Rect)
	_CGPathAddRects func(MutablePathRef, unsafe.Pointer, unsafe.Pointer, uintptr)
	_CGPathAddRelativeArc func(MutablePathRef, unsafe.Pointer, Float, Float, Float, Float, Float)
	_CGPathAddRoundedRect func(MutablePathRef, unsafe.Pointer, Rect, Float, Float)
	_CGPathContainsPoint func(PathRef, unsafe.Pointer, Point, bool) bool
	_CGPathCreateCopyByDashingPath func(PathRef, unsafe.Pointer, Float, []float64, uintptr) PathRef
	_CGPathCreateCopyByFlattening func(PathRef, Float) PathRef
	_CGPathCreateCopyByIntersectingPath func(PathRef, PathRef, bool) PathRef
	_CGPathCreateCopyByNormalizing func(PathRef, bool) PathRef
	_CGPathCreateCopyByStrokingPath func(PathRef, unsafe.Pointer, Float, LineCap, LineJoin, Float) PathRef
	_CGPathCreateCopyBySubtractingPath func(PathRef, PathRef, bool) PathRef
	_CGPathCreateCopyBySymmetricDifferenceOfPath func(PathRef, PathRef, bool) PathRef
	_CGPathCreateCopyByUnioningPath func(PathRef, PathRef, bool) PathRef
	_CGPathCreateCopyOfLineByIntersectingPath func(PathRef, PathRef, bool) PathRef
	_CGPathCreateCopyOfLineBySubtractingPath func(PathRef, PathRef, bool) PathRef
	_CGPathCreateSeparateComponents func(PathRef, bool) ArrayRef
	_CGPathEqualToPath func(PathRef, PathRef) bool
	_CGPathIntersectsPath func(PathRef, PathRef, bool) bool
	_CGPathMoveToPoint func(MutablePathRef, unsafe.Pointer, Float, Float)
	_CGPathRelease func(PathRef)
	_CGPathRetain func(PathRef) PathRef
	_CGPatternCreate func(unsafe.Pointer, Rect, AffineTransform, Float, Float, PatternTiling, bool, unsafe.Pointer) PatternRef
	_CGPatternGetTypeID func() TypeID
	_CGPatternRelease func(PatternRef)
	_CGPatternRetain func(PatternRef) PatternRef
	_CGPointApplyAffineTransform func(Point, AffineTransform) Point
	_CGPointCreateDictionaryRepresentation func(Point) DictionaryRef
	_CGPointEqualToPoint func(Point, Point) bool
	_CGPointMakeWithDictionaryRepresentation func(DictionaryRef, unsafe.Pointer) bool
	_CGPostKeyboardEvent func(CharCode, KeyCode, unsafe.Pointer) Error
	_CGPostMouseEvent func(Point, unsafe.Pointer, ButtonCount, unsafe.Pointer) Error
	_CGPostScrollWheelEvent func(WheelCount, int32) Error
	_CGPreflightListenEventAccess func() bool
	_CGPreflightPostEventAccess func() bool
	_CGPreflightScreenCaptureAccess func() bool
	_CGRectApplyAffineTransform func(Rect, AffineTransform) Rect
	_CGRectContainsPoint func(Rect, Point) bool
	_CGRectContainsRect func(Rect, Rect) bool
	_CGRectCreateDictionaryRepresentation func(Rect) DictionaryRef
	_CGRectDivide func(Rect, unsafe.Pointer, unsafe.Pointer, Float, RectEdge)
	_CGRectEqualToRect func(Rect, Rect) bool
	_CGRectGetHeight func(Rect) Float
	_CGRectGetMaxX func(Rect) Float
	_CGRectGetMaxY func(Rect) Float
	_CGRectGetMidX func(Rect) Float
	_CGRectGetMidY func(Rect) Float
	_CGRectGetMinX func(Rect) Float
	_CGRectGetMinY func(Rect) Float
	_CGRectGetWidth func(Rect) Float
	_CGRectInset func(Rect, Float, Float) Rect
	_CGRectIntegral func(Rect) Rect
	_CGRectIntersection func(Rect, Rect) Rect
	_CGRectIntersectsRect func(Rect, Rect) bool
	_CGRectIsEmpty func(Rect) bool
	_CGRectIsInfinite func(Rect) bool
	_CGRectIsNull func(Rect) bool
	_CGRectMakeWithDictionaryRepresentation func(DictionaryRef, unsafe.Pointer) bool
	_CGRectOffset func(Rect, Float, Float) Rect
	_CGRectStandardize func(Rect) Rect
	_CGRectUnion func(Rect, Rect) Rect
	_CGRegisterScreenRefreshCallback func(ScreenRefreshCallback, unsafe.Pointer) Error
	_CGReleaseAllDisplays func() Error
	_CGReleaseDisplayFadeReservation func(DisplayFadeReservationToken) Error
	_CGReleaseScreenRefreshRects func(unsafe.Pointer)
	_CGRenderingBufferLockBytePtr func(RenderingBufferProviderRef) unsafe.Pointer
	_CGRenderingBufferProviderCreate func(unsafe.Pointer, uintptr) RenderingBufferProviderRef
	_CGRenderingBufferProviderCreateWithCFData func(MutableDataRef) RenderingBufferProviderRef
	_CGRenderingBufferProviderGetSize func(RenderingBufferProviderRef) uintptr
	_CGRenderingBufferProviderGetTypeID func() TypeID
	_CGRenderingBufferUnlockBytePtr func(RenderingBufferProviderRef)
	_CGRequestListenEventAccess func() bool
	_CGRequestPostEventAccess func() bool
	_CGRequestScreenCaptureAccess func() bool
	_CGRestorePermanentDisplayConfiguration func()
	_CGScreenRegisterMoveCallback func(ScreenUpdateMoveCallback, unsafe.Pointer) Error
	_CGScreenUnregisterMoveCallback func(ScreenUpdateMoveCallback, unsafe.Pointer)
	_CGSessionCopyCurrentDictionary func() DictionaryRef
	_CGSetDisplayTransferByByteTable func(DirectDisplayID, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) Error
	_CGSetDisplayTransferByFormula func(DirectDisplayID, GammaValue, GammaValue, GammaValue, GammaValue, GammaValue, GammaValue, GammaValue, GammaValue, GammaValue) Error
	_CGSetDisplayTransferByTable func(DirectDisplayID, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) Error
	_CGSetLocalEventsFilterDuringSuppressionState func(EventFilterMask, EventSuppressionState) Error
	_CGSetLocalEventsSuppressionInterval func(TimeInterval) Error
	_CGShadingGetContentHeadroom func(ShadingRef) float32
	_CGShadingCreateAxialWithContentHeadroom func(float32, ColorSpaceRef, Point, Point, FunctionRef, bool, bool) ShadingRef
	_CGShadingCreateAxial func(ColorSpaceRef, Point, Point, FunctionRef, bool, bool) ShadingRef
	_CGShadingCreateRadialWithContentHeadroom func(float32, ColorSpaceRef, Point, Float, Point, Float, FunctionRef, bool, bool) ShadingRef
	_CGShadingCreateRadial func(ColorSpaceRef, Point, Float, Point, Float, FunctionRef, bool, bool) ShadingRef
	_CGShadingGetTypeID func() TypeID
	_CGShadingRelease func(ShadingRef)
	_CGShadingRetain func(ShadingRef) ShadingRef
	_CGShieldingWindowID func(DirectDisplayID) WindowID
	_CGShieldingWindowLevel func() WindowLevel
	_CGSizeApplyAffineTransform func(Size, AffineTransform) Size
	_CGSizeCreateDictionaryRepresentation func(Size) DictionaryRef
	_CGSizeEqualToSize func(Size, Size) bool
	_CGSizeMakeWithDictionaryRepresentation func(DictionaryRef, unsafe.Pointer) bool
	_CGUnregisterScreenRefreshCallback func(ScreenRefreshCallback, unsafe.Pointer)
	_CGWaitForScreenRefreshRects func(unsafe.Pointer, []uint32) Error
	_CGWaitForScreenUpdateRects func(ScreenUpdateOperation, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) Error
	_CGWarpMouseCursorPosition func(Point) Error
	_CGWindowLevelForKey func(WindowLevelKey) WindowLevel
	_CGWindowListCopyWindowInfo func(WindowListOption, WindowID) ArrayRef
	_CGWindowListCreate func(WindowListOption, WindowID) ArrayRef
	_CGWindowListCreateDescriptionFromArray func(ArrayRef) ArrayRef
	_CGWindowListCreateImage func(Rect, WindowListOption, WindowID, WindowImageOption) ImageRef
	_CGWindowServerCFMachPort func() MachPortRef
	_CGWindowServerCreateServerPort func() MachPortRef
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
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
	tryRegister(&_CGColorSpaceRelease, lib, "CGColorSpaceRelease")
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
	tryRegister(&_CGContextSetInterpolationQuality, lib, "CGContextSetInterpolationQuality")
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



// Reserves the fade hardware for a specified time interval.
//
// Added in macOS 10.2.
// Reserves the fade hardware for a specified time interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAcquireDisplayFadeReservation(_:_:)
func CGAcquireDisplayFadeReservation(seconds DisplayReservationInterval, token unsafe.Pointer) Error {
	return _CGAcquireDisplayFadeReservation(seconds, token)
}

// Returns an affine transformation matrix constructed by combining two existing affine transforms.
//
// Added in macOS 10.0.
// Returns an affine transformation matrix constructed by combining two existing affine transforms.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformConcat(_:_:)
func CGAffineTransformConcat(t1 AffineTransform, t2 AffineTransform) AffineTransform {
	return _CGAffineTransformConcat(t1, t2)
}

// CGAffineTransformDecompose is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformDecompose
func CGAffineTransformDecompose(transform AffineTransform) corefoundation.AffineTransformComponents {
	return _CGAffineTransformDecompose(transform)
}

// Checks whether two affine transforms are equal.
//
// Added in macOS 10.4.
// Checks whether two affine transforms are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformEqualToTransform(_:_:)
func CGAffineTransformEqualToTransform(t1 AffineTransform, t2 AffineTransform) bool {
	return _CGAffineTransformEqualToTransform(t1, t2)
}

// Returns an affine transformation matrix constructed by inverting an existing affine transform.
//
// Added in macOS 10.0.
// Returns an affine transformation matrix constructed by inverting an existing affine transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformInvert(_:)
func CGAffineTransformInvert(t AffineTransform) AffineTransform {
	return _CGAffineTransformInvert(t)
}

// Checks whether an affine transform is the identity transform.
//
// Added in macOS 10.4.
// Checks whether an affine transform is the identity transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformIsIdentity(_:)
func CGAffineTransformIsIdentity(t AffineTransform) bool {
	return _CGAffineTransformIsIdentity(t)
}

// Returns an affine transformation matrix constructed from values you provide.
//
// Added in macOS 10.0.
// Returns an affine transformation matrix constructed from values you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMake(_:_:_:_:_:_:)
func CGAffineTransformMake(a Float, b Float, c Float, d Float, tx Float, ty Float) AffineTransform {
	return _CGAffineTransformMake(a, b, c, d, tx, ty)
}

// Returns an affine transformation matrix constructed from a rotation value you provide.
//
// Added in macOS 10.0.
// Returns an affine transformation matrix constructed from a rotation value you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeRotation(_:)
func CGAffineTransformMakeRotation(angle Float) AffineTransform {
	return _CGAffineTransformMakeRotation(angle)
}

// Returns an affine transformation matrix constructed from scaling values you provide.
//
// Added in macOS 10.0.
// Returns an affine transformation matrix constructed from scaling values you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeScale(_:_:)
func CGAffineTransformMakeScale(sx Float, sy Float) AffineTransform {
	return _CGAffineTransformMakeScale(sx, sy)
}

// Returns an affine transformation matrix constructed from translation values you provide.
//
// Added in macOS 10.0.
// Returns an affine transformation matrix constructed from translation values you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeTranslation(_:_:)
func CGAffineTransformMakeTranslation(tx Float, ty Float) AffineTransform {
	return _CGAffineTransformMakeTranslation(tx, ty)
}

// CGAffineTransformMakeWithComponents is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeWithComponents
func CGAffineTransformMakeWithComponents(components corefoundation.AffineTransformComponents) AffineTransform {
	return _CGAffineTransformMakeWithComponents(components)
}

// Returns an affine transformation matrix constructed by rotating an existing affine transform.
//
// Added in macOS 10.0.
// Returns an affine transformation matrix constructed by rotating an existing affine transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformRotate(_:_:)
func CGAffineTransformRotate(t AffineTransform, angle Float) AffineTransform {
	return _CGAffineTransformRotate(t, angle)
}

// Returns an affine transformation matrix constructed by scaling an existing affine transform.
//
// Added in macOS 10.0.
// Returns an affine transformation matrix constructed by scaling an existing affine transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformScale(_:_:_:)
func CGAffineTransformScale(t AffineTransform, sx Float, sy Float) AffineTransform {
	return _CGAffineTransformScale(t, sx, sy)
}

// Returns an affine transformation matrix constructed by translating an existing affine transform.
//
// Added in macOS 10.0.
// Returns an affine transformation matrix constructed by translating an existing affine transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformTranslate(_:_:_:)
func CGAffineTransformTranslate(t AffineTransform, tx Float, ty Float) AffineTransform {
	return _CGAffineTransformTranslate(t, tx, ty)
}

// Connects or disconnects the mouse and cursor while an application is in the foreground.
//
// Added in macOS 10.0.
// Connects or disconnects the mouse and cursor while an application is in the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGAssociateMouseAndMouseCursorPosition(_:)
func CGAssociateMouseAndMouseCursorPosition(connected unsafe.Pointer) Error {
	return _CGAssociateMouseAndMouseCursorPosition(connected)
}

// Begins a new set of display configuration changes.
//
// Added in macOS 10.0.
// Begins a new set of display configuration changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBeginDisplayConfiguration(_:)
func CGBeginDisplayConfiguration(config unsafe.Pointer) Error {
	return _CGBeginDisplayConfiguration(config)
}

// CGBitmapContextCreateAdaptive is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapContextCreateAdaptive
func CGBitmapContextCreateAdaptive(width uintptr, height uintptr, auxiliaryInfo DictionaryRef, onResolve bool) ContextRef {
	return _CGBitmapContextCreateAdaptive(width, height, auxiliaryInfo, onResolve)
}

// Cancels a set of display configuration changes.
//
// Added in macOS 10.0.
// Cancels a set of display configuration changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCancelDisplayConfiguration(_:)
func CGCancelDisplayConfiguration(config DisplayConfigRef) Error {
	return _CGCancelDisplayConfiguration(config)
}

// Obtains exclusive use of all active displays, preventing other applications and system services from using the display or changing its configuration.
//
// Added in macOS 10.0.
// Obtains exclusive use of all active displays, preventing other applications and system services from using the display or changing its configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCaptureAllDisplays()
func CGCaptureAllDisplays() Error {
	return _CGCaptureAllDisplays()
}

// Captures all attached displays, using the specified options.
//
// Added in macOS 10.3.
// Captures all attached displays, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCaptureAllDisplaysWithOptions(_:)
func CGCaptureAllDisplaysWithOptions(options CaptureOptions) Error {
	return _CGCaptureAllDisplaysWithOptions(options)
}

// Returns the value of the alpha component associated with a color.
//
// Added in macOS 10.3.
// Returns the value of the alpha component associated with a color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/alpha
func CGColorGetAlpha(color ColorRef) Float {
	return _CGColorGetAlpha(color)
}

// Returns the color space associated with a color.
//
// Added in macOS 10.3.
// Returns the color space associated with a color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/colorSpace
func CGColorGetColorSpace(color ColorRef) ColorSpaceRef {
	return _CGColorGetColorSpace(color)
}

// CGColorGetContentHeadroom is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/contentHeadroom
func CGColorGetContentHeadroom(color ColorRef) float32 {
	return _CGColorGetContentHeadroom(color)
}

// Creates a new color in a different color space that matches the provided color.
//
// Added in macOS 10.11.
// Creates a new color in a different color space that matches the provided color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/converted(to:intent:options:)
func CGColorCreateCopyByMatchingToColorSpace(p0 ColorSpaceRef, intent ColorRenderingIntent, color ColorRef, options DictionaryRef) ColorRef {
	return _CGColorCreateCopyByMatchingToColorSpace(p0, intent, color, options)
}

// Creates a copy of an existing color.
//
// Added in macOS 10.3.
// Creates a copy of an existing color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/copy()
func CGColorCreateCopy(color ColorRef) ColorRef {
	return _CGColorCreateCopy(color)
}

// Creates a copy of an existing color, substituting a new alpha value.
//
// Added in macOS 10.3.
// Creates a copy of an existing color, substituting a new alpha value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/copy(alpha:)
func CGColorCreateCopyWithAlpha(color ColorRef, alpha Float) ColorRef {
	return _CGColorCreateCopyWithAlpha(color, alpha)
}

// Creates a color using a list of intensity values (including alpha) and an associated color space.
//
// Added in macOS 10.3.
// Creates a color using a list of intensity values (including alpha) and an associated color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(colorSpace:components:)
func CGColorCreate(space ColorSpaceRef, components []float64) ColorRef {
	return _CGColorCreate(space, components)
}

// Creates a color in the Generic CMYK color space.
//
// Added in macOS 10.5.
// Creates a color in the Generic CMYK color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(genericCMYKCyan:magenta:yellow:black:alpha:)
func CGColorCreateGenericCMYK(cyan Float, magenta Float, yellow Float, black Float, alpha Float) ColorRef {
	return _CGColorCreateGenericCMYK(cyan, magenta, yellow, black, alpha)
}

// Creates a color in the Generic gray color space with a gamma ramp of 2.2.
//
// Added in macOS 10.15.
// Creates a color in the Generic gray color space with a gamma ramp of 2.2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(genericGrayGamma2_2Gray:alpha:)
func CGColorCreateGenericGrayGamma2_2(gray Float, alpha Float) ColorRef {
	return _CGColorCreateGenericGrayGamma2_2(gray, alpha)
}

// Creates a color in the Generic gray color space.
//
// Added in macOS 10.5.
// Creates a color in the Generic gray color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(gray:alpha:)
func CGColorCreateGenericGray(gray Float, alpha Float) ColorRef {
	return _CGColorCreateGenericGray(gray, alpha)
}

// CGColorCreateWithContentHeadroom is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(headroom:colorSpace:red:green:blue:alpha:)
func CGColorCreateWithContentHeadroom(headroom float32, space ColorSpaceRef, red Float, green Float, blue Float, alpha Float) ColorRef {
	return _CGColorCreateWithContentHeadroom(headroom, space, red, green, blue, alpha)
}

// Creates a color using a list of intensity values (including alpha), a pattern color space, and a pattern.
//
// Added in macOS 10.3.
// Creates a color using a list of intensity values (including alpha), a pattern color space, and a pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(patternSpace:pattern:components:)
func CGColorCreateWithPattern(space ColorSpaceRef, pattern PatternRef, components []float64) ColorRef {
	return _CGColorCreateWithPattern(space, pattern, components)
}

// Creates a color in the Generic RGB color space.
//
// Added in macOS 10.5.
// Creates a color in the Generic RGB color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(red:green:blue:alpha:)
func CGColorCreateGenericRGB(red Float, green Float, blue Float, alpha Float) ColorRef {
	return _CGColorCreateGenericRGB(red, green, blue, alpha)
}

// Creates a color in the sRGB color space.
//
// Added in macOS 10.15.
// Creates a color in the sRGB color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(srgbRed:green:blue:alpha:)
func CGColorCreateSRGB(red Float, green Float, blue Float, alpha Float) ColorRef {
	return _CGColorCreateSRGB(red, green, blue, alpha)
}

// Returns the number of color components (including alpha) associated with a color.
//
// Added in macOS 10.3.
// Returns the number of color components (including alpha) associated with a color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/numberOfComponents
func CGColorGetNumberOfComponents(color ColorRef) uintptr {
	return _CGColorGetNumberOfComponents(color)
}

// Returns the pattern associated with a color in a pattern color space.
//
// Added in macOS 10.3.
// Returns the pattern associated with a color in a pattern color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/pattern
func CGColorGetPattern(color ColorRef) PatternRef {
	return _CGColorGetPattern(color)
}

// Returns the Core Foundation type identifier for a color data type.
//
// Added in macOS 10.3.
// Returns the Core Foundation type identifier for a color data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor/typeID
func CGColorGetTypeID() TypeID {
	return _CGColorGetTypeID()
}

// CGColorConversionInfoConvertData is a CoreGraphics function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/convert(width:height:to:format:from:format:options:)
func CGColorConversionInfoConvertData(info ColorConversionInfoRef, width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorBufferFormat, src_data unsafe.Pointer, src_format CGColorBufferFormat, options DictionaryRef) bool {
	return _CGColorConversionInfoConvertData(info, width, height, dst_data, dst_format, src_data, src_format, options)
}

// CGColorConversionInfoCreateWithOptions is a CoreGraphics function.
//
// Added in macOS 10.14.6.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/init(optionsSrc:dst:options:)
func CGColorConversionInfoCreateWithOptions(src ColorSpaceRef, dst ColorSpaceRef, options DictionaryRef) ColorConversionInfoRef {
	return _CGColorConversionInfoCreateWithOptions(src, dst, options)
}

// Creates a conversion between two specified color spaces.
//
// Added in macOS 10.12.
// Creates a conversion between two specified color spaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/init(src:dst:)
func CGColorConversionInfoCreate(src ColorSpaceRef, dst ColorSpaceRef) ColorConversionInfoRef {
	return _CGColorConversionInfoCreate(src, dst)
}

// CGColorConversionInfoCreateForToneMapping is a CoreGraphics function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/init(src:srcHeadroom:dst:dstHeadroom:toneMapping:options:_:)
func CGColorConversionInfoCreateForToneMapping(from ColorSpaceRef, source_headroom float32, to ColorSpaceRef, target_headroom float32, method ToneMapping, options DictionaryRef, error_ unsafe.Pointer) ColorConversionInfoRef {
	return _CGColorConversionInfoCreateForToneMapping(from, source_headroom, to, target_headroom, method, options, error_)
}

// Returns the Core Foundation type identifier for a color conversion info data type.
//
// Added in macOS .
// Returns the Core Foundation type identifier for a color conversion info data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/typeID
func CGColorConversionInfoGetTypeID() TypeID {
	return _CGColorConversionInfoGetTypeID()
}

// Creates a conversion between an arbitrary number of specified color spaces.
//
// Added in macOS 10.12.
// Creates a conversion between an arbitrary number of specified color spaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoCreateFromList
func CGColorConversionInfoCreateFromList(options DictionaryRef, p1 ColorSpaceRef, p2 ColorConversionInfoTransformType, p3 ColorRenderingIntent) ColorConversionInfoRef {
	return _CGColorConversionInfoCreateFromList(options, p1, p2, p3)
}

// CGColorConversionInfoCreateFromListWithArguments is a CoreGraphics function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoCreateFromListWithArguments
func CGColorConversionInfoCreateFromListWithArguments(options DictionaryRef, p1 ColorSpaceRef, p2 ColorConversionInfoTransformType, p3 ColorRenderingIntent, p4 unsafe.Pointer) ColorConversionInfoRef {
	return _CGColorConversionInfoCreateFromListWithArguments(options, p1, p2, p3, p4)
}

// Indicates whether two colors are equal.
//
// Added in macOS 10.3.
// Indicates whether two colors are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorEqualToColor
func CGColorEqualToColor(color1 ColorRef, color2 ColorRef) bool {
	return _CGColorEqualToColor(color1, color2)
}

// Returns the values of the color components (including alpha) associated with a color.
//
// Added in macOS 10.3.
// Returns the values of the color components (including alpha) associated with a color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorGetComponents
func CGColorGetComponents(color ColorRef) []float64 {
	return _CGColorGetComponents(color)
}

// Returns a color object that represents a constant color.
//
// Added in macOS 10.5.
// Returns a color object that represents a constant color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorGetConstantColor
func CGColorGetConstantColor(colorName StringRef) ColorRef {
	return _CGColorGetConstantColor(colorName)
}

// Decrements the retain count of a color.
//
// Added in macOS 10.3.
// Decrements the retain count of a color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRelease
func CGColorRelease(color ColorRef) {
	_CGColorRelease(color)
}

// Increments the retain count of a color.
//
// Added in macOS 10.3.
// Increments the retain count of a color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRetain
func CGColorRetain(color ColorRef) ColorRef {
	return _CGColorRetain(color)
}

// Returns the base color space of a pattern or indexed color space.
//
// Added in macOS 10.5.
// Returns the base color space of a pattern or indexed color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/baseColorSpace
func CGColorSpaceGetBaseColorSpace(space ColorSpaceRef) ColorSpaceRef {
	return _CGColorSpaceGetBaseColorSpace(space)
}

// Returns a copy of the ICC profile data of the provided color space.
//
// Added in macOS 10.12.
// Returns a copy of the ICC profile data of the provided color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/copyICCData()
func CGColorSpaceCopyICCData(space ColorSpaceRef) DataRef {
	return _CGColorSpaceCopyICCData(space)
}

// Returns a copy of the color space’s properties.
//
// Added in macOS 10.12.
// Returns a copy of the color space’s properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/copyPropertyList()
func CGColorSpaceCopyPropertyList(space ColorSpaceRef) PropertyListRef {
	return _CGColorSpaceCopyPropertyList(space)
}

// Returns a copy of the ICC profile of the provided color space.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.5.
// Returns a copy of the ICC profile of the provided color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/iccData
func CGColorSpaceCopyICCProfile(space ColorSpaceRef) DataRef {
	return _CGColorSpaceCopyICCProfile(space)
}

// Creates a calibrated grayscale color space.
//
// Added in macOS 10.0.
// Creates a calibrated grayscale color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(calibratedGrayWhitePoint:blackPoint:gamma:)
func CGColorSpaceCreateCalibratedGray(whitePoint unsafe.Pointer, blackPoint unsafe.Pointer, gamma unsafe.Pointer) ColorSpaceRef {
	return _CGColorSpaceCreateCalibratedGray(whitePoint, blackPoint, gamma)
}

// Creates a calibrated RGB color space.
//
// Added in macOS 10.0.
// Creates a calibrated RGB color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(calibratedRGBWhitePoint:blackPoint:gamma:matrix:)
func CGColorSpaceCreateCalibratedRGB(whitePoint unsafe.Pointer, blackPoint unsafe.Pointer, gamma unsafe.Pointer, matrix unsafe.Pointer, p4 unsafe.Pointer) ColorSpaceRef {
	return _CGColorSpaceCreateCalibratedRGB(whitePoint, blackPoint, gamma, matrix, p4)
}

// Creates a device-independent color space that is defined according to the ICC color profile specification.
//
// Added in macOS 10.0.
// Creates a device-independent color space that is defined according to the ICC color profile specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(iccBasedNComponents:range:profile:alternate:)
func CGColorSpaceCreateICCBased(nComponents uintptr, range_ []float64, profile DataProviderRef, alternate ColorSpaceRef) ColorSpaceRef {
	return _CGColorSpaceCreateICCBased(nComponents, range_, profile, alternate)
}

// Creates an ICC-based color space using the ICC profile contained in the specified data.
//
// Added in macOS 10.12.
// Creates an ICC-based color space using the ICC profile contained in the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(iccData:)
func CGColorSpaceCreateWithICCData(data TypeRef) ColorSpaceRef {
	return _CGColorSpaceCreateWithICCData(data)
}

// Creates an ICC-based color space using the ICC profile contained in the specified data.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.5.
// Creates an ICC-based color space using the ICC profile contained in the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(iccProfileData:)
func CGColorSpaceCreateWithICCProfile(data DataRef) ColorSpaceRef {
	return _CGColorSpaceCreateWithICCProfile(data)
}

// Creates an indexed color space, consisting of colors specified by a color lookup table.
//
// Added in macOS 10.0.
// Creates an indexed color space, consisting of colors specified by a color lookup table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(indexedBaseSpace:last:colorTable:)
func CGColorSpaceCreateIndexed(baseSpace ColorSpaceRef, lastIndex uintptr, colorTable unsafe.Pointer) ColorSpaceRef {
	return _CGColorSpaceCreateIndexed(baseSpace, lastIndex, colorTable)
}

// Creates a device-independent color space that is relative to human color perception, according to the CIE L*a*b* standard.
//
// Added in macOS 10.0.
// Creates a device-independent color space that is relative to human color perception, according to the CIE L*a*b* standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(labWhitePoint:blackPoint:range:)
func CGColorSpaceCreateLab(whitePoint unsafe.Pointer, blackPoint unsafe.Pointer, range_ unsafe.Pointer, p3 unsafe.Pointer) ColorSpaceRef {
	return _CGColorSpaceCreateLab(whitePoint, blackPoint, range_, p3)
}

// Creates a specified type of Quartz color space.
//
// Added in macOS 10.2.
// Creates a specified type of Quartz color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(name:)
func CGColorSpaceCreateWithName(name StringRef) ColorSpaceRef {
	return _CGColorSpaceCreateWithName(name)
}

// Creates a pattern color space.
//
// Added in macOS 10.0.
// Creates a pattern color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(patternBaseSpace:)
func CGColorSpaceCreatePattern(baseSpace ColorSpaceRef) ColorSpaceRef {
	return _CGColorSpaceCreatePattern(baseSpace)
}

// Creates a platform-specific color space.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
// Creates a platform-specific color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(platformColorSpaceRef:)
func CGColorSpaceCreateWithPlatformColorSpace(ref unsafe.Pointer) ColorSpaceRef {
	return _CGColorSpaceCreateWithPlatformColorSpace(ref)
}

// Creates a color space from a property list.
//
// Added in macOS 10.12.
// Creates a color space from a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(propertyListPlist:)
func CGColorSpaceCreateWithPropertyList(plist PropertyListRef) ColorSpaceRef {
	return _CGColorSpaceCreateWithPropertyList(plist)
}

// CGColorSpaceIsHDR is a CoreGraphics function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/isHDR()
func CGColorSpaceIsHDR(p0 ColorSpaceRef) bool {
	return _CGColorSpaceIsHDR(p0)
}

// Returns whether the RGB color space covers a significant portion of the NTSC color gamut.
//
// Added in macOS 10.12.
// Returns whether the RGB color space covers a significant portion of the NTSC color gamut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/isWideGamutRGB
func CGColorSpaceIsWideGamutRGB(p0 ColorSpaceRef) bool {
	return _CGColorSpaceIsWideGamutRGB(p0)
}

// Returns the color space model of the provided color space.
//
// Added in macOS 10.5.
// Returns the color space model of the provided color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/model
func CGColorSpaceGetModel(space ColorSpaceRef) ColorSpaceModel {
	return _CGColorSpaceGetModel(space)
}

// Returns the name used to create the specified color space.
//
// Added in macOS 10.6.
// Returns the name used to create the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/name
func CGColorSpaceCopyName(space ColorSpaceRef) StringRef {
	return _CGColorSpaceCopyName(space)
}

// Returns the number of color components in a color space.
//
// Added in macOS 10.0.
// Returns the number of color components in a color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/numberOfComponents
func CGColorSpaceGetNumberOfComponents(space ColorSpaceRef) uintptr {
	return _CGColorSpaceGetNumberOfComponents(space)
}

// Returns a Boolean indicating whether the color space can be used as a destination color space.
//
// Added in macOS 10.12.
// Returns a Boolean indicating whether the color space can be used as a destination color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/supportsOutput
func CGColorSpaceSupportsOutput(space ColorSpaceRef) bool {
	return _CGColorSpaceSupportsOutput(space)
}

// Returns the Core Foundation type identifier for Quartz color spaces.
//
// Added in macOS 10.2.
// Returns the Core Foundation type identifier for Quartz color spaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/typeID
func CGColorSpaceGetTypeID() TypeID {
	return _CGColorSpaceGetTypeID()
}

// CGColorSpaceCopyBaseColorSpace is a CoreGraphics function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCopyBaseColorSpace(_:)
func CGColorSpaceCopyBaseColorSpace(space ColorSpaceRef) ColorSpaceRef {
	return _CGColorSpaceCopyBaseColorSpace(space)
}

// CGColorSpaceCreateCopyWithStandardRange is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateCopyWithStandardRange(_:)
func CGColorSpaceCreateCopyWithStandardRange(space ColorSpaceRef) ColorSpaceRef {
	return _CGColorSpaceCreateCopyWithStandardRange(space)
}

// Creates a device-dependent CMYK color space.
//
// Added in macOS 10.0.
// Creates a device-dependent CMYK color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateDeviceCMYK()
func CGColorSpaceCreateDeviceCMYK() ColorSpaceRef {
	return _CGColorSpaceCreateDeviceCMYK()
}

// Creates a device-dependent grayscale color space.
//
// Added in macOS 10.0.
// Creates a device-dependent grayscale color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateDeviceGray()
func CGColorSpaceCreateDeviceGray() ColorSpaceRef {
	return _CGColorSpaceCreateDeviceGray()
}

// Creates a device-dependent RGB color space.
//
// Added in macOS 10.0.
// Creates a device-dependent RGB color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateDeviceRGB()
func CGColorSpaceCreateDeviceRGB() ColorSpaceRef {
	return _CGColorSpaceCreateDeviceRGB()
}

// CGColorSpaceCreateExtended is a CoreGraphics function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateExtended(_:)
func CGColorSpaceCreateExtended(space ColorSpaceRef) ColorSpaceRef {
	return _CGColorSpaceCreateExtended(space)
}

// CGColorSpaceCreateExtendedLinearized is a CoreGraphics function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateExtendedLinearized(_:)
func CGColorSpaceCreateExtendedLinearized(space ColorSpaceRef) ColorSpaceRef {
	return _CGColorSpaceCreateExtendedLinearized(space)
}

// CGColorSpaceCreateLinearized is a CoreGraphics function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateLinearized(_:)
func CGColorSpaceCreateLinearized(space ColorSpaceRef) ColorSpaceRef {
	return _CGColorSpaceCreateLinearized(space)
}

// CGColorSpaceCreateWithColorSyncProfile is a CoreGraphics function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateWithColorSyncProfile(_:_:)
func CGColorSpaceCreateWithColorSyncProfile(p0 ColorSyncProfileRef, options DictionaryRef) ColorSpaceRef {
	return _CGColorSpaceCreateWithColorSyncProfile(p0, options)
}

// Copies the entries in the color table of an indexed color space.
//
// Added in macOS 10.5.
// Copies the entries in the color table of an indexed color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceGetColorTable
func CGColorSpaceGetColorTable(space ColorSpaceRef, table unsafe.Pointer) {
	_CGColorSpaceGetColorTable(space, table)
}

// Returns the number of entries in the color table of an indexed color space.
//
// Added in macOS 10.5.
// Returns the number of entries in the color table of an indexed color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceGetColorTableCount
func CGColorSpaceGetColorTableCount(space ColorSpaceRef) uintptr {
	return _CGColorSpaceGetColorTableCount(space)
}

// CGColorSpaceGetName is a CoreGraphics function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceGetName
func CGColorSpaceGetName(space ColorSpaceRef) StringRef {
	return _CGColorSpaceGetName(space)
}

// CGColorSpaceIsHLGBased is a CoreGraphics function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceIsHLGBased(_:)
func CGColorSpaceIsHLGBased(s ColorSpaceRef) bool {
	return _CGColorSpaceIsHLGBased(s)
}

// CGColorSpaceIsPQBased is a CoreGraphics function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceIsPQBased(_:)
func CGColorSpaceIsPQBased(s ColorSpaceRef) bool {
	return _CGColorSpaceIsPQBased(s)
}

// Decrements the retain count of a color space.
//
// Added in macOS 10.0.
// Decrements the retain count of a color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceRelease
func CGColorSpaceRelease(space ColorSpaceRef) {
	_CGColorSpaceRelease(space)
}

// Increments the retain count of a color space.
//
// Added in macOS 10.0.
// Increments the retain count of a color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceRetain
func CGColorSpaceRetain(space ColorSpaceRef) ColorSpaceRef {
	return _CGColorSpaceRetain(space)
}

// CGColorSpaceUsesExtendedRange is a CoreGraphics function.
//
// Added in macOS 10.12.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceUsesExtendedRange(_:)
func CGColorSpaceUsesExtendedRange(space ColorSpaceRef) bool {
	return _CGColorSpaceUsesExtendedRange(space)
}

// CGColorSpaceUsesITUR_2100TF is a CoreGraphics function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceUsesITUR_2100TF(_:)
func CGColorSpaceUsesITUR_2100TF(p0 ColorSpaceRef) bool {
	return _CGColorSpaceUsesITUR_2100TF(p0)
}

// Completes a set of display configuration changes.
//
// Added in macOS 10.0.
// Completes a set of display configuration changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCompleteDisplayConfiguration(_:_:)
func CGCompleteDisplayConfiguration(config DisplayConfigRef, option ConfigureOption) Error {
	return _CGCompleteDisplayConfiguration(config, option)
}

// Modifies the settings of the built-in fade effect that occurs during a display configuration.
//
// Added in macOS 10.2.
// Modifies the settings of the built-in fade effect that occurs during a display configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayFadeEffect(_:_:_:_:_:_:)
func CGConfigureDisplayFadeEffect(config DisplayConfigRef, fadeOutSeconds DisplayFadeInterval, fadeInSeconds DisplayFadeInterval, fadeRed float32, fadeGreen float32, fadeBlue float32) Error {
	return _CGConfigureDisplayFadeEffect(config, fadeOutSeconds, fadeInSeconds, fadeRed, fadeGreen, fadeBlue)
}

// Changes the configuration of a mirroring set.
//
// Added in macOS 10.2.
// Changes the configuration of a mirroring set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayMirrorOfDisplay(_:_:_:)
func CGConfigureDisplayMirrorOfDisplay(config DisplayConfigRef, display DirectDisplayID, master DirectDisplayID) Error {
	return _CGConfigureDisplayMirrorOfDisplay(config, display, master)
}

// Configures the display mode of a display.

// Configures the display mode of a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayMode(_:_:_:)
func CGConfigureDisplayMode(config DisplayConfigRef, display DirectDisplayID, mode DictionaryRef) Error {
	return _CGConfigureDisplayMode(config, display, mode)
}

// Configures the origin of a display relative to the global display coordinate space.
//
// Added in macOS 10.0.
// Configures the origin of a display relative to the global display coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayOrigin(_:_:_:_:)
func CGConfigureDisplayOrigin(config DisplayConfigRef, display DirectDisplayID, x int32, y int32) Error {
	return _CGConfigureDisplayOrigin(config, display, x, y)
}

// Enables or disables stereo operation for a display, as part of a display configuration.
//
// Added in macOS 10.4.
// Enables or disables stereo operation for a display, as part of a display configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayStereoOperation(_:_:_:_:)
func CGConfigureDisplayStereoOperation(config DisplayConfigRef, display DirectDisplayID, stereo unsafe.Pointer, forceBlueLine unsafe.Pointer) Error {
	return _CGConfigureDisplayStereoOperation(config, display, stereo, forceBlueLine)
}

// Configures the display mode of a display.
//
// Added in macOS 10.6.
// Configures the display mode of a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayWithDisplayMode(_:_:_:_:)
func CGConfigureDisplayWithDisplayMode(config DisplayConfigRef, display DirectDisplayID, mode DisplayModeRef, options DictionaryRef) Error {
	return _CGConfigureDisplayWithDisplayMode(config, display, mode, options)
}

// Sets a destination to jump to when a point in the current page of a PDF graphics context is clicked.
//
// Added in macOS 10.4.
// Sets a destination to jump to when a point in the current page of a PDF graphics context is clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/addDestination(_:at:)
func CGPDFContextAddDestinationAtPoint(context ContextRef, name StringRef, point Point) {
	_CGPDFContextAddDestinationAtPoint(context, name, point)
}

// Associates custom metadata with the PDF document.
//
// Added in macOS 10.7.
// Associates custom metadata with the PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/addDocumentMetadata(_:)
func CGPDFContextAddDocumentMetadata(context ContextRef, metadata DataRef) {
	_CGPDFContextAddDocumentMetadata(context, metadata)
}

// Adds an ellipse that fits inside the specified rectangle.
//
// Added in macOS 10.4.
// Adds an ellipse that fits inside the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/addEllipse(in:)
func CGContextAddEllipseInRect(c ContextRef, rect Rect) {
	_CGContextAddEllipseInRect(c, rect)
}

// Adds a previously created path object to the current path in a graphics context.
//
// Added in macOS 10.2.
// Adds a previously created path object to the current path in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/addPath(_:)
func CGContextAddPath(c ContextRef, path PathRef) {
	_CGContextAddPath(c, path)
}

// Adds a rectangular path to the current path.
//
// Added in macOS 10.0.
// Adds a rectangular path to the current path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/addRect(_:)
func CGContextAddRect(c ContextRef, rect Rect) {
	_CGContextAddRect(c, rect)
}

// Returns the alpha information associated with the context, which indicates how a bitmap context handles the alpha component.
//
// Added in macOS 10.2.
// Returns the alpha information associated with the context, which indicates how a bitmap context handles the alpha component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/alphaInfo
func CGBitmapContextGetAlphaInfo(context ContextRef) ImageAlphaInfo {
	return _CGBitmapContextGetAlphaInfo(context)
}

// Begins a new page in a PDF graphics context.
//
// Added in macOS 10.4.
// Begins a new page in a PDF graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginPDFPage(_:)
func CGPDFContextBeginPage(context ContextRef, pageInfo DictionaryRef) {
	_CGPDFContextBeginPage(context, pageInfo)
}

// Starts a new page in a page-based graphics context.
//
// Added in macOS 10.0.
// Starts a new page in a page-based graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginPage(mediaBox:)
func CGContextBeginPage(c ContextRef, mediaBox unsafe.Pointer) {
	_CGContextBeginPage(c, mediaBox)
}

// Creates a new empty path in a graphics context.
//
// Added in macOS 10.0.
// Creates a new empty path in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginPath()
func CGContextBeginPath(c ContextRef) {
	_CGContextBeginPath(c)
}

// Begins a transparency layer.
//
// Added in macOS 10.3.
// Begins a transparency layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginTransparencyLayer(auxiliaryInfo:)
func CGContextBeginTransparencyLayer(c ContextRef, auxiliaryInfo DictionaryRef) {
	_CGContextBeginTransparencyLayer(c, auxiliaryInfo)
}

// Begins a transparency layer whose contents are bounded by the specified rectangle.
//
// Added in macOS 10.5.
// Begins a transparency layer whose contents are bounded by the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginTransparencyLayer(in:auxiliaryInfo:)
func CGContextBeginTransparencyLayerWithRect(c ContextRef, rect Rect, auxInfo DictionaryRef) {
	_CGContextBeginTransparencyLayerWithRect(c, rect, auxInfo)
}

// Obtains the bitmap information associated with a bitmap graphics context.
//
// Added in macOS 10.4.
// Obtains the bitmap information associated with a bitmap graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/bitmapInfo
func CGBitmapContextGetBitmapInfo(context ContextRef) BitmapInfo {
	return _CGBitmapContextGetBitmapInfo(context)
}

// Returns the bits per component of a bitmap context.
//
// Added in macOS 10.2.
// Returns the bits per component of a bitmap context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/bitsPerComponent
func CGBitmapContextGetBitsPerComponent(context ContextRef) uintptr {
	return _CGBitmapContextGetBitsPerComponent(context)
}

// Returns the bits per pixel of a bitmap context.
//
// Added in macOS 10.2.
// Returns the bits per pixel of a bitmap context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/bitsPerPixel
func CGBitmapContextGetBitsPerPixel(context ContextRef) uintptr {
	return _CGBitmapContextGetBitsPerPixel(context)
}

// Returns the bounding box of a clipping path.
//
// Added in macOS 10.3.
// Returns the bounding box of a clipping path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/boundingBoxOfClipPath
func CGContextGetClipBoundingBox(c ContextRef) Rect {
	return _CGContextGetClipBoundingBox(c)
}

// Returns the smallest rectangle that contains the current path.
//
// Added in macOS 10.0.
// Returns the smallest rectangle that contains the current path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/boundingBoxOfPath
func CGContextGetPathBoundingBox(c ContextRef) Rect {
	return _CGContextGetPathBoundingBox(c)
}

// Returns the bytes per row of a bitmap context.
//
// Added in macOS 10.2.
// Returns the bytes per row of a bitmap context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/bytesPerRow
func CGBitmapContextGetBytesPerRow(context ContextRef) uintptr {
	return _CGBitmapContextGetBytesPerRow(context)
}

// Paints a transparent rectangle.
//
// Added in macOS 10.0.
// Paints a transparent rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/clear(_:)
func CGContextClearRect(c ContextRef, rect Rect) {
	_CGContextClearRect(c, rect)
}

// Sets the clipping path to the intersection of the current clipping path with the area defined by the specified rectangle.
//
// Added in macOS 10.0.
// Sets the clipping path to the intersection of the current clipping path with the area defined by the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/clip(to:)-7cbwq
func CGContextClipToRect(c ContextRef, rect Rect) {
	_CGContextClipToRect(c, rect)
}

// Maps a mask into the specified rectangle and intersects it with the current clipping area of the graphics context.
//
// Added in macOS 10.4.
// Maps a mask into the specified rectangle and intersects it with the current clipping area of the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/clip(to:mask:)
func CGContextClipToMask(c ContextRef, rect Rect, mask ImageRef) {
	_CGContextClipToMask(c, rect, mask)
}

// Closes a PDF document.
//
// Added in macOS 10.5.
// Closes a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/closePDF()
func CGPDFContextClose(context ContextRef) {
	_CGPDFContextClose(context)
}

// Closes and terminates the current path’s subpath.
//
// Added in macOS 10.0.
// Closes and terminates the current path’s subpath.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/closePath()
func CGContextClosePath(c ContextRef) {
	_CGContextClosePath(c)
}

// Returns the color space of a bitmap context.
//
// Added in macOS 10.2.
// Returns the color space of a bitmap context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/colorSpace
func CGBitmapContextGetColorSpace(context ContextRef) ColorSpaceRef {
	return _CGBitmapContextGetColorSpace(context)
}

// Transforms the user coordinate system in a context using a specified matrix.
//
// Added in macOS 10.0.
// Transforms the user coordinate system in a context using a specified matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/concatenate(_:)
func CGContextConcatCTM(c ContextRef, transform AffineTransform) {
	_CGContextConcatCTM(c, transform)
}

// Returns a size that is transformed from user space coordinates to device space coordinates.
//
// Added in macOS 10.4.
// Returns a size that is transformed from user space coordinates to device space coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToDeviceSpace(_:)-224h2
func CGContextConvertSizeToDeviceSpace(c ContextRef, size Size) Size {
	return _CGContextConvertSizeToDeviceSpace(c, size)
}

// Returns a point that is transformed from user space coordinates to device space coordinates.
//
// Added in macOS 10.4.
// Returns a point that is transformed from user space coordinates to device space coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToDeviceSpace(_:)-53m7u
func CGContextConvertPointToDeviceSpace(c ContextRef, point Point) Point {
	return _CGContextConvertPointToDeviceSpace(c, point)
}

// Returns a rectangle that is transformed from user space coordinate to device space coordinates.
//
// Added in macOS 10.4.
// Returns a rectangle that is transformed from user space coordinate to device space coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToDeviceSpace(_:)-91x5g
func CGContextConvertRectToDeviceSpace(c ContextRef, rect Rect) Rect {
	return _CGContextConvertRectToDeviceSpace(c, rect)
}

// Returns a rectangle that is transformed from device space coordinate to user space coordinates.
//
// Added in macOS 10.4.
// Returns a rectangle that is transformed from device space coordinate to user space coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToUserSpace(_:)-1hk5r
func CGContextConvertRectToUserSpace(c ContextRef, rect Rect) Rect {
	return _CGContextConvertRectToUserSpace(c, rect)
}

// Returns a point that is transformed from device space coordinates to user space coordinates.
//
// Added in macOS 10.4.
// Returns a point that is transformed from device space coordinates to user space coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToUserSpace(_:)-3mtg3
func CGContextConvertPointToUserSpace(c ContextRef, point Point) Point {
	return _CGContextConvertPointToUserSpace(c, point)
}

// Returns a size that is transformed from device space coordinates to user space coordinates.
//
// Added in macOS 10.4.
// Returns a size that is transformed from device space coordinates to user space coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToUserSpace(_:)-693ur
func CGContextConvertSizeToUserSpace(c ContextRef, size Size) Size {
	return _CGContextConvertSizeToUserSpace(c, size)
}

// Returns the current transformation matrix.
//
// Added in macOS 10.0.
// Returns the current transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/ctm
func CGContextGetCTM(c ContextRef) AffineTransform {
	return _CGContextGetCTM(c)
}

// Returns the current point in a non-empty path.
//
// Added in macOS 10.0.
// Returns the current point in a non-empty path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/currentPointOfPath
func CGContextGetPathCurrentPoint(c ContextRef) Point {
	return _CGContextGetPathCurrentPoint(c)
}

// Returns a pointer to the image data associated with a bitmap context.
//
// Added in macOS 10.2.
// Returns a pointer to the image data associated with a bitmap context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/data
func CGBitmapContextGetData(context ContextRef) unsafe.Pointer {
	return _CGBitmapContextGetData(context)
}

// Paints a gradient fill that varies along the line defined by the provided starting and ending points.
//
// Added in macOS 10.5.
// Paints a gradient fill that varies along the line defined by the provided starting and ending points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawLinearGradient(_:start:end:options:)
func CGContextDrawLinearGradient(c ContextRef, gradient GradientRef, startPoint Point, endPoint Point, options GradientDrawingOptions) {
	_CGContextDrawLinearGradient(c, gradient, startPoint, endPoint, options)
}

// Draws the content of a PDF page into the current graphics context.
//
// Added in macOS 10.3.
// Draws the content of a PDF page into the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawPDFPage(_:)
func CGContextDrawPDFPage(c ContextRef, page PDFPageRef) {
	_CGContextDrawPDFPage(c, page)
}

// Draws the current path using the provided drawing mode.
//
// Added in macOS 10.0.
// Draws the current path using the provided drawing mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawPath(using:)
func CGContextDrawPath(c ContextRef, mode PathDrawingMode) {
	_CGContextDrawPath(c, mode)
}

// Paints a gradient fill that varies along the area defined by the provided starting and ending circles.
//
// Added in macOS 10.5.
// Paints a gradient fill that varies along the area defined by the provided starting and ending circles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawRadialGradient(_:startCenter:startRadius:endCenter:endRadius:options:)
func CGContextDrawRadialGradient(c ContextRef, gradient GradientRef, startCenter Point, startRadius Float, endCenter Point, endRadius Float, options GradientDrawingOptions) {
	_CGContextDrawRadialGradient(c, gradient, startCenter, startRadius, endCenter, endRadius, options)
}

// Fills the clipping path of a context with the specified shading.
//
// Added in macOS 10.2.
// Fills the clipping path of a context with the specified shading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawShading(_:)
func CGContextDrawShading(c ContextRef, shading ShadingRef) {
	_CGContextDrawShading(c, shading)
}

// Ends the current page in the PDF graphics context.
//
// Added in macOS 10.4.
// Ends the current page in the PDF graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/endPDFPage()
func CGPDFContextEndPage(context ContextRef) {
	_CGPDFContextEndPage(context)
}

// Ends the current page in a page-based graphics context.
//
// Added in macOS 10.0.
// Ends the current page in a page-based graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/endPage()
func CGContextEndPage(c ContextRef) {
	_CGContextEndPage(c)
}

// Ends a transparency layer.
//
// Added in macOS 10.3.
// Ends a transparency layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/endTransparencyLayer()
func CGContextEndTransparencyLayer(c ContextRef) {
	_CGContextEndTransparencyLayer(c)
}

// Paints the area contained within the provided rectangle, using the fill color in the current graphics state.
//
// Added in macOS 10.0.
// Paints the area contained within the provided rectangle, using the fill color in the current graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/fill(_:)-7a0rk
func CGContextFillRect(c ContextRef, rect Rect) {
	_CGContextFillRect(c, rect)
}

// Paints the area of the ellipse that fits inside the provided rectangle, using the fill color in the current graphics state.
//
// Added in macOS 10.4.
// Paints the area of the ellipse that fits inside the provided rectangle, using the fill color in the current graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/fillEllipse(in:)
func CGContextFillEllipseInRect(c ContextRef, rect Rect) {
	_CGContextFillEllipseInRect(c, rect)
}

// Forces all pending drawing operations in a window context to be rendered immediately to the destination device.
//
// Added in macOS 10.0.
// Forces all pending drawing operations in a window context to be rendered immediately to the destination device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/flush()
func CGContextFlush(c ContextRef) {
	_CGContextFlush(c)
}

// Returns the height in pixels of a bitmap context.
//
// Added in macOS 10.2.
// Returns the height in pixels of a bitmap context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/height
func CGBitmapContextGetHeight(context ContextRef) uintptr {
	return _CGBitmapContextGetHeight(context)
}

// Creates a URL-based PDF graphics context.
//
// Added in macOS 10.0.
// Creates a URL-based PDF graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(_:mediaBox:_:)
func CGPDFContextCreateWithURL(url URLRef, mediaBox unsafe.Pointer, auxiliaryInfo DictionaryRef) ContextRef {
	return _CGPDFContextCreateWithURL(url, mediaBox, auxiliaryInfo)
}

// Creates a PDF graphics context.
//
// Added in macOS 10.0.
// Creates a PDF graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(consumer:mediaBox:_:)
func CGPDFContextCreate(consumer DataConsumerRef, mediaBox unsafe.Pointer, auxiliaryInfo DictionaryRef) ContextRef {
	return _CGPDFContextCreate(consumer, mediaBox, auxiliaryInfo)
}

// CGBitmapContextCreate is a CoreGraphics function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(data:width:height:bitsPerComponent:bytesPerRow:space:bitmapInfo:)-10b3i
func CGBitmapContextCreate(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space ColorSpaceRef, bitmapInfo BitmapInfo) ContextRef {
	return _CGBitmapContextCreate(data, width, height, bitsPerComponent, bytesPerRow, space, bitmapInfo)
}

// CGBitmapContextCreateWithData is a CoreGraphics function.
//
// Added in macOS 10.6.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(data:width:height:bitsPerComponent:bytesPerRow:space:bitmapInfo:releaseCallback:releaseInfo:)-4yzt5
func CGBitmapContextCreateWithData(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space ColorSpaceRef, bitmapInfo BitmapInfo, releaseCallback BitmapContextReleaseDataCallback, releaseInfo unsafe.Pointer) ContextRef {
	return _CGBitmapContextCreateWithData(data, width, height, bitsPerComponent, bytesPerRow, space, bitmapInfo, releaseCallback, releaseInfo)
}

// Returns the current level of interpolation quality for a graphics context.
//
// Added in macOS 10.0.
// Returns the current level of interpolation quality for a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/interpolationQuality
func CGContextGetInterpolationQuality(c ContextRef) InterpolationQuality {
	return _CGContextGetInterpolationQuality(c)
}

// Indicates whether the current path contains any subpaths.
//
// Added in macOS 10.0.
// Indicates whether the current path contains any subpaths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/isPathEmpty
func CGContextIsPathEmpty(c ContextRef) bool {
	return _CGContextIsPathEmpty(c)
}

// Creates and returns a CGImage from the pixel data in a bitmap graphics context.
//
// Added in macOS 10.4.
// Creates and returns a CGImage from the pixel data in a bitmap graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/makeImage()
func CGBitmapContextCreateImage(context ContextRef) ImageRef {
	return _CGBitmapContextCreateImage(context)
}

// Returns a path object built from the current path information in a graphics context.
//
// Added in macOS 10.2.
// Returns a path object built from the current path information in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/path
func CGContextCopyPath(c ContextRef) PathRef {
	return _CGContextCopyPath(c)
}

// Checks to see whether the specified point is contained in the current path.
//
// Added in macOS 10.4.
// Checks to see whether the specified point is contained in the current path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/pathContains(_:mode:)
func CGContextPathContainsPoint(c ContextRef, point Point, mode PathDrawingMode) bool {
	return _CGContextPathContainsPoint(c, point, mode)
}

// Replaces the path in the graphics context with the stroked version of the path.
//
// Added in macOS 10.4.
// Replaces the path in the graphics context with the stroked version of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/replacePathWithStrokedPath()
func CGContextReplacePathWithStrokedPath(c ContextRef) {
	_CGContextReplacePathWithStrokedPath(c)
}

// CGContextResetClip is a CoreGraphics function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/resetClip()
func CGContextResetClip(c ContextRef) {
	_CGContextResetClip(c)
}

// Sets the current graphics state to the state most recently saved.
//
// Added in macOS 10.0.
// Sets the current graphics state to the state most recently saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/restoreGState()
func CGContextRestoreGState(c ContextRef) {
	_CGContextRestoreGState(c)
}

// Rotates the user coordinate system in a context.
//
// Added in macOS 10.0.
// Rotates the user coordinate system in a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/rotate(by:)
func CGContextRotateCTM(c ContextRef, angle Float) {
	_CGContextRotateCTM(c, angle)
}

// Pushes a copy of the current graphics state onto the graphics state stack for the context.
//
// Added in macOS 10.0.
// Pushes a copy of the current graphics state onto the graphics state stack for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/saveGState()
func CGContextSaveGState(c ContextRef) {
	_CGContextSaveGState(c)
}

// Changes the scale of the user coordinate system in a context.
//
// Added in macOS 10.0.
// Changes the scale of the user coordinate system in a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/scaleBy(x:y:)
func CGContextScaleCTM(c ContextRef, sx Float, sy Float) {
	_CGContextScaleCTM(c, sx, sy)
}

// Sets the font and font size in a graphics context.

// Sets the font and font size in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/selectFont(name:size:textEncoding:)
func CGContextSelectFont(c ContextRef, name unsafe.Pointer, size Float, textEncoding TextEncoding) {
	_CGContextSelectFont(c, name, size, textEncoding)
}

// Sets whether or not to allow antialiasing for a graphics context.
//
// Added in macOS 10.4.
// Sets whether or not to allow antialiasing for a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsAntialiasing(_:)
func CGContextSetAllowsAntialiasing(c ContextRef, allowsAntialiasing bool) {
	_CGContextSetAllowsAntialiasing(c, allowsAntialiasing)
}

// Sets whether or not to allow font smoothing for a graphics context.
//
// Added in macOS 10.2.
// Sets whether or not to allow font smoothing for a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsFontSmoothing(_:)
func CGContextSetAllowsFontSmoothing(c ContextRef, allowsFontSmoothing bool) {
	_CGContextSetAllowsFontSmoothing(c, allowsFontSmoothing)
}

// Sets whether or not to allow subpixel positioning for a graphics context.
//
// Added in macOS 10.5.
// Sets whether or not to allow subpixel positioning for a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsFontSubpixelPositioning(_:)
func CGContextSetAllowsFontSubpixelPositioning(c ContextRef, allowsFontSubpixelPositioning bool) {
	_CGContextSetAllowsFontSubpixelPositioning(c, allowsFontSubpixelPositioning)
}

// Sets whether or not to allow subpixel quantization for a graphics context.
//
// Added in macOS 10.5.
// Sets whether or not to allow subpixel quantization for a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsFontSubpixelQuantization(_:)
func CGContextSetAllowsFontSubpixelQuantization(c ContextRef, allowsFontSubpixelQuantization bool) {
	_CGContextSetAllowsFontSubpixelQuantization(c, allowsFontSubpixelQuantization)
}

// Sets the opacity level for objects drawn in a graphics context.
//
// Added in macOS 10.0.
// Sets the opacity level for objects drawn in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAlpha(_:)
func CGContextSetAlpha(c ContextRef, alpha Float) {
	_CGContextSetAlpha(c, alpha)
}

// Sets how sample values are composited by a graphics context.
//
// Added in macOS 10.4.
// Sets how sample values are composited by a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setBlendMode(_:)
func CGContextSetBlendMode(c ContextRef, mode BlendMode) {
	_CGContextSetBlendMode(c, mode)
}

// Sets the current character spacing.
//
// Added in macOS 10.0.
// Sets the current character spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setCharacterSpacing(_:)
func CGContextSetCharacterSpacing(c ContextRef, spacing Float) {
	_CGContextSetCharacterSpacing(c, spacing)
}

// Sets a destination to jump to when a rectangle in the current PDF page is clicked.
//
// Added in macOS 10.4.
// Sets a destination to jump to when a rectangle in the current PDF page is clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setDestination(_:for:)
func CGPDFContextSetDestinationForRect(context ContextRef, name StringRef, rect Rect) {
	_CGPDFContextSetDestinationForRect(context, name, rect)
}

// CGContextSetEDRTargetHeadroom is a CoreGraphics function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setEDRTargetHeadroom(_:)
func CGContextSetEDRTargetHeadroom(c ContextRef, headroom float32) bool {
	return _CGContextSetEDRTargetHeadroom(c, headroom)
}

// Sets the current fill color.
//
// Added in macOS 10.0.
// Sets the current fill color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(_:)-756dy
func CGContextSetFillColor(c ContextRef, components []float64) {
	_CGContextSetFillColor(c, components)
}

// Sets the current fill color in a graphics context, using a CGColor.
//
// Added in macOS 10.3.
// Sets the current fill color in a graphics context, using a CGColor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(_:)-8lhn8
func CGContextSetFillColorWithColor(c ContextRef, color ColorRef) {
	_CGContextSetFillColorWithColor(c, color)
}

// Sets the current fill color to a value in the DeviceCMYK color space.
//
// Added in macOS 10.0.
// Sets the current fill color to a value in the DeviceCMYK color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(cyan:magenta:yellow:black:alpha:)
func CGContextSetCMYKFillColor(c ContextRef, cyan Float, magenta Float, yellow Float, black Float, alpha Float) {
	_CGContextSetCMYKFillColor(c, cyan, magenta, yellow, black, alpha)
}

// Sets the current fill color to a value in the DeviceGray color space.
//
// Added in macOS 10.0.
// Sets the current fill color to a value in the DeviceGray color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(gray:alpha:)
func CGContextSetGrayFillColor(c ContextRef, gray Float, alpha Float) {
	_CGContextSetGrayFillColor(c, gray, alpha)
}

// Sets the current fill color to a value in the DeviceRGB color space.
//
// Added in macOS 10.0.
// Sets the current fill color to a value in the DeviceRGB color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(red:green:blue:alpha:)
func CGContextSetRGBFillColor(c ContextRef, red Float, green Float, blue Float, alpha Float) {
	_CGContextSetRGBFillColor(c, red, green, blue, alpha)
}

// Sets the fill color space in a graphics context.
//
// Added in macOS 10.0.
// Sets the fill color space in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColorSpace(_:)
func CGContextSetFillColorSpace(c ContextRef, space ColorSpaceRef) {
	_CGContextSetFillColorSpace(c, space)
}

// Sets the fill pattern in the specified graphics context.
//
// Added in macOS 10.0.
// Sets the fill pattern in the specified graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillPattern(_:colorComponents:)
func CGContextSetFillPattern(c ContextRef, pattern PatternRef, components []float64) {
	_CGContextSetFillPattern(c, pattern, components)
}

// Sets the accuracy of curved paths in a graphics context.
//
// Added in macOS 10.0.
// Sets the accuracy of curved paths in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFlatness(_:)
func CGContextSetFlatness(c ContextRef, flatness Float) {
	_CGContextSetFlatness(c, flatness)
}

// Sets the platform font in a graphics context.
//
// Added in macOS 10.0.
// Sets the platform font in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFont(_:)
func CGContextSetFont(c ContextRef, font FontRef) {
	_CGContextSetFont(c, font)
}

// Sets the current font size.
//
// Added in macOS 10.0.
// Sets the current font size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFontSize(_:)
func CGContextSetFontSize(c ContextRef, size Float) {
	_CGContextSetFontSize(c, size)
}

// Sets the style for the endpoints of lines drawn in a graphics context.
//
// Added in macOS 10.0.
// Sets the style for the endpoints of lines drawn in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setLineCap(_:)
func CGContextSetLineCap(c ContextRef, cap_ LineCap) {
	_CGContextSetLineCap(c, cap_)
}

// Sets the style for the joins of connected lines in a graphics context.
//
// Added in macOS 10.0.
// Sets the style for the joins of connected lines in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setLineJoin(_:)
func CGContextSetLineJoin(c ContextRef, join LineJoin) {
	_CGContextSetLineJoin(c, join)
}

// Sets the line width for a graphics context.
//
// Added in macOS 10.0.
// Sets the line width for a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setLineWidth(_:)
func CGContextSetLineWidth(c ContextRef, width Float) {
	_CGContextSetLineWidth(c, width)
}

// Sets the miter limit for the joins of connected lines in a graphics context.
//
// Added in macOS 10.0.
// Sets the miter limit for the joins of connected lines in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setMiterLimit(_:)
func CGContextSetMiterLimit(c ContextRef, limit Float) {
	_CGContextSetMiterLimit(c, limit)
}

// Sets the pattern phase of a context.
//
// Added in macOS 10.0.
// Sets the pattern phase of a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setPatternPhase(_:)
func CGContextSetPatternPhase(c ContextRef, phase Size) {
	_CGContextSetPatternPhase(c, phase)
}

// Sets the rendering intent in the current graphics state.
//
// Added in macOS 10.0.
// Sets the rendering intent in the current graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setRenderingIntent(_:)
func CGContextSetRenderingIntent(c ContextRef, intent ColorRenderingIntent) {
	_CGContextSetRenderingIntent(c, intent)
}

// Enables shadowing in a graphics context.
//
// Added in macOS 10.3.
// Enables shadowing in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShadow(offset:blur:)
func CGContextSetShadow(c ContextRef, offset Size, blur Float) {
	_CGContextSetShadow(c, offset, blur)
}

// Enables shadowing with color a graphics context.
//
// Added in macOS 10.3.
// Enables shadowing with color a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShadow(offset:blur:color:)
func CGContextSetShadowWithColor(c ContextRef, offset Size, blur Float, color ColorRef) {
	_CGContextSetShadowWithColor(c, offset, blur, color)
}

// Sets antialiasing on or off for a graphics context.
//
// Added in macOS 10.0.
// Sets antialiasing on or off for a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldAntialias(_:)
func CGContextSetShouldAntialias(c ContextRef, shouldAntialias bool) {
	_CGContextSetShouldAntialias(c, shouldAntialias)
}

// Enables or disables font smoothing in a graphics context.
//
// Added in macOS 10.2.
// Enables or disables font smoothing in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldSmoothFonts(_:)
func CGContextSetShouldSmoothFonts(c ContextRef, shouldSmoothFonts bool) {
	_CGContextSetShouldSmoothFonts(c, shouldSmoothFonts)
}

// Enables or disables subpixel positioning in a graphics context.
//
// Added in macOS 10.5.
// Enables or disables subpixel positioning in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldSubpixelPositionFonts(_:)
func CGContextSetShouldSubpixelPositionFonts(c ContextRef, shouldSubpixelPositionFonts bool) {
	_CGContextSetShouldSubpixelPositionFonts(c, shouldSubpixelPositionFonts)
}

// Enables or disables subpixel quantization in a graphics context.
//
// Added in macOS 10.5.
// Enables or disables subpixel quantization in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldSubpixelQuantizeFonts(_:)
func CGContextSetShouldSubpixelQuantizeFonts(c ContextRef, shouldSubpixelQuantizeFonts bool) {
	_CGContextSetShouldSubpixelQuantizeFonts(c, shouldSubpixelQuantizeFonts)
}

// Sets the current stroke color in a context, using a CGColor.
//
// Added in macOS 10.3.
// Sets the current stroke color in a context, using a CGColor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(_:)-1sskg
func CGContextSetStrokeColorWithColor(c ContextRef, color ColorRef) {
	_CGContextSetStrokeColorWithColor(c, color)
}

// Sets the current stroke color.
//
// Added in macOS 10.0.
// Sets the current stroke color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(_:)-4pd8p
func CGContextSetStrokeColor(c ContextRef, components []float64) {
	_CGContextSetStrokeColor(c, components)
}

// Sets the current stroke color to a value in the DeviceCMYK color space.
//
// Added in macOS 10.0.
// Sets the current stroke color to a value in the DeviceCMYK color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(cyan:magenta:yellow:black:alpha:)
func CGContextSetCMYKStrokeColor(c ContextRef, cyan Float, magenta Float, yellow Float, black Float, alpha Float) {
	_CGContextSetCMYKStrokeColor(c, cyan, magenta, yellow, black, alpha)
}

// Sets the current stroke color to a value in the DeviceGray color space.
//
// Added in macOS 10.0.
// Sets the current stroke color to a value in the DeviceGray color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(gray:alpha:)
func CGContextSetGrayStrokeColor(c ContextRef, gray Float, alpha Float) {
	_CGContextSetGrayStrokeColor(c, gray, alpha)
}

// Sets the current stroke color to a value in the DeviceRGB color space.
//
// Added in macOS 10.0.
// Sets the current stroke color to a value in the DeviceRGB color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(red:green:blue:alpha:)
func CGContextSetRGBStrokeColor(c ContextRef, red Float, green Float, blue Float, alpha Float) {
	_CGContextSetRGBStrokeColor(c, red, green, blue, alpha)
}

// Sets the stroke color space in a graphics context.
//
// Added in macOS 10.0.
// Sets the stroke color space in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColorSpace(_:)
func CGContextSetStrokeColorSpace(c ContextRef, space ColorSpaceRef) {
	_CGContextSetStrokeColorSpace(c, space)
}

// Sets the stroke pattern in the specified graphics context.
//
// Added in macOS 10.0.
// Sets the stroke pattern in the specified graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokePattern(_:colorComponents:)
func CGContextSetStrokePattern(c ContextRef, pattern PatternRef, components []float64) {
	_CGContextSetStrokePattern(c, pattern, components)
}

// Sets the current text drawing mode.
//
// Added in macOS 10.0.
// Sets the current text drawing mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setTextDrawingMode(_:)
func CGContextSetTextDrawingMode(c ContextRef, mode TextDrawingMode) {
	_CGContextSetTextDrawingMode(c, mode)
}

// Sets the URL associated with a rectangle in a PDF graphics context.
//
// Added in macOS 10.4.
// Sets the URL associated with a rectangle in a PDF graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/setURL(_:for:)
func CGPDFContextSetURLForRect(context ContextRef, url URLRef, rect Rect) {
	_CGPDFContextSetURLForRect(context, url, rect)
}

// Displays an array of glyphs at the current text position.

// Displays an array of glyphs at the current text position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/showGlyphs(g:count:)
func CGContextShowGlyphs(c ContextRef, g unsafe.Pointer, count uintptr) {
	_CGContextShowGlyphs(c, g, count)
}

// Displays an array of glyphs at a position you specify.

// Displays an array of glyphs at a position you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/showGlyphsAtPoint(x:y:glyphs:count:)
func CGContextShowGlyphsAtPoint(c ContextRef, x Float, y Float, glyphs unsafe.Pointer, count uintptr) {
	_CGContextShowGlyphsAtPoint(c, x, y, glyphs, count)
}

// Draws an array of glyphs with varying offsets.

// Draws an array of glyphs with varying offsets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/showGlyphsWithAdvances(glyphs:advances:count:)
func CGContextShowGlyphsWithAdvances(c ContextRef, glyphs unsafe.Pointer, advances unsafe.Pointer, count uintptr) {
	_CGContextShowGlyphsWithAdvances(c, glyphs, advances, count)
}

// Displays a character array at the current text position, a point specified by the current text matrix.

// Displays a character array at the current text position, a point specified by the current text matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/showText(string:length:)
func CGContextShowText(c ContextRef, string_ unsafe.Pointer, length uintptr) {
	_CGContextShowText(c, string_, length)
}

// Displays a character string at a position you specify.

// Displays a character string at a position you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/showTextAtPoint(x:y:string:length:)
func CGContextShowTextAtPoint(c ContextRef, x Float, y Float, string_ unsafe.Pointer, length uintptr) {
	_CGContextShowTextAtPoint(c, x, y, string_, length)
}

// Paints a rectangular path.
//
// Added in macOS 10.0.
// Paints a rectangular path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/stroke(_:)
func CGContextStrokeRect(c ContextRef, rect Rect) {
	_CGContextStrokeRect(c, rect)
}

// Paints a rectangular path, using the specified line width.
//
// Added in macOS 10.0.
// Paints a rectangular path, using the specified line width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/stroke(_:width:)
func CGContextStrokeRectWithWidth(c ContextRef, rect Rect, width Float) {
	_CGContextStrokeRectWithWidth(c, rect, width)
}

// Strokes an ellipse that fits inside the specified rectangle.
//
// Added in macOS 10.4.
// Strokes an ellipse that fits inside the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/strokeEllipse(in:)
func CGContextStrokeEllipseInRect(c ContextRef, rect Rect) {
	_CGContextStrokeEllipseInRect(c, rect)
}

// Paints a line along the current path.
//
// Added in macOS 10.0.
// Paints a line along the current path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/strokePath()
func CGContextStrokePath(c ContextRef) {
	_CGContextStrokePath(c)
}

// Marks a window context for update.
//
// Added in macOS 10.0.
// Marks a window context for update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/synchronize()
func CGContextSynchronize(c ContextRef) {
	_CGContextSynchronize(c)
}

// CGContextSynchronizeAttributes is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/synchronizeAttributes()
func CGContextSynchronizeAttributes(c ContextRef) {
	_CGContextSynchronizeAttributes(c)
}

// Returns the current text matrix.
//
// Added in macOS 10.0.
// Returns the current text matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/textMatrix
func CGContextGetTextMatrix(c ContextRef) AffineTransform {
	return _CGContextGetTextMatrix(c)
}

// Changes the origin of the user coordinate system in a context.
//
// Added in macOS 10.0.
// Changes the origin of the user coordinate system in a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/translateBy(x:y:)
func CGContextTranslateCTM(c ContextRef, tx Float, ty Float) {
	_CGContextTranslateCTM(c, tx, ty)
}

// Returns the type identifier for a graphics context.
//
// Added in macOS 10.2.
// Returns the type identifier for a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/typeID
func CGContextGetTypeID() TypeID {
	return _CGContextGetTypeID()
}

// Returns an affine transform that maps user space coordinates to device space coordinates.
//
// Added in macOS 10.4.
// Returns an affine transform that maps user space coordinates to device space coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/userSpaceToDeviceSpaceTransform
func CGContextGetUserSpaceToDeviceSpaceTransform(c ContextRef) AffineTransform {
	return _CGContextGetUserSpaceToDeviceSpaceTransform(c)
}

// Returns the width in pixels of a bitmap context.
//
// Added in macOS 10.2.
// Returns the width in pixels of a bitmap context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContext/width
func CGBitmapContextGetWidth(context ContextRef) uintptr {
	return _CGBitmapContextGetWidth(context)
}

// Adds an arc of a circle to the current path, possibly preceded by a straight line segment
//
// Added in macOS 10.0.
// Adds an arc of a circle to the current path, possibly preceded by a straight line segment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextAddArc
func CGContextAddArc(c ContextRef, x Float, y Float, radius Float, startAngle Float, endAngle Float, clockwise int) {
	_CGContextAddArc(c, x, y, radius, startAngle, endAngle, clockwise)
}

// Adds an arc of a circle to the current path, using a radius and tangent points.
//
// Added in macOS 10.0.
// Adds an arc of a circle to the current path, using a radius and tangent points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextAddArcToPoint
func CGContextAddArcToPoint(c ContextRef, x1 Float, y1 Float, x2 Float, y2 Float, radius Float) {
	_CGContextAddArcToPoint(c, x1, y1, x2, y2, radius)
}

// Appends a cubic Bézier curve from the current point, using the provided control points and end point .
//
// Added in macOS 10.0.
// Appends a cubic Bézier curve from the current point, using the provided control points and end point .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextAddCurveToPoint
func CGContextAddCurveToPoint(c ContextRef, cp1x Float, cp1y Float, cp2x Float, cp2y Float, x Float, y Float) {
	_CGContextAddCurveToPoint(c, cp1x, cp1y, cp2x, cp2y, x, y)
}

// Appends a straight line segment from the current point to the provided point .
//
// Added in macOS 10.0.
// Appends a straight line segment from the current point to the provided point .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextAddLineToPoint
func CGContextAddLineToPoint(c ContextRef, x Float, y Float) {
	_CGContextAddLineToPoint(c, x, y)
}

// Adds a sequence of connected straight-line segments to the current path.
//
// Added in macOS 10.0.
// Adds a sequence of connected straight-line segments to the current path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextAddLines
func CGContextAddLines(c ContextRef, points unsafe.Pointer, count uintptr) {
	_CGContextAddLines(c, points, count)
}

// Appends a quadratic Bézier curve from the current point, using a control point and an end point you specify.
//
// Added in macOS 10.0.
// Appends a quadratic Bézier curve from the current point, using a control point and an end point you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextAddQuadCurveToPoint
func CGContextAddQuadCurveToPoint(c ContextRef, cpx Float, cpy Float, x Float, y Float) {
	_CGContextAddQuadCurveToPoint(c, cpx, cpy, x, y)
}

// Adds a set of rectangular paths to the current path.
//
// Added in macOS 10.0.
// Adds a set of rectangular paths to the current path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextAddRects
func CGContextAddRects(c ContextRef, rects unsafe.Pointer, count uintptr) {
	_CGContextAddRects(c, rects, count)
}

// Modifies the current clipping path, using the nonzero winding number rule.
//
// Added in macOS 10.0.
// Modifies the current clipping path, using the nonzero winding number rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextClip
func CGContextClip(c ContextRef) {
	_CGContextClip(c)
}

// Sets the clipping path to the intersection of the current clipping path with the region defined by an array of rectangles.
//
// Added in macOS 10.0.
// Sets the clipping path to the intersection of the current clipping path with the region defined by an array of rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextClipToRects
func CGContextClipToRects(c ContextRef, rects unsafe.Pointer, count uintptr) {
	_CGContextClipToRects(c, rects, count)
}

// CGContextDrawConicGradient is a CoreGraphics function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawConicGradient(_:_:_:_:)
func CGContextDrawConicGradient(c ContextRef, gradient GradientRef, center Point, angle Float) {
	_CGContextDrawConicGradient(c, gradient, center, angle)
}

// Draws an image into a graphics context.
//
// Added in macOS 10.0.
// Draws an image into a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawImage
func CGContextDrawImage(c ContextRef, rect Rect, image ImageRef) {
	_CGContextDrawImage(c, rect, image)
}

// CGContextDrawImageApplyingToneMapping is a CoreGraphics function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawImageApplyingToneMapping
func CGContextDrawImageApplyingToneMapping(c ContextRef, r Rect, image ImageRef, method ToneMapping, options DictionaryRef) bool {
	return _CGContextDrawImageApplyingToneMapping(c, r, image, method, options)
}

// Draws the contents of a CGLayer object at the specified point.
//
// Added in macOS 10.4.
// Draws the contents of a CGLayer object at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawLayerAtPoint
func CGContextDrawLayerAtPoint(context ContextRef, point Point, layer LayerRef) {
	_CGContextDrawLayerAtPoint(context, point, layer)
}

// Draws the contents of a layer object into the specified rectangle.
//
// Added in macOS 10.4.
// Draws the contents of a layer object into the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawLayerInRect
func CGContextDrawLayerInRect(context ContextRef, rect Rect, layer LayerRef) {
	_CGContextDrawLayerInRect(context, rect, layer)
}

// CGContextDrawPDFDocument is a CoreGraphics function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawPDFDocument
func CGContextDrawPDFDocument(c ContextRef, rect Rect, document PDFDocumentRef, page int) {
	_CGContextDrawPDFDocument(c, rect, document, page)
}

// Repeatedly draws an image, scaled to the provided rectangle, to fill the current clip region.
//
// Added in macOS 10.5.
// Repeatedly draws an image, scaled to the provided rectangle, to fill the current clip region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawTiledImage
func CGContextDrawTiledImage(c ContextRef, rect Rect, image ImageRef) {
	_CGContextDrawTiledImage(c, rect, image)
}

// Modifies the current clipping path, using the even-odd rule.
//
// Added in macOS 10.0.
// Modifies the current clipping path, using the even-odd rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextEOClip
func CGContextEOClip(c ContextRef) {
	_CGContextEOClip(c)
}

// Paints the area within the current path, using the even-odd fill rule.
//
// Added in macOS 10.0.
// Paints the area within the current path, using the even-odd fill rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextEOFillPath
func CGContextEOFillPath(c ContextRef) {
	_CGContextEOFillPath(c)
}

// Paints the area within the current path, using the nonzero winding number rule.
//
// Added in macOS 10.0.
// Paints the area within the current path, using the nonzero winding number rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextFillPath
func CGContextFillPath(c ContextRef) {
	_CGContextFillPath(c)
}

// Paints the areas contained within the provided rectangles, using the fill color in the current graphics state.
//
// Added in macOS 10.0.
// Paints the areas contained within the provided rectangles, using the fill color in the current graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextFillRects
func CGContextFillRects(c ContextRef, rects unsafe.Pointer, count uintptr) {
	_CGContextFillRects(c, rects, count)
}

// CGContextGetContentToneMappingInfo is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextGetContentToneMappingInfo
func CGContextGetContentToneMappingInfo(c ContextRef) CGContentToneMappingInfo {
	return _CGContextGetContentToneMappingInfo(c)
}

// CGContextGetEDRTargetHeadroom is a CoreGraphics function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextGetEDRTargetHeadroom(_:)
func CGContextGetEDRTargetHeadroom(c ContextRef) float32 {
	return _CGContextGetEDRTargetHeadroom(c)
}

// CGContextGetTextPosition is a CoreGraphics function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextGetTextPosition
func CGContextGetTextPosition(c ContextRef) Point {
	return _CGContextGetTextPosition(c)
}

// Begins a new subpath at the point you specify.
//
// Added in macOS 10.0.
// Begins a new subpath at the point you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextMoveToPoint
func CGContextMoveToPoint(c ContextRef, x Float, y Float) {
	_CGContextMoveToPoint(c, x, y)
}

// Decrements the retain count of a graphics context.
//
// Added in macOS 10.0.
// Decrements the retain count of a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextRelease
func CGContextRelease(c ContextRef) {
	_CGContextRelease(c)
}

// Increments the retain count of a graphics context.
//
// Added in macOS 10.0.
// Increments the retain count of a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextRetain
func CGContextRetain(c ContextRef) ContextRef {
	return _CGContextRetain(c)
}

// CGContextSetContentToneMappingInfo is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextSetContentToneMappingInfo
func CGContextSetContentToneMappingInfo(c ContextRef, info CGContentToneMappingInfo) {
	_CGContextSetContentToneMappingInfo(c, info)
}

// Sets the level of interpolation quality for a graphics context.
//
// Added in macOS 10.0.
// Sets the level of interpolation quality for a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextSetInterpolationQuality
func CGContextSetInterpolationQuality(c ContextRef, quality InterpolationQuality) {
	_CGContextSetInterpolationQuality(c, quality)
}

// Sets the pattern for dashed lines in a graphics context.
//
// Added in macOS 10.0.
// Sets the pattern for dashed lines in a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextSetLineDash
func CGContextSetLineDash(c ContextRef, phase Float, lengths []float64, count uintptr) {
	_CGContextSetLineDash(c, phase, lengths, count)
}

// Sets the current text matrix.
//
// Added in macOS 10.0.
// Sets the current text matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextSetTextMatrix
func CGContextSetTextMatrix(c ContextRef, t AffineTransform) {
	_CGContextSetTextMatrix(c, t)
}

// Sets the location at which text is drawn.
//
// Added in macOS 10.0.
// Sets the location at which text is drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextSetTextPosition
func CGContextSetTextPosition(c ContextRef, x Float, y Float) {
	_CGContextSetTextPosition(c, x, y)
}

// Draws glyphs at the provided position.
//
// Added in macOS 10.5.
// Draws glyphs at the provided position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextShowGlyphsAtPositions
func CGContextShowGlyphsAtPositions(c ContextRef, glyphs unsafe.Pointer, Lpositions unsafe.Pointer, count uintptr) {
	_CGContextShowGlyphsAtPositions(c, glyphs, Lpositions, count)
}

// Strokes a sequence of line segments.
//
// Added in macOS 10.4.
// Strokes a sequence of line segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGContextStrokeLineSegments
func CGContextStrokeLineSegments(c ContextRef, points unsafe.Pointer, count uintptr) {
	_CGContextStrokeLineSegments(c, points, count)
}

// CGConvertColorDataWithFormat is a CoreGraphics function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConvertColorDataWithFormat(_:_:_:_:_:_:_:)
func CGConvertColorDataWithFormat(width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorDataFormat, src_data unsafe.Pointer, src_format CGColorDataFormat, options DictionaryRef) bool {
	return _CGConvertColorDataWithFormat(width, height, dst_data, dst_format, src_data, src_format, options)
}

// Returns a Boolean value indicating whether the mouse cursor is drawn in framebuffer memory.

// Returns a Boolean value indicating whether the mouse cursor is drawn in framebuffer memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCursorIsDrawnInFramebuffer()
func CGCursorIsDrawnInFramebuffer() unsafe.Pointer {
	return _CGCursorIsDrawnInFramebuffer()
}

// Returns a Boolean value indicating whether the mouse cursor is visible.

// Returns a Boolean value indicating whether the mouse cursor is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCursorIsVisible()
func CGCursorIsVisible() unsafe.Pointer {
	return _CGCursorIsVisible()
}

// Creates a data consumer that writes to a CFData object.
//
// Added in macOS 10.4.
// Creates a data consumer that writes to a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/init(data:)
func CGDataConsumerCreateWithCFData(data MutableDataRef) DataConsumerRef {
	return _CGDataConsumerCreateWithCFData(data)
}

// Creates a data consumer that uses callback functions to write data.
//
// Added in macOS 10.0.
// Creates a data consumer that uses callback functions to write data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/init(info:cbks:)
func CGDataConsumerCreate(info unsafe.Pointer, cbks unsafe.Pointer) DataConsumerRef {
	return _CGDataConsumerCreate(info, cbks)
}

// Creates a data consumer that writes data to a location specified by a URL.
//
// Added in macOS 10.0.
// Creates a data consumer that writes data to a location specified by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/init(url:)
func CGDataConsumerCreateWithURL(url URLRef) DataConsumerRef {
	return _CGDataConsumerCreateWithURL(url)
}

// Returns the Core Foundation type identifier for Core Graphics data consumers.
//
// Added in macOS 10.2.
// Returns the Core Foundation type identifier for Core Graphics data consumers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/typeID
func CGDataConsumerGetTypeID() TypeID {
	return _CGDataConsumerGetTypeID()
}

// Decrements the retain count of a data consumer.
//
// Added in macOS 10.0.
// Decrements the retain count of a data consumer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumerRelease
func CGDataConsumerRelease(consumer DataConsumerRef) {
	_CGDataConsumerRelease(consumer)
}

// Increments the retain count of a data consumer.
//
// Added in macOS 10.0.
// Increments the retain count of a data consumer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumerRetain
func CGDataConsumerRetain(consumer DataConsumerRef) DataConsumerRef {
	return _CGDataConsumerRetain(consumer)
}

// Returns a copy of the provider’s data.
//
// Added in macOS 10.3.
// Returns a copy of the provider’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/data
func CGDataProviderCopyData(provider DataProviderRef) DataRef {
	return _CGDataProviderCopyData(provider)
}

// CGDataProviderGetInfo is a CoreGraphics function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/info
func CGDataProviderGetInfo(provider DataProviderRef) unsafe.Pointer {
	return _CGDataProviderGetInfo(provider)
}

// Creates a data provider that reads from a CFData object.
//
// Added in macOS 10.4.
// Creates a data provider that reads from a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(data:)
func CGDataProviderCreateWithCFData(data DataRef) DataProviderRef {
	return _CGDataProviderCreateWithCFData(data)
}

// Creates a direct-access data provider that uses data your program supplies.
//
// Added in macOS 10.0.
// Creates a direct-access data provider that uses data your program supplies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(dataInfo:data:size:releaseData:)
func CGDataProviderCreateWithData(info unsafe.Pointer, data unsafe.Pointer, size uintptr, releaseData DataProviderReleaseDataCallback) DataProviderRef {
	return _CGDataProviderCreateWithData(info, data, size, releaseData)
}

// Creates a direct-access data provider.
//
// Added in macOS 10.5.
// Creates a direct-access data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(directInfo:size:callbacks:)
func CGDataProviderCreateDirect(info unsafe.Pointer, size unsafe.Pointer, callbacks unsafe.Pointer) DataProviderRef {
	return _CGDataProviderCreateDirect(info, size, callbacks)
}

// Creates a direct-access data provider that uses a file to supply data.
//
// Added in macOS 10.0.
// Creates a direct-access data provider that uses a file to supply data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(filename:)
func CGDataProviderCreateWithFilename(filename unsafe.Pointer) DataProviderRef {
	return _CGDataProviderCreateWithFilename(filename)
}

// Creates a sequential-access data provider.
//
// Added in macOS 10.5.
// Creates a sequential-access data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(sequentialInfo:callbacks:)
func CGDataProviderCreateSequential(info unsafe.Pointer, callbacks unsafe.Pointer) DataProviderRef {
	return _CGDataProviderCreateSequential(info, callbacks)
}

// Creates a direct-access data provider that uses a URL to supply data.
//
// Added in macOS 10.0.
// Creates a direct-access data provider that uses a URL to supply data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(url:)
func CGDataProviderCreateWithURL(url URLRef) DataProviderRef {
	return _CGDataProviderCreateWithURL(url)
}

// Returns the Core Foundation type identifier for data providers.
//
// Added in macOS 10.2.
// Returns the Core Foundation type identifier for data providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/typeID
func CGDataProviderGetTypeID() TypeID {
	return _CGDataProviderGetTypeID()
}

// Decrements the retain count of a data provider.
//
// Added in macOS 10.0.
// Decrements the retain count of a data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProviderRelease
func CGDataProviderRelease(provider DataProviderRef) {
	_CGDataProviderRelease(provider)
}

// Increments the retain count of a data provider.
//
// Added in macOS 10.0.
// Increments the retain count of a data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDataProviderRetain
func CGDataProviderRetain(provider DataProviderRef) DataProviderRef {
	return _CGDataProviderRetain(provider)
}

// Returns the GPU device instance that’s currently driving a display.
//
// Added in macOS 10.11.
// Returns the GPU device instance that’s currently driving a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDirectDisplayCopyCurrentMetalDevice(_:)
func CGDirectDisplayCopyCurrentMetalDevice(display DirectDisplayID) unsafe.Pointer {
	return _CGDirectDisplayCopyCurrentMetalDevice(display)
}

// Returns information about the currently available display modes.

// Returns information about the currently available display modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayAvailableModes(_:)
func CGDisplayAvailableModes(dsp DirectDisplayID) ArrayRef {
	return _CGDisplayAvailableModes(dsp)
}

// Returns information about the display mode closest to a specified depth and screen size.

// Returns information about the display mode closest to a specified depth and screen size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayBestModeForParameters(_:_:_:_:_:)
func CGDisplayBestModeForParameters(display DirectDisplayID, bitsPerPixel uintptr, width uintptr, height uintptr, exactMatch unsafe.Pointer) DictionaryRef {
	return _CGDisplayBestModeForParameters(display, bitsPerPixel, width, height, exactMatch)
}

// Returns information about the display mode closest to a specified depth, screen size, and refresh rate.

// Returns information about the display mode closest to a specified depth, screen size, and refresh rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayBestModeForParametersAndRefreshRate(_:_:_:_:_:_:)
func CGDisplayBestModeForParametersAndRefreshRate(display DirectDisplayID, bitsPerPixel uintptr, width uintptr, height uintptr, refreshRate RefreshRate, exactMatch unsafe.Pointer) DictionaryRef {
	return _CGDisplayBestModeForParametersAndRefreshRate(display, bitsPerPixel, width, height, refreshRate, exactMatch)
}

// Returns the bounds of a display in the global display coordinate space.
//
// Added in macOS 10.0.
// Returns the bounds of a display in the global display coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayBounds(_:)
func CGDisplayBounds(display DirectDisplayID) Rect {
	return _CGDisplayBounds(display)
}

// Obtains exclusive use of a display, preventing other applications and system services from using the display or changing its configuration.
//
// Added in macOS 10.0.
// Obtains exclusive use of a display, preventing other applications and system services from using the display or changing its configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCapture(_:)
func CGDisplayCapture(display DirectDisplayID) Error {
	return _CGDisplayCapture(display)
}

// Obtains exclusive use of a display for an application using the options you specify.
//
// Added in macOS 10.3.
// Obtains exclusive use of a display for an application using the options you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCaptureWithOptions(_:_:)
func CGDisplayCaptureWithOptions(display DirectDisplayID, options CaptureOptions) Error {
	return _CGDisplayCaptureWithOptions(display, options)
}

// Returns information about the currently available display modes.
//
// Added in macOS 10.6.
// Returns information about the currently available display modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCopyAllDisplayModes(_:_:)
func CGDisplayCopyAllDisplayModes(display DirectDisplayID, options DictionaryRef) ArrayRef {
	return _CGDisplayCopyAllDisplayModes(display, options)
}

// Returns the color space for a display.
//
// Added in macOS 10.5.
// Returns the color space for a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCopyColorSpace(_:)
func CGDisplayCopyColorSpace(display DirectDisplayID) ColorSpaceRef {
	return _CGDisplayCopyColorSpace(display)
}

// Returns information about a display’s current configuration.
//
// Added in macOS 10.6.
// Returns information about a display’s current configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCopyDisplayMode(_:)
func CGDisplayCopyDisplayMode(display DirectDisplayID) DisplayModeRef {
	return _CGDisplayCopyDisplayMode(display)
}

// Returns an image containing the contents of the specified display.

// Returns an image containing the contents of the specified display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCreateImage(_:)
func CGDisplayCreateImage(displayID DirectDisplayID) ImageRef {
	return _CGDisplayCreateImage(displayID)
}

// Returns an image containing the contents of a portion of the specified display.

// Returns an image containing the contents of a portion of the specified display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCreateImage(_:rect:)
func CGDisplayCreateImageForRect(display DirectDisplayID, rect Rect) ImageRef {
	return _CGDisplayCreateImageForRect(display, rect)
}

// Returns information about the current display mode.

// Returns information about the current display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCurrentMode(_:)
func CGDisplayCurrentMode(display DirectDisplayID) DictionaryRef {
	return _CGDisplayCurrentMode(display)
}

// Performs a single fade operation.
//
// Added in macOS 10.2.
// Performs a single fade operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayFade(_:_:_:_:_:_:_:_:)
func CGDisplayFade(token DisplayFadeReservationToken, duration DisplayFadeInterval, startBlend DisplayBlendFraction, endBlend DisplayBlendFraction, redBlend float32, greenBlend float32, blueBlend float32, synchronous unsafe.Pointer) Error {
	return _CGDisplayFade(token, duration, startBlend, endBlend, redBlend, greenBlend, blueBlend, synchronous)
}

// Returns a Boolean value indicating whether a fade operation is currently in progress.

// Returns a Boolean value indicating whether a fade operation is currently in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayFadeOperationInProgress()
func CGDisplayFadeOperationInProgress() unsafe.Pointer {
	return _CGDisplayFadeOperationInProgress()
}

// Returns the capacity, or number of entries, in the gamma table for a display.
//
// Added in macOS 10.3.
// Returns the capacity, or number of entries, in the gamma table for a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayGammaTableCapacity(_:)
func CGDisplayGammaTableCapacity(display DirectDisplayID) uint32 {
	return _CGDisplayGammaTableCapacity(display)
}

// Returns a graphics context suitable for drawing to a captured display.
//
// Added in macOS 10.3.
// Returns a graphics context suitable for drawing to a captured display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayGetDrawingContext(_:)
func CGDisplayGetDrawingContext(display DirectDisplayID) ContextRef {
	return _CGDisplayGetDrawingContext(display)
}

// Hides the mouse cursor, and increments the hide cursor count.
//
// Added in macOS 10.0.
// Hides the mouse cursor, and increments the hide cursor count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayHideCursor(_:)
func CGDisplayHideCursor(display DirectDisplayID) Error {
	return _CGDisplayHideCursor(display)
}

// Maps a display ID to an OpenGL display mask.
//
// Added in macOS 10.0.
// Maps a display ID to an OpenGL display mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIDToOpenGLDisplayMask(_:)
func CGDisplayIDToOpenGLDisplayMask(display DirectDisplayID) OpenGLDisplayMask {
	return _CGDisplayIDToOpenGLDisplayMask(display)
}

// Returns the I/O Kit service port of the specified display.

// Returns the I/O Kit service port of the specified display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIOServicePort(_:)
func CGDisplayIOServicePort(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayIOServicePort(display)
}

// Returns a Boolean value indicating whether a display is active.
//
// Added in macOS 10.2.
// Returns a Boolean value indicating whether a display is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsActive(_:)
func CGDisplayIsActive(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayIsActive(display)
}

// Returns a Boolean value indicating whether a display is always in a mirroring set.
//
// Added in macOS 10.2.
// Returns a Boolean value indicating whether a display is always in a mirroring set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsAlwaysInMirrorSet(_:)
func CGDisplayIsAlwaysInMirrorSet(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayIsAlwaysInMirrorSet(display)
}

// Returns a Boolean value indicating whether a display is sleeping (and is therefore not drawable).
//
// Added in macOS 10.2.
// Returns a Boolean value indicating whether a display is sleeping (and is therefore not drawable).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsAsleep(_:)
func CGDisplayIsAsleep(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayIsAsleep(display)
}

// Returns a Boolean value indicating whether a display is built-in, such as the internal display in portable systems.
//
// Added in macOS 10.2.
// Returns a Boolean value indicating whether a display is built-in, such as the internal display in portable systems.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsBuiltin(_:)
func CGDisplayIsBuiltin(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayIsBuiltin(display)
}

// Returns a Boolean value indicating whether a display is captured.

// Returns a Boolean value indicating whether a display is captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsCaptured(_:)
func CGDisplayIsCaptured(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayIsCaptured(display)
}

// Returns a Boolean value indicating whether a display is in a hardware mirroring set.
//
// Added in macOS 10.2.
// Returns a Boolean value indicating whether a display is in a hardware mirroring set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsInHWMirrorSet(_:)
func CGDisplayIsInHWMirrorSet(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayIsInHWMirrorSet(display)
}

// Returns a Boolean value indicating whether a display is in a mirroring set.
//
// Added in macOS 10.2.
// Returns a Boolean value indicating whether a display is in a mirroring set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsInMirrorSet(_:)
func CGDisplayIsInMirrorSet(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayIsInMirrorSet(display)
}

// Returns a Boolean value indicating whether a display is the main display.
//
// Added in macOS 10.2.
// Returns a Boolean value indicating whether a display is the main display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsMain(_:)
func CGDisplayIsMain(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayIsMain(display)
}

// Returns a Boolean value indicating whether a display is connected or online.
//
// Added in macOS 10.2.
// Returns a Boolean value indicating whether a display is connected or online.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsOnline(_:)
func CGDisplayIsOnline(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayIsOnline(display)
}

// Returns a Boolean value indicating whether a display is running in a stereo graphics mode.
//
// Added in macOS 10.4.
// Returns a Boolean value indicating whether a display is running in a stereo graphics mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsStereo(_:)
func CGDisplayIsStereo(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayIsStereo(display)
}

// For a secondary display in a mirroring set, returns the primary display.
//
// Added in macOS 10.2.
// For a secondary display in a mirroring set, returns the primary display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMirrorsDisplay(_:)
func CGDisplayMirrorsDisplay(display DirectDisplayID) DirectDisplayID {
	return _CGDisplayMirrorsDisplay(display)
}

// Returns the height of the specified display mode.
//
// Added in macOS 10.6.
// Returns the height of the specified display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/height
func CGDisplayModeGetHeight(mode DisplayModeRef) uintptr {
	return _CGDisplayModeGetHeight(mode)
}

// Returns the I/O Kit display mode ID of the specified display mode.
//
// Added in macOS 10.6.
// Returns the I/O Kit display mode ID of the specified display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/ioDisplayModeID
func CGDisplayModeGetIODisplayModeID(mode DisplayModeRef) int32 {
	return _CGDisplayModeGetIODisplayModeID(mode)
}

// Returns the I/O Kit flags of the specified display mode.
//
// Added in macOS 10.6.
// Returns the I/O Kit flags of the specified display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/ioFlags
func CGDisplayModeGetIOFlags(mode DisplayModeRef) uint32 {
	return _CGDisplayModeGetIOFlags(mode)
}

// Returns a Boolean value indicating whether the specified display mode is usable for a desktop graphical user interface.
//
// Added in macOS 10.6.
// Returns a Boolean value indicating whether the specified display mode is usable for a desktop graphical user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/isUsableForDesktopGUI()
func CGDisplayModeIsUsableForDesktopGUI(mode DisplayModeRef) bool {
	return _CGDisplayModeIsUsableForDesktopGUI(mode)
}

// Returns the pixel encoding of the specified display mode.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.6.
// Returns the pixel encoding of the specified display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/pixelEncoding
func CGDisplayModeCopyPixelEncoding(mode DisplayModeRef) StringRef {
	return _CGDisplayModeCopyPixelEncoding(mode)
}

// CGDisplayModeGetPixelHeight is a CoreGraphics function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/pixelHeight
func CGDisplayModeGetPixelHeight(mode DisplayModeRef) uintptr {
	return _CGDisplayModeGetPixelHeight(mode)
}

// CGDisplayModeGetPixelWidth is a CoreGraphics function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/pixelWidth
func CGDisplayModeGetPixelWidth(mode DisplayModeRef) uintptr {
	return _CGDisplayModeGetPixelWidth(mode)
}

// Returns the refresh rate of the specified display mode.
//
// Added in macOS 10.6.
// Returns the refresh rate of the specified display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/refreshRate
func CGDisplayModeGetRefreshRate(mode DisplayModeRef) float64 {
	return _CGDisplayModeGetRefreshRate(mode)
}

// Returns the type identifier of Quartz display modes.
//
// Added in macOS 10.6.
// Returns the type identifier of Quartz display modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/typeID
func CGDisplayModeGetTypeID() TypeID {
	return _CGDisplayModeGetTypeID()
}

// Returns the width of the specified display mode.
//
// Added in macOS 10.6.
// Returns the width of the specified display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/width
func CGDisplayModeGetWidth(mode DisplayModeRef) uintptr {
	return _CGDisplayModeGetWidth(mode)
}

// Releases a Core Graphics display mode.
//
// Added in macOS 10.6.
// Releases a Core Graphics display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayModeRelease
func CGDisplayModeRelease(mode DisplayModeRef) {
	_CGDisplayModeRelease(mode)
}

// Retains a Core Graphics display mode.
//
// Added in macOS 10.6.
// Retains a Core Graphics display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayModeRetain
func CGDisplayModeRetain(mode DisplayModeRef) DisplayModeRef {
	return _CGDisplayModeRetain(mode)
}

// Returns the model number of a display monitor.
//
// Added in macOS 10.2.
// Returns the model number of a display monitor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayModelNumber(_:)
func CGDisplayModelNumber(display DirectDisplayID) uint32 {
	return _CGDisplayModelNumber(display)
}

// Moves the mouse cursor to a specified point relative to the upper-left corner of the display.
//
// Added in macOS 10.0.
// Moves the mouse cursor to a specified point relative to the upper-left corner of the display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMoveCursorToPoint(_:_:)
func CGDisplayMoveCursorToPoint(display DirectDisplayID, point Point) Error {
	return _CGDisplayMoveCursorToPoint(display, point)
}

// Returns the display height in pixel units.
//
// Added in macOS 10.0.
// Returns the display height in pixel units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayPixelsHigh(_:)
func CGDisplayPixelsHigh(display DirectDisplayID) uintptr {
	return _CGDisplayPixelsHigh(display)
}

// Returns the display width in pixel units.
//
// Added in macOS 10.0.
// Returns the display width in pixel units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayPixelsWide(_:)
func CGDisplayPixelsWide(display DirectDisplayID) uintptr {
	return _CGDisplayPixelsWide(display)
}

// Returns the primary display in a hardware mirroring set.
//
// Added in macOS 10.2.
// Returns the primary display in a hardware mirroring set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayPrimaryDisplay(_:)
func CGDisplayPrimaryDisplay(display DirectDisplayID) DirectDisplayID {
	return _CGDisplayPrimaryDisplay(display)
}

// Registers a callback function to be invoked whenever a local display is reconfigured.
//
// Added in macOS 10.3.
// Registers a callback function to be invoked whenever a local display is reconfigured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRegisterReconfigurationCallback(_:_:)
func CGDisplayRegisterReconfigurationCallback(callback DisplayReconfigurationCallBack, userInfo unsafe.Pointer) Error {
	return _CGDisplayRegisterReconfigurationCallback(callback, userInfo)
}

// Releases a captured display.
//
// Added in macOS 10.0.
// Releases a captured display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRelease(_:)
func CGDisplayRelease(display DirectDisplayID) Error {
	return _CGDisplayRelease(display)
}

// Removes the registration of a callback function that’s invoked whenever a local display is reconfigured.
//
// Added in macOS 10.3.
// Removes the registration of a callback function that’s invoked whenever a local display is reconfigured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRemoveReconfigurationCallback(_:_:)
func CGDisplayRemoveReconfigurationCallback(callback DisplayReconfigurationCallBack, userInfo unsafe.Pointer) Error {
	return _CGDisplayRemoveReconfigurationCallback(callback, userInfo)
}

// Restores the gamma tables to the values in the user’s ColorSync display profile.
//
// Added in macOS 10.0.
// Restores the gamma tables to the values in the user’s ColorSync display profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRestoreColorSyncSettings()
func CGDisplayRestoreColorSyncSettings() {
	_CGDisplayRestoreColorSyncSettings()
}

// Returns the rotation angle of a display in degrees.
//
// Added in macOS 10.5.
// Returns the rotation angle of a display in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRotation(_:)
func CGDisplayRotation(display DirectDisplayID) float64 {
	return _CGDisplayRotation(display)
}

// Returns the width and height of a display in millimeters.
//
// Added in macOS 10.3.
// Returns the width and height of a display in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayScreenSize(_:)
func CGDisplayScreenSize(display DirectDisplayID) Size {
	return _CGDisplayScreenSize(display)
}

// Returns the serial number of a display monitor.
//
// Added in macOS 10.2.
// Returns the serial number of a display monitor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplaySerialNumber(_:)
func CGDisplaySerialNumber(display DirectDisplayID) uint32 {
	return _CGDisplaySerialNumber(display)
}

// Switches a display to a different mode.
//
// Added in macOS 10.6.
// Switches a display to a different mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplaySetDisplayMode(_:_:_:)
func CGDisplaySetDisplayMode(display DirectDisplayID, mode DisplayModeRef, options DictionaryRef) Error {
	return _CGDisplaySetDisplayMode(display, mode, options)
}

// Immediately enables or disables stereo operation for a display.
//
// Added in macOS 10.4.
// Immediately enables or disables stereo operation for a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplaySetStereoOperation(_:_:_:_:)
func CGDisplaySetStereoOperation(display DirectDisplayID, stereo unsafe.Pointer, forceBlueLine unsafe.Pointer, option ConfigureOption) Error {
	return _CGDisplaySetStereoOperation(display, stereo, forceBlueLine, option)
}

// Decrements the hide cursor count, and shows the mouse cursor if the count is .
//
// Added in macOS 10.0.
// Decrements the hide cursor count, and shows the mouse cursor if the count is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayShowCursor(_:)
func CGDisplayShowCursor(display DirectDisplayID) Error {
	return _CGDisplayShowCursor(display)
}

// Creates a new display stream whose updates are delivered to a dispatch queue.

// Creates a new display stream whose updates are delivered to a dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/init(dispatchQueueDisplay:outputWidth:outputHeight:pixelFormat:properties:queue:handler:)
func CGDisplayStreamCreateWithDispatchQueue(display DirectDisplayID, outputWidth uintptr, outputHeight uintptr, pixelFormat int32, properties DictionaryRef, queue unsafe.Pointer, handler DisplayStreamFrameAvailableHandler) DisplayStreamRef {
	return _CGDisplayStreamCreateWithDispatchQueue(display, outputWidth, outputHeight, pixelFormat, properties, queue, handler)
}

// Creates a new display stream to be used with a .

// Creates a new display stream to be used with a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/init(display:outputWidth:outputHeight:pixelFormat:properties:handler:)
func CGDisplayStreamCreate(display DirectDisplayID, outputWidth uintptr, outputHeight uintptr, pixelFormat int32, properties DictionaryRef, handler DisplayStreamFrameAvailableHandler) DisplayStreamRef {
	return _CGDisplayStreamCreate(display, outputWidth, outputHeight, pixelFormat, properties, handler)
}

// Gets the run loop source for a display stream.

// Gets the run loop source for a display stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/runLoopSource
func CGDisplayStreamGetRunLoopSource(displayStream DisplayStreamRef) RunLoopSourceRef {
	return _CGDisplayStreamGetRunLoopSource(displayStream)
}

// Tells a stream to start sending updates.

// Tells a stream to start sending updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/start()
func CGDisplayStreamStart(displayStream DisplayStreamRef) Error {
	return _CGDisplayStreamStart(displayStream)
}

// Tells a stream to stop sending updates.

// Tells a stream to stop sending updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/stop()
func CGDisplayStreamStop(displayStream DisplayStreamRef) Error {
	return _CGDisplayStreamStop(displayStream)
}

// Returns the type identifier of a Quartz display stream.

// Returns the type identifier of a Quartz display stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/typeID
func CGDisplayStreamGetTypeID() TypeID {
	return _CGDisplayStreamGetTypeID()
}

// Returns the number of frames that have been dropped since the last call to your update handler.

// Returns the number of frames that have been dropped since the last call to your update handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdate/dropCount
func CGDisplayStreamUpdateGetDropCount(updateRef DisplayStreamUpdateRef) uintptr {
	return _CGDisplayStreamUpdateGetDropCount(updateRef)
}

// Return the movement delta values for a single update.

// Return the movement delta values for a single update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdate/getMovedRectsDelta(dx:dy:)
func CGDisplayStreamUpdateGetMovedRectsDelta(updateRef DisplayStreamUpdateRef, dx []float64, dy []float64) {
	_CGDisplayStreamUpdateGetMovedRectsDelta(updateRef, dx, dy)
}

// Returns an array of rectangles that describe where the frame has changed since the previous frame.

// Returns an array of rectangles that describe where the frame has changed since the previous frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdate/getRects(_:rectCount:)
func CGDisplayStreamUpdateGetRects(updateRef DisplayStreamUpdateRef, rectType DisplayStreamUpdateRectType, rectCount unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayStreamUpdateGetRects(updateRef, rectType, rectCount)
}

// Combines two updates into a new update that includes the metadata for both source updates.

// Combines two updates into a new update that includes the metadata for both source updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdate/init(mergedUpdateFirstUpdate:secondUpdate:)
func CGDisplayStreamUpdateCreateMergedUpdate(firstUpdate DisplayStreamUpdateRef, secondUpdate DisplayStreamUpdateRef) DisplayStreamUpdateRef {
	return _CGDisplayStreamUpdateCreateMergedUpdate(firstUpdate, secondUpdate)
}

// Returns the type identifier of a Quartz display stream update.

// Returns the type identifier of a Quartz display stream update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdate/typeID
func CGDisplayStreamUpdateGetTypeID() TypeID {
	return _CGDisplayStreamUpdateGetTypeID()
}

// Switches a display to a different mode.

// Switches a display to a different mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplaySwitchToMode(_:_:)
func CGDisplaySwitchToMode(display DirectDisplayID, mode DictionaryRef) Error {
	return _CGDisplaySwitchToMode(display, mode)
}

// Returns the logical unit number of a display.
//
// Added in macOS 10.2.
// Returns the logical unit number of a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayUnitNumber(_:)
func CGDisplayUnitNumber(display DirectDisplayID) uint32 {
	return _CGDisplayUnitNumber(display)
}

// Returns a Boolean value indicating whether Quartz is using OpenGL-based window acceleration (Quartz Extreme) to render in a display.
//
// Added in macOS 10.2.
// Returns a Boolean value indicating whether Quartz is using OpenGL-based window acceleration (Quartz Extreme) to render in a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayUsesOpenGLAcceleration(_:)
func CGDisplayUsesOpenGLAcceleration(display DirectDisplayID) unsafe.Pointer {
	return _CGDisplayUsesOpenGLAcceleration(display)
}

// Returns the vendor number of the specified display’s monitor.
//
// Added in macOS 10.2.
// Returns the vendor number of the specified display’s monitor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayVendorNumber(_:)
func CGDisplayVendorNumber(display DirectDisplayID) uint32 {
	return _CGDisplayVendorNumber(display)
}

// CGEXRToneMappingGammaGetDefaultOptions is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEXRToneMappingGammaGetDefaultOptions
func CGEXRToneMappingGammaGetDefaultOptions() DictionaryRef {
	return _CGEXRToneMappingGammaGetDefaultOptions()
}

// Enables or disables the merging of actual key and mouse state with the application-specified state in a synthetic event.

// Enables or disables the merging of actual key and mouse state with the application-specified state in a synthetic event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEnableEventStateCombining(_:)
func CGEnableEventStateCombining(combineState unsafe.Pointer) Error {
	return _CGEnableEventStateCombining(combineState)
}

// CGErrorSetCallback is a CoreGraphics function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGErrorSetCallback(_:)
func CGErrorSetCallback(callback ErrorCallback) {
	_CGErrorSetCallback(callback)
}

// Returns a copy of an existing Quartz event.
//
// Added in macOS 10.4.
// Returns a copy of an existing Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/copy()
func CGEventCreateCopy(event EventRef) EventRef {
	return _CGEventCreateCopy(event)
}

// Returns the event flags of a Quartz event.
//
// Added in macOS 10.4.
// Returns the event flags of a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/flags
func CGEventGetFlags(event EventRef) EventFlags {
	return _CGEventGetFlags(event)
}

// Returns the floating-point value of a field in a Quartz event.
//
// Added in macOS 10.4.
// Returns the floating-point value of a field in a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/getDoubleValueField(_:)
func CGEventGetDoubleValueField(event EventRef, field EventField) float64 {
	return _CGEventGetDoubleValueField(event, field)
}

// Returns the integer value of a field in a Quartz event.
//
// Added in macOS 10.4.
// Returns the integer value of a field in a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/getIntegerValueField(_:)
func CGEventGetIntegerValueField(event EventRef, field EventField) int64 {
	return _CGEventGetIntegerValueField(event, field)
}

// Returns a new Quartz keyboard event.
//
// Added in macOS 10.4.
// Returns a new Quartz keyboard event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(keyboardEventSource:virtualKey:keyDown:)
func CGEventCreateKeyboardEvent(source EventSourceRef, virtualKey KeyCode, keyDown bool) EventRef {
	return _CGEventCreateKeyboardEvent(source, virtualKey, keyDown)
}

// Returns a new Quartz mouse event.
//
// Added in macOS 10.4.
// Returns a new Quartz mouse event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(mouseEventSource:mouseType:mouseCursorPosition:mouseButton:)
func CGEventCreateMouseEvent(source EventSourceRef, mouseType EventType, mouseCursorPosition Point, mouseButton MouseButton) EventRef {
	return _CGEventCreateMouseEvent(source, mouseType, mouseCursorPosition, mouseButton)
}

// CGEventCreateScrollWheelEvent2 is a CoreGraphics function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(scrollWheelEvent2Source:units:wheelCount:wheel1:wheel2:wheel3:)
func CGEventCreateScrollWheelEvent2(source EventSourceRef, units ScrollEventUnit, wheelCount uint32, wheel1 int32, wheel2 int32, wheel3 int32) EventRef {
	return _CGEventCreateScrollWheelEvent2(source, units, wheelCount, wheel1, wheel2, wheel3)
}

// Returns a new Quartz event.
//
// Added in macOS 10.4.
// Returns a new Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(source:)
func CGEventCreate(source EventSourceRef) EventRef {
	return _CGEventCreate(source)
}

// Returns a Quartz event created from a flattened data representation of the event.
//
// Added in macOS 10.4.
// Returns a Quartz event created from a flattened data representation of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(withDataAllocator:data:)
func CGEventCreateFromData(allocator AllocatorRef, data DataRef) EventRef {
	return _CGEventCreateFromData(allocator, data)
}

// Returns the Unicode string associated with a Quartz keyboard event.
//
// Added in macOS 10.4.
// Returns the Unicode string associated with a Quartz keyboard event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/keyboardGetUnicodeString(maxStringLength:actualStringLength:unicodeString:)
func CGEventKeyboardGetUnicodeString(event EventRef, maxStringLength unsafe.Pointer, actualStringLength unsafe.Pointer, unicodeString unsafe.Pointer) {
	_CGEventKeyboardGetUnicodeString(event, maxStringLength, actualStringLength, unicodeString)
}

// Sets the Unicode string associated with a Quartz keyboard event.
//
// Added in macOS 10.4.
// Sets the Unicode string associated with a Quartz keyboard event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/keyboardSetUnicodeString(stringLength:unicodeString:)
func CGEventKeyboardSetUnicodeString(event EventRef, stringLength unsafe.Pointer, unicodeString unsafe.Pointer) {
	_CGEventKeyboardSetUnicodeString(event, stringLength, unicodeString)
}

// Returns the location of a Quartz mouse event.
//
// Added in macOS 10.4.
// Returns the location of a Quartz mouse event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/location
func CGEventGetLocation(event EventRef) Point {
	return _CGEventGetLocation(event)
}

// Posts a Quartz event into the event stream at a specified location.
//
// Added in macOS 10.4.
// Posts a Quartz event into the event stream at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/post(tap:)
func CGEventPost(tap EventTapLocation, event EventRef) {
	_CGEventPost(tap, event)
}

// Posts a Quartz event into the event stream for a specific application.
//
// Added in macOS 10.4.
// Posts a Quartz event into the event stream for a specific application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/postToPSN(processSerialNumber:)
func CGEventPostToPSN(processSerialNumber unsafe.Pointer, event EventRef) {
	_CGEventPostToPSN(processSerialNumber, event)
}

// CGEventPostToPid is a CoreGraphics function.
//
// Added in macOS 10.11.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/postToPid(_:)
func CGEventPostToPid(pid unsafe.Pointer, event EventRef) {
	_CGEventPostToPid(pid, event)
}

// Sets the floating-point value of a field in a Quartz event.
//
// Added in macOS 10.4.
// Sets the floating-point value of a field in a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/setDoubleValueField(_:value:)
func CGEventSetDoubleValueField(event EventRef, field EventField, value float64) {
	_CGEventSetDoubleValueField(event, field, value)
}

// Sets the integer value of a field in a Quartz event.
//
// Added in macOS 10.4.
// Sets the integer value of a field in a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/setIntegerValueField(_:value:)
func CGEventSetIntegerValueField(event EventRef, field EventField, value int64) {
	_CGEventSetIntegerValueField(event, field, value)
}

// Sets the event source of a Quartz event.
//
// Added in macOS 10.4.
// Sets the event source of a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/setSource(_:)
func CGEventSetSource(event EventRef, source EventSourceRef) {
	_CGEventSetSource(event, source)
}

// Creates an event tap.
//
// Added in macOS 10.4.
// Creates an event tap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapCreate(tap:place:options:eventsOfInterest:callback:userInfo:)
func CGEventTapCreate(tap EventTapLocation, place EventTapPlacement, options EventTapOptions, eventsOfInterest EventMask, callback EventTapCallBack, userInfo unsafe.Pointer) MachPortRef {
	return _CGEventTapCreate(tap, place, options, eventsOfInterest, callback, userInfo)
}

// Creates an event tap for a specified process.
//
// Added in macOS 10.4.
// Creates an event tap for a specified process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapCreateForPSN(processSerialNumber:place:options:eventsOfInterest:callback:userInfo:)
func CGEventTapCreateForPSN(processSerialNumber unsafe.Pointer, place EventTapPlacement, options EventTapOptions, eventsOfInterest EventMask, callback EventTapCallBack, userInfo unsafe.Pointer) MachPortRef {
	return _CGEventTapCreateForPSN(processSerialNumber, place, options, eventsOfInterest, callback, userInfo)
}

// CGEventTapCreateForPid is a CoreGraphics function.
//
// Added in macOS 10.11.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapCreateForPid(pid:place:options:eventsOfInterest:callback:userInfo:)
func CGEventTapCreateForPid(pid unsafe.Pointer, place EventTapPlacement, options EventTapOptions, eventsOfInterest EventMask, callback EventTapCallBack, userInfo unsafe.Pointer) MachPortRef {
	return _CGEventTapCreateForPid(pid, place, options, eventsOfInterest, callback, userInfo)
}

// Enables or disables an event tap.
//
// Added in macOS 10.4.
// Enables or disables an event tap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapEnable(tap:enable:)
func CGEventTapEnable(tap MachPortRef, enable bool) {
	_CGEventTapEnable(tap, enable)
}

// Returns a Boolean value indicating whether an event tap is enabled.
//
// Added in macOS 10.4.
// Returns a Boolean value indicating whether an event tap is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapIsEnabled(tap:)
func CGEventTapIsEnabled(tap MachPortRef) bool {
	return _CGEventTapIsEnabled(tap)
}

// Posts a Quartz event from an event tap into the event stream.
//
// Added in macOS 10.4.
// Posts a Quartz event from an event tap into the event stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapPostEvent(_:)
func CGEventTapPostEvent(proxy EventTapProxy, event EventRef) {
	_CGEventTapPostEvent(proxy, event)
}

// Returns the timestamp of a Quartz event.
//
// Added in macOS 10.4.
// Returns the timestamp of a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/timestamp
func CGEventGetTimestamp(event EventRef) EventTimestamp {
	return _CGEventGetTimestamp(event)
}

// Returns the event type of a Quartz event (left mouse down, for example).
//
// Added in macOS 10.4.
// Returns the event type of a Quartz event (left mouse down, for example).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/type
func CGEventGetType(event EventRef) EventType {
	return _CGEventGetType(event)
}

// Returns the type identifier for the opaque type .
//
// Added in macOS 10.4.
// Returns the type identifier for the opaque type .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/typeID
func CGEventGetTypeID() TypeID {
	return _CGEventGetTypeID()
}

// Returns the location of a Quartz mouse event.
//
// Added in macOS 10.5.
// Returns the location of a Quartz mouse event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEvent/unflippedLocation
func CGEventGetUnflippedLocation(event EventRef) Point {
	return _CGEventGetUnflippedLocation(event)
}

// Returns a flattened data representation of a Quartz event.
//
// Added in macOS 10.4.
// Returns a flattened data representation of a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventCreateData
func CGEventCreateData(allocator AllocatorRef, event EventRef) DataRef {
	return _CGEventCreateData(allocator, event)
}

// Returns a new Quartz scrolling event.
//
// Added in macOS 10.5.
// Returns a new Quartz scrolling event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventCreateScrollWheelEvent
func CGEventCreateScrollWheelEvent(source EventSourceRef, units ScrollEventUnit, wheelCount uint32, wheel1 int32) EventRef {
	return _CGEventCreateScrollWheelEvent(source, units, wheelCount, wheel1)
}

// Sets the event flags of a Quartz event.
//
// Added in macOS 10.4.
// Sets the event flags of a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSetFlags
func CGEventSetFlags(event EventRef, flags EventFlags) {
	_CGEventSetFlags(event, flags)
}

// Sets the location of a Quartz mouse event.
//
// Added in macOS 10.4.
// Sets the location of a Quartz mouse event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSetLocation
func CGEventSetLocation(event EventRef, location Point) {
	_CGEventSetLocation(event, location)
}

// Sets the timestamp of a Quartz event.
//
// Added in macOS 10.4.
// Sets the timestamp of a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSetTimestamp
func CGEventSetTimestamp(event EventRef, timestamp EventTimestamp) {
	_CGEventSetTimestamp(event, timestamp)
}

// Sets the event type of a Quartz event (left mouse down, for example).
//
// Added in macOS 10.4.
// Sets the event type of a Quartz event (left mouse down, for example).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSetType
func CGEventSetType(event EventRef, type_ EventType) {
	_CGEventSetType(event, type_)
}

// Returns a Boolean value indicating the current button state of a Quartz event source.
//
// Added in macOS 10.4.
// Returns a Boolean value indicating the current button state of a Quartz event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/buttonState(_:button:)
func CGEventSourceButtonState(stateID EventSourceStateID, button MouseButton) bool {
	return _CGEventSourceButtonState(stateID, button)
}

// Returns a count of events of a given type seen since the window server started.
//
// Added in macOS 10.4.
// Returns a count of events of a given type seen since the window server started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/counterForEventType(_:eventType:)
func CGEventSourceCounterForEventType(stateID EventSourceStateID, eventType EventType) uint32 {
	return _CGEventSourceCounterForEventType(stateID, eventType)
}

// Returns the current flags of a Quartz event source.
//
// Added in macOS 10.4.
// Returns the current flags of a Quartz event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/flagsState(_:)
func CGEventSourceFlagsState(stateID EventSourceStateID) EventFlags {
	return _CGEventSourceFlagsState(stateID)
}

// Returns the mask that indicates which classes of local hardware events are enabled during event suppression.
//
// Added in macOS 10.4.
// Returns the mask that indicates which classes of local hardware events are enabled during event suppression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/getLocalEventsFilterDuringSuppressionState(_:)
func CGEventSourceGetLocalEventsFilterDuringSuppressionState(source EventSourceRef, state EventSuppressionState) EventFilterMask {
	return _CGEventSourceGetLocalEventsFilterDuringSuppressionState(source, state)
}

// Returns a Quartz event source created from an existing Quartz event.
//
// Added in macOS 10.4.
// Returns a Quartz event source created from an existing Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/init(event:)
func CGEventCreateSourceFromEvent(event EventRef) EventSourceRef {
	return _CGEventCreateSourceFromEvent(event)
}

// Returns a Quartz event source created with a specified source state.
//
// Added in macOS 10.4.
// Returns a Quartz event source created with a specified source state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/init(stateID:)
func CGEventSourceCreate(stateID EventSourceStateID) EventSourceRef {
	return _CGEventSourceCreate(stateID)
}

// Returns a Boolean value indicating the current keyboard state of a Quartz event source.
//
// Added in macOS 10.4.
// Returns a Boolean value indicating the current keyboard state of a Quartz event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/keyState(_:key:)
func CGEventSourceKeyState(stateID EventSourceStateID, key KeyCode) bool {
	return _CGEventSourceKeyState(stateID, key)
}

// Returns the keyboard type to be used with a Quartz event source.
//
// Added in macOS 10.4.
// Returns the keyboard type to be used with a Quartz event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/keyboardType
func CGEventSourceGetKeyboardType(source EventSourceRef) EventSourceKeyboardType {
	return _CGEventSourceGetKeyboardType(source)
}

// Returns the interval that local hardware events may be suppressed following the posting of a Quartz event.
//
// Added in macOS 10.4.
// Returns the interval that local hardware events may be suppressed following the posting of a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/localEventsSuppressionInterval
func CGEventSourceGetLocalEventsSuppressionInterval(source EventSourceRef) TimeInterval {
	return _CGEventSourceGetLocalEventsSuppressionInterval(source)
}

// Gets the scale of pixels per line in a scrolling event source.
//
// Added in macOS 10.5.
// Gets the scale of pixels per line in a scrolling event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/pixelsPerLine
func CGEventSourceGetPixelsPerLine(source EventSourceRef) float64 {
	return _CGEventSourceGetPixelsPerLine(source)
}

// Returns the elapsed time since the last event for a Quartz event source.
//
// Added in macOS 10.4.
// Returns the elapsed time since the last event for a Quartz event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/secondsSinceLastEventType(_:eventType:)
func CGEventSourceSecondsSinceLastEventType(stateID EventSourceStateID, eventType EventType) TimeInterval {
	return _CGEventSourceSecondsSinceLastEventType(stateID, eventType)
}

// Sets the mask that indicates which classes of local hardware events are enabled during event suppression.
//
// Added in macOS 10.4.
// Sets the mask that indicates which classes of local hardware events are enabled during event suppression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/setLocalEventsFilterDuringSuppressionState(_:state:)
func CGEventSourceSetLocalEventsFilterDuringSuppressionState(source EventSourceRef, filter EventFilterMask, state EventSuppressionState) {
	_CGEventSourceSetLocalEventsFilterDuringSuppressionState(source, filter, state)
}

// Returns the source state associated with a Quartz event source.
//
// Added in macOS 10.4.
// Returns the source state associated with a Quartz event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/sourceStateID
func CGEventSourceGetSourceStateID(source EventSourceRef) EventSourceStateID {
	return _CGEventSourceGetSourceStateID(source)
}

// Returns the type identifier for the opaque type .
//
// Added in macOS 10.4.
// Returns the type identifier for the opaque type .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/typeID
func CGEventSourceGetTypeID() TypeID {
	return _CGEventSourceGetTypeID()
}

// Returns the 64-bit user-specified data for a Quartz event source.
//
// Added in macOS 10.4.
// Returns the 64-bit user-specified data for a Quartz event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/userData
func CGEventSourceGetUserData(source EventSourceRef) int64 {
	return _CGEventSourceGetUserData(source)
}

// Sets the keyboard type to be used with a Quartz event source.
//
// Added in macOS 10.4.
// Sets the keyboard type to be used with a Quartz event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetKeyboardType
func CGEventSourceSetKeyboardType(source EventSourceRef, keyboardType EventSourceKeyboardType) {
	_CGEventSourceSetKeyboardType(source, keyboardType)
}

// Sets the interval that local hardware events may be suppressed following the posting of a Quartz event.
//
// Added in macOS 10.4.
// Sets the interval that local hardware events may be suppressed following the posting of a Quartz event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetLocalEventsSuppressionInterval
func CGEventSourceSetLocalEventsSuppressionInterval(source EventSourceRef, seconds TimeInterval) {
	_CGEventSourceSetLocalEventsSuppressionInterval(source, seconds)
}

// Sets the scale of pixels per line in a scrolling event source.
//
// Added in macOS 10.5.
// Sets the scale of pixels per line in a scrolling event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetPixelsPerLine
func CGEventSourceSetPixelsPerLine(source EventSourceRef, pixelsPerLine float64) {
	_CGEventSourceSetPixelsPerLine(source, pixelsPerLine)
}

// Sets the 64-bit user-specified data for a Quartz event source.
//
// Added in macOS 10.4.
// Sets the 64-bit user-specified data for a Quartz event source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetUserData
func CGEventSourceSetUserData(source EventSourceRef, userData int64) {
	_CGEventSourceSetUserData(source, userData)
}

// Returns the ascent of a font.
//
// Added in macOS 10.5.
// Returns the ascent of a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/ascent
func CGFontGetAscent(font FontRef) int {
	return _CGFontGetAscent(font)
}

// Determines whether Core Graphics can create a subset of the font in PostScript format.
//
// Added in macOS 10.4.
// Determines whether Core Graphics can create a subset of the font in PostScript format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/canCreatePostScriptSubset(_:)
func CGFontCanCreatePostScriptSubset(font FontRef, format FontPostScriptFormat) bool {
	return _CGFontCanCreatePostScriptSubset(font, format)
}

// Returns the cap height of a font.
//
// Added in macOS 10.5.
// Returns the cap height of a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/capHeight
func CGFontGetCapHeight(font FontRef) int {
	return _CGFontGetCapHeight(font)
}

// Creates a copy of a font using a variation specification dictionary.
//
// Added in macOS 10.4.
// Creates a copy of a font using a variation specification dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/copy(withVariations:)
func CGFontCreateCopyWithVariations(font FontRef, variations DictionaryRef) FontRef {
	return _CGFontCreateCopyWithVariations(font, variations)
}

// Creates a PostScript encoding of a font.
//
// Added in macOS 10.4.
// Creates a PostScript encoding of a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/createPostScriptEncoding(encoding:)
func CGFontCreatePostScriptEncoding(font FontRef, encoding unsafe.Pointer, p2 unsafe.Pointer) DataRef {
	return _CGFontCreatePostScriptEncoding(font, encoding, p2)
}

// Creates a subset of the font in the specified PostScript format.
//
// Added in macOS 10.4.
// Creates a subset of the font in the specified PostScript format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/createPostScriptSubset(subsetName:format:glyphs:count:encoding:)
func CGFontCreatePostScriptSubset(font FontRef, subsetName StringRef, format FontPostScriptFormat, glyphs unsafe.Pointer, count uintptr, encoding unsafe.Pointer, p6 unsafe.Pointer) DataRef {
	return _CGFontCreatePostScriptSubset(font, subsetName, format, glyphs, count, encoding, p6)
}

// Returns the descent of a font.
//
// Added in macOS 10.5.
// Returns the descent of a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/descent
func CGFontGetDescent(font FontRef) int {
	return _CGFontGetDescent(font)
}

// Returns the bounding box of a font.
//
// Added in macOS 10.5.
// Returns the bounding box of a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/fontBBox
func CGFontGetFontBBox(font FontRef) Rect {
	return _CGFontGetFontBBox(font)
}

// Returns the full name associated with a font object.
//
// Added in macOS 10.5.
// Returns the full name associated with a font object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/fullName
func CGFontCopyFullName(font FontRef) StringRef {
	return _CGFontCopyFullName(font)
}

// Gets the advance width of each glyph in the provided array.
//
// Added in macOS 10.0.
// Gets the advance width of each glyph in the provided array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/getGlyphAdvances(glyphs:count:advances:)
func CGFontGetGlyphAdvances(font FontRef, glyphs unsafe.Pointer, count uintptr, advances []int) bool {
	return _CGFontGetGlyphAdvances(font, glyphs, count, advances)
}

// Get the bounding box of each glyph in an array.
//
// Added in macOS 10.5.
// Get the bounding box of each glyph in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/getGlyphBBoxes(glyphs:count:bboxes:)
func CGFontGetGlyphBBoxes(font FontRef, glyphs unsafe.Pointer, count uintptr, bboxes unsafe.Pointer) bool {
	return _CGFontGetGlyphBBoxes(font, glyphs, count, bboxes)
}

// Returns the glyph for the glyph name associated with the specified font object.
//
// Added in macOS 10.5.
// Returns the glyph for the glyph name associated with the specified font object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/getGlyphWithGlyphName(name:)
func CGFontGetGlyphWithGlyphName(font FontRef, name StringRef) Glyph {
	return _CGFontGetGlyphWithGlyphName(font, name)
}

// Creates a font object corresponding to the font specified by a PostScript or full name.
//
// Added in macOS 10.5.
// Creates a font object corresponding to the font specified by a PostScript or full name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/init(_:)-1p4b
func CGFontCreateWithFontName(name StringRef) FontRef {
	return _CGFontCreateWithFontName(name)
}

// Creates a font object from data supplied from a data provider.
//
// Added in macOS 10.5.
// Creates a font object from data supplied from a data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/init(_:)-9aour
func CGFontCreateWithDataProvider(provider DataProviderRef) FontRef {
	return _CGFontCreateWithDataProvider(provider)
}

// Returns the italic angle of a font.
//
// Added in macOS 10.5.
// Returns the italic angle of a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/italicAngle
func CGFontGetItalicAngle(font FontRef) Float {
	return _CGFontGetItalicAngle(font)
}

// Returns the leading of a font.
//
// Added in macOS 10.5.
// Returns the leading of a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/leading
func CGFontGetLeading(font FontRef) int {
	return _CGFontGetLeading(font)
}

// Returns the glyph name of the specified glyph in the specified font.
//
// Added in macOS 10.5.
// Returns the glyph name of the specified glyph in the specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/name(for:)
func CGFontCopyGlyphNameForGlyph(font FontRef, glyph Glyph) StringRef {
	return _CGFontCopyGlyphNameForGlyph(font, glyph)
}

// Returns the number of glyphs in a font.
//
// Added in macOS 10.0.
// Returns the number of glyphs in a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/numberOfGlyphs
func CGFontGetNumberOfGlyphs(font FontRef) uintptr {
	return _CGFontGetNumberOfGlyphs(font)
}

// Obtains the PostScript name of a font.
//
// Added in macOS 10.4.
// Obtains the PostScript name of a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/postScriptName
func CGFontCopyPostScriptName(font FontRef) StringRef {
	return _CGFontCopyPostScriptName(font)
}

// Returns the thickness of the dominant vertical stems of glyphs in a font.
//
// Added in macOS 10.5.
// Returns the thickness of the dominant vertical stems of glyphs in a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/stemV
func CGFontGetStemV(font FontRef) Float {
	return _CGFontGetStemV(font)
}

// Returns the font table that corresponds to the provided tag.
//
// Added in macOS 10.5.
// Returns the font table that corresponds to the provided tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/table(for:)
func CGFontCopyTableForTag(font FontRef, tag uint32) DataRef {
	return _CGFontCopyTableForTag(font, tag)
}

// Returns an array of tags that correspond to the font tables for a font.
//
// Added in macOS 10.5.
// Returns an array of tags that correspond to the font tables for a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/tableTags
func CGFontCopyTableTags(font FontRef) ArrayRef {
	return _CGFontCopyTableTags(font)
}

// Returns the Core Foundation type identifier for Core Graphics fonts.
//
// Added in macOS 10.2.
// Returns the Core Foundation type identifier for Core Graphics fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/typeID
func CGFontGetTypeID() TypeID {
	return _CGFontGetTypeID()
}

// Returns the number of glyph space units per em for the provided font.
//
// Added in macOS 10.0.
// Returns the number of glyph space units per em for the provided font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/unitsPerEm
func CGFontGetUnitsPerEm(font FontRef) int {
	return _CGFontGetUnitsPerEm(font)
}

// Returns an array of the variation axis dictionaries for a font.
//
// Added in macOS 10.4.
// Returns an array of the variation axis dictionaries for a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/variationAxes
func CGFontCopyVariationAxes(font FontRef) ArrayRef {
	return _CGFontCopyVariationAxes(font)
}

// Returns the variation specification dictionary for a font.
//
// Added in macOS 10.4.
// Returns the variation specification dictionary for a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/variations
func CGFontCopyVariations(font FontRef) DictionaryRef {
	return _CGFontCopyVariations(font)
}

// Returns the x-height of a font.
//
// Added in macOS 10.5.
// Returns the x-height of a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFont/xHeight
func CGFontGetXHeight(font FontRef) int {
	return _CGFontGetXHeight(font)
}

// Creates a font object from an Apple Type Services (ATS) font.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Creates a font object from an Apple Type Services (ATS) font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFontCreateWithPlatformFont
func CGFontCreateWithPlatformFont(platformFontReference unsafe.Pointer) FontRef {
	return _CGFontCreateWithPlatformFont(platformFontReference)
}

// Decrements the retain count of a font.
//
// Added in macOS 10.0.
// Decrements the retain count of a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFontRelease
func CGFontRelease(font FontRef) {
	_CGFontRelease(font)
}

// Increments the retain count of a font.
//
// Added in macOS 10.0.
// Increments the retain count of a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFontRetain
func CGFontRetain(font FontRef) FontRef {
	return _CGFontRetain(font)
}

// Creates a Core Graphics function.
//
// Added in macOS 10.2.
// Creates a Core Graphics function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFunction/init(info:domainDimension:domain:rangeDimension:range:callbacks:)
func CGFunctionCreate(info unsafe.Pointer, domainDimension uintptr, domain []float64, rangeDimension uintptr, range_ []float64, callbacks unsafe.Pointer) FunctionRef {
	return _CGFunctionCreate(info, domainDimension, domain, rangeDimension, range_, callbacks)
}

// Returns the type identifier for Core Graphics function objects.
//
// Added in macOS 10.2.
// Returns the type identifier for Core Graphics function objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFunction/typeID
func CGFunctionGetTypeID() TypeID {
	return _CGFunctionGetTypeID()
}

// Decrements the retain count of a function object.
//
// Added in macOS 10.2.
// Decrements the retain count of a function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFunctionRelease
func CGFunctionRelease(function FunctionRef) {
	_CGFunctionRelease(function)
}

// Increments the retain count of a function object.
//
// Added in macOS 10.2.
// Increments the retain count of a function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFunctionRetain
func CGFunctionRetain(function FunctionRef) FunctionRef {
	return _CGFunctionRetain(function)
}

// Provides a list of displays that are active for drawing.
//
// Added in macOS 10.0.
// Provides a list of displays that are active for drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGetActiveDisplayList(_:_:_:)
func CGGetActiveDisplayList(maxDisplays uint32, activeDisplays unsafe.Pointer, displayCount []uint32) Error {
	return _CGGetActiveDisplayList(maxDisplays, activeDisplays, displayCount)
}

// Gets the coefficients of the gamma transfer formula for a display.
//
// Added in macOS 10.0.
// Gets the coefficients of the gamma transfer formula for a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplayTransferByFormula(_:_:_:_:_:_:_:_:_:_:)
func CGGetDisplayTransferByFormula(display DirectDisplayID, redMin unsafe.Pointer, redMax unsafe.Pointer, redGamma unsafe.Pointer, greenMin unsafe.Pointer, greenMax unsafe.Pointer, greenGamma unsafe.Pointer, blueMin unsafe.Pointer, blueMax unsafe.Pointer, blueGamma unsafe.Pointer) Error {
	return _CGGetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
}

// Gets the values in the RGB gamma tables for a display.
//
// Added in macOS 10.0.
// Gets the values in the RGB gamma tables for a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplayTransferByTable(_:_:_:_:_:_:)
func CGGetDisplayTransferByTable(display DirectDisplayID, capacity uint32, redTable unsafe.Pointer, greenTable unsafe.Pointer, blueTable unsafe.Pointer, sampleCount []uint32) Error {
	return _CGGetDisplayTransferByTable(display, capacity, redTable, greenTable, blueTable, sampleCount)
}

// Provides a list of displays that corresponds to the bits set in an OpenGL display mask.
//
// Added in macOS 10.0.
// Provides a list of displays that corresponds to the bits set in an OpenGL display mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplaysWithOpenGLDisplayMask(_:_:_:_:)
func CGGetDisplaysWithOpenGLDisplayMask(mask OpenGLDisplayMask, maxDisplays uint32, displays unsafe.Pointer, matchingDisplayCount []uint32) Error {
	return _CGGetDisplaysWithOpenGLDisplayMask(mask, maxDisplays, displays, matchingDisplayCount)
}

// Provides a list of online displays with bounds that include the specified point.
//
// Added in macOS 10.0.
// Provides a list of online displays with bounds that include the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplaysWithPoint(_:_:_:_:)
func CGGetDisplaysWithPoint(point Point, maxDisplays uint32, displays unsafe.Pointer, matchingDisplayCount []uint32) Error {
	return _CGGetDisplaysWithPoint(point, maxDisplays, displays, matchingDisplayCount)
}

// Gets a list of online displays with bounds that intersect the specified rectangle.
//
// Added in macOS 10.0.
// Gets a list of online displays with bounds that intersect the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplaysWithRect(_:_:_:_:)
func CGGetDisplaysWithRect(rect Rect, maxDisplays uint32, displays unsafe.Pointer, matchingDisplayCount []uint32) Error {
	return _CGGetDisplaysWithRect(rect, maxDisplays, displays, matchingDisplayCount)
}

// Gets a list of currently installed event taps.
//
// Added in macOS 10.4.
// Gets a list of currently installed event taps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGetEventTapList(_:_:_:)
func CGGetEventTapList(maxNumberOfTaps uint32, tapList unsafe.Pointer, eventTapCount []uint32) Error {
	return _CGGetEventTapList(maxNumberOfTaps, tapList, eventTapCount)
}

// Reports the change in mouse position since the last mouse movement event received by the application.
//
// Added in macOS 10.0.
// Reports the change in mouse position since the last mouse movement event received by the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGetLastMouseDelta
func CGGetLastMouseDelta(deltaX unsafe.Pointer, deltaY unsafe.Pointer) {
	_CGGetLastMouseDelta(deltaX, deltaY)
}

// Provides a list of displays that are online (active, mirrored, or sleeping).
//
// Added in macOS 10.2.
// Provides a list of displays that are online (active, mirrored, or sleeping).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGetOnlineDisplayList(_:_:_:)
func CGGetOnlineDisplayList(maxDisplays uint32, onlineDisplays unsafe.Pointer, displayCount []uint32) Error {
	return _CGGetOnlineDisplayList(maxDisplays, onlineDisplays, displayCount)
}

// CGGradientGetContentHeadroom is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradient/contentHeadroom
func CGGradientGetContentHeadroom(gradient GradientRef) float32 {
	return _CGGradientGetContentHeadroom(gradient)
}

// Creates a CGGradient object from a color space and the provided color components and locations.
//
// Added in macOS 10.5.
// Creates a CGGradient object from a color space and the provided color components and locations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradient/init(colorSpace:colorComponents:locations:count:)
func CGGradientCreateWithColorComponents(space ColorSpaceRef, components []float64, locations []float64, count uintptr) GradientRef {
	return _CGGradientCreateWithColorComponents(space, components, locations, count)
}

// Creates a gradient object from a color space and the provided color objects and locations.
//
// Added in macOS 10.5.
// Creates a gradient object from a color space and the provided color objects and locations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradient/init(colorsSpace:colors:locations:)
func CGGradientCreateWithColors(space ColorSpaceRef, colors ArrayRef, locations []float64) GradientRef {
	return _CGGradientCreateWithColors(space, colors, locations)
}

// CGGradientCreateWithContentHeadroom is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradient/init(headroom:colorSpace:colorComponents:locations:count:)
func CGGradientCreateWithContentHeadroom(headroom float32, space ColorSpaceRef, components []float64, locations []float64, count uintptr) GradientRef {
	return _CGGradientCreateWithContentHeadroom(headroom, space, components, locations, count)
}

// Returns the Core Foundation type identifier for CGGradient objects.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier for CGGradient objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradient/typeID
func CGGradientGetTypeID() TypeID {
	return _CGGradientGetTypeID()
}

// Decrements the retain count of a CGGradient object.
//
// Added in macOS 10.5.
// Decrements the retain count of a CGGradient object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradientRelease
func CGGradientRelease(gradient GradientRef) {
	_CGGradientRelease(gradient)
}

// Increments the retain count of a CGGradient object.
//
// Added in macOS 10.5.
// Increments the retain count of a CGGradient object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradientRetain
func CGGradientRetain(gradient GradientRef) GradientRef {
	return _CGGradientRetain(gradient)
}

// Returns the alpha channel information for a bitmap image.
//
// Added in macOS 10.0.
// Returns the alpha channel information for a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/alphaInfo
func CGImageGetAlphaInfo(image ImageRef) ImageAlphaInfo {
	return _CGImageGetAlphaInfo(image)
}

// Returns the bitmap information for a bitmap image.
//
// Added in macOS 10.4.
// Returns the bitmap information for a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/bitmapInfo
func CGImageGetBitmapInfo(image ImageRef) BitmapInfo {
	return _CGImageGetBitmapInfo(image)
}

// Returns the number of bits allocated for a single color component of a bitmap image.
//
// Added in macOS 10.0.
// Returns the number of bits allocated for a single color component of a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/bitsPerComponent
func CGImageGetBitsPerComponent(image ImageRef) uintptr {
	return _CGImageGetBitsPerComponent(image)
}

// Returns the number of bits allocated for a single pixel in a bitmap image.
//
// Added in macOS 10.0.
// Returns the number of bits allocated for a single pixel in a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/bitsPerPixel
func CGImageGetBitsPerPixel(image ImageRef) uintptr {
	return _CGImageGetBitsPerPixel(image)
}

// CGImageGetByteOrderInfo is a CoreGraphics function.
//
// Added in macOS 10.14.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/byteOrderInfo
func CGImageGetByteOrderInfo(image ImageRef) ImageByteOrderInfo {
	return _CGImageGetByteOrderInfo(image)
}

// Returns the number of bytes allocated for a single row of a bitmap image.
//
// Added in macOS 10.0.
// Returns the number of bytes allocated for a single row of a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/bytesPerRow
func CGImageGetBytesPerRow(image ImageRef) uintptr {
	return _CGImageGetBytesPerRow(image)
}

// CGImageCalculateContentAverageLightLevel is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/calculatedContentAverageLightLevel
func CGImageCalculateContentAverageLightLevel(image ImageRef) float32 {
	return _CGImageCalculateContentAverageLightLevel(image)
}

// CGImageCalculateContentHeadroom is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/calculatedContentHeadroom
func CGImageCalculateContentHeadroom(image ImageRef) float32 {
	return _CGImageCalculateContentHeadroom(image)
}

// Return the color space for a bitmap image.
//
// Added in macOS 10.0.
// Return the color space for a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/colorSpace
func CGImageGetColorSpace(image ImageRef) ColorSpaceRef {
	return _CGImageGetColorSpace(image)
}

// CGImageContainsImageSpecificToneMappingMetadata is a CoreGraphics function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/containsImageSpecificToneMappingMetadata
func CGImageContainsImageSpecificToneMappingMetadata(image ImageRef) bool {
	return _CGImageContainsImageSpecificToneMappingMetadata(image)
}

// CGImageGetContentAverageLightLevel is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/contentAverageLightLevel
func CGImageGetContentAverageLightLevel(image ImageRef) float32 {
	return _CGImageGetContentAverageLightLevel(image)
}

// CGImageGetContentHeadroom is a CoreGraphics function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/contentHeadroom
func CGImageGetContentHeadroom(image ImageRef) float32 {
	return _CGImageGetContentHeadroom(image)
}

// Creates a copy of a bitmap image.
//
// Added in macOS 10.4.
// Creates a copy of a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/copy()
func CGImageCreateCopy(image ImageRef) ImageRef {
	return _CGImageCreateCopy(image)
}

// Creates a copy of a bitmap image, replacing its colorspace.
//
// Added in macOS 10.3.
// Creates a copy of a bitmap image, replacing its colorspace.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/copy(colorSpace:)
func CGImageCreateCopyWithColorSpace(image ImageRef, space ColorSpaceRef) ImageRef {
	return _CGImageCreateCopyWithColorSpace(image, space)
}

// CGImageCreateCopyWithContentAverageLightLevel is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/copy(contentAverageLightLevel:)
func CGImageCreateCopyWithContentAverageLightLevel(image ImageRef, avll float32) ImageRef {
	return _CGImageCreateCopyWithContentAverageLightLevel(image, avll)
}

// CGImageCreateCopyWithCalculatedHDRStats is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/copyWithCalculatedHDRStats()
func CGImageCreateCopyWithCalculatedHDRStats(image ImageRef) ImageRef {
	return _CGImageCreateCopyWithCalculatedHDRStats(image)
}

// Creates a bitmap image using the data contained within a subregion of an existing bitmap image.
//
// Added in macOS 10.4.
// Creates a bitmap image using the data contained within a subregion of an existing bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/cropping(to:)
func CGImageCreateWithImageInRect(image ImageRef, rect Rect) ImageRef {
	return _CGImageCreateWithImageInRect(image, rect)
}

// Returns the data provider for a bitmap image or image mask.
//
// Added in macOS 10.0.
// Returns the data provider for a bitmap image or image mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/dataProvider
func CGImageGetDataProvider(image ImageRef) DataProviderRef {
	return _CGImageGetDataProvider(image)
}

// Returns the decode array for a bitmap image.
//
// Added in macOS 10.0.
// Returns the decode array for a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/decode
func CGImageGetDecode(image ImageRef) []float64 {
	return _CGImageGetDecode(image)
}

// Returns the height of a bitmap image.
//
// Added in macOS 10.0.
// Returns the height of a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/height
func CGImageGetHeight(image ImageRef) uintptr {
	return _CGImageGetHeight(image)
}

// CGImageCreateWithContentHeadroom is a CoreGraphics function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(headroom:width:height:bitsPerComponent:bitsPerPixel:bytesPerRow:space:bitmapInfo:provider:decode:shouldInterpolate:intent:)
func CGImageCreateWithContentHeadroom(headroom float32, width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space ColorSpaceRef, bitmapInfo BitmapInfo, provider DataProviderRef, decode []float64, shouldInterpolate bool, intent ColorRenderingIntent) ImageRef {
	return _CGImageCreateWithContentHeadroom(headroom, width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, bitmapInfo, provider, decode, shouldInterpolate, intent)
}

// Creates a bitmap image using JPEG-encoded data supplied by a data provider.
//
// Added in macOS 10.1.
// Creates a bitmap image using JPEG-encoded data supplied by a data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(jpegDataProviderSource:decode:shouldInterpolate:intent:)
func CGImageCreateWithJPEGDataProvider(source DataProviderRef, decode []float64, shouldInterpolate bool, intent ColorRenderingIntent) ImageRef {
	return _CGImageCreateWithJPEGDataProvider(source, decode, shouldInterpolate, intent)
}

// Creates a bitmap image mask from data supplied by a data provider.
//
// Added in macOS 10.0.
// Creates a bitmap image mask from data supplied by a data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(maskWidth:height:bitsPerComponent:bitsPerPixel:bytesPerRow:provider:decode:shouldInterpolate:)
func CGImageMaskCreate(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, provider DataProviderRef, decode []float64, shouldInterpolate bool) ImageRef {
	return _CGImageMaskCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, provider, decode, shouldInterpolate)
}

// Creates a bitmap image using PNG-encoded data supplied by a data provider.
//
// Added in macOS 10.2.
// Creates a bitmap image using PNG-encoded data supplied by a data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(pngDataProviderSource:decode:shouldInterpolate:intent:)
func CGImageCreateWithPNGDataProvider(source DataProviderRef, decode []float64, shouldInterpolate bool, intent ColorRenderingIntent) ImageRef {
	return _CGImageCreateWithPNGDataProvider(source, decode, shouldInterpolate, intent)
}

// Creates a bitmap image from data supplied by a data provider.
//
// Added in macOS 10.0.
// Creates a bitmap image from data supplied by a data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(width:height:bitsPerComponent:bitsPerPixel:bytesPerRow:space:bitmapInfo:provider:decode:shouldInterpolate:intent:)
func CGImageCreate(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space ColorSpaceRef, bitmapInfo BitmapInfo, provider DataProviderRef, decode []float64, shouldInterpolate bool, intent ColorRenderingIntent) ImageRef {
	return _CGImageCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, bitmapInfo, provider, decode, shouldInterpolate, intent)
}

// Returns a composite image of the specified windows.

// Returns a composite image of the specified windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(windowListFromArrayScreenBounds:windowArray:imageOption:)
func CGWindowListCreateImageFromArray(screenBounds Rect, windowArray ArrayRef, imageOption WindowImageOption) ImageRef {
	return _CGWindowListCreateImageFromArray(screenBounds, windowArray, imageOption)
}

// Returns whether a bitmap image is an image mask.
//
// Added in macOS 10.0.
// Returns whether a bitmap image is an image mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/isMask
func CGImageIsMask(image ImageRef) bool {
	return _CGImageIsMask(image)
}

// Creates a bitmap image from an existing image and an image mask.
//
// Added in macOS 10.4.
// Creates a bitmap image from an existing image and an image mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/masking(_:)
func CGImageCreateWithMask(image ImageRef, mask ImageRef) ImageRef {
	return _CGImageCreateWithMask(image, mask)
}

// CGImageGetPixelFormatInfo is a CoreGraphics function.
//
// Added in macOS 10.14.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/pixelFormatInfo
func CGImageGetPixelFormatInfo(image ImageRef) ImagePixelFormatInfo {
	return _CGImageGetPixelFormatInfo(image)
}

// Returns the rendering intent setting for a bitmap image.
//
// Added in macOS 10.0.
// Returns the rendering intent setting for a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/renderingIntent
func CGImageGetRenderingIntent(image ImageRef) ColorRenderingIntent {
	return _CGImageGetRenderingIntent(image)
}

// Returns the interpolation setting for a bitmap image.
//
// Added in macOS 10.0.
// Returns the interpolation setting for a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/shouldInterpolate
func CGImageGetShouldInterpolate(image ImageRef) bool {
	return _CGImageGetShouldInterpolate(image)
}

// CGImageShouldToneMap is a CoreGraphics function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/shouldToneMap
func CGImageShouldToneMap(image ImageRef) bool {
	return _CGImageShouldToneMap(image)
}

// Returns the type identifier for CGImage objects.
//
// Added in macOS 10.2.
// Returns the type identifier for CGImage objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/typeID
func CGImageGetTypeID() TypeID {
	return _CGImageGetTypeID()
}

// The Universal Type Identifier for the image.
//
// Added in macOS 10.11.
// The Universal Type Identifier for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/utType
func CGImageGetUTType(image ImageRef) StringRef {
	return _CGImageGetUTType(image)
}

// Returns the width of a bitmap image, in pixels.
//
// Added in macOS 10.0.
// Returns the width of a bitmap image, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage/width
func CGImageGetWidth(image ImageRef) uintptr {
	return _CGImageGetWidth(image)
}

// CGImageCreateCopyWithContentHeadroom is a CoreGraphics function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageCreateCopyWithContentHeadroom(_:_:)
func CGImageCreateCopyWithContentHeadroom(headroom float32, image ImageRef) ImageRef {
	return _CGImageCreateCopyWithContentHeadroom(headroom, image)
}

// Creates a bitmap image by masking an existing bitmap image with the provided color values.
//
// Added in macOS 10.4.
// Creates a bitmap image by masking an existing bitmap image with the provided color values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageCreateWithMaskingColors
func CGImageCreateWithMaskingColors(image ImageRef, components []float64) ImageRef {
	return _CGImageCreateWithMaskingColors(image, components)
}

// Decrements the retain count of a bitmap image.
//
// Added in macOS 10.0.
// Decrements the retain count of a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageRelease
func CGImageRelease(image ImageRef) {
	_CGImageRelease(image)
}

// Increments the retain count of a bitmap image.
//
// Added in macOS 10.0.
// Increments the retain count of a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageRetain
func CGImageRetain(image ImageRef) ImageRef {
	return _CGImageRetain(image)
}

// Turns off local hardware events in the current session.

// Turns off local hardware events in the current session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInhibitLocalEvents(_:)
func CGInhibitLocalEvents(inhibit unsafe.Pointer) Error {
	return _CGInhibitLocalEvents(inhibit)
}

// Returns the graphics context associated with a layer object.
//
// Added in macOS 10.4.
// Returns the graphics context associated with a layer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLayer/context
func CGLayerGetContext(layer LayerRef) ContextRef {
	return _CGLayerGetContext(layer)
}

// Creates a layer object that is associated with a graphics context.
//
// Added in macOS 10.4.
// Creates a layer object that is associated with a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLayer/init(_:size:auxiliaryInfo:)
func CGLayerCreateWithContext(context ContextRef, size Size, auxiliaryInfo DictionaryRef) LayerRef {
	return _CGLayerCreateWithContext(context, size, auxiliaryInfo)
}

// Returns the width and height of a layer object.
//
// Added in macOS 10.4.
// Returns the width and height of a layer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLayer/size
func CGLayerGetSize(layer LayerRef) Size {
	return _CGLayerGetSize(layer)
}

// Returns the unique type identifier used for objects.
//
// Added in macOS 10.4.
// Returns the unique type identifier used for objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLayer/typeID
func CGLayerGetTypeID() TypeID {
	return _CGLayerGetTypeID()
}

// Decrements the retain count of a layer object.
//
// Added in macOS 10.4.
// Decrements the retain count of a layer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLayerRelease
func CGLayerRelease(layer LayerRef) {
	_CGLayerRelease(layer)
}

// Increments the retain count of a layer object.
//
// Added in macOS 10.4.
// Increments the retain count of a layer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLayerRetain
func CGLayerRetain(layer LayerRef) LayerRef {
	return _CGLayerRetain(layer)
}

// Returns the display ID of the main display.
//
// Added in macOS 10.2.
// Returns the display ID of the main display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMainDisplayID()
func CGMainDisplayID() DirectDisplayID {
	return _CGMainDisplayID()
}

// Closes and completes a subpath in a mutable graphics path.
//
// Added in macOS 10.2.
// Closes and completes a subpath in a mutable graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMutablePath/closeSubpath()
func CGPathCloseSubpath(path MutablePathRef) {
	_CGPathCloseSubpath(path)
}

// Creates a mutable graphics path.
//
// Added in macOS 10.2.
// Creates a mutable graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMutablePath/init()
func CGPathCreateMutable() MutablePathRef {
	return _CGPathCreateMutable()
}

// Maps an OpenGL display mask to a display ID.
//
// Added in macOS 10.2.
// Maps an OpenGL display mask to a display ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGOpenGLDisplayMaskToDisplayID(_:)
func CGOpenGLDisplayMaskToDisplayID(mask OpenGLDisplayMask) DirectDisplayID {
	return _CGOpenGLDisplayMaskToDisplayID(mask)
}

// CGPDFArrayApplyBlock is a CoreGraphics function.
//
// Added in macOS 10.14.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayApplyBlock(_:_:_:)
func CGPDFArrayApplyBlock(array PDFArrayRef, block PDFArrayApplierBlock, info unsafe.Pointer) {
	_CGPDFArrayApplyBlock(array, block, info)
}

// Returns whether an object at a given index in a PDF array is another PDF array and, if so, retrieves that array.
//
// Added in macOS 10.3.
// Returns whether an object at a given index in a PDF array is another PDF array and, if so, retrieves that array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetArray(_:_:_:)
func CGPDFArrayGetArray(array PDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetArray(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF Boolean and, if so, retrieves that Boolean.
//
// Added in macOS 10.3.
// Returns whether an object at a given index in a PDF array is a PDF Boolean and, if so, retrieves that Boolean.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetBoolean(_:_:_:)
func CGPDFArrayGetBoolean(array PDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetBoolean(array, index, value)
}

// Returns the number of items in a PDF array.
//
// Added in macOS 10.3.
// Returns the number of items in a PDF array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetCount(_:)
func CGPDFArrayGetCount(array PDFArrayRef) uintptr {
	return _CGPDFArrayGetCount(array)
}

// Returns whether an object at a given index in a PDF array is a PDF dictionary and, if so, retrieves that dictionary.
//
// Added in macOS 10.3.
// Returns whether an object at a given index in a PDF array is a PDF dictionary and, if so, retrieves that dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetDictionary(_:_:_:)
func CGPDFArrayGetDictionary(array PDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetDictionary(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF integer and, if so, retrieves that object.
//
// Added in macOS 10.3.
// Returns whether an object at a given index in a PDF array is a PDF integer and, if so, retrieves that object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetInteger(_:_:_:)
func CGPDFArrayGetInteger(array PDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetInteger(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF name reference (represented as a constant C string) and, if so, retrieves that name.
//
// Added in macOS 10.3.
// Returns whether an object at a given index in a PDF array is a PDF name reference (represented as a constant C string) and, if so, retrieves that name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetName(_:_:_:)
func CGPDFArrayGetName(array PDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetName(array, index, value)
}

// Returns whether an object at a given index in a Quartz PDF array is a PDF null.
//
// Added in macOS 10.3.
// Returns whether an object at a given index in a Quartz PDF array is a PDF null.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetNull(_:_:)
func CGPDFArrayGetNull(array PDFArrayRef, index uintptr) bool {
	return _CGPDFArrayGetNull(array, index)
}

// Returns whether an object at a given index in a PDF array is a PDF number and, if so, retrieves that object.
//
// Added in macOS 10.3.
// Returns whether an object at a given index in a PDF array is a PDF number and, if so, retrieves that object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetNumber(_:_:_:)
func CGPDFArrayGetNumber(array PDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetNumber(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF object and, if so, retrieves that object.
//
// Added in macOS 10.3.
// Returns whether an object at a given index in a PDF array is a PDF object and, if so, retrieves that object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetObject(_:_:_:)
func CGPDFArrayGetObject(array PDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetObject(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF stream and, if so, retrieves that stream.
//
// Added in macOS 10.3.
// Returns whether an object at a given index in a PDF array is a PDF stream and, if so, retrieves that stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetStream(_:_:_:)
func CGPDFArrayGetStream(array PDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetStream(array, index, value)
}

// Returns whether an object at a given index in a PDF array is a PDF string and, if so, retrieves that string.
//
// Added in macOS 10.3.
// Returns whether an object at a given index in a PDF array is a PDF string and, if so, retrieves that string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetString(_:_:_:)
func CGPDFArrayGetString(array PDFArrayRef, index uintptr, value unsafe.Pointer) bool {
	return _CGPDFArrayGetString(array, index, value)
}

// Creates a content stream object from a PDF page object.
//
// Added in macOS 10.4.
// Creates a content stream object from a PDF page object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamCreateWithPage(_:)
func CGPDFContentStreamCreateWithPage(page PDFPageRef) PDFContentStreamRef {
	return _CGPDFContentStreamCreateWithPage(page)
}

// Creates a PDF content stream object from an existing PDF content stream object.
//
// Added in macOS 10.4.
// Creates a PDF content stream object from an existing PDF content stream object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamCreateWithStream(_:_:_:)
func CGPDFContentStreamCreateWithStream(stream PDFStreamRef, streamResources PDFDictionaryRef, parent PDFContentStreamRef) PDFContentStreamRef {
	return _CGPDFContentStreamCreateWithStream(stream, streamResources, parent)
}

// Gets the specified resource from a PDF content stream object.
//
// Added in macOS 10.4.
// Gets the specified resource from a PDF content stream object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamGetResource(_:_:_:)
func CGPDFContentStreamGetResource(cs PDFContentStreamRef, category unsafe.Pointer, name unsafe.Pointer) PDFObjectRef {
	return _CGPDFContentStreamGetResource(cs, category, name)
}

// Gets the array of PDF content streams contained in a PDF content stream object.
//
// Added in macOS 10.4.
// Gets the array of PDF content streams contained in a PDF content stream object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamGetStreams(_:)
func CGPDFContentStreamGetStreams(cs PDFContentStreamRef) ArrayRef {
	return _CGPDFContentStreamGetStreams(cs)
}

// Decrements the retain count of a PDF content stream object.
//
// Added in macOS 10.4.
// Decrements the retain count of a PDF content stream object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamRelease(_:)
func CGPDFContentStreamRelease(cs PDFContentStreamRef) {
	_CGPDFContentStreamRelease(cs)
}

// Increments the retain count of a PDF content stream object.
//
// Added in macOS 10.4.
// Increments the retain count of a PDF content stream object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamRetain(_:)
func CGPDFContentStreamRetain(cs PDFContentStreamRef) PDFContentStreamRef {
	return _CGPDFContentStreamRetain(cs)
}

// CGPDFContextBeginTag is a CoreGraphics function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextBeginTag(_:_:_:)
func CGPDFContextBeginTag(context ContextRef, tagType PDFTagType, tagProperties DictionaryRef) {
	_CGPDFContextBeginTag(context, tagType, tagProperties)
}

// CGPDFContextEndTag is a CoreGraphics function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextEndTag(_:)
func CGPDFContextEndTag(context ContextRef) {
	_CGPDFContextEndTag(context)
}

// CGPDFContextSetIDTree is a CoreGraphics function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetIDTree(_:_:)
func CGPDFContextSetIDTree(context ContextRef, IDTreeDictionary PDFDictionaryRef) {
	_CGPDFContextSetIDTree(context, IDTreeDictionary)
}

// CGPDFContextSetOutline is a CoreGraphics function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetOutline(_:_:)
func CGPDFContextSetOutline(context ContextRef, outline DictionaryRef) {
	_CGPDFContextSetOutline(context, outline)
}

// CGPDFContextSetPageTagStructureTree is a CoreGraphics function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetPageTagStructureTree(_:_:)
func CGPDFContextSetPageTagStructureTree(context ContextRef, pageTagStructureTreeDictionary DictionaryRef) {
	_CGPDFContextSetPageTagStructureTree(context, pageTagStructureTreeDictionary)
}

// CGPDFContextSetParentTree is a CoreGraphics function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetParentTree(_:_:)
func CGPDFContextSetParentTree(context ContextRef, parentTreeDictionary PDFDictionaryRef) {
	_CGPDFContextSetParentTree(context, parentTreeDictionary)
}

// CGPDFDictionaryApplyBlock is a CoreGraphics function.
//
// Added in macOS 10.14.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryApplyBlock(_:_:_:)
func CGPDFDictionaryApplyBlock(dict PDFDictionaryRef, block PDFDictionaryApplierBlock, info unsafe.Pointer) {
	_CGPDFDictionaryApplyBlock(dict, block, info)
}

// Applies a function to each entry in a dictionary.
//
// Added in macOS 10.3.
// Applies a function to each entry in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryApplyFunction(_:_:_:)
func CGPDFDictionaryApplyFunction(dict PDFDictionaryRef, function PDFDictionaryApplierFunction, info unsafe.Pointer) {
	_CGPDFDictionaryApplyFunction(dict, function, info)
}

// Returns whether there is a PDF array associated with a specified key in a PDF dictionary and, if so, retrieves that array.
//
// Added in macOS 10.3.
// Returns whether there is a PDF array associated with a specified key in a PDF dictionary and, if so, retrieves that array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetArray(_:_:_:)
func CGPDFDictionaryGetArray(dict PDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetArray(dict, key, value)
}

// Returns whether there is a PDF Boolean value associated with a specified key in a PDF dictionary and, if so, retrieves the Boolean value.
//
// Added in macOS 10.3.
// Returns whether there is a PDF Boolean value associated with a specified key in a PDF dictionary and, if so, retrieves the Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetBoolean(_:_:_:)
func CGPDFDictionaryGetBoolean(dict PDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetBoolean(dict, key, value)
}

// Returns the number of entries in a PDF dictionary.
//
// Added in macOS 10.3.
// Returns the number of entries in a PDF dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetCount(_:)
func CGPDFDictionaryGetCount(dict PDFDictionaryRef) uintptr {
	return _CGPDFDictionaryGetCount(dict)
}

// Returns whether there is another PDF dictionary associated with a specified key in a PDF dictionary and, if so, retrieves that dictionary.
//
// Added in macOS 10.3.
// Returns whether there is another PDF dictionary associated with a specified key in a PDF dictionary and, if so, retrieves that dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetDictionary(_:_:_:)
func CGPDFDictionaryGetDictionary(dict PDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetDictionary(dict, key, value)
}

// Returns whether there is a PDF integer associated with a specified key in a PDF dictionary and, if so, retrieves that integer.
//
// Added in macOS 10.3.
// Returns whether there is a PDF integer associated with a specified key in a PDF dictionary and, if so, retrieves that integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetInteger(_:_:_:)
func CGPDFDictionaryGetInteger(dict PDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetInteger(dict, key, value)
}

// Returns whether an object with a specified key in a PDF dictionary is a PDF name reference (represented as a constant C string) and, if so, retrieves that name.
//
// Added in macOS 10.3.
// Returns whether an object with a specified key in a PDF dictionary is a PDF name reference (represented as a constant C string) and, if so, retrieves that name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetName(_:_:_:)
func CGPDFDictionaryGetName(dict PDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetName(dict, key, value)
}

// Returns whether there is a PDF number associated with a specified key in a PDF dictionary and, if so, retrieves that number.
//
// Added in macOS 10.3.
// Returns whether there is a PDF number associated with a specified key in a PDF dictionary and, if so, retrieves that number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetNumber(_:_:_:)
func CGPDFDictionaryGetNumber(dict PDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetNumber(dict, key, value)
}

// Returns whether there is a PDF object associated with a specified key in a PDF dictionary and, if so, retrieves that object.
//
// Added in macOS 10.3.
// Returns whether there is a PDF object associated with a specified key in a PDF dictionary and, if so, retrieves that object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetObject(_:_:_:)
func CGPDFDictionaryGetObject(dict PDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetObject(dict, key, value)
}

// Returns whether there is a PDF stream associated with a specified key in a PDF dictionary and, if so, retrieves that stream.
//
// Added in macOS 10.3.
// Returns whether there is a PDF stream associated with a specified key in a PDF dictionary and, if so, retrieves that stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetStream(_:_:_:)
func CGPDFDictionaryGetStream(dict PDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetStream(dict, key, value)
}

// Returns whether there is a PDF string associated with a specified key in a PDF dictionary and, if so, retrieves that string.
//
// Added in macOS 10.3.
// Returns whether there is a PDF string associated with a specified key in a PDF dictionary and, if so, retrieves that string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetString(_:_:_:)
func CGPDFDictionaryGetString(dict PDFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _CGPDFDictionaryGetString(dict, key, value)
}

// CGPDFDocumentGetAccessPermissions is a CoreGraphics function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/accessPermissions
func CGPDFDocumentGetAccessPermissions(document PDFDocumentRef) PDFAccessPermissions {
	return _CGPDFDocumentGetAccessPermissions(document)
}

// Returns whether the specified PDF document allows copying.
//
// Added in macOS 10.2.
// Returns whether the specified PDF document allows copying.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/allowsCopying
func CGPDFDocumentAllowsCopying(document PDFDocumentRef) bool {
	return _CGPDFDocumentAllowsCopying(document)
}

// Returns whether a PDF document allows printing.
//
// Added in macOS 10.2.
// Returns whether a PDF document allows printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/allowsPrinting
func CGPDFDocumentAllowsPrinting(document PDFDocumentRef) bool {
	return _CGPDFDocumentAllowsPrinting(document)
}

// Returns the document catalog of a Core Graphics PDF document.
//
// Added in macOS 10.3.
// Returns the document catalog of a Core Graphics PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/catalog
func CGPDFDocumentGetCatalog(document PDFDocumentRef) PDFDictionaryRef {
	return _CGPDFDocumentGetCatalog(document)
}

// Gets the file identifier for a PDF document.
//
// Added in macOS 10.4.
// Gets the file identifier for a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/fileIdentifier
func CGPDFDocumentGetID(document PDFDocumentRef) PDFArrayRef {
	return _CGPDFDocumentGetID(document)
}

// Returns the major and minor version numbers of a Core Graphics PDF document.
//
// Added in macOS 10.3.
// Returns the major and minor version numbers of a Core Graphics PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/getVersion(majorVersion:minorVersion:)
func CGPDFDocumentGetVersion(document PDFDocumentRef, majorVersion []int, minorVersion []int) {
	_CGPDFDocumentGetVersion(document, majorVersion, minorVersion)
}

// Gets the information dictionary for a PDF document.
//
// Added in macOS 10.4.
// Gets the information dictionary for a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/info
func CGPDFDocumentGetInfo(document PDFDocumentRef) PDFDictionaryRef {
	return _CGPDFDocumentGetInfo(document)
}

// Creates a Core Graphics PDF document using data specified by a URL.
//
// Added in macOS 10.0.
// Creates a Core Graphics PDF document using data specified by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/init(_:)-2gtsd
func CGPDFDocumentCreateWithURL(url URLRef) PDFDocumentRef {
	return _CGPDFDocumentCreateWithURL(url)
}

// Creates a Core Graphics PDF document using a data provider.
//
// Added in macOS 10.0.
// Creates a Core Graphics PDF document using a data provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/init(_:)-gbq6
func CGPDFDocumentCreateWithProvider(provider DataProviderRef) PDFDocumentRef {
	return _CGPDFDocumentCreateWithProvider(provider)
}

// Returns whether the specified PDF file is encrypted.
//
// Added in macOS 10.2.
// Returns whether the specified PDF file is encrypted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/isEncrypted
func CGPDFDocumentIsEncrypted(document PDFDocumentRef) bool {
	return _CGPDFDocumentIsEncrypted(document)
}

// Returns whether the specified PDF document is currently unlocked.
//
// Added in macOS 10.2.
// Returns whether the specified PDF document is currently unlocked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/isUnlocked
func CGPDFDocumentIsUnlocked(document PDFDocumentRef) bool {
	return _CGPDFDocumentIsUnlocked(document)
}

// Returns the number of pages in a PDF document.
//
// Added in macOS 10.0.
// Returns the number of pages in a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/numberOfPages
func CGPDFDocumentGetNumberOfPages(document PDFDocumentRef) uintptr {
	return _CGPDFDocumentGetNumberOfPages(document)
}

// CGPDFDocumentGetOutline is a CoreGraphics function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/outline
func CGPDFDocumentGetOutline(document PDFDocumentRef) DictionaryRef {
	return _CGPDFDocumentGetOutline(document)
}

// Returns a page from a Core Graphics PDF document.
//
// Added in macOS 10.3.
// Returns a page from a Core Graphics PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/page(at:)
func CGPDFDocumentGetPage(document PDFDocumentRef, pageNumber uintptr) PDFPageRef {
	return _CGPDFDocumentGetPage(document, pageNumber)
}

// Returns the type identifier for Core Graphics PDF documents.
//
// Added in macOS 10.2.
// Returns the type identifier for Core Graphics PDF documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/typeID
func CGPDFDocumentGetTypeID() TypeID {
	return _CGPDFDocumentGetTypeID()
}

// Unlocks an encrypted PDF document when a valid password is supplied.
//
// Added in macOS 10.2.
// Unlocks an encrypted PDF document when a valid password is supplied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/unlockWithPassword(_:)
func CGPDFDocumentUnlockWithPassword(document PDFDocumentRef, password unsafe.Pointer) bool {
	return _CGPDFDocumentUnlockWithPassword(document, password)
}

// Returns the art box of a page in a PDF document.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// Returns the art box of a page in a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetArtBox
func CGPDFDocumentGetArtBox(document PDFDocumentRef, page int) Rect {
	return _CGPDFDocumentGetArtBox(document, page)
}

// Returns the bleed box of a page in a PDF document.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// Returns the bleed box of a page in a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetBleedBox
func CGPDFDocumentGetBleedBox(document PDFDocumentRef, page int) Rect {
	return _CGPDFDocumentGetBleedBox(document, page)
}

// Returns the crop box of a page in a PDF document.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// Returns the crop box of a page in a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetCropBox
func CGPDFDocumentGetCropBox(document PDFDocumentRef, page int) Rect {
	return _CGPDFDocumentGetCropBox(document, page)
}

// Returns the media box of a page in a PDF document.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// Returns the media box of a page in a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetMediaBox
func CGPDFDocumentGetMediaBox(document PDFDocumentRef, page int) Rect {
	return _CGPDFDocumentGetMediaBox(document, page)
}

// Returns the rotation angle of a page in a PDF document.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// Returns the rotation angle of a page in a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetRotationAngle
func CGPDFDocumentGetRotationAngle(document PDFDocumentRef, page int) int {
	return _CGPDFDocumentGetRotationAngle(document, page)
}

// Returns the trim box of a page in a PDF document.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// Returns the trim box of a page in a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetTrimBox
func CGPDFDocumentGetTrimBox(document PDFDocumentRef, page int) Rect {
	return _CGPDFDocumentGetTrimBox(document, page)
}

// Decrements the retain count of a PDF document.
//
// Added in macOS 10.0.
// Decrements the retain count of a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentRelease
func CGPDFDocumentRelease(document PDFDocumentRef) {
	_CGPDFDocumentRelease(document)
}

// Increments the retain count of a Core Graphics PDF document.
//
// Added in macOS 10.0.
// Increments the retain count of a Core Graphics PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentRetain
func CGPDFDocumentRetain(document PDFDocumentRef) PDFDocumentRef {
	return _CGPDFDocumentRetain(document)
}

// Returns the PDF type identifier of an object.
//
// Added in macOS 10.3.
// Returns the PDF type identifier of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectGetType(_:)
func CGPDFObjectGetType(object PDFObjectRef) PDFObjectType {
	return _CGPDFObjectGetType(object)
}

// Returns whether an object is of a given type and if it is, retrieves its value.
//
// Added in macOS 10.3.
// Returns whether an object is of a given type and if it is, retrieves its value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectGetValue(_:_:_:)
func CGPDFObjectGetValue(object PDFObjectRef, type_ PDFObjectType, value unsafe.Pointer) bool {
	return _CGPDFObjectGetValue(object, type_, value)
}

// Creates an empty PDF operator table.
//
// Added in macOS 10.4.
// Creates an empty PDF operator table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableCreate()
func CGPDFOperatorTableCreate() PDFOperatorTableRef {
	return _CGPDFOperatorTableCreate()
}

// Decrements the retain count of a CGPDFOperatorTable object.
//
// Added in macOS 10.4.
// Decrements the retain count of a CGPDFOperatorTable object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableRelease(_:)
func CGPDFOperatorTableRelease(table PDFOperatorTableRef) {
	_CGPDFOperatorTableRelease(table)
}

// Increments the retain count of a CGPDFOperatorTable object.
//
// Added in macOS 10.4.
// Increments the retain count of a CGPDFOperatorTable object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableRetain(_:)
func CGPDFOperatorTableRetain(table PDFOperatorTableRef) PDFOperatorTableRef {
	return _CGPDFOperatorTableRetain(table)
}

// Sets a callback function for a PDF operator.
//
// Added in macOS 10.4.
// Sets a callback function for a PDF operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableSetCallback(_:_:_:)
func CGPDFOperatorTableSetCallback(table PDFOperatorTableRef, name unsafe.Pointer, callback PDFOperatorCallback) {
	_CGPDFOperatorTableSetCallback(table, name, callback)
}

// Returns the dictionary of a PDF page.
//
// Added in macOS 10.3.
// Returns the dictionary of a PDF page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/dictionary
func CGPDFPageGetDictionary(page PDFPageRef) PDFDictionaryRef {
	return _CGPDFPageGetDictionary(page)
}

// Returns the document for a page.
//
// Added in macOS 10.3.
// Returns the document for a page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/document
func CGPDFPageGetDocument(page PDFPageRef) PDFDocumentRef {
	return _CGPDFPageGetDocument(page)
}

// Returns the rectangle that represents a type of box for a content region or page dimensions of a PDF page.
//
// Added in macOS 10.3.
// Returns the rectangle that represents a type of box for a content region or page dimensions of a PDF page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/getBoxRect(_:)
func CGPDFPageGetBoxRect(page PDFPageRef, box PDFBox) Rect {
	return _CGPDFPageGetBoxRect(page, box)
}

// Returns the affine transform that maps a box to a given rectangle on a PDF page.
//
// Added in macOS 10.3.
// Returns the affine transform that maps a box to a given rectangle on a PDF page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/getDrawingTransform(_:rect:rotate:preserveAspectRatio:)
func CGPDFPageGetDrawingTransform(page PDFPageRef, box PDFBox, rect Rect, rotate int, preserveAspectRatio bool) AffineTransform {
	return _CGPDFPageGetDrawingTransform(page, box, rect, rotate, preserveAspectRatio)
}

// Returns the page number of the specified PDF page.
//
// Added in macOS 10.3.
// Returns the page number of the specified PDF page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/pageNumber
func CGPDFPageGetPageNumber(page PDFPageRef) uintptr {
	return _CGPDFPageGetPageNumber(page)
}

// Returns the rotation angle of a PDF page, in degrees.
//
// Added in macOS 10.3.
// Returns the rotation angle of a PDF page, in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/rotationAngle
func CGPDFPageGetRotationAngle(page PDFPageRef) int {
	return _CGPDFPageGetRotationAngle(page)
}

// Returns the CFType ID for PDF page objects.
//
// Added in macOS 10.3.
// Returns the CFType ID for PDF page objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/typeID
func CGPDFPageGetTypeID() TypeID {
	return _CGPDFPageGetTypeID()
}

// Decrements the retain count of a PDF page.
//
// Added in macOS 10.3.
// Decrements the retain count of a PDF page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFPageRelease
func CGPDFPageRelease(page PDFPageRef) {
	_CGPDFPageRelease(page)
}

// Increments the retain count of a PDF page.
//
// Added in macOS 10.3.
// Increments the retain count of a PDF page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFPageRetain
func CGPDFPageRetain(page PDFPageRef) PDFPageRef {
	return _CGPDFPageRetain(page)
}

// Creates a PDF scanner.
//
// Added in macOS 10.4.
// Creates a PDF scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerCreate(_:_:_:)
func CGPDFScannerCreate(cs PDFContentStreamRef, table PDFOperatorTableRef, info unsafe.Pointer) PDFScannerRef {
	return _CGPDFScannerCreate(cs, table, info)
}

// Returns the content stream associated with a PDF scanner object.
//
// Added in macOS 10.4.
// Returns the content stream associated with a PDF scanner object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerGetContentStream(_:)
func CGPDFScannerGetContentStream(scanner PDFScannerRef) PDFContentStreamRef {
	return _CGPDFScannerGetContentStream(scanner)
}

// Retrieves an array object from the scanner stack.
//
// Added in macOS 10.4.
// Retrieves an array object from the scanner stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopArray(_:_:)
func CGPDFScannerPopArray(scanner PDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopArray(scanner, value)
}

// Retrieves a Boolean object from the scanner stack.
//
// Added in macOS 10.4.
// Retrieves a Boolean object from the scanner stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopBoolean(_:_:)
func CGPDFScannerPopBoolean(scanner PDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopBoolean(scanner, value)
}

// Retrieves a PDF dictionary object from the scanner stack.
//
// Added in macOS 10.4.
// Retrieves a PDF dictionary object from the scanner stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopDictionary(_:_:)
func CGPDFScannerPopDictionary(scanner PDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopDictionary(scanner, value)
}

// Retrieves an integer object from the scanner stack.
//
// Added in macOS 10.4.
// Retrieves an integer object from the scanner stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopInteger(_:_:)
func CGPDFScannerPopInteger(scanner PDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopInteger(scanner, value)
}

// Retrieves a character string from the scanner stack.
//
// Added in macOS 10.4.
// Retrieves a character string from the scanner stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopName(_:_:)
func CGPDFScannerPopName(scanner PDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopName(scanner, value)
}

// Retrieves a real value object from the scanner stack.
//
// Added in macOS 10.4.
// Retrieves a real value object from the scanner stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopNumber(_:_:)
func CGPDFScannerPopNumber(scanner PDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopNumber(scanner, value)
}

// Retrieves an object from the scanner stack.
//
// Added in macOS 10.4.
// Retrieves an object from the scanner stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopObject(_:_:)
func CGPDFScannerPopObject(scanner PDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopObject(scanner, value)
}

// Retrieves a PDF stream object from the scanner stack.
//
// Added in macOS 10.4.
// Retrieves a PDF stream object from the scanner stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopStream(_:_:)
func CGPDFScannerPopStream(scanner PDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopStream(scanner, value)
}

// Retrieves a string object from the scanner stack.
//
// Added in macOS 10.4.
// Retrieves a string object from the scanner stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopString(_:_:)
func CGPDFScannerPopString(scanner PDFScannerRef, value unsafe.Pointer) bool {
	return _CGPDFScannerPopString(scanner, value)
}

// Decrements the retain count of a scanner object.
//
// Added in macOS 10.4.
// Decrements the retain count of a scanner object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerRelease(_:)
func CGPDFScannerRelease(scanner PDFScannerRef) {
	_CGPDFScannerRelease(scanner)
}

// Increments the retain count of a scanner object.
//
// Added in macOS 10.4.
// Increments the retain count of a scanner object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerRetain(_:)
func CGPDFScannerRetain(scanner PDFScannerRef) PDFScannerRef {
	return _CGPDFScannerRetain(scanner)
}

// Parses the content stream of a PDF scanner object.
//
// Added in macOS 10.4.
// Parses the content stream of a PDF scanner object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerScan(_:)
func CGPDFScannerScan(scanner PDFScannerRef) bool {
	return _CGPDFScannerScan(scanner)
}

// CGPDFScannerStop is a CoreGraphics function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerStop(_:)
func CGPDFScannerStop(s PDFScannerRef) {
	_CGPDFScannerStop(s)
}

// Returns the data associated with a PDF stream.
//
// Added in macOS 10.3.
// Returns the data associated with a PDF stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFStreamCopyData(_:_:)
func CGPDFStreamCopyData(stream PDFStreamRef, format unsafe.Pointer) DataRef {
	return _CGPDFStreamCopyData(stream, format)
}

// Returns the dictionary associated with a PDF stream.
//
// Added in macOS 10.3.
// Returns the dictionary associated with a PDF stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFStreamGetDictionary(_:)
func CGPDFStreamGetDictionary(stream PDFStreamRef) PDFDictionaryRef {
	return _CGPDFStreamGetDictionary(stream)
}

// Converts a string to a date.
//
// Added in macOS 10.4.
// Converts a string to a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringCopyDate(_:)
func CGPDFStringCopyDate(string_ PDFStringRef) DateRef {
	return _CGPDFStringCopyDate(string_)
}

// Returns a CFString object that represents a PDF string as a text string.
//
// Added in macOS 10.3.
// Returns a CFString object that represents a PDF string as a text string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringCopyTextString(_:)
func CGPDFStringCopyTextString(string_ PDFStringRef) StringRef {
	return _CGPDFStringCopyTextString(string_)
}

// Returns a pointer to the bytes of a PDF string.
//
// Added in macOS 10.3.
// Returns a pointer to the bytes of a PDF string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringGetBytePtr(_:)
func CGPDFStringGetBytePtr(string_ PDFStringRef) unsafe.Pointer {
	return _CGPDFStringGetBytePtr(string_)
}

// Returns the number of bytes in a PDF string.
//
// Added in macOS 10.3.
// Returns the number of bytes in a PDF string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringGetLength(_:)
func CGPDFStringGetLength(string_ PDFStringRef) uintptr {
	return _CGPDFStringGetLength(string_)
}

// CGPDFTagTypeGetName is a CoreGraphics function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/name
func CGPDFTagTypeGetName(tagType PDFTagType) unsafe.Pointer {
	return _CGPDFTagTypeGetName(tagType)
}

// Tells a PostScript converter to abort a conversion at the next available opportunity.
//
// Added in macOS 10.3.
// Tells a PostScript converter to abort a conversion at the next available opportunity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/abort()
func CGPSConverterAbort(converter PSConverterRef) bool {
	return _CGPSConverterAbort(converter)
}

// Uses a PostScript converter to convert PostScript data to PDF data.
//
// Added in macOS 10.3.
// Uses a PostScript converter to convert PostScript data to PDF data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/convert(_:consumer:options:)
func CGPSConverterConvert(converter PSConverterRef, provider DataProviderRef, consumer DataConsumerRef, options DictionaryRef) bool {
	return _CGPSConverterConvert(converter, provider, consumer, options)
}

// Creates a new PostScript converter.
//
// Added in macOS 10.3.
// Creates a new PostScript converter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/init(info:callbacks:options:)
func CGPSConverterCreate(info unsafe.Pointer, callbacks unsafe.Pointer, options DictionaryRef) PSConverterRef {
	return _CGPSConverterCreate(info, callbacks, options)
}

// Checks whether the converter is currently converting data.
//
// Added in macOS 10.3.
// Checks whether the converter is currently converting data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/isConverting
func CGPSConverterIsConverting(converter PSConverterRef) bool {
	return _CGPSConverterIsConverting(converter)
}

// Returns the Core Foundation type identifier for PostScript converters.
//
// Added in macOS 10.3.
// Returns the Core Foundation type identifier for PostScript converters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/typeID
func CGPSConverterGetTypeID() TypeID {
	return _CGPSConverterGetTypeID()
}

// For each element in a graphics path, calls a custom applier function.
//
// Added in macOS 10.2.
// For each element in a graphics path, calls a custom applier function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/apply(info:function:)
func CGPathApply(path PathRef, info unsafe.Pointer, function PathApplierFunction) {
	_CGPathApply(path, info, function)
}

// CGPathApplyWithBlock is a CoreGraphics function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/applyWithBlock(_:)
func CGPathApplyWithBlock(path PathRef, block PathApplyBlock) {
	_CGPathApplyWithBlock(path, block)
}

// Returns the bounding box containing all points in a graphics path.
//
// Added in macOS 10.2.
// Returns the bounding box containing all points in a graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/boundingBox
func CGPathGetBoundingBox(path PathRef) Rect {
	return _CGPathGetBoundingBox(path)
}

// Returns the bounding box of a graphics path.
//
// Added in macOS 10.6.
// Returns the bounding box of a graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/boundingBoxOfPath
func CGPathGetPathBoundingBox(path PathRef) Rect {
	return _CGPathGetPathBoundingBox(path)
}

// Creates an immutable copy of a graphics path.
//
// Added in macOS 10.2.
// Creates an immutable copy of a graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/copy()
func CGPathCreateCopy(path PathRef) PathRef {
	return _CGPathCreateCopy(path)
}

// Creates an immutable copy of a graphics path transformed by a transformation matrix.
//
// Added in macOS 10.7.
// Creates an immutable copy of a graphics path transformed by a transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/copy(using:)
func CGPathCreateCopyByTransformingPath(path PathRef, transform unsafe.Pointer) PathRef {
	return _CGPathCreateCopyByTransformingPath(path, transform)
}

// Returns the current point in a graphics path.
//
// Added in macOS 10.2.
// Returns the current point in a graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/currentPoint
func CGPathGetCurrentPoint(path PathRef) Point {
	return _CGPathGetCurrentPoint(path)
}

// Create an immutable path of an ellipse.
//
// Added in macOS 10.7.
// Create an immutable path of an ellipse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/init(ellipseIn:transform:)
func CGPathCreateWithEllipseInRect(rect Rect, transform unsafe.Pointer) PathRef {
	return _CGPathCreateWithEllipseInRect(rect, transform)
}

// Create an immutable path of a rectangle.
//
// Added in macOS 10.5.
// Create an immutable path of a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/init(rect:transform:)
func CGPathCreateWithRect(rect Rect, transform unsafe.Pointer) PathRef {
	return _CGPathCreateWithRect(rect, transform)
}

// Create an immutable path of a rounded rectangle.
//
// Added in macOS 10.9.
// Create an immutable path of a rounded rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/init(roundedRect:cornerWidth:cornerHeight:transform:)
func CGPathCreateWithRoundedRect(rect Rect, cornerWidth Float, cornerHeight Float, transform unsafe.Pointer) PathRef {
	return _CGPathCreateWithRoundedRect(rect, cornerWidth, cornerHeight, transform)
}

// Indicates whether or not a graphics path is empty.
//
// Added in macOS 10.2.
// Indicates whether or not a graphics path is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/isEmpty
func CGPathIsEmpty(path PathRef) bool {
	return _CGPathIsEmpty(path)
}

// Indicates whether or not a graphics path represents a rectangle.
//
// Added in macOS 10.2.
// Indicates whether or not a graphics path represents a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/isRect(_:)
func CGPathIsRect(path PathRef, rect unsafe.Pointer) bool {
	return _CGPathIsRect(path, rect)
}

// Creates a mutable copy of an existing graphics path.
//
// Added in macOS 10.2.
// Creates a mutable copy of an existing graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/mutableCopy()
func CGPathCreateMutableCopy(path PathRef) MutablePathRef {
	return _CGPathCreateMutableCopy(path)
}

// Creates a mutable copy of a graphics path transformed by a transformation matrix.
//
// Added in macOS 10.7.
// Creates a mutable copy of a graphics path transformed by a transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/mutableCopy(using:)
func CGPathCreateMutableCopyByTransformingPath(path PathRef, transform unsafe.Pointer) MutablePathRef {
	return _CGPathCreateMutableCopyByTransformingPath(path, transform)
}

// Returns the Core Foundation type identifier for Core Graphics paths.
//
// Added in macOS 10.2.
// Returns the Core Foundation type identifier for Core Graphics paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPath/typeID
func CGPathGetTypeID() TypeID {
	return _CGPathGetTypeID()
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment.
//
// Added in macOS 10.2.
// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddArc
func CGPathAddArc(path MutablePathRef, m unsafe.Pointer, x Float, y Float, radius Float, startAngle Float, endAngle Float, clockwise bool) {
	_CGPathAddArc(path, m, x, y, radius, startAngle, endAngle, clockwise)
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment.
//
// Added in macOS 10.2.
// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddArcToPoint
func CGPathAddArcToPoint(path MutablePathRef, m unsafe.Pointer, x1 Float, y1 Float, x2 Float, y2 Float, radius Float) {
	_CGPathAddArcToPoint(path, m, x1, y1, x2, y2, radius)
}

// Appends a cubic Bézier curve to a mutable graphics path.
//
// Added in macOS 10.2.
// Appends a cubic Bézier curve to a mutable graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddCurveToPoint
func CGPathAddCurveToPoint(path MutablePathRef, m unsafe.Pointer, cp1x Float, cp1y Float, cp2x Float, cp2y Float, x Float, y Float) {
	_CGPathAddCurveToPoint(path, m, cp1x, cp1y, cp2x, cp2y, x, y)
}

// Adds to a path an ellipse that fits inside a rectangle.
//
// Added in macOS 10.4.
// Adds to a path an ellipse that fits inside a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddEllipseInRect
func CGPathAddEllipseInRect(path MutablePathRef, m unsafe.Pointer, rect Rect) {
	_CGPathAddEllipseInRect(path, m, rect)
}

// Appends a line segment to a mutable graphics path.
//
// Added in macOS 10.2.
// Appends a line segment to a mutable graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddLineToPoint
func CGPathAddLineToPoint(path MutablePathRef, m unsafe.Pointer, x Float, y Float) {
	_CGPathAddLineToPoint(path, m, x, y)
}

// Appends an array of new line segments to a mutable graphics path.
//
// Added in macOS 10.2.
// Appends an array of new line segments to a mutable graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddLines
func CGPathAddLines(path MutablePathRef, m unsafe.Pointer, points unsafe.Pointer, count uintptr) {
	_CGPathAddLines(path, m, points, count)
}

// Appends a path to onto a mutable graphics path.
//
// Added in macOS 10.2.
// Appends a path to onto a mutable graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddPath
func CGPathAddPath(path1 MutablePathRef, m unsafe.Pointer, path2 PathRef) {
	_CGPathAddPath(path1, m, path2)
}

// Appends a quadratic Bézier curve to a mutable graphics path.
//
// Added in macOS 10.2.
// Appends a quadratic Bézier curve to a mutable graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddQuadCurveToPoint
func CGPathAddQuadCurveToPoint(path MutablePathRef, m unsafe.Pointer, cpx Float, cpy Float, x Float, y Float) {
	_CGPathAddQuadCurveToPoint(path, m, cpx, cpy, x, y)
}

// Appends a rectangle to a mutable graphics path.
//
// Added in macOS 10.2.
// Appends a rectangle to a mutable graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddRect
func CGPathAddRect(path MutablePathRef, m unsafe.Pointer, rect Rect) {
	_CGPathAddRect(path, m, rect)
}

// Appends an array of rectangles to a mutable graphics path.
//
// Added in macOS 10.2.
// Appends an array of rectangles to a mutable graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddRects
func CGPathAddRects(path MutablePathRef, m unsafe.Pointer, rects unsafe.Pointer, count uintptr) {
	_CGPathAddRects(path, m, rects, count)
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment.
//
// Added in macOS 10.7.
// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddRelativeArc
func CGPathAddRelativeArc(path MutablePathRef, matrix unsafe.Pointer, x Float, y Float, radius Float, startAngle Float, delta Float) {
	_CGPathAddRelativeArc(path, matrix, x, y, radius, startAngle, delta)
}

// Appends a rounded rectangle to a mutable graphics path.
//
// Added in macOS 10.9.
// Appends a rounded rectangle to a mutable graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathAddRoundedRect
func CGPathAddRoundedRect(path MutablePathRef, transform unsafe.Pointer, rect Rect, cornerWidth Float, cornerHeight Float) {
	_CGPathAddRoundedRect(path, transform, rect, cornerWidth, cornerHeight)
}

// Checks whether a point is contained in a graphics path.
//
// Added in macOS 10.4.
// Checks whether a point is contained in a graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathContainsPoint
func CGPathContainsPoint(path PathRef, m unsafe.Pointer, point Point, eoFill bool) bool {
	return _CGPathContainsPoint(path, m, point, eoFill)
}

// Creates a dashed copy of another path.
//
// Added in macOS 10.7.
// Creates a dashed copy of another path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByDashingPath
func CGPathCreateCopyByDashingPath(path PathRef, transform unsafe.Pointer, phase Float, lengths []float64, count uintptr) PathRef {
	return _CGPathCreateCopyByDashingPath(path, transform, phase, lengths, count)
}

// CGPathCreateCopyByFlattening is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByFlattening
func CGPathCreateCopyByFlattening(path PathRef, flatteningThreshold Float) PathRef {
	return _CGPathCreateCopyByFlattening(path, flatteningThreshold)
}

// CGPathCreateCopyByIntersectingPath is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByIntersectingPath
func CGPathCreateCopyByIntersectingPath(path PathRef, maskPath PathRef, evenOddFillRule bool) PathRef {
	return _CGPathCreateCopyByIntersectingPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateCopyByNormalizing is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByNormalizing
func CGPathCreateCopyByNormalizing(path PathRef, evenOddFillRule bool) PathRef {
	return _CGPathCreateCopyByNormalizing(path, evenOddFillRule)
}

// Creates a stroked copy of another path.
//
// Added in macOS 10.7.
// Creates a stroked copy of another path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByStrokingPath
func CGPathCreateCopyByStrokingPath(path PathRef, transform unsafe.Pointer, lineWidth Float, lineCap LineCap, lineJoin LineJoin, miterLimit Float) PathRef {
	return _CGPathCreateCopyByStrokingPath(path, transform, lineWidth, lineCap, lineJoin, miterLimit)
}

// CGPathCreateCopyBySubtractingPath is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyBySubtractingPath
func CGPathCreateCopyBySubtractingPath(path PathRef, maskPath PathRef, evenOddFillRule bool) PathRef {
	return _CGPathCreateCopyBySubtractingPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateCopyBySymmetricDifferenceOfPath is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyBySymmetricDifferenceOfPath
func CGPathCreateCopyBySymmetricDifferenceOfPath(path PathRef, maskPath PathRef, evenOddFillRule bool) PathRef {
	return _CGPathCreateCopyBySymmetricDifferenceOfPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateCopyByUnioningPath is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByUnioningPath
func CGPathCreateCopyByUnioningPath(path PathRef, maskPath PathRef, evenOddFillRule bool) PathRef {
	return _CGPathCreateCopyByUnioningPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateCopyOfLineByIntersectingPath is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyOfLineByIntersectingPath
func CGPathCreateCopyOfLineByIntersectingPath(path PathRef, maskPath PathRef, evenOddFillRule bool) PathRef {
	return _CGPathCreateCopyOfLineByIntersectingPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateCopyOfLineBySubtractingPath is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyOfLineBySubtractingPath
func CGPathCreateCopyOfLineBySubtractingPath(path PathRef, maskPath PathRef, evenOddFillRule bool) PathRef {
	return _CGPathCreateCopyOfLineBySubtractingPath(path, maskPath, evenOddFillRule)
}

// CGPathCreateSeparateComponents is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateSeparateComponents
func CGPathCreateSeparateComponents(path PathRef, evenOddFillRule bool) ArrayRef {
	return _CGPathCreateSeparateComponents(path, evenOddFillRule)
}

// Indicates whether two graphics paths are equivalent.
//
// Added in macOS 10.2.
// Indicates whether two graphics paths are equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathEqualToPath
func CGPathEqualToPath(path1 PathRef, path2 PathRef) bool {
	return _CGPathEqualToPath(path1, path2)
}

// CGPathIntersectsPath is a CoreGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathIntersectsPath
func CGPathIntersectsPath(path1 PathRef, path2 PathRef, evenOddFillRule bool) bool {
	return _CGPathIntersectsPath(path1, path2, evenOddFillRule)
}

// Starts a new subpath at a specified location in a mutable graphics path.
//
// Added in macOS 10.2.
// Starts a new subpath at a specified location in a mutable graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathMoveToPoint
func CGPathMoveToPoint(path MutablePathRef, m unsafe.Pointer, x Float, y Float) {
	_CGPathMoveToPoint(path, m, x, y)
}

// Decrements the retain count of a graphics path.
//
// Added in macOS 10.2.
// Decrements the retain count of a graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathRelease
func CGPathRelease(path PathRef) {
	_CGPathRelease(path)
}

// Increments the retain count of a graphics path.
//
// Added in macOS 10.2.
// Increments the retain count of a graphics path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathRetain
func CGPathRetain(path PathRef) PathRef {
	return _CGPathRetain(path)
}

// Creates a pattern object.
//
// Added in macOS 10.0.
// Creates a pattern object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPattern/init(info:bounds:matrix:xStep:yStep:tiling:isColored:callbacks:)
func CGPatternCreate(info unsafe.Pointer, bounds Rect, matrix AffineTransform, xStep Float, yStep Float, tiling PatternTiling, isColored bool, callbacks unsafe.Pointer) PatternRef {
	return _CGPatternCreate(info, bounds, matrix, xStep, yStep, tiling, isColored, callbacks)
}

// Returns the type identifier for Core Graphics patterns.
//
// Added in macOS 10.2.
// Returns the type identifier for Core Graphics patterns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPattern/typeID
func CGPatternGetTypeID() TypeID {
	return _CGPatternGetTypeID()
}

// Decrements the retain count of a Core Graphics pattern.
//
// Added in macOS 10.0.
// Decrements the retain count of a Core Graphics pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternRelease
func CGPatternRelease(pattern PatternRef) {
	_CGPatternRelease(pattern)
}

// Increments the retain count of a Core Graphics pattern.
//
// Added in macOS 10.0.
// Increments the retain count of a Core Graphics pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternRetain
func CGPatternRetain(pattern PatternRef) PatternRef {
	return _CGPatternRetain(pattern)
}

// Returns the point resulting from an affine transformation of an existing point.
//
// Added in macOS 10.0.
// Returns the point resulting from an affine transformation of an existing point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPointApplyAffineTransform(_:_:)
func CGPointApplyAffineTransform(point Point, t AffineTransform) Point {
	return _CGPointApplyAffineTransform(point, t)
}

// Returns a dictionary representation of the specified point.
//
// Added in macOS 10.5.
// Returns a dictionary representation of the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPointCreateDictionaryRepresentation(_:)
func CGPointCreateDictionaryRepresentation(point Point) DictionaryRef {
	return _CGPointCreateDictionaryRepresentation(point)
}

// Returns whether two points are equal.
//
// Added in macOS 10.0.
// Returns whether two points are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPointEqualToPoint(_:_:)
func CGPointEqualToPoint(point1 Point, point2 Point) bool {
	return _CGPointEqualToPoint(point1, point2)
}

// Fills in a point using the contents of the specified dictionary.
//
// Added in macOS 10.5.
// Fills in a point using the contents of the specified dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPointMakeWithDictionaryRepresentation(_:_:)
func CGPointMakeWithDictionaryRepresentation(dict DictionaryRef, point unsafe.Pointer) bool {
	return _CGPointMakeWithDictionaryRepresentation(dict, point)
}

// Synthesizes a low-level keyboard event on the local machine.

// Synthesizes a low-level keyboard event on the local machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPostKeyboardEvent(_:_:_:)
func CGPostKeyboardEvent(keyChar CharCode, virtualKey KeyCode, keyDown unsafe.Pointer) Error {
	return _CGPostKeyboardEvent(keyChar, virtualKey, keyDown)
}

// Synthesizes a low-level mouse-button event on the local machine.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Synthesizes a low-level mouse-button event on the local machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPostMouseEvent
func CGPostMouseEvent(mouseCursorPosition Point, updateMouseCursorPosition unsafe.Pointer, buttonCount ButtonCount, mouseButtonDown unsafe.Pointer) Error {
	return _CGPostMouseEvent(mouseCursorPosition, updateMouseCursorPosition, buttonCount, mouseButtonDown)
}

// Synthesizes a low-level scrolling event on the local machine.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Synthesizes a low-level scrolling event on the local machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPostScrollWheelEvent
func CGPostScrollWheelEvent(wheelCount WheelCount, wheel1 int32) Error {
	return _CGPostScrollWheelEvent(wheelCount, wheel1)
}

// CGPreflightListenEventAccess is a CoreGraphics function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPreflightListenEventAccess()
func CGPreflightListenEventAccess() bool {
	return _CGPreflightListenEventAccess()
}

// CGPreflightPostEventAccess is a CoreGraphics function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPreflightPostEventAccess()
func CGPreflightPostEventAccess() bool {
	return _CGPreflightPostEventAccess()
}

// CGPreflightScreenCaptureAccess is a CoreGraphics function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPreflightScreenCaptureAccess()
func CGPreflightScreenCaptureAccess() bool {
	return _CGPreflightScreenCaptureAccess()
}

// Applies an affine transform to a rectangle.
//
// Added in macOS 10.4.
// Applies an affine transform to a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectApplyAffineTransform(_:_:)
func CGRectApplyAffineTransform(rect Rect, t AffineTransform) Rect {
	return _CGRectApplyAffineTransform(rect, t)
}

// Returns whether a rectangle contains a specified point.
//
// Added in macOS 10.0.
// Returns whether a rectangle contains a specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectContainsPoint(_:_:)
func CGRectContainsPoint(rect Rect, point Point) bool {
	return _CGRectContainsPoint(rect, point)
}

// Returns whether the first rectangle contains the second rectangle.
//
// Added in macOS 10.0.
// Returns whether the first rectangle contains the second rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectContainsRect(_:_:)
func CGRectContainsRect(rect1 Rect, rect2 Rect) bool {
	return _CGRectContainsRect(rect1, rect2)
}

// Returns a dictionary representation of the provided rectangle.
//
// Added in macOS 10.5.
// Returns a dictionary representation of the provided rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectCreateDictionaryRepresentation(_:)
func CGRectCreateDictionaryRepresentation(p0 Rect) DictionaryRef {
	return _CGRectCreateDictionaryRepresentation(p0)
}

// Divides a source rectangle into two component rectangles.
//
// Added in macOS 10.0.
// Divides a source rectangle into two component rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectDivide
func CGRectDivide(rect Rect, slice unsafe.Pointer, remainder unsafe.Pointer, amount Float, edge RectEdge) {
	_CGRectDivide(rect, slice, remainder, amount, edge)
}

// Returns whether two rectangles are equal in size and position.
//
// Added in macOS 10.0.
// Returns whether two rectangles are equal in size and position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectEqualToRect(_:_:)
func CGRectEqualToRect(rect1 Rect, rect2 Rect) bool {
	return _CGRectEqualToRect(rect1, rect2)
}

// Returns the height of a rectangle.
//
// Added in macOS 10.0.
// Returns the height of a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectGetHeight(_:)
func CGRectGetHeight(rect Rect) Float {
	return _CGRectGetHeight(rect)
}

// Returns the largest value of the x-coordinate for the rectangle.
//
// Added in macOS 10.0.
// Returns the largest value of the x-coordinate for the rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMaxX(_:)
func CGRectGetMaxX(rect Rect) Float {
	return _CGRectGetMaxX(rect)
}

// Returns the largest value for the y-coordinate of the rectangle.
//
// Added in macOS 10.0.
// Returns the largest value for the y-coordinate of the rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMaxY(_:)
func CGRectGetMaxY(rect Rect) Float {
	return _CGRectGetMaxY(rect)
}

// Returns the x- coordinate that establishes the center of a rectangle.
//
// Added in macOS 10.0.
// Returns the x- coordinate that establishes the center of a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMidX(_:)
func CGRectGetMidX(rect Rect) Float {
	return _CGRectGetMidX(rect)
}

// Returns the y-coordinate that establishes the center of the rectangle.
//
// Added in macOS 10.0.
// Returns the y-coordinate that establishes the center of the rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMidY(_:)
func CGRectGetMidY(rect Rect) Float {
	return _CGRectGetMidY(rect)
}

// Returns the smallest value for the x-coordinate of the rectangle.
//
// Added in macOS 10.0.
// Returns the smallest value for the x-coordinate of the rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMinX(_:)
func CGRectGetMinX(rect Rect) Float {
	return _CGRectGetMinX(rect)
}

// Returns the smallest value for the y-coordinate of the rectangle.
//
// Added in macOS 10.0.
// Returns the smallest value for the y-coordinate of the rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMinY(_:)
func CGRectGetMinY(rect Rect) Float {
	return _CGRectGetMinY(rect)
}

// Returns the width of a rectangle.
//
// Added in macOS 10.0.
// Returns the width of a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectGetWidth(_:)
func CGRectGetWidth(rect Rect) Float {
	return _CGRectGetWidth(rect)
}

// Returns a rectangle that is smaller or larger than the source rectangle, with the same center point.
//
// Added in macOS 10.0.
// Returns a rectangle that is smaller or larger than the source rectangle, with the same center point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectInset(_:_:_:)
func CGRectInset(rect Rect, dx Float, dy Float) Rect {
	return _CGRectInset(rect, dx, dy)
}

// Returns the smallest rectangle that results from converting the source rectangle values to integers.
//
// Added in macOS 10.0.
// Returns the smallest rectangle that results from converting the source rectangle values to integers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectIntegral(_:)
func CGRectIntegral(rect Rect) Rect {
	return _CGRectIntegral(rect)
}

// Returns the intersection of two rectangles.
//
// Added in macOS 10.0.
// Returns the intersection of two rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectIntersection(_:_:)
func CGRectIntersection(r1 Rect, r2 Rect) Rect {
	return _CGRectIntersection(r1, r2)
}

// Returns whether two rectangles intersect.
//
// Added in macOS 10.0.
// Returns whether two rectangles intersect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectIntersectsRect(_:_:)
func CGRectIntersectsRect(rect1 Rect, rect2 Rect) bool {
	return _CGRectIntersectsRect(rect1, rect2)
}

// Returns whether a rectangle has zero width or height, or is a null rectangle.
//
// Added in macOS 10.0.
// Returns whether a rectangle has zero width or height, or is a null rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectIsEmpty(_:)
func CGRectIsEmpty(rect Rect) bool {
	return _CGRectIsEmpty(rect)
}

// Returns whether a rectangle is infinite.
//
// Added in macOS 10.4.
// Returns whether a rectangle is infinite.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectIsInfinite(_:)
func CGRectIsInfinite(rect Rect) bool {
	return _CGRectIsInfinite(rect)
}

// Returns whether the rectangle is equal to the null rectangle.
//
// Added in macOS 10.0.
// Returns whether the rectangle is equal to the null rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectIsNull(_:)
func CGRectIsNull(rect Rect) bool {
	return _CGRectIsNull(rect)
}

// Fills in a rectangle using the contents of the specified dictionary.
//
// Added in macOS 10.5.
// Fills in a rectangle using the contents of the specified dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectMakeWithDictionaryRepresentation(_:_:)
func CGRectMakeWithDictionaryRepresentation(dict DictionaryRef, rect unsafe.Pointer) bool {
	return _CGRectMakeWithDictionaryRepresentation(dict, rect)
}

// Returns a rectangle with an origin that is offset from that of the source rectangle.
//
// Added in macOS 10.0.
// Returns a rectangle with an origin that is offset from that of the source rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectOffset(_:_:_:)
func CGRectOffset(rect Rect, dx Float, dy Float) Rect {
	return _CGRectOffset(rect, dx, dy)
}

// Returns a rectangle with a positive width and height.
//
// Added in macOS 10.0.
// Returns a rectangle with a positive width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectStandardize(_:)
func CGRectStandardize(rect Rect) Rect {
	return _CGRectStandardize(rect)
}

// Returns the smallest rectangle that contains the two source rectangles.
//
// Added in macOS 10.0.
// Returns the smallest rectangle that contains the two source rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRectUnion(_:_:)
func CGRectUnion(r1 Rect, r2 Rect) Rect {
	return _CGRectUnion(r1, r2)
}

// Registers a callback function to be invoked when local displays are refreshed or modified.

// Registers a callback function to be invoked when local displays are refreshed or modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRegisterScreenRefreshCallback(_:_:)
func CGRegisterScreenRefreshCallback(callback ScreenRefreshCallback, userInfo unsafe.Pointer) Error {
	return _CGRegisterScreenRefreshCallback(callback, userInfo)
}

// Releases all captured displays.
//
// Added in macOS 10.0.
// Releases all captured displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGReleaseAllDisplays()
func CGReleaseAllDisplays() Error {
	return _CGReleaseAllDisplays()
}

// Releases a display fade reservation, and unfades the display if needed.
//
// Added in macOS 10.2.
// Releases a display fade reservation, and unfades the display if needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGReleaseDisplayFadeReservation(_:)
func CGReleaseDisplayFadeReservation(token DisplayFadeReservationToken) Error {
	return _CGReleaseDisplayFadeReservation(token)
}

// Deallocates a list of rectangles that represent changed areas on local displays.

// Deallocates a list of rectangles that represent changed areas on local displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGReleaseScreenRefreshRects(_:)
func CGReleaseScreenRefreshRects(rects unsafe.Pointer) {
	_CGReleaseScreenRefreshRects(rects)
}

// CGRenderingBufferLockBytePtr is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferLockBytePtr
func CGRenderingBufferLockBytePtr(provider RenderingBufferProviderRef) unsafe.Pointer {
	return _CGRenderingBufferLockBytePtr(provider)
}

// CGRenderingBufferProviderCreate is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderCreate
func CGRenderingBufferProviderCreate(info unsafe.Pointer, size uintptr) RenderingBufferProviderRef {
	return _CGRenderingBufferProviderCreate(info, size)
}

// CGRenderingBufferProviderCreateWithCFData is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderCreateWithCFData
func CGRenderingBufferProviderCreateWithCFData(data MutableDataRef) RenderingBufferProviderRef {
	return _CGRenderingBufferProviderCreateWithCFData(data)
}

// CGRenderingBufferProviderGetSize is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderGetSize
func CGRenderingBufferProviderGetSize(provider RenderingBufferProviderRef) uintptr {
	return _CGRenderingBufferProviderGetSize(provider)
}

// CGRenderingBufferProviderGetTypeID is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderGetTypeID
func CGRenderingBufferProviderGetTypeID() TypeID {
	return _CGRenderingBufferProviderGetTypeID()
}

// CGRenderingBufferUnlockBytePtr is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferUnlockBytePtr
func CGRenderingBufferUnlockBytePtr(provider RenderingBufferProviderRef) {
	_CGRenderingBufferUnlockBytePtr(provider)
}

// CGRequestListenEventAccess is a CoreGraphics function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRequestListenEventAccess()
func CGRequestListenEventAccess() bool {
	return _CGRequestListenEventAccess()
}

// CGRequestPostEventAccess is a CoreGraphics function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRequestPostEventAccess()
func CGRequestPostEventAccess() bool {
	return _CGRequestPostEventAccess()
}

// CGRequestScreenCaptureAccess is a CoreGraphics function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRequestScreenCaptureAccess()
func CGRequestScreenCaptureAccess() bool {
	return _CGRequestScreenCaptureAccess()
}

// Restores the permanent display configuration settings for the current user.
//
// Added in macOS 10.2.
// Restores the permanent display configuration settings for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGRestorePermanentDisplayConfiguration()
func CGRestorePermanentDisplayConfiguration() {
	_CGRestorePermanentDisplayConfiguration()
}

// Registers a callback function to be invoked when an area of the display is moved.

// Registers a callback function to be invoked when an area of the display is moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenRegisterMoveCallback(_:_:)
func CGScreenRegisterMoveCallback(callback ScreenUpdateMoveCallback, userInfo unsafe.Pointer) Error {
	return _CGScreenRegisterMoveCallback(callback, userInfo)
}

// Removes a previously registered callback function invoked when an area of the display is moved.

// Removes a previously registered callback function invoked when an area of the display is moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenUnregisterMoveCallback(_:_:)
func CGScreenUnregisterMoveCallback(callback ScreenUpdateMoveCallback, userInfo unsafe.Pointer) {
	_CGScreenUnregisterMoveCallback(callback, userInfo)
}

// Returns information about the caller’s window server session.
//
// Added in macOS 10.3.
// Returns information about the caller’s window server session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGSessionCopyCurrentDictionary()
func CGSessionCopyCurrentDictionary() DictionaryRef {
	return _CGSessionCopyCurrentDictionary()
}

// Sets the byte values in the 8-bit RGB gamma tables for a display.
//
// Added in macOS 10.0.
// Sets the byte values in the 8-bit RGB gamma tables for a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGSetDisplayTransferByByteTable(_:_:_:_:_:)
func CGSetDisplayTransferByByteTable(display DirectDisplayID, tableSize uint32, redTable unsafe.Pointer, greenTable unsafe.Pointer, blueTable unsafe.Pointer) Error {
	return _CGSetDisplayTransferByByteTable(display, tableSize, redTable, greenTable, blueTable)
}

// Sets the gamma function for a display by specifying the coefficients of the gamma transfer formula.
//
// Added in macOS 10.0.
// Sets the gamma function for a display by specifying the coefficients of the gamma transfer formula.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGSetDisplayTransferByFormula(_:_:_:_:_:_:_:_:_:_:)
func CGSetDisplayTransferByFormula(display DirectDisplayID, redMin GammaValue, redMax GammaValue, redGamma GammaValue, greenMin GammaValue, greenMax GammaValue, greenGamma GammaValue, blueMin GammaValue, blueMax GammaValue, blueGamma GammaValue) Error {
	return _CGSetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
}

// Sets the color gamma function for a display by specifying the values in the RGB gamma tables.
//
// Added in macOS 10.0.
// Sets the color gamma function for a display by specifying the values in the RGB gamma tables.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGSetDisplayTransferByTable(_:_:_:_:_:)
func CGSetDisplayTransferByTable(display DirectDisplayID, tableSize uint32, redTable unsafe.Pointer, greenTable unsafe.Pointer, blueTable unsafe.Pointer) Error {
	return _CGSetDisplayTransferByTable(display, tableSize, redTable, greenTable, blueTable)
}

// Filters local hardware events from the keyboard and mouse during the short interval after a synthetic event is posted.

// Filters local hardware events from the keyboard and mouse during the short interval after a synthetic event is posted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGSetLocalEventsFilterDuringSuppressionState(_:_:)
func CGSetLocalEventsFilterDuringSuppressionState(filter EventFilterMask, state EventSuppressionState) Error {
	return _CGSetLocalEventsFilterDuringSuppressionState(filter, state)
}

// Sets the time interval in seconds that local hardware events are suppressed after posting a synthetic event.

// Sets the time interval in seconds that local hardware events are suppressed after posting a synthetic event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGSetLocalEventsSuppressionInterval(_:)
func CGSetLocalEventsSuppressionInterval(seconds TimeInterval) Error {
	return _CGSetLocalEventsSuppressionInterval(seconds)
}

// CGShadingGetContentHeadroom is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGShading/contentHeadroom
func CGShadingGetContentHeadroom(shading ShadingRef) float32 {
	return _CGShadingGetContentHeadroom(shading)
}

// CGShadingCreateAxialWithContentHeadroom is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(axialHeadroom:space:start:end:function:extendStart:extendEnd:)
func CGShadingCreateAxialWithContentHeadroom(headroom float32, space ColorSpaceRef, start Point, end Point, function FunctionRef, extendStart bool, extendEnd bool) ShadingRef {
	return _CGShadingCreateAxialWithContentHeadroom(headroom, space, start, end, function, extendStart, extendEnd)
}

// Creates a shading object to use for axial shading.
//
// Added in macOS 10.2.
// Creates a shading object to use for axial shading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(axialSpace:start:end:function:extendStart:extendEnd:)
func CGShadingCreateAxial(space ColorSpaceRef, start Point, end Point, function FunctionRef, extendStart bool, extendEnd bool) ShadingRef {
	return _CGShadingCreateAxial(space, start, end, function, extendStart, extendEnd)
}

// CGShadingCreateRadialWithContentHeadroom is a CoreGraphics function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(radialHeadroom:space:start:startRadius:end:endRadius:function:extendStart:extendEnd:)
func CGShadingCreateRadialWithContentHeadroom(headroom float32, space ColorSpaceRef, start Point, startRadius Float, end Point, endRadius Float, function FunctionRef, extendStart bool, extendEnd bool) ShadingRef {
	return _CGShadingCreateRadialWithContentHeadroom(headroom, space, start, startRadius, end, endRadius, function, extendStart, extendEnd)
}

// Creates a shading object to use for radial shading.
//
// Added in macOS 10.2.
// Creates a shading object to use for radial shading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(radialSpace:start:startRadius:end:endRadius:function:extendStart:extendEnd:)
func CGShadingCreateRadial(space ColorSpaceRef, start Point, startRadius Float, end Point, endRadius Float, function FunctionRef, extendStart bool, extendEnd bool) ShadingRef {
	return _CGShadingCreateRadial(space, start, startRadius, end, endRadius, function, extendStart, extendEnd)
}

// Returns the Core Foundation type identifier for Core Graphics shading objects.
//
// Added in macOS 10.2.
// Returns the Core Foundation type identifier for Core Graphics shading objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGShading/typeID
func CGShadingGetTypeID() TypeID {
	return _CGShadingGetTypeID()
}

// Decrements the retain count of a shading object.
//
// Added in macOS 10.2.
// Decrements the retain count of a shading object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGShadingRelease
func CGShadingRelease(shading ShadingRef) {
	_CGShadingRelease(shading)
}

// Increments the retain count of a shading object.
//
// Added in macOS 10.2.
// Increments the retain count of a shading object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGShadingRetain
func CGShadingRetain(shading ShadingRef) ShadingRef {
	return _CGShadingRetain(shading)
}

// Returns the window ID of the shield window for a captured display.
//
// Added in macOS 10.0.
// Returns the window ID of the shield window for a captured display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGShieldingWindowID(_:)
func CGShieldingWindowID(display DirectDisplayID) WindowID {
	return _CGShieldingWindowID(display)
}

// Returns the window level of the shield window for a captured display.
//
// Added in macOS 10.0.
// Returns the window level of the shield window for a captured display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGShieldingWindowLevel()
func CGShieldingWindowLevel() WindowLevel {
	return _CGShieldingWindowLevel()
}

// Returns the height and width resulting from a transformation of an existing height and width.
//
// Added in macOS 10.0.
// Returns the height and width resulting from a transformation of an existing height and width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGSizeApplyAffineTransform(_:_:)
func CGSizeApplyAffineTransform(size Size, t AffineTransform) Size {
	return _CGSizeApplyAffineTransform(size, t)
}

// Returns a dictionary representation of the specified size.
//
// Added in macOS 10.5.
// Returns a dictionary representation of the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGSizeCreateDictionaryRepresentation(_:)
func CGSizeCreateDictionaryRepresentation(size Size) DictionaryRef {
	return _CGSizeCreateDictionaryRepresentation(size)
}

// Returns whether two sizes are equal.
//
// Added in macOS 10.0.
// Returns whether two sizes are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGSizeEqualToSize(_:_:)
func CGSizeEqualToSize(size1 Size, size2 Size) bool {
	return _CGSizeEqualToSize(size1, size2)
}

// Fills in a size using the contents of the specified dictionary.
//
// Added in macOS 10.5.
// Fills in a size using the contents of the specified dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGSizeMakeWithDictionaryRepresentation(_:_:)
func CGSizeMakeWithDictionaryRepresentation(dict DictionaryRef, size unsafe.Pointer) bool {
	return _CGSizeMakeWithDictionaryRepresentation(dict, size)
}

// Removes a previously registered callback function invoked when local displays are refreshed or modified.

// Removes a previously registered callback function invoked when local displays are refreshed or modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGUnregisterScreenRefreshCallback(_:_:)
func CGUnregisterScreenRefreshCallback(callback ScreenRefreshCallback, userInfo unsafe.Pointer) {
	_CGUnregisterScreenRefreshCallback(callback, userInfo)
}

// Waits for screen refresh operations.

// Waits for screen refresh operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWaitForScreenRefreshRects(_:_:)
func CGWaitForScreenRefreshRects(rects unsafe.Pointer, count []uint32) Error {
	return _CGWaitForScreenRefreshRects(rects, count)
}

// Waits for screen update operations.

// Waits for screen update operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWaitForScreenUpdateRects(_:_:_:_:_:)
func CGWaitForScreenUpdateRects(requestedOperations ScreenUpdateOperation, currentOperation unsafe.Pointer, rects unsafe.Pointer, rectCount unsafe.Pointer, delta unsafe.Pointer) Error {
	return _CGWaitForScreenUpdateRects(requestedOperations, currentOperation, rects, rectCount, delta)
}

// Moves the mouse cursor without generating events.
//
// Added in macOS 10.0.
// Moves the mouse cursor without generating events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWarpMouseCursorPosition(_:)
func CGWarpMouseCursorPosition(newCursorPosition Point) Error {
	return _CGWarpMouseCursorPosition(newCursorPosition)
}

// Returns the window level that corresponds to one of the standard window types.
//
// Added in macOS 10.0.
// Returns the window level that corresponds to one of the standard window types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelForKey(_:)
func CGWindowLevelForKey(key WindowLevelKey) WindowLevel {
	return _CGWindowLevelForKey(key)
}

// Generates and returns information about the selected windows in the current user session.
//
// Added in macOS 10.5.
// Generates and returns information about the selected windows in the current user session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCopyWindowInfo(_:_:)
func CGWindowListCopyWindowInfo(option WindowListOption, relativeToWindow WindowID) ArrayRef {
	return _CGWindowListCopyWindowInfo(option, relativeToWindow)
}

// Returns the list of window IDs associated with the specified windows in the current user session.
//
// Added in macOS 10.5.
// Returns the list of window IDs associated with the specified windows in the current user session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCreate
func CGWindowListCreate(option WindowListOption, relativeToWindow WindowID) ArrayRef {
	return _CGWindowListCreate(option, relativeToWindow)
}

// Generates and returns information about windows with the specified window IDs.
//
// Added in macOS 10.5.
// Generates and returns information about windows with the specified window IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCreateDescriptionFromArray(_:)
func CGWindowListCreateDescriptionFromArray(windowArray ArrayRef) ArrayRef {
	return _CGWindowListCreateDescriptionFromArray(windowArray)
}

// Returns a composite image based on a dynamically generated list of windows.

// Returns a composite image based on a dynamically generated list of windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCreateImage(_:_:_:_:)
func CGWindowListCreateImage(screenBounds Rect, listOption WindowListOption, windowID WindowID, imageOption WindowImageOption) ImageRef {
	return _CGWindowListCreateImage(screenBounds, listOption, windowID, imageOption)
}

// Returns a Core Foundation Mach port (CFMachPort) that corresponds to the macOS window server.

// Returns a Core Foundation Mach port (CFMachPort) that corresponds to the macOS window server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowServerCFMachPort()
func CGWindowServerCFMachPort() MachPortRef {
	return _CGWindowServerCFMachPort()
}

// CGWindowServerCreateServerPort is a CoreGraphics function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowServerCreateServerPort()
func CGWindowServerCreateServerPort() MachPortRef {
	return _CGWindowServerCreateServerPort()
}



