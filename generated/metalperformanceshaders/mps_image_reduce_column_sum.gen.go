// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageReduceColumnSum] class.
var (
	ImageReduceColumnSumClass     _ImageReduceColumnSumClass
	ImageReduceColumnSumClassOnce sync.Once
)

func getImageReduceColumnSumClass() _ImageReduceColumnSumClass {
	ImageReduceColumnSumClassOnce.Do(func() {
		ImageReduceColumnSumClass = _ImageReduceColumnSumClass{objc.GetClass("MPSImageReduceColumnSum")}
	})
	return ImageReduceColumnSumClass
}

type _ImageReduceColumnSumClass struct {
	class objc.Class
}





// An interface definition for the [ImageReduceColumnSum] class.
type IImageReduceColumnSum interface {
	IImageReduceUnary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageReduceColumnSumClass) Alloc() ImageReduceColumnSum {
	rv := objc.Send[ImageReduceColumnSum](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceColumnSumClass) New() ImageReduceColumnSum {
	rv := objc.Send[ImageReduceColumnSum](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceColumnSum) Init() ImageReduceColumnSum {
	rv := objc.Send[ImageReduceColumnSum](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceColumnSum) Autorelease() ImageReduceColumnSum {
	rv := objc.Send[ImageReduceColumnSum](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceColumnSum creates a new ImageReduceColumnSum instance.
func NewImageReduceColumnSum() ImageReduceColumnSum {
	return getImageReduceColumnSumClass().New()
}





// A filter that returns the sum of all values for a column in an image.


// A filter that returns the sum of all values for a column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnSum
type ImageReduceColumnSum struct {
	ImageReduceUnary
}

// ImageReduceColumnSumFrom constructs a [ImageReduceColumnSum] from an unsafe.Pointer.
//
// A filter that returns the sum of all values for a column in an image.
func ImageReduceColumnSumFrom(ptr unsafe.Pointer) ImageReduceColumnSum {
	return ImageReduceColumnSum{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnsum/2942321-initwithdevice
func NewImageReduceColumnSumWithDevice(device unsafe.Pointer) ImageReduceColumnSum {
	instance := getImageReduceColumnSumClass().Alloc()
	rv := objc.Send[ImageReduceColumnSum](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























