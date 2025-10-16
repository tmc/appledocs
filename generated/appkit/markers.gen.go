
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [markers] class.
var markersClass _markersClass

func init() {
	markersClass = _markersClass{objc.GetClass("markers")}
}

type _markersClass struct {
	objc.Class
}

// An interface definition for the [markers] class.
type Imarkers interface {
	ID() objc.ID
}

type markers struct {
	id objc.ID
}

func markersFrom(ptr unsafe.Pointer) markers {
	return markers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ markers) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _markersClass) Alloc() markers {
	rv := objc.Send[markers](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _markersClass) New() markers {
	rv := objc.Send[markers](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newmarkers creates and returns a new initialized instance.
func Newmarkers() markers {
	return markersClass.New()
}

// Init initializes the instance.
func (m_ markers) Init() markers {
	rv := objc.Send[markers](m_.ID(), selInit)
	return rv
}
