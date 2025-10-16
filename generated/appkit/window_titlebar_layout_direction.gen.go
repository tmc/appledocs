
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [windowTitlebarLayoutDirection] class.
var windowTitlebarLayoutDirectionClass _windowTitlebarLayoutDirectionClass

func init() {
	windowTitlebarLayoutDirectionClass = _windowTitlebarLayoutDirectionClass{objc.GetClass("windowTitlebarLayoutDirection")}
}

type _windowTitlebarLayoutDirectionClass struct {
	objc.Class
}

// An interface definition for the [windowTitlebarLayoutDirection] class.
type IwindowTitlebarLayoutDirection interface {
	ID() objc.ID
}

type windowTitlebarLayoutDirection struct {
	id objc.ID
}

func windowTitlebarLayoutDirectionFrom(ptr unsafe.Pointer) windowTitlebarLayoutDirection {
	return windowTitlebarLayoutDirection{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ windowTitlebarLayoutDirection) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _windowTitlebarLayoutDirectionClass) Alloc() windowTitlebarLayoutDirection {
	rv := objc.Send[windowTitlebarLayoutDirection](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _windowTitlebarLayoutDirectionClass) New() windowTitlebarLayoutDirection {
	rv := objc.Send[windowTitlebarLayoutDirection](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwindowTitlebarLayoutDirection creates and returns a new initialized instance.
func NewwindowTitlebarLayoutDirection() windowTitlebarLayoutDirection {
	return windowTitlebarLayoutDirectionClass.New()
}

// Init initializes the instance.
func (w_ windowTitlebarLayoutDirection) Init() windowTitlebarLayoutDirection {
	rv := objc.Send[windowTitlebarLayoutDirection](w_.ID(), selInit)
	return rv
}
