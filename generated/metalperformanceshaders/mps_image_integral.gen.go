// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ImageIntegral] class.
var (
	ImageIntegralClass     _ImageIntegralClass
	ImageIntegralClassOnce sync.Once
)

func getImageIntegralClass() _ImageIntegralClass {
	ImageIntegralClassOnce.Do(func() {
		ImageIntegralClass = _ImageIntegralClass{objc.GetClass("MPSImageIntegral")}
	})
	return ImageIntegralClass
}

type _ImageIntegralClass struct {
	class objc.Class
}





// An interface definition for the [ImageIntegral] class.
type IImageIntegral interface {
	IUnaryImageKernel
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageIntegralClass) Alloc() ImageIntegral {
	rv := objc.Send[ImageIntegral](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageIntegralClass) New() ImageIntegral {
	rv := objc.Send[ImageIntegral](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageIntegral) Init() ImageIntegral {
	rv := objc.Send[ImageIntegral](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageIntegral) Autorelease() ImageIntegral {
	rv := objc.Send[ImageIntegral](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageIntegral creates a new ImageIntegral instance.
func NewImageIntegral() ImageIntegral {
	return getImageIntegralClass().New()
}





// A filter that calculates the sum of pixels over a specified region in an image.
//
// The value at each position is the sum of all pixels in a source image rectangle, The following listing shows the pseudocode used to calculate . Listing 1. Pseudocode for sumRect If the channels in the source image are normalized, half-float or floating values, the destination image is recommended to be a 32-bit floating-point image. If the channels in the source image are integer values, it is recommended that an appropriate 32-bit integer image destination format is used.


// A filter that calculates the sum of pixels over a specified region in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageIntegral
type ImageIntegral struct {
	UnaryImageKernel
}

// ImageIntegralFrom constructs a [ImageIntegral] from an unsafe.Pointer.
//
// A filter that calculates the sum of pixels over a specified region in an image.
func ImageIntegralFrom(ptr unsafe.Pointer) ImageIntegral {
	return ImageIntegral{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}































