// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo

// Type aliases and typedefs
// CVBufferRef - A reference to a Core Video buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVBuffer
// CVBufferRef has base type: struct __CVBuffer *
type CVBufferRef uintptr
// CVDisplayLinkRef - A reference to a display link object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLink
// CVDisplayLinkRef has base type: struct __CVDisplayLink *
type CVDisplayLinkRef uintptr
// CVDisplayLinkOutputCallback - A type for a display link callback function that the system invokes when it’s time for the app to output a video frame.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkOutputCallback
// CVDisplayLinkOutputCallback has base type: int (*)(struct __CVDisplayLink *, const CVTimeStamp *, const CVTimeStamp *, unsigned long long, unsigned long long *, void *)
type CVDisplayLinkOutputCallback uintptr
// CVEAGLContext - A type that resolves to an   pointer when appropriate.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVEAGLContext
// CVEAGLContext has base type: void *
type CVEAGLContext uintptr
// CVImageBufferRef - A reference to a Core Video image buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVImageBuffer
// CVImageBufferRef has base type: CVBufferRef
type CVImageBufferRef uintptr
// CVMetalBufferRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBuffer
// CVMetalBufferRef has base type: CVBufferRef
type CVMetalBufferRef uintptr
// CVMetalBufferCacheRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCache
// CVMetalBufferCacheRef has base type: struct __CVMetalBufferCache *
type CVMetalBufferCacheRef uintptr
// CVMetalTextureRef - A reference to a CoreVideo Metal texture-based image buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTexture
// CVMetalTextureRef has base type: CVImageBufferRef
type CVMetalTextureRef uintptr
// CVMetalTextureCacheRef - A reference to a Core Video Metal texture cache.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCache
// CVMetalTextureCacheRef has base type: struct __CVMetalTextureCache *
type CVMetalTextureCacheRef uintptr
// CVOpenGLBufferRef - A reference to a Core Video OpenGL buffer object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBuffer
// CVOpenGLBufferRef has base type: CVImageBufferRef
type CVOpenGLBufferRef uintptr
// CVOpenGLBufferPoolRef - A reference to an OpenGL buffer pool object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPool
// CVOpenGLBufferPoolRef has base type: struct __CVOpenGLBufferPool *
type CVOpenGLBufferPoolRef uintptr
// CVOpenGLESTextureRef - A reference to a Core Video texture-based image buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTexture
// CVOpenGLESTextureRef has base type: CVImageBufferRef
type CVOpenGLESTextureRef uintptr
// CVOpenGLESTextureCacheRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLESTextureCache
// CVOpenGLESTextureCacheRef has base type: struct __CVOpenGLESTextureCache *
type CVOpenGLESTextureCacheRef uintptr
// CVOpenGLTextureRef - A reference to an OpenGL texture-based image buffer object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTexture
// CVOpenGLTextureRef has base type: CVImageBufferRef
type CVOpenGLTextureRef uintptr
// CVOpenGLTextureCacheRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCache
// CVOpenGLTextureCacheRef has base type: struct __CVOpenGLTextureCache *
type CVOpenGLTextureCacheRef uintptr
// CVOptionFlags - The flags to be used for the display link output callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVOptionFlags
// CVOptionFlags has base type: uint64_t
type CVOptionFlags uintptr
// CVPixelBufferRef - A reference to a Core Video pixel buffer object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
// CVPixelBufferRef has base type: CVImageBufferRef
type CVPixelBufferRef uintptr
// CVPixelBufferPoolRef - A reference to a pixel buffer pool object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPool
// CVPixelBufferPoolRef has base type: struct __CVPixelBufferPool *
type CVPixelBufferPoolRef uintptr
// CVPixelBufferReleaseBytesCallback - A type that defines a release callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferReleaseBytesCallback
// CVPixelBufferReleaseBytesCallback has base type: void (*)(void *, const void *)
type CVPixelBufferReleaseBytesCallback uintptr
// CVPixelBufferReleasePlanarBytesCallback - Defines a pointer to a pixel buffer release callback function, which is called when a pixel buffer created by   is released.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferReleasePlanarBytesCallback
// CVPixelBufferReleasePlanarBytesCallback has base type: void (*)(void *, const void *, unsigned long, unsigned long, const void **)
type CVPixelBufferReleasePlanarBytesCallback uintptr
// CVReturn - A Core Video error type return value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVReturn
// CVReturn has base type: int32_t
type CVReturn uintptr

