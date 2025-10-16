
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ComboBoxCell] class.
var ComboBoxCellClass _ComboBoxCellClass

func init() {
	ComboBoxCellClass = _ComboBoxCellClass{objc.GetClass("NSComboBoxCell")}
}

type _ComboBoxCellClass struct {
	objc.Class
}

// An interface definition for the [ComboBoxCell] class.
type IComboBoxCell interface {
	ID() objc.ID
}

type ComboBoxCell struct {
	id objc.ID
}

func ComboBoxCellFrom(ptr unsafe.Pointer) ComboBoxCell {
	return ComboBoxCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ComboBoxCell) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ComboBoxCellClass) Alloc() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ComboBoxCellClass) New() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewComboBoxCell creates and returns a new initialized instance.
func NewComboBoxCell() ComboBoxCell {
	return ComboBoxCellClass.New()
}

// Init initializes the instance.
func (c_ ComboBoxCell) Init() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](c_.ID(), selInit)
	return rv
}
