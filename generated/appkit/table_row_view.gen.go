// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TableRowView] class.
var (
	tableRowViewClass     _TableRowViewClass
	tableRowViewClassOnce sync.Once
)

func getTableRowViewClass() _TableRowViewClass {
	tableRowViewClassOnce.Do(func() {
		tableRowViewClass = _TableRowViewClass{objc.GetClass("NSTableRowView")}
	})
	return tableRowViewClass
}

type _TableRowViewClass struct {
	class objc.Class
}

// An interface definition for the [TableRowView] class.
type ITableRowView interface {
	IView
	DrawBackgroundInRect(dirtyRect unsafe.Pointer)
	DrawDraggingDestinationFeedbackInRect(dirtyRect unsafe.Pointer)
	DrawSelectionInRect(dirtyRect unsafe.Pointer)
	DrawSeparatorInRect(dirtyRect unsafe.Pointer)
	ViewAtColumn(column int) objc.ID
}

// The view shown for a row in a table view. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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


// Draws the background of the row in the rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawBackground(in:)
func (t_ TableRowView) DrawBackgroundInRect(dirtyRect unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawBackgroundInRect:"), dirtyRect)
}
// Draws the row’s dragging destination feedback when the entire row is a drop target. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawDraggingDestinationFeedback(in:)
func (t_ TableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawDraggingDestinationFeedbackInRect:"), dirtyRect)
}
// Draws the selected row. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawSelection(in:)
func (t_ TableRowView) DrawSelectionInRect(dirtyRect unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawSelectionInRect:"), dirtyRect)
}
// Draws the horizontal separator between table rows. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawSeparator(in:)
func (t_ TableRowView) DrawSeparatorInRect(dirtyRect unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawSeparatorInRect:"), dirtyRect)
}
// Provides access to the given view at a particular column. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableRowView/view(atColumn:)
func (t_ TableRowView) ViewAtColumn(column int) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("viewAtColumn:"), column)
	return rv
}


