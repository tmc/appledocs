// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SliderAccessory] class.
var SliderAccessoryClass objc.Class

func init() {
	SliderAccessoryClass = objc.GetClass("NSSliderAccessory")
}

type SliderAccessory struct {
	objc.ID
}

func SliderAccessoryFrom(ptr unsafe.Pointer) SliderAccessory {
	return SliderAccessory{
		ID: objc.ID(ptr),
	}
}




