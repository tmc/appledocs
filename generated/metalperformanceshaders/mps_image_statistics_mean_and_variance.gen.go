// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageStatisticsMeanAndVariance] class.
var (
	ImageStatisticsMeanAndVarianceClass     _ImageStatisticsMeanAndVarianceClass
	ImageStatisticsMeanAndVarianceClassOnce sync.Once
)

func getImageStatisticsMeanAndVarianceClass() _ImageStatisticsMeanAndVarianceClass {
	ImageStatisticsMeanAndVarianceClassOnce.Do(func() {
		ImageStatisticsMeanAndVarianceClass = _ImageStatisticsMeanAndVarianceClass{objc.GetClass("MPSImageStatisticsMeanAndVariance")}
	})
	return ImageStatisticsMeanAndVarianceClass
}

type _ImageStatisticsMeanAndVarianceClass struct {
	class objc.Class
}





// An interface definition for the [ImageStatisticsMeanAndVariance] class.
type IImageStatisticsMeanAndVariance interface {
	IUnaryImageKernel
	

	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageStatisticsMeanAndVarianceClass) Alloc() ImageStatisticsMeanAndVariance {
	rv := objc.Send[ImageStatisticsMeanAndVariance](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageStatisticsMeanAndVarianceClass) New() ImageStatisticsMeanAndVariance {
	rv := objc.Send[ImageStatisticsMeanAndVariance](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageStatisticsMeanAndVariance) Init() ImageStatisticsMeanAndVariance {
	rv := objc.Send[ImageStatisticsMeanAndVariance](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageStatisticsMeanAndVariance) Autorelease() ImageStatisticsMeanAndVariance {
	rv := objc.Send[ImageStatisticsMeanAndVariance](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageStatisticsMeanAndVariance creates a new ImageStatisticsMeanAndVariance instance.
func NewImageStatisticsMeanAndVariance() ImageStatisticsMeanAndVariance {
	return getImageStatisticsMeanAndVarianceClass().New()
}





// A kernel that computes the mean and variance for a given region of an image.
//
// The mean and variance values are written to the destination image at the following pixel locations: Mean value is written at pixel location Variance value is written at pixel location


// A kernel that computes the mean and variance for a given region of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMeanAndVariance
type ImageStatisticsMeanAndVariance struct {
	UnaryImageKernel
}

// ImageStatisticsMeanAndVarianceFrom constructs a [ImageStatisticsMeanAndVariance] from an unsafe.Pointer.
//
// A kernel that computes the mean and variance for a given region of an image.
func ImageStatisticsMeanAndVarianceFrom(ptr unsafe.Pointer) ImageStatisticsMeanAndVariance {
	return ImageStatisticsMeanAndVariance{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance/2867044-initwithcoder
func NewImageStatisticsMeanAndVarianceWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageStatisticsMeanAndVariance {
	instance := getImageStatisticsMeanAndVarianceClass().Alloc()
	rv := objc.Send[ImageStatisticsMeanAndVariance](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance/2867165-initwithdevice
func NewImageStatisticsMeanAndVarianceWithDevice(device unsafe.Pointer) ImageStatisticsMeanAndVariance {
	instance := getImageStatisticsMeanAndVarianceClass().Alloc()
	rv := objc.Send[ImageStatisticsMeanAndVariance](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance/2867131-cliprectsource
func (i_ ImageStatisticsMeanAndVariance) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("clipRectSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance/2867131-cliprectsource
func (i_ ImageStatisticsMeanAndVariance) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}







