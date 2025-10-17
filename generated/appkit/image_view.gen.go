// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageView] class.
var imageViewClass = _ImageViewClass{objc.GetClass("NSImageView")}

type _ImageViewClass struct {
	class objc.Class
}

// A display of image data in a frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView

type ImageView struct {
	Control
}

// ImageViewFrom constructs a [ImageView] from an unsafe.Pointer.
//
// A display of image data in a frame.
func ImageViewFrom(ptr unsafe.Pointer) ImageView {
	return ImageView{
		Control: ControlFrom(ptr),
	}
}



