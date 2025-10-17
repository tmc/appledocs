// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DatePickerCell] class.
var DatePickerCellClass objc.Class

func init() {
	DatePickerCellClass = objc.GetClass("NSDatePickerCell")
}

type DatePickerCell struct {
	objc.ID
}

func DatePickerCellFrom(ptr unsafe.Pointer) DatePickerCell {
	return DatePickerCell{
		ID: objc.ID(ptr),
	}
}




