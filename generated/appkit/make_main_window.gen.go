
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [makeMainWindow] class.
var makeMainWindowClass _makeMainWindowClass

func init() {
	makeMainWindowClass = _makeMainWindowClass{objc.GetClass("makeMainWindow")}
}

type _makeMainWindowClass struct {
	objc.Class
}

// An interface definition for the [makeMainWindow] class.
type ImakeMainWindow interface {
	ID() objc.ID
}

type makeMainWindow struct {
	id objc.ID
}

func makeMainWindowFrom(ptr unsafe.Pointer) makeMainWindow {
	return makeMainWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ makeMainWindow) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _makeMainWindowClass) Alloc() makeMainWindow {
	rv := objc.Send[makeMainWindow](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _makeMainWindowClass) New() makeMainWindow {
	rv := objc.Send[makeMainWindow](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmakeMainWindow creates and returns a new initialized instance.
func NewmakeMainWindow() makeMainWindow {
	return makeMainWindowClass.New()
}

// Init initializes the instance.
func (m_ makeMainWindow) Init() makeMainWindow {
	rv := objc.Send[makeMainWindow](m_.ID(), selInit)
	return rv
}
