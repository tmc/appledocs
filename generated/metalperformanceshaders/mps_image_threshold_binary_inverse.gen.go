// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [ImageThresholdBinaryInverse] class.
type IImageThresholdBinaryInverse interface {
	IUnaryImageKernel
	// properties:
	MaximumValue() float32 /* primitive/slice/pointer. */
	ThresholdValue() float32 /* primitive/slice/pointer. */
	SetThresholdValue(value float32 /* primitive/slice/pointer. */)
	Transform() float32 /* primitive/slice/pointer. */
	SetTransform(value float32 /* primitive/slice/pointer. */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ic _ImageThresholdBinaryInverseClass) Alloc() ImageThresholdBinaryInverse {
	rv := objc.Send[ImageThresholdBinaryInverse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The maximum value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdBinaryInverse/maximumValue
func (i_ ImageThresholdBinaryInverse) MaximumValue() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](i_.ID, objc.Sel("maximumValue"))
	return rv
}


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/thresholdvalue
func (i_ ImageThresholdBinaryInverse) ThresholdValue() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](i_.ID, objc.Sel("thresholdValue"))
	return rv
}


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/thresholdvalue
func (i_ ImageThresholdBinaryInverse) SetThresholdValue(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThresholdValue:"), value)
}


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/transform
func (i_ ImageThresholdBinaryInverse) Transform() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](i_.ID, objc.Sel("transform"))
	return rv
}


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/transform
func (i_ ImageThresholdBinaryInverse) SetTransform(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransform:"), value)
}



