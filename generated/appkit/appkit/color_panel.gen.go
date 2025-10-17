// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorPanel] class.
var ColorPanelClass objc.Class

func init() {
	ColorPanelClass = objc.GetClass("NSColorPanel")
}

type ColorPanel struct {
	objc.ID
}

func ColorPanelFrom(ptr unsafe.Pointer) ColorPanel {
	return ColorPanel{
		ID: objc.ID(ptr),
	}
}



