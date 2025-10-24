// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageReduceRowMean] class.
var (
	ImageReduceRowMeanClass     _ImageReduceRowMeanClass
	ImageReduceRowMeanClassOnce sync.Once
)

func getImageReduceRowMeanClass() _ImageReduceRowMeanClass {
	ImageReduceRowMeanClassOnce.Do(func() {
		ImageReduceRowMeanClass = _ImageReduceRowMeanClass{objc.GetClass("MPSImageReduceRowMean")}
	})
	return ImageReduceRowMeanClass
}

type _ImageReduceRowMeanClass struct {
	class objc.Class
}





// An interface definition for the [ImageReduceRowMean] class.
type IImageReduceRowMean interface {
	IImageReduceUnary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageReduceRowMeanClass) Alloc() ImageReduceRowMean {
	rv := objc.Send[ImageReduceRowMean](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceRowMeanClass) New() ImageReduceRowMean {
	rv := objc.Send[ImageReduceRowMean](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceRowMean) Init() ImageReduceRowMean {
	rv := objc.Send[ImageReduceRowMean](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceRowMean) Autorelease() ImageReduceRowMean {
	rv := objc.Send[ImageReduceRowMean](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceRowMean creates a new ImageReduceRowMean instance.
func NewImageReduceRowMean() ImageReduceRowMean {
	return getImageReduceRowMeanClass().New()
}





// A filter that returns the mean value for each row in an image.


// A filter that returns the mean value for each row in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMean
type ImageReduceRowMean struct {
	ImageReduceUnary
}

// ImageReduceRowMeanFrom constructs a [ImageReduceRowMean] from an unsafe.Pointer.
//
// A filter that returns the mean value for each row in an image.
func ImageReduceRowMeanFrom(ptr unsafe.Pointer) ImageReduceRowMean {
	return ImageReduceRowMean{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowmean/2942322-initwithdevice
func NewImageReduceRowMeanWithDevice(device unsafe.Pointer) ImageReduceRowMean {
	instance := getImageReduceRowMeanClass().Alloc()
	rv := objc.Send[ImageReduceRowMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























