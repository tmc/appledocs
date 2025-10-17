// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ComboBoxCell] class.
var comboBoxCellClass = _ComboBoxCellClass{objc.GetClass("NSComboBoxCell")}

type _ComboBoxCellClass struct {
	class objc.Class
}

// An interface definition for the [ComboBoxCell] class.
type IComboBoxCell interface {
	ITextFieldCell
}

// The user interface of a combo box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell

type ComboBoxCell struct {
	TextFieldCell
}

// ComboBoxCellFrom constructs a [ComboBoxCell] from an unsafe.Pointer.
//
// The user interface of a combo box.
func ComboBoxCellFrom(ptr unsafe.Pointer) ComboBoxCell {
	return ComboBoxCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}



