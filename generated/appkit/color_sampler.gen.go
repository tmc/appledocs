// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ColorSampler] class.
var colorSamplerClass = _ColorSamplerClass{objc.GetClass("NSColorSampler")}

type _ColorSamplerClass struct {
	class objc.Class
}

// An object that displays the system’s color-sampling interface and returns the selected color to your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSampler

type ColorSampler struct {
	objectivec.Object
}

// ColorSamplerFrom constructs a [ColorSampler] from an unsafe.Pointer.
//
// An object that displays the system’s color-sampling interface and returns the selected color to your app.
func ColorSamplerFrom(ptr unsafe.Pointer) ColorSampler {
	return ColorSampler{objectivec.Object{objc.ID(ptr)}}
}



