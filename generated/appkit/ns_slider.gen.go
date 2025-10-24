// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSSlider */


/* debug [class_header]: Header for NSSlider */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Slider */
// An interface definition for the [Slider] class.
type ISlider interface {
	IControl
	
/* debug [class_interface_properties]: Properties for Slider */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Slider */
	// methods:
	AcceptsFirstMouse(event IEvent) bool
	ClosestTickMarkValueToValue(value float64) float64
	IndexOfTickMarkAtPoint(point vision.Point) int
	RectOfTickMarkAtIndex(index int) Rect /* not a class type */
	TickMarkValueAtIndex(index int) float64
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Slider */
// Alloc allocates a new instance without initialization.
func (sc _SliderClass) Alloc() Slider {
	rv := objc.Send[Slider](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Slider */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Slider */

// Creates a continuous horizontal slider whose values range from to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/init(target:action:)
func NewSliderWithTargetAction(target objc.IObject, action objc.SEL) Slider {
	rv := objc.Send[Slider](objc.ID(getSliderClass().class), objc.Sel("sliderWithTarget:action:"), target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewSliderWithTargetAction */


// Creates a continuous horizontal slider that represents values over the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/init(value:minValue:maxValue:target:action:)
func NewSliderWithValueMinValueMaxValueTargetAction(value float64, minValue float64, maxValue float64, target objc.IObject, action objc.SEL) Slider {
	rv := objc.Send[Slider](objc.ID(getSliderClass().class), objc.Sel("sliderWithValue:minValue:maxValue:target:action:"), value, minValue, maxValue, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewSliderWithValueMinValueMaxValueTargetAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Slider */

// Creates a continuous horizontal slider whose values range from to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/init(target:action:)
func (sc _SliderClass) SliderWithTargetAction(target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("sliderWithTarget:action:"), target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SliderWithTargetAction) */


// Creates a continuous horizontal slider that represents values over the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/init(value:minValue:maxValue:target:action:)
func (sc _SliderClass) SliderWithValueMinValueMaxValueTargetAction(value float64, minValue float64, maxValue float64, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("sliderWithValue:minValue:maxValue:target:action:"), value, minValue, maxValue, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SliderWithValueMinValueMaxValueTargetAction) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Slider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Slider */

// Returns a Boolean value indicating whether a mouse-down event both activates the window and starts dragging the slider’s knob.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/acceptsFirstMouse(for:)
func (s_ Slider) AcceptsFirstMouse(event IEvent) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("acceptsFirstMouse:"), event)
	return rv
}/* debug [instance_methods/method]: AcceptsFirstMouse */


// Returns the value of the tick mark closest to the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/closestTickMarkValue(toValue:)
func (s_ Slider) ClosestTickMarkValueToValue(value float64) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("closestTickMarkValueToValue:"), value)
	return rv
}/* debug [instance_methods/method]: ClosestTickMarkValueToValue */


// Returns the index of the tick mark closest to the location of the slider represented by the given point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/indexOfTickMark(at:)
func (s_ Slider) IndexOfTickMarkAtPoint(point vision.Point) int {
	rv := objc.Send[int](s_.ID, objc.Sel("indexOfTickMarkAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: IndexOfTickMarkAtPoint */


// Returns the bounding rectangle of the tick mark at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/rectOfTickMark(at:)
func (s_ Slider) RectOfTickMarkAtIndex(index int) Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: RectOfTickMarkAtIndex */


// Returns the slider’s value represented by the tick mark at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/tickMarkValue(at:)
func (s_ Slider) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: TickMarkValueAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Slider */

// A Boolean value that indicates whether the slider fixes its values to those values represented by its tick marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/allowsTickMarkValuesOnly
func (s_ Slider) AllowsTickMarkValuesOnly() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsTickMarkValuesOnly"))
	return rv
}/* debug [instance_properties/getter]: allowsTickMarkValuesOnly */


// A Boolean value that indicates whether the slider fixes its values to those values represented by its tick marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/allowsTickMarkValuesOnly
func (s_ Slider) SetAllowsTickMarkValuesOnly(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsTickMarkValuesOnly:"), value)
}/* debug [instance_properties/setter]: allowsTickMarkValuesOnly */


// The amount by which the slider changes its value when the user Option-drags the slider knob.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/altIncrementValue
func (s_ Slider) AltIncrementValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("altIncrementValue"))
	return rv
}/* debug [instance_properties/getter]: altIncrementValue */


// The amount by which the slider changes its value when the user Option-drags the slider knob.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/altIncrementValue
func (s_ Slider) SetAltIncrementValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAltIncrementValue:"), value)
}/* debug [instance_properties/setter]: altIncrementValue */


// An integer indicating the orientation (horizontal or vertical) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/isVertical
func (s_ Slider) Vertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("vertical"))
	return rv
}/* debug [instance_properties/getter]: vertical */


// An integer indicating the orientation (horizontal or vertical) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/isVertical
func (s_ Slider) SetVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVertical:"), value)
}/* debug [instance_properties/setter]: vertical */


// The knob’s thickness, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/knobThickness
func (s_ Slider) KnobThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("knobThickness"))
	return rv
}/* debug [instance_properties/getter]: knobThickness */


// The maximum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/maxValue
func (s_ Slider) MaxValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxValue"))
	return rv
}/* debug [instance_properties/getter]: maxValue */


// The maximum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/maxValue
func (s_ Slider) SetMaxValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}/* debug [instance_properties/setter]: maxValue */


// The minimum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/minValue
func (s_ Slider) MinValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minValue"))
	return rv
}/* debug [instance_properties/getter]: minValue */


// The minimum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/minValue
func (s_ Slider) SetMinValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinValue:"), value)
}/* debug [instance_properties/setter]: minValue */


// The value this slider will be filled from. This slider will be filled from its to its current value. If has not been explicitly set before, access to will return .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/neutralValue
func (s_ Slider) NeutralValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("neutralValue"))
	return rv
}/* debug [instance_properties/getter]: neutralValue */


// The value this slider will be filled from. This slider will be filled from its to its current value. If has not been explicitly set before, access to will return .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/neutralValue
func (s_ Slider) SetNeutralValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNeutralValue:"), value)
}/* debug [instance_properties/setter]: neutralValue */


// The number of tick marks associated with the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/numberOfTickMarks
func (s_ Slider) NumberOfTickMarks() int {
	rv := objc.Send[int](s_.ID, objc.Sel("numberOfTickMarks"))
	return rv
}/* debug [instance_properties/getter]: numberOfTickMarks */


// The number of tick marks associated with the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/numberOfTickMarks
func (s_ Slider) SetNumberOfTickMarks(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNumberOfTickMarks:"), value)
}/* debug [instance_properties/setter]: numberOfTickMarks */


// The type of the slider, such as vertical or circular.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/sliderType-swift.property
func (s_ Slider) SliderType() SliderType {
	rv := objc.Send[SliderType](s_.ID, objc.Sel("sliderType"))
	return rv
}/* debug [instance_properties/getter]: sliderType */


// The type of the slider, such as vertical or circular.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/sliderType-swift.property
func (s_ Slider) SetSliderType(value SliderType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSliderType:"), value)
}/* debug [instance_properties/setter]: sliderType */


// Determines where the slider’s tick marks are displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/tickMarkPosition-swift.property
func (s_ Slider) TickMarkPosition() TickMarkPosition {
	rv := objc.Send[TickMarkPosition](s_.ID, objc.Sel("tickMarkPosition"))
	return rv
}/* debug [instance_properties/getter]: tickMarkPosition */


// Determines where the slider’s tick marks are displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/tickMarkPosition-swift.property
func (s_ Slider) SetTickMarkPosition(value TickMarkPosition) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTickMarkPosition:"), value)
}/* debug [instance_properties/setter]: tickMarkPosition */


// The tint prominence of the slider. The automatic behavior for a regular slider tints its track fill, while a slider with tick marks is untinted. Setting the tint prominence will override this default behavior and choose an explicit track fill tint behavior. See for a list of possible values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/tintProminence
func (s_ Slider) TintProminence() TintProminence {
	rv := objc.Send[TintProminence](s_.ID, objc.Sel("tintProminence"))
	return rv
}/* debug [instance_properties/getter]: tintProminence */


// The tint prominence of the slider. The automatic behavior for a regular slider tints its track fill, while a slider with tick marks is untinted. Setting the tint prominence will override this default behavior and choose an explicit track fill tint behavior. See for a list of possible values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/tintProminence
func (s_ Slider) SetTintProminence(value TintProminence) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTintProminence:"), value)
}/* debug [instance_properties/setter]: tintProminence */


// The color of the filled portion of the slider track, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/trackFillColor
func (s_ Slider) TrackFillColor() IColor {
	rv := objc.Send[Color](s_.ID, objc.Sel("trackFillColor"))
	return rv
}/* debug [instance_properties/getter]: trackFillColor */


// The color of the filled portion of the slider track, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/trackFillColor
func (s_ Slider) SetTrackFillColor(value IColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTrackFillColor:"), value)
}/* debug [instance_properties/setter]: trackFillColor */


// An integer indicating the orientation (horizontal or vertical) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/isvertical
func (s_ Slider) IsVertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVertical"))
	return rv
}/* debug [instance_properties/getter]: isVertical */


// An integer indicating the orientation (horizontal or vertical) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/isvertical
func (s_ Slider) SetIsVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVertical:"), value)
}/* debug [instance_properties/setter]: isVertical */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSlider */


