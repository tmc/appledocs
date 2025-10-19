// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GridView] class.
var (
	gridViewClass     _GridViewClass
	gridViewClassOnce sync.Once
)

func getGridViewClass() _GridViewClass {
	gridViewClassOnce.Do(func() {
		gridViewClass = _GridViewClass{objc.GetClass("NSGridView")}
	})
	return gridViewClass
}

type _GridViewClass struct {
	class objc.Class
}

// An interface definition for the [GridView] class.
type IGridView interface {
	IView
	AddColumnWithViews(views unsafe.Pointer) unsafe.Pointer
	AddRowWithViews(views unsafe.Pointer) unsafe.Pointer
	CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) unsafe.Pointer
	CellForView(view unsafe.Pointer) unsafe.Pointer
	ColumnAtIndex(index int) unsafe.Pointer
	IndexOfColumn(column unsafe.Pointer) int
	IndexOfRow(row unsafe.Pointer) int
	InsertColumnAtIndexWithViews(index int, views unsafe.Pointer) unsafe.Pointer
	InsertRowAtIndexWithViews(index int, views unsafe.Pointer) unsafe.Pointer
	MergeCellsInHorizontalRangeVerticalRange(hRange unsafe.Pointer, vRange unsafe.Pointer)
	MoveColumnAtIndexToIndex(fromIndex int, toIndex int)
	MoveRowAtIndexToIndex(fromIndex int, toIndex int)
	RemoveColumnAtIndex(index int)
	RemoveRowAtIndex(index int)
	RowAtIndex(index int) unsafe.Pointer
}

// A container that aligns views in a flexible grid of rows and columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView

type GridView struct {
	View
}

// GridViewFrom constructs a [GridView] from an unsafe.Pointer.
//
// A container that aligns views in a flexible grid of rows and columns.
func GridViewFrom(ptr unsafe.Pointer) GridView {
	return GridView{
		View: ViewFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (gc _GridViewClass) Alloc() GridView {
	rv := objc.Send[GridView](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (gc _GridViewClass) New() GridView {
	rv := objc.Send[GridView](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GridView) Init() GridView {
	rv := objc.Send[GridView](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GridView) Autorelease() GridView {
	rv := objc.Send[GridView](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGridView creates a new GridView instance.
func NewGridView() GridView {
	return getGridViewClass().New()
}


// Creates a newly allocated grid view object from the coder. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(coder:)
func NewGridViewWithCoder(coder unsafe.Pointer) GridView {
	instance := getGridViewClass().Alloc()
	rv := objc.Send[GridView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}
// Creates a newly allocated grid view object with the specified frame rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(frame:)
func NewGridViewWithFrame(frameRect unsafe.Pointer) GridView {
	instance := getGridViewClass().Alloc()
	rv := objc.Send[GridView](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}
// Creates a newly allocated grid view object with the specified number of columns and rows. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(numberOfColumns:rows:)
func NewGridViewWithNumberOfColumnsRows(columnCount int, rowCount int) GridView {
	rv := objc.Send[GridView](objc.ID(getGridViewClass().class), objc.Sel("gridViewWithNumberOfColumns:rows:"), columnCount, rowCount)
	rv.Autorelease()
	return rv
}
// Creates a newly allocated grid view object with the specified array of arrays of views. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(views:)
func NewGridViewWithViews(rows unsafe.Pointer) GridView {
	rv := objc.Send[GridView](objc.ID(getGridViewClass().class), objc.Sel("gridViewWithViews:"), rows)
	rv.Autorelease()
	return rv
}


// Creates a newly allocated grid view object with the specified number of columns and rows. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(numberOfColumns:rows:)
func (gc _GridViewClass) GridViewWithNumberOfColumnsRows(columnCount int, rowCount int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("gridViewWithNumberOfColumns:rows:"), columnCount, rowCount)
	return rv
}
// Creates a newly allocated grid view object with the specified array of arrays of views. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(views:)
func (gc _GridViewClass) GridViewWithViews(rows unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("gridViewWithViews:"), rows)
	return rv
}
// Adds a new column containing the array of views. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/addColumn(with:)
func (g_ GridView) AddColumnWithViews(views unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("addColumnWithViews:"), views)
	return rv
}
// Adds an array of views to a new row. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/addRow(with:)
func (g_ GridView) AddRowWithViews(views unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("addRowWithViews:"), views)
	return rv
}
// Returns the grid cell object at the specified column and row index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/cell(atColumnIndex:rowIndex:)
func (g_ GridView) CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("cellAtColumnIndex:rowIndex:"), columnIndex, rowIndex)
	return rv
}
// Returns the grid cell object that contains the given view or one of its ancestors. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/cell(for:)
func (g_ GridView) CellForView(view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("cellForView:"), view)
	return rv
}
// Returns the grid column object at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/column(at:)
func (g_ GridView) ColumnAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("columnAtIndex:"), index)
	return rv
}
// Returns the index of the specified grid column. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/index(of:)-32sdd
func (g_ GridView) IndexOfColumn(column unsafe.Pointer) int {
	rv := objc.Send[int](g_.ID, objc.Sel("indexOfColumn:"), column)
	return rv
}
// Returns the index of the specified grid row. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/index(of:)-6zs2o
func (g_ GridView) IndexOfRow(row unsafe.Pointer) int {
	rv := objc.Send[int](g_.ID, objc.Sel("indexOfRow:"), row)
	return rv
}
// Inserts the array of view objects at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/insertColumn(at:with:)
func (g_ GridView) InsertColumnAtIndexWithViews(index int, views unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("insertColumnAtIndex:withViews:"), index, views)
	return rv
}
// Inserts the array of view objects into the grid view at the index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/insertRow(at:with:)
func (g_ GridView) InsertRowAtIndexWithViews(index int, views unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("insertRowAtIndex:withViews:"), index, views)
	return rv
}
// Expands the cell at the top-leading corner of the horizontal and vertical range to cover the entire area. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/mergeCells(inHorizontalRange:verticalRange:)
func (g_ GridView) MergeCellsInHorizontalRangeVerticalRange(hRange unsafe.Pointer, vRange unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("mergeCellsInHorizontalRange:verticalRange:"), hRange, vRange)
}
// Moves the specified column to a new column location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/moveColumn(at:to:)
func (g_ GridView) MoveColumnAtIndexToIndex(fromIndex int, toIndex int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("moveColumnAtIndex:toIndex:"), fromIndex, toIndex)
}
// Moves the specified row to the new row location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/moveRow(at:to:)
func (g_ GridView) MoveRowAtIndexToIndex(fromIndex int, toIndex int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("moveRowAtIndex:toIndex:"), fromIndex, toIndex)
}
// Removes the column from the grid view at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/removeColumn(at:)
func (g_ GridView) RemoveColumnAtIndex(index int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("removeColumnAtIndex:"), index)
}
// Removes the row from the grid view at the index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/removeRow(at:)
func (g_ GridView) RemoveRowAtIndex(index int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("removeRowAtIndex:"), index)
}
// Returns the grid row object at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/row(at:)
func (g_ GridView) RowAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("rowAtIndex:"), index)
	return rv
}

