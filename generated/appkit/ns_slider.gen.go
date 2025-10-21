// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	SetTitleFont(fontObj IFont)
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

// Sets the font used to draw the slider’s title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/setTitleFont:
func (s_ Slider) SetTitleFont(fontObj IFont) {
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
func (s_ Slider) TrackFillColor() NSColor {
	rv := objc.Send[NSColor](s_.ID, objc.Sel("trackFillColor"))
	return rv
}


// SetTrackFillColor sets the value of the trackFillColor property.
// The color of the filled portion of the slider track, in appearances that support it.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/trackFillColor
func (s_ Slider) SetTrackFillColor(value IColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTrackFillColor:"), value)
}

// A Boolean value that indicates whether the slider fixes its values to those values represented by its tick marks.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/allowstickmarkvaluesonly
func (s_ Slider) AllowsTickMarkValuesOnly() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsTickMarkValuesOnly"))
	return rv
}


// SetAllowsTickMarkValuesOnly sets the value of the allowsTickMarkValuesOnly property.
// A Boolean value that indicates whether the slider fixes its values to those values represented by its tick marks.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/allowstickmarkvaluesonly
func (s_ Slider) SetAllowsTickMarkValuesOnly(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsTickMarkValuesOnly:"), value)
}

// The amount by which the slider changes its value when the user Option-drags the slider knob.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/altincrementvalue
func (s_ Slider) AltIncrementValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("altIncrementValue"))
	return rv
}


// SetAltIncrementValue sets the value of the altIncrementValue property.
// The amount by which the slider changes its value when the user Option-drags the slider knob.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/altincrementvalue
func (s_ Slider) SetAltIncrementValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAltIncrementValue:"), value)
}

// An integer indicating the orientation (horizontal or vertical) of the slider.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/isvertical
func (s_ Slider) IsVertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVertical"))
	return rv
}


// SetIsVertical sets the value of the isVertical property.
// An integer indicating the orientation (horizontal or vertical) of the slider.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/isvertical
func (s_ Slider) SetIsVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVertical:"), value)
}

// The minimum value the slider can send to its target.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/minvalue
func (s_ Slider) MinValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minValue"))
	return rv
}


// SetMinValue sets the value of the minValue property.
// The minimum value the slider can send to its target.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/minvalue
func (s_ Slider) SetMinValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinValue:"), value)
}

// The value this slider will be filled from. This slider will be filled from its
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/neutralvalue
func (s_ Slider) NeutralValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("neutralValue"))
	return rv
}


// SetNeutralValue sets the value of the neutralValue property.
// The value this slider will be filled from. This slider will be filled from its

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/neutralvalue
func (s_ Slider) SetNeutralValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNeutralValue:"), value)
}

// The number of tick marks associated with the slider.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/numberoftickmarks
func (s_ Slider) NumberOfTickMarks() int {
	rv := objc.Send[int](s_.ID, objc.Sel("numberOfTickMarks"))
	return rv
}


// SetNumberOfTickMarks sets the value of the numberOfTickMarks property.
// The number of tick marks associated with the slider.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/numberoftickmarks
func (s_ Slider) SetNumberOfTickMarks(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNumberOfTickMarks:"), value)
}

// The type of the slider, such as vertical or circular.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/slidertype-swift.property
func (s_ Slider) SliderType() SliderType {
	rv := objc.Send[SliderType](s_.ID, objc.Sel("sliderType"))
	return rv
}


// SetSliderType sets the value of the sliderType property.
// The type of the slider, such as vertical or circular.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/slidertype-swift.property
func (s_ Slider) SetSliderType(value SliderType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSliderType:"), value)
}

// Determines where the slider’s tick marks are displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/tickmarkposition-swift.property
func (s_ Slider) TickMarkPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("tickMarkPosition"))
	return rv
}


// SetTickMarkPosition sets the value of the tickMarkPosition property.
// Determines where the slider’s tick marks are displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/tickmarkposition-swift.property
func (s_ Slider) SetTickMarkPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTickMarkPosition:"), value)
}

// The tint prominence of the slider. The automatic behavior for a regular slider tints its track fill, while a slider with tick marks is untinted. Setting the tint prominence will override this default behavior and choose an explicit track fill tint behavior. See
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/tintprominence
func (s_ Slider) TintProminence() TintProminence {
	rv := objc.Send[TintProminence](s_.ID, objc.Sel("tintProminence"))
	return rv
}


// SetTintProminence sets the value of the tintProminence property.
// The tint prominence of the slider. The automatic behavior for a regular slider tints its track fill, while a slider with tick marks is untinted. Setting the tint prominence will override this default behavior and choose an explicit track fill tint behavior. See

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/tintprominence
func (s_ Slider) SetTintProminence(value ITintProminence) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTintProminence:"), value)
}



