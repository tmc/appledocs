// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GridCell] class.
var gridCellClass = _GridCellClass{objc.GetClass("NSGridCell")}

type _GridCellClass struct {
	class objc.Class
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



