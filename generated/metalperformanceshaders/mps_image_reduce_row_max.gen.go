// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageReduceRowMax] class.
var (
	ImageReduceRowMaxClass     _ImageReduceRowMaxClass
	ImageReduceRowMaxClassOnce sync.Once
)

func getImageReduceRowMaxClass() _ImageReduceRowMaxClass {
	ImageReduceRowMaxClassOnce.Do(func() {
		ImageReduceRowMaxClass = _ImageReduceRowMaxClass{objc.GetClass("MPSImageReduceRowMax")}
	})
	return ImageReduceRowMaxClass
}

type _ImageReduceRowMaxClass struct {
	class objc.Class
}





// An interface definition for the [ImageReduceRowMax] class.
type IImageReduceRowMax interface {
	IImageReduceUnary
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageReduceRowMaxClass) Alloc() ImageReduceRowMax {
	rv := objc.Send[ImageReduceRowMax](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceRowMaxClass) New() ImageReduceRowMax {
	rv := objc.Send[ImageReduceRowMax](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceRowMax) Init() ImageReduceRowMax {
	rv := objc.Send[ImageReduceRowMax](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceRowMax) Autorelease() ImageReduceRowMax {
	rv := objc.Send[ImageReduceRowMax](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceRowMax creates a new ImageReduceRowMax instance.
func NewImageReduceRowMax() ImageReduceRowMax {
	return getImageReduceRowMaxClass().New()
}





// A filter that returns the maximum value for each row in an image.


// A filter that returns the maximum value for each row in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMax
type ImageReduceRowMax struct {
	ImageReduceUnary
}

// ImageReduceRowMaxFrom constructs a [ImageReduceRowMax] from an unsafe.Pointer.
//
// A filter that returns the maximum value for each row in an image.
func ImageReduceRowMaxFrom(ptr unsafe.Pointer) ImageReduceRowMax {
	return ImageReduceRowMax{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowmax/2942328-initwithdevice
func NewImageReduceRowMaxWithDevice(device unsafe.Pointer) ImageReduceRowMax {
	instance := getImageReduceRowMaxClass().Alloc()
	rv := objc.Send[ImageReduceRowMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























