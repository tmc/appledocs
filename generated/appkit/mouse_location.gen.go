
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [mouseLocation] class.
var mouseLocationClass _mouseLocationClass

func init() {
	mouseLocationClass = _mouseLocationClass{objc.GetClass("mouseLocation")}
}

type _mouseLocationClass struct {
	objc.Class
}

// An interface definition for the [mouseLocation] class.
type ImouseLocation interface {
	ID() objc.ID
}

type mouseLocation struct {
	id objc.ID
}

func mouseLocationFrom(ptr unsafe.Pointer) mouseLocation {
	return mouseLocation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ mouseLocation) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _mouseLocationClass) Alloc() mouseLocation {
	rv := objc.Send[mouseLocation](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _mouseLocationClass) New() mouseLocation {
	rv := objc.Send[mouseLocation](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmouseLocation creates and returns a new initialized instance.
func NewmouseLocation() mouseLocation {
	return mouseLocationClass.New()
}

// Init initializes the instance.
func (m_ mouseLocation) Init() mouseLocation {
	rv := objc.Send[mouseLocation](m_.ID(), selInit)
	return rv
}
