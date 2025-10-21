// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberImageItemView] class.
var (
	ScrubberImageItemViewClass     _ScrubberImageItemViewClass
	ScrubberImageItemViewClassOnce sync.Once
)

func getScrubberImageItemViewClass() _ScrubberImageItemViewClass {
	ScrubberImageItemViewClassOnce.Do(func() {
		ScrubberImageItemViewClass = _ScrubberImageItemViewClass{objc.GetClass("NSScrubberImageItemView")}
	})
	return ScrubberImageItemViewClass
}

type _ScrubberImageItemViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberImageItemView] class.
type IScrubberImageItemView interface {
	IScrubberItemView
}

// A concrete view subclass for displaying images in a scrubber items.
//
// Provide the image you want to display in the scrubber item to the property. If you want finer control over the appearance of the image, you can access the underlying image view using the property. The image is scaled proportionally to fit the view’s frame. Use the property to determine how the scaled image is cropped within that frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView
type ScrubberImageItemView struct {
	ScrubberItemView
}

// ScrubberImageItemViewFrom constructs a [ScrubberImageItemView] from an unsafe.Pointer.
//
// A concrete view subclass for displaying images in a scrubber items.
func ScrubberImageItemViewFrom(ptr unsafe.Pointer) ScrubberImageItemView {
	return ScrubberImageItemView{
		ScrubberItemView: ScrubberItemViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberImageItemViewClass) Alloc() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrubberImageItemViewClass) New() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberImageItemView) Init() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberImageItemView) Autorelease() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberImageItemView creates a new ScrubberImageItemView instance.
func NewScrubberImageItemView() ScrubberImageItemView {
	return getScrubberImageItemViewClass().New()
}


// The alignment of the image within the scrubber item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView/imageAlignment
func (s_ ScrubberImageItemView) ImageAlignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("imageAlignment"))
	return rv
}


// SetImageAlignment sets the value of the imageAlignment property.
// The alignment of the image within the scrubber item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView/imageAlignment
func (s_ ScrubberImageItemView) SetImageAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImageAlignment:"), value)
}

// The image view that the scrubber item uses to display its image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView/imageView
func (s_ ScrubberImageItemView) ImageView() NSImageView {
	rv := objc.Send[NSImageView](s_.ID, objc.Sel("imageView"))
	return rv
}

// The image displayed by the scrubber item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/image
func (s_ ScrubberImageItemView) Image() Image {
	rv := objc.Send[Image](s_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The image displayed by the scrubber item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/image
func (s_ ScrubberImageItemView) SetImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImage:"), value)
}



