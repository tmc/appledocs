// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PopUpButtonCell] class.
var PopUpButtonCellClass objc.Class

func init() {
	PopUpButtonCellClass = objc.GetClass("NSPopUpButtonCell")
}

type PopUpButtonCell struct {
	objc.ID
}

func PopUpButtonCellFrom(ptr unsafe.Pointer) PopUpButtonCell {
	return PopUpButtonCell{
		ID: objc.ID(ptr),
	}
}



