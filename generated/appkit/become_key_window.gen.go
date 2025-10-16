
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [becomeKeyWindow] class.
var becomeKeyWindowClass _becomeKeyWindowClass

func init() {
	becomeKeyWindowClass = _becomeKeyWindowClass{objc.GetClass("becomeKeyWindow")}
}

type _becomeKeyWindowClass struct {
	objc.Class
}

// An interface definition for the [becomeKeyWindow] class.
type IbecomeKeyWindow interface {
	ID() objc.ID
}

type becomeKeyWindow struct {
	id objc.ID
}

func becomeKeyWindowFrom(ptr unsafe.Pointer) becomeKeyWindow {
	return becomeKeyWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ becomeKeyWindow) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _becomeKeyWindowClass) Alloc() becomeKeyWindow {
	rv := objc.Send[becomeKeyWindow](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _becomeKeyWindowClass) New() becomeKeyWindow {
	rv := objc.Send[becomeKeyWindow](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbecomeKeyWindow creates and returns a new initialized instance.
func NewbecomeKeyWindow() becomeKeyWindow {
	return becomeKeyWindowClass.New()
}

// Init initializes the instance.
func (b_ becomeKeyWindow) Init() becomeKeyWindow {
	rv := objc.Send[becomeKeyWindow](b_.ID(), selInit)
	return rv
}
