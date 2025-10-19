// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TableView] class.
var (
	tableViewClass     _TableViewClass
	tableViewClassOnce sync.Once
)

func getTableViewClass() _TableViewClass {
	tableViewClassOnce.Do(func() {
		tableViewClass = _TableViewClass{objc.GetClass("NSTableView")}
	})
	return tableViewClass
}

type _TableViewClass struct {
	class objc.Class
}

// An interface definition for the [TableView] class.
type ITableView interface {
	IControl
	RectOfRow(row int) unsafe.Pointer
	RowAtPoint(point unsafe.Pointer) int
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
// Alloc allocates a new instance without initialization.
func (tc _TableViewClass) Alloc() TableView {
	rv := objc.Send[TableView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TableViewClass) New() TableView {
	rv := objc.Send[TableView](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableView) Init() TableView {
	rv := objc.Send[TableView](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableView) Autorelease() TableView {
	rv := objc.Send[TableView](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableView creates a new TableView instance.
func NewTableView() TableView {
	return getTableViewClass().New()
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


