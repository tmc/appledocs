// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GroupTouchBarItem] class.
var groupTouchBarItemClass = _GroupTouchBarItemClass{objc.GetClass("NSGroupTouchBarItem")}

type _GroupTouchBarItemClass struct {
	class objc.Class
}

// A bar item that provides a bar to contain other items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem

type GroupTouchBarItem struct {
	TouchBarItem
}

// GroupTouchBarItemFrom constructs a [GroupTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a bar to contain other items.
func GroupTouchBarItemFrom(ptr unsafe.Pointer) GroupTouchBarItem {
	return GroupTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}



