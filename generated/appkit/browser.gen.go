// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Browser] class.
var browserClass = _BrowserClass{objc.GetClass("NSBrowser")}

type _BrowserClass struct {
	class objc.Class
}

// An interface that displays a hierarchically organized list of data items that can be navigated and selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser

type Browser struct {
	Control
}

// BrowserFrom constructs a [Browser] from an unsafe.Pointer.
//
// An interface that displays a hierarchically organized list of data items that can be navigated and selected.
func BrowserFrom(ptr unsafe.Pointer) Browser {
	return Browser{
		Control: ControlFrom(ptr),
	}
}



