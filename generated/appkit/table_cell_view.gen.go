// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TableCellView] class.
var tableCellViewClass = _TableCellViewClass{objc.GetClass("NSTableCellView")}

type _TableCellViewClass struct {
	class objc.Class
}

// A reusable container view shown for a particular cell in a table view that uses rows for content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView

type TableCellView struct {
	View
}

// TableCellViewFrom constructs a [TableCellView] from an unsafe.Pointer.
//
// A reusable container view shown for a particular cell in a table view that uses rows for content.
func TableCellViewFrom(ptr unsafe.Pointer) TableCellView {
	return TableCellView{
		View: ViewFrom(ptr),
	}
}



