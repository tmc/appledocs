// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageThresholdToZeroInverse */


/* debug [class_header]: Header for MPSImageThresholdToZeroInverse */
// The class instance for the [ImageThresholdToZeroInverse] class.
var (
	ImageThresholdToZeroInverseClass     _ImageThresholdToZeroInverseClass
	ImageThresholdToZeroInverseClassOnce sync.Once
)

func getImageThresholdToZeroInverseClass() _ImageThresholdToZeroInverseClass {
	ImageThresholdToZeroInverseClassOnce.Do(func() {
		ImageThresholdToZeroInverseClass = _ImageThresholdToZeroInverseClass{objc.GetClass("MPSImageThresholdToZeroInverse")}
	})
	return ImageThresholdToZeroInverseClass
}

type _ImageThresholdToZeroInverseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageThresholdToZeroInverse */
// An interface definition for the [ImageThresholdToZeroInverse] class.
type IImageThresholdToZeroInverse interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageThresholdToZeroInverse */
	// properties:
	Transform() objectivec.IObject
	SetTransform(value objectivec.IObject)
	ThresholdValue() objectivec.IObject
	SetThresholdValue(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageThresholdToZeroInverse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageThresholdToZeroInverse */
// Alloc allocates a new instance without initialization.
func (ic _ImageThresholdToZeroInverseClass) Alloc() ImageThresholdToZeroInverse {
	rv := objc.Send[ImageThresholdToZeroInverse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageThresholdToZeroInverseClass) New() ImageThresholdToZeroInverse {
	rv := objc.Send[ImageThresholdToZeroInverse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageThresholdToZeroInverse) Init() ImageThresholdToZeroInverse {
	rv := objc.Send[ImageThresholdToZeroInverse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageThresholdToZeroInverse) Autorelease() ImageThresholdToZeroInverse {
	rv := objc.Send[ImageThresholdToZeroInverse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageThresholdToZeroInverse creates a new ImageThresholdToZeroInverse instance.
func NewImageThresholdToZeroInverse() ImageThresholdToZeroInverse {
	return getImageThresholdToZeroInverseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageThresholdToZeroInverse */
// A filter that returns 0 for each pixel with a value greater than a specified threshold or the original value otherwise.
//
// An filter converts a single channel image to a binary image. If the input image is not a single channel image, the function first converts the input image into a single channel luminance image using the linear gray color transform, and then it applies the threshold. The following listing shows the threshold to zero inverse function. Listing 1. Threshold to zero inverse function


// A filter that returns 0 for each pixel with a value greater than a specified threshold or the original value otherwise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZeroInverse
type ImageThresholdToZeroInverse struct {
	UnaryImageKernel
}

// ImageThresholdToZeroInverseFrom constructs a [ImageThresholdToZeroInverse] from an unsafe.Pointer.
//
// A filter that returns 0 for each pixel with a value greater than a specified threshold or the original value otherwise.
func ImageThresholdToZeroInverseFrom(ptr unsafe.Pointer) ImageThresholdToZeroInverse {
	return ImageThresholdToZeroInverse{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageThresholdToZeroInverse */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/2865667-initwithcoder
func NewImageThresholdToZeroInverseWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageThresholdToZeroInverse {
	instance := getImageThresholdToZeroInverseClass().Alloc()
	rv := objc.Send[ImageThresholdToZeroInverse](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageThresholdToZeroInverseWithCoderDevice */


// Initializes the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618911-initwithdevice
func NewImageThresholdToZeroInverseWithDeviceThresholdValueLinearGrayColorTransform(device unsafe.Pointer, thresholdValue float32, transform objectivec.IObject) ImageThresholdToZeroInverse {
	instance := getImageThresholdToZeroInverseClass().Alloc()
	rv := objc.Send[ImageThresholdToZeroInverse](instance.ID, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), device, thresholdValue, transform)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageThresholdToZeroInverseWithDeviceThresholdValueLinearGrayColorTransform */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageThresholdToZeroInverse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageThresholdToZeroInverse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageThresholdToZeroInverse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageThresholdToZeroInverse */

// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618828-transform
func (i_ ImageThresholdToZeroInverse) Transform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("transform"))
	return rv
}/* debug [instance_properties/getter]: transform */


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618828-transform
func (i_ ImageThresholdToZeroInverse) SetTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransform:"), value)
}/* debug [instance_properties/setter]: transform */


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618914-thresholdvalue
func (i_ ImageThresholdToZeroInverse) ThresholdValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("thresholdValue"))
	return rv
}/* debug [instance_properties/getter]: thresholdValue */


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618914-thresholdvalue
func (i_ ImageThresholdToZeroInverse) SetThresholdValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThresholdValue:"), value)
}/* debug [instance_properties/setter]: thresholdValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageThresholdToZeroInverse */


