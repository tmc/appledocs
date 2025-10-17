// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TableView] class.
var tableViewClass = _TableViewClass{objc.GetClass("NSTableView")}

type _TableViewClass struct {
	class objc.Class
}

// A set of related records, displayed in rows that represent individual records and columns that represent the attributes of those records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView

type TableView struct {
	Control
}

// TableViewFrom constructs a [TableView] from an unsafe.Pointer.
//
// A set of related records, displayed in rows that represent individual records and columns that represent the attributes of those records.
func TableViewFrom(ptr unsafe.Pointer) TableView {
	return TableView{
		Control: ControlFrom(ptr),
	}
}

// Returns the rectangle containing the row at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rect(ofRow:)
func (t_ TableView) RectOfRow(row int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("rectOfRow:"), row)
	return rv
}
// Returns the index of the row the specified point lies in. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/row(at:)
func (t_ TableView) RowAtPoint(point unsafe.Pointer) int {
	rv := objc.Send[int](t_.ID, objc.Sel("rowAtPoint:"), point)
	return rv
}


