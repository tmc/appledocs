// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EPSImageRep] class.
var ePSImageRepClass = _EPSImageRepClass{objc.GetClass("NSEPSImageRep")}

type _EPSImageRepClass struct {
	class objc.Class
}

// An object that can render an image from encapsulated PostScript (EPS) code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEPSImageRep

type EPSImageRep struct {
	ImageRep
}

// EPSImageRepFrom constructs a [EPSImageRep] from an unsafe.Pointer.
//
// An object that can render an image from encapsulated PostScript (EPS) code.
func EPSImageRepFrom(ptr unsafe.Pointer) EPSImageRep {
	return EPSImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}



