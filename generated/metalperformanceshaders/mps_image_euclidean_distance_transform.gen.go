// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageEuclideanDistanceTransform */


/* debug [class_header]: Header for MPSImageEuclideanDistanceTransform */
// The class instance for the [ImageEuclideanDistanceTransform] class.
var (
	ImageEuclideanDistanceTransformClass     _ImageEuclideanDistanceTransformClass
	ImageEuclideanDistanceTransformClassOnce sync.Once
)

func getImageEuclideanDistanceTransformClass() _ImageEuclideanDistanceTransformClass {
	ImageEuclideanDistanceTransformClassOnce.Do(func() {
		ImageEuclideanDistanceTransformClass = _ImageEuclideanDistanceTransformClass{objc.GetClass("MPSImageEuclideanDistanceTransform")}
	})
	return ImageEuclideanDistanceTransformClass
}

type _ImageEuclideanDistanceTransformClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageEuclideanDistanceTransform */
// An interface definition for the [ImageEuclideanDistanceTransform] class.
type IImageEuclideanDistanceTransform interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageEuclideanDistanceTransform */
	// properties:
	SearchLimitRadius() objectivec.IObject
	SetSearchLimitRadius(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageEuclideanDistanceTransform */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageEuclideanDistanceTransform */
// Alloc allocates a new instance without initialization.
func (ic _ImageEuclideanDistanceTransformClass) Alloc() ImageEuclideanDistanceTransform {
	rv := objc.Send[ImageEuclideanDistanceTransform](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageEuclideanDistanceTransformClass) New() ImageEuclideanDistanceTransform {
	rv := objc.Send[ImageEuclideanDistanceTransform](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageEuclideanDistanceTransform) Init() ImageEuclideanDistanceTransform {
	rv := objc.Send[ImageEuclideanDistanceTransform](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageEuclideanDistanceTransform) Autorelease() ImageEuclideanDistanceTransform {
	rv := objc.Send[ImageEuclideanDistanceTransform](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageEuclideanDistanceTransform creates a new ImageEuclideanDistanceTransform instance.
func NewImageEuclideanDistanceTransform() ImageEuclideanDistanceTransform {
	return getImageEuclideanDistanceTransformClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageEuclideanDistanceTransform */
// A filter that performs a Euclidean distance transform on an image.


// A filter that performs a Euclidean distance transform on an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEuclideanDistanceTransform
type ImageEuclideanDistanceTransform struct {
	UnaryImageKernel
}

// ImageEuclideanDistanceTransformFrom constructs a [ImageEuclideanDistanceTransform] from an unsafe.Pointer.
//
// A filter that performs a Euclidean distance transform on an image.
func ImageEuclideanDistanceTransformFrom(ptr unsafe.Pointer) ImageEuclideanDistanceTransform {
	return ImageEuclideanDistanceTransform{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageEuclideanDistanceTransform */

// Creates a Euclidean distance transform that uses a specified decoder for your data and runs on a specified device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageeuclideandistancetransform/2953972-initwithcoder
func NewImageEuclideanDistanceTransformWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageEuclideanDistanceTransform {
	instance := getImageEuclideanDistanceTransformClass().Alloc()
	rv := objc.Send[ImageEuclideanDistanceTransform](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageEuclideanDistanceTransformWithCoderDevice */


// Creates a Euclidean distance transform that runs on a specified device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageeuclideandistancetransform/2953973-initwithdevice
func NewImageEuclideanDistanceTransformWithDevice(device unsafe.Pointer) ImageEuclideanDistanceTransform {
	instance := getImageEuclideanDistanceTransformClass().Alloc()
	rv := objc.Send[ImageEuclideanDistanceTransform](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageEuclideanDistanceTransformWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageEuclideanDistanceTransform */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageEuclideanDistanceTransform */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageEuclideanDistanceTransform */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageEuclideanDistanceTransform */

// Limits the search in an image from a pixel to the closest nonzero pixel within a specified radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageeuclideandistancetransform/3547977-searchlimitradius
func (i_ ImageEuclideanDistanceTransform) SearchLimitRadius() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("searchLimitRadius"))
	return rv
}/* debug [instance_properties/getter]: searchLimitRadius */


// Limits the search in an image from a pixel to the closest nonzero pixel within a specified radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageeuclideandistancetransform/3547977-searchlimitradius
func (i_ ImageEuclideanDistanceTransform) SetSearchLimitRadius(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSearchLimitRadius:"), value)
}/* debug [instance_properties/setter]: searchLimitRadius */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageEuclideanDistanceTransform */


