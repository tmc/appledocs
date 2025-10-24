// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ImageThresholdToZeroInverse] class.
type IImageThresholdToZeroInverse interface {
	IUnaryImageKernel
	

	// properties:
	Transform() objectivec.IObject
	SetTransform(value objectivec.IObject)
	ThresholdValue() objectivec.IObject
	SetThresholdValue(value objectivec.IObject)


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/2865667-initwithcoder
func NewImageThresholdToZeroInverseWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageThresholdToZeroInverse {
	instance := getImageThresholdToZeroInverseClass().Alloc()
	rv := objc.Send[ImageThresholdToZeroInverse](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618911-initwithdevice
func NewImageThresholdToZeroInverseWithDeviceThresholdValueLinearGrayColorTransform(device unsafe.Pointer, thresholdValue float32, transform objectivec.IObject) ImageThresholdToZeroInverse {
	instance := getImageThresholdToZeroInverseClass().Alloc()
	rv := objc.Send[ImageThresholdToZeroInverse](instance.ID, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), device, thresholdValue, transform)
	rv.Autorelease()
	return rv
}






















// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618828-transform
func (i_ ImageThresholdToZeroInverse) Transform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("transform"))
	return rv
}


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618828-transform
func (i_ ImageThresholdToZeroInverse) SetTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransform:"), value)
}


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618914-thresholdvalue
func (i_ ImageThresholdToZeroInverse) ThresholdValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("thresholdValue"))
	return rv
}


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618914-thresholdvalue
func (i_ ImageThresholdToZeroInverse) SetThresholdValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThresholdValue:"), value)
}







