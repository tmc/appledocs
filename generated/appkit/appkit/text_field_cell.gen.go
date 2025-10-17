// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextFieldCell] class.
var TextFieldCellClass objc.Class

func init() {
	TextFieldCellClass = objc.GetClass("NSTextFieldCell")
}

type TextFieldCell struct {
	objc.ID
}

func TextFieldCellFrom(ptr unsafe.Pointer) TextFieldCell {
	return TextFieldCell{
		ID: objc.ID(ptr),
	}
}




