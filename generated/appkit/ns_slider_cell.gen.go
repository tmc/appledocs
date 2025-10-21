// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

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

// An interface definition for the [SliderCell] class.
type ISliderCell interface {
	IActionCell
	KnobRectFlipped(flipped bool) coregraphics.CGRect
}

// The appearance and behavior of an object.
//
// You can customize an to a certain degree, using its properties. If this doesn’t give you sufficient flexibility, you can create a subclass. In that subclass, you can override any of the following methods: , , , and .
//
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

// Alloc allocates a new instance without initialization.
func (sc _SliderCellClass) Alloc() SliderCell {
	rv := objc.Send[SliderCell](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the rectangle in which the slider knob is drawn.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/knobRect(flipped:)
func (s_ SliderCell) KnobRectFlipped(flipped bool) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("knobRectFlipped:"), flipped)
	return rv
}

// The slider type, either linear or circular.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/sliderType
func (s_ SliderCell) SliderType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sliderType"))
	return rv
}


// SetSliderType sets the value of the sliderType property.
// The slider type, either linear or circular.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/sliderType
func (s_ SliderCell) SetSliderType(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSliderType:"), value)
}

// A Boolean value indicating whether the receiver fixes its values to those values represented by its tick marks.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/allowstickmarkvaluesonly
func (s_ SliderCell) AllowsTickMarkValuesOnly() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsTickMarkValuesOnly"))
	return rv
}


// SetAllowsTickMarkValuesOnly sets the value of the allowsTickMarkValuesOnly property.
// A Boolean value indicating whether the receiver fixes its values to those values represented by its tick marks.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/allowstickmarkvaluesonly
func (s_ SliderCell) SetAllowsTickMarkValuesOnly(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsTickMarkValuesOnly:"), value)
}

// The amount by which the slider changes its value when the user Option-drags the knob.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/altincrementvalue
func (s_ SliderCell) AltIncrementValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("altIncrementValue"))
	return rv
}


// SetAltIncrementValue sets the value of the altIncrementValue property.
// The amount by which the slider changes its value when the user Option-drags the knob.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/altincrementvalue
func (s_ SliderCell) SetAltIncrementValue(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAltIncrementValue:"), value)
}

// An integer indicating the orientation (vertical or horizontal) of the slider.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/isvertical
func (s_ SliderCell) IsVertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVertical"))
	return rv
}


// SetIsVertical sets the value of the isVertical property.
// An integer indicating the orientation (vertical or horizontal) of the slider.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/isvertical
func (s_ SliderCell) SetIsVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVertical:"), value)
}

// The thickness of the slider knob, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/knobthickness
func (s_ SliderCell) KnobThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("knobThickness"))
	return rv
}


// SetKnobThickness sets the value of the knobThickness property.
// The thickness of the slider knob, in pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/knobthickness
func (s_ SliderCell) SetKnobThickness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKnobThickness:"), value)
}

// The maximum value the slider can send to its target.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/maxvalue
func (s_ SliderCell) MaxValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("maxValue"))
	return rv
}


// SetMaxValue sets the value of the maxValue property.
// The maximum value the slider can send to its target.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/maxvalue
func (s_ SliderCell) SetMaxValue(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}

// The minimum value the slider can send to its target.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/minvalue
func (s_ SliderCell) MinValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("minValue"))
	return rv
}


// SetMinValue sets the value of the minValue property.
// The minimum value the slider can send to its target.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/minvalue
func (s_ SliderCell) SetMinValue(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinValue:"), value)
}

// The number of tick marks associated with the slider, including the tick marks assigned to the minimum and maximum values.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/numberoftickmarks
func (s_ SliderCell) NumberOfTickMarks() int {
	rv := objc.Send[int](s_.ID, objc.Sel("numberOfTickMarks"))
	return rv
}


// SetNumberOfTickMarks sets the value of the numberOfTickMarks property.
// The number of tick marks associated with the slider, including the tick marks assigned to the minimum and maximum values.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/numberoftickmarks
func (s_ SliderCell) SetNumberOfTickMarks(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNumberOfTickMarks:"), value)
}

// The position of the tick marks relative to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/tickmarkposition
func (s_ SliderCell) TickMarkPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("tickMarkPosition"))
	return rv
}


// SetTickMarkPosition sets the value of the tickMarkPosition property.
// The position of the tick marks relative to the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/tickmarkposition
func (s_ SliderCell) SetTickMarkPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTickMarkPosition:"), value)
}

// The rectangle within which the cell tracks the pointer while the mouse button is down.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/trackrect
func (s_ SliderCell) TrackRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("trackRect"))
	return rv
}


// SetTrackRect sets the value of the trackRect property.
// The rectangle within which the cell tracks the pointer while the mouse button is down.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslidercell/trackrect
func (s_ SliderCell) SetTrackRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTrackRect:"), value)
}



