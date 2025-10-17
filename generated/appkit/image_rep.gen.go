// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ImageRep] class.
var ImageRepClass objc.Class

func init() {
	ImageRepClass = objc.GetClass("NSImageRep")
}

type ImageRep struct {
	objc.ID
}

func ImageRepFrom(ptr unsafe.Pointer) ImageRep {
	return ImageRep{
		ID: objc.ID(ptr),
	}
}



