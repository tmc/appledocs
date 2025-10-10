// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo

// CoreVideo Functions
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

// Discovered functions (52 total):

// CVBufferCopyAttachment(buffer _, key :  CVBuffer,  _, attachmentMode :  CFString,  _, :  UnsafeMutablePointer< CVAttachmentMode>?) ->  CFTypeRef?) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// CVBufferPropagateAttachments(sourceBuffer _, destinationBuffer :  CVBuffer,  _, :  CVBuffer) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CVColorPrimariesGetStringForIntegerCodePoint(colorPrimariesCodePoint _, :  Int32) ->  Unmanaged< CFString>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+


// CVDisplayLinkGetOutputVideoLatency(displayLink _, :  CVDisplayLink) ->  CVTime) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - macOS 10.4+ (Deprecated in 15.0)
//
// Deprecated: This function is deprecated.

// CVDisplayLinkSetCurrentCGDisplay(displayLink _, displayID :  CVDisplayLink,  _, :  CGDirectDisplayID) ->  CVReturn) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - macOS 10.4+ (Deprecated in 15.0)
//
// Deprecated: This function is deprecated.

// CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(displayLink _, cglContext :  CVDisplayLink,  _, cglPixelFormat :  CGLContextObj,  _, :  CGLPixelFormatObj) ->  CVReturn) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - macOS 10.4+ (Deprecated in 15.0)
//
// Deprecated: This function is deprecated.


// CVImageBufferGetDisplaySize(imageBuffer _, :  CVImageBuffer) ->  CGSize) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CVImageBufferGetEncodedSize(imageBuffer _, :  CVImageBuffer) ->  CGSize) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CVIsCompressedPixelFormatAvailable(pixelFormatType _, :  OSType) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// CVMetalBufferCacheCreate(allocator _, cacheAttributes :  CFAllocator?,  _, metalDevice :  CFDictionary?,  _, cacheOut : any  MTLDevice,  _, :  UnsafeMutablePointer< CVMetalBufferCache?>) ->  CVReturn) func
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+

// CVMetalBufferCacheCreateBufferFromImage(allocator _, bufferCache :  CFAllocator?,  _, imageBuffer :  CVMetalBufferCache,  _, bufferOut :  CVImageBuffer,  _, :  UnsafeMutablePointer< CVMetalBuffer?>) ->  CVReturn) func
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+

// CVMetalBufferCacheFlush(bufferCache _, options :  CVMetalBufferCache,  _, :  CVOptionFlags) func
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+


// CVMetalBufferCacheGetTypeID() func
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+

// CVMetalBufferGetBuffer(buffer _, :  CVMetalBuffer) -> (any  MTLBuffer)?) func
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+

// CVMetalBufferGetTypeID() func
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+


// CVMetalTextureCacheCreate(allocator _, cacheAttributes :  CFAllocator?,  _, metalDevice :  CFDictionary?,  _, textureAttributes : any  MTLDevice,  _, cacheOut :  CFDictionary?,  _, :  UnsafeMutablePointer< CVMetalTextureCache?>) ->  CVReturn) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CVMetalTextureCacheFlush(textureCache _, options :  CVMetalTextureCache,  _, :  CVOptionFlags) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CVMetalTextureGetCleanTexCoords(image _, lowerLeft :  CVMetalTexture,  _, lowerRight :  UnsafeMutablePointer< Float>,  _, upperRight :  UnsafeMutablePointer< Float>,  _, upperLeft :  UnsafeMutablePointer< Float>,  _, :  UnsafeMutablePointer< Float>)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+


// CVMetalTextureGetTexture(image _, :  CVMetalTexture) -> (any  MTLTexture)?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CVMetalTextureGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CVOpenGLBufferCreate(allocator _, width :  CFAllocator?,  _, height :  Int,  _, attributes :  Int,  _, bufferOut :  CFDictionary?,  _, :  UnsafeMutablePointer< CVOpenGLBuffer?>) ->  CVReturn) func
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.


// CVOpenGLBufferPoolCreate(allocator _, poolAttributes :  CFAllocator?,  _, openGLBufferAttributes :  CFDictionary?,  _, poolOut :  CFDictionary?,  _, :  UnsafeMutablePointer< CVOpenGLBufferPool?>) ->  CVReturn) func
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// CVOpenGLBufferPoolGetAttributes(pool _, :  CVOpenGLBufferPool) ->  Unmanaged< CFDictionary>?) func
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// CVOpenGLBufferPoolGetOpenGLBufferAttributes(pool _, :  CVOpenGLBufferPool) ->  Unmanaged< CFDictionary>?) func
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.


// CVOpenGLBufferPoolRetain(openGLBufferPool CVOpenGLBufferPoolRef, );) extern   CVOpenGLBufferPoolRef
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// CVOpenGLBufferRetain(buffer CVOpenGLBufferRef, );) extern   CVOpenGLBufferRef
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// CVOpenGLESTextureCacheCreate(allocator _, cacheAttributes :  CFAllocator?,  _, eaglContext :  CFDictionary?,  _, textureAttributes :  CVEAGLContext,  _, cacheOut :  CFDictionary?,  _, :  UnsafeMutablePointer< CVOpenGLESTextureCache?>) ->  CVReturn) func
//
// Availability:
//   - iOS 5.0+ (Deprecated in 12.0)
//   - iPadOS 5.0+ (Deprecated in 12.0)
//   - tvOS 9.0+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.


// CVOpenGLESTextureCacheCreateTextureFromImage(allocator _, textureCache :  CFAllocator?,  _, sourceImage :  CVOpenGLESTextureCache,  _, textureAttributes :  CVImageBuffer,  _, target :  CFDictionary?,  _, internalFormat :  GLenum,  _, width :  GLint,  _, height :  GLsizei,  _, format :  GLsizei,  _, type :  GLenum,  _, planeIndex :  GLenum,  _, textureOut :  Int,  _, :  UnsafeMutablePointer< CVOpenGLESTexture?>) ->  CVReturn) func
//
// Availability:
//   - iOS 5.0+ (Deprecated in 12.0)
//   - iPadOS 5.0+ (Deprecated in 12.0)
//   - tvOS 9.0+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// CVOpenGLESTextureGetCleanTexCoords(image _, lowerLeft :  CVOpenGLESTexture,  _, lowerRight :  UnsafeMutablePointer< GLfloat>,  _, upperRight :  UnsafeMutablePointer< GLfloat>,  _, upperLeft :  UnsafeMutablePointer< GLfloat>,  _, :  UnsafeMutablePointer< GLfloat>)) func
//
// Availability:
//   - iOS 5.0+ (Deprecated in 12.0)
//   - iPadOS 5.0+ (Deprecated in 12.0)
//   - tvOS 9.0+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// CVOpenGLESTextureGetName(image _, :  CVOpenGLESTexture) ->  GLuint) func
//
// Availability:
//   - iOS 5.0+ (Deprecated in 12.0)
//   - iPadOS 5.0+ (Deprecated in 12.0)
//   - tvOS 9.0+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.


// CVOpenGLESTextureGetTarget(image _, :  CVOpenGLESTexture) ->  GLenum) func
//
// Availability:
//   - iOS 5.0+ (Deprecated in 12.0)
//   - iPadOS 5.0+ (Deprecated in 12.0)
//   - tvOS 9.0+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// CVOpenGLESTextureGetTypeID() func
//
// Availability:
//   - iOS 5.0+ (Deprecated in 12.0)
//   - iPadOS 5.0+ (Deprecated in 12.0)
//   - tvOS 9.0+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// CVOpenGLESTextureIsFlipped(image _, :  CVOpenGLESTexture) ->  Bool) func
//
// Availability:
//   - iOS 5.0+ (Deprecated in 12.0)
//   - iPadOS 5.0+ (Deprecated in 12.0)
//   - tvOS 9.0+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.


// CVOpenGLTextureCacheRelease(textureCache CVOpenGLTextureCacheRef, );) extern   void
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// CVOpenGLTextureCacheRetain(textureCache CVOpenGLTextureCacheRef, );) extern   CVOpenGLTextureCacheRef
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// CVOpenGLTextureIsFlipped(image _, :  CVOpenGLTexture) ->  Bool) func
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.


// CVOpenGLTextureRetain(texture CVOpenGLTextureRef, );) extern   CVOpenGLTextureRef
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// CVPixelBufferCreateResolvedAttributesDictionary(allocator _, attributes :  CFAllocator?,  _, resolvedDictionaryOut :  CFArray?,  _, :  UnsafeMutablePointer< CFDictionary?>) ->  CVReturn) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CVPixelBufferFillExtendedPixels(pixelBuffer _, :  CVPixelBuffer) ->  CVReturn) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CVPixelBufferGetBaseAddressOfPlane(pixelBuffer _, planeIndex :  CVPixelBuffer,  _, :  Int) ->  UnsafeMutableRawPointer?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CVPixelBufferGetPixelFormatType(pixelBuffer _, :  CVPixelBuffer) ->  OSType) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CVPixelBufferGetPlaneCount(pixelBuffer _, :  CVPixelBuffer) ->  Int) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CVPixelBufferIsCompatibleWithAttributes(pixelBuffer _, attributes :  CVPixelBuffer,  _, :  CFDictionary?) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 4.0+
//   - visionOS 1.0+
//   - watchOS 1.0+

// CVPixelBufferIsPlanar(pixelBuffer _, :  CVPixelBuffer) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CVPixelBufferLockBaseAddress(pixelBuffer _, lockFlags :  CVPixelBuffer,  _, :  CVPixelBufferLockFlags) ->  CVReturn) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CVPixelBufferPoolCreate(allocator _, poolAttributes :  CFAllocator?,  _, pixelBufferAttributes :  CFDictionary?,  _, poolOut :  CFDictionary?,  _, :  UnsafeMutablePointer< CVPixelBufferPool?>) ->  CVReturn) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CVPixelBufferPoolGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CVPixelBufferPoolRelease(pixelBufferPool CVPixelBufferPoolRef, );) extern   void
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+


// CVPixelBufferUnlockBaseAddress(pixelBuffer _, unlockFlags :  CVPixelBuffer,  _, :  CVPixelBufferLockFlags) ->  CVReturn) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CVPixelFormatTypeCopyFourCharCodeString(pixelFormat _, :  OSType) ->  CFString) func
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CVTransferFunctionGetStringForIntegerCodePoint(transferFunctionCodePoint _, :  Int32) ->  Unmanaged< CFString>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+


// CVYCbCrMatrixGetIntegerCodePointForString(yCbCrMatrixString _, :  CFString?) ->  Int32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+

