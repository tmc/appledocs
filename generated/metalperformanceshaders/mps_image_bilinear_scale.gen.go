// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageBilinearScale] class.
var (
	ImageBilinearScaleClass     _ImageBilinearScaleClass
	ImageBilinearScaleClassOnce sync.Once
)

func getImageBilinearScaleClass() _ImageBilinearScaleClass {
	ImageBilinearScaleClassOnce.Do(func() {
		ImageBilinearScaleClass = _ImageBilinearScaleClass{objc.GetClass("MPSImageBilinearScale")}
	})
	return ImageBilinearScaleClass
}

type _ImageBilinearScaleClass struct {
	class objc.Class
}

// An interface definition for the [ImageBilinearScale] class.
type IImageBilinearScale interface {
	IImageScale
}

// A filter that resizes and changes the aspect ratio of an image using Bilinear resampling.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBilinearScale
type ImageBilinearScale struct {
	ImageScale
}

// ImageBilinearScaleFrom constructs a [ImageBilinearScale] from an unsafe.Pointer.
//
// A filter that resizes and changes the aspect ratio of an image using Bilinear resampling.
func ImageBilinearScaleFrom(ptr unsafe.Pointer) ImageBilinearScale {
	return ImageBilinearScale{
		ImageScale: ImageScaleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageBilinearScaleClass) Alloc() ImageBilinearScale {
	rv := objc.Send[ImageBilinearScale](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageBilinearScaleClass) New() ImageBilinearScale {
	rv := objc.Send[ImageBilinearScale](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageBilinearScale) Init() ImageBilinearScale {
	rv := objc.Send[ImageBilinearScale](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageBilinearScale) Autorelease() ImageBilinearScale {
	rv := objc.Send[ImageBilinearScale](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageBilinearScale creates a new ImageBilinearScale instance.
func NewImageBilinearScale() ImageBilinearScale {
	return getImageBilinearScaleClass().New()
}




