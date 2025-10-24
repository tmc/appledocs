// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo
import (
"unsafe"
)

// Type aliases and typedefs
// MetalBufferRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBuffer
// CVMetalBufferRef has base type: CVBufferRef
type MetalBufferRef uintptr
// MetalBufferCacheRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCache
// CVMetalBufferCacheRef has base type: struct __CVMetalBufferCache *
type MetalBufferCacheRef uintptr
// BufferRef - A reference to a Core Video buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBuffer
// CVBufferRef has base type: struct __CVBuffer *
type BufferRef uintptr
// DisplayLinkRef - A reference to a display link object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLink
// CVDisplayLinkRef has base type: struct __CVDisplayLink *
type DisplayLinkRef uintptr
// DisplayLinkOutputCallback - A type for a display link callback function that the system invokes when it’s time for the app to output a video frame.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkOutputCallback
// CVDisplayLinkOutputCallback is a callback function
// C type: int (*)(struct __CVDisplayLink *, const CVTimeStamp *, const CVTimeStamp *, unsigned long long, unsigned long long *, void *)
type DisplayLinkOutputCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint64, uint64, unsafe.Pointer) int32
// EAGLContext - A type that resolves to an   pointer when appropriate.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVEAGLContext
// CVEAGLContext has base type: void *
type EAGLContext uintptr
// FillExtendedPixelsCallBack - Defines a pointer to a custom extended pixel-fill function, which is called whenever the system needs to pad a buffer holding your custom pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVFillExtendedPixelsCallBack
// CVFillExtendedPixelsCallBack is a callback function
// C type: unsigned char (*)(struct __CVBuffer *, void *)
type FillExtendedPixelsCallBack = func(unsafe.Pointer, unsafe.Pointer) uint8
// ImageBufferRef - A reference to a Core Video image buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBuffer
// CVImageBufferRef has base type: CVBufferRef
type ImageBufferRef uintptr
// MetalTextureRef - A reference to a CoreVideo Metal texture-based image buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTexture
// CVMetalTextureRef has base type: CVImageBufferRef
type MetalTextureRef uintptr
// MetalTextureCacheRef - A reference to a Core Video Metal texture cache.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCache
// CVMetalTextureCacheRef has base type: struct __CVMetalTextureCache *
type MetalTextureCacheRef uintptr
// OpenGLBufferRef - A reference to a Core Video OpenGL buffer object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBuffer
// CVOpenGLBufferRef has base type: CVImageBufferRef
type OpenGLBufferRef uintptr
// OpenGLBufferPoolRef - A reference to an OpenGL buffer pool object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPool
// CVOpenGLBufferPoolRef has base type: struct __CVOpenGLBufferPool *
type OpenGLBufferPoolRef uintptr
// OpenGLESTextureRef - A reference to a Core Video texture-based image buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTexture
// CVOpenGLESTextureRef has base type: CVImageBufferRef
type OpenGLESTextureRef uintptr
// OpenGLESTextureCacheRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureCache
// CVOpenGLESTextureCacheRef has base type: struct __CVOpenGLESTextureCache *
type OpenGLESTextureCacheRef uintptr
// OpenGLTextureRef - A reference to an OpenGL texture-based image buffer object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTexture
// CVOpenGLTextureRef has base type: CVImageBufferRef
type OpenGLTextureRef uintptr
// OpenGLTextureCacheRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCache
// CVOpenGLTextureCacheRef has base type: struct __CVOpenGLTextureCache *
type OpenGLTextureCacheRef uintptr
// OptionFlags - The flags to be used for the display link output callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOptionFlags
// CVOptionFlags has base type: uint64_t
type OptionFlags uintptr
// PixelBufferRef - A reference to a Core Video pixel buffer object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
// CVPixelBufferRef has base type: CVImageBufferRef
type PixelBufferRef uintptr
// PixelBufferPoolRef - A reference to a pixel buffer pool object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPool
// CVPixelBufferPoolRef has base type: struct __CVPixelBufferPool *
type PixelBufferPoolRef uintptr
// PixelBufferReleaseBytesCallback - A type that defines a release callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferReleaseBytesCallback
// CVPixelBufferReleaseBytesCallback is a callback function
// C type: void (*)(void *, const void *)
type PixelBufferReleaseBytesCallback = func(unsafe.Pointer, unsafe.Pointer)
// PixelBufferReleasePlanarBytesCallback - Defines a pointer to a pixel buffer release callback function, which is called when a pixel buffer created by   is released.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferReleasePlanarBytesCallback
// CVPixelBufferReleasePlanarBytesCallback is a callback function
// C type: void (*)(void *, const void *, unsigned long, unsigned long, const void **)
type PixelBufferReleasePlanarBytesCallback = func(unsafe.Pointer, unsafe.Pointer, uint, uint, unsafe.Pointer)
// Return - A Core Video error type return value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVReturn
// CVReturn has base type: int32_t
type Return uintptr

