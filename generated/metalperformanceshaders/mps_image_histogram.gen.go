// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageHistogram] class.
var (
	ImageHistogramClass     _ImageHistogramClass
	ImageHistogramClassOnce sync.Once
)

func getImageHistogramClass() _ImageHistogramClass {
	ImageHistogramClassOnce.Do(func() {
		ImageHistogramClass = _ImageHistogramClass{objc.GetClass("MPSImageHistogram")}
	})
	return ImageHistogramClass
}

type _ImageHistogramClass struct {
	class objc.Class
}

// An interface definition for the [ImageHistogram] class.
type IImageHistogram interface {
	IKernel
	// properties:
	ClipRectSource() objc.IObject /* cross-framework: MTLRegion */
	SetClipRectSource(value objc.IObject /* cross-framework: MTLRegion */)
	HistogramInfo() ImageHistogramInfo /* not a class type */
	SetHistogramInfo(value ImageHistogramInfo /* not a class type */)
	MinPixelThresholdValue() unsafe.Pointer
	SetMinPixelThresholdValue(value unsafe.Pointer)
	ZeroHistogram() bool
	SetZeroHistogram(value bool)
	// methods:
}

// A filter that computes the histogram of an image.
//
// Typically, you use an filter to calculate an image’s histogram that is passed to a subsequent filter such as or . The following listing shows how you can create a histogram filter to calculate the histogram of the , . The filter is passed an instance of that specifies information to compute the histogram for the channels of an image. After encoding, contains the histogram information and can be used for further operations such as equalization or specification. Listing 1. Creating a histogram filter


// A filter that computes the histogram of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram
type ImageHistogram struct {
	Kernel
}

// ImageHistogramFrom constructs a [ImageHistogram] from an unsafe.Pointer.
//
// A filter that computes the histogram of an image.
func ImageHistogramFrom(ptr unsafe.Pointer) ImageHistogram {
	return ImageHistogram{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageHistogramClass) Alloc() ImageHistogram {
	rv := objc.Send[ImageHistogram](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageHistogramClass) New() ImageHistogram {
	rv := objc.Send[ImageHistogram](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageHistogram) Init() ImageHistogram {
	rv := objc.Send[ImageHistogram](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageHistogram) Autorelease() ImageHistogram {
	rv := objc.Send[ImageHistogram](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageHistogram creates a new ImageHistogram instance.
func NewImageHistogram() ImageHistogram {
	return getImageHistogramClass().New()
}



// The source rectangle to use when reading data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/cliprectsource
func (i_ ImageHistogram) ClipRectSource() objc.IObject /* cross-framework: MTLRegion */ {
	rv := objc.Send[Region](i_.ID, objc.Sel("clipRectSource"))
	return rv
}


// The source rectangle to use when reading data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/cliprectsource
func (i_ ImageHistogram) SetClipRectSource(value objc.IObject /* cross-framework: MTLRegion */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}


// A structure describing the histogram content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/histograminfo
func (i_ ImageHistogram) HistogramInfo() ImageHistogramInfo /* not a class type */ {
	rv := objc.Send[ImageHistogramInfo](i_.ID, objc.Sel("histogramInfo"))
	return rv
}


// A structure describing the histogram content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/histograminfo
func (i_ ImageHistogram) SetHistogramInfo(value ImageHistogramInfo /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHistogramInfo:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/minpixelthresholdvalue
func (i_ ImageHistogram) MinPixelThresholdValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("minPixelThresholdValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/minpixelthresholdvalue
func (i_ ImageHistogram) SetMinPixelThresholdValue(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMinPixelThresholdValue:"), value)
}


// Determines whether to zero-initialize the histogram results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/zerohistogram
func (i_ ImageHistogram) ZeroHistogram() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("zeroHistogram"))
	return rv
}


// Determines whether to zero-initialize the histogram results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/zerohistogram
func (i_ ImageHistogram) SetZeroHistogram(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setZeroHistogram:"), value)
}



