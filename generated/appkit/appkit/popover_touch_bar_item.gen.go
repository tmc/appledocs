// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PopoverTouchBarItem] class.
var PopoverTouchBarItemClass objc.Class

func init() {
	PopoverTouchBarItemClass = objc.GetClass("NSPopoverTouchBarItem")
}

type PopoverTouchBarItem struct {
	objc.ID
}

func PopoverTouchBarItemFrom(ptr unsafe.Pointer) PopoverTouchBarItem {
	return PopoverTouchBarItem{
		ID: objc.ID(ptr),
	}
}




