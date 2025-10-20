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


