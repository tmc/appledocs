// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo


import (
	"unsafe"

	"github.com/ebitengine/purego"
	corefoundation "github.com/tmc/appledocs/generated/corefoundation"
)


// CoreVideo Functions (121 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CVMetalBufferCacheCreate func(AllocatorRef, DictionaryRef, unsafe.Pointer, unsafe.Pointer) Return
	_CVMetalBufferCacheCreateBufferFromImage func(AllocatorRef, MetalBufferCacheRef, ImageBufferRef, unsafe.Pointer) Return
	_CVMetalBufferCacheFlush func(MetalBufferCacheRef, OptionFlags)
	_CVMetalBufferCacheGetTypeID func() TypeID
	_CVMetalBufferGetBuffer func(MetalBufferRef) unsafe.Pointer
	_CVMetalBufferGetTypeID func() TypeID
	_CVPixelBufferIsCompatibleWithAttributes func(PixelBufferRef, DictionaryRef) unsafe.Pointer
	_CVPixelFormatTypeCopyFourCharCodeString func(unsafe.Pointer) StringRef
	_CVBufferCopyAttachment func(BufferRef, StringRef, unsafe.Pointer) TypeRef
	_CVBufferCopyAttachments func(BufferRef, AttachmentMode) DictionaryRef
	_CVBufferGetAttachment func(BufferRef, StringRef, unsafe.Pointer) TypeRef
	_CVBufferGetAttachments func(BufferRef, AttachmentMode) DictionaryRef
	_CVBufferHasAttachment func(BufferRef, StringRef) unsafe.Pointer
	_CVBufferPropagateAttachments func(BufferRef, BufferRef)
	_CVBufferRemoveAllAttachments func(BufferRef)
	_CVBufferRemoveAttachment func(BufferRef, StringRef)
	_CVBufferSetAttachment func(BufferRef, StringRef, TypeRef, AttachmentMode)
	_CVBufferSetAttachments func(BufferRef, DictionaryRef, AttachmentMode)
	_CVColorPrimariesGetIntegerCodePointForString func(StringRef) int
	_CVColorPrimariesGetStringForIntegerCodePoint func(int) StringRef
	_CVDisplayLinkCreateWithActiveCGDisplays func(unsafe.Pointer) Return
	_CVDisplayLinkCreateWithCGDisplay func(DirectDisplayID, unsafe.Pointer) Return
	_CVDisplayLinkCreateWithCGDisplays func(unsafe.Pointer, Index, unsafe.Pointer) Return
	_CVDisplayLinkCreateWithOpenGLDisplayMask func(OpenGLDisplayMask, unsafe.Pointer) Return
	_CVDisplayLinkGetActualOutputVideoRefreshPeriod func(DisplayLinkRef) float64
	_CVDisplayLinkGetCurrentCGDisplay func(DisplayLinkRef) DirectDisplayID
	_CVDisplayLinkGetCurrentTime func(DisplayLinkRef, unsafe.Pointer) Return
	_CVDisplayLinkGetNominalOutputVideoRefreshPeriod func(DisplayLinkRef) Time
	_CVDisplayLinkGetOutputVideoLatency func(DisplayLinkRef) Time
	_CVDisplayLinkGetTypeID func() TypeID
	_CVDisplayLinkIsRunning func(DisplayLinkRef) unsafe.Pointer
	_CVDisplayLinkSetCurrentCGDisplay func(DisplayLinkRef, DirectDisplayID) Return
	_CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext func(DisplayLinkRef, LContextObj, LPixelFormatObj) Return
	_CVDisplayLinkSetOutputCallback func(DisplayLinkRef, DisplayLinkOutputCallback, unsafe.Pointer) Return
	_CVDisplayLinkSetOutputHandler func(DisplayLinkRef, DisplayLinkOutputHandler) Return
	_CVDisplayLinkStart func(DisplayLinkRef) Return
	_CVDisplayLinkStop func(DisplayLinkRef) Return
	_CVDisplayLinkTranslateTime func(DisplayLinkRef, unsafe.Pointer, unsafe.Pointer) Return
	_CVGetCurrentHostTime func() uint64
	_CVGetHostClockFrequency func() float64
	_CVGetHostClockMinimumTimeDelta func() uint32
	_CVImageBufferCreateColorSpaceFromAttachments func(DictionaryRef) ColorSpaceRef
	_CVImageBufferGetCleanRect func(ImageBufferRef) corefoundation.CGRect
	_CVImageBufferGetColorSpace func(ImageBufferRef) ColorSpaceRef
	_CVImageBufferGetDisplaySize func(ImageBufferRef) corefoundation.CGSize
	_CVImageBufferGetEncodedSize func(ImageBufferRef) corefoundation.CGSize
	_CVImageBufferIsFlipped func(ImageBufferRef) unsafe.Pointer
	_CVIsCompressedPixelFormatAvailable func(unsafe.Pointer) unsafe.Pointer
	_CVMetalTextureCacheCreate func(AllocatorRef, DictionaryRef, unsafe.Pointer, DictionaryRef, unsafe.Pointer) Return
	_CVMetalTextureCacheCreateTextureFromImage func(AllocatorRef, MetalTextureCacheRef, ImageBufferRef, DictionaryRef, PixelFormat, uintptr, uintptr, uintptr, unsafe.Pointer) Return
	_CVMetalTextureCacheFlush func(MetalTextureCacheRef, OptionFlags)
	_CVMetalTextureCacheGetTypeID func() TypeID
	_CVMetalTextureGetCleanTexCoords func(MetalTextureRef, float32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CVMetalTextureGetTexture func(MetalTextureRef) unsafe.Pointer
	_CVMetalTextureGetTypeID func() TypeID
	_CVMetalTextureIsFlipped func(MetalTextureRef) unsafe.Pointer
	_CVOpenGLBufferAttach func(OpenGLBufferRef, LContextObj, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) Return
	_CVOpenGLBufferCreate func(AllocatorRef, uintptr, uintptr, DictionaryRef, unsafe.Pointer) Return
	_CVOpenGLBufferGetAttributes func(OpenGLBufferRef) DictionaryRef
	_CVOpenGLBufferGetTypeID func() TypeID
	_CVOpenGLBufferPoolCreate func(AllocatorRef, DictionaryRef, DictionaryRef, unsafe.Pointer) Return
	_CVOpenGLBufferPoolCreateOpenGLBuffer func(AllocatorRef, OpenGLBufferPoolRef, unsafe.Pointer) Return
	_CVOpenGLBufferPoolGetAttributes func(OpenGLBufferPoolRef) DictionaryRef
	_CVOpenGLBufferPoolGetOpenGLBufferAttributes func(OpenGLBufferPoolRef) DictionaryRef
	_CVOpenGLBufferPoolGetTypeID func() TypeID
	_CVOpenGLESTextureCacheCreate func(AllocatorRef, DictionaryRef, EAGLContext, DictionaryRef, unsafe.Pointer) Return
	_CVOpenGLESTextureCacheCreateTextureFromImage func(AllocatorRef, OpenGLESTextureCacheRef, ImageBufferRef, DictionaryRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer) Return
	_CVOpenGLESTextureCacheFlush func(OpenGLESTextureCacheRef, OptionFlags)
	_CVOpenGLESTextureCacheGetTypeID func() TypeID
	_CVOpenGLESTextureGetCleanTexCoords func(OpenGLESTextureRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CVOpenGLESTextureGetName func(OpenGLESTextureRef) unsafe.Pointer
	_CVOpenGLESTextureGetTarget func(OpenGLESTextureRef) unsafe.Pointer
	_CVOpenGLESTextureGetTypeID func() TypeID
	_CVOpenGLESTextureIsFlipped func(OpenGLESTextureRef) unsafe.Pointer
	_CVOpenGLTextureCacheCreate func(AllocatorRef, DictionaryRef, LContextObj, LPixelFormatObj, DictionaryRef, unsafe.Pointer) Return
	_CVOpenGLTextureCacheCreateTextureFromImage func(AllocatorRef, OpenGLTextureCacheRef, ImageBufferRef, DictionaryRef, unsafe.Pointer) Return
	_CVOpenGLTextureCacheFlush func(OpenGLTextureCacheRef, OptionFlags)
	_CVOpenGLTextureCacheGetTypeID func() TypeID
	_CVOpenGLTextureGetCleanTexCoords func(OpenGLTextureRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CVOpenGLTextureGetName func(OpenGLTextureRef) unsafe.Pointer
	_CVOpenGLTextureGetTarget func(OpenGLTextureRef) unsafe.Pointer
	_CVOpenGLTextureGetTypeID func() TypeID
	_CVOpenGLTextureIsFlipped func(OpenGLTextureRef) unsafe.Pointer
	_CVPixelBufferCopyCreationAttributes func(PixelBufferRef) DictionaryRef
	_CVPixelBufferCreate func(AllocatorRef, uintptr, uintptr, unsafe.Pointer, DictionaryRef, unsafe.Pointer) Return
	_CVPixelBufferCreateResolvedAttributesDictionary func(AllocatorRef, ArrayRef, unsafe.Pointer) Return
	_CVPixelBufferCreateWithBytes func(AllocatorRef, uintptr, uintptr, unsafe.Pointer, unsafe.Pointer, uintptr, PixelBufferReleaseBytesCallback, unsafe.Pointer, DictionaryRef, unsafe.Pointer) Return
	_CVPixelBufferCreateWithIOSurface func(AllocatorRef, SurfaceRef, DictionaryRef, unsafe.Pointer) Return
	_CVPixelBufferCreateWithPlanarBytes func(AllocatorRef, uintptr, uintptr, unsafe.Pointer, unsafe.Pointer, uintptr, uintptr, unsafe.Pointer, uintptr, uintptr, uintptr, PixelBufferReleasePlanarBytesCallback, unsafe.Pointer, DictionaryRef, unsafe.Pointer) Return
	_CVPixelBufferFillExtendedPixels func(PixelBufferRef) Return
	_CVPixelBufferGetBaseAddress func(PixelBufferRef) unsafe.Pointer
	_CVPixelBufferGetBaseAddressOfPlane func(PixelBufferRef, uintptr) unsafe.Pointer
	_CVPixelBufferGetBytesPerRow func(PixelBufferRef) uintptr
	_CVPixelBufferGetBytesPerRowOfPlane func(PixelBufferRef, uintptr) uintptr
	_CVPixelBufferGetDataSize func(PixelBufferRef) uintptr
	_CVPixelBufferGetExtendedPixels func(PixelBufferRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CVPixelBufferGetHeight func(PixelBufferRef) uintptr
	_CVPixelBufferGetHeightOfPlane func(PixelBufferRef, uintptr) uintptr
	_CVPixelBufferGetIOSurface func(PixelBufferRef) SurfaceRef
	_CVPixelBufferGetPixelFormatType func(PixelBufferRef) unsafe.Pointer
	_CVPixelBufferGetPlaneCount func(PixelBufferRef) uintptr
	_CVPixelBufferGetTypeID func() TypeID
	_CVPixelBufferGetWidth func(PixelBufferRef) uintptr
	_CVPixelBufferGetWidthOfPlane func(PixelBufferRef, uintptr) uintptr
	_CVPixelBufferIsPlanar func(PixelBufferRef) unsafe.Pointer
	_CVPixelBufferLockBaseAddress func(PixelBufferRef, PixelBufferLockFlags) Return
	_CVPixelBufferPoolCreate func(AllocatorRef, DictionaryRef, DictionaryRef, unsafe.Pointer) Return
	_CVPixelBufferPoolCreatePixelBuffer func(AllocatorRef, PixelBufferPoolRef, unsafe.Pointer) Return
	_CVPixelBufferPoolCreatePixelBufferWithAuxAttributes func(AllocatorRef, PixelBufferPoolRef, DictionaryRef, unsafe.Pointer) Return
	_CVPixelBufferPoolFlush func(PixelBufferPoolRef, PixelBufferPoolFlushFlags)
	_CVPixelBufferPoolGetAttributes func(PixelBufferPoolRef) DictionaryRef
	_CVPixelBufferPoolGetPixelBufferAttributes func(PixelBufferPoolRef) DictionaryRef
	_CVPixelBufferPoolGetTypeID func() TypeID
	_CVPixelBufferUnlockBaseAddress func(PixelBufferRef, PixelBufferLockFlags) Return
	_CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes func(AllocatorRef) ArrayRef
	_CVPixelFormatDescriptionCreateWithPixelFormatType func(AllocatorRef, unsafe.Pointer) DictionaryRef
	_CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType func(DictionaryRef, unsafe.Pointer)
	_CVTransferFunctionGetIntegerCodePointForString func(StringRef) int
	_CVTransferFunctionGetStringForIntegerCodePoint func(int) StringRef
	_CVYCbCrMatrixGetIntegerCodePointForString func(StringRef) int
	_CVYCbCrMatrixGetStringForIntegerCodePoint func(int) StringRef
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CVMetalBufferCacheCreate, lib, "CVMetalBufferCacheCreate")
	tryRegister(&_CVMetalBufferCacheCreateBufferFromImage, lib, "CVMetalBufferCacheCreateBufferFromImage")
	tryRegister(&_CVMetalBufferCacheFlush, lib, "CVMetalBufferCacheFlush")
	tryRegister(&_CVMetalBufferCacheGetTypeID, lib, "CVMetalBufferCacheGetTypeID")
	tryRegister(&_CVMetalBufferGetBuffer, lib, "CVMetalBufferGetBuffer")
	tryRegister(&_CVMetalBufferGetTypeID, lib, "CVMetalBufferGetTypeID")
	tryRegister(&_CVPixelBufferIsCompatibleWithAttributes, lib, "CVPixelBufferIsCompatibleWithAttributes")
	tryRegister(&_CVPixelFormatTypeCopyFourCharCodeString, lib, "CVPixelFormatTypeCopyFourCharCodeString")
	tryRegister(&_CVBufferCopyAttachment, lib, "CVBufferCopyAttachment")
	tryRegister(&_CVBufferCopyAttachments, lib, "CVBufferCopyAttachments")
	tryRegister(&_CVBufferGetAttachment, lib, "CVBufferGetAttachment")
	tryRegister(&_CVBufferGetAttachments, lib, "CVBufferGetAttachments")
	tryRegister(&_CVBufferHasAttachment, lib, "CVBufferHasAttachment")
	tryRegister(&_CVBufferPropagateAttachments, lib, "CVBufferPropagateAttachments")
	tryRegister(&_CVBufferRemoveAllAttachments, lib, "CVBufferRemoveAllAttachments")
	tryRegister(&_CVBufferRemoveAttachment, lib, "CVBufferRemoveAttachment")
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
	tryRegister(&_CVOpenGLTextureGetCleanTexCoords, lib, "CVOpenGLTextureGetCleanTexCoords")
	tryRegister(&_CVOpenGLTextureGetName, lib, "CVOpenGLTextureGetName")
	tryRegister(&_CVOpenGLTextureGetTarget, lib, "CVOpenGLTextureGetTarget")
	tryRegister(&_CVOpenGLTextureGetTypeID, lib, "CVOpenGLTextureGetTypeID")
	tryRegister(&_CVOpenGLTextureIsFlipped, lib, "CVOpenGLTextureIsFlipped")
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
	tryRegister(&_CVPixelBufferIsPlanar, lib, "CVPixelBufferIsPlanar")
	tryRegister(&_CVPixelBufferLockBaseAddress, lib, "CVPixelBufferLockBaseAddress")
	tryRegister(&_CVPixelBufferPoolCreate, lib, "CVPixelBufferPoolCreate")
	tryRegister(&_CVPixelBufferPoolCreatePixelBuffer, lib, "CVPixelBufferPoolCreatePixelBuffer")
	tryRegister(&_CVPixelBufferPoolCreatePixelBufferWithAuxAttributes, lib, "CVPixelBufferPoolCreatePixelBufferWithAuxAttributes")
	tryRegister(&_CVPixelBufferPoolFlush, lib, "CVPixelBufferPoolFlush")
	tryRegister(&_CVPixelBufferPoolGetAttributes, lib, "CVPixelBufferPoolGetAttributes")
	tryRegister(&_CVPixelBufferPoolGetPixelBufferAttributes, lib, "CVPixelBufferPoolGetPixelBufferAttributes")
	tryRegister(&_CVPixelBufferPoolGetTypeID, lib, "CVPixelBufferPoolGetTypeID")
	tryRegister(&_CVPixelBufferUnlockBaseAddress, lib, "CVPixelBufferUnlockBaseAddress")
	tryRegister(&_CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes, lib, "CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes")
	tryRegister(&_CVPixelFormatDescriptionCreateWithPixelFormatType, lib, "CVPixelFormatDescriptionCreateWithPixelFormatType")
	tryRegister(&_CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType, lib, "CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType")
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



// CVMetalBufferCacheCreate is a CoreVideo function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheCreate(_:_:_:_:)
func CVMetalBufferCacheCreate(allocator AllocatorRef, cacheAttributes DictionaryRef, metalDevice unsafe.Pointer, cacheOut unsafe.Pointer) Return {
	return _CVMetalBufferCacheCreate(allocator, cacheAttributes, metalDevice, cacheOut)
}

// CVMetalBufferCacheCreateBufferFromImage is a CoreVideo function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheCreateBufferFromImage(_:_:_:_:)
func CVMetalBufferCacheCreateBufferFromImage(allocator AllocatorRef, bufferCache MetalBufferCacheRef, imageBuffer ImageBufferRef, bufferOut unsafe.Pointer) Return {
	return _CVMetalBufferCacheCreateBufferFromImage(allocator, bufferCache, imageBuffer, bufferOut)
}

// CVMetalBufferCacheFlush is a CoreVideo function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheFlush(_:_:)
func CVMetalBufferCacheFlush(bufferCache MetalBufferCacheRef, options OptionFlags) {
	_CVMetalBufferCacheFlush(bufferCache, options)
}

// CVMetalBufferCacheGetTypeID is a CoreVideo function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheGetTypeID()
func CVMetalBufferCacheGetTypeID() TypeID {
	return _CVMetalBufferCacheGetTypeID()
}

// CVMetalBufferGetBuffer is a CoreVideo function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferGetBuffer(_:)
func CVMetalBufferGetBuffer(buffer MetalBufferRef) unsafe.Pointer {
	return _CVMetalBufferGetBuffer(buffer)
}

// CVMetalBufferGetTypeID is a CoreVideo function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferGetTypeID()
func CVMetalBufferGetTypeID() TypeID {
	return _CVMetalBufferGetTypeID()
}

// CVPixelBufferIsCompatibleWithAttributes is a CoreVideo function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferIsCompatibleWithAttributes(_:_:)
func CVPixelBufferIsCompatibleWithAttributes(pixelBuffer PixelBufferRef, attributes DictionaryRef) unsafe.Pointer {
	return _CVPixelBufferIsCompatibleWithAttributes(pixelBuffer, attributes)
}

// CVPixelFormatTypeCopyFourCharCodeString is a CoreVideo function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatTypeCopyFourCharCodeString(_:)
func CVPixelFormatTypeCopyFourCharCodeString(pixelFormat unsafe.Pointer) StringRef {
	return _CVPixelFormatTypeCopyFourCharCodeString(pixelFormat)
}

// Returns a copy of an attachment from a Core Video buffer.
//
// Added in macOS 12.0.
// Returns a copy of an attachment from a Core Video buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferCopyAttachment(_:_:_:)
func CVBufferCopyAttachment(buffer BufferRef, key StringRef, attachmentMode unsafe.Pointer) TypeRef {
	return _CVBufferCopyAttachment(buffer, key, attachmentMode)
}

// Returns a copy of all attachments from a Core Video buffer.
//
// Added in macOS 12.0.
// Returns a copy of all attachments from a Core Video buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferCopyAttachments(_:_:)
func CVBufferCopyAttachments(buffer BufferRef, attachmentMode AttachmentMode) DictionaryRef {
	return _CVBufferCopyAttachments(buffer, attachmentMode)
}

// Retrieves a specific attachment of a Core Video buffer.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.4.
// Retrieves a specific attachment of a Core Video buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferGetAttachment(_:_:_:)
func CVBufferGetAttachment(buffer BufferRef, key StringRef, attachmentMode unsafe.Pointer) TypeRef {
	return _CVBufferGetAttachment(buffer, key, attachmentMode)
}

// Retrieves all attachments of a Core Video buffer.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.4.
// Retrieves all attachments of a Core Video buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferGetAttachments(_:_:)
func CVBufferGetAttachments(buffer BufferRef, attachmentMode AttachmentMode) DictionaryRef {
	return _CVBufferGetAttachments(buffer, attachmentMode)
}

// Returns a Boolean value that indicates whether a Core Video buffer contains a specified attachment.
//
// Added in macOS 12.0.
// Returns a Boolean value that indicates whether a Core Video buffer contains a specified attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferHasAttachment(_:_:)
func CVBufferHasAttachment(buffer BufferRef, key StringRef) unsafe.Pointer {
	return _CVBufferHasAttachment(buffer, key)
}

// Copies all attachments that Core Video can propagate from one buffer to another.
//
// Added in macOS 10.4.
// Copies all attachments that Core Video can propagate from one buffer to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferPropagateAttachments(_:_:)
func CVBufferPropagateAttachments(sourceBuffer BufferRef, destinationBuffer BufferRef) {
	_CVBufferPropagateAttachments(sourceBuffer, destinationBuffer)
}

// Removes all attachments from a Core Video buffer.
//
// Added in macOS 10.4.
// Removes all attachments from a Core Video buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferRemoveAllAttachments(_:)
func CVBufferRemoveAllAttachments(buffer BufferRef) {
	_CVBufferRemoveAllAttachments(buffer)
}

// Removes the attachment you specify from a Core Video buffer.
//
// Added in macOS 10.4.
// Removes the attachment you specify from a Core Video buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferRemoveAttachment(_:_:)
func CVBufferRemoveAttachment(buffer BufferRef, key StringRef) {
	_CVBufferRemoveAttachment(buffer, key)
}

// Sets or adds an attachment to a Core Video buffer.
//
// Added in macOS 10.4.
// Sets or adds an attachment to a Core Video buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferSetAttachment(_:_:_:_:)
func CVBufferSetAttachment(buffer BufferRef, key StringRef, value TypeRef, attachmentMode AttachmentMode) {
	_CVBufferSetAttachment(buffer, key, value, attachmentMode)
}

// Sets a dictionary of attachments on a Core Video buffer.
//
// Added in macOS 10.4.
// Sets a dictionary of attachments on a Core Video buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBufferSetAttachments(_:_:_:)
func CVBufferSetAttachments(buffer BufferRef, theAttachments DictionaryRef, attachmentMode AttachmentMode) {
	_CVBufferSetAttachments(buffer, theAttachments, attachmentMode)
}

// Returns the standard integer code point corresponding to the Core Video color primaries constant string that you specify.
//
// Added in macOS 10.13.
// Returns the standard integer code point corresponding to the Core Video color primaries constant string that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVColorPrimariesGetIntegerCodePointForString(_:)
func CVColorPrimariesGetIntegerCodePointForString(colorPrimariesString StringRef) int {
	return _CVColorPrimariesGetIntegerCodePointForString(colorPrimariesString)
}

// Returns the Core Video color primaries string corresponding to the standard integer code point that you specify.
//
// Added in macOS 10.13.
// Returns the Core Video color primaries string corresponding to the standard integer code point that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVColorPrimariesGetStringForIntegerCodePoint(_:)
func CVColorPrimariesGetStringForIntegerCodePoint(colorPrimariesCodePoint int) StringRef {
	return _CVColorPrimariesGetStringForIntegerCodePoint(colorPrimariesCodePoint)
}

// Creates a display link capable of being used with all active displays.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Creates a display link capable of being used with all active displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithActiveCGDisplays(_:)
func CVDisplayLinkCreateWithActiveCGDisplays(displayLinkOut unsafe.Pointer) Return {
	return _CVDisplayLinkCreateWithActiveCGDisplays(displayLinkOut)
}

// Creates a display link for a single display.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Creates a display link for a single display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithCGDisplay(_:_:)
func CVDisplayLinkCreateWithCGDisplay(displayID DirectDisplayID, displayLinkOut unsafe.Pointer) Return {
	return _CVDisplayLinkCreateWithCGDisplay(displayID, displayLinkOut)
}

// Creates a display link for an array of displays.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Creates a display link for an array of displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithCGDisplays(_:_:_:)
func CVDisplayLinkCreateWithCGDisplays(displayArray unsafe.Pointer, count Index, displayLinkOut unsafe.Pointer) Return {
	return _CVDisplayLinkCreateWithCGDisplays(displayArray, count, displayLinkOut)
}

// Creates a display link from an OpenGL display mask.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Creates a display link from an OpenGL display mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithOpenGLDisplayMask(_:_:)
func CVDisplayLinkCreateWithOpenGLDisplayMask(mask OpenGLDisplayMask, displayLinkOut unsafe.Pointer) Return {
	return _CVDisplayLinkCreateWithOpenGLDisplayMask(mask, displayLinkOut)
}

// Retrieves the actual output refresh period of a display as measured by the system time.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Retrieves the actual output refresh period of a display as measured by the system time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetActualOutputVideoRefreshPeriod(_:)
func CVDisplayLinkGetActualOutputVideoRefreshPeriod(displayLink DisplayLinkRef) float64 {
	return _CVDisplayLinkGetActualOutputVideoRefreshPeriod(displayLink)
}

// Gets the current display associated with a display link.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Gets the current display associated with a display link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetCurrentCGDisplay(_:)
func CVDisplayLinkGetCurrentCGDisplay(displayLink DisplayLinkRef) DirectDisplayID {
	return _CVDisplayLinkGetCurrentCGDisplay(displayLink)
}

// Retrieves the current (“now”) time of a given display link.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Retrieves the current (“now”) time of a given display link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetCurrentTime(_:_:)
func CVDisplayLinkGetCurrentTime(displayLink DisplayLinkRef, outTime unsafe.Pointer) Return {
	return _CVDisplayLinkGetCurrentTime(displayLink, outTime)
}

// Retrieves the nominal refresh period of a display link.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Retrieves the nominal refresh period of a display link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetNominalOutputVideoRefreshPeriod(_:)
func CVDisplayLinkGetNominalOutputVideoRefreshPeriod(displayLink DisplayLinkRef) Time {
	return _CVDisplayLinkGetNominalOutputVideoRefreshPeriod(displayLink)
}

// Retrieves the nominal latency of a display link.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Retrieves the nominal latency of a display link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetOutputVideoLatency(_:)
func CVDisplayLinkGetOutputVideoLatency(displayLink DisplayLinkRef) Time {
	return _CVDisplayLinkGetOutputVideoLatency(displayLink)
}

// Obtains the Core Foundation ID for the display link data type.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Obtains the Core Foundation ID for the display link data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetTypeID()
func CVDisplayLinkGetTypeID() TypeID {
	return _CVDisplayLinkGetTypeID()
}

// Indicates whether a given display link is running.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Indicates whether a given display link is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkIsRunning(_:)
func CVDisplayLinkIsRunning(displayLink DisplayLinkRef) unsafe.Pointer {
	return _CVDisplayLinkIsRunning(displayLink)
}

// Sets the current display of a display link.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Sets the current display of a display link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetCurrentCGDisplay(_:_:)
func CVDisplayLinkSetCurrentCGDisplay(displayLink DisplayLinkRef, displayID DirectDisplayID) Return {
	return _CVDisplayLinkSetCurrentCGDisplay(displayLink, displayID)
}

// Selects the display link most optimal for the current renderer of an OpenGL context.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Selects the display link most optimal for the current renderer of an OpenGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(_:_:_:)
func CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(displayLink DisplayLinkRef, cglContext LContextObj, cglPixelFormat LPixelFormatObj) Return {
	return _CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(displayLink, cglContext, cglPixelFormat)
}

// Sets the renderer output callback function.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Sets the renderer output callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetOutputCallback(_:_:_:)
func CVDisplayLinkSetOutputCallback(displayLink DisplayLinkRef, callback DisplayLinkOutputCallback, userInfo unsafe.Pointer) Return {
	return _CVDisplayLinkSetOutputCallback(displayLink, callback, userInfo)
}

// CVDisplayLinkSetOutputHandler is a CoreVideo function.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetOutputHandler(_:_:)
func CVDisplayLinkSetOutputHandler(displayLink DisplayLinkRef, handler DisplayLinkOutputHandler) Return {
	return _CVDisplayLinkSetOutputHandler(displayLink, handler)
}

// Activates a display link.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Activates a display link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkStart(_:)
func CVDisplayLinkStart(displayLink DisplayLinkRef) Return {
	return _CVDisplayLinkStart(displayLink)
}

// Stops a display link.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Stops a display link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkStop(_:)
func CVDisplayLinkStop(displayLink DisplayLinkRef) Return {
	return _CVDisplayLinkStop(displayLink)
}

// Translates the time in the display link’s time base from one representation to another.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Translates the time in the display link’s time base from one representation to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkTranslateTime(_:_:_:)
func CVDisplayLinkTranslateTime(displayLink DisplayLinkRef, inTime unsafe.Pointer, outTime unsafe.Pointer) Return {
	return _CVDisplayLinkTranslateTime(displayLink, inTime, outTime)
}

// Returns the current system time.
//
// Added in macOS 10.4.
// Returns the current system time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVGetCurrentHostTime()
func CVGetCurrentHostTime() uint64 {
	return _CVGetCurrentHostTime()
}

// Returns the frequency of updates to the system time.
//
// Added in macOS 10.4.
// Returns the frequency of updates to the system time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVGetHostClockFrequency()
func CVGetHostClockFrequency() float64 {
	return _CVGetHostClockFrequency()
}

// Returns the smallest possible increment in the system time.
//
// Added in macOS 10.4.
// Returns the smallest possible increment in the system time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVGetHostClockMinimumTimeDelta()
func CVGetHostClockMinimumTimeDelta() uint32 {
	return _CVGetHostClockMinimumTimeDelta()
}

// Attempts to create a Core Graphics color space from the image buffer’s attachments that you specify.
//
// Added in macOS 10.8.
// Attempts to create a Core Graphics color space from the image buffer’s attachments that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferCreateColorSpaceFromAttachments(_:)
func CVImageBufferCreateColorSpaceFromAttachments(attachments DictionaryRef) ColorSpaceRef {
	return _CVImageBufferCreateColorSpaceFromAttachments(attachments)
}

// Returns the source rectangle of a Core Video image buffer that represents the clean aperture of the buffer in encoded pixels.
//
// Added in macOS 10.4.
// Returns the source rectangle of a Core Video image buffer that represents the clean aperture of the buffer in encoded pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetCleanRect(_:)
func CVImageBufferGetCleanRect(imageBuffer ImageBufferRef) corefoundation.CGRect {
	return _CVImageBufferGetCleanRect(imageBuffer)
}

// Returns the color space of a Core Video image buffer.
//
// Added in macOS 10.4.
// Returns the color space of a Core Video image buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetColorSpace(_:)
func CVImageBufferGetColorSpace(imageBuffer ImageBufferRef) ColorSpaceRef {
	return _CVImageBufferGetColorSpace(imageBuffer)
}

// Returns the nominal output display size, in square pixels, of a Core Video image buffer.
//
// Added in macOS 10.4.
// Returns the nominal output display size, in square pixels, of a Core Video image buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetDisplaySize(_:)
func CVImageBufferGetDisplaySize(imageBuffer ImageBufferRef) corefoundation.CGSize {
	return _CVImageBufferGetDisplaySize(imageBuffer)
}

// Returns the full encoded dimensions of a Core Video image buffer.
//
// Added in macOS 10.4.
// Returns the full encoded dimensions of a Core Video image buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetEncodedSize(_:)
func CVImageBufferGetEncodedSize(imageBuffer ImageBufferRef) corefoundation.CGSize {
	return _CVImageBufferGetEncodedSize(imageBuffer)
}

// Returns a Boolean value indicating whether the image is vertically flipped.
//
// Added in macOS 10.4.
// Returns a Boolean value indicating whether the image is vertically flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBufferIsFlipped(_:)
func CVImageBufferIsFlipped(imageBuffer ImageBufferRef) unsafe.Pointer {
	return _CVImageBufferIsFlipped(imageBuffer)
}

// CVIsCompressedPixelFormatAvailable is a CoreVideo function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVIsCompressedPixelFormatAvailable(_:)
func CVIsCompressedPixelFormatAvailable(pixelFormatType unsafe.Pointer) unsafe.Pointer {
	return _CVIsCompressedPixelFormatAvailable(pixelFormatType)
}

// Creates a new texture cache.
//
// Added in macOS 10.11.
// Creates a new texture cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheCreate(_:_:_:_:_:)
func CVMetalTextureCacheCreate(allocator AllocatorRef, cacheAttributes DictionaryRef, metalDevice unsafe.Pointer, textureAttributes DictionaryRef, cacheOut unsafe.Pointer) Return {
	return _CVMetalTextureCacheCreate(allocator, cacheAttributes, metalDevice, textureAttributes, cacheOut)
}

// Creates a Core Video Metal texture buffer from an existing image buffer.
//
// Added in macOS 10.11.
// Creates a Core Video Metal texture buffer from an existing image buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheCreateTextureFromImage(_:_:_:_:_:_:_:_:_:)
func CVMetalTextureCacheCreateTextureFromImage(allocator AllocatorRef, textureCache MetalTextureCacheRef, sourceImage ImageBufferRef, textureAttributes DictionaryRef, pixelFormat PixelFormat, width uintptr, height uintptr, planeIndex uintptr, textureOut unsafe.Pointer) Return {
	return _CVMetalTextureCacheCreateTextureFromImage(allocator, textureCache, sourceImage, textureAttributes, pixelFormat, width, height, planeIndex, textureOut)
}

// Manually flushes the contents of the provided texture cache.
//
// Added in macOS 10.11.
// Manually flushes the contents of the provided texture cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheFlush(_:_:)
func CVMetalTextureCacheFlush(textureCache MetalTextureCacheRef, options OptionFlags) {
	_CVMetalTextureCacheFlush(textureCache, options)
}

// Returns the Core Foundation type identifier for a Core Video Metal texture cache.
//
// Added in macOS 10.11.
// Returns the Core Foundation type identifier for a Core Video Metal texture cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheGetTypeID()
func CVMetalTextureCacheGetTypeID() TypeID {
	return _CVMetalTextureCacheGetTypeID()
}

// Returns convenient normalized texture coordinates for the part of the image that should be displayed.
//
// Added in macOS 10.11.
// Returns convenient normalized texture coordinates for the part of the image that should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetCleanTexCoords(_:_:_:_:_:)
func CVMetalTextureGetCleanTexCoords(image MetalTextureRef, lowerLeft float32, lowerRight unsafe.Pointer, upperRight unsafe.Pointer, upperLeft unsafe.Pointer, p5 unsafe.Pointer) {
	_CVMetalTextureGetCleanTexCoords(image, lowerLeft, lowerRight, upperRight, upperLeft, p5)
}

// Returns the Metal texture object for the image buffer.
//
// Added in macOS 10.11.
// Returns the Metal texture object for the image buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetTexture(_:)
func CVMetalTextureGetTexture(image MetalTextureRef) unsafe.Pointer {
	return _CVMetalTextureGetTexture(image)
}

// Returns the Core Foundation type identifier for a CoreVideo Metal texture-based image buffer.
//
// Added in macOS 10.11.
// Returns the Core Foundation type identifier for a CoreVideo Metal texture-based image buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetTypeID()
func CVMetalTextureGetTypeID() TypeID {
	return _CVMetalTextureGetTypeID()
}

// Returns a Boolean value indicating whether the texture image is vertically flipped.
//
// Added in macOS 10.11.
// Returns a Boolean value indicating whether the texture image is vertically flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureIsFlipped(_:)
func CVMetalTextureIsFlipped(image MetalTextureRef) unsafe.Pointer {
	return _CVMetalTextureIsFlipped(image)
}

// Attaches an OpenGL context to a Core Video OpenGL buffer.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Attaches an OpenGL context to a Core Video OpenGL buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferAttach(_:_:_:_:_:)
func CVOpenGLBufferAttach(openGLBuffer OpenGLBufferRef, cglContext LContextObj, face unsafe.Pointer, level unsafe.Pointer, screen unsafe.Pointer) Return {
	return _CVOpenGLBufferAttach(openGLBuffer, cglContext, face, level, screen)
}

// Creates a new Core Video OpenGL buffer that can be used for OpenGL rendering purposes
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Creates a new Core Video OpenGL buffer that can be used for OpenGL rendering purposes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferCreate(_:_:_:_:_:)
func CVOpenGLBufferCreate(allocator AllocatorRef, width uintptr, height uintptr, attributes DictionaryRef, bufferOut unsafe.Pointer) Return {
	return _CVOpenGLBufferCreate(allocator, width, height, attributes, bufferOut)
}

// Obtains the attributes of a Core Video OpenGL buffer.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Obtains the attributes of a Core Video OpenGL buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferGetAttributes(_:)
func CVOpenGLBufferGetAttributes(openGLBuffer OpenGLBufferRef) DictionaryRef {
	return _CVOpenGLBufferGetAttributes(openGLBuffer)
}

// Obtains the Core Foundation type ID for the OpenGL buffer type.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Obtains the Core Foundation type ID for the OpenGL buffer type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferGetTypeID()
func CVOpenGLBufferGetTypeID() TypeID {
	return _CVOpenGLBufferGetTypeID()
}

// Creates a new OpenGL buffer pool.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Creates a new OpenGL buffer pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolCreate(_:_:_:_:)
func CVOpenGLBufferPoolCreate(allocator AllocatorRef, poolAttributes DictionaryRef, openGLBufferAttributes DictionaryRef, poolOut unsafe.Pointer) Return {
	return _CVOpenGLBufferPoolCreate(allocator, poolAttributes, openGLBufferAttributes, poolOut)
}

// Creates a new OpenGL buffer from an OpenGL buffer pool.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Creates a new OpenGL buffer from an OpenGL buffer pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolCreateOpenGLBuffer(_:_:_:)
func CVOpenGLBufferPoolCreateOpenGLBuffer(allocator AllocatorRef, openGLBufferPool OpenGLBufferPoolRef, openGLBufferOut unsafe.Pointer) Return {
	return _CVOpenGLBufferPoolCreateOpenGLBuffer(allocator, openGLBufferPool, openGLBufferOut)
}

// Returns the pool attributes dictionary for an Open GL buffer pool.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Returns the pool attributes dictionary for an Open GL buffer pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolGetAttributes(_:)
func CVOpenGLBufferPoolGetAttributes(pool OpenGLBufferPoolRef) DictionaryRef {
	return _CVOpenGLBufferPoolGetAttributes(pool)
}

// Returns the attributes of OpenGL buffers that will be created from a buffer pool.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Returns the attributes of OpenGL buffers that will be created from a buffer pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolGetOpenGLBufferAttributes(_:)
func CVOpenGLBufferPoolGetOpenGLBufferAttributes(pool OpenGLBufferPoolRef) DictionaryRef {
	return _CVOpenGLBufferPoolGetOpenGLBufferAttributes(pool)
}

// Obtains the Core Foundation ID for the OpenGL buffer pool type.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Obtains the Core Foundation ID for the OpenGL buffer pool type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolGetTypeID()
func CVOpenGLBufferPoolGetTypeID() TypeID {
	return _CVOpenGLBufferPoolGetTypeID()
}

// Creates a new Core Video texture cache.

// Creates a new Core Video texture cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureCacheCreate(_:_:_:_:_:)
func CVOpenGLESTextureCacheCreate(allocator AllocatorRef, cacheAttributes DictionaryRef, eaglContext EAGLContext, textureAttributes DictionaryRef, cacheOut unsafe.Pointer) Return {
	return _CVOpenGLESTextureCacheCreate(allocator, cacheAttributes, eaglContext, textureAttributes, cacheOut)
}

// Creates a object from an existing .

// Creates a object from an existing .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureCacheCreateTextureFromImage(_:_:_:_:_:_:_:_:_:_:_:_:)
func CVOpenGLESTextureCacheCreateTextureFromImage(allocator AllocatorRef, textureCache OpenGLESTextureCacheRef, sourceImage ImageBufferRef, textureAttributes DictionaryRef, target unsafe.Pointer, internalFormat unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, format unsafe.Pointer, type_ unsafe.Pointer, planeIndex uintptr, textureOut unsafe.Pointer) Return {
	return _CVOpenGLESTextureCacheCreateTextureFromImage(allocator, textureCache, sourceImage, textureAttributes, target, internalFormat, width, height, format, type_, planeIndex, textureOut)
}

// Performs internal housekeeping/recycling operations on a texture cache.

// Performs internal housekeeping/recycling operations on a texture cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureCacheFlush(_:_:)
func CVOpenGLESTextureCacheFlush(textureCache OpenGLESTextureCacheRef, options OptionFlags) {
	_CVOpenGLESTextureCacheFlush(textureCache, options)
}

// Returns the Core Foundation type identifier for a Core Video texture cache.

// Returns the Core Foundation type identifier for a Core Video texture cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureCacheGetTypeID()
func CVOpenGLESTextureCacheGetTypeID() TypeID {
	return _CVOpenGLESTextureCacheGetTypeID()
}

// Returns convenient normalized texture coordinates for the part of the image that should be displayed.

// Returns convenient normalized texture coordinates for the part of the image that should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureGetCleanTexCoords(_:_:_:_:_:)
func CVOpenGLESTextureGetCleanTexCoords(image OpenGLESTextureRef, lowerLeft unsafe.Pointer, lowerRight unsafe.Pointer, upperRight unsafe.Pointer, upperLeft unsafe.Pointer, p5 unsafe.Pointer) {
	_CVOpenGLESTextureGetCleanTexCoords(image, lowerLeft, lowerRight, upperRight, upperLeft, p5)
}

// Returns the texture target name for a .

// Returns the texture target name for a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureGetName(_:)
func CVOpenGLESTextureGetName(image OpenGLESTextureRef) unsafe.Pointer {
	return _CVOpenGLESTextureGetName(image)
}

// Returns the texture target for a .

// Returns the texture target for a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureGetTarget(_:)
func CVOpenGLESTextureGetTarget(image OpenGLESTextureRef) unsafe.Pointer {
	return _CVOpenGLESTextureGetTarget(image)
}

// Returns the Core Foundation type identifier for a Core Video texture-based image buffer.

// Returns the Core Foundation type identifier for a Core Video texture-based image buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureGetTypeID()
func CVOpenGLESTextureGetTypeID() TypeID {
	return _CVOpenGLESTextureGetTypeID()
}

// Returns whether the image is flipped vertically or not.

// Returns whether the image is flipped vertically or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureIsFlipped(_:)
func CVOpenGLESTextureIsFlipped(image OpenGLESTextureRef) unsafe.Pointer {
	return _CVOpenGLESTextureIsFlipped(image)
}

// Creates a new texture cache.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Creates a new texture cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheCreate(_:_:_:_:_:_:)
func CVOpenGLTextureCacheCreate(allocator AllocatorRef, cacheAttributes DictionaryRef, cglContext LContextObj, cglPixelFormat LPixelFormatObj, textureAttributes DictionaryRef, cacheOut unsafe.Pointer) Return {
	return _CVOpenGLTextureCacheCreate(allocator, cacheAttributes, cglContext, cglPixelFormat, textureAttributes, cacheOut)
}

// Creates a CVOpenGLTexture object from an existing .
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Creates a CVOpenGLTexture object from an existing .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheCreateTextureFromImage(_:_:_:_:_:)
func CVOpenGLTextureCacheCreateTextureFromImage(allocator AllocatorRef, textureCache OpenGLTextureCacheRef, sourceImage ImageBufferRef, attributes DictionaryRef, textureOut unsafe.Pointer) Return {
	return _CVOpenGLTextureCacheCreateTextureFromImage(allocator, textureCache, sourceImage, attributes, textureOut)
}

// Performs internal housekeeping/recycling operations on the cache.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Performs internal housekeeping/recycling operations on the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheFlush(_:_:)
func CVOpenGLTextureCacheFlush(textureCache OpenGLTextureCacheRef, options OptionFlags) {
	_CVOpenGLTextureCacheFlush(textureCache, options)
}

// Returns the Core Foundation type identifier for a the texture cache.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Returns the Core Foundation type identifier for a the texture cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheGetTypeID()
func CVOpenGLTextureCacheGetTypeID() TypeID {
	return _CVOpenGLTextureCacheGetTypeID()
}

// Returns the texture coordinates for the part of the image that should be displayed.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Returns the texture coordinates for the part of the image that should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureGetCleanTexCoords(_:_:_:_:_:)
func CVOpenGLTextureGetCleanTexCoords(image OpenGLTextureRef, lowerLeft unsafe.Pointer, lowerRight unsafe.Pointer, upperRight unsafe.Pointer, upperLeft unsafe.Pointer, p5 unsafe.Pointer) {
	_CVOpenGLTextureGetCleanTexCoords(image, lowerLeft, lowerRight, upperRight, upperLeft, p5)
}

// Returns the texture target name of a CoreVideo OpenGL texture.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Returns the texture target name of a CoreVideo OpenGL texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureGetName(_:)
func CVOpenGLTextureGetName(image OpenGLTextureRef) unsafe.Pointer {
	return _CVOpenGLTextureGetName(image)
}

// Returns the texture target (for example, ) of an OpenGL texture.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Returns the texture target (for example, ) of an OpenGL texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureGetTarget(_:)
func CVOpenGLTextureGetTarget(image OpenGLTextureRef) unsafe.Pointer {
	return _CVOpenGLTextureGetTarget(image)
}

// Obtains the Core Foundation ID for the Core Video OpenGL texture type.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Obtains the Core Foundation ID for the Core Video OpenGL texture type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureGetTypeID()
func CVOpenGLTextureGetTypeID() TypeID {
	return _CVOpenGLTextureGetTypeID()
}

// Determines whether an OpenGL texture is flipped vertically.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.4.
// Determines whether an OpenGL texture is flipped vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureIsFlipped(_:)
func CVOpenGLTextureIsFlipped(image OpenGLTextureRef) unsafe.Pointer {
	return _CVOpenGLTextureIsFlipped(image)
}

// CVPixelBufferCopyCreationAttributes is a CoreVideo function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCopyCreationAttributes(_:)
func CVPixelBufferCopyCreationAttributes(pixelBuffer PixelBufferRef) DictionaryRef {
	return _CVPixelBufferCopyCreationAttributes(pixelBuffer)
}

// Creates a single pixel buffer for a given size and pixel format.
//
// Added in macOS 10.4.
// Creates a single pixel buffer for a given size and pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreate(_:_:_:_:_:_:)
func CVPixelBufferCreate(allocator AllocatorRef, width uintptr, height uintptr, pixelFormatType unsafe.Pointer, pixelBufferAttributes DictionaryRef, pixelBufferOut unsafe.Pointer) Return {
	return _CVPixelBufferCreate(allocator, width, height, pixelFormatType, pixelBufferAttributes, pixelBufferOut)
}

// Resolves an array of objects describing various pixel buffer attributes into a single dictionary.
//
// Added in macOS 10.4.
// Resolves an array of objects describing various pixel buffer attributes into a single dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateResolvedAttributesDictionary(_:_:_:)
func CVPixelBufferCreateResolvedAttributesDictionary(allocator AllocatorRef, attributes ArrayRef, resolvedDictionaryOut unsafe.Pointer) Return {
	return _CVPixelBufferCreateResolvedAttributesDictionary(allocator, attributes, resolvedDictionaryOut)
}

// Creates a pixel buffer for a given size and pixel format containing data specified by a memory location.
//
// Added in macOS 10.4.
// Creates a pixel buffer for a given size and pixel format containing data specified by a memory location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithBytes(_:_:_:_:_:_:_:_:_:_:)
func CVPixelBufferCreateWithBytes(allocator AllocatorRef, width uintptr, height uintptr, pixelFormatType unsafe.Pointer, baseAddress unsafe.Pointer, bytesPerRow uintptr, releaseCallback PixelBufferReleaseBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes DictionaryRef, pixelBufferOut unsafe.Pointer) Return {
	return _CVPixelBufferCreateWithBytes(allocator, width, height, pixelFormatType, baseAddress, bytesPerRow, releaseCallback, releaseRefCon, pixelBufferAttributes, pixelBufferOut)
}

// Creates a single pixel buffer for the IO surface that you specify.
//
// Added in macOS 10.6.
// Creates a single pixel buffer for the IO surface that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithIOSurface(_:_:_:_:)
func CVPixelBufferCreateWithIOSurface(allocator AllocatorRef, surface SurfaceRef, pixelBufferAttributes DictionaryRef, pixelBufferOut unsafe.Pointer) Return {
	return _CVPixelBufferCreateWithIOSurface(allocator, surface, pixelBufferAttributes, pixelBufferOut)
}

// Creates a single pixel buffer in planar format for a given size and pixel format containing data specified by a memory location.
//
// Added in macOS 10.4.
// Creates a single pixel buffer in planar format for a given size and pixel format containing data specified by a memory location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithPlanarBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CVPixelBufferCreateWithPlanarBytes(allocator AllocatorRef, width uintptr, height uintptr, pixelFormatType unsafe.Pointer, dataPtr unsafe.Pointer, dataSize uintptr, numberOfPlanes uintptr, planeBaseAddress unsafe.Pointer, planeWidth uintptr, planeHeight uintptr, planeBytesPerRow uintptr, releaseCallback PixelBufferReleasePlanarBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes DictionaryRef, pixelBufferOut unsafe.Pointer) Return {
	return _CVPixelBufferCreateWithPlanarBytes(allocator, width, height, pixelFormatType, dataPtr, dataSize, numberOfPlanes, planeBaseAddress, planeWidth, planeHeight, planeBytesPerRow, releaseCallback, releaseRefCon, pixelBufferAttributes, pixelBufferOut)
}

// Fills the extended pixels of the pixel buffer.
//
// Added in macOS 10.4.
// Fills the extended pixels of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferFillExtendedPixels(_:)
func CVPixelBufferFillExtendedPixels(pixelBuffer PixelBufferRef) Return {
	return _CVPixelBufferFillExtendedPixels(pixelBuffer)
}

// Returns the base address of the pixel buffer.
//
// Added in macOS 10.4.
// Returns the base address of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBaseAddress(_:)
func CVPixelBufferGetBaseAddress(pixelBuffer PixelBufferRef) unsafe.Pointer {
	return _CVPixelBufferGetBaseAddress(pixelBuffer)
}

// Returns the base address of the plane at the specified plane index.
//
// Added in macOS 10.4.
// Returns the base address of the plane at the specified plane index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBaseAddressOfPlane(_:_:)
func CVPixelBufferGetBaseAddressOfPlane(pixelBuffer PixelBufferRef, planeIndex uintptr) unsafe.Pointer {
	return _CVPixelBufferGetBaseAddressOfPlane(pixelBuffer, planeIndex)
}

// Returns the number of bytes per row of the pixel buffer.
//
// Added in macOS 10.4.
// Returns the number of bytes per row of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBytesPerRow(_:)
func CVPixelBufferGetBytesPerRow(pixelBuffer PixelBufferRef) uintptr {
	return _CVPixelBufferGetBytesPerRow(pixelBuffer)
}

// Returns the number of bytes per row for a plane at the specified index in the pixel buffer.
//
// Added in macOS 10.4.
// Returns the number of bytes per row for a plane at the specified index in the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBytesPerRowOfPlane(_:_:)
func CVPixelBufferGetBytesPerRowOfPlane(pixelBuffer PixelBufferRef, planeIndex uintptr) uintptr {
	return _CVPixelBufferGetBytesPerRowOfPlane(pixelBuffer, planeIndex)
}

// Returns the data size for contiguous planes of the pixel buffer.
//
// Added in macOS 10.4.
// Returns the data size for contiguous planes of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetDataSize(_:)
func CVPixelBufferGetDataSize(pixelBuffer PixelBufferRef) uintptr {
	return _CVPixelBufferGetDataSize(pixelBuffer)
}

// Returns the amount of extended pixel padding in the pixel buffer.
//
// Added in macOS 10.4.
// Returns the amount of extended pixel padding in the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetExtendedPixels(_:_:_:_:_:)
func CVPixelBufferGetExtendedPixels(pixelBuffer PixelBufferRef, extraColumnsOnLeft unsafe.Pointer, extraColumnsOnRight unsafe.Pointer, extraRowsOnTop unsafe.Pointer, extraRowsOnBottom unsafe.Pointer) {
	_CVPixelBufferGetExtendedPixels(pixelBuffer, extraColumnsOnLeft, extraColumnsOnRight, extraRowsOnTop, extraRowsOnBottom)
}

// Returns the height of the pixel buffer.
//
// Added in macOS 10.4.
// Returns the height of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetHeight(_:)
func CVPixelBufferGetHeight(pixelBuffer PixelBufferRef) uintptr {
	return _CVPixelBufferGetHeight(pixelBuffer)
}

// Returns the height of the plane at planeIndex in the pixel buffer.
//
// Added in macOS 10.4.
// Returns the height of the plane at planeIndex in the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetHeightOfPlane(_:_:)
func CVPixelBufferGetHeightOfPlane(pixelBuffer PixelBufferRef, planeIndex uintptr) uintptr {
	return _CVPixelBufferGetHeightOfPlane(pixelBuffer, planeIndex)
}

// Returns the IOSurface backing the pixel buffer, or if it is not backed by an IOSurface.
//
// Added in macOS 10.6.
// Returns the IOSurface backing the pixel buffer, or if it is not backed by an IOSurface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetIOSurface(_:)
func CVPixelBufferGetIOSurface(pixelBuffer PixelBufferRef) SurfaceRef {
	return _CVPixelBufferGetIOSurface(pixelBuffer)
}

// Returns the pixel format type of the pixel buffer.
//
// Added in macOS 10.4.
// Returns the pixel format type of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetPixelFormatType(_:)
func CVPixelBufferGetPixelFormatType(pixelBuffer PixelBufferRef) unsafe.Pointer {
	return _CVPixelBufferGetPixelFormatType(pixelBuffer)
}

// Returns number of planes of the pixel buffer.
//
// Added in macOS 10.4.
// Returns number of planes of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetPlaneCount(_:)
func CVPixelBufferGetPlaneCount(pixelBuffer PixelBufferRef) uintptr {
	return _CVPixelBufferGetPlaneCount(pixelBuffer)
}

// Returns the Core Foundation type identifier of the pixel buffer type.
//
// Added in macOS 10.4.
// Returns the Core Foundation type identifier of the pixel buffer type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetTypeID()
func CVPixelBufferGetTypeID() TypeID {
	return _CVPixelBufferGetTypeID()
}

// Returns the width of the pixel buffer.
//
// Added in macOS 10.4.
// Returns the width of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetWidth(_:)
func CVPixelBufferGetWidth(pixelBuffer PixelBufferRef) uintptr {
	return _CVPixelBufferGetWidth(pixelBuffer)
}

// Returns the width of the plane at a given index in the pixel buffer.
//
// Added in macOS 10.4.
// Returns the width of the plane at a given index in the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetWidthOfPlane(_:_:)
func CVPixelBufferGetWidthOfPlane(pixelBuffer PixelBufferRef, planeIndex uintptr) uintptr {
	return _CVPixelBufferGetWidthOfPlane(pixelBuffer, planeIndex)
}

// Determines whether the pixel buffer is planar.
//
// Added in macOS 10.4.
// Determines whether the pixel buffer is planar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferIsPlanar(_:)
func CVPixelBufferIsPlanar(pixelBuffer PixelBufferRef) unsafe.Pointer {
	return _CVPixelBufferIsPlanar(pixelBuffer)
}

// Locks the base address of the pixel buffer.
//
// Added in macOS 10.4.
// Locks the base address of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferLockBaseAddress(_:_:)
func CVPixelBufferLockBaseAddress(pixelBuffer PixelBufferRef, lockFlags PixelBufferLockFlags) Return {
	return _CVPixelBufferLockBaseAddress(pixelBuffer, lockFlags)
}

// Creates a pixel buffer pool using the allocator and attributes that you specify.
//
// Added in macOS 10.4.
// Creates a pixel buffer pool using the allocator and attributes that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreate(_:_:_:_:)
func CVPixelBufferPoolCreate(allocator AllocatorRef, poolAttributes DictionaryRef, pixelBufferAttributes DictionaryRef, poolOut unsafe.Pointer) Return {
	return _CVPixelBufferPoolCreate(allocator, poolAttributes, pixelBufferAttributes, poolOut)
}

// Creates a pixel buffer from a pixel buffer pool, using the allocator that you specify.
//
// Added in macOS 10.4.
// Creates a pixel buffer from a pixel buffer pool, using the allocator that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreatePixelBuffer(_:_:_:)
func CVPixelBufferPoolCreatePixelBuffer(allocator AllocatorRef, pixelBufferPool PixelBufferPoolRef, pixelBufferOut unsafe.Pointer) Return {
	return _CVPixelBufferPoolCreatePixelBuffer(allocator, pixelBufferPool, pixelBufferOut)
}

// Creates a new pixel buffer with auxiliary attributes from the pool.
//
// Added in macOS 10.7.
// Creates a new pixel buffer with auxiliary attributes from the pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(_:_:_:_:)
func CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(allocator AllocatorRef, pixelBufferPool PixelBufferPoolRef, auxAttributes DictionaryRef, pixelBufferOut unsafe.Pointer) Return {
	return _CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(allocator, pixelBufferPool, auxAttributes, pixelBufferOut)
}

// Frees pixel buffers from the pool based on the options that you specify.
//
// Added in macOS 10.4.
// Frees pixel buffers from the pool based on the options that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolFlush(_:_:)
func CVPixelBufferPoolFlush(pool PixelBufferPoolRef, options PixelBufferPoolFlushFlags) {
	_CVPixelBufferPoolFlush(pool, options)
}

// The pool attributes dictionary for a pixel buffer pool.
//
// Added in macOS 10.4.
// The pool attributes dictionary for a pixel buffer pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetAttributes(_:)
func CVPixelBufferPoolGetAttributes(pool PixelBufferPoolRef) DictionaryRef {
	return _CVPixelBufferPoolGetAttributes(pool)
}

// The attributes of pixel buffers which the system creates using the pool you specify.
//
// Added in macOS 10.4.
// The attributes of pixel buffers which the system creates using the pool you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetPixelBufferAttributes(_:)
func CVPixelBufferPoolGetPixelBufferAttributes(pool PixelBufferPoolRef) DictionaryRef {
	return _CVPixelBufferPoolGetPixelBufferAttributes(pool)
}

// Returns the Core Foundation type identifier of the pixel buffer pool type.
//
// Added in macOS 10.4.
// Returns the Core Foundation type identifier of the pixel buffer pool type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetTypeID()
func CVPixelBufferPoolGetTypeID() TypeID {
	return _CVPixelBufferPoolGetTypeID()
}

// Unlocks the base address of the pixel buffer.
//
// Added in macOS 10.4.
// Unlocks the base address of the pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferUnlockBaseAddress(_:_:)
func CVPixelBufferUnlockBaseAddress(pixelBuffer PixelBufferRef, unlockFlags PixelBufferLockFlags) Return {
	return _CVPixelBufferUnlockBaseAddress(pixelBuffer, unlockFlags)
}

// Returns all the pixel format descriptions known to Core Video.
//
// Added in macOS 10.4.
// Returns all the pixel format descriptions known to Core Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes(_:)
func CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes(allocator AllocatorRef) ArrayRef {
	return _CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes(allocator)
}

// Creates a pixel format description from a given identifier.
//
// Added in macOS 10.4.
// Creates a pixel format description from a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatDescriptionCreateWithPixelFormatType(_:_:)
func CVPixelFormatDescriptionCreateWithPixelFormatType(allocator AllocatorRef, pixelFormat unsafe.Pointer) DictionaryRef {
	return _CVPixelFormatDescriptionCreateWithPixelFormatType(allocator, pixelFormat)
}

// Registers a pixel format description with Core Video.
//
// Added in macOS 10.4.
// Registers a pixel format description with Core Video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType(_:_:)
func CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType(description DictionaryRef, pixelFormat unsafe.Pointer) {
	_CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType(description, pixelFormat)
}

// Returns the standard integer code point corresponding to the Core Video transfer function string that you specify.
//
// Added in macOS 10.13.
// Returns the standard integer code point corresponding to the Core Video transfer function string that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTransferFunctionGetIntegerCodePointForString(_:)
func CVTransferFunctionGetIntegerCodePointForString(transferFunctionString StringRef) int {
	return _CVTransferFunctionGetIntegerCodePointForString(transferFunctionString)
}

// Returns the Core Video transfer function string corresponding to the standard integer code point that you specify.
//
// Added in macOS 10.13.
// Returns the Core Video transfer function string corresponding to the standard integer code point that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTransferFunctionGetStringForIntegerCodePoint(_:)
func CVTransferFunctionGetStringForIntegerCodePoint(transferFunctionCodePoint int) StringRef {
	return _CVTransferFunctionGetStringForIntegerCodePoint(transferFunctionCodePoint)
}

// Returns the standard integer code point corresponding to the Core Video YCbCr matrix string that you specify.
//
// Added in macOS 10.13.
// Returns the standard integer code point corresponding to the Core Video YCbCr matrix string that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVYCbCrMatrixGetIntegerCodePointForString(_:)
func CVYCbCrMatrixGetIntegerCodePointForString(yCbCrMatrixString StringRef) int {
	return _CVYCbCrMatrixGetIntegerCodePointForString(yCbCrMatrixString)
}

// Returns the Core Video YCbCr matrix string corresponding to the standard integer code point that you specify.
//
// Added in macOS 10.13.
// Returns the Core Video YCbCr matrix string corresponding to the standard integer code point that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVYCbCrMatrixGetStringForIntegerCodePoint(_:)
func CVYCbCrMatrixGetStringForIntegerCodePoint(yCbCrMatrixCodePoint int) StringRef {
	return _CVYCbCrMatrixGetStringForIntegerCodePoint(yCbCrMatrixCodePoint)
}




