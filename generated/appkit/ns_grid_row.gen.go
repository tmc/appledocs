// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSGridRow */


/* debug [class_header]: Header for NSGridRow */
// The class instance for the [GridRow] class.
var (
	GridRowClass     _GridRowClass
	GridRowClassOnce sync.Once
)

func getGridRowClass() _GridRowClass {
	GridRowClassOnce.Do(func() {
		GridRowClass = _GridRowClass{objc.GetClass("NSGridRow")}
	})
	return GridRowClass
}

type _GridRowClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GridRow */
// An interface definition for the [GridRow] class.
type IGridRow interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GridRow */
	// properties:
	BottomPadding() float64
	SetBottomPadding(value float64)
	GridView() IGridView
	Height() float64
	SetHeight(value float64)
	Hidden() bool
	SetHidden(value bool)
	NumberOfCells() int
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	TopPadding() float64
	SetTopPadding(value float64)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
	IsHidden() bool
	SetIsHidden(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GridRow */
	// methods:
	CellAtIndex(index int) IGridCell
	MergeCellsInRange(range_ corefoundation.Range)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GridRow */
// Alloc allocates a new instance without initialization.
func (gc _GridRowClass) Alloc() GridRow {
	rv := objc.Send[GridRow](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GridRowClass) New() GridRow {
	rv := objc.Send[GridRow](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GridRow) Init() GridRow {
	rv := objc.Send[GridRow](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GridRow) Autorelease() GridRow {
	rv := objc.Send[GridRow](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGridRow creates a new GridRow instance.
func NewGridRow() GridRow {
	return getGridRowClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GridRow */
// A row within a grid view.


// A row within a grid view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow
type GridRow struct {
	objectivec.Object
}

// GridRowFrom constructs a [GridRow] from an unsafe.Pointer.
//
// A row within a grid view.
func GridRowFrom(ptr unsafe.Pointer) GridRow {
	return GridRow{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GridRow *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GridRow */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GridRow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GridRow */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/cell(at:)
func (g_ GridRow) CellAtIndex(index int) IGridCell {
	rv := objc.Send[GridCell](g_.ID, objc.Sel("cellAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: CellAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/mergeCells(in:)
func (g_ GridRow) MergeCellsInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](g_.ID, objc.Sel("mergeCellsInRange:"), range_)
}/* debug [instance_methods/method]: MergeCellsInRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GridRow */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/bottomPadding
func (g_ GridRow) BottomPadding() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("bottomPadding"))
	return rv
}/* debug [instance_properties/getter]: bottomPadding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/bottomPadding
func (g_ GridRow) SetBottomPadding(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBottomPadding:"), value)
}/* debug [instance_properties/setter]: bottomPadding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/gridView
func (g_ GridRow) GridView() IGridView {
	rv := objc.Send[GridView](g_.ID, objc.Sel("gridView"))
	return rv
}/* debug [instance_properties/getter]: gridView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/height
func (g_ GridRow) Height() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/height
func (g_ GridRow) SetHeight(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/isHidden
func (g_ GridRow) Hidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hidden"))
	return rv
}/* debug [instance_properties/getter]: hidden */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/isHidden
func (g_ GridRow) SetHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setHidden:"), value)
}/* debug [instance_properties/setter]: hidden */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/numberOfCells
func (g_ GridRow) NumberOfCells() int {
	rv := objc.Send[int](g_.ID, objc.Sel("numberOfCells"))
	return rv
}/* debug [instance_properties/getter]: numberOfCells */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/rowAlignment
func (g_ GridRow) RowAlignment() GridRowAlignment {
	rv := objc.Send[GridRowAlignment](g_.ID, objc.Sel("rowAlignment"))
	return rv
}/* debug [instance_properties/getter]: rowAlignment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/rowAlignment
func (g_ GridRow) SetRowAlignment(value GridRowAlignment) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRowAlignment:"), value)
}/* debug [instance_properties/setter]: rowAlignment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/topPadding
func (g_ GridRow) TopPadding() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("topPadding"))
	return rv
}/* debug [instance_properties/getter]: topPadding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/topPadding
func (g_ GridRow) SetTopPadding(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTopPadding:"), value)
}/* debug [instance_properties/setter]: topPadding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/yPlacement
func (g_ GridRow) YPlacement() GridCellPlacement {
	rv := objc.Send[GridCellPlacement](g_.ID, objc.Sel("yPlacement"))
	return rv
}/* debug [instance_properties/getter]: yPlacement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/yPlacement
func (g_ GridRow) SetYPlacement(value GridCellPlacement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setYPlacement:"), value)
}/* debug [instance_properties/setter]: yPlacement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/ishidden
func (g_ GridRow) IsHidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/ishidden
func (g_ GridRow) SetIsHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSGridRow */



