// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ImageLaplacianPyramidSubtract] class.
var (
	ImageLaplacianPyramidSubtractClass     _ImageLaplacianPyramidSubtractClass
	ImageLaplacianPyramidSubtractClassOnce sync.Once
)

func getImageLaplacianPyramidSubtractClass() _ImageLaplacianPyramidSubtractClass {
	ImageLaplacianPyramidSubtractClassOnce.Do(func() {
		ImageLaplacianPyramidSubtractClass = _ImageLaplacianPyramidSubtractClass{objc.GetClass("MPSImageLaplacianPyramidSubtract")}
	})
	return ImageLaplacianPyramidSubtractClass
}

type _ImageLaplacianPyramidSubtractClass struct {
	class objc.Class
}





// An interface definition for the [ImageLaplacianPyramidSubtract] class.
type IImageLaplacianPyramidSubtract interface {
	IImageLaplacianPyramid
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageLaplacianPyramidSubtractClass) Alloc() ImageLaplacianPyramidSubtract {
	rv := objc.Send[ImageLaplacianPyramidSubtract](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageLaplacianPyramidSubtractClass) New() ImageLaplacianPyramidSubtract {
	rv := objc.Send[ImageLaplacianPyramidSubtract](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageLaplacianPyramidSubtract) Init() ImageLaplacianPyramidSubtract {
	rv := objc.Send[ImageLaplacianPyramidSubtract](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageLaplacianPyramidSubtract) Autorelease() ImageLaplacianPyramidSubtract {
	rv := objc.Send[ImageLaplacianPyramidSubtract](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageLaplacianPyramidSubtract creates a new ImageLaplacianPyramidSubtract instance.
func NewImageLaplacianPyramidSubtract() ImageLaplacianPyramidSubtract {
	return getImageLaplacianPyramidSubtractClass().New()
}





// A filter that convolves an image with a subtractive Laplacian pyramid.


// A filter that convolves an image with a subtractive Laplacian pyramid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacianPyramidSubtract
type ImageLaplacianPyramidSubtract struct {
	ImageLaplacianPyramid
}

// ImageLaplacianPyramidSubtractFrom constructs a [ImageLaplacianPyramidSubtract] from an unsafe.Pointer.
//
// A filter that convolves an image with a subtractive Laplacian pyramid.
func ImageLaplacianPyramidSubtractFrom(ptr unsafe.Pointer) ImageLaplacianPyramidSubtract {
	return ImageLaplacianPyramidSubtract{
		ImageLaplacianPyramid: ImageLaplacianPyramidFrom(ptr),
	}
}































