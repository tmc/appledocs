// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SliderTouchBarItem] class.
var sliderTouchBarItemClass = _SliderTouchBarItemClass{objc.GetClass("NSSliderTouchBarItem")}

type _SliderTouchBarItemClass struct {
	class objc.Class
}

// A bar item that provides a slider control for choosing a value in a range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderTouchBarItem

type SliderTouchBarItem struct {
	TouchBarItem
}

// SliderTouchBarItemFrom constructs a [SliderTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a slider control for choosing a value in a range.
func SliderTouchBarItemFrom(ptr unsafe.Pointer) SliderTouchBarItem {
	return SliderTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}



