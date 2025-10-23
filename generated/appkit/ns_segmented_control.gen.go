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
	// properties:
	SelectedSegmentBezelColor() IColor
	SetSelectedSegmentBezelColor(value IColor)
	ActiveCompressionOptions() objc.IObject /* cross-framework: UserInterfaceCompressionOptions */
	SetActiveCompressionOptions(value objc.IObject /* cross-framework: UserInterfaceCompressionOptions */)
	BorderShape() unsafe.Pointer
	SetBorderShape(value unsafe.Pointer)
	DoubleValueForSelectedSegment() float64 /* primitive/slice/pointer. */
	SetDoubleValueForSelectedSegment(value float64 /* primitive/slice/pointer. */)
	IndexOfSelectedItem() int /* primitive/slice/pointer. */
	SetIndexOfSelectedItem(value int /* primitive/slice/pointer. */)
	IsSpringLoaded() bool /* primitive/slice/pointer. */
	SetIsSpringLoaded(value bool /* primitive/slice/pointer. */)
	SegmentCount() int /* primitive/slice/pointer. */
	SetSegmentCount(value int /* primitive/slice/pointer. */)
	SegmentDistribution() unsafe.Pointer
	SetSegmentDistribution(value unsafe.Pointer)
	SegmentStyle() unsafe.Pointer
	SetSegmentStyle(value unsafe.Pointer)
	SelectedSegment() int /* primitive/slice/pointer. */
	SetSelectedSegment(value int /* primitive/slice/pointer. */)
	TrackingMode() unsafe.Pointer
	SetTrackingMode(value unsafe.Pointer)
	// methods:
}

// Display one or more buttons in a single horizontal group.
//
// The class uses an class to implement much of the control’s functionality. Most methods in are simply cover methods that call the corresponding method in . The methods of that do not have covers relate to accessing and setting values for tags and tooltips, programatically setting the key segment, and establishing the mode of the control. The features of a segmented control include the following: A segment can have an image, text (label), menu, tooltip, and tag. A segmented control can contain images or text, but not both. Either the control or individual segments can be enabled or disabled. Segmented controls have four tracking modes, described in . You use these modes with the property. Each segment can be either a fixed width or autosized to fit the contents. If a segment has text and is marked as autosizing, then the text may be truncated so that the control completely fits. If an image is too large to fit in a segment, it is clipped. If Full Keyboard Access is enabled in System Preferences > Keyboard, the keyboard may be used to move between and select segments.


// Display one or more buttons in a single horizontal group.
//
// [Full Topic]
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



// The color of the selected segment’s bezel, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectedSegmentBezelColor
func (s_ SegmentedControl) SelectedSegmentBezelColor() IColor {
	rv := objc.Send[Color](s_.ID, objc.Sel("selectedSegmentBezelColor"))
	return rv
}


// The color of the selected segment’s bezel, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectedSegmentBezelColor
func (s_ SegmentedControl) SetSelectedSegmentBezelColor(value IColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectedSegmentBezelColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/activecompressionoptions
func (s_ SegmentedControl) ActiveCompressionOptions() objc.IObject /* cross-framework: UserInterfaceCompressionOptions */ {
	rv := objc.Send[UserInterfaceCompressionOptions](s_.ID, objc.Sel("activeCompressionOptions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/activecompressionoptions
func (s_ SegmentedControl) SetActiveCompressionOptions(value objc.IObject /* cross-framework: UserInterfaceCompressionOptions */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setActiveCompressionOptions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/bordershape
func (s_ SegmentedControl) BorderShape() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("borderShape"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/bordershape
func (s_ SegmentedControl) SetBorderShape(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBorderShape:"), value)
}


// When the tracking mode for the control is set to use a momentary accelerator, returns a value for the selected segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/doublevalueforselectedsegment
func (s_ SegmentedControl) DoubleValueForSelectedSegment() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("doubleValueForSelectedSegment"))
	return rv
}


// When the tracking mode for the control is set to use a momentary accelerator, returns a value for the selected segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/doublevalueforselectedsegment
func (s_ SegmentedControl) SetDoubleValueForSelectedSegment(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDoubleValueForSelectedSegment:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/indexofselecteditem
func (s_ SegmentedControl) IndexOfSelectedItem() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](s_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/indexofselecteditem
func (s_ SegmentedControl) SetIndexOfSelectedItem(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIndexOfSelectedItem:"), value)
}


// A Boolean value that indicates whether spring loading is enabled for the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/isspringloaded
func (s_ SegmentedControl) IsSpringLoaded() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSpringLoaded"))
	return rv
}


// A Boolean value that indicates whether spring loading is enabled for the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/isspringloaded
func (s_ SegmentedControl) SetIsSpringLoaded(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSpringLoaded:"), value)
}


// The number of segments in the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentcount
func (s_ SegmentedControl) SegmentCount() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](s_.ID, objc.Sel("segmentCount"))
	return rv
}


// The number of segments in the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentcount
func (s_ SegmentedControl) SetSegmentCount(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentdistribution
func (s_ SegmentedControl) SegmentDistribution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("segmentDistribution"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentdistribution
func (s_ SegmentedControl) SetSegmentDistribution(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentDistribution:"), value)
}


// The visual style used to display the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentstyle
func (s_ SegmentedControl) SegmentStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("segmentStyle"))
	return rv
}


// The visual style used to display the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/segmentstyle
func (s_ SegmentedControl) SetSegmentStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentStyle:"), value)
}


// The index of the selected segment of the control, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/selectedsegment
func (s_ SegmentedControl) SelectedSegment() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](s_.ID, objc.Sel("selectedSegment"))
	return rv
}


// The index of the selected segment of the control, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/selectedsegment
func (s_ SegmentedControl) SetSelectedSegment(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectedSegment:"), value)
}


// The type of tracking behavior the control exhibits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (s_ SegmentedControl) TrackingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("trackingMode"))
	return rv
}


// The type of tracking behavior the control exhibits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (s_ SegmentedControl) SetTrackingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTrackingMode:"), value)
}



