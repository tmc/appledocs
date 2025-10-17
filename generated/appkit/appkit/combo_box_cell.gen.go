// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ComboBoxCell] class.
var ComboBoxCellClass objc.Class

func init() {
	ComboBoxCellClass = objc.GetClass("NSComboBoxCell")
}

type ComboBoxCell struct {
	objc.ID
}

func ComboBoxCellFrom(ptr unsafe.Pointer) ComboBoxCell {
	return ComboBoxCell{
		ID: objc.ID(ptr),
	}
}



