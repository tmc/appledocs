
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [mouseDownCanMoveWindow] class.
var mouseDownCanMoveWindowClass _mouseDownCanMoveWindowClass

func init() {
	mouseDownCanMoveWindowClass = _mouseDownCanMoveWindowClass{objc.GetClass("mouseDownCanMoveWindow")}
}

type _mouseDownCanMoveWindowClass struct {
	objc.Class
}

// An interface definition for the [mouseDownCanMoveWindow] class.
type ImouseDownCanMoveWindow interface {
	ID() objc.ID
}

type mouseDownCanMoveWindow struct {
	id objc.ID
}

func mouseDownCanMoveWindowFrom(ptr unsafe.Pointer) mouseDownCanMoveWindow {
	return mouseDownCanMoveWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ mouseDownCanMoveWindow) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _mouseDownCanMoveWindowClass) Alloc() mouseDownCanMoveWindow {
	rv := objc.Send[mouseDownCanMoveWindow](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _mouseDownCanMoveWindowClass) New() mouseDownCanMoveWindow {
	rv := objc.Send[mouseDownCanMoveWindow](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmouseDownCanMoveWindow creates and returns a new initialized instance.
func NewmouseDownCanMoveWindow() mouseDownCanMoveWindow {
	return mouseDownCanMoveWindowClass.New()
}

// Init initializes the instance.
func (m_ mouseDownCanMoveWindow) Init() mouseDownCanMoveWindow {
	rv := objc.Send[mouseDownCanMoveWindow](m_.ID(), selInit)
	return rv
}
