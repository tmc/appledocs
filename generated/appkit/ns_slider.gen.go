// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	AllowsTickMarkValuesOnly() bool
	SetAllowsTickMarkValuesOnly(value bool)
	AltIncrementValue() float64
	SetAltIncrementValue(value float64)
	Vertical() bool
	SetVertical(value bool)
	KnobThickness() float64
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	NeutralValue() float64
	SetNeutralValue(value float64)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	SliderType() SliderType
	SetSliderType(value SliderType)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
	TintProminence() TintProminence
	SetTintProminence(value TintProminence)
	TrackFillColor() IColor
	SetTrackFillColor(value IColor)
	IsVertical() bool
	SetIsVertical(value bool)
	// methods:
	AcceptsFirstMouse(event IEvent) bool
	ClosestTickMarkValueToValue(value float64) float64
	IndexOfTickMarkAtPoint(point objc.IObject /* cross-framework: Point */) int
	RectOfTickMarkAtIndex(index int) objc.IObject /* cross-framework: Rect */
	TickMarkValueAtIndex(index int) float64
}

// A display of a bar representing a continuous range of numerical values and a knob representing the currently selected value.
//
// A slider is a UI element that displays a range of values in the app. Sliders can be vertical or horizontal bars or circular dials. An indicator, or knob, notes the current setting. The user can move the knob in the slider’s bar—or rotate the knob in a circular slider—to change the setting. The class uses the class to implement its user interface.


// A display of a bar representing a continuous range of numerical values and a knob representing the currently selected value.
//
// [Full Topic]
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



// Creates a continuous horizontal slider whose values range from to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/init(target:action:)
func NewSliderWithTargetAction(target objc.IObject, action objc.SEL) Slider {
	rv := objc.Send[Slider](objc.ID(getSliderClass().class), objc.Sel("sliderWithTarget:action:"), target, action)
	return rv
}


// Creates a continuous horizontal slider that represents values over the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/init(value:minValue:maxValue:target:action:)
func NewSliderWithValueMinValueMaxValueTargetAction(value float64, minValue float64, maxValue float64, target objc.IObject, action objc.SEL) Slider {
	rv := objc.Send[Slider](objc.ID(getSliderClass().class), objc.Sel("sliderWithValue:minValue:maxValue:target:action:"), value, minValue, maxValue, target, action)
	return rv
}



// Creates a continuous horizontal slider whose values range from to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/init(target:action:)
func (sc _SliderClass) SliderWithTargetAction(target objc.IObject, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sliderWithTarget:action:"), target, action)
	return rv
}


// Creates a continuous horizontal slider that represents values over the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/init(value:minValue:maxValue:target:action:)
func (sc _SliderClass) SliderWithValueMinValueMaxValueTargetAction(value float64, minValue float64, maxValue float64, target objc.IObject, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sliderWithValue:minValue:maxValue:target:action:"), value, minValue, maxValue, target, action)
	return rv
}


// Returns a Boolean value indicating whether a mouse-down event both activates the window and starts dragging the slider’s knob.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/acceptsFirstMouse(for:)
func (s_ Slider) AcceptsFirstMouse(event IEvent) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("acceptsFirstMouse:"), event)
	return rv
}


// Returns the value of the tick mark closest to the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/closestTickMarkValue(toValue:)
func (s_ Slider) ClosestTickMarkValueToValue(value float64) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("closestTickMarkValueToValue:"), value)
	return rv
}


// Returns the index of the tick mark closest to the location of the slider represented by the given point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/indexOfTickMark(at:)
func (s_ Slider) IndexOfTickMarkAtPoint(point objc.IObject /* cross-framework: Point */) int {
	rv := objc.Send[int](s_.ID, objc.Sel("indexOfTickMarkAtPoint:"), point)
	return rv
}


// Returns the bounding rectangle of the tick mark at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/rectOfTickMark(at:)
func (s_ Slider) RectOfTickMarkAtIndex(index int) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](s_.ID, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return rv
}


// Returns the slider’s value represented by the tick mark at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/tickMarkValue(at:)
func (s_ Slider) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}


// A Boolean value that indicates whether the slider fixes its values to those values represented by its tick marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/allowsTickMarkValuesOnly
func (s_ Slider) AllowsTickMarkValuesOnly() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsTickMarkValuesOnly"))
	return rv
}


// A Boolean value that indicates whether the slider fixes its values to those values represented by its tick marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/allowsTickMarkValuesOnly
func (s_ Slider) SetAllowsTickMarkValuesOnly(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsTickMarkValuesOnly:"), value)
}


// The amount by which the slider changes its value when the user Option-drags the slider knob.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/altIncrementValue
func (s_ Slider) AltIncrementValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("altIncrementValue"))
	return rv
}


// The amount by which the slider changes its value when the user Option-drags the slider knob.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/altIncrementValue
func (s_ Slider) SetAltIncrementValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAltIncrementValue:"), value)
}


// An integer indicating the orientation (horizontal or vertical) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/isVertical
func (s_ Slider) Vertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("vertical"))
	return rv
}


// An integer indicating the orientation (horizontal or vertical) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/isVertical
func (s_ Slider) SetVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVertical:"), value)
}


// The knob’s thickness, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/knobThickness
func (s_ Slider) KnobThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("knobThickness"))
	return rv
}


// The maximum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/maxValue
func (s_ Slider) MaxValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxValue"))
	return rv
}


// The maximum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/maxValue
func (s_ Slider) SetMaxValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}


// The minimum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/minValue
func (s_ Slider) MinValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minValue"))
	return rv
}


// The minimum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/minValue
func (s_ Slider) SetMinValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinValue:"), value)
}


// The value this slider will be filled from. This slider will be filled from its to its current value. If has not been explicitly set before, access to will return .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/neutralValue
func (s_ Slider) NeutralValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("neutralValue"))
	return rv
}


// The value this slider will be filled from. This slider will be filled from its to its current value. If has not been explicitly set before, access to will return .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/neutralValue
func (s_ Slider) SetNeutralValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNeutralValue:"), value)
}


// The number of tick marks associated with the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/numberOfTickMarks
func (s_ Slider) NumberOfTickMarks() int {
	rv := objc.Send[int](s_.ID, objc.Sel("numberOfTickMarks"))
	return rv
}


// The number of tick marks associated with the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/numberOfTickMarks
func (s_ Slider) SetNumberOfTickMarks(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNumberOfTickMarks:"), value)
}


// The type of the slider, such as vertical or circular.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/sliderType-swift.property
func (s_ Slider) SliderType() SliderType {
	rv := objc.Send[SliderType](s_.ID, objc.Sel("sliderType"))
	return rv
}


// The type of the slider, such as vertical or circular.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/sliderType-swift.property
func (s_ Slider) SetSliderType(value SliderType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSliderType:"), value)
}


// Determines where the slider’s tick marks are displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/tickMarkPosition-swift.property
func (s_ Slider) TickMarkPosition() TickMarkPosition {
	rv := objc.Send[TickMarkPosition](s_.ID, objc.Sel("tickMarkPosition"))
	return rv
}


// Determines where the slider’s tick marks are displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/tickMarkPosition-swift.property
func (s_ Slider) SetTickMarkPosition(value TickMarkPosition) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTickMarkPosition:"), value)
}


// The tint prominence of the slider. The automatic behavior for a regular slider tints its track fill, while a slider with tick marks is untinted. Setting the tint prominence will override this default behavior and choose an explicit track fill tint behavior. See for a list of possible values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/tintProminence
func (s_ Slider) TintProminence() TintProminence {
	rv := objc.Send[TintProminence](s_.ID, objc.Sel("tintProminence"))
	return rv
}


// The tint prominence of the slider. The automatic behavior for a regular slider tints its track fill, while a slider with tick marks is untinted. Setting the tint prominence will override this default behavior and choose an explicit track fill tint behavior. See for a list of possible values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/tintProminence
func (s_ Slider) SetTintProminence(value TintProminence) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTintProminence:"), value)
}


// The color of the filled portion of the slider track, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/trackFillColor
func (s_ Slider) TrackFillColor() IColor {
	rv := objc.Send[Color](s_.ID, objc.Sel("trackFillColor"))
	return rv
}


// The color of the filled portion of the slider track, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/trackFillColor
func (s_ Slider) SetTrackFillColor(value IColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTrackFillColor:"), value)
}


// An integer indicating the orientation (horizontal or vertical) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/isvertical
func (s_ Slider) IsVertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVertical"))
	return rv
}


// An integer indicating the orientation (horizontal or vertical) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/isvertical
func (s_ Slider) SetIsVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVertical:"), value)
}


