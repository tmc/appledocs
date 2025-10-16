
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [becomeMainWindow] class.
var becomeMainWindowClass _becomeMainWindowClass

func init() {
	becomeMainWindowClass = _becomeMainWindowClass{objc.GetClass("becomeMainWindow")}
}

type _becomeMainWindowClass struct {
	objc.Class
}

// An interface definition for the [becomeMainWindow] class.
type IbecomeMainWindow interface {
	ID() objc.ID
}

type becomeMainWindow struct {
	id objc.ID
}

func becomeMainWindowFrom(ptr unsafe.Pointer) becomeMainWindow {
	return becomeMainWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ becomeMainWindow) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _becomeMainWindowClass) Alloc() becomeMainWindow {
	rv := objc.Send[becomeMainWindow](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _becomeMainWindowClass) New() becomeMainWindow {
	rv := objc.Send[becomeMainWindow](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbecomeMainWindow creates and returns a new initialized instance.
func NewbecomeMainWindow() becomeMainWindow {
	return becomeMainWindowClass.New()
}

// Init initializes the instance.
func (b_ becomeMainWindow) Init() becomeMainWindow {
	rv := objc.Send[becomeMainWindow](b_.ID(), selInit)
	return rv
}
