// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [GridRow] class.
type IGridRow interface {
	objectivec.IObject
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
	// methods:
	CellAtIndex(index int) IGridCell
	MergeCellsInRange(range_ corefoundation.Range)
}

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

// Alloc allocates a new instance without initialization.
func (gc _GridRowClass) Alloc() GridRow {
	rv := objc.Send[GridRow](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/cell(at:)
func (g_ GridRow) CellAtIndex(index int) IGridCell {
	rv := objc.Send[GridCell](g_.ID, objc.Sel("cellAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/mergeCells(in:)
func (g_ GridRow) MergeCellsInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](g_.ID, objc.Sel("mergeCellsInRange:"), range_)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/bottomPadding
func (g_ GridRow) BottomPadding() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("bottomPadding"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/bottomPadding
func (g_ GridRow) SetBottomPadding(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBottomPadding:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/gridView
func (g_ GridRow) GridView() IGridView {
	rv := objc.Send[GridView](g_.ID, objc.Sel("gridView"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/height
func (g_ GridRow) Height() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("height"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/height
func (g_ GridRow) SetHeight(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setHeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/isHidden
func (g_ GridRow) Hidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hidden"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/isHidden
func (g_ GridRow) SetHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setHidden:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/numberOfCells
func (g_ GridRow) NumberOfCells() int {
	rv := objc.Send[int](g_.ID, objc.Sel("numberOfCells"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/rowAlignment
func (g_ GridRow) RowAlignment() GridRowAlignment {
	rv := objc.Send[GridRowAlignment](g_.ID, objc.Sel("rowAlignment"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/rowAlignment
func (g_ GridRow) SetRowAlignment(value GridRowAlignment) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRowAlignment:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/topPadding
func (g_ GridRow) TopPadding() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("topPadding"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/topPadding
func (g_ GridRow) SetTopPadding(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTopPadding:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/yPlacement
func (g_ GridRow) YPlacement() GridCellPlacement {
	rv := objc.Send[GridCellPlacement](g_.ID, objc.Sel("yPlacement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/yPlacement
func (g_ GridRow) SetYPlacement(value GridCellPlacement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setYPlacement:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/ishidden
func (g_ GridRow) IsHidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isHidden"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridrow/ishidden
func (g_ GridRow) SetIsHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsHidden:"), value)
}



