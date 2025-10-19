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
	openGLPixelFormatClass     _OpenGLPixelFormatClass
	openGLPixelFormatClassOnce sync.Once
)

func getOpenGLPixelFormatClass() _OpenGLPixelFormatClass {
	openGLPixelFormatClassOnce.Do(func() {
		openGLPixelFormatClass = _OpenGLPixelFormatClass{objc.GetClass("NSOpenGLPixelFormat")}
	})
	return openGLPixelFormatClass
}

type _OpenGLPixelFormatClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLPixelFormat] class.
type IOpenGLPixelFormat interface {
	objectivec.IObject
}

// An object that specifies the types of buffers and other attributes of the OpenGL context. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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




