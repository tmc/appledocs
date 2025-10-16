
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [rowAlignment] class.
var rowAlignmentClass _rowAlignmentClass

func init() {
	rowAlignmentClass = _rowAlignmentClass{objc.GetClass("rowAlignment")}
}

type _rowAlignmentClass struct {
	objc.Class
}

// An interface definition for the [rowAlignment] class.
type IrowAlignment interface {
	ID() objc.ID
}

type rowAlignment struct {
	id objc.ID
}

func rowAlignmentFrom(ptr unsafe.Pointer) rowAlignment {
	return rowAlignment{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ rowAlignment) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _rowAlignmentClass) Alloc() rowAlignment {
	rv := objc.Send[rowAlignment](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _rowAlignmentClass) New() rowAlignment {
	rv := objc.Send[rowAlignment](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrowAlignment creates and returns a new initialized instance.
func NewrowAlignment() rowAlignment {
	return rowAlignmentClass.New()
}

// Init initializes the instance.
func (r_ rowAlignment) Init() rowAlignment {
	rv := objc.Send[rowAlignment](r_.ID(), selInit)
	return rv
}
