// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSTableView */


/* debug [class_header]: Header for NSTableView */
// The class instance for the [TableView] class.
var (
	TableViewClass     _TableViewClass
	TableViewClassOnce sync.Once
)

func getTableViewClass() _TableViewClass {
	TableViewClassOnce.Do(func() {
		TableViewClass = _TableViewClass{objc.GetClass("NSTableView")}
	})
	return TableViewClass
}

type _TableViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TableView */
// An interface definition for the [TableView] class.
type ITableView interface {
	IControl
	
/* debug [class_interface_properties]: Properties for TableView */
	// properties:
	AllowsColumnReordering() bool
	SetAllowsColumnReordering(value bool)
	AllowsColumnResizing() bool
	SetAllowsColumnResizing(value bool)
	AllowsColumnSelection() bool
	SetAllowsColumnSelection(value bool)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	AllowsTypeSelect() bool
	SetAllowsTypeSelect(value bool)
	AutosaveName() TableViewAutosaveName /* typedef */
	SetAutosaveName(value TableViewAutosaveName /* typedef */)
	AutosaveTableColumns() bool
	SetAutosaveTableColumns(value bool)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	ClickedColumn() int
	ClickedRow() int
	ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle
	SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle)
	CornerView() IView
	SetCornerView(value IView)
	DataSource() unsafe.Pointer
	SetDataSource(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DoubleAction() objc.SEL
	SetDoubleAction(value objc.SEL)
	DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle)
	EditedColumn() int
	EditedRow() int
	EffectiveRowSizeStyle() TableViewRowSizeStyle
	EffectiveStyle() TableViewStyle
	FloatsGroupRows() bool
	SetFloatsGroupRows(value bool)
	GridColor() IColor
	SetGridColor(value IColor)
	GridStyleMask() TableViewGridLineStyle
	SetGridStyleMask(value TableViewGridLineStyle)
	HeaderView() ITableHeaderView
	SetHeaderView(value ITableHeaderView)
	HiddenRowIndexes() foundation.IndexSet
	HighlightedTableColumn() ITableColumn
	SetHighlightedTableColumn(value ITableColumn)
	IntercellSpacing() Size /* not a class type */
	SetIntercellSpacing(value Size /* not a class type */)
	NumberOfColumns() int
	NumberOfRows() int
	NumberOfSelectedColumns() int
	NumberOfSelectedRows() int
	RegisteredNibsByIdentifier() foundation.IDictionary
	RowActionsVisible() bool
	SetRowActionsVisible(value bool)
	RowHeight() float64
	SetRowHeight(value float64)
	RowSizeStyle() TableViewRowSizeStyle
	SetRowSizeStyle(value TableViewRowSizeStyle)
	SelectedColumn() int
	SelectedColumnIndexes() foundation.IndexSet
	SelectedRow() int
	SelectedRowIndexes() foundation.IndexSet
	SelectionHighlightStyle() TableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle)
	SortDescriptors() []objc.IObject
	SetSortDescriptors(value []objc.IObject)
	Style() TableViewStyle
	SetStyle(value TableViewStyle)
	TableColumns() []TableColumn
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	UsesAlternatingRowBackgroundColors() bool
	SetUsesAlternatingRowBackgroundColors(value bool)
	UsesAutomaticRowHeights() bool
	SetUsesAutomaticRowHeights(value bool)
	UsesStaticContents() bool
	SetUsesStaticContents(value bool)
	VerticalMotionCanBeginDrag() bool
	SetVerticalMotionCanBeginDrag(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TableView */
	// methods:
	AddTableColumn(tableColumn ITableColumn)
	BeginUpdates()
	CanDragRowsWithIndexesAtPoint(rowIndexes foundation.IndexSet, mouseDownPoint vision.Point) bool
	ColumnAtPoint(point vision.Point) int
	ColumnForView(view IView) int
	ColumnWithIdentifier(identifier UserInterfaceItemIdentifier /* typedef */) int
	ColumnIndexesInRect(rect Rect /* not a class type */) foundation.IndexSet
	DeselectAll(sender objc.IObject)
	DeselectColumn(column int)
	DeselectRow(row int)
	DidAddRowViewForRow(rowView ITableRowView, row int)
	DidRemoveRowViewForRow(rowView ITableRowView, row int)
	DragImageForRowsWithIndexesTableColumnsEventOffset(dragRows foundation.IndexSet, tableColumns []TableColumn, dragEvent IEvent, dragImageOffset PointPointer /* not a class type */) IImage
	DrawBackgroundInClipRect(clipRect Rect /* not a class type */)
	DrawGridInClipRect(clipRect Rect /* not a class type */)
	DrawRowClipRect(row int, clipRect Rect /* not a class type */)
	EditColumnRowWithEventSelect(column int, row int, event IEvent, select_ bool)
	EndUpdates()
	EnumerateAvailableRowViewsUsingBlock(handler unsafe.Pointer)
	FrameOfCellAtColumnRow(column int, row int) Rect /* not a class type */
	HideRowsAtIndexesWithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions)
	HighlightSelectionInClipRect(clipRect Rect /* not a class type */)
	IndicatorImageInTableColumn(tableColumn ITableColumn) IImage
	InsertRowsAtIndexesWithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions)
	IsColumnSelected(column int) bool
	IsRowSelected(row int) bool
	MakeViewWithIdentifierOwner(identifier UserInterfaceItemIdentifier /* typedef */, owner objc.IObject) IView
	MoveColumnToColumn(oldIndex int, newIndex int)
	MoveRowAtIndexToIndex(oldIndex int, newIndex int)
	NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IndexSet)
	NoteNumberOfRowsChanged()
	RectOfColumn(column int) Rect /* not a class type */
	RectOfRow(row int) Rect /* not a class type */
	RegisterNibForIdentifier(nib INib, identifier UserInterfaceItemIdentifier /* typedef */)
	ReloadData()
	ReloadDataForRowIndexesColumnIndexes(rowIndexes foundation.IndexSet, columnIndexes foundation.IndexSet)
	RemoveRowsAtIndexesWithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions)
	RemoveTableColumn(tableColumn ITableColumn)
	RowAtPoint(point vision.Point) int
	RowForView(view IView) int
	RowsInRect(rect Rect /* not a class type */) corefoundation.Range
	RowViewAtRowMakeIfNecessary(row int, makeIfNecessary bool) ITableRowView
	ScrollColumnToVisible(column int)
	ScrollRowToVisible(row int)
	SelectAll(sender objc.IObject)
	SelectColumnIndexesByExtendingSelection(indexes foundation.IndexSet, extend bool)
	SelectRowIndexesByExtendingSelection(indexes foundation.IndexSet, extend bool)
	SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool)
	SetDropRowDropOperation(row int, dropOperation TableViewDropOperation)
	SetIndicatorImageInTableColumn(image IImage, tableColumn ITableColumn)
	SizeLastColumnToFit()
	SizeToFit()
	TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier /* typedef */) ITableColumn
	Tile()
	UnhideRowsAtIndexesWithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions)
	ViewAtColumnRowMakeIfNecessary(column int, row int, makeIfNecessary bool) IView
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TableView */
// Alloc allocates a new instance without initialization.
func (tc _TableViewClass) Alloc() TableView {
	rv := objc.Send[TableView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TableView */
// A set of related records, displayed in rows that represent individual records and columns that represent the attributes of those records.
//
// Table views are displayed in scroll views. Beginning with macOS v10.7, you can use objects (most commonly customized objects) instead of cells for specifying rows and columns. You can still use objects for each row and column item if you prefer. A table view does not store its own data; it retrieves data values as needed from a data source to which it has a weak reference. You should not, therefore, directly set data values programmatically in the table view; instead, modify the values in the data source and allow the changes to be reflected in the table view. To learn about the methods that an object uses to provide and access the contents of its data source object, see . To customize a table view’s behavior without subclassing , use the methods defined by the protocol. For example, the delegate supports table column management, type-to-select functionality, row selection and editing, custom tracking, and custom views for individual columns and rows. To learn more about the table view delegate, see .


// A set of related records, displayed in rows that represent individual records and columns that represent the attributes of those records.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TableView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/init(coder:)
func NewTableViewWithCoder(coder foundation.Coder) TableView {
	instance := getTableViewClass().Alloc()
	rv := objc.Send[TableView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTableViewWithCoder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/init(frame:)
func NewTableViewWithFrame(frameRect Rect /* not a class type */) TableView {
	instance := getTableViewClass().Alloc()
	rv := objc.Send[TableView](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTableViewWithFrame */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TableView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TableView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TableView */

// Adds the specified column as the last column of the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/addTableColumn(_:)
func (t_ TableView) AddTableColumn(tableColumn ITableColumn) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addTableColumn:"), tableColumn)
}/* debug [instance_methods/method]: AddTableColumn */


// Begins a group of updates for the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/beginUpdates()
func (t_ TableView) BeginUpdates() {
	objc.Send[objc.ID](t_.ID, objc.Sel("beginUpdates"))
}/* debug [instance_methods/method]: BeginUpdates */


// Returns a Boolean value indicating whether the table view allows dragging the rows with the drag initiated at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/canDragRows(with:at:)
func (t_ TableView) CanDragRowsWithIndexesAtPoint(rowIndexes foundation.IndexSet, mouseDownPoint vision.Point) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canDragRowsWithIndexes:atPoint:"), rowIndexes, mouseDownPoint)
	return rv
}/* debug [instance_methods/method]: CanDragRowsWithIndexesAtPoint */


// Returns the index of the column the specified point lies in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/column(at:)
func (t_ TableView) ColumnAtPoint(point vision.Point) int {
	rv := objc.Send[int](t_.ID, objc.Sel("columnAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: ColumnAtPoint */


// Returns the column index for the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/column(for:)
func (t_ TableView) ColumnForView(view IView) int {
	rv := objc.Send[int](t_.ID, objc.Sel("columnForView:"), view)
	return rv
}/* debug [instance_methods/method]: ColumnForView */


// Returns the index of the first column in the table view whose identifier is equal to the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/column(withIdentifier:)
func (t_ TableView) ColumnWithIdentifier(identifier UserInterfaceItemIdentifier /* typedef */) int {
	rv := objc.Send[int](t_.ID, objc.Sel("columnWithIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: ColumnWithIdentifier */


// Returns the indexes of the table view’s columns that intersect the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/columnIndexes(in:)
func (t_ TableView) ColumnIndexesInRect(rect Rect /* not a class type */) foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](t_.ID, objc.Sel("columnIndexesInRect:"), rect)
	return rv
}/* debug [instance_methods/method]: ColumnIndexesInRect */


// Deselects all selected rows or columns if empty selection is allowed; otherwise does nothing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/deselectAll(_:)
func (t_ TableView) DeselectAll(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("deselectAll:"), sender)
}/* debug [instance_methods/method]: DeselectAll */


// Deselects the column at the specified index if it’s selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/deselectColumn(_:)
func (t_ TableView) DeselectColumn(column int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("deselectColumn:"), column)
}/* debug [instance_methods/method]: DeselectColumn */


// Deselects the row at the specified index if it’s selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/deselectRow(_:)
func (t_ TableView) DeselectRow(row int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("deselectRow:"), row)
}/* debug [instance_methods/method]: DeselectRow */


// Invoked when a row view is added to the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/didAdd(_:forRow:)
func (t_ TableView) DidAddRowViewForRow(rowView ITableRowView, row int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("didAddRowView:forRow:"), rowView, row)
}/* debug [instance_methods/method]: DidAddRowViewForRow */


// Invoked when a row view is removed from the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/didRemove(_:forRow:)
func (t_ TableView) DidRemoveRowViewForRow(rowView ITableRowView, row int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("didRemoveRowView:forRow:"), rowView, row)
}/* debug [instance_methods/method]: DidRemoveRowViewForRow */


// Computes and returns an image to use for dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/dragImageForRows(with:tableColumns:event:offset:)
func (t_ TableView) DragImageForRowsWithIndexesTableColumnsEventOffset(dragRows foundation.IndexSet, tableColumns []TableColumn, dragEvent IEvent, dragImageOffset PointPointer /* not a class type */) IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("dragImageForRowsWithIndexes:tableColumns:event:offset:"), dragRows, tableColumns, dragEvent, dragImageOffset)
	return rv
}/* debug [instance_methods/method]: DragImageForRowsWithIndexesTableColumnsEventOffset */


// Draws the background of the table view in the clip rect specified by the rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/drawBackground(inClipRect:)
func (t_ TableView) DrawBackgroundInClipRect(clipRect Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawBackgroundInClipRect:"), clipRect)
}/* debug [instance_methods/method]: DrawBackgroundInClipRect */


// Draws the grid lines within the supplied rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/drawGrid(inClipRect:)
func (t_ TableView) DrawGridInClipRect(clipRect Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawGridInClipRect:"), clipRect)
}/* debug [instance_methods/method]: DrawGridInClipRect */


// Draws the cells for the row at in the columns that intersect .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/drawRow(_:clipRect:)
func (t_ TableView) DrawRowClipRect(row int, clipRect Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawRow:clipRect:"), row, clipRect)
}/* debug [instance_methods/method]: DrawRowClipRect */


// Edits the cell at the specified column and row using the specified event and selection behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/editColumn(_:row:with:select:)
func (t_ TableView) EditColumnRowWithEventSelect(column int, row int, event IEvent, select_ bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("editColumn:row:withEvent:select:"), column, row, event, select_)
}/* debug [instance_methods/method]: EditColumnRowWithEventSelect */


// Ends the group of updates for the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/endUpdates()
func (t_ TableView) EndUpdates() {
	objc.Send[objc.ID](t_.ID, objc.Sel("endUpdates"))
}/* debug [instance_methods/method]: EndUpdates */


// Allows the enumeration of all the table rows that are known to the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/enumerateAvailableRowViews(_:)
func (t_ TableView) EnumerateAvailableRowViewsUsingBlock(handler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("enumerateAvailableRowViewsUsingBlock:"), handler)
}/* debug [instance_methods/method]: EnumerateAvailableRowViewsUsingBlock */


// Returns a rectangle locating the cell that lies at the intersection of the specified column and row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/frameOfCell(atColumn:row:)
func (t_ TableView) FrameOfCellAtColumnRow(column int, row int) Rect /* not a class type */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("frameOfCellAtColumn:row:"), column, row)
	return rv
}/* debug [instance_methods/method]: FrameOfCellAtColumnRow */


// Hides the specified table rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/hideRows(at:withAnimation:)
func (t_ TableView) HideRowsAtIndexesWithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("hideRowsAtIndexes:withAnimation:"), indexes, rowAnimation)
}/* debug [instance_methods/method]: HideRowsAtIndexesWithAnimation */


// Highlights the region of the table view in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/highlightSelection(inClipRect:)
func (t_ TableView) HighlightSelectionInClipRect(clipRect Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("highlightSelectionInClipRect:"), clipRect)
}/* debug [instance_methods/method]: HighlightSelectionInClipRect */


// Returns the indicator image of the specified table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/indicatorImage(in:)
func (t_ TableView) IndicatorImageInTableColumn(tableColumn ITableColumn) IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("indicatorImageInTableColumn:"), tableColumn)
	return rv
}/* debug [instance_methods/method]: IndicatorImageInTableColumn */


// Inserts the rows using the specified animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/insertRows(at:withAnimation:)
func (t_ TableView) InsertRowsAtIndexesWithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertRowsAtIndexes:withAnimation:"), indexes, animationOptions)
}/* debug [instance_methods/method]: InsertRowsAtIndexesWithAnimation */


// Returns a Boolean value that indicates whether the column at the specified index is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/isColumnSelected(_:)
func (t_ TableView) IsColumnSelected(column int) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isColumnSelected:"), column)
	return rv
}/* debug [instance_methods/method]: IsColumnSelected */


// Returns a Boolean value that indicates whether the row at the specified index is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/isRowSelected(_:)
func (t_ TableView) IsRowSelected(row int) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRowSelected:"), row)
	return rv
}/* debug [instance_methods/method]: IsRowSelected */


// Returns a new or existing view with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/makeView(withIdentifier:owner:)
func (t_ TableView) MakeViewWithIdentifierOwner(identifier UserInterfaceItemIdentifier /* typedef */, owner objc.IObject) IView {
	rv := objc.Send[View](t_.ID, objc.Sel("makeViewWithIdentifier:owner:"), identifier, owner)
	return rv
}/* debug [instance_methods/method]: MakeViewWithIdentifierOwner */


// Moves the column and heading at the specified index to the new specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/moveColumn(_:toColumn:)
func (t_ TableView) MoveColumnToColumn(oldIndex int, newIndex int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("moveColumn:toColumn:"), oldIndex, newIndex)
}/* debug [instance_methods/method]: MoveColumnToColumn */


// Moves the specified row to the new row location using animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/moveRow(at:to:)
func (t_ TableView) MoveRowAtIndexToIndex(oldIndex int, newIndex int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("moveRowAtIndex:toIndex:"), oldIndex, newIndex)
}/* debug [instance_methods/method]: MoveRowAtIndexToIndex */


// Informs the table view that the rows specified in have changed height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/noteHeightOfRows(withIndexesChanged:)
func (t_ TableView) NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IndexSet) {
	objc.Send[objc.ID](t_.ID, objc.Sel("noteHeightOfRowsWithIndexesChanged:"), indexSet)
}/* debug [instance_methods/method]: NoteHeightOfRowsWithIndexesChanged */


// Informs the table view that the number of records in its data source has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/noteNumberOfRowsChanged()
func (t_ TableView) NoteNumberOfRowsChanged() {
	objc.Send[objc.ID](t_.ID, objc.Sel("noteNumberOfRowsChanged"))
}/* debug [instance_methods/method]: NoteNumberOfRowsChanged */


// Returns the rectangle containing the column at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rect(ofColumn:)
func (t_ TableView) RectOfColumn(column int) Rect /* not a class type */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("rectOfColumn:"), column)
	return rv
}/* debug [instance_methods/method]: RectOfColumn */


// Returns the rectangle containing the row at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rect(ofRow:)
func (t_ TableView) RectOfRow(row int) Rect /* not a class type */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("rectOfRow:"), row)
	return rv
}/* debug [instance_methods/method]: RectOfRow */


// Registers a NIB for the specified identifier, so that view-based table views can use it to instantiate views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/register(_:forIdentifier:)
func (t_ TableView) RegisterNibForIdentifier(nib INib, identifier UserInterfaceItemIdentifier /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("registerNib:forIdentifier:"), nib, identifier)
}/* debug [instance_methods/method]: RegisterNibForIdentifier */


// Marks the table view as needing redisplay, so it will reload the data for visible cells and draw the new values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/reloadData()
func (t_ TableView) ReloadData() {
	objc.Send[objc.ID](t_.ID, objc.Sel("reloadData"))
}/* debug [instance_methods/method]: ReloadData */


// Reloads the data for only the specified rows and columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/reloadData(forRowIndexes:columnIndexes:)
func (t_ TableView) ReloadDataForRowIndexesColumnIndexes(rowIndexes foundation.IndexSet, columnIndexes foundation.IndexSet) {
	objc.Send[objc.ID](t_.ID, objc.Sel("reloadDataForRowIndexes:columnIndexes:"), rowIndexes, columnIndexes)
}/* debug [instance_methods/method]: ReloadDataForRowIndexesColumnIndexes */


// Removes the rows using the specified animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/removeRows(at:withAnimation:)
func (t_ TableView) RemoveRowsAtIndexesWithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeRowsAtIndexes:withAnimation:"), indexes, animationOptions)
}/* debug [instance_methods/method]: RemoveRowsAtIndexesWithAnimation */


// Removes the specified column from the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/removeTableColumn(_:)
func (t_ TableView) RemoveTableColumn(tableColumn ITableColumn) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeTableColumn:"), tableColumn)
}/* debug [instance_methods/method]: RemoveTableColumn */


// Returns the index of the row the specified point lies in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/row(at:)
func (t_ TableView) RowAtPoint(point vision.Point) int {
	rv := objc.Send[int](t_.ID, objc.Sel("rowAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: RowAtPoint */


// Returns the index of the row for the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/row(for:)
func (t_ TableView) RowForView(view IView) int {
	rv := objc.Send[int](t_.ID, objc.Sel("rowForView:"), view)
	return rv
}/* debug [instance_methods/method]: RowForView */


// Returns a range of indexes for the rows that lie wholly or partially within the vertical boundaries of the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rows(in:)
func (t_ TableView) RowsInRect(rect Rect /* not a class type */) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rowsInRect:"), rect)
	return rv
}/* debug [instance_methods/method]: RowsInRect */


// Returns a row view at the specified index, creating one if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowView(atRow:makeIfNecessary:)
func (t_ TableView) RowViewAtRowMakeIfNecessary(row int, makeIfNecessary bool) ITableRowView {
	rv := objc.Send[TableRowView](t_.ID, objc.Sel("rowViewAtRow:makeIfNecessary:"), row, makeIfNecessary)
	return rv
}/* debug [instance_methods/method]: RowViewAtRowMakeIfNecessary */


// Scrolls the view so the specified column is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/scrollColumnToVisible(_:)
func (t_ TableView) ScrollColumnToVisible(column int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("scrollColumnToVisible:"), column)
}/* debug [instance_methods/method]: ScrollColumnToVisible */


// Scrolls the view so the specified row is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/scrollRowToVisible(_:)
func (t_ TableView) ScrollRowToVisible(row int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("scrollRowToVisible:"), row)
}/* debug [instance_methods/method]: ScrollRowToVisible */


// Selects all rows or all columns, according to whether rows or columns were most recently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectAll(_:)
func (t_ TableView) SelectAll(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectAll:"), sender)
}/* debug [instance_methods/method]: SelectAll */


// Sets the column selection using possibly extending the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectColumnIndexes(_:byExtendingSelection:)
func (t_ TableView) SelectColumnIndexesByExtendingSelection(indexes foundation.IndexSet, extend bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectColumnIndexes:byExtendingSelection:"), indexes, extend)
}/* debug [instance_methods/method]: SelectColumnIndexesByExtendingSelection */


// Sets the row selection using extending the selection if specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectRowIndexes(_:byExtendingSelection:)
func (t_ TableView) SelectRowIndexesByExtendingSelection(indexes foundation.IndexSet, extend bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectRowIndexes:byExtendingSelection:"), indexes, extend)
}/* debug [instance_methods/method]: SelectRowIndexesByExtendingSelection */


// Sets the default operation mask returned by to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/setDraggingSourceOperationMask(_:forLocal:)
func (t_ TableView) SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}/* debug [instance_methods/method]: SetDraggingSourceOperationMaskForLocal */


// Retargets the proposed drop operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/setDropRow(_:dropOperation:)
func (t_ TableView) SetDropRowDropOperation(row int, dropOperation TableViewDropOperation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDropRow:dropOperation:"), row, dropOperation)
}/* debug [instance_methods/method]: SetDropRowDropOperation */


// Sets the indicator image of the specified column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/setIndicatorImage(_:in:)
func (t_ TableView) SetIndicatorImageInTableColumn(image IImage, tableColumn ITableColumn) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIndicatorImage:inTableColumn:"), image, tableColumn)
}/* debug [instance_methods/method]: SetIndicatorImageInTableColumn */


// Resizes the last column so the table view fits exactly within its enclosing clip view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/sizeLastColumnToFit()
func (t_ TableView) SizeLastColumnToFit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("sizeLastColumnToFit"))
}/* debug [instance_methods/method]: SizeLastColumnToFit */


// Sizes the table view based on a uniform column autoresizing style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/sizeToFit()
func (t_ TableView) SizeToFit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("sizeToFit"))
}/* debug [instance_methods/method]: SizeToFit */


// Returns the object for the first column whose identifier is equal to the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/tableColumn(withIdentifier:)
func (t_ TableView) TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier /* typedef */) ITableColumn {
	rv := objc.Send[TableColumn](t_.ID, objc.Sel("tableColumnWithIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: TableColumnWithIdentifier */


// Properly sizes the table view and its header view and marks it as needing display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/tile()
func (t_ TableView) Tile() {
	objc.Send[objc.ID](t_.ID, objc.Sel("tile"))
}/* debug [instance_methods/method]: Tile */


// Unhides the specified table rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/unhideRows(at:withAnimation:)
func (t_ TableView) UnhideRowsAtIndexesWithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("unhideRowsAtIndexes:withAnimation:"), indexes, rowAnimation)
}/* debug [instance_methods/method]: UnhideRowsAtIndexesWithAnimation */


// Returns a view at the specified row and column indexes, creating one if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/view(atColumn:row:makeIfNecessary:)
func (t_ TableView) ViewAtColumnRowMakeIfNecessary(column int, row int, makeIfNecessary bool) IView {
	rv := objc.Send[View](t_.ID, objc.Sel("viewAtColumn:row:makeIfNecessary:"), column, row, makeIfNecessary)
	return rv
}/* debug [instance_methods/method]: ViewAtColumnRowMakeIfNecessary */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TableView */

// A Boolean value indicating whether the table view allows the user to rearrange columns by dragging their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnReordering
func (t_ TableView) AllowsColumnReordering() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColumnReordering"))
	return rv
}/* debug [instance_properties/getter]: allowsColumnReordering */


// A Boolean value indicating whether the table view allows the user to rearrange columns by dragging their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnReordering
func (t_ TableView) SetAllowsColumnReordering(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColumnReordering:"), value)
}/* debug [instance_properties/setter]: allowsColumnReordering */


// A Boolean value indicating whether the table view allows the user to resize columns by dragging between their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnResizing
func (t_ TableView) AllowsColumnResizing() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColumnResizing"))
	return rv
}/* debug [instance_properties/getter]: allowsColumnResizing */


// A Boolean value indicating whether the table view allows the user to resize columns by dragging between their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnResizing
func (t_ TableView) SetAllowsColumnResizing(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColumnResizing:"), value)
}/* debug [instance_properties/setter]: allowsColumnResizing */


// A Boolean value indicating whether the table view allows the user to select columns by clicking their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnSelection
func (t_ TableView) AllowsColumnSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColumnSelection"))
	return rv
}/* debug [instance_properties/getter]: allowsColumnSelection */


// A Boolean value indicating whether the table view allows the user to select columns by clicking their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnSelection
func (t_ TableView) SetAllowsColumnSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColumnSelection:"), value)
}/* debug [instance_properties/setter]: allowsColumnSelection */


// A Boolean value indicating whether the table view allows the user to select zero columns or rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsEmptySelection
func (t_ TableView) AllowsEmptySelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}/* debug [instance_properties/getter]: allowsEmptySelection */


// A Boolean value indicating whether the table view allows the user to select zero columns or rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsEmptySelection
func (t_ TableView) SetAllowsEmptySelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsEmptySelection:"), value)
}/* debug [instance_properties/setter]: allowsEmptySelection */


// A Boolean value indicating whether the table view allows the user to select more than one column or row at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsMultipleSelection
func (t_ TableView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}/* debug [instance_properties/getter]: allowsMultipleSelection */


// A Boolean value indicating whether the table view allows the user to select more than one column or row at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsMultipleSelection
func (t_ TableView) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}/* debug [instance_properties/setter]: allowsMultipleSelection */


// A Boolean value indicating whether the table view allows the user to type characters to select rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsTypeSelect
func (t_ TableView) AllowsTypeSelect() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsTypeSelect"))
	return rv
}/* debug [instance_properties/getter]: allowsTypeSelect */


// A Boolean value indicating whether the table view allows the user to type characters to select rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsTypeSelect
func (t_ TableView) SetAllowsTypeSelect(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsTypeSelect:"), value)
}/* debug [instance_properties/setter]: allowsTypeSelect */


// The name under which table information is automatically saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/autosaveName-swift.property
func (t_ TableView) AutosaveName() TableViewAutosaveName /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("autosaveName"))
	return rv
}/* debug [instance_properties/getter]: autosaveName */


// The name under which table information is automatically saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/autosaveName-swift.property
func (t_ TableView) SetAutosaveName(value TableViewAutosaveName /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutosaveName:"), value)
}/* debug [instance_properties/setter]: autosaveName */


// A Boolean value indicating whether the order and width of the table view’s columns are automatically saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/autosaveTableColumns
func (t_ TableView) AutosaveTableColumns() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("autosaveTableColumns"))
	return rv
}/* debug [instance_properties/getter]: autosaveTableColumns */


// A Boolean value indicating whether the order and width of the table view’s columns are automatically saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/autosaveTableColumns
func (t_ TableView) SetAutosaveTableColumns(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutosaveTableColumns:"), value)
}/* debug [instance_properties/setter]: autosaveTableColumns */


// The color used to draw the background of the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/backgroundColor
func (t_ TableView) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The color used to draw the background of the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/backgroundColor
func (t_ TableView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The index of the column the user clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/clickedColumn
func (t_ TableView) ClickedColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("clickedColumn"))
	return rv
}/* debug [instance_properties/getter]: clickedColumn */


// The index of the row the user clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/clickedRow
func (t_ TableView) ClickedRow() int {
	rv := objc.Send[int](t_.ID, objc.Sel("clickedRow"))
	return rv
}/* debug [instance_properties/getter]: clickedRow */


// The table view’s column autoresizing style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/columnAutoresizingStyle-swift.property
func (t_ TableView) ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle {
	rv := objc.Send[TableViewColumnAutoresizingStyle](t_.ID, objc.Sel("columnAutoresizingStyle"))
	return rv
}/* debug [instance_properties/getter]: columnAutoresizingStyle */


// The table view’s column autoresizing style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/columnAutoresizingStyle-swift.property
func (t_ TableView) SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColumnAutoresizingStyle:"), value)
}/* debug [instance_properties/setter]: columnAutoresizingStyle */


// The view used to draw the area to the right of the column headers and above the vertical scroller of the enclosing scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/cornerView
func (t_ TableView) CornerView() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("cornerView"))
	return rv
}/* debug [instance_properties/getter]: cornerView */


// The view used to draw the area to the right of the column headers and above the vertical scroller of the enclosing scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/cornerView
func (t_ TableView) SetCornerView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCornerView:"), value)
}/* debug [instance_properties/setter]: cornerView */


// The object that provides the data displayed by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/dataSource
func (t_ TableView) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// The object that provides the data displayed by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/dataSource
func (t_ TableView) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// The table view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/delegate
func (t_ TableView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The table view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/delegate
func (t_ TableView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The message sent to the table view’s target when the user double-clicks a cell or column header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/doubleAction
func (t_ TableView) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](t_.ID, objc.Sel("doubleAction"))
	return rv
}/* debug [instance_properties/getter]: doubleAction */


// The message sent to the table view’s target when the user double-clicks a cell or column header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/doubleAction
func (t_ TableView) SetDoubleAction(value objc.SEL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDoubleAction:"), value)
}/* debug [instance_properties/setter]: doubleAction */


// The feedback style displayed when the user drags over the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/draggingDestinationFeedbackStyle-swift.property
func (t_ TableView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := objc.Send[TableViewDraggingDestinationFeedbackStyle](t_.ID, objc.Sel("draggingDestinationFeedbackStyle"))
	return rv
}/* debug [instance_properties/getter]: draggingDestinationFeedbackStyle */


// The feedback style displayed when the user drags over the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/draggingDestinationFeedbackStyle-swift.property
func (t_ TableView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDraggingDestinationFeedbackStyle:"), value)
}/* debug [instance_properties/setter]: draggingDestinationFeedbackStyle */


// The index of the column being edited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/editedColumn
func (t_ TableView) EditedColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("editedColumn"))
	return rv
}/* debug [instance_properties/getter]: editedColumn */


// The index of the row being edited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/editedRow
func (t_ TableView) EditedRow() int {
	rv := objc.Send[int](t_.ID, objc.Sel("editedRow"))
	return rv
}/* debug [instance_properties/getter]: editedRow */


// The effective row size style for the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/effectiveRowSizeStyle
func (t_ TableView) EffectiveRowSizeStyle() TableViewRowSizeStyle {
	rv := objc.Send[TableViewRowSizeStyle](t_.ID, objc.Sel("effectiveRowSizeStyle"))
	return rv
}/* debug [instance_properties/getter]: effectiveRowSizeStyle */


// The effective style that the table uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/effectiveStyle
func (t_ TableView) EffectiveStyle() TableViewStyle {
	rv := objc.Send[TableViewStyle](t_.ID, objc.Sel("effectiveStyle"))
	return rv
}/* debug [instance_properties/getter]: effectiveStyle */


// A Boolean value indicating whether the table view draws grouped rows as if they are floating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/floatsGroupRows
func (t_ TableView) FloatsGroupRows() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("floatsGroupRows"))
	return rv
}/* debug [instance_properties/getter]: floatsGroupRows */


// A Boolean value indicating whether the table view draws grouped rows as if they are floating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/floatsGroupRows
func (t_ TableView) SetFloatsGroupRows(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFloatsGroupRows:"), value)
}/* debug [instance_properties/setter]: floatsGroupRows */


// The color used to draw grid lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridColor
func (t_ TableView) GridColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("gridColor"))
	return rv
}/* debug [instance_properties/getter]: gridColor */


// The color used to draw grid lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridColor
func (t_ TableView) SetGridColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGridColor:"), value)
}/* debug [instance_properties/setter]: gridColor */


// The grid lines drawn by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridStyleMask
func (t_ TableView) GridStyleMask() TableViewGridLineStyle {
	rv := objc.Send[TableViewGridLineStyle](t_.ID, objc.Sel("gridStyleMask"))
	return rv
}/* debug [instance_properties/getter]: gridStyleMask */


// The grid lines drawn by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridStyleMask
func (t_ TableView) SetGridStyleMask(value TableViewGridLineStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGridStyleMask:"), value)
}/* debug [instance_properties/setter]: gridStyleMask */


// The view object used to draw headers over columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/headerView
func (t_ TableView) HeaderView() ITableHeaderView {
	rv := objc.Send[TableHeaderView](t_.ID, objc.Sel("headerView"))
	return rv
}/* debug [instance_properties/getter]: headerView */


// The view object used to draw headers over columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/headerView
func (t_ TableView) SetHeaderView(value ITableHeaderView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeaderView:"), value)
}/* debug [instance_properties/setter]: headerView */


// The indexes of all hidden table rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/hiddenRowIndexes
func (t_ TableView) HiddenRowIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](t_.ID, objc.Sel("hiddenRowIndexes"))
	return rv
}/* debug [instance_properties/getter]: hiddenRowIndexes */


// The column highlighted in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/highlightedTableColumn
func (t_ TableView) HighlightedTableColumn() ITableColumn {
	rv := objc.Send[TableColumn](t_.ID, objc.Sel("highlightedTableColumn"))
	return rv
}/* debug [instance_properties/getter]: highlightedTableColumn */


// The column highlighted in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/highlightedTableColumn
func (t_ TableView) SetHighlightedTableColumn(value ITableColumn) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHighlightedTableColumn:"), value)
}/* debug [instance_properties/setter]: highlightedTableColumn */


// The horizontal and vertical spacing between cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/intercellSpacing
func (t_ TableView) IntercellSpacing() Size /* not a class type */ {
	rv := objc.Send[Size](t_.ID, objc.Sel("intercellSpacing"))
	return rv
}/* debug [instance_properties/getter]: intercellSpacing */


// The horizontal and vertical spacing between cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/intercellSpacing
func (t_ TableView) SetIntercellSpacing(value Size /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIntercellSpacing:"), value)
}/* debug [instance_properties/setter]: intercellSpacing */


// The number of columns in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfColumns
func (t_ TableView) NumberOfColumns() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfColumns"))
	return rv
}/* debug [instance_properties/getter]: numberOfColumns */


// The number of rows in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfRows
func (t_ TableView) NumberOfRows() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfRows"))
	return rv
}/* debug [instance_properties/getter]: numberOfRows */


// The number of selected columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfSelectedColumns
func (t_ TableView) NumberOfSelectedColumns() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfSelectedColumns"))
	return rv
}/* debug [instance_properties/getter]: numberOfSelectedColumns */


// The number of selected rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfSelectedRows
func (t_ TableView) NumberOfSelectedRows() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfSelectedRows"))
	return rv
}/* debug [instance_properties/getter]: numberOfSelectedRows */


// The dictionary of all registered nib files for view-based table view identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/registeredNibsByIdentifier
func (t_ TableView) RegisteredNibsByIdentifier() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("registeredNibsByIdentifier"))
	return rv
}/* debug [instance_properties/getter]: registeredNibsByIdentifier */


// A Boolean value indicating whether a table row’s actions are visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowActionsVisible
func (t_ TableView) RowActionsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rowActionsVisible"))
	return rv
}/* debug [instance_properties/getter]: rowActionsVisible */


// A Boolean value indicating whether a table row’s actions are visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowActionsVisible
func (t_ TableView) SetRowActionsVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowActionsVisible:"), value)
}/* debug [instance_properties/setter]: rowActionsVisible */


// The height of each row in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowHeight
func (t_ TableView) RowHeight() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("rowHeight"))
	return rv
}/* debug [instance_properties/getter]: rowHeight */


// The height of each row in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowHeight
func (t_ TableView) SetRowHeight(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowHeight:"), value)
}/* debug [instance_properties/setter]: rowHeight */


// The row size style (small, medium, large, or custom) used by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowSizeStyle-swift.property
func (t_ TableView) RowSizeStyle() TableViewRowSizeStyle {
	rv := objc.Send[TableViewRowSizeStyle](t_.ID, objc.Sel("rowSizeStyle"))
	return rv
}/* debug [instance_properties/getter]: rowSizeStyle */


// The row size style (small, medium, large, or custom) used by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowSizeStyle-swift.property
func (t_ TableView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowSizeStyle:"), value)
}/* debug [instance_properties/setter]: rowSizeStyle */


// The index of the last selected column (or the last column added to the selection).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectedColumn
func (t_ TableView) SelectedColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("selectedColumn"))
	return rv
}/* debug [instance_properties/getter]: selectedColumn */


// An index set containing the indexes of the selected columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectedColumnIndexes
func (t_ TableView) SelectedColumnIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](t_.ID, objc.Sel("selectedColumnIndexes"))
	return rv
}/* debug [instance_properties/getter]: selectedColumnIndexes */


// The index of the last selected row (or the last row added to the selection).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectedRow
func (t_ TableView) SelectedRow() int {
	rv := objc.Send[int](t_.ID, objc.Sel("selectedRow"))
	return rv
}/* debug [instance_properties/getter]: selectedRow */


// An index set containing the indexes of the selected rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectedRowIndexes
func (t_ TableView) SelectedRowIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](t_.ID, objc.Sel("selectedRowIndexes"))
	return rv
}/* debug [instance_properties/getter]: selectedRowIndexes */


// The selection highlight style used by the table view to indicate row and column selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionHighlightStyle-swift.property
func (t_ TableView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := objc.Send[TableViewSelectionHighlightStyle](t_.ID, objc.Sel("selectionHighlightStyle"))
	return rv
}/* debug [instance_properties/getter]: selectionHighlightStyle */


// The selection highlight style used by the table view to indicate row and column selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionHighlightStyle-swift.property
func (t_ TableView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionHighlightStyle:"), value)
}/* debug [instance_properties/setter]: selectionHighlightStyle */


// The table view’s sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/sortDescriptors
func (t_ TableView) SortDescriptors() []objc.IObject {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("sortDescriptors"))
	return rv
}/* debug [instance_properties/getter]: sortDescriptors */


// The table view’s sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/sortDescriptors
func (t_ TableView) SetSortDescriptors(value []objc.IObject) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setSortDescriptors:"), nsArray)
}/* debug [instance_properties/setter]: sortDescriptors */


// The style that the table view uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/style-swift.property
func (t_ TableView) Style() TableViewStyle {
	rv := objc.Send[TableViewStyle](t_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// The style that the table view uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/style-swift.property
func (t_ TableView) SetStyle(value TableViewStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */


// An array containing the current table column objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/tableColumns
func (t_ TableView) TableColumns() []TableColumn {
	rv := objc.Send[[]TableColumn](t_.ID, objc.Sel("tableColumns"))
	return rv
}/* debug [instance_properties/getter]: tableColumns */


// The layout direction of the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/userInterfaceLayoutDirection
func (t_ TableView) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](t_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}/* debug [instance_properties/getter]: userInterfaceLayoutDirection */


// The layout direction of the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/userInterfaceLayoutDirection
func (t_ TableView) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}/* debug [instance_properties/setter]: userInterfaceLayoutDirection */


// A Boolean value indicating whether the table view uses alternating row colors for its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesAlternatingRowBackgroundColors
func (t_ TableView) UsesAlternatingRowBackgroundColors() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesAlternatingRowBackgroundColors"))
	return rv
}/* debug [instance_properties/getter]: usesAlternatingRowBackgroundColors */


// A Boolean value indicating whether the table view uses alternating row colors for its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesAlternatingRowBackgroundColors
func (t_ TableView) SetUsesAlternatingRowBackgroundColors(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesAlternatingRowBackgroundColors:"), value)
}/* debug [instance_properties/setter]: usesAlternatingRowBackgroundColors */


// A Boolean value that indicates whether the table view uses autolayout to calculate the height of rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesAutomaticRowHeights
func (t_ TableView) UsesAutomaticRowHeights() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesAutomaticRowHeights"))
	return rv
}/* debug [instance_properties/getter]: usesAutomaticRowHeights */


// A Boolean value that indicates whether the table view uses autolayout to calculate the height of rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesAutomaticRowHeights
func (t_ TableView) SetUsesAutomaticRowHeights(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesAutomaticRowHeights:"), value)
}/* debug [instance_properties/setter]: usesAutomaticRowHeights */


// A Boolean value indicating whether the table uses static data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesStaticContents
func (t_ TableView) UsesStaticContents() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesStaticContents"))
	return rv
}/* debug [instance_properties/getter]: usesStaticContents */


// A Boolean value indicating whether the table uses static data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesStaticContents
func (t_ TableView) SetUsesStaticContents(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesStaticContents:"), value)
}/* debug [instance_properties/setter]: usesStaticContents */


// A Boolean value indicating whether vertical motion is treated as a drag or selection change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/verticalMotionCanBeginDrag
func (t_ TableView) VerticalMotionCanBeginDrag() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("verticalMotionCanBeginDrag"))
	return rv
}/* debug [instance_properties/getter]: verticalMotionCanBeginDrag */


// A Boolean value indicating whether vertical motion is treated as a drag or selection change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/verticalMotionCanBeginDrag
func (t_ TableView) SetVerticalMotionCanBeginDrag(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVerticalMotionCanBeginDrag:"), value)
}/* debug [instance_properties/setter]: verticalMotionCanBeginDrag */


// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/isenabled
func (t_ TableView) IsEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/isenabled
func (t_ TableView) SetIsEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTableView */


