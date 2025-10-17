// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SliderCell] class.
var sliderCellClass = _SliderCellClass{objc.GetClass("NSSliderCell")}

type _SliderCellClass struct {
	class objc.Class
}

// The appearance and behavior of an object. [Full Topic]
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



