// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ImageCell] class.
var (
	ImageCellClass     _ImageCellClass
	ImageCellClassOnce sync.Once
)

func getImageCellClass() _ImageCellClass {
	ImageCellClassOnce.Do(func() {
		ImageCellClass = _ImageCellClass{objc.GetClass("NSImageCell")}
	})
	return ImageCellClass
}

type _ImageCellClass struct {
	class objc.Class
}





// An interface definition for the [ImageCell] class.
type IImageCell interface {
	ICell
	

	// properties:
	ImageAlignment() ImageAlignment
	SetImageAlignment(value ImageAlignment)
	ImageFrameStyle() ImageFrameStyle
	SetImageFrameStyle(value ImageFrameStyle)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageCellClass) Alloc() ImageCell {
	rv := objc.Send[ImageCell](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageCellClass) New() ImageCell {
	rv := objc.Send[ImageCell](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageCell) Init() ImageCell {
	rv := objc.Send[ImageCell](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageCell) Autorelease() ImageCell {
	rv := objc.Send[ImageCell](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageCell creates a new ImageCell instance.
func NewImageCell() ImageCell {
	return getImageCellClass().New()
}





// An object displays a single image (encapsulated in an object) in a frame. This class provides methods for choosing the frame and for aligning and scaling the image to fit the frame.
//
// The object value of an object must be an object, so if you use the method of , be sure to supply an object as an argument. Because an object does not need to be converted for display, do not use the methods relating to formatters. An object is usually associated with some kind of control object. For example, an or an .


// An object displays a single image (encapsulated in an object) in a frame. This class provides methods for choosing the frame and for aligning and scaling the image to fit the frame.
//
// [Full Topic]
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

























// The alignment of the receiver’s image relative to its frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageAlignment
func (i_ ImageCell) ImageAlignment() ImageAlignment {
	rv := objc.Send[ImageAlignment](i_.ID, objc.Sel("imageAlignment"))
	return rv
}


// The alignment of the receiver’s image relative to its frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageAlignment
func (i_ ImageCell) SetImageAlignment(value ImageAlignment) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageAlignment:"), value)
}


// The style of the frame that borders the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageFrameStyle
func (i_ ImageCell) ImageFrameStyle() ImageFrameStyle {
	rv := objc.Send[ImageFrameStyle](i_.ID, objc.Sel("imageFrameStyle"))
	return rv
}


// The style of the frame that borders the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageFrameStyle
func (i_ ImageCell) SetImageFrameStyle(value ImageFrameStyle) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageFrameStyle:"), value)
}


// The scaling mode used to fit the receiver’s image into the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageScaling
func (i_ ImageCell) ImageScaling() ImageScaling {
	rv := objc.Send[ImageScaling](i_.ID, objc.Sel("imageScaling"))
	return rv
}


// The scaling mode used to fit the receiver’s image into the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageScaling
func (i_ ImageCell) SetImageScaling(value ImageScaling) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageScaling:"), value)
}








