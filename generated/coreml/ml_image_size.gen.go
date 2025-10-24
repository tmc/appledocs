// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLImageSize */


/* debug [class_header]: Header for MLImageSize */
// The class instance for the [ImageSize] class.
var (
	ImageSizeClass     _ImageSizeClass
	ImageSizeClassOnce sync.Once
)

func getImageSizeClass() _ImageSizeClass {
	ImageSizeClassOnce.Do(func() {
		ImageSizeClass = _ImageSizeClass{objc.GetClass("MLImageSize")}
	})
	return ImageSizeClass
}

type _ImageSizeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageSize */
// An interface definition for the [ImageSize] class.
type IImageSize interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ImageSize */
	// properties:
	PixelsHigh() int
	PixelsWide() int
	EnumeratedImageSizes() IMLImageSize
	SetEnumeratedImageSizes(value IMLImageSize)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageSize */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageSize */
// Alloc allocates a new instance without initialization.
func (ic _ImageSizeClass) Alloc() ImageSize {
	rv := objc.Send[ImageSize](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageSizeClass) New() ImageSize {
	rv := objc.Send[ImageSize](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageSize) Init() ImageSize {
	rv := objc.Send[ImageSize](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageSize) Autorelease() ImageSize {
	rv := objc.Send[ImageSize](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageSize creates a new ImageSize instance.
func NewImageSize() ImageSize {
	return getImageSizeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageSize */
// The width and height of an image feature size.


// The width and height of an image feature size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSize
type ImageSize struct {
	objectivec.Object
}

// ImageSizeFrom constructs a [ImageSize] from an unsafe.Pointer.
//
// The width and height of an image feature size.
func ImageSizeFrom(ptr unsafe.Pointer) ImageSize {
	return ImageSize{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageSize *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageSize */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageSize */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageSize */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageSize */

// The height of an image feature in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSize/pixelsHigh
func (i_ ImageSize) PixelsHigh() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsHigh"))
	return rv
}/* debug [instance_properties/getter]: pixelsHigh */


// The width of an image feature in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSize/pixelsWide
func (i_ ImageSize) PixelsWide() int {
	rv := objc.Send[int](i_.ID, objc.Sel("pixelsWide"))
	return rv
}/* debug [instance_properties/getter]: pixelsWide */


// An array of image sizes a model’s image feature accepts as input or produces as output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesizeconstraint/enumeratedimagesizes
func (i_ ImageSize) EnumeratedImageSizes() IMLImageSize {
	rv := objc.Send[ImageSize](i_.ID, objc.Sel("enumeratedImageSizes"))
	return rv
}/* debug [instance_properties/getter]: enumeratedImageSizes */


// An array of image sizes a model’s image feature accepts as input or produces as output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlimagesizeconstraint/enumeratedimagesizes
func (i_ ImageSize) SetEnumeratedImageSizes(value IMLImageSize) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEnumeratedImageSizes:"), value)
}/* debug [instance_properties/setter]: enumeratedImageSizes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLImageSize */



