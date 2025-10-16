
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [windowNumber] class.
var windowNumberClass _windowNumberClass

func init() {
	windowNumberClass = _windowNumberClass{objc.GetClass("windowNumber")}
}

type _windowNumberClass struct {
	objc.Class
}

// An interface definition for the [windowNumber] class.
type IwindowNumber interface {
	ID() objc.ID
}

type windowNumber struct {
	id objc.ID
}

func windowNumberFrom(ptr unsafe.Pointer) windowNumber {
	return windowNumber{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ windowNumber) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _windowNumberClass) Alloc() windowNumber {
	rv := objc.Send[windowNumber](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _windowNumberClass) New() windowNumber {
	rv := objc.Send[windowNumber](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwindowNumber creates and returns a new initialized instance.
func NewwindowNumber() windowNumber {
	return windowNumberClass.New()
}

// Init initializes the instance.
func (w_ windowNumber) Init() windowNumber {
	rv := objc.Send[windowNumber](w_.ID(), selInit)
	return rv
}
