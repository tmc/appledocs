// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ColorPickerTouchBarItem] class.
var colorPickerTouchBarItemClass = _ColorPickerTouchBarItemClass{objc.GetClass("NSColorPickerTouchBarItem")}

type _ColorPickerTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [ColorPickerTouchBarItem] class.
type IColorPickerTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a system-defined color picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem

type ColorPickerTouchBarItem struct {
	TouchBarItem
}

// ColorPickerTouchBarItemFrom constructs a [ColorPickerTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a system-defined color picker.
func ColorPickerTouchBarItemFrom(ptr unsafe.Pointer) ColorPickerTouchBarItem {
	return ColorPickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}



