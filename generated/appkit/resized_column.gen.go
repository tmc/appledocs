
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [resizedColumn] class.
var resizedColumnClass _resizedColumnClass

func init() {
	resizedColumnClass = _resizedColumnClass{objc.GetClass("resizedColumn")}
}

type _resizedColumnClass struct {
	objc.Class
}

// An interface definition for the [resizedColumn] class.
type IresizedColumn interface {
	ID() objc.ID
}

type resizedColumn struct {
	id objc.ID
}

func resizedColumnFrom(ptr unsafe.Pointer) resizedColumn {
	return resizedColumn{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ resizedColumn) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _resizedColumnClass) Alloc() resizedColumn {
	rv := objc.Send[resizedColumn](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _resizedColumnClass) New() resizedColumn {
	rv := objc.Send[resizedColumn](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewresizedColumn creates and returns a new initialized instance.
func NewresizedColumn() resizedColumn {
	return resizedColumnClass.New()
}

// Init initializes the instance.
func (r_ resizedColumn) Init() resizedColumn {
	rv := objc.Send[resizedColumn](r_.ID(), selInit)
	return rv
}
