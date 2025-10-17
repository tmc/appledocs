// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageCell] class.
var imageCellClass = _ImageCellClass{objc.GetClass("NSImageCell")}

type _ImageCellClass struct {
	class objc.Class
}

// An interface definition for the [ImageCell] class.
type IImageCell interface {
	ICell
}

// An object displays a single image (encapsulated in an object) in a frame. This class provides methods for choosing the frame and for aligning and scaling the image to fit the frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell

type ImageCell struct {
	Cell
}

// ImageCellFrom constructs a [ImageCell] from an unsafe.Pointer.
//
// An object displays a single image (encapsulated in an object) in a frame. This class provides methods for choosing the frame and for aligning and scaling the image to fit the frame.
func ImageCellFrom(ptr unsafe.Pointer) ImageCell {
	return ImageCell{
		Cell: CellFrom(ptr),
	}
}



