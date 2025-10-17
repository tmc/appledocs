
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ComboBox] class.
var ComboBoxClass _ComboBoxClass

func init() {
	ComboBoxClass = _ComboBoxClass{objc.GetClass("NSComboBox")}
}

type _ComboBoxClass struct {
	objc.Class
}

// An interface definition for the [ComboBox] class.
type IComboBox interface {
	ID() objc.ID
}

type ComboBox struct {
	id objc.ID
}

func ComboBoxFrom(ptr unsafe.Pointer) ComboBox {
	return ComboBox{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ComboBox) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ComboBoxClass) Alloc() ComboBox {
	rv := objc.Send[ComboBox](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ComboBoxClass) New() ComboBox {
	rv := objc.Send[ComboBox](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewComboBox creates and returns a new initialized instance.
func NewComboBox() ComboBox {
	return ComboBoxClass.New()
}

// Init initializes the instance.
func (c_ ComboBox) Init() ComboBox {
	rv := objc.Send[ComboBox](c_.ID(), selInit)
	return rv
}
