// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Slider] class.
var (
	SliderClass     _SliderClass
	SliderClassOnce sync.Once
)

func getSliderClass() _SliderClass {
	SliderClassOnce.Do(func() {
		SliderClass = _SliderClass{objc.GetClass("NSSlider")}
	})
	return SliderClass
}

type _SliderClass struct {
	class objc.Class
}

// An interface definition for the [Slider] class.
type ISlider interface {
	IControl
	IndexOfTickMarkAtPoint(point coregraphics.CGPoint) int
	SetKnobThickness(thickness float64)
	SetTitleFont(fontObj unsafe.Pointer)
}

// A display of a bar representing a continuous range of numerical values and a knob representing the currently selected value.
//
// A slider is a UI element that displays a range of values in the app. Sliders can be vertical or horizontal bars or circular dials. An indicator, or knob, notes the current setting. The user can move the knob in the slider’s bar—or rotate the knob in a circular slider—to change the setting. The class uses the class to implement its user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider
type Slider struct {
	Control
}

// SliderFrom constructs a [Slider] from an unsafe.Pointer.
//
// A display of a bar representing a continuous range of numerical values and a knob representing the currently selected value.
func SliderFrom(ptr unsafe.Pointer) Slider {
	return Slider{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SliderClass) Alloc() Slider {
	rv := objc.Send[Slider](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SliderClass) New() Slider {
	rv := objc.Send[Slider](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Slider) Init() Slider {
	rv := objc.Send[Slider](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Slider) Autorelease() Slider {
	rv := objc.Send[Slider](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSlider creates a new Slider instance.
func NewSlider() Slider {
	return getSliderClass().New()
}

// Returns the index of the tick mark closest to the location of the slider represented by the given point.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/indexOfTickMark(at:)
func (s_ Slider) IndexOfTickMarkAtPoint(point coregraphics.CGPoint) int {
	rv := objc.Send[int](s_.ID, objc.Sel("indexOfTickMarkAtPoint:"), point)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/setKnobThickness:
func (s_ Slider) SetKnobThickness(thickness float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKnobThickness:"), thickness)
}

// Sets the font used to draw the slider’s title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/setTitleFont:
func (s_ Slider) SetTitleFont(fontObj unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitleFont:"), fontObj)
}

// An integer indicating the orientation (horizontal or vertical) of the slider.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/isVertical
func (s_ Slider) Vertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("vertical"))
	return rv
}

// SetVertical sets the value of the vertical property.
// An integer indicating the orientation (horizontal or vertical) of the slider.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/isVertical
func (s_ Slider) SetVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVertical:"), value)
}

// The knob’s thickness, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/knobThickness
func (s_ Slider) KnobThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("knobThickness"))
	return rv
}

// The maximum value the slider can send to its target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/maxValue
func (s_ Slider) MaxValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxValue"))
	return rv
}

// SetMaxValue sets the value of the maxValue property.
// The maximum value the slider can send to its target.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/maxValue
func (s_ Slider) SetMaxValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}

// The color of the filled portion of the slider track, in appearances that support it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/trackFillColor
func (s_ Slider) TrackFillColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("trackFillColor"))
	return rv
}

// SetTrackFillColor sets the value of the trackFillColor property.
// The color of the filled portion of the slider track, in appearances that support it.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/trackFillColor
func (s_ Slider) SetTrackFillColor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTrackFillColor:"), value)
}
