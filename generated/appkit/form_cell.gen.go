
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FormCell] class.
var FormCellClass _FormCellClass

func init() {
	FormCellClass = _FormCellClass{objc.GetClass("NSFormCell")}
}

type _FormCellClass struct {
	objc.Class
}

// An interface definition for the [FormCell] class.
type IFormCell interface {
	ID() objc.ID
}

type FormCell struct {
	id objc.ID
}

func FormCellFrom(ptr unsafe.Pointer) FormCell {
	return FormCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ FormCell) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _FormCellClass) Alloc() FormCell {
	rv := objc.Send[FormCell](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _FormCellClass) New() FormCell {
	rv := objc.Send[FormCell](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewFormCell creates and returns a new initialized instance.
func NewFormCell() FormCell {
	return FormCellClass.New()
}

// Init initializes the instance.
func (f_ FormCell) Init() FormCell {
	rv := objc.Send[FormCell](f_.ID(), selInit)
	return rv
}
