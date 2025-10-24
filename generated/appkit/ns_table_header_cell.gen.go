// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSTableHeaderCell */


/* debug [class_header]: Header for NSTableHeaderCell */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TableHeaderCell */
// An interface definition for the [TableHeaderCell] class.
type ITableHeaderCell interface {
	ITextFieldCell
	
/* debug [class_interface_properties]: Properties for TableHeaderCell */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TableHeaderCell */
	// methods:
	DrawSortIndicatorWithFrameInViewAscendingPriority(cellFrame Rect /* not a class type */, controlView IView, ascending bool, priority int)
	SortIndicatorRectForBounds(rect Rect /* not a class type */) Rect /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TableHeaderCell */
// Alloc allocates a new instance without initialization.
func (tc _TableHeaderCellClass) Alloc() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TableHeaderCell */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TableHeaderCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TableHeaderCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TableHeaderCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TableHeaderCell */

// Draws a sorting indicator given a cell frame contained inside a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderCell/drawSortIndicator(withFrame:in:ascending:priority:)
func (t_ TableHeaderCell) DrawSortIndicatorWithFrameInViewAscendingPriority(cellFrame Rect /* not a class type */, controlView IView, ascending bool, priority int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawSortIndicatorWithFrame:inView:ascending:priority:"), cellFrame, controlView, ascending, priority)
}/* debug [instance_methods/method]: DrawSortIndicatorWithFrameInViewAscendingPriority */


// Returns the location to display the sorting indicator given .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderCell/sortIndicatorRect(forBounds:)
func (t_ TableHeaderCell) SortIndicatorRectForBounds(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("sortIndicatorRectForBounds:"), rect)
	return rv
}/* debug [instance_methods/method]: SortIndicatorRectForBounds */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TableHeaderCell */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTableHeaderCell */



