// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageSubtract] class.
var (
	ImageSubtractClass     _ImageSubtractClass
	ImageSubtractClassOnce sync.Once
)

func getImageSubtractClass() _ImageSubtractClass {
	ImageSubtractClassOnce.Do(func() {
		ImageSubtractClass = _ImageSubtractClass{objc.GetClass("MPSImageSubtract")}
	})
	return ImageSubtractClass
}

type _ImageSubtractClass struct {
	class objc.Class
}





// An interface definition for the [ImageSubtract] class.
type IImageSubtract interface {
	IImageArithmetic
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageSubtractClass) Alloc() ImageSubtract {
	rv := objc.Send[ImageSubtract](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageSubtractClass) New() ImageSubtract {
	rv := objc.Send[ImageSubtract](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageSubtract) Init() ImageSubtract {
	rv := objc.Send[ImageSubtract](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageSubtract) Autorelease() ImageSubtract {
	rv := objc.Send[ImageSubtract](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageSubtract creates a new ImageSubtract instance.
func NewImageSubtract() ImageSubtract {
	return getImageSubtractClass().New()
}





// A filter that returns the element-wise difference of its two input images.


// A filter that returns the element-wise difference of its two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSubtract
type ImageSubtract struct {
	ImageArithmetic
}

// ImageSubtractFrom constructs a [ImageSubtract] from an unsafe.Pointer.
//
// A filter that returns the element-wise difference of its two input images.
func ImageSubtractFrom(ptr unsafe.Pointer) ImageSubtract {
	return ImageSubtract{
		ImageArithmetic: ImageArithmeticFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesubtract/2866613-initwithdevice
func NewImageSubtractWithDevice(device unsafe.Pointer) ImageSubtract {
	instance := getImageSubtractClass().Alloc()
	rv := objc.Send[ImageSubtract](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























