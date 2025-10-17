// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SliderAccessory] class.
var sliderAccessoryClass = _SliderAccessoryClass{objc.GetClass("NSSliderAccessory")}

type _SliderAccessoryClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessory

type SliderAccessory struct {
	objectivec.Object
}

// SliderAccessoryFrom constructs a [SliderAccessory] from an unsafe.Pointer.
func SliderAccessoryFrom(ptr unsafe.Pointer) SliderAccessory {
	return SliderAccessory{objectivec.Object{objc.ID(ptr)}}
}



