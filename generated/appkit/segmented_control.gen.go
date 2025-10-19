// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SegmentedControl] class.
var (
	segmentedControlClass     _SegmentedControlClass
	segmentedControlClassOnce sync.Once
)

func getSegmentedControlClass() _SegmentedControlClass {
	segmentedControlClassOnce.Do(func() {
		segmentedControlClass = _SegmentedControlClass{objc.GetClass("NSSegmentedControl")}
	})
	return segmentedControlClass
}

type _SegmentedControlClass struct {
	class objc.Class
}

// An interface definition for the [SegmentedControl] class.
type ISegmentedControl interface {
	IControl
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


