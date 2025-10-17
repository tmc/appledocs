// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PickerTouchBarItem] class.
var pickerTouchBarItemClass = _PickerTouchBarItemClass{objc.GetClass("NSPickerTouchBarItem")}

type _PickerTouchBarItemClass struct {
	class objc.Class
}

// A bar item that provides a picker control with multiple options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem

type PickerTouchBarItem struct {
	TouchBarItem
}

// PickerTouchBarItemFrom constructs a [PickerTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a picker control with multiple options.
func PickerTouchBarItemFrom(ptr unsafe.Pointer) PickerTouchBarItem {
	return PickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}



