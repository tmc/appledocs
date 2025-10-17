// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableView] class.
var TableViewClass objc.Class

func init() {
	TableViewClass = objc.GetClass("NSTableView")
}

type TableView struct {
	objc.ID
}

func TableViewFrom(ptr unsafe.Pointer) TableView {
	return TableView{
		ID: objc.ID(ptr),
	}
}


// Returns the rectangle containing the row at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableView/rect(ofRow:)
func (t_ TableView) RectOfRow(row int) unsafe.Pointer {
	sel := objc.RegisterName("rectOfRow:")
	ret := t_.ID.Send(sel, row)
	return unsafe.Pointer(ret)
}
// Returns the index of the row the specified point lies in. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableView/row(at:)
func (t_ TableView) RowAtPoint(point unsafe.Pointer) int {
	sel := objc.RegisterName("rowAtPoint:")
	ret := t_.ID.Send(sel, point)
	return int(ret)
}

