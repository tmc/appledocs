
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextFieldCell] class.
var TextFieldCellClass _TextFieldCellClass

func init() {
	TextFieldCellClass = _TextFieldCellClass{objc.GetClass("NSTextFieldCell")}
}

type _TextFieldCellClass struct {
	objc.Class
}

// An interface definition for the [TextFieldCell] class.
type ITextFieldCell interface {
	ID() objc.ID
}

type TextFieldCell struct {
	id objc.ID
}

func TextFieldCellFrom(ptr unsafe.Pointer) TextFieldCell {
	return TextFieldCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextFieldCell) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextFieldCellClass) Alloc() TextFieldCell {
	rv := objc.Send[TextFieldCell](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextFieldCellClass) New() TextFieldCell {
	rv := objc.Send[TextFieldCell](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextFieldCell creates and returns a new initialized instance.
func NewTextFieldCell() TextFieldCell {
	return TextFieldCellClass.New()
}

// Init initializes the instance.
func (t_ TextFieldCell) Init() TextFieldCell {
	rv := objc.Send[TextFieldCell](t_.ID(), selInit)
	return rv
}
