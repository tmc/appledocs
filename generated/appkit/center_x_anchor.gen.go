
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [centerXAnchor] class.
var centerXAnchorClass _centerXAnchorClass

func init() {
	centerXAnchorClass = _centerXAnchorClass{objc.GetClass("centerXAnchor")}
}

type _centerXAnchorClass struct {
	objc.Class
}

// An interface definition for the [centerXAnchor] class.
type IcenterXAnchor interface {
	ID() objc.ID
}

type centerXAnchor struct {
	id objc.ID
}

func centerXAnchorFrom(ptr unsafe.Pointer) centerXAnchor {
	return centerXAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ centerXAnchor) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _centerXAnchorClass) Alloc() centerXAnchor {
	rv := objc.Send[centerXAnchor](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _centerXAnchorClass) New() centerXAnchor {
	rv := objc.Send[centerXAnchor](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcenterXAnchor creates and returns a new initialized instance.
func NewcenterXAnchor() centerXAnchor {
	return centerXAnchorClass.New()
}

// Init initializes the instance.
func (c_ centerXAnchor) Init() centerXAnchor {
	rv := objc.Send[centerXAnchor](c_.ID(), selInit)
	return rv
}
