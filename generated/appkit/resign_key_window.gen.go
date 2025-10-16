
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [resignKeyWindow] class.
var resignKeyWindowClass _resignKeyWindowClass

func init() {
	resignKeyWindowClass = _resignKeyWindowClass{objc.GetClass("resignKeyWindow")}
}

type _resignKeyWindowClass struct {
	objc.Class
}

// An interface definition for the [resignKeyWindow] class.
type IresignKeyWindow interface {
	ID() objc.ID
}

type resignKeyWindow struct {
	id objc.ID
}

func resignKeyWindowFrom(ptr unsafe.Pointer) resignKeyWindow {
	return resignKeyWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ resignKeyWindow) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _resignKeyWindowClass) Alloc() resignKeyWindow {
	rv := objc.Send[resignKeyWindow](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _resignKeyWindowClass) New() resignKeyWindow {
	rv := objc.Send[resignKeyWindow](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewresignKeyWindow creates and returns a new initialized instance.
func NewresignKeyWindow() resignKeyWindow {
	return resignKeyWindowClass.New()
}

// Init initializes the instance.
func (r_ resignKeyWindow) Init() resignKeyWindow {
	rv := objc.Send[resignKeyWindow](r_.ID(), selInit)
	return rv
}
