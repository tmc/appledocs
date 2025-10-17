// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PopoverTouchBarItem] class.
var popoverTouchBarItemClass = _PopoverTouchBarItemClass{objc.GetClass("NSPopoverTouchBarItem")}

type _PopoverTouchBarItemClass struct {
	class objc.Class
}

// A bar item that provides a two-state control that can expand into its second state, showing the contents of a bar that it owns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem

type PopoverTouchBarItem struct {
	TouchBarItem
}

// PopoverTouchBarItemFrom constructs a [PopoverTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a two-state control that can expand into its second state, showing the contents of a bar that it owns.
func PopoverTouchBarItemFrom(ptr unsafe.Pointer) PopoverTouchBarItem {
	return PopoverTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}



