// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ComboBox] class.
var comboBoxClass = _ComboBoxClass{objc.GetClass("NSComboBox")}

type _ComboBoxClass struct {
	class objc.Class
}

// A view that displays a list of values in a pop-up menu where the user selects a value or types in a custom value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBox

type ComboBox struct {
	TextField
}

// ComboBoxFrom constructs a [ComboBox] from an unsafe.Pointer.
//
// A view that displays a list of values in a pop-up menu where the user selects a value or types in a custom value.
func ComboBoxFrom(ptr unsafe.Pointer) ComboBox {
	return ComboBox{
		TextField: TextFieldFrom(ptr),
	}
}



