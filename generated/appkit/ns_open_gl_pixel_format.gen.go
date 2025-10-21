// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OpenGLPixelFormat] class.
var (
	OpenGLPixelFormatClass     _OpenGLPixelFormatClass
	OpenGLPixelFormatClassOnce sync.Once
)

func getOpenGLPixelFormatClass() _OpenGLPixelFormatClass {
	OpenGLPixelFormatClassOnce.Do(func() {
		OpenGLPixelFormatClass = _OpenGLPixelFormatClass{objc.GetClass("NSOpenGLPixelFormat")}
	})
	return OpenGLPixelFormatClass
}

type _OpenGLPixelFormatClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLPixelFormat] class.
type IOpenGLPixelFormat interface {
	objectivec.IObject
}

// An object that specifies the types of buffers and other attributes of the OpenGL context.
//
// To render with OpenGL into an , you must specify the context’s pixel format. Every object wraps a low-level, platform-specific Core OpenGL (CGL) pixel format object. Your application can retrieve the CGL pixel format object by calling the method. For more information on the underling CGL pixel format object, see .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelFormat
type OpenGLPixelFormat struct {
	objectivec.Object
}

// OpenGLPixelFormatFrom constructs a [OpenGLPixelFormat] from an unsafe.Pointer.
//
// An object that specifies the types of buffers and other attributes of the OpenGL context.
func OpenGLPixelFormatFrom(ptr unsafe.Pointer) OpenGLPixelFormat {
	return OpenGLPixelFormat{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OpenGLPixelFormatClass) Alloc() OpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OpenGLPixelFormatClass) New() OpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenGLPixelFormat) Init() OpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenGLPixelFormat) Autorelease() OpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenGLPixelFormat creates a new OpenGLPixelFormat instance.
func NewOpenGLPixelFormat() OpenGLPixelFormat {
	return getOpenGLPixelFormatClass().New()
}

// The low-level, platform-specific Core OpenGL (CGL) pixel format object represented by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelFormat/cglPixelFormatObj
func (o_ OpenGLPixelFormat) CGLPixelFormatObj() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("CGLPixelFormatObj"))
	return rv
}
