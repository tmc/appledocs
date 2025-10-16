
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [frameCenterRotation] class.
var frameCenterRotationClass _frameCenterRotationClass

func init() {
	frameCenterRotationClass = _frameCenterRotationClass{objc.GetClass("frameCenterRotation")}
}

type _frameCenterRotationClass struct {
	objc.Class
}

// An interface definition for the [frameCenterRotation] class.
type IframeCenterRotation interface {
	ID() objc.ID
}

type frameCenterRotation struct {
	id objc.ID
}

func frameCenterRotationFrom(ptr unsafe.Pointer) frameCenterRotation {
	return frameCenterRotation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ frameCenterRotation) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _frameCenterRotationClass) Alloc() frameCenterRotation {
	rv := objc.Send[frameCenterRotation](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _frameCenterRotationClass) New() frameCenterRotation {
	rv := objc.Send[frameCenterRotation](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewframeCenterRotation creates and returns a new initialized instance.
func NewframeCenterRotation() frameCenterRotation {
	return frameCenterRotationClass.New()
}

// Init initializes the instance.
func (f_ frameCenterRotation) Init() frameCenterRotation {
	rv := objc.Send[frameCenterRotation](f_.ID(), selInit)
	return rv
}
