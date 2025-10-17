// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextFieldCell] class.
var textFieldCellClass = _TextFieldCellClass{objc.GetClass("NSTextFieldCell")}

type _TextFieldCellClass struct {
	class objc.Class
}

// An interface definition for the [TextFieldCell] class.
type ITextFieldCell interface {
	IActionCell
}

// An object that enhances the text display capabilities of a cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFieldCell

type TextFieldCell struct {
	ActionCell
}

// TextFieldCellFrom constructs a [TextFieldCell] from an unsafe.Pointer.
//
// An object that enhances the text display capabilities of a cell.
func TextFieldCellFrom(ptr unsafe.Pointer) TextFieldCell {
	return TextFieldCell{
		ActionCell: ActionCellFrom(ptr),
	}
}



