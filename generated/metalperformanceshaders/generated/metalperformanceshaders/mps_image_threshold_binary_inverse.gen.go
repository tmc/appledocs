// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageThresholdBinaryInverse */


/* debug [class_header]: Header for MPSImageThresholdBinaryInverse */
// The class instance for the [ImageThresholdBinaryInverse] class.
var (
	ImageThresholdBinaryInverseClass     _ImageThresholdBinaryInverseClass
	ImageThresholdBinaryInverseClassOnce sync.Once
)

func getImageThresholdBinaryInverseClass() _ImageThresholdBinaryInverseClass {
	ImageThresholdBinaryInverseClassOnce.Do(func() {
		ImageThresholdBinaryInverseClass = _ImageThresholdBinaryInverseClass{objc.GetClass("MPSImageThresholdBinaryInverse")}
	})
	return ImageThresholdBinaryInverseClass
}

type _ImageThresholdBinaryInverseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageThresholdBinaryInverse */
// An interface definition for the [ImageThresholdBinaryInverse] class.
type IImageThresholdBinaryInverse interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageThresholdBinaryInverse */
	// properties:
	ThresholdValue() objectivec.IObject
	SetThresholdValue(value objectivec.IObject)
	Transform() objectivec.IObject
	SetTransform(value objectivec.IObject)
	MaximumValue() objectivec.IObject
	SetMaximumValue(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageThresholdBinaryInverse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageThresholdBinaryInverse */
// Alloc allocates a new instance without initialization.
func (ic _ImageThresholdBinaryInverseClass) Alloc() ImageThresholdBinaryInverse {
	rv := objc.Send[ImageThresholdBinaryInverse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageThresholdBinaryInverseClass) New() ImageThresholdBinaryInverse {
	rv := objc.Send[ImageThresholdBinaryInverse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageThresholdBinaryInverse) Init() ImageThresholdBinaryInverse {
	rv := objc.Send[ImageThresholdBinaryInverse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageThresholdBinaryInverse) Autorelease() ImageThresholdBinaryInverse {
	rv := objc.Send[ImageThresholdBinaryInverse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageThresholdBinaryInverse creates a new ImageThresholdBinaryInverse instance.
func NewImageThresholdBinaryInverse() ImageThresholdBinaryInverse {
	return getImageThresholdBinaryInverseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageThresholdBinaryInverse */
// A filter that returns 0 for each pixel with a value greater than a specified threshold or a specified value otherwise.
//
// An function converts a single channel image to a binary image. If the input image is not a single channel image, the function first converts the input image into a single channel luminance image using the linear gray color transform, and then it applies the threshold. The following listing shows the threshold binary inverse function. Listing 1. Threshold binary inverse function


// A filter that returns 0 for each pixel with a value greater than a specified threshold or a specified value otherwise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinaryInverse
type ImageThresholdBinaryInverse struct {
	UnaryImageKernel
}

// ImageThresholdBinaryInverseFrom constructs a [ImageThresholdBinaryInverse] from an unsafe.Pointer.
//
// A filter that returns 0 for each pixel with a value greater than a specified threshold or a specified value otherwise.
func ImageThresholdBinaryInverseFrom(ptr unsafe.Pointer) ImageThresholdBinaryInverse {
	return ImageThresholdBinaryInverse{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageThresholdBinaryInverse */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/2865666-initwithcoder
func NewImageThresholdBinaryInverseWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageThresholdBinaryInverse {
	instance := getImageThresholdBinaryInverseClass().Alloc()
	rv := objc.Send[ImageThresholdBinaryInverse](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageThresholdBinaryInverseWithCoderDevice */


// Initializes the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/1618903-initwithdevice
func NewImageThresholdBinaryInverseWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device unsafe.Pointer, thresholdValue float32, maximumValue float32, transform objectivec.IObject) ImageThresholdBinaryInverse {
	instance := getImageThresholdBinaryInverseClass().Alloc()
	rv := objc.Send[ImageThresholdBinaryInverse](instance.ID, objc.Sel("initWithDevice:thresholdValue:maximumValue:linearGrayColorTransform:"), device, thresholdValue, maximumValue, transform)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageThresholdBinaryInverseWithDeviceThresholdValueMaximumValueLinearGrayColorTransform */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageThresholdBinaryInverse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageThresholdBinaryInverse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageThresholdBinaryInverse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageThresholdBinaryInverse */

// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/1618845-thresholdvalue
func (i_ ImageThresholdBinaryInverse) ThresholdValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("thresholdValue"))
	return rv
}/* debug [instance_properties/getter]: thresholdValue */


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/1618845-thresholdvalue
func (i_ ImageThresholdBinaryInverse) SetThresholdValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThresholdValue:"), value)
}/* debug [instance_properties/setter]: thresholdValue */


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/1618904-transform
func (i_ ImageThresholdBinaryInverse) Transform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("transform"))
	return rv
}/* debug [instance_properties/getter]: transform */


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/1618904-transform
func (i_ ImageThresholdBinaryInverse) SetTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransform:"), value)
}/* debug [instance_properties/setter]: transform */


// The maximum value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/1618906-maximumvalue
func (i_ ImageThresholdBinaryInverse) MaximumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("maximumValue"))
	return rv
}/* debug [instance_properties/getter]: maximumValue */


// The maximum value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/1618906-maximumvalue
func (i_ ImageThresholdBinaryInverse) SetMaximumValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaximumValue:"), value)
}/* debug [instance_properties/setter]: maximumValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageThresholdBinaryInverse */


