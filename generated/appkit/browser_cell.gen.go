// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BrowserCell] class.
var BrowserCellClass objc.Class

func init() {
	BrowserCellClass = objc.GetClass("NSBrowserCell")
}

type BrowserCell struct {
	objc.ID
}

func BrowserCellFrom(ptr unsafe.Pointer) BrowserCell {
	return BrowserCell{
		ID: objc.ID(ptr),
	}
}



