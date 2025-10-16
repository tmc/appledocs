
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [wantsExtendedDynamicRangeOpenGLSurface] class.
var wantsExtendedDynamicRangeOpenGLSurfaceClass _wantsExtendedDynamicRangeOpenGLSurfaceClass

func init() {
	wantsExtendedDynamicRangeOpenGLSurfaceClass = _wantsExtendedDynamicRangeOpenGLSurfaceClass{objc.GetClass("wantsExtendedDynamicRangeOpenGLSurface")}
}

type _wantsExtendedDynamicRangeOpenGLSurfaceClass struct {
	objc.Class
}

// An interface definition for the [wantsExtendedDynamicRangeOpenGLSurface] class.
type IwantsExtendedDynamicRangeOpenGLSurface interface {
	ID() objc.ID
}

type wantsExtendedDynamicRangeOpenGLSurface struct {
	id objc.ID
}

func wantsExtendedDynamicRangeOpenGLSurfaceFrom(ptr unsafe.Pointer) wantsExtendedDynamicRangeOpenGLSurface {
	return wantsExtendedDynamicRangeOpenGLSurface{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ wantsExtendedDynamicRangeOpenGLSurface) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _wantsExtendedDynamicRangeOpenGLSurfaceClass) Alloc() wantsExtendedDynamicRangeOpenGLSurface {
	rv := objc.Send[wantsExtendedDynamicRangeOpenGLSurface](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _wantsExtendedDynamicRangeOpenGLSurfaceClass) New() wantsExtendedDynamicRangeOpenGLSurface {
	rv := objc.Send[wantsExtendedDynamicRangeOpenGLSurface](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwantsExtendedDynamicRangeOpenGLSurface creates and returns a new initialized instance.
func NewwantsExtendedDynamicRangeOpenGLSurface() wantsExtendedDynamicRangeOpenGLSurface {
	return wantsExtendedDynamicRangeOpenGLSurfaceClass.New()
}

// Init initializes the instance.
func (w_ wantsExtendedDynamicRangeOpenGLSurface) Init() wantsExtendedDynamicRangeOpenGLSurface {
	rv := objc.Send[wantsExtendedDynamicRangeOpenGLSurface](w_.ID(), selInit)
	return rv
}
