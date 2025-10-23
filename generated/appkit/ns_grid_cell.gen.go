// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GridCell] class.
var (
	GridCellClass     _GridCellClass
	GridCellClassOnce sync.Once
)

func getGridCellClass() _GridCellClass {
	GridCellClassOnce.Do(func() {
		GridCellClass = _GridCellClass{objc.GetClass("NSGridCell")}
	})
	return GridCellClass
}

type _GridCellClass struct {
	class objc.Class
}

// An interface definition for the [GridCell] class.
type IGridCell interface {
	objectivec.IObject
	Column() NSGridColumn
	SetColumn(value IGridColumn)
	ContentView() NSView
	SetContentView(value IView)
	CustomPlacementConstraints() NSLayoutConstraint
	SetCustomPlacementConstraints(value ILayoutConstraint)
	Row() unsafe.Pointer
	SetRow(value unsafe.Pointer)
	RowAlignment() unsafe.Pointer
	SetRowAlignment(value unsafe.Pointer)
	XPlacement() unsafe.Pointer
	SetXPlacement(value unsafe.Pointer)
	YPlacement() unsafe.Pointer
	SetYPlacement(value unsafe.Pointer)
}

// An individual content area within a grid view, typically at the intersection of a row and a column.
//
// Use a grid cell to specify the content view to display and to position the content view within the cell’s area.


// An individual content area within a grid view, typically at the intersection of a row and a column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell
type GridCell struct {
	objectivec.Object
}

// GridCellFrom constructs a [GridCell] from an unsafe.Pointer.
//
// An individual content area within a grid view, typically at the intersection of a row and a column.
func GridCellFrom(ptr unsafe.Pointer) GridCell {
	return GridCell{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GridCellClass) Alloc() GridCell {
	rv := objc.Send[GridCell](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GridCellClass) New() GridCell {
	rv := objc.Send[GridCell](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GridCell) Init() GridCell {
	rv := objc.Send[GridCell](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GridCell) Autorelease() GridCell {
	rv := objc.Send[GridCell](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGridCell creates a new GridCell instance.
func NewGridCell() GridCell {
	return getGridCellClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/column
func (g_ GridCell) Column() NSGridColumn {
	rv := objc.Send[NSGridColumn](g_.ID, objc.Sel("column"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/column
func (g_ GridCell) SetColumn(value IGridColumn) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColumn:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/contentview
func (g_ GridCell) ContentView() NSView {
	rv := objc.Send[NSView](g_.ID, objc.Sel("contentView"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/contentview
func (g_ GridCell) SetContentView(value IView) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setContentView:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/customplacementconstraints
func (g_ GridCell) CustomPlacementConstraints() NSLayoutConstraint {
	rv := objc.Send[NSLayoutConstraint](g_.ID, objc.Sel("customPlacementConstraints"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/customplacementconstraints
func (g_ GridCell) SetCustomPlacementConstraints(value ILayoutConstraint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCustomPlacementConstraints:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/row
func (g_ GridCell) Row() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("row"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/row
func (g_ GridCell) SetRow(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRow:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/rowalignment
func (g_ GridCell) RowAlignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("rowAlignment"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/rowalignment
func (g_ GridCell) SetRowAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRowAlignment:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/xplacement
func (g_ GridCell) XPlacement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("xPlacement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/xplacement
func (g_ GridCell) SetXPlacement(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setXPlacement:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/yplacement
func (g_ GridCell) YPlacement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("yPlacement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcell/yplacement
func (g_ GridCell) SetYPlacement(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setYPlacement:"), value)
}



