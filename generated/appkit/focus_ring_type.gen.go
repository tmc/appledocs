
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [focusRingType] class.
var focusRingTypeClass _focusRingTypeClass

func init() {
	focusRingTypeClass = _focusRingTypeClass{objc.GetClass("focusRingType")}
}

type _focusRingTypeClass struct {
	objc.Class
}

// An interface definition for the [focusRingType] class.
type IfocusRingType interface {
	ID() objc.ID
}

type focusRingType struct {
	id objc.ID
}

func focusRingTypeFrom(ptr unsafe.Pointer) focusRingType {
	return focusRingType{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ focusRingType) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _focusRingTypeClass) Alloc() focusRingType {
	rv := objc.Send[focusRingType](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _focusRingTypeClass) New() focusRingType {
	rv := objc.Send[focusRingType](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfocusRingType creates and returns a new initialized instance.
func NewfocusRingType() focusRingType {
	return focusRingTypeClass.New()
}

// Init initializes the instance.
func (f_ focusRingType) Init() focusRingType {
	rv := objc.Send[focusRingType](f_.ID(), selInit)
	return rv
}
