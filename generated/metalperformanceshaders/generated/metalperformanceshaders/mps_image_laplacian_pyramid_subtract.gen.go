// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSImageLaplacianPyramidSubtract */


/* debug [class_header]: Header for MPSImageLaplacianPyramidSubtract */
// The class instance for the [ImageLaplacianPyramidSubtract] class.
var (
	ImageLaplacianPyramidSubtractClass     _ImageLaplacianPyramidSubtractClass
	ImageLaplacianPyramidSubtractClassOnce sync.Once
)

func getImageLaplacianPyramidSubtractClass() _ImageLaplacianPyramidSubtractClass {
	ImageLaplacianPyramidSubtractClassOnce.Do(func() {
		ImageLaplacianPyramidSubtractClass = _ImageLaplacianPyramidSubtractClass{objc.GetClass("MPSImageLaplacianPyramidSubtract")}
	})
	return ImageLaplacianPyramidSubtractClass
}

type _ImageLaplacianPyramidSubtractClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageLaplacianPyramidSubtract */
// An interface definition for the [ImageLaplacianPyramidSubtract] class.
type IImageLaplacianPyramidSubtract interface {
	IImageLaplacianPyramid
	
/* debug [class_interface_properties]: Properties for ImageLaplacianPyramidSubtract */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageLaplacianPyramidSubtract */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageLaplacianPyramidSubtract */
// Alloc allocates a new instance without initialization.
func (ic _ImageLaplacianPyramidSubtractClass) Alloc() ImageLaplacianPyramidSubtract {
	rv := objc.Send[ImageLaplacianPyramidSubtract](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageLaplacianPyramidSubtractClass) New() ImageLaplacianPyramidSubtract {
	rv := objc.Send[ImageLaplacianPyramidSubtract](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageLaplacianPyramidSubtract) Init() ImageLaplacianPyramidSubtract {
	rv := objc.Send[ImageLaplacianPyramidSubtract](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageLaplacianPyramidSubtract) Autorelease() ImageLaplacianPyramidSubtract {
	rv := objc.Send[ImageLaplacianPyramidSubtract](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageLaplacianPyramidSubtract creates a new ImageLaplacianPyramidSubtract instance.
func NewImageLaplacianPyramidSubtract() ImageLaplacianPyramidSubtract {
	return getImageLaplacianPyramidSubtractClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageLaplacianPyramidSubtract */
// A filter that convolves an image with a subtractive Laplacian pyramid.


// A filter that convolves an image with a subtractive Laplacian pyramid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacianPyramidSubtract
type ImageLaplacianPyramidSubtract struct {
	ImageLaplacianPyramid
}

// ImageLaplacianPyramidSubtractFrom constructs a [ImageLaplacianPyramidSubtract] from an unsafe.Pointer.
//
// A filter that convolves an image with a subtractive Laplacian pyramid.
func ImageLaplacianPyramidSubtractFrom(ptr unsafe.Pointer) ImageLaplacianPyramidSubtract {
	return ImageLaplacianPyramidSubtract{
		ImageLaplacianPyramid: ImageLaplacianPyramidFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageLaplacianPyramidSubtract *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageLaplacianPyramidSubtract */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageLaplacianPyramidSubtract */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageLaplacianPyramidSubtract */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageLaplacianPyramidSubtract */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageLaplacianPyramidSubtract */



