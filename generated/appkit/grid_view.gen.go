
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GridView] class.
var GridViewClass _GridViewClass

func init() {
	GridViewClass = _GridViewClass{objc.GetClass("NSGridView")}
}

type _GridViewClass struct {
	objc.Class
}

// An interface definition for the [GridView] class.
type IGridView interface {
	ID() objc.ID
	AddColumnWithViews(views unsafe.Pointer) unsafe.Pointer
	AddRowWithViews(views unsafe.Pointer) unsafe.Pointer
	CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) unsafe.Pointer
	CellForView(view unsafe.Pointer) unsafe.Pointer
	ColumnAtIndex(index int) unsafe.Pointer
	IndexOfColumn(column unsafe.Pointer) int
	IndexOfRow(row unsafe.Pointer) int
	InitWithCoder(coder unsafe.Pointer) unsafe.Pointer
	InitWithFrame(frameRect unsafe.Pointer) unsafe.Pointer
	InsertColumnAtIndexWithViews(index int, views unsafe.Pointer) unsafe.Pointer
	InsertRowAtIndexWithViews(index int, views unsafe.Pointer) unsafe.Pointer
	MergeCellsInHorizontalRangeVerticalRange(hRange unsafe.Pointer, vRange unsafe.Pointer)
	MoveColumnAtIndexToIndex(fromIndex int, toIndex int)
	MoveRowAtIndexToIndex(fromIndex int, toIndex int)
	RemoveColumnAtIndex(index int)
	RemoveRowAtIndex(index int)
	RowAtIndex(index int) unsafe.Pointer
}

type GridView struct {
	id objc.ID
}

func GridViewFrom(ptr unsafe.Pointer) GridView {
	return GridView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ GridView) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _GridViewClass) Alloc() GridView {
	rv := objc.Send[GridView](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _GridViewClass) New() GridView {
	rv := objc.Send[GridView](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewGridView creates and returns a new initialized instance.
func NewGridView() GridView {
	return GridViewClass.New()
}

// Init initializes the instance.
func (g_ GridView) Init() GridView {
	rv := objc.Send[GridView](g_.ID(), selInit)
	return rv
}
// Creates a newly allocated grid view object with the specified number of columns and rows. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/init(numberOfColumns:rows:)
func (gc _GridViewClass) GridViewWithNumberOfColumnsRows(columnCount int, rowCount int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.Class), objc.RegisterName("gridViewWithNumberOfColumns:rows:"), columnCount, rowCount)
	return rv
}

// GridView_GridViewWithNumberOfColumnsRows creates a new instance via class method. [Full Topic]
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/init(numberOfColumns:rows:)
func GridView_GridViewWithNumberOfColumnsRows(columnCount int, rowCount int) unsafe.Pointer {
	return GridViewClass.GridViewWithNumberOfColumnsRows(columnCount, rowCount)
}
// Creates a newly allocated grid view object with the specified array of arrays of views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/init(views:)
func (gc _GridViewClass) GridViewWithViews(rows unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.Class), objc.RegisterName("gridViewWithViews:"), rows)
	return rv
}

// GridView_GridViewWithViews creates a new instance via class method. [Full Topic]
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/init(views:)
func GridView_GridViewWithViews(rows unsafe.Pointer) unsafe.Pointer {
	return GridViewClass.GridViewWithViews(rows)
}
// Adds a new column containing the array of views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/addColumn(with:)
func (g_ GridView) AddColumnWithViews(views unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("addColumnWithViews:"), views)
	return rv
}
// Adds an array of views to a new row. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/addRow(with:)
func (g_ GridView) AddRowWithViews(views unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("addRowWithViews:"), views)
	return rv
}
// Returns the grid cell object at the specified column and row index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/cell(atColumnIndex:rowIndex:)
func (g_ GridView) CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("cellAtColumnIndex:rowIndex:"), columnIndex, rowIndex)
	return rv
}
// Returns the grid cell object that contains the given view or one of its ancestors. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/cell(for:)
func (g_ GridView) CellForView(view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("cellForView:"), view)
	return rv
}
// Returns the grid column object at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/column(at:)
func (g_ GridView) ColumnAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("columnAtIndex:"), index)
	return rv
}
// Returns the index of the specified grid column. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/index(of:)-32sdd
func (g_ GridView) IndexOfColumn(column unsafe.Pointer) int {
	rv := objc.Send[int](g_.ID(), objc.RegisterName("indexOfColumn:"), column)
	return rv
}
// Returns the index of the specified grid row. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/index(of:)-6zs2o
func (g_ GridView) IndexOfRow(row unsafe.Pointer) int {
	rv := objc.Send[int](g_.ID(), objc.RegisterName("indexOfRow:"), row)
	return rv
}
// Creates a newly allocated grid view object from the coder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/init(coder:)
func (g_ GridView) InitWithCoder(coder unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("initWithCoder:"), coder)
	return rv
}
// Creates a newly allocated grid view object with the specified frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/init(frame:)
func (g_ GridView) InitWithFrame(frameRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("initWithFrame:"), frameRect)
	return rv
}
// Inserts the array of view objects at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/insertColumn(at:with:)
func (g_ GridView) InsertColumnAtIndexWithViews(index int, views unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("insertColumnAtIndex:withViews:"), index, views)
	return rv
}
// Inserts the array of view objects into the grid view at the index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/insertRow(at:with:)
func (g_ GridView) InsertRowAtIndexWithViews(index int, views unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("insertRowAtIndex:withViews:"), index, views)
	return rv
}
// Expands the cell at the top-leading corner of the horizontal and vertical range to cover the entire area. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/mergeCells(inHorizontalRange:verticalRange:)
func (g_ GridView) MergeCellsInHorizontalRangeVerticalRange(hRange unsafe.Pointer, vRange unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("mergeCellsInHorizontalRange:verticalRange:"), hRange, vRange)
}
// Moves the specified column to a new column location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/moveColumn(at:to:)
func (g_ GridView) MoveColumnAtIndexToIndex(fromIndex int, toIndex int) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("moveColumnAtIndex:toIndex:"), fromIndex, toIndex)
}
// Moves the specified row to the new row location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/moveRow(at:to:)
func (g_ GridView) MoveRowAtIndexToIndex(fromIndex int, toIndex int) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("moveRowAtIndex:toIndex:"), fromIndex, toIndex)
}
// Removes the column from the grid view at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/removeColumn(at:)
func (g_ GridView) RemoveColumnAtIndex(index int) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("removeColumnAtIndex:"), index)
}
// Removes the row from the grid view at the index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/removeRow(at:)
func (g_ GridView) RemoveRowAtIndex(index int) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("removeRowAtIndex:"), index)
}
// Returns the grid row object at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/row(at:)
func (g_ GridView) RowAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("rowAtIndex:"), index)
	return rv
}
// The column spacing for the grid view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/columnSpacing
func (g_ GridView) ColumnSpacing() float64 {
	rv := objc.Send[float64](g_.ID(), objc.RegisterName("columnSpacing"))
	return rv
}
// SetColumnSpacing sets the value of the columnSpacing property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/columnSpacing
func (g_ GridView) SetColumnSpacing(value float64) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setColumnSpacing:"), value)
}
// The number of columns in the grid view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/numberOfColumns
func (g_ GridView) NumberOfColumns() int {
	rv := objc.Send[int](g_.ID(), objc.RegisterName("numberOfColumns"))
	return rv
}
// The number of rows in the grid view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/numberOfRows
func (g_ GridView) NumberOfRows() int {
	rv := objc.Send[int](g_.ID(), objc.RegisterName("numberOfRows"))
	return rv
}
// The row alignment for the grid view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/rowAlignment
func (g_ GridView) RowAlignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("rowAlignment"))
	return rv
}
// SetRowAlignment sets the value of the rowAlignment property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/rowAlignment
func (g_ GridView) SetRowAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setRowAlignment:"), value)
}
// The row spacing for the grid view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/rowSpacing
func (g_ GridView) RowSpacing() float64 {
	rv := objc.Send[float64](g_.ID(), objc.RegisterName("rowSpacing"))
	return rv
}
// SetRowSpacing sets the value of the rowSpacing property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/rowSpacing
func (g_ GridView) SetRowSpacing(value float64) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setRowSpacing:"), value)
}
// The placement of the cell within the grid column. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/xPlacement
func (g_ GridView) XPlacement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("xPlacement"))
	return rv
}
// SetXPlacement sets the value of the xPlacement property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/xPlacement
func (g_ GridView) SetXPlacement(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setXPlacement:"), value)
}
// The placement of the cell within the grid row. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/yPlacement
func (g_ GridView) YPlacement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("yPlacement"))
	return rv
}
// SetYPlacement sets the value of the yPlacement property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGridView/yPlacement
func (g_ GridView) SetYPlacement(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setYPlacement:"), value)
}
