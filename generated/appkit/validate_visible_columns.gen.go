
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [validateVisibleColumns] class.
var validateVisibleColumnsClass _validateVisibleColumnsClass

func init() {
	validateVisibleColumnsClass = _validateVisibleColumnsClass{objc.GetClass("validateVisibleColumns")}
}

type _validateVisibleColumnsClass struct {
	objc.Class
}

// An interface definition for the [validateVisibleColumns] class.
type IvalidateVisibleColumns interface {
	ID() objc.ID
}

type validateVisibleColumns struct {
	id objc.ID
}

func validateVisibleColumnsFrom(ptr unsafe.Pointer) validateVisibleColumns {
	return validateVisibleColumns{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ validateVisibleColumns) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _validateVisibleColumnsClass) Alloc() validateVisibleColumns {
	rv := objc.Send[validateVisibleColumns](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _validateVisibleColumnsClass) New() validateVisibleColumns {
	rv := objc.Send[validateVisibleColumns](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewvalidateVisibleColumns creates and returns a new initialized instance.
func NewvalidateVisibleColumns() validateVisibleColumns {
	return validateVisibleColumnsClass.New()
}

// Init initializes the instance.
func (v_ validateVisibleColumns) Init() validateVisibleColumns {
	rv := objc.Send[validateVisibleColumns](v_.ID(), selInit)
	return rv
}
