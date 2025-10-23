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
	ObjectValue() unsafe.Pointer
	SetObjectValue(value unsafe.Pointer)
	ImageAlignment() unsafe.Pointer
	SetImageAlignment(value unsafe.Pointer)
	ImageFrameStyle() unsafe.Pointer
	SetImageFrameStyle(value unsafe.Pointer)
	ImageScaling() ImageScaling
	SetImageScaling(value IImageScaling)
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

// Alloc allocates a new instance without initialization.
func (ic _ImageCellClass) Alloc() ImageCell {
	rv := objc.Send[ImageCell](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The cell’s value as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/objectvalue
func (i_ ImageCell) ObjectValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("objectValue"))
	return rv
}


// The cell’s value as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/objectvalue
func (i_ ImageCell) SetObjectValue(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setObjectValue:"), value)
}


// The alignment of the receiver’s image relative to its frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/imagealignment
func (i_ ImageCell) ImageAlignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageAlignment"))
	return rv
}


// The alignment of the receiver’s image relative to its frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/imagealignment
func (i_ ImageCell) SetImageAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageAlignment:"), value)
}


// The style of the frame that borders the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/imageframestyle
func (i_ ImageCell) ImageFrameStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageFrameStyle"))
	return rv
}


// The style of the frame that borders the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/imageframestyle
func (i_ ImageCell) SetImageFrameStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageFrameStyle:"), value)
}


// The scaling mode used to fit the receiver’s image into the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/imagescaling
func (i_ ImageCell) ImageScaling() ImageScaling {
	rv := objc.Send[ImageScaling](i_.ID, objc.Sel("imageScaling"))
	return rv
}


// The scaling mode used to fit the receiver’s image into the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagecell/imagescaling
func (i_ ImageCell) SetImageScaling(value IImageScaling) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageScaling:"), value)
}



