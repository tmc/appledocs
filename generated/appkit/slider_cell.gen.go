// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SliderCell] class.
var SliderCellClass objc.Class

func init() {
	SliderCellClass = objc.GetClass("NSSliderCell")
}

type SliderCell struct {
	objc.ID
}

func SliderCellFrom(ptr unsafe.Pointer) SliderCell {
	return SliderCell{
		ID: objc.ID(ptr),
	}
}



