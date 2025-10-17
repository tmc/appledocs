// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ComboButton] class.
var ComboButtonClass objc.Class

func init() {
	ComboButtonClass = objc.GetClass("NSComboButton")
}

type ComboButton struct {
	objc.ID
}

func ComboButtonFrom(ptr unsafe.Pointer) ComboButton {
	return ComboButton{
		ID: objc.ID(ptr),
	}
}



