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

// An interface definition for the [SegmentedCell] class.
type ISegmentedCell interface {
	IActionCell
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
// Alloc allocates a new instance without initialization.
func (sc _SegmentedCellClass) Alloc() SegmentedCell {
	rv := objc.Send[SegmentedCell](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SegmentedCellClass) New() SegmentedCell {
	rv := objc.Send[SegmentedCell](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SegmentedCell) Init() SegmentedCell {
	rv := objc.Send[SegmentedCell](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SegmentedCell) Autorelease() SegmentedCell {
	rv := objc.Send[SegmentedCell](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSegmentedCell creates a new SegmentedCell instance.
func NewSegmentedCell() SegmentedCell {
	return segmentedCellClass.New()
}




