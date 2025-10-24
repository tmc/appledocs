// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [ImageThresholdToZero] class.
type IImageThresholdToZero interface {
	IUnaryImageKernel
	// properties:
	ThresholdValue() float32 /* primitive/slice/pointer. */
	Transform() float32 /* primitive/slice/pointer. */
	SetTransform(value float32 /* primitive/slice/pointer. */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ic _ImageThresholdToZeroClass) Alloc() ImageThresholdToZero {
	rv := objc.Send[ImageThresholdToZero](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdToZero/thresholdValue
func (i_ ImageThresholdToZero) ThresholdValue() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](i_.ID, objc.Sel("thresholdValue"))
	return rv
}


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero/transform
func (i_ ImageThresholdToZero) Transform() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](i_.ID, objc.Sel("transform"))
	return rv
}


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero/transform
func (i_ ImageThresholdToZero) SetTransform(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransform:"), value)
}



