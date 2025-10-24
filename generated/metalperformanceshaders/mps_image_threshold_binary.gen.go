// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageThresholdBinary */


/* debug [class_header]: Header for MPSImageThresholdBinary */
// The class instance for the [ImageThresholdBinary] class.
var (
	ImageThresholdBinaryClass     _ImageThresholdBinaryClass
	ImageThresholdBinaryClassOnce sync.Once
)

func getImageThresholdBinaryClass() _ImageThresholdBinaryClass {
	ImageThresholdBinaryClassOnce.Do(func() {
		ImageThresholdBinaryClass = _ImageThresholdBinaryClass{objc.GetClass("MPSImageThresholdBinary")}
	})
	return ImageThresholdBinaryClass
}

type _ImageThresholdBinaryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageThresholdBinary */
// An interface definition for the [ImageThresholdBinary] class.
type IImageThresholdBinary interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageThresholdBinary */
	// properties:
	Transform() objectivec.IObject
	SetTransform(value objectivec.IObject)
	ThresholdValue() objectivec.IObject
	SetThresholdValue(value objectivec.IObject)
	MaximumValue() objectivec.IObject
	SetMaximumValue(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageThresholdBinary */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageThresholdBinary */
// Alloc allocates a new instance without initialization.
func (ic _ImageThresholdBinaryClass) Alloc() ImageThresholdBinary {
	rv := objc.Send[ImageThresholdBinary](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageThresholdBinaryClass) New() ImageThresholdBinary {
	rv := objc.Send[ImageThresholdBinary](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageThresholdBinary) Init() ImageThresholdBinary {
	rv := objc.Send[ImageThresholdBinary](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageThresholdBinary) Autorelease() ImageThresholdBinary {
	rv := objc.Send[ImageThresholdBinary](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageThresholdBinary creates a new ImageThresholdBinary instance.
func NewImageThresholdBinary() ImageThresholdBinary {
	return getImageThresholdBinaryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageThresholdBinary */
// A filter that returns a specified value for each pixel with a value greater than a specified threshold or 0 otherwise.
//
// An filter converts a single channel image to a binary image. If the input image is not a single channel image, the function first converts the input image into a single channel luminance image using the linear gray color transform, and then it applies the threshold. The following listing shows the threshold binary function. Listing 1. Threshold binary function


// A filter that returns a specified value for each pixel with a value greater than a specified threshold or 0 otherwise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinary
type ImageThresholdBinary struct {
	UnaryImageKernel
}

// ImageThresholdBinaryFrom constructs a [ImageThresholdBinary] from an unsafe.Pointer.
//
// A filter that returns a specified value for each pixel with a value greater than a specified threshold or 0 otherwise.
func ImageThresholdBinaryFrom(ptr unsafe.Pointer) ImageThresholdBinary {
	return ImageThresholdBinary{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageThresholdBinary */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/2865668-initwithcoder
func NewImageThresholdBinaryWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageThresholdBinary {
	instance := getImageThresholdBinaryClass().Alloc()
	rv := objc.Send[ImageThresholdBinary](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageThresholdBinaryWithCoderDevice */


// Initializes the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/1618855-initwithdevice
func NewImageThresholdBinaryWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device unsafe.Pointer, thresholdValue float32, maximumValue float32, transform objectivec.IObject) ImageThresholdBinary {
	instance := getImageThresholdBinaryClass().Alloc()
	rv := objc.Send[ImageThresholdBinary](instance.ID, objc.Sel("initWithDevice:thresholdValue:maximumValue:linearGrayColorTransform:"), device, thresholdValue, maximumValue, transform)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageThresholdBinaryWithDeviceThresholdValueMaximumValueLinearGrayColorTransform */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageThresholdBinary */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageThresholdBinary */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageThresholdBinary */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageThresholdBinary */

// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/1618744-transform
func (i_ ImageThresholdBinary) Transform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("transform"))
	return rv
}/* debug [instance_properties/getter]: transform */


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/1618744-transform
func (i_ ImageThresholdBinary) SetTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransform:"), value)
}/* debug [instance_properties/setter]: transform */


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/1618851-thresholdvalue
func (i_ ImageThresholdBinary) ThresholdValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("thresholdValue"))
	return rv
}/* debug [instance_properties/getter]: thresholdValue */


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/1618851-thresholdvalue
func (i_ ImageThresholdBinary) SetThresholdValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThresholdValue:"), value)
}/* debug [instance_properties/setter]: thresholdValue */


// The maximum value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/1618852-maximumvalue
func (i_ ImageThresholdBinary) MaximumValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("maximumValue"))
	return rv
}/* debug [instance_properties/getter]: maximumValue */


// The maximum value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/1618852-maximumvalue
func (i_ ImageThresholdBinary) SetMaximumValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaximumValue:"), value)
}/* debug [instance_properties/setter]: maximumValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageThresholdBinary */


