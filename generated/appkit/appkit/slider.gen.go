// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Slider] class.
var SliderClass objc.Class

func init() {
	SliderClass = objc.GetClass("NSSlider")
}

type Slider struct {
	objc.ID
}

func SliderFrom(ptr unsafe.Pointer) Slider {
	return Slider{
		ID: objc.ID(ptr),
	}
}


// Sets the font used to draw the slider’s title. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSlider/setTitleFont:
func (s_ Slider) SetTitleFont(fontObj unsafe.Pointer) {
	sel := objc.RegisterName("setTitleFont:")
	s_.ID.Send(sel, fontObj)
}


