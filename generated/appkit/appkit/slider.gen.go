// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Slider] class.
var (
	sliderClass     _SliderClass
	sliderClassOnce sync.Once
)

func getSliderClass() _SliderClass {
	sliderClassOnce.Do(func() {
		sliderClass = _SliderClass{objc.GetClass("NSSlider")}
	})
	return sliderClass
}

type _SliderClass struct {
	class objc.Class
}

// An interface definition for the [Slider] class.
type ISlider interface {
	IControl
	SetTitleFont(fontObj unsafe.Pointer)
}

// A display of a bar representing a continuous range of numerical values and a knob representing the currently selected value. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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


// Sets the font used to draw the slider’s title. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/setTitleFont:
func (s_ Slider) SetTitleFont(fontObj unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitleFont:"), fontObj)
}


