
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canBecomeKeyWindow] class.
var canBecomeKeyWindowClass _canBecomeKeyWindowClass

func init() {
	canBecomeKeyWindowClass = _canBecomeKeyWindowClass{objc.GetClass("canBecomeKeyWindow")}
}

type _canBecomeKeyWindowClass struct {
	objc.Class
}

// An interface definition for the [canBecomeKeyWindow] class.
type IcanBecomeKeyWindow interface {
	ID() objc.ID
}

type canBecomeKeyWindow struct {
	id objc.ID
}

func canBecomeKeyWindowFrom(ptr unsafe.Pointer) canBecomeKeyWindow {
	return canBecomeKeyWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canBecomeKeyWindow) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canBecomeKeyWindowClass) Alloc() canBecomeKeyWindow {
	rv := objc.Send[canBecomeKeyWindow](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canBecomeKeyWindowClass) New() canBecomeKeyWindow {
	rv := objc.Send[canBecomeKeyWindow](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanBecomeKeyWindow creates and returns a new initialized instance.
func NewcanBecomeKeyWindow() canBecomeKeyWindow {
	return canBecomeKeyWindowClass.New()
}

// Init initializes the instance.
func (c_ canBecomeKeyWindow) Init() canBecomeKeyWindow {
	rv := objc.Send[canBecomeKeyWindow](c_.ID(), selInit)
	return rv
}
