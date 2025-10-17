
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PopUpButtonCell] class.
var PopUpButtonCellClass _PopUpButtonCellClass

func init() {
	PopUpButtonCellClass = _PopUpButtonCellClass{objc.GetClass("NSPopUpButtonCell")}
}

type _PopUpButtonCellClass struct {
	objc.Class
}

// An interface definition for the [PopUpButtonCell] class.
type IPopUpButtonCell interface {
	ID() objc.ID
}

type PopUpButtonCell struct {
	id objc.ID
}

func PopUpButtonCellFrom(ptr unsafe.Pointer) PopUpButtonCell {
	return PopUpButtonCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PopUpButtonCell) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PopUpButtonCellClass) Alloc() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PopUpButtonCellClass) New() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPopUpButtonCell creates and returns a new initialized instance.
func NewPopUpButtonCell() PopUpButtonCell {
	return PopUpButtonCellClass.New()
}

// Init initializes the instance.
func (p_ PopUpButtonCell) Init() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](p_.ID(), selInit)
	return rv
}
