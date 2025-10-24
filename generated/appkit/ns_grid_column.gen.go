// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSGridColumn */


/* debug [class_header]: Header for NSGridColumn */
// The class instance for the [GridColumn] class.
var (
	GridColumnClass     _GridColumnClass
	GridColumnClassOnce sync.Once
)

func getGridColumnClass() _GridColumnClass {
	GridColumnClassOnce.Do(func() {
		GridColumnClass = _GridColumnClass{objc.GetClass("NSGridColumn")}
	})
	return GridColumnClass
}

type _GridColumnClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GridColumn */
// An interface definition for the [GridColumn] class.
type IGridColumn interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GridColumn */
	// properties:
	GridView() IGridView
	Hidden() bool
	SetHidden(value bool)
	LeadingPadding() float64
	SetLeadingPadding(value float64)
	NumberOfCells() int
	TrailingPadding() float64
	SetTrailingPadding(value float64)
	Width() float64
	SetWidth(value float64)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
	IsHidden() bool
	SetIsHidden(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GridColumn */
	// methods:
	CellAtIndex(index int) IGridCell
	MergeCellsInRange(range_ corefoundation.Range)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GridColumn */
// Alloc allocates a new instance without initialization.
func (gc _GridColumnClass) Alloc() GridColumn {
	rv := objc.Send[GridColumn](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GridColumnClass) New() GridColumn {
	rv := objc.Send[GridColumn](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GridColumn) Init() GridColumn {
	rv := objc.Send[GridColumn](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GridColumn) Autorelease() GridColumn {
	rv := objc.Send[GridColumn](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGridColumn creates a new GridColumn instance.
func NewGridColumn() GridColumn {
	return getGridColumnClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GridColumn */
// A column within a grid view.


// A column within a grid view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn
type GridColumn struct {
	objectivec.Object
}

// GridColumnFrom constructs a [GridColumn] from an unsafe.Pointer.
//
// A column within a grid view.
func GridColumnFrom(ptr unsafe.Pointer) GridColumn {
	return GridColumn{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GridColumn *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GridColumn */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GridColumn */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GridColumn */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/cell(at:)
func (g_ GridColumn) CellAtIndex(index int) IGridCell {
	rv := objc.Send[GridCell](g_.ID, objc.Sel("cellAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: CellAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/mergeCells(in:)
func (g_ GridColumn) MergeCellsInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](g_.ID, objc.Sel("mergeCellsInRange:"), range_)
}/* debug [instance_methods/method]: MergeCellsInRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GridColumn */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/gridView
func (g_ GridColumn) GridView() IGridView {
	rv := objc.Send[GridView](g_.ID, objc.Sel("gridView"))
	return rv
}/* debug [instance_properties/getter]: gridView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/isHidden
func (g_ GridColumn) Hidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hidden"))
	return rv
}/* debug [instance_properties/getter]: hidden */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/isHidden
func (g_ GridColumn) SetHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setHidden:"), value)
}/* debug [instance_properties/setter]: hidden */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/leadingPadding
func (g_ GridColumn) LeadingPadding() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("leadingPadding"))
	return rv
}/* debug [instance_properties/getter]: leadingPadding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/leadingPadding
func (g_ GridColumn) SetLeadingPadding(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeadingPadding:"), value)
}/* debug [instance_properties/setter]: leadingPadding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/numberOfCells
func (g_ GridColumn) NumberOfCells() int {
	rv := objc.Send[int](g_.ID, objc.Sel("numberOfCells"))
	return rv
}/* debug [instance_properties/getter]: numberOfCells */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/trailingPadding
func (g_ GridColumn) TrailingPadding() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("trailingPadding"))
	return rv
}/* debug [instance_properties/getter]: trailingPadding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/trailingPadding
func (g_ GridColumn) SetTrailingPadding(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTrailingPadding:"), value)
}/* debug [instance_properties/setter]: trailingPadding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/width
func (g_ GridColumn) Width() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/width
func (g_ GridColumn) SetWidth(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/xPlacement
func (g_ GridColumn) XPlacement() GridCellPlacement {
	rv := objc.Send[GridCellPlacement](g_.ID, objc.Sel("xPlacement"))
	return rv
}/* debug [instance_properties/getter]: xPlacement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/xPlacement
func (g_ GridColumn) SetXPlacement(value GridCellPlacement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setXPlacement:"), value)
}/* debug [instance_properties/setter]: xPlacement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/ishidden
func (g_ GridColumn) IsHidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/ishidden
func (g_ GridColumn) SetIsHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSGridColumn */



