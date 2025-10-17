// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ColorPanel] class.
var colorPanelClass = _ColorPanelClass{objc.GetClass("NSColorPanel")}

type _ColorPanelClass struct {
	class objc.Class
}

// An interface definition for the [ColorPanel] class.
type IColorPanel interface {
	IPanel
}

// A standard user interface for selecting color in an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel

type ColorPanel struct {
	Panel
}

// ColorPanelFrom constructs a [ColorPanel] from an unsafe.Pointer.
//
// A standard user interface for selecting color in an app.
func ColorPanelFrom(ptr unsafe.Pointer) ColorPanel {
	return ColorPanel{
		Panel: PanelFrom(ptr),
	}
}



