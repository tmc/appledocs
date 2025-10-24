// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSTableRowView */


/* debug [class_header]: Header for NSTableRowView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TableRowView */
// An interface definition for the [TableRowView] class.
type ITableRowView interface {
	IView
	
/* debug [class_interface_properties]: Properties for TableRowView */
	// properties:
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle)
	IndentationForDropOperation() float64
	SetIndentationForDropOperation(value float64)
	InteriorBackgroundStyle() BackgroundStyle
	Emphasized() bool
	SetEmphasized(value bool)
	Floating() bool
	SetFloating(value bool)
	GroupRowStyle() bool
	SetGroupRowStyle(value bool)
	NextRowSelected() bool
	SetNextRowSelected(value bool)
	PreviousRowSelected() bool
	SetPreviousRowSelected(value bool)
	Selected() bool
	SetSelected(value bool)
	TargetForDropOperation() bool
	SetTargetForDropOperation(value bool)
	NumberOfColumns() int
	SelectionHighlightStyle() TableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle)
	IsEmphasized() bool
	SetIsEmphasized(value bool)
	IsFloating() bool
	SetIsFloating(value bool)
	IsGroupRowStyle() bool
	SetIsGroupRowStyle(value bool)
	IsNextRowSelected() bool
	SetIsNextRowSelected(value bool)
	IsPreviousRowSelected() bool
	SetIsPreviousRowSelected(value bool)
	IsSelected() bool
	SetIsSelected(value bool)
	IsTargetForDropOperation() bool
	SetIsTargetForDropOperation(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TableRowView */
	// methods:
	DrawBackgroundInRect(dirtyRect Rect /* not a class type */)
	DrawDraggingDestinationFeedbackInRect(dirtyRect Rect /* not a class type */)
	DrawSelectionInRect(dirtyRect Rect /* not a class type */)
	DrawSeparatorInRect(dirtyRect Rect /* not a class type */)
	ViewAtColumn(column int) objc.ID
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TableRowView */
// Alloc allocates a new instance without initialization.
func (tc _TableRowViewClass) Alloc() TableRowView {
	rv := objc.Send[TableRowView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TableRowView */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TableRowView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TableRowView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TableRowView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TableRowView */

// Draws the background of the row in the rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawBackground(in:)
func (t_ TableRowView) DrawBackgroundInRect(dirtyRect Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawBackgroundInRect:"), dirtyRect)
}/* debug [instance_methods/method]: DrawBackgroundInRect */


// Draws the row’s dragging destination feedback when the entire row is a drop target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawDraggingDestinationFeedback(in:)
func (t_ TableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawDraggingDestinationFeedbackInRect:"), dirtyRect)
}/* debug [instance_methods/method]: DrawDraggingDestinationFeedbackInRect */


// Draws the selected row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawSelection(in:)
func (t_ TableRowView) DrawSelectionInRect(dirtyRect Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawSelectionInRect:"), dirtyRect)
}/* debug [instance_methods/method]: DrawSelectionInRect */


// Draws the horizontal separator between table rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawSeparator(in:)
func (t_ TableRowView) DrawSeparatorInRect(dirtyRect Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawSeparatorInRect:"), dirtyRect)
}/* debug [instance_methods/method]: DrawSeparatorInRect */


// Provides access to the given view at a particular column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/view(atColumn:)
func (t_ TableRowView) ViewAtColumn(column int) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("viewAtColumn:"), column)
	return rv
}/* debug [instance_methods/method]: ViewAtColumn */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TableRowView */

// The background color of the row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/backgroundColor
func (t_ TableRowView) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The background color of the row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/backgroundColor
func (t_ TableRowView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// Specifies the dragging destination feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/draggingDestinationFeedbackStyle
func (t_ TableRowView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := objc.Send[TableViewDraggingDestinationFeedbackStyle](t_.ID, objc.Sel("draggingDestinationFeedbackStyle"))
	return rv
}/* debug [instance_properties/getter]: draggingDestinationFeedbackStyle */


// Specifies the dragging destination feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/draggingDestinationFeedbackStyle
func (t_ TableRowView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDraggingDestinationFeedbackStyle:"), value)
}/* debug [instance_properties/setter]: draggingDestinationFeedbackStyle */


// Defines the amount the drag target for a row should be indented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/indentationForDropOperation
func (t_ TableRowView) IndentationForDropOperation() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("indentationForDropOperation"))
	return rv
}/* debug [instance_properties/getter]: indentationForDropOperation */


// Defines the amount the drag target for a row should be indented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/indentationForDropOperation
func (t_ TableRowView) SetIndentationForDropOperation(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIndentationForDropOperation:"), value)
}/* debug [instance_properties/setter]: indentationForDropOperation */


// Specifies how the subviews should draw.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/interiorBackgroundStyle
func (t_ TableRowView) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.Send[BackgroundStyle](t_.ID, objc.Sel("interiorBackgroundStyle"))
	return rv
}/* debug [instance_properties/getter]: interiorBackgroundStyle */


// Determines whether the row will draw with the alternate or secondary color (unless overridden).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isEmphasized
func (t_ TableRowView) Emphasized() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("emphasized"))
	return rv
}/* debug [instance_properties/getter]: emphasized */


// Determines whether the row will draw with the alternate or secondary color (unless overridden).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isEmphasized
func (t_ TableRowView) SetEmphasized(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEmphasized:"), value)
}/* debug [instance_properties/setter]: emphasized */


// Specifies whether the row is drawn using the floating style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isFloating
func (t_ TableRowView) Floating() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("floating"))
	return rv
}/* debug [instance_properties/getter]: floating */


// Specifies whether the row is drawn using the floating style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isFloating
func (t_ TableRowView) SetFloating(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFloating:"), value)
}/* debug [instance_properties/setter]: floating */


// Specifies whether this row view is a group row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isGroupRowStyle
func (t_ TableRowView) GroupRowStyle() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("groupRowStyle"))
	return rv
}/* debug [instance_properties/getter]: groupRowStyle */


// Specifies whether this row view is a group row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isGroupRowStyle
func (t_ TableRowView) SetGroupRowStyle(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGroupRowStyle:"), value)
}/* debug [instance_properties/setter]: groupRowStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isNextRowSelected
func (t_ TableRowView) NextRowSelected() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("nextRowSelected"))
	return rv
}/* debug [instance_properties/getter]: nextRowSelected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isNextRowSelected
func (t_ TableRowView) SetNextRowSelected(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNextRowSelected:"), value)
}/* debug [instance_properties/setter]: nextRowSelected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isPreviousRowSelected
func (t_ TableRowView) PreviousRowSelected() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("previousRowSelected"))
	return rv
}/* debug [instance_properties/getter]: previousRowSelected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isPreviousRowSelected
func (t_ TableRowView) SetPreviousRowSelected(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreviousRowSelected:"), value)
}/* debug [instance_properties/setter]: previousRowSelected */


// Determines whether the row is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isSelected
func (t_ TableRowView) Selected() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("selected"))
	return rv
}/* debug [instance_properties/getter]: selected */


// Determines whether the row is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isSelected
func (t_ TableRowView) SetSelected(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelected:"), value)
}/* debug [instance_properties/setter]: selected */


// Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isTargetForDropOperation
func (t_ TableRowView) TargetForDropOperation() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("targetForDropOperation"))
	return rv
}/* debug [instance_properties/getter]: targetForDropOperation */


// Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/isTargetForDropOperation
func (t_ TableRowView) SetTargetForDropOperation(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTargetForDropOperation:"), value)
}/* debug [instance_properties/setter]: targetForDropOperation */


// Returns the number of columns represented by views in the table row view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/numberOfColumns
func (t_ TableRowView) NumberOfColumns() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfColumns"))
	return rv
}/* debug [instance_properties/getter]: numberOfColumns */


// Specifies the selection highlight style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/selectionHighlightStyle
func (t_ TableRowView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := objc.Send[TableViewSelectionHighlightStyle](t_.ID, objc.Sel("selectionHighlightStyle"))
	return rv
}/* debug [instance_properties/getter]: selectionHighlightStyle */


// Specifies the selection highlight style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/selectionHighlightStyle
func (t_ TableRowView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionHighlightStyle:"), value)
}/* debug [instance_properties/setter]: selectionHighlightStyle */


// Determines whether the row will draw with the alternate or secondary color (unless overridden).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isemphasized
func (t_ TableRowView) IsEmphasized() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEmphasized"))
	return rv
}/* debug [instance_properties/getter]: isEmphasized */


// Determines whether the row will draw with the alternate or secondary color (unless overridden).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isemphasized
func (t_ TableRowView) SetIsEmphasized(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEmphasized:"), value)
}/* debug [instance_properties/setter]: isEmphasized */


// Specifies whether the row is drawn using the floating style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isfloating
func (t_ TableRowView) IsFloating() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFloating"))
	return rv
}/* debug [instance_properties/getter]: isFloating */


// Specifies whether the row is drawn using the floating style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isfloating
func (t_ TableRowView) SetIsFloating(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFloating:"), value)
}/* debug [instance_properties/setter]: isFloating */


// Specifies whether this row view is a group row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isgrouprowstyle
func (t_ TableRowView) IsGroupRowStyle() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isGroupRowStyle"))
	return rv
}/* debug [instance_properties/getter]: isGroupRowStyle */


// Specifies whether this row view is a group row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isgrouprowstyle
func (t_ TableRowView) SetIsGroupRowStyle(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsGroupRowStyle:"), value)
}/* debug [instance_properties/setter]: isGroupRowStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isnextrowselected
func (t_ TableRowView) IsNextRowSelected() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isNextRowSelected"))
	return rv
}/* debug [instance_properties/getter]: isNextRowSelected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isnextrowselected
func (t_ TableRowView) SetIsNextRowSelected(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsNextRowSelected:"), value)
}/* debug [instance_properties/setter]: isNextRowSelected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/ispreviousrowselected
func (t_ TableRowView) IsPreviousRowSelected() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isPreviousRowSelected"))
	return rv
}/* debug [instance_properties/getter]: isPreviousRowSelected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/ispreviousrowselected
func (t_ TableRowView) SetIsPreviousRowSelected(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsPreviousRowSelected:"), value)
}/* debug [instance_properties/setter]: isPreviousRowSelected */


// Determines whether the row is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isselected
func (t_ TableRowView) IsSelected() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelected"))
	return rv
}/* debug [instance_properties/getter]: isSelected */


// Determines whether the row is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/isselected
func (t_ TableRowView) SetIsSelected(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelected:"), value)
}/* debug [instance_properties/setter]: isSelected */


// Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/istargetfordropoperation
func (t_ TableRowView) IsTargetForDropOperation() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isTargetForDropOperation"))
	return rv
}/* debug [instance_properties/getter]: isTargetForDropOperation */


// Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/istargetfordropoperation
func (t_ TableRowView) SetIsTargetForDropOperation(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsTargetForDropOperation:"), value)
}/* debug [instance_properties/setter]: isTargetForDropOperation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTableRowView */



