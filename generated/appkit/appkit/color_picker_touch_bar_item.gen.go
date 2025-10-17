// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorPickerTouchBarItem] class.
var ColorPickerTouchBarItemClass objc.Class

func init() {
	ColorPickerTouchBarItemClass = objc.GetClass("NSColorPickerTouchBarItem")
}

type ColorPickerTouchBarItem struct {
	objc.ID
}

func ColorPickerTouchBarItemFrom(ptr unsafe.Pointer) ColorPickerTouchBarItem {
	return ColorPickerTouchBarItem{
		ID: objc.ID(ptr),
	}
}




