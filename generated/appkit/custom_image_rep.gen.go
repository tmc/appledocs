// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CustomImageRep] class.
var customImageRepClass = _CustomImageRepClass{objc.GetClass("NSCustomImageRep")}

type _CustomImageRepClass struct {
	class objc.Class
}

// An interface definition for the [CustomImageRep] class.
type ICustomImageRep interface {
	IImageRep
}

// An object that uses a delegate object to render an image from a custom format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep

type CustomImageRep struct {
	ImageRep
}

// CustomImageRepFrom constructs a [CustomImageRep] from an unsafe.Pointer.
//
// An object that uses a delegate object to render an image from a custom format.
func CustomImageRepFrom(ptr unsafe.Pointer) CustomImageRep {
	return CustomImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}



