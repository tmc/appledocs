// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextField] class.
var TextFieldClass objc.Class

func init() {
	TextFieldClass = objc.GetClass("NSTextField")
}

type TextField struct {
	objc.ID
}

func TextFieldFrom(ptr unsafe.Pointer) TextField {
	return TextField{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc TextField) Alloc() TextField {
	ret := objc.ID(TextFieldClass).Send(objc.RegisterName("alloc"))
	return TextField{ret}
}

// New creates and returns a new initialized instance.
func (tc TextField) New() TextField {
	ret := objc.ID(TextFieldClass).Send(objc.RegisterName("new"))
	return TextField{ret}
}

// NewTextField creates and returns a new initialized instance.
func NewTextField() TextField {
	ret := objc.ID(TextFieldClass).Send(objc.RegisterName("new"))
	return TextField{ret}
}

// Init initializes the instance.
func (t_ TextField) Init() TextField {
	ret := t_.ID.Send(objc.RegisterName("init"))
	return TextField{ret}
}

