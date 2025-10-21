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
	RemoveRowsAtIndexesWithAnimation(indexes unsafe.Pointer, animationOptions unsafe.Pointer)
	RowAtPoint(point coregraphics.CGPoint) int
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
func (t_ TableView) RemoveRowsAtIndexesWithAnimation(indexes unsafe.Pointer, animationOptions unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeRowsAtIndexes:withAnimation:"), indexes, animationOptions)
}

// Returns the index of the row the specified point lies in.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/row(at:)
func (t_ TableView) RowAtPoint(point coregraphics.CGPoint) int {
	rv := objc.Send[int](t_.ID, objc.Sel("rowAtPoint:"), point)
	return rv
}

// The color used to draw the background of the table.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/backgroundColor
func (t_ TableView) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The color used to draw the background of the table.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/backgroundColor
func (t_ TableView) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}

// The table view’s column autoresizing style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/columnAutoresizingStyle-swift.property
func (t_ TableView) ColumnAutoresizingStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("columnAutoresizingStyle"))
	return rv
}


// SetColumnAutoresizingStyle sets the value of the columnAutoresizingStyle property.
// The table view’s column autoresizing style.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/columnAutoresizingStyle-swift.property
func (t_ TableView) SetColumnAutoresizingStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColumnAutoresizingStyle:"), value)
}

// The feedback style displayed when the user drags over the table view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/draggingDestinationFeedbackStyle-swift.property
func (t_ TableView) DraggingDestinationFeedbackStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("draggingDestinationFeedbackStyle"))
	return rv
}


// SetDraggingDestinationFeedbackStyle sets the value of the draggingDestinationFeedbackStyle property.
// The feedback style displayed when the user drags over the table view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/draggingDestinationFeedbackStyle-swift.property
func (t_ TableView) SetDraggingDestinationFeedbackStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDraggingDestinationFeedbackStyle:"), value)
}

// The effective row size style for the table.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/effectiveRowSizeStyle
func (t_ TableView) EffectiveRowSizeStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("effectiveRowSizeStyle"))
	return rv
}

// The grid lines drawn by the table view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridStyleMask
func (t_ TableView) GridStyleMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("gridStyleMask"))
	return rv
}


// SetGridStyleMask sets the value of the gridStyleMask property.
// The grid lines drawn by the table view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/gridStyleMask
func (t_ TableView) SetGridStyleMask(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGridStyleMask:"), value)
}

// The row size style (small, medium, large, or custom) used by the table view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowSizeStyle-swift.property
func (t_ TableView) RowSizeStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("rowSizeStyle"))
	return rv
}


// SetRowSizeStyle sets the value of the rowSizeStyle property.
// The row size style (small, medium, large, or custom) used by the table view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/rowSizeStyle-swift.property
func (t_ TableView) SetRowSizeStyle(value unsafe.Pointer) {
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



