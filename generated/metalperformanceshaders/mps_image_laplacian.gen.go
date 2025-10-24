// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageLaplacian] class.
var (
	ImageLaplacianClass     _ImageLaplacianClass
	ImageLaplacianClassOnce sync.Once
)

func getImageLaplacianClass() _ImageLaplacianClass {
	ImageLaplacianClassOnce.Do(func() {
		ImageLaplacianClass = _ImageLaplacianClass{objc.GetClass("MPSImageLaplacian")}
	})
	return ImageLaplacianClass
}

type _ImageLaplacianClass struct {
	class objc.Class
}

// An interface definition for the [ImageLaplacian] class.
type IImageLaplacian interface {
	IUnaryImageKernel
	// properties:
	Bias() float32
	SetBias(value float32)
	// methods:
}

// An optimized Laplacian filter, provided for ease of use.
//
// This filter uses an optimized convolution filter with a 3x3 kernel with the following weights:


// An optimized Laplacian filter, provided for ease of use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacian
type ImageLaplacian struct {
	UnaryImageKernel
}

// ImageLaplacianFrom constructs a [ImageLaplacian] from an unsafe.Pointer.
//
// An optimized Laplacian filter, provided for ease of use.
func ImageLaplacianFrom(ptr unsafe.Pointer) ImageLaplacian {
	return ImageLaplacian{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageLaplacianClass) Alloc() ImageLaplacian {
	rv := objc.Send[ImageLaplacian](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageLaplacianClass) New() ImageLaplacian {
	rv := objc.Send[ImageLaplacian](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageLaplacian) Init() ImageLaplacian {
	rv := objc.Send[ImageLaplacian](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageLaplacian) Autorelease() ImageLaplacian {
	rv := objc.Send[ImageLaplacian](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageLaplacian creates a new ImageLaplacian instance.
func NewImageLaplacian() ImageLaplacian {
	return getImageLaplacianClass().New()
}



// The value added to a convolved pixel before it is converted back to its intended storage format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacian/bias
func (i_ ImageLaplacian) Bias() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("bias"))
	return rv
}


// The value added to a convolved pixel before it is converted back to its intended storage format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacian/bias
func (i_ ImageLaplacian) SetBias(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBias:"), value)
}



