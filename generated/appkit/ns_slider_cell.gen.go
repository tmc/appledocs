// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSSliderCell */


/* debug [class_header]: Header for NSSliderCell */
// The class instance for the [SliderCell] class.
var (
	SliderCellClass     _SliderCellClass
	SliderCellClassOnce sync.Once
)

func getSliderCellClass() _SliderCellClass {
	SliderCellClassOnce.Do(func() {
		SliderCellClass = _SliderCellClass{objc.GetClass("NSSliderCell")}
	})
	return SliderCellClass
}

type _SliderCellClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SliderCell */
// An interface definition for the [SliderCell] class.
type ISliderCell interface {
	IActionCell
	
/* debug [class_interface_properties]: Properties for SliderCell */
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
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	SliderType() SliderType
	SetSliderType(value SliderType)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
	TrackRect() Rect /* not a class type */
	IsVertical() bool
	SetIsVertical(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SliderCell */
	// methods:
	BarRectFlipped(flipped bool) Rect /* not a class type */
	ClosestTickMarkValueToValue(value float64) float64
	DrawBarInsideFlipped(rect Rect /* not a class type */, flipped bool)
	DrawKnob()
	DrawKnobWithKnobRect(knobRect Rect /* not a class type */)
	DrawTickMarks()
	IndexOfTickMarkAtPoint(point vision.Point) int
	KnobRectFlipped(flipped bool) Rect /* not a class type */
	RectOfTickMarkAtIndex(index int) Rect /* not a class type */
	TickMarkValueAtIndex(index int) float64
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SliderCell */
// Alloc allocates a new instance without initialization.
func (sc _SliderCellClass) Alloc() SliderCell {
	rv := objc.Send[SliderCell](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SliderCellClass) New() SliderCell {
	rv := objc.Send[SliderCell](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SliderCell) Init() SliderCell {
	rv := objc.Send[SliderCell](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SliderCell) Autorelease() SliderCell {
	rv := objc.Send[SliderCell](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSliderCell creates a new SliderCell instance.
func NewSliderCell() SliderCell {
	return getSliderCellClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SliderCell */
// The appearance and behavior of an object.
//
// You can customize an to a certain degree, using its properties. If this doesn’t give you sufficient flexibility, you can create a subclass. In that subclass, you can override any of the following methods: , , , and .


// The appearance and behavior of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell
type SliderCell struct {
	ActionCell
}

// SliderCellFrom constructs a [SliderCell] from an unsafe.Pointer.
//
// The appearance and behavior of an object.
func SliderCellFrom(ptr unsafe.Pointer) SliderCell {
	return SliderCell{
		ActionCell: ActionCellFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SliderCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SliderCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SliderCell */

// Returns a Boolean value indicating whether the continues to track the pointer until the next mouse up.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/prefersTrackingUntilMouseUp
func (sc _SliderCellClass) PrefersTrackingUntilMouseUp() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("prefersTrackingUntilMouseUp"))
	return rv
}/* debug [class_properties_class/property]: prefersTrackingUntilMouseUp */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SliderCell */

// Returns the rectangle in which the bar is drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/barRect(flipped:)
func (s_ SliderCell) BarRectFlipped(flipped bool) Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("barRectFlipped:"), flipped)
	return rv
}/* debug [instance_methods/method]: BarRectFlipped */


// Returns the value of the tick mark closest to the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/closestTickMarkValue(toValue:)
func (s_ SliderCell) ClosestTickMarkValueToValue(value float64) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("closestTickMarkValueToValue:"), value)
	return rv
}/* debug [instance_methods/method]: ClosestTickMarkValueToValue */


// Draws the slider’s bar—but not its bezel or knob—inside the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/drawBar(inside:flipped:)
func (s_ SliderCell) DrawBarInsideFlipped(rect Rect /* not a class type */, flipped bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawBarInside:flipped:"), rect, flipped)
}/* debug [instance_methods/method]: DrawBarInsideFlipped */


// Calculates the rectangle in which the knob should be drawn, then calls to actually draw the knob.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/drawKnob()
func (s_ SliderCell) DrawKnob() {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawKnob"))
}/* debug [instance_methods/method]: DrawKnob */


// Draws the slider knob in the given rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/drawKnob(_:)
func (s_ SliderCell) DrawKnobWithKnobRect(knobRect Rect /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawKnob:"), knobRect)
}/* debug [instance_methods/method]: DrawKnobWithKnobRect */


// Draws the slider’s tick marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/drawTickMarks()
func (s_ SliderCell) DrawTickMarks() {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawTickMarks"))
}/* debug [instance_methods/method]: DrawTickMarks */


// Returns the index of the tick mark closest to the location of the slider represented by the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/indexOfTickMark(at:)
func (s_ SliderCell) IndexOfTickMarkAtPoint(point vision.Point) int {
	rv := objc.Send[int](s_.ID, objc.Sel("indexOfTickMarkAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: IndexOfTickMarkAtPoint */


// Returns the rectangle in which the slider knob is drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/knobRect(flipped:)
func (s_ SliderCell) KnobRectFlipped(flipped bool) Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("knobRectFlipped:"), flipped)
	return rv
}/* debug [instance_methods/method]: KnobRectFlipped */


// Returns the bounding rectangle of the tick mark at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/rectOfTickMark(at:)
func (s_ SliderCell) RectOfTickMarkAtIndex(index int) Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: RectOfTickMarkAtIndex */


// Returns the receiver’s value represented by the tick mark at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/tickMarkValue(at:)
func (s_ SliderCell) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: TickMarkValueAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SliderCell */

// A Boolean value indicating whether the receiver fixes its values to those values represented by its tick marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/allowsTickMarkValuesOnly
func (s_ SliderCell) AllowsTickMarkValuesOnly() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsTickMarkValuesOnly"))
	return rv
}/* debug [instance_properties/getter]: allowsTickMarkValuesOnly */


// A Boolean value indicating whether the receiver fixes its values to those values represented by its tick marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/allowsTickMarkValuesOnly
func (s_ SliderCell) SetAllowsTickMarkValuesOnly(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsTickMarkValuesOnly:"), value)
}/* debug [instance_properties/setter]: allowsTickMarkValuesOnly */


// The amount by which the slider changes its value when the user Option-drags the knob.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/altIncrementValue
func (s_ SliderCell) AltIncrementValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("altIncrementValue"))
	return rv
}/* debug [instance_properties/getter]: altIncrementValue */


// The amount by which the slider changes its value when the user Option-drags the knob.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/altIncrementValue
func (s_ SliderCell) SetAltIncrementValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAltIncrementValue:"), value)
}/* debug [instance_properties/setter]: altIncrementValue */


// An integer indicating the orientation (vertical or horizontal) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/isVertical
func (s_ SliderCell) Vertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("vertical"))
	return rv
}/* debug [instance_properties/getter]: vertical */


// An integer indicating the orientation (vertical or horizontal) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/isVertical
func (s_ SliderCell) SetVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVertical:"), value)
}/* debug [instance_properties/setter]: vertical */


// The thickness of the slider knob, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/knobThickness
func (s_ SliderCell) KnobThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("knobThickness"))
	return rv
}/* debug [instance_properties/getter]: knobThickness */


// The maximum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/maxValue
func (s_ SliderCell) MaxValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxValue"))
	return rv
}/* debug [instance_properties/getter]: maxValue */


// The maximum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/maxValue
func (s_ SliderCell) SetMaxValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}/* debug [instance_properties/setter]: maxValue */


// The minimum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/minValue
func (s_ SliderCell) MinValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minValue"))
	return rv
}/* debug [instance_properties/getter]: minValue */


// The minimum value the slider can send to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/minValue
func (s_ SliderCell) SetMinValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinValue:"), value)
}/* debug [instance_properties/setter]: minValue */


// The number of tick marks associated with the slider, including the tick marks assigned to the minimum and maximum values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/numberOfTickMarks
func (s_ SliderCell) NumberOfTickMarks() int {
	rv := objc.Send[int](s_.ID, objc.Sel("numberOfTickMarks"))
	return rv
}/* debug [instance_properties/getter]: numberOfTickMarks */


// The number of tick marks associated with the slider, including the tick marks assigned to the minimum and maximum values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/numberOfTickMarks
func (s_ SliderCell) SetNumberOfTickMarks(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNumberOfTickMarks:"), value)
}/* debug [instance_properties/setter]: numberOfTickMarks */


// Returns a Boolean value indicating whether the continues to track the pointer until the next mouse up.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/prefersTrackingUntilMouseUp
func (s_ SliderCell) PrefersTrackingUntilMouseUp() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("prefersTrackingUntilMouseUp"))
	return rv
}/* debug [instance_properties/getter]: prefersTrackingUntilMouseUp */


// The slider type, either linear or circular.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/sliderType
func (s_ SliderCell) SliderType() SliderType {
	rv := objc.Send[SliderType](s_.ID, objc.Sel("sliderType"))
	return rv
}/* debug [instance_properties/getter]: sliderType */


// The slider type, either linear or circular.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/sliderType
func (s_ SliderCell) SetSliderType(value SliderType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSliderType:"), value)
}/* debug [instance_properties/setter]: sliderType */


// The position of the tick marks relative to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/tickMarkPosition
func (s_ SliderCell) TickMarkPosition() TickMarkPosition {
	rv := objc.Send[TickMarkPosition](s_.ID, objc.Sel("tickMarkPosition"))
	return rv
}/* debug [instance_properties/getter]: tickMarkPosition */


// The position of the tick marks relative to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/tickMarkPosition
func (s_ SliderCell) SetTickMarkPosition(value TickMarkPosition) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTickMarkPosition:"), value)
}/* debug [instance_properties/setter]: tickMarkPosition */


// The rectangle within which the cell tracks the pointer while the mouse button is down.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/trackRect
func (s_ SliderCell) TrackRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("trackRect"))
	return rv
}/* debug [instance_properties/getter]: trackRect */


// An integer indicating the orientation (vertical or horizontal) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/isvertical
func (s_ SliderCell) IsVertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVertical"))
	return rv
}/* debug [instance_properties/getter]: isVertical */


// An integer indicating the orientation (vertical or horizontal) of the slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/isvertical
func (s_ SliderCell) SetIsVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVertical:"), value)
}/* debug [instance_properties/setter]: isVertical */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSliderCell */



