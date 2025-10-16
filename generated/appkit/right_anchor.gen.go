
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [rightAnchor] class.
var rightAnchorClass _rightAnchorClass

func init() {
	rightAnchorClass = _rightAnchorClass{objc.GetClass("rightAnchor")}
}

type _rightAnchorClass struct {
	objc.Class
}

// An interface definition for the [rightAnchor] class.
type IrightAnchor interface {
	ID() objc.ID
}

type rightAnchor struct {
	id objc.ID
}

func rightAnchorFrom(ptr unsafe.Pointer) rightAnchor {
	return rightAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ rightAnchor) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _rightAnchorClass) Alloc() rightAnchor {
	rv := objc.Send[rightAnchor](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _rightAnchorClass) New() rightAnchor {
	rv := objc.Send[rightAnchor](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrightAnchor creates and returns a new initialized instance.
func NewrightAnchor() rightAnchor {
	return rightAnchorClass.New()
}

// Init initializes the instance.
func (r_ rightAnchor) Init() rightAnchor {
	rv := objc.Send[rightAnchor](r_.ID(), selInit)
	return rv
}
