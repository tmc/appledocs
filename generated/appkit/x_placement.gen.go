
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [xPlacement] class.
var xPlacementClass _xPlacementClass

func init() {
	xPlacementClass = _xPlacementClass{objc.GetClass("xPlacement")}
}

type _xPlacementClass struct {
	objc.Class
}

// An interface definition for the [xPlacement] class.
type IxPlacement interface {
	ID() objc.ID
}

type xPlacement struct {
	id objc.ID
}

func xPlacementFrom(ptr unsafe.Pointer) xPlacement {
	return xPlacement{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (x_ xPlacement) ID() objc.ID {
	return x_.id
}

// Alloc allocates a new instance without initialization.
func (xc _xPlacementClass) Alloc() xPlacement {
	rv := objc.Send[xPlacement](objc.ID(xc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (xc _xPlacementClass) New() xPlacement {
	rv := objc.Send[xPlacement](objc.ID(xc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewxPlacement creates and returns a new initialized instance.
func NewxPlacement() xPlacement {
	return xPlacementClass.New()
}

// Init initializes the instance.
func (x_ xPlacement) Init() xPlacement {
	rv := objc.Send[xPlacement](x_.ID(), selInit)
	return rv
}
