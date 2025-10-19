// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageView] class.
var (
	imageViewClass     _ImageViewClass
	imageViewClassOnce sync.Once
)

func getImageViewClass() _ImageViewClass {
	imageViewClassOnce.Do(func() {
		imageViewClass = _ImageViewClass{objc.GetClass("NSImageView")}
	})
	return imageViewClass
}

type _ImageViewClass struct {
	class objc.Class
}

// An interface definition for the [ImageView] class.
type IImageView interface {
	IControl
}

// A display of image data in a frame.
//
// Image views can be static or editable. A static image view only displays the image that you specify. An editable image view object lets the user change the displayed image. You can also configure an image view to allow copying, pasting, deleting, and dragging of the image.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getImageViewClass().New()
}




