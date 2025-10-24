// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageLanczosScale] class.
var (
	ImageLanczosScaleClass     _ImageLanczosScaleClass
	ImageLanczosScaleClassOnce sync.Once
)

func getImageLanczosScaleClass() _ImageLanczosScaleClass {
	ImageLanczosScaleClassOnce.Do(func() {
		ImageLanczosScaleClass = _ImageLanczosScaleClass{objc.GetClass("MPSImageLanczosScale")}
	})
	return ImageLanczosScaleClass
}

type _ImageLanczosScaleClass struct {
	class objc.Class
}





// An interface definition for the [ImageLanczosScale] class.
type IImageLanczosScale interface {
	IImageScale
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageLanczosScaleClass) Alloc() ImageLanczosScale {
	rv := objc.Send[ImageLanczosScale](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageLanczosScaleClass) New() ImageLanczosScale {
	rv := objc.Send[ImageLanczosScale](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageLanczosScale) Init() ImageLanczosScale {
	rv := objc.Send[ImageLanczosScale](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageLanczosScale) Autorelease() ImageLanczosScale {
	rv := objc.Send[ImageLanczosScale](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageLanczosScale creates a new ImageLanczosScale instance.
func NewImageLanczosScale() ImageLanczosScale {
	return getImageLanczosScaleClass().New()
}





// A filter that resizes and changes the aspect ratio of an image using Lanczos resampling.
//
// You can use this filter to enlarge or reduce the size of an image, or to change the aspect ratio of an image. The filter uses a Lanczos resampling algorithm, that typically produces better quality for photographs, but is slower than linear sampling that uses GPU texture units. Lanczos downsampling does not require a low pass filter to be applied before it is used. Because the resampling function has negative lobes, Lanczos can result in ringing artifacts near sharp edges, making it less suitable for vector art.


// A filter that resizes and changes the aspect ratio of an image using Lanczos resampling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLanczosScale
type ImageLanczosScale struct {
	ImageScale
}

// ImageLanczosScaleFrom constructs a [ImageLanczosScale] from an unsafe.Pointer.
//
// A filter that resizes and changes the aspect ratio of an image using Lanczos resampling.
func ImageLanczosScaleFrom(ptr unsafe.Pointer) ImageLanczosScale {
	return ImageLanczosScale{
		ImageScale: ImageScaleFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelanczosscale/2867140-initwithcoder
func NewImageLanczosScaleWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageLanczosScale {
	instance := getImageLanczosScaleClass().Alloc()
	rv := objc.Send[ImageLanczosScale](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelanczosscale/2867001-initwithdevice
func NewImageLanczosScaleWithDevice(device unsafe.Pointer) ImageLanczosScale {
	instance := getImageLanczosScaleClass().Alloc()
	rv := objc.Send[ImageLanczosScale](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























