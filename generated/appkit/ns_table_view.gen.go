// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/cloudkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [TableView] class.
type ITableView interface {
	IControl
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
	AutosaveName() objc.IObject /* cross-framework: TableViewAutosaveName */
	SetAutosaveName(value objc.IObject /* cross-framework: TableViewAutosaveName */)
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
	DataSource() objc.ID
	SetDataSource(value objc.ID)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
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
	IntercellSpacing() objc.IObject /* cross-framework: Size */
	SetIntercellSpacing(value objc.IObject /* cross-framework: Size */)
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
	// methods:
	AddTableColumn(tableColumn ITableColumn)
	BeginUpdates()
	CanDragRowsWithIndexesAtPoint(rowIndexes foundation.IndexSet, mouseDownPoint objc.IObject /* cross-framework: Point */) bool
	ColumnAtPoint(point objc.IObject /* cross-framework: Point */) int
	ColumnForView(view IView) int
	ColumnWithIdentifier(identifier objc.IObject /* cross-framework: UserInterfaceItemIdentifier */) int
	ColumnIndexesInRect(rect objc.IObject /* cross-framework: Rect */) foundation.IndexSet
	DeselectAll(sender objc.IObject)
	DeselectColumn(column int)
	DeselectRow(row int)
	DidAddRowViewForRow(rowView ITableRowView, row int)
	DidRemoveRowViewForRow(rowView ITableRowView, row int)
	DragImageForRowsWithIndexesTableColumnsEventOffset(dragRows foundation.IndexSet, tableColumns []TableColumn, dragEvent IEvent, dragImageOffset PointPointer /* not a class type */) IImage
	DrawBackgroundInClipRect(clipRect objc.IObject /* cross-framework: Rect */)
	DrawGridInClipRect(clipRect objc.IObject /* cross-framework: Rect */)
	DrawRowClipRect(row int, clipRect objc.IObject /* cross-framework: Rect */)
	EditColumnRowWithEventSelect(column int, row int, event IEvent, select_ bool)
	EndUpdates()
	EnumerateAvailableRowViewsUsingBlock(handler unsafe.Pointer)
	FrameOfCellAtColumnRow(column int, row int) objc.IObject /* cross-framework: Rect */
	HideRowsAtIndexesWithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions)
	HighlightSelectionInClipRect(clipRect objc.IObject /* cross-framework: Rect */)
	IndicatorImageInTableColumn(tableColumn ITableColumn) IImage
	InsertRowsAtIndexesWithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions)
	IsColumnSelected(column int) bool
	IsRowSelected(row int) bool
	MakeViewWithIdentifierOwner(identifier objc.IObject /* cross-framework: UserInterfaceItemIdentifier */, owner objc.IObject) IView
	MoveColumnToColumn(oldIndex int, newIndex int)
	MoveRowAtIndexToIndex(oldIndex int, newIndex int)
	NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IndexSet)
	NoteNumberOfRowsChanged()
	RectOfColumn(column int) objc.IObject /* cross-framework: Rect */
	RectOfRow(row int) objc.IObject /* cross-framework: Rect */
	RegisterNibForIdentifier(nib INib, identifier objc.IObject /* cross-framework: UserInterfaceItemIdentifier */)
	ReloadData()
	ReloadDataForRowIndexesColumnIndexes(rowIndexes foundation.IndexSet, columnIndexes foundation.IndexSet)
	RemoveRowsAtIndexesWithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions)
	RemoveTableColumn(tableColumn ITableColumn)
	RowAtPoint(point objc.IObject /* cross-framework: Point */) int
	RowForView(view IView) int
	RowsInRect(rect objc.IObject /* cross-framework: Rect */) corefoundation.Range
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
	TableColumnWithIdentifier(identifier objc.IObject /* cross-framework: UserInterfaceItemIdentifier */) ITableColumn
	Tile()
	UnhideRowsAtIndexesWithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions)
	ViewAtColumnRowMakeIfNecessary(column int, row int, makeIfNecessary bool) IView
}

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

// Alloc allocates a new instance without initialization.
func (tc _TableViewClass) Alloc() TableView {
	rv := objc.Send[TableView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/init(coder:)
func NewTableViewWithCoder(coder foundation.Coder) TableView {
	instance := getTableViewClass().Alloc()
	rv := objc.Send[TableView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/init(frame:)
func NewTableViewWithFrame(frameRect objc.IObject /* cross-framework: Rect */) TableView {
	instance := getTableViewClass().Alloc()
	rv := objc.Send[TableView](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}



// Adds the specified column as the last column of the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/addTableColumn(_:)
func (t_ TableView) AddTableColumn(tableColumn ITableColumn) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addTableColumn:"), tableColumn)
}


// Begins a group of updates for the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/beginUpdates()
func (t_ TableView) BeginUpdates() {
	objc.Send[objc.ID](t_.ID, objc.Sel("beginUpdates"))
}


// Returns a Boolean value indicating whether the table view allows dragging the rows with the drag initiated at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/canDragRows(with:at:)
func (t_ TableView) CanDragRowsWithIndexesAtPoint(rowIndexes foundation.IndexSet, mouseDownPoint objc.IObject /* cross-framework: Point */) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canDragRowsWithIndexes:atPoint:"), rowIndexes, mouseDownPoint)
	return rv
}


// Returns the index of the column the specified point lies in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/column(at:)
func (t_ TableView) ColumnAtPoint(point objc.IObject /* cross-framework: Point */) int {
	rv := objc.Send[int](t_.ID, objc.Sel("columnAtPoint:"), point)
	return rv
}


// Returns the column index for the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/column(for:)
func (t_ TableView) ColumnForView(view IView) int {
	rv := objc.Send[int](t_.ID, objc.Sel("columnForView:"), view)
	return rv
}


// Returns the index of the first column in the table view whose identifier is equal to the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/column(withIdentifier:)
func (t_ TableView) ColumnWithIdentifier(identifier objc.IObject /* cross-framework: UserInterfaceItemIdentifier */) int {
	rv := objc.Send[int](t_.ID, objc.Sel("columnWithIdentifier:"), identifier)
	return rv
}


// Returns the indexes of the table view’s columns that intersect the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/columnIndexes(in:)
func (t_ TableView) ColumnIndexesInRect(rect objc.IObject /* cross-framework: Rect */) foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](t_.ID, objc.Sel("columnIndexesInRect:"), rect)
	return rv
}


// Deselects all selected rows or columns if empty selection is allowed; otherwise does nothing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/deselectAll(_:)
func (t_ TableView) DeselectAll(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("deselectAll:"), sender)
}


// Deselects the column at the specified index if it’s selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/deselectColumn(_:)
func (t_ TableView) DeselectColumn(column int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("deselectColumn:"), column)
}


// Deselects the row at the specified index if it’s selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/deselectRow(_:)
func (t_ TableView) DeselectRow(row int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("deselectRow:"), row)
}


// Invoked when a row view is added to the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/didAdd(_:forRow:)
func (t_ TableView) DidAddRowViewForRow(rowView ITableRowView, row int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("didAddRowView:forRow:"), rowView, row)
}


// Invoked when a row view is removed from the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/didRemove(_:forRow:)
func (t_ TableView) DidRemoveRowViewForRow(rowView ITableRowView, row int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("didRemoveRowView:forRow:"), rowView, row)
}


// Computes and returns an image to use for dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/dragImageForRows(with:tableColumns:event:offset:)
func (t_ TableView) DragImageForRowsWithIndexesTableColumnsEventOffset(dragRows foundation.IndexSet, tableColumns []TableColumn, dragEvent IEvent, dragImageOffset PointPointer /* not a class type */) IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("dragImageForRowsWithIndexes:tableColumns:event:offset:"), dragRows, tableColumns, dragEvent, dragImageOffset)
	return rv
}


// Draws the background of the table view in the clip rect specified by the rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/drawBackground(inClipRect:)
func (t_ TableView) DrawBackgroundInClipRect(clipRect objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawBackgroundInClipRect:"), clipRect)
}


// Draws the grid lines within the supplied rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/drawGrid(inClipRect:)
func (t_ TableView) DrawGridInClipRect(clipRect objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawGridInClipRect:"), clipRect)
}


// Draws the cells for the row at in the columns that intersect .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/drawRow(_:clipRect:)
func (t_ TableView) DrawRowClipRect(row int, clipRect objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawRow:clipRect:"), row, clipRect)
}


// Edits the cell at the specified column and row using the specified event and selection behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/editColumn(_:row:with:select:)
func (t_ TableView) EditColumnRowWithEventSelect(column int, row int, event IEvent, select_ bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("editColumn:row:withEvent:select:"), column, row, event, select_)
}


// Ends the group of updates for the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/endUpdates()
func (t_ TableView) EndUpdates() {
	objc.Send[objc.ID](t_.ID, objc.Sel("endUpdates"))
}


// Allows the enumeration of all the table rows that are known to the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/enumerateAvailableRowViews(_:)
func (t_ TableView) EnumerateAvailableRowViewsUsingBlock(handler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("enumerateAvailableRowViewsUsingBlock:"), handler)
}


// Returns a rectangle locating the cell that lies at the intersection of the specified column and row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/frameOfCell(atColumn:row:)
func (t_ TableView) FrameOfCellAtColumnRow(column int, row int) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](t_.ID, objc.Sel("frameOfCellAtColumn:row:"), column, row)
	return rv
}


// Hides the specified table rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/hideRows(at:withAnimation:)
func (t_ TableView) HideRowsAtIndexesWithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("hideRowsAtIndexes:withAnimation:"), indexes, rowAnimation)
}


// Highlights the region of the table view in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/highlightSelection(inClipRect:)
func (t_ TableView) HighlightSelectionInClipRect(clipRect objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("highlightSelectionInClipRect:"), clipRect)
}


// Returns the indicator image of the specified table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/indicatorImage(in:)
func (t_ TableView) IndicatorImageInTableColumn(tableColumn ITableColumn) IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("indicatorImageInTableColumn:"), tableColumn)
	return rv
}


// Inserts the rows using the specified animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/insertRows(at:withAnimation:)
func (t_ TableView) InsertRowsAtIndexesWithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertRowsAtIndexes:withAnimation:"), indexes, animationOptions)
}


// Returns a Boolean value that indicates whether the column at the specified index is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/isColumnSelected(_:)
func (t_ TableView) IsColumnSelected(column int) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isColumnSelected:"), column)
	return rv
}


// Returns a Boolean value that indicates whether the row at the specified index is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/isRowSelected(_:)
func (t_ TableView) IsRowSelected(row int) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRowSelected:"), row)
	return rv
}


// Returns a new or existing view with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/makeView(withIdentifier:owner:)
func (t_ TableView) MakeViewWithIdentifierOwner(identifier objc.IObject /* cross-framework: UserInterfaceItemIdentifier */, owner objc.IObject) IView {
	rv := objc.Send[View](t_.ID, objc.Sel("makeViewWithIdentifier:owner:"), identifier, owner)
	return rv
}


// Moves the column and heading at the specified index to the new specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/moveColumn(_:toColumn:)
func (t_ TableView) MoveColumnToColumn(oldIndex int, newIndex int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("moveColumn:toColumn:"), oldIndex, newIndex)
}


// Moves the specified row to the new row location using animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/moveRow(at:to:)
func (t_ TableView) MoveRowAtIndexToIndex(oldIndex int, newIndex int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("moveRowAtIndex:toIndex:"), oldIndex, newIndex)
}


// Informs the table view that the rows specified in have changed height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/noteHeightOfRows(withIndexesChanged:)
func (t_ TableView) NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IndexSet) {
	objc.Send[objc.ID](t_.ID, objc.Sel("noteHeightOfRowsWithIndexesChanged:"), indexSet)
}


// Informs the table view that the number of records in its data source has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/noteNumberOfRowsChanged()
func (t_ TableView) NoteNumberOfRowsChanged() {
	objc.Send[objc.ID](t_.ID, objc.Sel("noteNumberOfRowsChanged"))
}


// Returns the rectangle containing the column at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rect(ofColumn:)
func (t_ TableView) RectOfColumn(column int) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](t_.ID, objc.Sel("rectOfColumn:"), column)
	return rv
}


// Returns the rectangle containing the row at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rect(ofRow:)
func (t_ TableView) RectOfRow(row int) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](t_.ID, objc.Sel("rectOfRow:"), row)
	return rv
}


// Registers a NIB for the specified identifier, so that view-based table views can use it to instantiate views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/register(_:forIdentifier:)
func (t_ TableView) RegisterNibForIdentifier(nib INib, identifier objc.IObject /* cross-framework: UserInterfaceItemIdentifier */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("registerNib:forIdentifier:"), nib, identifier)
}


// Marks the table view as needing redisplay, so it will reload the data for visible cells and draw the new values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/reloadData()
func (t_ TableView) ReloadData() {
	objc.Send[objc.ID](t_.ID, objc.Sel("reloadData"))
}


// Reloads the data for only the specified rows and columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/reloadData(forRowIndexes:columnIndexes:)
func (t_ TableView) ReloadDataForRowIndexesColumnIndexes(rowIndexes foundation.IndexSet, columnIndexes foundation.IndexSet) {
	objc.Send[objc.ID](t_.ID, objc.Sel("reloadDataForRowIndexes:columnIndexes:"), rowIndexes, columnIndexes)
}


// Removes the rows using the specified animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/removeRows(at:withAnimation:)
func (t_ TableView) RemoveRowsAtIndexesWithAnimation(indexes foundation.IndexSet, animationOptions TableViewAnimationOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeRowsAtIndexes:withAnimation:"), indexes, animationOptions)
}


// Removes the specified column from the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/removeTableColumn(_:)
func (t_ TableView) RemoveTableColumn(tableColumn ITableColumn) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeTableColumn:"), tableColumn)
}


// Returns the index of the row the specified point lies in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/row(at:)
func (t_ TableView) RowAtPoint(point objc.IObject /* cross-framework: Point */) int {
	rv := objc.Send[int](t_.ID, objc.Sel("rowAtPoint:"), point)
	return rv
}


// Returns the index of the row for the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/row(for:)
func (t_ TableView) RowForView(view IView) int {
	rv := objc.Send[int](t_.ID, objc.Sel("rowForView:"), view)
	return rv
}


// Returns a range of indexes for the rows that lie wholly or partially within the vertical boundaries of the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rows(in:)
func (t_ TableView) RowsInRect(rect objc.IObject /* cross-framework: Rect */) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rowsInRect:"), rect)
	return rv
}


// Returns a row view at the specified index, creating one if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowView(atRow:makeIfNecessary:)
func (t_ TableView) RowViewAtRowMakeIfNecessary(row int, makeIfNecessary bool) ITableRowView {
	rv := objc.Send[TableRowView](t_.ID, objc.Sel("rowViewAtRow:makeIfNecessary:"), row, makeIfNecessary)
	return rv
}


// Scrolls the view so the specified column is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/scrollColumnToVisible(_:)
func (t_ TableView) ScrollColumnToVisible(column int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("scrollColumnToVisible:"), column)
}


// Scrolls the view so the specified row is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/scrollRowToVisible(_:)
func (t_ TableView) ScrollRowToVisible(row int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("scrollRowToVisible:"), row)
}


// Selects all rows or all columns, according to whether rows or columns were most recently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectAll(_:)
func (t_ TableView) SelectAll(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectAll:"), sender)
}


// Sets the column selection using possibly extending the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectColumnIndexes(_:byExtendingSelection:)
func (t_ TableView) SelectColumnIndexesByExtendingSelection(indexes foundation.IndexSet, extend bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectColumnIndexes:byExtendingSelection:"), indexes, extend)
}


// Sets the row selection using extending the selection if specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectRowIndexes(_:byExtendingSelection:)
func (t_ TableView) SelectRowIndexesByExtendingSelection(indexes foundation.IndexSet, extend bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectRowIndexes:byExtendingSelection:"), indexes, extend)
}


// Sets the default operation mask returned by to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/setDraggingSourceOperationMask(_:forLocal:)
func (t_ TableView) SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}


// Retargets the proposed drop operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/setDropRow(_:dropOperation:)
func (t_ TableView) SetDropRowDropOperation(row int, dropOperation TableViewDropOperation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDropRow:dropOperation:"), row, dropOperation)
}


// Sets the indicator image of the specified column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/setIndicatorImage(_:in:)
func (t_ TableView) SetIndicatorImageInTableColumn(image IImage, tableColumn ITableColumn) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIndicatorImage:inTableColumn:"), image, tableColumn)
}


// Resizes the last column so the table view fits exactly within its enclosing clip view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/sizeLastColumnToFit()
func (t_ TableView) SizeLastColumnToFit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("sizeLastColumnToFit"))
}


// Sizes the table view based on a uniform column autoresizing style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/sizeToFit()
func (t_ TableView) SizeToFit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("sizeToFit"))
}


// Returns the object for the first column whose identifier is equal to the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/tableColumn(withIdentifier:)
func (t_ TableView) TableColumnWithIdentifier(identifier objc.IObject /* cross-framework: UserInterfaceItemIdentifier */) ITableColumn {
	rv := objc.Send[TableColumn](t_.ID, objc.Sel("tableColumnWithIdentifier:"), identifier)
	return rv
}


// Properly sizes the table view and its header view and marks it as needing display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/tile()
func (t_ TableView) Tile() {
	objc.Send[objc.ID](t_.ID, objc.Sel("tile"))
}


// Unhides the specified table rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/unhideRows(at:withAnimation:)
func (t_ TableView) UnhideRowsAtIndexesWithAnimation(indexes foundation.IndexSet, rowAnimation TableViewAnimationOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("unhideRowsAtIndexes:withAnimation:"), indexes, rowAnimation)
}


// Returns a view at the specified row and column indexes, creating one if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/view(atColumn:row:makeIfNecessary:)
func (t_ TableView) ViewAtColumnRowMakeIfNecessary(column int, row int, makeIfNecessary bool) IView {
	rv := objc.Send[View](t_.ID, objc.Sel("viewAtColumn:row:makeIfNecessary:"), column, row, makeIfNecessary)
	return rv
}


// A Boolean value indicating whether the table view allows the user to rearrange columns by dragging their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnReordering
func (t_ TableView) AllowsColumnReordering() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColumnReordering"))
	return rv
}


// A Boolean value indicating whether the table view allows the user to rearrange columns by dragging their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnReordering
func (t_ TableView) SetAllowsColumnReordering(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColumnReordering:"), value)
}


// A Boolean value indicating whether the table view allows the user to resize columns by dragging between their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnResizing
func (t_ TableView) AllowsColumnResizing() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColumnResizing"))
	return rv
}


// A Boolean value indicating whether the table view allows the user to resize columns by dragging between their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnResizing
func (t_ TableView) SetAllowsColumnResizing(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColumnResizing:"), value)
}


// A Boolean value indicating whether the table view allows the user to select columns by clicking their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnSelection
func (t_ TableView) AllowsColumnSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColumnSelection"))
	return rv
}


// A Boolean value indicating whether the table view allows the user to select columns by clicking their headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnSelection
func (t_ TableView) SetAllowsColumnSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColumnSelection:"), value)
}


// A Boolean value indicating whether the table view allows the user to select zero columns or rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsEmptySelection
func (t_ TableView) AllowsEmptySelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}


// A Boolean value indicating whether the table view allows the user to select zero columns or rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsEmptySelection
func (t_ TableView) SetAllowsEmptySelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsEmptySelection:"), value)
}


// A Boolean value indicating whether the table view allows the user to select more than one column or row at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsMultipleSelection
func (t_ TableView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}


// A Boolean value indicating whether the table view allows the user to select more than one column or row at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsMultipleSelection
func (t_ TableView) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}


// A Boolean value indicating whether the table view allows the user to type characters to select rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsTypeSelect
func (t_ TableView) AllowsTypeSelect() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsTypeSelect"))
	return rv
}


// A Boolean value indicating whether the table view allows the user to type characters to select rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/allowsTypeSelect
func (t_ TableView) SetAllowsTypeSelect(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsTypeSelect:"), value)
}


// The name under which table information is automatically saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/autosaveName-swift.property
func (t_ TableView) AutosaveName() objc.IObject /* cross-framework: TableViewAutosaveName */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("autosaveName"))
	return rv
}


// The name under which table information is automatically saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/autosaveName-swift.property
func (t_ TableView) SetAutosaveName(value objc.IObject /* cross-framework: TableViewAutosaveName */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutosaveName:"), value)
}


// A Boolean value indicating whether the order and width of the table view’s columns are automatically saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/autosaveTableColumns
func (t_ TableView) AutosaveTableColumns() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("autosaveTableColumns"))
	return rv
}


// A Boolean value indicating whether the order and width of the table view’s columns are automatically saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/autosaveTableColumns
func (t_ TableView) SetAutosaveTableColumns(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutosaveTableColumns:"), value)
}


// The color used to draw the background of the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/backgroundColor
func (t_ TableView) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The color used to draw the background of the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/backgroundColor
func (t_ TableView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The index of the column the user clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/clickedColumn
func (t_ TableView) ClickedColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("clickedColumn"))
	return rv
}


// The index of the row the user clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/clickedRow
func (t_ TableView) ClickedRow() int {
	rv := objc.Send[int](t_.ID, objc.Sel("clickedRow"))
	return rv
}


// The table view’s column autoresizing style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/columnAutoresizingStyle-swift.property
func (t_ TableView) ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle {
	rv := objc.Send[TableViewColumnAutoresizingStyle](t_.ID, objc.Sel("columnAutoresizingStyle"))
	return rv
}


// The table view’s column autoresizing style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/columnAutoresizingStyle-swift.property
func (t_ TableView) SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColumnAutoresizingStyle:"), value)
}


// The view used to draw the area to the right of the column headers and above the vertical scroller of the enclosing scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/cornerView
func (t_ TableView) CornerView() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("cornerView"))
	return rv
}


// The view used to draw the area to the right of the column headers and above the vertical scroller of the enclosing scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/cornerView
func (t_ TableView) SetCornerView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCornerView:"), value)
}


// The object that provides the data displayed by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/dataSource
func (t_ TableView) DataSource() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("dataSource"))
	return rv
}


// The object that provides the data displayed by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/dataSource
func (t_ TableView) SetDataSource(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDataSource:"), value)
}


// The table view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/delegate
func (t_ TableView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}


// The table view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/delegate
func (t_ TableView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// The message sent to the table view’s target when the user double-clicks a cell or column header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/doubleAction
func (t_ TableView) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](t_.ID, objc.Sel("doubleAction"))
	return rv
}


// The message sent to the table view’s target when the user double-clicks a cell or column header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/doubleAction
func (t_ TableView) SetDoubleAction(value objc.SEL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDoubleAction:"), value)
}


// The feedback style displayed when the user drags over the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/draggingDestinationFeedbackStyle-swift.property
func (t_ TableView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := objc.Send[TableViewDraggingDestinationFeedbackStyle](t_.ID, objc.Sel("draggingDestinationFeedbackStyle"))
	return rv
}


// The feedback style displayed when the user drags over the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/draggingDestinationFeedbackStyle-swift.property
func (t_ TableView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDraggingDestinationFeedbackStyle:"), value)
}


// The index of the column being edited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/editedColumn
func (t_ TableView) EditedColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("editedColumn"))
	return rv
}


// The index of the row being edited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/editedRow
func (t_ TableView) EditedRow() int {
	rv := objc.Send[int](t_.ID, objc.Sel("editedRow"))
	return rv
}


// The effective row size style for the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/effectiveRowSizeStyle
func (t_ TableView) EffectiveRowSizeStyle() TableViewRowSizeStyle {
	rv := objc.Send[TableViewRowSizeStyle](t_.ID, objc.Sel("effectiveRowSizeStyle"))
	return rv
}


// The effective style that the table uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/effectiveStyle
func (t_ TableView) EffectiveStyle() TableViewStyle {
	rv := objc.Send[TableViewStyle](t_.ID, objc.Sel("effectiveStyle"))
	return rv
}


// A Boolean value indicating whether the table view draws grouped rows as if they are floating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/floatsGroupRows
func (t_ TableView) FloatsGroupRows() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("floatsGroupRows"))
	return rv
}


// A Boolean value indicating whether the table view draws grouped rows as if they are floating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/floatsGroupRows
func (t_ TableView) SetFloatsGroupRows(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFloatsGroupRows:"), value)
}


// The color used to draw grid lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridColor
func (t_ TableView) GridColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("gridColor"))
	return rv
}


// The color used to draw grid lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridColor
func (t_ TableView) SetGridColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGridColor:"), value)
}


// The grid lines drawn by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridStyleMask
func (t_ TableView) GridStyleMask() TableViewGridLineStyle {
	rv := objc.Send[TableViewGridLineStyle](t_.ID, objc.Sel("gridStyleMask"))
	return rv
}


// The grid lines drawn by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridStyleMask
func (t_ TableView) SetGridStyleMask(value TableViewGridLineStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGridStyleMask:"), value)
}


// The view object used to draw headers over columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/headerView
func (t_ TableView) HeaderView() ITableHeaderView {
	rv := objc.Send[TableHeaderView](t_.ID, objc.Sel("headerView"))
	return rv
}


// The view object used to draw headers over columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/headerView
func (t_ TableView) SetHeaderView(value ITableHeaderView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeaderView:"), value)
}


// The indexes of all hidden table rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/hiddenRowIndexes
func (t_ TableView) HiddenRowIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](t_.ID, objc.Sel("hiddenRowIndexes"))
	return rv
}


// The column highlighted in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/highlightedTableColumn
func (t_ TableView) HighlightedTableColumn() ITableColumn {
	rv := objc.Send[TableColumn](t_.ID, objc.Sel("highlightedTableColumn"))
	return rv
}


// The column highlighted in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/highlightedTableColumn
func (t_ TableView) SetHighlightedTableColumn(value ITableColumn) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHighlightedTableColumn:"), value)
}


// The horizontal and vertical spacing between cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/intercellSpacing
func (t_ TableView) IntercellSpacing() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](t_.ID, objc.Sel("intercellSpacing"))
	return rv
}


// The horizontal and vertical spacing between cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/intercellSpacing
func (t_ TableView) SetIntercellSpacing(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIntercellSpacing:"), value)
}


// The number of columns in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfColumns
func (t_ TableView) NumberOfColumns() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfColumns"))
	return rv
}


// The number of rows in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfRows
func (t_ TableView) NumberOfRows() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfRows"))
	return rv
}


// The number of selected columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfSelectedColumns
func (t_ TableView) NumberOfSelectedColumns() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfSelectedColumns"))
	return rv
}


// The number of selected rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfSelectedRows
func (t_ TableView) NumberOfSelectedRows() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfSelectedRows"))
	return rv
}


// The dictionary of all registered nib files for view-based table view identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/registeredNibsByIdentifier
func (t_ TableView) RegisteredNibsByIdentifier() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("registeredNibsByIdentifier"))
	return rv
}


// A Boolean value indicating whether a table row’s actions are visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowActionsVisible
func (t_ TableView) RowActionsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rowActionsVisible"))
	return rv
}


// A Boolean value indicating whether a table row’s actions are visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowActionsVisible
func (t_ TableView) SetRowActionsVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowActionsVisible:"), value)
}


// The height of each row in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowHeight
func (t_ TableView) RowHeight() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("rowHeight"))
	return rv
}


// The height of each row in the table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowHeight
func (t_ TableView) SetRowHeight(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowHeight:"), value)
}


// The row size style (small, medium, large, or custom) used by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowSizeStyle-swift.property
func (t_ TableView) RowSizeStyle() TableViewRowSizeStyle {
	rv := objc.Send[TableViewRowSizeStyle](t_.ID, objc.Sel("rowSizeStyle"))
	return rv
}


// The row size style (small, medium, large, or custom) used by the table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowSizeStyle-swift.property
func (t_ TableView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowSizeStyle:"), value)
}


// The index of the last selected column (or the last column added to the selection).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectedColumn
func (t_ TableView) SelectedColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("selectedColumn"))
	return rv
}


// An index set containing the indexes of the selected columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectedColumnIndexes
func (t_ TableView) SelectedColumnIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](t_.ID, objc.Sel("selectedColumnIndexes"))
	return rv
}


// The index of the last selected row (or the last row added to the selection).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectedRow
func (t_ TableView) SelectedRow() int {
	rv := objc.Send[int](t_.ID, objc.Sel("selectedRow"))
	return rv
}


// An index set containing the indexes of the selected rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectedRowIndexes
func (t_ TableView) SelectedRowIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](t_.ID, objc.Sel("selectedRowIndexes"))
	return rv
}


// The selection highlight style used by the table view to indicate row and column selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionHighlightStyle-swift.property
func (t_ TableView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := objc.Send[TableViewSelectionHighlightStyle](t_.ID, objc.Sel("selectionHighlightStyle"))
	return rv
}


// The selection highlight style used by the table view to indicate row and column selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionHighlightStyle-swift.property
func (t_ TableView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionHighlightStyle:"), value)
}


// The table view’s sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/sortDescriptors
func (t_ TableView) SortDescriptors() []objc.IObject {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// The table view’s sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/sortDescriptors
func (t_ TableView) SetSortDescriptors(value []objc.IObject) {
	// Convert Go slice to NSArray
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
}


// The style that the table view uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/style-swift.property
func (t_ TableView) Style() TableViewStyle {
	rv := objc.Send[TableViewStyle](t_.ID, objc.Sel("style"))
	return rv
}


// The style that the table view uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/style-swift.property
func (t_ TableView) SetStyle(value TableViewStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStyle:"), value)
}


// An array containing the current table column objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/tableColumns
func (t_ TableView) TableColumns() []TableColumn {
	rv := objc.Send[[]TableColumn](t_.ID, objc.Sel("tableColumns"))
	return rv
}


// The layout direction of the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/userInterfaceLayoutDirection
func (t_ TableView) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](t_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// The layout direction of the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/userInterfaceLayoutDirection
func (t_ TableView) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}


// A Boolean value indicating whether the table view uses alternating row colors for its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesAlternatingRowBackgroundColors
func (t_ TableView) UsesAlternatingRowBackgroundColors() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesAlternatingRowBackgroundColors"))
	return rv
}


// A Boolean value indicating whether the table view uses alternating row colors for its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesAlternatingRowBackgroundColors
func (t_ TableView) SetUsesAlternatingRowBackgroundColors(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesAlternatingRowBackgroundColors:"), value)
}


// A Boolean value that indicates whether the table view uses autolayout to calculate the height of rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesAutomaticRowHeights
func (t_ TableView) UsesAutomaticRowHeights() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesAutomaticRowHeights"))
	return rv
}


// A Boolean value that indicates whether the table view uses autolayout to calculate the height of rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesAutomaticRowHeights
func (t_ TableView) SetUsesAutomaticRowHeights(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesAutomaticRowHeights:"), value)
}


// A Boolean value indicating whether the table uses static data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesStaticContents
func (t_ TableView) UsesStaticContents() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesStaticContents"))
	return rv
}


// A Boolean value indicating whether the table uses static data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/usesStaticContents
func (t_ TableView) SetUsesStaticContents(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesStaticContents:"), value)
}


// A Boolean value indicating whether vertical motion is treated as a drag or selection change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/verticalMotionCanBeginDrag
func (t_ TableView) VerticalMotionCanBeginDrag() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("verticalMotionCanBeginDrag"))
	return rv
}


// A Boolean value indicating whether vertical motion is treated as a drag or selection change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/verticalMotionCanBeginDrag
func (t_ TableView) SetVerticalMotionCanBeginDrag(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVerticalMotionCanBeginDrag:"), value)
}


// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/isenabled
func (t_ TableView) IsEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/isenabled
func (t_ TableView) SetIsEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEnabled:"), value)
}


