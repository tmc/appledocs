
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [yPlacement] class.
var yPlacementClass _yPlacementClass

func init() {
	yPlacementClass = _yPlacementClass{objc.GetClass("yPlacement")}
}

type _yPlacementClass struct {
	objc.Class
}

// An interface definition for the [yPlacement] class.
type IyPlacement interface {
	ID() objc.ID
}

type yPlacement struct {
	id objc.ID
}

func yPlacementFrom(ptr unsafe.Pointer) yPlacement {
	return yPlacement{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (y_ yPlacement) ID() objc.ID {
	return y_.id
}

// Alloc allocates a new instance without initialization.
func (yc _yPlacementClass) Alloc() yPlacement {
	rv := objc.Send[yPlacement](objc.ID(yc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (yc _yPlacementClass) New() yPlacement {
	rv := objc.Send[yPlacement](objc.ID(yc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewyPlacement creates and returns a new initialized instance.
func NewyPlacement() yPlacement {
	return yPlacementClass.New()
}

// Init initializes the instance.
func (y_ yPlacement) Init() yPlacement {
	rv := objc.Send[yPlacement](y_.ID(), selInit)
	return rv
}
