
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [controlSize] class.
var controlSizeClass _controlSizeClass

func init() {
	controlSizeClass = _controlSizeClass{objc.GetClass("controlSize")}
}

type _controlSizeClass struct {
	objc.Class
}

// An interface definition for the [controlSize] class.
type IcontrolSize interface {
	ID() objc.ID
}

type controlSize struct {
	id objc.ID
}

func controlSizeFrom(ptr unsafe.Pointer) controlSize {
	return controlSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ controlSize) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _controlSizeClass) Alloc() controlSize {
	rv := objc.Send[controlSize](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _controlSizeClass) New() controlSize {
	rv := objc.Send[controlSize](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontrolSize creates and returns a new initialized instance.
func NewcontrolSize() controlSize {
	return controlSizeClass.New()
}

// Init initializes the instance.
func (c_ controlSize) Init() controlSize {
	rv := objc.Send[controlSize](c_.ID(), selInit)
	return rv
}
