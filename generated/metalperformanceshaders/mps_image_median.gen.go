// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ImageMedian] class.
type IImageMedian interface {
	IUnaryImageKernel
	// properties:
	KernelDiameter() uint /* primitive/slice/pointer. */
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ic _ImageMedianClass) Alloc() ImageMedian {
	rv := objc.Send[ImageMedian](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian/init(coder:device:)
func NewImageMedianWithCoderDevice(aDecoder objc.IObject /* cross-framework: Coder */, device objectivec.IObject) ImageMedian {
	instance := getImageMedianClass().Alloc()
	rv := objc.Send[ImageMedian](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes a filter for a particular kernel size and device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian/init(device:kernelDiameter:)
func NewImageMedianWithDeviceKernelDiameter(device objectivec.IObject, kernelDiameter uint /* primitive/slice/pointer. */) ImageMedian {
	instance := getImageMedianClass().Alloc()
	rv := objc.Send[ImageMedian](instance.ID, objc.Sel("initWithDevice:kernelDiameter:"), device, kernelDiameter)
	rv.Autorelease()
	return rv
}



// Queries the maximum diameter, in pixels, of the filter window supported by the median filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian/maxKernelDiameter()
func (ic _ImageMedianClass) MaxKernelDiameter() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](objc.ID(ic.class), objc.Sel("maxKernelDiameter"))
	return rv
}


// Queries the minimum diameter, in pixels, of the filter window supported by the median filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian/minKernelDiameter()
func (ic _ImageMedianClass) MinKernelDiameter() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](objc.ID(ic.class), objc.Sel("minKernelDiameter"))
	return rv
}


// The diameter, in pixels, of the filter window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMedian/kernelDiameter
func (i_ ImageMedian) KernelDiameter() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("kernelDiameter"))
	return rv
}


