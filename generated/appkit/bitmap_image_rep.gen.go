// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BitmapImageRep] class.
var bitmapImageRepClass = _BitmapImageRepClass{objc.GetClass("NSBitmapImageRep")}

type _BitmapImageRepClass struct {
	class objc.Class
}

// An interface definition for the [BitmapImageRep] class.
type IBitmapImageRep interface {
	IImageRep
}

// An object that renders an image from bitmap data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep

type BitmapImageRep struct {
	ImageRep
}

// BitmapImageRepFrom constructs a [BitmapImageRep] from an unsafe.Pointer.
//
// An object that renders an image from bitmap data.
func BitmapImageRepFrom(ptr unsafe.Pointer) BitmapImageRep {
	return BitmapImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}



