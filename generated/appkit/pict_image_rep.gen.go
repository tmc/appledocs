// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PICTImageRep] class.
var pICTImageRepClass = _PICTImageRepClass{objc.GetClass("NSPICTImageRep")}

type _PICTImageRepClass struct {
	class objc.Class
}

// An object that renders an image from a PICT format data stream of version 1, version 2, and extended version 2. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPICTImageRep

type PICTImageRep struct {
	ImageRep
}

// PICTImageRepFrom constructs a [PICTImageRep] from an unsafe.Pointer.
//
// An object that renders an image from a PICT format data stream of version 1, version 2, and extended version 2.
func PICTImageRepFrom(ptr unsafe.Pointer) PICTImageRep {
	return PICTImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}



