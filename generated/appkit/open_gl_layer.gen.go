
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenGLLayer] class.
var OpenGLLayerClass _OpenGLLayerClass

func init() {
	OpenGLLayerClass = _OpenGLLayerClass{objc.GetClass("NSOpenGLLayer")}
}

type _OpenGLLayerClass struct {
	objc.Class
}

// An interface definition for the [OpenGLLayer] class.
type IOpenGLLayer interface {
	ID() objc.ID
}

type OpenGLLayer struct {
	id objc.ID
}

func OpenGLLayerFrom(ptr unsafe.Pointer) OpenGLLayer {
	return OpenGLLayer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ OpenGLLayer) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _OpenGLLayerClass) Alloc() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _OpenGLLayerClass) New() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewOpenGLLayer creates and returns a new initialized instance.
func NewOpenGLLayer() OpenGLLayer {
	return OpenGLLayerClass.New()
}

// Init initializes the instance.
func (o_ OpenGLLayer) Init() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](o_.ID(), selInit)
	return rv
}
