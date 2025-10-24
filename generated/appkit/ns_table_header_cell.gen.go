// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

// The class instance for the [TableHeaderCell] class.
var (
	TableHeaderCellClass     _TableHeaderCellClass
	TableHeaderCellClassOnce sync.Once
)

func getTableHeaderCellClass() _TableHeaderCellClass {
	TableHeaderCellClassOnce.Do(func() {
		TableHeaderCellClass = _TableHeaderCellClass{objc.GetClass("NSTableHeaderCell")}
	})
	return TableHeaderCellClass
}

type _TableHeaderCellClass struct {
	class objc.Class
}

// An interface definition for the [TableHeaderCell] class.
type ITableHeaderCell interface {
	ITextFieldCell
	// properties:
	// methods:
	DrawSortIndicatorWithFrameInViewAscendingPriority(cellFrame objc.IObject /* cross-framework: Rect */, controlView IView, ascending bool, priority int)
	SortIndicatorRectForBounds(rect objc.IObject /* cross-framework: Rect */) objc.IObject /* cross-framework: Rect */
}

// An object that a table header view uses to draw the content of the column headers.
//
// Subclasses of the class can override the , , and methods to change the way headers appear. This specific subclass is responsible for drawing the sort indicators. See the class specification for information on overriding these methods. See the and for more information.


// An object that a table header view uses to draw the content of the column headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderCell
type TableHeaderCell struct {
	TextFieldCell
}

// TableHeaderCellFrom constructs a [TableHeaderCell] from an unsafe.Pointer.
//
// An object that a table header view uses to draw the content of the column headers.
func TableHeaderCellFrom(ptr unsafe.Pointer) TableHeaderCell {
	return TableHeaderCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TableHeaderCellClass) Alloc() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TableHeaderCellClass) New() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableHeaderCell) Init() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableHeaderCell) Autorelease() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableHeaderCell creates a new TableHeaderCell instance.
func NewTableHeaderCell() TableHeaderCell {
	return getTableHeaderCellClass().New()
}



// Draws a sorting indicator given a cell frame contained inside a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderCell/drawSortIndicator(withFrame:in:ascending:priority:)
func (t_ TableHeaderCell) DrawSortIndicatorWithFrameInViewAscendingPriority(cellFrame objc.IObject /* cross-framework: Rect */, controlView IView, ascending bool, priority int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawSortIndicatorWithFrame:inView:ascending:priority:"), cellFrame, controlView, ascending, priority)
}


// Returns the location to display the sorting indicator given .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderCell/sortIndicatorRect(forBounds:)
func (t_ TableHeaderCell) SortIndicatorRectForBounds(rect objc.IObject /* cross-framework: Rect */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](t_.ID, objc.Sel("sortIndicatorRectForBounds:"), rect)
	return rv
}



