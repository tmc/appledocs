// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ComboBox] class.
var ComboBoxClass objc.Class

func init() {
	ComboBoxClass = objc.GetClass("NSComboBox")
}

type ComboBox struct {
	objc.ID
}

func ComboBoxFrom(ptr unsafe.Pointer) ComboBox {
	return ComboBox{
		ID: objc.ID(ptr),
	}
}




