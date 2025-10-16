
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [wantsBestResolutionOpenGLSurface] class.
var wantsBestResolutionOpenGLSurfaceClass _wantsBestResolutionOpenGLSurfaceClass

func init() {
	wantsBestResolutionOpenGLSurfaceClass = _wantsBestResolutionOpenGLSurfaceClass{objc.GetClass("wantsBestResolutionOpenGLSurface")}
}

type _wantsBestResolutionOpenGLSurfaceClass struct {
	objc.Class
}

// An interface definition for the [wantsBestResolutionOpenGLSurface] class.
type IwantsBestResolutionOpenGLSurface interface {
	ID() objc.ID
}

type wantsBestResolutionOpenGLSurface struct {
	id objc.ID
}

func wantsBestResolutionOpenGLSurfaceFrom(ptr unsafe.Pointer) wantsBestResolutionOpenGLSurface {
	return wantsBestResolutionOpenGLSurface{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ wantsBestResolutionOpenGLSurface) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _wantsBestResolutionOpenGLSurfaceClass) Alloc() wantsBestResolutionOpenGLSurface {
	rv := objc.Send[wantsBestResolutionOpenGLSurface](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _wantsBestResolutionOpenGLSurfaceClass) New() wantsBestResolutionOpenGLSurface {
	rv := objc.Send[wantsBestResolutionOpenGLSurface](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwantsBestResolutionOpenGLSurface creates and returns a new initialized instance.
func NewwantsBestResolutionOpenGLSurface() wantsBestResolutionOpenGLSurface {
	return wantsBestResolutionOpenGLSurfaceClass.New()
}

// Init initializes the instance.
func (w_ wantsBestResolutionOpenGLSurface) Init() wantsBestResolutionOpenGLSurface {
	rv := objc.Send[wantsBestResolutionOpenGLSurface](w_.ID(), selInit)
	return rv
}
