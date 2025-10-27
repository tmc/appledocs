// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [TableHeaderView] class.
var (
	TableHeaderViewClass     _TableHeaderViewClass
	TableHeaderViewClassOnce sync.Once
)

func getTableHeaderViewClass() _TableHeaderViewClass {
	TableHeaderViewClassOnce.Do(func() {
		TableHeaderViewClass = _TableHeaderViewClass{objc.GetClass("NSTableHeaderView")}
	})
	return TableHeaderViewClass
}

type _TableHeaderViewClass struct {
	class objc.Class
}





// An interface definition for the [TableHeaderView] class.
type ITableHeaderView interface {
	IView
	

	// properties:
	DraggedColumn() int
	DraggedDistance() float64
	ResizedColumn() int
	TableView() ITableView
	SetTableView(value ITableView)


	

	// methods:
	ColumnAtPoint(point corefoundation.CGPoint) int
	HeaderRectOfColumn(column int) corefoundation.CGRect


}





// Alloc allocates a new instance without initialization.
func (tc _TableHeaderViewClass) Alloc() TableHeaderView {
	rv := objc.Send[TableHeaderView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TableHeaderViewClass) New() TableHeaderView {
	rv := objc.Send[TableHeaderView](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableHeaderView) Init() TableHeaderView {
	rv := objc.Send[TableHeaderView](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableHeaderView) Autorelease() TableHeaderView {
	rv := objc.Send[TableHeaderView](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableHeaderView creates a new TableHeaderView instance.
func NewTableHeaderView() TableHeaderView {
	return getTableHeaderViewClass().New()
}





// An object that draws headers over a table view’s columns and handles mouse events in those headers.
//
// uses to implement its user interface.


// An object that draws headers over a table view’s columns and handles mouse events in those headers.
//
// [Full Topic]
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




















// Returns the index of the column whose header lies under in the receiver, or –1 if no such column is found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/column(at:)
func (t_ TableHeaderView) ColumnAtPoint(point corefoundation.CGPoint) int {
	rv := objc.Send[int](t_.ID, objc.Sel("columnAtPoint:"), point)
	return rv
}


// Returns the rectangle containing the header tile for the column at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/headerRect(ofColumn:)
func (t_ TableHeaderView) HeaderRectOfColumn(column int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("headerRectOfColumn:"), column)
	return rv
}







// The index of the column that the user is dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/draggedColumn
func (t_ TableHeaderView) DraggedColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("draggedColumn"))
	return rv
}


// The horizontal distance that the user has dragged a column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/draggedDistance
func (t_ TableHeaderView) DraggedDistance() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("draggedDistance"))
	return rv
}


// The index of the column that the user is resizing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/resizedColumn
func (t_ TableHeaderView) ResizedColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("resizedColumn"))
	return rv
}


// The instance that this table header view belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/tableView
func (t_ TableHeaderView) TableView() ITableView {
	rv := objc.Send[TableView](t_.ID, objc.Sel("tableView"))
	return rv
}


// The instance that this table header view belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/tableView
func (t_ TableHeaderView) SetTableView(value ITableView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTableView:"), value)
}








