// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BrowserCell] class.
var browserCellClass = _BrowserCellClass{objc.GetClass("NSBrowserCell")}

type _BrowserCellClass struct {
	class objc.Class
}

// The user interface of a browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell

type BrowserCell struct {
	Cell
}

// BrowserCellFrom constructs a [BrowserCell] from an unsafe.Pointer.
//
// The user interface of a browser.
func BrowserCellFrom(ptr unsafe.Pointer) BrowserCell {
	return BrowserCell{
		Cell: CellFrom(ptr),
	}
}



