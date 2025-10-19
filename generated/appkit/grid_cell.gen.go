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
	gridCellClass     _GridCellClass
	gridCellClassOnce sync.Once
)

func getGridCellClass() _GridCellClass {
	gridCellClassOnce.Do(func() {
		gridCellClass = _GridCellClass{objc.GetClass("NSGridCell")}
	})
	return gridCellClass
}

type _GridCellClass struct {
	class objc.Class
}

// An interface definition for the [GridCell] class.
type IGridCell interface {
	objectivec.IObject
}

// An individual content area within a grid view, typically at the intersection of a row and a column. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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




