// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PopUpButton] class.
var PopUpButtonClass objc.Class

func init() {
	PopUpButtonClass = objc.GetClass("NSPopUpButton")
}

type PopUpButton struct {
	objc.ID
}

func PopUpButtonFrom(ptr unsafe.Pointer) PopUpButton {
	return PopUpButton{
		ID: objc.ID(ptr),
	}
}



