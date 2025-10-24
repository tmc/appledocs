// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSImageIntegralOfSquares */


/* debug [class_header]: Header for MPSImageIntegralOfSquares */
// The class instance for the [ImageIntegralOfSquares] class.
var (
	ImageIntegralOfSquaresClass     _ImageIntegralOfSquaresClass
	ImageIntegralOfSquaresClassOnce sync.Once
)

func getImageIntegralOfSquaresClass() _ImageIntegralOfSquaresClass {
	ImageIntegralOfSquaresClassOnce.Do(func() {
		ImageIntegralOfSquaresClass = _ImageIntegralOfSquaresClass{objc.GetClass("MPSImageIntegralOfSquares")}
	})
	return ImageIntegralOfSquaresClass
}

type _ImageIntegralOfSquaresClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageIntegralOfSquares */
// An interface definition for the [ImageIntegralOfSquares] class.
type IImageIntegralOfSquares interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageIntegralOfSquares */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageIntegralOfSquares */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageIntegralOfSquares */
// Alloc allocates a new instance without initialization.
func (ic _ImageIntegralOfSquaresClass) Alloc() ImageIntegralOfSquares {
	rv := objc.Send[ImageIntegralOfSquares](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageIntegralOfSquaresClass) New() ImageIntegralOfSquares {
	rv := objc.Send[ImageIntegralOfSquares](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageIntegralOfSquares) Init() ImageIntegralOfSquares {
	rv := objc.Send[ImageIntegralOfSquares](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageIntegralOfSquares) Autorelease() ImageIntegralOfSquares {
	rv := objc.Send[ImageIntegralOfSquares](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageIntegralOfSquares creates a new ImageIntegralOfSquares instance.
func NewImageIntegralOfSquares() ImageIntegralOfSquares {
	return getImageIntegralOfSquaresClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageIntegralOfSquares */
// A filter that calculates the sum of squared pixels over a specified region in an image.
//
// The value at each position is the sum of all squared pixels in a source image rectangle, The following listing shows the pseudocode used to calculate . Listing 1. Pseudocode for sumRect If the channels in the source image are normalized, half-float or floating values, the destination image is recommended to be a 32-bit floating-point image. If the channels in the source image are integer values, it is recommended that an appropriate 32-bit integer image destination format is used.


// A filter that calculates the sum of squared pixels over a specified region in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageIntegralOfSquares
type ImageIntegralOfSquares struct {
	UnaryImageKernel
}

// ImageIntegralOfSquaresFrom constructs a [ImageIntegralOfSquares] from an unsafe.Pointer.
//
// A filter that calculates the sum of squared pixels over a specified region in an image.
func ImageIntegralOfSquaresFrom(ptr unsafe.Pointer) ImageIntegralOfSquares {
	return ImageIntegralOfSquares{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageIntegralOfSquares *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageIntegralOfSquares */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageIntegralOfSquares */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageIntegralOfSquares */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageIntegralOfSquares */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageIntegralOfSquares */



