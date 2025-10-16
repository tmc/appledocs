
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [noteFocusRingMaskChanged] class.
var noteFocusRingMaskChangedClass _noteFocusRingMaskChangedClass

func init() {
	noteFocusRingMaskChangedClass = _noteFocusRingMaskChangedClass{objc.GetClass("noteFocusRingMaskChanged")}
}

type _noteFocusRingMaskChangedClass struct {
	objc.Class
}

// An interface definition for the [noteFocusRingMaskChanged] class.
type InoteFocusRingMaskChanged interface {
	ID() objc.ID
}

type noteFocusRingMaskChanged struct {
	id objc.ID
}

func noteFocusRingMaskChangedFrom(ptr unsafe.Pointer) noteFocusRingMaskChanged {
	return noteFocusRingMaskChanged{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ noteFocusRingMaskChanged) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _noteFocusRingMaskChangedClass) Alloc() noteFocusRingMaskChanged {
	rv := objc.Send[noteFocusRingMaskChanged](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _noteFocusRingMaskChangedClass) New() noteFocusRingMaskChanged {
	rv := objc.Send[noteFocusRingMaskChanged](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnoteFocusRingMaskChanged creates and returns a new initialized instance.
func NewnoteFocusRingMaskChanged() noteFocusRingMaskChanged {
	return noteFocusRingMaskChangedClass.New()
}

// Init initializes the instance.
func (n_ noteFocusRingMaskChanged) Init() noteFocusRingMaskChanged {
	rv := objc.Send[noteFocusRingMaskChanged](n_.ID(), selInit)
	return rv
}
