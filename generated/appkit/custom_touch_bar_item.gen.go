// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CustomTouchBarItem] class.
var customTouchBarItemClass = _CustomTouchBarItemClass{objc.GetClass("NSCustomTouchBarItem")}

type _CustomTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [CustomTouchBarItem] class.
type ICustomTouchBarItem interface {
	ITouchBarItem
}

// A bar item that contains a responder of your choice, such as a view, a button, or a scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem

type CustomTouchBarItem struct {
	TouchBarItem
}

// CustomTouchBarItemFrom constructs a [CustomTouchBarItem] from an unsafe.Pointer.
//
// A bar item that contains a responder of your choice, such as a view, a button, or a scrubber.
func CustomTouchBarItemFrom(ptr unsafe.Pointer) CustomTouchBarItem {
	return CustomTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}



