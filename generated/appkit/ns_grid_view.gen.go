// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSGridView */


/* debug [class_header]: Header for NSGridView */
// The class instance for the [GridView] class.
var (
	GridViewClass     _GridViewClass
	GridViewClassOnce sync.Once
)

func getGridViewClass() _GridViewClass {
	GridViewClassOnce.Do(func() {
		GridViewClass = _GridViewClass{objc.GetClass("NSGridView")}
	})
	return GridViewClass
}

type _GridViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GridView */
// An interface definition for the [GridView] class.
type IGridView interface {
	IView
	
/* debug [class_interface_properties]: Properties for GridView */
	// properties:
	ColumnSpacing() float64
	SetColumnSpacing(value float64)
	NumberOfRows() int
	NumberOfColumns() int
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	RowSpacing() float64
	SetRowSpacing(value float64)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GridView */
	// methods:
	AddColumnWithViews(views []View) IGridColumn
	AddRowWithViews(views []View) IGridRow
	CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) IGridCell
	CellForView(view IView) IGridCell
	ColumnAtIndex(index int) IGridColumn
	IndexOfColumn(column IGridColumn) int
	IndexOfRow(row IGridRow) int
	InsertColumnAtIndexWithViews(index int, views []View) IGridColumn
	InsertRowAtIndexWithViews(index int, views []View) IGridRow
	MergeCellsInHorizontalRangeVerticalRange(hRange corefoundation.Range, vRange corefoundation.Range)
	MoveColumnAtIndexToIndex(fromIndex int, toIndex int)
	MoveRowAtIndexToIndex(fromIndex int, toIndex int)
	RemoveColumnAtIndex(index int)
	RemoveRowAtIndex(index int)
	RowAtIndex(index int) IGridRow
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GridView */
// Alloc allocates a new instance without initialization.
func (gc _GridViewClass) Alloc() GridView {
	rv := objc.Send[GridView](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GridView */
// A container that aligns views in a flexible grid of rows and columns.
//
// A grid view helps you lay out content, such as photos or thumbnails, in a row-column arrangement similar to a spreadsheet. Within a grid view, an item that occupies a single row-column intersection is represented by an object.


// A container that aligns views in a flexible grid of rows and columns.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GridView */

// Creates a newly allocated grid view object from the coder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(coder:)
func NewGridViewWithCoder(coder foundation.Coder) GridView {
	instance := getGridViewClass().Alloc()
	rv := objc.Send[GridView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGridViewWithCoder */


// Creates a newly allocated grid view object with the specified frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(frame:)
func NewGridViewWithFrame(frameRect Rect /* not a class type */) GridView {
	instance := getGridViewClass().Alloc()
	rv := objc.Send[GridView](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGridViewWithFrame */


// Creates a newly allocated grid view object with the specified number of columns and rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(numberOfColumns:rows:)
func NewGridViewWithNumberOfColumnsRows(columnCount int, rowCount int) GridView {
	rv := objc.Send[GridView](objc.ID(getGridViewClass().class), objc.Sel("gridViewWithNumberOfColumns:rows:"), columnCount, rowCount)
	return rv
}/* debug [class_init_methods/constructor]: NewGridViewWithNumberOfColumnsRows */


// Creates a newly allocated grid view object with the specified array of arrays of views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(views:)
func NewGridViewWithViews(rows []foundation.Array) GridView {
	rv := objc.Send[GridView](objc.ID(getGridViewClass().class), objc.Sel("gridViewWithViews:"), rows)
	return rv
}/* debug [class_init_methods/constructor]: NewGridViewWithViews */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GridView */

// Creates a newly allocated grid view object with the specified number of columns and rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(numberOfColumns:rows:)
func (gc _GridViewClass) GridViewWithNumberOfColumnsRows(columnCount int, rowCount int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("gridViewWithNumberOfColumns:rows:"), columnCount, rowCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GridViewWithNumberOfColumnsRows) */


// Creates a newly allocated grid view object with the specified array of arrays of views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/init(views:)
func (gc _GridViewClass) GridViewWithViews(rows []foundation.Array) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("gridViewWithViews:"), rows)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GridViewWithViews) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GridView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GridView */

// Adds a new column containing the array of views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/addColumn(with:)
func (g_ GridView) AddColumnWithViews(views []View) IGridColumn {
	rv := objc.Send[GridColumn](g_.ID, objc.Sel("addColumnWithViews:"), views)
	return rv
}/* debug [instance_methods/method]: AddColumnWithViews */


// Adds an array of views to a new row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/addRow(with:)
func (g_ GridView) AddRowWithViews(views []View) IGridRow {
	rv := objc.Send[GridRow](g_.ID, objc.Sel("addRowWithViews:"), views)
	return rv
}/* debug [instance_methods/method]: AddRowWithViews */


// Returns the grid cell object at the specified column and row index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/cell(atColumnIndex:rowIndex:)
func (g_ GridView) CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) IGridCell {
	rv := objc.Send[GridCell](g_.ID, objc.Sel("cellAtColumnIndex:rowIndex:"), columnIndex, rowIndex)
	return rv
}/* debug [instance_methods/method]: CellAtColumnIndexRowIndex */


// Returns the grid cell object that contains the given view or one of its ancestors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/cell(for:)
func (g_ GridView) CellForView(view IView) IGridCell {
	rv := objc.Send[GridCell](g_.ID, objc.Sel("cellForView:"), view)
	return rv
}/* debug [instance_methods/method]: CellForView */


// Returns the grid column object at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/column(at:)
func (g_ GridView) ColumnAtIndex(index int) IGridColumn {
	rv := objc.Send[GridColumn](g_.ID, objc.Sel("columnAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ColumnAtIndex */


// Returns the index of the specified grid column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/index(of:)-32sdd
func (g_ GridView) IndexOfColumn(column IGridColumn) int {
	rv := objc.Send[int](g_.ID, objc.Sel("indexOfColumn:"), column)
	return rv
}/* debug [instance_methods/method]: IndexOfColumn */


// Returns the index of the specified grid row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/index(of:)-6zs2o
func (g_ GridView) IndexOfRow(row IGridRow) int {
	rv := objc.Send[int](g_.ID, objc.Sel("indexOfRow:"), row)
	return rv
}/* debug [instance_methods/method]: IndexOfRow */


// Inserts the array of view objects at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/insertColumn(at:with:)
func (g_ GridView) InsertColumnAtIndexWithViews(index int, views []View) IGridColumn {
	rv := objc.Send[GridColumn](g_.ID, objc.Sel("insertColumnAtIndex:withViews:"), index, views)
	return rv
}/* debug [instance_methods/method]: InsertColumnAtIndexWithViews */


// Inserts the array of view objects into the grid view at the index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/insertRow(at:with:)
func (g_ GridView) InsertRowAtIndexWithViews(index int, views []View) IGridRow {
	rv := objc.Send[GridRow](g_.ID, objc.Sel("insertRowAtIndex:withViews:"), index, views)
	return rv
}/* debug [instance_methods/method]: InsertRowAtIndexWithViews */


// Expands the cell at the top-leading corner of the horizontal and vertical range to cover the entire area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/mergeCells(inHorizontalRange:verticalRange:)
func (g_ GridView) MergeCellsInHorizontalRangeVerticalRange(hRange corefoundation.Range, vRange corefoundation.Range) {
	objc.Send[objc.ID](g_.ID, objc.Sel("mergeCellsInHorizontalRange:verticalRange:"), hRange, vRange)
}/* debug [instance_methods/method]: MergeCellsInHorizontalRangeVerticalRange */


// Moves the specified column to a new column location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/moveColumn(at:to:)
func (g_ GridView) MoveColumnAtIndexToIndex(fromIndex int, toIndex int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("moveColumnAtIndex:toIndex:"), fromIndex, toIndex)
}/* debug [instance_methods/method]: MoveColumnAtIndexToIndex */


// Moves the specified row to the new row location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/moveRow(at:to:)
func (g_ GridView) MoveRowAtIndexToIndex(fromIndex int, toIndex int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("moveRowAtIndex:toIndex:"), fromIndex, toIndex)
}/* debug [instance_methods/method]: MoveRowAtIndexToIndex */


// Removes the column from the grid view at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/removeColumn(at:)
func (g_ GridView) RemoveColumnAtIndex(index int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("removeColumnAtIndex:"), index)
}/* debug [instance_methods/method]: RemoveColumnAtIndex */


// Removes the row from the grid view at the index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/removeRow(at:)
func (g_ GridView) RemoveRowAtIndex(index int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("removeRowAtIndex:"), index)
}/* debug [instance_methods/method]: RemoveRowAtIndex */


// Returns the grid row object at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/row(at:)
func (g_ GridView) RowAtIndex(index int) IGridRow {
	rv := objc.Send[GridRow](g_.ID, objc.Sel("rowAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: RowAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GridView */

// The column spacing for the grid view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/columnSpacing
func (g_ GridView) ColumnSpacing() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("columnSpacing"))
	return rv
}/* debug [instance_properties/getter]: columnSpacing */


// The column spacing for the grid view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/columnSpacing
func (g_ GridView) SetColumnSpacing(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColumnSpacing:"), value)
}/* debug [instance_properties/setter]: columnSpacing */


// The number of rows in the grid view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/numberOfRows
func (g_ GridView) NumberOfRows() int {
	rv := objc.Send[int](g_.ID, objc.Sel("numberOfRows"))
	return rv
}/* debug [instance_properties/getter]: numberOfRows */


// The number of columns in the grid view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/numberOfColumns
func (g_ GridView) NumberOfColumns() int {
	rv := objc.Send[int](g_.ID, objc.Sel("numberOfColumns"))
	return rv
}/* debug [instance_properties/getter]: numberOfColumns */


// The row alignment for the grid view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/rowAlignment
func (g_ GridView) RowAlignment() GridRowAlignment {
	rv := objc.Send[GridRowAlignment](g_.ID, objc.Sel("rowAlignment"))
	return rv
}/* debug [instance_properties/getter]: rowAlignment */


// The row alignment for the grid view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/rowAlignment
func (g_ GridView) SetRowAlignment(value GridRowAlignment) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRowAlignment:"), value)
}/* debug [instance_properties/setter]: rowAlignment */


// The row spacing for the grid view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/rowSpacing
func (g_ GridView) RowSpacing() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("rowSpacing"))
	return rv
}/* debug [instance_properties/getter]: rowSpacing */


// The row spacing for the grid view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/rowSpacing
func (g_ GridView) SetRowSpacing(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRowSpacing:"), value)
}/* debug [instance_properties/setter]: rowSpacing */


// The placement of the cell within the grid column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/xPlacement
func (g_ GridView) XPlacement() GridCellPlacement {
	rv := objc.Send[GridCellPlacement](g_.ID, objc.Sel("xPlacement"))
	return rv
}/* debug [instance_properties/getter]: xPlacement */


// The placement of the cell within the grid column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/xPlacement
func (g_ GridView) SetXPlacement(value GridCellPlacement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setXPlacement:"), value)
}/* debug [instance_properties/setter]: xPlacement */


// The placement of the cell within the grid row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/yPlacement
func (g_ GridView) YPlacement() GridCellPlacement {
	rv := objc.Send[GridCellPlacement](g_.ID, objc.Sel("yPlacement"))
	return rv
}/* debug [instance_properties/getter]: yPlacement */


// The placement of the cell within the grid row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridView/yPlacement
func (g_ GridView) SetYPlacement(value GridCellPlacement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setYPlacement:"), value)
}/* debug [instance_properties/setter]: yPlacement */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSGridView */


