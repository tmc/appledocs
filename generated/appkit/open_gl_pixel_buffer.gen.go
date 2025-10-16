
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenGLPixelBuffer] class.
var OpenGLPixelBufferClass _OpenGLPixelBufferClass

func init() {
	OpenGLPixelBufferClass = _OpenGLPixelBufferClass{objc.GetClass("NSOpenGLPixelBuffer")}
}

type _OpenGLPixelBufferClass struct {
	objc.Class
}

// An interface definition for the [OpenGLPixelBuffer] class.
type IOpenGLPixelBuffer interface {
	ID() objc.ID
}

type OpenGLPixelBuffer struct {
	id objc.ID
}

func OpenGLPixelBufferFrom(ptr unsafe.Pointer) OpenGLPixelBuffer {
	return OpenGLPixelBuffer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ OpenGLPixelBuffer) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _OpenGLPixelBufferClass) Alloc() OpenGLPixelBuffer {
	rv := objc.Send[OpenGLPixelBuffer](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _OpenGLPixelBufferClass) New() OpenGLPixelBuffer {
	rv := objc.Send[OpenGLPixelBuffer](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewOpenGLPixelBuffer creates and returns a new initialized instance.
func NewOpenGLPixelBuffer() OpenGLPixelBuffer {
	return OpenGLPixelBufferClass.New()
}

// Init initializes the instance.
func (o_ OpenGLPixelBuffer) Init() OpenGLPixelBuffer {
	rv := objc.Send[OpenGLPixelBuffer](o_.ID(), selInit)
	return rv
}
