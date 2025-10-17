
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ComboButton] class.
var ComboButtonClass _ComboButtonClass

func init() {
	ComboButtonClass = _ComboButtonClass{objc.GetClass("NSComboButton")}
}

type _ComboButtonClass struct {
	objc.Class
}

// An interface definition for the [ComboButton] class.
type IComboButton interface {
	ID() objc.ID
}

type ComboButton struct {
	id objc.ID
}

func ComboButtonFrom(ptr unsafe.Pointer) ComboButton {
	return ComboButton{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ComboButton) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ComboButtonClass) Alloc() ComboButton {
	rv := objc.Send[ComboButton](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ComboButtonClass) New() ComboButton {
	rv := objc.Send[ComboButton](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewComboButton creates and returns a new initialized instance.
func NewComboButton() ComboButton {
	return ComboButtonClass.New()
}

// Init initializes the instance.
func (c_ ComboButton) Init() ComboButton {
	rv := objc.Send[ComboButton](c_.ID(), selInit)
	return rv
}
