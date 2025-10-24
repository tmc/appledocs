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

// The class instance for the [Browser] class.
var (
	BrowserClass     _BrowserClass
	BrowserClassOnce sync.Once
)

func getBrowserClass() _BrowserClass {
	BrowserClassOnce.Do(func() {
		BrowserClass = _BrowserClass{objc.GetClass("NSBrowser")}
	})
	return BrowserClass
}

type _BrowserClass struct {
	class objc.Class
}

// An interface definition for the [Browser] class.
type IBrowser interface {
	IControl
	// properties:
	AllowsBranchSelection() bool
	SetAllowsBranchSelection(value bool)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	AllowsTypeSelect() bool
	SetAllowsTypeSelect(value bool)
	AutohidesScroller() bool
	SetAutohidesScroller(value bool)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	CellPrototype() objc.ID
	SetCellPrototype(value objc.ID)
	ClickedColumn() int
	ClickedRow() int
	ColumnResizingType() BrowserColumnResizingType
	SetColumnResizingType(value BrowserColumnResizingType)
	ColumnsAutosaveName() objc.IObject /* cross-framework: BrowserColumnsAutosaveName */
	SetColumnsAutosaveName(value objc.IObject /* cross-framework: BrowserColumnsAutosaveName */)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	DoubleAction() objc.SEL
	SetDoubleAction(value objc.SEL)
	FirstVisibleColumn() int
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	Loaded() bool
	Titled() bool
	SetTitled(value bool)
	LastColumn() int
	SetLastColumn(value int)
	LastVisibleColumn() int
	MaxVisibleColumns() int
	SetMaxVisibleColumns(value int)
	MinColumnWidth() float64
	SetMinColumnWidth(value float64)
	NumberOfVisibleColumns() int
	PathSeparator() objc.IObject /* cross-framework: NSString */
	SetPathSeparator(value objc.IObject /* cross-framework: NSString */)
	PrefersAllColumnUserResizing() bool
	SetPrefersAllColumnUserResizing(value bool)
	ReusesColumns() bool
	SetReusesColumns(value bool)
	RowHeight() float64
	SetRowHeight(value float64)
	SelectedCell() objc.ID
	SelectedCells() []Cell
	SelectedColumn() int
	SelectionIndexPath() foundation.IndexPath
	SetSelectionIndexPath(value foundation.IndexPath)
	SelectionIndexPaths() []foundation.IndexPath
	SetSelectionIndexPaths(value []foundation.IndexPath)
	SendsActionOnArrowKeys() bool
	SetSendsActionOnArrowKeys(value bool)
	SeparatesColumns() bool
	SetSeparatesColumns(value bool)
	TakesTitleFromPreviousColumn() bool
	SetTakesTitleFromPreviousColumn(value bool)
	TitleHeight() float64
	IsLoaded() bool
	SetIsLoaded(value bool)
	IsTitled() bool
	SetIsTitled(value bool)
	IsOpaque() bool
	SetIsOpaque(value bool)
	// methods:
	AddColumn()
	CanDragRowsWithIndexesInColumnWithEvent(rowIndexes foundation.IndexSet, column int, event IEvent) bool
	ColumnContentWidthForColumnWidth(columnWidth float64) float64
	ColumnWidthForColumnContentWidth(columnContentWidth float64) float64
	DefaultColumnWidth() float64
	DoClick(sender objc.IObject)
	DoDoubleClick(sender objc.IObject)
	DraggingImageForRowsWithIndexesInColumnWithEventOffset(rowIndexes foundation.IndexSet, column int, event IEvent, dragImageOffset PointPointer /* not a class type */) IImage
	DrawTitleOfColumnInRect(column int, rect objc.IObject /* cross-framework: Rect */)
	EditItemAtIndexPathWithEventSelect(indexPath foundation.IndexPath, event IEvent, select_ bool)
	FrameOfColumn(column int) objc.IObject /* cross-framework: Rect */
	FrameOfInsideOfColumn(column int) objc.IObject /* cross-framework: Rect */
	FrameOfRowInColumn(row int, column int) objc.IObject /* cross-framework: Rect */
	GetRowColumnForPoint(row int, column int, point objc.IObject /* cross-framework: Point */) bool
	IndexPathForColumn(column int) foundation.IndexPath
	IsLeafItem(item objc.IObject) bool
	ItemAtIndexPath(indexPath foundation.IndexPath) objc.ID
	ItemAtRowInColumn(row int, column int) objc.ID
	LoadColumnZero()
	LoadedCellAtRowColumn(row int, col int) objc.ID
	NoteHeightOfRowsWithIndexesChangedInColumn(indexSet foundation.IndexSet, columnIndex int)
	ParentForItemsInColumn(column int) objc.ID
	Path() foundation.String
	PathToColumn(column int) foundation.String
	ReloadColumn(column int)
	ReloadDataForRowIndexesInColumn(rowIndexes foundation.IndexSet, column int)
	ScrollColumnToVisible(column int)
	ScrollColumnsLeftBy(shiftAmount int)
	ScrollColumnsRightBy(shiftAmount int)
	ScrollRowToVisibleInColumn(row int, column int)
	SelectAll(sender objc.IObject)
	SelectRowInColumn(row int, column int)
	SelectRowIndexesInColumn(indexes foundation.IndexSet, column int)
	SelectedCellInColumn(column int) objc.ID
	SelectedRowInColumn(column int) int
	SelectedRowIndexesInColumn(column int) foundation.IndexSet
	SendAction() bool
	SetDefaultColumnWidth(columnWidth float64)
	SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool)
	SetPath(path objc.IObject /* cross-framework: NSString */) bool
	SetTitleOfColumn(string_ objc.IObject /* cross-framework: NSString */, column int)
	SetWidthOfColumn(columnWidth float64, columnIndex int)
	Tile()
	TitleOfColumn(column int) foundation.String
	TitleFrameOfColumn(column int) objc.IObject /* cross-framework: Rect */
	ValidateVisibleColumns()
	WidthOfColumn(column int) float64
}

// An interface that displays a hierarchically organized list of data items that can be navigated and selected.
//
// A browser displays information using a set of columns, which are indexed from left to right. Each successive column displays the next level down in the data hierarchy. This class uses the class to implement its user interface. Browsers have the following components: Columns Scroll views Matrices Browser cells To the user, browsers display data in columns and rows within each column. These components are arranged in the following component hierarchy:


// An interface that displays a hierarchically organized list of data items that can be navigated and selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser
type Browser struct {
	Control
}

// BrowserFrom constructs a [Browser] from an unsafe.Pointer.
//
// An interface that displays a hierarchically organized list of data items that can be navigated and selected.
func BrowserFrom(ptr unsafe.Pointer) Browser {
	return Browser{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BrowserClass) Alloc() Browser {
	rv := objc.Send[Browser](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BrowserClass) New() Browser {
	rv := objc.Send[Browser](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ Browser) Init() Browser {
	rv := objc.Send[Browser](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ Browser) Autorelease() Browser {
	rv := objc.Send[Browser](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBrowser creates a new Browser instance.
func NewBrowser() Browser {
	return getBrowserClass().New()
}



// Removes the column configuration data stored under the given name from the application’s user defaults.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/removeSavedColumns(withAutosaveName:)
func (bc _BrowserClass) RemoveSavedColumnsWithAutosaveName(name objc.IObject /* cross-framework: BrowserColumnsAutosaveName */) {
	objc.Send[objc.ID](objc.ID(bc.class), objc.Sel("removeSavedColumnsWithAutosaveName:"), name)
}


// Returns the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/cellClass
func (bc _BrowserClass) CellClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(bc.class), objc.Sel("cellClass"))
	return rv
}

// Adds a column to the right of the last column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/addColumn()
func (b_ Browser) AddColumn() {
	objc.Send[objc.ID](b_.ID, objc.Sel("addColumn"))
}


// Indicates whether the browser can attempt to initiate a drag of the given rows for the given event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/canDragRows(with:inColumn:with:)
func (b_ Browser) CanDragRowsWithIndexesInColumnWithEvent(rowIndexes foundation.IndexSet, column int, event IEvent) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("canDragRowsWithIndexes:inColumn:withEvent:"), rowIndexes, column, event)
	return rv
}


// Returns the content width for a given column width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/columnContentWidth(forColumnWidth:)
func (b_ Browser) ColumnContentWidthForColumnWidth(columnWidth float64) float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("columnContentWidthForColumnWidth:"), columnWidth)
	return rv
}


// Returns the column width for the width of the given column’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/columnWidth(forColumnContentWidth:)
func (b_ Browser) ColumnWidthForColumnContentWidth(columnContentWidth float64) float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("columnWidthForColumnContentWidth:"), columnContentWidth)
	return rv
}


// Returns the default column width of the browser’s columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/defaultColumnWidth()
func (b_ Browser) DefaultColumnWidth() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("defaultColumnWidth"))
	return rv
}


// Responds to (single) mouse clicks in a column of the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/doClick(_:)
func (b_ Browser) DoClick(sender objc.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("doClick:"), sender)
}


// Responds to double clicks in a column of the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/doDoubleClick(_:)
func (b_ Browser) DoDoubleClick(sender objc.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("doDoubleClick:"), sender)
}


// Provides an image to represent dragged rows during a drag operation on the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/draggingImageForRows(with:inColumn:with:offset:)
func (b_ Browser) DraggingImageForRowsWithIndexesInColumnWithEventOffset(rowIndexes foundation.IndexSet, column int, event IEvent, dragImageOffset PointPointer /* not a class type */) IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("draggingImageForRowsWithIndexes:inColumn:withEvent:offset:"), rowIndexes, column, event, dragImageOffset)
	return rv
}


// Draws the title for the specified column within the given rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/drawTitle(ofColumn:in:)
func (b_ Browser) DrawTitleOfColumnInRect(column int, rect objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("drawTitleOfColumn:inRect:"), column, rect)
}


// Begins editing the item at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/editItem(at:with:select:)
func (b_ Browser) EditItemAtIndexPathWithEventSelect(indexPath foundation.IndexPath, event IEvent, select_ bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("editItemAtIndexPath:withEvent:select:"), indexPath, event, select_)
}


// Returns the rectangle containing the given column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/frame(ofColumn:)
func (b_ Browser) FrameOfColumn(column int) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](b_.ID, objc.Sel("frameOfColumn:"), column)
	return rv
}


// Returns the rectangle containing the specified column, not including borders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/frame(ofInsideOfColumn:)
func (b_ Browser) FrameOfInsideOfColumn(column int) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](b_.ID, objc.Sel("frameOfInsideOfColumn:"), column)
	return rv
}


// Returns the frame of the cell at the specified location, including the expandable arrow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/frame(ofRow:inColumn:)
func (b_ Browser) FrameOfRowInColumn(row int, column int) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](b_.ID, objc.Sel("frameOfRow:inColumn:"), row, column)
	return rv
}


// Gets the row and column coordinates for the specified point, if a cell exists at that point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/getRow(_:column:for:)
func (b_ Browser) GetRowColumnForPoint(row int, column int, point objc.IObject /* cross-framework: Point */) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("getRow:column:forPoint:"), row, column, point)
	return rv
}


// Returns the index path of the item whose children are displayed in the given column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/indexPath(forColumn:)
func (b_ Browser) IndexPathForColumn(column int) foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](b_.ID, objc.Sel("indexPathForColumn:"), column)
	return rv
}


// Returns whether the specified item is a leaf item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/isLeafItem(_:)
func (b_ Browser) IsLeafItem(item objc.IObject) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isLeafItem:"), item)
	return rv
}


// Returns the item at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/item(at:)
func (b_ Browser) ItemAtIndexPath(indexPath foundation.IndexPath) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("itemAtIndexPath:"), indexPath)
	return rv
}


// Returns the item located at the specified row and column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/item(atRow:inColumn:)
func (b_ Browser) ItemAtRowInColumn(row int, column int) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("itemAtRow:inColumn:"), row, column)
	return rv
}


// Loads column 0; unloads previously loaded columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/loadColumnZero()
func (b_ Browser) LoadColumnZero() {
	objc.Send[objc.ID](b_.ID, objc.Sel("loadColumnZero"))
}


// Loads, if necessary, and returns the cell at the specified row and column location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/loadedCell(atRow:column:)
func (b_ Browser) LoadedCellAtRowColumn(row int, col int) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("loadedCellAtRow:column:"), row, col)
	return rv
}


// Immediately retiles the browser’s columns using row heights specified by the browser’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/noteHeightOfRowsWithIndexesChanged(_:inColumn:)
func (b_ Browser) NoteHeightOfRowsWithIndexesChangedInColumn(indexSet foundation.IndexSet, columnIndex int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("noteHeightOfRowsWithIndexesChanged:inColumn:"), indexSet, columnIndex)
}


// Returns the item that contains the children located in the specified column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/parentForItems(inColumn:)
func (b_ Browser) ParentForItemsInColumn(column int) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("parentForItemsInColumn:"), column)
	return rv
}


// Returns a string representing the browser’s current path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/path()
func (b_ Browser) Path() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("path"))
	return rv
}


// Returns a string representing the path from the first column up to, but not including, the column at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/path(toColumn:)
func (b_ Browser) PathToColumn(column int) foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("pathToColumn:"), column)
	return rv
}


// Reloads the given column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/reloadColumn(_:)
func (b_ Browser) ReloadColumn(column int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("reloadColumn:"), column)
}


// Updates the rows in the column with the specified column index with indexes in the specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/reloadData(forRowIndexes:inColumn:)
func (b_ Browser) ReloadDataForRowIndexesInColumn(rowIndexes foundation.IndexSet, column int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("reloadDataForRowIndexes:inColumn:"), rowIndexes, column)
}


// Scrolls to make the specified column visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/scrollColumnToVisible(_:)
func (b_ Browser) ScrollColumnToVisible(column int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("scrollColumnToVisible:"), column)
}


// Scrolls columns left by the specified number of columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/scrollColumnsLeft(by:)
func (b_ Browser) ScrollColumnsLeftBy(shiftAmount int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("scrollColumnsLeftBy:"), shiftAmount)
}


// Scrolls columns right by the specified number of columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/scrollColumnsRight(by:)
func (b_ Browser) ScrollColumnsRightBy(shiftAmount int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("scrollColumnsRightBy:"), shiftAmount)
}


// Scrolls the specified row to be visible within the specified column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/scrollRowToVisible(_:inColumn:)
func (b_ Browser) ScrollRowToVisibleInColumn(row int, column int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("scrollRowToVisible:inColumn:"), row, column)
}


// Selects all cells in the last column of the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectAll(_:)
func (b_ Browser) SelectAll(sender objc.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("selectAll:"), sender)
}


// Selects the cell at the specified row and column index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectRow(_:inColumn:)
func (b_ Browser) SelectRowInColumn(row int, column int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("selectRow:inColumn:"), row, column)
}


// Specifies the selected rows in a given column of the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectRowIndexes(_:inColumn:)
func (b_ Browser) SelectRowIndexesInColumn(indexes foundation.IndexSet, column int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("selectRowIndexes:inColumn:"), indexes, column)
}


// Returns the last (lowest) cell selected in the given column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectedCell(inColumn:)
func (b_ Browser) SelectedCellInColumn(column int) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("selectedCellInColumn:"), column)
	return rv
}


// Returns the row index of the selected cell in the specified column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectedRow(inColumn:)
func (b_ Browser) SelectedRowInColumn(column int) int {
	rv := objc.Send[int](b_.ID, objc.Sel("selectedRowInColumn:"), column)
	return rv
}


// Provides the indexes of the selected rows in a given column of the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectedRowIndexes(inColumn:)
func (b_ Browser) SelectedRowIndexesInColumn(column int) foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](b_.ID, objc.Sel("selectedRowIndexesInColumn:"), column)
	return rv
}


// Sends the action message to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/sendAction()
func (b_ Browser) SendAction() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("sendAction"))
	return rv
}


// Sets the default column width for new browser columns that do not otherwise have an initial width from defaults or the browser’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/setDefaultColumnWidth(_:)
func (b_ Browser) SetDefaultColumnWidth(columnWidth float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDefaultColumnWidth:"), columnWidth)
}


// Specifies the drag-operation mask for dragging operations with local or external destinations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/setDraggingSourceOperationMask(_:forLocal:)
func (b_ Browser) SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}


// Sets the path to be displayed by the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/setPath(_:)
func (b_ Browser) SetPath(path objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("setPath:"), path)
	return rv
}


// Sets the title of the given column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/setTitle(_:ofColumn:)
func (b_ Browser) SetTitleOfColumn(string_ objc.IObject /* cross-framework: NSString */, column int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:ofColumn:"), string_, column)
}


// Sets the width of the specified column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/setWidth(_:ofColumn:)
func (b_ Browser) SetWidthOfColumn(columnWidth float64, columnIndex int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setWidth:ofColumn:"), columnWidth, columnIndex)
}


// Adjusts the various subviews of the browser—scrollers, columns, titles, and so on—without redrawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/tile()
func (b_ Browser) Tile() {
	objc.Send[objc.ID](b_.ID, objc.Sel("tile"))
}


// Returns the title displayed for the given column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/title(ofColumn:)
func (b_ Browser) TitleOfColumn(column int) foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("titleOfColumn:"), column)
	return rv
}


// Returns the bounds of the title frame for the specified column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/titleFrame(ofColumn:)
func (b_ Browser) TitleFrameOfColumn(column int) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](b_.ID, objc.Sel("titleFrameOfColumn:"), column)
	return rv
}


// Validates the browser’s visible columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/validateVisibleColumns()
func (b_ Browser) ValidateVisibleColumns() {
	objc.Send[objc.ID](b_.ID, objc.Sel("validateVisibleColumns"))
}


// Returns the width of the specified column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/width(ofColumn:)
func (b_ Browser) WidthOfColumn(column int) float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("widthOfColumn:"), column)
	return rv
}


// A Boolean that indicates whether the user can select branch items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsBranchSelection
func (b_ Browser) AllowsBranchSelection() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsBranchSelection"))
	return rv
}


// A Boolean that indicates whether the user can select branch items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsBranchSelection
func (b_ Browser) SetAllowsBranchSelection(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsBranchSelection:"), value)
}


// A Boolean that indicates whether there can be nothing selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsEmptySelection
func (b_ Browser) AllowsEmptySelection() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}


// A Boolean that indicates whether there can be nothing selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsEmptySelection
func (b_ Browser) SetAllowsEmptySelection(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsEmptySelection:"), value)
}


// A Boolean that indicates whether the user can select multiple items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsMultipleSelection
func (b_ Browser) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}


// A Boolean that indicates whether the user can select multiple items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsMultipleSelection
func (b_ Browser) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}


// A Boolean that indicates whether the browser allows keystroke-based selection (type select).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsTypeSelect
func (b_ Browser) AllowsTypeSelect() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsTypeSelect"))
	return rv
}


// A Boolean that indicates whether the browser allows keystroke-based selection (type select).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsTypeSelect
func (b_ Browser) SetAllowsTypeSelect(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsTypeSelect:"), value)
}


// A Boolean that indicates whether the browser automatically hides its scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/autohidesScroller
func (b_ Browser) AutohidesScroller() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("autohidesScroller"))
	return rv
}


// A Boolean that indicates whether the browser automatically hides its scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/autohidesScroller
func (b_ Browser) SetAutohidesScroller(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAutohidesScroller:"), value)
}


// The browser’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/backgroundColor
func (b_ Browser) BackgroundColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The browser’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/backgroundColor
func (b_ Browser) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackgroundColor:"), value)
}


// Returns the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/cellClass
func (b_ Browser) CellClass() objc.Class {
	rv := objc.Send[objc.Class](b_.ID, objc.Sel("cellClass"))
	return rv
}


// The prototype for displaying items in the matrices in the columns of the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/cellPrototype
func (b_ Browser) CellPrototype() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("cellPrototype"))
	return rv
}


// The prototype for displaying items in the matrices in the columns of the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/cellPrototype
func (b_ Browser) SetCellPrototype(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCellPrototype:"), value)
}


// The column number of the cell that the user clicked to display a context menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/clickedColumn
func (b_ Browser) ClickedColumn() int {
	rv := objc.Send[int](b_.ID, objc.Sel("clickedColumn"))
	return rv
}


// The row number of the cell that the user clicked to display a context menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/clickedRow
func (b_ Browser) ClickedRow() int {
	rv := objc.Send[int](b_.ID, objc.Sel("clickedRow"))
	return rv
}


// A constant indicating the browser’s column resizing type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/columnResizingType-swift.property
func (b_ Browser) ColumnResizingType() BrowserColumnResizingType {
	rv := objc.Send[BrowserColumnResizingType](b_.ID, objc.Sel("columnResizingType"))
	return rv
}


// A constant indicating the browser’s column resizing type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/columnResizingType-swift.property
func (b_ Browser) SetColumnResizingType(value BrowserColumnResizingType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setColumnResizingType:"), value)
}


// The name used to automatically save the browser’s column configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/columnsAutosaveName-swift.property
func (b_ Browser) ColumnsAutosaveName() objc.IObject /* cross-framework: BrowserColumnsAutosaveName */ {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("columnsAutosaveName"))
	return rv
}


// The name used to automatically save the browser’s column configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/columnsAutosaveName-swift.property
func (b_ Browser) SetColumnsAutosaveName(value objc.IObject /* cross-framework: BrowserColumnsAutosaveName */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setColumnsAutosaveName:"), value)
}


// The browser’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/delegate
func (b_ Browser) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}


// The browser’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/delegate
func (b_ Browser) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}


// The browser’s double-click action method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/doubleAction
func (b_ Browser) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](b_.ID, objc.Sel("doubleAction"))
	return rv
}


// The browser’s double-click action method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/doubleAction
func (b_ Browser) SetDoubleAction(value objc.SEL) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDoubleAction:"), value)
}


// The index of the first visible column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/firstVisibleColumn
func (b_ Browser) FirstVisibleColumn() int {
	rv := objc.Send[int](b_.ID, objc.Sel("firstVisibleColumn"))
	return rv
}


// A Boolean that indicates whether the browser has a horizontal scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/hasHorizontalScroller
func (b_ Browser) HasHorizontalScroller() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("hasHorizontalScroller"))
	return rv
}


// A Boolean that indicates whether the browser has a horizontal scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/hasHorizontalScroller
func (b_ Browser) SetHasHorizontalScroller(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setHasHorizontalScroller:"), value)
}


// A Boolean that indicates whether column 0 is loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/isLoaded
func (b_ Browser) Loaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("loaded"))
	return rv
}


// A Boolean that indicates whether columns display titles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/isTitled
func (b_ Browser) Titled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("titled"))
	return rv
}


// A Boolean that indicates whether columns display titles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/isTitled
func (b_ Browser) SetTitled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitled:"), value)
}


// The index of the last column loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/lastColumn
func (b_ Browser) LastColumn() int {
	rv := objc.Send[int](b_.ID, objc.Sel("lastColumn"))
	return rv
}


// The index of the last column loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/lastColumn
func (b_ Browser) SetLastColumn(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLastColumn:"), value)
}


// The index of the last visible column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/lastVisibleColumn
func (b_ Browser) LastVisibleColumn() int {
	rv := objc.Send[int](b_.ID, objc.Sel("lastVisibleColumn"))
	return rv
}


// The maximum number of visible columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/maxVisibleColumns
func (b_ Browser) MaxVisibleColumns() int {
	rv := objc.Send[int](b_.ID, objc.Sel("maxVisibleColumns"))
	return rv
}


// The maximum number of visible columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/maxVisibleColumns
func (b_ Browser) SetMaxVisibleColumns(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMaxVisibleColumns:"), value)
}


// The minimum column width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/minColumnWidth
func (b_ Browser) MinColumnWidth() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("minColumnWidth"))
	return rv
}


// The minimum column width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/minColumnWidth
func (b_ Browser) SetMinColumnWidth(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMinColumnWidth:"), value)
}


// The number of visible columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/numberOfVisibleColumns
func (b_ Browser) NumberOfVisibleColumns() int {
	rv := objc.Send[int](b_.ID, objc.Sel("numberOfVisibleColumns"))
	return rv
}


// The path separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/pathSeparator
func (b_ Browser) PathSeparator() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("pathSeparator"))
	return rv
}


// The path separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/pathSeparator
func (b_ Browser) SetPathSeparator(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPathSeparator:"), value)
}


// A Boolean that indicates whether the browser is set to resize all columns simultaneously rather than resizing a single column at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/prefersAllColumnUserResizing
func (b_ Browser) PrefersAllColumnUserResizing() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("prefersAllColumnUserResizing"))
	return rv
}


// A Boolean that indicates whether the browser is set to resize all columns simultaneously rather than resizing a single column at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/prefersAllColumnUserResizing
func (b_ Browser) SetPrefersAllColumnUserResizing(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrefersAllColumnUserResizing:"), value)
}


// A Boolean that indicates whether the browser reuses matrix objects after their columns are unloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/reusesColumns
func (b_ Browser) ReusesColumns() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("reusesColumns"))
	return rv
}


// A Boolean that indicates whether the browser reuses matrix objects after their columns are unloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/reusesColumns
func (b_ Browser) SetReusesColumns(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setReusesColumns:"), value)
}


// The height of the browser’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/rowHeight
func (b_ Browser) RowHeight() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("rowHeight"))
	return rv
}


// The height of the browser’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/rowHeight
func (b_ Browser) SetRowHeight(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setRowHeight:"), value)
}


// The last (rightmost and lowest) selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectedCell
func (b_ Browser) SelectedCell() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("selectedCell"))
	return rv
}


// All cells selected in the rightmost column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectedCells
func (b_ Browser) SelectedCells() []Cell {
	rv := objc.Send[[]Cell](b_.ID, objc.Sel("selectedCells"))
	return rv
}


// The index of the last column with a selected item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectedColumn
func (b_ Browser) SelectedColumn() int {
	rv := objc.Send[int](b_.ID, objc.Sel("selectedColumn"))
	return rv
}


// The index path of the item selected in the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectionIndexPath
func (b_ Browser) SelectionIndexPath() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](b_.ID, objc.Sel("selectionIndexPath"))
	return rv
}


// The index path of the item selected in the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectionIndexPath
func (b_ Browser) SetSelectionIndexPath(value foundation.IndexPath) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSelectionIndexPath:"), value)
}


// An array containing the index paths of all items selected in the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectionIndexPaths
func (b_ Browser) SelectionIndexPaths() []foundation.IndexPath {
	rv := objc.Send[[]foundation.IndexPath](b_.ID, objc.Sel("selectionIndexPaths"))
	return rv
}


// An array containing the index paths of all items selected in the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/selectionIndexPaths
func (b_ Browser) SetSelectionIndexPaths(value []foundation.IndexPath) {
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
	objc.Send[objc.ID](b_.ID, objc.Sel("setSelectionIndexPaths:"), nsArray)
}


// A Boolean that indicates whether pressing an arrow key causes an action message to be sent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/sendsActionOnArrowKeys
func (b_ Browser) SendsActionOnArrowKeys() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("sendsActionOnArrowKeys"))
	return rv
}


// A Boolean that indicates whether pressing an arrow key causes an action message to be sent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/sendsActionOnArrowKeys
func (b_ Browser) SetSendsActionOnArrowKeys(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSendsActionOnArrowKeys:"), value)
}


// A Boolean that indicates whether columns are separated by bezeled borders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/separatesColumns
func (b_ Browser) SeparatesColumns() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("separatesColumns"))
	return rv
}


// A Boolean that indicates whether columns are separated by bezeled borders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/separatesColumns
func (b_ Browser) SetSeparatesColumns(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSeparatesColumns:"), value)
}


// A Boolean that indicates whether a column takes its title from the selected cell in the previous column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/takesTitleFromPreviousColumn
func (b_ Browser) TakesTitleFromPreviousColumn() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("takesTitleFromPreviousColumn"))
	return rv
}


// A Boolean that indicates whether a column takes its title from the selected cell in the previous column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/takesTitleFromPreviousColumn
func (b_ Browser) SetTakesTitleFromPreviousColumn(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTakesTitleFromPreviousColumn:"), value)
}


// The height of the column titles for the browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/titleHeight
func (b_ Browser) TitleHeight() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("titleHeight"))
	return rv
}


// A Boolean that indicates whether column 0 is loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/isloaded
func (b_ Browser) IsLoaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isLoaded"))
	return rv
}


// A Boolean that indicates whether column 0 is loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/isloaded
func (b_ Browser) SetIsLoaded(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsLoaded:"), value)
}


// A Boolean that indicates whether columns display titles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/istitled
func (b_ Browser) IsTitled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTitled"))
	return rv
}


// A Boolean that indicates whether columns display titles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/istitled
func (b_ Browser) SetIsTitled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsTitled:"), value)
}


// A Boolean value indicating whether the view fills its frame rectangle with opaque content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/isopaque
func (b_ Browser) IsOpaque() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isOpaque"))
	return rv
}


// A Boolean value indicating whether the view fills its frame rectangle with opaque content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/isopaque
func (b_ Browser) SetIsOpaque(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsOpaque:"), value)
}



