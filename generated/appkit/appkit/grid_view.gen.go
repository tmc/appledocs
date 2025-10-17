// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GridView] class.
var GridViewClass objc.Class

func init() {
	GridViewClass = objc.GetClass("NSGridView")
}

type GridView struct {
	objc.ID
}

func GridViewFrom(ptr unsafe.Pointer) GridView {
	return GridView{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc GridView) Alloc() GridView {
	ret := objc.ID(GridViewClass).Send(objc.RegisterName("alloc"))
	return GridView{ret}
}

// Init initializes the instance.
func (g_ GridView) Init() GridView {
	ret := g_.ID.Send(objc.RegisterName("init"))
	return GridView{ret}
}
// Creates a newly allocated grid view object from the coder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/init(coder:)
func NewGridViewWithCoder(coder unsafe.Pointer) GridView {
	instance := GridView{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = GridView{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Creates a newly allocated grid view object with the specified frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/init(frame:)
func NewGridViewWithFrame(frameRect unsafe.Pointer) GridView {
	instance := GridView{}.Alloc()
	sel := objc.RegisterName("initWithFrame:")
	ret := instance.ID.Send(sel, frameRect)
	instance = GridView{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Creates a newly allocated grid view object with the specified number of columns and rows. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/init(numberOfColumns:rows:)
func (gc GridView) GridViewWithNumberOfColumnsRows(columnCount int, rowCount int) unsafe.Pointer {
	sel := objc.RegisterName("gridViewWithNumberOfColumns:rows:")
	ret := objc.ID(GridViewClass).Send(sel, columnCount, rowCount)
	return unsafe.Pointer(ret)
}
// Creates a newly allocated grid view object with the specified array of arrays of views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/init(views:)
func (gc GridView) GridViewWithViews(rows unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("gridViewWithViews:")
	ret := objc.ID(GridViewClass).Send(sel, rows)
	return unsafe.Pointer(ret)
}
// Adds a new column containing the array of views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/addColumn(with:)
func (g_ GridView) AddColumnWithViews(views unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("addColumnWithViews:")
	ret := g_.ID.Send(sel, views)
	return unsafe.Pointer(ret)
}
// Adds an array of views to a new row. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/addRow(with:)
func (g_ GridView) AddRowWithViews(views unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("addRowWithViews:")
	ret := g_.ID.Send(sel, views)
	return unsafe.Pointer(ret)
}
// Returns the grid cell object at the specified column and row index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/cell(atColumnIndex:rowIndex:)
func (g_ GridView) CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) unsafe.Pointer {
	sel := objc.RegisterName("cellAtColumnIndex:rowIndex:")
	ret := g_.ID.Send(sel, columnIndex, rowIndex)
	return unsafe.Pointer(ret)
}
// Returns the grid cell object that contains the given view or one of its ancestors. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/cell(for:)
func (g_ GridView) CellForView(view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("cellForView:")
	ret := g_.ID.Send(sel, view)
	return unsafe.Pointer(ret)
}
// Returns the grid column object at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/column(at:)
func (g_ GridView) ColumnAtIndex(index int) unsafe.Pointer {
	sel := objc.RegisterName("columnAtIndex:")
	ret := g_.ID.Send(sel, index)
	return unsafe.Pointer(ret)
}
// Returns the index of the specified grid column. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/index(of:)-32sdd
func (g_ GridView) IndexOfColumn(column unsafe.Pointer) int {
	sel := objc.RegisterName("indexOfColumn:")
	ret := g_.ID.Send(sel, column)
	return int(ret)
}
// Returns the index of the specified grid row. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/index(of:)-6zs2o
func (g_ GridView) IndexOfRow(row unsafe.Pointer) int {
	sel := objc.RegisterName("indexOfRow:")
	ret := g_.ID.Send(sel, row)
	return int(ret)
}
// Inserts the array of view objects at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/insertColumn(at:with:)
func (g_ GridView) InsertColumnAtIndexWithViews(index int, views unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("insertColumnAtIndex:withViews:")
	ret := g_.ID.Send(sel, index, views)
	return unsafe.Pointer(ret)
}
// Inserts the array of view objects into the grid view at the index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/insertRow(at:with:)
func (g_ GridView) InsertRowAtIndexWithViews(index int, views unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("insertRowAtIndex:withViews:")
	ret := g_.ID.Send(sel, index, views)
	return unsafe.Pointer(ret)
}
// Expands the cell at the top-leading corner of the horizontal and vertical range to cover the entire area. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/mergeCells(inHorizontalRange:verticalRange:)
func (g_ GridView) MergeCellsInHorizontalRangeVerticalRange(hRange unsafe.Pointer, vRange unsafe.Pointer) {
	sel := objc.RegisterName("mergeCellsInHorizontalRange:verticalRange:")
	g_.ID.Send(sel, hRange, vRange)
}
// Moves the specified column to a new column location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/moveColumn(at:to:)
func (g_ GridView) MoveColumnAtIndexToIndex(fromIndex int, toIndex int) {
	sel := objc.RegisterName("moveColumnAtIndex:toIndex:")
	g_.ID.Send(sel, fromIndex, toIndex)
}
// Moves the specified row to the new row location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/moveRow(at:to:)
func (g_ GridView) MoveRowAtIndexToIndex(fromIndex int, toIndex int) {
	sel := objc.RegisterName("moveRowAtIndex:toIndex:")
	g_.ID.Send(sel, fromIndex, toIndex)
}
// Removes the column from the grid view at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/removeColumn(at:)
func (g_ GridView) RemoveColumnAtIndex(index int) {
	sel := objc.RegisterName("removeColumnAtIndex:")
	g_.ID.Send(sel, index)
}
// Removes the row from the grid view at the index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/removeRow(at:)
func (g_ GridView) RemoveRowAtIndex(index int) {
	sel := objc.RegisterName("removeRowAtIndex:")
	g_.ID.Send(sel, index)
}
// Returns the grid row object at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/row(at:)
func (g_ GridView) RowAtIndex(index int) unsafe.Pointer {
	sel := objc.RegisterName("rowAtIndex:")
	ret := g_.ID.Send(sel, index)
	return unsafe.Pointer(ret)
}

