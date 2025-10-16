
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [locationInWindow] class.
var locationInWindowClass _locationInWindowClass

func init() {
	locationInWindowClass = _locationInWindowClass{objc.GetClass("locationInWindow")}
}

type _locationInWindowClass struct {
	objc.Class
}

// An interface definition for the [locationInWindow] class.
type IlocationInWindow interface {
	ID() objc.ID
}

type locationInWindow struct {
	id objc.ID
}

func locationInWindowFrom(ptr unsafe.Pointer) locationInWindow {
	return locationInWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ locationInWindow) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _locationInWindowClass) Alloc() locationInWindow {
	rv := objc.Send[locationInWindow](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _locationInWindowClass) New() locationInWindow {
	rv := objc.Send[locationInWindow](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlocationInWindow creates and returns a new initialized instance.
func NewlocationInWindow() locationInWindow {
	return locationInWindowClass.New()
}

// Init initializes the instance.
func (l_ locationInWindow) Init() locationInWindow {
	rv := objc.Send[locationInWindow](l_.ID(), selInit)
	return rv
}
