// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageView] class.
var (
	ImageViewClass     _ImageViewClass
	ImageViewClassOnce sync.Once
)

func getImageViewClass() _ImageViewClass {
	ImageViewClassOnce.Do(func() {
		ImageViewClass = _ImageViewClass{objc.GetClass("NSImageView")}
	})
	return ImageViewClass
}

type _ImageViewClass struct {
	class objc.Class
}

// An interface definition for the [ImageView] class.
type IImageView interface {
	IControl
	AddSymbolEffectOptions(symbolEffect unsafe.Pointer, options unsafe.Pointer)
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


// Adds a symbol effect to the image view with the specified options and default animation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/addSymbolEffect:options:
func (i_ ImageView) AddSymbolEffectOptions(symbolEffect unsafe.Pointer, options unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("addSymbolEffect:options:"), symbolEffect, options)
}

// A Boolean value indicating whether the image view automatically plays animated images.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/animates
func (i_ ImageView) Animates() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("animates"))
	return rv
}


// SetAnimates sets the value of the animates property.
// A Boolean value indicating whether the image view automatically plays animated images.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/animates
func (i_ ImageView) SetAnimates(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAnimates:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/symbolConfiguration
func (i_ ImageView) SymbolConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("symbolConfiguration"))
	return rv
}


// SetSymbolConfiguration sets the value of the symbolConfiguration property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/symbolConfiguration
func (i_ ImageView) SetSymbolConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSymbolConfiguration:"), value)
}

// A Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/allowscutcopypaste
func (i_ ImageView) AllowsCutCopyPaste() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("allowsCutCopyPaste"))
	return rv
}


// SetAllowsCutCopyPaste sets the value of the allowsCutCopyPaste property.
// A Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/allowscutcopypaste
func (i_ ImageView) SetAllowsCutCopyPaste(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAllowsCutCopyPaste:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/contenttintcolor
func (i_ ImageView) ContentTintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("contentTintColor"))
	return rv
}


// SetContentTintColor sets the value of the contentTintColor property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/contenttintcolor
func (i_ ImageView) SetContentTintColor(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContentTintColor:"), value)
}

// The image displayed by the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/image
func (i_ ImageView) Image() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The image displayed by the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/image
func (i_ ImageView) SetImage(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImage:"), value)
}

// The alignment of the cell’s image inside the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/imagealignment
func (i_ ImageView) ImageAlignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageAlignment"))
	return rv
}


// SetImageAlignment sets the value of the imageAlignment property.
// The alignment of the cell’s image inside the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/imagealignment
func (i_ ImageView) SetImageAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageAlignment:"), value)
}

// The resolved dynamic range of the fully resolved image content.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/imagedynamicrange
func (i_ ImageView) ImageDynamicRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageDynamicRange"))
	return rv
}


// SetImageDynamicRange sets the value of the imageDynamicRange property.
// The resolved dynamic range of the fully resolved image content.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/imagedynamicrange
func (i_ ImageView) SetImageDynamicRange(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageDynamicRange:"), value)
}

// The style of frame that appears around the image.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/imageframestyle
func (i_ ImageView) ImageFrameStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageFrameStyle"))
	return rv
}


// SetImageFrameStyle sets the value of the imageFrameStyle property.
// The style of frame that appears around the image.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/imageframestyle
func (i_ ImageView) SetImageFrameStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageFrameStyle:"), value)
}

// The scaling mode applied to make the cell’s image fit the frame of the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/imagescaling
func (i_ ImageView) ImageScaling() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageScaling"))
	return rv
}


// SetImageScaling sets the value of the imageScaling property.
// The scaling mode applied to make the cell’s image fit the frame of the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/imagescaling
func (i_ ImageView) SetImageScaling(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageScaling:"), value)
}

// A Boolean value indicating whether the user can drag a new image into the image view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/iseditable
func (i_ ImageView) IsEditable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isEditable"))
	return rv
}


// SetIsEditable sets the value of the isEditable property.
// A Boolean value indicating whether the user can drag a new image into the image view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/iseditable
func (i_ ImageView) SetIsEditable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsEditable:"), value)
}

// The preferred dynamic range when displaying an image in the receiving image view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/preferredimagedynamicrange
func (i_ ImageView) PreferredImageDynamicRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("preferredImageDynamicRange"))
	return rv
}


// SetPreferredImageDynamicRange sets the value of the preferredImageDynamicRange property.
// The preferred dynamic range when displaying an image in the receiving image view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/preferredimagedynamicrange
func (i_ ImageView) SetPreferredImageDynamicRange(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredImageDynamicRange:"), value)
}



