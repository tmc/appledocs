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

// Discovered functions (23 total):

// CGContextSetInterpolationQuality(CGContextRef c,  CGInterpolationQuality quality);)
// CGBitmapContextCreateAdaptive(width size_t, height ,  size_t, auxiliaryInfo ,  CFDictionaryRef, onResolve ,  bool (^, )( const  CGContentInfo *, ,  CGBitmapParameters *, onAllocate ),  CGRenderingBufferProviderRef (^, )( const  CGContentInfo *, ,  const  CGBitmapParameters *, onFree ),  void (^, )( CGRenderingBufferProviderRef, ,  const  CGContentInfo *, ,  const  CGBitmapParameters *, onError ),  void (^, )( CFErrorRef, ,  const  CGContentInfo *, ,  const  CGBitmapParameters *, ));) extern   CGContextRef
// CGBitmapInfoMake(alpha CGImageAlphaInfo, component ,  CGImageComponentInfo, byteOrder ,  CGImageByteOrderInfo, pixelFormat ,  CGImagePixelFormatInfo, );) static   CGBitmapInfo
// contentHeadroom() var
// synchronizeAttributes() func
// CGContextGetContentToneMappingInfo(c CGContextRef, );) extern   CGContentToneMappingInfo

// CGContextSetContentToneMappingInfo(c CGContextRef, info ,  CGContentToneMappingInfo, );) extern   void
// CGEXRToneMappingGammaGetDefaultOptions() extern   CFDictionaryRef
// contentHeadroom() var
// CGRenderingBufferLockBytePtr(provider CGRenderingBufferProviderRef, );) extern   void  *

// CGRenderingBufferProviderCreate(info void *, size ,  size_t, lockPointer ,  void * (^, info )( void *, unlockPointer ),  void (^, info )( void *, pointer ,  void *, releaseInfo ),  void (^, info )( void *, ));) extern   CGRenderingBufferProviderRef
// CGRenderingBufferProviderCreateWithCFData(data CFMutableDataRef, );) extern   CGRenderingBufferProviderRef
// CGRenderingBufferProviderGetSize(provider CGRenderingBufferProviderRef, );) extern   size_t
// CGRenderingBufferProviderGetTypeID() extern   CFTypeID
// CGRenderingBufferUnlockBytePtr(provider CGRenderingBufferProviderRef, );) extern   void

// contentHeadroom() var
