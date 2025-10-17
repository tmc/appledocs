// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ButtonTouchBarItem] class.
var buttonTouchBarItemClass = _ButtonTouchBarItemClass{objc.GetClass("NSButtonTouchBarItem")}

type _ButtonTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [ButtonTouchBarItem] class.
type IButtonTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem

type ButtonTouchBarItem struct {
	TouchBarItem
}

// ButtonTouchBarItemFrom constructs a [ButtonTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a button.
func ButtonTouchBarItemFrom(ptr unsafe.Pointer) ButtonTouchBarItem {
	return ButtonTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}



