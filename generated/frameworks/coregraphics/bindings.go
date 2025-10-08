// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import (
	"unsafe"
	"github.com/ebitengine/purego"
)

// Common types
type CGFloat float64
type CGContextRef unsafe.Pointer
type CGColorRef unsafe.Pointer
type CGColorSpaceRef unsafe.Pointer
type CGPathRef unsafe.Pointer
type CGImageRef unsafe.Pointer

var lib uintptr

func init() {
	var err error
	lib, err = purego.Dlopen("/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

// CoreGraphics Functions

// CGContextSetInterpolationQuality( CGContextRef c,  CGInterpolationQuality quality);) -> void
// CGBitmapContextCreateAdaptive(width size_t, height ,  size_t, auxiliaryInfo ,  CFDictionaryRef, onResolve ,  bool (^,  )( const  CGContentInfo *,  ,  CGBitmapParameters *, onAllocate ),  CGRenderingBufferProviderRef (^,  )( const  CGContentInfo *,  ,  const  CGBitmapParameters *, onFree ),  void (^,  )( CGRenderingBufferProviderRef,  ,  const  CGContentInfo *,  ,  const  CGBitmapParameters *, onError ),  void (^,  )( CFErrorRef,  ,  const  CGContentInfo *,  ,  const  CGBitmapParameters *,  ));) -> extern   CGContextRef
// CGBitmapInfoMake(alpha CGImageAlphaInfo, component ,  CGImageComponentInfo, byteOrder ,  CGImageByteOrderInfo, pixelFormat ,  CGImagePixelFormatInfo,  );) -> static   CGBitmapInfo
// contentHeadroom -> var
// synchronizeAttributes -> func
// CGContextGetContentToneMappingInfo(c CGContextRef,  );) -> extern   CGContentToneMappingInfo
// CGContextSetContentToneMappingInfo(c CGContextRef, info ,  CGContentToneMappingInfo,  );) -> extern   void
// CGEXRToneMappingGammaGetDefaultOptions -> extern   CFDictionaryRef
// contentHeadroom -> var
// CGRenderingBufferLockBytePtr(provider CGRenderingBufferProviderRef,  );) -> extern   void  *
// CGRenderingBufferProviderCreate(info void *, size ,  size_t, lockPointer ,  void * (^, info )( void *, unlockPointer ),  void (^, info )( void *, pointer ,  void *, releaseInfo ),  void (^, info )( void *,  ));) -> extern   CGRenderingBufferProviderRef
// CGRenderingBufferProviderCreateWithCFData(data CFMutableDataRef,  );) -> extern   CGRenderingBufferProviderRef
// CGRenderingBufferProviderGetSize(provider CGRenderingBufferProviderRef,  );) -> extern   size_t
// CGRenderingBufferProviderGetTypeID -> extern   CFTypeID
// CGRenderingBufferUnlockBytePtr(provider CGRenderingBufferProviderRef,  );) -> extern   void

// Add bindings here using purego.RegisterLibFunc
