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
	// properties:
	Column() IGridColumn
	ContentView() IView
	SetContentView(value IView)
	CustomPlacementConstraints() []LayoutConstraint
	SetCustomPlacementConstraints(value []LayoutConstraint)
	Row() IGridRow
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
	// methods:
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/emptyContentView
func (gc _GridCellClass) EmptyContentView() IView {
	rv := objc.Send[View](objc.ID(gc.class), objc.Sel("emptyContentView"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/column
func (g_ GridCell) Column() IGridColumn {
	rv := objc.Send[GridColumn](g_.ID, objc.Sel("column"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/contentView
func (g_ GridCell) ContentView() IView {
	rv := objc.Send[View](g_.ID, objc.Sel("contentView"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/contentView
func (g_ GridCell) SetContentView(value IView) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setContentView:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/customPlacementConstraints
func (g_ GridCell) CustomPlacementConstraints() []LayoutConstraint {
	rv := objc.Send[[]LayoutConstraint](g_.ID, objc.Sel("customPlacementConstraints"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/customPlacementConstraints
func (g_ GridCell) SetCustomPlacementConstraints(value []LayoutConstraint) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](g_.ID, objc.Sel("setCustomPlacementConstraints:"), nsArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/emptyContentView
func (g_ GridCell) EmptyContentView() IView {
	rv := objc.Send[View](g_.ID, objc.Sel("emptyContentView"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/row
func (g_ GridCell) Row() IGridRow {
	rv := objc.Send[GridRow](g_.ID, objc.Sel("row"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/rowAlignment
func (g_ GridCell) RowAlignment() GridRowAlignment {
	rv := objc.Send[GridRowAlignment](g_.ID, objc.Sel("rowAlignment"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/rowAlignment
func (g_ GridCell) SetRowAlignment(value GridRowAlignment) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRowAlignment:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/xPlacement
func (g_ GridCell) XPlacement() GridCellPlacement {
	rv := objc.Send[GridCellPlacement](g_.ID, objc.Sel("xPlacement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/xPlacement
func (g_ GridCell) SetXPlacement(value GridCellPlacement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setXPlacement:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/yPlacement
func (g_ GridCell) YPlacement() GridCellPlacement {
	rv := objc.Send[GridCellPlacement](g_.ID, objc.Sel("yPlacement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/yPlacement
func (g_ GridCell) SetYPlacement(value GridCellPlacement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setYPlacement:"), value)
}



