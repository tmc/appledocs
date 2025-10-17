// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DatePicker] class.
var datePickerClass = _DatePickerClass{objc.GetClass("NSDatePicker")}

type _DatePickerClass struct {
	class objc.Class
}

// An interface definition for the [DatePicker] class.
type IDatePicker interface {
	IControl
}

// A display of a calendar date with controls for editing the date value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker

type DatePicker struct {
	Control
}

// DatePickerFrom constructs a [DatePicker] from an unsafe.Pointer.
//
// A display of a calendar date with controls for editing the date value.
func DatePickerFrom(ptr unsafe.Pointer) DatePicker {
	return DatePicker{
		Control: ControlFrom(ptr),
	}
}



