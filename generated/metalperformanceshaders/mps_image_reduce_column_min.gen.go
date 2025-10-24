// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageReduceColumnMin] class.
var (
	ImageReduceColumnMinClass     _ImageReduceColumnMinClass
	ImageReduceColumnMinClassOnce sync.Once
)

func getImageReduceColumnMinClass() _ImageReduceColumnMinClass {
	ImageReduceColumnMinClassOnce.Do(func() {
		ImageReduceColumnMinClass = _ImageReduceColumnMinClass{objc.GetClass("MPSImageReduceColumnMin")}
	})
	return ImageReduceColumnMinClass
}

type _ImageReduceColumnMinClass struct {
	class objc.Class
}





// An interface definition for the [ImageReduceColumnMin] class.
type IImageReduceColumnMin interface {
	IImageReduceUnary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageReduceColumnMinClass) Alloc() ImageReduceColumnMin {
	rv := objc.Send[ImageReduceColumnMin](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceColumnMinClass) New() ImageReduceColumnMin {
	rv := objc.Send[ImageReduceColumnMin](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceColumnMin) Init() ImageReduceColumnMin {
	rv := objc.Send[ImageReduceColumnMin](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceColumnMin) Autorelease() ImageReduceColumnMin {
	rv := objc.Send[ImageReduceColumnMin](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceColumnMin creates a new ImageReduceColumnMin instance.
func NewImageReduceColumnMin() ImageReduceColumnMin {
	return getImageReduceColumnMinClass().New()
}





// A filter that returns the minimum value for each column in an image.


// A filter that returns the minimum value for each column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMin
type ImageReduceColumnMin struct {
	ImageReduceUnary
}

// ImageReduceColumnMinFrom constructs a [ImageReduceColumnMin] from an unsafe.Pointer.
//
// A filter that returns the minimum value for each column in an image.
func ImageReduceColumnMinFrom(ptr unsafe.Pointer) ImageReduceColumnMin {
	return ImageReduceColumnMin{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnmin/2942333-initwithdevice
func NewImageReduceColumnMinWithDevice(device unsafe.Pointer) ImageReduceColumnMin {
	instance := getImageReduceColumnMinClass().Alloc()
	rv := objc.Send[ImageReduceColumnMin](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























