// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SegmentedCell] class.
var (
	SegmentedCellClass     _SegmentedCellClass
	SegmentedCellClassOnce sync.Once
)

func getSegmentedCellClass() _SegmentedCellClass {
	SegmentedCellClassOnce.Do(func() {
		SegmentedCellClass = _SegmentedCellClass{objc.GetClass("NSSegmentedCell")}
	})
	return SegmentedCellClass
}

type _SegmentedCellClass struct {
	class objc.Class
}

// An interface definition for the [SegmentedCell] class.
type ISegmentedCell interface {
	IActionCell
}

// An object implements the appearance and behavior of a horizontal button divided into multiple segments. This class is used in conjunction with the class to implement a segmented control.
//
// Use the methods of to customize the attributes of a segmented control. To customize the appearance of individual segments, you can also subclass and override the method.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getSegmentedCellClass().New()
}


// The number of segments in the segmented control.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/segmentcount
func (s_ SegmentedCell) SegmentCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("segmentCount"))
	return rv
}


// SetSegmentCount sets the value of the segmentCount property.
// The number of segments in the segmented control.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/segmentcount
func (s_ SegmentedCell) SetSegmentCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentCount:"), value)
}

// The visual style used to display the segmented control.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/segmentstyle
func (s_ SegmentedCell) SegmentStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("segmentStyle"))
	return rv
}


// SetSegmentStyle sets the value of the segmentStyle property.
// The visual style used to display the segmented control.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/segmentstyle
func (s_ SegmentedCell) SetSegmentStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentStyle:"), value)
}

// The index of the selected segment of the control, or
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/selectedsegment
func (s_ SegmentedCell) SelectedSegment() int {
	rv := objc.Send[int](s_.ID, objc.Sel("selectedSegment"))
	return rv
}


// SetSelectedSegment sets the value of the selectedSegment property.
// The index of the selected segment of the control, or

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/selectedsegment
func (s_ SegmentedCell) SetSelectedSegment(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectedSegment:"), value)
}

// The tracking mode used for the segments of the control.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/trackingmode
func (s_ SegmentedCell) TrackingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("trackingMode"))
	return rv
}


// SetTrackingMode sets the value of the trackingMode property.
// The tracking mode used for the segments of the control.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcell/trackingmode
func (s_ SegmentedCell) SetTrackingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTrackingMode:"), value)
}



