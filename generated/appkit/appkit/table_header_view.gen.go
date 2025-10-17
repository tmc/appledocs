// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableHeaderView] class.
var TableHeaderViewClass objc.Class

func init() {
	TableHeaderViewClass = objc.GetClass("NSTableHeaderView")
}

type TableHeaderView struct {
	objc.ID
}

func TableHeaderViewFrom(ptr unsafe.Pointer) TableHeaderView {
	return TableHeaderView{
		ID: objc.ID(ptr),
	}
}


// Returns the index of the column whose header lies under   in the receiver, or –1 if no such column is found. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableHeaderView/column(at:)
func (t_ TableHeaderView) ColumnAtPoint(point unsafe.Pointer) int {
	sel := objc.RegisterName("columnAtPoint:")
	ret := t_.ID.Send(sel, point)
	return int(ret)
}
// Returns the rectangle containing the header tile for the column at  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableHeaderView/headerRect(ofColumn:)
func (t_ TableHeaderView) HeaderRectOfColumn(column int) unsafe.Pointer {
	sel := objc.RegisterName("headerRectOfColumn:")
	ret := t_.ID.Send(sel, column)
	return unsafe.Pointer(ret)
}


