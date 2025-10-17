// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorSampler] class.
var ColorSamplerClass objc.Class

func init() {
	ColorSamplerClass = objc.GetClass("NSColorSampler")
}

type ColorSampler struct {
	objc.ID
}

func ColorSamplerFrom(ptr unsafe.Pointer) ColorSampler {
	return ColorSampler{
		ID: objc.ID(ptr),
	}
}




