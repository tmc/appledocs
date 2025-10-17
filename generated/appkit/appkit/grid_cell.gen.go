// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GridCell] class.
var GridCellClass objc.Class

func init() {
	GridCellClass = objc.GetClass("NSGridCell")
}

type GridCell struct {
	objc.ID
}

func GridCellFrom(ptr unsafe.Pointer) GridCell {
	return GridCell{
		ID: objc.ID(ptr),
	}
}




