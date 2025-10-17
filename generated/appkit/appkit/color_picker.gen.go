// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorPicker] class.
var ColorPickerClass objc.Class

func init() {
	ColorPickerClass = objc.GetClass("NSColorPicker")
}

type ColorPicker struct {
	objc.ID
}

func ColorPickerFrom(ptr unsafe.Pointer) ColorPicker {
	return ColorPicker{
		ID: objc.ID(ptr),
	}
}




