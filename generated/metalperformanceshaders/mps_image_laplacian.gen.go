// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageLaplacian */


/* debug [class_header]: Header for MPSImageLaplacian */
// The class instance for the [ImageLaplacian] class.
var (
	ImageLaplacianClass     _ImageLaplacianClass
	ImageLaplacianClassOnce sync.Once
)

func getImageLaplacianClass() _ImageLaplacianClass {
	ImageLaplacianClassOnce.Do(func() {
		ImageLaplacianClass = _ImageLaplacianClass{objc.GetClass("MPSImageLaplacian")}
	})
	return ImageLaplacianClass
}

type _ImageLaplacianClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageLaplacian */
// An interface definition for the [ImageLaplacian] class.
type IImageLaplacian interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageLaplacian */
	// properties:
	Bias() objectivec.IObject
	SetBias(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageLaplacian */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageLaplacian */
// Alloc allocates a new instance without initialization.
func (ic _ImageLaplacianClass) Alloc() ImageLaplacian {
	rv := objc.Send[ImageLaplacian](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageLaplacianClass) New() ImageLaplacian {
	rv := objc.Send[ImageLaplacian](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageLaplacian) Init() ImageLaplacian {
	rv := objc.Send[ImageLaplacian](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageLaplacian) Autorelease() ImageLaplacian {
	rv := objc.Send[ImageLaplacian](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageLaplacian creates a new ImageLaplacian instance.
func NewImageLaplacian() ImageLaplacian {
	return getImageLaplacianClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageLaplacian */
// An optimized Laplacian filter, provided for ease of use.
//
// This filter uses an optimized convolution filter with a 3x3 kernel with the following weights:


// An optimized Laplacian filter, provided for ease of use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacian
type ImageLaplacian struct {
	UnaryImageKernel
}

// ImageLaplacianFrom constructs a [ImageLaplacian] from an unsafe.Pointer.
//
// An optimized Laplacian filter, provided for ease of use.
func ImageLaplacianFrom(ptr unsafe.Pointer) ImageLaplacian {
	return ImageLaplacian{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageLaplacian *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageLaplacian */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageLaplacian */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageLaplacian */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageLaplacian */

// The value added to a convolved pixel before it is converted back to its intended storage format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacian/1648929-bias
func (i_ ImageLaplacian) Bias() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("bias"))
	return rv
}/* debug [instance_properties/getter]: bias */


// The value added to a convolved pixel before it is converted back to its intended storage format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacian/1648929-bias
func (i_ ImageLaplacian) SetBias(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBias:"), value)
}/* debug [instance_properties/setter]: bias */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageLaplacian */



