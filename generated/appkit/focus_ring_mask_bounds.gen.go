
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [focusRingMaskBounds] class.
var focusRingMaskBoundsClass _focusRingMaskBoundsClass

func init() {
	focusRingMaskBoundsClass = _focusRingMaskBoundsClass{objc.GetClass("focusRingMaskBounds")}
}

type _focusRingMaskBoundsClass struct {
	objc.Class
}

// An interface definition for the [focusRingMaskBounds] class.
type IfocusRingMaskBounds interface {
	ID() objc.ID
}

type focusRingMaskBounds struct {
	id objc.ID
}

func focusRingMaskBoundsFrom(ptr unsafe.Pointer) focusRingMaskBounds {
	return focusRingMaskBounds{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ focusRingMaskBounds) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _focusRingMaskBoundsClass) Alloc() focusRingMaskBounds {
	rv := objc.Send[focusRingMaskBounds](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _focusRingMaskBoundsClass) New() focusRingMaskBounds {
	rv := objc.Send[focusRingMaskBounds](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfocusRingMaskBounds creates and returns a new initialized instance.
func NewfocusRingMaskBounds() focusRingMaskBounds {
	return focusRingMaskBoundsClass.New()
}

// Init initializes the instance.
func (f_ focusRingMaskBounds) Init() focusRingMaskBounds {
	rv := objc.Send[focusRingMaskBounds](f_.ID(), selInit)
	return rv
}
