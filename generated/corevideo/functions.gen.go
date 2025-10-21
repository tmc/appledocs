// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// CoreVideo Functions (134 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CVBufferCopyAttachment func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVBufferCopyAttachments func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVBufferGetAttachment func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVBufferGetAttachments func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVBufferHasAttachment func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVBufferPropagateAttachments func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVBufferRelease func(unsafe.Pointer) unsafe.Pointer
	_CVBufferRemoveAllAttachments func(unsafe.Pointer) unsafe.Pointer
	_CVBufferRemoveAttachment func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVBufferRetain func(unsafe.Pointer) unsafe.Pointer
	_CVBufferSetAttachment func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVBufferSetAttachments func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVColorPrimariesGetIntegerCodePointForString func(unsafe.Pointer) unsafe.Pointer
	_CVColorPrimariesGetStringForIntegerCodePoint func(unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkCreateWithActiveCGDisplays func(unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkCreateWithCGDisplay func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkCreateWithCGDisplays func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkCreateWithOpenGLDisplayMask func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkGetActualOutputVideoRefreshPeriod func(unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkGetCurrentCGDisplay func(unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkGetCurrentTime func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkGetNominalOutputVideoRefreshPeriod func(unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkGetOutputVideoLatency func(unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkGetTypeID func() unsafe.Pointer
	_CVDisplayLinkIsRunning func(unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkRelease func(unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkRetain func(unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkSetCurrentCGDisplay func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkSetOutputCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkSetOutputHandler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkStart func(unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkStop func(unsafe.Pointer) unsafe.Pointer
	_CVDisplayLinkTranslateTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVGetCurrentHostTime func() unsafe.Pointer
	_CVGetHostClockFrequency func() unsafe.Pointer
	_CVGetHostClockMinimumTimeDelta func() unsafe.Pointer
	_CVImageBufferCreateColorSpaceFromAttachments func(unsafe.Pointer) coregraphics.CGColorSpaceRef
	_CVImageBufferGetCleanRect func(unsafe.Pointer) coregraphics.CGRect
	_CVImageBufferGetColorSpace func(unsafe.Pointer) coregraphics.CGColorSpaceRef
	_CVImageBufferGetDisplaySize func(unsafe.Pointer) coregraphics.CGSize
	_CVImageBufferGetEncodedSize func(unsafe.Pointer) coregraphics.CGSize
	_CVImageBufferIsFlipped func(unsafe.Pointer) unsafe.Pointer
	_CVIsCompressedPixelFormatAvailable func(unsafe.Pointer) unsafe.Pointer
	_CVMetalBufferCacheCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVMetalBufferCacheCreateBufferFromImage func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVMetalBufferCacheFlush func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVMetalBufferCacheGetTypeID func() unsafe.Pointer
	_CVMetalBufferGetBuffer func(unsafe.Pointer) unsafe.Pointer
	_CVMetalBufferGetTypeID func() unsafe.Pointer
	_CVMetalTextureCacheCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVMetalTextureCacheCreateTextureFromImage func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVMetalTextureCacheFlush func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVMetalTextureCacheGetTypeID func() unsafe.Pointer
	_CVMetalTextureGetCleanTexCoords func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVMetalTextureGetTexture func(unsafe.Pointer) unsafe.Pointer
	_CVMetalTextureGetTypeID func() unsafe.Pointer
	_CVMetalTextureIsFlipped func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLBufferAttach func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLBufferCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLBufferGetAttributes func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLBufferGetTypeID func() unsafe.Pointer
	_CVOpenGLBufferPoolCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLBufferPoolCreateOpenGLBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLBufferPoolGetAttributes func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLBufferPoolGetOpenGLBufferAttributes func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLBufferPoolGetTypeID func() unsafe.Pointer
	_CVOpenGLBufferPoolRelease func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLBufferPoolRetain func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLBufferRelease func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLBufferRetain func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLESTextureCacheCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLESTextureCacheCreateTextureFromImage func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLESTextureCacheFlush func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLESTextureCacheGetTypeID func() unsafe.Pointer
	_CVOpenGLESTextureGetCleanTexCoords func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLESTextureGetName func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLESTextureGetTarget func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLESTextureGetTypeID func() unsafe.Pointer
	_CVOpenGLESTextureIsFlipped func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLTextureCacheCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLTextureCacheCreateTextureFromImage func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLTextureCacheFlush func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLTextureCacheGetTypeID func() unsafe.Pointer
	_CVOpenGLTextureCacheRelease func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLTextureCacheRetain func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLTextureGetCleanTexCoords func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVOpenGLTextureGetName func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLTextureGetTarget func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLTextureGetTypeID func() unsafe.Pointer
	_CVOpenGLTextureIsFlipped func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLTextureRelease func(unsafe.Pointer) unsafe.Pointer
	_CVOpenGLTextureRetain func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferCopyCreationAttributes func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferCreateResolvedAttributesDictionary func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferCreateWithBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferCreateWithIOSurface func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferCreateWithPlanarBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferFillExtendedPixels func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetBaseAddress func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetBaseAddressOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetBytesPerRow func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetBytesPerRowOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetDataSize func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetExtendedPixels func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetHeight func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetHeightOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetIOSurface func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetPixelFormatType func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetPlaneCount func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetTypeID func() unsafe.Pointer
	_CVPixelBufferGetWidth func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferGetWidthOfPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferIsCompatibleWithAttributes func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferIsPlanar func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferLockBaseAddress func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferPoolCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferPoolCreatePixelBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferPoolCreatePixelBufferWithAuxAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferPoolFlush func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferPoolGetAttributes func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferPoolGetPixelBufferAttributes func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferPoolGetTypeID func() unsafe.Pointer
	_CVPixelBufferPoolRelease func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferPoolRetain func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferRelease func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferRetain func(unsafe.Pointer) unsafe.Pointer
	_CVPixelBufferUnlockBaseAddress func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CVPixelFormatTypeCopyFourCharCodeString func(unsafe.Pointer) unsafe.Pointer
	_CVTransferFunctionGetIntegerCodePointForString func(unsafe.Pointer) unsafe.Pointer
	_CVTransferFunctionGetStringForIntegerCodePoint func(unsafe.Pointer) unsafe.Pointer
	_CVYCbCrMatrixGetIntegerCodePointForString func(unsafe.Pointer) unsafe.Pointer
	_CVYCbCrMatrixGetStringForIntegerCodePoint func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CVBufferCopyAttachment, lib, "CVBufferCopyAttachment")
	tryRegister(&_CVBufferCopyAttachments, lib, "CVBufferCopyAttachments")
	tryRegister(&_CVBufferGetAttachment, lib, "CVBufferGetAttachment")
	tryRegister(&_CVBufferGetAttachments, lib, "CVBufferGetAttachments")
	tryRegister(&_CVBufferHasAttachment, lib, "CVBufferHasAttachment")
	tryRegister(&_CVBufferPropagateAttachments, lib, "CVBufferPropagateAttachments")
	tryRegister(&_CVBufferRelease, lib, "CVBufferRelease")
	tryRegister(&_CVBufferRemoveAllAttachments, lib, "CVBufferRemoveAllAttachments")
	tryRegister(&_CVBufferRemoveAttachment, lib, "CVBufferRemoveAttachment")
	tryRegister(&_CVBufferRetain, lib, "CVBufferRetain")
	tryRegister(&_CVBufferSetAttachment, lib, "CVBufferSetAttachment")
	tryRegister(&_CVBufferSetAttachments, lib, "CVBufferSetAttachments")
	tryRegister(&_CVColorPrimariesGetIntegerCodePointForString, lib, "CVColorPrimariesGetIntegerCodePointForString")
	tryRegister(&_CVColorPrimariesGetStringForIntegerCodePoint, lib, "CVColorPrimariesGetStringForIntegerCodePoint")
	tryRegister(&_CVDisplayLinkCreateWithActiveCGDisplays, lib, "CVDisplayLinkCreateWithActiveCGDisplays")
	tryRegister(&_CVDisplayLinkCreateWithCGDisplay, lib, "CVDisplayLinkCreateWithCGDisplay")
	tryRegister(&_CVDisplayLinkCreateWithCGDisplays, lib, "CVDisplayLinkCreateWithCGDisplays")
	tryRegister(&_CVDisplayLinkCreateWithOpenGLDisplayMask, lib, "CVDisplayLinkCreateWithOpenGLDisplayMask")
	tryRegister(&_CVDisplayLinkGetActualOutputVideoRefreshPeriod, lib, "CVDisplayLinkGetActualOutputVideoRefreshPeriod")
	tryRegister(&_CVDisplayLinkGetCurrentCGDisplay, lib, "CVDisplayLinkGetCurrentCGDisplay")
	tryRegister(&_CVDisplayLinkGetCurrentTime, lib, "CVDisplayLinkGetCurrentTime")
	tryRegister(&_CVDisplayLinkGetNominalOutputVideoRefreshPeriod, lib, "CVDisplayLinkGetNominalOutputVideoRefreshPeriod")
	tryRegister(&_CVDisplayLinkGetOutputVideoLatency, lib, "CVDisplayLinkGetOutputVideoLatency")
	tryRegister(&_CVDisplayLinkGetTypeID, lib, "CVDisplayLinkGetTypeID")
	tryRegister(&_CVDisplayLinkIsRunning, lib, "CVDisplayLinkIsRunning")
	tryRegister(&_CVDisplayLinkRelease, lib, "CVDisplayLinkRelease")
	tryRegister(&_CVDisplayLinkRetain, lib, "CVDisplayLinkRetain")
	tryRegister(&_CVDisplayLinkSetCurrentCGDisplay, lib, "CVDisplayLinkSetCurrentCGDisplay")
	tryRegister(&_CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext, lib, "CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext")
	tryRegister(&_CVDisplayLinkSetOutputCallback, lib, "CVDisplayLinkSetOutputCallback")
	tryRegister(&_CVDisplayLinkSetOutputHandler, lib, "CVDisplayLinkSetOutputHandler")
	tryRegister(&_CVDisplayLinkStart, lib, "CVDisplayLinkStart")
	tryRegister(&_CVDisplayLinkStop, lib, "CVDisplayLinkStop")
	tryRegister(&_CVDisplayLinkTranslateTime, lib, "CVDisplayLinkTranslateTime")
	tryRegister(&_CVGetCurrentHostTime, lib, "CVGetCurrentHostTime")
	tryRegister(&_CVGetHostClockFrequency, lib, "CVGetHostClockFrequency")
	tryRegister(&_CVGetHostClockMinimumTimeDelta, lib, "CVGetHostClockMinimumTimeDelta")
	tryRegister(&_CVImageBufferCreateColorSpaceFromAttachments, lib, "CVImageBufferCreateColorSpaceFromAttachments")
	tryRegister(&_CVImageBufferGetCleanRect, lib, "CVImageBufferGetCleanRect")
	tryRegister(&_CVImageBufferGetColorSpace, lib, "CVImageBufferGetColorSpace")
	tryRegister(&_CVImageBufferGetDisplaySize, lib, "CVImageBufferGetDisplaySize")
	tryRegister(&_CVImageBufferGetEncodedSize, lib, "CVImageBufferGetEncodedSize")
	tryRegister(&_CVImageBufferIsFlipped, lib, "CVImageBufferIsFlipped")
	tryRegister(&_CVIsCompressedPixelFormatAvailable, lib, "CVIsCompressedPixelFormatAvailable")
	tryRegister(&_CVMetalBufferCacheCreate, lib, "CVMetalBufferCacheCreate")
	tryRegister(&_CVMetalBufferCacheCreateBufferFromImage, lib, "CVMetalBufferCacheCreateBufferFromImage")
	tryRegister(&_CVMetalBufferCacheFlush, lib, "CVMetalBufferCacheFlush")
	tryRegister(&_CVMetalBufferCacheGetTypeID, lib, "CVMetalBufferCacheGetTypeID")
	tryRegister(&_CVMetalBufferGetBuffer, lib, "CVMetalBufferGetBuffer")
	tryRegister(&_CVMetalBufferGetTypeID, lib, "CVMetalBufferGetTypeID")
	tryRegister(&_CVMetalTextureCacheCreate, lib, "CVMetalTextureCacheCreate")
	tryRegister(&_CVMetalTextureCacheCreateTextureFromImage, lib, "CVMetalTextureCacheCreateTextureFromImage")
	tryRegister(&_CVMetalTextureCacheFlush, lib, "CVMetalTextureCacheFlush")
	tryRegister(&_CVMetalTextureCacheGetTypeID, lib, "CVMetalTextureCacheGetTypeID")
	tryRegister(&_CVMetalTextureGetCleanTexCoords, lib, "CVMetalTextureGetCleanTexCoords")
	tryRegister(&_CVMetalTextureGetTexture, lib, "CVMetalTextureGetTexture")
	tryRegister(&_CVMetalTextureGetTypeID, lib, "CVMetalTextureGetTypeID")
	tryRegister(&_CVMetalTextureIsFlipped, lib, "CVMetalTextureIsFlipped")
	tryRegister(&_CVOpenGLBufferAttach, lib, "CVOpenGLBufferAttach")
	tryRegister(&_CVOpenGLBufferCreate, lib, "CVOpenGLBufferCreate")
	tryRegister(&_CVOpenGLBufferGetAttributes, lib, "CVOpenGLBufferGetAttributes")
	tryRegister(&_CVOpenGLBufferGetTypeID, lib, "CVOpenGLBufferGetTypeID")
	tryRegister(&_CVOpenGLBufferPoolCreate, lib, "CVOpenGLBufferPoolCreate")
	tryRegister(&_CVOpenGLBufferPoolCreateOpenGLBuffer, lib, "CVOpenGLBufferPoolCreateOpenGLBuffer")
	tryRegister(&_CVOpenGLBufferPoolGetAttributes, lib, "CVOpenGLBufferPoolGetAttributes")
	tryRegister(&_CVOpenGLBufferPoolGetOpenGLBufferAttributes, lib, "CVOpenGLBufferPoolGetOpenGLBufferAttributes")
	tryRegister(&_CVOpenGLBufferPoolGetTypeID, lib, "CVOpenGLBufferPoolGetTypeID")
	tryRegister(&_CVOpenGLBufferPoolRelease, lib, "CVOpenGLBufferPoolRelease")
	tryRegister(&_CVOpenGLBufferPoolRetain, lib, "CVOpenGLBufferPoolRetain")
	tryRegister(&_CVOpenGLBufferRelease, lib, "CVOpenGLBufferRelease")
	tryRegister(&_CVOpenGLBufferRetain, lib, "CVOpenGLBufferRetain")
	tryRegister(&_CVOpenGLESTextureCacheCreate, lib, "CVOpenGLESTextureCacheCreate")
	tryRegister(&_CVOpenGLESTextureCacheCreateTextureFromImage, lib, "CVOpenGLESTextureCacheCreateTextureFromImage")
	tryRegister(&_CVOpenGLESTextureCacheFlush, lib, "CVOpenGLESTextureCacheFlush")
	tryRegister(&_CVOpenGLESTextureCacheGetTypeID, lib, "CVOpenGLESTextureCacheGetTypeID")
	tryRegister(&_CVOpenGLESTextureGetCleanTexCoords, lib, "CVOpenGLESTextureGetCleanTexCoords")
	tryRegister(&_CVOpenGLESTextureGetName, lib, "CVOpenGLESTextureGetName")
	tryRegister(&_CVOpenGLESTextureGetTarget, lib, "CVOpenGLESTextureGetTarget")
	tryRegister(&_CVOpenGLESTextureGetTypeID, lib, "CVOpenGLESTextureGetTypeID")
	tryRegister(&_CVOpenGLESTextureIsFlipped, lib, "CVOpenGLESTextureIsFlipped")
	tryRegister(&_CVOpenGLTextureCacheCreate, lib, "CVOpenGLTextureCacheCreate")
	tryRegister(&_CVOpenGLTextureCacheCreateTextureFromImage, lib, "CVOpenGLTextureCacheCreateTextureFromImage")
	tryRegister(&_CVOpenGLTextureCacheFlush, lib, "CVOpenGLTextureCacheFlush")
	tryRegister(&_CVOpenGLTextureCacheGetTypeID, lib, "CVOpenGLTextureCacheGetTypeID")
	tryRegister(&_CVOpenGLTextureCacheRelease, lib, "CVOpenGLTextureCacheRelease")
	tryRegister(&_CVOpenGLTextureCacheRetain, lib, "CVOpenGLTextureCacheRetain")
	tryRegister(&_CVOpenGLTextureGetCleanTexCoords, lib, "CVOpenGLTextureGetCleanTexCoords")
	tryRegister(&_CVOpenGLTextureGetName, lib, "CVOpenGLTextureGetName")
	tryRegister(&_CVOpenGLTextureGetTarget, lib, "CVOpenGLTextureGetTarget")
	tryRegister(&_CVOpenGLTextureGetTypeID, lib, "CVOpenGLTextureGetTypeID")
	tryRegister(&_CVOpenGLTextureIsFlipped, lib, "CVOpenGLTextureIsFlipped")
	tryRegister(&_CVOpenGLTextureRelease, lib, "CVOpenGLTextureRelease")
	tryRegister(&_CVOpenGLTextureRetain, lib, "CVOpenGLTextureRetain")
	tryRegister(&_CVPixelBufferCopyCreationAttributes, lib, "CVPixelBufferCopyCreationAttributes")
	tryRegister(&_CVPixelBufferCreate, lib, "CVPixelBufferCreate")
	tryRegister(&_CVPixelBufferCreateResolvedAttributesDictionary, lib, "CVPixelBufferCreateResolvedAttributesDictionary")
	tryRegister(&_CVPixelBufferCreateWithBytes, lib, "CVPixelBufferCreateWithBytes")
	tryRegister(&_CVPixelBufferCreateWithIOSurface, lib, "CVPixelBufferCreateWithIOSurface")
	tryRegister(&_CVPixelBufferCreateWithPlanarBytes, lib, "CVPixelBufferCreateWithPlanarBytes")
	tryRegister(&_CVPixelBufferFillExtendedPixels, lib, "CVPixelBufferFillExtendedPixels")
	tryRegister(&_CVPixelBufferGetBaseAddress, lib, "CVPixelBufferGetBaseAddress")
	tryRegister(&_CVPixelBufferGetBaseAddressOfPlane, lib, "CVPixelBufferGetBaseAddressOfPlane")
	tryRegister(&_CVPixelBufferGetBytesPerRow, lib, "CVPixelBufferGetBytesPerRow")
	tryRegister(&_CVPixelBufferGetBytesPerRowOfPlane, lib, "CVPixelBufferGetBytesPerRowOfPlane")
	tryRegister(&_CVPixelBufferGetDataSize, lib, "CVPixelBufferGetDataSize")
	tryRegister(&_CVPixelBufferGetExtendedPixels, lib, "CVPixelBufferGetExtendedPixels")
	tryRegister(&_CVPixelBufferGetHeight, lib, "CVPixelBufferGetHeight")
	tryRegister(&_CVPixelBufferGetHeightOfPlane, lib, "CVPixelBufferGetHeightOfPlane")
	tryRegister(&_CVPixelBufferGetIOSurface, lib, "CVPixelBufferGetIOSurface")
	tryRegister(&_CVPixelBufferGetPixelFormatType, lib, "CVPixelBufferGetPixelFormatType")
	tryRegister(&_CVPixelBufferGetPlaneCount, lib, "CVPixelBufferGetPlaneCount")
	tryRegister(&_CVPixelBufferGetTypeID, lib, "CVPixelBufferGetTypeID")
	tryRegister(&_CVPixelBufferGetWidth, lib, "CVPixelBufferGetWidth")
	tryRegister(&_CVPixelBufferGetWidthOfPlane, lib, "CVPixelBufferGetWidthOfPlane")
	tryRegister(&_CVPixelBufferIsCompatibleWithAttributes, lib, "CVPixelBufferIsCompatibleWithAttributes")
	tryRegister(&_CVPixelBufferIsPlanar, lib, "CVPixelBufferIsPlanar")
	tryRegister(&_CVPixelBufferLockBaseAddress, lib, "CVPixelBufferLockBaseAddress")
	tryRegister(&_CVPixelBufferPoolCreate, lib, "CVPixelBufferPoolCreate")
	tryRegister(&_CVPixelBufferPoolCreatePixelBuffer, lib, "CVPixelBufferPoolCreatePixelBuffer")
	tryRegister(&_CVPixelBufferPoolCreatePixelBufferWithAuxAttributes, lib, "CVPixelBufferPoolCreatePixelBufferWithAuxAttributes")
	tryRegister(&_CVPixelBufferPoolFlush, lib, "CVPixelBufferPoolFlush")
	tryRegister(&_CVPixelBufferPoolGetAttributes, lib, "CVPixelBufferPoolGetAttributes")
	tryRegister(&_CVPixelBufferPoolGetPixelBufferAttributes, lib, "CVPixelBufferPoolGetPixelBufferAttributes")
	tryRegister(&_CVPixelBufferPoolGetTypeID, lib, "CVPixelBufferPoolGetTypeID")
	tryRegister(&_CVPixelBufferPoolRelease, lib, "CVPixelBufferPoolRelease")
	tryRegister(&_CVPixelBufferPoolRetain, lib, "CVPixelBufferPoolRetain")
	tryRegister(&_CVPixelBufferRelease, lib, "CVPixelBufferRelease")
	tryRegister(&_CVPixelBufferRetain, lib, "CVPixelBufferRetain")
	tryRegister(&_CVPixelBufferUnlockBaseAddress, lib, "CVPixelBufferUnlockBaseAddress")
	tryRegister(&_CVPixelFormatTypeCopyFourCharCodeString, lib, "CVPixelFormatTypeCopyFourCharCodeString")
	tryRegister(&_CVTransferFunctionGetIntegerCodePointForString, lib, "CVTransferFunctionGetIntegerCodePointForString")
	tryRegister(&_CVTransferFunctionGetStringForIntegerCodePoint, lib, "CVTransferFunctionGetStringForIntegerCodePoint")
	tryRegister(&_CVYCbCrMatrixGetIntegerCodePointForString, lib, "CVYCbCrMatrixGetIntegerCodePointForString")
	tryRegister(&_CVYCbCrMatrixGetStringForIntegerCodePoint, lib, "CVYCbCrMatrixGetStringForIntegerCodePoint")
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



// Returns a copy of an attachment from a Core Video buffer. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferCopyAttachment(_:_:_:)
func CVBufferCopyAttachment(buffer unsafe.Pointer, key unsafe.Pointer, attachmentMode unsafe.Pointer) unsafe.Pointer {
	return _CVBufferCopyAttachment(buffer, key, attachmentMode)
	}


// Returns a copy of all attachments from a Core Video buffer. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferCopyAttachments(_:_:)
func CVBufferCopyAttachments(buffer unsafe.Pointer, attachmentMode unsafe.Pointer) unsafe.Pointer {
	return _CVBufferCopyAttachments(buffer, attachmentMode)
	}


// Retrieves a specific attachment of a Core Video buffer. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferGetAttachment(_:_:_:)
func CVBufferGetAttachment(buffer unsafe.Pointer, key unsafe.Pointer, attachmentMode unsafe.Pointer) unsafe.Pointer {
	return _CVBufferGetAttachment(buffer, key, attachmentMode)
	}


// Retrieves all attachments of a Core Video buffer. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferGetAttachments(_:_:)
func CVBufferGetAttachments(buffer unsafe.Pointer, attachmentMode unsafe.Pointer) unsafe.Pointer {
	return _CVBufferGetAttachments(buffer, attachmentMode)
	}


// Returns a Boolean value that indicates whether a Core Video buffer contains a specified attachment. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferHasAttachment(_:_:)
func CVBufferHasAttachment(buffer unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CVBufferHasAttachment(buffer, key)
	}


// Copies all attachments that Core Video can propagate from one buffer to another. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferPropagateAttachments(_:_:)
func CVBufferPropagateAttachments(sourceBuffer unsafe.Pointer, destinationBuffer unsafe.Pointer) {
	_CVBufferPropagateAttachments(sourceBuffer, destinationBuffer)
	}


// Releases a Core Video buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferRelease
func CVBufferRelease(buffer unsafe.Pointer) {
	_CVBufferRelease(buffer)
	}


// Removes all attachments from a Core Video buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferRemoveAllAttachments(_:)
func CVBufferRemoveAllAttachments(buffer unsafe.Pointer) {
	_CVBufferRemoveAllAttachments(buffer)
	}


// Removes the attachment you specify from a Core Video buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferRemoveAttachment(_:_:)
func CVBufferRemoveAttachment(buffer unsafe.Pointer, key unsafe.Pointer) {
	_CVBufferRemoveAttachment(buffer, key)
	}


// Retains a Core Video buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferRetain
func CVBufferRetain(buffer unsafe.Pointer) unsafe.Pointer {
	return _CVBufferRetain(buffer)
	}


// Sets or adds an attachment to a Core Video buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferSetAttachment(_:_:_:_:)
func CVBufferSetAttachment(buffer unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer, attachmentMode unsafe.Pointer) {
	_CVBufferSetAttachment(buffer, key, value, attachmentMode)
	}


// Sets a dictionary of attachments on a Core Video buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferSetAttachments(_:_:_:)
func CVBufferSetAttachments(buffer unsafe.Pointer, theAttachments unsafe.Pointer, attachmentMode unsafe.Pointer) {
	_CVBufferSetAttachments(buffer, theAttachments, attachmentMode)
	}


// Returns the standard integer code point corresponding to the Core Video color primaries constant string that you specify. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVColorPrimariesGetIntegerCodePointForString(_:)
func CVColorPrimariesGetIntegerCodePointForString(colorPrimariesString unsafe.Pointer) unsafe.Pointer {
	return _CVColorPrimariesGetIntegerCodePointForString(colorPrimariesString)
	}


// Returns the Core Video color primaries string corresponding to the standard integer code point that you specify. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVColorPrimariesGetStringForIntegerCodePoint(_:)
func CVColorPrimariesGetStringForIntegerCodePoint(colorPrimariesCodePoint unsafe.Pointer) unsafe.Pointer {
	return _CVColorPrimariesGetStringForIntegerCodePoint(colorPrimariesCodePoint)
	}


// Creates a display link capable of being used with all active displays. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithActiveCGDisplays(_:)
func CVDisplayLinkCreateWithActiveCGDisplays(displayLinkOut unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkCreateWithActiveCGDisplays(displayLinkOut)
	}


// Creates a display link for a single display. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithCGDisplay(_:_:)
func CVDisplayLinkCreateWithCGDisplay(displayID unsafe.Pointer, displayLinkOut unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkCreateWithCGDisplay(displayID, displayLinkOut)
	}


// Creates a display link for an array of displays. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithCGDisplays(_:_:_:)
func CVDisplayLinkCreateWithCGDisplays(displayArray unsafe.Pointer, count unsafe.Pointer, displayLinkOut unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkCreateWithCGDisplays(displayArray, count, displayLinkOut)
	}


// Creates a display link from an OpenGL display mask. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithOpenGLDisplayMask(_:_:)
func CVDisplayLinkCreateWithOpenGLDisplayMask(mask unsafe.Pointer, displayLinkOut unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkCreateWithOpenGLDisplayMask(mask, displayLinkOut)
	}


// Retrieves the actual output refresh period of a display as measured by the system time. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetActualOutputVideoRefreshPeriod(_:)
func CVDisplayLinkGetActualOutputVideoRefreshPeriod(displayLink unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkGetActualOutputVideoRefreshPeriod(displayLink)
	}


// Gets the current display associated with a display link. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetCurrentCGDisplay(_:)
func CVDisplayLinkGetCurrentCGDisplay(displayLink unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkGetCurrentCGDisplay(displayLink)
	}


// Retrieves the current (“now”) time of a given display link. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetCurrentTime(_:_:)
func CVDisplayLinkGetCurrentTime(displayLink unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkGetCurrentTime(displayLink, outTime)
	}


// Retrieves the nominal refresh period of a display link. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetNominalOutputVideoRefreshPeriod(_:)
func CVDisplayLinkGetNominalOutputVideoRefreshPeriod(displayLink unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkGetNominalOutputVideoRefreshPeriod(displayLink)
	}


// Retrieves the nominal latency of a display link. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetOutputVideoLatency(_:)
func CVDisplayLinkGetOutputVideoLatency(displayLink unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkGetOutputVideoLatency(displayLink)
	}


// Obtains the Core Foundation ID for the display link data type. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetTypeID()
func CVDisplayLinkGetTypeID() unsafe.Pointer {
	return _CVDisplayLinkGetTypeID()
	}


// Indicates whether a given display link is running. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkIsRunning(_:)
func CVDisplayLinkIsRunning(displayLink unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkIsRunning(displayLink)
	}


// Releases a display link. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkRelease
func CVDisplayLinkRelease(displayLink unsafe.Pointer) {
	_CVDisplayLinkRelease(displayLink)
	}


// Retains a display link. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkRetain
func CVDisplayLinkRetain(displayLink unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkRetain(displayLink)
	}


// Sets the current display of a display link. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetCurrentCGDisplay(_:_:)
func CVDisplayLinkSetCurrentCGDisplay(displayLink unsafe.Pointer, displayID unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkSetCurrentCGDisplay(displayLink, displayID)
	}


// Selects the display link most optimal for the current renderer of an OpenGL context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(_:_:_:)
func CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(displayLink unsafe.Pointer, cglContext unsafe.Pointer, cglPixelFormat unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(displayLink, cglContext, cglPixelFormat)
	}


// Sets the renderer output callback function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetOutputCallback(_:_:_:)
func CVDisplayLinkSetOutputCallback(displayLink unsafe.Pointer, callback unsafe.Pointer, userInfo unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkSetOutputCallback(displayLink, callback, userInfo)
	}


// CVDisplayLinkSetOutputHandler is a CoreVideo function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetOutputHandler(_:_:)
func CVDisplayLinkSetOutputHandler(displayLink unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkSetOutputHandler(displayLink, handler)
	}


// Activates a display link. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkStart(_:)
func CVDisplayLinkStart(displayLink unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkStart(displayLink)
	}


// Stops a display link. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkStop(_:)
func CVDisplayLinkStop(displayLink unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkStop(displayLink)
	}


// Translates the time in the display link’s time base from one representation to another. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkTranslateTime(_:_:_:)
func CVDisplayLinkTranslateTime(displayLink unsafe.Pointer, inTime unsafe.Pointer, outTime unsafe.Pointer) unsafe.Pointer {
	return _CVDisplayLinkTranslateTime(displayLink, inTime, outTime)
	}


// Returns the current system time. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVGetCurrentHostTime()
func CVGetCurrentHostTime() unsafe.Pointer {
	return _CVGetCurrentHostTime()
	}


// Returns the frequency of updates to the system time. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVGetHostClockFrequency()
func CVGetHostClockFrequency() unsafe.Pointer {
	return _CVGetHostClockFrequency()
	}


// Returns the smallest possible increment in the system time. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVGetHostClockMinimumTimeDelta()
func CVGetHostClockMinimumTimeDelta() unsafe.Pointer {
	return _CVGetHostClockMinimumTimeDelta()
	}


// Attempts to create a Core Graphics color space from the image buffer’s attachments that you specify. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferCreateColorSpaceFromAttachments(_:)
func CVImageBufferCreateColorSpaceFromAttachments(attachments unsafe.Pointer) coregraphics.CGColorSpaceRef {
	return _CVImageBufferCreateColorSpaceFromAttachments(attachments)
	}


// Returns the source rectangle of a Core Video image buffer that represents the clean aperture of the buffer in encoded pixels. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetCleanRect(_:)
func CVImageBufferGetCleanRect(imageBuffer unsafe.Pointer) coregraphics.CGRect {
	return _CVImageBufferGetCleanRect(imageBuffer)
	}


// Returns the color space of a Core Video image buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetColorSpace(_:)
func CVImageBufferGetColorSpace(imageBuffer unsafe.Pointer) coregraphics.CGColorSpaceRef {
	return _CVImageBufferGetColorSpace(imageBuffer)
	}


// Returns the nominal output display size, in square pixels, of a Core Video image buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetDisplaySize(_:)
func CVImageBufferGetDisplaySize(imageBuffer unsafe.Pointer) coregraphics.CGSize {
	return _CVImageBufferGetDisplaySize(imageBuffer)
	}


// Returns the full encoded dimensions of a Core Video image buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetEncodedSize(_:)
func CVImageBufferGetEncodedSize(imageBuffer unsafe.Pointer) coregraphics.CGSize {
	return _CVImageBufferGetEncodedSize(imageBuffer)
	}


// Returns a Boolean value indicating whether the image is vertically flipped. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferIsFlipped(_:)
func CVImageBufferIsFlipped(imageBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVImageBufferIsFlipped(imageBuffer)
	}


// CVIsCompressedPixelFormatAvailable is a CoreVideo function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVIsCompressedPixelFormatAvailable(_:)
func CVIsCompressedPixelFormatAvailable(pixelFormatType unsafe.Pointer) unsafe.Pointer {
	return _CVIsCompressedPixelFormatAvailable(pixelFormatType)
	}


// CVMetalBufferCacheCreate is a CoreVideo function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheCreate(_:_:_:_:)
func CVMetalBufferCacheCreate(allocator unsafe.Pointer, cacheAttributes unsafe.Pointer, metalDevice unsafe.Pointer, cacheOut unsafe.Pointer) unsafe.Pointer {
	return _CVMetalBufferCacheCreate(allocator, cacheAttributes, metalDevice, cacheOut)
	}


// CVMetalBufferCacheCreateBufferFromImage is a CoreVideo function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheCreateBufferFromImage(_:_:_:_:)
func CVMetalBufferCacheCreateBufferFromImage(allocator unsafe.Pointer, bufferCache unsafe.Pointer, imageBuffer unsafe.Pointer, bufferOut unsafe.Pointer) unsafe.Pointer {
	return _CVMetalBufferCacheCreateBufferFromImage(allocator, bufferCache, imageBuffer, bufferOut)
	}


// CVMetalBufferCacheFlush is a CoreVideo function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheFlush(_:_:)
func CVMetalBufferCacheFlush(bufferCache unsafe.Pointer, options unsafe.Pointer) {
	_CVMetalBufferCacheFlush(bufferCache, options)
	}


// CVMetalBufferCacheGetTypeID is a CoreVideo function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheGetTypeID()
func CVMetalBufferCacheGetTypeID() unsafe.Pointer {
	return _CVMetalBufferCacheGetTypeID()
	}


// CVMetalBufferGetBuffer is a CoreVideo function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferGetBuffer(_:)
func CVMetalBufferGetBuffer(buffer unsafe.Pointer) unsafe.Pointer {
	return _CVMetalBufferGetBuffer(buffer)
	}


// CVMetalBufferGetTypeID is a CoreVideo function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferGetTypeID()
func CVMetalBufferGetTypeID() unsafe.Pointer {
	return _CVMetalBufferGetTypeID()
	}


// Creates a new texture cache. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheCreate(_:_:_:_:_:)
func CVMetalTextureCacheCreate(allocator unsafe.Pointer, cacheAttributes unsafe.Pointer, metalDevice unsafe.Pointer, textureAttributes unsafe.Pointer, cacheOut unsafe.Pointer) unsafe.Pointer {
	return _CVMetalTextureCacheCreate(allocator, cacheAttributes, metalDevice, textureAttributes, cacheOut)
	}


// Creates a Core Video Metal texture buffer from an existing image buffer. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheCreateTextureFromImage(_:_:_:_:_:_:_:_:_:)
func CVMetalTextureCacheCreateTextureFromImage(allocator unsafe.Pointer, textureCache unsafe.Pointer, sourceImage unsafe.Pointer, textureAttributes unsafe.Pointer, pixelFormat unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, planeIndex unsafe.Pointer, textureOut unsafe.Pointer) unsafe.Pointer {
	return _CVMetalTextureCacheCreateTextureFromImage(allocator, textureCache, sourceImage, textureAttributes, pixelFormat, width, height, planeIndex, textureOut)
	}


// Manually flushes the contents of the provided texture cache. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheFlush(_:_:)
func CVMetalTextureCacheFlush(textureCache unsafe.Pointer, options unsafe.Pointer) {
	_CVMetalTextureCacheFlush(textureCache, options)
	}


// Returns the Core Foundation type identifier for a Core Video Metal texture cache. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheGetTypeID()
func CVMetalTextureCacheGetTypeID() unsafe.Pointer {
	return _CVMetalTextureCacheGetTypeID()
	}


// Returns convenient normalized texture coordinates for the part of the image that should be displayed. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetCleanTexCoords(_:_:_:_:_:)
func CVMetalTextureGetCleanTexCoords(image unsafe.Pointer, lowerLeft unsafe.Pointer, lowerRight unsafe.Pointer, upperRight unsafe.Pointer, upperLeft unsafe.Pointer, p5 unsafe.Pointer) {
	_CVMetalTextureGetCleanTexCoords(image, lowerLeft, lowerRight, upperRight, upperLeft, p5)
	}


// Returns the Metal texture object for the image buffer. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetTexture(_:)
func CVMetalTextureGetTexture(image unsafe.Pointer) unsafe.Pointer {
	return _CVMetalTextureGetTexture(image)
	}


// Returns the Core Foundation type identifier for a CoreVideo Metal texture-based image buffer. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetTypeID()
func CVMetalTextureGetTypeID() unsafe.Pointer {
	return _CVMetalTextureGetTypeID()
	}


// Returns a Boolean value indicating whether the texture image is vertically flipped. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureIsFlipped(_:)
func CVMetalTextureIsFlipped(image unsafe.Pointer) unsafe.Pointer {
	return _CVMetalTextureIsFlipped(image)
	}


// Attaches an OpenGL context to a Core Video OpenGL buffer. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferAttach(_:_:_:_:_:)
func CVOpenGLBufferAttach(openGLBuffer unsafe.Pointer, cglContext unsafe.Pointer, face unsafe.Pointer, level unsafe.Pointer, screen unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLBufferAttach(openGLBuffer, cglContext, face, level, screen)
	}


// Creates a new Core Video OpenGL buffer that can be used for OpenGL rendering purposes [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferCreate(_:_:_:_:_:)
func CVOpenGLBufferCreate(allocator unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, attributes unsafe.Pointer, bufferOut unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLBufferCreate(allocator, width, height, attributes, bufferOut)
	}


// Obtains the attributes of a Core Video OpenGL buffer. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferGetAttributes(_:)
func CVOpenGLBufferGetAttributes(openGLBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLBufferGetAttributes(openGLBuffer)
	}


// Obtains the Core Foundation type ID for the OpenGL buffer type. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferGetTypeID()
func CVOpenGLBufferGetTypeID() unsafe.Pointer {
	return _CVOpenGLBufferGetTypeID()
	}


// Creates a new OpenGL buffer pool. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolCreate(_:_:_:_:)
func CVOpenGLBufferPoolCreate(allocator unsafe.Pointer, poolAttributes unsafe.Pointer, openGLBufferAttributes unsafe.Pointer, poolOut unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLBufferPoolCreate(allocator, poolAttributes, openGLBufferAttributes, poolOut)
	}


// Creates a new OpenGL buffer from an OpenGL buffer pool. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolCreateOpenGLBuffer(_:_:_:)
func CVOpenGLBufferPoolCreateOpenGLBuffer(allocator unsafe.Pointer, openGLBufferPool unsafe.Pointer, openGLBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLBufferPoolCreateOpenGLBuffer(allocator, openGLBufferPool, openGLBufferOut)
	}


// Returns the pool attributes dictionary for an Open GL buffer pool. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolGetAttributes(_:)
func CVOpenGLBufferPoolGetAttributes(pool unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLBufferPoolGetAttributes(pool)
	}


// Returns the attributes of OpenGL buffers that will be created from a buffer pool. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolGetOpenGLBufferAttributes(_:)
func CVOpenGLBufferPoolGetOpenGLBufferAttributes(pool unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLBufferPoolGetOpenGLBufferAttributes(pool)
	}


// Obtains the Core Foundation ID for the OpenGL buffer pool type. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolGetTypeID()
func CVOpenGLBufferPoolGetTypeID() unsafe.Pointer {
	return _CVOpenGLBufferPoolGetTypeID()
	}


// Releases an OpenGL buffer pool. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolRelease
func CVOpenGLBufferPoolRelease(openGLBufferPool unsafe.Pointer) {
	_CVOpenGLBufferPoolRelease(openGLBufferPool)
	}


// Retains an OpenGL buffer pool. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolRetain
func CVOpenGLBufferPoolRetain(openGLBufferPool unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLBufferPoolRetain(openGLBufferPool)
	}


// Releases a Core Video OpenGL buffer. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferRelease
func CVOpenGLBufferRelease(buffer unsafe.Pointer) {
	_CVOpenGLBufferRelease(buffer)
	}


// Retains a Core Video OpenGL buffer. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferRetain
func CVOpenGLBufferRetain(buffer unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLBufferRetain(buffer)
	}


// Creates a new Core Video texture cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureCacheCreate(_:_:_:_:_:)
func CVOpenGLESTextureCacheCreate(allocator unsafe.Pointer, cacheAttributes unsafe.Pointer, eaglContext unsafe.Pointer, textureAttributes unsafe.Pointer, cacheOut unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLESTextureCacheCreate(allocator, cacheAttributes, eaglContext, textureAttributes, cacheOut)
	}


// Creates a object from an existing . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureCacheCreateTextureFromImage(_:_:_:_:_:_:_:_:_:_:_:_:)
func CVOpenGLESTextureCacheCreateTextureFromImage(allocator unsafe.Pointer, textureCache unsafe.Pointer, sourceImage unsafe.Pointer, textureAttributes unsafe.Pointer, target unsafe.Pointer, internalFormat unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, format unsafe.Pointer, type_ unsafe.Pointer, planeIndex unsafe.Pointer, textureOut unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLESTextureCacheCreateTextureFromImage(allocator, textureCache, sourceImage, textureAttributes, target, internalFormat, width, height, format, type_, planeIndex, textureOut)
	}


// Performs internal housekeeping/recycling operations on a texture cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureCacheFlush(_:_:)
func CVOpenGLESTextureCacheFlush(textureCache unsafe.Pointer, options unsafe.Pointer) {
	_CVOpenGLESTextureCacheFlush(textureCache, options)
	}


// Returns the Core Foundation type identifier for a Core Video texture cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureCacheGetTypeID()
func CVOpenGLESTextureCacheGetTypeID() unsafe.Pointer {
	return _CVOpenGLESTextureCacheGetTypeID()
	}


// Returns convenient normalized texture coordinates for the part of the image that should be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureGetCleanTexCoords(_:_:_:_:_:)
func CVOpenGLESTextureGetCleanTexCoords(image unsafe.Pointer, lowerLeft unsafe.Pointer, lowerRight unsafe.Pointer, upperRight unsafe.Pointer, upperLeft unsafe.Pointer, p5 unsafe.Pointer) {
	_CVOpenGLESTextureGetCleanTexCoords(image, lowerLeft, lowerRight, upperRight, upperLeft, p5)
	}


// Returns the texture target name for a . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureGetName(_:)
func CVOpenGLESTextureGetName(image unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLESTextureGetName(image)
	}


// Returns the texture target for a . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureGetTarget(_:)
func CVOpenGLESTextureGetTarget(image unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLESTextureGetTarget(image)
	}


// Returns the Core Foundation type identifier for a Core Video texture-based image buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureGetTypeID()
func CVOpenGLESTextureGetTypeID() unsafe.Pointer {
	return _CVOpenGLESTextureGetTypeID()
	}


// Returns whether the image is flipped vertically or not. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureIsFlipped(_:)
func CVOpenGLESTextureIsFlipped(image unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLESTextureIsFlipped(image)
	}


// Creates a new texture cache. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheCreate(_:_:_:_:_:_:)
func CVOpenGLTextureCacheCreate(allocator unsafe.Pointer, cacheAttributes unsafe.Pointer, cglContext unsafe.Pointer, cglPixelFormat unsafe.Pointer, textureAttributes unsafe.Pointer, cacheOut unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLTextureCacheCreate(allocator, cacheAttributes, cglContext, cglPixelFormat, textureAttributes, cacheOut)
	}


// Creates a CVOpenGLTexture object from an existing . [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheCreateTextureFromImage(_:_:_:_:_:)
func CVOpenGLTextureCacheCreateTextureFromImage(allocator unsafe.Pointer, textureCache unsafe.Pointer, sourceImage unsafe.Pointer, attributes unsafe.Pointer, textureOut unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLTextureCacheCreateTextureFromImage(allocator, textureCache, sourceImage, attributes, textureOut)
	}


// Performs internal housekeeping/recycling operations on the cache. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheFlush(_:_:)
func CVOpenGLTextureCacheFlush(textureCache unsafe.Pointer, options unsafe.Pointer) {
	_CVOpenGLTextureCacheFlush(textureCache, options)
	}


// Returns the Core Foundation type identifier for a the texture cache. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheGetTypeID()
func CVOpenGLTextureCacheGetTypeID() unsafe.Pointer {
	return _CVOpenGLTextureCacheGetTypeID()
	}


// Releases a texture cache object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheRelease
func CVOpenGLTextureCacheRelease(textureCache unsafe.Pointer) {
	_CVOpenGLTextureCacheRelease(textureCache)
	}


// Retains a texture cache object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheRetain
func CVOpenGLTextureCacheRetain(textureCache unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLTextureCacheRetain(textureCache)
	}


// Returns the texture coordinates for the part of the image that should be displayed. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureGetCleanTexCoords(_:_:_:_:_:)
func CVOpenGLTextureGetCleanTexCoords(image unsafe.Pointer, lowerLeft unsafe.Pointer, lowerRight unsafe.Pointer, upperRight unsafe.Pointer, upperLeft unsafe.Pointer, p5 unsafe.Pointer) {
	_CVOpenGLTextureGetCleanTexCoords(image, lowerLeft, lowerRight, upperRight, upperLeft, p5)
	}


// Returns the texture target name of a CoreVideo OpenGL texture. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureGetName(_:)
func CVOpenGLTextureGetName(image unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLTextureGetName(image)
	}


// Returns the texture target (for example, ) of an OpenGL texture. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureGetTarget(_:)
func CVOpenGLTextureGetTarget(image unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLTextureGetTarget(image)
	}


// Obtains the Core Foundation ID for the Core Video OpenGL texture type. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureGetTypeID()
func CVOpenGLTextureGetTypeID() unsafe.Pointer {
	return _CVOpenGLTextureGetTypeID()
	}


// Determines whether an OpenGL texture is flipped vertically. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureIsFlipped(_:)
func CVOpenGLTextureIsFlipped(image unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLTextureIsFlipped(image)
	}


// Releases a Core Video OpenGL texture. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureRelease
func CVOpenGLTextureRelease(texture unsafe.Pointer) {
	_CVOpenGLTextureRelease(texture)
	}


// Retains a Core Video OpenGL texture. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureRetain
func CVOpenGLTextureRetain(texture unsafe.Pointer) unsafe.Pointer {
	return _CVOpenGLTextureRetain(texture)
	}


// CVPixelBufferCopyCreationAttributes is a CoreVideo function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCopyCreationAttributes(_:)
func CVPixelBufferCopyCreationAttributes(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferCopyCreationAttributes(pixelBuffer)
	}


// Creates a single pixel buffer for a given size and pixel format. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreate(_:_:_:_:_:_:)
func CVPixelBufferCreate(allocator unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, pixelFormatType unsafe.Pointer, pixelBufferAttributes unsafe.Pointer, pixelBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferCreate(allocator, width, height, pixelFormatType, pixelBufferAttributes, pixelBufferOut)
	}


// Resolves an array of objects describing various pixel buffer attributes into a single dictionary. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateResolvedAttributesDictionary(_:_:_:)
func CVPixelBufferCreateResolvedAttributesDictionary(allocator unsafe.Pointer, attributes unsafe.Pointer, resolvedDictionaryOut unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferCreateResolvedAttributesDictionary(allocator, attributes, resolvedDictionaryOut)
	}


// Creates a pixel buffer for a given size and pixel format containing data specified by a memory location. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithBytes(_:_:_:_:_:_:_:_:_:_:)
func CVPixelBufferCreateWithBytes(allocator unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, pixelFormatType unsafe.Pointer, baseAddress unsafe.Pointer, bytesPerRow unsafe.Pointer, releaseCallback unsafe.Pointer, releaseRefCon unsafe.Pointer, pixelBufferAttributes unsafe.Pointer, pixelBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferCreateWithBytes(allocator, width, height, pixelFormatType, baseAddress, bytesPerRow, releaseCallback, releaseRefCon, pixelBufferAttributes, pixelBufferOut)
	}


// Creates a single pixel buffer for the IO surface that you specify. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithIOSurface(_:_:_:_:)
func CVPixelBufferCreateWithIOSurface(allocator unsafe.Pointer, surface unsafe.Pointer, pixelBufferAttributes unsafe.Pointer, pixelBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferCreateWithIOSurface(allocator, surface, pixelBufferAttributes, pixelBufferOut)
	}


// Creates a single pixel buffer in planar format for a given size and pixel format containing data specified by a memory location. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithPlanarBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CVPixelBufferCreateWithPlanarBytes(allocator unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, pixelFormatType unsafe.Pointer, dataPtr unsafe.Pointer, dataSize unsafe.Pointer, numberOfPlanes unsafe.Pointer, planeBaseAddress unsafe.Pointer, planeWidth unsafe.Pointer, planeHeight unsafe.Pointer, planeBytesPerRow unsafe.Pointer, releaseCallback unsafe.Pointer, releaseRefCon unsafe.Pointer, pixelBufferAttributes unsafe.Pointer, pixelBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferCreateWithPlanarBytes(allocator, width, height, pixelFormatType, dataPtr, dataSize, numberOfPlanes, planeBaseAddress, planeWidth, planeHeight, planeBytesPerRow, releaseCallback, releaseRefCon, pixelBufferAttributes, pixelBufferOut)
	}


// Fills the extended pixels of the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferFillExtendedPixels(_:)
func CVPixelBufferFillExtendedPixels(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferFillExtendedPixels(pixelBuffer)
	}


// Returns the base address of the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBaseAddress(_:)
func CVPixelBufferGetBaseAddress(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetBaseAddress(pixelBuffer)
	}


// Returns the base address of the plane at the specified plane index. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBaseAddressOfPlane(_:_:)
func CVPixelBufferGetBaseAddressOfPlane(pixelBuffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetBaseAddressOfPlane(pixelBuffer, planeIndex)
	}


// Returns the number of bytes per row of the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBytesPerRow(_:)
func CVPixelBufferGetBytesPerRow(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetBytesPerRow(pixelBuffer)
	}


// Returns the number of bytes per row for a plane at the specified index in the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBytesPerRowOfPlane(_:_:)
func CVPixelBufferGetBytesPerRowOfPlane(pixelBuffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetBytesPerRowOfPlane(pixelBuffer, planeIndex)
	}


// Returns the data size for contiguous planes of the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetDataSize(_:)
func CVPixelBufferGetDataSize(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetDataSize(pixelBuffer)
	}


// Returns the amount of extended pixel padding in the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetExtendedPixels(_:_:_:_:_:)
func CVPixelBufferGetExtendedPixels(pixelBuffer unsafe.Pointer, extraColumnsOnLeft unsafe.Pointer, extraColumnsOnRight unsafe.Pointer, extraRowsOnTop unsafe.Pointer, extraRowsOnBottom unsafe.Pointer) {
	_CVPixelBufferGetExtendedPixels(pixelBuffer, extraColumnsOnLeft, extraColumnsOnRight, extraRowsOnTop, extraRowsOnBottom)
	}


// Returns the height of the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetHeight(_:)
func CVPixelBufferGetHeight(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetHeight(pixelBuffer)
	}


// Returns the height of the plane at planeIndex in the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetHeightOfPlane(_:_:)
func CVPixelBufferGetHeightOfPlane(pixelBuffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetHeightOfPlane(pixelBuffer, planeIndex)
	}


// Returns the IOSurface backing the pixel buffer, or if it is not backed by an IOSurface. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetIOSurface(_:)
func CVPixelBufferGetIOSurface(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetIOSurface(pixelBuffer)
	}


// Returns the pixel format type of the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetPixelFormatType(_:)
func CVPixelBufferGetPixelFormatType(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetPixelFormatType(pixelBuffer)
	}


// Returns number of planes of the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetPlaneCount(_:)
func CVPixelBufferGetPlaneCount(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetPlaneCount(pixelBuffer)
	}


// Returns the Core Foundation type identifier of the pixel buffer type. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetTypeID()
func CVPixelBufferGetTypeID() unsafe.Pointer {
	return _CVPixelBufferGetTypeID()
	}


// Returns the width of the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetWidth(_:)
func CVPixelBufferGetWidth(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetWidth(pixelBuffer)
	}


// Returns the width of the plane at a given index in the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetWidthOfPlane(_:_:)
func CVPixelBufferGetWidthOfPlane(pixelBuffer unsafe.Pointer, planeIndex unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferGetWidthOfPlane(pixelBuffer, planeIndex)
	}


// CVPixelBufferIsCompatibleWithAttributes is a CoreVideo function. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferIsCompatibleWithAttributes(_:_:)
func CVPixelBufferIsCompatibleWithAttributes(pixelBuffer unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferIsCompatibleWithAttributes(pixelBuffer, attributes)
	}


// Determines whether the pixel buffer is planar. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferIsPlanar(_:)
func CVPixelBufferIsPlanar(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferIsPlanar(pixelBuffer)
	}


// Locks the base address of the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferLockBaseAddress(_:_:)
func CVPixelBufferLockBaseAddress(pixelBuffer unsafe.Pointer, lockFlags unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferLockBaseAddress(pixelBuffer, lockFlags)
	}


// Creates a pixel buffer pool using the allocator and attributes that you specify. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreate(_:_:_:_:)
func CVPixelBufferPoolCreate(allocator unsafe.Pointer, poolAttributes unsafe.Pointer, pixelBufferAttributes unsafe.Pointer, poolOut unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferPoolCreate(allocator, poolAttributes, pixelBufferAttributes, poolOut)
	}


// Creates a pixel buffer from a pixel buffer pool, using the allocator that you specify. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreatePixelBuffer(_:_:_:)
func CVPixelBufferPoolCreatePixelBuffer(allocator unsafe.Pointer, pixelBufferPool unsafe.Pointer, pixelBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferPoolCreatePixelBuffer(allocator, pixelBufferPool, pixelBufferOut)
	}


// Creates a new pixel buffer with auxiliary attributes from the pool. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(_:_:_:_:)
func CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(allocator unsafe.Pointer, pixelBufferPool unsafe.Pointer, auxAttributes unsafe.Pointer, pixelBufferOut unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(allocator, pixelBufferPool, auxAttributes, pixelBufferOut)
	}


// Frees pixel buffers from the pool based on the options that you specify. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolFlush(_:_:)
func CVPixelBufferPoolFlush(pool unsafe.Pointer, options unsafe.Pointer) {
	_CVPixelBufferPoolFlush(pool, options)
	}


// The pool attributes dictionary for a pixel buffer pool. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetAttributes(_:)
func CVPixelBufferPoolGetAttributes(pool unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferPoolGetAttributes(pool)
	}


// The attributes of pixel buffers which the system creates using the pool you specify. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetPixelBufferAttributes(_:)
func CVPixelBufferPoolGetPixelBufferAttributes(pool unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferPoolGetPixelBufferAttributes(pool)
	}


// Returns the Core Foundation type identifier of the pixel buffer pool type. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetTypeID()
func CVPixelBufferPoolGetTypeID() unsafe.Pointer {
	return _CVPixelBufferPoolGetTypeID()
	}


// Releases a pixel buffer pool. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolRelease
func CVPixelBufferPoolRelease(pixelBufferPool unsafe.Pointer) {
	_CVPixelBufferPoolRelease(pixelBufferPool)
	}


// Retains the pixel buffer pool that you specify. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolRetain
func CVPixelBufferPoolRetain(pixelBufferPool unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferPoolRetain(pixelBufferPool)
	}


// Releases a pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferRelease
func CVPixelBufferRelease(texture unsafe.Pointer) {
	_CVPixelBufferRelease(texture)
	}


// Retains a pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferRetain
func CVPixelBufferRetain(texture unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferRetain(texture)
	}


// Unlocks the base address of the pixel buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferUnlockBaseAddress(_:_:)
func CVPixelBufferUnlockBaseAddress(pixelBuffer unsafe.Pointer, unlockFlags unsafe.Pointer) unsafe.Pointer {
	return _CVPixelBufferUnlockBaseAddress(pixelBuffer, unlockFlags)
	}


// CVPixelFormatTypeCopyFourCharCodeString is a CoreVideo function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatTypeCopyFourCharCodeString(_:)
func CVPixelFormatTypeCopyFourCharCodeString(pixelFormat unsafe.Pointer) unsafe.Pointer {
	return _CVPixelFormatTypeCopyFourCharCodeString(pixelFormat)
	}


// Returns the standard integer code point corresponding to the Core Video transfer function string that you specify. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTransferFunctionGetIntegerCodePointForString(_:)
func CVTransferFunctionGetIntegerCodePointForString(transferFunctionString unsafe.Pointer) unsafe.Pointer {
	return _CVTransferFunctionGetIntegerCodePointForString(transferFunctionString)
	}


// Returns the Core Video transfer function string corresponding to the standard integer code point that you specify. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTransferFunctionGetStringForIntegerCodePoint(_:)
func CVTransferFunctionGetStringForIntegerCodePoint(transferFunctionCodePoint unsafe.Pointer) unsafe.Pointer {
	return _CVTransferFunctionGetStringForIntegerCodePoint(transferFunctionCodePoint)
	}


// Returns the standard integer code point corresponding to the Core Video YCbCr matrix string that you specify. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVYCbCrMatrixGetIntegerCodePointForString(_:)
func CVYCbCrMatrixGetIntegerCodePointForString(yCbCrMatrixString unsafe.Pointer) unsafe.Pointer {
	return _CVYCbCrMatrixGetIntegerCodePointForString(yCbCrMatrixString)
	}


// Returns the Core Video YCbCr matrix string corresponding to the standard integer code point that you specify. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVYCbCrMatrixGetStringForIntegerCodePoint(_:)
func CVYCbCrMatrixGetStringForIntegerCodePoint(yCbCrMatrixCodePoint unsafe.Pointer) unsafe.Pointer {
	return _CVYCbCrMatrixGetStringForIntegerCodePoint(yCbCrMatrixCodePoint)
	}




