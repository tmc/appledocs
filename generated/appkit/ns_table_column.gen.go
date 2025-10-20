// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TableColumn] class.
var (
	TableColumnClass     _TableColumnClass
	TableColumnClassOnce sync.Once
)

func getTableColumnClass() _TableColumnClass {
	TableColumnClassOnce.Do(func() {
		TableColumnClass = _TableColumnClass{objc.GetClass("NSTableColumn")}
	})
	return TableColumnClass
}

type _TableColumnClass struct {
	class objc.Class
}

// An interface definition for the [TableColumn] class.
type ITableColumn interface {
	objectivec.IObject
}

// The display characteristics and identifier for a column in a table view.
//
// A table column object determines the width (including the maximum and minimum widths) of its column in the table view and specifies the column’s resizing and editing behavior. A table column stores two cell objects: the header cell, which is used to draw the column header, and the data cell, which is used to draw the values for each row. In a cell-based table, you can control the display of the column by specifying subclasses of to use and by setting the font and other display characteristics for these cells. For example, you can use an to display string values or substitute an to display pictures.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn
type TableColumn struct {
	objectivec.Object
}

// TableColumnFrom constructs a [TableColumn] from an unsafe.Pointer.
//
// The display characteristics and identifier for a column in a table view.
func TableColumnFrom(ptr unsafe.Pointer) TableColumn {
	return TableColumn{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TableColumnClass) Alloc() TableColumn {
	rv := objc.Send[TableColumn](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TableColumnClass) New() TableColumn {
	rv := objc.Send[TableColumn](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableColumn) Init() TableColumn {
	rv := objc.Send[TableColumn](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableColumn) Autorelease() TableColumn {
	rv := objc.Send[TableColumn](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableColumn creates a new TableColumn instance.
func NewTableColumn() TableColumn {
	return getTableColumnClass().New()
}


// The cell prototype used by the table column to draw individual cells.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/dataCell
func (t_ TableColumn) DataCell() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("dataCell"))
	return rv
}


// SetDataCell sets the value of the dataCell property.
// The cell prototype used by the table column to draw individual cells.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/dataCell
func (t_ TableColumn) SetDataCell(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDataCell:"), value)
}
// A Boolean that indicates whether a cell-based table’s column cells are user editable.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/isEditable
func (t_ TableColumn) Editable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("editable"))
	return rv
}


// SetEditable sets the value of the editable property.
// A Boolean that indicates whether a cell-based table’s column cells are user editable.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/isEditable
func (t_ TableColumn) SetEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditable:"), value)
}


