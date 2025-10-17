
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenGLPixelFormat] class.
var OpenGLPixelFormatClass _OpenGLPixelFormatClass

func init() {
	OpenGLPixelFormatClass = _OpenGLPixelFormatClass{objc.GetClass("NSOpenGLPixelFormat")}
}

type _OpenGLPixelFormatClass struct {
	objc.Class
}

// An interface definition for the [OpenGLPixelFormat] class.
type IOpenGLPixelFormat interface {
	ID() objc.ID
}

type OpenGLPixelFormat struct {
	id objc.ID
}

func OpenGLPixelFormatFrom(ptr unsafe.Pointer) OpenGLPixelFormat {
	return OpenGLPixelFormat{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ OpenGLPixelFormat) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _OpenGLPixelFormatClass) Alloc() OpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _OpenGLPixelFormatClass) New() OpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewOpenGLPixelFormat creates and returns a new initialized instance.
func NewOpenGLPixelFormat() OpenGLPixelFormat {
	return OpenGLPixelFormatClass.New()
}

// Init initializes the instance.
func (o_ OpenGLPixelFormat) Init() OpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](o_.ID(), selInit)
	return rv
}
