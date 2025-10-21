// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
func NewImageNormalizedHistogramWithCoderDevice(aDecoder foundation.ICoder, device objectivec.IObject) ImageNormalizedHistogram {
	instance := getImageNormalizedHistogramClass().Alloc()
	rv := objc.Send[ImageNormalizedHistogram](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/cliprectsource
func (i_ ImageNormalizedHistogram) ClipRectSource() corelocation.Region {
	rv := objc.Send[corelocation.Region](i_.ID, objc.Sel("clipRectSource"))
	return rv
}


// SetClipRectSource sets the value of the clipRectSource property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/cliprectsource
func (i_ ImageNormalizedHistogram) SetClipRectSource(value corelocation.IRegion) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/histograminfo
func (i_ ImageNormalizedHistogram) HistogramInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("histogramInfo"))
	return rv
}


// SetHistogramInfo sets the value of the histogramInfo property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/histograminfo
func (i_ ImageNormalizedHistogram) SetHistogramInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHistogramInfo:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/zerohistogram
func (i_ ImageNormalizedHistogram) ZeroHistogram() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("zeroHistogram"))
	return rv
}


// SetZeroHistogram sets the value of the zeroHistogram property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/zerohistogram
func (i_ ImageNormalizedHistogram) SetZeroHistogram(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setZeroHistogram:"), value)
}


