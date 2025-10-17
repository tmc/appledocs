// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ColorWell] class.
var colorWellClass = _ColorWellClass{objc.GetClass("NSColorWell")}

type _ColorWellClass struct {
	class objc.Class
}

// A control that displays a color value and lets the user change that color value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell

type ColorWell struct {
	Control
}

// ColorWellFrom constructs a [ColorWell] from an unsafe.Pointer.
//
// A control that displays a color value and lets the user change that color value.
func ColorWellFrom(ptr unsafe.Pointer) ColorWell {
	return ColorWell{
		Control: ControlFrom(ptr),
	}
}



