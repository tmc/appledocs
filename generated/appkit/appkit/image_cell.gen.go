// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ImageCell] class.
var ImageCellClass objc.Class

func init() {
	ImageCellClass = objc.GetClass("NSImageCell")
}

type ImageCell struct {
	objc.ID
}

func ImageCellFrom(ptr unsafe.Pointer) ImageCell {
	return ImageCell{
		ID: objc.ID(ptr),
	}
}



