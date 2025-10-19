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

// An interface definition for the [ImageView] class.
type IImageView interface {
	IControl
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
// Alloc allocates a new instance without initialization.
func (ic _ImageViewClass) Alloc() ImageView {
	rv := objc.Send[ImageView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ic _ImageViewClass) New() ImageView {
	rv := objc.Send[ImageView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageView) Init() ImageView {
	rv := objc.Send[ImageView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageView) Autorelease() ImageView {
	rv := objc.Send[ImageView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageView creates a new ImageView instance.
func NewImageView() ImageView {
	return imageViewClass.New()
}




