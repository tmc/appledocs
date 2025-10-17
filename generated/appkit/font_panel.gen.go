// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FontPanel] class.
var fontPanelClass = _FontPanelClass{objc.GetClass("NSFontPanel")}

type _FontPanelClass struct {
	class objc.Class
}

// An interface definition for the [FontPanel] class.
type IFontPanel interface {
	IPanel
}

// The Font panel—a user interface object that displays a list of available fonts, letting the user preview them and change the font used to display text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel

type FontPanel struct {
	Panel
}

// FontPanelFrom constructs a [FontPanel] from an unsafe.Pointer.
//
// The Font panel—a user interface object that displays a list of available fonts, letting the user preview them and change the font used to display text.
func FontPanelFrom(ptr unsafe.Pointer) FontPanel {
	return FontPanel{
		Panel: PanelFrom(ptr),
	}
}



