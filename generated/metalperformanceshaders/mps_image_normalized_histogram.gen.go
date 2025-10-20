// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageNormalizedHistogram] class.
var (
	ImageNormalizedHistogramClass     _ImageNormalizedHistogramClass
	ImageNormalizedHistogramClassOnce sync.Once
)

func getImageNormalizedHistogramClass() _ImageNormalizedHistogramClass {
	ImageNormalizedHistogramClassOnce.Do(func() {
		ImageNormalizedHistogramClass = _ImageNormalizedHistogramClass{objc.GetClass("MPSImageNormalizedHistogram")}
	})
	return ImageNormalizedHistogramClass
}

type _ImageNormalizedHistogramClass struct {
	class objc.Class
}

// An interface definition for the [ImageNormalizedHistogram] class.
type IImageNormalizedHistogram interface {
	IKernel
}

// A filter that computes the normalized histogram of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram
type ImageNormalizedHistogram struct {
	Kernel
}

// ImageNormalizedHistogramFrom constructs a [ImageNormalizedHistogram] from an unsafe.Pointer.
//
// A filter that computes the normalized histogram of an image.
func ImageNormalizedHistogramFrom(ptr unsafe.Pointer) ImageNormalizedHistogram {
	return ImageNormalizedHistogram{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageNormalizedHistogramClass) Alloc() ImageNormalizedHistogram {
	rv := objc.Send[ImageNormalizedHistogram](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageNormalizedHistogramClass) New() ImageNormalizedHistogram {
	rv := objc.Send[ImageNormalizedHistogram](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageNormalizedHistogram) Init() ImageNormalizedHistogram {
	rv := objc.Send[ImageNormalizedHistogram](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageNormalizedHistogram) Autorelease() ImageNormalizedHistogram {
	rv := objc.Send[ImageNormalizedHistogram](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageNormalizedHistogram creates a new ImageNormalizedHistogram instance.
func NewImageNormalizedHistogram() ImageNormalizedHistogram {
	return getImageNormalizedHistogramClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram/init(coder:device:)
func NewImageNormalizedHistogramWithCoderDevice(aDecoder unsafe.Pointer, device objc.ID) ImageNormalizedHistogram {
	instance := getImageNormalizedHistogramClass().Alloc()
	rv := objc.Send[ImageNormalizedHistogram](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}
