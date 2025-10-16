
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [makeKeyWindow] class.
var makeKeyWindowClass _makeKeyWindowClass

func init() {
	makeKeyWindowClass = _makeKeyWindowClass{objc.GetClass("makeKeyWindow")}
}

type _makeKeyWindowClass struct {
	objc.Class
}

// An interface definition for the [makeKeyWindow] class.
type ImakeKeyWindow interface {
	ID() objc.ID
}

type makeKeyWindow struct {
	id objc.ID
}

func makeKeyWindowFrom(ptr unsafe.Pointer) makeKeyWindow {
	return makeKeyWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ makeKeyWindow) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _makeKeyWindowClass) Alloc() makeKeyWindow {
	rv := objc.Send[makeKeyWindow](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _makeKeyWindowClass) New() makeKeyWindow {
	rv := objc.Send[makeKeyWindow](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmakeKeyWindow creates and returns a new initialized instance.
func NewmakeKeyWindow() makeKeyWindow {
	return makeKeyWindowClass.New()
}

// Init initializes the instance.
func (m_ makeKeyWindow) Init() makeKeyWindow {
	rv := objc.Send[makeKeyWindow](m_.ID(), selInit)
	return rv
}
