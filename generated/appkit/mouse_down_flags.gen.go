
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [mouseDownFlags] class.
var mouseDownFlagsClass _mouseDownFlagsClass

func init() {
	mouseDownFlagsClass = _mouseDownFlagsClass{objc.GetClass("mouseDownFlags")}
}

type _mouseDownFlagsClass struct {
	objc.Class
}

// An interface definition for the [mouseDownFlags] class.
type ImouseDownFlags interface {
	ID() objc.ID
}

type mouseDownFlags struct {
	id objc.ID
}

func mouseDownFlagsFrom(ptr unsafe.Pointer) mouseDownFlags {
	return mouseDownFlags{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ mouseDownFlags) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _mouseDownFlagsClass) Alloc() mouseDownFlags {
	rv := objc.Send[mouseDownFlags](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _mouseDownFlagsClass) New() mouseDownFlags {
	rv := objc.Send[mouseDownFlags](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmouseDownFlags creates and returns a new initialized instance.
func NewmouseDownFlags() mouseDownFlags {
	return mouseDownFlagsClass.New()
}

// Init initializes the instance.
func (m_ mouseDownFlags) Init() mouseDownFlags {
	rv := objc.Send[mouseDownFlags](m_.ID(), selInit)
	return rv
}
