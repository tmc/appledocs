// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

// CoreGraphics Functions
//
// This file contains function declarations discovered from Apple's documentation.
// To use these functions, you need to:
//   1. Map C types to Go types
//   2. Create function variables
//   3. Register them with purego.RegisterLibFunc
//
// Example:
//   var CGContextSetRGBFillColor func(c CGContextRef, red, green, blue, alpha CGFloat)
//   purego.RegisterLibFunc(&CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")

// Discovered functions (210 total):

// CGContextSetInterpolationQuality(CGContextRef  c,  CGInterpolationQuality  quality)
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGAcquireDisplayFadeReservation(seconds CGDisplayReservationInterval, token ,  CGDisplayFadeReservationToken  *, ) CGError
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.2+

// CGAffineTransformConcat(t1 CGAffineTransform, t2 ,  CGAffineTransform, ) CGAffineTransform
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGAffineTransformMake(a CGFloat, b ,  CGFloat, c ,  CGFloat, d ,  CGFloat, tx ,  CGFloat, ty ,  CGFloat, ) CGAffineTransform
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGAffineTransformMakeTranslation(tx CGFloat, ty ,  CGFloat, ) CGAffineTransform
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGAffineTransformScale(t CGAffineTransform, sx ,  CGFloat, sy ,  CGFloat, ) CGAffineTransform
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGAffineTransformTranslate(t CGAffineTransform, tx ,  CGFloat, ty ,  CGFloat, ) CGAffineTransform
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGBitmapContextCreateAdaptive(width size_t, height ,  size_t, auxiliaryInfo ,  CFDictionaryRef, onResolve ,  bool  (^, )( const CGContentInfo  *, ,  CGBitmapParameters  *, onAllocate ),  CGRenderingBufferProviderRef  (^, )( const CGContentInfo  *, ,  const CGBitmapParameters  *, onFree ),  void  (^, )( CGRenderingBufferProviderRef, ,  const CGContentInfo  *, ,  const CGBitmapParameters  *, onError ),  void  (^, )( CFErrorRef, ,  const CGContentInfo  *, ,  const CGBitmapParameters  *, ) CGContextRef
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGColorGetColorSpace(color CGColorRef, ) CGColorSpaceRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGColorGetContentHeadroom(color CGColorRef, ) float
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGColorCreate(space CGColorSpaceRef, components ,  const CGFloat  *, ) CGColorRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGColorCreateGenericCMYK(cyan CGFloat, magenta ,  CGFloat, yellow ,  CGFloat, black ,  CGFloat, alpha ,  CGFloat, ) CGColorRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.5+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CGColorCreateWithContentHeadroom(headroom float, space ,  CGColorSpaceRef, red ,  CGFloat, green ,  CGFloat, blue ,  CGFloat, alpha ,  CGFloat, ) CGColorRef
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGColorGetPattern(color CGColorRef, ) CGPatternRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGColorConversionInfoConvertData(info CGColorConversionInfoRef, width ,  size_t, height ,  size_t, dst_data ,  void  *, dst_format ,  CGColorBufferFormat, src_data ,  const void  *, src_format ,  CGColorBufferFormat, options ,  CFDictionaryRef, ) bool
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+


// CGColorConversionInfoCreateForToneMapping(from CGColorSpaceRef, source_headroom ,  float, to ,  CGColorSpaceRef, target_headroom ,  float, method ,  CGToneMapping, options ,  CFDictionaryRef, error ,  CFErrorRef  *, ) CGColorConversionInfoRef
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+

// CGColorConversionInfoCreateFromList(options CFDictionaryRef, ,  CGColorSpaceRef, ,  CGColorConversionInfoTransformType, ,  CGColorRenderingIntent, , ...) CGColorConversionInfoRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+

// CGColorRelease(color CGColorRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGColorSpaceCreateCalibratedRGB(whitePoint const CGFloat, blackPoint [ 3 ],  const CGFloat, gamma [ 3 ],  const CGFloat, matrix [ 3 ],  const CGFloat, [ 9 ]) CGColorSpaceRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGColorSpaceCreatePattern(baseSpace CGColorSpaceRef, ) CGColorSpaceRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGColorSpaceIsWideGamutRGB(CGColorSpaceRef, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+


// CGColorSpaceCreateLinearized(space CGColorSpaceRef, ) CGColorSpaceRef
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// CGColorSpaceIsHLGBased(s CGColorSpaceRef, ) bool
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// CGColorSpaceUsesExtendedRange(space CGColorSpaceRef, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+


// CGContextAddPath(c CGContextRef, path ,  CGPathRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGBitmapContextGetBytesPerRow(context CGContextRef, ) size_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGContextDrawShading(c CGContextRef, shading ,  CGShadingRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFContextEndPage(context CGContextRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGBitmapContextCreate(data void  *, width ,  size_t, height ,  size_t, bitsPerComponent ,  size_t, bytesPerRow ,  size_t, space ,  CGColorSpaceRef, bitmapInfo ,  CGBitmapInfo, ) CGContextRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGBitmapContextCreateWithData(data void  *, width ,  size_t, height ,  size_t, bitsPerComponent ,  size_t, bytesPerRow ,  size_t, space ,  CGColorSpaceRef, bitmapInfo ,  CGBitmapInfo, releaseCallback ,  CGBitmapContextReleaseDataCallback, releaseInfo ,  void  *, ) CGContextRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGContextGetInterpolationQuality(c CGContextRef, ) CGInterpolationQuality
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGContextRotateCTM(c CGContextRef, angle ,  CGFloat, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGContextScaleCTM(c CGContextRef, sx ,  CGFloat, sy ,  CGFloat, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFContextSetDestinationForRect(context CGContextRef, name ,  CFStringRef, rect ,  CGRect, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGContextSetShouldSubpixelPositionFonts(c CGContextRef, shouldSubpixelPositionFonts ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGContextSetStrokeColor(c CGContextRef, components ,  const CGFloat  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGContextSynchronizeAttributes(c CGContextRef, )
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGContextGetTextMatrix(c CGContextRef, ) CGAffineTransform
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGContextGetContentToneMappingInfo(c CGContextRef, ) CGContentToneMappingInfo
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+


// CGContextMoveToPoint(c CGContextRef, x ,  CGFloat, y ,  CGFloat, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGContextSetContentToneMappingInfo(c CGContextRef, info ,  CGContentToneMappingInfo, )
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGContextSetLineDash(c CGContextRef, phase ,  CGFloat, lengths ,  const CGFloat  *, count ,  size_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGContextStrokeLineSegments(c CGContextRef, points ,  const CGPoint  *, count ,  size_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGDataConsumerCreateWithCFData(data CFMutableDataRef, ) CGDataConsumerRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGDataConsumerCreateWithURL(url CFURLRef, ) CGDataConsumerRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGDataConsumerRelease(consumer CGDataConsumerRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGDataConsumerRetain(consumer CGDataConsumerRef, ) CGDataConsumerRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGDataProviderCopyData(provider CGDataProviderRef, ) CFDataRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGDataProviderGetInfo(provider CGDataProviderRef, ) void  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+

// CGDataProviderCreateWithData(info void  *, data ,  const void  *, size ,  size_t, releaseData ,  CGDataProviderReleaseDataCallback, ) CGDataProviderRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGDataProviderCreateDirect(info void  *, size ,  off_t, callbacks ,  const CGDataProviderDirectCallbacks  *, ) CGDataProviderRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGDataProviderCreateWithFilename(filename const char  *, ) CGDataProviderRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGDataProviderCreateSequential(info void  *, callbacks ,  const CGDataProviderSequentialCallbacks  *, ) CGDataProviderRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGDataProviderRetain(provider CGDataProviderRef, ) CGDataProviderRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGDisplayCaptureWithOptions(display CGDirectDisplayID, options ,  CGCaptureOptions, ) CGError
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.3+

// CGDisplayIsActive(display CGDirectDisplayID, ) boolean_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.2+

// CGDisplayModeGetPixelHeight(mode CGDisplayModeRef, ) size_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.8+


// CGDisplayRemoveReconfigurationCallback(callback CGDisplayReconfigurationCallBack, userInfo ,  void  *, ) CGError
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.3+

// CGDisplayStreamUpdateGetRects(updateRef CGDisplayStreamUpdateRef, rectType ,  CGDisplayStreamUpdateRectType, rectCount ,  size_t  *, ) const CGRect  *

// CGEventCreateMouseEvent(source CGEventSourceRef, mouseType ,  CGEventType, mouseCursorPosition ,  CGPoint, mouseButton ,  CGMouseButton, ) CGEventRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+


// CGEventKeyboardGetUnicodeString(event CGEventRef, maxStringLength ,  UniCharCount, actualStringLength ,  UniCharCount  *, unicodeString ,  UniChar  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+

// CGEventKeyboardSetUnicodeString(event CGEventRef, stringLength ,  UniCharCount, unicodeString ,  const UniChar  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+

// CGEventTapCreate(tap CGEventTapLocation, place ,  CGEventTapPlacement, options ,  CGEventTapOptions, eventsOfInterest ,  CGEventMask, callback ,  CGEventTapCallBack, userInfo ,  void  *, ) CFMachPortRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+


// CGEventTapCreateForPSN(processSerialNumber void  *, place ,  CGEventTapPlacement, options ,  CGEventTapOptions, eventsOfInterest ,  CGEventMask, callback ,  CGEventTapCallBack, userInfo ,  void  *, ) CFMachPortRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+

// CGEventTapPostEvent(proxy CGEventTapProxy, event ,  CGEventRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+

// CGEventCreateScrollWheelEvent(source CGEventSourceRef, units ,  CGScrollEventUnit, wheelCount ,  uint32_t, wheel1 ,  int32_t, , ...) CGEventRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.5+


// CGEventSetFlags(event CGEventRef, flags ,  CGEventFlags, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+

// CGEventSourceCounterForEventType(stateID CGEventSourceStateID, eventType ,  CGEventType, ) uint32_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+

// CGEventSourceFlagsState(stateID CGEventSourceStateID, ) CGEventFlags
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+


// CGEventSourceGetLocalEventsSuppressionInterval(source CGEventSourceRef, ) CFTimeInterval
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+

// CGFontGetCapHeight(font CGFontRef, ) int
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGFontCreateCopyWithVariations(font CGFontRef, variations ,  CFDictionaryRef, ) CGFontRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGFontCreatePostScriptEncoding(font CGFontRef, encoding ,  const CGGlyph, [ 256 ]) CFDataRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGFontGetDescent(font CGFontRef, ) int
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGFontGetGlyphWithGlyphName(font CGFontRef, name ,  CFStringRef, ) CGGlyph
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGFontGetLeading(font CGFontRef, ) int
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGFontGetStemV(font CGFontRef, ) CGFloat
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGFontCopyTableForTag(font CGFontRef, tag ,  uint32_t, ) CFDataRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGFontGetXHeight(font CGFontRef, ) int
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGFontCreateWithPlatformFont(platformFontReference void  *, ) CGFontRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// CGFontRelease(font CGFontRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGFunctionCreate(info void  *, domainDimension ,  size_t, domain ,  const CGFloat  *, rangeDimension ,  size_t, range ,  const CGFloat  *, callbacks ,  const CGFunctionCallbacks  *, ) CGFunctionRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGFunctionRelease(function CGFunctionRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGFunctionRetain(function CGFunctionRef, ) CGFunctionRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGGradientGetContentHeadroom(gradient CGGradientRef, ) float
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGGradientCreateWithColorComponents(space CGColorSpaceRef, components ,  const CGFloat  *, locations ,  const CGFloat  *, count ,  size_t, ) CGGradientRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGGradientCreateWithColors(space CGColorSpaceRef, colors ,  CFArrayRef, locations ,  const CGFloat  *, ) CGGradientRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGGradientCreateWithContentHeadroom(headroom float, space ,  CGColorSpaceRef, components ,  const CGFloat  *, locations ,  const CGFloat  *, count ,  size_t, ) CGGradientRef
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGGradientRelease(gradient CGGradientRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGGradientRetain(gradient CGGradientRef, ) CGGradientRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGImageGetAlphaInfo(image CGImageRef, ) CGImageAlphaInfo
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageGetBitsPerComponent(image CGImageRef, ) size_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageGetByteOrderInfo(image CGImageRef, ) CGImageByteOrderInfo
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// CGImageGetBytesPerRow(image CGImageRef, ) size_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageCalculateContentAverageLightLevel(image CGImageRef, ) float
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGImageCreateWithJPEGDataProvider(source CGDataProviderRef, decode ,  const CGFloat  *, shouldInterpolate ,  bool, intent ,  CGColorRenderingIntent, ) CGImageRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.1+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGImageCreate(width size_t, height ,  size_t, bitsPerComponent ,  size_t, bitsPerPixel ,  size_t, bytesPerRow ,  size_t, space ,  CGColorSpaceRef, bitmapInfo ,  CGBitmapInfo, provider ,  CGDataProviderRef, decode ,  const CGFloat  *, shouldInterpolate ,  bool, intent ,  CGColorRenderingIntent, ) CGImageRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGWindowListCreateImageFromArray(screenBounds CGRect, windowArray ,  CFArrayRef, imageOption ,  CGWindowImageOption, ) CGImageRef

// CGImageIsMask(image CGImageRef, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGImageCreateWithMask(image CGImageRef, mask ,  CGImageRef, ) CGImageRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageRetain(image CGImageRef, ) CGImageRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGLayerGetContext(layer CGLayerRef, ) CGContextRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGLayerCreateWithContext(context CGContextRef, size ,  CGSize, auxiliaryInfo ,  CFDictionaryRef, ) CGLayerRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGLayerGetSize(layer CGLayerRef, ) CGSize
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGLayerRelease(layer CGLayerRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGLayerRetain(layer CGLayerRef, ) CGLayerRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPathCloseSubpath(path CGMutablePathRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFArrayGetBoolean(array CGPDFArrayRef, index ,  size_t, value ,  CGPDFBoolean  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFArrayGetDictionary(array CGPDFArrayRef, index ,  size_t, value ,  CGPDFDictionaryRef  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFArrayGetInteger(array CGPDFArrayRef, index ,  size_t, value ,  CGPDFInteger  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFArrayGetNull(array CGPDFArrayRef, index ,  size_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFArrayGetObject(array CGPDFArrayRef, index ,  size_t, value ,  CGPDFObjectRef  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFArrayGetString(array CGPDFArrayRef, index ,  size_t, value ,  CGPDFStringRef  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFContentStreamGetResource(cs CGPDFContentStreamRef, category ,  const char  *, name ,  const char  *, ) CGPDFObjectRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFContentStreamRelease(cs CGPDFContentStreamRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFContentStreamRetain(cs CGPDFContentStreamRef, ) CGPDFContentStreamRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFContextBeginTag(context CGContextRef, tagType ,  CGPDFTagType, tagProperties ,  CFDictionaryRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CGPDFDictionaryApplyFunction(dict CGPDFDictionaryRef, function ,  CGPDFDictionaryApplierFunction, info ,  void  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFDictionaryGetArray(dict CGPDFDictionaryRef, key ,  const char  *, value ,  CGPDFArrayRef  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFDictionaryGetBoolean(dict CGPDFDictionaryRef, key ,  const char  *, value ,  CGPDFBoolean  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFDictionaryGetCount(dict CGPDFDictionaryRef, ) size_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFDictionaryGetDictionary(dict CGPDFDictionaryRef, key ,  const char  *, value ,  CGPDFDictionaryRef  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFDictionaryGetObject(dict CGPDFDictionaryRef, key ,  const char  *, value ,  CGPDFObjectRef  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFDictionaryGetStream(dict CGPDFDictionaryRef, key ,  const char  *, value ,  CGPDFStreamRef  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFDictionaryGetString(dict CGPDFDictionaryRef, key ,  const char  *, value ,  CGPDFStringRef  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFDocumentGetAccessPermissions(document CGPDFDocumentRef, ) CGPDFAccessPermissions
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+


// CGPDFDocumentAllowsCopying(document CGPDFDocumentRef, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFDocumentGetID(document CGPDFDocumentRef, ) CGPDFArrayRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFDocumentCreateWithProvider(provider CGDataProviderRef, ) CGPDFDocumentRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFDocumentGetNumberOfPages(document CGPDFDocumentRef, ) size_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFDocumentGetOutline(document CGPDFDocumentRef, ) CFDictionaryRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+

// CGPDFDocumentUnlockWithPassword(document CGPDFDocumentRef, password ,  const char  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFDocumentGetCropBox(document CGPDFDocumentRef, page ,  int, ) CGRect
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.5)
//
// Deprecated: This function is deprecated.

// CGPDFDocumentGetTrimBox(document CGPDFDocumentRef, page ,  int, ) CGRect
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.5)
//
// Deprecated: This function is deprecated.

// CGPDFDocumentRetain(document CGPDFDocumentRef, ) CGPDFDocumentRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFObjectGetType(object CGPDFObjectRef, ) CGPDFObjectType
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFObjectGetValue(object CGPDFObjectRef, type ,  CGPDFObjectType, value ,  void  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFOperatorTableRetain(table CGPDFOperatorTableRef, ) CGPDFOperatorTableRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFOperatorTableSetCallback(table CGPDFOperatorTableRef, name ,  const char  *, callback ,  CGPDFOperatorCallback, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFPageGetDictionary(page CGPDFPageRef, ) CGPDFDictionaryRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFPageGetDocument(page CGPDFPageRef, ) CGPDFDocumentRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFPageGetBoxRect(page CGPDFPageRef, box ,  CGPDFBox, ) CGRect
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFPageGetDrawingTransform(page CGPDFPageRef, box ,  CGPDFBox, rect ,  CGRect, rotate ,  int, preserveAspectRatio ,  bool, ) CGAffineTransform
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFPageGetPageNumber(page CGPDFPageRef, ) size_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFPageGetRotationAngle(page CGPDFPageRef, ) int
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFPageRelease(page CGPDFPageRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFPageRetain(page CGPDFPageRef, ) CGPDFPageRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFScannerCreate(cs CGPDFContentStreamRef, table ,  CGPDFOperatorTableRef, info ,  void  *, ) CGPDFScannerRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFScannerPopArray(scanner CGPDFScannerRef, value ,  CGPDFArrayRef  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFScannerPopDictionary(scanner CGPDFScannerRef, value ,  CGPDFDictionaryRef  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFScannerPopName(scanner CGPDFScannerRef, value ,  const char  * *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFScannerPopNumber(scanner CGPDFScannerRef, value ,  CGPDFReal  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFScannerPopString(scanner CGPDFScannerRef, value ,  CGPDFStringRef  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFScannerRelease(scanner CGPDFScannerRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFScannerRetain(scanner CGPDFScannerRef, ) CGPDFScannerRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFStreamCopyData(stream CGPDFStreamRef, format ,  CGPDFDataFormat  *, ) CFDataRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFStreamGetDictionary(stream CGPDFStreamRef, ) CGPDFDictionaryRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFStringCopyDate(string CGPDFStringRef, ) CFDateRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFStringCopyTextString(string CGPDFStringRef, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPDFStringGetBytePtr(string CGPDFStringRef, ) const unsigned char  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPDFStringGetLength(string CGPDFStringRef, ) size_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.3+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPSConverterAbort(converter CGPSConverterRef, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.3+


// CGPSConverterConvert(converter CGPSConverterRef, provider ,  CGDataProviderRef, consumer ,  CGDataConsumerRef, options ,  CFDictionaryRef, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.3+

// CGPSConverterCreate(info void  *, callbacks ,  const CGPSConverterCallbacks  *, options ,  CFDictionaryRef, ) CGPSConverterRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.3+

// CGPathGetBoundingBox(path CGPathRef, ) CGRect
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPathGetCurrentPoint(path CGPathRef, ) CGPoint
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPathCreateWithRect(rect CGRect, transform ,  const CGAffineTransform  *, ) CGPathRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPathCreateWithRoundedRect(rect CGRect, cornerWidth ,  CGFloat, cornerHeight ,  CGFloat, transform ,  const CGAffineTransform  *, ) CGPathRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPathIsRect(path CGPathRef, rect ,  CGRect  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPathCreateMutableCopy(path CGPathRef, ) CGMutablePathRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPathAddArcToPoint(path CGMutablePathRef, m ,  const CGAffineTransform  *, x1 ,  CGFloat, y1 ,  CGFloat, x2 ,  CGFloat, y2 ,  CGFloat, radius ,  CGFloat, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPathAddCurveToPoint(path CGMutablePathRef, m ,  const CGAffineTransform  *, cp1x ,  CGFloat, cp1y ,  CGFloat, cp2x ,  CGFloat, cp2y ,  CGFloat, x ,  CGFloat, y ,  CGFloat, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPathAddLineToPoint(path CGMutablePathRef, m ,  const CGAffineTransform  *, x ,  CGFloat, y ,  CGFloat, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPathAddLines(path CGMutablePathRef, m ,  const CGAffineTransform  *, points ,  const CGPoint  *, count ,  size_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPathAddPath(path1 CGMutablePathRef, m ,  const CGAffineTransform  *, path2 ,  CGPathRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPathAddQuadCurveToPoint(path CGMutablePathRef, m ,  const CGAffineTransform  *, cpx ,  CGFloat, cpy ,  CGFloat, x ,  CGFloat, y ,  CGFloat, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPathCreateCopyByDashingPath(path CGPathRef, transform ,  const CGAffineTransform  *, phase ,  CGFloat, lengths ,  const CGFloat  *, count ,  size_t, ) CGPathRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPathCreateCopyByFlattening(path CGPathRef, flatteningThreshold ,  CGFloat, ) CGPathRef
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// CGPathCreateCopyByStrokingPath(path CGPathRef, transform ,  const CGAffineTransform  *, lineWidth ,  CGFloat, lineCap ,  CGLineCap, lineJoin ,  CGLineJoin, miterLimit ,  CGFloat, ) CGPathRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPathCreateCopyBySymmetricDifferenceOfPath(path CGPathRef, maskPath ,  CGPathRef, evenOddFillRule ,  bool, ) CGPathRef
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+


// CGPathMoveToPoint(path CGMutablePathRef, m ,  const CGAffineTransform  *, x ,  CGFloat, y ,  CGFloat, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPatternCreate(info void  *, bounds ,  CGRect, matrix ,  CGAffineTransform, xStep ,  CGFloat, yStep ,  CGFloat, tiling ,  CGPatternTiling, isColored ,  bool, callbacks ,  const CGPatternCallbacks  *, ) CGPatternRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPatternRelease(pattern CGPatternRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGPatternRetain(pattern CGPatternRef, ) CGPatternRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGPostMouseEvent(mouseCursorPosition CGPoint, updateMouseCursorPosition ,  boolean_t, buttonCount ,  CGButtonCount, mouseButtonDown ,  boolean_t, , ...) CGError
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// CGRectApplyAffineTransform(rect CGRect, t ,  CGAffineTransform, ) CGRect
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGRectEqualToRect(rect1 CGRect, rect2 ,  CGRect, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGRectGetMaxY(rect CGRect, ) CGFloat
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGRectGetMinX(rect CGRect, ) CGFloat
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGRectGetMinY(rect CGRect, ) CGFloat
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGRectIsNull(rect CGRect, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGRectOffset(rect CGRect, dx ,  CGFloat, dy ,  CGFloat, ) CGRect
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGRenderingBufferLockBytePtr(provider CGRenderingBufferProviderRef, ) void  *
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGRenderingBufferProviderCreate(info void  *, size ,  size_t, lockPointer ,  void  * (^, info )( void  *, unlockPointer ),  void  (^, info )( void  *, pointer ,  void  *, releaseInfo ),  void  (^, info )( void  *, ) CGRenderingBufferProviderRef
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGRenderingBufferProviderCreateWithCFData(data CFMutableDataRef, ) CGRenderingBufferProviderRef
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+


// CGRenderingBufferProviderGetSize(provider CGRenderingBufferProviderRef, ) size_t
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGRenderingBufferUnlockBytePtr(provider CGRenderingBufferProviderRef, )
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGShadingGetContentHeadroom(shading CGShadingRef, ) float
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+


// CGShadingCreateAxialWithContentHeadroom(headroom float, space ,  CGColorSpaceRef, start ,  CGPoint, end ,  CGPoint, function ,  CGFunctionRef, extendStart ,  bool, extendEnd ,  bool, ) CGShadingRef
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CGShadingCreateAxial(space CGColorSpaceRef, start ,  CGPoint, end ,  CGPoint, function ,  CGFunctionRef, extendStart ,  bool, extendEnd ,  bool, ) CGShadingRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGShadingCreateRadialWithContentHeadroom(headroom float, space ,  CGColorSpaceRef, start ,  CGPoint, startRadius ,  CGFloat, end ,  CGPoint, endRadius ,  CGFloat, function ,  CGFunctionRef, extendStart ,  bool, extendEnd ,  bool, ) CGShadingRef
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+


// CGShadingCreateRadial(space CGColorSpaceRef, start ,  CGPoint, startRadius ,  CGFloat, end ,  CGPoint, endRadius ,  CGFloat, function ,  CGFunctionRef, extendStart ,  bool, extendEnd ,  bool, ) CGShadingRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGShadingRelease(shading CGShadingRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGShadingRetain(shading CGShadingRef, ) CGShadingRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGSizeApplyAffineTransform(size CGSize, t ,  CGAffineTransform, ) CGSize
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGSizeCreateDictionaryRepresentation(size CGSize, ) CFDictionaryRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGSizeMakeWithDictionaryRepresentation(dict CFDictionaryRef, size ,  CGSize  *, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGWaitForScreenRefreshRects(rects CGRect  * *, count ,  uint32_t  *, ) CGError
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//
// Deprecated: This function is deprecated.

// CGWindowListCopyWindowInfo(option CGWindowListOption, relativeToWindow ,  CGWindowID, ) CFArrayRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.5+

// CGWindowListCreateDescriptionFromArray(windowArray CFArrayRef, ) CFArrayRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.5+


