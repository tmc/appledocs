// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SegmentedCell] class.
var segmentedCellClass = _SegmentedCellClass{objc.GetClass("NSSegmentedCell")}

type _SegmentedCellClass struct {
	class objc.Class
}

// An object implements the appearance and behavior of a horizontal button divided into multiple segments. This class is used in conjunction with the class to implement a segmented control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell

type SegmentedCell struct {
	ActionCell
}

// SegmentedCellFrom constructs a [SegmentedCell] from an unsafe.Pointer.
//
// An object implements the appearance and behavior of a horizontal button divided into multiple segments. This class is used in conjunction with the class to implement a segmented control.
func SegmentedCellFrom(ptr unsafe.Pointer) SegmentedCell {
	return SegmentedCell{
		ActionCell: ActionCellFrom(ptr),
	}
}



