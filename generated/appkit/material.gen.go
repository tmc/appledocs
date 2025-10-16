
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [material] class.
var materialClass _materialClass

func init() {
	materialClass = _materialClass{objc.GetClass("material")}
}

type _materialClass struct {
	objc.Class
}

// An interface definition for the [material] class.
type Imaterial interface {
	ID() objc.ID
}

type material struct {
	id objc.ID
}

func materialFrom(ptr unsafe.Pointer) material {
	return material{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ material) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _materialClass) Alloc() material {
	rv := objc.Send[material](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _materialClass) New() material {
	rv := objc.Send[material](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newmaterial creates and returns a new initialized instance.
func Newmaterial() material {
	return materialClass.New()
}

// Init initializes the instance.
func (m_ material) Init() material {
	rv := objc.Send[material](m_.ID(), selInit)
	return rv
}
