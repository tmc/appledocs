
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableRowView] class.
var TableRowViewClass _TableRowViewClass

func init() {
	TableRowViewClass = _TableRowViewClass{objc.GetClass("NSTableRowView")}
}

type _TableRowViewClass struct {
	objc.Class
}

// An interface definition for the [TableRowView] class.
type ITableRowView interface {
	ID() objc.ID
	DrawBackgroundInRect(dirtyRect unsafe.Pointer)
	DrawDraggingDestinationFeedbackInRect(dirtyRect unsafe.Pointer)
	DrawSelectionInRect(dirtyRect unsafe.Pointer)
	DrawSeparatorInRect(dirtyRect unsafe.Pointer)
	ViewAtColumn(column int) objc.ID
}

type TableRowView struct {
	id objc.ID
}

func TableRowViewFrom(ptr unsafe.Pointer) TableRowView {
	return TableRowView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TableRowView) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TableRowViewClass) Alloc() TableRowView {
	rv := objc.Send[TableRowView](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TableRowViewClass) New() TableRowView {
	rv := objc.Send[TableRowView](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTableRowView creates and returns a new initialized instance.
func NewTableRowView() TableRowView {
	return TableRowViewClass.New()
}

// Init initializes the instance.
func (t_ TableRowView) Init() TableRowView {
	rv := objc.Send[TableRowView](t_.ID(), selInit)
	return rv
}
// Draws the background of the row in the rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/drawBackground(in:)
func (t_ TableRowView) DrawBackgroundInRect(dirtyRect unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("drawBackgroundInRect:"), dirtyRect)
}
// Draws the row’s dragging destination feedback when the entire row is a drop target. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/drawDraggingDestinationFeedback(in:)
func (t_ TableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("drawDraggingDestinationFeedbackInRect:"), dirtyRect)
}
// Draws the selected row. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/drawSelection(in:)
func (t_ TableRowView) DrawSelectionInRect(dirtyRect unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("drawSelectionInRect:"), dirtyRect)
}
// Draws the horizontal separator between table rows. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/drawSeparator(in:)
func (t_ TableRowView) DrawSeparatorInRect(dirtyRect unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("drawSeparatorInRect:"), dirtyRect)
}
// Provides access to the given view at a particular column. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/view(atColumn:)
func (t_ TableRowView) ViewAtColumn(column int) objc.ID {
	rv := objc.Send[objc.ID](t_.ID(), objc.RegisterName("viewAtColumn:"), column)
	return rv
}
// The background color of the row. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/backgroundColor
func (t_ TableRowView) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("backgroundColor"))
	return rv
}
// SetBackgroundColor sets the value of the backgroundColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/backgroundColor
func (t_ TableRowView) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setBackgroundColor:"), value)
}
// Specifies the dragging destination feedback style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/draggingDestinationFeedbackStyle
func (t_ TableRowView) DraggingDestinationFeedbackStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("draggingDestinationFeedbackStyle"))
	return rv
}
// SetDraggingDestinationFeedbackStyle sets the value of the draggingDestinationFeedbackStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/draggingDestinationFeedbackStyle
func (t_ TableRowView) SetDraggingDestinationFeedbackStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setDraggingDestinationFeedbackStyle:"), value)
}
// Defines the amount the drag target for a row should be indented. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/indentationForDropOperation
func (t_ TableRowView) IndentationForDropOperation() float64 {
	rv := objc.Send[float64](t_.ID(), objc.RegisterName("indentationForDropOperation"))
	return rv
}
// SetIndentationForDropOperation sets the value of the indentationForDropOperation property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/indentationForDropOperation
func (t_ TableRowView) SetIndentationForDropOperation(value float64) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setIndentationForDropOperation:"), value)
}
// Specifies how the subviews should draw. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/interiorBackgroundStyle
func (t_ TableRowView) InteriorBackgroundStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("interiorBackgroundStyle"))
	return rv
}
// Determines whether the row will draw with the alternate or secondary color (unless overridden). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isEmphasized
func (t_ TableRowView) Emphasized() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("emphasized"))
	return rv
}
// SetEmphasized sets the value of the emphasized property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isEmphasized
func (t_ TableRowView) SetEmphasized(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setEmphasized:"), value)
}
// Specifies whether the row is drawn using the floating style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isFloating
func (t_ TableRowView) Floating() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("floating"))
	return rv
}
// SetFloating sets the value of the floating property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isFloating
func (t_ TableRowView) SetFloating(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setFloating:"), value)
}
// Specifies whether this row view is a group row. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isGroupRowStyle
func (t_ TableRowView) GroupRowStyle() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("groupRowStyle"))
	return rv
}
// SetGroupRowStyle sets the value of the groupRowStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isGroupRowStyle
func (t_ TableRowView) SetGroupRowStyle(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setGroupRowStyle:"), value)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isNextRowSelected
func (t_ TableRowView) NextRowSelected() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("nextRowSelected"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isNextRowSelected
func (t_ TableRowView) SetNextRowSelected(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setNextRowSelected:"), value)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isPreviousRowSelected
func (t_ TableRowView) PreviousRowSelected() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("previousRowSelected"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isPreviousRowSelected
func (t_ TableRowView) SetPreviousRowSelected(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setPreviousRowSelected:"), value)
}
// Determines whether the row is selected. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isSelected
func (t_ TableRowView) Selected() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("selected"))
	return rv
}
// SetSelected sets the value of the selected property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isSelected
func (t_ TableRowView) SetSelected(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setSelected:"), value)
}
// Specifies whether this row will draw a drop indicator based on the current dragging feedback style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isTargetForDropOperation
func (t_ TableRowView) TargetForDropOperation() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("targetForDropOperation"))
	return rv
}
// SetTargetForDropOperation sets the value of the targetForDropOperation property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/isTargetForDropOperation
func (t_ TableRowView) SetTargetForDropOperation(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setTargetForDropOperation:"), value)
}
// Returns the number of columns represented by views in the table row view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/numberOfColumns
func (t_ TableRowView) NumberOfColumns() int {
	rv := objc.Send[int](t_.ID(), objc.RegisterName("numberOfColumns"))
	return rv
}
// Specifies the selection highlight style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/selectionHighlightStyle
func (t_ TableRowView) SelectionHighlightStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("selectionHighlightStyle"))
	return rv
}
// SetSelectionHighlightStyle sets the value of the selectionHighlightStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/selectionHighlightStyle
func (t_ TableRowView) SetSelectionHighlightStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setSelectionHighlightStyle:"), value)
}
