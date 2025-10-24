// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OpenGLPixelBuffer] class.
var (
	OpenGLPixelBufferClass     _OpenGLPixelBufferClass
	OpenGLPixelBufferClassOnce sync.Once
)

func getOpenGLPixelBufferClass() _OpenGLPixelBufferClass {
	OpenGLPixelBufferClassOnce.Do(func() {
		OpenGLPixelBufferClass = _OpenGLPixelBufferClass{objc.GetClass("NSOpenGLPixelBuffer")}
	})
	return OpenGLPixelBufferClass
}

type _OpenGLPixelBufferClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLPixelBuffer] class.
type IOpenGLPixelBuffer interface {
	objectivec.IObject
	// properties:
	CGLPBufferObj() LPBufferObj /* not a class type */
	PixelsHigh() unsafe.Pointer
	PixelsWide() unsafe.Pointer
	TextureInternalFormat() unsafe.Pointer
	TextureMaxMipMapLevel() unsafe.Pointer
	TextureTarget() unsafe.Pointer
	// methods:
}

// An object that provides access to accelerated offscreen rendering.
//
// Using offscreen rendering you could, for example, draw into the pixel buffer, then use the contents as a texture map elsewhere. Typically you initialize an object using the method and attach the resulting object to an OpenGL context with the method of . Every object wraps a low-level, platform-specific Core OpenGL (CGL) pixel buffer object. Your application can retrieve the CGL pixel buffer by calling the method. For more information on the underling CGL pixel buffer, see .


// An object that provides access to accelerated offscreen rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelBuffer
type OpenGLPixelBuffer struct {
	objectivec.Object
}

// OpenGLPixelBufferFrom constructs a [OpenGLPixelBuffer] from an unsafe.Pointer.
//
// An object that provides access to accelerated offscreen rendering.
func OpenGLPixelBufferFrom(ptr unsafe.Pointer) OpenGLPixelBuffer {
	return OpenGLPixelBuffer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OpenGLPixelBufferClass) Alloc() OpenGLPixelBuffer {
	rv := objc.Send[OpenGLPixelBuffer](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OpenGLPixelBufferClass) New() OpenGLPixelBuffer {
	rv := objc.Send[OpenGLPixelBuffer](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenGLPixelBuffer) Init() OpenGLPixelBuffer {
	rv := objc.Send[OpenGLPixelBuffer](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenGLPixelBuffer) Autorelease() OpenGLPixelBuffer {
	rv := objc.Send[OpenGLPixelBuffer](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenGLPixelBuffer creates a new OpenGLPixelBuffer instance.
func NewOpenGLPixelBuffer() OpenGLPixelBuffer {
	return getOpenGLPixelBufferClass().New()
}



// Initializes and returns an OpenGL pixel buffer object that encapsulates an existing CGL pixel buffer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelBuffer/initWithCGLPBufferObj:
func NewOpenGLPixelBufferWithCGLPBufferObj(pbuffer LPBufferObj /* not a class type */) OpenGLPixelBuffer {
	instance := getOpenGLPixelBufferClass().Alloc()
	rv := objc.Send[OpenGLPixelBuffer](instance.ID, objc.Sel("initWithCGLPBufferObj:"), pbuffer)
	rv.Autorelease()
	return rv
}


// Returns an object initialized with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelBuffer/initWithTextureTarget:textureInternalFormat:textureMaxMipMapLevel:pixelsWide:pixelsHigh:
func NewOpenGLPixelBufferWithTextureTargetTextureInternalFormatTextureMaxMipMapLevelPixelsWidePixelsHigh(target unsafe.Pointer, format unsafe.Pointer, maxLevel unsafe.Pointer, pixelsWide unsafe.Pointer, pixelsHigh unsafe.Pointer) OpenGLPixelBuffer {
	instance := getOpenGLPixelBufferClass().Alloc()
	rv := objc.Send[OpenGLPixelBuffer](instance.ID, objc.Sel("initWithTextureTarget:textureInternalFormat:textureMaxMipMapLevel:pixelsWide:pixelsHigh:"), target, format, maxLevel, pixelsWide, pixelsHigh)
	rv.Autorelease()
	return rv
}



// The underlying CGL pixel buffer object associated with the OpenGL pixel buffer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelBuffer/CGLPBufferObj
func (o_ OpenGLPixelBuffer) CGLPBufferObj() LPBufferObj /* not a class type */ {
	rv := objc.Send[LPBufferObj](o_.ID, objc.Sel("CGLPBufferObj"))
	return rv
}


// The height of the OpenGL pixel buffer’s texture (in pixels).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelBuffer/pixelsHigh
func (o_ OpenGLPixelBuffer) PixelsHigh() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("pixelsHigh"))
	return rv
}


// The width of the OpenGL pixel buffer’s texture, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelBuffer/pixelsWide
func (o_ OpenGLPixelBuffer) PixelsWide() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("pixelsWide"))
	return rv
}


// The internal format of the OpenGL pixel buffer’s texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelBuffer/textureInternalFormat
func (o_ OpenGLPixelBuffer) TextureInternalFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("textureInternalFormat"))
	return rv
}


// The maximum mipmap level of the OpenGL pixel buffer’s texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelBuffer/textureMaxMipMapLevel
func (o_ OpenGLPixelBuffer) TextureMaxMipMapLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("textureMaxMipMapLevel"))
	return rv
}


// The texture target of the OpenGL pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelBuffer/textureTarget
func (o_ OpenGLPixelBuffer) TextureTarget() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("textureTarget"))
	return rv
}


