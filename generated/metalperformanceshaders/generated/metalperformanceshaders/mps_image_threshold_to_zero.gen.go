// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageThresholdToZero */


/* debug [class_header]: Header for MPSImageThresholdToZero */
// The class instance for the [ImageThresholdToZero] class.
var (
	ImageThresholdToZeroClass     _ImageThresholdToZeroClass
	ImageThresholdToZeroClassOnce sync.Once
)

func getImageThresholdToZeroClass() _ImageThresholdToZeroClass {
	ImageThresholdToZeroClassOnce.Do(func() {
		ImageThresholdToZeroClass = _ImageThresholdToZeroClass{objc.GetClass("MPSImageThresholdToZero")}
	})
	return ImageThresholdToZeroClass
}

type _ImageThresholdToZeroClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageThresholdToZero */
// An interface definition for the [ImageThresholdToZero] class.
type IImageThresholdToZero interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageThresholdToZero */
	// properties:
	ThresholdValue() objectivec.IObject
	SetThresholdValue(value objectivec.IObject)
	Transform() objectivec.IObject
	SetTransform(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageThresholdToZero */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageThresholdToZero */
// Alloc allocates a new instance without initialization.
func (ic _ImageThresholdToZeroClass) Alloc() ImageThresholdToZero {
	rv := objc.Send[ImageThresholdToZero](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageThresholdToZeroClass) New() ImageThresholdToZero {
	rv := objc.Send[ImageThresholdToZero](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageThresholdToZero) Init() ImageThresholdToZero {
	rv := objc.Send[ImageThresholdToZero](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageThresholdToZero) Autorelease() ImageThresholdToZero {
	rv := objc.Send[ImageThresholdToZero](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageThresholdToZero creates a new ImageThresholdToZero instance.
func NewImageThresholdToZero() ImageThresholdToZero {
	return getImageThresholdToZeroClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageThresholdToZero */
// A filter that returns the original value for each pixel with a value greater than a specified threshold or 0 otherwise.
//
// An filter converts a single channel image to a binary image. If the input image is not a single channel image, the function first converts the input image into a single channel luminance image using the linear gray color transform, and then it applies the threshold. The following listing shows the threshold to zero function. Listing 1. Threshold to zero function


// A filter that returns the original value for each pixel with a value greater than a specified threshold or 0 otherwise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZero
type ImageThresholdToZero struct {
	UnaryImageKernel
}

// ImageThresholdToZeroFrom constructs a [ImageThresholdToZero] from an unsafe.Pointer.
//
// A filter that returns the original value for each pixel with a value greater than a specified threshold or 0 otherwise.
func ImageThresholdToZeroFrom(ptr unsafe.Pointer) ImageThresholdToZero {
	return ImageThresholdToZero{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageThresholdToZero */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero/2865665-initwithcoder
func NewImageThresholdToZeroWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageThresholdToZero {
	instance := getImageThresholdToZeroClass().Alloc()
	rv := objc.Send[ImageThresholdToZero](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageThresholdToZeroWithCoderDevice */


// Initializes the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero/1618865-initwithdevice
func NewImageThresholdToZeroWithDeviceThresholdValueLinearGrayColorTransform(device unsafe.Pointer, thresholdValue float32, transform objectivec.IObject) ImageThresholdToZero {
	instance := getImageThresholdToZeroClass().Alloc()
	rv := objc.Send[ImageThresholdToZero](instance.ID, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), device, thresholdValue, transform)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageThresholdToZeroWithDeviceThresholdValueLinearGrayColorTransform */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageThresholdToZero */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageThresholdToZero */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageThresholdToZero */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageThresholdToZero */

// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero/1618767-thresholdvalue
func (i_ ImageThresholdToZero) ThresholdValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("thresholdValue"))
	return rv
}/* debug [instance_properties/getter]: thresholdValue */


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero/1618767-thresholdvalue
func (i_ ImageThresholdToZero) SetThresholdValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThresholdValue:"), value)
}/* debug [instance_properties/setter]: thresholdValue */


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero/1618823-transform
func (i_ ImageThresholdToZero) Transform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("transform"))
	return rv
}/* debug [instance_properties/getter]: transform */


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero/1618823-transform
func (i_ ImageThresholdToZero) SetTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransform:"), value)
}/* debug [instance_properties/setter]: transform */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageThresholdToZero */


