// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [GridColumn] class.
type IGridColumn interface {
	objectivec.IObject
	LeadingPadding() float64
	SetLeadingPadding(value float64)
	TrailingPadding() float64
	SetTrailingPadding(value float64)
	XPlacement() unsafe.Pointer
	SetXPlacement(value unsafe.Pointer)
	GridView() NSGridView
	SetGridView(value IGridView)
	IsHidden() bool
	SetIsHidden(value bool)
	NumberOfCells() int
	SetNumberOfCells(value int)
	Width() float64
	SetWidth(value float64)
}

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

// Alloc allocates a new instance without initialization.
func (gc _GridColumnClass) Alloc() GridColumn {
	rv := objc.Send[GridColumn](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/leadingPadding

func (g_ GridColumn) LeadingPadding() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("leadingPadding"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/leadingPadding

func (g_ GridColumn) SetLeadingPadding(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeadingPadding:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/trailingPadding

func (g_ GridColumn) TrailingPadding() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("trailingPadding"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/trailingPadding

func (g_ GridColumn) SetTrailingPadding(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTrailingPadding:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/xPlacement

func (g_ GridColumn) XPlacement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("xPlacement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridColumn/xPlacement

func (g_ GridColumn) SetXPlacement(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setXPlacement:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/gridview

func (g_ GridColumn) GridView() NSGridView {
	rv := objc.Send[NSGridView](g_.ID, objc.Sel("gridView"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/gridview

func (g_ GridColumn) SetGridView(value IGridView) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGridView:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/ishidden

func (g_ GridColumn) IsHidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isHidden"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/ishidden

func (g_ GridColumn) SetIsHidden(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsHidden:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/numberofcells

func (g_ GridColumn) NumberOfCells() int {
	rv := objc.Send[int](g_.ID, objc.Sel("numberOfCells"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/numberofcells

func (g_ GridColumn) SetNumberOfCells(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setNumberOfCells:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/width

func (g_ GridColumn) Width() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("width"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridcolumn/width

func (g_ GridColumn) SetWidth(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWidth:"), value)
}



