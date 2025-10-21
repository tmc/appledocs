// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SegmentedControl] class.
var (
	SegmentedControlClass     _SegmentedControlClass
	SegmentedControlClassOnce sync.Once
)

func getSegmentedControlClass() _SegmentedControlClass {
	SegmentedControlClassOnce.Do(func() {
		SegmentedControlClass = _SegmentedControlClass{objc.GetClass("NSSegmentedControl")}
	})
	return SegmentedControlClass
}

type _SegmentedControlClass struct {
	class objc.Class
}

// An interface definition for the [SegmentedControl] class.
type ISegmentedControl interface {
	IControl
	SelectSegmentWithTag(tag int) bool
}

// Display one or more buttons in a single horizontal group.
//
// The class uses an class to implement much of the control’s functionality. Most methods in are simply cover methods that call the corresponding method in . The methods of that do not have covers relate to accessing and setting values for tags and tooltips, programatically setting the key segment, and establishing the mode of the control. The features of a segmented control include the following: A segment can have an image, text (label), menu, tooltip, and tag. A segmented control can contain images or text, but not both. Either the control or individual segments can be enabled or disabled. Segmented controls have four tracking modes, described in . You use these modes with the property. Each segment can be either a fixed width or autosized to fit the contents. If a segment has text and is marked as autosizing, then the text may be truncated so that the control completely fits. If an image is too large to fit in a segment, it is clipped. If Full Keyboard Access is enabled in System Preferences > Keyboard, the keyboard may be used to move between and select segments.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl
type SegmentedControl struct {
	Control
}

// SegmentedControlFrom constructs a [SegmentedControl] from an unsafe.Pointer.
//
// Display one or more buttons in a single horizontal group.
func SegmentedControlFrom(ptr unsafe.Pointer) SegmentedControl {
	return SegmentedControl{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SegmentedControlClass) Alloc() SegmentedControl {
	rv := objc.Send[SegmentedControl](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SegmentedControlClass) New() SegmentedControl {
	rv := objc.Send[SegmentedControl](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SegmentedControl) Init() SegmentedControl {
	rv := objc.Send[SegmentedControl](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SegmentedControl) Autorelease() SegmentedControl {
	rv := objc.Send[SegmentedControl](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSegmentedControl creates a new SegmentedControl instance.
func NewSegmentedControl() SegmentedControl {
	return getSegmentedControlClass().New()
}


// Selects the segment with the specified tag.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectSegment(withTag:)
func (s_ SegmentedControl) SelectSegmentWithTag(tag int) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("selectSegmentWithTag:"), tag)
	return rv
}

// The color of the selected segment’s bezel, in appearances that support it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectedSegmentBezelColor
func (s_ SegmentedControl) SelectedSegmentBezelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("selectedSegmentBezelColor"))
	return rv
}


// SetSelectedSegmentBezelColor sets the value of the selectedSegmentBezelColor property.
// The color of the selected segment’s bezel, in appearances that support it.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectedSegmentBezelColor
func (s_ SegmentedControl) SetSelectedSegmentBezelColor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectedSegmentBezelColor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/activecompressionoptions
func (s_ SegmentedControl) ActiveCompressionOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("activeCompressionOptions"))
	return rv
}


// SetActiveCompressionOptions sets the value of the activeCompressionOptions property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/activecompressionoptions
func (s_ SegmentedControl) SetActiveCompressionOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setActiveCompressionOptions:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/bordershape
func (s_ SegmentedControl) BorderShape() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("borderShape"))
	return rv
}


// SetBorderShape sets the value of the borderShape property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/bordershape
func (s_ SegmentedControl) SetBorderShape(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBorderShape:"), value)
}

// When the tracking mode for the control is set to use a momentary accelerator, returns a value for the selected segment.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/doublevalueforselectedsegment
func (s_ SegmentedControl) DoubleValueForSelectedSegment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("doubleValueForSelectedSegment"))
	return rv
}


// SetDoubleValueForSelectedSegment sets the value of the doubleValueForSelectedSegment property.
// When the tracking mode for the control is set to use a momentary accelerator, returns a value for the selected segment.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/doublevalueforselectedsegment
func (s_ SegmentedControl) SetDoubleValueForSelectedSegment(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDoubleValueForSelectedSegment:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/indexofselecteditem
func (s_ SegmentedControl) IndexOfSelectedItem() int {
	rv := objc.Send[int](s_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}


// SetIndexOfSelectedItem sets the value of the indexOfSelectedItem property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/indexofselecteditem
func (s_ SegmentedControl) SetIndexOfSelectedItem(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIndexOfSelectedItem:"), value)
}

// A Boolean value that indicates whether spring loading is enabled for the control.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/isspringloaded
func (s_ SegmentedControl) IsSpringLoaded() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSpringLoaded"))
	return rv
}


// SetIsSpringLoaded sets the value of the isSpringLoaded property.
// A Boolean value that indicates whether spring loading is enabled for the control.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/isspringloaded
func (s_ SegmentedControl) SetIsSpringLoaded(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSpringLoaded:"), value)
}

// The number of segments in the control.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentcount
func (s_ SegmentedControl) SegmentCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("segmentCount"))
	return rv
}


// SetSegmentCount sets the value of the segmentCount property.
// The number of segments in the control.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentcount
func (s_ SegmentedControl) SetSegmentCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentdistribution
func (s_ SegmentedControl) SegmentDistribution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("segmentDistribution"))
	return rv
}


// SetSegmentDistribution sets the value of the segmentDistribution property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentdistribution
func (s_ SegmentedControl) SetSegmentDistribution(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentDistribution:"), value)
}

// The visual style used to display the control.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentstyle
func (s_ SegmentedControl) SegmentStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("segmentStyle"))
	return rv
}


// SetSegmentStyle sets the value of the segmentStyle property.
// The visual style used to display the control.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentstyle
func (s_ SegmentedControl) SetSegmentStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentStyle:"), value)
}

// The index of the selected segment of the control, or
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/selectedsegment
func (s_ SegmentedControl) SelectedSegment() int {
	rv := objc.Send[int](s_.ID, objc.Sel("selectedSegment"))
	return rv
}


// SetSelectedSegment sets the value of the selectedSegment property.
// The index of the selected segment of the control, or

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/selectedsegment
func (s_ SegmentedControl) SetSelectedSegment(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectedSegment:"), value)
}

// The type of tracking behavior the control exhibits.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (s_ SegmentedControl) TrackingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("trackingMode"))
	return rv
}


// SetTrackingMode sets the value of the trackingMode property.
// The type of tracking behavior the control exhibits.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (s_ SegmentedControl) SetTrackingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTrackingMode:"), value)
}



