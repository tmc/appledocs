// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableRowView] class.
var TableRowViewClass objc.Class

func init() {
	TableRowViewClass = objc.GetClass("NSTableRowView")
}

type TableRowView struct {
	objc.ID
}

func TableRowViewFrom(ptr unsafe.Pointer) TableRowView {
	return TableRowView{
		ID: objc.ID(ptr),
	}
}


// Draws the background of the row in the rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/drawBackground(in:)
func (t_ TableRowView) DrawBackgroundInRect(dirtyRect unsafe.Pointer) {
	sel := objc.RegisterName("drawBackgroundInRect:")
	t_.ID.Send(sel, dirtyRect)
}
// Draws the row’s dragging destination feedback when the entire row is a drop target. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/drawDraggingDestinationFeedback(in:)
func (t_ TableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect unsafe.Pointer) {
	sel := objc.RegisterName("drawDraggingDestinationFeedbackInRect:")
	t_.ID.Send(sel, dirtyRect)
}
// Draws the selected row. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/drawSelection(in:)
func (t_ TableRowView) DrawSelectionInRect(dirtyRect unsafe.Pointer) {
	sel := objc.RegisterName("drawSelectionInRect:")
	t_.ID.Send(sel, dirtyRect)
}
// Draws the horizontal separator between table rows. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/drawSeparator(in:)
func (t_ TableRowView) DrawSeparatorInRect(dirtyRect unsafe.Pointer) {
	sel := objc.RegisterName("drawSeparatorInRect:")
	t_.ID.Send(sel, dirtyRect)
}
// Provides access to the given view at a particular column. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableRowView/view(atColumn:)
func (t_ TableRowView) ViewAtColumn(column int) objc.ID {
	sel := objc.RegisterName("viewAtColumn:")
	ret := t_.ID.Send(sel, column)
	return ret
}

