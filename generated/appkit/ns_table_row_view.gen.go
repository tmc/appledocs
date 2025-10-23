// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle /* not a class type */
	SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle /* not a class type */)
	IndentationForDropOperation() float64 /* primitive/slice/pointer. */
	SetIndentationForDropOperation(value float64 /* primitive/slice/pointer. */)
	InteriorBackgroundStyle() BackgroundStyle
	Emphasized() bool /* primitive/slice/pointer. */
	SetEmphasized(value bool /* primitive/slice/pointer. */)
	Floating() bool /* primitive/slice/pointer. */
	SetFloating(value bool /* primitive/slice/pointer. */)
	GroupRowStyle() bool /* primitive/slice/pointer. */
	SetGroupRowStyle(value bool /* primitive/slice/pointer. */)
	NextRowSelected() bool /* primitive/slice/pointer. */
	SetNextRowSelected(value bool /* primitive/slice/pointer. */)
	PreviousRowSelected() bool /* primitive/slice/pointer. */
	SetPreviousRowSelected(value bool /* primitive/slice/pointer. */)
	Selected() bool /* primitive/slice/pointer. */
	SetSelected(value bool /* primitive/slice/pointer. */)
	TargetForDropOperation() bool /* primitive/slice/pointer. */
	SetTargetForDropOperation(value bool /* primitive/slice/pointer. */)
	NumberOfColumns() int /* primitive/slice/pointer. */
	SelectionHighlightStyle() TableViewSelectionHighlightStyle /* not a class type */
	SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle /* not a class type */)
	IsEmphasized() bool /* primitive/slice/pointer. */
	SetIsEmphasized(value bool /* primitive/slice/pointer. */)
	IsFloating() bool /* primitive/slice/pointer. */
	SetIsFloating(value bool /* primitive/slice/pointer. */)
	IsGroupRowStyle() bool /* primitive/slice/pointer. */
	SetIsGroupRowStyle(value bool /* primitive/slice/pointer. */)
	IsNextRowSelected() bool /* primitive/slice/pointer. */
	SetIsNextRowSelected(value bool /* primitive/slice/pointer. */)
	IsPreviousRowSelected() bool /* primitive/slice/pointer. */
	SetIsPreviousRowSelected(value bool /* primitive/slice/pointer. */)
	IsSelected() bool /* primitive/slice/pointer. */
	SetIsSelected(value bool /* primitive/slice/pointer. */)
	IsTargetForDropOperation() bool /* primitive/slice/pointer. */
	SetIsTargetForDropOperation(value bool /* primitive/slice/pointer. */)
	// methods:
	DrawBackgroundInRect(dirtyRect objc.IObject /* cross-framework Rect */)
	DrawDraggingDestinationFeedbackInRect(dirtyRect objc.IObject /* cross-framework Rect */)
	DrawSelectionInRect(dirtyRect objc.IObject /* cross-framework Rect */)
	DrawSeparatorInRect(dirtyRect objc.IObject /* cross-framework Rect */)
	ViewAtColumn(column int /* primitive/slice/pointer. */) objc.ID
}

// The view shown for a row in a table view.
//
// is responsible for displaying attributes associated with the row, including the selection highlight, and group row look.


// The view shown for a row in a table view.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawBackground(in:)
func (t_ TableRowView) DrawBackgroundInRect(dirtyRect objc.IObject /* cross-framework Rect */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawBackgroundInRect:"), dirtyRect)
}


// Draws the row’s dragging destination feedback when the entire row is a drop target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawDraggingDestinationFeedback(in:)
func (t_ TableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect objc.IObject /* cross-framework Rect */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawDraggingDestinationFeedbackInRect:"), dirtyRect)
}


// Draws the selected row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawSelection(in:)
func (t_ TableRowView) DrawSelectionInRect(dirtyRect objc.IObject /* cross-framework Rect */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawSelectionInRect:"), dirtyRect)
}


// Draws the horizontal separator between table rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawSeparator(in:)
func (t_ TableRowView) DrawSeparatorInRect(dirtyRect objc.IObject /* cross-framework Rect */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawSeparatorInRect:"), dirtyRect)
}


// Provides access to the given view at a particular column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/view(atColumn:)
func (t_ TableRowView) ViewAtColumn(column int /* primitive/slice/pointer. */) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("viewAtColumn:"), column)
	return rv
}


// The background color of the row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/backgroundColor
func (t_ TableRowView) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color of the row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/backgroundColor
func (t_ TableRowView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// Specifies the dragging destination feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/draggingDestinationFeedbackStyle
func (t_ TableRowView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle /* not a class type */ {
	rv := objc.Send[TableViewDraggingDestinationFeedbackStyle](t_.ID, objc.Sel("draggingDestinationFeedbackStyle"))
	return rv
}


// Specifies the dragging destination feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/draggingDestinationFeedbackStyle
func (t_ TableRowView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDraggingDestinationFeedbackStyle:"), value)
}


// Defines the amount the drag target for a row should be indented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/indentationForDropOperation
func (t_ TableRowView) IndentationForDropOperation() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](t_.ID, objc.Sel("indentationForDropOperation"))
	return rv
}


// Defines the amount the drag target for a row should be indented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/indentationForDropOperation
func (t_ TableRowView) SetIndentationForDropOperation(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIndentationForDropOperation:"), value)
}


// Specifies how the subviews should draw.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/interiorBackgroundStyle
func (t_ TableRowView) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.Send[BackgroundStyle](t_.ID, objc.Sel("interiorBackgroundStyle"))
	return rv
}


// Determines whether the row will draw with the alternate or secondary color (unless overridden).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isEmphasized
func (t_ TableRowView) Emphasized() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("emphasized"))
	return rv
}


// Determines whether the row will draw with the alternate or secondary color (unless overridden).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isEmphasized
func (t_ TableRowView) SetEmphasized(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEmphasized:"), value)
}


// Specifies whether the row is drawn using the floating style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isFloating
func (t_ TableRowView) Floating() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("floating"))
	return rv
}


// Specifies whether the row is drawn using the floating style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isFloating
func (t_ TableRowView) SetFloating(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFloating:"), value)
}


// Specifies whether this row view is a group row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isGroupRowStyle
func (t_ TableRowView) GroupRowStyle() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("groupRowStyle"))
	return rv
}


// Specifies whether this row view is a group row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isGroupRowStyle
func (t_ TableRowView) SetGroupRowStyle(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGroupRowStyle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isNextRowSelected
func (t_ TableRowView) NextRowSelected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("nextRowSelected"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isNextRowSelected
func (t_ TableRowView) SetNextRowSelected(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNextRowSelected:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isPreviousRowSelected
func (t_ TableRowView) PreviousRowSelected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("previousRowSelected"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isPreviousRowSelected
func (t_ TableRowView) SetPreviousRowSelected(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreviousRowSelected:"), value)
}


// Determines whether the row is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isSelected
func (t_ TableRowView) Selected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("selected"))
	return rv
}


// Determines whether the row is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isSelected
func (t_ TableRowView) SetSelected(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelected:"), value)
}


// Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isTargetForDropOperation
func (t_ TableRowView) TargetForDropOperation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("targetForDropOperation"))
	return rv
}


// Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isTargetForDropOperation
func (t_ TableRowView) SetTargetForDropOperation(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTargetForDropOperation:"), value)
}


// Returns the number of columns represented by views in the table row view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/numberOfColumns
func (t_ TableRowView) NumberOfColumns() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfColumns"))
	return rv
}


// Specifies the selection highlight style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/selectionHighlightStyle
func (t_ TableRowView) SelectionHighlightStyle() TableViewSelectionHighlightStyle /* not a class type */ {
	rv := objc.Send[TableViewSelectionHighlightStyle](t_.ID, objc.Sel("selectionHighlightStyle"))
	return rv
}


// Specifies the selection highlight style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/selectionHighlightStyle
func (t_ TableRowView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionHighlightStyle:"), value)
}


// Determines whether the row will draw with the alternate or secondary color (unless overridden).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isemphasized
func (t_ TableRowView) IsEmphasized() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEmphasized"))
	return rv
}


// Determines whether the row will draw with the alternate or secondary color (unless overridden).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isemphasized
func (t_ TableRowView) SetIsEmphasized(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEmphasized:"), value)
}


// Specifies whether the row is drawn using the floating style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isfloating
func (t_ TableRowView) IsFloating() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFloating"))
	return rv
}


// Specifies whether the row is drawn using the floating style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isfloating
func (t_ TableRowView) SetIsFloating(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFloating:"), value)
}


// Specifies whether this row view is a group row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isgrouprowstyle
func (t_ TableRowView) IsGroupRowStyle() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isGroupRowStyle"))
	return rv
}


// Specifies whether this row view is a group row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isgrouprowstyle
func (t_ TableRowView) SetIsGroupRowStyle(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsGroupRowStyle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isnextrowselected
func (t_ TableRowView) IsNextRowSelected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isNextRowSelected"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isnextrowselected
func (t_ TableRowView) SetIsNextRowSelected(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsNextRowSelected:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/ispreviousrowselected
func (t_ TableRowView) IsPreviousRowSelected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isPreviousRowSelected"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/ispreviousrowselected
func (t_ TableRowView) SetIsPreviousRowSelected(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsPreviousRowSelected:"), value)
}


// Determines whether the row is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isselected
func (t_ TableRowView) IsSelected() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelected"))
	return rv
}


// Determines whether the row is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isselected
func (t_ TableRowView) SetIsSelected(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelected:"), value)
}


// Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/istargetfordropoperation
func (t_ TableRowView) IsTargetForDropOperation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isTargetForDropOperation"))
	return rv
}


// Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/istargetfordropoperation
func (t_ TableRowView) SetIsTargetForDropOperation(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsTargetForDropOperation:"), value)
}



