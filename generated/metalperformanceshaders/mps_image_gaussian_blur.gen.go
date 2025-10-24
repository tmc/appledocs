// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageGaussianBlur] class.
var (
	ImageGaussianBlurClass     _ImageGaussianBlurClass
	ImageGaussianBlurClassOnce sync.Once
)

func getImageGaussianBlurClass() _ImageGaussianBlurClass {
	ImageGaussianBlurClassOnce.Do(func() {
		ImageGaussianBlurClass = _ImageGaussianBlurClass{objc.GetClass("MPSImageGaussianBlur")}
	})
	return ImageGaussianBlurClass
}

type _ImageGaussianBlurClass struct {
	class objc.Class
}

// An interface definition for the [ImageGaussianBlur] class.
type IImageGaussianBlur interface {
	IUnaryImageKernel
	// properties:
	Sigma() float32
	SetSigma(value float32)
	// methods:
}

// A filter that convolves an image with a Gaussian blur of a given sigma in both the x and y directions.


// A filter that convolves an image with a Gaussian blur of a given sigma in both the x and y directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianBlur
type ImageGaussianBlur struct {
	UnaryImageKernel
}

// ImageGaussianBlurFrom constructs a [ImageGaussianBlur] from an unsafe.Pointer.
//
// A filter that convolves an image with a Gaussian blur of a given sigma in both the x and y directions.
func ImageGaussianBlurFrom(ptr unsafe.Pointer) ImageGaussianBlur {
	return ImageGaussianBlur{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageGaussianBlurClass) Alloc() ImageGaussianBlur {
	rv := objc.Send[ImageGaussianBlur](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageGaussianBlurClass) New() ImageGaussianBlur {
	rv := objc.Send[ImageGaussianBlur](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageGaussianBlur) Init() ImageGaussianBlur {
	rv := objc.Send[ImageGaussianBlur](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageGaussianBlur) Autorelease() ImageGaussianBlur {
	rv := objc.Send[ImageGaussianBlur](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageGaussianBlur creates a new ImageGaussianBlur instance.
func NewImageGaussianBlur() ImageGaussianBlur {
	return getImageGaussianBlurClass().New()
}



// Initializes a Gaussian blur filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianBlur/init(device:sigma:)
func NewImageGaussianBlurWithDeviceSigma(device objectivec.IObject, sigma float32) ImageGaussianBlur {
	instance := getImageGaussianBlurClass().Alloc()
	rv := objc.Send[ImageGaussianBlur](instance.ID, objc.Sel("initWithDevice:sigma:"), device, sigma)
	rv.Autorelease()
	return rv
}



// The sigma value with which the filter was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagegaussianblur/sigma
func (i_ ImageGaussianBlur) Sigma() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("sigma"))
	return rv
}


// The sigma value with which the filter was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagegaussianblur/sigma
func (i_ ImageGaussianBlur) SetSigma(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSigma:"), value)
}


