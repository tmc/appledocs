// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CachedImageRep] class.
var cachedImageRepClass = _CachedImageRepClass{objc.GetClass("NSCachedImageRep")}

type _CachedImageRepClass struct {
	class objc.Class
}

// An object that stores image data in a form that can be readily transferred to the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCachedImageRep

type CachedImageRep struct {
	ImageRep
}

// CachedImageRepFrom constructs a [CachedImageRep] from an unsafe.Pointer.
//
// An object that stores image data in a form that can be readily transferred to the screen.
func CachedImageRepFrom(ptr unsafe.Pointer) CachedImageRep {
	return CachedImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}



