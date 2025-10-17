// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DatePickerCell] class.
var datePickerCellClass = _DatePickerCellClass{objc.GetClass("NSDatePickerCell")}

type _DatePickerCellClass struct {
	class objc.Class
}

// An interface definition for the [DatePickerCell] class.
type IDatePickerCell interface {
	IActionCell
}

// An object that controls the behavior of a date picker, or of a single date picker cell in a matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell

type DatePickerCell struct {
	ActionCell
}

// DatePickerCellFrom constructs a [DatePickerCell] from an unsafe.Pointer.
//
// An object that controls the behavior of a date picker, or of a single date picker cell in a matrix.
func DatePickerCellFrom(ptr unsafe.Pointer) DatePickerCell {
	return DatePickerCell{
		ActionCell: ActionCellFrom(ptr),
	}
}



