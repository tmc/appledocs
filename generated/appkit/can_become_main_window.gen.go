
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canBecomeMainWindow] class.
var canBecomeMainWindowClass _canBecomeMainWindowClass

func init() {
	canBecomeMainWindowClass = _canBecomeMainWindowClass{objc.GetClass("canBecomeMainWindow")}
}

type _canBecomeMainWindowClass struct {
	objc.Class
}

// An interface definition for the [canBecomeMainWindow] class.
type IcanBecomeMainWindow interface {
	ID() objc.ID
}

type canBecomeMainWindow struct {
	id objc.ID
}

func canBecomeMainWindowFrom(ptr unsafe.Pointer) canBecomeMainWindow {
	return canBecomeMainWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canBecomeMainWindow) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canBecomeMainWindowClass) Alloc() canBecomeMainWindow {
	rv := objc.Send[canBecomeMainWindow](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canBecomeMainWindowClass) New() canBecomeMainWindow {
	rv := objc.Send[canBecomeMainWindow](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanBecomeMainWindow creates and returns a new initialized instance.
func NewcanBecomeMainWindow() canBecomeMainWindow {
	return canBecomeMainWindowClass.New()
}

// Init initializes the instance.
func (c_ canBecomeMainWindow) Init() canBecomeMainWindow {
	rv := objc.Send[canBecomeMainWindow](c_.ID(), selInit)
	return rv
}
