// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ColorPicker] class.
var colorPickerClass = _ColorPickerClass{objc.GetClass("NSColorPicker")}

type _ColorPickerClass struct {
	class objc.Class
}

// An abstract superclass that implements the default color picking protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker

type ColorPicker struct {
	objectivec.Object
}

// ColorPickerFrom constructs a [ColorPicker] from an unsafe.Pointer.
//
// An abstract superclass that implements the default color picking protocol.
func ColorPickerFrom(ptr unsafe.Pointer) ColorPicker {
	return ColorPicker{objectivec.Object{objc.ID(ptr)}}
}



