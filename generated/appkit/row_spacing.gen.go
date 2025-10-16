
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [rowSpacing] class.
var rowSpacingClass _rowSpacingClass

func init() {
	rowSpacingClass = _rowSpacingClass{objc.GetClass("rowSpacing")}
}

type _rowSpacingClass struct {
	objc.Class
}

// An interface definition for the [rowSpacing] class.
type IrowSpacing interface {
	ID() objc.ID
}

type rowSpacing struct {
	id objc.ID
}

func rowSpacingFrom(ptr unsafe.Pointer) rowSpacing {
	return rowSpacing{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ rowSpacing) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _rowSpacingClass) Alloc() rowSpacing {
	rv := objc.Send[rowSpacing](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _rowSpacingClass) New() rowSpacing {
	rv := objc.Send[rowSpacing](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrowSpacing creates and returns a new initialized instance.
func NewrowSpacing() rowSpacing {
	return rowSpacingClass.New()
}

// Init initializes the instance.
func (r_ rowSpacing) Init() rowSpacing {
	rv := objc.Send[rowSpacing](r_.ID(), selInit)
	return rv
}
