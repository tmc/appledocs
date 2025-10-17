// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Slider] class.
var sliderClass = _SliderClass{objc.GetClass("NSSlider")}

type _SliderClass struct {
	class objc.Class
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

// Sets the font used to draw the slider’s title. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/setTitleFont:
func (s_ Slider) SetTitleFont(fontObj unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitleFont:"), fontObj)
}


