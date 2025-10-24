// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageMedian */


/* debug [class_header]: Header for MPSImageMedian */
// The class instance for the [ImageMedian] class.
var (
	ImageMedianClass     _ImageMedianClass
	ImageMedianClassOnce sync.Once
)

func getImageMedianClass() _ImageMedianClass {
	ImageMedianClassOnce.Do(func() {
		ImageMedianClass = _ImageMedianClass{objc.GetClass("MPSImageMedian")}
	})
	return ImageMedianClass
}

type _ImageMedianClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageMedian */
// An interface definition for the [ImageMedian] class.
type IImageMedian interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageMedian */
	// properties:
	KernelDiameter() objectivec.IObject
	SetKernelDiameter(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageMedian */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageMedian */
// Alloc allocates a new instance without initialization.
func (ic _ImageMedianClass) Alloc() ImageMedian {
	rv := objc.Send[ImageMedian](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageMedianClass) New() ImageMedian {
	rv := objc.Send[ImageMedian](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageMedian) Init() ImageMedian {
	rv := objc.Send[ImageMedian](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageMedian) Autorelease() ImageMedian {
	rv := objc.Send[ImageMedian](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageMedian creates a new ImageMedian instance.
func NewImageMedian() ImageMedian {
	return getImageMedianClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageMedian */
// A filter that applies a median filter in a square region centered around each pixel in the source image.
//
// An filter finds the median color value for each channel within a window surrounding the pixel of interest. It is a common means of noise reduction and also as a smoothing filter with edge preserving qualities.


// A filter that applies a median filter in a square region centered around each pixel in the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian
type ImageMedian struct {
	UnaryImageKernel
}

// ImageMedianFrom constructs a [ImageMedian] from an unsafe.Pointer.
//
// A filter that applies a median filter in a square region centered around each pixel in the source image.
func ImageMedianFrom(ptr unsafe.Pointer) ImageMedian {
	return ImageMedian{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageMedian */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/2865529-initwithcoder
func NewImageMedianWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageMedian {
	instance := getImageMedianClass().Alloc()
	rv := objc.Send[ImageMedian](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageMedianWithCoderDevice */


// Initializes a filter for a particular kernel size and device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/1618837-initwithdevice
func NewImageMedianWithDeviceKernelDiameter(device unsafe.Pointer, kernelDiameter uint) ImageMedian {
	instance := getImageMedianClass().Alloc()
	rv := objc.Send[ImageMedian](instance.ID, objc.Sel("initWithDevice:kernelDiameter:"), device, kernelDiameter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageMedianWithDeviceKernelDiameter */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageMedian */

// Queries the maximum diameter, in pixels, of the filter window supported by the median filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/1618830-maxkerneldiameter
func (ic _ImageMedianClass) MaxKernelDiameter() {
	objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("maxKernelDiameter"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MaxKernelDiameter) */


// Queries the minimum diameter, in pixels, of the filter window supported by the median filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/1618864-minkerneldiameter
func (ic _ImageMedianClass) MinKernelDiameter() {
	objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("minKernelDiameter"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MinKernelDiameter) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageMedian */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageMedian */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageMedian */

// The diameter, in pixels, of the filter window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/1618909-kerneldiameter
func (i_ ImageMedian) KernelDiameter() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelDiameter"))
	return rv
}/* debug [instance_properties/getter]: kernelDiameter */


// The diameter, in pixels, of the filter window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/1618909-kerneldiameter
func (i_ ImageMedian) SetKernelDiameter(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelDiameter:"), value)
}/* debug [instance_properties/setter]: kernelDiameter */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageMedian */


