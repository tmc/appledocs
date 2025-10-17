// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorSpace] class.
var ColorSpaceClass objc.Class

func init() {
	ColorSpaceClass = objc.GetClass("NSColorSpace")
}

type ColorSpace struct {
	objc.ID
}

func ColorSpaceFrom(ptr unsafe.Pointer) ColorSpace {
	return ColorSpace{
		ID: objc.ID(ptr),
	}
}



