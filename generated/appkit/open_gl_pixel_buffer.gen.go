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
	openGLPixelBufferClass     _OpenGLPixelBufferClass
	openGLPixelBufferClassOnce sync.Once
)

func getOpenGLPixelBufferClass() _OpenGLPixelBufferClass {
	openGLPixelBufferClassOnce.Do(func() {
		openGLPixelBufferClass = _OpenGLPixelBufferClass{objc.GetClass("NSOpenGLPixelBuffer")}
	})
	return openGLPixelBufferClass
}

type _OpenGLPixelBufferClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLPixelBuffer] class.
type IOpenGLPixelBuffer interface {
	objectivec.IObject
}

// An object that provides access to accelerated offscreen rendering.
//
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




