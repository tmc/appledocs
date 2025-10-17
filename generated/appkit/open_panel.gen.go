// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OpenPanel] class.
var openPanelClass = _OpenPanelClass{objc.GetClass("NSOpenPanel")}

type _OpenPanelClass struct {
	class objc.Class
}

// A panel that prompts the user to select a file to open. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel

type OpenPanel struct {
	SavePanel
}

// OpenPanelFrom constructs a [OpenPanel] from an unsafe.Pointer.
//
// A panel that prompts the user to select a file to open.
func OpenPanelFrom(ptr unsafe.Pointer) OpenPanel {
	return OpenPanel{
		SavePanel: SavePanelFrom(ptr),
	}
}



