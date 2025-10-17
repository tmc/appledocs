// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SharingServicePickerTouchBarItem] class.
var SharingServicePickerTouchBarItemClass objc.Class

func init() {
	SharingServicePickerTouchBarItemClass = objc.GetClass("NSSharingServicePickerTouchBarItem")
}

type SharingServicePickerTouchBarItem struct {
	objc.ID
}

func SharingServicePickerTouchBarItemFrom(ptr unsafe.Pointer) SharingServicePickerTouchBarItem {
	return SharingServicePickerTouchBarItem{
		ID: objc.ID(ptr),
	}
}




