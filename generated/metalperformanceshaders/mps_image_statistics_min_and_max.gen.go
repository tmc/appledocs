// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageStatisticsMinAndMax] class.
var (
	ImageStatisticsMinAndMaxClass     _ImageStatisticsMinAndMaxClass
	ImageStatisticsMinAndMaxClassOnce sync.Once
)

func getImageStatisticsMinAndMaxClass() _ImageStatisticsMinAndMaxClass {
	ImageStatisticsMinAndMaxClassOnce.Do(func() {
		ImageStatisticsMinAndMaxClass = _ImageStatisticsMinAndMaxClass{objc.GetClass("MPSImageStatisticsMinAndMax")}
	})
	return ImageStatisticsMinAndMaxClass
}

type _ImageStatisticsMinAndMaxClass struct {
	class objc.Class
}

// An interface definition for the [ImageStatisticsMinAndMax] class.
type IImageStatisticsMinAndMax interface {
	IUnaryImageKernel
	// properties:
	ClipRectSource() objc.IObject /* cross-framework: MTLRegion */
	SetClipRectSource(value objc.IObject /* cross-framework: MTLRegion */)
	// methods:
}

// A kernel that computes the minimum and maximum pixel values for a given region of an image.
//
// The minimum and maximum values are written to the destination image at the following pixel locations: Minimum value is written at pixel location Maximum value is written at pixel location


// A kernel that computes the minimum and maximum pixel values for a given region of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMinAndMax
type ImageStatisticsMinAndMax struct {
	UnaryImageKernel
}

// ImageStatisticsMinAndMaxFrom constructs a [ImageStatisticsMinAndMax] from an unsafe.Pointer.
//
// A kernel that computes the minimum and maximum pixel values for a given region of an image.
func ImageStatisticsMinAndMaxFrom(ptr unsafe.Pointer) ImageStatisticsMinAndMax {
	return ImageStatisticsMinAndMax{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageStatisticsMinAndMaxClass) Alloc() ImageStatisticsMinAndMax {
	rv := objc.Send[ImageStatisticsMinAndMax](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageStatisticsMinAndMaxClass) New() ImageStatisticsMinAndMax {
	rv := objc.Send[ImageStatisticsMinAndMax](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageStatisticsMinAndMax) Init() ImageStatisticsMinAndMax {
	rv := objc.Send[ImageStatisticsMinAndMax](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageStatisticsMinAndMax) Autorelease() ImageStatisticsMinAndMax {
	rv := objc.Send[ImageStatisticsMinAndMax](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageStatisticsMinAndMax creates a new ImageStatisticsMinAndMax instance.
func NewImageStatisticsMinAndMax() ImageStatisticsMinAndMax {
	return getImageStatisticsMinAndMaxClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMinAndMax/init(coder:device:)
func NewImageStatisticsMinAndMaxWithCoderDevice(aDecoder objc.IObject /* cross-framework: Coder */, device objectivec.IObject) ImageStatisticsMinAndMax {
	instance := getImageStatisticsMinAndMaxClass().Alloc()
	rv := objc.Send[ImageStatisticsMinAndMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMinAndMax/init(device:)
func NewImageStatisticsMinAndMaxWithDevice(device objectivec.IObject) ImageStatisticsMinAndMax {
	instance := getImageStatisticsMinAndMaxClass().Alloc()
	rv := objc.Send[ImageStatisticsMinAndMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMinAndMax/clipRectSource
func (i_ ImageStatisticsMinAndMax) ClipRectSource() objc.IObject /* cross-framework: MTLRegion */ {
	rv := objc.Send[Region](i_.ID, objc.Sel("clipRectSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMinAndMax/clipRectSource
func (i_ ImageStatisticsMinAndMax) SetClipRectSource(value objc.IObject /* cross-framework: MTLRegion */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}


