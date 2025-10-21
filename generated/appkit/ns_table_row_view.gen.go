// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [TableRowView] class.
var (
	TableRowViewClass     _TableRowViewClass
	TableRowViewClassOnce sync.Once
)

func getTableRowViewClass() _TableRowViewClass {
	TableRowViewClassOnce.Do(func() {
		TableRowViewClass = _TableRowViewClass{objc.GetClass("NSTableRowView")}
	})
	return TableRowViewClass
}

type _TableRowViewClass struct {
	class objc.Class
}

// An interface definition for the [TableRowView] class.
type ITableRowView interface {
	IView
	DrawBackgroundInRect(dirtyRect coregraphics.CGRect)
	DrawDraggingDestinationFeedbackInRect(dirtyRect coregraphics.CGRect)
	DrawSelectionInRect(dirtyRect coregraphics.CGRect)
	DrawSeparatorInRect(dirtyRect coregraphics.CGRect)
	ViewAtColumn(column int) objc.ID
}

// The view shown for a row in a table view.
//
// is responsible for displaying attributes associated with the row, including the selection highlight, and group row look.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView
type TableRowView struct {
	View
}

// TableRowViewFrom constructs a [TableRowView] from an unsafe.Pointer.
//
// The view shown for a row in a table view.
func TableRowViewFrom(ptr unsafe.Pointer) TableRowView {
	return TableRowView{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TableRowViewClass) Alloc() TableRowView {
	rv := objc.Send[TableRowView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TableRowViewClass) New() TableRowView {
	rv := objc.Send[TableRowView](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableRowView) Init() TableRowView {
	rv := objc.Send[TableRowView](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableRowView) Autorelease() TableRowView {
	rv := objc.Send[TableRowView](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableRowView creates a new TableRowView instance.
func NewTableRowView() TableRowView {
	return getTableRowViewClass().New()
}


// Draws the background of the row in the rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawBackground(in:)
func (t_ TableRowView) DrawBackgroundInRect(dirtyRect coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawBackgroundInRect:"), dirtyRect)
}

// Draws the row’s dragging destination feedback when the entire row is a drop target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawDraggingDestinationFeedback(in:)
func (t_ TableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawDraggingDestinationFeedbackInRect:"), dirtyRect)
}

// Draws the selected row.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawSelection(in:)
func (t_ TableRowView) DrawSelectionInRect(dirtyRect coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawSelectionInRect:"), dirtyRect)
}

// Draws the horizontal separator between table rows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawSeparator(in:)
func (t_ TableRowView) DrawSeparatorInRect(dirtyRect coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawSeparatorInRect:"), dirtyRect)
}

// Provides access to the given view at a particular column.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/view(atColumn:)
func (t_ TableRowView) ViewAtColumn(column int) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("viewAtColumn:"), column)
	return rv
}

// The background color of the row.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/backgroundColor
func (t_ TableRowView) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The background color of the row.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/backgroundColor
func (t_ TableRowView) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}

// Specifies the dragging destination feedback style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/draggingDestinationFeedbackStyle
func (t_ TableRowView) DraggingDestinationFeedbackStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("draggingDestinationFeedbackStyle"))
	return rv
}


// SetDraggingDestinationFeedbackStyle sets the value of the draggingDestinationFeedbackStyle property.
// Specifies the dragging destination feedback style.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/draggingDestinationFeedbackStyle
func (t_ TableRowView) SetDraggingDestinationFeedbackStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDraggingDestinationFeedbackStyle:"), value)
}

// Defines the amount the drag target for a row should be indented.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/indentationForDropOperation
func (t_ TableRowView) IndentationForDropOperation() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("indentationForDropOperation"))
	return rv
}


// SetIndentationForDropOperation sets the value of the indentationForDropOperation property.
// Defines the amount the drag target for a row should be indented.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/indentationForDropOperation
func (t_ TableRowView) SetIndentationForDropOperation(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIndentationForDropOperation:"), value)
}

// Specifies how the subviews should draw.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/interiorBackgroundStyle
func (t_ TableRowView) InteriorBackgroundStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("interiorBackgroundStyle"))
	return rv
}

// Determines whether the row will draw with the alternate or secondary color (unless overridden).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isEmphasized
func (t_ TableRowView) Emphasized() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("emphasized"))
	return rv
}


// SetEmphasized sets the value of the emphasized property.
// Determines whether the row will draw with the alternate or secondary color (unless overridden).

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isEmphasized
func (t_ TableRowView) SetEmphasized(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEmphasized:"), value)
}

// Specifies whether the row is drawn using the floating style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isFloating
func (t_ TableRowView) Floating() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("floating"))
	return rv
}


// SetFloating sets the value of the floating property.
// Specifies whether the row is drawn using the floating style.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isFloating
func (t_ TableRowView) SetFloating(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFloating:"), value)
}

// Specifies whether this row view is a group row.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isGroupRowStyle
func (t_ TableRowView) GroupRowStyle() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("groupRowStyle"))
	return rv
}


// SetGroupRowStyle sets the value of the groupRowStyle property.
// Specifies whether this row view is a group row.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isGroupRowStyle
func (t_ TableRowView) SetGroupRowStyle(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGroupRowStyle:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isNextRowSelected
func (t_ TableRowView) NextRowSelected() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("nextRowSelected"))
	return rv
}


// SetNextRowSelected sets the value of the nextRowSelected property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isNextRowSelected
func (t_ TableRowView) SetNextRowSelected(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNextRowSelected:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isPreviousRowSelected
func (t_ TableRowView) PreviousRowSelected() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("previousRowSelected"))
	return rv
}


// SetPreviousRowSelected sets the value of the previousRowSelected property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isPreviousRowSelected
func (t_ TableRowView) SetPreviousRowSelected(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreviousRowSelected:"), value)
}

// Determines whether the row is selected.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isSelected
func (t_ TableRowView) Selected() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("selected"))
	return rv
}


// SetSelected sets the value of the selected property.
// Determines whether the row is selected.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isSelected
func (t_ TableRowView) SetSelected(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelected:"), value)
}

// Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isTargetForDropOperation
func (t_ TableRowView) TargetForDropOperation() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("targetForDropOperation"))
	return rv
}


// SetTargetForDropOperation sets the value of the targetForDropOperation property.
// Specifies whether this row will draw a drop indicator based on the current dragging feedback style.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isTargetForDropOperation
func (t_ TableRowView) SetTargetForDropOperation(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTargetForDropOperation:"), value)
}

// Returns the number of columns represented by views in the table row view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/numberOfColumns
func (t_ TableRowView) NumberOfColumns() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfColumns"))
	return rv
}

// Specifies the selection highlight style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/selectionHighlightStyle
func (t_ TableRowView) SelectionHighlightStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("selectionHighlightStyle"))
	return rv
}


// SetSelectionHighlightStyle sets the value of the selectionHighlightStyle property.
// Specifies the selection highlight style.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/selectionHighlightStyle
func (t_ TableRowView) SetSelectionHighlightStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionHighlightStyle:"), value)
}



