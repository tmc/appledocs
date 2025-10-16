
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [modifierFlags] class.
var modifierFlagsClass _modifierFlagsClass

func init() {
	modifierFlagsClass = _modifierFlagsClass{objc.GetClass("modifierFlags")}
}

type _modifierFlagsClass struct {
	objc.Class
}

// An interface definition for the [modifierFlags] class.
type ImodifierFlags interface {
	ID() objc.ID
}

type modifierFlags struct {
	id objc.ID
}

func modifierFlagsFrom(ptr unsafe.Pointer) modifierFlags {
	return modifierFlags{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ modifierFlags) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _modifierFlagsClass) Alloc() modifierFlags {
	rv := objc.Send[modifierFlags](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _modifierFlagsClass) New() modifierFlags {
	rv := objc.Send[modifierFlags](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmodifierFlags creates and returns a new initialized instance.
func NewmodifierFlags() modifierFlags {
	return modifierFlagsClass.New()
}

// Init initializes the instance.
func (m_ modifierFlags) Init() modifierFlags {
	rv := objc.Send[modifierFlags](m_.ID(), selInit)
	return rv
}
