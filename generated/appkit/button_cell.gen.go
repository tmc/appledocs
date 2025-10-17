
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ButtonCell] class.
var ButtonCellClass _ButtonCellClass

func init() {
	ButtonCellClass = _ButtonCellClass{objc.GetClass("NSButtonCell")}
}

type _ButtonCellClass struct {
	objc.Class
}

// An interface definition for the [ButtonCell] class.
type IButtonCell interface {
	ID() objc.ID
}

type ButtonCell struct {
	id objc.ID
}

func ButtonCellFrom(ptr unsafe.Pointer) ButtonCell {
	return ButtonCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ ButtonCell) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _ButtonCellClass) Alloc() ButtonCell {
	rv := objc.Send[ButtonCell](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _ButtonCellClass) New() ButtonCell {
	rv := objc.Send[ButtonCell](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewButtonCell creates and returns a new initialized instance.
func NewButtonCell() ButtonCell {
	return ButtonCellClass.New()
}

// Init initializes the instance.
func (b_ ButtonCell) Init() ButtonCell {
	rv := objc.Send[ButtonCell](b_.ID(), selInit)
	return rv
}
