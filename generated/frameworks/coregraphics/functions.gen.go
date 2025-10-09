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

// Discovered functions (22 total):

// CGBitmapContextCreateAdaptive(width size_t, height ,  size_t, auxiliaryInfo ,  CFDictionaryRef, onResolve ,  bool (^, )( const  CGContentInfo *, ,  CGBitmapParameters *, onAllocate ),  CGRenderingBufferProviderRef (^, )( const  CGContentInfo *, ,  const  CGBitmapParameters *, onFree ),  void (^, )( CGRenderingBufferProviderRef, ,  const  CGContentInfo *, ,  const  CGBitmapParameters *, onError ),  void (^, )( CFErrorRef, ,  const  CGContentInfo *, ,  const  CGBitmapParameters *, ));) extern   CGContextRef
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+

// CGBitmapInfoMake(alpha CGImageAlphaInfo, component ,  CGImageComponentInfo, byteOrder ,  CGImageByteOrderInfo, pixelFormat ,  CGImagePixelFormatInfo, );) static   CGBitmapInfo
//
// Availability:
//   - macOS 10.14+
//   - iOS 12.0+
//   - watchOS 5.0+
//   - tvOS 12.0+

// contentHeadroom() var
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+


// synchronizeAttributes() func
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+

// CGContextGetContentToneMappingInfo(c CGContextRef, );) extern   CGContentToneMappingInfo
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+


// CGContextSetContentToneMappingInfo(c CGContextRef, info ,  CGContentToneMappingInfo, );) extern   void
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+

// CGEXRToneMappingGammaGetDefaultOptions() extern   CFDictionaryRef
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+

// contentHeadroom() var
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+


// CGRenderingBufferLockBytePtr(provider CGRenderingBufferProviderRef, );) extern   void  *
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+

// CGRenderingBufferProviderCreate(info void *, size ,  size_t, lockPointer ,  void * (^, info )( void *, unlockPointer ),  void (^, info )( void *, pointer ,  void *, releaseInfo ),  void (^, info )( void *, ));) extern   CGRenderingBufferProviderRef
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+


// CGRenderingBufferProviderCreateWithCFData(data CFMutableDataRef, );) extern   CGRenderingBufferProviderRef
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+

// CGRenderingBufferProviderGetSize(provider CGRenderingBufferProviderRef, );) extern   size_t
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+

// CGRenderingBufferProviderGetTypeID() extern   CFTypeID
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+


// CGRenderingBufferUnlockBytePtr(provider CGRenderingBufferProviderRef, );) extern   void
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+

// contentHeadroom() var
//
// Availability:
//   - macOS 26.0+
//   - iOS 26.0+
//   - watchOS 26.0+
//   - tvOS 26.0+

