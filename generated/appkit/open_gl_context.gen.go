
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenGLContext] class.
var OpenGLContextClass _OpenGLContextClass

func init() {
	OpenGLContextClass = _OpenGLContextClass{objc.GetClass("NSOpenGLContext")}
}

type _OpenGLContextClass struct {
	objc.Class
}

// An interface definition for the [OpenGLContext] class.
type IOpenGLContext interface {
	ID() objc.ID
}

type OpenGLContext struct {
	id objc.ID
}

func OpenGLContextFrom(ptr unsafe.Pointer) OpenGLContext {
	return OpenGLContext{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ OpenGLContext) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _OpenGLContextClass) Alloc() OpenGLContext {
	rv := objc.Send[OpenGLContext](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _OpenGLContextClass) New() OpenGLContext {
	rv := objc.Send[OpenGLContext](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewOpenGLContext creates and returns a new initialized instance.
func NewOpenGLContext() OpenGLContext {
	return OpenGLContextClass.New()
}

// Init initializes the instance.
func (o_ OpenGLContext) Init() OpenGLContext {
	rv := objc.Send[OpenGLContext](o_.ID(), selInit)
	return rv
}
