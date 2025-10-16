
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [columnSpacing] class.
var columnSpacingClass _columnSpacingClass

func init() {
	columnSpacingClass = _columnSpacingClass{objc.GetClass("columnSpacing")}
}

type _columnSpacingClass struct {
	objc.Class
}

// An interface definition for the [columnSpacing] class.
type IcolumnSpacing interface {
	ID() objc.ID
}

type columnSpacing struct {
	id objc.ID
}

func columnSpacingFrom(ptr unsafe.Pointer) columnSpacing {
	return columnSpacing{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ columnSpacing) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _columnSpacingClass) Alloc() columnSpacing {
	rv := objc.Send[columnSpacing](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _columnSpacingClass) New() columnSpacing {
	rv := objc.Send[columnSpacing](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcolumnSpacing creates and returns a new initialized instance.
func NewcolumnSpacing() columnSpacing {
	return columnSpacingClass.New()
}

// Init initializes the instance.
func (c_ columnSpacing) Init() columnSpacing {
	rv := objc.Send[columnSpacing](c_.ID(), selInit)
	return rv
}
