
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PopUpButton] class.
var PopUpButtonClass _PopUpButtonClass

func init() {
	PopUpButtonClass = _PopUpButtonClass{objc.GetClass("NSPopUpButton")}
}

type _PopUpButtonClass struct {
	objc.Class
}

// An interface definition for the [PopUpButton] class.
type IPopUpButton interface {
	ID() objc.ID
}

type PopUpButton struct {
	id objc.ID
}

func PopUpButtonFrom(ptr unsafe.Pointer) PopUpButton {
	return PopUpButton{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PopUpButton) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PopUpButtonClass) Alloc() PopUpButton {
	rv := objc.Send[PopUpButton](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PopUpButtonClass) New() PopUpButton {
	rv := objc.Send[PopUpButton](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPopUpButton creates and returns a new initialized instance.
func NewPopUpButton() PopUpButton {
	return PopUpButtonClass.New()
}

// Init initializes the instance.
func (p_ PopUpButton) Init() PopUpButton {
	rv := objc.Send[PopUpButton](p_.ID(), selInit)
	return rv
}
