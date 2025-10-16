
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [controlTint] class.
var controlTintClass _controlTintClass

func init() {
	controlTintClass = _controlTintClass{objc.GetClass("controlTint")}
}

type _controlTintClass struct {
	objc.Class
}

// An interface definition for the [controlTint] class.
type IcontrolTint interface {
	ID() objc.ID
}

type controlTint struct {
	id objc.ID
}

func controlTintFrom(ptr unsafe.Pointer) controlTint {
	return controlTint{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ controlTint) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _controlTintClass) Alloc() controlTint {
	rv := objc.Send[controlTint](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _controlTintClass) New() controlTint {
	rv := objc.Send[controlTint](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontrolTint creates and returns a new initialized instance.
func NewcontrolTint() controlTint {
	return controlTintClass.New()
}

// Init initializes the instance.
func (c_ controlTint) Init() controlTint {
	rv := objc.Send[controlTint](c_.ID(), selInit)
	return rv
}
