// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [ImageThresholdBinary] class.
type IImageThresholdBinary interface {
	IUnaryImageKernel
	MaximumValue() float32
	SetMaximumValue(value float32)
	ThresholdValue() float32
	SetThresholdValue(value float32)
	Transform() float32
	SetTransform(value float32)
}

// A filter that returns a specified value for each pixel with a value greater than a specified threshold or 0 otherwise.
//
// An filter converts a single channel image to a binary image. If the input image is not a single channel image, the function first converts the input image into a single channel luminance image using the linear gray color transform, and then it applies the threshold. The following listing shows the threshold binary function. Listing 1. Threshold binary function
//
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

// Alloc allocates a new instance without initialization.
func (ic _ImageThresholdBinaryClass) Alloc() ImageThresholdBinary {
	rv := objc.Send[ImageThresholdBinary](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The maximum value used to initialize the threshold filter.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/maximumvalue
func (i_ ImageThresholdBinary) MaximumValue() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("maximumValue"))
	return rv
}


// SetMaximumValue sets the value of the maximumValue property.
// The maximum value used to initialize the threshold filter.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/maximumvalue
func (i_ ImageThresholdBinary) SetMaximumValue(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaximumValue:"), value)
}

// The threshold value used to initialize the threshold filter.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/thresholdvalue
func (i_ ImageThresholdBinary) ThresholdValue() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("thresholdValue"))
	return rv
}


// SetThresholdValue sets the value of the thresholdValue property.
// The threshold value used to initialize the threshold filter.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/thresholdvalue
func (i_ ImageThresholdBinary) SetThresholdValue(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThresholdValue:"), value)
}

// The color transform used to initialize the threshold filter.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/transform
func (i_ ImageThresholdBinary) Transform() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("transform"))
	return rv
}


// SetTransform sets the value of the transform property.
// The color transform used to initialize the threshold filter.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/transform
func (i_ ImageThresholdBinary) SetTransform(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransform:"), value)
}



