// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ImageView] class.
var ImageViewClass objc.Class

func init() {
	ImageViewClass = objc.GetClass("NSImageView")
}

type ImageView struct {
	objc.ID
}

func ImageViewFrom(ptr unsafe.Pointer) ImageView {
	return ImageView{
		ID: objc.ID(ptr),
	}
}



