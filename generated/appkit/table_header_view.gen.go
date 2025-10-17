
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// The class instance for the [TableHeaderView] class.
var TableHeaderViewClass _TableHeaderViewClass

func init() {
	TableHeaderViewClass = _TableHeaderViewClass{objc.GetClass("NSTableHeaderView")}
}

type _TableHeaderViewClass struct {
	objc.Class
}

// An interface definition for the [TableHeaderView] class.
type ITableHeaderView interface {
	ID() objc.ID
	ColumnAtPoint(point foundation.Point) int
	HeaderRectOfColumn(column int) foundation.Rect
}

type TableHeaderView struct {
	id objc.ID
}

func TableHeaderViewFrom(ptr unsafe.Pointer) TableHeaderView {
	return TableHeaderView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TableHeaderView) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TableHeaderViewClass) Alloc() TableHeaderView {
	rv := objc.Send[TableHeaderView](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TableHeaderViewClass) New() TableHeaderView {
	rv := objc.Send[TableHeaderView](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTableHeaderView creates and returns a new initialized instance.
func NewTableHeaderView() TableHeaderView {
	return TableHeaderViewClass.New()
}

// Init initializes the instance.
func (t_ TableHeaderView) Init() TableHeaderView {
	rv := objc.Send[TableHeaderView](t_.ID(), selInit)
	return rv
}
// Returns the index of the column whose header lies under   in the receiver, or –1 if no such column is found. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableHeaderView/column(at:)
func (t_ TableHeaderView) ColumnAtPoint(point foundation.Point) int {
	rv := objc.Send[int](t_.ID(), objc.RegisterName("columnAtPoint:"), point)
	return rv
}
// Returns the rectangle containing the header tile for the column at  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableHeaderView/headerRect(ofColumn:)
func (t_ TableHeaderView) HeaderRectOfColumn(column int) foundation.Rect {
	rv := objc.Send[foundation.Rect](t_.ID(), objc.RegisterName("headerRectOfColumn:"), column)
	return rv
}
// The index of the column that the user is dragging. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableHeaderView/draggedColumn
func (t_ TableHeaderView) DraggedColumn() int {
	rv := objc.Send[int](t_.ID(), objc.RegisterName("draggedColumn"))
	return rv
}
// The horizontal distance that the user has dragged a column. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableHeaderView/draggedDistance
func (t_ TableHeaderView) DraggedDistance() float64 {
	rv := objc.Send[float64](t_.ID(), objc.RegisterName("draggedDistance"))
	return rv
}
// The index of the column that the user is resizing. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableHeaderView/resizedColumn
func (t_ TableHeaderView) ResizedColumn() int {
	rv := objc.Send[int](t_.ID(), objc.RegisterName("resizedColumn"))
	return rv
}
// The   instance that this table header view belongs to. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableHeaderView/tableView
func (t_ TableHeaderView) TableView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("tableView"))
	return rv
}
// SetTableView sets the value of the tableView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableHeaderView/tableView
func (t_ TableHeaderView) SetTableView(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setTableView:"), value)
}
