// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ButtonCell] class.
var buttonCellClass = _ButtonCellClass{objc.GetClass("NSButtonCell")}

type _ButtonCellClass struct {
	class objc.Class
}

// An object that defines the user interface of a button or other clickable region of a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell

type ButtonCell struct {
	ActionCell
}

// ButtonCellFrom constructs a [ButtonCell] from an unsafe.Pointer.
//
// An object that defines the user interface of a button or other clickable region of a view.
func ButtonCellFrom(ptr unsafe.Pointer) ButtonCell {
	return ButtonCell{
		ActionCell: ActionCellFrom(ptr),
	}
}



