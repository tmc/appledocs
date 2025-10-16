
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [windowRef] class.
var windowRefClass _windowRefClass

func init() {
	windowRefClass = _windowRefClass{objc.GetClass("windowRef")}
}

type _windowRefClass struct {
	objc.Class
}

// An interface definition for the [windowRef] class.
type IwindowRef interface {
	ID() objc.ID
}

type windowRef struct {
	id objc.ID
}

func windowRefFrom(ptr unsafe.Pointer) windowRef {
	return windowRef{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ windowRef) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _windowRefClass) Alloc() windowRef {
	rv := objc.Send[windowRef](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _windowRefClass) New() windowRef {
	rv := objc.Send[windowRef](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwindowRef creates and returns a new initialized instance.
func NewwindowRef() windowRef {
	return windowRefClass.New()
}

// Init initializes the instance.
func (w_ windowRef) Init() windowRef {
	rv := objc.Send[windowRef](w_.ID(), selInit)
	return rv
}
