
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [boundsRotation] class.
var boundsRotationClass _boundsRotationClass

func init() {
	boundsRotationClass = _boundsRotationClass{objc.GetClass("boundsRotation")}
}

type _boundsRotationClass struct {
	objc.Class
}

// An interface definition for the [boundsRotation] class.
type IboundsRotation interface {
	ID() objc.ID
}

type boundsRotation struct {
	id objc.ID
}

func boundsRotationFrom(ptr unsafe.Pointer) boundsRotation {
	return boundsRotation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ boundsRotation) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _boundsRotationClass) Alloc() boundsRotation {
	rv := objc.Send[boundsRotation](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _boundsRotationClass) New() boundsRotation {
	rv := objc.Send[boundsRotation](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewboundsRotation creates and returns a new initialized instance.
func NewboundsRotation() boundsRotation {
	return boundsRotationClass.New()
}

// Init initializes the instance.
func (b_ boundsRotation) Init() boundsRotation {
	rv := objc.Send[boundsRotation](b_.ID(), selInit)
	return rv
}
