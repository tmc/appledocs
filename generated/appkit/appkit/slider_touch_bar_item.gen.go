// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SliderTouchBarItem] class.
var SliderTouchBarItemClass objc.Class

func init() {
	SliderTouchBarItemClass = objc.GetClass("NSSliderTouchBarItem")
}

type SliderTouchBarItem struct {
	objc.ID
}

func SliderTouchBarItemFrom(ptr unsafe.Pointer) SliderTouchBarItem {
	return SliderTouchBarItem{
		ID: objc.ID(ptr),
	}
}




