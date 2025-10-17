// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ButtonCell] class.
var ButtonCellClass objc.Class

func init() {
	ButtonCellClass = objc.GetClass("NSButtonCell")
}

type ButtonCell struct {
	objc.ID
}

func ButtonCellFrom(ptr unsafe.Pointer) ButtonCell {
	return ButtonCell{
		ID: objc.ID(ptr),
	}
}



