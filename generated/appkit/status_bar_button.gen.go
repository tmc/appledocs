// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [StatusBarButton] class.
var statusBarButtonClass = _StatusBarButtonClass{objc.GetClass("NSStatusBarButton")}

type _StatusBarButtonClass struct {
	class objc.Class
}

// The appearance and behavior of an item in the systemwide menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBarButton

type StatusBarButton struct {
	Button
}

// StatusBarButtonFrom constructs a [StatusBarButton] from an unsafe.Pointer.
//
// The appearance and behavior of an item in the systemwide menu bar.
func StatusBarButtonFrom(ptr unsafe.Pointer) StatusBarButton {
	return StatusBarButton{
		Button: ButtonFrom(ptr),
	}
}



