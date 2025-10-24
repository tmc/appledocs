// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSImageTranspose */


/* debug [class_header]: Header for MPSImageTranspose */
// The class instance for the [ImageTranspose] class.
var (
	ImageTransposeClass     _ImageTransposeClass
	ImageTransposeClassOnce sync.Once
)

func getImageTransposeClass() _ImageTransposeClass {
	ImageTransposeClassOnce.Do(func() {
		ImageTransposeClass = _ImageTransposeClass{objc.GetClass("MPSImageTranspose")}
	})
	return ImageTransposeClass
}

type _ImageTransposeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageTranspose */
// An interface definition for the [ImageTranspose] class.
type IImageTranspose interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageTranspose */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageTranspose */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageTranspose */
// Alloc allocates a new instance without initialization.
func (ic _ImageTransposeClass) Alloc() ImageTranspose {
	rv := objc.Send[ImageTranspose](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageTransposeClass) New() ImageTranspose {
	rv := objc.Send[ImageTranspose](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageTranspose) Init() ImageTranspose {
	rv := objc.Send[ImageTranspose](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageTranspose) Autorelease() ImageTranspose {
	rv := objc.Send[ImageTranspose](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageTranspose creates a new ImageTranspose instance.
func NewImageTranspose() ImageTranspose {
	return getImageTransposeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageTranspose */
// A filter that transposes an image.
//
// An filter applies a matrix transposition to the source image by exchanging its rows with its columns.


// A filter that transposes an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageTranspose
type ImageTranspose struct {
	UnaryImageKernel
}

// ImageTransposeFrom constructs a [ImageTranspose] from an unsafe.Pointer.
//
// A filter that transposes an image.
func ImageTransposeFrom(ptr unsafe.Pointer) ImageTranspose {
	return ImageTranspose{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageTranspose *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageTranspose */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageTranspose */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageTranspose */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageTranspose */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageTranspose */



