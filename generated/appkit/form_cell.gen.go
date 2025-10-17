// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FormCell] class.
var formCellClass = _FormCellClass{objc.GetClass("NSFormCell")}

type _FormCellClass struct {
	class objc.Class
}

// The class is used to implement text entry fields in a form. The left part of an object contains a title. The right part contains an editable text entry field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFormCell

type FormCell struct {
	ActionCell
}

// FormCellFrom constructs a [FormCell] from an unsafe.Pointer.
//
// The class is used to implement text entry fields in a form. The left part of an object contains a title. The right part contains an editable text entry field.
func FormCellFrom(ptr unsafe.Pointer) FormCell {
	return FormCell{
		ActionCell: ActionCellFrom(ptr),
	}
}



