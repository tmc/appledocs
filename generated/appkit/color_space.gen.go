// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ColorSpace] class.
var colorSpaceClass = _ColorSpaceClass{objc.GetClass("NSColorSpace")}

type _ColorSpaceClass struct {
	class objc.Class
}

// An object that represents a custom color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace

type ColorSpace struct {
	objectivec.Object
}

// ColorSpaceFrom constructs a [ColorSpace] from an unsafe.Pointer.
//
// An object that represents a custom color space.
func ColorSpaceFrom(ptr unsafe.Pointer) ColorSpace {
	return ColorSpace{objectivec.Object{objc.ID(ptr)}}
}



