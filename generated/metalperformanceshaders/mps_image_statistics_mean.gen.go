// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageStatisticsMean] class.
var (
	ImageStatisticsMeanClass     _ImageStatisticsMeanClass
	ImageStatisticsMeanClassOnce sync.Once
)

func getImageStatisticsMeanClass() _ImageStatisticsMeanClass {
	ImageStatisticsMeanClassOnce.Do(func() {
		ImageStatisticsMeanClass = _ImageStatisticsMeanClass{objc.GetClass("MPSImageStatisticsMean")}
	})
	return ImageStatisticsMeanClass
}

type _ImageStatisticsMeanClass struct {
	class objc.Class
}





// An interface definition for the [ImageStatisticsMean] class.
type IImageStatisticsMean interface {
	IUnaryImageKernel
	

	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageStatisticsMeanClass) Alloc() ImageStatisticsMean {
	rv := objc.Send[ImageStatisticsMean](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageStatisticsMeanClass) New() ImageStatisticsMean {
	rv := objc.Send[ImageStatisticsMean](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageStatisticsMean) Init() ImageStatisticsMean {
	rv := objc.Send[ImageStatisticsMean](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageStatisticsMean) Autorelease() ImageStatisticsMean {
	rv := objc.Send[ImageStatisticsMean](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageStatisticsMean creates a new ImageStatisticsMean instance.
func NewImageStatisticsMean() ImageStatisticsMean {
	return getImageStatisticsMeanClass().New()
}





// A kernel that computes the mean for a given region of an image.


// A kernel that computes the mean for a given region of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMean
type ImageStatisticsMean struct {
	UnaryImageKernel
}

// ImageStatisticsMeanFrom constructs a [ImageStatisticsMean] from an unsafe.Pointer.
//
// A kernel that computes the mean for a given region of an image.
func ImageStatisticsMeanFrom(ptr unsafe.Pointer) ImageStatisticsMean {
	return ImageStatisticsMean{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean/2867124-initwithcoder
func NewImageStatisticsMeanWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageStatisticsMean {
	instance := getImageStatisticsMeanClass().Alloc()
	rv := objc.Send[ImageStatisticsMean](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean/2867156-initwithdevice
func NewImageStatisticsMeanWithDevice(device unsafe.Pointer) ImageStatisticsMean {
	instance := getImageStatisticsMeanClass().Alloc()
	rv := objc.Send[ImageStatisticsMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean/2867093-cliprectsource
func (i_ ImageStatisticsMean) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("clipRectSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean/2867093-cliprectsource
func (i_ ImageStatisticsMean) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}







