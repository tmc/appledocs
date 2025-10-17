// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CIImageRep] class.
var cIImageRepClass = _CIImageRepClass{objc.GetClass("NSCIImageRep")}

type _CIImageRepClass struct {
	class objc.Class
}

// An interface definition for the [CIImageRep] class.
type ICIImageRep interface {
	IImageRep
}

// An object that can render an image from a Core Image object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCIImageRep

type CIImageRep struct {
	ImageRep
}

// CIImageRepFrom constructs a [CIImageRep] from an unsafe.Pointer.
//
// An object that can render an image from a Core Image object.
func CIImageRepFrom(ptr unsafe.Pointer) CIImageRep {
	return CIImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}



