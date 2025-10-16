
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [centerYAnchor] class.
var centerYAnchorClass _centerYAnchorClass

func init() {
	centerYAnchorClass = _centerYAnchorClass{objc.GetClass("centerYAnchor")}
}

type _centerYAnchorClass struct {
	objc.Class
}

// An interface definition for the [centerYAnchor] class.
type IcenterYAnchor interface {
	ID() objc.ID
}

type centerYAnchor struct {
	id objc.ID
}

func centerYAnchorFrom(ptr unsafe.Pointer) centerYAnchor {
	return centerYAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ centerYAnchor) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _centerYAnchorClass) Alloc() centerYAnchor {
	rv := objc.Send[centerYAnchor](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _centerYAnchorClass) New() centerYAnchor {
	rv := objc.Send[centerYAnchor](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcenterYAnchor creates and returns a new initialized instance.
func NewcenterYAnchor() centerYAnchor {
	return centerYAnchorClass.New()
}

// Init initializes the instance.
func (c_ centerYAnchor) Init() centerYAnchor {
	rv := objc.Send[centerYAnchor](c_.ID(), selInit)
	return rv
}
