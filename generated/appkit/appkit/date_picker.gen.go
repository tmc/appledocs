// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DatePicker] class.
var DatePickerClass objc.Class

func init() {
	DatePickerClass = objc.GetClass("NSDatePicker")
}

type DatePicker struct {
	objc.ID
}

func DatePickerFrom(ptr unsafe.Pointer) DatePicker {
	return DatePicker{
		ID: objc.ID(ptr),
	}
}



