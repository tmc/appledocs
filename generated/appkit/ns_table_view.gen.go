// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	RectOfRow(row int) coregraphics.CGRect
	RemoveRowsAtIndexesWithAnimation(indexes unsafe.Pointer, animationOptions TableViewAnimationOptions)
	RowAtPoint(point coregraphics.CGPoint) int
	SetDraggingSourceOperationMaskForLocal(mask IDragOperation, isLocal bool)
}

// A set of related records, displayed in rows that represent individual records and columns that represent the attributes of those records.
//
// Table views are displayed in scroll views. Beginning with macOS v10.7, you can use objects (most commonly customized objects) instead of cells for specifying rows and columns. You can still use objects for each row and column item if you prefer. A table view does not store its own data; it retrieves data values as needed from a data source to which it has a weak reference. You should not, therefore, directly set data values programmatically in the table view; instead, modify the values in the data source and allow the changes to be reflected in the table view. To learn about the methods that an object uses to provide and access the contents of its data source object, see . To customize a table view’s behavior without subclassing , use the methods defined by the protocol. For example, the delegate supports table column management, type-to-select functionality, row selection and editing, custom tracking, and custom views for individual columns and rows. To learn more about the table view delegate, see .
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


// Returns the rectangle containing the row at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rect(ofRow:)
func (t_ TableView) RectOfRow(row int) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("rectOfRow:"), row)
	return rv
}

// Removes the rows using the specified animation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/removeRows(at:withAnimation:)
func (t_ TableView) RemoveRowsAtIndexesWithAnimation(indexes unsafe.Pointer, animationOptions TableViewAnimationOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeRowsAtIndexes:withAnimation:"), indexes, animationOptions)
}

// Returns the index of the row the specified point lies in.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/row(at:)
func (t_ TableView) RowAtPoint(point coregraphics.CGPoint) int {
	rv := objc.Send[int](t_.ID, objc.Sel("rowAtPoint:"), point)
	return rv
}

// Sets the default operation mask returned by to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/setDraggingSourceOperationMask(_:forLocal:)
func (t_ TableView) SetDraggingSourceOperationMaskForLocal(mask IDragOperation, isLocal bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}

// The color used to draw the background of the table.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/backgroundColor
func (t_ TableView) BackgroundColor() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The color used to draw the background of the table.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/backgroundColor
func (t_ TableView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}

// The table view’s column autoresizing style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/columnAutoresizingStyle-swift.property
func (t_ TableView) ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle {
	rv := objc.Send[TableViewColumnAutoresizingStyle](t_.ID, objc.Sel("columnAutoresizingStyle"))
	return rv
}


// SetColumnAutoresizingStyle sets the value of the columnAutoresizingStyle property.
// The table view’s column autoresizing style.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/columnAutoresizingStyle-swift.property
func (t_ TableView) SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColumnAutoresizingStyle:"), value)
}

// The feedback style displayed when the user drags over the table view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/draggingDestinationFeedbackStyle-swift.property
func (t_ TableView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := objc.Send[TableViewDraggingDestinationFeedbackStyle](t_.ID, objc.Sel("draggingDestinationFeedbackStyle"))
	return rv
}


// SetDraggingDestinationFeedbackStyle sets the value of the draggingDestinationFeedbackStyle property.
// The feedback style displayed when the user drags over the table view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/draggingDestinationFeedbackStyle-swift.property
func (t_ TableView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDraggingDestinationFeedbackStyle:"), value)
}

// The effective row size style for the table.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/effectiveRowSizeStyle
func (t_ TableView) EffectiveRowSizeStyle() TableViewRowSizeStyle {
	rv := objc.Send[TableViewRowSizeStyle](t_.ID, objc.Sel("effectiveRowSizeStyle"))
	return rv
}

// The grid lines drawn by the table view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridStyleMask
func (t_ TableView) GridStyleMask() TableViewGridLineStyle {
	rv := objc.Send[TableViewGridLineStyle](t_.ID, objc.Sel("gridStyleMask"))
	return rv
}


// SetGridStyleMask sets the value of the gridStyleMask property.
// The grid lines drawn by the table view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridStyleMask
func (t_ TableView) SetGridStyleMask(value TableViewGridLineStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGridStyleMask:"), value)
}

// The row size style (small, medium, large, or custom) used by the table view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowSizeStyle-swift.property
func (t_ TableView) RowSizeStyle() TableViewRowSizeStyle {
	rv := objc.Send[TableViewRowSizeStyle](t_.ID, objc.Sel("rowSizeStyle"))
	return rv
}


// SetRowSizeStyle sets the value of the rowSizeStyle property.
// The row size style (small, medium, large, or custom) used by the table view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowSizeStyle-swift.property
func (t_ TableView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowSizeStyle:"), value)
}

// The selection highlight style used by the table view to indicate row and column selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionHighlightStyle-swift.property
func (t_ TableView) SelectionHighlightStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("selectionHighlightStyle"))
	return rv
}


// SetSelectionHighlightStyle sets the value of the selectionHighlightStyle property.
// The selection highlight style used by the table view to indicate row and column selection.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionHighlightStyle-swift.property
func (t_ TableView) SetSelectionHighlightStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionHighlightStyle:"), value)
}

// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/isenabled
func (t_ TableView) IsEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// A Boolean value that indicates whether the receiver reacts to mouse events.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/isenabled
func (t_ TableView) SetIsEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEnabled:"), value)
}

// A Boolean value indicating whether the table view allows the user to rearrange columns by dragging their headers.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowscolumnreordering
func (t_ TableView) AllowsColumnReordering() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColumnReordering"))
	return rv
}


// SetAllowsColumnReordering sets the value of the allowsColumnReordering property.
// A Boolean value indicating whether the table view allows the user to rearrange columns by dragging their headers.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowscolumnreordering
func (t_ TableView) SetAllowsColumnReordering(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColumnReordering:"), value)
}

// A Boolean value indicating whether the table view allows the user to resize columns by dragging between their headers.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowscolumnresizing
func (t_ TableView) AllowsColumnResizing() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColumnResizing"))
	return rv
}


// SetAllowsColumnResizing sets the value of the allowsColumnResizing property.
// A Boolean value indicating whether the table view allows the user to resize columns by dragging between their headers.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowscolumnresizing
func (t_ TableView) SetAllowsColumnResizing(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColumnResizing:"), value)
}

// A Boolean value indicating whether the table view allows the user to select columns by clicking their headers.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowscolumnselection
func (t_ TableView) AllowsColumnSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColumnSelection"))
	return rv
}


// SetAllowsColumnSelection sets the value of the allowsColumnSelection property.
// A Boolean value indicating whether the table view allows the user to select columns by clicking their headers.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowscolumnselection
func (t_ TableView) SetAllowsColumnSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColumnSelection:"), value)
}

// A Boolean value indicating whether the table view allows the user to select zero columns or rows.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowsemptyselection
func (t_ TableView) AllowsEmptySelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}


// SetAllowsEmptySelection sets the value of the allowsEmptySelection property.
// A Boolean value indicating whether the table view allows the user to select zero columns or rows.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowsemptyselection
func (t_ TableView) SetAllowsEmptySelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsEmptySelection:"), value)
}

// A Boolean value indicating whether the table view allows the user to select more than one column or row at a time.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowsmultipleselection
func (t_ TableView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}


// SetAllowsMultipleSelection sets the value of the allowsMultipleSelection property.
// A Boolean value indicating whether the table view allows the user to select more than one column or row at a time.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowsmultipleselection
func (t_ TableView) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}

// A Boolean value indicating whether the table view allows the user to type characters to select rows.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowstypeselect
func (t_ TableView) AllowsTypeSelect() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsTypeSelect"))
	return rv
}


// SetAllowsTypeSelect sets the value of the allowsTypeSelect property.
// A Boolean value indicating whether the table view allows the user to type characters to select rows.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/allowstypeselect
func (t_ TableView) SetAllowsTypeSelect(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsTypeSelect:"), value)
}

// The name under which table information is automatically saved.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/autosavename-swift.property
func (t_ TableView) AutosaveName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("autosaveName"))
	return rv
}


// SetAutosaveName sets the value of the autosaveName property.
// The name under which table information is automatically saved.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/autosavename-swift.property
func (t_ TableView) SetAutosaveName(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutosaveName:"), value)
}

// A Boolean value indicating whether the order and width of the table view’s columns are automatically saved.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/autosavetablecolumns
func (t_ TableView) AutosaveTableColumns() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("autosaveTableColumns"))
	return rv
}


// SetAutosaveTableColumns sets the value of the autosaveTableColumns property.
// A Boolean value indicating whether the order and width of the table view’s columns are automatically saved.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/autosavetablecolumns
func (t_ TableView) SetAutosaveTableColumns(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutosaveTableColumns:"), value)
}

// The index of the column the user clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/clickedcolumn
func (t_ TableView) ClickedColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("clickedColumn"))
	return rv
}


// SetClickedColumn sets the value of the clickedColumn property.
// The index of the column the user clicked.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/clickedcolumn
func (t_ TableView) SetClickedColumn(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setClickedColumn:"), value)
}

// The index of the row the user clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/clickedrow
func (t_ TableView) ClickedRow() int {
	rv := objc.Send[int](t_.ID, objc.Sel("clickedRow"))
	return rv
}


// SetClickedRow sets the value of the clickedRow property.
// The index of the row the user clicked.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/clickedrow
func (t_ TableView) SetClickedRow(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setClickedRow:"), value)
}

// The view used to draw the area to the right of the column headers and above the vertical scroller of the enclosing scroll view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/cornerview
func (t_ TableView) CornerView() NSView {
	rv := objc.Send[NSView](t_.ID, objc.Sel("cornerView"))
	return rv
}


// SetCornerView sets the value of the cornerView property.
// The view used to draw the area to the right of the column headers and above the vertical scroller of the enclosing scroll view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/cornerview
func (t_ TableView) SetCornerView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCornerView:"), value)
}

// The object that provides the data displayed by the table view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/datasource
func (t_ TableView) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("dataSource"))
	return rv
}


// SetDataSource sets the value of the dataSource property.
// The object that provides the data displayed by the table view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/datasource
func (t_ TableView) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDataSource:"), value)
}

// The table view’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/delegate
func (t_ TableView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The table view’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/delegate
func (t_ TableView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}

// The message sent to the table view’s target when the user double-clicks a cell or column header.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/doubleaction
func (t_ TableView) DoubleAction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("doubleAction"))
	return rv
}


// SetDoubleAction sets the value of the doubleAction property.
// The message sent to the table view’s target when the user double-clicks a cell or column header.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/doubleaction
func (t_ TableView) SetDoubleAction(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDoubleAction:"), value)
}

// The index of the column being edited.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/editedcolumn
func (t_ TableView) EditedColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("editedColumn"))
	return rv
}


// SetEditedColumn sets the value of the editedColumn property.
// The index of the column being edited.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/editedcolumn
func (t_ TableView) SetEditedColumn(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditedColumn:"), value)
}

// The index of the row being edited.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/editedrow
func (t_ TableView) EditedRow() int {
	rv := objc.Send[int](t_.ID, objc.Sel("editedRow"))
	return rv
}


// SetEditedRow sets the value of the editedRow property.
// The index of the row being edited.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/editedrow
func (t_ TableView) SetEditedRow(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditedRow:"), value)
}

// The effective style that the table uses.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/effectivestyle
func (t_ TableView) EffectiveStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("effectiveStyle"))
	return rv
}


// SetEffectiveStyle sets the value of the effectiveStyle property.
// The effective style that the table uses.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/effectivestyle
func (t_ TableView) SetEffectiveStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEffectiveStyle:"), value)
}

// A Boolean value indicating whether the table view draws grouped rows as if they are floating.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/floatsgrouprows
func (t_ TableView) FloatsGroupRows() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("floatsGroupRows"))
	return rv
}


// SetFloatsGroupRows sets the value of the floatsGroupRows property.
// A Boolean value indicating whether the table view draws grouped rows as if they are floating.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/floatsgrouprows
func (t_ TableView) SetFloatsGroupRows(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFloatsGroupRows:"), value)
}

// The color used to draw grid lines.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/gridcolor
func (t_ TableView) GridColor() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("gridColor"))
	return rv
}


// SetGridColor sets the value of the gridColor property.
// The color used to draw grid lines.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/gridcolor
func (t_ TableView) SetGridColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGridColor:"), value)
}

// The view object used to draw headers over columns.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/headerview
func (t_ TableView) HeaderView() NSTableHeaderView {
	rv := objc.Send[NSTableHeaderView](t_.ID, objc.Sel("headerView"))
	return rv
}


// SetHeaderView sets the value of the headerView property.
// The view object used to draw headers over columns.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/headerview
func (t_ TableView) SetHeaderView(value ITableHeaderView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeaderView:"), value)
}

// The indexes of all hidden table rows.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/hiddenrowindexes
func (t_ TableView) HiddenRowIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("hiddenRowIndexes"))
	return rv
}


// SetHiddenRowIndexes sets the value of the hiddenRowIndexes property.
// The indexes of all hidden table rows.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/hiddenrowindexes
func (t_ TableView) SetHiddenRowIndexes(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHiddenRowIndexes:"), value)
}

// The column highlighted in the table.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/highlightedtablecolumn
func (t_ TableView) HighlightedTableColumn() NSTableColumn {
	rv := objc.Send[NSTableColumn](t_.ID, objc.Sel("highlightedTableColumn"))
	return rv
}


// SetHighlightedTableColumn sets the value of the highlightedTableColumn property.
// The column highlighted in the table.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/highlightedtablecolumn
func (t_ TableView) SetHighlightedTableColumn(value ITableColumn) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHighlightedTableColumn:"), value)
}

// The horizontal and vertical spacing between cells.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/intercellspacing
func (t_ TableView) IntercellSpacing() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("intercellSpacing"))
	return rv
}


// SetIntercellSpacing sets the value of the intercellSpacing property.
// The horizontal and vertical spacing between cells.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/intercellspacing
func (t_ TableView) SetIntercellSpacing(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIntercellSpacing:"), value)
}

// The number of columns in the table.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/numberofcolumns
func (t_ TableView) NumberOfColumns() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfColumns"))
	return rv
}


// SetNumberOfColumns sets the value of the numberOfColumns property.
// The number of columns in the table.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/numberofcolumns
func (t_ TableView) SetNumberOfColumns(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNumberOfColumns:"), value)
}

// The number of rows in the table.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/numberofrows
func (t_ TableView) NumberOfRows() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfRows"))
	return rv
}


// SetNumberOfRows sets the value of the numberOfRows property.
// The number of rows in the table.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/numberofrows
func (t_ TableView) SetNumberOfRows(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNumberOfRows:"), value)
}

// The number of selected columns.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/numberofselectedcolumns
func (t_ TableView) NumberOfSelectedColumns() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfSelectedColumns"))
	return rv
}


// SetNumberOfSelectedColumns sets the value of the numberOfSelectedColumns property.
// The number of selected columns.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/numberofselectedcolumns
func (t_ TableView) SetNumberOfSelectedColumns(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNumberOfSelectedColumns:"), value)
}

// The number of selected rows.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/numberofselectedrows
func (t_ TableView) NumberOfSelectedRows() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfSelectedRows"))
	return rv
}


// SetNumberOfSelectedRows sets the value of the numberOfSelectedRows property.
// The number of selected rows.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/numberofselectedrows
func (t_ TableView) SetNumberOfSelectedRows(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNumberOfSelectedRows:"), value)
}

// The dictionary of all registered nib files for view-based table view identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/registerednibsbyidentifier
func (t_ TableView) RegisteredNibsByIdentifier() NSNib {
	rv := objc.Send[NSNib](t_.ID, objc.Sel("registeredNibsByIdentifier"))
	return rv
}


// SetRegisteredNibsByIdentifier sets the value of the registeredNibsByIdentifier property.
// The dictionary of all registered nib files for view-based table view identifiers.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/registerednibsbyidentifier
func (t_ TableView) SetRegisteredNibsByIdentifier(value INib) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRegisteredNibsByIdentifier:"), value)
}

// A Boolean value indicating whether a table row’s actions are visible.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/rowactionsvisible
func (t_ TableView) RowActionsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rowActionsVisible"))
	return rv
}


// SetRowActionsVisible sets the value of the rowActionsVisible property.
// A Boolean value indicating whether a table row’s actions are visible.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/rowactionsvisible
func (t_ TableView) SetRowActionsVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowActionsVisible:"), value)
}

// The height of each row in the table.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/rowheight
func (t_ TableView) RowHeight() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("rowHeight"))
	return rv
}


// SetRowHeight sets the value of the rowHeight property.
// The height of each row in the table.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/rowheight
func (t_ TableView) SetRowHeight(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowHeight:"), value)
}

// The index of the last selected column (or the last column added to the selection).
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/selectedcolumn
func (t_ TableView) SelectedColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("selectedColumn"))
	return rv
}


// SetSelectedColumn sets the value of the selectedColumn property.
// The index of the last selected column (or the last column added to the selection).

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/selectedcolumn
func (t_ TableView) SetSelectedColumn(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedColumn:"), value)
}

// An index set containing the indexes of the selected columns.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/selectedcolumnindexes
func (t_ TableView) SelectedColumnIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("selectedColumnIndexes"))
	return rv
}


// SetSelectedColumnIndexes sets the value of the selectedColumnIndexes property.
// An index set containing the indexes of the selected columns.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/selectedcolumnindexes
func (t_ TableView) SetSelectedColumnIndexes(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedColumnIndexes:"), value)
}

// The index of the last selected row (or the last row added to the selection).
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/selectedrow
func (t_ TableView) SelectedRow() int {
	rv := objc.Send[int](t_.ID, objc.Sel("selectedRow"))
	return rv
}


// SetSelectedRow sets the value of the selectedRow property.
// The index of the last selected row (or the last row added to the selection).

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/selectedrow
func (t_ TableView) SetSelectedRow(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRow:"), value)
}

// An index set containing the indexes of the selected rows.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/selectedrowindexes
func (t_ TableView) SelectedRowIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("selectedRowIndexes"))
	return rv
}


// SetSelectedRowIndexes sets the value of the selectedRowIndexes property.
// An index set containing the indexes of the selected rows.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/selectedrowindexes
func (t_ TableView) SetSelectedRowIndexes(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRowIndexes:"), value)
}

// The table view’s sort descriptors.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/sortdescriptors
func (t_ TableView) SortDescriptors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// SetSortDescriptors sets the value of the sortDescriptors property.
// The table view’s sort descriptors.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/sortdescriptors
func (t_ TableView) SetSortDescriptors(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSortDescriptors:"), value)
}

// The style that the table view uses.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/style-swift.property
func (t_ TableView) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// The style that the table view uses.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/style-swift.property
func (t_ TableView) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStyle:"), value)
}

// An array containing the current table column objects.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/tablecolumns
func (t_ TableView) TableColumns() NSTableColumn {
	rv := objc.Send[NSTableColumn](t_.ID, objc.Sel("tableColumns"))
	return rv
}


// SetTableColumns sets the value of the tableColumns property.
// An array containing the current table column objects.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/tablecolumns
func (t_ TableView) SetTableColumns(value ITableColumn) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTableColumns:"), value)
}

// The layout direction of the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/userinterfacelayoutdirection
func (t_ TableView) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](t_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// SetUserInterfaceLayoutDirection sets the value of the userInterfaceLayoutDirection property.
// The layout direction of the user interface.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/userinterfacelayoutdirection
func (t_ TableView) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}

// A Boolean value indicating whether the table view uses alternating row colors for its background.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/usesalternatingrowbackgroundcolors
func (t_ TableView) UsesAlternatingRowBackgroundColors() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesAlternatingRowBackgroundColors"))
	return rv
}


// SetUsesAlternatingRowBackgroundColors sets the value of the usesAlternatingRowBackgroundColors property.
// A Boolean value indicating whether the table view uses alternating row colors for its background.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/usesalternatingrowbackgroundcolors
func (t_ TableView) SetUsesAlternatingRowBackgroundColors(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesAlternatingRowBackgroundColors:"), value)
}

// A Boolean value that indicates whether the table view uses autolayout to calculate the height of rows.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/usesautomaticrowheights
func (t_ TableView) UsesAutomaticRowHeights() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesAutomaticRowHeights"))
	return rv
}


// SetUsesAutomaticRowHeights sets the value of the usesAutomaticRowHeights property.
// A Boolean value that indicates whether the table view uses autolayout to calculate the height of rows.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/usesautomaticrowheights
func (t_ TableView) SetUsesAutomaticRowHeights(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesAutomaticRowHeights:"), value)
}

// A Boolean value indicating whether the table uses static data.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/usesstaticcontents
func (t_ TableView) UsesStaticContents() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesStaticContents"))
	return rv
}


// SetUsesStaticContents sets the value of the usesStaticContents property.
// A Boolean value indicating whether the table uses static data.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/usesstaticcontents
func (t_ TableView) SetUsesStaticContents(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesStaticContents:"), value)
}

// A Boolean value indicating whether vertical motion is treated as a drag or selection change.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/verticalmotioncanbegindrag
func (t_ TableView) VerticalMotionCanBeginDrag() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("verticalMotionCanBeginDrag"))
	return rv
}


// SetVerticalMotionCanBeginDrag sets the value of the verticalMotionCanBeginDrag property.
// A Boolean value indicating whether vertical motion is treated as a drag or selection change.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/verticalmotioncanbegindrag
func (t_ TableView) SetVerticalMotionCanBeginDrag(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVerticalMotionCanBeginDrag:"), value)
}



