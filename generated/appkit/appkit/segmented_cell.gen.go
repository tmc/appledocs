// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SegmentedCell] class.
var SegmentedCellClass objc.Class

func init() {
	SegmentedCellClass = objc.GetClass("NSSegmentedCell")
}

type SegmentedCell struct {
	objc.ID
}

func SegmentedCellFrom(ptr unsafe.Pointer) SegmentedCell {
	return SegmentedCell{
		ID: objc.ID(ptr),
	}
}




