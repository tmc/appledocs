// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PickerTouchBarItem] class.
var PickerTouchBarItemClass objc.Class

func init() {
	PickerTouchBarItemClass = objc.GetClass("NSPickerTouchBarItem")
}

type PickerTouchBarItem struct {
	objc.ID
}

func PickerTouchBarItemFrom(ptr unsafe.Pointer) PickerTouchBarItem {
	return PickerTouchBarItem{
		ID: objc.ID(ptr),
	}
}



