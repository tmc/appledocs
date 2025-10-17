// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TableHeaderView] class.
var tableHeaderViewClass = _TableHeaderViewClass{objc.GetClass("NSTableHeaderView")}

type _TableHeaderViewClass struct {
	class objc.Class
}

// An interface definition for the [TableHeaderView] class.
type ITableHeaderView interface {
	IView
	ColumnAtPoint(point unsafe.Pointer) int
	HeaderRectOfColumn(column int) unsafe.Pointer
}

// An object that draws headers over a table view’s columns and handles mouse events in those headers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderView

type TableHeaderView struct {
	View
}

// TableHeaderViewFrom constructs a [TableHeaderView] from an unsafe.Pointer.
//
// An object that draws headers over a table view’s columns and handles mouse events in those headers.
func TableHeaderViewFrom(ptr unsafe.Pointer) TableHeaderView {
	return TableHeaderView{
		View: ViewFrom(ptr),
	}
}

// Returns the index of the column whose header lies under in the receiver, or –1 if no such column is found. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/column(at:)
func (t_ TableHeaderView) ColumnAtPoint(point unsafe.Pointer) int {
	rv := objc.Send[int](t_.ID, objc.Sel("columnAtPoint:"), point)
	return rv
}
// Returns the rectangle containing the header tile for the column at . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/headerRect(ofColumn:)
func (t_ TableHeaderView) HeaderRectOfColumn(column int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("headerRectOfColumn:"), column)
	return rv
}


