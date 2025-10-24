// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSImageIntegral */


/* debug [class_header]: Header for MPSImageIntegral */
// The class instance for the [ImageIntegral] class.
var (
	ImageIntegralClass     _ImageIntegralClass
	ImageIntegralClassOnce sync.Once
)

func getImageIntegralClass() _ImageIntegralClass {
	ImageIntegralClassOnce.Do(func() {
		ImageIntegralClass = _ImageIntegralClass{objc.GetClass("MPSImageIntegral")}
	})
	return ImageIntegralClass
}

type _ImageIntegralClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageIntegral */
// An interface definition for the [ImageIntegral] class.
type IImageIntegral interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageIntegral */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageIntegral */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageIntegral */
// Alloc allocates a new instance without initialization.
func (ic _ImageIntegralClass) Alloc() ImageIntegral {
	rv := objc.Send[ImageIntegral](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageIntegralClass) New() ImageIntegral {
	rv := objc.Send[ImageIntegral](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageIntegral) Init() ImageIntegral {
	rv := objc.Send[ImageIntegral](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageIntegral) Autorelease() ImageIntegral {
	rv := objc.Send[ImageIntegral](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageIntegral creates a new ImageIntegral instance.
func NewImageIntegral() ImageIntegral {
	return getImageIntegralClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageIntegral */
// A filter that calculates the sum of pixels over a specified region in an image.
//
// The value at each position is the sum of all pixels in a source image rectangle, The following listing shows the pseudocode used to calculate . Listing 1. Pseudocode for sumRect If the channels in the source image are normalized, half-float or floating values, the destination image is recommended to be a 32-bit floating-point image. If the channels in the source image are integer values, it is recommended that an appropriate 32-bit integer image destination format is used.


// A filter that calculates the sum of pixels over a specified region in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageIntegral
type ImageIntegral struct {
	UnaryImageKernel
}

// ImageIntegralFrom constructs a [ImageIntegral] from an unsafe.Pointer.
//
// A filter that calculates the sum of pixels over a specified region in an image.
func ImageIntegralFrom(ptr unsafe.Pointer) ImageIntegral {
	return ImageIntegral{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageIntegral *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageIntegral */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageIntegral */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageIntegral */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageIntegral */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageIntegral */



