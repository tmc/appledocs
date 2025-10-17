
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Gradient] class.
var GradientClass _GradientClass

func init() {
	GradientClass = _GradientClass{objc.GetClass("NSGradient")}
}

type _GradientClass struct {
	objc.Class
}

// An interface definition for the [Gradient] class.
type IGradient interface {
	ID() objc.ID
}

type Gradient struct {
	id objc.ID
}

func GradientFrom(ptr unsafe.Pointer) Gradient {
	return Gradient{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ Gradient) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _GradientClass) Alloc() Gradient {
	rv := objc.Send[Gradient](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _GradientClass) New() Gradient {
	rv := objc.Send[Gradient](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewGradient creates and returns a new initialized instance.
func NewGradient() Gradient {
	return GradientClass.New()
}

// Init initializes the instance.
func (g_ Gradient) Init() Gradient {
	rv := objc.Send[Gradient](g_.ID(), selInit)
	return rv
}
