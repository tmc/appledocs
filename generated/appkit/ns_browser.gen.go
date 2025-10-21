// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An interface that displays a hierarchically organized list of data items that can be navigated and selected.
//
// A browser displays information using a set of columns, which are indexed from left to right. Each successive column displays the next level down in the data hierarchy. This class uses the class to implement its user interface. Browsers have the following components: Columns Scroll views Matrices Browser cells To the user, browsers display data in columns and rows within each column. These components are arranged in the following component hierarchy:
//
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


// A Boolean that indicates whether the user can select branch items.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/allowsbranchselection
func (b_ Browser) AllowsBranchSelection() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsBranchSelection"))
	return rv
}


// SetAllowsBranchSelection sets the value of the allowsBranchSelection property.
// A Boolean that indicates whether the user can select branch items.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/allowsbranchselection
func (b_ Browser) SetAllowsBranchSelection(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsBranchSelection:"), value)
}

// A Boolean that indicates whether there can be nothing selected.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/allowsemptyselection
func (b_ Browser) AllowsEmptySelection() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}


// SetAllowsEmptySelection sets the value of the allowsEmptySelection property.
// A Boolean that indicates whether there can be nothing selected.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/allowsemptyselection
func (b_ Browser) SetAllowsEmptySelection(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsEmptySelection:"), value)
}

// A Boolean that indicates whether the user can select multiple items.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/allowsmultipleselection
func (b_ Browser) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}


// SetAllowsMultipleSelection sets the value of the allowsMultipleSelection property.
// A Boolean that indicates whether the user can select multiple items.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/allowsmultipleselection
func (b_ Browser) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}

// A Boolean that indicates whether the browser allows keystroke-based selection (type select).
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/allowstypeselect
func (b_ Browser) AllowsTypeSelect() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsTypeSelect"))
	return rv
}


// SetAllowsTypeSelect sets the value of the allowsTypeSelect property.
// A Boolean that indicates whether the browser allows keystroke-based selection (type select).

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/allowstypeselect
func (b_ Browser) SetAllowsTypeSelect(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsTypeSelect:"), value)
}

// A Boolean that indicates whether the browser automatically hides its scroller.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/autohidesscroller
func (b_ Browser) AutohidesScroller() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("autohidesScroller"))
	return rv
}


// SetAutohidesScroller sets the value of the autohidesScroller property.
// A Boolean that indicates whether the browser automatically hides its scroller.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/autohidesscroller
func (b_ Browser) SetAutohidesScroller(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAutohidesScroller:"), value)
}

// The browser’s background color.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/backgroundcolor
func (b_ Browser) BackgroundColor() NSColor {
	rv := objc.Send[NSColor](b_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The browser’s background color.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/backgroundcolor
func (b_ Browser) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackgroundColor:"), value)
}

// The prototype
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/cellprototype
func (b_ Browser) CellPrototype() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("cellPrototype"))
	return rv
}


// SetCellPrototype sets the value of the cellPrototype property.
// The prototype

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/cellprototype
func (b_ Browser) SetCellPrototype(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCellPrototype:"), value)
}

// The column number of the cell that the user clicked to display a context menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/clickedcolumn
func (b_ Browser) ClickedColumn() int {
	rv := objc.Send[int](b_.ID, objc.Sel("clickedColumn"))
	return rv
}


// SetClickedColumn sets the value of the clickedColumn property.
// The column number of the cell that the user clicked to display a context menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/clickedcolumn
func (b_ Browser) SetClickedColumn(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setClickedColumn:"), value)
}

// The row number of the cell that the user clicked to display a context menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/clickedrow
func (b_ Browser) ClickedRow() int {
	rv := objc.Send[int](b_.ID, objc.Sel("clickedRow"))
	return rv
}


// SetClickedRow sets the value of the clickedRow property.
// The row number of the cell that the user clicked to display a context menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/clickedrow
func (b_ Browser) SetClickedRow(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setClickedRow:"), value)
}

// A constant indicating the browser’s column resizing type.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/columnresizingtype-swift.property
func (b_ Browser) ColumnResizingType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("columnResizingType"))
	return rv
}


// SetColumnResizingType sets the value of the columnResizingType property.
// A constant indicating the browser’s column resizing type.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/columnresizingtype-swift.property
func (b_ Browser) SetColumnResizingType(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setColumnResizingType:"), value)
}

// The name used to automatically save the browser’s column configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/columnsautosavename-swift.property
func (b_ Browser) ColumnsAutosaveName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("columnsAutosaveName"))
	return rv
}


// SetColumnsAutosaveName sets the value of the columnsAutosaveName property.
// The name used to automatically save the browser’s column configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/columnsautosavename-swift.property
func (b_ Browser) SetColumnsAutosaveName(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setColumnsAutosaveName:"), value)
}

// The browser’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/delegate
func (b_ Browser) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The browser’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/delegate
func (b_ Browser) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}

// The browser’s double-click action method.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/doubleaction
func (b_ Browser) DoubleAction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("doubleAction"))
	return rv
}


// SetDoubleAction sets the value of the doubleAction property.
// The browser’s double-click action method.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/doubleaction
func (b_ Browser) SetDoubleAction(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDoubleAction:"), value)
}

// The index of the first visible column.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/firstvisiblecolumn
func (b_ Browser) FirstVisibleColumn() int {
	rv := objc.Send[int](b_.ID, objc.Sel("firstVisibleColumn"))
	return rv
}


// SetFirstVisibleColumn sets the value of the firstVisibleColumn property.
// The index of the first visible column.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/firstvisiblecolumn
func (b_ Browser) SetFirstVisibleColumn(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFirstVisibleColumn:"), value)
}

// A Boolean that indicates whether the browser has a horizontal scroller.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/hashorizontalscroller
func (b_ Browser) HasHorizontalScroller() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("hasHorizontalScroller"))
	return rv
}


// SetHasHorizontalScroller sets the value of the hasHorizontalScroller property.
// A Boolean that indicates whether the browser has a horizontal scroller.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/hashorizontalscroller
func (b_ Browser) SetHasHorizontalScroller(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setHasHorizontalScroller:"), value)
}

// A Boolean that indicates whether column 0 is loaded.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/isloaded
func (b_ Browser) IsLoaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isLoaded"))
	return rv
}


// SetIsLoaded sets the value of the isLoaded property.
// A Boolean that indicates whether column 0 is loaded.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/isloaded
func (b_ Browser) SetIsLoaded(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsLoaded:"), value)
}

// A Boolean that indicates whether columns display titles.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/istitled
func (b_ Browser) IsTitled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTitled"))
	return rv
}


// SetIsTitled sets the value of the isTitled property.
// A Boolean that indicates whether columns display titles.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/istitled
func (b_ Browser) SetIsTitled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsTitled:"), value)
}

// The index of the last column loaded.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/lastcolumn
func (b_ Browser) LastColumn() int {
	rv := objc.Send[int](b_.ID, objc.Sel("lastColumn"))
	return rv
}


// SetLastColumn sets the value of the lastColumn property.
// The index of the last column loaded.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/lastcolumn
func (b_ Browser) SetLastColumn(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLastColumn:"), value)
}

// The index of the last visible column.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/lastvisiblecolumn
func (b_ Browser) LastVisibleColumn() int {
	rv := objc.Send[int](b_.ID, objc.Sel("lastVisibleColumn"))
	return rv
}


// SetLastVisibleColumn sets the value of the lastVisibleColumn property.
// The index of the last visible column.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/lastvisiblecolumn
func (b_ Browser) SetLastVisibleColumn(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLastVisibleColumn:"), value)
}

// The maximum number of visible columns.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/maxvisiblecolumns
func (b_ Browser) MaxVisibleColumns() int {
	rv := objc.Send[int](b_.ID, objc.Sel("maxVisibleColumns"))
	return rv
}


// SetMaxVisibleColumns sets the value of the maxVisibleColumns property.
// The maximum number of visible columns.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/maxvisiblecolumns
func (b_ Browser) SetMaxVisibleColumns(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMaxVisibleColumns:"), value)
}

// The minimum column width, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/mincolumnwidth
func (b_ Browser) MinColumnWidth() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("minColumnWidth"))
	return rv
}


// SetMinColumnWidth sets the value of the minColumnWidth property.
// The minimum column width, in pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/mincolumnwidth
func (b_ Browser) SetMinColumnWidth(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMinColumnWidth:"), value)
}

// The number of visible columns.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/numberofvisiblecolumns
func (b_ Browser) NumberOfVisibleColumns() int {
	rv := objc.Send[int](b_.ID, objc.Sel("numberOfVisibleColumns"))
	return rv
}


// SetNumberOfVisibleColumns sets the value of the numberOfVisibleColumns property.
// The number of visible columns.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/numberofvisiblecolumns
func (b_ Browser) SetNumberOfVisibleColumns(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNumberOfVisibleColumns:"), value)
}

// The path separator.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/pathseparator
func (b_ Browser) PathSeparator() string {
	rv := objc.Send[string](b_.ID, objc.Sel("pathSeparator"))
	return rv
}


// SetPathSeparator sets the value of the pathSeparator property.
// The path separator.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/pathseparator
func (b_ Browser) SetPathSeparator(value string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPathSeparator:"), objc.String(value))
}

// A Boolean that indicates whether the browser is set to resize all columns simultaneously rather than resizing a single column at a time.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/prefersallcolumnuserresizing
func (b_ Browser) PrefersAllColumnUserResizing() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("prefersAllColumnUserResizing"))
	return rv
}


// SetPrefersAllColumnUserResizing sets the value of the prefersAllColumnUserResizing property.
// A Boolean that indicates whether the browser is set to resize all columns simultaneously rather than resizing a single column at a time.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/prefersallcolumnuserresizing
func (b_ Browser) SetPrefersAllColumnUserResizing(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrefersAllColumnUserResizing:"), value)
}

// A Boolean that indicates whether the browser reuses matrix objects after their columns are unloaded.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/reusescolumns
func (b_ Browser) ReusesColumns() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("reusesColumns"))
	return rv
}


// SetReusesColumns sets the value of the reusesColumns property.
// A Boolean that indicates whether the browser reuses matrix objects after their columns are unloaded.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/reusescolumns
func (b_ Browser) SetReusesColumns(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setReusesColumns:"), value)
}

// The height of the browser’s rows.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/rowheight
func (b_ Browser) RowHeight() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("rowHeight"))
	return rv
}


// SetRowHeight sets the value of the rowHeight property.
// The height of the browser’s rows.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/rowheight
func (b_ Browser) SetRowHeight(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setRowHeight:"), value)
}

// All cells selected in the rightmost column.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/selectedcells
func (b_ Browser) SelectedCells() NSCell {
	rv := objc.Send[NSCell](b_.ID, objc.Sel("selectedCells"))
	return rv
}


// SetSelectedCells sets the value of the selectedCells property.
// All cells selected in the rightmost column.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/selectedcells
func (b_ Browser) SetSelectedCells(value ICell) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSelectedCells:"), value)
}

// The index of the last column with a selected item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/selectedcolumn
func (b_ Browser) SelectedColumn() int {
	rv := objc.Send[int](b_.ID, objc.Sel("selectedColumn"))
	return rv
}


// SetSelectedColumn sets the value of the selectedColumn property.
// The index of the last column with a selected item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/selectedcolumn
func (b_ Browser) SetSelectedColumn(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSelectedColumn:"), value)
}

// The index path of the item selected in the browser.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/selectionindexpath
func (b_ Browser) SelectionIndexPath() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](b_.ID, objc.Sel("selectionIndexPath"))
	return rv
}


// SetSelectionIndexPath sets the value of the selectionIndexPath property.
// The index path of the item selected in the browser.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/selectionindexpath
func (b_ Browser) SetSelectionIndexPath(value foundation.IIndexPath) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSelectionIndexPath:"), value)
}

// An array containing the index paths of all items selected in the browser.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/selectionindexpaths
func (b_ Browser) SelectionIndexPaths() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](b_.ID, objc.Sel("selectionIndexPaths"))
	return rv
}


// SetSelectionIndexPaths sets the value of the selectionIndexPaths property.
// An array containing the index paths of all items selected in the browser.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/selectionindexpaths
func (b_ Browser) SetSelectionIndexPaths(value foundation.IIndexPath) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSelectionIndexPaths:"), value)
}

// A Boolean that indicates whether pressing an arrow key causes an action message to be sent.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/sendsactiononarrowkeys
func (b_ Browser) SendsActionOnArrowKeys() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("sendsActionOnArrowKeys"))
	return rv
}


// SetSendsActionOnArrowKeys sets the value of the sendsActionOnArrowKeys property.
// A Boolean that indicates whether pressing an arrow key causes an action message to be sent.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/sendsactiononarrowkeys
func (b_ Browser) SetSendsActionOnArrowKeys(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSendsActionOnArrowKeys:"), value)
}

// A Boolean that indicates whether columns are separated by bezeled borders.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/separatescolumns
func (b_ Browser) SeparatesColumns() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("separatesColumns"))
	return rv
}


// SetSeparatesColumns sets the value of the separatesColumns property.
// A Boolean that indicates whether columns are separated by bezeled borders.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/separatescolumns
func (b_ Browser) SetSeparatesColumns(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSeparatesColumns:"), value)
}

// A Boolean that indicates whether a column takes its title from the selected cell in the previous column.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/takestitlefrompreviouscolumn
func (b_ Browser) TakesTitleFromPreviousColumn() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("takesTitleFromPreviousColumn"))
	return rv
}


// SetTakesTitleFromPreviousColumn sets the value of the takesTitleFromPreviousColumn property.
// A Boolean that indicates whether a column takes its title from the selected cell in the previous column.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/takestitlefrompreviouscolumn
func (b_ Browser) SetTakesTitleFromPreviousColumn(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTakesTitleFromPreviousColumn:"), value)
}

// The height of the column titles for the browser.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/titleheight
func (b_ Browser) TitleHeight() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("titleHeight"))
	return rv
}


// SetTitleHeight sets the value of the titleHeight property.
// The height of the column titles for the browser.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowser/titleheight
func (b_ Browser) SetTitleHeight(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitleHeight:"), value)
}

// A Boolean value indicating whether the view fills its frame rectangle with opaque content.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/isopaque
func (b_ Browser) IsOpaque() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isOpaque"))
	return rv
}


// SetIsOpaque sets the value of the isOpaque property.
// A Boolean value indicating whether the view fills its frame rectangle with opaque content.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/isopaque
func (b_ Browser) SetIsOpaque(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsOpaque:"), value)
}



