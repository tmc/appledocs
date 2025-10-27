// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	AllowsCutCopyPaste() bool
	SetAllowsCutCopyPaste(value bool)
	Animates() bool
	SetAnimates(value bool)
	ContentTintColor() IColor
	SetContentTintColor(value IColor)
	Image() IImage
	SetImage(value IImage)
	ImageAlignment() ImageAlignment
	SetImageAlignment(value ImageAlignment)
	ImageDynamicRange() ImageDynamicRange
	ImageFrameStyle() ImageFrameStyle
	SetImageFrameStyle(value ImageFrameStyle)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	Editable() bool
	SetEditable(value bool)
	PreferredImageDynamicRange() ImageDynamicRange
	SetPreferredImageDynamicRange(value ImageDynamicRange)
	SymbolConfiguration() IImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
	IsEditable() bool
	SetIsEditable(value bool)


	

	// methods:
	AddSymbolEffect(symbolEffect SymbolEffect /* not a class type */)
	AddSymbolEffectOptions(symbolEffect SymbolEffect /* not a class type */, options SymbolEffectOptions /* not a class type */)
	AddSymbolEffectOptionsAnimated(symbolEffect SymbolEffect /* not a class type */, options SymbolEffectOptions /* not a class type */, animated bool)
	RemoveAllSymbolEffects()
	RemoveAllSymbolEffectsWithOptions(options SymbolEffectOptions /* not a class type */)
	RemoveAllSymbolEffectsWithOptionsAnimated(options SymbolEffectOptions /* not a class type */, animated bool)
	RemoveSymbolEffectOfType(symbolEffect SymbolEffect /* not a class type */)
	RemoveSymbolEffectOfTypeOptions(symbolEffect SymbolEffect /* not a class type */, options SymbolEffectOptions /* not a class type */)
	RemoveSymbolEffectOfTypeOptionsAnimated(symbolEffect SymbolEffect /* not a class type */, options SymbolEffectOptions /* not a class type */, animated bool)
	SetSymbolImageWithContentTransition(symbolImage IImage, transition SymbolContentTransition /* not a class type */)
	SetSymbolImageWithContentTransitionOptions(symbolImage IImage, transition SymbolContentTransition /* not a class type */, options SymbolEffectOptions /* not a class type */)


}





// Alloc allocates a new instance without initialization.
func (ic _ImageViewClass) Alloc() ImageView {
	rv := objc.Send[ImageView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A display of image data in a frame.
//
// Image views can be static or editable. A static image view only displays the image that you specify. An editable image view object lets the user change the displayed image. You can also configure an image view to allow copying, pasting, deleting, and dragging of the image.


// A display of image data in a frame.
//
// [Full Topic]
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/init(image:)
func NewImageViewWithImage(image IImage) ImageView {
	rv := objc.Send[ImageView](objc.ID(getImageViewClass().class), objc.Sel("imageViewWithImage:"), image)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/init(image:)
func (ic _ImageViewClass) ImageViewWithImage(image IImage) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageViewWithImage:"), image)
	return rv
}







// The default preferred image dynamic range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/defaultPreferredImageDynamicRange
func (ic _ImageViewClass) DefaultPreferredImageDynamicRange() ImageDynamicRange {
	rv := objc.Send[ImageDynamicRange](objc.ID(ic.class), objc.Sel("defaultPreferredImageDynamicRange"))
	return rv
}






// Adds a symbol effect to the image view with default options and animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/addSymbolEffect:
func (i_ ImageView) AddSymbolEffect(symbolEffect SymbolEffect /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("addSymbolEffect:"), symbolEffect)
}


// Adds a symbol effect to the image view with the specified options and default animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/addSymbolEffect:options:
func (i_ ImageView) AddSymbolEffectOptions(symbolEffect SymbolEffect /* not a class type */, options SymbolEffectOptions /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("addSymbolEffect:options:"), symbolEffect, options)
}


// Adds a symbol effect to the image view with the specified options and animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/addSymbolEffect:options:animated:
func (i_ ImageView) AddSymbolEffectOptionsAnimated(symbolEffect SymbolEffect /* not a class type */, options SymbolEffectOptions /* not a class type */, animated bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("addSymbolEffect:options:animated:"), symbolEffect, options, animated)
}


// Removes all symbol effects from the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/removeAllSymbolEffects
func (i_ ImageView) RemoveAllSymbolEffects() {
	objc.Send[objc.ID](i_.ID, objc.Sel("removeAllSymbolEffects"))
}


// Removes all symbol effects from the image view, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/removeAllSymbolEffectsWithOptions:
func (i_ ImageView) RemoveAllSymbolEffectsWithOptions(options SymbolEffectOptions /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("removeAllSymbolEffectsWithOptions:"), options)
}


// Removes all symbol effects from the image view, using the specified options and animation setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/removeAllSymbolEffectsWithOptions:animated:
func (i_ ImageView) RemoveAllSymbolEffectsWithOptionsAnimated(options SymbolEffectOptions /* not a class type */, animated bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("removeAllSymbolEffectsWithOptions:animated:"), options, animated)
}


// Removes the symbol effect that matches the specified effect type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/removeSymbolEffectOfType:
func (i_ ImageView) RemoveSymbolEffectOfType(symbolEffect SymbolEffect /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("removeSymbolEffectOfType:"), symbolEffect)
}


// Removes the symbol effect that matches the specified effect type, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/removeSymbolEffectOfType:options:
func (i_ ImageView) RemoveSymbolEffectOfTypeOptions(symbolEffect SymbolEffect /* not a class type */, options SymbolEffectOptions /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("removeSymbolEffectOfType:options:"), symbolEffect, options)
}


// Removes the symbol effect that matches the specified effect type, using the specified options and animation setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/removeSymbolEffectOfType:options:animated:
func (i_ ImageView) RemoveSymbolEffectOfTypeOptionsAnimated(symbolEffect SymbolEffect /* not a class type */, options SymbolEffectOptions /* not a class type */, animated bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("removeSymbolEffectOfType:options:animated:"), symbolEffect, options, animated)
}


// Sets a symbol image using the specified content-transition effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/setSymbolImage:withContentTransition:
func (i_ ImageView) SetSymbolImageWithContentTransition(symbolImage IImage, transition SymbolContentTransition /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSymbolImage:withContentTransition:"), symbolImage, transition)
}


// Sets a symbol image using the specified content-transition effect and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/setSymbolImage:withContentTransition:options:
func (i_ ImageView) SetSymbolImageWithContentTransitionOptions(symbolImage IImage, transition SymbolContentTransition /* not a class type */, options SymbolEffectOptions /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSymbolImage:withContentTransition:options:"), symbolImage, transition, options)
}







// A Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/allowsCutCopyPaste
func (i_ ImageView) AllowsCutCopyPaste() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("allowsCutCopyPaste"))
	return rv
}


// A Boolean value indicating whether the image view lets the user cut, copy, and paste the image contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/allowsCutCopyPaste
func (i_ ImageView) SetAllowsCutCopyPaste(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAllowsCutCopyPaste:"), value)
}


// A Boolean value indicating whether the image view automatically plays animated images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/animates
func (i_ ImageView) Animates() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("animates"))
	return rv
}


// A Boolean value indicating whether the image view automatically plays animated images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/animates
func (i_ ImageView) SetAnimates(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAnimates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/contentTintColor
func (i_ ImageView) ContentTintColor() IColor {
	rv := objc.Send[Color](i_.ID, objc.Sel("contentTintColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/contentTintColor
func (i_ ImageView) SetContentTintColor(value IColor) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContentTintColor:"), value)
}


// The default preferred image dynamic range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/defaultPreferredImageDynamicRange
func (i_ ImageView) DefaultPreferredImageDynamicRange() ImageDynamicRange {
	rv := objc.Send[ImageDynamicRange](i_.ID, objc.Sel("defaultPreferredImageDynamicRange"))
	return rv
}


// The default preferred image dynamic range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/defaultPreferredImageDynamicRange
func (i_ ImageView) SetDefaultPreferredImageDynamicRange(value ImageDynamicRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDefaultPreferredImageDynamicRange:"), value)
}


// The image displayed by the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/image
func (i_ ImageView) Image() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("image"))
	return rv
}


// The image displayed by the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/image
func (i_ ImageView) SetImage(value IImage) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImage:"), value)
}


// The alignment of the cell’s image inside the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/imageAlignment
func (i_ ImageView) ImageAlignment() ImageAlignment {
	rv := objc.Send[ImageAlignment](i_.ID, objc.Sel("imageAlignment"))
	return rv
}


// The alignment of the cell’s image inside the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/imageAlignment
func (i_ ImageView) SetImageAlignment(value ImageAlignment) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageAlignment:"), value)
}


// The resolved dynamic range of the fully resolved image content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/imageDynamicRange
func (i_ ImageView) ImageDynamicRange() ImageDynamicRange {
	rv := objc.Send[ImageDynamicRange](i_.ID, objc.Sel("imageDynamicRange"))
	return rv
}


// The style of frame that appears around the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/imageFrameStyle
func (i_ ImageView) ImageFrameStyle() ImageFrameStyle {
	rv := objc.Send[ImageFrameStyle](i_.ID, objc.Sel("imageFrameStyle"))
	return rv
}


// The style of frame that appears around the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/imageFrameStyle
func (i_ ImageView) SetImageFrameStyle(value ImageFrameStyle) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageFrameStyle:"), value)
}


// The scaling mode applied to make the cell’s image fit the frame of the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/imageScaling
func (i_ ImageView) ImageScaling() ImageScaling {
	rv := objc.Send[ImageScaling](i_.ID, objc.Sel("imageScaling"))
	return rv
}


// The scaling mode applied to make the cell’s image fit the frame of the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/imageScaling
func (i_ ImageView) SetImageScaling(value ImageScaling) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageScaling:"), value)
}


// A Boolean value indicating whether the user can drag a new image into the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/isEditable
func (i_ ImageView) Editable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("editable"))
	return rv
}


// A Boolean value indicating whether the user can drag a new image into the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/isEditable
func (i_ ImageView) SetEditable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEditable:"), value)
}


// The preferred dynamic range when displaying an image in the receiving image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/preferredImageDynamicRange
func (i_ ImageView) PreferredImageDynamicRange() ImageDynamicRange {
	rv := objc.Send[ImageDynamicRange](i_.ID, objc.Sel("preferredImageDynamicRange"))
	return rv
}


// The preferred dynamic range when displaying an image in the receiving image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/preferredImageDynamicRange
func (i_ ImageView) SetPreferredImageDynamicRange(value ImageDynamicRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredImageDynamicRange:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/symbolConfiguration
func (i_ ImageView) SymbolConfiguration() IImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](i_.ID, objc.Sel("symbolConfiguration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/symbolConfiguration
func (i_ ImageView) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSymbolConfiguration:"), value)
}


// A Boolean value indicating whether the user can drag a new image into the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/iseditable
func (i_ ImageView) IsEditable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean value indicating whether the user can drag a new image into the image view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/iseditable
func (i_ ImageView) SetIsEditable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsEditable:"), value)
}







