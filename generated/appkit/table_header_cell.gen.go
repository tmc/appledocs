// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TableHeaderCell] class.
var tableHeaderCellClass = _TableHeaderCellClass{objc.GetClass("NSTableHeaderCell")}

type _TableHeaderCellClass struct {
	class objc.Class
}

// An object that a table header view uses to draw the content of the column headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderCell

type TableHeaderCell struct {
	TextFieldCell
}

// TableHeaderCellFrom constructs a [TableHeaderCell] from an unsafe.Pointer.
//
// An object that a table header view uses to draw the content of the column headers.
func TableHeaderCellFrom(ptr unsafe.Pointer) TableHeaderCell {
	return TableHeaderCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}



