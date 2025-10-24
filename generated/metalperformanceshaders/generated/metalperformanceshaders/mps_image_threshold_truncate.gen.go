// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageThresholdTruncate */


/* debug [class_header]: Header for MPSImageThresholdTruncate */
// The class instance for the [ImageThresholdTruncate] class.
var (
	ImageThresholdTruncateClass     _ImageThresholdTruncateClass
	ImageThresholdTruncateClassOnce sync.Once
)

func getImageThresholdTruncateClass() _ImageThresholdTruncateClass {
	ImageThresholdTruncateClassOnce.Do(func() {
		ImageThresholdTruncateClass = _ImageThresholdTruncateClass{objc.GetClass("MPSImageThresholdTruncate")}
	})
	return ImageThresholdTruncateClass
}

type _ImageThresholdTruncateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageThresholdTruncate */
// An interface definition for the [ImageThresholdTruncate] class.
type IImageThresholdTruncate interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageThresholdTruncate */
	// properties:
	Transform() objectivec.IObject
	SetTransform(value objectivec.IObject)
	ThresholdValue() objectivec.IObject
	SetThresholdValue(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageThresholdTruncate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageThresholdTruncate */
// Alloc allocates a new instance without initialization.
func (ic _ImageThresholdTruncateClass) Alloc() ImageThresholdTruncate {
	rv := objc.Send[ImageThresholdTruncate](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageThresholdTruncateClass) New() ImageThresholdTruncate {
	rv := objc.Send[ImageThresholdTruncate](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageThresholdTruncate) Init() ImageThresholdTruncate {
	rv := objc.Send[ImageThresholdTruncate](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageThresholdTruncate) Autorelease() ImageThresholdTruncate {
	rv := objc.Send[ImageThresholdTruncate](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageThresholdTruncate creates a new ImageThresholdTruncate instance.
func NewImageThresholdTruncate() ImageThresholdTruncate {
	return getImageThresholdTruncateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageThresholdTruncate */
// A filter that clamps the return value to an upper specified value.
//
// An filter converts a single channel image to a binary image. If the input image is not a single channel image, the function first converts the input image into a single channel luminance image using the linear gray color transform, and then it applies the threshold. The following listing shows the threshold truncate function. Listing 1. Threshold truncate function


// A filter that clamps the return value to an upper specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdTruncate
type ImageThresholdTruncate struct {
	UnaryImageKernel
}

// ImageThresholdTruncateFrom constructs a [ImageThresholdTruncate] from an unsafe.Pointer.
//
// A filter that clamps the return value to an upper specified value.
func ImageThresholdTruncateFrom(ptr unsafe.Pointer) ImageThresholdTruncate {
	return ImageThresholdTruncate{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageThresholdTruncate */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/2865664-initwithcoder
func NewImageThresholdTruncateWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageThresholdTruncate {
	instance := getImageThresholdTruncateClass().Alloc()
	rv := objc.Send[ImageThresholdTruncate](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageThresholdTruncateWithCoderDevice */


// Initializes the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618818-initwithdevice
func NewImageThresholdTruncateWithDeviceThresholdValueLinearGrayColorTransform(device unsafe.Pointer, thresholdValue float32, transform objectivec.IObject) ImageThresholdTruncate {
	instance := getImageThresholdTruncateClass().Alloc()
	rv := objc.Send[ImageThresholdTruncate](instance.ID, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), device, thresholdValue, transform)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageThresholdTruncateWithDeviceThresholdValueLinearGrayColorTransform */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageThresholdTruncate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageThresholdTruncate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageThresholdTruncate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageThresholdTruncate */

// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618787-transform
func (i_ ImageThresholdTruncate) Transform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("transform"))
	return rv
}/* debug [instance_properties/getter]: transform */


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618787-transform
func (i_ ImageThresholdTruncate) SetTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransform:"), value)
}/* debug [instance_properties/setter]: transform */


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618882-thresholdvalue
func (i_ ImageThresholdTruncate) ThresholdValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("thresholdValue"))
	return rv
}/* debug [instance_properties/getter]: thresholdValue */


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618882-thresholdvalue
func (i_ ImageThresholdTruncate) SetThresholdValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThresholdValue:"), value)
}/* debug [instance_properties/setter]: thresholdValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageThresholdTruncate */


