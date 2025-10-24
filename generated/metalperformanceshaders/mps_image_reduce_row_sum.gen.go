// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageReduceRowSum] class.
var (
	ImageReduceRowSumClass     _ImageReduceRowSumClass
	ImageReduceRowSumClassOnce sync.Once
)

func getImageReduceRowSumClass() _ImageReduceRowSumClass {
	ImageReduceRowSumClassOnce.Do(func() {
		ImageReduceRowSumClass = _ImageReduceRowSumClass{objc.GetClass("MPSImageReduceRowSum")}
	})
	return ImageReduceRowSumClass
}

type _ImageReduceRowSumClass struct {
	class objc.Class
}





// An interface definition for the [ImageReduceRowSum] class.
type IImageReduceRowSum interface {
	IImageReduceUnary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageReduceRowSumClass) Alloc() ImageReduceRowSum {
	rv := objc.Send[ImageReduceRowSum](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceRowSumClass) New() ImageReduceRowSum {
	rv := objc.Send[ImageReduceRowSum](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceRowSum) Init() ImageReduceRowSum {
	rv := objc.Send[ImageReduceRowSum](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceRowSum) Autorelease() ImageReduceRowSum {
	rv := objc.Send[ImageReduceRowSum](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceRowSum creates a new ImageReduceRowSum instance.
func NewImageReduceRowSum() ImageReduceRowSum {
	return getImageReduceRowSumClass().New()
}





// A filter that returns the sum of all values for a row in an image.


// A filter that returns the sum of all values for a row in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowSum
type ImageReduceRowSum struct {
	ImageReduceUnary
}

// ImageReduceRowSumFrom constructs a [ImageReduceRowSum] from an unsafe.Pointer.
//
// A filter that returns the sum of all values for a row in an image.
func ImageReduceRowSumFrom(ptr unsafe.Pointer) ImageReduceRowSum {
	return ImageReduceRowSum{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowsum/2942334-initwithdevice
func NewImageReduceRowSumWithDevice(device unsafe.Pointer) ImageReduceRowSum {
	instance := getImageReduceRowSumClass().Alloc()
	rv := objc.Send[ImageReduceRowSum](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























