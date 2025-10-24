// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSScrubberImageItemView */


/* debug [class_header]: Header for NSScrubberImageItemView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrubberImageItemView */
// An interface definition for the [ScrubberImageItemView] class.
type IScrubberImageItemView interface {
	IScrubberItemView
	
/* debug [class_interface_properties]: Properties for ScrubberImageItemView */
	// properties:
	Image() IImage
	SetImage(value IImage)
	ImageAlignment() ImageAlignment
	SetImageAlignment(value ImageAlignment)
	ImageView() IImageView
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrubberImageItemView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrubberImageItemView */
// Alloc allocates a new instance without initialization.
func (sc _ScrubberImageItemViewClass) Alloc() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrubberImageItemView */
// A concrete view subclass for displaying images in a scrubber items.
//
// Provide the image you want to display in the scrubber item to the property. If you want finer control over the appearance of the image, you can access the underlying image view using the property. The image is scaled proportionally to fit the view’s frame. Use the property to determine how the scaled image is cropped within that frame.


// A concrete view subclass for displaying images in a scrubber items.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrubberImageItemView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrubberImageItemView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrubberImageItemView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrubberImageItemView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrubberImageItemView */

// The image displayed by the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView/image
func (s_ ScrubberImageItemView) Image() IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// The image displayed by the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView/image
func (s_ ScrubberImageItemView) SetImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImage:"), value)
}/* debug [instance_properties/setter]: image */


// The alignment of the image within the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView/imageAlignment
func (s_ ScrubberImageItemView) ImageAlignment() ImageAlignment {
	rv := objc.Send[ImageAlignment](s_.ID, objc.Sel("imageAlignment"))
	return rv
}/* debug [instance_properties/getter]: imageAlignment */


// The alignment of the image within the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView/imageAlignment
func (s_ ScrubberImageItemView) SetImageAlignment(value ImageAlignment) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImageAlignment:"), value)
}/* debug [instance_properties/setter]: imageAlignment */


// The image view that the scrubber item uses to display its image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView/imageView
func (s_ ScrubberImageItemView) ImageView() IImageView {
	rv := objc.Send[ImageView](s_.ID, objc.Sel("imageView"))
	return rv
}/* debug [instance_properties/getter]: imageView */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrubberImageItemView */



